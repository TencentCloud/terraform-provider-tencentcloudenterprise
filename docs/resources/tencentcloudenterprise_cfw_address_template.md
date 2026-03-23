---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_address_template"
sidebar_current: "docs-tencentcloudenterprise-resource-cfw_address_template"
description: |-
  Provides a resource to create a cloud firewall (cfw) address template.
---

# tencentcloudenterprise_cfw_address_template

Provides a resource to create a cloud firewall (cfw) address template.

## Example Usage

```hcl
resource "tencentcloudenterprise_cfw_address_template" "example" {
  name      = "tf_example"
  detail    = "terraform example"
  ip_string = "1.1.1.1,2.2.2.2"
  type      = 1
}
```

## Argument Reference

The following arguments are supported:

* `detail` - (Required, String) Template Detail.
* `ip_string` - (Required, String) Type is 1, ip template eg: 1.1.1.1,2.2.2.2; Type is 5, domain name template eg: www.qq.com, www.tencent.com.
* `name` - (Required, String) Template name.
* `type` - (Required, Int) 1: ip template; 5: domain name templates.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cfw_address_template can be imported using the id, e.g.

```
Cloud firewall address template can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_address_template.example template_id
```
```

