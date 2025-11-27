package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudTdmqRocketmqTopicDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTdmqRocketmqTopic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tdmq_rocketmq_topic.topic"),
					resource.TestCheckResourceAttr("data.cloud_tdmq_rocketmq_topic.topic", "topics.#", "1"),
				),
			},
		},
	})
}

const testAccDataSourceTdmqRocketmqTopic = `

resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
	cluster_name = "test_rocketmq_datasource_topic"
	remark = "test recket mq"
}

resource "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespace" {
	cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	namespace_name = "test_namespace_datasource_topic"
	ttl = 65000
	retention_time = 65000
	remark = "test namespace"
  }
  
  resource "tencentcloudenterprise_tdmq_rocketmq_topic" "topic" {
	topic_name = "test_rocketmq_topic"
	namespace_name = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
	type = "Normal"
	cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	remark = "test rocketmq topic"
  }
  
  data "tencentcloudenterprise_tdmq_rocketmq_topic" "topic" {
	cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	namespace_id = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
	filter_name = cloud_tdmq_rocketmq_topic.topic.topic_name
  }
`
