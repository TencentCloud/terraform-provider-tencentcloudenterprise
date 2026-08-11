---
subcategory: "Virtual Private Cloud DNS(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_zone_record"
sidebar_current: "docs-tencentcloudenterprise-resource-vpcdns_zone_record"
description: |-
  Provide a resource to create a Private Dns Record.
---

# tencentcloudenterprise_vpcdns_zone_record

Provide a resource to create a Private Dns Record.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpcdns_zone_record" "foo" {
  zone_id      = "zone-rqndjnki"
  record_type  = "A"
  record_value = "192.168.1.2"
  sub_domain   = "www"
  ttl          = 300
  weight       = 1
  mx           = 0
  remark       = "test"
}
```

## Argument Reference

The following arguments are supported:

* `record_type` - (Required, String) Record type. Valid values: "A", "AAAA", "CNAME", "MX", "TXT", "PTR".
* `record_value` - (Required, String) Record value, such as IP: 192.168.10.2, CNAME: cname.qcloud.com, and MX: mail.qcloud.com..
* `sub_domain` - (Required, String) Subdomain, such as "www", "m", and "@".
* `zone_id` - (Required, String, ForceNew) Private domain ID.
* `mx` - (Optional, Int) MX priority, which is required when the record type is MX. Valid values: 5, 10, 15, 20, 30, 40, 50.
* `remark` - (Optional, String) Remarks.
* `status` - (Optional, String) Record status. Valid values: enabled, disabled.
* `ttl` - (Optional, Int) Record cache time. The smaller the value, the faster the record will take effect. Value range: 1~86400s.
* `weight` - (Optional, Int) Record weight. Value range: 1~100.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_vpcdns_zone_record can be imported using the id, e.g.

```
Private Dns Record can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_zone_record.foo zone_id#record_id
```
```

