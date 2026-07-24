package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Regression: updating only remark must not re-send Name (OrganizationMemberNameUsed).
func TestAccTencentCloudOrganizationOrgMemberResource_updateRemarkOnly(t *testing.T) {
	t.Parallel()
	name := fmt.Sprintf("tf-om-%s", acctest.RandString(6))
	resourceName := "tencentcloudenterprise_organization_org_member.example"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationOrgMemberBasic(name, "remark-v1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "remark", "remark-v1"),
					resource.TestCheckResourceAttr(resourceName, "policy_type", "Financial"),
				),
			},
			{
				Config: testAccOrganizationOrgMemberBasic(name, "remark-v2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "remark", "remark-v2"),
				),
			},
		},
	})
}

func testAccOrganizationOrgMemberBasic(name, remark string) string {
	return fmt.Sprintf(`
data "tencentcloudenterprise_organization_nodes" "nodes" {}

resource "tencentcloudenterprise_organization_org_member" "example" {
  name           = "%s"
  node_id        = data.tencentcloudenterprise_organization_nodes.nodes.items.0.node_id
  policy_type    = "Financial"
  permission_ids = [1, 2]
  remark         = "%s"
}
`, name, remark)
}
