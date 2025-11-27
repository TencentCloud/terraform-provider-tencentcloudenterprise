/*
Use this data source to query detailed information of SSM secret
Example Usage
```hcl

	data "tencentcloudenterprise_ssm_secrets" "foo" {
	  secret_name = "test"
	  order_type = 1
	  state = 1
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ssm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm/v20190923"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_ssm_secrets", CNDescription{
		TerraformTypeCN: "SSM凭据",
		DescriptionCN:   "用于查询SSM凭据的详细信息",
		AttributesCN: map[string]string{
			"order_type":         "凭据创建时间的排序方式，0-降序，1-升序，默认为0",
			"state":              "凭据状态过滤器",
			"search_secret_name": "根据凭据名称进行过滤",
			"tag_filters":        "标签过滤器",
			"secret_type":        "凭据类型过滤器",
			"result_output_file": "用于保存结果",
			"secret_list":        "凭据列表",
			"secret_name":        "凭据名称",
			"description":        "凭据描述",
			"kms_key_id":         "KMS密钥ID",
			"create_uin":         "创建者UIN",
			"status":             "凭据状态",
			"delete_time":        "删除时间",
			"create_time":        "创建时间",
			"kms_key_type":       "KMS密钥类型",
		},
	})
}

func dataSourceTencentCloudSsmSecrets() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudSsmSecretsRead,
		Description: "Use this data source to query detailed information of SSM secrets",
		Schema: map[string]*schema.Schema{
			"order_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The order to sort the create time of secret. `0` - desc, `1` - asc. Default value is `0`.",
			},
			"state": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Filter by state of secret. `0` - all secrets are queried, `1` - only Enabled secrets are queried, `2` - only Disabled secrets are queried, `3` - only PendingDelete secrets are queried.",
			},
			"secret_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret name used to filter result.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tags to filter secret.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"secret_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of SSM secrets.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secret_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of secret.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of secret.",
						},
						"kms_key_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "KMS keyId used to encrypt secret.",
						},
						"create_uin": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Uin of Creator.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of secret.",
						},
						"delete_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete time of CMK.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Create time of secret.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudSsmSecretsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ssm_secrets.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	param := make(map[string]interface{})
	if v, ok := d.GetOk("order_type"); ok {
		param["order_type"] = v.(int)
	}
	if v, ok := d.GetOk("state"); ok {
		param["state"] = v.(int)
	}
	if v, ok := d.GetOk("secret_name"); ok {
		param["secret_name"] = v.(string)
	}
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		param["tag_filter"] = tags
	}

	ssmService := SsmService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var secrets []*ssm.SecretMetadata
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := ssmService.DescribeSecretsByFilter(ctx, param)
		if e != nil {
			return retryError(e)
		}
		secrets = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read SSM secrets failed, reason:%+v", logId, err)
		return err
	}
	secretList := make([]map[string]interface{}, 0, len(secrets))
	secretNames := make([]string, 0, len(secrets))
	for _, secret := range secrets {
		mapping := map[string]interface{}{
			"secret_name": secret.SecretName,
			"description": secret.Description,
			"kms_key_id":  secret.KmsKeyId,
			"create_uin":  secret.CreateUin,
			"status":      secret.Status,
			"delete_time": secret.DeleteTime,
			"create_time": secret.CreateTime,
		}

		secretList = append(secretList, mapping)
		secretNames = append(secretNames, *secret.SecretName)
	}

	d.SetId(helper.DataResourceIdsHash(secretNames))
	if e := d.Set("secret_list", secretList); e != nil {
		log.Printf("[CRITAL]%s provider set SSM secret list fail, reason:%+v", logId, e)
		return e
	}
	if output, ok := d.GetOk("result_output_file"); ok && output.(string) != "" {
		return writeToFile(output.(string), secretList)
	}
	return nil
}
