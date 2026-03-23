---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_image_manifests"
sidebar_current: "docs-tencentcloudenterprise-datasource-tcr_image_manifests"
description: |-
  Use this data source to query detailed information of tcr image_manifests
---

# tencentcloudenterprise_tcr_image_manifests

Use this data source to query detailed information of tcr image_manifests

## Example Usage

```hcl
data "tencentcloudenterprise_tcr_image_manifests" "image_manifests" {
  registry_id     = "%s"
  namespace_name  = "%s"
  repository_name = "%s"
  image_version   = "v1"
}
```

## Argument Reference

The following arguments are supported:

* `image_version` - (Required, String) Mirror version.
* `namespace_name` - (Required, String) Namespace name.
* `registry_id` - (Required, String) Instance ID.
* `repository_name` - (Required, String) Mirror warehouse name.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `config` - Configuration information of the image.
* `manifest` - Manifest information of the image.

