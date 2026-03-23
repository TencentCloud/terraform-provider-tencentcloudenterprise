---
subcategory: "Cloud Connect Network(CCN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ccn_instances_reset_attach"
sidebar_current: "docs-tencentcloudenterprise-resource-ccn_instances_reset_attach"
description: |-
  Provides a resource to reset expired cross-account CCN attachment requests.
---

# tencentcloudenterprise_ccn_instances_reset_attach

Provides a resource to reset expired cross-account CCN attachment requests.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn_instances_reset_attach" "example" {
  ccn_id  = "ccn-39lqkygf"
  ccn_uin = "100022975249"

  instances {
    instance_id     = "vpc-j9yhbzpn"
    instance_region = "ap-guangzhou"
    instance_type   = "VPC"
  }
}
```

## Argument Reference

The following arguments are supported:

* `ccn_id` - (Required, String, ForceNew) CCN instance ID.
* `ccn_uin` - (Required, String, ForceNew) Root account UIN of the target CCN.
* `instances` - (Required, List, ForceNew) List of attachment instances to reset.

The `instances` object supports the following:

* `instance_id` - (Required, String) Attachment instance ID.
* `instance_region` - (Required, String) Region of the attachment instance.
* `instance_type` - (Optional, String) Attachment instance type. Valid values: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


