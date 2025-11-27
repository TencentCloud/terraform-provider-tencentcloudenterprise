package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudEipsDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEipsDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_eips.data_eips"),
					resource.TestCheckResourceAttr("data.cloud_eips.data_eips", "eip_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_eips.data_eips", "eip_list.0.eip_id"),
					resource.TestCheckResourceAttr("data.cloud_eips.data_eips", "eip_list.0.eip_name", "tf-test-eip"),
					resource.TestCheckResourceAttrSet("data.cloud_eips.data_eips", "eip_list.0.eip_type"),
					resource.TestCheckResourceAttrSet("data.cloud_eips.data_eips", "eip_list.0.status"),
					resource.TestCheckResourceAttrSet("data.cloud_eips.data_eips", "eip_list.0.public_ip"),
					resource.TestCheckResourceAttrSet("data.cloud_eips.data_eips", "eip_list.0.create_time"),

					testAccCheckTencentCloudDataSourceID("data.cloud_eips.tags"),
					resource.TestCheckResourceAttr("data.cloud_eips.tags", "eip_list.0.tags.test", "test"),
				),
			},
		},
	})
}

const testAccEipsDataSource = `
resource "tencentcloudenterprise_eip" "eip" {
  name = "tf-test-eip"

  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_eips" "data_eips" {
  eip_id = cloud_eip.eip.id
}

data "tencentcloudenterprise_eips" "tags" {
  tags = cloud_eip.eip.tags
}
`
