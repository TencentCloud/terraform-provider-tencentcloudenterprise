package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCamUser_basic(t *testing.T) {
	t.Parallel()

	rName := fmt.Sprintf("tf-subacct-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamUserBasic(rName, "test", "acc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamUserExists("tencentcloudenterprise_cam_user.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_user.test", "name", rName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_user.test", "remark", "test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_user.test", "console_login", "false"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_user.test", "tags.env", "acc"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_user.test", "uid"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_user.test", "uin"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_user.test", "secret_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_user.test", "secret_key"),
				),
			},
			{
				ResourceName:            "tencentcloudenterprise_cam_user.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password", "secret_id", "secret_key"},
			},
		},
	})
}

func testAccCheckCamUserExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM user][Exists] resource %s not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM user][Exists] ID is not set")
		}

		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		instance, err := camService.DescribeUserById(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance == nil || instance.Uid == nil {
			return fmt.Errorf("[CHECK][CAM user][Exists] user not found")
		}
		return nil
	}
}

func testAccCheckCamUserDestroy(s *terraform.State) error {
	camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_user" {
			continue
		}
		instance, err := camService.DescribeUserById(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance != nil && instance.Uid != nil {
			return fmt.Errorf("[CHECK][CAM user][Destroy] user still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCamUserBasic(name, remark, tag string) string {
	phone := "13800000000"
	email := fmt.Sprintf("%s@test.local", name)
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_user" "test" {
  name                = "%s"
  remark              = "%s"
  console_login       = false
  need_reset_password = true
  use_api             = true
  phone_num           = "%s"
  country_code        = "86"
  email               = "%s"
  tags = {
    env = "%s"
  }
}
`, name, remark, phone, email, tag)
}

func testAccCamUserUpdate(name, remark, tag string) string {
	phone := "13800000001"
	email := fmt.Sprintf("%s-update@test.local", name)
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_user" "test" {
  name                = "%s"
  remark              = "%s"
  console_login       = true
  need_reset_password = false
  use_api             = true
  phone_num           = "%s"
  country_code        = "86"
  email               = "%s"
  tags = {
    env = "%s"
  }
}
`, name, remark, phone, email, tag)
}
