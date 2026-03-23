/*
Use this data source to query detailed information of tcr images

Example Usage

```hcl
data "tencentcloudenterprise_tcr_images" "images" {
  registry_id = "tcr-xxx"
  namespace_name = "ns"
  repository_name = "repo"
  image_version = "v1"
  digest = "sha256:xxxxx"
  exact_match = false
  }
```
*/
package tencentcloud

import (
	"context"
	"strings"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tcr_images", CNDescription{
		TerraformTypeCN: "TCR镜像",
		DescriptionCN:   "用于查询TCR镜像的详细信息",
		AttributesCN: map[string]string{
			"registry_id":        "实例ID",
			"namespace_name":     "命名空间名称",
			"repository_name":    "仓库名称",
			"image_version":      "镜像版本名称，默认为模糊匹配",
			"digest":             "指定镜像摘要进行查找",
			"exact_match":        "指定是否为精确匹配，true为精确匹配，不填为模糊匹配",
			"image_info_list":    "容器镜像信息列表",
			"size":               "镜像大小",
			"update_time":        "更新时间",
			"kind":               "镜像类型",
			"kms_signature":      "KMS签名信息",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudTcrImages() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudTcrImagesRead,
		Description: "Use this data source to query detailed information of TCR images",
		Schema: map[string]*schema.Schema{
			"registry_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Instance id.",
			},

			"namespace_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Namespace name.",
			},

			"repository_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Repository name.",
			},

			"image_version": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Image version name, default is fuzzy match.",
			},

			"digest": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify image digest for lookup.",
			},

			"exact_match": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Specifies whether it is an exact match, true is an exact match, and not filled is a fuzzy match.",
			},

			"image_info_list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Container image information list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"digest": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hash value.",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Image size (unit: byte).",
						},
						"image_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tag name.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time.",
						},
						"kind": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Product type,note: this field may return null, indicating that no valid value can be obtained.",
						},
						"kms_signature": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Kms signature information,note: this field may return null, indicating that no valid value can be obtained.",
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

func dataSourceTencentCloudTcrImagesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_tcr_images.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId         = getLogId(contextNil)
		ctx           = context.WithValue(context.TODO(), logIdKey, logId)
		registryId    string
		namespaceName string
		repoName      string
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("registry_id"); ok {
		paramMap["registry_id"] = helper.String(v.(string))
		registryId = v.(string)
	}

	if v, ok := d.GetOk("namespace_name"); ok {
		paramMap["namespace_name"] = helper.String(v.(string))
		namespaceName = v.(string)
	}

	if v, ok := d.GetOk("repository_name"); ok {
		paramMap["repository_name"] = helper.String(v.(string))
		repoName = v.(string)
	}

	if v, ok := d.GetOk("image_version"); ok {
		paramMap["image_version"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("digest"); ok {
		paramMap["digest"] = helper.String(v.(string))
	}

	if v, _ := d.GetOk("exact_match"); v != nil {
		paramMap["exact_match"] = helper.Bool(v.(bool))
	}

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}

	var imageInfoList []*tcr.TcrImageInfo

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTcrImagesByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		imageInfoList = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(imageInfoList))
	tmpList := make([]map[string]interface{}, 0, len(imageInfoList))

	if imageInfoList != nil {
		for _, tcrImageInfo := range imageInfoList {
			tcrImageInfoMap := map[string]interface{}{}

			if tcrImageInfo.Digest != nil {
				tcrImageInfoMap["digest"] = tcrImageInfo.Digest
			}

			if tcrImageInfo.Size != nil {
				tcrImageInfoMap["size"] = tcrImageInfo.Size
			}

			if tcrImageInfo.ImageVersion != nil {
				tcrImageInfoMap["image_version"] = tcrImageInfo.ImageVersion
			}

			if tcrImageInfo.UpdateTime != nil {
				tcrImageInfoMap["update_time"] = tcrImageInfo.UpdateTime
			}

			if tcrImageInfo.Kind != nil {
				tcrImageInfoMap["kind"] = tcrImageInfo.Kind
			}

			if tcrImageInfo.KmsSignature != nil {
				tcrImageInfoMap["kms_signature"] = tcrImageInfo.KmsSignature
			}

			ids = append(ids, strings.Join([]string{registryId, namespaceName, repoName, *tcrImageInfo.ImageVersion}, FILED_SP))
			tmpList = append(tmpList, tcrImageInfoMap)
		}

		_ = d.Set("image_info_list", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}
