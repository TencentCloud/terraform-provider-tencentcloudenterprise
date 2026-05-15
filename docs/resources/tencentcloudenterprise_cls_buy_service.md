---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_buy_service"
sidebar_current: "docs-tencentcloudenterprise-resource-cls_buy_service"
description: |-
  Provides a resource to create a cls buy_service
---

# tencentcloudenterprise_cls_buy_service

Provides a resource to create a cls buy_service

## Example Usage

```hcl
resource "tencentcloudenterprise_cls_buy_service" "buy_service" {}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `status` - Account status. 0: not activated, 1: normal, 2: overdue, 3: destroyed.

