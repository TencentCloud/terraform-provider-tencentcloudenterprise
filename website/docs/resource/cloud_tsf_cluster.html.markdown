---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_cluster"
sidebar_current: "docs-cloud-resource-tsf_cluster"
description: |-
  Provides a resource to create a tsf cluster
---

# cloud_tsf_cluster

Provides a resource to create a tsf cluster

## Example Usage

```hcl
resource "cloud_tsf_cluster" "cluster" {
  cluster_name  = "terraform-test"
  cluster_type  = "V"
  vpc_id        = "vpc-2l7uk2q1"
  cluster_desc  = "test"
  tsf_region_id = "chongqing"
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name` - (Required, String) Cluster name.
* `cluster_type` - (Required, String) Cluster type. Valid values: `C`: Container cluster; `V`: Virtual machine cluster.
* `vpc_id` - (Required, String) Vpc id.
* `cluster_cidr` - (Optional, String) CIDR assigned to cluster containers and service IP.
* `cluster_desc` - (Optional, String) Cluster notes.
* `cluster_remark_name` - (Optional, String) Cluster remark name.
* `cluster_version` - (Optional, String) Cluster version.
* `max_cluster_service_num` - (Optional, Int) The maximum number of services in the cluster. The value ranges from 32 to 32768. If it is not a power of 2, the nearest power of 2 will be taken up.
* `max_node_pod_num` - (Optional, Int) The maximum number of Pods on each Node in the cluster. The value ranges from 4 to 256. When the value is not a power of 2, the nearest power of 2 will be taken up.
* `program_id` - (Optional, String) The dataset ID to be bound.
* `subnet_id` - (Optional, String) Subnet id.
* `tags` - (Optional, Map) Tag description list.
* `tsf_region_id` - (Optional, String) The TSF region to which the cluster belongs.
* `tsf_zone_id` - (Optional, String) The TSF availability zone to which the cluster belongs.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `abnormal_group_count` - Abnormal number of deployment groups.
* `cluster_id` - Cluster ID.
* `cluster_limit_cpu` - Cluster remaining cpu limit.
* `cluster_limit_mem` - Cluster remaining memory limit.
* `cluster_status` - Cluster status.
* `cluster_total_cpu` - The total CPU of the cluster, unit: core.
* `cluster_total_mem` - The total memory of the cluster, unit: G.
* `cluster_used_cpu` - CPU used by the cluster, unit: core.
* `cluster_used_mem` - The memory used by the cluster, unit: G.
* `create_time` - Create time.
* `delete_flag_reason` - Reasons why clusters cannot be deleted.
* `delete_flag` - Delete flag: `true`: can be deleted; `false`: can not be deleted.
* `group_count` - Total number of deployment groups.
* `instance_count` - Number of cluster machine instances.
* `normal_instance_count` - The number of machine instances in the normal state of the cluster.
* `operation_info` - Control information for buttons on the front end.
  * `add_instance` - Add the control information of the instance button.
    * `disabled_reason` - Reason for not showing.
    * `enabled` - Is the button clickable.
    * `supported` - Whether to show the button.
  * `destroy` - Destroy the control information of the machine.
    * `disabled_reason` - Reason for not showing.
    * `enabled` - Is the button clickable.
    * `supported` - Whether to show the button.
  * `init` - Initialize the control information of the button.
    * `disabled_reason` - Reason for not showing.
    * `enabled` - Is the button clickable.
    * `supported` - Whether to show the button.
* `run_group_count` - Number of Deployment Groups in progress.
* `run_instance_count` - Number of machine instances running in the cluster.
* `run_service_instance_count` - Number of running service instances.
* `stop_group_count` - Number of deployment groups in stop.
* `tsf_region_name` - Name of the TSF region to which the cluster belongs.
* `tsf_zone_name` - The name of the TSF availability zone to which the cluster belongs.
* `update_time` - Update time.


