---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_vpc_fw_switches"
sidebar_current: "docs-tencentcloudenterprise-datasource-cfw_vpc_fw_switches"
description: |-
  Use this data source to query detailed information of cloud firewall (cfw) VPC firewall switches.
---

# tencentcloudenterprise_cfw_vpc_fw_switches

Use this data source to query detailed information of cloud firewall (cfw) VPC firewall switches.

## Example Usage

```hcl
data "tencentcloudenterprise_cfw_vpc_fw_switches" "example" {
  vpc_ins_id = "cfwins-xxxxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `vpc_ins_id` - (Required, String) Firewall instance id.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `switch_list` - Switch list.
  * `enable` - Switch status 0: off, 1: on.
  * `status` - Switch status 0: normal, 1: switching.
  * `switch_id` - Firewall switch ID.
  * `switch_mode` - switch mode.
  * `switch_name` - Firewall switch name.

