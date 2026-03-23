package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCamUserPolicyAttachment_basic(t *testing.T) {
	t.Parallel()

	rUser := fmt.Sprintf("tf-user-%s", acctest.RandString(6))
	rPolicy := fmt.Sprintf("tf-policy-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamUserPolicyAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamUserPolicyAttachmentBasic(rUser, rPolicy),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamUserPolicyAttachmentExists("tencentcloudenterprise_cam_user_policy_attachment.user_policy_attachment_basic"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_user_policy_attachment.user_policy_attachment_basic", "user_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_user_policy_attachment.user_policy_attachment_basic", "policy_id"),
				),
			},
			{
				ResourceName:            "tencentcloudenterprise_cam_user_policy_attachment.user_policy_attachment_basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"policy_type"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["tencentcloudenterprise_cam_user_policy_attachment.user_policy_attachment_basic"]
					if !ok {
						return "", fmt.Errorf("resource not found in state")
					}
					return fmt.Sprintf("%s#%s", rs.Primary.Attributes["user_id"], rs.Primary.Attributes["policy_id"]), nil
				},
			},
		},
	})
}

func testAccCheckCamUserPolicyAttachmentDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_user_policy_attachment" {
			continue
		}

		instance, err := camService.DescribeUserPolicyAttachmentById(ctx, rs.Primary.ID)
		if err == nil && instance != nil {
			return fmt.Errorf("[CHECK][CAM user policy attachment][Destroy] check: CAM user policy attachment still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckCamUserPolicyAttachmentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM user policy attachment][Exists] check: CAM user policy attachment %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM user policy attachment][Exists] check: CAM user policy attachment id is not set")
		}
		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		instance, err := camService.DescribeUserPolicyAttachmentById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance == nil {
			return fmt.Errorf("[CHECK][CAM user policy attachment][Exists] check: CAM user policy attachment %s is not exist", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCamUserPolicyAttachmentBasic(userName, policyName string) string {
	phone := "13800000000"
	email := fmt.Sprintf("%s@test.local", userName)
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_user" "user" {
  name                = "%s"
  remark              = "acc user"
  console_login       = false
  need_reset_password = true
  use_api             = true
  phone_num           = "%s"
  country_code        = "86"
  email               = "%s"
}

resource "tencentcloudenterprise_cam_policy" "policy" {
  name        = "%s"
  description = "acc policy"
  document    = "{\"version\":\"2.0\",\"statement\":[{\"action\":[\"cos:*\"],\"resource\":[\"*\"],\"effect\":\"allow\"},{\"effect\":\"allow\",\"action\":[\"monitor:*\",\"cam:ListUsersForGroup\",\"cam:ListGroups\",\"cam:GetGroup\"],\"resource\":[\"*\"]}]}"
}

resource "tencentcloudenterprise_cam_user_policy_attachment" "user_policy_attachment_basic" {
  user_id   = tencentcloudenterprise_cam_user.user.id
  policy_id = tencentcloudenterprise_cam_policy.policy.id
}

`, userName, phone, email, policyName)
}
