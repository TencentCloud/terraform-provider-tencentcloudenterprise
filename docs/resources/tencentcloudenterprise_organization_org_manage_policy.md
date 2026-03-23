---
subcategory: "Tencent Cloud Organization"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_organization_org_manage_policy"
sidebar_current: "docs-tencentcloudenterprise-resource-organization_org_manage_policy"
description: |-
  Provides a resource to create an organization manage policy
---

# tencentcloudenterprise_organization_org_manage_policy

Provides a resource to create an organization manage policy

## Example Usage

```hcl
resource "tencentcloudenterprise_organization_org_manage_policy" "example" {
  name        = "example-policy"
  type        = "SERVICE_CONTROL_POLICY"
  description = "example policy description"
  content     = <<EOF
	{
	    "version": "2.0",
	    "statement": [
	        {
	            "effect": "deny",
	            "action": [
	                "account:*"
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

* `content` - (Required, String) Policy content. Refer to the CAM policy syntax.
* `name` - (Required, String) Policy name.
The length is 1~128 characters, which can include Chinese characters, English letters, numbers, and underscores.
* `description` - (Optional, String) Policy description.
* `type` - (Optional, String) Policy type. Default value is SERVICE_CONTROL_POLICY.
Valid values:
  - `SERVICE_CONTROL_POLICY`: Service control policy.
  - `TAG_POLICY`: Tag policy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `policy_id` - Policy Id.

