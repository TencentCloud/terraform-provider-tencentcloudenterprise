package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbListenerDefaultDomain_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListenerDefaultDomain_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_default_domain.example", "clb_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_default_domain.example", "listener_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_default_domain.example", "domain", "www.example.com"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_clb_listener_default_domain.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccClbListenerDefaultDomain_basic = `
resource "tencentcloudenterprise_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-listener-default-domain-test"
}

resource "tencentcloudenterprise_clb_listener" "listener_basic" {
  clb_id        = tencentcloudenterprise_clb_instance.clb_basic.id
  port          = 80
  protocol      = "HTTP"
  listener_name = "listener_basic"
}

resource "tencentcloudenterprise_clb_listener_rule" "rule_basic" {
  clb_id      = tencentcloudenterprise_clb_instance.clb_basic.id
  listener_id = tencentcloudenterprise_clb_listener.listener_basic.listener_id
  domain      = "www.example.com"
  url         = "/"
}

resource "tencentcloudenterprise_clb_listener_default_domain" "example" {
  clb_id      = tencentcloudenterprise_clb_instance.clb_basic.id
  listener_id = tencentcloudenterprise_clb_listener.listener_basic.listener_id
  domain      = "www.example.com"
  depends_on  = [tencentcloudenterprise_clb_listener_rule.rule_basic]
}
`
