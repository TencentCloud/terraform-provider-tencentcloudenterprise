---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_path_rewrite"
sidebar_current: "docs-cloud-resources-tsf_path_rewrite"
description: |-
  Provides a resource to create a tsf path_rewrite
---

# cloud_tsf_path_rewrite

Provides a resource to create a tsf path_rewrite

## Example Usage

```hcl
resource "cloud_tsf_path_rewrite" "path_rewrite" {
  gateway_group_id = "group-a2j9zxpv"
  regex            = "/test"
  replacement      = "/tt"
  blocked          = "N"
  order            = 2
}
```

## Argument Reference

The following arguments are supported:

* `blocked` - (Required, String) Whether to shield the mapped path, Y: Yes N: No.
* `gateway_group_id` - (Required, String) Gateway deployment group ID.
* `order` - (Required, Int) Rule order, the smaller the higher the priority.
* `regex` - (Required, String) Regular expression.
* `replacement` - (Required, String) Content to replace.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `path_rewrite_id` - Path rewrite rule ID.


## Import

tsf path_rewrite can be imported using the id, e.g.

```
terraform import cloud_tsf_path_rewrite.path_rewrite rewrite-nygq33v2
```

