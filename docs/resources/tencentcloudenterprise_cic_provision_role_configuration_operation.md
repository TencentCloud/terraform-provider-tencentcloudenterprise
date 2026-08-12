---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_provision_role_configuration_operation"
sidebar_current: "docs-tencentcloudenterprise-resource-cic_provision_role_configuration_operation"
description: |-
  Provides a CIC permission configuration deployment operation. With no target
arguments, the resource discovers the current identity center and deploys every
permission configuration whose status is DeployedRequired.
---

# tencentcloudenterprise_cic_provision_role_configuration_operation

Provides a CIC permission configuration deployment operation. With no target
arguments, the resource discovers the current identity center and deploys every
permission configuration whose status is DeployedRequired.

If automatic discovery finds no matching record, the operation succeeds and
exports `provisioned_count = 0` without submitting a deployment task.

`role_configuration_id`, `target_type`, and `target_uin` must either all be
specified or all be omitted. Because this is an operation resource, Terraform
runs it when the resource is created. Use `terraform apply -replace=...` to run
the same configuration again.

## Example Usage

```hcl
# Deploy every role configuration whose status is DeployedRequired.
resource "tencentcloudenterprise_cic_provision_role_configuration_operation" "all_required" {}
```

### # Deploy or redeploy one specific target

```hcl
resource "tencentcloudenterprise_cic_provision_role_configuration_operation" "target" {
  role_configuration_id = "rc-xxxxxxxxxxxx"
  target_type           = "MemberUin"
  target_uin            = 100001234567
}
```

## Argument Reference

The following arguments are supported:

* `role_configuration_id` - (Optional, String, ForceNew) Permission configuration ID for a targeted deployment. Must be specified together with target_type and target_uin.
* `target_type` - (Optional, String, ForceNew) Type of the synchronized target account for a targeted deployment. Valid values: ManagerUin and MemberUin. Must be specified together with role_configuration_id and target_uin.
* `target_uin` - (Optional, Int, ForceNew) UIN of the target account for a targeted deployment. Must be specified together with role_configuration_id and target_type.
* `zone_id` - (Optional, String, ForceNew) Space ID. If omitted, the provider reads it from the current CIC identity center.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `provisioned_count` - Number of targets provisioned by this operation.
* `provisioned` - Details of targets provisioned by this operation.
  * `deployment_status` - Deployment status after this operation completes.
  * `role_configuration_id` - Permission configuration ID deployed by this operation.
  * `role_configuration_name` - Permission configuration name returned by CIC.
  * `target_name` - Name of the target account returned by CIC.
  * `target_type` - Type of the target account deployed by this operation.
  * `target_uin` - UIN of the target account deployed by this operation.
