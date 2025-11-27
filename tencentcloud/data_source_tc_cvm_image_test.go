package tencentcloud

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDataSourceImageBase(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudDataSourceImageBase,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_image.public_image"),
					resource.TestMatchResourceAttr("data.cloud_cvm_image.public_image", "image_id", regexp.MustCompile("^img-")),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_image.public_image", "image_name"),
				),
			},
			{
				Config: testAccTencentCloudDataSourceImageBaseWithFilter,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_image.public_image"),
					resource.TestMatchResourceAttr("data.cloud_cvm_image.public_image", "image_id", regexp.MustCompile("^img-")),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_image.public_image", "image_name"),
				),
			},
			{
				Config: testAccTencentCloudDataSourceImageBaseWithOsName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_image.public_image"),
					resource.TestMatchResourceAttr("data.cloud_cvm_image.public_image", "image_id", regexp.MustCompile("^img-")),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_image.public_image", "image_name"),
				),
			},
			{
				Config: testAccTencentCloudDataSourceImageBaseWithImageNameRegex,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_cvm_image.public_image"),
					resource.TestMatchResourceAttr("data.cloud_cvm_image.public_image", "image_id", regexp.MustCompile("^img-")),
					resource.TestCheckResourceAttrSet("data.cloud_cvm_image.public_image", "image_name"),
				),
			},
		},
	})
}

const testAccTencentCloudDataSourceImageBase = `
data "tencentcloudenterprise_cvm_image" "public_image" {
}
`

const testAccTencentCloudDataSourceImageBaseWithFilter = `
data "tencentcloudenterprise_cvm_image" "public_image" {
  filter {
    name   = "image-type"
    values = ["PUBLIC_IMAGE"]
  }
}
`

const testAccTencentCloudDataSourceImageBaseWithOsName = `
data "tencentcloudenterprise_cvm_image" "public_image" {
  # os_name = "TencentOS Server 3.2"
  os_name = "centos"

  filter {
    name   = "image-type"
    values = ["PUBLIC_IMAGE"]
  }
}
`

const testAccTencentCloudDataSourceImageBaseWithImageNameRegex = `
data "tencentcloudenterprise_cvm_image" "public_image" {
  image_name_regex = "^Windows\\s.*$"

  filter {
    name   = "image-type"
    values = ["PUBLIC_IMAGE"]
  }
}
`
