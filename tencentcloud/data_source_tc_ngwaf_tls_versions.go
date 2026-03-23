package tencentcloud

import (
	"context"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTencentCloudNgwafTlsVersions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudNgwafTlsVersionsRead,
		Schema: map[string]*schema.Schema{
			"tls": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "TLS key value.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TLS version IDNote: This field may return null, indicating that a valid value cannot be obtained.",
						},
						"version_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tls version nameNote: This field may return null, indicating that a valid value cannot be obtained.",
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

func dataSourceTencentCloudNgwafTlsVersionsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_ngwaf_tls_versions.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
		tLS     []*ngwaf.TLSVersion
	)

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeWafTlsVersionsByFilter(ctx)
		if e != nil {
			return retryError(e)
		}

		tLS = result
		return nil
	})

	if err != nil {
		return err
	}

	ids := make([]string, 0, len(tLS))
	tmpList := make([]map[string]interface{}, 0, len(tLS))

	if tLS != nil {
		for _, tLSVersion := range tLS {
			tLSVersionMap := map[string]interface{}{}

			if tLSVersion.VersionId != nil {
				tLSVersionMap["version_id"] = tLSVersion.VersionId
			}

		if tLSVersion.VersionName != nil {
			tLSVersionMap["version_name"] = tLSVersion.VersionName
			ids = append(ids, *tLSVersion.VersionName)
		}

		tmpList = append(tmpList, tLSVersionMap)
		}

		_ = d.Set("tls", tmpList)
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
