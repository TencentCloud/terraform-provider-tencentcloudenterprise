package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudOrganizationOrgNode_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_SMS) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationOrgNode,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_organization_org_node.org_node", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_organization_org_node.org_node", "parent_node_id", "2003721"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_organization_org_node.org_node", "name", "terraform_test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_organization_org_node.org_node", "remark", "for terraform test"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_organization_org_node.org_node",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccOrganizationOrgNode = `

resource "tencentcloudenterprise_organization_org_node" "org_node" {
  name           = "terraform_test"
  parent_node_id = 2003721
  remark         = "for terraform test"
}

`
