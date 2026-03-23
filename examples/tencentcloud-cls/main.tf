resource "tencentcloudenterprise_cls_logset" "logset" {
  logset_name = "tf-topic-test"
  tags        = {
    "test" = "test"
  }
}

resource "tencentcloudenterprise_cls_topic" "topic" {
  auto_split           = true
  logset_id            = tencentcloudenterprise_cls_logset.logset.id
  max_split_partitions = 20
  partition_count      = 1
  period               = 10
  storage_type         = "hot"
  tags                 = {
    "test" = "test"
  }
  topic_name           = "tf-topic-test"
}

resource "tencentcloudenterprise_cls_machine_group" "group" {
  group_name        = "tf-basic-group"
  service_logging   = true
  auto_update       = true
  update_end_time   = "19:05:00"
  update_start_time = "17:05:00"

  machine_group_type {
    type   = "ip"
    values = [
      "192.168.1.1",
      "192.168.1.2",
    ]
  }
}