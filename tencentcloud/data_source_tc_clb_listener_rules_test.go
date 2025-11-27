package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbListenerRulesDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListenerRulesDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckClbListenerRuleExists("tencentcloudenterprise_clb_listener_rule.rule"),
					resource.TestCheckResourceAttr("data.cloud_clb_listener_rules.rules", "rule_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_listener_rules.rules", "rule_list.0.clb_id"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_listener_rules.rules", "rule_list.0.listener_id"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_listener_rules.rules", "rule_list.0.rule_id"),
					resource.TestCheckResourceAttr("data.cloud_clb_listener_rules.rules", "rule_list.0.session_expire_time", "30"),
					resource.TestCheckResourceAttr("data.cloud_clb_listener_rules.rules", "rule_list.0.scheduler", "WRR"),
				),
			},
		},
	})
}

const testAccClbListenerRulesDataSource = `
resource "tencentcloudenterprise_clb_instance" "clb" {
  network_type = "OPEN"
  clb_name     = "tf-clb-rules"
}

resource "tencentcloudenterprise_clb_listener" "listener" {
  clb_id        = cloud_clb_instance.clb.id
  port          = 1
  protocol      = "HTTP"
  listener_name = "mylistener1234"
}

resource "tencentcloudenterprise_clb_listener_rule" "rule" {
  clb_id              = cloud_clb_instance.clb.id
  listener_id         = cloud_clb_listener.listener.listener_id
  domain              = "abcde.com"
  url                 = "/"
  session_expire_time = 30
  scheduler           = "WRR"
}

data "tencentcloudenterprise_clb_listener_rules" "rules" {
  clb_id      = cloud_clb_instance.clb.id
  listener_id = cloud_clb_listener.listener.listener_id
  domain      = cloud_clb_listener_rule.rule.domain
  url         = cloud_clb_listener_rule.rule.url
}
`
