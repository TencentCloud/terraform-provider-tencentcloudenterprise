---
subcategory: "Cloud Load Balancer(CLB)"
layout: "cloud"
page_title: "TencentCloud: cloud_clb_instance"
sidebar_current: "docs-cloud-resource-clb_instance"
description: |-
  Provides a resource to create a CLB instance.
---

# cloud_clb_instance

Provides a resource to create a CLB instance.

## Example Usage

### INTERNAL CLB

```hcl
resource "cloud_clb_instance" "internal_clb" {
  network_type = "INTERNAL"
  clb_name     = "myclb"
  project_id   = 0
  vpc_id       = "vpc-7007ll7q"
  subnet_id    = "subnet-12rastkr"

  tags = {
    test = "tf"
  }
}
```

### OPEN CLB

```hcl
resource "cloud_clb_instance" "open_clb" {
  network_type    = "OPEN"
  clb_name        = "myclb"
  project_id      = 0
  vpc_id          = "vpc-da7ffa61"
  security_groups = ["sg-o0ek7r93"]

  tags = {
    test = "tf"
  }
}
```

### Default enable

```hcl
resource "cloud_vpc_subnet" "subnet" {
  availability_zone = "ap-guangzhou-1"
  name              = "sdk-feature-test"
  vpc_id            = cloud_vpc.foo.id
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

resource "cloud_vpc_security_group" "sglab" {
  name        = "sg_o0ek7r93"
  description = "favourite sg"
  project_id  = 0
}

resource "cloud_vpc" "foo" {
  name       = "for-my-open-clb"
  cidr_block = "10.0.0.0/16"

  tags = {
    "test" = "mytest"
  }
}

resource "cloud_clb_instance" "open_clb" {
  network_type    = "OPEN"
  clb_name        = "my-open-clb"
  project_id      = 0
  vpc_id          = cloud_vpc.foo.id
  security_groups = [cloud_vpc_security_group.sglab.id]

  tags = {
    test = "open"
  }
}
```

### CREATE multiple instance

```hcl
resource "cloud_clb_instance" "open_clb1" {
  network_type   = "OPEN"
  clb_name       = "hello"
  master_zone_id = "ap-guangzhou-3"
}
```

## Argument Reference

The following arguments are supported:

* `clb_name` - (Required, String) Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.
* `network_type` - (Required, String, ForceNew) Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
* `address_ip_version` - (Optional, String) IP version, only applicable to open CLB. Valid values are `IPV4`, `IPV6` and `IPv6FullChain`, Default is IPV4.
* `internet_bandwidth_max_out` - (Optional, Int) Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2000]. Unit is Mbps.
* `internet_charge_type` - (Optional, String) Internet charge type, only applicable to open CLB. Valid values are only `TRAFFIC_POSTPAID_BY_HOUR`.
* `log_set_id` - (Optional, String) The id of log set.
* `log_topic_id` - (Optional, String) The id of log topic.
* `master_zone_id` - (Optional, String) Setting master zone id of cross available zone disaster recovery.
* `project_id` - (Optional, Int, ForceNew) ID of the project within the CLB instance, `0` - Default Project.
* `security_groups` - (Optional, List: [`String`]) Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
* `slave_zone_id` - (Optional, String) Setting slave zone id of cross available zone disaster recovery. this zone will undertake traffic when the master is down.
* `subnet_id` - (Optional, String, ForceNew) Subnet ID of the CLB. Effective only for CLB within the VPC.
* `tags` - (Optional, Map) The available tags within this CLB.
* `target_region_info_region` - (Optional, String) Region of the target region for cross-region CLB.
* `target_region_info_vpc_id` - (Optional, String) VPC ID of the target region for cross-region CLB.
* `vpc_id` - (Optional, String, ForceNew) VPC ID of the CLB.
* `zone_id` - (Optional, String) Available zone id, only applicable to open CLB.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `clb_vips` - The virtual service address table of the CLB.
* `instance_id` - CLB instance ID.
* `vip_isp` - Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).


## Import

CLB instance can be imported using the id, e.g.

```
$ terraform import cloud_clb_instance.foo lb-7a0t6zqb
```

