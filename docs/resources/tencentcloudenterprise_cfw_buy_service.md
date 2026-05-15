---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_buy_service"
sidebar_current: "docs-tencentcloudenterprise-resource-cfw_buy_service"
description: |-
  Provide a resource to activate CFW (Cloud Firewall) service
---

# tencentcloudenterprise_cfw_buy_service

Provide a resource to activate CFW (Cloud Firewall) service

## Example Usage

```hcl
resource "tencentcloudenterprise_cfw_buy_service" "example" {
  region_id = "50000001"
  zone_id   = "50010001"
  vpc_spec  = "2"
}
```

## Argument Reference

The following arguments are supported:

* `region_id` - (Required, String, ForceNew) Region ID for CFW service activation.
* `vpc_spec` - (Required, String, ForceNew) VPC spec, 1: standard edition, 2: professional edition.
* `zone_id` - (Required, String, ForceNew) Zone ID for CFW service activation.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `pay_mode` - Billing mode.
* `resource_id` - Resource instance ID.
* `status` - Service status, 0: not activated, 1: activated.

