---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "cloud"
page_title: "TencentCloud: cloud_tke_kubernetes_charts"
sidebar_current: "docs-cloud-datasource-tke_kubernetes_charts"
description: |-
  Use this data source to query detailed information of kubernetes cluster addons.
---

# cloud_tke_kubernetes_charts

Use this data source to query detailed information of kubernetes cluster addons.

## Example Usage

```hcl
data "cloud_tke_kubernetes_charts" "name" {}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.
* `search` - (Optional, String) Search of chart.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - App chart list.
  * `label` - Label of chart.
  * `latest_version` - Chart latest version.
  * `name` - Name of chart.


