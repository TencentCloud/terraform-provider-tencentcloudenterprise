/*
Provide a resource to activate CFW (Cloud Firewall) service

# Example Usage

```hcl

	resource "tencentcloudenterprise_cfw_buy_service" "example" {
	  region_id = "50000001"
	  zone_id   = "50010001"
	  vpc_spec  = "2"
	}

```
*/
package tencentcloud

import (
	"encoding/json"
	"fmt"
	"log"

	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfw_buy_service", CNDescription{
		TerraformTypeCN: "CFW 开通",
		DescriptionCN:   "提供 CFW 云防火墙服务开通资源，用于开通云防火墙服务。",
		AttributesCN: map[string]string{
			"region_id":    "地域ID",
			"zone_id":      "可用区ID",
			"vpc_spec":     "VPC规格，1:标准版 2:专业版",
			"resource_id":  "资源实例ID",
			"status":       "服务状态，0:未开通 1:已开通",
			"pay_mode":     "计费模式",
		},
	})
}

func resourceTencentCloudCfwBuyService() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to activate CFW (Cloud Firewall) service.",
		Create:      resourceTencentCloudCfwBuyServiceCreate,
		Read:        resourceTencentCloudCfwBuyServiceRead,
		Delete:      resourceTencentCloudCfwBuyServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Region ID for CFW service activation.",
			},

			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone ID for CFW service activation.",
			},

			"vpc_spec": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VPC spec, 1: standard edition, 2: professional edition.",
			},

			// computed
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource instance ID.",
			},

			"status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Service status, 0: not activated, 1: activated.",
			},

			"pay_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Billing mode.",
			},
		},
	}
}

func resourceTencentCloudCfwBuyServiceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_buy_service.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId = getLogId(contextNil)
	)

	// Build ReqParams JSON for CreateBilling command
	reqParams := map[string]interface{}{
		"Cmd":      "CreateBilling",
		"RegionId": d.Get("region_id").(string),
		"ZoneId":   d.Get("zone_id").(string),
		"GoodsParams": map[string]string{
			"VpcSpec": d.Get("vpc_spec").(string),
		},
	}

	reqParamsJson, err := json.Marshal(reqParams)
	if err != nil {
		return fmt.Errorf("marshal ReqParams failed: %s", err)
	}

	request := cfw.NewDescribeApiDispatchRequest()
	request.ReqParams = helper.String(string(reqParamsJson))

	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().DescribeApiDispatch(request)
		if e != nil {
			// Already activated, treat as success
			if sdkErr, ok := e.(*errors.CloudSDKError); ok && sdkErr.GetCode() == "FailedOperation" {
				log.Printf("[WARN]%s cfw buy service already activated, treat as success: %s", logId, e.Error())
				return nil
			}
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("Create CFW buy service failed, Response is nil."))
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s create cfw buy service failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	// Wait for activation by polling DescribeConfig
	reqErr = resource.Retry(readRetryTimeout*7, func() *resource.RetryError {
		configReq := cfw.NewDescribeConfigRequest()
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().DescribeConfig(configReq)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, response body [%s]\n", logId, configReq.GetAction(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.Data == nil {
			return resource.RetryableError(fmt.Errorf("CFW service is still activating..."))
		}

		var configData struct {
			Status     int    `json:"status"`
			ResourceID string `json:"resourceID"`
		}
		if err := json.Unmarshal([]byte(*result.Response.Data), &configData); err != nil {
			return resource.NonRetryableError(fmt.Errorf("unmarshal DescribeConfig Data failed: %s", err))
		}

		if configData.Status == 1 {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("CFW service is still activating... status is %d", configData.Status))
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s wait cfw buy service activation failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	// Use region_id as the resource ID (CFW service is one per region)
	d.SetId(d.Get("region_id").(string))

	return resourceTencentCloudCfwBuyServiceRead(d, meta)
}

func resourceTencentCloudCfwBuyServiceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_buy_service.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId = getLogId(contextNil)
	)

	request := cfw.NewDescribeConfigRequest()
	result, err := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().DescribeConfig(request)
	if err != nil {
		return err
	}

	if result == nil || result.Response == nil || result.Response.Data == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_cfw_buy_service` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	var configData struct {
		Status     int    `json:"status"`
		ResourceID string `json:"resourceID"`
		VpcSpec    string `json:"vpcSpec"`
		PayMode    string `json:"payMode"`
	}
	if err := json.Unmarshal([]byte(*result.Response.Data), &configData); err != nil {
		return fmt.Errorf("unmarshal DescribeConfig Data failed: %s", err)
	}

	if configData.Status != 1 {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_cfw_buy_service` [%s] not activated, status is %d.\n", logId, d.Id(), configData.Status)
		d.SetId("")
		return nil
	}

	_ = d.Set("resource_id", configData.ResourceID)
	_ = d.Set("status", configData.Status)
	_ = d.Set("pay_mode", configData.PayMode)
	if configData.VpcSpec != "" {
		_ = d.Set("vpc_spec", configData.VpcSpec)
	}

	return nil
}

func resourceTencentCloudCfwBuyServiceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_buy_service.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
