---
subcategory: "Cloud Distributed SQL For PostGreSQL(TBASE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tbase_instances"
sidebar_current: "docs-tencentcloudenterprise-datasource-tbase_instances"
description: |-
  Use this data source to query detailed information of dcdb instances
---

# tencentcloudenterprise_tbase_instances

Use this data source to query detailed information of dcdb instances

## Example Usage

```hcl
data "tencentcloudenterprise_tbase_instances" "foo" {
  instance_ids       = ["tdpg-8xg0uenw", "tdpg-4uq7ufzw"]
  result_output_file = "foo.json"
}
```

## Argument Reference

The following arguments are supported:

* `instance_ids` - (Optional, Set: [`String`]) Instance ids, if not specified, all instances will be returned.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - A list of tbase instances. Each element contains the following attributes.


