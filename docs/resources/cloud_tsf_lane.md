---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_lane"
sidebar_current: "docs-cloud-resources-tsf_lane"
description: |-
  Provides a resource to create a tsf lane
---

# cloud_tsf_lane

Provides a resource to create a tsf lane

## Example Usage

```hcl
resource "cloud_tsf_lane" "lane" {
  lane_name = "lane-name-1"
  remark    = "lane desc1"
  lane_group_list {
    group_id = "group-yn7j5l8a"
    entrance = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `lane_group_list` - (Required, List) Swimlane Deployment Group Information.
* `lane_name` - (Required, String) Lane name.
* `remark` - (Required, String) Lane Remarks.

The `lane_group_list` object supports the following:

* `entrance` - (Required, Bool) Whether to enter the application.
* `group_id` - (Required, String) Deployment group ID.
* `application_id` - (Optional, String) Application ID.
* `application_name` - (Optional, String) Application name.
* `cluster_type` - (Optional, String) Cluster type.
* `create_time` - (Optional, Int) Creation time.
* `group_name` - (Optional, String) Deployment group name.
* `lane_group_id` - (Optional, String) Swimlane deployment group ID.
* `lane_id` - (Optional, String) Lane ID.
* `namespace_id` - (Optional, String) Namespace ID.
* `namespace_name` - (Optional, String) Namespace name.
* `update_time` - (Optional, Int) Update time.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time.
* `entrance` - Whether to enter the application.
* `lane_id` - Lane id.
* `namespace_id_list` - A list of namespaces to which the swimlane has associated deployment groups.
* `update_time` - Update time.


