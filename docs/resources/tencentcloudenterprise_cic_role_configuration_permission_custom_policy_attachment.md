---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment"
sidebar_current: "docs-tencentcloudenterprise-resource-cic_role_configuration_permission_custom_policy_attachment"
description: |-
  Provides a resource to create an organization cic_role_configuration_permission_custom_policy_attachment
---

# tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment

Provides a resource to create an organization cic_role_configuration_permission_custom_policy_attachment

## Example Usage

```hcl
resource "tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment" "cic_role_configuration_permission_custom_policy_attachment" {
  zone_id               = "z-xxxxxx"
  role_configuration_id = "rc-xxxxxx"
  role_policy_name      = "CustomPolicy"
  role_policy_document  = <<-EOF
{
    "version": "2.0",
    "statement": [
        {
            "effect": "allow",
            "action": [
                "vpc:DescribeVpcs"
            ],
            "resource": [
                "*"
            ]
        }
    ]
}
EOF
}
```

## Argument Reference

The following arguments are supported:

* `role_configuration_id` - (Required, String, ForceNew) Permission configuration ID.
* `role_policy_document` - (Required, String, ForceNew) Role policy document.
* `role_policy_name` - (Required, String, ForceNew) Role policy name.
* `zone_id` - (Required, String, ForceNew) Space ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `add_time` - Role policy add time.
* `role_policy_type` - Role policy type.

## Import

tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment can be imported using the id, e.g.

```
organization cic_role_configuration_permission_custom_policy_attachment can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment.cic_role_configuration_permission_custom_policy_attachment ${zoneId}#${roleConfigurationId}#${rolePolicyName}
```
```

