package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudCspBucketNotification_basic tests without SASL:
// create → update events → update with filter_prefix → import → destroy
func TestAccTencentCloudCspBucketNotification_basic(t *testing.T) {
	rName := "cloud_csp_bucket_notification.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCspBucketNotificationDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with basic config (no SASL, no filter)
				Config: testAccCspBucketNotification_noSasl_noFilter(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "bucket", "est123-1255000115"),
					resource.TestCheckResourceAttr(rName, "notification_rule.#", "1"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.id", "test-basic"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.topic", "csp"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.ckafka_instance_id", "ckafka-7k3pve8e"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.events.#", "2"),
					resource.TestCheckResourceAttrSet(rName, "notification_rule.0.endpoint"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.filter_prefix", ""),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.filter_suffix", ""),
				),
			},
			{
				// Step 2: Update - change events to single event
				Config: testAccCspBucketNotification_noSasl_updateEvents(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.events.#", "1"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.events.0", "cos:ObjectCreated:*"),
				),
			},
			{
				// Step 3: Update - add filter_prefix
				Config: testAccCspBucketNotification_noSasl_withFilter(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.filter_prefix", "logs/"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.filter_suffix", ".json"),
				),
			},
			{
				// Step 4: Import
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccTencentCloudCspBucketNotification_sasl tests with SASL authentication:
// create with SASL → update remove SASL → import → destroy
// Note: SASL route (AccessType=1) may need to be created, so this test may take longer.
func TestAccTencentCloudCspBucketNotification_sasl(t *testing.T) {
	rName := "cloud_csp_bucket_notification.test_sasl"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCspBucketNotificationDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with SASL
				Config: testAccCspBucketNotification_withSasl(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "bucket", "est123-1255000115"),
					resource.TestCheckResourceAttr(rName, "notification_rule.#", "1"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.id", "test-sasl"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.topic", "csp"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.ckafka_instance_id", "ckafka-7k3pve8e"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.events.#", "2"),
					resource.TestCheckResourceAttrSet(rName, "notification_rule.0.endpoint"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.sasl_user", "test123"),
				),
			},
			{
				// Step 2: Update - remove SASL (switch to non-SASL route)
				Config: testAccCspBucketNotification_saslRemoved(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.sasl_user", ""),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.sasl_password", ""),
					resource.TestCheckResourceAttrSet(rName, "notification_rule.0.endpoint"),
				),
			},
			{
				// Step 3: Import
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"notification_rule.0.sasl_password"},
			},
		},
	})
}

// TestAccTencentCloudCspBucketNotification_saslWithFilter tests SASL + filter combined.
func TestAccTencentCloudCspBucketNotification_saslWithFilter(t *testing.T) {
	rName := "cloud_csp_bucket_notification.test_sasl_filter"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCspBucketNotificationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCspBucketNotification_saslWithFilter(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketNotificationExists(rName),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.sasl_user", "test123"),
					resource.TestCheckResourceAttr(rName, "notification_rule.0.filter_prefix", "data/"),
					resource.TestCheckResourceAttrSet(rName, "notification_rule.0.endpoint"),
				),
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
		if rs.Type != "cloud_csp_bucket_notification" {
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

// --- Without SASL configs ---

func testAccCspBucketNotification_noSasl_noFilter() string {
	return `
resource "cloud_csp_bucket_notification" "test" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-basic"
    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
    topic              = "csp"
    ckafka_instance_id = "ckafka-7k3pve8e"
  }
}
`
}

func testAccCspBucketNotification_noSasl_updateEvents() string {
	return `
resource "cloud_csp_bucket_notification" "test" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-basic"
    events             = ["cos:ObjectCreated:*"]
    topic              = "csp"
    ckafka_instance_id = "ckafka-7k3pve8e"
  }
}
`
}

func testAccCspBucketNotification_noSasl_withFilter() string {
	return `
resource "cloud_csp_bucket_notification" "test" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-basic"
    events             = ["cos:ObjectCreated:*"]
    topic              = "csp"
    ckafka_instance_id = "ckafka-7k3pve8e"
    filter_prefix      = "logs/"
    filter_suffix      = ".json"
  }
}
`
}

// --- With SASL configs ---

func testAccCspBucketNotification_withSasl() string {
	return `
resource "cloud_csp_bucket_notification" "test_sasl" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-sasl"
    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
    topic              = "csp"
    ckafka_instance_id = "ckafka-7k3pve8e"
    sasl_user          = "test123"
    sasl_password      = "test123456"
  }
}
`
}

func testAccCspBucketNotification_saslRemoved() string {
	return `
resource "cloud_csp_bucket_notification" "test_sasl" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-sasl"
    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
    topic              = "csp"
    ckafka_instance_id = "ckafka-7k3pve8e"
  }
}
`
}

// --- SASL + Filter ---

func testAccCspBucketNotification_saslWithFilter() string {
	return `
resource "cloud_csp_bucket_notification" "test_sasl_filter" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "test-sasl-filter"
    events             = ["cos:ObjectCreated:*"]
    topic              = "csp"
    ckafka_instance_id = "ckafka-7k3pve8e"
    sasl_user          = "test123"
    sasl_password      = "test123456"
    filter_prefix      = "data/"
  }
}
`
}
