---
subcategory: "Virtual Private Cloud Domain Name System(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_domains"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpcdns_domains"
description: |-
  Provide a resource to query VPCDNS domain.
---

# tencentcloudenterprise_vpcdns_domains

Provide a resource to query VPCDNS domain.

## Example Usage

```hcl
data "tencentcloudenterprise_vpcdns_domains" "foo" {
  domain = "brucezylin.cc"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String) The domain of the VpcDns.
* `result_output_file` - (Optional, String) The file path to output the result.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - List of domain.
  * `create_time` - Create time of domain.
  * `dns_forward_status` - Dns forward status of domain.
  * `domain_id` - ID of the domain.
  * `domain` - Content of domain.
  * `record_count` - Record count of domain.
  * `remark` - Remark of domain.
  * `update_time` - Update time of domain.


## Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_domain.test domain_id
```

