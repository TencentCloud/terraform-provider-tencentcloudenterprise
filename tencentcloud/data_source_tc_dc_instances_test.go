package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTencentCloudDcV3InstancesBasic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccDataSourceTencentCloudDcInstancesNoFilter,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_dc_instances.list_all"),
					resource.TestCheckResourceAttrSet("data.cloud_dc_instances.list_all", "instance_list.#"),
				),
			},
			{
				Config: TestAccDataSourceTencentCloudDcInstancesByName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_dc_instances.name_select"),
					resource.TestCheckResourceAttrSet("data.cloud_dc_instances.name_select", "name"),
				),
			},
			{
				Config: TestAccDataSourceTencentCloudDcInstancesById,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_dc_instances.id_select"),
					resource.TestCheckResourceAttrSet("data.cloud_dc_instances.id_select", "dc_id"),
				),
			},
		},
	})
}

const TestAccDataSourceTencentCloudDcInstancesNoFilter = `
data cloud_dc_instances list_all {
}
`

const TestAccDataSourceTencentCloudDcInstancesByName = `
data cloud_dc_instances name_select {
    name = "zww-test"
}
`

const TestAccDataSourceTencentCloudDcInstancesById = `
data cloud_dc_instances id_select {
    dc_id = "dc-Lmb262v7"
}
`
