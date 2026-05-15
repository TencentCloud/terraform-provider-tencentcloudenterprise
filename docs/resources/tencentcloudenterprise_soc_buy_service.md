---
subcategory: "SOC"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_soc_buy_service"
sidebar_current: "docs-tencentcloudenterprise-resource-soc_buy_service"
description: |-
  Provides a resource to create a soc buy_service
---

# tencentcloudenterprise_soc_buy_service

Provides a resource to create a soc buy_service

## Example Usage

```hcl
resource "tencentcloudenterprise_soc_buy_service" "buy_service" {
  type = "premium"
  tce_area {
    region_id = 50000001
    zone_id   = 50010001
  }
}
```

## Argument Reference

The following arguments are supported:

* `tce_area` - (Required, List, ForceNew) Region and zone info for service activation.
* `type` - (Optional, String, ForceNew) Billing model type, e.g. premium.

The `tce_area` object supports the following:

* `region_id` - (Required, Int) Region ID.
* `zone_id` - (Required, Int) Zone ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `buy_status` - Whether the SOC service is purchased. true: purchased; false: not purchased.
* `resource_id` - Resource ID returned after activation.

