---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_encryption_protection"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_encryption_protection"
description: |-
  Provide a resource to enable/disable TKE cluster encryption protection.
---

# tencentcloudenterprise_tke_kubernetes_encryption_protection

Provide a resource to enable/disable TKE cluster encryption protection.

## Example Usage

```hcl
resource "tencentcloudenterprise_tke_kubernetes_encryption_protection" "example" {
  cluster_id = "cls-xxxxxxxx"

  kms_configuration {
    key_id     = "kms-xxxxxxxx"
    kms_region = "ap-guangzhou"
  }
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) ID of the cluster.
* `kms_configuration` - (Required, List, ForceNew) KMS encryption configuration.

The `kms_configuration` object supports the following:

* `key_id` - (Required, String, ForceNew) KMS key ID.
* `kms_region` - (Required, String, ForceNew) KMS region.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_tke_kubernetes_encryption_protection can be imported using the id, e.g.

```
TKE encryption protection can be imported using the cluster_id, e.g.
```
$ terraform import tencentcloudenterprise_tke_kubernetes_encryption_protection.example cls-xxxxxxxx
```
```

