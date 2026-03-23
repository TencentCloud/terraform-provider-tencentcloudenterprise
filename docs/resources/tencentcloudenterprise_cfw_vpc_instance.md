---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_vpc_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-cfw_vpc_instance"
description: |-
  Provides a resource to create a cloud firewall (cfw) vpc instance.
---

# tencentcloudenterprise_cfw_vpc_instance

Provides a resource to create a cloud firewall (cfw) vpc instance.

## Example Usage

### # If mode is 0

```hcl
resource "tencentcloudenterprise_cfw_vpc_instance" "example" {
  name = "tf_example"
  mode = 0

  vpc_fw_instances {
    name = "fw_ins_example"
    vpc_ids = [
      "vpc-9tk1icg3",
      "vpc-e8wcbn67"
    ]
    fw_deploy {
      width         = 200
      cross_a_zone  = 1
      deploy_region = "ap-beijing-region-jcctest-ops"
    }
  }

  switch_mode = 1
  fw_vpc_cidr = "auto"
}
```

### # If mode is 1

```hcl
resource "tencentcloudenterprise_cfw_vpc_instance" "example" {
  name = "tf_example"
  mode = 1

  vpc_fw_instances {
    name = "fw_ins_example"
    fw_deploy {
      deploy_region = "ap-beijing-region-jcctest-ops"
      width         = 200
      cross_a_zone  = 0
    }
  }

  ccn_id      = "ccn-peihfqo7"
  switch_mode = 1
  fw_vpc_cidr = "auto"
}
```

## Argument Reference

The following arguments are supported:

* `mode` - (Required, Int) Mode 0: private network mode; 1: CCN cloud networking mode.
* `name` - (Required, String) VPC firewall (group) name.
* `switch_mode` - (Required, Int) Switch mode of firewall instance. 1: Single point intercommunication; 2: Multi-point communication; 4: Custom Routing.
* `vpc_fw_instances` - (Required, List) List of firewall instances under firewall (group).
* `ccn_id` - (Optional, String) Cloud networking id, suitable for cloud networking mode.
* `fw_cidr_info` - (Optional, List) Specify the network segment information used by the firewall.
* `fw_vpc_cidr` - (Optional, String) auto Automatically select the firewall network segment; 10.10.10.0/24 The firewall network segment entered by the user.

The `fw_cidr_info` object supports the following:

* `fw_cidr_type` - (Required, String) The type of network segment used by the firewall. The values VpcSelf/Assis/Custom respectively represent own network segment priority/extended network segment priority/custom.
* `com_fw_cidr` - (Optional, String) Other firewalls occupy the network segment, which is usually the network segment specified when the firewall needs to exclusively occupy the vpc.
* `fw_cidr_lst` - (Optional, List) Specify the network segment of the firewall for each vpc.

The `fw_cidr_lst` object supports the following:

* `fw_cidr` - (Required, String) Firewall network segment, at least /24 network segment.
* `vpc_id` - (Required, String) Vpc id.

The `fw_deploy` object supports the following:

* `deploy_region` - (Required, String) Firewall Deployment Region.
* `width` - (Required, Int) Bandwidth, unit: Mbps.
* `cdc_id` - (Optional, String) When it is a CDC firewall, fill in this ID.
* `cross_a_zone` - (Optional, Int) Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if it is empty, off-site disaster recovery will not be used by default.
* `zone_bak` - (Optional, String, ForceNew) Backup availability zone, if empty, the default availability zone is selected.
* `zone` - (Optional, String, ForceNew) main zone, use default available zone if empty.

The `vpc_fw_instances` object supports the following:

* `fw_deploy` - (Required, List) Deploy regional information.
* `name` - (Required, String) Firewall instance name.
* `vpc_ids` - (Required, Set) List of VpcIds accessed in private network mode; only used in private network mode.
* `fw_ins_id` - (Optional, String) Firewall instance ID (passed in editing scenario).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `fw_group_id` - Firewall Group ID.

## Import

tencentcloudenterprise_cfw_vpc_instance can be imported using the id, e.g.

```
Cloud firewall vpc group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_vpc_instance.example cfwg-4ee69507
```

