---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_subnet"
sidebar_current: "docs-tencentcloudenterprise-resources-vpc_subnet"
description: |-
  Provide a resource to create a VPC subnet.
---

# tencentcloudenterprise_vpc_subnet

Provide a resource to create a VPC subnet.

## Example Usage

```hcl
variable "availability_zone" {
  default = "ap-guangzhou-3"
}

resource "tencentcloudenterprise_vpc" "foo" {
  name       = "guagua-ci-temp-test"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  availability_zone = var.availability_zone
  name              = "guagua-ci-temp-test"
  vpc_id            = tencentcloudenterprise_vpc.foo.id
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, String, ForceNew) The availability zone within which the subnet should be created.
* `cidr_block` - (Required, String, ForceNew) A network address block of the subnet.
* `name` - (Required, String) The name of subnet to be created.
* `vpc_id` - (Required, String, ForceNew) ID of the VPC to be associated.
* `is_multicast` - (Optional, Bool) Indicates whether multicast is enabled. The default value is 'true'.
* `route_table_id` - (Optional, String) ID of a routing table to which the subnet should be associated.
* `subnet_type` - (Optional, Int) subnet type. 0: Normal VPC; 9: BMS. Default is 0.
* `tags` - (Optional, Map) Tags of the subnet.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `available_ip_count` - The number of available IPs.
* `create_time` - Creation time of subnet resource.
* `is_default` - Indicates whether it is the default VPC for this region.
* `subnet_id` - ID of the subnet.


## Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_subnet.test subnet_id
```

