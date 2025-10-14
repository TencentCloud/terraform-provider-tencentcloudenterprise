---
subcategory: "Cloud Log Service(CLS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cls_config_attachment"
sidebar_current: "docs-cloud-resources-cls_config_attachment"
description: |-
  Provides a resource to create a cls config attachment
---

# cloud_cls_config_attachment

Provides a resource to create a cls config attachment

## Example Usage

```hcl
resource "cloud_cls_config_attachment" "attach" {
  config_id = cloud_cls_config.config.id
  group_id  = "27752a9b-9918-440a-8ee7-9c84a14a47ed"
}

# Import

cls config_attachment can be imported using the id, e.g.
```

### terraform import cloud_cls_config_attachment.attach config_id#group_id

```hcl

```

## Argument Reference

The following arguments are supported:

* `config_id` - (Required, String, ForceNew) Collection configuration id.
* `group_id` - (Required, String, ForceNew) Machine group id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



