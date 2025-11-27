package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbRedirectionsDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbRedirectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbRedirectionsDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckClbRedirectionExists("tencentcloudenterprise_clb_redirection.redirection_basic"),
					resource.TestCheckResourceAttr("data.cloud_clb_redirections.redirections", "redirection_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_redirections.redirections", "redirection_list.0.clb_id"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_redirections.redirections", "redirection_list.0.source_listener_id"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_redirections.redirections", "redirection_list.0.target_listener_id"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_redirections.redirections", "redirection_list.0.source_rule_id"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_redirections.redirections", "redirection_list.0.target_rule_id"),
				),
			},
		},
	})
}

const testAccClbRedirectionsDataSource = `
resource "tencentcloudenterprise_clb_instance" "clb" {
  network_type = "OPEN"
  clb_name     = "tf-clb-redirections"
}

resource "tencentcloudenterprise_clb_listener" "listener_basic" {
  clb_id        = cloud_clb_instance.clb.id
  port          = 1
  protocol      = "HTTP"
  listener_name = "listener_basic"
}

resource "tencentcloudenterprise_clb_listener_rule" "rule_basic" {
  clb_id              = cloud_clb_instance.clb.id
  listener_id         = cloud_clb_listener.listener_basic.listener_id
  domain              = "abc.com"
  url                 = "/"
  session_expire_time = 30
  scheduler           = "WRR"
}

resource "tencentcloudenterprise_clb_listener" "listener_target" {
  clb_id        = cloud_clb_instance.clb.id
  port          = 44
  protocol      = "HTTP"
  listener_name = "listener_basic1"
}

resource "tencentcloudenterprise_clb_listener_rule" "rule_target" {
  clb_id              = cloud_clb_instance.clb.id
  listener_id         = cloud_clb_listener.listener_target.listener_id
  domain              = "abcd.com"
  url                 = "/"
  session_expire_time = 30
  scheduler           = "WRR"
}

resource "tencentcloudenterprise_clb_redirection" "redirection_basic" {
  clb_id             = cloud_clb_instance.clb.id
  source_listener_id = cloud_clb_listener.listener_basic.listener_id
  target_listener_id = cloud_clb_listener.listener_target.listener_id
  source_rule_id     = cloud_clb_listener_rule.rule_basic.rule_id
  target_rule_id     = cloud_clb_listener_rule.rule_target.rule_id
}

data "tencentcloudenterprise_clb_redirections" "redirections" {
  clb_id             = cloud_clb_instance.clb.id
  source_listener_id = cloud_clb_redirection.redirection_basic.source_listener_id
  source_rule_id     = cloud_clb_redirection.redirection_basic.source_rule_id
}
`
