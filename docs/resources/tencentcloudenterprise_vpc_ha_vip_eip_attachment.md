---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_ha_vip_eip_attachment"
sidebar_current: "docs-tencentcloudenterprise-resources-vpc_ha_vip_eip_attachment"
description: |-
  Provides a resource to create a HA VIP EIP attachment.
---

# tencentcloudenterprise_vpc_ha_vip_eip_attachment

Provides a resource to create a HA VIP EIP attachment.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_ha_vip_eip_attachment" "foo" {
  havip_id   = "havip-kjqwe4ba"
  address_ip = "1.1.1.1"
}
```

## Argument Reference

The following arguments are supported:

* `address_ip` - (Required, String, ForceNew) Public address of the EIP.
* `havip_id` - (Required, String, ForceNew) ID of the attached HA VIP.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

HA VIP EIP attachment can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_vpc_ha_vip_eip_attachment.foo havip-kjqwe4ba#1.1.1.1
```

