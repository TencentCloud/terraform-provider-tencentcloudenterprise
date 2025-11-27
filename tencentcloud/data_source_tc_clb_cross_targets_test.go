package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbCrossTargetsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbCrossTargetsDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_clb_cross_targets.cross_targets")),
			},
		},
	})
}

const testAccClbCrossTargetsDataSource = `

data "tencentcloudenterprise_clb_cross_targets" "cross_targets" {
  filters {
    name = "vpc-id"
    values = ["vpc-4owdpnwr"]
  }
}

`
