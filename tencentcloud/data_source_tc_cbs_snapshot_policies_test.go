package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCbsSnapshotPoliciesDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_PREPAY) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCbsSnapshotPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCbsSnapshotPoliciesDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSnapshotPolicyExists("tencentcloudenterprise_cbs_snapshot_policy.policy"),
					resource.TestCheckResourceAttr("data.cloud_cbs_snapshot_policies.policies", "snapshot_policy_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_cbs_snapshot_policies.policies", "snapshot_policy_list.0.snapshot_policy_id"),
					resource.TestCheckResourceAttr("data.cloud_cbs_snapshot_policies.policies", "snapshot_policy_list.0.snapshot_policy_name", "tf-test-snapshot-policy"),
					resource.TestCheckResourceAttr("data.cloud_cbs_snapshot_policies.policies", "snapshot_policy_list.0.repeat_weekdays.#", "2"),
					resource.TestCheckResourceAttr("data.cloud_cbs_snapshot_policies.policies", "snapshot_policy_list.0.repeat_hours.#", "1"),
					resource.TestCheckResourceAttr("data.cloud_cbs_snapshot_policies.policies", "snapshot_policy_list.0.retention_days", "30"),
				),
			},
		},
	})
}

const testAccCbsSnapshotPoliciesDataSource = `
resource "tencentcloudenterprise_cbs_snapshot_policy" "policy" {
  snapshot_policy_name = "tf-test-snapshot-policy"
  repeat_weekdays      = [0, 3]
  repeat_hours         = [0]
  retention_days       = 30
}

data "tencentcloudenterprise_cbs_snapshot_policies" "policies" {
  snapshot_policy_id = cloud_cbs_snapshot_policy.policy.id
  snapshot_policy_name = cloud_cbs_snapshot_policy.policy.snapshot_policy_name
}
`
