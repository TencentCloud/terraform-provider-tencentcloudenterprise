---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_bot_resource"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_bot_resource"
description: |-
  Provides a resource to create a NGWAF bot resource.
---

# tencentcloudenterprise_ngwaf_bot_resource

Provides a resource to create a NGWAF bot resource.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_bot_resource" "example" {
  instance_id = "waf-xxxxxxxx"
  region_id   = "1"
  pay_mode    = 1
  pid         = 1000
  project_id  = 0
  goods_num   = 1
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) WAF instance ID.
* `pay_mode` - (Required, Int, ForceNew) Payment mode, 0: post-pay, 1: pre-pay.
* `pid` - (Required, Int, ForceNew) Product ID.
* `project_id` - (Required, Int, ForceNew) Project ID.
* `region_id` - (Required, String, ForceNew) Region ID.
* `goods_num` - (Optional, Int, ForceNew) Goods number.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_ngwaf_bot_resource can be imported using the id, e.g.

```
NGWAF bot resource can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_bot_resource.example waf-xxxxxxxx
```
```

