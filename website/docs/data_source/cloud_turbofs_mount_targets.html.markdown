---
subcategory: "Parallel File Storage(TurboFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_mount_targets"
sidebar_current: "docs-cloud-datasource-turbofs_mount_targets"
description: |-
  Use this data source to query detailed information of turbofs mount_targets
---

# cloud_turbofs_mount_targets

Use this data source to query detailed information of turbofs mount_targets

## Example Usage

```hcl
data "cloud_turbofs_mount_targets" "mount_targets" {
  file_system_id = "cfs-iobiaxtj"
}
```

## Argument Reference

The following arguments are supported:

* `file_system_id` - (Required, String) File system ID.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `mount_targets` - Mount target details.
  * `ccn_id` - CCN instance ID used by Turbofs Turbo.
  * `cidr_block` - CCN IP range used by Turbofs Turbo.
  * `file_system_id` - File system ID.
  * `fs_id` - Mount root-directory.
  * `ip_address` - Mount target IP.
  * `life_cycle_state` - Mount target status.
  * `mount_target_id` - Mount target ID.
  * `network_interface` - Network type.
  * `subnet_id` - Subnet ID.
  * `subnet_name` - Subnet name.
  * `vpc_id` - VPC ID.
  * `vpc_name` - VPC name.


