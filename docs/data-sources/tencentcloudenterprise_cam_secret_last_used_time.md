---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_secret_last_used_time"
sidebar_current: "docs-tencentcloudenterprise-datasource-cam_secret_last_used_time"
description: |-
  使用Account SDK接口查询密钥最后使用时间。
---

# tencentcloudenterprise_cam_secret_last_used_time

使用Account SDK接口查询密钥最后使用时间。

Use Account SDK API to query secret last used time.

## Example Usage

```hcl
data "tencentcloudenterprise_cam_secret_last_used_time" "example" {
  secret_id_list = ["AKIDxxxxxxxxxxxx"]
}
```

## Argument Reference

The following arguments are supported:

* `secret_id_list` - (Required, Set: [`String`]) Query the key ID list. Supports up to 10.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `secret_id_last_used_rows` - Last used time list.
  * `last_secret_used_date` - Last used timestamp.
  * `last_used_date` - Last used date (with 1 day delay).
  * `secret_id` - Secret Id.

