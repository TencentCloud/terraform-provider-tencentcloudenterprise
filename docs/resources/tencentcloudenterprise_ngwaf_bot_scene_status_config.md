---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_bot_scene_status_config"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_bot_scene_status_config"
description: |-
  Provides a resource to manage the bot scene status configuration for a NGWAF protected domain.
---

# tencentcloudenterprise_ngwaf_bot_scene_status_config

Provides a resource to manage the bot scene status configuration for a NGWAF protected domain.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_bot_scene_status_config" "example" {
  domain   = "example.com"
  scene_id = "scene-xxxxxxxx"
  status   = true
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String, ForceNew) Domain.
* `scene_id` - (Required, String, ForceNew) Scene ID.
* `status` - (Required, Bool) Bot status. true - enable; false - disable.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `priority` - Priority.
* `scene_name` - Scene name.
* `type` - Scene type, default: Default scenario, custom: Non default scenario.

## Import

tencentcloudenterprise_ngwaf_bot_scene_status_config can be imported using the id, e.g.

```
NGWAF bot scene status config can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_bot_scene_status_config.example example.com#scene-xxxxxxxx
```
```

