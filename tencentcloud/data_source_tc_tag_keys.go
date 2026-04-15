/*
Use this data source to query tag keys

# Example Usage

```hcl

	data "tencentcloudenterprise_tag_keys" "example" {
	}

```
*/
package tencentcloud

import (
	"log"

	tag "terraform-provider-tencentcloudenterprise/sdk/tag/v20180813"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tag_keys", CNDescription{
		TerraformTypeCN: "标签键列表",
		DescriptionCN:   "查询标签键列表。",
		AttributesCN: map[string]string{
			"tags":               "标签键列表。",
			"result_output_file": "用于保存结果。",
		},
	})
}

func dataSourceTencentCloudTagKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudTagKeysRead,
		Schema: map[string]*schema.Schema{
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Tag key list.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func dataSourceTencentCloudTagKeysRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_tag_keys.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := tag.NewDescribeTagKeysRequest()

	var (
		offset uint64 = 0
		limit  int64  = 1000
	)

	result := make([]*string, 0)
	for {
		request.Offset = &offset
		request.Limit = &limit
		ratelimit.Check(request.GetAction())
		response, err := meta.(*TencentCloudClient).apiV3Conn.UseTagClient().DescribeTagKeys(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), err.Error())
			return err
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		if response == nil || response.Response == nil || len(response.Response.Tags) < 1 {
			break
		}

		result = append(result, response.Response.Tags...)
		if len(response.Response.Tags) < int(limit) {
			break
		}

		offset += uint64(limit)
	}

	if result != nil {
		_ = d.Set("tags", result)
	}

	d.SetId(helper.BuildToken())
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), d); e != nil {
			return e
		}
	}

	return nil
}
