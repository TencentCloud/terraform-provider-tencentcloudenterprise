---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_repository"
sidebar_current: "docs-tencentcloudenterprise-resource-tcr_repository"
description: |-
  Use this resource to create tcr repository.
---

# tencentcloudenterprise_tcr_repository

Use this resource to create tcr repository.

## Example Usage

```hcl
data "tencentcloudenterprise_tcr_instances" "test" {
  name = "test"
}

resource "tencentcloudenterprise_tcr_repository" "foo" {
  instance_id    = data.tencentcloudenterprise_tcr_instances.test.instance_list[0].id
  namespace_name = "exampleNamespace"
  name           = "example"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) ID of the TCR instance.
* `name` - (Required, String, ForceNew) Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
* `namespace_name` - (Required, String, ForceNew) Name of the TCR namespace.
* `brief_desc` - (Optional, String) Brief description of the repository. Valid length is [1~100].
* `description` - (Optional, String) Description of the repository. Valid length is [1~1000].

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time.
* `is_public` - Indicate the repository is public or not.
* `update_time` - Last updated time.
* `url` - URL of the repository.

## Import

tencentcloudenterprise_tcr_repository can be imported using the id, e.g.

```
tcr repository can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_tcr_repository.foo cls-cda1iex1#namespace#repository
```
```

