---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_role"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_role"
description: |-
  Provides a resource to create a CAM role.
---

# tencentcloudenterprise_cam_role

Provides a resource to create a CAM role.

## Example Usage

### # Create normally

```hcl
resource "tencentcloudenterprise_cam_role" "foo" {
  name     = "cam-role-test"
  document = <<EOF

	{
	  "version": "2.0",
	  "statement": [
	    {
	      "action": ["name/sts:AssumeRole"],
	      "effect": "allow",
	      "principal": {
	        "qcs": ["qcs::cam::uin/<your-account-id>:uin/<your-account-id>"]
	      }
	    }
	  ]
	}

EOF

  description   = "test"
  console_login = true
  tags = {
    test = "tf-cam-role",
  }
}
```

## Argument Reference

The following arguments are supported:

* `document` - (Required, String) Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string` and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be replaced with the uin it stands for.
* `name` - (Required, String, ForceNew) Name of CAM role.
* `console_login` - (Optional, Bool) Indicates whether the CAM role can login or not.
* `description` - (Optional, String) Description of the CAM role.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time of the CAM role.
* `update_time` - The last update time of the CAM role.

## Import

tencentcloudenterprise_cam_role can be imported using the id, e.g.

```
CAM role can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_role.foo 4611686018427733635
```
```

