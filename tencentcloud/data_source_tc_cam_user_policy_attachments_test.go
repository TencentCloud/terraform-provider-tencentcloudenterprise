package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamUserPolicyAttachmentsDataSource_basic(t *testing.T) {
	t.Parallel()

	rName := fmt.Sprintf("tf-cam-user-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamUserPolicyAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamUserPolicyAttachmentsDataSource_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.tencentcloudenterprise_cam_user_policy_attachments.user_policy_attachments", "user_policy_attachment_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_user_policy_attachments.user_policy_attachments", "user_policy_attachment_list.0.user_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_user_policy_attachments.user_policy_attachments", "user_policy_attachment_list.0.policy_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_user_policy_attachments.user_policy_attachments", "user_policy_attachment_list.0.policy_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_user_policy_attachments.user_policy_attachments", "user_policy_attachment_list.0.create_mode"),
				),
			},
		},
	})
}

func testAccCamUserPolicyAttachmentsDataSource_basic(rName string) string {
	return fmt.Sprintf(`
# Create a test user
resource "tencentcloudenterprise_cam_user" "test" {
  name          = "%s"
  remark        = "test user for data source"
  console_login = false
  phone_num     = "13800000000"
  country_code  = "86"
  email         = "%s@test.local"
}

# Create a test policy
resource "tencentcloudenterprise_cam_policy" "test" {
  name        = "%s-policy"
  description = "test policy for data source"
  document    = jsonencode({
    version = "2.0"
    statement = [{
      effect = "allow"
      action = ["cos:GetObject"]
      resource = ["*"]
    }]
  })
}

# Attach policy to user
resource "tencentcloudenterprise_cam_user_policy_attachment" "test" {
  user_id   = tencentcloudenterprise_cam_user.test.name
  policy_id = tencentcloudenterprise_cam_policy.test.id
}

# Query the attachment using data source
data "tencentcloudenterprise_cam_user_policy_attachments" "user_policy_attachments" {
  user_name = tencentcloudenterprise_cam_user.test.name
  
  depends_on = [tencentcloudenterprise_cam_user_policy_attachment.test]
}
`, rName, rName, rName)
}
