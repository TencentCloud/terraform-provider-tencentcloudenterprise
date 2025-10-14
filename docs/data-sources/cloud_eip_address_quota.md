---
subcategory: "Cloud Elastic IP(EIP)"
layout: "cloud"
page_title: "TencentCloud: cloud_eip_address_quota"
sidebar_current: "docs-cloud-datasource-eip_address_quota"
description: |-
  Use this data source to query detailed information of vpc address_quota
---

# cloud_eip_address_quota

Use this data source to query detailed information of vpc address_quota

## Example Usage

```hcl
data "cloud_eip_address_quota" "address_quota" {}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `quota_set` - The specified account EIP quota information.
  * `quota_current` - Current count.
  * `quota_id` - Quota name: TOTAL_EIP_QUOTA,DAILY_EIP_APPLY,DAILY_PUBLIC_IP_ASSIGN.
  * `quota_limit` - Quota count.


