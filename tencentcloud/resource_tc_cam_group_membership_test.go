package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCamGroupMembership_basic(t *testing.T) {
	t.Parallel()

	rGroup := fmt.Sprintf("tf-group-%s", acctest.RandString(6))
	rUser1 := fmt.Sprintf("tf-user-%s", acctest.RandString(6))
	rUser2 := fmt.Sprintf("tf-user-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamGroupMembershipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamGroupMembershipBasic(rGroup, rUser1, rUser2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamGroupMembershipExists("tencentcloudenterprise_cam_group_membership.group_membership_basic"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group_membership.group_membership_basic", "group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group_membership.group_membership_basic", "user_names.#", "2"),
				),
			},
			{
				Config: testAccCamGroupMembershipUpdate(rGroup, rUser1, rUser2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamGroupMembershipExists("tencentcloudenterprise_cam_group_membership.group_membership_basic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group_membership.group_membership_basic", "user_names.#", "1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cam_group_membership.group_membership_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCamGroupMembershipDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_group_membership" {
			continue
		}

		members, err := camService.DescribeGroupMembershipById(ctx, rs.Primary.ID)
		if err == nil && len(members) > 0 {
			return fmt.Errorf("[CHECK][CAM group membership][Destroy] check: CAM group membership still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckCamGroupMembershipExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM group membership][Exists] check: CAM group membership %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM group membership][Exists] check: CAM group membership id is not set")
		}
		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		members, err := camService.DescribeGroupMembershipById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if len(members) == 0 {
			return fmt.Errorf("[CHECK][CAM group membership][Exists] check: CAM group membership %s has no members", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCamGroupMembershipBasic(groupName, user1, user2 string) string {
	phone := "13800000000"
	email1 := fmt.Sprintf("%s@test.local", user1)
	email2 := fmt.Sprintf("%s@test.local", user2)
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_group" "group" {
  name   = "%s"
  remark = "test group for membership"
}

resource "tencentcloudenterprise_cam_user" "user1" {
  name                = "%s"
  remark              = "test user 1"
  console_login       = false
  need_reset_password = true
  use_api             = true
  phone_num           = "%s"
  country_code        = "86"
  email               = "%s"
}

resource "tencentcloudenterprise_cam_user" "user2" {
  name                = "%s"
  remark              = "test user 2"
  console_login       = false
  need_reset_password = true
  use_api             = true
  phone_num           = "%s"
  country_code        = "86"
  email               = "%s"
}

resource "tencentcloudenterprise_cam_group_membership" "group_membership_basic" {
  group_id   = tencentcloudenterprise_cam_group.group.id
  user_names = [tencentcloudenterprise_cam_user.user1.name, tencentcloudenterprise_cam_user.user2.name]
}
`, groupName, user1, phone, email1, user2, phone, email2)
}

func testAccCamGroupMembershipUpdate(groupName, user1, user2 string) string {
	phone := "13800000000"
	email1 := fmt.Sprintf("%s@test.local", user1)
	email2 := fmt.Sprintf("%s@test.local", user2)
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_group" "group" {
  name   = "%s"
  remark = "test group for membership"
}

resource "tencentcloudenterprise_cam_user" "user1" {
  name                = "%s"
  remark              = "test user 1"
  console_login       = false
  need_reset_password = true
  use_api             = true
  phone_num           = "%s"
  country_code        = "86"
  email               = "%s"
}

resource "tencentcloudenterprise_cam_user" "user2" {
  name                = "%s"
  remark              = "test user 2"
  console_login       = false
  need_reset_password = true
  use_api             = true
  phone_num           = "%s"
  country_code        = "86"
  email               = "%s"
}

resource "tencentcloudenterprise_cam_group_membership" "group_membership_basic" {
  group_id   = tencentcloudenterprise_cam_group.group.id
  user_names = [tencentcloudenterprise_cam_user.user1.name]
}
`, groupName, user1, phone, email1, user2, phone, email2)
}
