package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDataSourceInstancesBase(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudDataSourceInstancesBase_1,
				Check: resource.ComposeTestCheckFunc(
					//testAccCheckTencentCloudInstanceExists("tencentcloudenterprise_cvm_instance.default"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_instances.foo", "instance_list.0.instance_id"),
					//resource.TestCheckResourceAttr("data.cloud_cvm_instances.foo", "instance_list.0.instance_name", defaultInsName),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_instances.foo", "instance_list.0.instance_type"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_instances.foo", "instance_list.0.cpu"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_instances.foo", "instance_list.0.memory"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_instances.foo", "instance_list.0.availability_zone"),
					resource.TestCheckResourceAttr("data.cloud_cvm_instances.foo", "instance_list.0.project_id", "0"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_instances.foo", "instance_list.0.system_disk_type"),
				),
			},
		},
	})
}

const testAccTencentCloudDataSourceInstancesBase = instanceCommonTestCase + `
data "tencentcloudenterprise_cvm_instances" "foo" {
  instance_id = cloud_cvm_instance.default.id
  instance_name = cloud_cvm_instance.default.instance_name
}
`
const testAccTencentCloudDataSourceInstancesBase_1 = `
data "tencentcloudenterprise_cvm_instances" "foo" {
}
`
