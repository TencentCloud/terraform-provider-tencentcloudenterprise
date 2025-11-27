package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudTsfGroupGatewaysDataSource_basic -v
func TestAccTencentCloudTsfGroupGatewaysDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfGroupGatewaysDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tsf_group_gateways.group_gateways"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.total_count"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.acl_mode"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.api_count"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.auth_type"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.created_time"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.description"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.gateway_instance_id"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.gateway_instance_type"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.group_context"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.group_type"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.namespace_name_key_position"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.service_name_key_position"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.status"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_group_gateways.group_gateways", "result.0.content.0.updated_time"),
				),
			},
		},
	})
}

const testAccTsfGroupGatewaysDataSource = `

data "tencentcloudenterprise_tsf_group_gateways" "group_gateways" {
	gateway_deploy_group_id = "group-aeoej4qy"
	search_word = "test"
}

`
