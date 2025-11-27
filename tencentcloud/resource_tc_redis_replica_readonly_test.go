package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudRedisReplicaReadonlyResource_basic -v
func TestAccTencentCloudRedisReplicaReadonlyResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRedisReplicaReadonly,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_redis_replica_readonly.replica_readonly", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_redis_replica_readonly.replica_readonly", "instance_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_replica_readonly.replica_readonly", "operate", "enable"),
				),
			},
			{
				Config: testAccRedisReplicaReadonlyUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_redis_replica_readonly.replica_readonly", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_redis_replica_readonly.replica_readonly", "instance_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_replica_readonly.replica_readonly", "operate", "disable"),
				),
			},
		},
	})
}

const testAccRedisReplicaReadonly = `

resource "tencentcloudenterprise_redis_replica_readonly" "replica_readonly" {
	instance_id = "crs-2yypjrnv"
	readonly_policy = ["master"]
	operate = "enable"
}

`

const testAccRedisReplicaReadonlyUpdate = `

resource "tencentcloudenterprise_redis_replica_readonly" "replica_readonly" {
	instance_id = "crs-2yypjrnv"
	readonly_policy = ["master"]
	operate = "disable"
}

`
