package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamGroupPolicyAttachmentsDataSource_basic(t *testing.T) {
	t.Parallel()

	rName := fmt.Sprintf("tf-cam-group-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamGroupPolicyAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamGroupPolicyAttachmentsDataSource_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCamGroupPolicyAttachmentExists("tencentcloudenterprise_cam_group_policy_attachment.test"),
					resource.TestCheckResourceAttr("data.tencentcloudenterprise_cam_group_policy_attachments.group_policy_attachments", "group_policy_attachment_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_group_policy_attachments.group_policy_attachments", "group_policy_attachment_list.0.group_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_group_policy_attachments.group_policy_attachments", "group_policy_attachment_list.0.policy_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_group_policy_attachments.group_policy_attachments", "group_policy_attachment_list.0.policy_name"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_group_policy_attachments.group_policy_attachments", "group_policy_attachment_list.0.create_mode"),
				),
			},
		},
	})
}

func testAccCamGroupPolicyAttachmentsDataSource_basic(rName string) string {
	return fmt.Sprintf(`
# Create a test group
resource "tencentcloudenterprise_cam_group" "test" {
  name   = "%s"
  remark = "test group for data source"
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

# Attach policy to group
resource "tencentcloudenterprise_cam_group_policy_attachment" "test" {
  group_id  = tencentcloudenterprise_cam_group.test.id
  policy_id = tencentcloudenterprise_cam_policy.test.id
}

# Query the attachment using data source
data "tencentcloudenterprise_cam_group_policy_attachments" "group_policy_attachments" {
  group_id = tencentcloudenterprise_cam_group.test.id
  
  depends_on = [tencentcloudenterprise_cam_group_policy_attachment.test]
}
`, rName, rName)
}
