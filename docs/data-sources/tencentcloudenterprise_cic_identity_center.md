---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_identity_center"
sidebar_current: "docs-tencentcloudenterprise-datasource-cic_identity_center"
description: |-
  Use this data source to query detailed information of cic identity center
---

# tencentcloudenterprise_cic_identity_center

Use this data source to query detailed information of cic identity center

## Example Usage

```hcl
data "tencentcloudenterprise_cic_identity_center" "identity_center" {
}

output "zone_id" {
  value = data.tencentcloudenterprise_cic_identity_center.identity_center.zone_id
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `create_time` - Creation time.
* `scim_sync_status` - SCIM synchronization status. Enabled: enabled; Disabled: disabled.
* `service_status` - Service status. Disabled: not opened; Enabled: opened.
* `update_time` - Update time.
* `zone_id` - Space ID. It starts with the z- prefix, followed by 12 random digits/lowercase letters.
* `zone_name` - Space name, which must be globally unique. It contains lowercase letters, digits, and hyphens (-). It cannot start or end with a hyphen (-), and cannot have two consecutive hyphens (-). Length: 2-64 characters.

