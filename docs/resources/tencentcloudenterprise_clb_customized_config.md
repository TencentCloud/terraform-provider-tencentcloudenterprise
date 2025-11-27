---
subcategory: "Cloud Load Balancer(CLB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_clb_customized_config"
sidebar_current: "docs-tencentcloudenterprise-resources-clb_customized_config"
description: |-
  Provides a resource to create a CLB customized config.
---

# tencentcloudenterprise_clb_customized_config

Provides a resource to create a CLB customized config.

## Example Usage

```hcl
resource "tencentcloudenterprise_clb_customized_config" "foo" {
  config_content = "client_max_body_size 224M;\r\nclient_body_timeout 60s;"
  config_name    = "helloWorld"
  load_balancer_ids = [
    "${tencentcloudenterprise_clb_instance.internal_clb.id}",
    "${tencentcloudenterprise_clb_instance.internal_clb2.id}",
  ]
}
```

## Argument Reference

The following arguments are supported:

* `config_content` - (Required, String) Content of Customized Config.
* `config_name` - (Required, String) Name of Customized Config.
* `load_balancer_ids` - (Optional, Set: [`String`]) List of LoadBalancer Ids.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time of Customized Config.
* `update_time` - Update time of Customized Config.


## Import

CLB customized config can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_clb_customized_config.foo pz-diowqstq
```

