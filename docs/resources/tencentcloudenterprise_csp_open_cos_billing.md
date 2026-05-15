---
subcategory: "Cloud Storage Platform(CSP)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_csp_open_cos_billing"
sidebar_current: "docs-tencentcloudenterprise-resource-csp_open_cos_billing"
description: |-
  Provides a resource to create a csp open_cos_billing
---

# tencentcloudenterprise_csp_open_cos_billing

Provides a resource to create a csp open_cos_billing

## Example Usage

```hcl
resource "tencentcloudenterprise_csp_open_cos_billing" "example" {
  cos_region = "ap-guangzhou"
}
```

## Argument Reference

The following arguments are supported:

* `cos_region` - (Required, String, ForceNew) The COS region where the bucket resides.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


