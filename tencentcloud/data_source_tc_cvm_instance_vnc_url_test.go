package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCvmInstanceVncUrlDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCvmInstanceVncUrlDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_cvm_instance_vnc_url.instance_vnc_url")),
			},
		},
	})
}

const testAccCvmInstanceVncUrlDataSource = defaultCvmModificationVariable + `

data "tencentcloudenterprise_cvm_instance_vnc_url" "instance_vnc_url" {
  instance_id = var.cvm_id
  result_output_file = "testAccCvmInstanceVncUrlDataSource.txt"
}
`
