---
subcategory: "Cloud File Storage(CFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cfs_sign_up_cfs_service"
sidebar_current: "docs-cloud-resources-cfs_sign_up_cfs_service"
description: |-
  Provides a resource to create a cfs sign_up_cfs_service
---

# cloud_cfs_sign_up_cfs_service

Provides a resource to create a cfs sign_up_cfs_service

## Example Usage

```hcl
resource "cloud_cfs_sign_up_cfs_service" "sign_up_cfs_service" {}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `cfs_service_status` - Current status of the CFS service for this user. Valid values: creating (activating); created (activated).


