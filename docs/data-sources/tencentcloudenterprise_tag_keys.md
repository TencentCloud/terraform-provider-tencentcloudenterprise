---
subcategory: "Tag"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tag_keys"
sidebar_current: "docs-tencentcloudenterprise-datasource-tag_keys"
description: |-
  Use this data source to query tag keys
---

# tencentcloudenterprise_tag_keys

Use this data source to query tag keys

## Example Usage

```hcl
data "tencentcloudenterprise_tag_keys" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `tags` - Tag key list.

