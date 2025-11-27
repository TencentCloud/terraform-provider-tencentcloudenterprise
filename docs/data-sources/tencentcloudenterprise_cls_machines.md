---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_machines"
sidebar_current: "docs-tencentcloudenterprise-datasource-cls_machines"
description: |-
  Use this data source to query detailed information of cls machines
---

# tencentcloudenterprise_cls_machines

Use this data source to query detailed information of cls machines

## Example Usage

```hcl
resource "tencentcloudenterprise_cls_machine_group" "group" {
  group_name        = "tf-describe-mg-test"
  service_logging   = true
  auto_update       = true
  update_end_time   = "19:05:00"
  update_start_time = "17:05:00"

  machine_group_type {
    type = "ip"
    values = [
      "192.168.1.1",
      "192.168.1.2",
    ]
  }
}

data "tencentcloudenterprise_cls_machines" "machines" {
  group_id = tencentcloudenterprise_cls_machine_group.group.id
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required, String) Group id.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `machines` - Info of Machines.
  * `auto_update` - If open auto update flag.
  * `err_code` - Code of update operation.
  * `err_msg` - Msg of update operation.
  * `ip` - Ip of machine.
  * `offline_time` - Offline time of machine.
  * `status` - Status of machine.
  * `update_status` - Machine update status.
  * `version` - Current machine version.


