---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_microservice_api_version"
sidebar_current: "docs-cloud-datasource-tsf_microservice_api_version"
description: |-
  Use this data source to query detailed information of tsf microservice_api_version
---

# cloud_tsf_microservice_api_version

Use this data source to query detailed information of tsf microservice_api_version

## Example Usage

```hcl
data "cloud_tsf_microservice_api_version" "microservice_api_version" {
  microservice_id = "ms-yq3jo6jd"
  path            = ""
  method          = "get"
}
```

## Argument Reference

The following arguments are supported:

* `microservice_id` - (Required, String) Microservice ID.
* `method` - (Optional, String) Request method.
* `path` - (Optional, String) Api path.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Api version list.
  * `application_id` - Application ID.
  * `application_name` - Application Name.
  * `pkg_version` - Application pkg version.


