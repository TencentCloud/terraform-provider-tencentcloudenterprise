package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbTargetGroupListDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbTargetGroupListDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_clb_target_group_list.target_group_list")),
			},
		},
	})
}

const testAccClbTargetGroupListDataSource = `

data "tencentcloudenterprise_clb_target_group_list" "target_group_list" {
  filters {
    name = "TargetGroupName"
    values = ["keep-tgg"]
  }
}

`
