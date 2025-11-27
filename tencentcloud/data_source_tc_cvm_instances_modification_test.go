package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCvmInstancesModificationDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCvmInstancesModificationDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_instances_modification.foo"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_instances_modification.foo", "instance_type_config_status_list.#"),
				),
			},
		},
	})
}

const testAccCvmInstancesModificationDataSource = defaultCvmModificationVariable + `
data "tencentcloudenterprise_cvm_instances_modification" "foo" {
	instance_ids = [var.cvm_id]
}
`
