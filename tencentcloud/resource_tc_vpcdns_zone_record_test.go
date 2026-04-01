package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcDnsZoneRecord_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcDnsZoneRecord_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone_record.record", "weight", "1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpcdns_zone_record.record",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcDnsZoneRecord_basic = defaultInstanceVariable + `
resource "tencentcloudenterprise_vpcdns_zone" "zone" {
  dns_forward_status = "DISABLED"
  domain             = "domain.com"
  remark             = "test_record"
  tags = {
    "created-by" : "terraform",
  }
}

resource "tencentcloudenterprise_vpcdns_zone_record" "record" {
  mx           = 0
  record_type  = "A"
  record_value = "192.168.1.2"
  sub_domain   = "www"
  ttl          = 300
  weight       = 1
  zone_id      = tencentcloudenterprise_vpcdns_zone.zone.id
}
`
