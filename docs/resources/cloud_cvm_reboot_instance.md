---
subcategory: "Cloud Virtual Machine(CVM)"
layout: "cloud"
page_title: "TencentCloud: cloud_cvm_reboot_instance"
sidebar_current: "docs-cloud-resources-cvm_reboot_instance"
description: |-
  Provides a resource to create a cvm reboot_instance
---

# cloud_cvm_reboot_instance

Provides a resource to create a cvm reboot_instance

## Example Usage

```hcl
resource "cloud_cvm_reboot_instance" "reboot_instance" {
  instance_id = "ins-xxxxx"
  stop_type   = "SOFT"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) Instance ID.
* `stop_type` - (Optional, String, ForceNew) Shutdown type. Valid values: `SOFT`: soft shutdown; `HARD`: hard shutdown; `SOFT_FIRST`: perform a soft shutdown first, and perform a hard shutdown if the soft shutdown fails. Default value: SOFT.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



