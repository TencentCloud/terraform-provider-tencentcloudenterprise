package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcDnsZone_basic(t *testing.T) {
	t.Parallel()
	domain := fmt.Sprintf("tf-acc-vpcdns-zone-%s.com", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccVpcDnsZone_basic, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone.zone", "domain", domain),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpcdns_zone.zone", "id"),
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

func TestAccTencentCloudVpcDnsZone_concurrentSameSuffix(t *testing.T) {
	t.Parallel()
	suffix := acctest.RandString(8)
	domainA := fmt.Sprintf("tf-acc-vpcdns-a-%s.com", suffix)
	domainB := fmt.Sprintf("tf-acc-vpcdns-b-%s.com", suffix)
	domainC := fmt.Sprintf("tf-acc-vpcdns-c-%s.com", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccVpcDnsZone_concurrentSameSuffix, domainA, domainB, domainC),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone.zone_a", "domain", domainA),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone.zone_b", "domain", domainB),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone.zone_c", "domain", domainC),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpcdns_zone.zone_a", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpcdns_zone.zone_b", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpcdns_zone.zone_c", "id"),
				),
			},
		},
	})
}

const testAccVpcDnsZone_basic = `
resource "tencentcloudenterprise_vpcdns_zone" "zone" {
  dns_forward_status = "DISABLED"
  domain             = "%s"
  remark             = "test_zone"
  tags = {
    "created-by" : "terraform",
  }
}
`

const testAccVpcDnsZone_concurrentSameSuffix = `
resource "tencentcloudenterprise_vpcdns_zone" "zone_a" {
  dns_forward_status = "DISABLED"
  domain             = "%s"
  remark             = "concurrent-a"
}

resource "tencentcloudenterprise_vpcdns_zone" "zone_b" {
  dns_forward_status = "DISABLED"
  domain             = "%s"
  remark             = "concurrent-b"
}

resource "tencentcloudenterprise_vpcdns_zone" "zone_c" {
  dns_forward_status = "DISABLED"
  domain             = "%s"
  remark             = "concurrent-c"
}
`
