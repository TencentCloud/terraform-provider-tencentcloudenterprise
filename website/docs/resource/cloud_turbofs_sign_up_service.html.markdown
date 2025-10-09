---
subcategory: "Parallel File Storage(TurboFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_sign_up_service"
sidebar_current: "docs-cloud-resource-turbofs_sign_up_service"
description: |-
  Provides a resource to create a turbofs sign up service
---

# cloud_turbofs_sign_up_service

Provides a resource to create a turbofs sign up service

## Example Usage

```hcl
resource "cloud_turbofs_sign_up_service" "sign_up" {}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `cfs_service_status` - Current status of the TurboFS service for this user. Valid values: creating (activating); created (activated).


