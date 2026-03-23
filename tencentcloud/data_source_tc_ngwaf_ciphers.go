package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func dataSourceTencentCloudNgwafCiphers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudNgwafCiphersRead,
		Schema: map[string]*schema.Schema{
			"ciphers": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Encryption Suite InformationNote: This field may return null, indicating that a valid value cannot be obtained.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TLS version IDNote: This field may return null, indicating that a valid value cannot be obtained.",
						},
						"cipher_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Encryption Suite IDNote: This field may return null, indicating that a valid value cannot be obtained.",
						},
						"cipher_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Encryption Suite NameNote: This field may return null, indicating that a valid value cannot be obtained.",
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

func dataSourceTencentCloudNgwafCiphersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_ngwaf_ciphers.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
		ciphers []*ngwaf.TLSCiphers
	)

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeWafCiphersByFilter(ctx)
		if e != nil {
			return retryError(e)
		}

		ciphers = result
		return nil
	})

	if err != nil {
		return err
	}

	ids := make([]string, 0, len(ciphers))
	tmpList := make([]map[string]interface{}, 0, len(ciphers))

	if ciphers != nil {
		for _, tLSCiphers := range ciphers {
			tLSCiphersMap := map[string]interface{}{}

			if tLSCiphers.VersionId != nil {
				tLSCiphersMap["version_id"] = tLSCiphers.VersionId
			}

			if tLSCiphers.CipherId != nil {
				tLSCiphersMap["cipher_id"] = tLSCiphers.CipherId
			}

			if tLSCiphers.CipherName != nil {
				tLSCiphersMap["cipher_name"] = tLSCiphers.CipherName
			}

			ids = append(ids, *tLSCiphers.CipherName)
			tmpList = append(tmpList, tLSCiphersMap)
		}

		_ = d.Set("ciphers", tmpList)
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
