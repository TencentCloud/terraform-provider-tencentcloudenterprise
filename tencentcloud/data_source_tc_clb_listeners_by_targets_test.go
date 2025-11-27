package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbListenersByTargetsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListenersByTargetsDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_clb_listeners_by_targets.listeners_by_targets")),
			},
		},
	})
}

const testAccClbListenersByTargetsDataSource = `

data "tencentcloudenterprise_clb_listeners_by_targets" "listeners_by_targets" {
  backends {
    vpc_id     = "vpc-4owdpnwr"
    private_ip = "203.0.113.10"
  }
}
`
