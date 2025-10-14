---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_mariadb_sale_info"
sidebar_current: "docs-cloud-datasource-mariadb_sale_info"
description: |-
  Use this data source to query detailed information of mariadb sale_info
---

# cloud_mariadb_sale_info

Use this data source to query detailed information of mariadb sale_info

## Example Usage

```hcl
data "cloud_mariadb_sale_info" "sale_info" {
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `region_list` - list of sale region info.
  * `available_choice` - available zone choice.
    * `master_zone` - master zone.
      * `on_sale` - is zone on sale.
      * `zone_id` - zone id.
      * `zone_name` - zone name(zh).
      * `zone` - zone name(en).
    * `slave_zones` - slave zones.
      * `on_sale` - is zone on sale.
      * `zone_id` - zone id.
      * `zone_name` - zone name(zh).
      * `zone` - zone name(en).
  * `region_id` - region id.
  * `region_name` - region name(zh).
  * `region` - region name(en).
  * `zone_list` - list of az zone.
    * `on_sale` - is zone on sale.
    * `zone_id` - zone id.
    * `zone_name` - zone name(zh).
    * `zone` - zone name(en).


