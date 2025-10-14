---
subcategory: "Bare Metal Server(BMS)"
layout: "cloud"
page_title: "TencentCloud: cloud_bms_instance"
sidebar_current: "docs-cloud-resources-bms_instance"
description: |-
  Provide a resource to create a bms instance
---

# cloud_bms_instance

Provide a resource to create a bms instance

## Example Usage

```hcl
resource "cloud_bms_instance" "cloud_bms_instance-wuua" {
  operating_system      = "your operating system name"
  operating_system_type = "Linux"
  password              = "your password"
  subnet_id             = "your subnet id"
  vpc_id                = "your vpc id"
  flavor_id             = "your flavorid"
  instance_name         = "your instance"
  ipv6_address          = true
  raid_type             = "NORAID"
  availability_zone     = "your zone name"
  security_service      = true
  whistle_service       = true
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, String, ForceNew) The available zone for the bms instance.
* `flavor_id` - (Required, String) The flavor of the instance.
* `instance_name` - (Required, String) The name of the instance. The max length of instance_name is 60, and default value is `Unnamed`.
* `operating_system_type` - (Required, String) The operating system type of the instance.
* `operating_system` - (Required, String) The operate system of bms instance.
* `password` - (Required, String) Password for the instance. The password can only be set when creating a new instance. It is not modifiable after the instance is created.
* `raid_type` - (Required, String) The type of raid.
* `subnet_id` - (Required, String) The ID of a VPC subnet.
* `vpc_id` - (Required, String) The ID of a VPC network.
* `allocate_public_ip` - (Optional, Bool, ForceNew) Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
* `hostname` - (Optional, String) The hostname of the instance. The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
* `internet_max_bandwidth_out` - (Optional, Int) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when `allocate_public_ip` is false. The max values is 1000.
* `internet_service_provider` - (Optional, String) The Internet Service Provider (ISP) associated with the instance.
* `ipv6_address` - (Optional, Bool) Whether to allocate an IPv6 address. Defaults to not allocating if not specified.
* `placement_group_id` - (Optional, String, ForceNew) The ID of a placement group.
* `private_ip` - (Optional, Set: [`String`]) The private IP to be assigned to this instance, must be in the provided subnet and available.
* `security_service` - (Optional, Bool) Enhance service for security, it is disable by default.
* `tags` - (Optional, Map) A mapping of tags to assign to the resource.
* `whistle_service` - (Optional, Bool) Enhance service for whistle, it is disable by default.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time of the instance.
* `instance_status` - Current status of the instance.


## Import

Placement group can be imported using the id, e.g.

```
$ terraform import cloud_bms_placement_group.foo ps-ilan8vjf
```

