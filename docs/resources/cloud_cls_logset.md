---
subcategory: "Cloud Log Service(CLS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cls_logset"
sidebar_current: "docs-cloud-resources-cls_logset"
description: |-
  Provides a resource to create a cls logset
---

# cloud_cls_logset

Provides a resource to create a cls logset

## Example Usage

```hcl
resource "cloud_cls_logset" "logset" {
  logset_name = "demo"
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `logset_name` - (Required, String) Logset name.
* `period` - (Required, Int, ForceNew) Lifecycle in days. Value range: 1-366.
* `logset_id` - (Optional, String, ForceNew) Logset ID. Format: {user-defined-part}-{user-appid}. The user-defined part only supports lowercase letters, numbers and '-', cannot start or end with '-', and the length is 3 to 40 characters. The user appid needs to be appended at the end.
* `tags` - (Optional, Map) List of tag descriptions. By specifying this parameter, tags can be bound to the corresponding logset. A maximum of 10 tag key-value pairs is supported, and each tag key must be unique to the resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cls logset can be imported using the id, e.g.
```
$ terraform import cloud_cls_logset.logset logset_id
```

