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

func TestAccTencentCloudClbAttachmentResource_tcp(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbServerAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbServerAttachment_tcp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbServerAttachmentExists("tencentcloudenterprise_clb_attachment.foo"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_attachment.foo", "clb_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_attachment.foo", "listener_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "protocol_type", "TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "targets.#", "1"),
				),
			}, {
				Config: testAccClbServerAttachment_tcp_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbServerAttachmentExists("tencentcloudenterprise_clb_attachment.foo"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_attachment.foo", "clb_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_attachment.foo", "listener_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "protocol_type", "TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "targets.#", "1"),
				),
			}, {
				Config: testAccClbServerAttachment_tcp_update_ssl,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbServerAttachmentExists("tencentcloudenterprise_clb_attachment.foo"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_attachment.foo", "clb_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_attachment.foo", "listener_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "protocol_type", "TCP_SSL"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "targets.#", "1"),
				),
			},
		},
	})
}

func TestAccTencentCloudClbAttachmentResource_http(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbServerAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccClbServerAttachment_http, defaultSshCertificate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbServerAttachmentExists("tencentcloudenterprise_clb_attachment.foo"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_attachment.foo", "clb_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_attachment.foo", "listener_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "protocol_type", "HTTPS"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "targets.#", "1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_clb_attachment.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckClbServerAttachmentDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clbService := ClbService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_clb_attachment" {
			continue
		}
		time.Sleep(5 * time.Second)
		items := strings.Split(rs.Primary.ID, "#")
		if len(items) != 3 {
			return fmt.Errorf("[CHECK][CLB attachment][Destroy] check: id %s of resource.cloud_clb_attachment is not match loc-xxx#lbl-xxx#lb-xxx", rs.Primary.ID)
		}
		locationId := items[0]
		listenerId := items[1]
		clbId := items[2]
		instance, err := clbService.DescribeAttachmentByPara(ctx, clbId, listenerId, locationId)
		if (instance != nil && !(len(instance.Targets) == 0 && locationId == "") && !(len(instance.Rules) == 0 && locationId != "")) && err == nil {
			return fmt.Errorf("[CHECK][CLB attachment][Destroy] check: CLB Attachment still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func TestAccTencentCloudClbAttachmentResource_argetGroups(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbServerAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccClbServerAttachment_multiple, defaultSshCertificate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbServerAttachmentExists("tencentcloudenterprise_clb_attachment.foo"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "targets.#", "2"),
				),
			},
			{
				Config: fmt.Sprintf(testAccClbServerAttachment_multiple_update, defaultSshCertificate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbServerAttachmentExists("tencentcloudenterprise_clb_attachment.foo"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_attachment.foo", "targets.#", "1"),
				),
			},
		},
	})
}

func testAccCheckClbServerAttachmentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CLB attachment][Exists] check: CLB Attachment %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CLB attachment][Exists] check: CLB Attachment id is not set")
		}
		clbService := ClbService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		items := strings.Split(rs.Primary.ID, "#")
		if len(items) != 3 {
			return fmt.Errorf("[CHECK][CLB attachment][Exists] check: id %s of resource.cloud_clb_attachment is not match loc-xxx#lbl-xxx#lb-xxx", rs.Primary.ID)
		}
		locationId := items[0]
		listenerId := items[1]
		clbId := items[2]
		instance, err := clbService.DescribeAttachmentByPara(ctx, clbId, listenerId, locationId)
		if err != nil {
			return err
		}
		if instance == nil || (len(instance.Targets) == 0 && locationId == "") || (len(instance.Rules) == 0 && locationId != "") {
			return fmt.Errorf("[CHECK][CLB attachment][Exists] id %s is not exist", rs.Primary.ID)
		}
		return nil
	}
}

const testAccClbServerAttachment_tcp = instanceCommonTestCase + `
resource "tencentcloudenterprise_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = "tf-clb-attach-tcp-test"
  vpc_id       = var.cvm_vpc_id
}

resource "tencentcloudenterprise_clb_listener" "foo" {
  clb_id                     = cloud_clb_instance.foo.id
  listener_name              = "tf-clb-attach-tcp-test"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 30
  health_check_interval_time = 100
  health_check_health_num    = 2
  health_check_unhealth_num  = 2
  session_expire_time        = 30
  scheduler                  = "WRR"
}

resource "tencentcloudenterprise_clb_attachment" "foo" {
  clb_id      = cloud_clb_instance.foo.id
  listener_id = cloud_clb_listener.foo.listener_id

  targets {
    instance_id = cloud_cvm_instance.default.id
    port        = 23
    weight      = 10
  }
}
`

const testAccClbServerAttachment_tcp_update = instanceCommonTestCase + presetCVM + `
resource "tencentcloudenterprise_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = "tf-clb-attach-tcp-test"
  vpc_id       = var.cvm_vpc_id
  snat_pro     = true
}

resource "tencentcloudenterprise_clb_listener" "foo" {
  clb_id                     = cloud_clb_instance.foo.id
  listener_name              = "tf-clb-attach-tcp-test"
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 30
  health_check_interval_time = 100
  health_check_health_num    = 2
  health_check_unhealth_num  = 2
  session_expire_time        = 30
  scheduler                  = "WRR"
}

resource "tencentcloudenterprise_clb_attachment" "foo" {
  clb_id      = cloud_clb_instance.foo.id
  listener_id = cloud_clb_listener.foo.listener_id

  targets {
    eni_ip      = local.cvm_private_ip
    port        = 23
    weight      = 50
  }
}
`

const testAccClbServerAttachment_tcp_update_ssl = instanceCommonTestCase + presetCVM + `
data "tencentcloudenterprise_ssl_certificates" "foo" {
  name = "keep-svr-ca"
}

resource "tencentcloudenterprise_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = "tf-clb-attach-tcp-ssl"
  vpc_id       = var.cvm_vpc_id
  snat_pro     = true
}

# This is will force new as expected
resource "tencentcloudenterprise_clb_listener" "foo" {
  clb_id                     = cloud_clb_instance.foo.id
  listener_name              = "tf-clb-attach-tcp-ssl"
  port                       = 44
  protocol                   = "TCP_SSL"
  health_check_switch        = true
  scheduler                  = "WRR"
  certificate_ssl_mode       = "UNIDIRECTIONAL"
  certificate_id             = data.cloud_ssl_certificates.foo.certificates.0.id
}

resource "tencentcloudenterprise_clb_attachment" "foo" {
  clb_id      = cloud_clb_instance.foo.id
  listener_id = cloud_clb_listener.foo.listener_id

  # cross network target
  targets {
    eni_ip      = local.cvm_private_ip
    port        = 23
    weight      = 50
  }
}
`

const testAccClbServerAttachment_http = instanceCommonTestCase + presetCVM + `
resource "tencentcloudenterprise_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = "tf-clb-attach-http-test"
  vpc_id       = var.cvm_vpc_id
}

resource "tencentcloudenterprise_clb_listener" "foo" {
  clb_id               = cloud_clb_instance.foo.id
  listener_name        = "tf-clb-attach-http-test"
  port                 = 77
  protocol             = "HTTPS"
  certificate_ssl_mode = "UNIDIRECTIONAL"
  certificate_id       = "%s"
}

resource "tencentcloudenterprise_clb_listener_rule" "foo" {
  clb_id              = cloud_clb_instance.foo.id
  listener_id         = cloud_clb_listener.foo.listener_id
  domain              = "abc.com"
  url                 = "/"
  session_expire_time = 30
  scheduler           = "WRR"
}

resource "tencentcloudenterprise_clb_attachment" "foo" {
  clb_id      = cloud_clb_instance.foo.id
  listener_id = cloud_clb_listener.foo.listener_id
  rule_id     = cloud_clb_listener_rule.foo.rule_id

  targets {
    instance_id = cloud_cvm_instance.default.id
    port        = 23
    weight      = 10
  }
}
`

const testAccClbServerAttachment_multiple = instanceCommonTestCase + `
resource "tencentcloudenterprise_cvm_instance" "update" {
  instance_name              = var.instance_name_update
  availability_zone          = var.availability_cvm_zone
  image_id                   = data.cloud_cvm_images.default.images.0.image_id
  instance_type              = data.cloud_cvm_instance_types.default.instance_types.0.instance_type
  system_disk_type           = "CLOUD_PREMIUM"
  system_disk_size           = 50
  allocate_public_ip         = true
  internet_max_bandwidth_out = 10
  vpc_id                     = var.cvm_vpc_id
  subnet_id                  = var.cvm_subnet_id
}

resource "tencentcloudenterprise_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = "tf-clb-attach-multi-test"
  vpc_id       = var.cvm_vpc_id
}

resource "tencentcloudenterprise_clb_listener" "foo" {
  clb_id               = cloud_clb_instance.foo.id
  listener_name        = "tf-clb-attach-multi-test"
  port                 = 77
  protocol             = "HTTPS"
  certificate_ssl_mode = "UNIDIRECTIONAL"
  certificate_id       = "%s"
}

resource "tencentcloudenterprise_clb_listener_rule" "foo" {
  clb_id              = cloud_clb_instance.foo.id
  listener_id         = cloud_clb_listener.foo.listener_id
  domain              = "abc.com"
  url                 = "/"
  session_expire_time = 30
  scheduler           = "WRR"
}

resource "tencentcloudenterprise_clb_attachment" "foo" {
  clb_id      = cloud_clb_instance.foo.id
  listener_id = cloud_clb_listener.foo.listener_id
  rule_id     = cloud_clb_listener_rule.foo.rule_id

  targets {
    instance_id = cloud_cvm_instance.default.id
    port        = 23
    weight      = 10
  }
  targets {
    instance_id = cloud_cvm_instance.update.id
    port        = 24
    weight      = 10
  }
}
`

const testAccClbServerAttachment_multiple_update = instanceCommonTestCase + `

resource "tencentcloudenterprise_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = "tf-clb-attach-multi-test"
  vpc_id       = var.cvm_vpc_id
}

resource "tencentcloudenterprise_clb_listener" "foo" {
  clb_id               = cloud_clb_instance.foo.id
  listener_name        = "tf-clb-attach-multi-test"
  port                 = 77
  protocol             = "HTTPS"
  certificate_ssl_mode = "UNIDIRECTIONAL"
  certificate_id       = "%s"
}

resource "tencentcloudenterprise_clb_listener_rule" "foo" {
  clb_id              = cloud_clb_instance.foo.id
  listener_id         = cloud_clb_listener.foo.listener_id
  domain              = "abc.com"
  url                 = "/"
  session_expire_time = 30
  scheduler           = "WRR"
}

resource "tencentcloudenterprise_clb_attachment" "foo" {
  clb_id      = cloud_clb_instance.foo.id
  listener_id = cloud_clb_listener.foo.listener_id
  rule_id     = cloud_clb_listener_rule.foo.rule_id

  targets {
    instance_id = cloud_cvm_instance.default.id
    port        = 23
    weight      = 10
  }
}
`
