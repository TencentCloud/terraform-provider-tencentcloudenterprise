---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_resource"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_resource"
description: |-
  Provide a resource to create a DASB resource
---

# tencentcloudenterprise_dasb_resource

Provide a resource to create a DASB resource

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_resource" "example" {
  deploy_region    = "ap-guangzhou"
  vpc_id           = "vpc-example"
  subnet_id        = "subnet-example"
  resource_edition = "standard"
  resource_node    = 50
  auto_renew_flag  = 1
  deploy_zone      = "ap-guangzhou-1"
  cidr_block       = "10.0.0.0/24"
  vpc_cidr_block   = "10.0.0.0/16"
}
```

## Argument Reference

The following arguments are supported:

* `auto_renew_flag` - (Required, Int) Automatic renewal. 1 is auto renew flag, 0 is not.
* `cidr_block` - (Required, String) Subnet segments that require service activation.
* `deploy_region` - (Required, String) Deploy region.
* `deploy_zone` - (Required, String) Deploy zone.
* `resource_edition` - (Required, String) Resource type.Value:standard/pro.
* `resource_node` - (Required, Int) Number of resource nodes.
* `subnet_id` - (Required, String) Deploy resource subnetId.
* `vpc_cidr_block` - (Required, String) The network segment corresponding to the VPC that requires service activation.
* `vpc_id` - (Required, String) Deploy resource vpcId.
* `package_bandwidth` - (Optional, Int) Number of bandwidth expansion packets (4M), The set value is an integer multiple of 4.
* `time_span` - (Optional, Int) Billing time. This field is mandatory, with a minimum value of 1.
* `time_unit` - (Optional, String) Billing cycle, only support m: month. This field is mandatory, fill in m.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_dasb_resource can be imported using the id, e.g.

```
DASB resource can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_resource.example resource-example
```
```

