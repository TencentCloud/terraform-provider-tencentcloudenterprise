---
subcategory: "Cloud Kafka(ckafka)"
layout: "cloud"
page_title: "TencentCloud: cloud_ckafka_instances"
sidebar_current: "docs-cloud-datasource-ckafka_instances"
description: |-
  Use this data source to query detailed instance information of Ckafka
---

# cloud_ckafka_instances

Use this data source to query detailed instance information of Ckafka

## Example Usage

```hcl
data "cloud_ckafka_instances" "foo" {
  instance_ids = ["ckafka-vv7wpvae"]
}
```

## Argument Reference

The following arguments are supported:

* `filters` - (Optional, List) Filter. filter.name supports ('Ip', 'VpcId', 'SubNetId', 'InstanceType','InstanceId'), filter.values can pass up to 10 values.
* `instance_id` - (Optional, String) Filter by instance ID.
* `limit` - (Optional, Int) The number of pages, default is `10`.
* `offset` - (Optional, Int) The page start offset, default is `0`.
* `result_output_file` - (Optional, String) Used to save results.
* `search_word` - (Optional, String) Filter by instance name, support fuzzy query.
* `status` - (Optional, List: [`Int`]) (Filter Criteria) The status of the instance. 0: Create, 1: Run, 2: Delete, do not fill the default return all.
* `tag_key` - (Optional, String) Matches the tag key value.

The `filters` object supports the following:

* `name` - (Required, String) The field that needs to be filtered.
* `values` - (Required, List) The filtered value of the field.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `instance_list` - A list of ckafka users. Each element contains the following attributes:
  * `bandwidth` - Instance bandwidth, in Mbps.
  * `create_time` - The time when the instance was created.
  * `cvm` - ckafka sale type. Note: This field may return null, indicating that a valid value could not be retrieved.
  * `disk_size` - The storage size of the instance, in GB.
  * `disk_type` - Disk Type. Note: This field may return null, indicating that a valid value could not be retrieved.
  * `expire_time` - The instance expiration time.
  * `healthy_message` - Instance status information.
  * `healthy` - Instance status int: 1 indicates health, 2 indicates alarm, and 3 indicates abnormal instance status.
  * `instance_id` - The instance ID.
  * `instance_name` - The instance name.
  * `is_internal` - Whether it is an internal customer. A value of 1 indicates an internal customer.
  * `renew_flag` - Whether the instance is renewed, the int enumeration value: 1 indicates auto-renewal, and 2 indicates that it is not automatically renewed.
  * `status` - The status of the instance. 0: Created, 1: Running, 2: Delete: 5 Quarantined, -1 Creation failed.
  * `subnet_id` - Subnet id.
  * `tags` - Tag information.
    * `tag_key` - Tag Key.
    * `tag_value` - Tag Value.
  * `topic_num` - The number of topics.
  * `version` - Kafka version information. Note: This field may return null, indicating that a valid value could not be retrieved.
  * `vip_list` - Virtual IP entities.
    * `vip` - Virtual IP.
    * `vport` - Virtual PORT.
  * `vip` - Virtual IP.
  * `vpc_id` - VpcId, if empty, indicates that it is the underlying network.
  * `vport` - Virtual PORT.
  * `zone_id` - Availability Zone ID.
  * `zone_ids` - Across Availability Zones. Note: This field may return null, indicating that a valid value could not be retrieved.


