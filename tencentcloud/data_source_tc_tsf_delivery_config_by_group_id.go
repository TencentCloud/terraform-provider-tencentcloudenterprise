/*
Use this data source to query detailed information of tsf delivery_config_by_group_id

# Example Usage

```hcl

	data "tencentcloudenterprise_tsf_delivery_config_by_group_id" "delivery_config_by_group_id" {
	  group_id = "group-yrjkln9v"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tsf_delivery_config_by_group_id", CNDescription{
		TerraformTypeCN: "TSF分组下配置项",
		DescriptionCN:   "提供TSF分组下配置项数据源，用于查询TSF分组下配置项的详细信息。",
		AttributesCN: map[string]string{
			"group_id":           "分组ID",
			"result":             "配置项列表",
			"config_id":          "配置ID",
			"config_name":        "配置名称",
			"result_output_file": "用于保存结果",
		},
	})

}

func dataSourceTencentCloudTsfDeliveryConfigByGroupId() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides detailed information of tsf delivery_config_by_group_id",
		Read:        dataSourceTencentCloudTsfDeliveryConfigByGroupIdRead,
		Schema: map[string]*schema.Schema{
			"group_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "GroupId.",
			},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Configuration item for deliver to a Kafka.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Config ID. Note: This field may return null, which means that no valid value was obtained.",
						},
						"config_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Config Name. Note: This field may return null, which means that no valid value was obtained.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudTsfDeliveryConfigByGroupIdRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tsf_delivery_config_by_group_id.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var groupId string

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("group_id"); ok {
		groupId = v.(string)
		paramMap["GroupId"] = helper.String(v.(string))
	}

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	var deliveryConfig *tsf.SimpleKafkaDeliveryConfig
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTsfDeliveryConfigByGroupIdByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		deliveryConfig = result
		return nil
	})
	if err != nil {
		return err
	}

	simpleKafkaDeliveryConfigMap := map[string]interface{}{}
	if deliveryConfig != nil {
		if deliveryConfig.ConfigId != nil {
			simpleKafkaDeliveryConfigMap["config_id"] = deliveryConfig.ConfigId
		}

		if deliveryConfig.ConfigName != nil {
			simpleKafkaDeliveryConfigMap["config_name"] = deliveryConfig.ConfigName
		}

		_ = d.Set("result", []interface{}{simpleKafkaDeliveryConfigMap})
	}

	d.SetId(groupId)
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), simpleKafkaDeliveryConfigMap); e != nil {
			return e
		}
	}
	return nil
}
