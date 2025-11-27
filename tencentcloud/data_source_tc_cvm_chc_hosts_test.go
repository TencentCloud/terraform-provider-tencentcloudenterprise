package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCvmChcHostsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCvmChcHostsDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_cvm_chc_hosts.chc_hosts")),
			},
		},
	})
}

const testAccCvmChcHostsDataSource = `

data "tencentcloudenterprise_cvm_chc_hosts" "chc_hosts" {
	chc_ids = ["chc-0brmw3wl"]
  filters {
    name = "zone"
    values = ["ap-guangzhou-7"]
  }
}
`
