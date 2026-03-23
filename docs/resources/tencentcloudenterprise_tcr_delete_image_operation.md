---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_delete_image_operation"
sidebar_current: "docs-tencentcloudenterprise-resource-tcr_delete_image_operation"
description: |-
  Provides a resource to create a tcr delete_image_operation
---

# tencentcloudenterprise_tcr_delete_image_operation

Provides a resource to create a tcr delete_image_operation

## Example Usage

```hcl
resource "tencentcloudenterprise_tcr_delete_image_operation" "delete_image_operation" {
  registry_id     = "tcr-xxx"
  repository_name = "repo"
  image_version   = "v1"
  namespace_name  = "ns"
}
```

## Argument Reference

The following arguments are supported:

* `image_version` - (Required, String, ForceNew) Image version name.
* `namespace_name` - (Required, String, ForceNew) Namespace name.
* `registry_id` - (Required, String, ForceNew) Instance id.
* `repository_name` - (Required, String, ForceNew) Repository name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


