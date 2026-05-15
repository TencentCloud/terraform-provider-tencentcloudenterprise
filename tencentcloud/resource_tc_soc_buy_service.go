/*
Provides a resource to create a soc buy_service

Example Usage

```hcl
resource "tencentcloudenterprise_soc_buy_service" "buy_service" {
  type = "premium"
  tce_area {
    region_id = 50000001
    zone_id   = 50010001
  }
}
```
*/
package tencentcloud

import (
	"fmt"
	"log"
	"strings"

	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	yjn "terraform-provider-tencentcloudenterprise/sdk/yjn/v20240320"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_soc_buy_service", CNDescription{
		TerraformTypeCN: "开通SOC服务",
		DescriptionCN:   "提供SOC服务开通资源，用于开通安全运营中心（SOC）服务。",
		AttributesCN: map[string]string{
			"type":        "计费模型，如 premium",
			"tce_area":    "开通地域可用区信息",
			"region_id":   "地域ID",
			"zone_id":     "可用区ID",
			"resource_id": "资源ID",
			"buy_status":  "服务是否已开通",
		},
	})
}

func resourceTencentCloudSocBuyService() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a SOC buy_service.",
		Create:      resourceTencentCloudSocBuyServiceCreate,
		Read:        resourceTencentCloudSocBuyServiceRead,
		Delete:      resourceTencentCloudSocBuyServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"type": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Default:     "premium",
				Description: "Billing model type, e.g. premium.",
			},
			"tce_area": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Region and zone info for service activation.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Region ID.",
						},
						"zone_id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Zone ID.",
						},
					},
				},
			},
			"resource_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Resource ID returned after activation.",
			},
			"buy_status": {
				Computed:    true,
				Type:        schema.TypeBool,
				Description: "Whether the SOC service is purchased. true: purchased; false: not purchased.",
			},
		},
	}
}

func resourceTencentCloudSocBuyServiceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_soc_buy_service.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// Check if already activated via GetBuyStatus
	var regionId string
	if v, ok := d.GetOk("tce_area"); ok {
		areaList := v.([]interface{})
		if len(areaList) > 0 {
			m := areaList[0].(map[string]interface{})
			regionId = helper.Int64ToStr(int64(m["region_id"].(int)))
		}
	}
	if regionId == "" {
		regionId = "50000001"
	}

	statusReq := yjn.NewGetBuyStatusRequest()
	statusReq.OrderBillBuyStatusList = []*yjn.BillBuyStatusInfo{
		{RegionId: helper.String(regionId)},
	}
	statusResult, statusErr := meta.(*TencentCloudClient).apiV3Conn.UseYjnClient().GetBuyStatus(statusReq)
	if statusErr == nil && statusResult.Response != nil {
		activated := false
		if statusResult.Response.Data != nil && statusResult.Response.Data.Result != nil {
			activated = *statusResult.Response.Data.Result
		} else if statusResult.Response.Result != nil {
			activated = *statusResult.Response.Result
		}
		if activated {
			log.Printf("[WARN]%s soc buy service already activated, skip BillingCreate", logId)
			d.SetId("soc_service")
			return resourceTencentCloudSocBuyServiceRead(d, meta)
		}
	}

	request := yjn.NewBillingCreateRequest()

	if v, ok := d.GetOk("type"); ok {
		request.Type = helper.String(v.(string))
	}

	if v, ok := d.GetOk("tce_area"); ok {
		areaList := v.([]interface{})
		tceAreas := make([]*yjn.TceArea, 0, len(areaList))
		for _, item := range areaList {
			m := item.(map[string]interface{})
			tceArea := &yjn.TceArea{}
			if v, ok := m["region_id"]; ok {
				tceArea.RegionId = helper.Int64(int64(v.(int)))
			}
			if v, ok := m["zone_id"]; ok {
				tceArea.ZoneId = helper.Int64(int64(v.(int)))
			}
			tceAreas = append(tceAreas, tceArea)
		}
		request.TceArea = tceAreas
	}

	var resourceId string
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseYjnClient().BillingCreate(request)
		if e != nil {
			// Already activated, treat as success
			if sdkErr, ok := e.(*errors.CloudSDKError); ok {
				msg := sdkErr.GetMessage()
				if strings.Contains(msg, "-100") || strings.Contains(msg, "已开通") || strings.Contains(msg, "save resource error") {
					log.Printf("[WARN]%s soc BillingCreate already activated, treat as success: %s", logId, e.Error())
					return nil
				}
			}
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		if result.Response.Data != nil && result.Response.Data.ResourceIds != nil {
			for _, rs := range result.Response.Data.ResourceIds {
				if len(rs.Value) > 0 {
					resourceId = *rs.Value[0]
					break
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate yjn BillingCreate failed, reason:%+v", logId, err)
		return err
	}

	if resourceId == "" {
		resourceId = "soc_service"
	}
	d.SetId(resourceId)

	// Wait for activation by polling GetBuyStatus
	reqErr := resource.Retry(readRetryTimeout*3, func() *resource.RetryError {
		statusReq := yjn.NewGetBuyStatusRequest()
		statusReq.OrderBillBuyStatusList = []*yjn.BillBuyStatusInfo{
			{RegionId: helper.String(regionId)},
		}
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseYjnClient().GetBuyStatus(statusReq)
		if e != nil {
			return retryError(e)
		}
		activated := false
		if result.Response.Data != nil && result.Response.Data.Result != nil {
			activated = *result.Response.Data.Result
		} else if result.Response.Result != nil {
			activated = *result.Response.Result
		}
		if activated {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("SOC service is still activating..."))
	})
	if reqErr != nil {
		log.Printf("[CRITAL]%s wait SOC buy service activation failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return resourceTencentCloudSocBuyServiceRead(d, meta)
}

func resourceTencentCloudSocBuyServiceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_soc_buy_service.read")()

	logId := getLogId(contextNil)

	var regionId string
	if v, ok := d.GetOk("tce_area"); ok {
		areaList := v.([]interface{})
		if len(areaList) > 0 {
			m := areaList[0].(map[string]interface{})
			regionId = helper.Int64ToStr(int64(m["region_id"].(int)))
		}
	}

	if regionId == "" {
		regionId = "50000001"
	}

	request := yjn.NewGetBuyStatusRequest()
	request.OrderBillBuyStatusList = []*yjn.BillBuyStatusInfo{
		{RegionId: helper.String(regionId)},
	}

	var buyStatus bool
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseYjnClient().GetBuyStatus(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		if result.Response.Data != nil && result.Response.Data.Result != nil {
			buyStatus = *result.Response.Data.Result
		} else if result.Response.Result != nil {
			buyStatus = *result.Response.Result
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate yjn GetBuyStatus failed, reason:%+v", logId, err)
		return err
	}

	if !buyStatus {
		d.SetId("")
		return nil
	}

	_ = d.Set("buy_status", buyStatus)
	_ = d.Set("resource_id", d.Id())

	return nil
}

func resourceTencentCloudSocBuyServiceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_soc_buy_service.delete")()
	defer inconsistentCheck(d, meta)()

	// TODO: Delete requires BillId from GetBillList, but SDK ResultList type is incorrect (single pointer vs array).
	// Will implement after SDK is fixed. BillingUnBlock API is available in SDK.

	return nil
}

