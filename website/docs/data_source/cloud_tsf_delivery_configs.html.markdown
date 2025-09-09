---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_delivery_configs"
sidebar_current: "docs-cloud-datasource-tsf_delivery_configs"
description: |-
  Use this data source to query detailed information of tsf delivery_configs
---

# cloud_tsf_delivery_configs

Use this data source to query detailed information of tsf delivery_configs

## Example Usage

```hcl
data "cloud_tsf_delivery_configs" "delivery_configs" {
  search_word = "test"
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.
* `search_word` - (Optional, String) search word.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - deploy group information about the deployment group associated with a delivery item.Note: This field may return null, which means that no valid value was obtained.
  * `content` - content. Note: This field may return null, which means that no valid value was obtained.
    * `collect_path` - harvest log path. Note: This field may return null, which means that no valid value was obtained.
    * `config_id` - config id.
    * `config_name` - config name.
    * `create_time` - Creation time.Note: This field may return null, indicating that no valid values can be obtained.
    * `groups` - Associated deployment group information.Note: This field may return null, indicating that no valid values can be obtained.
      * `associate_time` - Associate Time. Note: This field may return null, indicating that no valid values can be obtained.
      * `cluster_id` - Cluster ID. Note: This field may return null, indicating that no valid values can be obtained.
      * `cluster_name` - Cluster Name. Note: This field may return null, indicating that no valid values can be obtained.
      * `cluster_type` - Cluster type.
      * `group_id` - Group Id.
      * `group_name` - Group Name.
      * `namespace_name` - Namespace Name. Note: This field may return null, indicating that no valid values can be obtained.
  * `total_count` - total count. Note: This field may return null, which means that no valid value was obtained.


