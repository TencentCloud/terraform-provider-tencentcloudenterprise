---
subcategory: "TDSQL for MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_security_group_attachment"
sidebar_current: "docs-cloud-resource-dcdb_security_group_attachment"
description: |-
  Provides a resource to create a dcdb security_group_attachment
---

# cloud_dcdb_security_group_attachment

Provides a resource to create a dcdb security_group_attachment

## Example Usage

```hcl
resource "cloud_dcdb_security_group_attachment" "security_group_attachment" {
  security_group_id = "sg-9s7k6qgw"
  instance_id       = "tdsqlshard-973xatu3"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) attached instance id.
* `security_group_id` - (Required, String, ForceNew) security group id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

dcdb security_group_attachment can be imported using the id, e.g.
```
$ terraform import cloud_dcdb_security_group_attachment.security_group_attachment securityGroupAttachment_id
```

