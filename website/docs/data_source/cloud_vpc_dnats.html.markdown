---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_dnats"
sidebar_current: "docs-cloud-datasource-vpc_dnats"
description: |-
  Use this data source to query detailed information of DNATs.
---

# cloud_vpc_dnats

Use this data source to query detailed information of DNATs.

## Example Usage

```hcl
# query by nat gateway id

data "cloud_vpc_dnats" "foo" {
  nat_id = "nat-xfaq1"
}

# query by vpc id

data "cloud_vpc_dnats" "foo" {
  vpc_id = "vpc-xfqag"
}

# query by elastic ip

data "cloud_vpc_dnats" "foo" {
  elastic_ip = "203.0.113.1"
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional, String) Description of the NAT forward.
* `elastic_ip` - (Optional, String) Network address of the EIP.
* `elastic_port` - (Optional, String) Port of the EIP.
* `nat_id` - (Optional, String) ID of the NAT gateway.
* `private_ip` - (Optional, String) Network address of the backend service.
* `private_port` - (Optional, String) Port of intranet.
* `result_output_file` - (Optional, String) Used to save results.
* `vpc_id` - (Optional, String) ID of the VPC.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `dnat_list` - Information list of the DNATs.
  * `elastic_ip` - Network address of the EIP.
  * `elastic_port` - Port of the EIP.
  * `nat_id` - ID of the NAT.
  * `private_ip` - Network address of the backend service.
  * `private_port` - Port of intranet.
  * `protocol` - Type of the network protocol. Valid values: `TCP` and `UDP`.
  * `vpc_id` - ID of the VPC.


