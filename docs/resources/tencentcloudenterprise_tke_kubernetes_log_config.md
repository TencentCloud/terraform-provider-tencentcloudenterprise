---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_log_config"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_log_config"
description: |-
  Provide a resource to manage TKE CLS log configs via CreateCLSLogConfig/DescribeLogConfigs/DeleteLogConfigs.
---

# tencentcloudenterprise_tke_kubernetes_log_config

Provide a resource to manage TKE CLS log configs via CreateCLSLogConfig/DescribeLogConfigs/DeleteLogConfigs.

## Example Usage

```hcl
resource "tencentcloudenterprise_tke_kubernetes_log_config" "log_config" {
  cluster_id      = "cls-xxxxxx"
  log_config_name = "test-tf"
  logset_id       = "9066db41-1758-4b73-a3a5-310882aba884"

  log_config = jsonencode({
    apiVersion = "cls.cloud.tencent.com/v1"
    kind       = "LogConfig"
    metadata = {
      name = "test-tf"
    }
    spec = {
      clsDetail = {
        logType   = "json_log"
        logFormat = "default"
        region    = "ap-qingyuan-region-devtest-ops"
        period    = 30
        indexs    = [{ indexName = "namespace" }]
        extractRule = {
          unMatchUpload = "true"
          unMatchedKey  = "LogParseFailure"
          backtracking  = "-1"
          jsonStandard  = "true"
          isGBK         = "false"
        }
      }
      inputDetail = {
        type = "container_stdout"
        containerStdout = {
          allContainers     = true
          namespace         = "default"
          metadataContainer = ["namespace", "container_id", "container_name", "image_name", "cluster_id"]
          metadataLabels    = ["__NULL__"]
        }
      }
    }
  })
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) Cluster ID.
* `log_config_name` - (Required, String, ForceNew) Log config name.
* `log_config` - (Required, String, ForceNew) JSON expression of log collection configuration.
* `cluster_type` - (Optional, String, ForceNew) Cluster type, supported values: tke, eks. Default is tke.
* `logset_id` - (Optional, String, ForceNew) CLS log set ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


