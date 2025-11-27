---
subcategory: "Cloud Virtual Machine(CVM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cvm_import_image_os"
sidebar_current: "docs-tencentcloudenterprise-datasource-cvm_import_image_os"
description: |-
  Use this data source to query detailed information of cvm import_image_os
---

# tencentcloudenterprise_cvm_import_image_os

Use this data source to query detailed information of cvm import_image_os

## Example Usage

```hcl
data "tencentcloudenterprise_cvm_import_image_os" "import_image_os" {
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `import_image_os_list_supported` - Supported operating system types of imported images.
  * `linux` - Supported Linux OS Note: This field may return null, indicating that no valid values can be obtained.
  * `windows` - Supported Windows OS Note: This field may return null, indicating that no valid values can be obtained.
* `import_image_os_version_set` - Supported operating system versions of imported images.
  * `architecture` - Supported operating system architecture.
  * `os_name` - Operating system type.
  * `os_versions` - Supported operating system versions.


