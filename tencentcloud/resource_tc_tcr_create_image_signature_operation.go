/*
Provides a resource to create a tcr image_signature_operation

Example Usage

```hcl
resource "tencentcloudenterprise_tcr_create_image_signature_operation" "image_signature_operation" {
  registry_id = "tcr-xxx"
  namespace_name = "ns"
  repository_name = "repo"
  image_version = "v1"

}
```

Import

tcr image_signature_operation can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_tcr_create_image_signature_operation.image_signature_operation image_signature_operation_id
```
*/
package tencentcloud

import (
	"log"
	"strings"

	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//func init() {
//	registerResourceDescriptionProvider("tencentcloudenterprise_tcr_create_image_signature_operation", CNDescription{
//		TerraformTypeCN: "TCR创建镜像签名操作",
//		DescriptionCN:   "提供TCR创建镜像签名操作资源，用于为镜像创建数字签名",
//		AttributesCN: map[string]string{
//			"registry_id":     "实例ID",
//			"namespace_name":  "命名空间名称",
//			"repository_name": "仓库名称",
//			"image_version":   "镜像版本",
//		},
//	})
//}

func resourceTencentCloudTcrCreateImageSignatureOperation() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTcrCreateImageSignatureOperationCreate,
		Read:        resourceTencentCloudTcrCreateImageSignatureOperationRead,
		Delete:      resourceTencentCloudTcrCreateImageSignatureOperationDelete,
		Description: "Provides a resource to create TCR image signature operation",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"registry_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance id.",
			},

			"namespace_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Namespace name.",
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
		},
	}
}

func resourceTencentCloudTcrCreateImageSignatureOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_create_image_signature_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request        = tcr.NewCreateSignatureRequest()
		registryId     string
		namespaceName  string
		repositoryName string
		imageVersion   string
	)
	if v, ok := d.GetOk("registry_id"); ok {
		request.RegistryId = helper.String(v.(string))
		registryId = v.(string)
	}

	if v, ok := d.GetOk("namespace_name"); ok {
		request.NamespaceName = helper.String(v.(string))
		namespaceName = v.(string)
	}

	if v, ok := d.GetOk("repository_name"); ok {
		request.RepositoryName = helper.String(v.(string))
		repositoryName = v.(string)
	}

	if v, ok := d.GetOk("image_version"); ok {
		request.ImageVersion = helper.String(v.(string))
		imageVersion = v.(string)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTCRClient().CreateSignature(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate tcr ImageSignatureOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{registryId, namespaceName, repositoryName, imageVersion}, FILED_SP))

	return resourceTencentCloudTcrCreateImageSignatureOperationRead(d, meta)
}

func resourceTencentCloudTcrCreateImageSignatureOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_create_image_signature_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudTcrCreateImageSignatureOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_create_image_signature_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
