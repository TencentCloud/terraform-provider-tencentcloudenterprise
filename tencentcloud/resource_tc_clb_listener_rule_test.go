package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudClbListenerRuleResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListenerRule_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerRuleExists("tencentcloudenterprise_clb_listener_rule.rule_basic"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_rule.rule_basic", "clb_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_rule.rule_basic", "listener_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "domain", "abc.com"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "session_expire_time", "30"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "url", "/"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "target_type", "TARGETGROUP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "forward_type", "HTTPS"),
				),
			},
			{
				Config: testAccClbListenerRule__basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerRuleExists("tencentcloudenterprise_clb_listener_rule.rule_basic"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_rule.rule_basic", "clb_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_rule.rule_basic", "listener_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "domain", "abc.com"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "session_expire_time", "30"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "url", "/"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_basic", "forward_type", "HTTP"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_clb_listener_rule.rule_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudClbListenerRuleResource_full(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccClbListenerRule_full, defaultSshCertificate, defaultSshCertificateB),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerRuleExists("tencentcloudenterprise_clb_listener_rule.rule_full"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_rule.rule_full", "clb_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_rule.rule_full", "listener_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "domain", "abc.com"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "session_expire_time", "30"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "url", "/"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_interval_time", "200"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_health_num", "3"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_unhealth_num", "3"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_http_method", "GET"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_http_code", "31"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_http_domain", "abc.com"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_http_path", "/"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "certificate_ssl_mode", "UNIDIRECTIONAL"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "certificate_id", defaultSshCertificateB),
				),
			}, {
				Config: fmt.Sprintf(testAccClbListenerRule_update, defaultSshCertificate, defaultSshCertificateB),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerRuleExists("tencentcloudenterprise_clb_listener_rule.rule_full"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_rule.rule_full", "clb_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_listener_rule.rule_full", "listener_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "domain", "abcd.com"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "session_expire_time", "60"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "url", "/"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_interval_time", "300"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_health_num", "6"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_unhealth_num", "6"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_http_method", "HEAD"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_http_code", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_http_domain", "abcd.com"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "health_check_http_path", "/"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "certificate_ssl_mode", "UNIDIRECTIONAL"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_listener_rule.rule_full", "certificate_id", defaultSshCertificateB),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_clb_listener_rule.rule_full",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckClbListenerRuleDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clbService := ClbService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_clb_listener_rule" {
			continue
		}
		resourceId := rs.Primary.ID
		items := strings.Split(resourceId, FILED_SP)
		itemLength := len(items)
		locationId := items[itemLength-1]
		listenerId := rs.Primary.Attributes["listener_id"]
		clbId := rs.Primary.Attributes["clb_id"]
		//this function is not supported by api, need to be travelled
		filter := map[string]string{"rule_id": locationId, "listener_id": listenerId, "clb_id": clbId}
		rules, err := clbService.DescribeRulesByFilter(ctx, filter)
		if len(rules) > 0 && err == nil {
			return fmt.Errorf("[CHECK][CLB listener rule][Destroy] check: CLB listener rule still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckClbListenerRuleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CLB listener rule][Exists] check: CLB listener rule %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CLB listener rule][Exists] check: CLB listener rule id is not set")
		}
		clbService := ClbService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		resourceId := rs.Primary.ID
		items := strings.Split(resourceId, FILED_SP)
		itemLength := len(items)
		locationId := items[itemLength-1]
		listenerId := rs.Primary.Attributes["listener_id"]
		clbId := rs.Primary.Attributes["clb_id"]
		filter := map[string]string{"rule_id": locationId, "listener_id": listenerId, "clb_id": clbId}
		rules, err := clbService.DescribeRulesByFilter(ctx, filter)
		if err != nil {
			return err
		}
		if len(rules) == 0 {
			return fmt.Errorf("[CHECK][CLB listener rule][Exists] id %s is not exist", rs.Primary.ID)
		}
		return nil
	}
}

const testAccClbListenerRule_basic = `
resource "tencentcloudenterprise_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-rule-basic"
}

resource "tencentcloudenterprise_clb_listener" "listener_basic" {
  clb_id        = cloud_clb_instance.clb_basic.id
  port          = 1
  protocol      = "HTTP"
  listener_name = "listener_basic"
}

resource "tencentcloudenterprise_clb_listener_rule" "rule_basic" {
  clb_id              = cloud_clb_instance.clb_basic.id
  listener_id         = cloud_clb_listener.listener_basic.listener_id
  domain              = "abc.com"
  url                 = "/"
  session_expire_time = 30
  scheduler           = "WRR"
  target_type         = "TARGETGROUP"
  forward_type        = "HTTPS"
}
`

const testAccClbListenerRule__basic_update = `
resource "tencentcloudenterprise_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-rule-basic"
}

resource "tencentcloudenterprise_clb_listener" "listener_basic" {
  clb_id        = cloud_clb_instance.clb_basic.id
  port          = 1
  protocol      = "HTTP"
  listener_name = "listener_basic"
}

resource "tencentcloudenterprise_clb_listener_rule" "rule_basic" {
  clb_id              = cloud_clb_instance.clb_basic.id
  listener_id         = cloud_clb_listener.listener_basic.listener_id
  domain              = "abc.com"
  url                 = "/"
  session_expire_time = 30
  scheduler           = "WRR"
  forward_type        = "HTTP"
}
`

const testAccClbListenerRule_full = `
resource "tencentcloudenterprise_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-rule-full"
}

resource "tencentcloudenterprise_clb_listener" "listener_basic" {
  clb_id               = cloud_clb_instance.clb_basic.id
  listener_name        = "listener_https"
  port                 = 77
  protocol             = "HTTPS"
  certificate_ssl_mode = "UNIDIRECTIONAL"
  certificate_id       = "%s"
  sni_switch = true
}

resource "tencentcloudenterprise_clb_listener_rule" "rule_full" {
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_id                = cloud_clb_listener.listener_basic.listener_id
  domain                     = "abc.com"
  url                        = "/"
  session_expire_time        = 30
  scheduler                  = "WRR"
  health_check_switch        = true
  health_check_interval_time = 200
  health_check_health_num    = 3
  health_check_unhealth_num  = 3
  health_check_http_path     = "/"
  health_check_http_domain   = "abc.com"
  health_check_http_code     = "31"
  health_check_http_method   = "GET"
  certificate_ssl_mode       = "UNIDIRECTIONAL"
  certificate_id             = "%s"
}
`

const testAccClbListenerRule_update = `
resource "tencentcloudenterprise_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-rule-full"
}

resource "tencentcloudenterprise_clb_listener" "listener_basic" {
  clb_id               = cloud_clb_instance.clb_basic.id
  listener_name        = "listener_https"
  port                 = 77
  protocol             = "HTTPS"
  certificate_ssl_mode = "UNIDIRECTIONAL"
  certificate_id       = "%s"
  sni_switch = true
}

resource "tencentcloudenterprise_clb_listener_rule" "rule_full" {
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_id                = cloud_clb_listener.listener_basic.listener_id
  domain                     = "abcd.com"
  url                        = "/"
  session_expire_time        = 60
  scheduler                  = "WRR"
  health_check_switch        = true
  health_check_interval_time = 300
  health_check_health_num    = 6
  health_check_unhealth_num  = 6
  health_check_time_out      = 4
  health_check_type          = "TCP"
  health_check_http_path     = "/"
  health_check_http_domain   = "abcd.com"
  health_check_http_code     = "1"
  health_check_http_method   = "HEAD"
  certificate_ssl_mode       = "UNIDIRECTIONAL"
  certificate_id             = "%s"
}
`
