package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpnDnsDomainDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudVpnDnsDomainDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_vpcdns_domains.foo"),
					resource.TestCheckResourceAttr("data.cloud_vpcdns_domains.foo", "list.#", "1"),
					resource.TestCheckResourceAttr("data.cloud_vpcdns_domains.foo",
						"list.0.domain", "brucelin.cc"),
				),
			},
		},
	})
}

const testAccTencentCloudVpnDnsDomainDataSourceConfig_basic = `
data "tencentcloudenterprise_vpcdns_domains" "foo" {
  result_output_file = "foo.json"
  domain      = "brucezylin.cc"
}
`
