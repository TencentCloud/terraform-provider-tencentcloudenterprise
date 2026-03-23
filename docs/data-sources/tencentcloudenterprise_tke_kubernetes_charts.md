---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_charts"
sidebar_current: "docs-tencentcloudenterprise-datasource-tke_kubernetes_charts"
description: |-
  Use this data source to query detailed information of kubernetes cluster addons.
---

# tencentcloudenterprise_tke_kubernetes_charts

Use this data source to query detailed information of kubernetes cluster addons.

## Example Usage

```hcl
data "tencentcloudenterprise_tke_kubernetes_charts" "name" {
  kind         = "network"
  arch         = "amd64"
  cluster_type = "tke"
}
```

## Argument Reference

The following arguments are supported:

* `arch` - (Optional, String) Operation system app supported. Available values: `arm32`, `arm64`, `amd64`.
* `cluster_type` - (Optional, String) Cluster type. Available values: `tke`, `eks`.
* `kind` - (Optional, String) Kind of app chart. Available values: `log`, `scheduler`, `network`, `storage`, `monitor`, `dns`, `image`, `other`, `invisible`.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `chart_list` - App chart list.
  * `label` - Label of chart.
  * `latest_version` - Chart latest version.
  * `name` - Name of chart.

