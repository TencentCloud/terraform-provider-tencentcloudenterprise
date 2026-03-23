---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_ciphers"
sidebar_current: "docs-tencentcloudenterprise-datasource-ngwaf_ciphers"
description: |-
  Use this data source to query available SSL/TLS ciphers supported by NGWAF.
---

# tencentcloudenterprise_ngwaf_ciphers

Use this data source to query available SSL/TLS ciphers supported by NGWAF.

## Example Usage

```hcl
data "tencentcloudenterprise_ngwaf_ciphers" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `ciphers` - Encryption Suite InformationNote: This field may return null, indicating that a valid value cannot be obtained.
  * `cipher_id` - Encryption Suite IDNote: This field may return null, indicating that a valid value cannot be obtained.
  * `cipher_name` - Encryption Suite NameNote: This field may return null, indicating that a valid value cannot be obtained.
  * `version_id` - TLS version IDNote: This field may return null, indicating that a valid value cannot be obtained.

