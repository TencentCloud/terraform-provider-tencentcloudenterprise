---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_application_file_config"
sidebar_current: "docs-cloud-datasource-tsf_application_file_config"
description: |-
  Use this data source to query detailed information of tsf application_file_config
---

# cloud_tsf_application_file_config

Use this data source to query detailed information of tsf application_file_config

## Example Usage

```hcl
data "cloud_tsf_application_file_config" "application_file_config" {
  config_id = "dcfg-f-4y4ekzqv"
  # config_id_list = [""]
  config_name    = "file-log1"
  application_id = "application-2vzk6n3v"
  config_version = "1.2"
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional, String) Application ID.
* `config_id_list` - (Optional, Set: [`String`]) List of configuration item ID.
* `config_id` - (Optional, String) Configuration ID.
* `config_name` - (Optional, String) Configuration item name.
* `config_version` - (Optional, String) Configuration item version.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - File configuration item list. Note: This field may return null, indicating that no valid values can be obtained.
  * `content` - File configuration array. Note: This field may return null, indicating that no valid values can be obtained.
    * `application_id` - Application Id. Note: This field may return null, indicating that no valid values can be obtained.
    * `application_name` - Application name. Note: This field may return null, indicating that no valid values can be obtained.
    * `config_file_code` - Configuration file code. Note: This field may return null, indicating that no valid values can be obtained.
    * `config_file_name` - Configuration item file name. Note: This field may return null, indicating that no valid values can be obtained.
    * `config_file_path` - File config path. Note: This field may return null, indicating that no valid values can be obtained.
    * `config_file_value_length` - Config item content length.  Note: This field may return null, indicating that no valid values can be obtained.
    * `config_file_value` - Configuration file content. Note: This field may return null, indicating that no valid values can be obtained.
    * `config_id` - Config ID. Note: This field may return null, indicating that no valid values can be obtained.
    * `config_name` - Configuration item name. Note: This field may return null, indicating that no valid values can be obtained.
    * `config_post_cmd` - Last update time.  Note: This field may return null, indicating that no valid values can be obtained.
    * `config_version_count` - Config version count.  Note: This field may return null, indicating that no valid values can be obtained.
    * `config_version_desc` - Configuration item version description. Note: This field may return null, indicating that no valid values can be obtained.
    * `config_version` - Configuration version. Note: This field may return null, indicating that no valid values can be obtained.
    * `creation_time` - CreationTime. Note: This field may return null, indicating that no valid values can be obtained.
    * `delete_flag` - Delete flag, true: allow delete; false: delete prohibit.
    * `last_update_time` - Last update time.  Note: This field may return null, indicating that no valid values can be obtained.
  * `total_count` - Total count.


