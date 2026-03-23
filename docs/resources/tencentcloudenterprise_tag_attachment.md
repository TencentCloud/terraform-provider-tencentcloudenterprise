---
subcategory: "Tag"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tag_attachment"
sidebar_current: "docs-tencentcloudenterprise-resource-tag_attachment"
description: |-
  Provides a resource to create a tag attachment
---

# tencentcloudenterprise_tag_attachment

Provides a resource to create a tag attachment

## Example Usage

```hcl
resource "tencentcloudenterprise_tag_attachment" "example" {
  tag_key   = "example-key"
  tag_value = "example-value"
  resource  = "qcs::cvm:ap-guangzhou:uin/123456789:instance/ins-xxxxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `resource` - (Required, String, ForceNew) Six-segment description of resources.
* `tag_key` - (Required, String, ForceNew) Tag key.
* `tag_value` - (Required, String, ForceNew) Tag value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


