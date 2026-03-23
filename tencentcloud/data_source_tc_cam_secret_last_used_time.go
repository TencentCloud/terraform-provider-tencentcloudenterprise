/*
使用Account SDK接口查询密钥最后使用时间。

Use Account SDK API to query secret last used time.

# Example Usage

```hcl

	data "tencentcloudenterprise_cam_secret_last_used_time" "example" {
	  secret_id_list = ["AKIDxxxxxxxxxxxx"]
	}

```
*/
package tencentcloud

import (
	account "terraform-provider-tencentcloudenterprise/sdk/account/v20190325"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cam_secret_last_used_time", CNDescription{
		TerraformTypeCN: "密钥最后使用时间",
		DescriptionCN:   "用于查询密钥最后使用时间。",
		AttributesCN: map[string]string{
			"secret_id_list":           "查询的密钥ID列表",
			"result_output_file":       "用于保存结果",
			"secret_id_last_used_rows": "最后使用时间列表",
			"secret_id":                "密钥ID",
			"last_used_date":           "最后使用日期（具有1天延迟）",
			"last_secret_used_date":    "最后使用时间戳",
		},
	})
}

func dataSourceTencentCloudCamSecretLastUsedTime() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamSecretLastUsedTimeRead,
		Schema: map[string]*schema.Schema{
			"secret_id_list": {
				Required:  true,
				Sensitive: true,
				Type:      schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Query the key ID list. Supports up to 10.",
			},

			"secret_id_last_used_rows": {
				Computed:    true,
				Sensitive:   true,
				Type:        schema.TypeList,
				Description: "Last used time list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secret_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secret Id.",
						},
						"last_used_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last used date (with 1 day delay).",
						},
						"last_secret_used_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last used timestamp.",
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

func dataSourceTencentCloudCamSecretLastUsedTimeRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cam_secret_last_used_time.read")()

	var secretIdList []*string
	if v, ok := d.GetOk("secret_id_list"); ok {
		secretIdListSet := v.(*schema.Set).List()
		secretIdList = helper.InterfacesStringsPoint(secretIdListSet)
	}

	client := meta.(*TencentCloudClient).apiV3Conn

	var secretIdLastUsedRows []*account.SecretIDLastUsedRow

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		request := account.NewGetSecurityLastUsedRequest()
		request.SecretIdList = secretIdList
		response, e := client.UseAccountClient().GetSecurityLastUsed(request)
		if e != nil {
			return retryError(e)
		}
		secretIdLastUsedRows = response.Response.SecretIdLastUsedRows
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(secretIdLastUsedRows))
	tmpList := make([]map[string]interface{}, 0, len(secretIdLastUsedRows))

	if secretIdLastUsedRows != nil {
		for _, secretIdLastUsed := range secretIdLastUsedRows {
			secretIdLastUsedMap := map[string]interface{}{}

			if secretIdLastUsed.SecretId != nil {
				secretIdLastUsedMap["secret_id"] = secretIdLastUsed.SecretId
			}

			if secretIdLastUsed.LastUsedDate != nil {
				secretIdLastUsedMap["last_used_date"] = secretIdLastUsed.LastUsedDate
			}

			if secretIdLastUsed.LastSecretUsedDate != nil {
				secretIdLastUsedMap["last_secret_used_date"] = *secretIdLastUsed.LastSecretUsedDate
			}

			ids = append(ids, *secretIdLastUsed.SecretId)
			tmpList = append(tmpList, secretIdLastUsedMap)
		}

		_ = d.Set("secret_id_last_used_rows", tmpList)
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
