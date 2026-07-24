---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_provision_role_configuration_operation"
sidebar_current: "docs-tencentcloudenterprise-resource-cic_provision_role_configuration_operation"
description: |-
  Provides a resource to create a cic provision_role_configuration_operation
---

# tencentcloudenterprise_cic_provision_role_configuration_operation

Provides a resource to create a cic provision_role_configuration_operation

## Example Usage

```hcl
data "tencentcloudenterprise_cic_identity_center" "center" {}

resource "tencentcloudenterprise_cic_role_configuration" "example" {
  zone_id                 = data.tencentcloudenterprise_cic_identity_center.center.zone_id
  role_configuration_name = "tf-provision-example"
  description             = "tf example"
}

resource "tencentcloudenterprise_cic_provision_role_configuration_operation" "example" {
  zone_id               = data.tencentcloudenterprise_cic_identity_center.center.zone_id
  role_configuration_id = tencentcloudenterprise_cic_role_configuration.example.role_configuration_id
  target_type           = "MemberUin"
  target_uin            = 100001234567
}
```

## Argument Reference

The following arguments are supported:

* `role_configuration_id` - (Required, String, ForceNew) Permission configuration ID.
* `target_type` - (Required, String, ForceNew) Type of the synchronized target account of the Tencent Cloud Organization. ManagerUin: admin account; MemberUin: member account.
* `target_uin` - (Required, Int, ForceNew) UIN of the target account of the Tencent Cloud Organization.
* `zone_id` - (Required, String, ForceNew) Space ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


