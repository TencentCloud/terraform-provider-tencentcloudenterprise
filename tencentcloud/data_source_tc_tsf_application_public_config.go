/*
Use this data source to query detailed information of tsf application_public_config

# Example Usage

```hcl

	data "tencentcloudenterprise_tsf_application_public_config" "application_public_config" {
	  config_id = "dcfg-p-evjrbgly"
	  # config_id_list = [""]
	  config_name = "dsadsa"
	  config_version = "123"
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
	registerDataDescriptionProvider("tencentcloudenterprise_tsf_application_public_config", CNDescription{
		TerraformTypeCN: "TSF应用公共配置",
		DescriptionCN:   "提供TSF应用公共配置数据源，用于查询TSF应用公共配置的详细信息。",
		AttributesCN: map[string]string{
			"config_id":            "配置ID",
			"config_name":          "配置名称",
			"config_version":       "配置版本",
			"config_version_desc":  "配置版本描述",
			"config_value":         "配置值",
			"config_type":          "配置类型",
			"creation_time":        "创建时间",
			"application_id":       "应用ID",
			"application_name":     "应用名称",
			"delete_flag":          "删除标识，true:允许删除；false:禁止删除。",
			"last_update_time":     "最后更新时间",
			"config_version_count": "配置版本数量",
			"result":               "分页全局配置列表。注意：此字段可能返回 null，表示取不到有效值。",
			"result_output_file":   "用于保存结果",
			"config_id_list":       "配置ID列表。不传则查询所有项，优先级低。",
		},
	})

}

func dataSourceTencentCloudTsfApplicationPublicConfig() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides detailed information of tsf application_public_config",
		Read:        dataSourceTencentCloudTsfApplicationPublicConfigRead,
		Schema: map[string]*schema.Schema{
			"config_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Config ID. Query all items if not passed, high priority.",
			},

			"config_id_list": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Config ID list. Query all items if not passed, low priority.",
			},

			"config_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Config name. Exact query. Query all items if not passed.",
			},

			"config_version": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Config version. Exact query. Query all items if not passed.",
			},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Paginated global configuration  list. Note: This field may return null, indicating that no valid value can be obtained.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TsfPageConfig.",
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
										Description: "Config ID. Note: This field may return null, indicating that no valid value can be obtained.",
									},
									"config_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Config name. Note: This field may return null, indicating that no valid value can be obtained.",
									},
									"config_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Config version. Note: This field may return null, indicating that no valid value can be obtained.",
									},
									"config_version_desc": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Config version description. Note: This field may return null, indicating that no valid value can be obtained.",
									},
									"config_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Config value. Note: This field may return null, indicating that no valid value can be obtained.",
									},
									"config_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Config type. Note: This field may return null, indicating that no valid value can be obtained.",
									},
									"creation_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "CreationTime. Note: This field may return null, indicating that no valid values can be obtained.",
									},
									"application_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Application Id. Note: This field may return null, indicating that no valid values can be obtained.",
									},
									"application_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Application Id. Note: This field may return null, indicating that no valid values can be obtained.",
									},
									"delete_flag": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Delete flag, true: allow delete; false: delete prohibit.",
									},
									"last_update_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Last update time.  Note: This field may return null, indicating that no valid values can be obtained.",
									},
									"config_version_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Config version count.  Note: This field may return null, indicating that no valid values can be obtained.",
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

func dataSourceTencentCloudTsfApplicationPublicConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tsf_application_public_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("config_id"); ok {
		paramMap["ConfigId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_id_list"); ok {
		configIdListSet := v.(*schema.Set).List()
		paramMap["ConfigIdList"] = helper.InterfacesStringsPoint(configIdListSet)
	}

	if v, ok := d.GetOk("config_name"); ok {
		paramMap["ConfigName"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_version"); ok {
		paramMap["ConfigVersion"] = helper.String(v.(string))
	}

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	var config *tsf.TsfPageConfig
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTsfApplicationPublicConfigByFilter(ctx, paramMap)
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
