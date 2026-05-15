---
subcategory: "Secrets Manager(SSM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ssm_buy_service"
sidebar_current: "docs-tencentcloudenterprise-resource-ssm_buy_service"
description: |-
  Provides a resource to create a ssm buy_service
---

# tencentcloudenterprise_ssm_buy_service

Provides a resource to create a ssm buy_service

## Example Usage

```hcl
resource "tencentcloudenterprise_ssm_buy_service" "buy_service" {}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `invalid_type` - Service unavailable type. 0: not purchased; 1: normal; 2: arrears; 3: resource released.
* `service_enabled` - Whether the SSM service is enabled. true: enabled; false: not enabled.

