---
subcategory: "TDSQL PostgreSQL (Tbase)"
layout: "cloud"
page_title: "TencentCloud: cloud_tbase_pg_instances"
sidebar_current: "docs-cloud-datasource-tbase_pg_instances"
description: |-
  Use this data source to query detailed information of dcdb instances
---

# cloud_tbase_pg_instances

Use this data source to query detailed information of dcdb instances

## Example Usage

```hcl
data "cloud_tbase_pg_instances" "foo" {
  instance_ids       = ["tdpg-8xg0uenw", "tdpg-4uq7ufzw"]
  result_output_file = "foo.json"
}
```

## Argument Reference

The following arguments are supported:

* `instance_ids` - (Optional, Set: [`String`]) instance ids, if not specified, all instances will be returned.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - A list of tbase pg instances. Each element contains the following attributes.


