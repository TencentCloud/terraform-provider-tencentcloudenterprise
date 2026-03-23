---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_bot_status_config"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_bot_status_config"
description: |-
  Provides a resource to manage the bot status configuration for a NGWAF protected domain.
---

# tencentcloudenterprise_ngwaf_bot_status_config

Provides a resource to manage the bot status configuration for a NGWAF protected domain.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_bot_status_config" "example" {
  instance_id = "waf-xxxxxxxx"
  domain      = "example.com"
  status      = "1"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String, ForceNew) Domain.
* `instance_id` - (Required, String, ForceNew) Instance ID.
* `status` - (Required, String) Bot status. 1 - enable; 0 - disable.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `current_global_scene` - The currently enabled scenario with a global matching range and the highest priority.
  * `priority` - Priority.
  * `scene_id` - Scene ID.
  * `scene_name` - Scene name.
  * `update_time` - Update time.
* `custom_rule_nums` - Total number of custom rules, excluding BOT whitelist.
* `scene_count` - Scene total count.
* `valid_scene_count` - Number of effective scenarios.

## Import

tencentcloudenterprise_ngwaf_bot_status_config can be imported using the id, e.g.

```
NGWAF bot status config can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_bot_status_config.example waf-xxxxxxxx#example.com
```
```

