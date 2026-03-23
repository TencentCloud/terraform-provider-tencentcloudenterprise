package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCamRole_basic(t *testing.T) {
	t.Parallel()

	rName := fmt.Sprintf("tf-role-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamRoleBasic(rName, "desc-1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamRoleExists("tencentcloudenterprise_cam_role.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role.test", "name", rName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role.test", "description", "desc-1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role.test", "console_login", "true"),
				),
			},
			{
				Config: testAccCamRoleBasic(rName, "desc-2", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamRoleExists("tencentcloudenterprise_cam_role.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role.test", "name", rName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role.test", "description", "desc-2"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role.test", "console_login", "false"),
				),
			},
			{
				ResourceName:            "tencentcloudenterprise_cam_role.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"document"},
			},
		},
	})
}

func testAccCheckCamRoleExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM role][Exists] resource %s not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM role][Exists] ID is not set")
		}

		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		instance, err := camService.DescribeRoleById(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance == nil || instance.RoleId == nil {
			return fmt.Errorf("[CHECK][CAM role][Exists] role not found")
		}
		return nil
	}
}

func testAccCheckCamRoleDestroy(s *terraform.State) error {
	camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_role" {
			continue
		}
		instance, err := camService.DescribeRoleById(context.Background(), rs.Primary.ID)
		if err != nil {
			if strings.Contains(err.Error(), "RoleNotExist") || strings.Contains(err.Error(), "NoSuchEntity") {
				continue
			}
			return err
		}
		if instance != nil && instance.RoleId != nil {
			return fmt.Errorf("[CHECK][CAM role][Destroy] role still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCamRoleBasic(name, desc string, consoleLogin bool) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_role" "test" {
  name          = "%s"
  description   = "%s"
  console_login = %t
  document      = <<EOF
{
  "version": "2.0",
  "statement": [
    {
      "action": ["sts:AssumeRole"],
      "effect": "allow",
      "principal": {
        "qcs": ["qcs::cam::uin/110000000285:uin/110000000285"]
      }
    }
  ]
}
EOF
}
`, name, desc, consoleLogin)
}
