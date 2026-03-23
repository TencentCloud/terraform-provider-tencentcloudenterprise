---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_export"
sidebar_current: "docs-tencentcloudenterprise-resource-cls_export"
description: |-
  Provides a resource to create a cls export
---

# tencentcloudenterprise_cls_export

Provides a resource to create a cls export

## Example Usage

```hcl
resource "tencentcloudenterprise_cls_export" "export" {
  topic_id  = "7e34a3a7-635e-4da8-9005-88106c1fde69"
  log_count = 2
  query     = "select count(*) as count"
  from      = 1607499107000
  to        = 1607499108000
  order     = "desc"
  format    = "json"
}
```

## Argument Reference

The following arguments are supported:

* `from` - (Required, Int, ForceNew) Log export start time, millisecond timestamp.
* `log_count` - (Required, Int, ForceNew) Number of logs to export, maximum value is 10 million.
* `query` - (Required, String, ForceNew) Log export search statement, SQL statements are not supported.
* `to` - (Required, Int, ForceNew) Log export end time, millisecond timestamp.
* `topic_id` - (Required, String, ForceNew) Log topic ID.
* `derived_fields` - (Optional, List: [`String`], ForceNew) Derived fields to export.
* `display_header` - (Optional, Bool, ForceNew) Whether to display header keys in CSV. Default is true.
* `escape_character` - (Optional, String, ForceNew) CSV escape character. Default is double quote.
* `fill_field` - (Optional, String, ForceNew) CSV fill field for missing values.
* `format` - (Optional, String, ForceNew) Log export data format. json, csv, default is json.
* `order` - (Optional, String, ForceNew) Log export time sorting. desc, asc, default is desc.
* `separator` - (Optional, String, ForceNew) CSV separator. Default is comma.
* `syntax_rule` - (Optional, Int, ForceNew) Syntax rule. 0: Lucene, 1: CQL, default is 0.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `export_id` - Log export ID.

## Import

tencentcloudenterprise_cls_export can be imported using the id, e.g.

```
cls export can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cls_export.export topic_id#export_id
```
```

