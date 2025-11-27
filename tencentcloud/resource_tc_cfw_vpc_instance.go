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
	      width         = 1024
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
	      width         = 1024
	      cross_a_zone  = 0
	    }
	  }

	  ccn_id      = "ccn-peihfqo7"
	  switch_mode = 1
	  fw_vpc_cidr = "auto"
	}

```

# Import

Cloud firewall vpc instance can be imported using the id, e.g.

```
$ terraform import cloud_cfw_vpc_policy.example cfwg-4ee69507
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfw_vpc_instance", CNDescription{
		TerraformTypeCN: "云防火墙VPC实例",
		DescriptionCN:   "提供云防火墙VPC实例资源，用于创建和管理云防火墙VPC间防护实例。",
		AttributesCN: map[string]string{
			"name":             "防火墙实例名称",
			"mode":             "防火墙模式",
			"vpc_fw_instances": "VPC防火墙实例配置",
			"switch_mode":      "开关模式",
			"fw_vpc_cidr":      "防火墙VPC网段",
			"ccn_id":           "云联网ID",
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
							Optional:    true,
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
									"deploy_region": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Firewall Deployment Region.",
									},
									"width": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validateIntegerMin(1024),
										Description:  "Bandwidth, unit: Mbps.",
									},
									"cross_a_zone": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validateAllowedIntValue(CROSS_A_ZONE),
										Default:      CROSS_A_ZONE_0,
										Description:  "Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if it is empty, off-site disaster recovery will not be used by default.",
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
		},
	}
}

func resourceTencentCloudCfwVpcInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfw_vpc_instance.create")()
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

	if v, ok := d.GetOkExists("mode"); ok {
		request.Mode = helper.IntInt64(v.(int))
		mode = v.(int)
	}

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
						return fmt.Errorf("If `mode` is 0, `vpc_ids` is required.")
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
						//crossAZone := v.(int)
						//if v, ok := fwDeployMap["zone_set"]; ok {
						//	zoneList := v.(*schema.Set).List()
						//	if crossAZone == CROSS_A_ZONE_0 {
						//		if len(zoneList) != 1 {
						//			return fmt.Errorf("if `cross_a_zone` is 0, `zone_set` only can be set one zone")
						//		}
						//
						//		fwDeploy.Zone = helper.String(zoneList[0].(string))
						//
						//	} else {
						//		if len(zoneList) != 2 {
						//			return fmt.Errorf("if `cross_a_zone` is 1, `zone_set` must be set two zones")
						//		}
						//
						//		fwDeploy.Zone = helper.String(zoneList[0].(string))
						//		fwDeploy.ZoneBak = helper.String(zoneList[1].(string))
						//	}
						//}
						fwDeploy.CrossAZone = helper.IntInt64(v.(int))
					}

					vpcFwInstance.FwDeploy = &fwDeploy
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

		fwCidrInfo := cfw.FwCidrInfo{}
		fwCidrInfo.FwCidrType = helper.String("VpcSelf")
		fwCidrInfo.ComFwCidr = helper.String("")
		request.FwCidrInfo = &fwCidrInfo

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
						crossAZone := v.(int)
						if v, ok := fwDeployMap["zone_set"]; ok {
							zoneList := v.(*schema.Set).List()
							if crossAZone == CROSS_A_ZONE_0 {
								if len(zoneList) != 1 {
									return fmt.Errorf("if `cross_a_zone` is 0, `zone_set` only can be set one zone")
								}

								fwDeploy.Zone = helper.String(zoneList[0].(string))

							} else {
								if len(zoneList) != 2 {
									return fmt.Errorf("if `cross_a_zone` is 1, `zone_set` must be set tow zones")
								}

								fwDeploy.Zone = helper.String(zoneList[0].(string))
								fwDeploy.ZoneBak = helper.String(zoneList[1].(string))
							}
						}

						fwDeploy.CrossAZone = helper.IntInt64(v.(int))
					}

					vpcFwInstance.FwDeploy = &fwDeploy
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

		fwCidrInfo := cfw.FwCidrInfo{}
		fwCidrInfo.FwCidrType = helper.String("Assis")
		fwCidrInfo.ComFwCidr = helper.String("")
		request.FwCidrInfo = &fwCidrInfo
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

	// wait
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
			return nil
		}

		return resource.RetryableError(fmt.Errorf("create cfw vpcInstance status is %d", *vpcFwGroupInfo.Status))
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cfw vpcInstance failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCfwVpcInstanceRead(d, meta)
}

func resourceTencentCloudCfwVpcInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfw_vpc_instance.read")()
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
					zone := ""
					zoneBak := ""
					if fwCvm.Region != nil {
						fwDeployMap["deploy_region"] = fwCvm.Region
					}

					if fwCvm.BandWidth != nil {
						fwDeployMap["width"] = fwCvm.BandWidth
					}

					//if fwCvm.ZoneZh != nil {
					//	zone = ZONE_MAP_CN2EN[*fwCvm.ZoneZh]
					//}
					//
					//if fwCvm.ZoneZhBack != nil {
					//	zoneBak = ZONE_MAP_CN2EN[*fwCvm.ZoneZhBack]
					//}
					//
					//if zone == zoneBak {
					//	fwDeployMap["cross_a_zone"] = CROSS_A_ZONE_0
					//
					//} else {
					//	fwDeployMap["cross_a_zone"] = CROSS_A_ZONE_1
					//}

					zoneList := []string{
						zone,
						zoneBak,
					}
					fwDeployMap["zone_set"] = zoneList

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

	return nil
}

func resourceTencentCloudCfwVpcInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfw_vpc_instance.update")()
	defer inconsistentCheck(d, meta)()
	logId := getLogId(contextNil)
	fwGroupId := d.Id()
	request := cfw.NewModifyVpcFwGroupRequest()

	immutableArgs := []string{"mode", "vpc_fw_instances", "switch_mode", "fw_vpc_cidr", "ccn_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	request.FwGroupId = &fwGroupId

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok {
			request.Name = helper.String(v.(string))
		}
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
		log.Printf("[CRITAL]%s update cfw vpcInstance failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCfwVpcInstanceRead(d, meta)
}

func resourceTencentCloudCfwVpcInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfw_vpc_instance.delete")()
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
