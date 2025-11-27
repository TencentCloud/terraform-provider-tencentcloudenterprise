---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_environments"
sidebar_current: "docs-tencentcloudenterprise-datasource-tdmq_pulsar_environments"
description: |-
  Use this data source to query detailed information of tdmq pulsar environments
---

# tencentcloudenterprise_tdmq_pulsar_environments

Use this data source to query detailed information of tdmq pulsar environments

## Example Usage

```hcl
data "tencentcloudenterprise_tdmq_pulsar_environments" "example" {
  environment_id = "keep-ns"
  cluster_id     = "pulsar-9n95ax58b9vn"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) ID of the Pulsar cluster.
* `environment_id` - (Optional, String) Environment (namespace) name.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - List of environments.
  * `create_time` - Creation time.
  * `environment_id` - Environment (namespace) name.
  * `msg_ttl` - Expiration time of unconsumed messages, unit second, maximum 1296000 (15 days).
  * `namespace_id` - Namespace ID.
  * `namespace_name` - Namespace name.
  * `remark` - Remarks, within 128 characters.
  * `topic_num` - Number of topics.
  * `update_time` - Last modification time.


