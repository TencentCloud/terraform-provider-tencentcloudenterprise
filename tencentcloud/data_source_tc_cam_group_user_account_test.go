package tencentcloud

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// 测试基础场景 - 使用 uid 查询子用户所在用户组
func TestAccTencentCloudCamGroupUserAccountDataSource_basic(t *testing.T) {
	t.Parallel()

	testUid := os.Getenv("TENCENTCLOUD_TEST_CAM_UID")
	if testUid == "" {
		t.Skip("TENCENTCLOUD_TEST_CAM_UID is not set, skipping test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamGroupUserAccountDataSourceByUid(testUid),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloudenterprise_cam_group_user_account.test"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_group_user_account.test", "total_num"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_group_user_account.test", "group_info.#"),
				),
			},
		},
	})
}

// 测试使用 uin 查询子用户所在用户组
func TestAccTencentCloudCamGroupUserAccountDataSource_byUin(t *testing.T) {
	t.Parallel()

	testUin := os.Getenv("TENCENTCLOUD_TEST_CAM_UIN")
	if testUin == "" {
		t.Skip("TENCENTCLOUD_TEST_CAM_UIN is not set, skipping test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamGroupUserAccountDataSourceByUin(testUin),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloudenterprise_cam_group_user_account.test"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_group_user_account.test", "total_num"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_group_user_account.test", "group_info.#"),
				),
			},
		},
	})
}

func testAccCamGroupUserAccountDataSourceByUid(uid string) string {
	return fmt.Sprintf(`
data "tencentcloudenterprise_cam_group_user_account" "test" {
  uid = %s
}
`, uid)
}

func testAccCamGroupUserAccountDataSourceByUin(uin string) string {
	return fmt.Sprintf(`
data "tencentcloudenterprise_cam_group_user_account" "test" {
  uin = %s
}
`, uin)
}
