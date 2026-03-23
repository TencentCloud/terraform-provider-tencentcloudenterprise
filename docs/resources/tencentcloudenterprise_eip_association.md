---
subcategory: "Cloud Elastic IP(EIP)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_eip_association"
sidebar_current: "docs-tencentcloudenterprise-resource-eip_association"
description: |-
  Provides an eip resource associated with other resource like CVM, CLB, ENI, BMS and HAVIP.
---

# tencentcloudenterprise_eip_association

Provides an eip resource associated with other resource like CVM, CLB, ENI, BMS and HAVIP.

~> **NOTE:** Please DO NOT define `allocate_public_ip` in `tencentcloudenterprise_cvm_instance` resource when using `tencentcloudenterprise_eip_association`.

## Example Usage

### Bind EIP to CVM instance

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_eip_association" "eip_cvm" {
  eip_id      = tencentcloudenterprise_eip.example.id
  instance_id = "ins-xxxxxx"
}
```

### Bind EIP to CLB instance

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_eip_association" "eip_clb" {
  eip_id      = tencentcloudenterprise_eip.example.id
  instance_id = "lb-xxxxxx"
}
```

### Bind EIP to ENI with specific private IP

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_vpc_eni" "example" {
  name      = "eni-example"
  vpc_id    = "vpc-xxxxxx"
  subnet_id = "subnet-xxxxxx"
}

resource "tencentcloudenterprise_eip_association" "eip_eni" {
  eip_id               = tencentcloudenterprise_eip.example.id
  network_interface_id = tencentcloudenterprise_vpc_eni.example.id
  private_ip           = "10.0.1.22"
}
```

### Bind EIP to BMS (Blackstone Metal Server)

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_eip_association" "eip_bms" {
  eip_id      = tencentcloudenterprise_eip.example.id
  instance_id = "bms-xxxxxx"
  private_ip  = "10.24.0.6"
}
```

### Bind EIP to HAVIP (High Availability Virtual IP)

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_havip" "example" {
  name      = "havip-example"
  vpc_id    = "vpc-xxxxxx"
  subnet_id = "subnet-xxxxxx"
}

resource "tencentcloudenterprise_eip_association" "eip_havip" {
  eip_id   = tencentcloudenterprise_eip.example.id
  havip_id = tencentcloudenterprise_havip.example.id
}
```

## Argument Reference

The following arguments are supported:

* `eip_id` - (Required, String, ForceNew) The ID of EIP.
* `havip_id` - (Optional, String, ForceNew) The HAVip id going to bind with the EIP. This field is conflict with `instance_id` and `network_interface_id`.
* `instance_id` - (Optional, String, ForceNew) The CVM, CLB or BMS instance id going to bind with the EIP. This field is conflict with `network_interface_id` and `havip_id`.
* `network_interface_id` - (Optional, String, ForceNew) Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.
* `private_ip` - (Optional, String, ForceNew) Indicates a private IP belongs to the `network_interface_id`, or a specific private IP to bind when using with BMS instance. This field is conflict with `havip_id`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_eip_association can be imported using the id, e.g.

```
Eip association can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_eip_association.eip_cvm eip-41s6jwy4::ins-34jwj3
$ terraform import tencentcloudenterprise_eip_association.eip_eni eip-41s6jwy4::eni-34jwj3::10.0.1.22
$ terraform import tencentcloudenterprise_eip_association.eip_havip eip-41s6jwy4::havip-34jwj3
```
```

