---
subcategory: "Tencent Service Framework(TSF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tsf_namespace"
sidebar_current: "docs-tencentcloudenterprise-resources-tsf_namespace"
description: |-
  Provides a resource to create a tsf namespace
---

# tencentcloudenterprise_tsf_namespace

Provides a resource to create a tsf namespace

## Example Usage

```hcl
resource "tencentcloudenterprise_tsf_namespace" "namespace" {
  namespace_name = "terraform-namespace-name"
  namespace_desc = "terraform-test"
  namespace_type = "DEF"
  is_ha_enable   = "0"
}
```

## Argument Reference

The following arguments are supported:

* `namespace_name` - (Required, String) Namespace name.
* `cluster_id` - (Optional, String) Cluster ID.
* `is_ha_enable` - (Optional, String) Whether to enable high availability.
* `namespace_desc` - (Optional, String) Namespace description.
* `namespace_id` - (Optional, String) Namespace ID.
* `namespace_resource_type` - (Optional, String) Namespace resource type (default is DEF).
* `namespace_type` - (Optional, String) Whether it is a global namespace (the default is DEF, which means a common namespace; GLOBAL means a global namespace).
* `program_id` - (Optional, String) ID of the dataset to be bound.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time.
* `delete_flag` - Delete ID.
* `is_default` - Default namespace.
* `namespace_code` - Namespace encoding.
* `namespace_status` - Namespace status.
* `update_time` - Update time.


