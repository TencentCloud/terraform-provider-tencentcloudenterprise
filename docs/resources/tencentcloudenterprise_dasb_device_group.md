---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_device_group"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_device_group"
description: |-
  Provide a resource to create a DASB device group
---

# tencentcloudenterprise_dasb_device_group

Provide a resource to create a DASB device group

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_device_group" "example" {
  name = "example-device-group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Device group name, the maximum length is 32 characters.
* `department_id` - (Optional, String) The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_dasb_device_group can be imported using the id, e.g.

```
DASB device group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_device_group.example 12345
```
```

