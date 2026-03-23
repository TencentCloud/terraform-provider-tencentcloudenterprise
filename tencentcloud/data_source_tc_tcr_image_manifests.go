/*
Use this data source to query detailed information of tcr image_manifests

Example Usage

```hcl
data "tencentcloudenterprise_tcr_image_manifests" "image_manifests" {
	registry_id = "%s"
	namespace_name = "%s"
	repository_name = "%s"
	image_version = "v1"
}
```
*/
package tencentcloud

import (
	"context"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tcr_image_manifests", CNDescription{
		TerraformTypeCN: "TCR镜像清单",
		DescriptionCN:   "用于查询TCR镜像清单的详细信息",
		AttributesCN: map[string]string{
			"registry_id":        "实例ID",
			"namespace_name":     "命名空间名称",
			"repository_name":    "镜像仓库名称",
			"image_version":      "镜像版本",
			"manifest":           "镜像的清单信息",
			"config":             "镜像的配置信息",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudTcrImageManifests() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudTcrImageManifestsRead,
		Description: "Use this data source to query detailed information of TCR image manifests",
		Schema: map[string]*schema.Schema{
			"registry_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Instance ID.",
			},

			"namespace_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Namespace name.",
			},

			"repository_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Mirror warehouse name.",
			},

			"image_version": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Mirror version.",
			},

			"manifest": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Manifest information of the image.",
			},

			"config": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Configuration information of the image.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudTcrImageManifestsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_tcr_image_manifests.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var (
		registryId     string
		namespaceName  string
		repositoryName string
		imageVersion   string
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
		repositoryName = v.(string)
	}

	if v, ok := d.GetOk("image_version"); ok {
		paramMap["image_version"] = helper.String(v.(string))
		imageVersion = v.(string)
	}

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		config   *string
		manifest *string
	)

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTcrImageManifestsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		config = result.Config
		manifest = result.Manifest
		return nil
	})
	if err != nil {
		return err
	}

	if manifest != nil {
		_ = d.Set("manifest", manifest)
	}

	if config != nil {
		_ = d.Set("config", config)
	}

	tmpList := []map[string]interface{}{
		{
			"manifest": manifest,
			"config":   config,
		},
	}

	d.SetId(helper.DataResourceIdsHash([]string{registryId, namespaceName, repositoryName, imageVersion}))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}
