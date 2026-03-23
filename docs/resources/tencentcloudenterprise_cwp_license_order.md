---
subcategory: "Cloud Workload Protection Platform(CWP)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cwp_license_order"
sidebar_current: "docs-tencentcloudenterprise-resource-cwp_license_order"
description: |-
  Provides a CWP (Cloud Workload Protection) license order resource.
---

# tencentcloudenterprise_cwp_license_order

Provides a CWP (Cloud Workload Protection) license order resource.

~> **NOTE:** This resource allows you to create and manage CWP license orders with automatic scaling support.
~> **NOTE:** When modifying `license_num`, the system will automatically handle expansion (increase) or shrinkage (decrease) operations.
~> **NOTE:** `auto_bind_switch` and `auto_repurchase_switch` are only available during resource creation.

## Example Usage

```hcl
# Basic CWP license order creation
resource "tencentcloudenterprise_cwp_license_order" "example" {
  region_id    = 50000001 # Chongqing region
  alias        = "Production Environment CWP License"
  license_type = 5 # 5 = FLAGSHIP version, 0 = PRO version
  license_num  = 1 # Initial license quantity

  # Auto-bind and auto-repurchase settings (creation only)
  auto_bind_switch       = true
  auto_repurchase_switch = true

  project_id = 0
  tags = {
    Environment = "production"
    Team        = "security"
    CostCenter  = "IT"
  }
}

# License expansion example - modify license_num to expand
resource "tencentcloudenterprise_cwp_license_order" "expansion_example" {
  region_id    = 50000001
  alias        = "Expanded CWP License"
  license_type = 5
  license_num  = 5 # Expand from 1 to 5 licenses (system will add 4 more)

  tags = {
    Environment = "production"
    Action      = "expansion"
  }
}

# License shrinkage example - modify license_num to shrink
resource "tencentcloudenterprise_cwp_license_order" "shrink_example" {
  region_id    = 50000001
  alias        = "Shrunk CWP License"
  license_type = 5
  license_num  = 2 # Shrink from 5 to 2 licenses (system will remove 3)

  tags = {
    Environment = "production"
    Action      = "shrinkage"
  }
}

# Modify alias only (no license quantity change)
resource "tencentcloudenterprise_cwp_license_order" "modify_alias" {
  region_id    = 50000001
  alias        = "Updated Resource Alias" # Only change alias
  license_type = 5
  license_num  = 2 # Keep same quantity

  tags = {
    Environment = "production"
    Modified    = "true"
  }
}

# Import existing license order
# terraform import tencentcloudenterprise_cwp_license_order.example cwplic-2c263017#50000001
```

## Argument Reference

The following arguments are supported:

* `region_id` - (Required, Int) Purchase order region.
* `alias` - (Optional, String) Resource alias.
* `auto_bind_switch` - (Optional, Bool) Auto bind switch for license order. Only available during creation. Default is false.
* `auto_repurchase_switch` - (Optional, Bool) Auto repurchase switch for license order. Only available during creation. Default is false.
* `license_num` - (Optional, Int) License quantity. Initial quantity for creation, and can be modified for expansion or shrinkage.
* `license_type` - (Optional, Int) LicenseType, 0 CWP Pro - Pay as you go, 1 CWP Pro - Monthly subscription, 2 CWP Ultimate - Monthly subscription. Default is 0.
* `project_id` - (Optional, Int) Project ID. Default is 0.
* `tags` - (Optional, Map) Tags of the license order.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `license_id` - license id.
* `resource_id` - resource id.
* `used_license_cnt` - used license count.

