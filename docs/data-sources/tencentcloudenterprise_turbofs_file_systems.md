---
subcategory: "Parallel File Storage(TurboFS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_turbofs_file_systems"
sidebar_current: "docs-tencentcloudenterprise-datasource-turbofs_file_systems"
description: |-
  Use this data source to query the detail information of TurboFS file systems.
---

# tencentcloudenterprise_turbofs_file_systems

Use this data source to query the detail information of TurboFS file systems.

## Example Usage

```hcl
data "tencentcloudenterprise_turbofs_file_systems" "file_systems" {
  file_system_id = "turbofs-6hgquxmj"
}
```

## Argument Reference

The following arguments are supported:

* `file_system_id` - (Optional, String) A specified file system ID used to query.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `file_system_list` - An information list of cloud file system. Each element contains the following attributes:
  * `alloced_space` - IP address of the mount target.
  * `app_id` - App id of the file system.
  * `auto_scale_up_rule` - Auto scale up rule.
  * `auto_snapshot_policy_id` - Auto snapshot policy id.
  * `bandwidth_limit` - Limit of the file system bandwidth.
  * `bandwidth_resource_pkg` - The bandwidth resource package bound to the file system. Each file system can only be bound to one bandwidth resource package.
  * `capacity` - Capacity of the TURBO file system in TiB. The value should be an integer multiple of the spec policy.
  * `cfs_version` - TurboFS file system version.
  * `creation_time` - Creation datetime of the file system.
  * `creation_token` - User-defined file system name, with lower priority than fs_name.
  * `encrypted` - Whether the file system is encrypted.
  * `file_system_id` - ID of the file system.
  * `fs_name` - Name of a file system.
  * `ip_address` - Allocated space of the file system.
  * `kms_key_id` - ID of the KMS key used for encryption. This field is required when `encrypted` is set to true.
  * `life_cycle_state` - State of the file system.
  * `pool_id` - ID of the resource pool. If not specified, the system will automatically select the default resource pool.
  * `protocol` - File service protocol. Valid value, the default is `TURBO`.
  * `size_byte` - Usage of the file system.
  * `size_limit_max` - Max limit of the usage of the file system.
  * `size_limit` - Max usage of the file system.
  * `snap_id` - Snapshot ID.
  * `snap_status` - Snapshot status.
  * `spec_policy` - Spec policy.
  * `storage_resource_pkg` - The storage resource package bound to the file system. Each file system can only be bound to one storage resource package.
  * `storage_type` - File service StorageType. Valid values are `TP` and `TB`.
  * `tag_id` - TurboFS pool id.
  * `tag_name` - TurboFS pool name.
  * `tags` - Tags.
  * `tiering_detail` - Tiering detail.
    * `tiering_size_in_bytes` - tiering size in bytes.
  * `tiering_state` - Tiering state.
  * `zone_id` - The zone id that the file system locates at.
  * `zone` - The zone that the file system locates at.


