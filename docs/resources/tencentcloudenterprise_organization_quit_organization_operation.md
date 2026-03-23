---
subcategory: "Tencent Cloud Organization"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_organization_quit_organization_operation"
sidebar_current: "docs-tencentcloudenterprise-resource-organization_quit_organization_operation"
description: |-
  Provides a resource to create an organization quit_organization_operation
---

# tencentcloudenterprise_organization_quit_organization_operation

Provides a resource to create an organization quit_organization_operation

## Example Usage

```hcl
resource "tencentcloudenterprise_organization_quit_organization_operation" "quit_organization_operation" {
  org_id = 45155
}
```

## Argument Reference

The following arguments are supported:

* `org_id` - (Required, Int, ForceNew) Organization ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_organization_quit_organization_operation can be imported using the id, e.g.

```
organization quit_organization_operation can be imported using the id, e.g.

```
terraform import tencenttencentcloudenterprise_organization_quit_organization_operation.quit_organization_operation quit_organization_operation_id
```
```

