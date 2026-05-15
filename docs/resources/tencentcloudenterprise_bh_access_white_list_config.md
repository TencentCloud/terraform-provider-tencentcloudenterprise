---
subcategory: "BH"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bh_access_white_list_config"
sidebar_current: "docs-tencentcloudenterprise-resource-bh_access_white_list_config"
description: |-
  Provide a resource to create a BH access white list config
---

# tencentcloudenterprise_bh_access_white_list_config

Provide a resource to create a BH access white list config

## Example Usage

```hcl
resource "tencentcloudenterprise_bh_access_white_list_config" "example" {
  allow_any  = false
  allow_auto = true
}
```

## Argument Reference

The following arguments are supported:

* `allow_any` - (Optional, Bool) true: allow all source IPs; false: do not allow all source IPs.
* `allow_auto` - (Optional, Bool) true: allow automatically added IPs; false: do not allow automatically added IPs.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_bh_access_white_list_config can be imported using the id, e.g.

```
BH access white list config can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_bh_access_white_list_config.example id
```
```

