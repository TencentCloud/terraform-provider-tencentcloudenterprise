---
subcategory: "Cloud Load Balancer(CLB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_clb_security_group_attachment"
sidebar_current: "docs-tencentcloudenterprise-resources-clb_security_group_attachment"
description: |-
  Provides a resource to create a clb security_group_attachment
---

# tencentcloudenterprise_clb_security_group_attachment

Provides a resource to create a clb security_group_attachment

## Example Usage

```hcl
resource "tencentcloudenterprise_clb_security_group_attachment" "security_group_attachment" {
  security_group    = "sg-ijato2x1"
  load_balancer_ids = ["lb-5dnrkgry"]
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_ids` - (Required, Set: [`String`], ForceNew) Array of CLB instance IDs. Only support set one security group now.
* `security_group` - (Required, String, ForceNew) Security group ID, such as esg-12345678.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

clb security_group_attachment can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_clb_security_group_attachment.security_group_attachment security_group_id#clb_id
```

