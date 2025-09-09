---
subcategory: "Cloud Virtual Machine(CVM)"
layout: "cloud"
page_title: "TencentCloud: cloud_cvm_image_quota"
sidebar_current: "docs-cloud-datasource-cvm_image_quota"
description: |-
  Use this data source to query detailed information of cvm image_quota
---

# cloud_cvm_image_quota

Use this data source to query detailed information of cvm image_quota

## Example Usage

```hcl
data "cloud_cvm_image_quota" "image_quota" {
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `image_num_quota` - The image quota of an account.


