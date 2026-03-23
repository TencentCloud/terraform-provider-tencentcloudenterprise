---
subcategory: "Cloud Kafka(ckafka)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ckafka_route"
sidebar_current: "docs-tencentcloudenterprise-resource-ckafka_route"
description: |-
  Provides a resource to create a ckafka route.
---

# tencentcloudenterprise_ckafka_route

Provides a resource to create a ckafka route.

## Example Usage

```hcl
resource "tencentcloudenterprise_ckafka_route" "example" {
  instance_id = "ckafka-qz8w8rxz"
  vip_type    = 3
  vpc_id      = "vpc-k1azc3mv"
  subnet_id   = "subnet-or7ddsbc"
  access_type = 0
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) CKafka instance id.
* `vip_type` - (Required, Int, ForceNew) Routing network type (3: vpc routing; 4: standard support routing; 7: professional support routing).
* `access_type` - (Optional, Int) Access type. Valid values:
- 0: PLAINTEXT (in clear text, supported by both the old version and the community version without user information)
- 1: SASL_PLAINTEXT (in clear text, but at the beginning of the data, authentication will be logged in through SASL, which is only supported by the community version)
- 2: SSL (SSL encrypted communication without user information, supported by both older and community versions)
- 3: SASL_SSL (SSL encrypted communication. When the data starts, authentication will be logged in through SASL. Only the community version supports it).
* `auth_flag` - (Optional, Int) Auth flag.
* `caller_appid` - (Optional, Int) Caller appid.
* `subnet_id` - (Optional, String) Subnet id.
* `vpc_id` - (Optional, String) Vpc id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `route_id` - Route ID, available after route creation is complete.
* `vip_list` - Virtual IP list.
  * `vip` - Virtual IP.
  * `vport` - Virtual port.

## Import

tencentcloudenterprise_ckafka_route can be imported using the id, e.g.

```
ckafka route can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ckafka_route.example ckafka-qz8w8rxz#14256
```
```

