package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCbsSnapshotSharePermissionResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckCommon(t, ACCOUNT_TYPE_PREPAY)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCbsSnapshotSharePermission,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cbs_snapshot_share_permission.snapshot_share_permission", "id"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cbs_snapshot_share_permission.snapshot_share_permission",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccCbsSnapshotSharePermission = `
resource "tencentcloudenterprise_cbs_snapshot_share_permission" "snapshot_share_permission" {
	account_ids = ["100004603073"]
	snapshot_id = "snap-7o0uk1pt"
}
`
