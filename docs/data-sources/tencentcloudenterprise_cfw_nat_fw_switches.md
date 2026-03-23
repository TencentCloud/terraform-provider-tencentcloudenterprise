---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_nat_fw_switches"
sidebar_current: "docs-tencentcloudenterprise-datasource-cfw_nat_fw_switches"
description: |-
  Use this data source to query detailed information of cloud firewall (cfw) NAT firewall switches.
---

# tencentcloudenterprise_cfw_nat_fw_switches

Use this data source to query detailed information of cloud firewall (cfw) NAT firewall switches.

## Example Usage

```hcl
# Query all NAT firewall switches

data "tencentcloudenterprise_cfw_nat_fw_switches" "example" {
  nat_ins_id = "cfwnat-xxxxxxxx"
}

# Query NAT firewall switches with specific enable status

data "tencentcloudenterprise_cfw_nat_fw_switches" "enabled_only" {
  nat_ins_id = "cfwnat-xxxxxxxx"
  enable     = 1
}
```

## Argument Reference

The following arguments are supported:

* `enable` - (Optional, Int) Switch enable status, 1 open; 0 close.
* `nat_ins_id` - (Optional, String) Filter the NAT firewall instance to which the NAT firewall subnet switch belongs.
* `result_output_file` - (Optional, String) Used to save results.
* `status` - (Optional, Int, **Deprecated**) It has been deprecated from version 1.82.37. Please use `enable` instead. Switch status, 1 open; 0 close.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `data` - NAT border firewall switch list data.
  * `abnormal` - Whether the switch is abnormal, 0: normal, 1: abnormal.
  * `cvm_num` - Cvm Num.
  * `enable` - Effective status.
  * `id` - ID.
  * `nat_id` - NAT gatway Id.
  * `nat_ins_id` - NAT firewall instance Id.
  * `nat_ins_name` - NAT firewall instance name.
  * `nat_name` - NAT gatway name.
  * `region` - Region.
  * `route_id` - Route Id.
  * `route_name` - Route Name.
  * `status` - Switch status.
  * `subnet_cidr` - IPv4 CIDR.
  * `subnet_id` - Subnet Id.
  * `subnet_name` - Subnet Name.
  * `vpc_id` - Vpc Id.
  * `vpc_name` - Vpc Name.

