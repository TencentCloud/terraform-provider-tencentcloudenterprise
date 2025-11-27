---
subcategory: "Tencent Service Framework(TSF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tsf_usable_unit_namespaces"
sidebar_current: "docs-tencentcloudenterprise-datasource-tsf_usable_unit_namespaces"
description: |-
  Use this data source to query detailed information of tsf usable_unit_namespaces
---

# tencentcloudenterprise_tsf_usable_unit_namespaces

Use this data source to query detailed information of tsf usable_unit_namespaces

## Example Usage

```hcl
data "tencentcloudenterprise_tsf_usable_unit_namespaces" "usable_unit_namespaces" {
  search_word = ""
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.
* `search_word` - (Optional, String) Search by namespace id or namespace Name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Namespace object list.
  * `content` - Namespace list.
    * `id` - Unit namespace ID. Note: This field may return null, indicating that no valid value was found.
    * `namespace_id` - Namespace id.
    * `namespace_name` - Namespace name.
  * `total_count` - Total count.


