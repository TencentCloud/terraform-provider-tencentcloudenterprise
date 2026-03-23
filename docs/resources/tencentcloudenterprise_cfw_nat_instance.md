---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_nat_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-cfw_nat_instance"
description: |-
  Provides a resource to create a cloud firewall (cfw) NAT firewall instance.
---

# tencentcloudenterprise_cfw_nat_instance

Provides a resource to create a cloud firewall (cfw) NAT firewall instance.

## Example Usage

```hcl
resource "tencentcloudenterprise_cfw_nat_instance" "example" {
  name  = "cfw-nat-example"
  width = 20
  mode  = 0

  new_mode_items {
    vpc_list  = ["vpc-3skwc52h"]
    eips      = []
    add_count = 1
  }

  cross_a_zone = 0

  fw_cidr_info {
    fw_cidr_type = "VpcSelf"
  }
}
```

## Argument Reference

The following arguments are supported:

* `mode` - (Required, Int, ForceNew) Access mode, 0: new mode, 1: access mode.
* `name` - (Required, String) Firewall instance name.
* `width` - (Required, Int) Bandwidth.
* `cross_a_zone` - (Optional, Int) Cross-region disaster recovery 1: use cross-region disaster recovery; 0: do not use cross-region disaster recovery; if empty, cross-region disaster recovery is not used by default.
* `domain` - (Optional, String) Required if you want to create a domain name.
* `fw_cidr_info` - (Optional, List) Specify the network segment information used by the firewall.
* `nat_gw_list` - (Optional, List: [`String`]) A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.
* `new_mode_items` - (Optional, List) New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.
* `zone_bak` - (Optional, String) Backup availability zone, if empty, the default availability zone is selected.
* `zone` - (Optional, String, ForceNew) main zone, use default available zone if empty.

The `fw_cidr_info` object supports the following:

* `fw_cidr_type` - (Required, String) The type of network segment used by the firewall. The values VpcSelf/Assis/Custom respectively represent own network segment priority/extended network segment priority/custom.
* `com_fw_cidr` - (Optional, String) Other firewalls occupy the network segment, which is usually the network segment specified when the firewall needs to exclusively occupy the vpc.
* `fw_cidr_lst` - (Optional, List) Specify the network segment of the firewall for each vpc.

The `fw_cidr_lst` object supports the following:

* `fw_cidr` - (Required, String) Firewall network segment, at least /24 network segment.
* `vpc_id` - (Required, String) Vpc id.

The `new_mode_items` object supports the following:

* `add_count` - (Optional, Int) Number of EIPs to create. If eips is specified, this will be calculated from eips length.
* `eips` - (Optional, List) Elastic public IP list.
* `vpc_list` - (Optional, List) VPC list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `cfw_ins_id` - Nat firewall instance id.
* `status` - Instance status. 0: normal, 1: initializing.

## Import

tencentcloudenterprise_cfw_nat_instance can be imported using the id, e.g.

```
Cloud firewall nat instance can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_nat_instance.example cfwnat-xxxxxxxx
```
```

