package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudInstanceTypesDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudInstanceTypesDataSourceConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloud_cvm_instance_types.t4c8g", "instance_types.0.cpu_core_count", "4"),
					resource.TestCheckResourceAttr("data.cloud_cvm_instance_types.t4c8g", "instance_types.0.memory_size", "8"),
					resource.TestCheckResourceAttr("data.cloud_cvm_instance_types.t4c8g", "instance_types.0.availability_zone", "ap-guangzhou-3"),
				),
			},
		},
	})
}

func TestAccTencentCloudInstanceTypesDataSource_sell(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudInstanceTypesDataSourceConfigSell,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloud_cvm_instance_types.t4c8g", "instance_types.0.cpu_core_count", "2"),
					resource.TestCheckResourceAttr("data.cloud_cvm_instance_types.t4c8g", "instance_types.0.memory_size", "2"),
					resource.TestCheckResourceAttr("data.cloud_cvm_instance_types.t4c8g", "instance_types.0.availability_zone", "ap-guangzhou-3"),
					resource.TestCheckResourceAttr("data.cloud_cvm_instance_types.t4c8g", "instance_types.0.family", "SA2"),
				),
			},
		},
	})
}

const testAccTencentCloudInstanceTypesDataSourceConfigBasic = `
data "tencentcloudenterprise_cvm_instance_types" "t4c8g" {
  availability_zone = "yfm18"
  cpu_core_count = 4
  memory_size    = 8
}
`

const testAccTencentCloudInstanceTypesDataSourceConfigSell = `
data "tencentcloudenterprise_cvm_instance_types" "t4c8g" {
  cpu_core_count = 2
  memory_size    = 2
  exclude_sold_out = true

  filter{
	name = "instance-family"
    values = ["SA2"]
  }

  filter{
	name = "zone"
    values = ["ap-guangzhou-3"]
  }
}
`
