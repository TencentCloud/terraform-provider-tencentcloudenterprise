---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_webhook_trigger_logs"
sidebar_current: "docs-tencentcloudenterprise-datasource-tcr_webhook_trigger_logs"
description: |-
  Use this data source to query detailed information of tencentcloudenterprise_tcr_webhook_trigger_logs
---

# tencentcloudenterprise_tcr_webhook_trigger_logs

Use this data source to query detailed information of tencentcloudenterprise_tcr_webhook_trigger_logs

## Example Usage

```hcl
data "tencentcloudenterprise_tcr_webhook_trigger_logs" "my_logs" {
  registry_id = local.tcr_id
  namespace   = var.tcr_namespace
  trigger_id  = var.trigger_id
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required, String) Namespace.
* `registry_id` - (Required, String) Instance Id.
* `trigger_id` - (Required, Int) Trigger id.
* `result_output_file` - (Optional, String) Used to save results.
* `tags` - (Optional, Map) Tag description list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `logs` - Log list.
  * `creation_time` - Creation time.
  * `detail` - Webhook trigger detail.
  * `event_type` - Event type.
  * `id` - Log id.
  * `notify_type` - Notification type.
  * `status` - Status.
  * `trigger_id` - Trigger Id.
  * `update_time` - Update time.

