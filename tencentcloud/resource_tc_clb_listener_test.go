package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudClbListener_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListener_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_basic"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_basic", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_basic", "protocol", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_basic", "listener_name", "listener_basic"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_basic", "session_expire_time", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_basic", "port", "1"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_basic", "scheduler", "WRR"),
				),
			},
			{
				ResourceName:      "cloud_clb_listener.listener_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudClbListener_tcp_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListener_tcp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcp"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcp", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "protocol", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "listener_name", "listener_tcp"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "session_expire_time", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_time_out", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_interval_time", "100"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_health_num", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_unhealth_num", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "target_type", "NODE"),
				),
			},
			{
				Config: testAccClbListener_tcp_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcp"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcp", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "protocol", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "listener_name", "listener_tcp_update"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "session_expire_time", "60"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_time_out", "20"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_interval_time", "200"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_health_num", "3"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_unhealth_num", "3"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "target_type", "TARGETGROUP"),
				),
			},
			{
				ResourceName:      "cloud_clb_listener.listener_tcp",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudClbListenerTCPWithTCP(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListener_tcp_tcp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcp"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcp", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "protocol", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "listener_name", "listener_tcp"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "session_expire_time", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_time_out", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_interval_time", "100"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_health_num", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_unhealth_num", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_type", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_port", "200"),
				),
			},
			{
				Config: testAccClbListener_tcp_update_tcp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcp"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcp", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "protocol", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "listener_name", "listener_tcp_update"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "session_expire_time", "60"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_time_out", "20"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_interval_time", "200"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_health_num", "3"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_unhealth_num", "3"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_type", "TCP"),
				),
			},
			{
				ResourceName:      "cloud_clb_listener.listener_tcp",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudClbListenerTCPWithHTTP(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListener_tcp_http,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcp"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcp", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "protocol", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "listener_name", "listener_tcp"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "session_expire_time", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_time_out", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_interval_time", "100"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_health_num", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_unhealth_num", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_type", "HTTP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_code", "16"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_path", "/"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_domain", "www.tencent.com"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_method", "HEAD"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_version", "HTTP/1.1"),
				),
			},
			{
				Config: testAccClbListener_tcp_update_http,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcp"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcp", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "protocol", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "listener_name", "listener_tcp_update"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "session_expire_time", "60"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_time_out", "20"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_interval_time", "200"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_health_num", "3"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_unhealth_num", "3"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_type", "HTTP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_port", "200"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_code", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_path", ""),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_domain", ""),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_method", "GET"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_http_version", "HTTP/1.0"),
				),
			},
			{
				ResourceName:      "cloud_clb_listener.listener_tcp",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudClbListenerTCPWithCustomer(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListener_tcp_customer,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcp"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcp", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "protocol", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "listener_name", "listener_tcp"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "session_expire_time", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_time_out", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_interval_time", "100"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_health_num", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_unhealth_num", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_type", "CUSTOM"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_context_type", "HEX"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_send_context", "0123456789ABCDEF"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_recv_context", "ABCD"),
				),
			},
			{
				Config: testAccClbListener_tcp_customer_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcp"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcp", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "protocol", "TCP"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "listener_name", "listener_tcp_update"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "session_expire_time", "60"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_time_out", "20"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_interval_time", "200"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_health_num", "3"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_unhealth_num", "3"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_type", "CUSTOM"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_context_type", "TEXT"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_send_context", "/get/test"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcp", "health_check_recv_context", "http_1xx"),
				),
			},
			{
				ResourceName:      "cloud_clb_listener.listener_tcp",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudClbListener_https(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccClbListener_https, defaultSshCertificate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_https"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_https", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "protocol", "HTTPS"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "listener_name", "listener_https"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "port", "77"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "certificate_ssl_mode", "UNIDIRECTIONAL"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "certificate_id", defaultSshCertificate),
				),
			},
			{
				Config: fmt.Sprintf(testAccClbListener_https_update, defaultSshCertificateB),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_https"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_https", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "protocol", "HTTPS"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "listener_name", "listener_https_update"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "port", "33"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "certificate_ssl_mode", "UNIDIRECTIONAL"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_https", "certificate_id", defaultSshCertificateB),
				),
			},
			{
				ResourceName:      "cloud_clb_listener.listener_https",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudClbListener_tcpssl(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccClbListener_tcpssl, defaultSshCertificate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcpssl"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcpssl", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "protocol", "TCP_SSL"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "listener_name", "listener_tcpssl"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "certificate_ssl_mode", "UNIDIRECTIONAL"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "certificate_id", defaultSshCertificate),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_time_out", "30"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_interval_time", "100"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_health_num", "2"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_unhealth_num", "2"),
				),
			},
			{
				Config: fmt.Sprintf(testAccClbListener_tcpssl_update, defaultSshCertificateB),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener_tcpssl"),
					resource.TestCheckResourceAttrSet("cloud_clb_listener.listener_tcpssl", "clb_id"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "protocol", "TCP_SSL"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "listener_name", "listener_tcpssl_update"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "certificate_ssl_mode", "UNIDIRECTIONAL"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "certificate_id", defaultSshCertificateB),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "port", "44"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "scheduler", "WRR"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_switch", "true"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_time_out", "20"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_interval_time", "200"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_health_num", "3"),
					resource.TestCheckResourceAttr("cloud_clb_listener.listener_tcpssl", "health_check_unhealth_num", "3"),
				),
			},
			{
				ResourceName:      "cloud_clb_listener.listener_tcpssl",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckClbListenerDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clbService := ClbService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloud_clb_listener" {
			continue
		}
		time.Sleep(5 * time.Second)
		resourceId := rs.Primary.ID
		items := strings.Split(resourceId, FILED_SP)
		itemLength := len(items)
		listenerId := items[itemLength-1]
		clbId := rs.Primary.Attributes["clb_id"]
		if itemLength == 2 && clbId != "" {
			clbId = items[0]
		}
		instance, err := clbService.DescribeListenerById(ctx, listenerId, clbId)
		if instance != nil && err == nil {
			return fmt.Errorf("[CHECK][CLB listener][Destroy] check: CLB listener still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckClbListenerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CLB listener][Exists] check: CLB listener %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CLB listener][Exists] check: CLB listener id is not set")
		}
		clbService := ClbService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		resourceId := rs.Primary.ID
		items := strings.Split(resourceId, FILED_SP)
		itemLength := len(items)
		listenerId := items[itemLength-1]
		clbId := rs.Primary.Attributes["clb_id"]
		if itemLength == 2 && clbId != "" {
			clbId = items[0]
		}
		instance, err := clbService.DescribeListenerById(ctx, listenerId, clbId)
		if err != nil {
			return err
		}
		if instance == nil {
			return fmt.Errorf("[CHECK][CLB listener][Exists] id %s is not exist", rs.Primary.ID)
		}
		return nil
	}
}

const testAccClbListener_basic = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-listener-basic"
}

resource "cloud_clb_listener" "listener_basic" {
  clb_id              = cloud_clb_instance.clb_basic.id
  port                = 1
  protocol            = "TCP"
  listener_name       = "listener_basic"
  session_expire_time = 30
  scheduler           = "WRR"
  target_type         = "TARGETGROUP"
}
`

const testAccClbListener_tcp = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-listener-tcp"
}

resource "cloud_clb_listener" "listener_tcp" {
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcp"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 30
  health_check_interval_time = 100
  health_check_health_num    = 2
  health_check_unhealth_num  = 2
  session_expire_time        = 30
  scheduler                  = "WRR"
  target_type         = "NODE"
}
`

const testAccClbListener_tcp_update = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-listener-tcp"
}

resource "cloud_clb_listener" "listener_tcp"{
  clb_id = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcp_update"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 20
  health_check_interval_time = 200
  health_check_health_num    = 3
  health_check_unhealth_num  = 3
  session_expire_time        = 60
  scheduler                  = "WRR"
  target_type         = "TARGETGROUP"
}
`

const testAccClbListener_tcpssl = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-tcpssl"
}

resource "cloud_clb_listener" "listener_tcpssl" {
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcpssl"
  port                       = 44
  protocol                   = "TCP_SSL"
  certificate_ssl_mode       = "UNIDIRECTIONAL"
  certificate_id             = "%s"
  health_check_switch        = true
  health_check_time_out      = 30
  health_check_interval_time = 100
  health_check_health_num    = 2
  health_check_unhealth_num  = 2
  scheduler                  = "WRR"
  target_type         = "TARGETGROUP"
}
`
const testAccClbListener_tcpssl_update = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-tcpssl"
}

resource "cloud_clb_listener" "listener_tcpssl"{
  clb_id = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcpssl_update"
  port                       = 44
  protocol                   = "TCP_SSL"
  certificate_ssl_mode       = "UNIDIRECTIONAL"
  certificate_id             = "%s"
  health_check_switch        = true
  health_check_time_out      = 20
  health_check_interval_time = 200
  health_check_health_num    = 3
  health_check_unhealth_num  = 3
  scheduler                  = "WRR"
  target_type         = "TARGETGROUP"
}
`
const testAccClbListener_https = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-https"
}

resource "cloud_clb_listener" "listener_https" {
  clb_id               = cloud_clb_instance.clb_basic.id
  listener_name        = "listener_https"
  port                 = 77
  protocol             = "HTTPS"
  certificate_ssl_mode = "UNIDIRECTIONAL"
  certificate_id       = "%s"
  sni_switch           = false
}
`

const testAccClbListener_https_update = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-https"
}

resource "cloud_clb_listener" "listener_https" {
  clb_id               = cloud_clb_instance.clb_basic.id
  listener_name        = "listener_https_update"
  port                 = 33
  protocol             = "HTTPS"
  certificate_ssl_mode = "UNIDIRECTIONAL"
  certificate_id       = "%s"
  sni_switch           = false
}
`

const clb_tcp_instance = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-listener-tcp-tcp"
}
`

const testAccClbListener_tcp_tcp = clb_tcp_instance + `
resource "cloud_clb_listener" "listener_tcp" {
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcp"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 30
  health_check_interval_time = 100
  health_check_health_num    = 2
  health_check_unhealth_num  = 2
  session_expire_time        = 30
  scheduler                  = "WRR"
  health_check_type          = "TCP"
  health_check_port          = 200
}
`

const testAccClbListener_tcp_update_tcp = clb_tcp_instance + `
resource "cloud_clb_listener" "listener_tcp"{
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcp_update"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 20
  health_check_interval_time = 200
  health_check_health_num    = 3
  health_check_unhealth_num  = 3
  session_expire_time        = 60
  scheduler                  = "WRR"
  health_check_type          = "TCP"
  health_check_port          = 333
}
`

const clb_http_instance = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-listener-tcp-http"
}
`

const testAccClbListener_tcp_http = clb_http_instance + `
resource "cloud_clb_listener" "listener_tcp" {
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcp"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 30
  health_check_interval_time = 100
  health_check_health_num    = 2
  health_check_unhealth_num  = 2
  session_expire_time        = 30
  scheduler                  = "WRR"
  health_check_type          = "HTTP"
  health_check_http_domain   = "www.tencent.com"
  health_check_http_code     = 16
  health_check_http_version  = "HTTP/1.1"
  health_check_http_method   = "HEAD"
  health_check_http_path     = "/"
}
`

const testAccClbListener_tcp_update_http = clb_http_instance + `
resource "cloud_clb_listener" "listener_tcp"{
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcp_update"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 20
  health_check_interval_time = 200
  health_check_health_num    = 3
  health_check_unhealth_num  = 3
  session_expire_time        = 60
  scheduler                  = "WRR"
  health_check_port          = 200
  health_check_type          = "HTTP"
  health_check_http_code     = 2
  health_check_http_version  = "HTTP/1.0"
  health_check_http_method   = "GET"
}
`

const clb_customer_instance = `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-listener-tcp-customer"
}
`

const testAccClbListener_tcp_customer = clb_customer_instance + `
resource "cloud_clb_listener" "listener_tcp"{
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcp"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 30
  health_check_interval_time = 100
  health_check_health_num    = 2
  health_check_unhealth_num  = 2
  session_expire_time        = 30
  scheduler                  = "WRR"
  health_check_type          = "CUSTOM"
  health_check_context_type  = "HEX"
  health_check_send_context  = "0123456789ABCDEF"
  health_check_recv_context  = "ABCD"
}
`

const testAccClbListener_tcp_customer_update = clb_customer_instance + `
resource "cloud_clb_listener" "listener_tcp"{
  clb_id                     = cloud_clb_instance.clb_basic.id
  listener_name              = "listener_tcp_update"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 20
  health_check_interval_time = 200
  health_check_health_num    = 3
  health_check_unhealth_num  = 3
  session_expire_time        = 60
  scheduler                  = "WRR"
  health_check_type          = "CUSTOM"
  health_check_context_type  = "TEXT"
  health_check_send_context  = "/get/test"
  health_check_recv_context  = "http_1xx"
}
`
