---
subcategory: "Cloud Load Balancer(CLB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_clb_certificates"
sidebar_current: "docs-tencentcloudenterprise-datasource-clb_certificates"
description: |-
  Use this data source to query SSL certificate.
---

# tencentcloudenterprise_clb_certificates

Use this data source to query SSL certificate.

## Example Usage

```hcl
data "tencentcloudenterprise_clb_certificates" "foo" {
  result_output_file = "foo2.json"
  with_cert          = true
  cert_ids           = ["76cXqklW"]
}
```

## Argument Reference

The following arguments are supported:

* `cert_ids` - (Optional, Set: [`String`]) ID of the SSL certificate to be queried.
* `cert_type` - (Optional, String) Type of the SSL certificate to be queried. Available values includes: `CA` and `SVR`.
* `result_output_file` - (Optional, String) Used to save results.
* `with_cert` - (Optional, Bool) Whether to return the content of the SSL certificate.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `certificates` - An information list of certificate. Each element contains the following attributes:
  * `begin_time` - Beginning time of the SSL certificate.
  * `cert` - Content of the SSL certificate.
  * `create_time` - Creation time of the SSL certificate.
  * `domain` - Primary domain of the SSL certificate.
  * `end_time` - Ending time of the SSL certificate.
  * `id` - ID of the SSL certificate.
  * `key` - Key of the SSL certificate.
  * `name` - Name of the SSL certificate.
  * `product_zh_name` - Certificate authority.
  * `project_id` - Project ID of the SSL certificate.
  * `status` - Status of the SSL certificate.
  * `subject_names` - ALL domains included in the SSL certificate. Including the primary domain name.
  * `type` - Type of the SSL certificate.


