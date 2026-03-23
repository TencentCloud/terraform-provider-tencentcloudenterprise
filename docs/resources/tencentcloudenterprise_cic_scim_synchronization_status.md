---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_scim_synchronization_status"
sidebar_current: "docs-tencentcloudenterprise-resource-cic_scim_synchronization_status"
description: |-
  Provides a resource to manage identity center scim synchronization status
---

# tencentcloudenterprise_cic_scim_synchronization_status

Provides a resource to manage identity center scim synchronization status

## Example Usage

```hcl
resource "tencentcloudenterprise_cic_scim_synchronization_status" "cic_scim_synchronization_status" {
  zone_id                     = "z-xxxxxx"
  scim_synchronization_status = "Enabled"
}
```

## Argument Reference

The following arguments are supported:

* `scim_synchronization_status` - (Required, String) SCIM synchronization status. Enabled-enabled. Disabled-disables.
* `zone_id` - (Required, String, ForceNew) Space ID. z-prefix starts with 12 random digits/lowercase letters.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cic_scim_synchronization_status can be imported using the id, e.g.

```
organization cic_scim_synchronization_status can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_scim_synchronization_status.cic_scim_synchronization_status ${zone_id}
```
```

