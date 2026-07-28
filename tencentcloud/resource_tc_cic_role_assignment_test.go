package tencentcloud

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCicRoleAssignmentResource_concurrentCreate(t *testing.T) {
	t.Parallel()

	zoneId := os.Getenv("TF_ACC_CIC_ZONE_ID")
	targetUin := os.Getenv("TF_ACC_CIC_TARGET_UIN")
	if zoneId == "" || targetUin == "" {
		t.Skip("set TF_ACC_CIC_ZONE_ID and TF_ACC_CIC_TARGET_UIN to run CIC role assignment acceptance test")
	}

	suffix := acctest.RandString(6)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Three assignments created in parallel (no depends_on between them).
				Config: testAccCicRoleAssignmentConcurrentCreate(suffix, zoneId, targetUin),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g1", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g2", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g3", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g1", "principal_name"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g2", "principal_name"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g3", "principal_name"),
				),
			},
		},
	})
}

func TestAccTencentCloudCicRoleAssignmentResource_multiGroupDestroy(t *testing.T) {
	t.Parallel()

	zoneId := os.Getenv("TF_ACC_CIC_ZONE_ID")
	targetUin := os.Getenv("TF_ACC_CIC_TARGET_UIN")
	if zoneId == "" || targetUin == "" {
		t.Skip("set TF_ACC_CIC_ZONE_ID and TF_ACC_CIC_TARGET_UIN to run CIC role assignment acceptance test")
	}

	suffix := acctest.RandString(6)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Parallel create + concurrent destroy with DeprovisionForLast (lock + AuthorizationExist swallow).
				Config: testAccCicRoleAssignmentConcurrentCreate(suffix, zoneId, targetUin),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g1", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g2", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g3", "id"),
				),
			},
		},
	})
}

func TestAccTencentCloudCicRoleAssignmentResource_concurrentLastDeprovision(t *testing.T) {
	t.Parallel()

	zoneId := os.Getenv("TF_ACC_CIC_ZONE_ID")
	targetUin := os.Getenv("TF_ACC_CIC_TARGET_UIN")
	if zoneId == "" || targetUin == "" {
		t.Skip("set TF_ACC_CIC_ZONE_ID and TF_ACC_CIC_TARGET_UIN to run CIC role assignment acceptance test")
	}

	suffix := acctest.RandString(6)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// All assignments ask for last-deprovision; concurrent destroy must not fail when
				// more than one destroyer races on dismantle.
				Config: testAccCicRoleAssignmentConcurrentLastDeprovision(suffix, zoneId, targetUin),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g1", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cic_role_assignment.g2", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cic_role_assignment.g1", "deprovision_strategy", "DeprovisionForLastRoleAssignmentOnAccount"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cic_role_assignment.g2", "deprovision_strategy", "DeprovisionForLastRoleAssignmentOnAccount"),
				),
			},
		},
	})
}

func testAccCicRoleAssignmentConcurrentCreate(suffix, zoneId, targetUin string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cic_group" "g1" {
  zone_id     = "%[1]s"
  group_name  = "tf-ra-c1-%[2]s"
  description = "tf concurrent role assignment"
}

resource "tencentcloudenterprise_cic_group" "g2" {
  zone_id     = "%[1]s"
  group_name  = "tf-ra-c2-%[2]s"
  description = "tf concurrent role assignment"
}

resource "tencentcloudenterprise_cic_group" "g3" {
  zone_id     = "%[1]s"
  group_name  = "tf-ra-c3-%[2]s"
  description = "tf concurrent role assignment"
}

resource "tencentcloudenterprise_cic_role_configuration" "example" {
  zone_id                 = "%[1]s"
  role_configuration_name = "tf-ra-c-%[2]s"
  description             = "tf concurrent role assignment"
}

resource "tencentcloudenterprise_cic_role_assignment" "g1" {
  zone_id               = "%[1]s"
  role_configuration_id = tencentcloudenterprise_cic_role_configuration.example.role_configuration_id
  target_type           = "MemberUin"
  target_uin            = %[3]s
  principal_type        = "Group"
  principal_id          = tencentcloudenterprise_cic_group.g1.group_id
  deprovision_strategy  = "DeprovisionForLastRoleAssignmentOnAccount"
}

resource "tencentcloudenterprise_cic_role_assignment" "g2" {
  zone_id               = "%[1]s"
  role_configuration_id = tencentcloudenterprise_cic_role_configuration.example.role_configuration_id
  target_type           = "MemberUin"
  target_uin            = %[3]s
  principal_type        = "Group"
  principal_id          = tencentcloudenterprise_cic_group.g2.group_id
  deprovision_strategy  = "DeprovisionForLastRoleAssignmentOnAccount"
}

resource "tencentcloudenterprise_cic_role_assignment" "g3" {
  zone_id               = "%[1]s"
  role_configuration_id = tencentcloudenterprise_cic_role_configuration.example.role_configuration_id
  target_type           = "MemberUin"
  target_uin            = %[3]s
  principal_type        = "Group"
  principal_id          = tencentcloudenterprise_cic_group.g3.group_id
  deprovision_strategy  = "DeprovisionForLastRoleAssignmentOnAccount"
}
`, zoneId, suffix, targetUin)
}

func testAccCicRoleAssignmentConcurrentLastDeprovision(suffix, zoneId, targetUin string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cic_group" "g1" {
  zone_id     = "%[1]s"
  group_name  = "tf-ra-d1-%[2]s"
  description = "tf concurrent last deprovision"
}

resource "tencentcloudenterprise_cic_group" "g2" {
  zone_id     = "%[1]s"
  group_name  = "tf-ra-d2-%[2]s"
  description = "tf concurrent last deprovision"
}

resource "tencentcloudenterprise_cic_role_configuration" "example" {
  zone_id                 = "%[1]s"
  role_configuration_name = "tf-ra-d-%[2]s"
  description             = "tf concurrent last deprovision"
}

resource "tencentcloudenterprise_cic_role_assignment" "g1" {
  zone_id               = "%[1]s"
  role_configuration_id = tencentcloudenterprise_cic_role_configuration.example.role_configuration_id
  target_type           = "MemberUin"
  target_uin            = %[3]s
  principal_type        = "Group"
  principal_id          = tencentcloudenterprise_cic_group.g1.group_id
  deprovision_strategy  = "DeprovisionForLastRoleAssignmentOnAccount"
}

resource "tencentcloudenterprise_cic_role_assignment" "g2" {
  zone_id               = "%[1]s"
  role_configuration_id = tencentcloudenterprise_cic_role_configuration.example.role_configuration_id
  target_type           = "MemberUin"
  target_uin            = %[3]s
  principal_type        = "Group"
  principal_id          = tencentcloudenterprise_cic_group.g2.group_id
  deprovision_strategy  = "DeprovisionForLastRoleAssignmentOnAccount"
}
`, zoneId, suffix, targetUin)
}
