---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "cloud"
page_title: "TencentCloud: cloud_tke_kubernetes_available_cluster_versions"
sidebar_current: "docs-cloud-datasource-tke_kubernetes_available_cluster_versions"
description: |-
  Use this data source to query detailed information of kubernetes available_cluster_versions
---

# cloud_tke_kubernetes_available_cluster_versions

Use this data source to query detailed information of kubernetes available_cluster_versions

## Example Usage

```hcl
data "cloud_tke_kubernetes_available_cluster_versions" "query_by_id" {
  cluster_id = "xxx"
}

output "versions_id" {
  description = "Query versions from id."
  value       = data.cloud_tke_kubernetes_available_cluster_versions.query_by_id.versions
}

data "cloud_tke_kubernetes_available_cluster_versions" "query_by_ids" {
  cluster_ids = ["xxx"]
}

output "versions_ids" {
  description = "Query versions from ids."
  value       = data.cloud_tke_kubernetes_available_cluster_versions.query_by_ids.clusters
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional, String) Cluster Id.  cluster_id and cluster_ids must be set only one.
* `cluster_ids` - (Optional, Set: [`String`]) List of cluster IDs. cluster_id and cluster_ids must be set only one.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `clusters` - Cluster information. Note: This field may return null, indicating that no valid value can be obtained.
  * `cluster_id` - Cluster ID.
  * `versions` - List of cluster major version numbers, for example 1.18.4.
* `versions` - Upgradable cluster version number. Note: This field may return null, indicating that no valid value can be obtained.


