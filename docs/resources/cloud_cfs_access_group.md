---
subcategory: "Cloud File Storage(CFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cfs_access_group"
sidebar_current: "docs-cloud-resources-cfs_access_group"
description: |-
  Provides a resource to create a CFS access group.
---

# cloud_cfs_access_group

Provides a resource to create a CFS access group.

## Example Usage

```hcl
resource "cloud_cfs_access_group" "foo" {
  name        = "test_access_group"
  description = "test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of the access group, and max length is 64.
* `description` - (Optional, String) Description of the access group, and max length is 255.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time of the access group.
* `p_group_id` - PGroupId of the access group.


## Import

CFS access group can be imported using the id, e.g.

```
$ terraform import cloud_cfs_access_group.foo pgroup-7nx89k7l
```

