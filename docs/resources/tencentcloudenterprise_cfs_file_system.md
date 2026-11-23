---
subcategory: "Cloud File Storage(CFS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfs_file_system"
sidebar_current: "docs-tencentcloudenterprise-resource-cfs_file_system"
description: |-
  Manages a CFS file system.
---

# tencentcloudenterprise_cfs_file_system

Manages a CFS file system.

## Example Usage

```hcl
variable "availability_zone" { type = string }
variable "subnet_id" { type = string }
variable "vpc_id" { type = string }
variable "tag_id" { type = number }

resource "tencentcloudenterprise_cfs_file_system" "example" {
  availability_zone = var.availability_zone
  name              = "tf-example"
  subnet_id         = var.subnet_id
  vpc_id            = var.vpc_id
  tag_id            = var.tag_id
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, String, ForceNew) The available zone that the file system locates at.
* `name` - (Required, String) Name of a file system.
* `subnet_id` - (Required, String, ForceNew) ID of a subnet.
* `tag_id` - (Required, Int) ID of a CFS storage pool in the availability zone. Use DescribeCfsTags to select a pool offering the requested storage type.
* `vpc_id` - (Required, String, ForceNew) ID of a VPC network.
* `access_group_id` - (Optional, String) ID of a access group.
* `mount_ip` - (Optional, String, ForceNew) IP of mount point.
* `project_id` - (Optional, String, ForceNew) Project ID.
* `protocol` - (Optional, String, ForceNew) File service protocol. Valid values are `NFS` and `CIFS`. and the default is `NFS`.
* `snap_id` - (Optional, String) Snapshot ID.
* `storage_type` - (Optional, String, ForceNew) File service StorageType. Valid values are `SD` and `EP`. The default is `SD`.
* `tags` - (Optional, Map) Instance tags.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time of the file system.
* `fs_id` - Mount root-directory.

