---
subcategory: "Cloud Load Balancer(CLB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_clb_listener_default_domain"
sidebar_current: "docs-tencentcloudenterprise-resource-clb_listener_default_domain"
description: |-
  Provides a resource to set CLB listener default domain.
---

# tencentcloudenterprise_clb_listener_default_domain

Provides a resource to set CLB listener default domain.

## Example Usage

```hcl
resource "tencentcloudenterprise_clb_listener_default_domain" "example" {
  clb_id      = "lb-7a0t6zqb"
  listener_id = "lbl-hh141sn9"
  domain      = "www.example.com"
}
```

## Argument Reference

The following arguments are supported:

* `clb_id` - (Required, String, ForceNew) ID of CLB instance.
* `domain` - (Required, String) Domain name of the listener rule. Single domain rules are passed to `domain`, and multi domain rules are passed to `domains`.
* `listener_id` - (Required, String, ForceNew) ID of CLB listener.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `rule_id` - ID of this CLB listener rule.

## Import

tencentcloudenterprise_clb_listener_default_domain can be imported using the id, e.g.

```
CLB listener default domain can be imported using the id (clb_id#listener_id), e.g.

```
$ terraform import tencentcloudenterprise_clb_listener_default_domain.example lb-7a0t6zqb#lbl-hh141sn9
```
```

