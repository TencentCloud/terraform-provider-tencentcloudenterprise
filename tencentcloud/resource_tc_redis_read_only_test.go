package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudRedisReadOnlyResource_basic -v
func TestAccTencentCloudRedisReadOnlyResource_basic(t *testing.T) {
	// t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRedisReadOnly,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_redis_read_only.read_only", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_read_only.read_only", "instance_id", defaultCrsInstanceId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_read_only.read_only", "input_mode", "0"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_redis_read_only.read_only",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccRedisReadOnlyVar = `
variable "instance_id" {
	default = "` + defaultCrsInstanceId + `"
}
`

const testAccRedisReadOnly = testAccRedisReadOnlyVar + `

resource "tencentcloudenterprise_redis_read_only" "read_only" {
	instance_id = var.instance_id
	input_mode = "0"
}

`
