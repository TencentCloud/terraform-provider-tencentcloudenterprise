---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_cluster_plugin"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_cluster_plugin"
description: |-
  Provide a resource to manage TKE application in cluster via ForwardApplicationRequestV3.
---

# tencentcloudenterprise_tke_kubernetes_cluster_plugin

Provide a resource to manage TKE application in cluster via ForwardApplicationRequestV3.

## Example Usage

### Repository chart:

```hcl
resource "tencentcloudenterprise_tke_kubernetes_cluster_plugin" "alertmanager" {
  cluster_id = "cls-rkeuubqw"
  namespace  = "tke-addon"

  request_body = jsonencode({
    kind       = "App"
    apiVersion = "application.tkestack.io/v1"
    metadata = {
      name      = "alertmanager"
      namespace = "tke-addon"
      labels = {
        "application.tkestack.io/type" = "internal-app"
      }
    }
    spec = {
      chart = {
        chartGroupName = "local"
        chartName      = "alertmanager"
        chartVersion   = "1.7.0"
        tenantID       = "local"
        importedRepo   = true
      }
      name          = "alertmanager"
      targetCluster = "cls-rkeuubqw"
      type          = "HelmV3"
      values = {
        rawValues     = ""
        rawValuesType = "yaml"
        values        = ["rootDir="]
      }
      dryRun = false
    }
  })
}
```

### Local chart package (tgz):

```hcl
locals {
  chart_b64  = filebase64("${path.module}/charts/alertmanager-1.7.0.tgz")
  values_yml = file("${path.module}/values.yaml")
}

resource "tencentcloudenterprise_tke_kubernetes_cluster_plugin" "alertmanager" {
  cluster_id = "cls-rkeuubqw"
  namespace  = "tke-cluster-inspection"

  request_body = jsonencode({
    kind       = "App"
    apiVersion = "application.tkestack.io/v1"
    metadata = {
      name      = "alertmanager"
      namespace = "tke-cluster-inspection"
      annotations = {
        "application.tkestack.io/chart" = local.chart_b64
      }
      labels = {
        "application.tkestack.io/type" = "internal-app"
      }
    }
    spec = {
      chart = {
        chartGroupName = "local"
        chartName      = "alertmanager"
        chartVersion   = "1.7.0"
        importedRepo   = true
      }
      name          = "alertmanager"
      targetCluster = "cls-rkeuubqw"
      type          = "HelmV3"
      values = {
        rawValuesType = "yaml"
        rawValues     = local.values_yml
      }
      dryRun = false
    }
  })
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) ID of the cluster.
* `namespace` - (Required, String, ForceNew) Namespace for the application.
* `request_body` - (Required, String, ForceNew) request_body.
* `path` - (Optional, String, ForceNew) Application endpoint path. If empty, defaults to /apis/application.tkestack.io/v1/namespaces/<namespace>/apps.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `apiversion` - apiversion.
* `app_id` - app_id.

