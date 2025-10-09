---
subcategory: "Cloud Kafka(CKafka)"
layout: "cloud"
page_title: "TencentCloud: cloud_ckafka_topic"
sidebar_current: "docs-cloud-resource-ckafka_topic"
description: |-
  Use this resource to create ckafka topic.
---

# cloud_ckafka_topic

Use this resource to create ckafka topic.

## Example Usage

```hcl
resource "cloud_ckafka_topic" "foo" {
  instance_id                    = "ckafka-f9ife4zz"
  topic_name                     = "example"
  note                           = "topic note"
  replica_num                    = 2
  partition_num                  = 1
  enable_white_list              = true
  ip_white_list                  = ["ip1", "ip2"]
  clean_up_policy                = "delete"
  sync_replica_min_num           = 1
  unclean_leader_election_enable = false
  segment                        = 3600000
  retention                      = 60000
  max_message_bytes              = 0
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) Ckafka instance ID.
* `partition_num` - (Required, Int) The number of partition.
* `replica_num` - (Required, Int) Number of replicas. The value cannot exceed the number of brokers. Maximum value: 3.
* `topic_name` - (Required, String, ForceNew) Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).
* `clean_up_policy` - (Optional, String) Policy or mode to clear logs. Default value: delete. delete": logs are deleted based on the corresponding retention period. "compact": logs are compressed based on the corresponding keys. "compact, delete": logs are compressed based on the corresponding keys and deleted based on the corresponding retention period.
* `enable_white_list` - (Optional, Bool) Whether to open the ip whitelist, `true`: open, `false`: close.
* `ip_white_list` - (Optional, List: [`String`]) IP allowlist used for specifying a quota. You are required to specify this parameter when EnableWhiteList is set to true.
* `max_message_bytes` - (Optional, Int) Maximum size of messages in a topic. Unit: Byte. Maximum value: 12582912. The maximum value is equal to 12 MB.
* `note` - (Optional, String) The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).
* `retention` - (Optional, Int) Message retention period. Unit: ms. Minimum value: 60000.
* `segment` - (Optional, Int) Period after which a segment is rolled. Unit: ms. Minimum value: 3600000.
* `sync_replica_min_num` - (Optional, Int) Min number of sync replicas, Default is `1`.
* `unclean_leader_election_enable` - (Optional, Bool) Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time of the CKafka topic.
* `forward_cos_bucket` - Data backup cos bucket: the bucket address that is dumped to cos.
* `forward_interval` - Periodic frequency of data backup to cos.
* `forward_status` - Data backup cos status. Valid values: `0`, `1`. `1`: do not open data backup, `0`: open data backup.
* `segment_bytes` - Number of bytes rolled by shard.


## Import

ckafka topic can be imported using the instance_id#topic_name, e.g.

```
$ terraform import cloud_ckafka_topic.foo ckafka-f9ife4zz#example
```

