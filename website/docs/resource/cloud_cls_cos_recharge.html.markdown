---
subcategory: "Cloud Log Service(CLS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cls_cos_recharge"
sidebar_current: "docs-cloud-resource-cls_cos_recharge"
description: |-
  Provides a resource to create a cls cos_recharge
---

# cloud_cls_cos_recharge

Provides a resource to create a cls cos_recharge

~> **NOTE:** This resource can not be deleted if you run `terraform destroy`.

## Example Usage

```hcl
resource "cloud_cls_cos_recharge" "cos_recharge" {
  bucket        = "cos-lock-1308919341"
  bucket_region = "ap-guangzhou"
  log_type      = "minimalist_log"
  logset_id     = "dd426d1a-95bc-4bca-b8c2-baa169261812"
  name          = "cos_recharge_for_test"
  prefix        = "test"
  topic_id      = "7e34a3a7-635e-4da8-9005-88106c1fde69"

  extract_rule_info {
    backtracking            = 0
    is_gbk                  = 0
    json_standard           = 0
    keys                    = []
    metadata_type           = 0
    un_match_up_load_switch = false

    filter_key_regex {
      key   = "__CONTENT__"
      regex = "dasd"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `bucket_region` - (Required, String) Cos bucket region.
* `bucket` - (Required, String) Cos bucket.
* `log_type` - (Required, String) Log type.
* `logset_id` - (Required, String) Logset id.
* `name` - (Required, String) Recharge name.
* `prefix` - (Required, String) Cos file prefix.
* `topic_id` - (Required, String, ForceNew) Topic id.
* `compress` - (Optional, String) Supported gzip, lzop, snappy.
* `extract_rule_info` - (Optional, List) Extract rule info.

The `extract_rule_info` object supports the following:

* `address` - (Optional, String) Syslog address.
* `backtracking` - (Optional, Int) backtracking data volume in incremental acquisition mode.
* `begin_regex` - (Optional, String) begin line regex.
* `delimiter` - (Optional, String) Log delimiter.
* `filter_key_regex` - (Optional, List) Rules that need to filter logs.
* `is_gbk` - (Optional, Int) Gbk encoding.
* `json_standard` - (Optional, Int) Is standard json.
* `keys` - (Optional, Set) Key list.
* `log_regex` - (Optional, String) Log regex.
* `meta_tags` - (Optional, List) Metadata tag list.
* `metadata_type` - (Optional, Int) Metadata type.
* `parse_protocol` - (Optional, String) Parse protocol.
* `path_regex` - (Optional, String) Metadata path regex.
* `protocol` - (Optional, String) Syslog protocol.
* `time_format` - (Optional, String) Time format.
* `time_key` - (Optional, String) Time key.
* `un_match_log_key` - (Optional, String) Parsing failure log key.
* `un_match_up_load_switch` - (Optional, Bool) Whether to upload the parsing failure log.

The `filter_key_regex` object supports the following:

* `key` - (Required, String) Need filter log key.
* `regex` - (Required, String) Need filter log regex.

The `meta_tags` object supports the following:

* `key` - (Optional, String) Metadata key.
* `value` - (Optional, String) Metadata value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cls cos_recharge can be imported using the id, e.g.

```
terraform import cloud_cls_cos_recharge.cos_recharge topic_id#cos_recharge_id
```

