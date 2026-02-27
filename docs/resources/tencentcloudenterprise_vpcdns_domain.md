---
subcategory: "Virtual Private Cloud DNS(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_domain"
sidebar_current: "docs-tencentcloudenterprise-resource-vpcdns_domain"
description: |-
  Provide a resource to create a VPCDNS domain.
---

# tencentcloudenterprise_vpcdns_domain

Provide a resource to create a VPCDNS domain.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpcdns_domain" "foo" {
  domain             = "brucezylin.cc"
  dns_forward_status = "ENABLED"
  tags = {
    "createdBy" = "terraform3"
  }
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String) The domain of the VpcDns.
* `dns_forward_status` - (Optional, String) DNS forward status, valid values: ENABLED, DISABLED. Default value: DISABLED.
* `tags` - (Optional, Map) Tags of the VPC.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of VpcDns Domain.
* `domain_id` - id of vpcdns domain

## Import

tencentcloudenterprise_vpcdns_domain can be imported using the id, e.g.

```
Vpc subnet instance can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_domain.test domain_id
```
```

