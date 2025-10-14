---
subcategory: "Cloud Redis®(Redis)"
layout: "cloud"
page_title: "TencentCloud: cloud_redis_backup_download_info"
sidebar_current: "docs-cloud-datasource-redis_backup_download_info"
description: |-
  Use this data source to query detailed information of backup_download_info
---

# cloud_redis_backup_download_info

Use this data source to query detailed information of backup_download_info

## Example Usage

```hcl
data "cloud_redis_backup_download_info" "backup_download_info" {
  instance_id = "crs-iw7d9wdd"
  backup_id   = "641186639-8362913-1516672770"
  # limit_type = "NoLimit"
  # vpc_comparison_symbol = "In"
  # ip_comparison_symbol = "In"
  # limit_vpc {
  # 	region = "ap-guangzhou"
  # 	vpc_list = [""]
  # }
  # limit_ip = [""]
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Required, String) The backup ID, which can be accessed via [DescribeInstanceBackups](https://cloud.tencent.com/document/product/239/20011) interface returns the parameter RedisBackupSet to get.
* `instance_id` - (Required, String) The ID of instance.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `backup_infos` - A list of backup file information.
  * `download_url` - Backup file download address on the Internet (6 hours).
  * `file_name` - Backup file name.
  * `file_size` - The backup file size is in unit B, if it is 0, it is invalid.
  * `inner_download_url` - Backup file intranet download address (6 hours).


