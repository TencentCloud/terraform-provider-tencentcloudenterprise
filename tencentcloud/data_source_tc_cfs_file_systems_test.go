package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCfsFileSystemsDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCfsFileSystemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCfsFileSystemsDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCfsFileSystemExists("tencentcloudenterprise_cfs_file_system.foo"),
					resource.TestCheckResourceAttr("data.cloud_cfs_file_systems.file_systems", "file_system_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_cfs_file_systems.file_systems", "file_system_list.0.file_system_id"),
					resource.TestCheckResourceAttrSet("data.cloud_cfs_file_systems.file_systems", "file_system_list.0.name"),
					resource.TestCheckResourceAttr("data.cloud_cfs_file_systems.file_systems", "file_system_list.0.availability_zone", "ap-guangzhou-3"),
					resource.TestCheckResourceAttr("data.cloud_cfs_file_systems.file_systems", "file_system_list.0.protocol", "NFS"),
					resource.TestCheckResourceAttrSet("data.cloud_cfs_file_systems.file_systems", "file_system_list.0.access_group_id"),
					resource.TestCheckResourceAttrSet("data.cloud_cfs_file_systems.file_systems", "file_system_list.0.status"),
					resource.TestCheckResourceAttrSet("data.cloud_cfs_file_systems.file_systems", "file_system_list.0.create_time"),
					resource.TestCheckResourceAttrSet("data.cloud_cfs_file_systems.file_systems", "file_system_list.0.mount_ip"),
				),
			},
		},
	})
}

const testAccCfsFileSystemsDataSource = defaultCfsAccessGroup + `
resource "tencentcloudenterprise_vpc" "vpc" {
  name       = "test-cfs-vpc"
  cidr_block = "10.2.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  vpc_id            = cloud_vpc.vpc.id
  name              = "test-cfs-subnet"
  cidr_block        = "10.2.11.0/24"
  availability_zone = "ap-guangzhou-3"
}

resource "tencentcloudenterprise_cfs_file_system" "foo" {
  name = "test_cfs_file_system"
  availability_zone = "ap-guangzhou-3"
  access_group_id = local.cfs_access_group_id
  protocol = "NFS"
  vpc_id = cloud_vpc.vpc.id
  subnet_id = cloud_vpc_subnet.subnet.id
}

data "tencentcloudenterprise_cfs_file_systems" "file_systems" {
  file_system_id = cloud_cfs_file_system.foo.id
  name = cloud_cfs_file_system.foo.name
}
`
