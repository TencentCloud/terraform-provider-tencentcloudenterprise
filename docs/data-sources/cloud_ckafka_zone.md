---
subcategory: "Cloud Kafka(CKafka)"
layout: "cloud"
page_title: "TencentCloud: cloud_ckafka_zone"
sidebar_current: "docs-cloud-datasource-ckafka_zone"
description: |-
  Use this data source to query detailed information of ckafka zone
---

# cloud_ckafka_zone

Use this data source to query detailed information of ckafka zone

## Example Usage

```hcl
data "cloud_ckafka_zone" "ckafka_zone" {
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Query result complex object entity.
  * `cluster_info` - User exclusive cluster information.
    * `available_band_width` - The current available bandwidth of the cluster in MBs.
    * `available_disk_size` - The current available disk of the cluster, in GB.
    * `cluster_id` - ClusterId.
    * `cluster_name` - ClusterName.
    * `max_band_width` - Maximum cluster bandwidth in MBs.
    * `max_disk_size` - The largest disk in the cluster, in GB.
    * `zone_id` - Availability zone to which the cluster belongs, indicating the availability zone to which the cluster belongs.
    * `zone_ids` - The availability zone where the cluster node is located. If the cluster is a cross-availability zone cluster, it includes multiple availability zones where the cluster node is located.
  * `max_bandwidth` - Maximum purchased bandwidth in Mbs.
  * `max_buy_instance_num` - The maximum number of purchased instances.
  * `message_price` - Postpaid message unit price.
    * `real_total_cost` - Discount price.
    * `total_cost` - Original price.
  * `unit_price` - Postpaid unit price.
    * `real_total_cost` - Discount price.
    * `total_cost` - Original price.
  * `zone_list` - Zone list.
    * `app_id` - App id.
    * `exflag` - Extra flag.
    * `flag` - Flag.
    * `is_internal_app` - Internal APP.
    * `sold_out` - Json object, key is model, value true is sold out, false is not sold out.
    * `zone_id` - Zone id.
    * `zone_name` - Zone name.
    * `zone_status` - Zone status.


