---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_security_group"
sidebar_current: "docs-cloud-datasource-vpc_security_group"
description: |-
  Use this data source to query detailed information of security group.
---

# cloud_vpc_security_group

Use this data source to query detailed information of security group.

~> **NOTE:** It has been deprecated and replaced by cloud_vpc_security_groups.

## Example Usage

```hcl
data "cloud_vpc_security_group" "sglab" {
  security_group_id = cloud_vpc_security_group.sglab.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, String) Name of the security group to be queried. Conflict with `security_group_id`.
* `security_group_id` - (Optional, String) ID of the security group to be queried. Conflict with `name`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `be_associate_count` - Number of security group binding resources.
* `create_time` - Creation time of security group.
* `description` - Description of the security group.
* `project_id` - Project ID of the security group.


