---
subcategory: "CFS Turbo(Turbofs)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_sign_up_cfs_service"
sidebar_current: "docs-cloud-resources-turbofs_sign_up_cfs_service"
description: |-
  Provides a resource to create a turbofs sign_up_cfs_service
---

# cloud_turbofs_sign_up_cfs_service

Provides a resource to create a turbofs sign_up_cfs_service

## Example Usage

```hcl
resource "cloud_turbofs_sign_up_cfs_service" "sign_up_cfs_service" {}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `cfs_service_status` - Current status of the CFS service for this user. Valid values: creating (activating); created (activated).


