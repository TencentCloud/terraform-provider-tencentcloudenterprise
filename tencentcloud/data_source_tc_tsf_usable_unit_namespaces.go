/*
Use this data source to query detailed information of tsf usable_unit_namespaces

# Example Usage

```hcl

	data "tencentcloudenterprise_tsf_usable_unit_namespaces" "usable_unit_namespaces" {
	  search_word = ""
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
	registerDataDescriptionProvider("tencentcloudenterprise_tsf_usable_unit_namespaces", CNDescription{
		TerraformTypeCN: "TSF可用单元命名空间",
		DescriptionCN:   "提供TSF可用单元命名空间数据源，用于查询TSF可用单元命名空间的详细信息。",
		AttributesCN: map[string]string{
			"search_word":    "搜索关键字（按命名空间ID或命名空间名称搜索）",
			"result":         "命名空间对象列表",
			"total_count":    "总数",
			"content":        "命名空间列表",
			"namespace_id":   "命名空间ID",
			"namespace_name": "命名空间名称",
			"id":             "单元命名空间ID。注意：此字段可能返回 null，表示取不到有效值。",
			//"gateway_instance_id": "网关实体ID。注意：此字段可能返回 null，表示取不到有效值。",
			//"created_time":       "创建时间。注意：此字段可能返回 null，表示取不到有效值。",
			//"updated_time":       "更新时间。注意：此字段可能返回 null，表示取不到有效值。",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudTsfUsableUnitNamespaces() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tsf usable_unit_namespaces",
		Read:        dataSourceTencentCloudTsfUsableUnitNamespacesRead,
		Schema: map[string]*schema.Schema{
			"search_word": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Search by namespace id or namespace Name.",
			},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Namespace object list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total count.",
						},
						"content": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Namespace list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"namespace_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Namespace id.",
									},
									"namespace_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Namespace name.",
									},
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Unit namespace ID. Note: This field may return null, indicating that no valid value was found.",
									},
									//"gateway_instance_id": {
									//	Type:        schema.TypeString,
									//	Computed:    true,
									//	Description: "Gateway instance id Note: This field may return null, indicating that no valid value was found.",
									//},
									//"created_time": {
									//	Type:        schema.TypeString,
									//	Computed:    true,
									//	Description: "Create time. Note: This field may return null, indicating that no valid value was found.",
									//},
									//"updated_time": {
									//	Type:        schema.TypeString,
									//	Computed:    true,
									//	Description: "Update time. Note: This field may return null, indicating that no valid value was found.",
									//},
								},
							},
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

func dataSourceTencentCloudTsfUsableUnitNamespacesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tsf_usable_unit_namespaces.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("search_word"); ok {
		paramMap["SearchWord"] = helper.String(v.(string))
	}

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	var result *tsf.TsfPageUnitNamespace
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		response, e := service.DescribeTsfUsableUnitNamespacesByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		result = response
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(result.Content))
	tsfPageUnitNamespaceMap := map[string]interface{}{}
	if result != nil {

		if result.TotalCount != nil {
			tsfPageUnitNamespaceMap["total_count"] = result.TotalCount
		}

		if result.Content != nil {
			contentList := []interface{}{}
			for _, content := range result.Content {
				contentMap := map[string]interface{}{}

				if content.NamespaceId != nil {
					contentMap["namespace_id"] = content.NamespaceId
				}

				if content.NamespaceName != nil {
					contentMap["namespace_name"] = content.NamespaceName
				}

				if content.Id != nil {
					contentMap["id"] = content.Id
				}

				//if content.GatewayInstanceId != nil {
				//	contentMap["gateway_instance_id"] = content.GatewayInstanceId
				//}
				//
				//if content.CreatedTime != nil {
				//	contentMap["created_time"] = content.CreatedTime
				//}
				//
				//if content.UpdatedTime != nil {
				//	contentMap["updated_time"] = content.UpdatedTime
				//}

				contentList = append(contentList, contentMap)
				ids = append(ids, *content.NamespaceId)
			}

			tsfPageUnitNamespaceMap["content"] = contentList
		}

		_ = d.Set("result", []interface{}{tsfPageUnitNamespaceMap})
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tsfPageUnitNamespaceMap); e != nil {
			return e
		}
	}
	return nil
}
