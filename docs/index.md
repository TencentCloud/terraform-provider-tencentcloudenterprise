---
layout: "tencentcloud"
page_title: "Provider: tencentcloudenterprise"
sidebar_current: "docs-tencentcloudenterprise-index"
description: The TencentCloudEnterprise provider is used to interact with many resources supported by TencentCloudEnterprise. The provider needs to be configured with the proper credentials before it can be used.
---

# TencentCloudEnterprise Provider

The TencentCloudEnterprise (TCE) provider is used to interact with many resources supported by [TencentCloudEnterprise](https://www.tencentcloud.com/solutions/tce).  
This provider is implemented as `terraform-provider-tencentcloudenterprise` and requires proper credential configuration before use.

Use the navigation on the left to read about the available resources.
> ℹ️**Note**: This provider is still under development.

## Example Usage

``` terraform
terraform {
  required_providers {
    tencentcloudenterprise = {
      source = "tencentcloudenterprise/cloud"
    }
  }
}

# Configure the TCE Provider
provider "tencentcloud" {
  secret_id  = "my-secret-id"
  secret_key = "my-secret-key"
  region     = "ap-guangzhou"
}

# Get availability images
data "cloud_cvm_images" "default" {
  image_type = ["PUBLIC_IMAGE"]
  os_name    = "centos"
}

# Get availability instance types
data "cloud_cvm_instance_types" "default" {
  cpu_core_count = 1
  memory_size    = 1
}

# Create a web server
resource "cloud_cvm_instance" "web" {
  instance_name              = "web server"
  availability_zone          = "ap-guangzhou"
  image_id                   = data.cloud_cvm_images.default.images.0.image_id
  instance_type              = data.cloud_cvm_instance_types.default.instance_types.0.instance_type
  system_disk_type           = "CLOUD_PREMIUM"
  system_disk_size           = 50
  allocate_public_ip         = true
  internet_max_bandwidth_out = 20
  security_groups            = [cloud_vpc_security_group.default.id]
  count                      = 1
}

# Create security group
resource "cloud_vpc_security_group" "default" {
  name        = "web accessibility"
  description = "make it accessible for both production and stage ports"
}

# Create security group rule allow web request
resource "cloud_security_group_rule" "web" {
  security_group_id = cloud_vpc_security_group.default.id
  type              = "ingress"
  cidr_ip           = "0.0.0.0/0"
  ip_protocol       = "tcp"
  port_range        = "80,8080"
  policy            = "accept"
}

# Create security group rule allow ssh request
resource "cloud_security_group_rule" "ssh" {
  security_group_id = cloud_vpc_security_group.default.id
  type              = "ingress"
  cidr_ip           = "0.0.0.0/0"
  ip_protocol       = "tcp"
  port_range        = "22"
  policy            = "accept"
}
```

## Authentication

The TCE provider offers a flexible means of providing credentials for authentication.
The following methods are supported and explained below:


### Static credentials

> ⚠️**Warning:** Hard-coding credentials into any Terraform configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding an `secret_id` `secret_key` and `region` in-line in the tencentcloud provider block:

Usage:

```hcl
provider "tencentcloudenterprise" {
  secret_id  = "my-secret-id"
  secret_key = "my-secret-key"
  region     = "ap-guangzhou"
}
```

```hcl
provider "tencentcloudenterprise" {
  secret_id  = "my-secret-id"
  secret_key = "my-secret-key"
  region     = "ap-guangzhou"
}
```

### Environment variables

You can provide your credentials via `TENCENTCLOUD_SECRET_ID` and `TENCENTCLOUD_SECRET_KEY` environment variables,
representing your TencentCloud Secret Id and Secret Key respectively. `TENCENTCLOUD_REGION` is also used, if applicable:

```hcl
provider "tencentcloudenterprise" {

}
```

Usage:

```shell
$ export TENCENTCLOUD_SECRET_ID="my-secret-id"
$ export TENCENTCLOUD_SECRET_KEY="my-secret-key"
$ export TENCENTCLOUD_REGION="ap-guangzhou"
$ terraform plan
```

## Argument Reference

In addition to generic provider arguments (e.g. alias and version), the following arguments are supported in the TencentCloud provider block:

* `secret_id` - (Optional) This is the TencentCloud secret id. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_ID` environment variable.
* `secret_key` - (Optional) This is the TencentCloud secret key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_KEY` environment variable.
* `region` - (Optional) This is the TencentCloud region. It must be provided, but it can also be sourced from the `TENCENTCLOUD_REGION` environment variables.
* `protocol` - (Optional) The protocol of the API request. Valid values: `HTTP` and `HTTPS`.
* `domain` - (Optional) The root domain of the API request.