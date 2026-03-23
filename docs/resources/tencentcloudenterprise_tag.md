---
subcategory: "Tag"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tag"
sidebar_current: "docs-tencentcloudenterprise-resource-tag"
description: |-
  Provides a resource to create a tag
---

# tencentcloudenterprise_tag

Provides a resource to create a tag

## Example Usage

```hcl
resource "tencentcloudenterprise_tag" "example" {
  tag_key   = "example-key"
  tag_value = "example-value"
}
```

## Argument Reference

The following arguments are supported:

* `tag_key` - (Required, String, ForceNew) Tag key.
* `tag_value` - (Required, String, ForceNew) Tag value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


