package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClsConfigExtra_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClsConfigExtra,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_cls_config_extra.extra", "name", "helloworld"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cls_config_extra.extra",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccClsConfigExtra = `
resource "tencentcloudenterprise_cls_logset" "logset" {
  logset_name = "tf-config-extra-test"
  tags        = {
    "test" = "test"
  }
}

resource "tencentcloudenterprise_cls_topic" "topic" {
  auto_split           = true
  logset_id            = cloud_cls_logset.logset.id
  max_split_partitions = 20
  partition_count      = 1
  period               = 10
  storage_type         = "hot"
  tags                 = {
    "test" = "test"
  }
  topic_name = "tf-config-extra-test"
}

resource "tencentcloudenterprise_cls_machine_group" "group" {
  group_name        = "tf-config-extra-test"
  service_logging   = true
  auto_update       = true
  update_end_time   = "19:05:00"
  update_start_time = "17:05:00"

  machine_group_type {
    type   = "ip"
    values = [
      "203.0.113.101",
      "203.0.113.102",
    ]
  }
}

resource "tencentcloudenterprise_cls_config_extra" "extra" {
  name        = "helloworld"
  topic_id    = cloud_cls_topic.topic.id
  type        = "container_file"
  log_type    = "json_log"
  config_flag = "label_k8s"
  logset_id   = cloud_cls_logset.logset.id
  logset_name = cloud_cls_logset.logset.logset_name
  topic_name  = cloud_cls_topic.topic.topic_name
  container_file {
    container    = "nginx"
    file_pattern = "log"
    log_path     = "/nginx"
    namespace    = "default"
    workload {
      container = "nginx"
      kind      = "deployment"
      name      = "nginx"
      namespace = "default"
    }
  }
  group_id = cloud_cls_machine_group.group.id
}


`
