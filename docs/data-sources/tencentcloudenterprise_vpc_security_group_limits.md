---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_security_group_limits"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpc_security_group_limits"
description: |-
  Use this data source to query detailed information of vpc security_group_limits
---

# tencentcloudenterprise_vpc_security_group_limits

Use this data source to query detailed information of vpc security_group_limits

## Example Usage

```hcl
data "tencentcloudenterprise_vpc_security_group_limits" "security_group_limits" {}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `security_group_limit_set` - Sg limit set.
  * `instance_security_group_limit` - Number of instances associated sg.
  * `referred_security_group_limit` - Number of sg can be referred.
  * `security_group_extended_policy_limit` - Number of sg extended policy.
  * `security_group_instance_limit` - Number of sg associated instances.
  * `security_group_limit` - Number of sg can be created.
  * `security_group_policy_limit` - Number of sg polciy can be created.
  * `security_group_referred_cvm_and_eni_limit` - Number of eni and cvm can be referred.
  * `security_group_referred_svc_limit` - Number of svc can be referred.


