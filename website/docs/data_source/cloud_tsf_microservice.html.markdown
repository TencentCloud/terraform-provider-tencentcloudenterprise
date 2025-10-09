---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_microservice"
sidebar_current: "docs-cloud-datasource-tsf_microservice"
description: |-
  Use this data source to query detailed information of tsf microservice
---

# cloud_tsf_microservice

Use this data source to query detailed information of tsf microservice

## Example Usage

```hcl
data "cloud_tsf_microservice" "microservice" {
  namespace_id = var.namespace_id
  # status =
  microservice_id_list   = ["ms-yq3jo6jd"]
  microservice_name_list = ["provider-demo"]
}
```

## Argument Reference

The following arguments are supported:

* `namespace_id` - (Required, String) Namespace id.
* `microservice_id_list` - (Optional, Set: [`String`]) Microservice id list.
* `microservice_name_list` - (Optional, Set: [`String`]) List of service names for search.
* `result_output_file` - (Optional, String) Used to save results.
* `status` - (Optional, Set: [`String`]) Status filter, online, offline, single_online.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Microservice paging list information. Note: This field may return null, indicating that no valid value can be obtained.
  * `content` - Microservice list information. Note: This field may return null, indicating that no valid value can be obtained.
    * `create_time` - CreationTime. Note: This field may return null, indicating that no valid values can be obtained.
    * `critical_instance_count` - Offline instance count.  Note: This field may return null, indicating that no valid values can be obtained.
    * `microservice_desc` - Microservice description. Note: This field may return null, indicating that no valid value can be obtained.
    * `microservice_id` - Microservice Id. Note: This field may return null, indicating that no valid value can be obtained.
    * `microservice_name` - Microservice name. Note: This field may return null, indicating that no valid value can be obtained.
    * `namespace_id` - Namespace Id.  Note: This field may return null, indicating that no valid values can be obtained.
    * `run_instance_count` - Run instance count in namespace.  Note: This field may return null, indicating that no valid values can be obtained.
    * `update_time` - Last update time.  Note: This field may return null, indicating that no valid values can be obtained.
  * `total_count` - Microservice paging list information. Note: This field may return null, indicating that no valid value can be obtained.


