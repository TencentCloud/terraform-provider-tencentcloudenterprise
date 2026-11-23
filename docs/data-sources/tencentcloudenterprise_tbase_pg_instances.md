---
subcategory: "TDSQL PostgreSQL(Tbase)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tbase_pg_instances"
sidebar_current: "docs-tencentcloudenterprise-datasource-tbase_pg_instances"
description: |-
  Queries tbase pg instances.
---

# tencentcloudenterprise_tbase_pg_instances

Queries tbase pg instances.

## Example Usage

```hcl
data "tencentcloudenterprise_tbase_pg_instances" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `instance_ids` - (Optional, Set: [`String`]) instance ids, if not specified, all instances will be returned.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - A list of tbase pg instances. Each element contains the following attributes.

