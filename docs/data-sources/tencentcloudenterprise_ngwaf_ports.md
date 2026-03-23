---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_ports"
sidebar_current: "docs-tencentcloudenterprise-datasource-ngwaf_ports"
description: |-
  Use this data source to query NGWAF ports configuration.
---

# tencentcloudenterprise_ngwaf_ports

Use this data source to query NGWAF ports configuration.

## Example Usage

```hcl
data "tencentcloudenterprise_ngwaf_ports" "example" {
  edition     = "clb-waf"
  instance_id = "waf-xxxxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `edition` - (Optional, String) Instance type, sparta-waf represents SAAS WAF, clb-waf represents CLB WAF.
* `instance_id` - (Optional, String) Instance unique ID.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `http_ports` - Http port list for instance.
* `https_ports` - Https port list for instance.

