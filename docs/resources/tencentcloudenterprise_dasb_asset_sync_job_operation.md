---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_asset_sync_job_operation"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_asset_sync_job_operation"
description: |-
  Provide a resource to create a DASB asset sync job operation
---

# tencentcloudenterprise_dasb_asset_sync_job_operation

Provide a resource to create a DASB asset sync job operation

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_asset_sync_job_operation" "example" {
  category = 1
}
```

## Argument Reference

The following arguments are supported:

* `category` - (Required, Int, ForceNew) Synchronize asset categories, 1- Host assets, 2- Database assets.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


