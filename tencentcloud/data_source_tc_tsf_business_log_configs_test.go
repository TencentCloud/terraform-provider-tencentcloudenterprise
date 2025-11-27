package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudTsfBusinessLogConfigsDataSource_basic -v
func TestAccTencentCloudTsfBusinessLogConfigsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfBusinessLogConfigsDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tsf_business_log_configs.business_log_configs"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.total_count"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_create_time"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_id"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_name"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_path"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_pipeline"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_schema.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_schema.0.schema_create_time"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_schema.0.schema_date_format"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_schema.0.schema_multiline_pattern"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_tags"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_business_log_configs.business_log_configs", "result.0.content.0.config_update_time"),

					testAccCheckTencentCloudDataSourceID("data.cloud_tsf_business_log_configs.business_log_configs")),
			},
		},
	})
}

const testAccTsfBusinessLogConfigsDataSource = `

data "tencentcloudenterprise_tsf_business_log_configs" "business_log_configs" {
  config_id_list = ["apm-busi-log-cfg-qv3x3rdv"]
}

`
