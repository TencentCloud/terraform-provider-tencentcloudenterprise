---
subcategory: "Tencent Service Framework(TSF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tsf_api_detail"
sidebar_current: "docs-tencentcloudenterprise-datasource-tsf_api_detail"
description: |-
  Use this data source to query detailed information of tsf api_detail
---

# tencentcloudenterprise_tsf_api_detail

Use this data source to query detailed information of tsf api_detail

## Example Usage

```hcl
data "tencentcloudenterprise_tsf_api_detail" "api_detail" {
  microservice_id = "ms-yq3jo6jd"
  path            = "/printRequest"
  method          = "GET"
  pkg_version     = "20210625192923"
  application_id  = "application-a24x29xv"
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Required, String) Application id.
* `method` - (Required, String) Request method.
* `microservice_id` - (Required, String) Microservice id.
* `path` - (Required, String) Api path.
* `pkg_version` - (Required, String) Pkg version.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Api detail.
  * `can_run` - Can debug or not.
  * `definitions` - Api data struct.
    * `name` - Object name.
    * `properties` - Object property list.
      * `description` - Property description.
      * `name` - Property name.
      * `type` - Property type.
  * `description` - API description. Note: This field may return null, indicating that no valid value can be obtained.
  * `request_content_type` - Api content type.
  * `request` - Api request description.
    * `default_value` - Default value.
    * `description` - Param description.
    * `in` - Param position.
    * `name` - Param name.
    * `required` - Require or not.
    * `type` - Type.
  * `response` - Api response.
    * `description` - Param description.
    * `name` - Param description.
    * `type` - Param type.
  * `status` - API status 0: offline 1: online, default 0. Note: This section may return null, indicating that no valid value can be obtained.


