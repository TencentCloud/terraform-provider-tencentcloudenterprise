---
subcategory: "Application Performance Management(APM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_apm_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-apm_instance"
description: |-
  Provides a resource to create a apm instance
---

# tencentcloudenterprise_apm_instance

Provides a resource to create a apm instance

## Example Usage

```hcl
resource "tencentcloudenterprise_apm_instance" "instance" {
  name           = "terraform-test"
  description    = "for terraform test"
  trace_duration = 15
}

output "instance_id" {
  value = tencentcloudenterprise_apm_instance.instance.instance_id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name Of Instance.
* `description` - (Optional, String) Description Of Instance.
* `trace_duration` - (Optional, Int) Duration of trace data retention in days. Valid values range from 1 to 30. Default is 7.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `instance_id` - Instance ID.

## Import

tencentcloudenterprise_apm_instance can be imported using the id, e.g.

```
apm instance can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_apm_instance.instance instance_id
```
```

