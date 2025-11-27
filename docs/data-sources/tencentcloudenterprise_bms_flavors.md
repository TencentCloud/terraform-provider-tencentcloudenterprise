---
subcategory: "Bare Metal Server(BMS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bms_flavors"
sidebar_current: "docs-tencentcloudenterprise-datasource-bms_flavors"
description: |-
  Use this data source to query bms placement group.
---

# tencentcloudenterprise_bms_flavors

Use this data source to query bms placement group.

## Example Usage

```hcl
data "tencentcloudenterprise_bms_flavors" "flavors" {
  result_output_file = "flavors.json"
}
```

## Argument Reference

The following arguments are supported:

* `cpu_arch` - (Optional, String) The cpu_arch of the flavor, `X86` and `ARM` are optional.
* `flavor_ids` - (Optional, Set: [`String`]) The IDs of the flavors.
* `result_output_file` - (Optional, String) The file path to output the result.
* `zone` - (Optional, String) The zone of the flavor.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - The list of flavors.
  * `cpu_arch` - The cpu_arch of the flavor.
  * `cpu` - The cpu of the flavor.
  * `create_time` - The create time of the flavor.
  * `flavor_id` - The ID of the flavor.
  * `flavor_name` - The name of the flavor.
  * `flavor_type` - The type of the flavor.
  * `memory` - The memory of the flavor.
  * `raid_type` - The raid type of the flavor.
  * `system_disk` - The system disk of the flavor.
  * `zone` - The zone of the flavor.


