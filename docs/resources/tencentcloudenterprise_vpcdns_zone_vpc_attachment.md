---
subcategory: "Virtual Private Cloud DNS(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_zone_vpc_attachment"
sidebar_current: "docs-tencentcloudenterprise-resource-vpcdns_zone_vpc_attachment"
description: |-
  Provide a resource to bindvpc for a Private Dns Zone.
---

# tencentcloudenterprise_vpcdns_zone_vpc_attachment

Provide a resource to bindvpc for a Private Dns Zone.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpcdns_zone_vpc_attachment" "foo" {
  zone_id = tencentcloudenterprise_vpcdns_zone.zone.id
  vpc_set {
    uniq_vpc_id = "vpc-xxxxx"
    region      = "ap-guangzhou"
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
Private Dns Zone Vpc Attachment can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_zone_vpc_attachment.foo zone_id#vpc_id
```
```

