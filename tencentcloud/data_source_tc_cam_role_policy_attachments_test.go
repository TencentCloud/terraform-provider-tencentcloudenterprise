package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamRolePolicyAttachmentsDataSource_basic(t *testing.T) {
	t.Parallel()

	rName := fmt.Sprintf("tf-cam-role-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamRolePolicyAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamRolePolicyAttachmentsDataSource_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCamRolePolicyAttachmentExists("tencentcloudenterprise_cam_role_policy_attachment.test"),
					resource.TestCheckResourceAttr("data.tencentcloudenterprise_cam_role_policy_attachments.role_policy_attachments", "role_policy_attachment_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_role_policy_attachments.role_policy_attachments", "role_policy_attachment_list.0.role_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_role_policy_attachments.role_policy_attachments", "role_policy_attachment_list.0.policy_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_role_policy_attachments.role_policy_attachments", "role_policy_attachment_list.0.policy_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_role_policy_attachments.role_policy_attachments", "role_policy_attachment_list.0.create_mode"),
				),
			},
		},
	})
}

func testAccCamRolePolicyAttachmentsDataSource_basic(rName string) string {
	return fmt.Sprintf(`
# Create a test role
resource "tencentcloudenterprise_cam_role" "test" {
  name          = "%s"
  description   = "test role for data source"
  console_login = false
  document      = <<EOF
{
  "version": "2.0",
  "statement": [
    {
      "action": ["sts:AssumeRole"],
      "effect": "allow",
      "principal": {
        "qcs": ["qcs::cam::uin/100000000001:uin/100000000001"]
      }
    }
  ]
}
EOF
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

# Attach policy to role
resource "tencentcloudenterprise_cam_role_policy_attachment" "test" {
  role_id   = tencentcloudenterprise_cam_role.test.id
  policy_id = tencentcloudenterprise_cam_policy.test.id
}

# Query the attachment using data source
data "tencentcloudenterprise_cam_role_policy_attachments" "role_policy_attachments" {
  role_id = tencentcloudenterprise_cam_role.test.id
  
  depends_on = [tencentcloudenterprise_cam_role_policy_attachment.test]
}
`, rName, rName)
}
