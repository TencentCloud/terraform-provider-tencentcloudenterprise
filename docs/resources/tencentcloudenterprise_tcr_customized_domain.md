---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_customized_domain"
sidebar_current: "docs-tencentcloudenterprise-resource-tcr_customized_domain"
description: |-
  Provides a resource to create a tcr customized_domain
---

# tencentcloudenterprise_tcr_customized_domain

Provides a resource to create a tcr customized_domain

## Example Usage

```hcl
resource "tencentcloudenterprise_tcr_customized_domain" "my_domain" {
  registry_id    = local.tcr_id
  domain_name    = "www.test.com"
  certificate_id = "%s"
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `certificate_id` - (Required, String, ForceNew) Certificate id.
* `domain_name` - (Required, String, ForceNew) Custom domain name.
* `registry_id` - (Required, String, ForceNew) Instance id.
* `tags` - (Optional, Map) Tag description list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_tcr_customized_domain can be imported using the id, e.g.

```
tcr customized_domain can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_tcr_customized_domain.customized_domain customized_domain_id
```
```

