---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpc"
description: |-
  Provides details about a specific VPC.
---

# tencentcloudenterprise_vpc

Provides details about a specific VPC.

This resource can prove useful when a module accepts a vpc id as an input variable and needs to, for example, determine the CIDR block of that VPC.

~> **NOTE:** It has been deprecated and replaced by tencentcloudenterprise_vpc_instances.

## Example Usage

```hcl
variable "vpc_id" {}

data "tencentcloudenterprise_vpc" "selected" {
  id = var.vpc_id
}

resource "tencentcloudenterprise_vpc_subnet" "main" {
  name              = "my test subnet"
  cidr_block        = cidrsubnet(data.tencentcloudenterprise_vpc.selected.cidr_block, 4, 1)
  availability_zone = "eu-frankfurt-1"
  vpc_id            = data.tencentcloudenterprise_vpc.selected.id
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional, String) The ID of the specific VPC to retrieve.
* `name` - (Optional, String) The name of the specific VPC to retrieve.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cidr_block` - The CIDR block of the VPC.
* `is_default` - Whether or not the default VPC.
* `is_multicast` - Whether or not the VPC has Multicast support.


