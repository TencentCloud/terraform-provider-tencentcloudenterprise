package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCamPolicyByName_basic(t *testing.T) {
	t.Parallel()

	rName := fmt.Sprintf("tf-policy-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamPolicyByNameDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamPolicyByNameBasic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamPolicyByNameExists("tencentcloudenterprise_cam_policy_by_name.policy"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_policy_by_name.policy", "name", rName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_policy_by_name.policy", "description", "test policy"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_policy_by_name.policy", "document"),
				),
			},
			{
				Config: testAccCamPolicyByNameUpdate(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamPolicyByNameExists("tencentcloudenterprise_cam_policy_by_name.policy"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_policy_by_name.policy", "description", "updated test policy"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cam_policy_by_name.policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCamPolicyByNameDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_policy_by_name" {
			continue
		}

		params := make(map[string]interface{})
		params["name"] = rs.Primary.ID
		policies, err := camService.DescribePoliciesByFilter(ctx, params)
		if err == nil && len(policies) > 0 {
			return fmt.Errorf("[CHECK][CAM policy by name][Destroy] check: CAM policy still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckCamPolicyByNameExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM policy by name][Exists] check: CAM policy %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM policy by name][Exists] check: CAM policy name is not set")
		}
		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		params := make(map[string]interface{})
		params["name"] = rs.Primary.ID
		policies, err := camService.DescribePoliciesByFilter(ctx, params)
		if err != nil {
			return err
		}
		if len(policies) == 0 {
			return fmt.Errorf("[CHECK][CAM policy by name][Exists] check: CAM policy %s does not exist", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCamPolicyByNameBasic(name string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_policy_by_name" "policy" {
  name        = "%s"
  description = "test policy"
  document    = <<EOF
{
  "version": "2.0",
  "statement": [
    {
      "action": [
        "name/sts:AssumeRole"
      ],
      "effect": "allow",
      "resource": [
        "*"
      ]
    }
  ]
}
EOF
}
`, name)
}

func testAccCamPolicyByNameUpdate(name string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_policy_by_name" "policy" {
  name        = "%s"
  description = "updated test policy"
  document    = <<EOF
{
  "version": "2.0",
  "statement": [
    {
      "action": [
        "name/sts:AssumeRole"
      ],
      "effect": "allow",
      "resource": [
        "*"
      ]
    }
  ]
}
EOF
}
`, name)
}
