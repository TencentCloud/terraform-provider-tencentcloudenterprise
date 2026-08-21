---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_user_permissions"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_user_permissions"
description: |-
  Provides a resource to manage TKE cluster user permissions via RBAC.
---

# tencentcloudenterprise_tke_kubernetes_user_permissions

Provides a resource to manage TKE cluster user permissions via RBAC.

## Example Usage

```hcl
resource "tencentcloudenterprise_tke_kubernetes_user_permissions" "example" {
  cluster_id = "cls-xxxxxxxx"
  target_uin = "110000000000"

  permissions {
    role_name = "tke:admin"
    role_type = "cluster"
  }

  permissions {
    role_name = "tke:dev"
    role_type = "namespace"
    namespace = "default"
  }
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) Cluster ID.
* `target_uin` - (Required, String, ForceNew) Unique identifier of the user to be authorized (supports sub-account UIN and role UIN).
* `permissions` - (Optional, Set) Complete list of permissions that the user should ultimately have. Uses declarative semantics, the passed list represents all permissions the user should ultimately have, the system will automatically calculate differences and perform necessary create/delete operations. When empty or not provided, all permissions for this user will be cleared.

The `permissions` object supports the following:

* `role_name` - (Required, String) Role name. Predefined roles include: tke:admin (cluster administrator), tke:ops (operations personnel), tke:dev (developer), tke:ro (read-only user), tke:ns:dev (namespace developer), tke:ns:ro (namespace read-only user), others are user-defined roles.
* `role_type` - (Required, String) Authorization type. Enum values: cluster (cluster-level permissions, corresponding to ClusterRoleBinding), namespace (namespace-level permissions, corresponding to RoleBinding).
* `is_custom` - (Optional, Bool) Whether it is a custom role, default false.
* `namespace` - (Optional, String) Namespace. Required when role_type is namespace.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_tke_kubernetes_user_permissions can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_tke_kubernetes_user_permissions.example cls-xxxxxxxx:110000000000
```

