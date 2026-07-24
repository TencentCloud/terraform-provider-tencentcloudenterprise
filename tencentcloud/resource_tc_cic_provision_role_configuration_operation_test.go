package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCicProvisionRoleConfigurationOperationResource_basic(t *testing.T) {
	t.Parallel()
	suffix := acctest.RandString(6)
	const (
		existingRoleConfigurationId = "rc-xxxxxx"
		existingTargetUin           = "100001234567"
	)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCicProvisionRoleConfigurationOperation(suffix, existingRoleConfigurationId, existingTargetUin),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_provision_role_configuration_operation.example", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_provision_role_configuration_operation.example", "zone_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cic_provision_role_configuration_operation.example", "role_configuration_id", existingRoleConfigurationId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cic_provision_role_configuration_operation.example", "target_type", "MemberUin"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cic_provision_role_configuration_operation.example", "target_uin", existingTargetUin),
				),
			},
		},
	})
}

func testAccCicProvisionRoleConfigurationOperation(suffix, roleConfigurationId, targetUin string) string {
	return fmt.Sprintf(`
data "tencentcloudenterprise_cic_identity_center" "center" {}

resource "tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment" "example" {
  zone_id               = data.tencentcloudenterprise_cic_identity_center.center.zone_id
  role_configuration_id = "%s"
  role_policy_name      = "TfProvisionCustomPolicy%s"
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

resource "tencentcloudenterprise_cic_provision_role_configuration_operation" "example" {
  zone_id               = data.tencentcloudenterprise_cic_identity_center.center.zone_id
  role_configuration_id = "%s"
  target_type           = "MemberUin"
  target_uin            = %s
  depends_on            = [tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment.example]
}
`, roleConfigurationId, suffix, roleConfigurationId, targetUin)
}
