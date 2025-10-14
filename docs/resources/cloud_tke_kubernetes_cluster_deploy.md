---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "cloud"
page_title: "TencentCloud: cloud_tke_kubernetes_cluster_deploy"
sidebar_current: "docs-cloud-resources-tke_kubernetes_cluster_deploy"
description: |-
  Provide a resource to increase instance to cluster
---

# cloud_tke_kubernetes_cluster_deploy

Provide a resource to increase instance to cluster

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

## Example Usage

```hcl
resource "cloud_kubernetes_cluster_deploy" "app-csp-sm" {
  cluster_id   = cloud_tke_kubernetes_cluster.cluster.id
  namespace    = "app-csp-sm"
  path         = "/apis/platform.tkestack.io/v1/clusters/cls-x8lxd2jx/apply"
  request_body = "{\"kind\":\"deploy\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"app-csp-sm\",\"annotations\":{\"description\":\"hkjc1\"}}}{\"kind\":\"deploy\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"qcloudregistrykey\",\"namespace\":\"app-csp-sm\",\"labels\":{\"qcloud-app\":\"qcloudregistrykey\"}},\"type\":\"kubernetes.io/dockercfg\",\"data\":{\".dockercfg\":\"eyJjY3IudGNlMzEwMHBvYy5mc3BoZXJlLmNuIjp7InVzZXJuYW1lIjoiMTAwMDA0NjAzMTU3IiwicGFzc3dvcmQiOiJ7QXBwbGljYXRpb25Ub2tlbjo0OGJlNzY2ZTVkZmRmN2JhZTAwZjdlZTQ3NTQyNDJlMX0iLCJlbWFpbCI6Im5vdEB2YWwuaWQiLCJhdXRoIjoiTVRBd01EQTBOakF6TVRVM09udEJjSEJzYVdOaGRHbHZibFJ2YTJWdU9qUTRZbVUzTmpabE5XUm1aR1kzWW1GbE1EQm1OMlZsTkRjMU5ESTBNbVV4ZlE9PSJ9fQ==\"}}"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) ID of the cluster.
* `deploy_name` - (Required, String, ForceNew) deploy_name.
* `namespace` - (Required, String, ForceNew) ns name.
* `request_body` - (Required, String, ForceNew) request_body.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



