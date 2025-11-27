---
subcategory: "Virtual Private Cloud Domain Name System(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_record"
sidebar_current: "docs-tencentcloudenterprise-resources-vpcdns_record"
description: |-
  Provide a resource to create a VPCDNS record.
---

# tencentcloudenterprise_vpcdns_record

Provide a resource to create a VPCDNS record.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpcdns_record" "foo" {
  domain_id   = "745"
  mx          = 0
  record_type = "A"
  sub_domain  = "www"
  value       = "192.168.1.3"
  weight      = "100"
}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Required, Int) Id of the VpcDns.
* `record_type` - (Required, String) Record type, valid values: A, CNAME, TXT, MX, SRV, AAAA, SPF. Default value: A.
* `sub_domain` - (Required, String) Sub domain.
* `value` - (Required, String) The value of the VpcDns Record.
* `mx` - (Optional, Int) The priority of the MX record. Value range: 1-50. Default value: 10.
* `weight` - (Optional, String) The weight of the SRV record. Value range: 1-100. Default value: 10.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `created_on` - The creation time of the VpcDns Record.
* `enabled` - Whether the VpcDns Record is enabled. Valid values: 0 (disabled), 1 (enabled).
* `extra` - The extra information of the VpcDns Record.
* `record_id` - The ID of the VpcDns Record.
* `status` - The status of the VpcDns Record. Valid values: 1, 0.
* `ttl` - The TTL of the VpcDns Record.
* `updated_on` - The last update time of the VpcDns Record.


## Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_record.test record_id
```

