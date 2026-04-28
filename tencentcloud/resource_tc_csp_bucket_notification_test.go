package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudCspBucketNotification_basic tests the full lifecycle (create/update/import/destroy)
// without SASL authentication.
func TestAccTencentCloudCspBucketNotification_basic(t *testing.T) {

	rName := "tencentcloudenterprise_csp_bucket_notification.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCspBucketNotificationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCspBucketNotificationBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "bucket", "est123-1255000115"),
					resource.TestCheckResourceAttr(rName, "notification_rule.#", "1"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.id", "test-rule-1"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.ckafka_instance_id", "ckafka-7k3pve8e"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.events.#", "2"),
					resource.TestCheckResourceAttrSet(rName, "notification_rule.0.endpoint"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.sasl_user", ""),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.sasl_password", ""),
				),
			},
			{
				Config: testAccCspBucketNotificationUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "notification_rule.#", "1"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.id", "test-rule-1"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.events.#", "1"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccTencentCloudCspBucketNotification_sasl tests the full lifecycle with SASL authentication.
func TestAccTencentCloudCspBucketNotification_sasl(t *testing.T) {

	rName := "tencentcloudenterprise_csp_bucket_notification.test_sasl"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCspBucketNotificationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCspBucketNotificationSasl(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "bucket", "est123-1255000115"),
					resource.TestCheckResourceAttr(rName, "notification_rule.#", "1"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.id", "test-rule-sasl"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.ckafka_instance_id", "ckafka-7k3pve8e"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.events.#", "2"),
					resource.TestCheckResourceAttrSet(rName, "notification_rule.0.endpoint"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.sasl_user", "123123"),
				),
			},
			{
				// Update: remove SASL
				Config: testAccCspBucketNotificationSaslRemoved(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.sasl_user", ""),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.sasl_password", ""),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"notification_rule.0.sasl_password"},
			},
		},
	})
}

func testAccCheckCspBucketNotificationExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource %s not found in state", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("resource %s has empty ID", n)
		}

		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)
		svc := CosService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn, useCspClient: true}

		config, err := svc.GetBucketNotification(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if config == nil || len(config.TopicConfigurations) == 0 {
			return fmt.Errorf("bucket notification not found for %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckCspBucketNotificationDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	svc := CosService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn, useCspClient: true}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_csp_bucket_notification" {
			continue
		}
		config, err := svc.GetBucketNotification(ctx, rs.Primary.ID)
		if err != nil {
			return nil
		}
		if config != nil && len(config.TopicConfigurations) > 0 {
			return fmt.Errorf("bucket notification still exists for %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCspBucketNotificationBasic() string {
	return `
resource "tencentcloudenterprise_csp_bucket_notification" "test" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-rule-1"
    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
    ckafka_instance_id = "ckafka-7k3pve8e"
  }
}
`
}

func testAccCspBucketNotificationUpdate() string {
	return `
resource "tencentcloudenterprise_csp_bucket_notification" "test" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-rule-1"
    events             = ["cos:ObjectCreated:*"]
    ckafka_instance_id = "ckafka-7k3pve8e"
  }
}
`
}

func testAccCspBucketNotificationSasl() string {
	return `
resource "tencentcloudenterprise_csp_bucket_notification" "test_sasl" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-rule-sasl"
    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
    ckafka_instance_id = "ckafka-7k3pve8e"
    sasl_user          = "123123"
    sasl_password      = "Tencent@321"
  }
}
`
}

func testAccCspBucketNotificationSaslRemoved() string {
	return `
resource "tencentcloudenterprise_csp_bucket_notification" "test_sasl" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-rule-sasl"
    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
    ckafka_instance_id = "ckafka-7k3pve8e"
  }
}
`
}
