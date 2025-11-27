package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudRedisBackupDownloadRestrictionResource_basic -v
func TestAccTencentCloudRedisBackupDownloadRestrictionResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRedisBackupDownloadRestriction,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_redis_backup_download_restriction.backup_download_restriction", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_backup_download_restriction.backup_download_restriction", "limit_type", "Customize"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_backup_download_restriction.backup_download_restriction", "vpc_comparison_symbol", "In"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_backup_download_restriction.backup_download_restriction", "ip_comparison_symbol", "In"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_backup_download_restriction.backup_download_restriction", "limit_vpc.#", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_backup_download_restriction.backup_download_restriction", "limit_vpc.0.region", "ap-guangzhou"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_backup_download_restriction.backup_download_restriction", "limit_vpc.0.vpc_list.#", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_backup_download_restriction.backup_download_restriction", "limit_ip.#", "2"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_redis_backup_download_restriction.backup_download_restriction",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccRedisBackupDownloadRestrictionVar = `
variable "vpc_id" {
	default = "` + defaultCrsVpcId + `"
}
`

const testAccRedisBackupDownloadRestriction = testAccRedisBackupDownloadRestrictionVar + `

resource "tencentcloudenterprise_redis_backup_download_restriction" "backup_download_restriction" {
	limit_type = "Customize"
	vpc_comparison_symbol = "In"
	ip_comparison_symbol = "In"
	limit_vpc {
		  region = "ap-guangzhou"
		  vpc_list = ["vpc-4owdpnwr"]
	}
	limit_ip = ["203.0.113.12", "203.0.113.13"]
}

`
