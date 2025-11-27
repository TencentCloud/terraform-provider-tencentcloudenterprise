---
subcategory: "Tencent Service Framework(TSF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tsf_application"
sidebar_current: "docs-tencentcloudenterprise-datasource-tsf_application"
description: |-
  Use this data source to query detailed information of tsf application
---

# tencentcloudenterprise_tsf_application

Use this data source to query detailed information of tsf application

## Example Usage

```hcl
data "tencentcloudenterprise_tsf_application" "application" {
  application_type  = "V"
  microservice_type = "N"
  # application_resource_type_list = [""]
  application_id_list = ["application-a24x29xv"]
}
```

## Argument Reference

The following arguments are supported:

* `application_id_list` - (Optional, Set: [`String`]) Id list.
* `application_resource_type_list` - (Optional, Set: [`String`]) An array of application resource types.
* `application_type` - (Optional, String) The application type. V OR C, V means VM, C means container.
* `microservice_type` - (Optional, String) The microservice type of the application.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - The application paging list information.
  * `content` - The list of application information.
    * `apigateway_service_id` - Gateway service id.
    * `application_desc` - The description of the application.
    * `application_id` - The ID of the application.
    * `application_name` - The name of the application.
    * `application_remark_name` - Remark name.
    * `application_resource_type` - Application resource type.
    * `application_runtime_type` - Application runtime type.
    * `application_type` - The type of the application.
    * `create_time` - Create time.
    * `microservice_type` - The microservice type of the application.
    * `prog_lang` - Programming language.
    * `service_config_list` - Service config list.
      * `health_check` - Health check setting.
        * `path` - Health check path.
      * `name` - ServiceName.
      * `ports` - Port list.
        * `protocol` - Protocol.
        * `target_port` - Service port.
    * `update_time` - Update time.
  * `total_count` - The total number of applications.


