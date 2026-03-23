---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_namespace"
sidebar_current: "docs-tencentcloudenterprise-resource-tcr_namespace"
description: |-
  Use this resource to create tcr namespace.
---

# tencentcloudenterprise_tcr_namespace

Use this resource to create tcr namespace.

## Example Usage

```hcl
resource "tencentcloudenterprise_tcr_namespace" "foo" {
  instance_id    = ""
  name           = "example"
  is_public      = true
  is_auto_scan   = true
  is_prevent_vul = true
  severity       = "medium"
  cve_whitelist_items {
    cve_id = "cve-xxxxx"
  }
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) ID of the TCR instance.
* `name` - (Required, String, ForceNew) Name of the TCR namespace. Valid length is [2~30]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`), and cannot start, end or continue with separators.
* `cve_whitelist_items` - (Optional, List) Vulnerability Whitelist.
* `is_auto_scan` - (Optional, Bool) Scanning level, `True` is automatic, `False` is manual. Default is `false`.
* `is_prevent_vul` - (Optional, Bool) Blocking switch, `True` is open, `False` is closed. Default is `false`.
* `is_public` - (Optional, Bool) Indicate that the namespace is public or not. Default is `false`.
* `severity` - (Optional, String) Block vulnerability level, currently only supports `low`, `medium`, `high`.

The `cve_whitelist_items` object supports the following:

* `cve_id` - (Optional, String) Vulnerability Whitelist ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_tcr_namespace can be imported using the id, e.g.

```
tcr namespace can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_tcr_namespace.foo cls-cda1iex1#namespace
```
```

