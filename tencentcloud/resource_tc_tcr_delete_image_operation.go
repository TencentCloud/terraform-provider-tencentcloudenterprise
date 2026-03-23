/*
Provides a resource to create a tcr delete_image_operation

Example Usage

```hcl
resource "tencentcloudenterprise_tcr_delete_image_operation" "delete_image_operation" {
  registry_id = "tcr-xxx"
  repository_name = "repo"
  image_version = "v1"
  namespace_name = "ns"
}
```

*/
package tencentcloud

import (
	"log"
	"strings"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tcr_delete_image_operation", CNDescription{
		TerraformTypeCN: "TCR删除镜像操作",
		DescriptionCN:   "提供TCR删除镜像操作资源，用于删除容器镜像仓库中的镜像",
		AttributesCN: map[string]string{
			"registry_id":     "实例ID",
			"namespace_name":  "命名空间名称",
			"repository_name": "仓库名称",
			"image_version":   "镜像版本",
		},
	})
}

func resourceTencentCloudTcrDeleteImageOperation() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTcrDeleteImageOperationCreate,
		Read:        resourceTencentCloudTcrDeleteImageOperationRead,
		Delete:      resourceTencentCloudTcrDeleteImageOperationDelete,
		Description: "Provides a resource to create TCR delete image operation",
		Schema: map[string]*schema.Schema{
			"registry_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance id.",
			},

			"repository_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Repository name.",
			},

			"image_version": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Image version name.",
			},

			"namespace_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Namespace name.",
			},
		},
	}
}

func resourceTencentCloudTcrDeleteImageOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_delete_image_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request        = tcr.NewDeleteImageRequest()
		registryId     string
		namespaceName  string
		repositoryName string
		imageVersion   string
	)
	if v, ok := d.GetOk("registry_id"); ok {
		request.RegistryId = helper.String(v.(string))
		registryId = v.(string)
	}

	if v, ok := d.GetOk("repository_name"); ok {
		request.RepositoryName = helper.String(v.(string))
		repositoryName = v.(string)
	}

	if v, ok := d.GetOk("image_version"); ok {
		request.ImageVersion = helper.String(v.(string))
		imageVersion = v.(string)
	}

	if v, ok := d.GetOk("namespace_name"); ok {
		request.NamespaceName = helper.String(v.(string))
		namespaceName = v.(string)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTCRClient().DeleteImage(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate tcr DeleteImageOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{registryId, namespaceName, repositoryName, imageVersion}, FILED_SP))

	return resourceTencentCloudTcrDeleteImageOperationRead(d, meta)
}

func resourceTencentCloudTcrDeleteImageOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_delete_image_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudTcrDeleteImageOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_delete_image_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
