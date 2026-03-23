---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_webhook_trigger"
sidebar_current: "docs-tencentcloudenterprise-resource-tcr_webhook_trigger"
description: |-
  Provides a resource to create a tcr webhook_trigger
---

# tencentcloudenterprise_tcr_webhook_trigger

Provides a resource to create a tcr webhook_trigger

## Example Usage

```hcl
resource "tencentcloudenterprise_tcr_instance" "mytcr_webhooktrigger" {
  name          = "tf-test-tcr-%s"
  instance_type = "basic"
  delete_bucket = true

  tags = {
    test = "test"
  }
}

resource "tencentcloudenterprise_tcr_namespace" "my_ns" {
  instance_id    = tencentcloudenterprise_tcr_instance.mytcr_webhooktrigger.id
  name           = "tf_test_ns_%s"
  is_public      = true
  is_auto_scan   = true
  is_prevent_vul = true
  severity       = "medium"
  cve_whitelist_items {
    cve_id = "cve-xxxxx"
  }
}

data "tencentcloudenterprise_tcr_namespaces" "id_test" {
  instance_id = tencentcloudenterprise_tcr_namespace.my_ns.instance_id
}

locals {
  ns_id = data.tencentcloudenterprise_tcr_namespaces.id_test.namespace_list.0.id
}

resource "tencentcloudenterprise_tcr_webhook_trigger" "my_trigger" {
  registry_id = tencentcloudenterprise_tcr_instance.mytcr_webhooktrigger.id
  namespace   = tencentcloudenterprise_tcr_namespace.my_ns.name
  trigger {
    name = "trigger-%s"
    targets {
      address = "http://example.org/post"
      headers {
        key    = "X-Custom-Header"
        values = ["a"]
      }
    }
    event_types  = ["pushImage"]
    condition    = ".*"
    enabled      = true
    description  = "this is trigger description"
    namespace_id = local.ns_id

  }
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required, String) Namespace name.
* `registry_id` - (Required, String) Instance Id.
* `trigger` - (Required, List) Trigger parameters.
* `tags` - (Optional, Map) Tag description list.

The `headers` object supports the following:

* `key` - (Required, String) Header Key.
* `values` - (Required, Set) Header Values.

The `targets` object supports the following:

* `address` - (Required, String) Target address.
* `headers` - (Optional, List) Custom Headers.

The `trigger` object supports the following:

* `condition` - (Required, String) Trigger rule.
* `enabled` - (Required, Bool) Enable trigger.
* `event_types` - (Required, Set) Trigger action.
* `name` - (Required, String) Trigger name.
* `targets` - (Required, List) Trigger target.
* `description` - (Optional, String) Trigger description.
* `namespace_id` - (Optional, Int) The namespace Id to which the trigger belongs.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_tcr_webhook_trigger can be imported using the id, e.g.

```
tcr webhook_trigger can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_tcr_webhook_trigger.webhook_trigger webhook_trigger_id
```
```

