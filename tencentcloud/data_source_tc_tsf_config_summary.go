/*
Use this data source to query detailed information of tsf config_summary

# Example Usage

```hcl

	data "tencentcloudenterprise_tsf_config_summary" "config_summary" {
		application_id = "application-a24x29xv"
		search_word = "terraform"
		order_by = "last_update_time"
		order_type = 0
		disable_program_auth_check = true
		config_id_list = ["dcfg-y54wzk3a"]
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
	registerDataDescriptionProvider("tencentcloudenterprise_tsf_config_summary", CNDescription{
		TerraformTypeCN: "TSF配置项概要",
		DescriptionCN:   "提供TSF配置概要数据源，用于查询TSF配置概要的详细信息。",
		AttributesCN: map[string]string{
			"application_id":             "应用ID。不传入时查询所有。",
			"search_word":                "查询关键字，模糊查询：应用名，配置项名称。不传入时查询所有。",
			"order_by":                   "排序字段。支持按时间排序：creation_time；或者按名称排序：config_name。",
			"order_type":                 "传入0时为升序，传入1时为降序。",
			"config_tag_list":            "config标签列表。",
			"disable_program_auth_check": "是否禁用程序鉴权。",
			"config_id_list":             "配置项ID列表。",
			"result":                     "配置项分页项。",
			"total_count":                "总条数。",
			"content":                    "配置项列表。",
			"config_id":                  "配置项ID。注意：此字段可能返回 null，表示取不到有效值。",
			"config_name":                "配置项名称。注意：此字段可能返回 null，表示取不到有效值。",
			"config_version":             "配置版本。注意：此字段可能返回 null，表示取不到有效值。",
			"config_version_desc":        "配置版本描述。注意：此字段可能返回 null，表示取不到有效值。",
			"config_value":               "配置值。注意：此字段可能返回 null，表示取不到有效值。",
			"config_type":                "配置类型。注意：此字段可能返回 null，表示取不到有效值。",
			"creation_time":              "创建时间。注意：此字段可能返回 null，表示取不到有效值。",
			"application_name":           "应用名称。注意：此字段可能返回 null，表示取不到有效值。",
			"delete_flag":                "删除标识，true:可删除；false:不可删除。注意：此字段可能返回 null，表示取不到有效值。",
			"last_update_time":           "最后更新时间。注意：此字段可能返回 null，表示取不到有效值。",
			"config_version_count":       "配置版本数量。注意：此字段可能返回 null，表示取不到有效值。",
			"result_output_file":         "用于保存结果。",
		},
	})

}

func dataSourceTencentCloudTsfConfigSummary() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides detailed information of tsf config_summary",
		Read:        dataSourceTencentCloudTsfConfigSummaryRead,
		Schema: map[string]*schema.Schema{
			"application_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Application ID. If not passed, the query will be for all.",
			},

			"search_word": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Query keyword, fuzzy query: application name, configuration item name. If not passed, the query will be for all.",
			},

			"order_by": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Order term. support Sort by time: creation_time; or Sort by name: config_name.",
			},

			"order_type": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Pass 0 for ascending order and 1 for descending order.",
			},

			"config_tag_list": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Config tag list.",
			},

			"disable_program_auth_check": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether to disable dataset authentication.",
			},

			"config_id_list": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Config Id List.",
			},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Config Page Item.",
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
							Description: "Config list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Configuration item ID.Note: This field may return null, indicating that no valid value was found.",
									},
									"config_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Configuration name.Note: This field may return null, indicating that no valid value was found.",
									},
									"config_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Configuration version. Note: This field may return null, indicating that no valid value was found.",
									},
									"config_version_desc": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Configuration version description.Note: This field may return null, indicating that no valid value was found.",
									},
									"config_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Configuration value.Note: This field may return null, indicating that no valid value was found.",
									},
									"config_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Config type. Note: This field may return null, indicating that no valid value was found.",
									},
									"creation_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Create time.Note: This field may return null, indicating that no valid value was found.",
									},
									"application_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Application ID.Note: This field may return null, indicating that no valid value was found.",
									},
									"application_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Application Name. Note: This field may return null, indicating that no valid value was found.",
									},
									"delete_flag": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Deletion flag, true: deletable; false: not deletable.Note: This field may return null, indicating that no valid value was found.",
									},
									"last_update_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Last update time.Note: This field may return null, indicating that no valid value was found.",
									},
									"config_version_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Configure version count.Note: This field may return null, indicating that no valid value was found.",
									},
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

func dataSourceTencentCloudTsfConfigSummaryRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tsf_config_summary.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("application_id"); ok {
		paramMap["ApplicationId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("search_word"); ok {
		paramMap["SearchWord"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("order_by"); ok {
		paramMap["OrderBy"] = helper.String(v.(string))
	}

	if v, _ := d.GetOk("order_type"); v != nil {
		paramMap["OrderType"] = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("config_tag_list"); ok {
		configTagListSet := v.(*schema.Set).List()
		paramMap["ConfigTagList"] = helper.InterfacesStringsPoint(configTagListSet)
	}

	if v, _ := d.GetOk("disable_program_auth_check"); v != nil {
		paramMap["DisableProgramAuthCheck"] = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("config_id_list"); ok {
		configIdListSet := v.(*schema.Set).List()
		paramMap["ConfigIdList"] = helper.InterfacesStringsPoint(configIdListSet)
	}

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	var config *tsf.TsfPageConfig
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTsfConfigSummaryByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		config = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(config.Content))
	tsfPageConfigMap := map[string]interface{}{}
	if config != nil {
		if config.TotalCount != nil {
			tsfPageConfigMap["total_count"] = config.TotalCount
		}

		if config.Content != nil {
			contentList := []interface{}{}
			for _, content := range config.Content {
				contentMap := map[string]interface{}{}

				if content.ConfigId != nil {
					contentMap["config_id"] = content.ConfigId
				}

				if content.ConfigName != nil {
					contentMap["config_name"] = content.ConfigName
				}

				if content.ConfigVersion != nil {
					contentMap["config_version"] = content.ConfigVersion
				}

				if content.ConfigVersionDesc != nil {
					contentMap["config_version_desc"] = content.ConfigVersionDesc
				}

				if content.ConfigValue != nil {
					contentMap["config_value"] = content.ConfigValue
				}

				if content.ConfigType != nil {
					contentMap["config_type"] = content.ConfigType
				}

				if content.CreationTime != nil {
					contentMap["creation_time"] = content.CreationTime
				}

				if content.ApplicationId != nil {
					contentMap["application_id"] = content.ApplicationId
				}

				if content.ApplicationName != nil {
					contentMap["application_name"] = content.ApplicationName
				}

				if content.DeleteFlag != nil {
					contentMap["delete_flag"] = content.DeleteFlag
				}

				if content.LastUpdateTime != nil {
					contentMap["last_update_time"] = content.LastUpdateTime
				}

				if content.ConfigVersionCount != nil {
					contentMap["config_version_count"] = content.ConfigVersionCount
				}

				contentList = append(contentList, contentMap)
				ids = append(ids, *content.ConfigId)
			}

			tsfPageConfigMap["content"] = contentList
		}

		_ = d.Set("result", []interface{}{tsfPageConfigMap})
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tsfPageConfigMap); e != nil {
			return e
		}
	}
	return nil
}
