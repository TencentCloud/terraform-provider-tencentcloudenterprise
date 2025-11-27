package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDataSourceImagesBase(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudDataSourceImagesBase,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_images.foo"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_images.foo", "images.#"),
				),
			},
			{
				Config: testAccTencentCloudDataSourceImagesBaseWithFilter,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_images.foo"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_images.foo", "images.#"),
				),
			},
			{
				Config: testAccTencentCloudDataSourceImagesBaseWithOsName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_images.foo"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_images.foo", "images.#"),
				),
			},
			{
				Config: testAccTencentCloudDataSourceImagesBaseWithImageNameRegex,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_images.foo"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_images.foo", "images.#"),
				),
			},
			{
				Config: testAccTencentCloudDataSourceImagesBaseWithInstanceType,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_images.foo"),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_images.foo", "images.#"),
				),
			},
		},
	})
}

const testAccTencentCloudDataSourceImagesBase = `
data "tencentcloudenterprise_cvm_images" "foo" {
	result_output_file = "data_source_tc_images_test.txt"
}
`

const testAccTencentCloudDataSourceImagesBaseWithFilter = `
data "tencentcloudenterprise_cvm_images" "foo" {
	image_type = ["PRIVATE_IMAGE"]
}
`

const testAccTencentCloudDataSourceImagesBaseWithOsName = `
data "tencentcloudenterprise_cvm_images" "foo" {
  image_type = ["PUBLIC_IMAGE"]
  os_name = "CentOS 7.5"
}
`

const testAccTencentCloudDataSourceImagesBaseWithImageNameRegex = `
data "tencentcloudenterprise_cvm_images" "foo" {
  image_type = ["PUBLIC_IMAGE"]
  image_name_regex = "^CentOS\\s+7\\.5\\s+64\\w*"
}
`

const testAccTencentCloudDataSourceImagesBaseWithInstanceType = `
data "tencentcloudenterprise_cvm_images" "foo" {
  //instance_type = "S1.SMALL1"
  instance_type = "S5l.SMALL1"
}
`
