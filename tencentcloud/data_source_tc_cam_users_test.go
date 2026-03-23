package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamUsersDataSource_basic(t *testing.T) {
	t.Parallel()

	rName := fmt.Sprintf("tf-cam-user-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamUsersDataSource_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.tencentcloudenterprise_cam_users.users", "user_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_users.users", "user_list.0.name"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_users.users", "user_list.0.user_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_users.users", "user_list.0.uin"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_users.users", "user_list.0.uid"),
				),
			},
			{
				Config: testAccCamUsersDataSource_all(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_users.all", "user_list.0.name"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_users.all", "user_list.0.uin"),
				),
			},
		},
	})
}

func TestAccTencentCloudCamUsersDataSource_all(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamUsersDataSource_allOnly,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_users.all", "user_list.0.name"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_users.all", "user_list.0.uin"),
				),
			},
		},
	})
}

func testAccCamUsersDataSource_basic(rName string) string {
	return testAccCamUserBasic(rName, "test user for data source", "acc") + `

data "tencentcloudenterprise_cam_users" "users" {
  name = tencentcloudenterprise_cam_user.test.name

  depends_on = [tencentcloudenterprise_cam_user.test]
}
`
}

func testAccCamUsersDataSource_all(rName string) string {
	return testAccCamUserBasic(rName, "test user for data source", "acc") + `

data "tencentcloudenterprise_cam_users" "all" {
  depends_on = [tencentcloudenterprise_cam_user.test]
}
`
}

const testAccCamUsersDataSource_allOnly = `
data "tencentcloudenterprise_cam_users" "all" {}
`
