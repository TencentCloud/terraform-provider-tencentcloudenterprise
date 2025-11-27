---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dcdb_flush_binlog_operation"
sidebar_current: "docs-tencentcloudenterprise-resources-dcdb_flush_binlog_operation"
description: |-
  Provides a resource to create a dcdb flush_binlog_operation
---

# tencentcloudenterprise_dcdb_flush_binlog_operation

Provides a resource to create a dcdb flush_binlog_operation

## Example Usage

```hcl
resource "tencentcloudenterprise_dcdb_flush_binlog_operation" "flush_operation" {
  instance_id = local.dcdb_id
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) Instance ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



