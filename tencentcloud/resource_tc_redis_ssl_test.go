package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudRedisSslResource_basic -v
func TestAccTencentCloudRedisSslResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRedisSsl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_redis_ssl.ssl", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_ssl.ssl", "instance_id", defaultCrsInstanceId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_ssl.ssl", "ssl_config", "enabled"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_redis_ssl.ssl",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccRedisSslUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_redis_ssl.ssl", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_ssl.ssl", "instance_id", defaultCrsInstanceId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_redis_ssl.ssl", "ssl_config", "disabled"),
				),
			},
		},
	})
}

const testAccRedisSslVar = `
variable "instance_id" {
	default = "` + defaultCrsInstanceId + `"
}
`

const testAccRedisSsl = testAccRedisSslVar + `

resource "tencentcloudenterprise_redis_ssl" "ssl" {
	instance_id = var.instance_id
	ssl_config = "enabled"
  }

`

const testAccRedisSslUpdate = testAccRedisSslVar + `

resource "tencentcloudenterprise_redis_ssl" "ssl" {
	instance_id = var.instance_id
	ssl_config = "disabled"
  }

`
