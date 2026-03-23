/*
Provides a CWP (Cloud Workload Protection) license order resource.

~> **NOTE:** This resource allows you to create and manage CWP license orders with automatic scaling support.
~> **NOTE:** When modifying `license_num`, the system will automatically handle expansion (increase) or shrinkage (decrease) operations.
~> **NOTE:** `auto_bind_switch` and `auto_repurchase_switch` are only available during resource creation.

Example Usage

```hcl
# Basic CWP license order creation
resource "tencentcloudenterprise_cwp_license_order" "example" {
  region_id    = 50000001  # Chongqing region
  alias        = "Production Environment CWP License"
  license_type = 5         # 5 = FLAGSHIP version, 0 = PRO version
  license_num  = 1         # Initial license quantity
  
  # Auto-bind and auto-repurchase settings (creation only)
  auto_bind_switch       = true
  auto_repurchase_switch = true
  
  project_id = 0
  tags = {
    Environment = "production"
    Team        = "security"
    CostCenter  = "IT"
  }
}

# License expansion example - modify license_num to expand
resource "tencentcloudenterprise_cwp_license_order" "expansion_example" {
  region_id    = 50000001
  alias        = "Expanded CWP License"
  license_type = 5
  license_num  = 5         # Expand from 1 to 5 licenses (system will add 4 more)
  
  tags = {
    Environment = "production"
    Action      = "expansion"
  }
}

# License shrinkage example - modify license_num to shrink
resource "tencentcloudenterprise_cwp_license_order" "shrink_example" {
  region_id    = 50000001
  alias        = "Shrunk CWP License"
  license_type = 5
  license_num  = 2         # Shrink from 5 to 2 licenses (system will remove 3)
  
  tags = {
    Environment = "production"
    Action      = "shrinkage"
  }
}

# Modify alias only (no license quantity change)
resource "tencentcloudenterprise_cwp_license_order" "modify_alias" {
  region_id    = 50000001
  alias        = "Updated Resource Alias"  # Only change alias
  license_type = 5
  license_num  = 2         # Keep same quantity
  
  tags = {
    Environment = "production"
    Modified    = "true"
  }
}

# Import existing license order
# terraform import tencentcloudenterprise_cwp_license_order.example cwplic-2c263017#50000001
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	cwp "terraform-provider-tencentcloudenterprise/sdk/cwp/v20180228"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cwp_license_order", CNDescription{
		TerraformTypeCN: "主机安全许可证订单",
		DescriptionCN:   "提供主机安全许可证订单资源，用于创建和管理主机安全许可证订单。",
		AttributesCN: map[string]string{
			"alias":                  "资源别名",
			"license_type":           "许可证类型",
			"license_num":            "许可证数量，创建时指定初始数量，修改时支持扩容或缩容",
			"region_id":              "购买订单地域",
			"project_id":             "项目ID",
			"tags":                   "许可证标签",
			"auto_bind_switch":       "自动绑定开关",
			"auto_repurchase_switch": "自动加购开关",
		},
	})
}

func resourceTencentCloudCwpLicenseOrder() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudCwpLicenseOrderCreate,
		Read:        resourceTencentCloudCwpLicenseOrderRead,
		Update:      resourceTencentCloudCwpLicenseOrderUpdate,
		Delete:      resourceTencentCloudCwpLicenseOrderDelete,
		Description: "Provides a resource to create and manage CWP license order",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"alias": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Resource alias.",
			},
			"license_type": {
				Optional:     true,
				Type:         schema.TypeInt,
				Default:      LICENSE_TYPE_0,
				Description:  "LicenseType, 0 CWP Pro - Pay as you go, 1 CWP Pro - Monthly subscription, 2 CWP Ultimate - Monthly subscription. Default is 0.",
			},
			"license_num": {
				Optional:    true,
				Type:        schema.TypeInt,
				Default:     1,
				Description: "License quantity. Initial quantity for creation, and can be modified for expansion or shrinkage.",
			},
			"region_id": {
				Required:     true,
				Type:         schema.TypeInt,
				Description:  "Purchase order region.",
			},
			"project_id": {
				Optional:    true,
				Type:        schema.TypeInt,
				Default:     0,
				Description: "Project ID. Default is 0.",
			},
			"tags": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "Tags of the license order.",
					},
			"auto_bind_switch": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "Auto bind switch for license order. Only available during creation. Default is false.",
				},
			"auto_repurchase_switch": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Auto repurchase switch for license order. Only available during creation. Default is false.",
			},
	
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "resource id.",
			},
			"license_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "license id.",
			},
			"used_license_cnt": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "used license count.",
			},
		},
	}
}

func resourceTencentCloudCwpLicenseOrderCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cwp_license_order.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		request     = cwp.NewCreateLicenseOrderRequest()
		response    = cwp.NewCreateLicenseOrderResponse()
		tagService  = TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		resourceId  string
		regionIdInt int
	)

	if v, ok := d.GetOkExists("license_type"); ok {
		request.LicenseType = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("license_num"); ok {
		request.LicenseNum = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("region_id"); ok {
		request.RegionId = helper.IntUint64(v.(int))
		regionIdInt = v.(int)
	}

	if v, ok := d.GetOkExists("project_id"); ok {
		request.ProjectId = helper.IntUint64(v.(int))
	}

	// Set BillingDefineParams if provided
	if d.Get("auto_bind_switch") != nil || d.Get("auto_repurchase_switch") != nil {
		billingParams := &cwp.BillingDefineParams{}
		if v, ok := d.GetOkExists("auto_bind_switch"); ok {
			billingParams.AutoBindSwitch = helper.Bool(v.(bool))
		}
		if v, ok := d.GetOkExists("auto_repurchase_switch"); ok {
			billingParams.AutoRepurchaseSwitch = helper.Bool(v.(bool))
		}
		request.BillingDefineParams = billingParams
	}



	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().CreateLicenseOrder(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || len(result.Response.ResourceIds) != 1 {
			e = fmt.Errorf("cwp licenseOrder not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cwp licenseOrder failed, reason:%+v", logId, err)
		return err
	}

	resourceId = *response.Response.ResourceIds[0]
	regionId := strconv.Itoa(regionIdInt)
	d.SetId(strings.Join([]string{resourceId, regionId}, FILED_SP))

	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::cwp:%s:uin/:order/%s", region, resourceId)
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	// set alias using ModifyOrderAttribute API
	if v, ok := d.GetOk("alias"); ok {
		aliasRequest := cwp.NewModifyOrderAttributeRequest()
		aliasRequest.AttrName = helper.String("alias")
		aliasRequest.AttrValue = helper.String(v.(string))
		aliasRequest.ResourceId = &resourceId
		
		if licenseType, ok := d.GetOkExists("license_type"); ok {
			aliasRequest.LicenseType = helper.IntUint64(licenseType.(int))
		}

		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().ModifyOrderAttribute(aliasRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, aliasRequest.GetAction(), aliasRequest.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s set cwp licenseOrder alias failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCwpLicenseOrderRead(d, meta)
}

func resourceTencentCloudCwpLicenseOrderRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cwp_license_order.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
		tagService = TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	resourceId := idSplit[0]
	regionId := idSplit[1]

	licenseOrder, err := service.DescribeCwpLicenseOrderById(ctx, resourceId)
	if err != nil {
		return err
	}

	if licenseOrder == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CwpLicenseOrder` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	regionIdInt, _ := strconv.Atoi(regionId)
	_ = d.Set("region_id", regionIdInt)

	if licenseOrder.Alias != nil {
		_ = d.Set("alias", licenseOrder.Alias)
	}

	if licenseOrder.ResourceId != nil {
		_ = d.Set("resource_id", licenseOrder.ResourceId)
	}

	if licenseOrder.LicenseId != nil {
		_ = d.Set("license_id", licenseOrder.LicenseId)
	}

	if licenseOrder.LicenseType != nil {
		_ = d.Set("license_type", licenseOrder.LicenseType)
	}

	if licenseOrder.LicenseCnt != nil {
		_ = d.Set("license_num", licenseOrder.LicenseCnt)
	}

	if licenseOrder.UsedLicenseCnt != nil {
		_ = d.Set("used_license_cnt", licenseOrder.UsedLicenseCnt)
	}

	if licenseOrder.ProjectId != nil {
		_ = d.Set("project_id", licenseOrder.ProjectId)
	}

	// Note: auto_bind_switch and auto_repurchase_switch are not returned in LicenseDetail
	// These fields are only used for creation/update operations

	region := meta.(*TencentCloudClient).apiV3Conn.Region
	tags, err := tagService.DescribeResourceTags(ctx, "cwp", "order", region, resourceId)
	if err != nil {
		return err
	}

	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudCwpLicenseOrderUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cwp_license_order.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		tagService = TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	immutableArgs := []string{"license_type", "region_id", "auto_bind_switch", "auto_repurchase_switch"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	resourceId := idSplit[0]

	// Handle alias changes using ModifyOrderAttribute API
	if d.HasChange("alias") {
		if v, ok := d.GetOk("alias"); ok {
			modifyAttrRequest := cwp.NewModifyOrderAttributeRequest()
			modifyAttrRequest.AttrName = helper.String("alias")
			modifyAttrRequest.AttrValue = helper.String(v.(string))
			modifyAttrRequest.ResourceId = &resourceId
			
			if licenseType, ok := d.GetOkExists("license_type"); ok {
				modifyAttrRequest.LicenseType = helper.IntUint64(licenseType.(int))
			}

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().ModifyOrderAttribute(modifyAttrRequest)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, modifyAttrRequest.GetAction(), modifyAttrRequest.ToJsonString(), result.ToJsonString())
				}
				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s update cwp licenseOrder alias failed, reason:%+v", logId, err)
				return err
			}
		}
	}

	// Handle project_id changes using ModifyOrderAttribute API
	if d.HasChange("project_id") {
		if v, ok := d.GetOkExists("project_id"); ok {
			modifyAttrRequest := cwp.NewModifyOrderAttributeRequest()
			modifyAttrRequest.AttrName = helper.String("projectId")
			modifyAttrRequest.AttrValue = helper.String(strconv.Itoa(v.(int)))
			modifyAttrRequest.ResourceId = &resourceId
			
			if licenseType, ok := d.GetOkExists("license_type"); ok {
				modifyAttrRequest.LicenseType = helper.IntUint64(licenseType.(int))
			}

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().ModifyOrderAttribute(modifyAttrRequest)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, modifyAttrRequest.GetAction(), modifyAttrRequest.ToJsonString(), result.ToJsonString())
				}
				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s update cwp licenseOrder project_id failed, reason:%+v", logId, err)
				return err
			}
		}
	}

// Handle license_num expansion/shrink using CreateLicenseOrder with ModifyConfig
	if d.HasChange("license_num") {
		oldNum, newNum := d.GetChange("license_num")
		oldNumInt := oldNum.(int)
		newNumInt := newNum.(int)
		
		// Calculate the difference (positive for expansion, negative for shrink)
		inquireNum := newNumInt - oldNumInt
		
		if inquireNum != 0 {
			createRequest := cwp.NewCreateLicenseOrderRequest()
			
			// Set basic required fields
			if v, ok := d.GetOkExists("license_type"); ok {
				createRequest.LicenseType = helper.IntUint64(v.(int))
			}
			
			if v, ok := d.GetOkExists("region_id"); ok {
				createRequest.RegionId = helper.IntUint64(v.(int))
			}
			
			// Set ModifyConfig for expansion/shrink
			modifyConfig := &cwp.OrderModifyObject{}
			modifyConfig.ResourceId = &resourceId
			
			// Get current license details to determine the product code
			service := CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
			licenseOrder, err := service.DescribeCwpLicenseOrderById(ctx, resourceId)
			if err != nil {
				return fmt.Errorf("failed to get current license details: %v", err)
			}
			
			// Determine product code based on license type
			// LicenseType 0 = PRO_VERSION, 5 = FLAGSHIP
			if licenseOrder != nil && licenseOrder.LicenseType != nil {
				switch *licenseOrder.LicenseType {
				case 0:
					modifyConfig.NewSubProductCode = helper.String("PRO_VERSION")
				case 5:
					modifyConfig.NewSubProductCode = helper.String("FLAGSHIP")
				default:
					// Default to FLAGSHIP for unknown types
					modifyConfig.NewSubProductCode = helper.String("FLAGSHIP")
				}
			} else {
				// Default to FLAGSHIP for expansion/shrink
				modifyConfig.NewSubProductCode = helper.String("FLAGSHIP")
			}
			
			// Set the difference (positive for expansion, negative for shrink)
			modifyConfig.InquireNum = helper.IntInt64(inquireNum)
			
			createRequest.ModifyConfig = modifyConfig

			err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().CreateLicenseOrder(createRequest)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, createRequest.GetAction(), createRequest.ToJsonString(), result.ToJsonString())
				}
				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s modify cwp licenseOrder config failed, reason:%+v", logId, err)
				return err
			}
		}
	}

	// Note: auto_bind_switch and auto_repurchase_switch are only available during creation
	// These cannot be modified after resource creation according to API design

	if d.HasChange("tags") {
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := tagService.DiffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("cwp", "order", "", resourceId)
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudCwpLicenseOrderRead(d, meta)
}

func resourceTencentCloudCwpLicenseOrderDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cwp_license_order.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		service     = CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
		licenseType *uint64
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	resourceId := idSplit[0]

	if v, ok := d.GetOkExists("license_type"); ok {
		licenseType = helper.IntUint64(v.(int))
	}

	if err := service.DeleteCwpLicenseOrderById(ctx, resourceId, licenseType); err != nil {
		return err
	}

	return nil
}