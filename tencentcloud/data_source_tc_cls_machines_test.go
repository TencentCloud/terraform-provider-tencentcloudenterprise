package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClsMachinesDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClsMachinesDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_cls_machines.machines")),
			},
		},
	})
}

const testAccClsMachinesDataSource = `

resource "tencentcloudenterprise_cls_machine_group" "group" {
  group_name        = "tf-describe-mg-test"
  service_logging   = true
  auto_update       = true
  update_end_time   = "19:05:00"
  update_start_time = "17:05:00"

  machine_group_type {
    type   = "ip"
    values = [
      "203.0.113.101",
      "203.0.113.102",
    ]
  }
}

data "tencentcloudenterprise_cls_machines" "machines" {
  group_id = cloud_cls_machine_group.group.id
}

`
