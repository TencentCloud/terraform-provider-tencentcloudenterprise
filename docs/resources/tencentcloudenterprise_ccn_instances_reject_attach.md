---
subcategory: "Cloud Connect Network(CCN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ccn_instances_reject_attach"
sidebar_current: "docs-tencentcloudenterprise-resource-ccn_instances_reject_attach"
description: |-
  Provides a resource to reject cross-account CCN attachment requests.
---

# tencentcloudenterprise_ccn_instances_reject_attach

Provides a resource to reject cross-account CCN attachment requests.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn_instances_reject_attach" "example" {
  ccn_id = "ccn-39lqkygf"

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
* `instances` - (Required, List, ForceNew) List of attachment instances to reject.

The `instances` object supports the following:

* `instance_id` - (Required, String) Attachment instance ID.
* `instance_region` - (Required, String) Region of the attachment instance.
* `instance_type` - (Optional, String) Attachment instance type. Valid values: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_ccn_instances_reject_attach can be imported using the id, e.g.

```
CCN instance reject-attach actions can be imported using the CCN ID, e.g.

```bash
terraform import tencentcloudenterprise_ccn_instances_reject_attach.example ccn-39lqkygf
```
```

