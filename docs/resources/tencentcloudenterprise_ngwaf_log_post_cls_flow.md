---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_log_post_cls_flow"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_log_post_cls_flow"
description: |-
  Provides a resource to manage NGWAF log posting to CLS (Cloud Log Service).
---

# tencentcloudenterprise_ngwaf_log_post_cls_flow

Provides a resource to manage NGWAF log posting to CLS (Cloud Log Service).

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_log_post_cls_flow" "example" {
  cls_region     = "ap-shanghai"
  logset_name    = "waf_post_logset"
  log_type       = 1
  log_topic_name = "waf_post_logtopic"
}
```

## Argument Reference

The following arguments are supported:

* `cls_region` - (Optional, String) The region where the CLS is delivered. The default value is ap-shanghai.
* `log_topic_name` - (Optional, String) The name of the log subject where the submitted CLS is located. The default value is waf_post_logtopic.
* `log_type` - (Optional, Int) 1- Access log, 2- Attack log, the default is access log.
* `logset_name` - (Optional, String) The name of the log set where the delivered CLS is located. The default value is waf_post_logset.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `flow_id` - Unique ID for post cls flow.
* `log_topic_id` - CLS log topic ID.
* `logset_id` - CLS logset ID.
* `status` - Status 0- Off 1- On.

## Import

tencentcloudenterprise_ngwaf_log_post_cls_flow can be imported using the id, e.g.

```
NGWAF log post cls flow can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_log_post_cls_flow.example flow_id#log_type
```
```

