---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_module_status"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_module_status"
description: |-
  Provides a resource to manage NGWAF module status for a domain.
---

# tencentcloudenterprise_ngwaf_module_status

Provides a resource to manage NGWAF module status for a domain.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_module_status" "example" {
  domain         = "example.com"
  web_security   = 1
  access_control = 1
  cc_protection  = 1
  api_protection = 0
  anti_tamper    = 0
  anti_leakage   = 0
}
```

## Argument Reference

The following arguments are supported:

* `access_control` - (Required, Int) ACL module status, 0:closed, 1:opened.
* `api_protection` - (Required, Int) API security module status, 0:closed, 1:opened.
* `cc_protection` - (Required, Int) CC module status, 0:closed, 1:opened.
* `domain` - (Required, String) Domain.
* `web_security` - (Required, Int) WEB security module status, 0:closed, 1:opened.
* `anti_leakage` - (Optional, Int) Anti leakage module status, 0:closed, 1:opened.
* `anti_tamper` - (Optional, Int) Anti tamper module status, 0:closed, 1:opened.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_ngwaf_module_status can be imported using the id, e.g.

```
NGWAF module status can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_module_status.example example.com
```
```

