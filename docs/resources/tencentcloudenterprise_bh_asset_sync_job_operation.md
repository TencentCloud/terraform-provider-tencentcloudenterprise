---
subcategory: "BH"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bh_asset_sync_job_operation"
sidebar_current: "docs-tencentcloudenterprise-resource-bh_asset_sync_job_operation"
description: |-
  Provide a resource to create a BH asset sync job operation
---

# tencentcloudenterprise_bh_asset_sync_job_operation

Provide a resource to create a BH asset sync job operation

## Example Usage

```hcl
resource "tencentcloudenterprise_bh_asset_sync_job_operation" "example" {
  category = 1
}
```

## Argument Reference

The following arguments are supported:

* `category` - (Required, Int, ForceNew) Asset synchronization category. 1 - host assets, 2 - database assets, 3 - Container assets.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


