---
subcategory: "Parallel File Storage(TurboFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_file_system"
sidebar_current: "docs-cloud-resources-turbofs_file_system"
description: |-
  Provides a resource to create a parallel file system(TurboFS).
---

# cloud_turbofs_file_system

Provides a resource to create a parallel file system(TurboFS).

## Example Usage

```hcl
resource "cloud_turbofs_file_system" "example" {
  fs_name      = "tf-test"
  zone         = "az"
  p_group_id   = "pgroupbasic"
  storage_type = "TP"
  capacity     = 2560 # unit GiB
  vpc_id       = "vpc-cvukkbpd"
  subnet_id    = "subnet-kt2ffuim"
  pool_id      = "pool-wzFg3qtSu"
}
```

## Argument Reference

The following arguments are supported:

* `capacity` - (Required, Int) Capacity of the TURBO file system in GiB. The value should be an integer multiple of the spec policy.
* `fs_name` - (Required, String) Name of a file system.
* `p_group_id` - (Required, String) ID of a permission group.
* `storage_type` - (Required, String, ForceNew) File service StorageType. Valid values are `TP` and `TB`.
* `subnet_id` - (Required, String, ForceNew) ID of a subnet.
* `vpc_id` - (Required, String, ForceNew) ID of a VPC network.
* `zone` - (Required, String, ForceNew) The available zone that the file system locates at.
* `bandwidth_resource_pkg_id` - (Optional, String) ID of the bandwidth resource package bound to the file system. Each file system can only be bound to one bandwidth resource package.
* `cfs_id` - (Optional, String) Id of the file system.
* `encrypted` - (Optional, Bool) Whether the file system is encrypted.
* `kms_key_id` - (Optional, String) ID of the KMS key used for encryption. This field is required when `encrypted` is set to true.
* `mount_ip` - (Optional, String) IP of mount point.
* `pool_id` - (Optional, String, ForceNew) ID of the resource pool. If not specified, the system will automatically select the default resource pool.
* `project_id` - (Optional, String) ID of the project.
* `resource_tags` - (Optional, List) TurboFS resource tag.
* `snapshot_id` - (Optional, String) Snapshot ID.
* `storage_resource_pkg_id` - (Optional, String) ID of the storage resource package bound to the file system. Each file system can only be bound to one storage resource package.
* `tag_id` - (Optional, Int) TurboFS pool id.

The `resource_tags` object supports the following:

* `tag_key` - (Required, String) Tag key.
* `tag_value` - (Required, String) Tag value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `creation_time` - Create time of the file system.
* `file_system_id` - Id of the file system.
* `life_cycle_state` - Life cycle state of the file system.
* `size_byte` - Usage of the file system.


## Import

Cloud file system can be imported using the id, e.g.

```
$ terraform import cloud_turbofs_file_system.foo turbofs-6hgquxmj
```

