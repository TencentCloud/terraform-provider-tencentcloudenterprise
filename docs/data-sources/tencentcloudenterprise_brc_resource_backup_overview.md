---
subcategory: "Backup and Recovery Center(BRC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_brc_resource_backup_overview"
sidebar_current: "docs-tencentcloudenterprise-datasource-brc_resource_backup_overview"
description: |-
  Use this data source to query detailed information of brc backup overviews
---

# tencentcloudenterprise_brc_resource_backup_overview

Use this data source to query detailed information of brc backup overviews

## Example Usage

```hcl
data "tencentcloudenterprise_brc_resource_backup_overview" "overview" {
  result_output_file = "backup_overview.json"
}
```

## Argument Reference

The following arguments are supported:

* `resource_type` - (Optional, String) Resource type filter. Valid values: INSTANCE, DISK, CFS, COS, CSP, MySQL_MariaDB, TDSQL_MySQL.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `backup_overview` - Backup overview list.
  * `backup_count` - Backup count.
  * `backup_resource_count` - Backup resource count.
  * `backup_size_mb` - Backup size in MB.
  * `resource_type` - Resource type.
* `overview_detail_set` - Overview detail set.
  * `is_open` - Whether the service is open.
  * `is_support` - Whether the service is supported.
  * `region` - Region.
  * `resource_overview` - Resource overview.
    * `backup_count` - Backup count.
    * `backup_resource_count` - Backup resource count.
    * `backup_size_mb` - Backup size in MB.
  * `resource_type` - Resource type.


