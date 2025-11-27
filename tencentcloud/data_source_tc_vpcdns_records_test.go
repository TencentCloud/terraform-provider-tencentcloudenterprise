package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpnDnsRecordDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudVpnDnsRecordDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_vpcdns_records.records"),
					resource.TestCheckResourceAttr("data.cloud_vpcdns_records.records", "list.#", "2"),
					//resource.TestCheckResourceAttr("data.cloud_vpcdns_domains.records",
					//"list.0.domain", "brucelin.cc"),
				),
			},
		},
	})
}

const testAccTencentCloudVpnDnsRecordDataSourceConfig_basic = `
data "tencentcloudenterprise_vpcdns_records" "records" {
  domain_id = 745
  result_output_file = "records.json"
}
`
