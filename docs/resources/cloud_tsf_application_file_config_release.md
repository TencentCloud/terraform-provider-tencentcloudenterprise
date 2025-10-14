---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_application_file_config_release"
sidebar_current: "docs-cloud-resources-tsf_application_file_config_release"
description: |-
  Provides a resource to create a tsf application_file_config_release
---

# cloud_tsf_application_file_config_release

Provides a resource to create a tsf application_file_config_release

## Example Usage

```hcl
resource "cloud_tsf_application_file_config_release" "application_file_config_release" {
  config_id    = "dcfg-f-123456"
  group_id     = "group-123456"
  release_desc = "product release"
}
```

## Argument Reference

The following arguments are supported:

* `config_id` - (Required, String, ForceNew) File config id.
* `group_id` - (Required, String, ForceNew) Group Id.
* `release_desc` - (Optional, String, ForceNew) Release Description.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

tsf applicationfile_config_release can be imported using the id, e.g.

```
terraform import cloud_tsf_application_file_config_release.application_file_config_release application_file_config_release_id
```

