---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_gateway_all_group_apis"
sidebar_current: "docs-cloud-datasource-tsf_gateway_all_group_apis"
description: |-
  Use this data source to query detailed information of tsf gateway_all_group_apis
---

# cloud_tsf_gateway_all_group_apis

Use this data source to query detailed information of tsf gateway_all_group_apis

## Example Usage

```hcl
data "cloud_tsf_gateway_all_group_apis" "gateway_all_group_apis" {
  gateway_deploy_group_id = "group-aeoej4qy"
  search_word             = "user"
}
```

## Argument Reference

The following arguments are supported:

* `gateway_deploy_group_id` - (Required, String) Gateway group Id.
* `result_output_file` - (Optional, String) Used to save results.
* `search_word` - (Optional, String) Search keyword, supports api group name or API path.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Gateway group and API list information.
  * `gateway_deploy_group_id` - Gateway deployment group ID.Note: This field may return null, which means no valid value was found.
  * `gateway_deploy_group_name` - Gateway deployment group name.Note: This field may return null, which means no valid value was found.
  * `group_num` - Gateway deployment api group number.Note: This field may return null, which means no valid value was found.
  * `groups` - Gateway deployment  api group list.Note: This field may return null, which means no valid value was found.
    * `group_api_count` - Number of APIs under the group. Note: This field may return null, which means no valid value was found.
    * `group_apis` - List of APIs under the group.Note: This field may return null, which means no valid value was found.
      * `api_id` - API ID.
      * `method` - API method.
      * `microservice_name` - API service name.
      * `namespace_name` - Namespace name.
      * `path` - API path.
    * `group_id` - Api group id.Note: This field may return null, which means no valid value was found.
    * `group_name` - Api group name.Note: This field may return null, which means no valid value was found.


