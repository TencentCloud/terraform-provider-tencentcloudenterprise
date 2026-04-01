package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcDnsZone_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcDnsZone_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone.zone", "domain", "domain.com"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpcdns_zone.zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcDnsZone_basic = defaultInstanceVariable + `
resource "tencentcloudenterprise_vpcdns_zone" "zone" {
  dns_forward_status = "DISABLED"
  domain             = "domain.com"
  remark             = "test_zone"
  vpc_set {
    region      = "ap-guangzhou"
    uniq_vpc_id = var.cvm_vpc_id
  }
  vpc_set {
    region      = "ap-guangzhou"
    uniq_vpc_id = var.vpc_id
  }
  tags = {
    "created-by" : "terraform",
  }
}
`
