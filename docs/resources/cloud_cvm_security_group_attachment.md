---
subcategory: "Cloud Virtual Machine(CVM)"
layout: "cloud"
page_title: "TencentCloud: cloud_cvm_security_group_attachment"
sidebar_current: "docs-cloud-resources-cvm_security_group_attachment"
description: |-
  Provides a resource to create a cvm security_group_attachment
---

# cloud_cvm_security_group_attachment

Provides a resource to create a cvm security_group_attachment

## Example Usage

```hcl
resource "cloud_cvm_security_group_attachment" "security_group_attachment" {
  security_group_id = "sg-xxxxxxx"
  instance_id       = "ins-xxxxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) Instance id.
* `security_group_id` - (Required, String, ForceNew) Security group id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cvm security_group_attachment can be imported using the id, e.g.

```
terraform import cloud_cvm_security_group_attachment.security_group_attachment ${instance_id}#${security_group_id}
```

