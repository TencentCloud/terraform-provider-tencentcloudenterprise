package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCicRoleConfigurationPermissionCustomPolicyAttachmentResource_basic(t *testing.T) {
	t.Parallel()
	rName := fmt.Sprintf("tf-cp-%s", acctest.RandString(6))
	resourceName := "tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment.example"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCicRoleConfigurationPermissionCustomPolicyAttachment(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_id"),
					resource.TestCheckResourceAttrSet(resourceName, "role_configuration_id"),
					resource.TestCheckResourceAttr(resourceName, "role_policy_name", "TfCustomPolicy"),
					resource.TestCheckResourceAttrSet(resourceName, "role_policy_document"),
					resource.TestCheckResourceAttr(resourceName, "role_policy_type", "Custom"),
				),
			},
		},
	})
}

func testAccCicRoleConfigurationPermissionCustomPolicyAttachment(name string) string {
	return fmt.Sprintf(`
data "tencentcloudenterprise_cic_identity_center" "center" {}

resource "tencentcloudenterprise_cic_role_configuration" "example" {
  zone_id                 = data.tencentcloudenterprise_cic_identity_center.center.zone_id
  role_configuration_name = "%s"
  description             = "tf custom policy test"
}

resource "tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment" "example" {
  zone_id               = data.tencentcloudenterprise_cic_identity_center.center.zone_id
  role_configuration_id = tencentcloudenterprise_cic_role_configuration.example.role_configuration_id
  role_policy_name      = "TfCustomPolicy"
  role_policy_document  = <<-EOF
{
  "version": "2.0",
  "statement": [
    {
      "effect": "allow",
      "action": [
        "vpc:DescribeVpcs"
      ],
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
