---
subcategory: "Cloud Object Storage(COS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cos_buckets"
sidebar_current: "docs-tencentcloudenterprise-datasource-cos_buckets"
description: |-
  Use this data source to query the COS buckets of the current Cloud user.
---

# tencentcloudenterprise_cos_buckets

Use this data source to query the COS buckets of the current Cloud user.

## Example Usage

```hcl
data "tencentcloudenterprise_cos_buckets" "cos_buckets" {
  result_output_file = "mytestpath"
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `bucket_list` - Cos bucket list.
  * `create_date` - bucket create date.
  * `name` - bucket name.
* `owner` - Information about the owner of the bucket, including the id and display_name fields, which represent the unique identifier and display name of the owner respectively.


