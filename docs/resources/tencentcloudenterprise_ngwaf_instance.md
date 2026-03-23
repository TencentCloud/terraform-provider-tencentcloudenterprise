---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_instance"
description: |-
  Provides a resource to create a NGWAF instance.
---

# tencentcloudenterprise_ngwaf_instance

Provides a resource to create a NGWAF instance.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_instance" "clb_waf" {
  instance_name  = "production-clb-waf"
  region_id      = "50000001"
  pay_mode       = 0
  pid            = 11543
  type           = 1
  project_id     = 0
  goods_num      = 1
  bot_expand     = 0
  interface_name = "qcloud.Hour.create"
}
```

## Argument Reference

The following arguments are supported:

* `bot_expand` - (Optional, Int, ForceNew) Bot expansion flag.
* `goods_num` - (Optional, Int, ForceNew) Goods number.
* `instance_name` - (Optional, String) Instance name.
* `interface_name` - (Optional, String, ForceNew) Interface name for API request.
* `pay_mode` - (Optional, Int, ForceNew) Payment mode, 0: post-pay, 1: pre-pay.
* `pid` - (Optional, Int, ForceNew) Product ID.
* `project_id` - (Optional, Int, ForceNew) Project ID.
* `region_id` - (Optional, String, ForceNew) Region ID.
* `type` - (Optional, Int, ForceNew) Instance type.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `attack_log_post` - Attack log post status.
* `begin_time` - Begin time.
* `domain_count` - Domain count.
* `edition` - Instance edition.
* `instance_id` - Instance unique ID.
* `main_domain_count` - Main domain count.
* `region` - Region name.
* `status` - Instance status.
* `valid_time` - Valid time.

## Import

tencentcloudenterprise_ngwaf_instance can be imported using the id, e.g.

```
NGWAF instance can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_instance.clb_waf waf-xxxxxxxx
```
```

