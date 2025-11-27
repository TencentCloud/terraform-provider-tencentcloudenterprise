/*
Use this data source to query detailed information of tsf application_attribute

# Example Usage

```hcl

	data "tencentcloudenterprise_tsf_application_attribute" "application_attribute" {
	  application_id = "application-a24x29xv"
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
	registerDataDescriptionProvider("tencentcloudenterprise_tsf_application_attribute", CNDescription{
		TerraformTypeCN: "TSF应用属性",
		DescriptionCN:   "提供TSF应用属性数据源，用于查询TSF应用属性的详细信息。",
		AttributesCN: map[string]string{
			"application_id":     "应用ID",
			"result":             "应用列表其他属性",
			"instance_count":     "实例总数",
			"run_instance_count": "运行实例数",
			"group_count":        "部署组数量",
			"result_output_file": "用于保存结果",
		},
	})

}

func dataSourceTencentCloudTsfApplicationAttribute() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides detailed information of tsf application_attribute",
		Read:        dataSourceTencentCloudTsfApplicationAttributeRead,
		Schema: map[string]*schema.Schema{
			"application_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Application Id.",
			},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Application list other attribute.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total number of instances.Note: This field may return null, indicating that no valid values can be obtained.",
						},
						"run_instance_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of running instances.Note: This field may return null, indicating that no valid values can be obtained.",
						},
						"group_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of deployment groups under the application.Note: This field may return null, indicating that no valid values can be obtained.",
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

func dataSourceTencentCloudTsfApplicationAttributeRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tsf_application_attribute.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	ids := ""

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("application_id"); ok {
		ids = v.(string)
		paramMap["ApplicationId"] = helper.String(v.(string))
	}

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	var attribute *tsf.ApplicationAttribute

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTsfApplicationAttributeByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		attribute = result
		return nil
	})
	if err != nil {
		return err
	}

	applicationAttributeMap := map[string]interface{}{}
	if attribute != nil {
		if attribute.InstanceCount != nil {
			applicationAttributeMap["instance_count"] = attribute.InstanceCount
		}

		if attribute.RunInstanceCount != nil {
			applicationAttributeMap["run_instance_count"] = attribute.RunInstanceCount
		}

		if attribute.GroupCount != nil {
			applicationAttributeMap["group_count"] = attribute.GroupCount
		}

		_ = d.Set("result", []interface{}{applicationAttributeMap})
	}

	d.SetId(ids)
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), applicationAttributeMap); e != nil {
			return e
		}
	}
	return nil
}
