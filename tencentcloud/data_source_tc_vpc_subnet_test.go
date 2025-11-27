package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTencentCloudSubnet_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccDataSourceTencentCloudSubnetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_subnet.foo"),
					resource.TestCheckResourceAttr("data.cloud_vpc_subnet.foo", "name", "tf-ci-test"),
					resource.TestCheckResourceAttr("data.cloud_vpc_subnet.foo", "availability_zone", "ap-guangzhou-3"),
				),
			},
		},
	})
}

const TestAccDataSourceTencentCloudSubnetConfig = `
variable "availability_zone" {
  default = "ap-guangzhou-3"
}

resource "tencentcloudenterprise_vpc" "foo" {
  name       = "tf-ci-test"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  availability_zone = var.availability_zone
  name              = "tf-ci-test"
  vpc_id            = cloud_vpc.foo.id
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

data "tencentcloudenterprise_vpc_subnet" "foo" {
  vpc_id    = cloud_vpc.foo.id
  subnet_id = cloud_vpc_subnet.subnet.id
}
`
