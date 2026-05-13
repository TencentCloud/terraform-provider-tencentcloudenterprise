---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_addon_config"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_addon_config"
description: "Provide a resource to configure addon that kubernetes comes with. Unlike tencentcloudenterprise_tke_kubernetes_addon which manages the full lifecycle (install/update/delete), this resource only manages addon configuration (update). It will not install or delete the addon."
---

# tencentcloudenterprise_tke_kubernetes_addon_config

Provide a resource to configure addon that kubernetes comes with.
Unlike tencentcloudenterprise_tke_kubernetes_addon which manages the full lifecycle (install/update/delete),
this resource only manages addon configuration (update). It will not install or delete the addon.

## Example Usage

```hcl
resource "tencentcloudenterprise_tke_kubernetes_addon_config" "example" {
  cluster_id = "cls-rkeuubqw"
  addon_name = "cluster-autoscaler"
  raw_values = jsonencode({
    extraArgs = {
      scale-down-enabled               = true
      max-empty-bulk-delete            = 11
      scale-down-delay-after-add       = "10mm"
      scale-down-unneeded-time         = "10mm"
      scale-down-utilization-threshold = 0.005
      ignore-daemonsets-utilization    = false
      skip-nodes-with-local-storage    = true
      skip-nodes-with-system-pods      = true
    }
  })
}
```

## Argument Reference

The following arguments are supported:

* `addon_name` - (Required, String, ForceNew) Name of addon.
* `cluster_id` - (Required, String, ForceNew) ID of cluster.
* `addon_version` - (Optional, String) Version of addon.
* `raw_values` - (Optional, String) Params of addon, base64 encoded json format.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `phase` - Status of addon.
* `reason` - Reason of addon failed.

