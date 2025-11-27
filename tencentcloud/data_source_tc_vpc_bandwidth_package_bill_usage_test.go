package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcBandwidthPackageBillUsageDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcBandwidthPackageBillUsageDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_vpc_bandwidth_package_bill_usage.bandwidth_package_bill_usage")),
			},
		},
	})
}

const testAccVpcBandwidthPackageBillUsageDataSource = `

resource "tencentcloudenterprise_vpc_bandwidth_package" "bandwidth_package" {
  network_type            = "BGP"
  charge_type             = "TOP5_POSTPAID_BY_MONTH"
  bandwidth_package_name  = "iac-test-data"
  tags = {
    "createdBy" = "terraform"
  }
}

data "tencentcloudenterprise_vpc_bandwidth_package_bill_usage" "bandwidth_package_bill_usage" {
  bandwidth_package_id =  cloud_vpc_bandwidth_package.bandwidth_package.id
}

`
