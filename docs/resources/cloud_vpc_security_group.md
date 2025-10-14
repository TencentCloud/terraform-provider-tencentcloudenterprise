---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_security_group"
sidebar_current: "docs-cloud-resources-vpc_security_group"
description: |-
  Provides a resource to create security group.
---

# cloud_vpc_security_group

Provides a resource to create security group.

## Example Usage

```hcl
resource "cloud_vpc_security_group" "sglab" {
  name        = "mysg"
  description = "favourite sg"
  project_id  = 0
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of the security group to be queried.
* `description` - (Optional, String) Description of the security group.
* `project_id` - (Optional, Int, ForceNew) Project ID of the security group.
* `tags` - (Optional, Map) Tags of the security group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `security_group_id` - ID of the security group to be queried.


## Import

Security group can be imported using the id, e.g.

```

	$ terraform import cloud_vpc_security_group.sglab sg-ey3wmiz1

```

