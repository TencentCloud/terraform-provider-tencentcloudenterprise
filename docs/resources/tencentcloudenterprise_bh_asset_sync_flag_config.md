---
subcategory: "BH"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bh_asset_sync_flag_config"
sidebar_current: "docs-tencentcloudenterprise-resource-bh_asset_sync_flag_config"
description: |-
  Provide a resource to create a BH asset sync flag config
---

# tencentcloudenterprise_bh_asset_sync_flag_config

Provide a resource to create a BH asset sync flag config

## Example Usage

```hcl
resource "tencentcloudenterprise_bh_asset_sync_flag_config" "example" {
  auto_sync = true
}
```

## Argument Reference

The following arguments are supported:

* `auto_sync` - (Required, Bool) Whether to enable asset auto-sync, false - disabled, true - enabled.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `role_granted` - Whether the role has been authorized, false - not authorized, true - authorized.

## Import

tencentcloudenterprise_bh_asset_sync_flag_config can be imported using the id, e.g.

```
BH asset sync flag config can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_bh_asset_sync_flag_config.example id
```
```

