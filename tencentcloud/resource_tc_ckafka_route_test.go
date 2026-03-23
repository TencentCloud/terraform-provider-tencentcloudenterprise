package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCkafkaRouteResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCkafkaRoute_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_ckafka_route.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ckafka_route.example", "vip_type", "3"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ckafka_route.example", "instance_id", defaultKafkaInstanceId),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_ckafka_route.example",
				ImportState:       true,
				ImportStateVerify: true,
				// vpc_id and subnet_id are not read back from API (local SDK Route struct lacks these fields)
				ImportStateVerifyIgnore: []string{"vpc_id", "subnet_id", "auth_flag", "caller_appid"},
			},
		},
	})
}

const testAccCkafkaRoute_basic = `
resource "tencentcloudenterprise_ckafka_route" "example" {
  instance_id = "` + defaultKafkaInstanceId + `"
  vip_type    = 3
  vpc_id      = "` + defaultKafkaVpcId + `"
  subnet_id   = "` + defaultKafkaSubnetId + `"
  access_type = 0
}
`
