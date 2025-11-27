---
subcategory: "Tencent Service Framework(TSF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tsf_delivery_configs"
sidebar_current: "docs-tencentcloudenterprise-datasource-tsf_delivery_configs"
description: |-
  Use this data source to query detailed information of tsf delivery_configs
---

# tencentcloudenterprise_tsf_delivery_configs

Use this data source to query detailed information of tsf delivery_configs

## Example Usage

```hcl
data "tencentcloudenterprise_tsf_delivery_configs" "delivery_configs" {
  search_word = "test"
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.
* `search_word` - (Optional, String) Search word.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Deploy group information about the deployment group associated with a delivery item.Note: This field may return null, which means that no valid value was obtained.
  * `content` - Content. Note: This field may return null, which means that no valid value was obtained.
    * `collect_path` - Harvest log path. Note: This field may return null, which means that no valid value was obtained.
    * `config_id` - Config id.
    * `config_name` - Config name.
    * `create_time` - Creation time.Note: This field may return null, indicating that no valid values can be obtained.
    * `groups` - Associated deployment group information.Note: This field may return null, indicating that no valid values can be obtained.
      * `associate_time` - Associate Time. Note: This field may return null, indicating that no valid values can be obtained.
      * `cluster_id` - Cluster ID. Note: This field may return null, indicating that no valid values can be obtained.
      * `cluster_name` - Cluster Name. Note: This field may return null, indicating that no valid values can be obtained.
      * `cluster_type` - Cluster type.
      * `group_id` - Group Id.
      * `group_name` - Group Name.
      * `namespace_name` - Namespace Name. Note: This field may return null, indicating that no valid values can be obtained.
  * `total_count` - Total count. Note: This field may return null, which means that no valid value was obtained.


