---
subcategory: "Cloud Connect Network(CCN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ccn"
sidebar_current: "docs-tencentcloudenterprise-resource-ccn"
description: |-
  Provides a resource to create a CCN instance.
---

# tencentcloudenterprise_ccn

Provides a resource to create a CCN instance.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn" "example" {
  name                 = "test-ccn"
  description          = "test ccn description"
  bandwidth_limit_type = "INTER_REGION_LIMIT"
  qos                  = "AU"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of the CCN to be queried, and maximum length does not exceed 60 bytes.
* `bandwidth_limit_type` - (Optional, String) The speed limit type. Valid values: `INTER_REGION_LIMIT`, `OUTER_REGION_LIMIT`. `OUTER_REGION_LIMIT` represents the regional export speed limit, `INTER_REGION_LIMIT` is the inter-regional speed limit. The default is `OUTER_REGION_LIMIT`.
* `charge_type` - (Optional, String, ForceNew) Billing mode. Valid values: `PREPAID`, `POSTPAID`. `PREPAID` means prepaid, which means annual and monthly subscription, `POSTPAID` means post-payment, which means billing by volume. The default is `POSTPAID`. The prepaid model only supports inter-regional speed limit, and the post-paid model supports inter-regional speed limit and regional export speed limit.
* `description` - (Optional, String) Description of CCN, and maximum length does not exceed 100 bytes.
* `qos` - (Optional, String, ForceNew) Service quality of CCN. Valid values: `PT`, `AU`, `AG`. The default is `AU`.
* `tags` - (Optional, Map) Instance tag.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of resource.
* `instance_count` - Number of attached instances.
* `state` - States of instance. Valid values: `ISOLATED`(arrears) and `AVAILABLE`.

## Import

tencentcloudenterprise_ccn can be imported using the id, e.g.

```
Ccn instance can be imported, e.g.

```
$ terraform import tencentcloudenterprise_ccn.example ccn-cwm743gl
```
```

