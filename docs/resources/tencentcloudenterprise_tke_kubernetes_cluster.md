---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_cluster"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_cluster"
description: |-
  Provide a resource to create a kubernetes cluster.
---

# tencentcloudenterprise_tke_kubernetes_cluster

Provide a resource to create a kubernetes cluster.

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

~> **NOTE:** We recommend this usage that uses the `tencentcloudenterprise_tke_kubernetes_cluster` resource to create a cluster without any `worker_config`, then adds nodes by the `tencentcloudenterprise_kubernetes_node_pool` resource.
It's more flexible than managing worker config directly with `tencentcloudenterprise_tke_kubernetes_cluster`, `tencentcloudenterprise_tke_kubernetes_scale_worker`, or existing node management of `tencentcloudenterprise_kubernetes_attachment`. The reason is that `worker_config` is unchangeable and may cause the whole cluster resource to `ForceNew`.

## Example Usage

```hcl
variable "availability_zone_first" {
  default = "ap-guangzhou-3"
}

variable "availability_zone_second" {
  default = "ap-guangzhou-4"
}

variable "cluster_cidr" {
  default = "10.31.0.0/16"
}

variable "default_instance_type" {
  default = "SA2.2XLARGE16"
}

data "tencentcloudenterprise_vpc_subnets" "vpc_first" {
  is_default        = true
  availability_zone = var.availability_zone_first
}

data "tencentcloudenterprise_vpc_subnets" "vpc_second" {
  is_default        = true
  availability_zone = var.availability_zone_second
}

resource "tencentcloudenterprise_tke_kubernetes_cluster" "managed_cluster" {
  vpc_id                  = data.tencentcloudenterprise_vpc_subnets.vpc_first.instance_list.0.vpc_id
  cluster_cidr            = var.cluster_cidr
  cluster_max_pod_num     = 32
  cluster_name            = "test"
  cluster_desc            = "test cluster desc"
  cluster_max_service_num = 32
  cluster_deploy_type     = "MANAGED_CLUSTER"

  worker_config {
    count                      = 1
    availability_zone          = var.availability_zone_first
    instance_type              = var.default_instance_type
    system_disk_type           = "CLOUD_SSD"
    system_disk_size           = 60
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 100
    public_ip_assigned         = true
    subnet_id                  = data.tencentcloudenterprise_vpc_subnets.vpc_first.instance_list.0.subnet_id
    img_id                     = "img-rkiynh11"

    data_disk {
      disk_type = "CLOUD_PREMIUM"
      disk_size = 50
    }

    enhanced_security_service = false
    enhanced_monitor_service  = false
    user_data                 = "dGVzdA=="
    # password                  = "ZZXXccvv1212" // Optional, should be set if key_ids not set.
    key_ids            = ["skey-11112222"]
    security_group_ids = ["sg-xxxxxxxx"]
  }

  worker_config {
    count                      = 1
    availability_zone          = var.availability_zone_second
    instance_type              = var.default_instance_type
    system_disk_type           = "CLOUD_SSD"
    system_disk_size           = 60
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 100
    public_ip_assigned         = true
    subnet_id                  = data.tencentcloudenterprise_vpc_subnets.vpc_second.instance_list.0.subnet_id

    data_disk {
      disk_type = "CLOUD_PREMIUM"
      disk_size = 50
    }

    enhanced_security_service = false
    enhanced_monitor_service  = false
    user_data                 = "dGVzdA=="
    # password                  = "ZZXXccvv1212" // Optional, should be set if key_ids not set.
    key_ids            = ["skey-11112222"]
    security_group_ids = ["sg-xxxxxxxx"]
  }

  tags = {
    "test1" = "test1"
    "test2" = "test2"
  }
}
```

### # Use Kubelet

```hcl
variable "availability_zone_first" {
  default = "ap-guangzhou-3"
}

variable "availability_zone_second" {
  default = "ap-guangzhou-4"
}

variable "cluster_cidr" {
  default = "10.31.0.0/16"
}

variable "default_instance_type" {
  default = "SA2.2XLARGE16"
}

data "tencentcloudenterprise_vpc_subnets" "vpc_first" {
  is_default        = true
  availability_zone = var.availability_zone_first
}

data "tencentcloudenterprise_vpc_subnets" "vpc_second" {
  is_default        = true
  availability_zone = var.availability_zone_second
}

resource "tencentcloudenterprise_tke_kubernetes_cluster" "managed_cluster" {
  vpc_id                  = data.tencentcloudenterprise_vpc_subnets.vpc_first.instance_list.0.vpc_id
  cluster_cidr            = var.cluster_cidr
  cluster_max_pod_num     = 32
  cluster_name            = "test"
  cluster_desc            = "test cluster desc"
  cluster_max_service_num = 32
  cluster_deploy_type     = "MANAGED_CLUSTER"

  worker_config {
    count                      = 1
    availability_zone          = var.availability_zone_first
    instance_type              = var.default_instance_type
    system_disk_type           = "CLOUD_SSD"
    system_disk_size           = 60
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 100
    public_ip_assigned         = true
    subnet_id                  = data.tencentcloudenterprise_vpc_subnets.vpc_first.instance_list.0.subnet_id

    data_disk {
      disk_type = "CLOUD_PREMIUM"
      disk_size = 50
    }

    enhanced_security_service = false
    enhanced_monitor_service  = false
    user_data                 = "dGVzdA=="
    # password                  = "ZZXXccvv1212" // Optional, should be set if key_ids not set.
    key_ids            = ["skey-11112222"]
    security_group_ids = ["sg-xxxxxxxx"]
  }

  worker_config {
    count                      = 1
    availability_zone          = var.availability_zone_second
    instance_type              = var.default_instance_type
    system_disk_type           = "CLOUD_SSD"
    system_disk_size           = 60
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 100
    public_ip_assigned         = true
    subnet_id                  = data.tencentcloudenterprise_vpc_subnets.vpc_second.instance_list.0.subnet_id

    data_disk {
      disk_type = "CLOUD_PREMIUM"
      disk_size = 50
    }

    enhanced_security_service = false
    enhanced_monitor_service  = false
    user_data                 = "dGVzdA=="
    # password                  = "ZZXXccvv1212" // Optional, should be set if key_ids not set.
    key_ids            = ["skey-11112222"]
    security_group_ids = ["sg-xxxxxxxx"]
  }

  tags = {
    "test1" = "test1"
    "test2" = "test2"
  }

  extra_args = [
    "root-dir=/var/lib/kubelet"
  ]
}
```

### Using VPC-CNI network type

```hcl
variable "availability_zone" {
  default = "ap-guangzhou-1"
}

variable "vpc" {
  default = "vpc-r1m1fyx5"
}

variable "default_instance_type" {
  default = "SA2.SMALL2"
}

resource "tencentcloudenterprise_tke_kubernetes_cluster" "managed_cluster" {
  vpc_id                  = var.vpc
  cluster_max_pod_num     = 32
  cluster_name            = "test"
  cluster_desc            = "test cluster desc"
  cluster_max_service_num = 256
  cluster_deploy_type     = "MANAGED_CLUSTER"
  network_type            = "VPC-CNI"
  eni_subnet_ids          = ["subnet-bk1etlyu"]
  service_cidr            = "10.1.0.0/24"

  worker_config {
    count                      = 1
    availability_zone          = var.availability_zone
    instance_type              = var.default_instance_type
    system_disk_type           = "CLOUD_PREMIUM"
    system_disk_size           = 60
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 100
    public_ip_assigned         = true
    subnet_id                  = "subnet-t5dv27rs"

    data_disk {
      disk_type = "CLOUD_PREMIUM"
      disk_size = 50
    }

    enhanced_security_service = false
    enhanced_monitor_service  = false
    user_data                 = "dGVzdA=="
    # password                  = "ZZXXccvv1212" // Optional, should be set if key_ids not set.
    key_ids            = ["skey-11112222"]
    security_group_ids = ["sg-xxxxxxxx"]
  }

  tags = {
    "test1" = "test1"
    "test2" = "test2"
  }
}
```

### # Independent Cluster with master_config

```hcl
variable "availability_zone" {
  default = "ap-guangzhou-3"
}

variable "vpc_id" {
  default = "vpc-xxxxxxxx"
}

variable "subnet_id" {
  default = "subnet-xxxxxxxx"
}

resource "tencentcloudenterprise_tke_kubernetes_cluster" "independent_cluster" {
  vpc_id              = var.vpc_id
  cluster_cidr        = "172.16.0.0/16"
  cluster_name        = "my-independent-cluster"
  cluster_desc        = "Independent cluster with master_config"
  cluster_deploy_type = "INDEPENDENT_CLUSTER"
  cluster_version     = "1.28.3"

  master_config {
    instance_type      = "S5.LARGE8"
    availability_zone  = var.availability_zone
    subnet_id          = var.subnet_id
    system_disk_type   = "CLOUD_SSD"
    system_disk_size   = 100
    security_group_ids = ["sg-xxxxxxxx"]
    password           = "YourPassword123!"

    data_disk {
      disk_type = "CLOUD_SSD"
      disk_size = 100
    }
  }

  master_config {
    instance_type      = "S5.LARGE8"
    availability_zone  = var.availability_zone
    subnet_id          = var.subnet_id
    system_disk_type   = "CLOUD_SSD"
    system_disk_size   = 100
    security_group_ids = ["sg-xxxxxxxx"]
    password           = "YourPassword123!"

    data_disk {
      disk_type = "CLOUD_SSD"
      disk_size = 100
    }
  }

  master_config {
    instance_type      = "S5.LARGE8"
    availability_zone  = var.availability_zone
    subnet_id          = var.subnet_id
    system_disk_type   = "CLOUD_SSD"
    system_disk_size   = 100
    security_group_ids = ["sg-xxxxxxxx"]
    password           = "YourPassword123!"

    data_disk {
      disk_type = "CLOUD_SSD"
      disk_size = 100
    }
  }
}
```

### for split-role deployments.

```hcl
resource "tencentcloudenterprise_tke_kubernetes_cluster" "independent_cluster" {
  vpc_id              = var.vpc_id
  cluster_cidr        = "172.16.0.0/16"
  cluster_name        = "my-independent-cluster"
  cluster_deploy_type = "INDEPENDENT_CLUSTER"

  # Existing 3 masters
  master_config {
    instance_type      = "S5.LARGE8"
    availability_zone  = "ap-guangzhou-3"
    subnet_id          = "subnet-xxxxxxxx"
    system_disk_type   = "CLOUD_SSD"
    system_disk_size   = 100
    security_group_ids = ["sg-xxxxxxxx"]
    password           = "YourPassword123!"
  }

  master_config {
    instance_type      = "S5.LARGE8"
    availability_zone  = "ap-guangzhou-3"
    subnet_id          = "subnet-xxxxxxxx"
    system_disk_type   = "CLOUD_SSD"
    system_disk_size   = 100
    security_group_ids = ["sg-xxxxxxxx"]
    password           = "YourPassword123!"
  }

  master_config {
    instance_type      = "S5.LARGE8"
    availability_zone  = "ap-guangzhou-3"
    subnet_id          = "subnet-xxxxxxxx"
    system_disk_type   = "CLOUD_SSD"
    system_disk_size   = 100
    security_group_ids = ["sg-xxxxxxxx"]
    password           = "YourPassword123!"
  }

  # Scale out: add 2 more master nodes
  master_config {
    instance_type      = "S5.LARGE8"
    availability_zone  = "ap-guangzhou-4"
    subnet_id          = "subnet-yyyyyyyy"
    system_disk_type   = "CLOUD_SSD"
    system_disk_size   = 100
    security_group_ids = ["sg-xxxxxxxx"]
    password           = "YourPassword123!"
    node_role          = "MASTER_ETCD"
  }

  master_config {
    instance_type      = "S5.LARGE8"
    availability_zone  = "ap-guangzhou-4"
    subnet_id          = "subnet-yyyyyyyy"
    system_disk_type   = "CLOUD_SSD"
    system_disk_size   = 100
    security_group_ids = ["sg-xxxxxxxx"]
    password           = "YourPassword123!"
    node_role          = "MASTER_ETCD"
  }
}
```

## Argument Reference

The following arguments are supported:

* `vpc_id` - (Required, String, ForceNew) Vpc Id of the cluster.
* `audit_enabled` - (Optional, Bool) Indicates whether cluster audit is enabled. Default is false.
* `audit_log_topic_id` - (Optional, String) Audit log topic ID.
* `audit_logset_id` - (Optional, String) Audit log logset ID.
* `auto_upgrade_cluster_level` - (Optional, Bool) Whether the cluster level auto upgraded, valid for managed cluster.
* `base_pod_num` - (Optional, Int, ForceNew) The number of basic pods. valid when enable_customized_pod_cidr=true.
* `claim_expired_seconds` - (Optional, Int) Claim expired seconds to recycle ENI. This field can only set when field `network_type` is 'VPC-CNI'. `claim_expired_seconds` must greater or equal than 300 and less than 15768000.
* `cluster_arch` - (Optional, String) Cpu arch of the cluster.
* `cluster_audit` - (Optional, List) Specify Cluster Audit config. NOTE: Please make sure your TKE CamRole have permission to access CLS service.
* `cluster_cidr` - (Optional, String) A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this vpc. Must be in  10./192.168/172.[16-31] segments.
* `cluster_deploy_type` - (Optional, String, ForceNew) Deployment type of the cluster, the available values include: 'MANAGED_CLUSTER' and 'INDEPENDENT_CLUSTER'. Default is 'MANAGED_CLUSTER'.
* `cluster_desc` - (Optional, String) Description of the cluster.
* `cluster_extra_args` - (Optional, List, ForceNew) Customized parameters for master component,such as kube-apiserver, kube-controller-manager, kube-scheduler.
* `cluster_ipvs` - (Optional, Bool, ForceNew) Indicates whether `ipvs` is enabled. Default is true. False means `iptables` is enabled.
* `cluster_level` - (Optional, String) Specify cluster level, valid for managed cluster, use data source `tencentcloudenterprise_kubernetes_cluster_levels` to query available levels. Available value examples `L5`, `L20`, `L50`, `L100`, etc.
* `cluster_max_pod_num` - (Optional, Int, ForceNew) The maximum number of Pods per node in the cluster. Default is 256. The minimum value is 4. When its power unequal to 2, it will round upward to the closest power of 2.
* `cluster_max_service_num` - (Optional, Int, ForceNew) The maximum number of services in the cluster. Default is 256. The range is from 32 to 32768. When its power unequal to 2, it will round upward to the closest power of 2.
* `cluster_name` - (Optional, String) Name of the cluster.
* `cluster_os_type` - (Optional, String, ForceNew) Image type of the cluster os, the available values include: 'GENERAL'. Default is 'GENERAL'.
* `cluster_os` - (Optional, String, ForceNew) Cluster operating system, supports setting public images (image Name) and custom images (image ID).
* `cluster_subnet_id` - (Optional, String, ForceNew) Control Plane Subnet Information. Required for some network plugins (for example, CiliumOverlay).
* `cluster_version` - (Optional, String) Version of the cluster, Default is '1.10.5'. Use `tencentcloudenterprise_tke_kubernetes_available_cluster_versions` to get the available versions.
* `container_runtime` - (Optional, String, ForceNew) Runtime type of the cluster, the available values include: 'docker' and 'containerd'.The Kubernetes v1.24 has removed dockershim, so please use containerd in v1.24 or higher.Default is 'docker'.
* `data_plane_v2` - (Optional, Bool, ForceNew) Whether to use data plane V2 (cilium v2). Default is false.
* `deletion_protection` - (Optional, Bool) Indicates whether cluster deletion protection is enabled. Default is false.
* `disable_addons` - (Optional, List: [`String`], ForceNew) List of add-on names to disable during cluster creation. For example: ["ip-masq-agent"].
* `docker_graph_path` - (Optional, String, ForceNew) Docker graph path. Default is `/var/lib/docker`.
* `enable_customized_pod_cidr` - (Optional, Bool, ForceNew) Whether to enable the custom mode of node podCIDR size. Default is false.
* `eni_subnet_ids` - (Optional, List: [`String`]) Subnet Ids for cluster with VPC-CNI network mode. This field can only set when field `network_type` is 'VPC-CNI'. `eni_subnet_ids` can not empty once be set.
* `event_persistence` - (Optional, List) Specify cluster Event Persistence config. NOTE: Please make sure your TKE CamRole have permission to access CLS service.
* `exist_instance` - (Optional, List, ForceNew) Create tke cluster by existed instances.
* `extension_addon` - (Optional, List, ForceNew) List of extension add-ons to be installed.
* `extra_args` - (Optional, List: [`String`], ForceNew) Custom parameter information related to the node.
* `globe_desired_pod_num` - (Optional, Int, ForceNew) Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it takes effect for all nodes.
* `ignore_cluster_cidr_conflict` - (Optional, Bool, ForceNew) Indicates whether to ignore the cluster cidr conflict error. Default is false.
* `ignore_service_cidr_conflict` - (Optional, Bool, ForceNew) Indicates whether to ignore the service cidr conflict error. Only valid in VPC-CNI mode. Default is false.
* `instance_data_disks` - (Optional, List, ForceNew) Data disks for instance advanced settings.
* `is_dual_stack` - (Optional, Bool, ForceNew) Indicates whether the cluster is dual stack (IPv4/IPv6). Default is false.
* `is_non_static_ip_mode` - (Optional, Bool, ForceNew) Indicates whether non-static ip mode is enabled. Default is false.
* `kube_proxy_mode` - (Optional, String, ForceNew) Cluster kube-proxy mode, the available values include: 'kube-proxy-bpf'. Default is not set.When set to kube-proxy-bpf, cluster version greater than 1.14 and with Linux 2.4 is required.
* `labels` - (Optional, Map, ForceNew) Labels of tke cluster nodes.
* `log_agent` - (Optional, List) Specify cluster log agent config.
* `master_config` - (Optional, List) Deploy the machine configuration information of the 'MASTER_ETCD' service, and create <=7 units for common users.
* `mount_target` - (Optional, String, ForceNew) Mount target. Default is not mounting.
* `need_work_security_group` - (Optional, Bool, ForceNew) Indicates whether to enable the default node security group. Default is false.
* `network_type` - (Optional, String, ForceNew) Cluster network type, GR or VPC-CNI. Default is GR.
* `node_name_type` - (Optional, String, ForceNew) Node name type of Cluster, the available values include: 'lan-ip' and 'hostname', Default is 'lan-ip'.
* `pre_start_user_script` - (Optional, String, ForceNew) Base64-encoded user script, executed before initializing the node, currently only effective for adding existing nodes.
* `project_id` - (Optional, Int) Project ID, default value is 0.
* `qgpu_share_enable` - (Optional, Bool, ForceNew) Indicates whether QGPU sharing is enabled. Default is false.
* `run_instances_for_node` - (Optional, List, ForceNew) RunInstancesForNode settings to build CreateCluster request directly.
* `runtime_version` - (Optional, String, ForceNew) Container runtime version.
* `service_cidr` - (Optional, String, ForceNew) A network address block of the service. Different from vpc cidr and cidr of other clusters within this vpc. Must be in  10./192.168/172.[16-31] segments.
* `tags` - (Optional, Map) The tags of the cluster.
* `taints` - (Optional, List, ForceNew) Node taint.
* `unschedulable` - (Optional, Int, ForceNew) Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
* `user_script` - (Optional, String, ForceNew) Base64-encoded user script executed after initializing the node.
* `vpc_cni_type` - (Optional, String) Distinguish between shared network card multi-IP mode and independent network card mode. Fill in `tke-route-eni` for shared network card multi-IP mode and `tke-direct-eni` for independent network card mode. The default is shared network card mode.
* `worker_config` - (Optional, List, ForceNew) Deploy the machine configuration information of the 'WORKER' service, and create <=20 units for common users. The other 'WORK' service are added by 'tencentcloudenterprise_kubernetes_worker'.

The `cluster_audit` object supports the following:

* `enabled` - (Required, Bool) Specify weather the Cluster Audit enabled. NOTE: Enable Cluster Audit will also auto install Log Agent.
* `delete_audit_log_and_topic` - (Optional, Bool) When you want to close the cluster audit log or delete the cluster, you can use this parameter to determine whether the audit log set and topic created by default will be deleted.
* `log_set_id` - (Optional, String) Specify id of existing CLS log set, or auto create a new set by leave it empty.
* `topic_id` - (Optional, String) Specify id of existing CLS log topic, or auto create a new topic by leave it empty.

The `cluster_extra_args` object supports the following:

* `kube_apiserver` - (Optional, List, ForceNew) The customized parameters for kube-apiserver.
* `kube_controller_manager` - (Optional, List, ForceNew) The customized parameters for kube-controller-manager.
* `kube_scheduler` - (Optional, List, ForceNew) The customized parameters for kube-scheduler.

The `data_disk` object supports the following:

* `disk_id` - (Required, String, ForceNew) Data disk ID.
* `disk_partition` - (Required, String, ForceNew) The device or partition name to mount.
* `mount_target` - (Required, String, ForceNew) Mount target.
* `auto_format_and_mount` - (Optional, Bool, ForceNew) Indicate whether to auto format and mount or not. Default is `false`.
* `disk_pool_group` - (Optional, String, ForceNew) disk pool group.
* `disk_size` - (Optional, Int, ForceNew) Volume of disk in GB. Default is `0`.
* `disk_type` - (Optional, String, ForceNew) Types of disk, available values: `CLOUD_PREMIUM` and `CLOUD_SSD` and `CLOUD_HSSD` and `CLOUD_TSSD`.
* `file_system` - (Optional, String, ForceNew) File system, e.g. `ext3/ext4/xfs`.
* `snapshot_id` - (Optional, String, ForceNew) Data disk snapshot ID.

The `data_disks` object supports the following:

* `mount_target` - (Required, String) Mount target.
* `auto_format_and_mount` - (Optional, Bool) Indicate whether to auto format and mount or not.
* `disk_id` - (Optional, String) Data disk ID.
* `disk_partition` - (Optional, String) The name of the device or partition to mount.
* `disk_size` - (Optional, Int) Volume of disk in GB.
* `disk_type` - (Optional, String) Types of disk.
* `file_system` - (Optional, String) File system, e.g. `ext3/ext4/xfs`.

The `enhanced_service` object supports the following:

* `automation_service` - (Optional, Bool, ForceNew) Enable automation service.
* `monitor_service` - (Optional, Bool, ForceNew) Enable cloud monitor service.
* `security_service` - (Optional, Bool, ForceNew) Enable cloud security service.

The `event_persistence` object supports the following:

* `enabled` - (Required, Bool) Specify weather the Event Persistence enabled.
* `delete_event_log_and_topic` - (Optional, Bool) When you want to close the cluster event persistence or delete the cluster, you can use this parameter to determine whether the event persistence log set and topic created by default will be deleted.
* `log_set_id` - (Optional, String) Specify id of existing CLS log set, or auto create a new set by leave it empty.
* `topic_id` - (Optional, String) Specify id of existing CLS log topic, or auto create a new topic by leave it empty.

The `exist_instance` object supports the following:

* `desired_pod_numbers` - (Optional, List, ForceNew) Custom mode cluster, you can specify the number of pods for each node. corresponding to the existed_instances_para.instance_ids parameter.
* `instances_para` - (Optional, List, ForceNew) Reinstallation parameters of an existing instance.
* `node_role` - (Optional, String, ForceNew) Role of existed node. value:MASTER_ETCD or WORKER.

The `extension_addon` object supports the following:

* `name` - (Required, String) Add-on name.
* `param` - (Optional, String) Add-on parameters (JSON string format), please check the example at the top of page for reference.

The `gpu_args` object supports the following:

* `cuda` - (Optional, Map) CUDA  version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.
* `cudnn` - (Optional, Map) cuDNN version. Format like: `{ version: String, name: String, doc_name: String, dev_name: String }`. `version`: cuDNN version; `name`: cuDNN name; `doc_name`: Doc name of cuDNN; `dev_name`: Dev name of cuDNN.
* `custom_driver` - (Optional, Map) Custom GPU driver. Format like: `{address: String}`. `address`: URL of custom GPU driver address.
* `driver` - (Optional, Map) GPU driver version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.
* `mig_enable` - (Optional, Bool) Whether to enable MIG.

The `instance_advanced_settings_overrides` object supports the following:

* `data_disks` - (Optional, List) Data disks for instance advanced settings override.
* `desired_pod_number` - (Optional, Int) Indicate to set desired pod number in node. valid when the cluster is podCIDR.
* `docker_graph_path` - (Optional, String) Docker graph path.
* `extra_args` - (Optional, List) Custom parameter information related to the node.
* `gpu_args` - (Optional, List) GPU driver parameters.
* `labels` - (Optional, List) Node label list.
* `pre_start_user_script` - (Optional, String) Base64-encoded user script, executed before initializing the node.
* `taints` - (Optional, List) Node taint.
* `unschedulable` - (Optional, Int) Sets whether the joining node participates in the schedule.
* `user_script` - (Optional, String) Base64-encoded user script executed after initializing the node.

The `instance_data_disks` object supports the following:

* `auto_format_and_mount` - (Optional, Bool) Indicate whether to auto format and mount or not.
* `disk_id` - (Optional, String) Data disk ID.
* `disk_partition` - (Optional, String) The name of the device or partition to mount.
* `disk_size` - (Optional, Int) Volume of disk in GB.
* `disk_type` - (Optional, String) Types of disk.
* `file_system` - (Optional, String) File system, e.g. `ext3/ext4/xfs`.
* `mount_target` - (Optional, String) Mount target.

The `instances_para` object supports the following:

* `instance_ids` - (Required, List, ForceNew) Cluster IDs.
* `enhanced_service` - (Optional, List, ForceNew) Enhanced service settings.
* `host_name` - (Optional, String, ForceNew) Host name of the instance.
* `security_group_ids` - (Optional, List, ForceNew) Security groups to which a CVM instance belongs.
* `skip_options` - (Optional, List, ForceNew) Skip options for existing instances.

The `labels` object supports the following:

* `name` - (Optional, String) Label name.
* `value` - (Optional, String) Label value.

The `log_agent` object supports the following:

* `enabled` - (Required, Bool) Whether the log agent enabled.
* `kubelet_root_dir` - (Optional, String) Kubelet root directory as the literal.

The `master_config` object supports the following:

* `instance_type` - (Required, String) Specified types of CVM instance.
* `subnet_id` - (Required, String) Private network ID.
* `availability_zone` - (Optional, String) Indicates which availability zone will be used.
* `cam_role_name` - (Optional, String) CAM role name authorized to access.
* `data_disk` - (Optional, List) Configurations of data disk.
* `desired_pod_num` - (Optional, Int) Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it override `[globe_]desired_pod_num` for current node. Either all the fields `desired_pod_num` or none.
* `enhanced_automation_service` - (Optional, Bool) To specify whether to enable automation service. Default is TRUE.
* `enhanced_monitor_service` - (Optional, Bool) To specify whether to enable cloud monitor service. Default is TRUE.
* `enhanced_security_service` - (Optional, Bool) To specify whether to enable cloud security service. Default is TRUE.
* `hostname` - (Optional, String) The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
* `img_id` - (Optional, String) The valid image id, format of img-xxx.
* `instance_charge_type_prepaid_period` - (Optional, Int) The tenancy (time unit is month) of the prepaid instance. NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
* `instance_charge_type_prepaid_renew_flag` - (Optional, String) Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
* `instance_charge_type` - (Optional, String) The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`. Note: Cloud International only supports `POSTPAID_BY_HOUR`, `PREPAID` instance will not terminated after cluster deleted, and may not allow to delete before expired.
* `instance_name` - (Optional, String) Name of the CVMs.
* `internet_charge_type` - (Optional, String) Charge types for network traffic. Available values include `TRAFFIC_POSTPAID_BY_HOUR`.
* `internet_max_bandwidth_out` - (Optional, Int) Max bandwidth of Internet access in Mbps. Default is 0.
* `key_ids` - (Optional, List) ID list of keys, should be set if `password` not set.
* `node_role` - (Optional, String) The role of the node. Valid values: `MASTER_ETCD` (default), `MASTER`, `ETCD`.
* `password` - (Optional, String) Password to access, should be set if `key_ids` not set.
* `pre_start_user_script` - (Optional, String) Base64-encoded user script executed before TKE initializes the node. All master configurations must use the same value during cluster creation; newly added masters may use different values during scale-out.
* `public_ip_assigned` - (Optional, Bool) Specify whether to assign an Internet IP address.
* `security_group_ids` - (Optional, List) Security groups to which a CVM instance belongs.
* `system_disk_pool_group` - (Optional, String) System disk pool group.
* `system_disk_size` - (Optional, Int) Volume of system disk in GB. Default is `50`.
* `system_disk_type` - (Optional, String) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage. NOTE: `CLOUD_BASIC`, `LOCAL_BASIC` and `LOCAL_SSD` are deprecated.
* `user_data` - (Optional, String) Ase64-encoded User Data text, the length limit is 16KB.
* `user_script` - (Optional, String) Base64-encoded user script executed after TKE initializes the node. All master configurations must use the same value during cluster creation; newly added masters may use different values during scale-out.

The `run_instances_for_node` object supports the following:

* `node_role` - (Required, String) Node role, value: MASTER_ETCD or WORKER.
* `run_instances_para` - (Required, List) CVM run instances parameters in JSON string format.
* `instance_advanced_settings_overrides` - (Optional, List) Per-node instance advanced settings overrides. Order must match run_instances_para.

The `taints` object supports the following:

* `effect` - (Optional, String) Effect of the taint.
* `key` - (Optional, String) Key of the taint.
* `value` - (Optional, String) Value of the taint.

The `taints` object supports the following:

* `effect` - (Optional, String) Effect of the taint. Valid values are: `NoSchedule`, `PreferNoSchedule`, `NoExecute`.
* `key` - (Optional, String) Key of the taint.
* `value` - (Optional, String) Value of the taint.

The `worker_config` object supports the following:

* `instance_type` - (Required, String, ForceNew) Specified types of CVM instance.
* `subnet_id` - (Required, String, ForceNew) Private network ID.
* `availability_zone` - (Optional, String, ForceNew) Indicates which availability zone will be used.
* `cam_role_name` - (Optional, String, ForceNew) CAM role name authorized to access.
* `count` - (Optional, Int, ForceNew) Number of cvm.
* `data_disk` - (Optional, List, ForceNew) Configurations of data disk.
* `desired_pod_num` - (Optional, Int, ForceNew) Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it override `[globe_]desired_pod_num` for current node. Either all the fields `desired_pod_num` or none.
* `enhanced_automation_service` - (Optional, Bool, ForceNew) To specify whether to enable automation service. Default is TRUE.
* `enhanced_monitor_service` - (Optional, Bool, ForceNew) To specify whether to enable cloud monitor service. Default is TRUE.
* `enhanced_security_service` - (Optional, Bool, ForceNew) To specify whether to enable cloud security service. Default is TRUE.
* `hostname` - (Optional, String, ForceNew) The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
* `img_id` - (Optional, String) The valid image id, format of img-xxx.
* `instance_charge_type_prepaid_period` - (Optional, Int, ForceNew) The tenancy (time unit is month) of the prepaid instance. NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
* `instance_charge_type_prepaid_renew_flag` - (Optional, String, ForceNew) Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
* `instance_charge_type` - (Optional, String, ForceNew) The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`. Note: Cloud International only supports `POSTPAID_BY_HOUR`, `PREPAID` instance will not terminated after cluster deleted, and may not allow to delete before expired.
* `instance_name` - (Optional, String, ForceNew) Name of the CVMs.
* `internet_charge_type` - (Optional, String, ForceNew) Charge types for network traffic. Available values include `TRAFFIC_POSTPAID_BY_HOUR`.
* `internet_max_bandwidth_out` - (Optional, Int) Max bandwidth of Internet access in Mbps. Default is 0.
* `key_ids` - (Optional, List, ForceNew) ID list of keys, should be set if `password` not set.
* `password` - (Optional, String, ForceNew) Password to access, should be set if `key_ids` not set.
* `public_ip_assigned` - (Optional, Bool, ForceNew) Specify whether to assign an Internet IP address.
* `security_group_ids` - (Optional, List, ForceNew) Security groups to which a CVM instance belongs.
* `system_disk_pool_group` - (Optional, String, ForceNew) System disk pool group.
* `system_disk_size` - (Optional, Int, ForceNew) Volume of system disk in GB. Default is `50`.
* `system_disk_type` - (Optional, String, ForceNew) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage. NOTE: `CLOUD_BASIC`, `LOCAL_BASIC` and `LOCAL_SSD` are deprecated.
* `user_data` - (Optional, String, ForceNew) Ase64-encoded User Data text, the length limit is 16KB.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `certification_authority` - The certificate used for access.
* `cluster_external_endpoint` - External network address to access.
* `cluster_node_num` - Number of nodes in the cluster.
* `domain` - Domain name for access.
* `kube_config_intranet` - Kubernetes config of private network.
* `kube_config` - Kubernetes config.
* `password` - Password of account.
* `pgw_endpoint` - The Intranet address used for access.
* `security_policy` - Access policy.
* `user_name` - User name of account.
* `worker_instances_list` - An information list of cvm within the 'WORKER' clusters. Each element contains the following attributes:
  * `failed_reason` - Information of the cvm when it is failed.
  * `instance_id` - ID of the cvm.
  * `instance_role` - Role of the cvm.
  * `instance_state` - State of the cvm.
  * `lan_ip` - LAN IP of the cvm.

