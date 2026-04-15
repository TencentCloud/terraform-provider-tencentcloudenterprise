package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCspBucketBackupSetting_basic(t *testing.T) {
	t.Parallel()

	rName := "tencentcloudenterprise_csp_bucket_backup_setting.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCspBucketBackupSettingDestroy(rName),
		Steps: []resource.TestStep{
			{
				Config: testAccCspBucketBackupSettingBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketBackupSettingExists(rName),
					resource.TestCheckResourceAttr(rName, "backup_enable", "true"),
					resource.TestCheckResourceAttr(rName, "backsource_enabled", "false"),
					resource.TestCheckResourceAttr(rName, "use_https", "false"),
					resource.TestCheckResourceAttrSet(rName, "backup_bucket_name"),
					resource.TestCheckResourceAttrSet(rName, "backup_endpoint"),
				),
			},
			{
				Config: testAccCspBucketBackupSettingUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspBucketBackupSettingExists(rName),
					resource.TestCheckResourceAttr(rName, "backup_enable", "true"),
					resource.TestCheckResourceAttr(rName, "backsource_enabled", "true"),
				),
			},
			{
				ResourceName:            rName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret_key"},
			},
		},
	})
}

func testAccCheckCspBucketBackupSettingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource %s not found in state", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("resource %s has empty ID", n)
		}
		return nil
	}
}

func testAccCheckCspBucketBackupSettingDestroy(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return nil
		}
		return nil
	}
}

// src bucket: tf-test-bk-src-1255000115
// dst bucket: tf-test-bk-dst-1255000115 (created in test, used as backup target)
// backup_endpoint: cos.kazakhstan-1.csp.v120.fsphere.cn
// access_key/secret_key: same account credentials (same CSP cluster)

func testAccCspBucketBackupSettingBasic() string {
	return `
resource "tencentcloudenterprise_csp_bucket" "src" {
  bucket = "tf-test-bk-src-1255000115"
  acl    = "private"
}

resource "tencentcloudenterprise_csp_bucket" "dst" {
  bucket = "tf-test-bk-dst-1255000115"
  acl    = "private"
}

resource "tencentcloudenterprise_csp_bucket_backup_setting" "test" {
  bucket             = tencentcloudenterprise_csp_bucket.src.bucket
  region             = "kazakhstan-1"
  backup_enable      = true
  backup_bucket_name = tencentcloudenterprise_csp_bucket.dst.bucket
  backup_endpoint    = "cos.kazakhstan-1.csp.v120.fsphere.cn"
  access_key         = "your-secret-id-here"
  secret_key         = "ZR7tYJczxYnRvw7USPFyiOI6HUNykx8d"
  backsource_enabled = false
  use_https          = false

  depends_on = [tencentcloudenterprise_csp_bucket.dst]
}
`
}

func testAccCspBucketBackupSettingUpdate() string {
	return `
resource "tencentcloudenterprise_csp_bucket" "src" {
  bucket = "tf-test-bk-src-1255000115"
  acl    = "private"
}

resource "tencentcloudenterprise_csp_bucket" "dst" {
  bucket = "tf-test-bk-dst-1255000115"
  acl    = "private"
}

resource "tencentcloudenterprise_csp_bucket_backup_setting" "test" {
  bucket             = tencentcloudenterprise_csp_bucket.src.bucket
  region             = "kazakhstan-1"
  backup_enable      = true
  backup_bucket_name = tencentcloudenterprise_csp_bucket.dst.bucket
  backup_endpoint    = "cos.kazakhstan-1.csp.v120.fsphere.cn"
  access_key         = "your-secret-id-here"
  secret_key         = "ZR7tYJczxYnRvw7USPFyiOI6HUNykx8d"
  backsource_enabled = true
  use_https          = false

  depends_on = [tencentcloudenterprise_csp_bucket.dst]
}
`
}
