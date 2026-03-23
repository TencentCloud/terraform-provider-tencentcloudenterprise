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
data "tencentcloudenterprise_cls_machines" "machines" {
  group_id = "80278829-74ed-4bde-a4bc-e9837fd8e0ef"
}

# Output all machines information
output "machines_list" {
  value = data.tencentcloudenterprise_cls_machines.machines.machines
}

# Output specific machine details
output "machine_details" {
  value = [for machine in data.tencentcloudenterprise_cls_machines.machines.machines : {
    ip          = machine.ip
    status      = machine.status
    version     = machine.version
    instance_id = machine.instance_id
  }]
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required, String) Queried machine group ID.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `flag` - TKE flag. The default value is an empty string. An empty string indicates logs are not from TKE. label_k8s indicates logs are from TKE.
* `group_auto_update` - Whether the machine group has enabled automatic upgrade function. 0: Automatic upgrade not enabled; 1: Automatic upgrade enabled.
* `latest_agent_version` - The latest LogListener version available to the current user.
* `machines` - Machine status information group.
  * `auto_update` - If open auto update flag.
  * `err_code` - Code of update operation.
  * `err_msg` - Msg of update operation.
  * `instance_id` - Machine instance ID.
  * `ip` - Ip of machine.
  * `offline_time` - Offline time of machine.
  * `status` - Status of machine.
  * `update_status` - Machine update status.
  * `version` - Current machine version.
* `service_logging` - Whether service logs are enabled. true: Service logs are enabled. false: Service logs are not enabled.
* `update_end_time` - Scheduled end time for the auto-upgrade feature of the machine group.
* `update_start_time` - Scheduled start time for the automatic upgrade feature of the machine group.

