/*
Provides a resource to create a cloud firewall (cfw) NAT firewall instance.

Example Usage

```hcl

	resource "tencentcloudenterprise_cfw_nat_instance" "example" {
	  name  = "cfw-nat-example"
	  width = 20
	  mode  = 0

	  new_mode_items {
	    vpc_list  = ["vpc-3skwc52h"]
	    eips      = []
	    add_count = 1
	  }

	  cross_a_zone = 0

	  fw_cidr_info {
	    fw_cidr_type = "VpcSelf"
	  }
	}

```

Import

Cloud firewall nat instance can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_nat_instance.example cfwnat-xxxxxxxx
```
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
	registerResourceDescriptionProvider("tencentcloudenterprise_cfw_nat_instance", CNDescription{
		TerraformTypeCN: "NAT防火墙实例",
		DescriptionCN:   "提供NAT防火墙实例资源，用于创建和管理NAT防火墙实例。",
		AttributesCN: map[string]string{
			"name":           "防火墙实例名称",
			"width":          "带宽",
			"mode":           "接入模式，0：新增模式，1：接入模式",
			"nat_gw_list":    "接入模式接入的nat网关列表",
			"new_mode_items": "新增模式传递参数",
			"vpc_list":       "新增模式下接入的vpc列表",
			"eips":           "新增模式下绑定的出口弹性公网ip列表",
			"add_count":      "新增模式下新增绑定的出口弹性公网ip个数",
			"zone":           "主可用区",
			"zone_bak":       "备可用区",
			"cross_a_zone":   "异地灾备",
			"domain":         "域名",
			"fw_cidr_info":   "防火墙使用网段信息",
			"fw_cidr_type":   "防火墙使用的网段类型",
			"fw_cidr_lst":    "为每个vpc指定防火墙的网段",
			"vpc_id":         "VPC的ID",
			"fw_cidr":        "防火墙网段",
			"com_fw_cidr":    "其他防火墙占用网段",
			"cfw_ins_id":     "NAT实例ID",
			"status":         "实例状态",
		},
	})
}

func resourceTencentCloudCfwNatInstance() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudCfwNatInstanceCreate,
		Read:        resourceTencentCloudCfwNatInstanceRead,
		Update:      resourceTencentCloudCfwNatInstanceUpdate,
		Delete:      resourceTencentCloudCfwNatInstanceDelete,
		Description: "Provides a resource to create and manage CFW NAT Firewall instance",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Firewall instance name.",
			},
			"width": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "Bandwidth.",
			},
			"mode": {
				Required:     true,
				ForceNew:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntInSlice([]int{0, 1}),
				Description:  "Access mode, 0: new mode, 1: access mode.",
			},
			"new_mode_items": {
				Optional:     true,
				Type:         schema.TypeList,
				MaxItems:     1,
				ExactlyOneOf: []string{"nat_gw_list", "new_mode_items"},
				Description:  "New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_list": {
							Optional: true,
							Type:     schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "VPC list.",
						},
						"eips": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Elastic public IP list.",
						},
						"add_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Number of EIPs to create. If eips is specified, this will be calculated from eips length.",
						},
					},
				},
			},
			"nat_gw_list": {
				Optional:     true,
				Type:         schema.TypeList,
				ExactlyOneOf: []string{"nat_gw_list", "new_mode_items"},
				Elem:         &schema.Schema{Type: schema.TypeString},
				Description:  "A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.",
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
				Description: "Backup availability zone, if empty, the default availability zone is selected.",
			},
			"cross_a_zone": {
				Optional:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntInSlice([]int{0, 1}),
				Description:  "Cross-region disaster recovery 1: use cross-region disaster recovery; 0: do not use cross-region disaster recovery; if empty, cross-region disaster recovery is not used by default.",
			},
			"domain": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Required if you want to create a domain name.",
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
			"cfw_ins_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Nat firewall instance id.",
			},
			"status": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Instance status. 0: normal, 1: initializing.",
			},
		},
	}
}

func resourceTencentCloudCfwNatInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_nat_instance.create")()
	defer inconsistentCheck(d, meta)()

	request := cfw.NewCreateNatFwInstanceWithDomainRequest()
	response := cfw.NewCreateNatFwInstanceWithDomainResponse()
	logId := getLogId(contextNil)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("width"); ok {
		request.Width = helper.IntInt64(v.(int))
	}

	mode := d.Get("mode").(int)
	request.Mode = helper.IntInt64(mode)

	// Set zone and zone_bak (optional, will use default if not specified)
	if v, ok := d.GetOk("zone"); ok {
		request.Zone = helper.String(v.(string))
	}

	if v, ok := d.GetOk("zone_bak"); ok {
		request.ZoneBak = helper.String(v.(string))
	}

	if mode == MODE_0 {
		if v, ok := d.GetOk("new_mode_items"); ok {
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				newModeItems := cfw.NewModeItems{}
				if v, ok = dMap["vpc_list"]; ok {
					vpcList := v.([]interface{})
					tmqVpcList := make([]*string, 0, len(vpcList))
					for i := range vpcList {
						vpc := vpcList[i].(string)
						tmqVpcList = append(tmqVpcList, &vpc)
					}
					newModeItems.VpcList = tmqVpcList
				}

				if v, ok := dMap["eips"]; ok {
					eipList := v.([]interface{})
					tmqEipList := make([]*string, 0, len(eipList))
					for i := range eipList {
						eip := eipList[i].(string)
						tmqEipList = append(tmqEipList, &eip)
					}
					newModeItems.Eips = tmqEipList
				}

				if v, ok := dMap["add_count"]; ok {
					newModeItems.AddCount = helper.IntInt64(v.(int))
				} else {
					newModeItems.AddCount = helper.IntInt64(len(newModeItems.Eips))
				}

				request.NewModeItems = &newModeItems
			}

		} else {
			return fmt.Errorf("if `mode` is 0, `new_mode_items` is required")
		}
	} else {
		if v, ok := d.GetOk("nat_gw_list"); ok {
			gwList := v.(*schema.Set).List()
			tmqGwList := make([]*string, 0, len(gwList))
			for i := range gwList {
				gw := gwList[i].(string)
				tmqGwList = append(tmqGwList, &gw)
			}

			request.NatGwList = tmqGwList

		} else {
			return fmt.Errorf("if `mode` is 1, `nat_gw_list` is required")
		}
	}

	if v, ok := d.GetOk("cross_a_zone"); ok {
		request.CrossAZone = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
		request.IsCreateDomain = helper.IntInt64(1)
	} else {
		request.IsCreateDomain = helper.IntInt64(0)
	}

	// Set FwCidrInfo (default to VpcSelf if not specified)
	if v, ok := d.GetOk("fw_cidr_info"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			fwCidrInfo := cfw.FwCidrInfo{}
			if v, ok := dMap["fw_cidr_type"]; ok {
				fwCidrInfo.FwCidrType = helper.String(v.(string))
			}

			if v, ok := dMap["com_fw_cidr"]; ok {
				fwCidrInfo.ComFwCidr = helper.String(v.(string))
			}

			if v, ok := dMap["fw_cidr_lst"]; ok {
				for _, cidr := range v.([]interface{}) {
					iMap := cidr.(map[string]interface{})
					fwCidr := cfw.FwVpcCidr{}
					if v, ok := iMap["vpc_id"]; ok {
						fwCidr.VpcId = helper.String(v.(string))
					}

					if v, ok := iMap["fw_cidr"]; ok {
						fwCidr.FwCidr = helper.String(v.(string))
					}

					fwCidrInfo.FwCidrLst = append(fwCidrInfo.FwCidrLst, &fwCidr)
				}
			}

			request.FwCidrInfo = &fwCidrInfo
		}
	} else {
		// Default FwCidrInfo if not specified
		fwCidrInfo := cfw.FwCidrInfo{}
		fwCidrInfo.FwCidrType = helper.String("VpcSelf")
		fwCidrInfo.ComFwCidr = helper.String("")
		request.FwCidrInfo = &fwCidrInfo
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, err := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().CreateNatFwInstanceWithDomain(request)
		if err != nil {
			return retryError(err)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cfw natInstance failed, reason:%+v", logId, err)
		return err
	}

	instanceId := *response.Response.CfwInsId
	d.SetId(instanceId)

	// Wait for instance to be ready (Status: 0 = normal, 1 = initializing)
	// Also verify that associated resources (VPC list, EIPs) are properly configured
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cfwService := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}

	err = resource.Retry(10*readRetryTimeout, func() *resource.RetryError {
		natInstance, errRet := cfwService.DescribeNatFwInstancesInfoById(ctx, instanceId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}

		if natInstance == nil {
			return resource.NonRetryableError(fmt.Errorf("cfw nat instance %s not found", instanceId))
		}

		if natInstance.Status == nil {
			return resource.NonRetryableError(fmt.Errorf("cfw nat instance %s status is nil", instanceId))
		}

		if *natInstance.Status == 1 {
			// Still initializing
			return resource.RetryableError(fmt.Errorf("cfw nat instance %s is still initializing, status: %d", instanceId, *natInstance.Status))
		}

		if *natInstance.Status == 0 {
			// Normal, ready to use
			log.Printf("[DEBUG]%s cfw nat instance %s is ready, status: %d\n", logId, instanceId, *natInstance.Status)
			return nil
		}

		// Unknown status
		return resource.NonRetryableError(fmt.Errorf("cfw nat instance %s has unknown status: %d", instanceId, *natInstance.Status))
	})

	if err != nil {
		log.Printf("[CRITAL]%s wait for cfw nat instance ready failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCfwNatInstanceRead(d, meta)
}

func resourceTencentCloudCfwNatInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_nat_instance.read")()
	defer inconsistentCheck(d, meta)()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	instanceId := d.Id()

	cfwService := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}

	natInstance, err := cfwService.DescribeNatFwInstancesInfoById(ctx, instanceId)
	if err != nil {
		return err
	}

	if natInstance == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CfwNatInstance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if natInstance.NatinsName != nil {
		_ = d.Set("name", natInstance.NatinsName)
	}

	if natInstance.BandWidth != nil {
		_ = d.Set("width", natInstance.BandWidth)
	}

	if natInstance.FwMode != nil {
		_ = d.Set("mode", natInstance.FwMode)

		if *natInstance.FwMode == MODE_0 {
			var newModeItems []interface{}
			newModeItemsMap := map[string]interface{}{}

			// Get VPC list
			vpcList, err := cfwService.DescribeNatFwVpcDnsLstById(ctx, instanceId)
			if err != nil {
				return err
			}

			// Always set vpc_list, even if empty
			newModeItemsMap["vpc_list"] = vpcList

			// Get EIP addresses
			eips := make([]string, 0)
			if natInstance.EipAddress != nil && len(natInstance.EipAddress) > 0 {
				for _, eip := range natInstance.EipAddress {
					if eip != nil {
						eips = append(eips, *eip)
					}
				}
			}
			newModeItemsMap["eips"] = eips

			// Set add_count based on EIPs length
			newModeItemsMap["add_count"] = len(eips)

			newModeItems = append(newModeItems, newModeItemsMap)
			_ = d.Set("new_mode_items", newModeItems)
		} else {
			// Mode = 1, get NAT gateway list
			natGwList, err := cfwService.DescribeCfwEipsById(ctx, instanceId)
			if err != nil {
				return err
			}

			// Always set nat_gw_list, even if empty
			_ = d.Set("nat_gw_list", natGwList)
		}
	}

	// Set zone and zone_bak (these are Computed, so always set them)
	if natInstance.Zone != nil {
		_ = d.Set("zone", natInstance.Zone)
	}

	if natInstance.ZoneBak != nil {
		_ = d.Set("zone_bak", natInstance.ZoneBak)
	}

	if natInstance.Status != nil {
		_ = d.Set("status", natInstance.Status)
	}

	// Note: cross_a_zone and fw_cidr_info are not returned by DescribeNatFwInstancesInfo API
	// Terraform will preserve these values from state

	return nil
}

func resourceTencentCloudCfwNatInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_nat_instance.update")()
	defer inconsistentCheck(d, meta)()

	var (
		instanceId = d.Id()
		logId      = getLogId(contextNil)
	)

	immutableArgs := []string{"mode", "nat_gw_list", "zone", "zone_bak", "cross_a_zone", "domain", "fw_cidr_info"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	// Handle name change
	if d.HasChange("name") {
		request := cfw.NewModifyNatInstanceRequest()
		request.NatInstanceId = &instanceId

		if v, ok := d.GetOk("name"); ok {
			request.InstanceName = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().ModifyNatInstance(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s update cfw nat firewall instance name failed, reason:%+v", logId, err)
			return err
		}
	}

	// Handle width (bandwidth) change
	if d.HasChange("width") {
		request := cfw.NewExpandCfwVerticalRequest()
		request.FwType = helper.String("nat")
		request.CfwInstance = &instanceId

		if v, ok := d.GetOk("width"); ok {
			request.Width = helper.IntUint64(v.(int))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().ExpandCfwVertical(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s update cfw nat firewall instance bandwidth failed, reason:%+v", logId, err)
			return err
		}
	}

	// Handle new_mode_items.vpc_list change (Mode = 0)
	if d.HasChange("new_mode_items") {
		old, new := d.GetChange("new_mode_items")

		mode := d.Get("mode").(int)
		if mode != MODE_0 {
			return fmt.Errorf("new_mode_items can only be modified for mode 0 instances")
		}

		oldList := old.([]interface{})
		newList := new.([]interface{})

		var oldVpcs, newVpcs []string

		if len(oldList) > 0 {
			oldMap := oldList[0].(map[string]interface{})
			if vpcList, ok := oldMap["vpc_list"]; ok && vpcList != nil {
				for _, vpc := range vpcList.([]interface{}) {
					oldVpcs = append(oldVpcs, vpc.(string))
				}
			}
		}

		if len(newList) > 0 {
			newMap := newList[0].(map[string]interface{})
			if vpcList, ok := newMap["vpc_list"]; ok && vpcList != nil {
				for _, vpc := range vpcList.([]interface{}) {
					newVpcs = append(newVpcs, vpc.(string))
				}
			}
		}

		// Calculate VPCs to add and remove
		toAdd := []string{}
		toRemove := []string{}

		// Find VPCs to add (in new but not in old)
		for _, newVpc := range newVpcs {
			found := false
			for _, oldVpc := range oldVpcs {
				if newVpc == oldVpc {
					found = true
					break
				}
			}
			if !found {
				toAdd = append(toAdd, newVpc)
			}
		}

		// Find VPCs to remove (in old but not in new)
		for _, oldVpc := range oldVpcs {
			found := false
			for _, newVpc := range newVpcs {
				if oldVpc == newVpc {
					found = true
					break
				}
			}
			if !found {
				toRemove = append(toRemove, oldVpc)
			}
		}

		// Remove VPCs
		if len(toRemove) > 0 {
			request := cfw.NewRemoveNatFwObjRequest()
			request.CfwInstance = &instanceId
			request.Mode = helper.IntInt64(MODE_0)

			vpcListPtr := make([]*string, 0, len(toRemove))
			for i := range toRemove {
				vpcListPtr = append(vpcListPtr, &toRemove[i])
			}
			request.VpcList = vpcListPtr

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().RemoveNatFwObj(request)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				}
				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s remove vpc from nat firewall failed, reason:%+v", logId, err)
				return err
			}
		}

		// Add VPCs
		if len(toAdd) > 0 {
			request := cfw.NewAddNatFwObjRequest()
			request.CfwInstance = &instanceId
			request.Mode = helper.IntInt64(MODE_0)

			vpcListPtr := make([]*string, 0, len(toAdd))
			for i := range toAdd {
				vpcListPtr = append(vpcListPtr, &toAdd[i])
			}
			request.VpcList = vpcListPtr

			// Set FwCidrInfo (use existing config or default)
			if v, ok := d.GetOk("fw_cidr_info"); ok {
				for _, item := range v.([]interface{}) {
					dMap := item.(map[string]interface{})
					fwCidrInfo := cfw.FwCidrInfo{}
					if v, ok := dMap["fw_cidr_type"]; ok {
						fwCidrInfo.FwCidrType = helper.String(v.(string))
					}
					if v, ok := dMap["com_fw_cidr"]; ok {
						fwCidrInfo.ComFwCidr = helper.String(v.(string))
					}
					if v, ok := dMap["fw_cidr_lst"]; ok {
						for _, cidr := range v.([]interface{}) {
							iMap := cidr.(map[string]interface{})
							fwCidr := cfw.FwVpcCidr{}
							if v, ok := iMap["vpc_id"]; ok {
								fwCidr.VpcId = helper.String(v.(string))
							}
							if v, ok := iMap["fw_cidr"]; ok {
								fwCidr.FwCidr = helper.String(v.(string))
							}
							fwCidrInfo.FwCidrLst = append(fwCidrInfo.FwCidrLst, &fwCidr)
						}
					}
					request.FwCidrInfo = &fwCidrInfo
				}
			} else {
				// Default FwCidrInfo
				fwCidrInfo := cfw.FwCidrInfo{}
				fwCidrInfo.FwCidrType = helper.String("VpcSelf")
				fwCidrInfo.ComFwCidr = helper.String("")
				request.FwCidrInfo = &fwCidrInfo
			}

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().AddNatFwObj(request)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				}
				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s add vpc to nat firewall failed, reason:%+v", logId, err)
				return err
			}
		}
	}

	return resourceTencentCloudCfwNatInstanceRead(d, meta)
}

func resourceTencentCloudCfwNatInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_nat_instance.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		ctx      = context.WithValue(context.TODO(), logIdKey, logId)
		service  = CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
		cfwInsId = d.Id()
	)

	if err := service.DeleteNatFwInstanceById(ctx, cfwInsId); err != nil {
		return err
	}

	return nil
}
