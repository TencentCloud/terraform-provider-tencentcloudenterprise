package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbInstanceByCertIdDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstanceByCertIdDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_clb_instance_by_cert_id.instance_by_cert_id")),
			},
		},
	})
}

const testAccClbInstanceByCertIdDataSource = `

data "cloud_clb_instance_by_cert_id" "instance_by_cert_id" {
  cert_ids = ["3a6B5y8v"]
}

`
