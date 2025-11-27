package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudBmsFlavorsDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudBmsFlavorsDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_bms_flavors.flavors"),
					//resource.TestCheckResourceAttr("data.cloud_bms_flavors.flavors", "list.#", "2"),
					//resource.TestCheckResourceAttr("data.cloud_vpcdns_domains.records",
					//"list.0.domain", "brucelin.cc"),
				),
			},
		},
	})
}

const testAccTencentCloudBmsFlavorsDataSourceConfig_basic = `
data "tencentcloudenterprise_bms_flavors" "flavors" {
  result_output_file = "flavors.json"
}
`
