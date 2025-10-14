---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_application_public_config"
sidebar_current: "docs-cloud-resources-tsf_application_public_config"
description: |-
  Provides a resource to create a tsf application_public_config
---

# cloud_tsf_application_public_config

Provides a resource to create a tsf application_public_config

## Example Usage

```hcl
resource "cloud_tsf_application_public_config" "application_public_config" {
  config_name         = "my_config"
  config_version      = "1.0"
  config_value        = "test: 1"
  config_version_desc = "product version"
  // config_type = "P"
  encode_with_base64 = true
  # program_id_list =
}
```

## Argument Reference

The following arguments are supported:

* `config_name` - (Required, String, ForceNew) Config Name.
* `config_value` - (Required, String, ForceNew) Config value, only yaml file allowed.
* `config_version` - (Required, String, ForceNew) Config version.
* `config_version_desc` - (Optional, String, ForceNew) Config version description.
* `encode_with_base64` - (Optional, Bool, ForceNew) The config value is encoded with base64 or not.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



