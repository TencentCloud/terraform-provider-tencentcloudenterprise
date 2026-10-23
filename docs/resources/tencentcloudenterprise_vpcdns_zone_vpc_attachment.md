---
subcategory: "Virtual Private Cloud DNS(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_zone_vpc_attachment"
sidebar_current: "docs-tencentcloudenterprise-resource-vpcdns_zone_vpc_attachment"
description: |-
  Manage a single VPC association with a private DNS zone.
---

# tencentcloudenterprise_vpcdns_zone_vpc_attachment

Manage a single VPC association with a private DNS zone.

An existing zone can be referenced by its ID without recreating or importing the zone. Other VPC associations are preserved when this attachment is created or removed.

If the zone's VPC list is managed by another Terraform configuration, ensure that configuration does not overwrite independently managed attachments. When delegating associations to attachment resources, exclude `vpc_set` from updates in the zone configuration, for example with `lifecycle { ignore_changes = [vpc_set] }`.

## Example Usage

### Attach a new VPC to an existing private DNS zone:

```hcl
variable "existing_zone_id" {
  type        = string
  description = "ID of the existing private DNS zone."
}

variable "region" {
  type        = string
  description = "Region of the new VPC."
}

variable "vpc_name" {
  type        = string
  description = "Name of the new VPC."
}

variable "vpc_cidr" {
  type        = string
  description = "IPv4 CIDR block of the new VPC."
}

resource "tencentcloudenterprise_vpc" "new" {
  name       = var.vpc_name
  cidr_block = var.vpc_cidr
}

resource "tencentcloudenterprise_vpcdns_zone_vpc_attachment" "existing_zone" {
  zone_id = var.existing_zone_id

  vpc_set {
    uniq_vpc_id = tencentcloudenterprise_vpc.new.id
    region      = var.region
  }
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required, String, ForceNew) PrivateZone ID.
* `account_vpc_set` - (Optional, List, ForceNew) New add account vpc info.
* `vpc_set` - (Optional, List, ForceNew) New add vpc info.

The `account_vpc_set` object supports the following:

* `region` - (Required, String, ForceNew) Vpc region.
* `uin` - (Required, String, ForceNew) Vpc owner uin. To grant role authorization to this account.
* `uniq_vpc_id` - (Required, String, ForceNew) Uniq Vpc Id.

The `vpc_set` object supports the following:

* `region` - (Required, String, ForceNew) Vpc region.
* `uniq_vpc_id` - (Required, String, ForceNew) Uniq Vpc Id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_vpcdns_zone_vpc_attachment can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_vpcdns_zone_vpc_attachment.existing_zone 'zone-id#vpc-id'
```

