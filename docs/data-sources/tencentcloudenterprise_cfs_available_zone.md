---
subcategory: "Cloud File Storage(CFS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfs_available_zone"
sidebar_current: "docs-tencentcloudenterprise-datasource-cfs_available_zone"
description: |-
  Use this data source to query detailed information of cfs available_zone
---

# tencentcloudenterprise_cfs_available_zone

Use this data source to query detailed information of cfs available_zone

## Example Usage

```hcl
data "tencentcloudenterprise_cfs_available_zone" "available_zone" {}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `region_zones` - Information such as resource availability in each AZ and the supported storage classes and protocols.
  * `region_cn_name` - Region chinese name, such as `Guangzhou`.
  * `region_name` - Region name, such as `bj`.
  * `region_status` - Region availability. If a region has at least one AZ where resources are purchasable, this value will be AVAILABLE; otherwise, it will be UNAVAILABLE.
  * `region` - Region name, such as `ap-beijing`.
  * `zones` - Array of AZs.
    * `types` - Array of classes.
      * `protocols` - Protocol and sale details.
        * `protocol` - Protocol type. Valid values: NFS, CIFS.
        * `sale_status` - 	Sale status. Valid values: sale_out (sold out), saling (purchasable), no_saling (non-purchasable).
      * `type` - Storage class. Valid values: SD (standard storage) and HP (high-performance storage).
    * `zone_cn_name` - Chinese name of an AZ.
    * `zone_en_name` - Chinese and English names of an AZ.
    * `zone_id` - AZ ID.
    * `zone` - AZ name.


