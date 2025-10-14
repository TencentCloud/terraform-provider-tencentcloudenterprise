---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "cloud"
page_title: "TencentCloud: cloud_tke_kubernetes_cluster_common_names"
sidebar_current: "docs-cloud-datasource-tke_kubernetes_cluster_common_names"
description: |-
  Provide a datasource to query cluster CommonNames.
---

# cloud_tke_kubernetes_cluster_common_names

Provide a datasource to query cluster CommonNames.

## Example Usage

```hcl
data "cloud_tke_kubernetes_cluster_common_names" "foo" {
  cluster_id      = "cls-tddgs8fl"
  subaccount_uins = ["100004610790", "0987654321"]
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) Cluster ID.
* `subaccount_uins` - (Required, List: [`String`]) List of sub-account. Up to 50 sub-accounts can be passed in at a time.
* `result_output_file` - (Optional, String) Used for save result.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - List of the CommonName in the certificate of the client corresponding to the sub-account UIN.
  * `cn` - The cn of the client certificate corresponding to the sub-account UIN.
  * `subaccount_uin` - User UIN.


