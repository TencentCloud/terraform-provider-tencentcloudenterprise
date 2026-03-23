/*
Provides a resource to create a cloud firewall (cfw) vpc instance.

# Example Usage

# If mode is 0

```hcl

	resource "tencentcloudenterprise_cfw_vpc_instance" "example" {
	  name = "tf_example"
	  mode = 0

	  vpc_fw_instances {
	    name    = "fw_ins_example"
	    vpc_ids = [
	      "vpc-9tk1icg3",
	      "vpc-e8wcbn67"
	    ]
	    fw_deploy {
	      width         = 200
	      cross_a_zone  = 1
	      deploy_region = "ap-beijing-region-jcctest-ops"
	    }
	  }

	  switch_mode = 1
	  fw_vpc_cidr = "auto"
	}

```

# If mode is 1

```hcl

	resource "tencentcloudenterprise_cfw_vpc_instance" "example" {
	  name = "tf_example"
	  mode = 1

	  vpc_fw_instances {
	    name = "fw_ins_example"
	    fw_deploy {
	      deploy_region = "ap-beijing-region-jcctest-ops"
	      width         = 200
	      cross_a_zone  = 0
	    }
	  }

	  ccn_id      = "ccn-peihfqo7"
	  switch_mode = 1
	  fw_vpc_cidr = "auto"
	}

```

# Import

Cloud firewall vpc group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_vpc_instance.example cfwg-4ee69507
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfw_vpc_instance", CNDescription{
		TerraformTypeCN: "VPC间防火墙实例",
		DescriptionCN:   "提供VPC间防火墙实例资源，用于创建和管理VPC间防火墙实例。",
		AttributesCN: map[string]string{
			"name":             "防火墙实例名称",
			"mode":             "防火墙模式",
			"vpc_fw_instances": "VPC防火墙实例配置",
			"fw_ins_id":        "防火墙实例ID",
			"vpc_ids":          "私有网络模式下接入的VpcId列表",
			"fw_deploy":        "部署地域信息",
			"deploy_region":    "防火墙部署地域",
			"width":            "带宽",
			"cross_a_zone":     "异地灾备",
			"zone":             "主可用区",
			"zone_bak":         "备可用区",
			"cdc_id":           "CDC专用集群ID",
			"switch_mode":      "开关模式",
			"fw_vpc_cidr":      "防火墙VPC网段",
			"ccn_id":           "云联网ID",
			"fw_cidr_info":     "指定防火墙使用网段信息",
			"fw_cidr_type":     "防火墙使用的网段类型",
			"fw_cidr_lst":      "为每个vpc指定防火墙的网段",
			"vpc_id":           "VPC ID",
			"fw_cidr":          "防火墙网段",
			"com_fw_cidr":      "其他防火墙占用网段",
			"fw_group_id":      "防火墙组ID",
		},
	})
}

func resourceTencentCloudCfwVpcInstance() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudCfwVpcInstanceCreate,
		Read:        resourceTencentCloudCfwVpcInstanceRead,
		Update:      resourceTencentCloudCfwVpcInstanceUpdate,
		Delete:      resourceTencentCloudCfwVpcInstanceDelete,
		Description: "Provides a resource to create and manage cloud firewall VPC instance for inter-VPC protection",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "VPC firewall (group) name.",
			},
			"mode": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue(MODE),
				Description:  "Mode 0: private network mode; 1: CCN cloud networking mode.",
			},
			"vpc_fw_instances": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "List of firewall instances under firewall (group).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fw_ins_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Firewall instance ID (passed in editing scenario).",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Firewall instance name.",
						},
						"vpc_ids": {
							Type:        schema.TypeSet,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "List of VpcIds accessed in private network mode; only used in private network mode.",
						},
						"fw_deploy": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "Deploy regional information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cdc_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "When it is a CDC firewall, fill in this ID.",
									},
									"deploy_region": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Firewall Deployment Region.",
									},
									"width": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntAtLeast(1),
										Description:  "Bandwidth, unit: Mbps.",
									},
									"cross_a_zone": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validateAllowedIntValue(CROSS_A_ZONE),
										Default:      CROSS_A_ZONE_0,
										Description:  "Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if it is empty, off-site disaster recovery will not be used by default.",
									},
									"zone": {
										Optional:    true,
										Computed:    true,
										Type:        schema.TypeString,
										ForceNew:    true,
										Description: "main zone, use default available zone if empty.",
									},
									"zone_bak": {
										Optional:    true,
										Computed:    true,
										Type:        schema.TypeString,
										ForceNew:    true,
										Description: "Backup availability zone, if empty, the default availability zone is selected.",
									},
								},
							},
						},
					},
				},
			},
			"switch_mode": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue(SWITCH_MODE),
				Description:  "Switch mode of firewall instance. 1: Single point intercommunication; 2: Multi-point communication; 4: Custom Routing.",
			},
			"fw_vpc_cidr": {
				Optional:    true,
				Type:        schema.TypeString,
				Default:     "auto",
				Description: "auto Automatically select the firewall network segment; 10.10.10.0/24 The firewall network segment entered by the user.",
			},
			"ccn_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Cloud networking id, suitable for cloud networking mode.",
			},
			"fw_cidr_info": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Specify the network segment information used by the firewall.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fw_cidr_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of network segment used by the firewall. The values VpcSelf/Assis/Custom respectively represent own network segment priority/extended network segment priority/custom.",
						},
						"fw_cidr_lst": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specify the network segment of the firewall for each vpc.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Vpc id.",
									},
									"fw_cidr": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Firewall network segment, at least /24 network segment.",
									},
								},
							},
						},
						"com_fw_cidr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Other firewalls occupy the network segment, which is usually the network segment specified when the firewall needs to exclusively occupy the vpc.",
						},
					},
				},
			},
			"fw_group_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Firewall Group ID",
			},
		},
	}
}

func resourceTencentCloudCfwVpcInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_vpc_instance.create")()
	defer inconsistentCheck(d, meta)()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var (
		request   = cfw.NewCreateVpcFwGroupRequest()
		response  = cfw.NewCreateVpcFwGroupResponse()
		fwGroupId string
		mode      int
	)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	mode = d.Get("mode").(int)
	request.Mode = helper.IntInt64(mode)

	if mode == MODE_0 {
		if v, ok := d.GetOk("vpc_fw_instances"); ok {
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				vpcFwInstance := cfw.VpcFwInstance{}
				if v, ok := dMap["name"]; ok {
					vpcFwInstance.Name = helper.String(v.(string))
				}

				if v, ok := dMap["vpc_ids"]; ok {
					vpcIdsSet := v.(*schema.Set).List()
					if len(vpcIdsSet) == 0 {
						return fmt.Errorf("if `mode` is 0, `vpc_ids` is required")
					}

					for i := range vpcIdsSet {
						vpcIds := vpcIdsSet[i].(string)
						vpcFwInstance.VpcIds = append(vpcFwInstance.VpcIds, &vpcIds)
					}
				}

				if fwDeployMap, ok := helper.InterfaceToMap(dMap, "fw_deploy"); ok {
					fwDeploy := cfw.FwDeploy{}
					if v, ok := fwDeployMap["deploy_region"]; ok {
						fwDeploy.DeployRegion = helper.String(v.(string))
					}

					if v, ok := fwDeployMap["width"]; ok {
						fwDeploy.Width = helper.IntInt64(v.(int))
					}

					if v, ok := fwDeployMap["cross_a_zone"]; ok {
						fwDeploy.CrossAZone = helper.IntInt64(v.(int))
					}

					if v, ok := fwDeployMap["zone"]; ok {
						fwDeploy.Zone = helper.String(v.(string))
					}

					if v, ok := fwDeployMap["zone_bak"]; ok {
						fwDeploy.ZoneBak = helper.String(v.(string))
					}

					if v, ok := fwDeployMap["cdc_id"]; ok {
						fwDeploy.CdcId = helper.String(v.(string))
					}

					vpcFwInstance.FwDeploy = &fwDeploy
				}

				if v, ok := dMap["fw_ins_id"]; ok {
					vpcFwInstance.FwInsId = helper.String(v.(string))
				}

				request.VpcFwInstances = append(request.VpcFwInstances, &vpcFwInstance)
			}
		}

		if _, ok := d.GetOk("ccn_id"); ok {
			return fmt.Errorf("if `mode` is 0, `ccn_id` is not supported")
		}

		if v, ok := d.GetOkExists("switch_mode"); ok {
			request.SwitchMode = helper.IntInt64(v.(int))
		}

		if v, ok := d.GetOk("fw_cidr_info"); ok {
			fwCidrInfoList := v.([]interface{})
			if len(fwCidrInfoList) > 0 {
				fwCidrInfoMap := fwCidrInfoList[0].(map[string]interface{})
				fwCidrInfo := cfw.FwCidrInfo{}

				if v, ok := fwCidrInfoMap["fw_cidr_type"]; ok {
					fwCidrInfo.FwCidrType = helper.String(v.(string))
				}

				if v, ok := fwCidrInfoMap["com_fw_cidr"]; ok {
					fwCidrInfo.ComFwCidr = helper.String(v.(string))
				}

				if v, ok := fwCidrInfoMap["fw_cidr_lst"]; ok {
					fwCidrLstList := v.([]interface{})
					for _, item := range fwCidrLstList {
						fwCidrMap := item.(map[string]interface{})
						fwCidrLst := cfw.FwVpcCidr{}

						if v, ok := fwCidrMap["vpc_id"]; ok {
							fwCidrLst.VpcId = helper.String(v.(string))
						}

						if v, ok := fwCidrMap["fw_cidr"]; ok {
							fwCidrLst.FwCidr = helper.String(v.(string))
						}

						fwCidrInfo.FwCidrLst = append(fwCidrInfo.FwCidrLst, &fwCidrLst)
					}
				}

				request.FwCidrInfo = &fwCidrInfo
			}
		} else {
			fwCidrInfo := cfw.FwCidrInfo{}
			fwCidrInfo.FwCidrType = helper.String("VpcSelf")
			fwCidrInfo.ComFwCidr = helper.String("")
			request.FwCidrInfo = &fwCidrInfo
		}

	} else {
		if v, ok := d.GetOk("vpc_fw_instances"); ok {
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				vpcFwInstance := cfw.VpcFwInstance{}
				if v, ok := dMap["name"]; ok {
					vpcFwInstance.Name = helper.String(v.(string))
				}

				if v, ok := dMap["vpc_ids"]; ok {
					vpcIdsSet := v.(*schema.Set).List()
					if len(vpcIdsSet) != 0 {
						return fmt.Errorf("if `mode` is 1, `vpc_ids` is not supported")
					}
				}

				if fwDeployMap, ok := helper.InterfaceToMap(dMap, "fw_deploy"); ok {
					fwDeploy := cfw.FwDeploy{}
					if v, ok := fwDeployMap["deploy_region"]; ok {
						fwDeploy.DeployRegion = helper.String(v.(string))
					}

					if v, ok := fwDeployMap["width"]; ok {
						fwDeploy.Width = helper.IntInt64(v.(int))
					}

					if v, ok := fwDeployMap["cross_a_zone"]; ok {
						fwDeploy.CrossAZone = helper.IntInt64(v.(int))
					}

					if v, ok := fwDeployMap["zone"]; ok {
						fwDeploy.Zone = helper.String(v.(string))
					}

					if v, ok := fwDeployMap["zone_bak"]; ok {
						fwDeploy.ZoneBak = helper.String(v.(string))
					}

					if v, ok := fwDeployMap["cdc_id"]; ok {
						fwDeploy.CdcId = helper.String(v.(string))
					}

					vpcFwInstance.FwDeploy = &fwDeploy
				}

				if v, ok := dMap["fw_ins_id"]; ok {
					vpcFwInstance.FwInsId = helper.String(v.(string))
				}

				request.VpcFwInstances = append(request.VpcFwInstances, &vpcFwInstance)
			}
		}

		if v, ok := d.GetOk("ccn_id"); ok {
			request.CcnId = helper.String(v.(string))
		} else {
			return fmt.Errorf("if `mode` is 1, `ccn_id` is required")
		}

		if v, ok := d.GetOk("switch_mode"); ok {
			switchMode := v.(int)
			if switchMode == SWITCH_MODE_2 {
				return fmt.Errorf("if `mode` is 1, `switch_mode` only support 1, 4")
			}

			request.SwitchMode = helper.IntInt64(v.(int))
		}

		if v, ok := d.GetOk("fw_cidr_info"); ok {
			fwCidrInfoList := v.([]interface{})
			if len(fwCidrInfoList) > 0 {
				fwCidrInfoMap := fwCidrInfoList[0].(map[string]interface{})
				fwCidrInfo := cfw.FwCidrInfo{}

				if v, ok := fwCidrInfoMap["fw_cidr_type"]; ok {
					fwCidrInfo.FwCidrType = helper.String(v.(string))
				}

				if v, ok := fwCidrInfoMap["com_fw_cidr"]; ok {
					fwCidrInfo.ComFwCidr = helper.String(v.(string))
				}

				if v, ok := fwCidrInfoMap["fw_cidr_lst"]; ok {
					fwCidrLstList := v.([]interface{})
					for _, item := range fwCidrLstList {
						fwCidrMap := item.(map[string]interface{})
						fwCidr := cfw.FwVpcCidr{}

						if v, ok := fwCidrMap["vpc_id"]; ok {
							fwCidr.VpcId = helper.String(v.(string))
						}

						if v, ok := fwCidrMap["fw_cidr"]; ok {
							fwCidr.FwCidr = helper.String(v.(string))
						}

						fwCidrInfo.FwCidrLst = append(fwCidrInfo.FwCidrLst, &fwCidr)
					}
				}

				request.FwCidrInfo = &fwCidrInfo
			}
		} else {
			fwCidrInfo := cfw.FwCidrInfo{}
			request.FwCidrInfo = &fwCidrInfo
		}
	}

	if v, ok := d.GetOk("fw_vpc_cidr"); ok {
		request.FwVpcCidr = helper.String(v.(string))
	}

	cfwService := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().CreateVpcFwGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cfw vpcInstance failed, reason:%+v", logId, err)
		return err
	}

	fwGroupId = *response.Response.FwGroupId
	d.SetId(fwGroupId)

	// wait for instance to be ready (Status: 0 = normal, 1 = initializing)
	err = resource.Retry(writeRetryTimeout*3, func() *resource.RetryError {
		vpcFwGroupInfo, e := cfwService.DescribeVpcFwGroupInstanceById(ctx, fwGroupId)
		if e != nil {
			return retryError(e)
		}

		if vpcFwGroupInfo == nil {
			e = fmt.Errorf("cfw vpc instance %s not exists", fwGroupId)
			return resource.NonRetryableError(e)
		}

		if *vpcFwGroupInfo.Status == 1 {
			// Still initializing
			log.Printf("[DEBUG]%s cfw vpc instance %s is still initializing, status: %d\n", logId, fwGroupId, *vpcFwGroupInfo.Status)
			return resource.RetryableError(fmt.Errorf("cfw vpc instance %s is still initializing, status: %d", fwGroupId, *vpcFwGroupInfo.Status))
		}

		if *vpcFwGroupInfo.Status == 0 {
			// Normal, ready to use
			log.Printf("[DEBUG]%s cfw vpc instance %s is ready, status: %d\n", logId, fwGroupId, *vpcFwGroupInfo.Status)
			return nil
		}

		return resource.NonRetryableError(fmt.Errorf("cfw vpc instance %s has unexpected status: %d", fwGroupId, *vpcFwGroupInfo.Status))
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cfw vpcInstance failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCfwVpcInstanceRead(d, meta)
}

func resourceTencentCloudCfwVpcInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_vpc_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	fwGroupId := d.Id()
	var mode int64

	cfwService := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}

	vpcInstance, err := cfwService.DescribeVpcFwGroupInstanceById(ctx, fwGroupId)
	if err != nil {
		return err
	}

	if vpcInstance == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CfwVpcInstance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if vpcInstance.FwGroupName != nil {
		_ = d.Set("name", vpcInstance.FwGroupName)
	}

	if vpcInstance.Mode != nil {
		_ = d.Set("mode", vpcInstance.Mode)
		mode = *vpcInstance.Mode
	}

	if vpcInstance.FwInstanceLst != nil {
		vpcFwInstancesList := []interface{}{}
		for _, vpcFwInstances := range vpcInstance.FwInstanceLst {
			vpcFwInstancesMap := map[string]interface{}{}

			if vpcFwInstances.FwInsId != nil {
				vpcFwInstancesMap["fw_ins_id"] = vpcFwInstances.FwInsId
			}

			if vpcFwInstances.FwInsName != nil {
				vpcFwInstancesMap["name"] = vpcFwInstances.FwInsName
			}

			if mode == MODE_0 {
				if vpcFwInstances.JoinInsIdLst != nil {
					vpcFwInstancesMap["vpc_ids"] = vpcFwInstances.JoinInsIdLst
				}
			}

			if vpcFwInstances.FwCvmLst != nil {
				tmpList := make([]map[string]interface{}, 0, len(vpcFwInstances.FwCvmLst))
				for _, fwCvm := range vpcFwInstances.FwCvmLst {
					fwDeployMap := map[string]interface{}{}
					if fwCvm.Region != nil {
						fwDeployMap["deploy_region"] = fwCvm.Region
					}

					if fwCvm.BandWidth != nil {
						fwDeployMap["width"] = fwCvm.BandWidth
					}

					if fwCvm.ZoneZh != nil {
						fwDeployMap["zone"] = fwCvm.ZoneZh
					}

					if fwCvm.ZoneZhBack != nil {
						fwDeployMap["zone_bak"] = fwCvm.ZoneZhBack
					}

					tmpList = append(tmpList, fwDeployMap)
				}

				vpcFwInstancesMap["fw_deploy"] = tmpList
			}

			vpcFwInstancesList = append(vpcFwInstancesList, vpcFwInstancesMap)

			if vpcFwInstances.CcnId != nil && len(vpcFwInstances.CcnId) != 0 {
				_ = d.Set("ccn_id", vpcFwInstances.CcnId[0])
			}
		}

		_ = d.Set("vpc_fw_instances", vpcFwInstancesList)
	}

	if vpcInstance.SwitchMode != nil {
		_ = d.Set("switch_mode", vpcInstance.SwitchMode)
	}

	if vpcInstance.FwVpcCidr != nil {
		_ = d.Set("fw_vpc_cidr", vpcInstance.FwVpcCidr)
	}

	if vpcInstance.FwGroupId != nil {
		_ = d.Set("fw_group_id", vpcInstance.FwGroupId)
	}

	return nil
}

func resourceTencentCloudCfwVpcInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_vpc_instance.update")()
	defer inconsistentCheck(d, meta)()
	var (
		fwGroupId = d.Id()
		logId     = getLogId(contextNil)
	)

	immutableArgs := []string{"mode", "switch_mode", "fw_vpc_cidr", "ccn_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	// Handle name change
	if d.HasChange("name") {
		request := cfw.NewModifyVpcFwGroupRequest()
		request.FwGroupId = &fwGroupId

		if v, ok := d.GetOk("name"); ok {
			request.Name = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().ModifyVpcFwGroup(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s update cfw vpc group name failed, reason:%+v", logId, err)
			return err
		}
	}

	// Handle vpc_fw_instances change (name, width, vpc_ids, and instance count modification is supported)
	if d.HasChange("vpc_fw_instances") {
		oldVal, newVal := d.GetChange("vpc_fw_instances")
		oldList := oldVal.([]interface{})
		newList := newVal.([]interface{})

		widthChanged := false
		needModifyVpcFwGroup := false // name, vpc_ids, or instance count changed

		// If instance count changed, need to use ModifyVpcFwGroup
		if len(oldList) != len(newList) {
			needModifyVpcFwGroup = true
			log.Printf("[DEBUG]%s vpc_fw_instances count changed from %d to %d, will use ModifyVpcFwGroup\n",
				logId, len(oldList), len(newList))
		} else {
			// Check if existing instances changed
			for i := 0; i < len(newList); i++ {
				oldMap := oldList[i].(map[string]interface{})
				newMap := newList[i].(map[string]interface{})

				// Check if only width changed
				oldFwDeploy := oldMap["fw_deploy"].([]interface{})[0].(map[string]interface{})
				newFwDeploy := newMap["fw_deploy"].([]interface{})[0].(map[string]interface{})

				// Check deploy_region change (not allowed)
				if oldFwDeploy["deploy_region"] != newFwDeploy["deploy_region"] {
					return fmt.Errorf("vpc_fw_instances deploy_region cannot be changed")
				}

				// Check name change
				if oldMap["name"] != newMap["name"] {
					needModifyVpcFwGroup = true
				}

				// Check VPC IDs change
				if fmt.Sprintf("%v", oldMap["vpc_ids"]) != fmt.Sprintf("%v", newMap["vpc_ids"]) {
					needModifyVpcFwGroup = true
				}

				// Check width change
				if oldFwDeploy["width"] != newFwDeploy["width"] {
					widthChanged = true
				}
			}
		}

		// Handle name, VPC IDs, or instance count modification using ModifyVpcFwGroup
		if needModifyVpcFwGroup {
			request := cfw.NewModifyVpcFwGroupRequest()
			request.FwGroupId = &fwGroupId

			mode := d.Get("mode").(int)
			if mode != 0 {
				return fmt.Errorf("vpc_ids modification is only supported for mode 0")
			}

			// Construct VpcFwInstances
			for _, item := range newList {
				dMap := item.(map[string]interface{})
				vpcFwInstance := cfw.VpcFwInstance{}

				if v, ok := dMap["fw_ins_id"]; ok {
					vpcFwInstance.FwInsId = helper.String(v.(string))
				}

				if v, ok := dMap["name"]; ok {
					vpcFwInstance.Name = helper.String(v.(string))
				}

				if v, ok := dMap["vpc_ids"]; ok {
					vpcIdsSet := v.(*schema.Set).List()
					for i := range vpcIdsSet {
						vpcId := vpcIdsSet[i].(string)
						vpcFwInstance.VpcIds = append(vpcFwInstance.VpcIds, &vpcId)
					}
				}

				if fwDeployMap, ok := helper.InterfaceToMap(dMap, "fw_deploy"); ok {
					fwDeploy := cfw.FwDeploy{}
					if v, ok := fwDeployMap["deploy_region"]; ok {
						fwDeploy.DeployRegion = helper.String(v.(string))
					}
					if v, ok := fwDeployMap["width"]; ok {
						fwDeploy.Width = helper.IntInt64(v.(int))
					}
					if v, ok := fwDeployMap["cross_a_zone"]; ok {
						fwDeploy.CrossAZone = helper.IntInt64(v.(int))
					}
					if v, ok := fwDeployMap["zone"]; ok {
						fwDeploy.Zone = helper.String(v.(string))
					}
					if v, ok := fwDeployMap["zone_bak"]; ok {
						fwDeploy.ZoneBak = helper.String(v.(string))
					}
					if v, ok := fwDeployMap["cdc_id"]; ok {
						fwDeploy.CdcId = helper.String(v.(string))
					}
					vpcFwInstance.FwDeploy = &fwDeploy
				}

				request.VpcFwInstances = append(request.VpcFwInstances, &vpcFwInstance)
			}

			// Set FwCidrInfo
			if v, ok := d.GetOk("fw_cidr_info"); ok {
				fwCidrInfoList := v.([]interface{})
				if len(fwCidrInfoList) > 0 {
					fwCidrInfoMap := fwCidrInfoList[0].(map[string]interface{})
					fwCidrInfo := cfw.FwCidrInfo{}

					if v, ok := fwCidrInfoMap["fw_cidr_type"]; ok {
						fwCidrInfo.FwCidrType = helper.String(v.(string))
					}
					if v, ok := fwCidrInfoMap["com_fw_cidr"]; ok {
						fwCidrInfo.ComFwCidr = helper.String(v.(string))
					}
					if v, ok := fwCidrInfoMap["fw_cidr_lst"]; ok {
						fwCidrLstList := v.([]interface{})
						for _, item := range fwCidrLstList {
							fwCidrMap := item.(map[string]interface{})
							fwCidr := cfw.FwVpcCidr{}
							if v, ok := fwCidrMap["vpc_id"]; ok {
								fwCidr.VpcId = helper.String(v.(string))
							}
							if v, ok := fwCidrMap["fw_cidr"]; ok {
								fwCidr.FwCidr = helper.String(v.(string))
							}
							fwCidrInfo.FwCidrLst = append(fwCidrInfo.FwCidrLst, &fwCidr)
						}
					}
					request.FwCidrInfo = &fwCidrInfo
				}
			} else {
				fwCidrInfo := cfw.FwCidrInfo{}
				fwCidrInfo.FwCidrType = helper.String("VpcSelf")
				fwCidrInfo.ComFwCidr = helper.String("")
				request.FwCidrInfo = &fwCidrInfo
			}

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().ModifyVpcFwGroup(request)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				}
				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s update cfw vpc group vpc_ids failed, reason:%+v", logId, err)
				return err
			}

			// Wait and verify VPC IDs and names have been updated
			cfwService := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
			ctx := context.WithValue(context.TODO(), logIdKey, logId)

			err = resource.Retry(readRetryTimeout*3, func() *resource.RetryError {
				vpcFwGroupInfo, e := cfwService.DescribeVpcFwGroupInstanceById(ctx, fwGroupId)
				if e != nil {
					return retryError(e)
				}

				if vpcFwGroupInfo == nil || vpcFwGroupInfo.FwInstanceLst == nil {
					return resource.NonRetryableError(fmt.Errorf("cfw vpc instance %s not found", fwGroupId))
				}

				// Verify all instances have correct names and VPC IDs
				for _, item := range newList {
					dMap := item.(map[string]interface{})
					expectedFwInsId := dMap["fw_ins_id"].(string)
					expectedName := dMap["name"].(string)
					expectedVpcIds := dMap["vpc_ids"].(*schema.Set).List()

					expectedVpcIdSet := make(map[string]bool)
					for _, vid := range expectedVpcIds {
						expectedVpcIdSet[vid.(string)] = true
					}

					// Find the instance in response
					found := false
					for _, fwInstance := range vpcFwGroupInfo.FwInstanceLst {
						if fwInstance.FwInsId != nil && *fwInstance.FwInsId == expectedFwInsId {
							found = true

							// Check name
							if fwInstance.FwInsName != nil {
								actualName := *fwInstance.FwInsName
								if actualName != expectedName {
									log.Printf("[DEBUG]%s instance %s name is %s, expected %s, retrying...\n",
										logId, expectedFwInsId, actualName, expectedName)
									return resource.RetryableError(fmt.Errorf("instance name not updated yet"))
								}
							}

							// Check VPC IDs
							if fwInstance.JoinInsIdLst != nil {
								actualVpcIdSet := make(map[string]bool)
								for _, vid := range fwInstance.JoinInsIdLst {
									if vid != nil {
										actualVpcIdSet[*vid] = true
									}
								}

								// Compare sets
								if len(expectedVpcIdSet) != len(actualVpcIdSet) {
									log.Printf("[DEBUG]%s instance %s VPC count mismatch, expected %d, got %d, retrying...\n",
										logId, expectedFwInsId, len(expectedVpcIdSet), len(actualVpcIdSet))
									return resource.RetryableError(fmt.Errorf("VPC IDs not updated yet"))
								}

								for vid := range expectedVpcIdSet {
									if !actualVpcIdSet[vid] {
										log.Printf("[DEBUG]%s instance %s missing VPC %s, retrying...\n", logId, expectedFwInsId, vid)
										return resource.RetryableError(fmt.Errorf("VPC IDs not updated yet"))
									}
								}

								log.Printf("[DEBUG]%s instance %s name and VPC IDs updated successfully\n", logId, expectedFwInsId)
							}
							break
						}
					}

					if !found {
						return resource.NonRetryableError(fmt.Errorf("instance %s not found in fw group", expectedFwInsId))
					}
				}

				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s verify cfw vpc group update failed, reason:%+v", logId, err)
				return err
			}
		}

		// Handle width modification using ModifyVpcCfwWidth (only if name and vpc_ids didn't change)
		if widthChanged && !needModifyVpcFwGroup {
			for i := 0; i < len(newList); i++ {
				oldMap := oldList[i].(map[string]interface{})
				newMap := newList[i].(map[string]interface{})

				oldFwDeploy := oldMap["fw_deploy"].([]interface{})[0].(map[string]interface{})
				newFwDeploy := newMap["fw_deploy"].([]interface{})[0].(map[string]interface{})
				fwInsId := newMap["fw_ins_id"].(string)

				if oldFwDeploy["width"] != newFwDeploy["width"] {
					// Modify width using ModifyVpcCfwWidth API
					request := cfw.NewModifyVpcCfwWidthRequest()
					request.FwType = helper.String("ew") // VPC firewall type
					request.CfwInstance = helper.String(fwInsId)

					fwDeploy := cfw.FwDeploy{}
					if v, ok := newFwDeploy["deploy_region"]; ok {
						fwDeploy.DeployRegion = helper.String(v.(string))
					}

					expectedWidth := 0
					if v, ok := newFwDeploy["width"]; ok {
						expectedWidth = v.(int)
						fwDeploy.Width = helper.IntInt64(expectedWidth)
					}

					request.FwDeploy = []*cfw.FwDeploy{&fwDeploy}

					err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
						result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().ModifyVpcCfwWidth(request)
						if e != nil {
							return retryError(e)
						} else {
							log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
						}

						return nil
					})

					if err != nil {
						log.Printf("[CRITAL]%s update cfw vpc instance width failed, reason:%+v", logId, err)
						return err
					}

					// Wait and verify the width has been updated
					cfwService := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
					ctx := context.WithValue(context.TODO(), logIdKey, logId)

					err = resource.Retry(readRetryTimeout*3, func() *resource.RetryError {
						vpcFwGroupInfo, e := cfwService.DescribeVpcFwGroupInstanceById(ctx, fwGroupId)
						if e != nil {
							return retryError(e)
						}

						if vpcFwGroupInfo == nil || vpcFwGroupInfo.FwInstanceLst == nil {
							return resource.NonRetryableError(fmt.Errorf("cfw vpc instance %s not found", fwGroupId))
						}

						// Find the instance and check width
						for _, fwInstance := range vpcFwGroupInfo.FwInstanceLst {
							if fwInstance.FwInsId != nil && *fwInstance.FwInsId == fwInsId {
								if fwInstance.FwCvmLst != nil && len(fwInstance.FwCvmLst) > 0 {
									if fwInstance.FwCvmLst[0].BandWidth != nil {
										actualWidth := int(*fwInstance.FwCvmLst[0].BandWidth)
										if actualWidth == expectedWidth {
											log.Printf("[DEBUG]%s cfw vpc instance %s width updated successfully to %d\n", logId, fwInsId, actualWidth)
											return nil
										}
										log.Printf("[DEBUG]%s cfw vpc instance %s width is %d, expected %d, retrying...\n", logId, fwInsId, actualWidth, expectedWidth)
										return resource.RetryableError(fmt.Errorf("width not updated yet, current: %d, expected: %d", actualWidth, expectedWidth))
									}
								}
								return resource.NonRetryableError(fmt.Errorf("cannot get bandwidth info for instance %s", fwInsId))
							}
						}

						return resource.NonRetryableError(fmt.Errorf("instance %s not found in fw group", fwInsId))
					})

					if err != nil {
						log.Printf("[CRITAL]%s verify cfw vpc instance width update failed, reason:%+v", logId, err)
						return err
					}
				}
			}
		}
	}

	return resourceTencentCloudCfwVpcInstanceRead(d, meta)
}

func resourceTencentCloudCfwVpcInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_vpc_instance.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cfwService := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
	fwGroupId := d.Id()

	if err := cfwService.DeleteCfwVpcInstanceById(ctx, fwGroupId); err != nil {
		return err
	}

	return nil
}
