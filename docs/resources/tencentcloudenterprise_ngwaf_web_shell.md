---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_web_shell"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_web_shell"
description: |-
  Provides a resource to manage NGWAF web shell detection status for a domain.
---

# tencentcloudenterprise_ngwaf_web_shell

Provides a resource to manage NGWAF web shell detection status for a domain.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_web_shell" "example" {
  domain = "example.com"
  status = 1
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String) Domain.
* `status` - (Required, Int) Webshell status, 1: open; 0: closed; 2: log.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_ngwaf_web_shell can be imported using the id, e.g.

```
NGWAF web shell can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_web_shell.example example.com
```
```

