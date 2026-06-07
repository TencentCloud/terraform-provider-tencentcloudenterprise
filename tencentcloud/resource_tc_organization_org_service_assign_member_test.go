package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudOrganizationOrgServiceAssignMemberResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOrgServiceAssignMember,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_organization_org_service_assign_member.demo", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_organization_org_service_assign_member.demo", "service_id", "24"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_organization_org_service_assign_member.demo", "member_uin", "110000003066"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_organization_org_service_assign_member.demo", "management_scope", "1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_organization_org_service_assign_member.demo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccOrgServiceAssignMember = `
resource "tencentcloudenterprise_organization_org_service_assign_member" "demo" {
  service_id       = 24
  member_uin       = 110000003066
  management_scope = 1
}
`
