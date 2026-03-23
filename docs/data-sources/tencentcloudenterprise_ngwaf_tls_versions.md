---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_tls_versions"
sidebar_current: "docs-tencentcloudenterprise-datasource-ngwaf_tls_versions"
description: |-
  Use this data source to query available TLS versions supported by NGWAF.
---

# tencentcloudenterprise_ngwaf_tls_versions

Use this data source to query available TLS versions supported by NGWAF.

## Example Usage

```hcl
data "tencentcloudenterprise_ngwaf_tls_versions" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `tls` - TLS key value.
  * `version_id` - TLS version IDNote: This field may return null, indicating that a valid value cannot be obtained.
  * `version_name` - Tls version nameNote: This field may return null, indicating that a valid value cannot be obtained.

