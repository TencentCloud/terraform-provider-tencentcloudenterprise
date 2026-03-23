---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_cc_session"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_cc_session"
description: |-
  Provides a resource to create a NGWAF CC session.
---

# tencentcloudenterprise_ngwaf_cc_session

Provides a resource to create a NGWAF CC session.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_cc_session" "location_example" {
  domain           = "example.com"
  source           = "cookie"
  category         = "location"
  key_or_start_mat = "session_id"
  start_offset     = "3"
  end_offset       = "35"
  edition          = "clb-waf"
  session_name     = "session-from-cookie"
}
```

## Argument Reference

The following arguments are supported:

* `category` - (Required, String) Session match pattern, Optional patterns are match, location.
* `domain` - (Required, String) Domain.
* `edition` - (Required, String) Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
* `end_mat` - (Required, String) Session end identifier, when Category is match.
* `end_offset` - (Required, String) End offset position, when Category is location.
* `key_or_start_mat` - (Required, String) Session identifier.
* `session_name` - (Required, String) Session Name.
* `source` - (Required, String) Session matching position, Optional locations are get, post, header, cookie.
* `start_offset` - (Required, String) Starting offset position, when Category is location.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `session_id` - Session ID.

## Import

tencentcloudenterprise_ngwaf_cc_session can be imported using the id, e.g.

```
NGWAF CC session can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_cc_session.location_example example.com#clb-waf#session-xxxxxxxx
```
```

