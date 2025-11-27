package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudTsfDeliveryConfigsDataSource_basic -v
func TestAccTencentCloudTsfDeliveryConfigsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfDeliveryConfigsDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tsf_delivery_configs.delivery_configs"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.0.config_id"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.0.config_name"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.0.create_time"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.0.enable_auth"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.0.kafka_infos.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.0.kafka_infos.0.path.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.0.kafka_infos.0.topic"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.0.enable_global_line_rule"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_delivery_configs.delivery_configs", "result.0.content.0.line_rule"),
				),
			},
		},
	})
}

const testAccTsfDeliveryConfigsDataSource = `

data "tencentcloudenterprise_tsf_delivery_configs" "delivery_configs" {
  search_word = "test"
}

`
