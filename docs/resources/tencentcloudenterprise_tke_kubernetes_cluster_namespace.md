---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_cluster_namespace"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_cluster_namespace"
description: |-
  Provide a resource to create/read/delete a namespace in cluster
---

# tencentcloudenterprise_tke_kubernetes_cluster_namespace

Provide a resource to create/read/delete a namespace in cluster

## Example Usage

### # JSON

```hcl
resource "tencentcloudenterprise_tke_kubernetes_cluster_namespace" "ns" {
  cluster_id = var.cluster_id
  namespace  = "tf-test"

  path = format("/apis/platform.tkestack.io/v1/clusters/%s/apply", var.cluster_id)

  request_body = jsonencode({
    apiVersion = "v1"
    kind       = "Namespace"
    metadata = {
      name = "tf-test"
      annotations = {
        description = "ddddd"
      }
    }
  })
}
```

### # YAML

```hcl
resource "tencentcloudenterprise_tke_kubernetes_cluster_namespace" "ns" {
  cluster_id = var.cluster_id
  namespace  = "tf-test"

  path = format("/apis/platform.tkestack.io/v1/clusters/%s/apply", var.cluster_id)

  request_body = <<-YAML

apiVersion: v1
kind: Namespace
metadata:

	name: tf-test
	labels:
	  env: dev
	annotations:
	  description: created-by-terraform

YAML
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) ID of the cluster.
* `namespace` - (Required, String, ForceNew) Namespace name used for read/delete.
* `path` - (Required, String, ForceNew) Apply endpoint for namespace creation, e.g. /apis/platform.tkestack.io/v1/clusters/<cluster_id>/apply.
* `request_body` - (Required, String, ForceNew) Namespace manifest in JSON or YAML. Must include apiVersion=v1 and kind=Namespace.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `apiversion` - apiversion.
* `message` - message.
* `status` - status.

