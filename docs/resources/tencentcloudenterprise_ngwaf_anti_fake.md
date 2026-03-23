---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_anti_fake"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_anti_fake"
description: |-
  Provides a resource to create a NGWAF anti fake URL rule.
---

# tencentcloudenterprise_ngwaf_anti_fake

Provides a resource to create a NGWAF anti fake URL rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_anti_fake" "example" {
  domain = "example.com"
  name   = "anti-fake-rule"
  uri    = "/index.html"
  status = 1
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String) Domain.
* `name` - (Required, String) Rule Name.
* `uri` - (Required, String) Uri.
* `status` - (Optional, Int) Status. 0: Turn off rules and log switches, 1: Turn on the rule switch and Turn off the log switch; 2: Turn off the rule switch and turn on the log switch; 3: Turn on the rule switch and turn on the log switch.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `protocol` - Protocol.
* `rule_id` - Rule ID.

## Import

tencentcloudenterprise_ngwaf_anti_fake can be imported using the id, e.g.

```
NGWAF anti fake rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_anti_fake.example rule_id#example.com
```
```

