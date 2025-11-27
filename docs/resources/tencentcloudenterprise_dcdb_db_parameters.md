---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dcdb_db_parameters"
sidebar_current: "docs-tencentcloudenterprise-resources-dcdb_db_parameters"
description: |-
  Provides a resource to create a dcdb db_parameters
---

# tencentcloudenterprise_dcdb_db_parameters

Provides a resource to create a dcdb db_parameters

## Example Usage

```hcl
resource "tencentcloudenterprise_dcdb_db_parameters" "db_parameters" {
  instance_id = "tdsqlshard-973xatu3"
  params {
    param = "audit_txsql_audit_mode"
    value = "OFF"
  }
  params {
    param = "audit_txsql_log_file_max_size"
    value = "536870912"
  }
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) The ID of instance.
* `params` - (Required, List) Parameter list, each element is a combination of Param and Value.

The `params` object supports the following:

* `param` - (Required, String) The name of parameter.
* `value` - (Required, String) The value of parameter.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

dcdb db_parameters can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_dcdb_db_parameters.db_parameters instanceId#paramName
```

