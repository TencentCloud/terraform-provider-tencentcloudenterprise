/*
Provide a resource to enable/disable TKE cluster encryption protection.

# Example Usage

```hcl
resource "tencentcloudenterprise_tke_kubernetes_encryption_protection" "example" {
  cluster_id = "cls-xxxxxxxx"

  kms_configuration {
    key_id     = "kms-xxxxxxxx"
    kms_region = "ap-guangzhou"
  }
}
```

# Import

TKE encryption protection can be imported using the cluster_id, e.g.
```
$ terraform import tencentcloudenterprise_tke_kubernetes_encryption_protection.example cls-xxxxxxxx
```
*/
package tencentcloud

import (
	"context"
	"log"

	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_encryption_protection", CNDescription{
		TerraformTypeCN: "集群加密保护",
		DescriptionCN:   "用于开启或关闭 TKE 集群加密保护。",
		AttributesCN: map[string]string{
			"cluster_id":        "集群 ID。",
			"kms_configuration": "KMS 加密配置。",
			"key_id":            "KMS 密钥 ID。",
			"kms_region":        "KMS 地域。",
		},
	})
}

func resourceTencentCloudKubernetesEncryptionProtection() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to enable/disable TKE cluster encryption protection.",
		Create:      resourceTencentCloudKubernetesEncryptionProtectionCreate,
		Read:        resourceTencentCloudKubernetesEncryptionProtectionRead,
		Delete:      resourceTencentCloudKubernetesEncryptionProtectionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the cluster.",
			},
			"kms_configuration": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "KMS encryption configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "KMS key ID.",
						},
						"kms_region": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "KMS region.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudKubernetesEncryptionProtectionCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_encryption_protection.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	_ = context.WithValue(context.TODO(), logIdKey, logId)

	client := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient()
	request := tke.NewEnableEncryptionProtectionRequest()

	clusterId := d.Get("cluster_id").(string)
	request.ClusterId = helper.String(clusterId)

	if v, ok := d.GetOk("kms_configuration"); ok {
		kmsConfigList := v.([]interface{})
		if len(kmsConfigList) > 0 {
			kmsConfigMap := kmsConfigList[0].(map[string]interface{})
			kmsConfiguration := tke.KMSConfiguration{}
			if v, ok := kmsConfigMap["key_id"]; ok {
				kmsConfiguration.KeyId = helper.String(v.(string))
			}
			if v, ok := kmsConfigMap["kms_region"]; ok {
				kmsConfiguration.KmsRegion = helper.String(v.(string))
			}
			request.KMSConfiguration = &kmsConfiguration
		}
	}

	log.Printf("[DEBUG]%s api[EnableEncryptionProtection] request: %s", logId, request.ToJsonString())

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := client.EnableEncryptionProtection(request)
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s enable encryption protection failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(clusterId)

	return resourceTencentCloudKubernetesEncryptionProtectionRead(d, meta)
}

func resourceTencentCloudKubernetesEncryptionProtectionRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_encryption_protection.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clusterId := d.Id()

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, has, err := service.DescribeCluster(ctx, clusterId)
	if err != nil {
		return err
	}
	if !has {
		d.SetId("")
		return nil
	}

	_ = d.Set("cluster_id", clusterId)

	return nil
}

func resourceTencentCloudKubernetesEncryptionProtectionDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_encryption_protection.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	_ = context.WithValue(context.TODO(), logIdKey, logId)

	client := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient()
	request := tke.NewDisableEncryptionProtectionRequest()

	clusterId := d.Id()
	request.ClusterId = helper.String(clusterId)

	log.Printf("[DEBUG]%s api[DisableEncryptionProtection] request: %s", logId, request.ToJsonString())

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := client.DisableEncryptionProtection(request)
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s disable encryption protection failed, reason:%+v", logId, err)
		return err
	}

	return nil
}
