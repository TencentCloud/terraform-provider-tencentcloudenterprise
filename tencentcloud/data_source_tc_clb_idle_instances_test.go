package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbIdleInstancesDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbIdleInstancesDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_clb_idle_instances.idle_instance")),
			},
		},
	})
}

const testAccClbIdleInstancesDataSource = `

data "tencentcloudenterprise_clb_idle_instances" "idle_instance" {
  load_balancer_region = "ap-guangzhou"
}
`
