package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudCamRolePolicyAttachmentResource_basic -v
func TestAccTencentCloudCamRolePolicyAttachmentResource_basic(t *testing.T) {
	// t.Parallel()
	rRole := fmt.Sprintf("tf-role-%s", acctest.RandString(6))
	rPolicy := fmt.Sprintf("tf-policy-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamRolePolicyAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamRolePolicyAttachment_basic(rRole, rPolicy),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamRolePolicyAttachmentExists("tencentcloudenterprise_cam_role_policy_attachment.role_policy_attachment_basic"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_role_policy_attachment.role_policy_attachment_basic", "role_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_role_policy_attachment.role_policy_attachment_basic", "policy_id"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cam_role_policy_attachment.role_policy_attachment_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCamRolePolicyAttachmentDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_role_policy_attachment" {
			continue
		}

		instance, err := camService.DescribeRolePolicyAttachmentById(ctx, rs.Primary.ID)
		if err == nil && instance != nil {
			return fmt.Errorf("[CHECK][CAM role policy attachment][Desctroy] check: CAM role policy attachment still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckCamRolePolicyAttachmentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM role policy attachment][Exist] check: CAM role policy attachment %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM role policy attachment][Exist] check: CAM role policy attachment id is not set")
		}
		camService := CamService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		instance, err := camService.DescribeRolePolicyAttachmentById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance == nil {
			return fmt.Errorf("[CHECK][CAM role policy attachment][Exist] check: CAM role policy attachment %s is not exist", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCamRolePolicyAttachment_basic(roleName, policyName string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_policy" "policy" {
  name        = "%s"
  description = "acc policy"
  document    = <<EOF
{
  "version": "2.0",
  "statement": [
    {
      "effect": "allow",
      "action": [
        "account:DescribeSubAccounts"
      ],
      "resource": ["*"]
    }
  ]
}
EOF
}

resource "tencentcloudenterprise_cam_role" "role" {
  name          = "%s"
  description   = "acc role"
  console_login = false
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

resource "tencentcloudenterprise_cam_role_policy_attachment" "role_policy_attachment_basic" {
  role_id   = tencentcloudenterprise_cam_role.role.id
  policy_id = tencentcloudenterprise_cam_policy.policy.id
}

`, policyName, roleName)

}
