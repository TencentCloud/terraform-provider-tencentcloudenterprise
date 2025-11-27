package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudEipDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudEipDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_eip.my_eip"),
					resource.TestCheckResourceAttrSet("data.cloud_eip.my_eip", "id"),
					resource.TestCheckResourceAttrSet("data.cloud_eip.my_eip", "public_ip"),
				),
			},
			{
				Config: testAccTencentCloudEipDataSourceConfig_filter,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_eip.my_eip"),
					resource.TestCheckResourceAttrSet("data.cloud_eip.my_eip", "id"),
					resource.TestCheckResourceAttrSet("data.cloud_eip.my_eip", "public_ip"),
					resource.TestCheckResourceAttr("data.cloud_eip.my_eip", "status", "UNBIND"),
				),
			},
		},
	})
}

const testAccTencentCloudEipDataSourceConfig_basic = `
resource "tencentcloudenterprise_eip" "my_eip" {
	name = "tf-ci-test"
}

data "tencentcloudenterprise_eip" "my_eip" {
	filter {
		name = "address-id"
		values = [cloud_eip.my_eip.id]
	}
}
`

const testAccTencentCloudEipDataSourceConfig_filter = `
resource "tencentcloudenterprise_eip" "my_eip" {
	name = "tf-ci-test"
}

data "tencentcloudenterprise_eip" "my_eip" {
	filter {
		name   = "address-name"
		values = [cloud_eip.my_eip.name]
	}
}
`
