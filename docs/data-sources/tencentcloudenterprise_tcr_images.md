---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_images"
sidebar_current: "docs-tencentcloudenterprise-datasource-tcr_images"
description: |-
  Use this data source to query detailed information of tcr images
---

# tencentcloudenterprise_tcr_images

Use this data source to query detailed information of tcr images

## Example Usage

```hcl
data "tencentcloudenterprise_tcr_images" "images" {
  registry_id     = "tcr-xxx"
  namespace_name  = "ns"
  repository_name = "repo"
  image_version   = "v1"
  digest          = "sha256:xxxxx"
  exact_match     = false
}
```

## Argument Reference

The following arguments are supported:

* `namespace_name` - (Required, String) Namespace name.
* `registry_id` - (Required, String) Instance id.
* `repository_name` - (Required, String) Repository name.
* `digest` - (Optional, String) Specify image digest for lookup.
* `exact_match` - (Optional, Bool) Specifies whether it is an exact match, true is an exact match, and not filled is a fuzzy match.
* `image_version` - (Optional, String) Image version name, default is fuzzy match.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `image_info_list` - Container image information list.
  * `digest` - Hash value.
  * `image_version` - Tag name.
  * `kind` - Product type,note: this field may return null, indicating that no valid value can be obtained.
  * `kms_signature` - Kms signature information,note: this field may return null, indicating that no valid value can be obtained.
  * `size` - Image size (unit: byte).
  * `update_time` - Update time.

