---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_nat_instance"
sidebar_current: "docs-tencentcloudenterprise-resources-cfw_nat_instance"
description: |-
  Provides a resource to create a cloud firewall (cfw) nat instance.
---

# tencentcloudenterprise_cfw_nat_instance

Provides a resource to create a cloud firewall (cfw) nat instance.

## Example Usage

```hcl
resource "tencentcloudenterprise_cfw_nat_instance" "example" {
  name  = "tf_example"
  width = 20
  mode  = 0
  new_mode_items {
    vpc_list  = ["vpc-3h7d5le1"]
    eips      = []
    add_count = 1
  }
  cross_a_zone = 1
  fw_cidr_info {
    fw_cidr_type = "VpcSelf"
  }
}
```

## Argument Reference

The following arguments are supported:

* `mode` - (Required, Int) Mode 1: access mode; 0: new mode.
* `name` - (Required, String) Firewall instance name.
* `width` - (Required, Int) Bandwidth.
* `cross_a_zone` - (Optional, Int) Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if empty, the default is not to use off-site disaster recovery.
* `domain` - (Optional, String) Required if you want to create a domain name.
* `fw_cidr_info` - (Optional, List) Specify the network segment information used by the firewall.
* `nat_gw_list` - (Optional, Set: [`String`]) A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.
* `new_mode_items` - (Optional, List) New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.
* `zone_bak` - (Optional, String) backup zone, use default available zone if empty.
* `zone` - (Optional, String, ForceNew) main zone, use default available zone if empty.

The `fw_cidr_info` object supports the following:

* `fw_cidr_type` - (Required, String) The type of network segment used by the firewall. The values VpcSelf/Assis/Custom respectively represent own network segment priority/extended network segment priority/custom.
* `com_fw_cidr` - (Optional, String) Other firewalls occupy the network segment, which is usually the network segment specified when the firewall needs to exclusively occupy the vpc.
* `fw_cidr_lst` - (Optional, List) Specify the network segment of the firewall for each vpc.

The `fw_cidr_lst` object supports the following:

* `fw_cidr` - (Required, String) Firewall network segment, at least /24 network segment.
* `vpc_id` - (Required, String) Vpc id.

The `new_mode_items` object supports the following:

* `vpc_list` - (Required, Set) List of vpcs connected in new mode.
* `add_count` - (Optional, Int) 
* `eips` - (Optional, Set) List of egress elastic public network IPs bound in the new mode.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

Cloud firewall nat instance can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_nat_instance.example cfwnat-54a21421```

