package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcDnsForwardRule_basic(t *testing.T) {
	t.Parallel()

	domain := fmt.Sprintf("tf-forward-%s.com", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccVpcDnsForwardRule_basic, domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpcdns_forward_rule.rule", "rule_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_forward_rule.rule", "remark", "forward_rule_foo"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpcdns_forward_rule.rule", "forward_address.#", "2"),
				),
			},
		},
	})
}

const testAccVpcDnsForwardRule_basic = `
resource "tencentcloudenterprise_vpcdns_zone" "zone" {
  dns_forward_status = "ENABLED"
  domain             = "%s"
  remark             = "test_forward_rule"
  tags = {
    "created-by" = "terraform"
  }
}

resource "tencentcloudenterprise_vpcdns_forward_rule" "rule" {
  remark          = "forward_rule_foo"
  zone_id         = tencentcloudenterprise_vpcdns_zone.zone.id
  forward_address = ["8.8.8.8:53", "1.1.1.1:53"]
}
`
