---
subcategory: "Tencent Service Framework(TSF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tsf_application_attribute"
sidebar_current: "docs-tencentcloudenterprise-datasource-tsf_application_attribute"
description: |-
  Use this data source to query detailed information of tsf application_attribute
---

# tencentcloudenterprise_tsf_application_attribute

Use this data source to query detailed information of tsf application_attribute

## Example Usage

```hcl
data "tencentcloudenterprise_tsf_application_attribute" "application_attribute" {
  application_id = "application-a24x29xv"
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Required, String) Application Id.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Application list other attribute.
  * `group_count` - Number of deployment groups under the application.Note: This field may return null, indicating that no valid values can be obtained.
  * `instance_count` - Total number of instances.Note: This field may return null, indicating that no valid values can be obtained.
  * `run_instance_count` - Number of running instances.Note: This field may return null, indicating that no valid values can be obtained.


