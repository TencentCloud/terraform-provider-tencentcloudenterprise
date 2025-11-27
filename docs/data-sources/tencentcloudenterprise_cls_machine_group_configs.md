---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_machine_group_configs"
sidebar_current: "docs-tencentcloudenterprise-datasource-cls_machine_group_configs"
description: |-
  Use this data source to query detailed information of cls machine_group_configs
---

# tencentcloudenterprise_cls_machine_group_configs

Use this data source to query detailed information of cls machine_group_configs

## Example Usage

```hcl
resource "tencentcloudenterprise_cls_machine_group" "group" {
  group_name        = "tf-describe-mg-config-test"
  service_logging   = true
  auto_update       = true
  update_end_time   = "19:05:00"
  update_start_time = "17:05:00"

  machine_group_type {
    type = "ip"
    values = [
      "203.0.113.101",
      "203.0.113.102",
    ]
  }
}

data "tencentcloudenterprise_cls_machine_group_configs" "machine_group_configs" {
  group_id = tencentcloudenterprise_cls_machine_group.group.id
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required, String) Group id.
* `result_output_file` - (Optional, String) path to save result file.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configs` - Collection configuration list.
  * `advanced_config` - Advanced config.
  * `config_extra_id` - Config extra id.
  * `config_flag` - Config flag.
  * `config_id` - Collection configuration ID.
  * `create_time` - Create time.
  * `exclude_paths` - Collection path blocklist.
  * `extract_rule` - Extraction rule. If ExtractRule is set, LogType must be set.
    * `address` - Syslog system log collection specifies the address and port that the collector listens to.
    * `backtracking` - Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection).
    * `begin_regex` - First-Line matching rule, which is valid only if log_type is multiline_log or fullregex_log.
    * `delimiter` - Delimiter for delimited log, which is valid only if log_type is delimiter_log.
    * `event_log_rules` - Windows event log collection rules, which only take effect when LogType is windows_event_log.
      * `event_channel` - Event channel, supports Application, Security, Setup, System, ALL.
      * `event_ids` - Event ID filter list. An empty list means no filtering. It supports positive filtering of single values or ranges, and also supports negative filtering of single values.
      * `time_stamp` - Time, when the user selects a custom time type, a specific time needs to be specified.
      * `time_type` - Time type, 1: user-defined, 2: current time.
    * `filter_key_regex` - Log keys to be filtered and the corresponding regex.
    * `is_gbk` - GBK encoding. Default 0.
    * `json_standard` - Standard json. Default 0.
    * `keys` - Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if log_type is delimiter_log. json_log logs use the key of JSON itself.
    * `log_regex` - Full log matching rule, which is valid only if log_type is fullregex_log.
    * `meta_tags` - Metadata tags.
      * `key` - Metadata tag key.
      * `value` - Metadata tag value.
    * `metadata_type` - Metadata type.
    * `parse_protocol` - Parse protocol.
    * `path_regex` - Metadata path regex.
    * `protocol` - Syslog protocol, tcp or udp.
    * `time_format` - Time field format. For more information, please see the output parameters of the time format description of the strftime function in C language.
    * `time_key` - The key name of the time field. time_key and time_format must appear in pairs.
    * `un_match_log_key` - Unmatched log key.
    * `un_match_up_load_switch` - Whether to upload the logs that failed to be parsed. Valid values: true: yes; false: no.
  * `log_format` - Log formatting method.
  * `log_type` - Type of the log to be collected. json_log represents JSON format logs, delimiter_log represents delimiter format logs, minimalist_log represents minimalist logs, multiline_log represents multi-line logs, fullregex_log represents full regex, default is minimalist_log.
  * `name` - Collection configuration name.
  * `output` - Topicid.
  * `path` - Log collection path, including filename.
  * `source` - Config source.
  * `update_time` - Update time.
  * `user_define_rule` - User define rule.


