/*
Use this data source to query detailed information of CAM SAML providers

# Example Usage

```hcl

	data "tencentcloudenterprise_cam_saml_providers" "foo" {
	  name = "cam-test-provider"
	}

```
*/
package tencentcloud

import (
	"fmt"
	"log"
	"strings"

	open "terraform-provider-tencentcloudenterprise/sdk/open/v20201202"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cam_saml_providers", CNDescription{
		TerraformTypeCN: "CAM SAML提供商列表",
		DescriptionCN:   "用于查询 CAM SAML 提供商信息的列表。",
		AttributesCN: map[string]string{
			"name":          "提供商名称",
			"description":   "描述",
			"provider_list": "提供商列表",
			"create_time":   "创建时间",
			"modify_time":   "最后修改时间",
		},
	})
}

func dataSourceTencentCloudCamSAMLProviders() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamSAMLProvidersRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the CAM SAML provider to be queried.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the CAM SAML provider.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"provider_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of CAM SAML providers. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of CAM SAML provider.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of CAM SAML provider.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of the CAM SAML provider.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last modify time of the CAM SAML provider.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCamSAMLProvidersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cam_saml_providers.read")()

	logId := getLogId(contextNil)

	nameFilter := ""
	descFilter := ""

	if v, ok := d.GetOk("name"); ok {
		nameFilter = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		descFilter = v.(string)
	}

	// Since Open SDK doesn't have ListSAMLProviders API, we can only get a single provider by name
	// If name filter is not provided, we cannot list all providers
	if nameFilter == "" {
		return fmt.Errorf("name parameter is required for querying SAML providers")
	}

	client := meta.(*TencentCloudClient).apiV3Conn
	cpf := client.NewClientProfileTce(300)
	openClient, _ := open.NewClient(client.CredentialTce, client.Region, cpf)
	openClient.WithHttpTransport(&connectivity.LogRoundTripper{})

	var providerList []map[string]interface{}
	var ids []string

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		request := open.NewGetSAMLProviderRequest()
		request.Name = &nameFilter

		response, e := openClient.GetSAMLProvider(request)
		if e != nil {
			// If provider not found, return empty list
			if strings.Contains(e.Error(), "ResourceNotFound") || strings.Contains(e.Error(), "NotFound") {
				log.Printf("[INFO]%s SAML provider %s not found\n", logId, nameFilter)
				return nil
			}
			log.Printf("[CRITAL]%s read CAM SAML provider failed, reason:%s\n", logId, e.Error())
			return retryError(e)
		}

		if response.Response == nil || response.Response.Name == nil {
			return nil
		}

		// Check description filter if provided
		if descFilter != "" && response.Response.Description != nil && *response.Response.Description != descFilter {
			// Description doesn't match, return empty list
			return nil
		}

		mapping := map[string]interface{}{
			"name":        *response.Response.Name,
			"create_time": *response.Response.CreateTime,
			"modify_time": *response.Response.ModifyTime,
		}
		if response.Response.Description != nil {
			mapping["description"] = *response.Response.Description
		}

		providerList = append(providerList, mapping)
		ids = append(ids, *response.Response.Name)
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s read CAM SAML providers failed, reason:%s\n", logId, err.Error())
		return err
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("provider_list", providerList); e != nil {
		log.Printf("[CRITAL]%s provider set provider list fail, reason:%s\n", logId, e.Error())
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), providerList); e != nil {
			return e
		}
	}

	return nil
}
