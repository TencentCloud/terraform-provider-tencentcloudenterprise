---
subcategory: "Cloud Virtual Machine(CVM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cvm_instance"
sidebar_current: "docs-tencentcloudenterprise-resources-cvm_instance"
description: |-
  Provides a CVM instance resource.
---

# tencentcloudenterprise_cvm_instance

Provides a CVM instance resource.

~> **NOTE:** You can launch an CVM instance for a VPC network via specifying parameter `vpc_id`. One instance can only belong to one VPC.

~> **NOTE:** At present, 'PREPAID' instance cannot be deleted and must wait it to be outdated and released automatically.

## Example Usage

```hcl
# variables.tf
variable "number" {
  default = "1"
}

variable "count_format" {
  default = "%02d"
}

variable "image_name_regex_centos" {
  default = "CentOS 7.4"
}

variable "terraform-instance-name" {
  default = "terraform-instance"
}

variable "instance_family" {
  default = "S2"
}

variable "cvm_password" {
  default = "Test@12345"
}

variable "internet_charge_type" {
  default = "TRAFFIC_POSTPAID_BY_HOUR"
}

variable "instance_charge_type" {
  default = "POSTPAID_BY_HOUR"
}

variable "internet_max_bandwidth_out" {
  default = 5
}

variable "disk_type" {
  default = "CLOUD_PREMIUM"
}

variable "disk_size" {
  default = "50"
}

variable "internet_service_provider" {
  default = "CTCC"
}

# main.tf
## Datasource Query
data "tencentcloudenterprise_cvm_instance_types" "instance_type" {
  filter {
    name = "instance-family"
    values = [var.instance_family]
  }
  exclude_sold_out = true
}

data "tencentcloudenterprise_cvm_images" "foo" {
  instance_type = data.tencentcloudenterprise_cvm_instance_types.instance_type.instance_types[0].instance_type
  image_type = ["PUBLIC_IMAGE"]
  image_name_regex = var.image_name_regex_centos
}

# Create VPC & Subnet 
resource "tencentcloudenterprise_vpc" "vpc" {
  cidr_block = "172.16.0.0/16"
  name       = var.terraform-instance-name
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  vpc_id            = tencentcloudenterprise_vpc.vpc.id
  cidr_block        = "172.16.0.0/24"
  availability_zone = data.tencentcloudenterprise_cvm_instance_types.instance_type.instance_types[0].availability_zone
  name              = "${var.terraform-instance-name}-subnet"
  is_multicast      = false
}

## Create security group & rule set
resource "tencentcloudenterprise_vpc_security_group" "group" {
  name        = var.terraform-instance-name
  description = "New security group"
  project_id  = 0
}

resource "tencentcloudenterprise_vpc_security_group_rule_set" "group_rule" {
  security_group_id = tencentcloudenterprise_vpc_security_group.group.id

  ingress {
    cidr_block  = "0.0.0.0/0"
    protocol    = "TCP"
    port        = "22"
    action      = "ACCEPT"
    description = "Allow SSH"
  }

  ingress {
    cidr_block  = "0.0.0.0/0"
    protocol    = "TCP"
    port        = "80"
    action      = "ACCEPT"
    description = "Allow HTTP"
  }

  ingress {
    cidr_block  = "0.0.0.0/0"
    protocol    = "TCP"
    port        = "443"
    action      = "ACCEPT"
    description = "Allow HTTPS"
  }

  ingress {
    cidr_block  = "0.0.0.0/0"
    protocol    = "ICMP"
    action      = "ACCEPT"
    description = "Allow ICMP"
  }

  egress {
    cidr_block  = "0.0.0.0/0"
    action      = "ACCEPT"
    description = "Allow all egress"
  }
}



## CVM Instance
resource "tencentcloudenterprise_cvm_instance" "instance" {
  instance_name   = "${var.terraform-instance-name}-instance-${format(var.count_format, count.index + 1)}"
  hostname        = "${var.terraform-instance-name}-host-${format(var.count_format, count.index + 1)}"
  image_id        = data.tencentcloudenterprise_cvm_images.foo.images[0].image_id
  instance_type   = data.tencentcloudenterprise_cvm_instance_types.instance_type.instance_types[0].instance_type
  count           = var.number
  availability_zone = data.tencentcloudenterprise_cvm_instance_types.instance_type.instance_types[0].availability_zone

  orderly_security_groups = [tencentcloudenterprise_vpc_security_group.group.id]
  vpc_id                  = tencentcloudenterprise_vpc.vpc.id
  subnet_id               = tencentcloudenterprise_vpc_subnet.subnet.id

  password = var.cvm_password

  instance_charge_type    = var.instance_charge_type
  system_disk_type        = var.disk_type
  system_disk_size        = 50

  tags = {
    type = "example"
    env  = "test"
  }
}

## Create EIP and attach to CVM(If public network configuration is applicable)
resource "tencentcloudenterprise_eip" "eip" {
  count                      = var.number
  name                       = "eip-${format(var.count_format, count.index + 1)}"
  internet_charge_type       = var.internet_charge_type
  internet_max_bandwidth_out = var.internet_max_bandwidth_out
  internet_service_provider = var.internet_service_provider
  type                       = "EIP"

  tags = {
    type = "example"
    env  = "test"
  }
}
resource "tencentcloudenterprise_eip_association" "eip_attachment" {
  count       = var.number
  eip_id      = tencentcloudenterprise_eip.eip[count.index].id
  instance_id = tencentcloudenterprise_cvm_instance.instance[count.index].id
}


## Create CBS data disk and attach to CVM(If applicable)
resource "tencentcloudenterprise_cbs_storage" "disk" {
  availability_zone = data.tencentcloudenterprise_cvm_instance_types.instance_type.instance_types[0].availability_zone
  storage_type      = var.disk_type
  storage_size      = var.disk_size
  storage_name      = "${var.terraform-instance-name}-disk-${format(var.count_format, count.index + 1)}"
  count             = var.number
}

resource "tencentcloudenterprise_cbs_storage_attachment" "instance_attachment" {
  count       = var.number
  storage_id  = tencentcloudenterprise_cbs_storage.disk.*.id[count.index]
  instance_id = tencentcloudenterprise_cvm_instance.instance.*.id[count.index]
}

# outputs.tf
output "hostname_list" {
  value = join(",", tencentcloudenterprise_cvm_instance.instance.*.instance_name)
}

output "cvm_ids" {
  value = join(",", tencentcloudenterprise_cvm_instance.instance.*.id)
}

output "cvm_public_ip" {
  value = join(",", tencentcloudenterprise_eip.eip.*.public_ip)
}

output "eip_ids" {
  value = join(",", tencentcloudenterprise_eip.eip.*.id)
}

output "cvm_private_ip" {
  value = join(",", tencentcloudenterprise_cvm_instance.instance.*.private_ip)
}

output "tags" {
  value = jsonencode(tencentcloudenterprise_cvm_instance.instance.*.tags)
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, String, ForceNew) The available zone for the CVM instance.
* `image_id` - (Required, String) The image to use for the instance. Changing `image_id` will cause the instance reset.
* `instance_type` - (Required, String) The type of the instance.
* `allocate_public_ip` - (Optional, Bool, ForceNew) Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
* `cam_role_name` - (Optional, String, ForceNew) CAM role name authorized to access.
* `cdh_host_id` - (Optional, String, ForceNew) Id of cdh instance. Note: it only works when instance_charge_type is set to `CDHPAID`.
* `cdh_instance_type` - (Optional, String) Type of instance created on cdh, the value of this parameter is in the format of CDH_XCXG based on the number of CPU cores and memory capacity. Note: it only works when instance_charge_type is set to `CDHPAID`.
* `data_disks` - (Optional, List, ForceNew) Settings for data disks.
* `disable_monitor_service` - (Optional, Bool) Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed. Modifying will cause the instance reset.
* `disable_security_service` - (Optional, Bool) Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed. Modifying will cause the instance reset.
* `force_delete` - (Optional, Bool) Indicate whether to force delete the instance. Default is `false`. If set true, the instance will be permanently deleted instead of being moved into the recycle bin. Note: only works for `PREPAID` instance.
* `hostname` - (Optional, String) The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
* `instance_charge_type_prepaid_period` - (Optional, Int) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
* `instance_charge_type_prepaid_renew_flag` - (Optional, String) Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
* `instance_charge_type` - (Optional, String) The charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR` and `CDHPAID`. The default is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR` and `CDHPAID`. `PREPAID` instance may not allow to delete before expired. `SPOTPAID` instance must set `spot_instance_type` and `spot_max_price` at the same time. `CDHPAID` instance must set `cdh_instance_type` and `cdh_host_id`.
* `instance_count` - (Optional, Int, **Deprecated**) It has been deprecated from version 1.59.18. Use built-in `count` instead. (Optional, Int, *Deprecated*) Please use `count` instead. The number of instances to be purchased. Value range:[1,100]; default value: 1.
* `instance_name` - (Optional, String) The name of the instance. The max length of instance_name is 60, and default value is `Terraform-CVM-Instance`.
* `internet_charge_type` - (Optional, String) Internet charge type of the instance, Valid values are `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`. This value takes NO Effect when changing and does not need to be set when `allocate_public_ip` is false.
* `internet_max_bandwidth_out` - (Optional, Int) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when `allocate_public_ip` is false.
* `keep_image_login` - (Optional, Bool) Whether to keep image login or not, default is `false`. When the image type is private or shared or imported, this parameter can be set `true`. Modifying will cause the instance reset.
* `key_ids` - (Optional, Set: [`String`]) The key pair to use for the instance, it looks like `skey-16jig7tx`. Modifying will cause the instance reset.
* `key_name` - (Optional, String, **Deprecated**) Please use `key_ids` instead. (Optional, String, *Deprecated*) Please use `key_ids` instead. The key pair to use for the instance, it looks like `skey-16jig7tx`. Modifying will cause the instance reset.
* `orderly_security_groups` - (Optional, List: [`String`]) A list of orderly security group IDs to associate with.
* `password` - (Optional, String) Password for the instance. In order for the new password to take effect, the instance will be restarted after the password change. Modifying will cause the instance reset.
* `placement_group_id` - (Optional, String, ForceNew) The ID of a placement group.
* `platform_project_id` - (Optional, String) The project the instance belongs to.
* `private_ip` - (Optional, String) The private IP to be assigned to this instance, must be in the provided subnet and available.
* `project_id` - (Optional, Int) The project the instance belongs to, default to 0.
* `resource_type` - (Optional, String) resource type, default instance.
* `running_flag` - (Optional, Bool) Set instance to running or stop. Default value is true, the instance will shutdown when this flag is false.
* `security_groups` - (Optional, Set: [`String`], **Deprecated**) It will be deprecated. Use `orderly_security_groups` instead. (Optional, Set, `Deprecated`) Please use `orderly_security_groups` instead. A list of security group IDs to associate with.
* `stopped_mode` - (Optional, String) Billing method of a pay-as-you-go instance after shutdown. Available values: `KEEP_CHARGING`,`STOP_CHARGING`. Default `KEEP_CHARGING`.
* `subnet_id` - (Optional, String) The ID of a VPC subnet. If you want to create instances in a VPC network, this parameter must be set.
* `system_disk_id` - (Optional, String) System disk snapshot ID used to initialize the system disk. When system disk type is `LOCAL_BASIC` and `LOCAL_SSD`, disk id is not supported.
* `system_disk_size` - (Optional, Int) Size of the system disk. unit is GB, Default is 50GB. If modified, the instance may force stop.
* `system_disk_type` - (Optional, String) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage, `CLOUD_BSSD`: Basic SSD. NOTE: If modified, the instance may force stop.
* `tags` - (Optional, Map) A mapping of tags to assign to the resource. For tag limits, please refer to [Use Limits](https://intl.cloud.com/document/product/651/13354).
* `user_data_raw` - (Optional, String, ForceNew) The user data to be injected into this instance, in plain text. Conflicts with `user_data`. Up to 16 KB after base64 encoded.
* `user_data` - (Optional, String, ForceNew) The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.
* `vpc_id` - (Optional, String) The ID of a VPC network. If you want to create instances in a VPC network, this parameter must be set.

The `data_disks` object supports the following:

* `data_disk_size` - (Required, Int) Size of the data disk, and unit is GB.
* `data_disk_type` - (Required, String, ForceNew) Data disk type. For more information about limits on different data disk types, see [Storage Overview](https://intl.cloud.com/document/product/213/4952). Valid values: LOCAL_BASIC: local disk, LOCAL_SSD: local SSD disk, LOCAL_NVME: local NVME disk, specified in the InstanceType, LOCAL_PRO: local HDD disk, specified in the InstanceType, CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD, CLOUD_BSSD: Balanced SSD.
* `data_disk_id` - (Optional, String) Data disk ID used to initialize the data disk. When data disk type is `LOCAL_BASIC` and `LOCAL_SSD`, disk id is not supported.
* `data_disk_snapshot_id` - (Optional, String, ForceNew) Snapshot ID of the data disk. The selected data disk snapshot size must be smaller than the data disk size.
* `delete_with_instance` - (Optional, Bool, ForceNew) Decides whether the disk is deleted with instance(only applied to `CLOUD_BASIC`, `CLOUD_SSD` and `CLOUD_PREMIUM` disk with `POSTPAID_BY_HOUR` instance), default is true.
* `throughput_performance` - (Optional, Int, ForceNew) Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time of the instance.
* `expired_time` - Expired time of the instance.
* `instance_id` - The ID of the instance.
* `instance_status` - Current status of the instance.
* `public_ip` - Public IP of the instance.


## Import

CVM instance can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cvm_instance.foo ins-2qol3a80
```

