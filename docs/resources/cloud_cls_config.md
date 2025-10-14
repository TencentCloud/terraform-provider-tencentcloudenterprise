---
subcategory: "Cloud Log Service(CLS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cls_config"
sidebar_current: "docs-cloud-resources-cls_config"
description: |-
  Provides a resource to create a cls config
---

# cloud_cls_config

Provides a resource to create a cls config

## Example Usage

```hcl
resource "cloud_cls_config" "config" {
  name     = "config_hello"
  output   = "4d07fba0-b93e-4e0b-9a7f-d58542560bbb"
  path     = "/var/log/kubernetes"
  log_type = "json_log"
  extract_rule {
    filter_key_regex {
      key   = "key1"
      regex = "value1"
    }
    filter_key_regex {
      key   = "key2"
      regex = "value2"
    }
    un_match_up_load_switch = true
    un_match_log_key        = "config"
    backtracking            = -1
  }
  exclude_paths {
    type  = "Path"
    value = "/data"
  }
  exclude_paths {
    type  = "File"
    value = "/file"
  }

  #  user_define_rule = ""
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Collection configuration name.
* `output` - (Required, String) Log topic ID (TopicId) of collection configuration.
* `advanced_config` - (Optional, String) Advanced collection configuration. JSON string, Key/Value definitions: ClsAgentFileTimeout (timeout attribute), value range: integer >= 0, 0 means no timeout; ClsAgentMaxDepth (maximum directory depth), value range: integer >= 0; ClsAgentParseFailMerge (merge parse failed logs), value range: true or false. Example: {"ClsAgentFileTimeout":0,"ClsAgentMaxDepth":10,"ClsAgentParseFailMerge":true}. Console default placeholder: {"ClsAgentDefault":0}.
* `config_extra_id` - (Optional, String) config_extra table primary key ID.
* `config_flag` - (Optional, String) Collection configuration tag.
* `exclude_paths` - (Optional, List) Collection blacklist path list.
* `extract_rule` - (Optional, List) Extraction rule. If ExtractRule is set, LogType must be set.
* `log_format` - (Optional, String) Log formatting method.
* `log_type` - (Optional, String) Type of the log to be collected. Default value: minimalist_log. Supported types: json_log: JSON file log; delimiter_log: delimiter file log; minimalist_log: single-line full text file log; fullregex_log: single-line full regex file log; multiline_log: multi-line full text file log; multiline_fullregex_log: multi-line full regex file log; user_define_log: composite parsing (for multi-format nested logs); service_syslog: syslog collection; windows_event_log: Windows event log.
* `path` - (Optional, String) Log collection path, including filename. Supports multiple paths separated by English commas. Required for file collection.
* `source` - (Optional, Int) Collection configuration source. 0: default source, 1: TKE.
* `user_define_rule` - (Optional, String) User-defined collection rule, JSON format serialized string. Required when LogType is user_define_log.

The `event_log_rules` object supports the following:

* `event_channel` - (Required, String) Event channel. Supported values: Application, Security, Setup, System, ALL.
* `time_type` - (Required, Int) Time type. 1: user-defined, 2: current time.
* `event_ids` - (Optional, List) Event ID filter list. Optional, empty means no filtering. Supports positive filtering for single values (e.g.: 20) or ranges (e.g.: 0-20), also supports negative filtering for single values (e.g.: -20). Multiple filter items can be separated by commas. For example: 1-200,-100 means collecting event logs in the range 1-200 excluding 100.
* `time_stamp` - (Optional, Int) Time. When the user selects the custom time type, a specific time needs to be specified.

The `exclude_paths` object supports the following:

* `type` - (Optional, String) Type. Valid values: File, Path.
* `value` - (Optional, String) Specific content corresponding to Type.

The `extract_rule` object supports the following:

* `address` - (Optional, String) Syslog system log collection specifies the address and port that the collector listens to.
* `backtracking` - (Optional, Int) Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection).
* `begin_regex` - (Optional, String) First-Line matching rule, which is valid only if log_type is multiline_log or fullregex_log.
* `delimiter` - (Optional, String) Delimiter for delimited log, which is valid only if log_type is delimiter_log.
* `event_log_rules` - (Optional, List) Windows event log collection rules. Only effective when LogType is windows_event_log, no need to fill for other types.
* `filter_key_regex` - (Optional, List) Log keys to be filtered and the corresponding regex.
* `is_gbk` - (Optional, Int) GBK encoding. Default 0.
* `json_standard` - (Optional, Int) Standard json. Default 0.
* `keys` - (Optional, List) Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if log_type is delimiter_log. json_log logs use the key of JSON itself.
* `log_regex` - (Optional, String) Full log matching rule, which is valid only if log_type is fullregex_log.
* `meta_tags` - (Optional, List) List of metadata tags for the machine group.
* `metadata_type` - (Optional, Int) Metadata type.
* `parse_protocol` - (Optional, String) Parse protocol.
* `path_regex` - (Optional, String) Metadata path regex.
* `protocol` - (Optional, String) Syslog protocol, tcp or udp.
* `time_format` - (Optional, String) Time field format. For more information, please see the output parameters of the time format description of the strftime function in C language.
* `time_key` - (Optional, String) Time field key name. time_key and time_format must appear in pair.
* `un_match_log_key` - (Optional, String) Unmatched log key.
* `un_match_up_load_switch` - (Optional, Bool) Whether to upload the logs that failed to be parsed. Valid values: true: yes; false: no.

The `filter_key_regex` object supports the following:

* `key` - (Required, String) Log key to be filtered.
* `regex` - (Required, String) Filter rule regex corresponding to key.

The `meta_tags` object supports the following:

* `key` - (Optional, String) Key of the metadata tag for the machine group.
* `value` - (Optional, String) Value of the metadata tag for the machine group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `config_id` - Collection configuration ID.


## Import

cls config can be imported using the id, e.g.

```
terraform import cloud_cls_config.config config_id
```

