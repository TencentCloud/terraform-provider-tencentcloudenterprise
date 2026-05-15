---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_cmd_template"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_cmd_template"
description: |-
  Provide a resource to create a DASB command template
---

# tencentcloudenterprise_dasb_cmd_template

Provide a resource to create a DASB command template

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_cmd_template" "example" {
  name     = "example-cmd-template"
  cmd_list = "rm -rf*"
}
```

## Argument Reference

The following arguments are supported:

* `cmd_list` - (Required, String) Command list, n separated, maximum length 32768 bytes.
* `name` - (Required, String) Template name, maximum length 32 characters, cannot contain blank characters.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_dasb_cmd_template can be imported using the id, e.g.

```
DASB command template can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_cmd_template.example 12345
```
```

