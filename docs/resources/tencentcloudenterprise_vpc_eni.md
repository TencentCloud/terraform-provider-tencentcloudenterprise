---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_eni"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_eni"
description: |-
  Provides a resource to create an ENI.
---

# tencentcloudenterprise_vpc_eni

Provides a resource to create an ENI.

## Example Usage

### Auto-assign IP (recommended):

```hcl
resource "tencentcloudenterprise_vpc_eni" "auto" {
  name        = "ci-test-eni"
  vpc_id      = tencentcloudenterprise_vpc.foo.id
  subnet_id   = tencentcloudenterprise_vpc_subnet.foo.id
  description = "eni with auto-assigned IP"
}
```

### Specify IP count:

```hcl
resource "tencentcloudenterprise_vpc_eni" "count" {
  name        = "ci-test-eni"
  vpc_id      = tencentcloudenterprise_vpc.foo.id
  subnet_id   = tencentcloudenterprise_vpc_subnet.foo.id
  description = "eni desc"
  ipv4_count  = 1
}
```

### Specify IPs manually:

```hcl
resource "tencentcloudenterprise_vpc_eni" "manual" {
  name        = "ci-test-eni"
  vpc_id      = tencentcloudenterprise_vpc.foo.id
  subnet_id   = tencentcloudenterprise_vpc_subnet.foo.id
  description = "eni desc"

  ipv4s {
    ip          = "10.0.0.10"
    primary     = true
    description = "primary IP"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of the ENI, maximum length 60.
* `subnet_id` - (Required, String, ForceNew) ID of the subnet within this vpc.
* `vpc_id` - (Required, String, ForceNew) ID of the vpc.
* `description` - (Optional, String) Description of the ENI, maximum length 60.
* `ipv4_count` - (Optional, Int) The number of intranet IPv4s. When it is greater than 1, there is only one primary intranet IP. The others are auxiliary intranet IPs, which conflict with `ipv4s`. If both `ipv4s` and `ipv4_count` are not specified, the cloud will automatically assign a primary IP.
* `ipv4s` - (Optional, Set) Applying for intranet IPv4s collection, conflict with `ipv4_count`. When there are multiple ipv4s, can only be one primary IP, and the maximum length of the array is 30. If both `ipv4s` and `ipv4_count` are not specified, the cloud will automatically assign a primary IP. Each element contains the following attributes:
* `security_groups` - (Optional, Set: [`String`]) A set of security group IDs.
* `tags` - (Optional, Map) Tags of the ENI.

The `ipv4s` object supports the following:

* `ip` - (Required, String) Intranet IP.
* `primary` - (Required, Bool) Indicates whether the IP is primary.
* `description` - (Optional, String) Description of the IP, maximum length 25.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of the ENI.
* `ipv4_info` - An information list of IPv4s. Each element contains the following attributes:
  * `description` - Description of the IP.
  * `ip` - Intranet IP.
  * `primary` - Indicates whether the IP is primary.
* `mac` - MAC address.
* `primary` - Indicates whether the IP is primary.
* `state` - State of the ENI.

## Import

tencentcloudenterprise_vpc_eni can be imported using the id, e.g.

```
ENI can be imported using the id, e.g.

```
  $ terraform import tencentcloudenterprise_vpc_eni.foo eni-qka182br
```
```

