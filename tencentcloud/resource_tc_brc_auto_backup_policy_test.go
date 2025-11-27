package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudBrcAutoBackupPolicy_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccBrcAutoBackupPolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_brc_autobackup_policy.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_brc_autobackup_policy.example", "resource_type", "INSTANCE"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_brc_autobackup_policy.example", "auto_backup_policy_name", "test-policy"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_brc_autobackup_policy.example", "is_permanent", "false"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_brc_autobackup_policy.example", "full_backup_interval", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_brc_autobackup_policy.example", "retention_amount", "5"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_brc_autobackup_policy.example", "is_activated", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_brc_autobackup_policy.example", "dry_run", "false"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_brc_autobackup_policy.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccBrcAutoBackupPolicy_basic = `
resource "tencentcloudenterprise_brc_autobackup_policy" "example" {
  resource_type = "INSTANCE"
  auto_backup_policy_name = "test-policy"
  
  policy {
    hour = [0, 12]
    interval_days = 7
  }
  
  is_permanent = false
  full_backup_interval = 1
  retention_amount = 5
  
  advanced_retention_policy {
    days = 1
    weeks = 1
    months = 1
    years = 1
  }
  
  is_activated = true
  dry_run = false
}
`
