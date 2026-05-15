---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_device_account"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_device_account"
description: |-
  Provide a resource to create a DASB device account
---

# tencentcloudenterprise_dasb_device_account

Provide a resource to create a DASB device account

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_device_account" "example" {
  device_id = 1
  account   = "root"
}
```

## Argument Reference

The following arguments are supported:

* `account` - (Required, String, ForceNew) Device account.
* `device_id` - (Required, Int, ForceNew) Device ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_dasb_device_account can be imported using the id, e.g.

```
DASB device account can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_device_account.example 12345
```
```

