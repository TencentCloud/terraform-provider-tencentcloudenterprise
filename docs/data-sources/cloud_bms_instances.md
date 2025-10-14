---
subcategory: "Bare Metal Server(BMS)"
layout: "cloud"
page_title: "TencentCloud: cloud_bms_instances"
sidebar_current: "docs-cloud-datasource-bms_instances"
description: |-
  Use this data source to query bms instances.
---

# cloud_bms_instances

Use this data source to query bms instances.

## Example Usage

```hcl
data "cloud_bms_instances" "bms" {
  result_output_file = "bms.json"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Optional, String) The ID of the instance.
* `instance_name` - (Optional, String) The name of the instance. The max length of instance_name is 60, and default value is `Terraform-bms-Instance`.
* `instance_state` - (Optional, String) The state of the instance.
* `operating_system_type` - (Optional, String) The operating system type of the instance.
* `private_ip_address` - (Optional, String) The private ip address of the instance.
* `result_output_file` - (Optional, String) The file path to output the result.
* `subnet_id` - (Optional, String) The subnet ID of the instance.
* `vpc_id` - (Optional, String) The VPC ID of the instance.
* `zone` - (Optional, String) The available zone for the bms instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `instance_list` - The list of instances.
  * `create_time` - Create time of the instance.
  * `instance_id` - The ID of the instance.
  * `instance_name` - The name of the instance.
  * `instance_status` - Current status of the instance.
  * `zone` - The available zone for the bms instance.


