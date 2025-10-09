---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_application_config"
sidebar_current: "docs-cloud-resource-tsf_application_config"
description: |-
  Provides a resource to create a tsf application_config
---

# cloud_tsf_application_config

Provides a resource to create a tsf application_config

## Example Usage

```hcl
resource "cloud_tsf_application_config" "application_config" {
  config_name         = "tf-test-config"
  config_version      = "1.0"
  config_value        = "name: \"name\""
  application_id      = "application-oydl6xa2"
  config_version_desc = "version desc"
  // config_type = ""
  // encode_with_base64 = false
  // program_id_list =
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Required, String) Application ID.
* `config_name` - (Required, String) Configuration item name.
* `config_value` - (Required, String) Configuration item value.
* `config_version` - (Required, String) Configuration item version.
* `config_type` - (Optional, String) Configuration item value type.
* `config_version_desc` - (Optional, String) Configuration item version description.
* `encode_with_base64` - (Optional, Bool) Base64 encoded configuration items.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



