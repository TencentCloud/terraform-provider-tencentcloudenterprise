---
subcategory: "Cloud Connect Network(CCN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ccn_attachment_v2"
sidebar_current: "docs-tencentcloudenterprise-resource-ccn_attachment_v2"
description: |-
  Provides a resource to create a CCN attachment instance.
---

# tencentcloudenterprise_ccn_attachment_v2

Provides a resource to create a CCN attachment instance.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn_attachment_v2" "example" {
  ccn_id          = "ccn-cwm743gl"
  instance_id     = "vpc-l40swg8d"
  instance_type   = "VPC"
  instance_region = "ap-qingyuan-region-devtest-ops"
  route_table_id  = "ccnrtb-gbaaugtl"
  description     = "test ccn attachment description"
}
```

## Argument Reference

The following arguments are supported:

* `ccn_id` - (Required, String, ForceNew) ID of the CCN.
* `instance_id` - (Required, String, ForceNew) ID of the attached instance.
* `instance_region` - (Required, String, ForceNew) The region where the attached instance is located.
* `instance_type` - (Required, String, ForceNew) Type of the attached instance. Valid values: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.
* `ccn_uin` - (Optional, String, ForceNew) UIN of the attached CCN owner. When unset, the current account is used.
* `description` - (Optional, String) Remark of the attachment.
* `route_table_id` - (Optional, String, ForceNew) CCN route table ID bound to the attachment.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `attached_time` - Attachment time.
* `cidr_block` - CIDR blocks of the attached instance.
* `route_ids` - Route ID list under the CCN.
* `state` - State of the attachment.

## Import

tencentcloudenterprise_ccn_attachment_v2 can be imported using the id, e.g.

```
CCN attachment instance can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ccn_attachment_v2.example ccn-cwm743gl#VPC#ap-qingyuan-region-devtest-ops#vpc-l40swg8d
```
```

