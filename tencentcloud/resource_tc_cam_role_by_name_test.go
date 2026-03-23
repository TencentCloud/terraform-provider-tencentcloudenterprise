package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCamRoleByName_basic(t *testing.T) {
	t.Parallel()

	rName := fmt.Sprintf("tf-role-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamRoleByNameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamRoleByNameBasic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamRoleByNameExists("tencentcloudenterprise_cam_role_by_name.role"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role_by_name.role", "name", rName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role_by_name.role", "description", "test role"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role_by_name.role", "console_login", "true"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_role_by_name.role", "document"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_role_by_name.role", "create_time"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_role_by_name.role", "update_time"),
				),
			},
			{
				Config: testAccCamRoleByNameUpdate(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamRoleByNameExists("tencentcloudenterprise_cam_role_by_name.role"),
					testAccCheckCamRoleByNameDescription("tencentcloudenterprise_cam_role_by_name.role", "updated test role"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_role_by_name.role", "description", "updated test role"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cam_role_by_name.role",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCamRoleByNameDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_role_by_name" {
			continue
		}

		params := make(map[string]interface{})
		params["name"] = rs.Primary.ID
		roles, err := camService.DescribeRolesByFilter(ctx, params)
		if err == nil && len(roles) > 0 {
			return fmt.Errorf("[CHECK][CAM role by name][Destroy] check: CAM role still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckCamRoleByNameExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM role by name][Exists] check: CAM role %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM role by name][Exists] check: CAM role name is not set")
		}
		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		params := make(map[string]interface{})
		params["name"] = rs.Primary.ID
		roles, err := camService.DescribeRolesByFilter(ctx, params)
		if err != nil {
			return err
		}
		if len(roles) == 0 {
			return fmt.Errorf("[CHECK][CAM role by name][Exists] check: CAM role %s does not exist", rs.Primary.ID)
		}

		log.Printf("[DEBUG] CAM role found: name=%s, description=%v, console_login=%v",
			*roles[0].RoleName, roles[0].Description, roles[0].ConsoleLogin)

		return nil
	}
}

func testAccCheckCamRoleByNameDescription(n string, expectedDescription string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM role by name][Description] check: CAM role %s is not found", n)
		}

		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		params := make(map[string]interface{})
		params["name"] = rs.Primary.ID
		roles, err := camService.DescribeRolesByFilter(ctx, params)
		if err != nil {
			return err
		}
		if len(roles) == 0 {
			return fmt.Errorf("[CHECK][CAM role by name][Description] check: CAM role %s does not exist", rs.Primary.ID)
		}

		actualDescription := ""
		if roles[0].Description != nil {
			actualDescription = *roles[0].Description
		}

		log.Printf("[DEBUG] Checking CAM role description: expected=%s, actual=%s", expectedDescription, actualDescription)

		if actualDescription != expectedDescription {
			return fmt.Errorf("[CHECK][CAM role by name][Description] check: expected description '%s', got '%s'",
				expectedDescription, actualDescription)
		}

		return nil
	}
}

func testAccCamRoleByNameBasic(name string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_role_by_name" "role" {
  name          = "%s"
  description   = "test role"
  console_login = true
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
`, name)
}

func testAccCamRoleByNameUpdate(name string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_role_by_name" "role" {
  name          = "%s"
  description   = "updated test role"
  console_login = true
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
`, name)
}
