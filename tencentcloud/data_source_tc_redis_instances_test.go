package tencentcloud

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudRedisInstancesDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudRedisInstancesDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.#"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.name"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.zone"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.project_id"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.type"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.mem_size"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.status"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.ip"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.port"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.create_time"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis", "instance_list.0.charge_type"),

					resource.TestMatchResourceAttr("data.cloud_redis_instances.redis-tags", "instance_list.#", regexp.MustCompile(`^[1-9]\d*$`)),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis-tags", "instance_list.0.name"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis-tags", "instance_list.0.zone"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis-tags", "instance_list.0.project_id"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis-tags", "instance_list.0.type"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis-tags", "instance_list.0.mem_size"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis-tags", "instance_list.0.status"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis-tags", "instance_list.0.ip"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis-tags", "instance_list.0.port"),
					resource.TestCheckResourceAttrSet("data.cloud_redis_instances.redis-tags", "instance_list.0.create_time"),
					resource.TestCheckResourceAttr("data.cloud_redis_instances.redis-tags", "instance_list.0.tags.test", "test"),
				),
			},
		},
	})
}

func testAccTencentCloudRedisInstancesDataSourceConfig() string {
	return defaultVpcSubnets + `
resource "tencentcloudenterprise_redis_instance" "redis_instance_test" {
  availability_zone = "ap-guangzhou-3"
  type_id           = 2
  password          = "test12345789"
  mem_size          = 8192
  name              = "terraform_test"
  port              = 6379
  vpc_id            = local.vpc_id
  subnet_id         = local.subnet_id
  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_redis_instances" "redis" {
  zone       = "ap-guangzhou-3"
  search_key = cloud_redis_instance.redis_instance_test.id
}

data "tencentcloudenterprise_redis_instances" "redis-tags" {
  zone = "ap-guangzhou-3"
  tags = cloud_redis_instance.redis_instance_test.tags
}
`
}
