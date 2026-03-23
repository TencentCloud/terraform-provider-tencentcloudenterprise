---
subcategory: "Cloud Load Balancer(CLB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_clb_log_set"
sidebar_current: "docs-tencentcloudenterprise-resource-clb_log_set"
description: |-
  Provides a resource to create an exclusive CLB Logset.
---

# tencentcloudenterprise_clb_log_set

Provides a resource to create an exclusive CLB Logset.

## Example Usage

```hcl
resource "tencentcloudenterprise_clb_log_set" "foo" {
  period = 7
}
```

## Argument Reference

The following arguments are supported:

* `period` - (Optional, Int, ForceNew) Logset retention period in days. Maximun value is `90`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Logset creation time.
* `name` - Logset name, which unique and fixed `clb_logset` among all CLS logsets.
* `topic_count` - Number of log topics in logset.

## Import

tencentcloudenterprise_clb_log_set can be imported using the id, e.g.

```
CLB log set can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_clb_logset.foo 4eb9e3a8-9c42-4b32-9ddf-e215e9c92764
```
```

