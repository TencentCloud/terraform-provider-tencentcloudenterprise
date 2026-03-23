---
subcategory: "Cloud Connect Network(CCN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ccn_route_table_associate_instance_config"
sidebar_current: "docs-tencentcloudenterprise-resource-ccn_route_table_associate_instance_config"
description: |-
  Provides a resource to manage the instance associations of a CCN route table.
---

# tencentcloudenterprise_ccn_route_table_associate_instance_config

Provides a resource to manage the instance associations of a CCN route table.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn_route_table_associate_instance_config" "example" {
  ccn_id         = "ccn-cwm743gl"
  route_table_id = "ccnrtb-1ydgdxt1"

  instances {
    instance_id   = "vpc-ayl8ggap"
    instance_type = "VPC"
  }
}
```

## Argument Reference

The following arguments are supported:

* `ccn_id` - (Required, String, ForceNew) ID of the CCN.
* `instances` - (Required, Set) Associated instance list.
* `route_table_id` - (Required, String, ForceNew) CCN route table ID.

The `instances` object supports the following:

* `instance_id` - (Required, String) Instance ID.
* `instance_type` - (Required, String) Instance type. Valid values include `VPC`, `DIRECTCONNECT`, `BMVPC`, `EDGE`, `EDGE_TUNNEL`, `EDGE_VPNGW`, `VPNGW`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_ccn_route_table_associate_instance_config can be imported using the id, e.g.

```
CCN route table associate instance config can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ccn_route_table_associate_instance_config.example ccn-cwm743gl#ccnrtb-1ydgdxt1
```
```

