package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCkafkaUsersDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_PREPAY) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCkafkaUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudDataSourceCkafkaUser,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCkafkaUserExists("tencentcloudenterprise_ckafka_user.foo"),
					resource.TestCheckResourceAttrSet("data.cloud_ckafka_users.foo", "instance_id"),
					resource.TestCheckResourceAttr("data.cloud_ckafka_users.foo", "user_list.0.account_name", "test1"),
					resource.TestCheckResourceAttrSet("data.cloud_ckafka_users.foo", "user_list.0.create_time"),
					resource.TestCheckResourceAttrSet("data.cloud_ckafka_users.foo", "user_list.0.update_time"),
				),
			},
		},
	})
}

const testAccTencentCloudDataSourceCkafkaUser = defaultKafkaVariable + `
resource "tencentcloudenterprise_ckafka_user" "foo" {
  instance_id  = var.instance_id
  account_name = "test1"
  password     = "test1234"
}

data "tencentcloudenterprise_ckafka_users" "foo" {
	instance_id  = cloud_ckafka_user.foo.instance_id
	account_name = cloud_ckafka_user.foo.account_name
}
`
