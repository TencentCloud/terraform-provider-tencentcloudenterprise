---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_health_check_policy"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_health_check_policy"
description: |-
  Provide a resource to create a TKE health check policy.
---

# tencentcloudenterprise_tke_kubernetes_health_check_policy

Provide a resource to create a TKE health check policy.

## Example Usage

```hcl
resource "tencentcloudenterprise_tke_kubernetes_health_check_policy" "example" {
  cluster_id = "cls-xxxxxxxx"
  name       = "health-policy-example"

  rules {
    name                = "OOMKilling"
    enabled             = true
    auto_repair_enabled = true
  }

  rules {
    name                = "KubeletUnhealthy"
    enabled             = true
    auto_repair_enabled = false
  }
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) ID of the cluster.
* `name` - (Required, String, ForceNew) Health check policy name.
* `rules` - (Required, List) Health check policy rules.

The `rules` object supports the following:

* `auto_repair_enabled` - (Required, Bool) Whether to enable auto repair.
* `enabled` - (Required, Bool) Whether to enable this check item.
* `name` - (Required, String) Health check rule name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_tke_kubernetes_health_check_policy can be imported using the id, e.g.

```
TKE health check policy can be imported using cluster_id#policy_name, e.g.
```
$ terraform import tencentcloudenterprise_tke_kubernetes_health_check_policy.example cls-xxxxxxxx#health-policy-example
```
```

