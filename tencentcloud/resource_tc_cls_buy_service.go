/*
Provides a resource to create a cls buy_service

Example Usage

```hcl
resource "tencentcloudenterprise_cls_buy_service" "buy_service" {}
```
*/
package tencentcloud

import (
	"log"
	"strings"

	bill "terraform-provider-tencentcloudenterprise/sdk/bill/v20181025"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_buy_service", CNDescription{
		TerraformTypeCN: "开通CLS服务",
		DescriptionCN:   "提供CLS服务开通资源，用于开通日志服务（CLS）。",
		AttributesCN: map[string]string{
			"status": "账户状态，0:未开通，1:正常，2:欠费，3:销毁",
		},
	})
}

func resourceTencentCloudClsBuyService() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a CLS buy_service.",
		Create:      resourceTencentCloudClsBuyServiceCreate,
		Read:        resourceTencentCloudClsBuyServiceRead,
		Delete:      resourceTencentCloudClsBuyServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"status": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Account status. 0: not activated, 1: normal, 2: overdue, 3: destroyed.",
			},
		},
	}
}

func resourceTencentCloudClsBuyServiceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_buy_service.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := bill.NewHourCreateRequest()
	goodsInfo := &bill.GoodsInfoList{
		PayMode:  helper.Int64(0),
		GoodsNum: helper.Int64(1),
		ProjectId: helper.Int64(0),
		Type:     helper.String("sp_cls"),
		RegionId: helper.Int64(50000001),
		ZoneId:   helper.Int64(50010001),
		GoodsDetail: helper.String(`{"productCode":"p_cls","subProductCode":"sp_cls","timeUnit":"p","timeSpan":1,"pid":0,"sv_cls_partition_count":1}`),
	}
	request.GoodsInfoList = []*bill.GoodsInfoList{goodsInfo}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBillClient().HourCreate(request)
		if e != nil {
			// Already activated, treat as success
			if sdkErr, ok := e.(*errors.CloudSDKError); ok {
				msg := sdkErr.GetMessage()
				if strings.Contains(msg, "-100") || strings.Contains(msg, "已开通") || strings.Contains(msg, "save resource error") {
					log.Printf("[WARN]%s cls buy service already activated, treat as success: %s", logId, e.Error())
					return nil
				}
			}
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate bill HourCreate failed, reason:%+v", logId, err)
		return err
	}

	d.SetId("cls_service")

	return resourceTencentCloudClsBuyServiceRead(d, meta)
}

func resourceTencentCloudClsBuyServiceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_buy_service.read")()

	logId := getLogId(contextNil)

	request := cls.NewDescribeAccountRequest()
	var status int64
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().DescribeAccount(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		if result.Response.Status != nil {
			status = *result.Response.Status
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate cls DescribeAccount failed, reason:%+v", logId, err)
		return err
	}

	if status == 0 {
		d.SetId("")
		return nil
	}

	_ = d.Set("status", status)

	return nil
}

func resourceTencentCloudClsBuyServiceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_buy_service.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
