---
subcategory: "Cloud Load Balancer(CLB)"
layout: "cloud"
page_title: "TencentCloud: cloud_clb_redirection"
sidebar_current: "docs-cloud-resource-clb_redirection"
description: |-
  Provides a resource to create a CLB redirection.
---

# cloud_clb_redirection

Provides a resource to create a CLB redirection.

## Example Usage

### # Manual Rewrite

```hcl
resource "cloud_clb_redirection" "foo" {
  clb_id             = "lb-p7olt9e5"
  source_listener_id = "lbl-jc1dx6ju"
  target_listener_id = "lbl-asj1hzuo"
  source_rule_id     = "loc-ft8fmngv"
  target_rule_id     = "loc-4xxr2cy7"
}
```

### # Auto Rewrite

```hcl
resource "cloud_clb_redirection" "foo" {
  clb_id             = "lb-p7olt9e5"
  target_listener_id = "lbl-asj1hzuo"
  target_rule_id     = "loc-4xxr2cy7"
  is_auto_rewrite    = true
}
```

## Argument Reference

The following arguments are supported:

* `clb_id` - (Required, String, ForceNew) ID of CLB instance.
* `source_listener_id` - (Required, String, ForceNew) ID of source listener.
* `source_rule_id` - (Required, String, ForceNew) Rule ID of source listener.
* `target_listener_id` - (Required, String, ForceNew) ID of target listener.
* `target_rule_id` - (Required, String, ForceNew) Rule ID of target listener.
* `delete_all_auto_rewrite` - (Optional, Bool) Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
* `is_auto_rewrite` - (Optional, Bool, ForceNew) Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location will not be used, the target listener and location must be https protocol and port is 443.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

CLB redirection can be imported using the id, e.g.

```
$ terraform import cloud_clb_redirection.foo loc-ft8fmngv#loc-4xxr2cy7#lbl-jc1dx6ju#lbl-asj1hzuo#lb-p7olt9e5
```

