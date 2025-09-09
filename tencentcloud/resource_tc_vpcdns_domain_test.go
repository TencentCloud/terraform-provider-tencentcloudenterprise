package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcDnsDomainResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcDnsDomain,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloud_vpc_traffic_package.traffic_package", "id")),
			},
			{
				ResourceName:      "cloud_vpc_traffic_package.traffic_package",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcDnsDomain = `

resource "cloud_vpc_traffic_package" "traffic_package" {
  traffic_amount = 10
}
`
