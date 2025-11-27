package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbTargetHealthDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbTargetHealthDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_clb_target_health.target_health")),
			},
		},
	})
}

const testAccClbTargetHealthDataSource = `

data "tencentcloudenterprise_clb_target_health" "target_health" {
  load_balancer_ids = ["lb-5dnrkgry"]
}
`
