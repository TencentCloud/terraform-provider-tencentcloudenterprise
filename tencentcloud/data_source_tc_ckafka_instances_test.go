package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCkafkaInstancesDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_PREPAY) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCkafkaUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudDataSourceCkafkaInstances,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloud_ckafka_instances.foo", "instance_list.0.instance_id", defaultKafkaInstanceId),
				),
			},
		},
	})
}

const testAccTencentCloudDataSourceCkafkaInstances = defaultKafkaVariable + `
data "tencentcloudenterprise_ckafka_instances" "foo" {
	instance_ids=[var.instance_id]
}
`
