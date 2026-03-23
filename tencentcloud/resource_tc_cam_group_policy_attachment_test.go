package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCamGroupPolicyAttachment_basic(t *testing.T) {
	t.Parallel()

	rGroup := fmt.Sprintf("tf-group-%s", acctest.RandString(6))
	rPolicy := fmt.Sprintf("tf-policy-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamGroupPolicyAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamGroupPolicyAttachmentBasic(rGroup, rPolicy),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamGroupPolicyAttachmentExists("tencentcloudenterprise_cam_group_policy_attachment.group_policy_attachment_basic"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group_policy_attachment.group_policy_attachment_basic", "group_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group_policy_attachment.group_policy_attachment_basic", "policy_id"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cam_group_policy_attachment.group_policy_attachment_basic",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["tencentcloudenterprise_cam_group_policy_attachment.group_policy_attachment_basic"]
					if !ok {
						return "", fmt.Errorf("resource not found in state")
					}
					return fmt.Sprintf("%s#%s", rs.Primary.Attributes["group_id"], rs.Primary.Attributes["policy_id"]), nil
				},
			},
		},
	})
}

func testAccCheckCamGroupPolicyAttachmentDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_group_policy_attachment" {
			continue
		}

		instance, err := camService.DescribeGroupPolicyAttachmentById(ctx, rs.Primary.ID)
		if err == nil && instance != nil {
			return fmt.Errorf("[CHECK][CAM group policy attachment][Destroy] check: CAM group policy attachment still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckCamGroupPolicyAttachmentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM group policy attachment][Exists] check: CAM group policy attachment %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM group policy attachment][Exists] check: CAM group policy attachment id is not set")
		}
		camService := CamService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		instance, err := camService.DescribeGroupPolicyAttachmentById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance == nil {
			return fmt.Errorf("[CHECK][CAM group policy attachment][Exists] check: CAM group policy attachment %s is not exist", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCamGroupPolicyAttachmentBasic(groupName, policyName string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_group" "group" {
  name        = "%s"
  remark      = "acc group"
}

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

resource "tencentcloudenterprise_cam_group_policy_attachment" "group_policy_attachment_basic" {
  group_id  = tencentcloudenterprise_cam_group.group.id
  policy_id = tencentcloudenterprise_cam_policy.policy.id
}

`, groupName, policyName)
}
