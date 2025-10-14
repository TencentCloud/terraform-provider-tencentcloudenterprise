---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_mariadb_flow"
sidebar_current: "docs-cloud-datasource-mariadb_flow"
description: |-
  Use this data source to query detailed information of mariadb flow
---

# cloud_mariadb_flow

Use this data source to query detailed information of mariadb flow

## Example Usage

```hcl
data "cloud_mariadb_flow" "flow" {
  flow_id = 1307
}
```

## Argument Reference

The following arguments are supported:

* `flow_id` - (Required, Int) Flow ID returned by async request API.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `status` - Flow status. 0: succeeded, 1: failed, 2: running.


