---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_group_gateways"
sidebar_current: "docs-cloud-datasource-tsf_group_gateways"
description: |-
  Use this data source to query detailed information of tsf group_gateways
---

# cloud_tsf_group_gateways

Use this data source to query detailed information of tsf group_gateways

## Example Usage

```hcl
data "cloud_tsf_group_gateways" "group_gateways" {
  gateway_deploy_group_id = "group-aeoej4qy"
  search_word             = "test"
}
```

## Argument Reference

The following arguments are supported:

* `gateway_deploy_group_id` - (Required, String) Gateway group Id.
* `result_output_file` - (Optional, String) Used to save results.
* `search_word` - (Optional, String) Search word.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Api group information.
  * `content` - Api group Info.
    * `acl_mode` - ACL type for accessing the group.Note: This field may return null, which means no valid value was found.
    * `api_count` - Number of APIs.Note: This field may return null, which means no valid value was found.
    * `auth_type` - Authentication type. secret: key authentication; none: no authentication.Note: This field may return null, which means no valid value was found.
    * `binded_gateway_deploy_groups` - Gateway deployment group bound to the API group.Note: This field may return null, which means no valid value was found.
      * `application_id` - Application ID.Note: This field may return null, which means no valid value was found.
      * `application_name` - Application name.Note: This field may return null, which means no valid value was found.
      * `application_type` - Application category: V: virtual machine application, C: container application.Note: This field may return null, which means no valid value was found.
      * `cluster_type` - Cluster type, with possible values: C: container, V: virtual machine.Note: This field may return null, which means no valid value was found.
      * `deploy_group_id` - Gateway deployment group ID.Note: This field may return null, which means no valid value was found.
      * `deploy_group_name` - Gateway deployment group name.Note: This field may return null, which means no valid value was found.
      * `group_status` - Application status of the deployment group, with possible values: Running, Waiting, Paused, Updating, RollingBack, Abnormal, Unknown.Note: This field may return null, which means no valid value was found.
    * `created_time` - Group creation time, such as: 2019-06-20 15:51:28.Note: This field may return null, which means no valid value was found.
    * `description` - Description.Note: This field may return null, which means no valid value was found.
    * `group_context` - Api group context.Note: This field may return null, which means no valid value was found.
    * `group_id` - Api group id.Note: This field may return null, which means no valid value was found.
    * `group_name` - Api group name.Note: This field may return null, which means no valid value was found.
    * `group_type` - Group type. ms: microservice group; external: external API group.This field may return null, which means no valid value was found.
    * `status` - Release status. drafted: not released. released: released.Note: This field may return null, which means no valid value was found.
    * `updated_time` - Group update time, such as: 2019-06-20 15:51:28.Note: This field may return null, which means no valid value was found.
  * `total_count` - Total count.


