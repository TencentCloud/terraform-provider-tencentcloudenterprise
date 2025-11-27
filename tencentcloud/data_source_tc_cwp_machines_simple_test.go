package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudNeedFixCwpMachinesSimpleDataSource_basic -v
func TestAccTencentCloudNeedFixCwpMachinesSimpleDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCwpMachinesSimpleDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cwp_machines_simple.example"),
					resource.TestCheckResourceAttrSet("data.cloud_cwp_machines_simple.example", "machine_type"),
					resource.TestCheckResourceAttrSet("data.cloud_cwp_machines_simple.example", "machine_region"),
				),
			},
		},
	})
}

const testAccCwpMachinesSimpleDataSource = `
data "tencentcloudenterprise_cwp_machines_simple" "example" {
  machine_type   = "CVM"
  machine_region = "ap-guangzhou"
  project_ids    = [0]
}
`
