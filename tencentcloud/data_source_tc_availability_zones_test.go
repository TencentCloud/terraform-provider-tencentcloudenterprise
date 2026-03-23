package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudAvailabilityZonesDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudAvailabilityZonesDataSourceConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloudenterprise_availability_zones.all"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_availability_zones.all", "zones.#"),
				),
			},
			{
				Config: testAccTencentCloudAvailabilityZonesDataSourceConfigFilterWithName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloudenterprise_availability_zones.filter"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_availability_zones.filter", "zones.#"),
				),
			},
			{
				Config: testAccTencentCloudAvailabilityZonesDataSourceConfigIncludeUnavailable,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloudenterprise_availability_zones.unavailable"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_availability_zones.unavailable", "zones.#"),
				),
			},
		},
	})
}

const testAccTencentCloudAvailabilityZonesDataSourceConfigBasic = `
data "tencentcloudenterprise_availability_zones" "all" {
}
`

const testAccTencentCloudAvailabilityZonesDataSourceConfigFilterWithName = defaultVpcVariable + `
data "tencentcloudenterprise_availability_zones" "filter" {
  name = var.availability_zone
}
`

const testAccTencentCloudAvailabilityZonesDataSourceConfigIncludeUnavailable = `
data "tencentcloudenterprise_availability_zones" "unavailable" {
  include_unavailable = true
}
`
