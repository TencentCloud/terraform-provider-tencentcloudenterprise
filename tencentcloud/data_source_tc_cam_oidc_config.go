/*
Use this data source to query detailed information of CAM OIDC SSO configuration.

# Example Usage

```hcl

	data "tencentcloudenterprise_cam_oidc_config" "example" {
	  name = "example-oidc-idp"
	}

```
*/
package tencentcloud

import (
	"encoding/json"
	"fmt"
	"log"

	open "terraform-provider-tencentcloudenterprise/sdk/open/v20201202"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cam_oidc_config", CNDescription{
		TerraformTypeCN: "CAM OIDC SSO配置",
		DescriptionCN:   "用于查询CAM OIDC SSO配置的详细信息。",
		AttributesCN: map[string]string{
			"name":          "OIDC身份提供商名称",
			"provider_type": "身份提供商类型，11表示角色身份提供商",
			"identity_url":  "身份提供商URL",
			"identity_key":  "签名公钥",
			"client_id":     "客户端ID列表",
			"status":        "状态。0: 未设置; 2: 已禁用; 11: 已启用",
			"description":   "描述信息",
		},
	})
}

func DataSourceTencentCloudCamOidcConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamOidcConfigRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Name.",
			},

			"provider_type": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "IdP type. 11: Role IdP.",
			},

			"identity_url": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "IdP URL.",
			},

			"identity_key": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Public key for signature.",
			},

			"client_id": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Client ID.",
			},

			"status": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Status. 0: Not set; 2: Disabled; 11: Enabled.",
			},

			"description": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Description.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCamOidcConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cam_oidc_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	var name string
	result := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	// Call ListIdentityProvider to get all IdP configs
	request := open.NewListIdentityProviderRequest()
	var response *open.ListIdentityProviderResponse

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOpenClient().ListIdentityProvider(request)
		if e != nil {
			return retryError(e)
		}
		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s read CAM OIDC config failed, reason:%s\n", logId, err.Error())
		return err
	}

	if response == nil || response.Response == nil || response.Response.Item == nil || response.Response.Item.Data == nil {
		return fmt.Errorf("ListIdentityProvider returns nil response")
	}

	// Find the OIDC config with matching name
	var found bool
	for _, item := range response.Response.Item.Data.List {
		if item == nil {
			continue
		}

		// Check if this is an OIDC provider
		if item.Oidc == nil || *item.Oidc == "" {
			continue
		}

		// Check if the name matches
		if item.Name != nil && *item.Name == name {
			// Parse OIDC JSON string
			var oidcData map[string]interface{}
			if err := json.Unmarshal([]byte(*item.Oidc), &oidcData); err != nil {
				log.Printf("[WARN]%s failed to parse OIDC JSON for item %v: %s\n", logId, item.Id, err.Error())
				continue
			}

			found = true

			// Set provider_type and status from Item
			if item.ProviderType != nil {
				_ = d.Set("provider_type", item.ProviderType)
				result["provider_type"] = item.ProviderType
			}

			if item.Status != nil {
				_ = d.Set("status", item.Status)
				result["status"] = item.Status
			}

			if item.Desc != nil {
				_ = d.Set("description", item.Desc)
				result["description"] = item.Desc
			}

			// Extract OIDC-specific fields from JSON
			if val, ok := oidcData["IdentityUrl"].(string); ok {
				_ = d.Set("identity_url", val)
				result["identity_url"] = val
			}

			if val, ok := oidcData["IdentityKey"].(string); ok {
				_ = d.Set("identity_key", val)
				result["identity_key"] = val
			}

			if val, ok := oidcData["ClientId"].(string); ok {
				clientIds := []string{val}
				_ = d.Set("client_id", clientIds)
				result["client_id"] = clientIds
			}

			break
		}
	}

	if !found {
		return fmt.Errorf("OIDC config with name %s not found", name)
	}

	d.SetId(name)
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), result); e != nil {
			return e
		}
	}
	return nil
}
