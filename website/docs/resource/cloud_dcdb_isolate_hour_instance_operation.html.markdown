---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_isolate_hour_instance_operation"
sidebar_current: "docs-cloud-resource-dcdb_isolate_hour_instance_operation"
description: |-
  Provides a resource to create a dcdb isolate_hour_instance_operation
---

# cloud_dcdb_isolate_hour_instance_operation

Provides a resource to create a dcdb isolate_hour_instance_operation

## Example Usage

```hcl
resource "cloud_dcdb_isolate_hour_instance_operation" "isolate_hour_instance_operation" {
  instance_ids = local.dcdb_id
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) Instance ID list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



