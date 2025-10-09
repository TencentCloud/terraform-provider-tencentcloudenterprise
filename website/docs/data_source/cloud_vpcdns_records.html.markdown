---
subcategory: "Virtual Private Cloud Domain Name System(VPCDNS)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpcdns_records"
sidebar_current: "docs-cloud-datasource-vpcdns_records"
description: |-
  Provide a resource to query VPCDNS records.
---

# cloud_vpcdns_records

Provide a resource to query VPCDNS records.

## Example Usage

```hcl
data "cloud_vpcdns_records" "foo" {
  domain_id = "xxx"
}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Required, Int) The domain id of the VpcDns.
* `sub_domain` - (Required, String) The domain of the VpcDns.
* `result_output_file` - (Optional, String) The file path to output the result.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - List of records.
  * `create_time` - Create time of record.
  * `domain_id` - ID of the domain.
  * `enabled` - Is enabled, 0: disabled, 1: enabled.
  * `mx` - MX.
  * `record_id` - ID of the record.
  * `record_type` - Type of record.
  * `status` - Status of record.
  * `sub_domain` - Sub domain of the record.
  * `ttl` - TTL.
  * `update_time` - Update time of record.


## Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import cloud_vpcdns_domain.test domain_id
```

