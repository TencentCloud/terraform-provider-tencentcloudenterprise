---
subcategory: "Cloud Kafka(ckafka)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ckafka_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-ckafka_instance"
description: |-
  Manages a CKafka instance.
---

# tencentcloudenterprise_ckafka_instance

Manages a CKafka instance.

## Example Usage

```hcl
variable "zone_id" { type = number }
variable "subnet_id" { type = string }
variable "vpc_id" { type = string }

resource "tencentcloudenterprise_ckafka_instance" "example" {
  instance_name = "tf-example"
  zone_id       = var.zone_id
  band_width    = 20
  disk_size     = 100
  subnet_id     = var.subnet_id
  vpc_id        = var.vpc_id
}
```

## Argument Reference

The following arguments are supported:

* `band_width` - (Required, Int) Instance bandwidth as listed by the CKafka type configuration.
* `disk_size` - (Required, Int) Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
* `instance_name` - (Required, String, ForceNew) Instance name.
* `subnet_id` - (Required, String) Subnet id.
* `vpc_id` - (Required, String) Vpc id.
* `zone_id` - (Required, Int) Available zone id.
* `cluster_id` - (Optional, Int, ForceNew) Cluster-ID represents  the cluster to which Ckafa belongs, associated with cluster_name only needs to select one of them.
* `instance_id` - (Optional, String) Ckafka instance ID.
* `msg_retention_time` - (Optional, Int) The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the default 0 is not filled, which means that the log retention time recovery policy is not enabled.
* `renew_flag` - (Optional, Int) Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear no automatic renewal (user setting).
* `tag_set` - (Optional, Map) Tag set of instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `vip` - Vip of instance.
* `vport` - Type of instance.

