package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcDnsZoneRecord_basic(t *testing.T) {
	t.Parallel()
	domain := fmt.Sprintf("tf-acc-vpcdns-record-%s.com", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccVpcDnsZoneRecord_basic, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone_record.record", "weight", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone_record.record", "remark", "tf-acc-remark"),
				),
			},
			{
				Config: fmt.Sprintf(testAccVpcDnsZoneRecord_update, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone_record.record", "weight", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_zone_record.record", "remark", "tf-acc-remark-updated"),
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

const testAccVpcDnsZoneRecord_basic = `
resource "tencentcloudenterprise_vpcdns_zone" "zone" {
  dns_forward_status = "DISABLED"
  domain             = "%s"
  remark             = "test_zone"
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
  remark       = "tf-acc-remark"
  zone_id      = tencentcloudenterprise_vpcdns_zone.zone.id
}
`

const testAccVpcDnsZoneRecord_update = `
resource "tencentcloudenterprise_vpcdns_zone" "zone" {
  dns_forward_status = "DISABLED"
  domain             = "%s"
  remark             = "test_zone"
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
  remark       = "tf-acc-remark-updated"
  zone_id      = tencentcloudenterprise_vpcdns_zone.zone.id
}
`
