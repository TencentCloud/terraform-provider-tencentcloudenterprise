---
subcategory: "Cloud Log Service(CLS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cls_index"
sidebar_current: "docs-cloud-resource-cls_index"
description: |-
  Provides a resource to create a cls index.
---

# cloud_cls_index

Provides a resource to create a cls index.

## Example Usage

```hcl
resource "cloud_cls_index" "complete_index" {

  topic_id       = "a22f98e2-7331-44af-adc3-e21abcac1ae3"
  coverage_field = "parse_error_field"

  rule {
    # 全文索引配置
    full_text {
      case_sensitive = true
      tokenizer      = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
      contain_z_h    = true
    }

    # 键值索引配置
    key_value {
      case_sensitive = true
      template_type  = ""

      key_values {
        key = "level"
        value {
          type        = "text"
          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
          sql_flag    = true
          contain_z_h = false
          alias       = "log_level"
        }
      }

      key_values {
        key = "timestamp"
        value {
          type     = "long"
          sql_flag = true
          alias    = "log_time"
        }
      }

      key_values {
        key = "message"
        value {
          type        = "text"
          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
          sql_flag    = true
          contain_z_h = true
          alias       = "log_message"
        }
      }
    }

    dynamic_index {
      status = true
    }

    # 元字段索引配置
    tag {
      case_sensitive = false

      key_values {
        key = "source"
        value {
          type        = "text"
          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
          sql_flag    = true
          contain_z_h = false
          alias       = "log_source"
        }
      }

      key_values {
        key = "environment"
        value {
          type     = "text"
          sql_flag = true
          alias    = "env"
        }
      }
    }
  }
  status                  = false
  include_internal_fields = true
  metadata_flag           = 1
}
```

## Argument Reference

The following arguments are supported:

* `rule` - (Required, List) Index rule. FullText, KeyValue, Tag parameters must input one valid parameter.
* `topic_id` - (Required, String, ForceNew) Log topic ID.
* `coverage_field` - (Optional, String) Custom log parsing exception storage field.
* `include_internal_fields` - (Optional, Bool) Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
* `metadata_flag` - (Optional, Int) Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields).
* `status` - (Optional, Bool) Whether to take effect. Default value: true.

The `dynamic_index` object supports the following:

* `status` - (Optional, Bool) Dynamic index configuration switch.

The `full_text` object supports the following:

* `case_sensitive` - (Required, Bool) Case sensitivity. true means case sensitive, false means case insensitive.
* `tokenizer` - (Required, String) Full-Text index delimiter. Each character in the string represents a delimiter.
* `contain_z_h` - (Optional, Bool) Whether Chinese characters are contained. true means contains Chinese, false means does not contain Chinese.

The `key_value` object supports the following:

* `case_sensitive` - (Required, Bool) Case sensitivity. true means case sensitive, false means case insensitive.
* `key_values` - (Optional, List) Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
* `template_type` - (Optional, String) Whether index enables dynamic template. If enabled, index will be configured based on reported key-value pairs, but all field types are text, case sensitive, analysis not supported, using default delimiters. Deprecated.

The `key_values` object supports the following:

* `key` - (Required, String) Field name for key-value or metafield index configuration. Only supports letters, numbers, underscores and -./@, and cannot start with underscore. Note: 1. Metafield (tag) Key does not need __TAG__. prefix, consistent with field Key when uploading logs, console will automatically add __TAG__. prefix for display. 2. Total number of Keys in key-value index (KeyValue) and metafield index (Tag) cannot exceed 300. 3. Key hierarchy cannot exceed 10 levels, e.g. a.b.c.d.e.f.g.h.j.k. 4. Cannot contain both json parent and child fields, e.g. a and a.b.
* `value` - (Required, List) Field index description information.

The `rule` object supports the following:

* `dynamic_index` - (Optional, List) Key-value index automatic configuration. Empty means this feature is not enabled. When enabled, fields in logs will be automatically added to key-value index, including new fields added later.
* `full_text` - (Optional, List) Full-Text index configuration. Empty means full-text index is not enabled.
* `key_value` - (Optional, List) Key-Value index configuration. Empty means key-value index is not enabled.
* `tag` - (Optional, List) Metafield index configuration. Empty means metafield index is not enabled.

The `tag` object supports the following:

* `case_sensitive` - (Required, Bool) Case sensitivity. true means case sensitive, false means case insensitive.
* `key_values` - (Required, List) Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.

The `value` object supports the following:

* `type` - (Required, String) Field type. Valid values: long, text, double.
* `alias` - (Optional, String) Index key-value alias.
* `contain_z_h` - (Optional, Bool) Whether Chinese characters are contained. true means contains Chinese, false means does not contain Chinese.
* `sql_flag` - (Optional, Bool) Whether the analysis feature is enabled for the field.
* `tokenizer` - (Optional, String) Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cls cos index can be imported using the id, e.g.

```
$ terraform import cloud_cls_index.index 0937e56f-4008-49d2-ad2d-69c52a9f11cc
```

