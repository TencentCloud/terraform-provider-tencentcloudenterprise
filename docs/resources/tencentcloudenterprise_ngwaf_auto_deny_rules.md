---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_auto_deny_rules"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_auto_deny_rules"
description: |-
  Provides a resource to create NGWAF auto deny rules.
---

# tencentcloudenterprise_ngwaf_auto_deny_rules

Provides a resource to create NGWAF auto deny rules.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_auto_deny_rules" "example" {
  domain              = "example.com"
  attack_threshold    = 10
  time_threshold      = 10
  deny_time_threshold = 30
}
```

## Argument Reference

The following arguments are supported:

* `attack_threshold` - (Required, Int, ForceNew) The threshold number of attacks that triggers IP autodeny, ranging from 2 to 100 times.
* `deny_time_threshold` - (Required, Int, ForceNew) The IP autodeny time after triggering the IP autodeny, ranging from 5 to 360 minutes.
* `domain` - (Required, String, ForceNew) Domain.
* `time_threshold` - (Required, Int, ForceNew) IP autodeny statistical time, ranging from 1-60 minutes.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_ngwaf_auto_deny_rules can be imported using the id, e.g.

```
NGWAF auto deny rules can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_auto_deny_rules.example example.com
```
```

