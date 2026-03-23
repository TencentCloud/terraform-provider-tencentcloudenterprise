---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_addon"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_addon"
description: |-
  Provide a resource to manage TKE addons via InstallAddon/DescribeAddon/UpdateAddon/DeleteAddon.
---

# tencentcloudenterprise_tke_kubernetes_addon

Provide a resource to manage TKE addons via InstallAddon/DescribeAddon/UpdateAddon/DeleteAddon.

## Example Usage

```hcl
resource "tencentcloudenterprise_tke_kubernetes_addon" "cbs" {
  cluster_id = "cls-rkeuubqw"
  addon_name = "cbs"

  # raw_values accepts JSON and will be base64-encoded by the provider.
  raw_values = jsonencode({
    tolerations = [
      {
        key      = "123"
        operator = "Exists"
      }
    ]
  })

  values = [
    "global.image.host=ccr.d12-x86.fsphere.cn",
    "global.cluster.id=cls-rkeuubqw",
    "global.cluster.appid=1255000044",
    "global.cluster.uin=110000000047",
    "global.cluster.subuin=110000000370",
    "global.cluster.type=tke",
    "global.cluster.clustertype=INDEPENDENT_CLUSTER",
    "global.cluster.kubeversion=1.22.5",
    "global.cluster.kubeminor=22.0",
    "cbs.url=cbs.api3.d12-x86.fsphere.cn",
    "cvm.url=cvm.api3.d12-x86.fsphere.cn",
    "cfs.url=cfs.api3.d12-x86.fsphere.cn",
    "metadata.url=http://product-cvm-metadata.ap-qingyuan-region-devtest-ops.d12-x86.fsphere.cn/meta-data",
  ]
}
```

## Argument Reference

The following arguments are supported:

* `addon_name` - (Required, String, ForceNew) Name of addon.
* `cluster_id` - (Required, String, ForceNew) ID of cluster.
* `addon_version` - (Optional, String) Version of addon. If no set, the latest version will be installed by default.
* `raw_values` - (Optional, String) Addon params in JSON format. Provider will base64-encode before sending.
* `values` - (Optional, List: [`String`]) Addon params list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `phase` - Status of addon.
* `reason` - Reason of addon failed.

