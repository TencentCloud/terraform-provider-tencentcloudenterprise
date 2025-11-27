/*
Use this data source to query detailed information of tsf repository

# Example Usage

```hcl

	data "tencentcloudenterprise_tsf_repository" "repository" {
	  search_word = "test"
	  repository_type = "default"
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
	registerDataDescriptionProvider("tencentcloudenterprise_tsf_repository", CNDescription{
		TerraformTypeCN: "TSF仓库",
		DescriptionCN:   "提供TSF仓库数据源，用于查询TSF仓库的详细信息。",
		AttributesCN: map[string]string{
			"search_word":        "查询关键字（按仓库名搜索）",
			"repository_type":    "仓库类型（默认仓库：default，私有仓库：private）",
			"result":             "符合查询条件的仓库信息列表",
			"total_count":        "仓库总数",
			"content":            "仓库信息列表。注意：此字段可能返回 null，表示取不到有效值。",
			"repository_id":      "仓库Id",
			"repository_name":    "仓库名称",
			"repository_desc":    "仓库描述（默认仓库：default，私有仓库：private）",
			"is_used":            "仓库是否正在被使用。注意：此字段可能返回 null，表示取不到有效值。",
			"create_time":        "创建时间。注意：此字段可能返回 null，表示取不到有效值。",
			"bucket_name":        "仓库bucket名称。注意：此字段可能返回 null，表示取不到有效值。",
			"bucket_region":      "仓库区域。注意：此字段可能返回 null，表示取不到有效值。",
			"directory":          "仓库目录。注意：此字段可能返回 null，表示取不到有效值。",
			"result_output_file": "用于保存结果",
		},
	})

}

func dataSourceTencentCloudTsfRepository() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tsf repository.",
		Read:        dataSourceTencentCloudTsfRepositoryRead,
		Schema: map[string]*schema.Schema{
			"search_word": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Query keywords (search by Repository name).",
			},

			"repository_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Repository type (default Repository: default, private Repository: private).",
			},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "A list of Repository information that meets the query criteria.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total Repository.",
						},
						"content": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Repository information list. Note: This field may return null, indicating that no valid value can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"repository_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Repository Id.",
									},
									"repository_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Repository Name.",
									},
									"repository_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Repository type (default Repository: default, private Repository: private).",
									},
									"repository_desc": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Repository description (default warehouse: default, private warehouse: private).",
									},
									"is_used": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the repository is being used. Note: This field may return null, indicating that no valid value can be obtained.",
									},
									"create_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "CreationTime. Note: This field may return null, indicating that no valid values can be obtained.",
									},
									"bucket_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Repository bucket name. Note: This field may return null, indicating that no valid value can be obtained.",
									},
									"bucket_region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Repository region. Note: This field may return null, indicating that no valid value can be obtained.",
									},
									"directory": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Repository Directory. Note: This field may return null, indicating that no valid value can be obtained.",
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

func dataSourceTencentCloudTsfRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tsf_repository.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("search_word"); ok {
		paramMap["SearchWord"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("repository_type"); ok {
		paramMap["RepositoryType"] = helper.String(v.(string))
	}

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	var repository *tsf.RepositoryList
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTsfRepositoryByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		repository = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(repository.Content))
	repositoryListMap := map[string]interface{}{}
	if repository != nil {
		if repository.TotalCount != nil {
			repositoryListMap["total_count"] = repository.TotalCount
		}

		if repository.Content != nil {
			contentList := []interface{}{}
			for _, content := range repository.Content {
				contentMap := map[string]interface{}{}

				if content.RepositoryId != nil {
					contentMap["repository_id"] = content.RepositoryId
				}

				if content.RepositoryName != nil {
					contentMap["repository_name"] = content.RepositoryName
				}

				if content.RepositoryType != nil {
					contentMap["repository_type"] = content.RepositoryType
				}

				if content.RepositoryDesc != nil {
					contentMap["repository_desc"] = content.RepositoryDesc
				}

				if content.IsUsed != nil {
					contentMap["is_used"] = content.IsUsed
				}

				if content.CreateTime != nil {
					contentMap["create_time"] = content.CreateTime
				}

				if content.BucketName != nil {
					contentMap["bucket_name"] = content.BucketName
				}

				if content.BucketRegion != nil {
					contentMap["bucket_region"] = content.BucketRegion
				}

				if content.Directory != nil {
					contentMap["directory"] = content.Directory
				}

				contentList = append(contentList, contentMap)
				ids = append(ids, *content.RepositoryId)
			}

			repositoryListMap["content"] = contentList
		}

		_ = d.Set("result", []interface{}{repositoryListMap})
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), repositoryListMap); e != nil {
			return e
		}
	}
	return nil
}
