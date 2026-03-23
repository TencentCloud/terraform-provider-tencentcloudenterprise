---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_anti_info_leak"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_anti_info_leak"
description: |-
  Provides a resource to create a NGWAF anti information leak rule.
---

# tencentcloudenterprise_ngwaf_anti_info_leak

Provides a resource to create a NGWAF anti information leak rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_anti_info_leak" "example" {
  domain      = "example.com"
  name        = "anti-info-leak-rule"
  action_type = 1
  uri         = "/api/user"
  status      = 1

  strategies {
    field   = "information"
    content = "phone"
  }
}
```

## Argument Reference

The following arguments are supported:

* `action_type` - (Required, Int) Rule Action, 0 (log), 1 (replace), 2 (only display the first four digits), 3 (only display the last four digits), 4 (deny).
* `domain` - (Required, String) Domain.
* `name` - (Required, String) Rule Name.
* `strategies` - (Required, List) Strategies detail.
* `uri` - (Required, String) Uri.
* `status` - (Optional, Int) status.

The `strategies` object supports the following:

* `content` - (Required, String) Matching content
          The following options are available when Field is set to information:
          idcard (ID card), phone (phone number), and bankcard (bank card).
          The following options are available when Field is set to returncode:
          400 (status code 400), 403 (status code 403), 404 (status code 404), 4xx (other 4xx status codes), 500 (status code 500), 501 (status code 501), 502 (status code 502), 504 (status code 504), and 5xx (other 5xx status codes).
          When Field is set to keywords, users need to input the matching content themselves.
* `field` - (Required, String) Matching Criteria, returncode (Response Code), keywords (Keywords), information (Sensitive Information).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_ngwaf_anti_info_leak can be imported using the id, e.g.

```
NGWAF anti info leak rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_anti_info_leak.example rule_id#example.com
```
```

