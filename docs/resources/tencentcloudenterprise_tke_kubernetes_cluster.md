---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_cluster"
sidebar_current: "docs-tencentcloudenterprise-resources-tke_kubernetes_cluster"
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
  cluster_internet        = true
  # managed_cluster_internet_security_policies = ["3.3.3.3", "1.1.1.1"]
  cluster_deploy_type = "MANAGED_CLUSTER"
  # node_pool_id							   	 = "xxx"

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
    key_ids = "skey-11112222"
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
    key_ids       = "skey-11112222"
    cam_role_name = "CVM_QcsRole"
  }

  labels = {
    "test1" = "test1",
    "test2" = "test2",
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
  cluster_internet        = true
  # managed_cluster_internet_security_policies = ["3.3.3.3", "1.1.1.1"]
  cluster_deploy_type = "MANAGED_CLUSTER"

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
    key_ids = "skey-11112222"
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
    cam_role_name = "CVM_QcsRole"
    key_ids       = "skey-11112222"
  }

  labels = {
    "test1" = "test1",
    "test2" = "test2",
  }

  extra_args = [
    "root-dir=/var/lib/kubelet"
  ]
}
```

### # Use extension addons

```hcl
variable "availability_zone_first" {
  default = "ap-guangzhou-3"
}

variable "cluster_cidr" {
  default = "10.31.0.0/16"
}

variable "default_instance_type" {
  default = "S5.SMALL1"
}

data "tencentcloudenterprise_vpc_subnets" "vpc_first" {
  is_default        = true
  availability_zone = var.availability_zone_first
}

# fetch latest addon(chart) versions
data "tencentcloudenterprise_tke_kubernetes_charts" "charts" {}

locals {
  chartNames    = data.tencentcloudenterprise_tke_kubernetes_charts.charts.chart_list.*.name
  chartVersions = data.tencentcloudenterprise_tke_kubernetes_charts.charts.chart_list.*.latest_version
  chartMap      = zipmap(local.chartNames, local.chartVersions)
}

resource "tencentcloudenterprise_tke_kubernetes_cluster" "cluster_with_addon" {
  vpc_id                  = data.tencentcloudenterprise_vpc_subnets.vpc_first.instance_list.0.vpc_id
  cluster_cidr            = var.cluster_cidr
  cluster_max_pod_num     = 32
  cluster_name            = "test"
  cluster_desc            = "test cluster desc"
  cluster_max_service_num = 32
  cluster_internet        = true
  # managed_cluster_internet_security_policies = ["3.3.3.3", "1.1.1.1"]
  cluster_deploy_type = "MANAGED_CLUSTER"

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
    enhanced_security_service  = false
    enhanced_monitor_service   = false
    user_data                  = "dGVzdA=="
    # password                  = "ZZXXccvv1212" // Optional, should be set if key_ids not set.
    key_ids = "skey-11112222"
  }

  extension_addon {
    name = "COS"
    param = jsonencode({
      "kind" : "App", "spec" : {
        "chart" : { "chartName" : "cos", "chartVersion" : local.chartMap["cos"] },
        "values" : { "values" : [], "rawValues" : "e30=", "rawValuesType" : "json" }
      }
    })
  }
  extension_addon {
    name = "SecurityGroupPolicy"
    param = jsonencode({
      "kind" : "App", "spec" : { "chart" : { "chartName" : "securitygrouppolicy", "chartVersion" : local.chartMap["securitygrouppolicy"] } }
    })
  }
  extension_addon {
    name = "OOMGuard"
    param = jsonencode({
      "kind" : "App", "spec" : { "chart" : { "chartName" : "oomguard", "chartVersion" : local.chartMap["oomguard"] } }
    })
  }
  extension_addon {
    name = "OLM"
    param = jsonencode({
      "kind" : "App", "spec" : { "chart" : { "chartName" : "olm", "chartVersion" : local.chartMap["olm"] } }
    })
  }
}
```

### # Use node pool global config

```hcl
variable "availability_zone" {
  default = "ap-guangzhou-3"
}

variable "vpc" {
  default = "vpc-dk8zmwuf"
}

variable "subnet" {
  default = "subnet-pqfek0t8"
}

variable "default_instance_type" {
  default = "SA1.LARGE8"
}

resource "tencentcloudenterprise_tke_kubernetes_cluster" "test_node_pool_global_config" {
  vpc_id                  = var.vpc
  cluster_cidr            = "10.1.0.0/16"
  cluster_max_pod_num     = 32
  cluster_name            = "test"
  cluster_desc            = "test cluster desc"
  cluster_max_service_num = 32
  cluster_internet        = true
  # managed_cluster_internet_security_policies = ["3.3.3.3", "1.1.1.1"]
  cluster_deploy_type = "MANAGED_CLUSTER"

  worker_config {
    count                      = 1
    availability_zone          = var.availability_zone
    instance_type              = var.default_instance_type
    system_disk_type           = "CLOUD_SSD"
    system_disk_size           = 60
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 100
    public_ip_assigned         = true
    subnet_id                  = var.subnet

    data_disk {
      disk_type = "CLOUD_PREMIUM"
      disk_size = 50
    }

    enhanced_security_service = false
    enhanced_monitor_service  = false
    user_data                 = "dGVzdA=="
    # password                  = "ZZXXccvv1212" // Optional, should be set if key_ids not set.
    key_ids = "skey-11112222"
  }

  node_pool_global_config {
    is_scale_in_enabled            = true
    expander                       = "random"
    ignore_daemon_sets_utilization = true
    max_concurrent_scale_in        = 5
    scale_in_delay                 = 15
    scale_in_unneeded_time         = 15
    scale_in_utilization_threshold = 30
    skip_nodes_with_local_storage  = false
    skip_nodes_with_system_pods    = true
  }

  labels = {
    "test1" = "test1",
    "test2" = "test2",
  }
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
  cluster_internet        = true
  # managed_cluster_internet_security_policies = ["3.3.3.3", "1.1.1.1"]
  cluster_deploy_type = "MANAGED_CLUSTER"
  network_type        = "VPC-CNI"
  eni_subnet_ids      = ["subnet-bk1etlyu"]
  service_cidr        = "10.1.0.0/24"

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
    key_ids = "skey-11112222"
  }

  labels = {
    "test1" = "test1",
    "test2" = "test2",
  }
}
```

### Using ops options

```hcl
resource "tencentcloudenterprise_tke_kubernetes_cluster" "managed_cluster" {
  # ...your basic fields

  log_agent {
    enabled          = true
    kubelet_root_dir = "" # optional
  }

  event_persistence {
    enabled       = true
    log_set_id    = "" # optional
    log_set_topic = "" # optional
  }

  cluster_audit {
    enabled       = true
    log_set_id    = "" # optional
    log_set_topic = "" # optional
  }
}
```

## Argument Reference

The following arguments are supported:

* `cluster_cidr` - (Required, String) A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this vpc. Must be in  10./192.168/172.[16-31] segments.
* `vpc_id` - (Required, String, ForceNew) Vpc Id of the cluster.
* `base_pod_num` - (Optional, Int, ForceNew) The number of basic pods. valid when enable_customized_pod_cidr=true.
* `claim_expired_seconds` - (Optional, Int) Claim expired seconds to recycle ENI. This field can only set when field `network_type` is 'VPC-CNI'. `claim_expired_seconds` must greater or equal than 300 and less than 15768000.
* `cluster_arch` - (Optional, String) Cpu arch of the cluster.
* `cluster_deploy_type` - (Optional, String, ForceNew) Deployment type of the cluster, the available values include: 'MANAGED_CLUSTER' and 'INDEPENDENT_CLUSTER'. Default is 'MANAGED_CLUSTER'.
* `cluster_desc` - (Optional, String) Description of the cluster.
* `cluster_extra_args` - (Optional, List, ForceNew) Customized parameters for master component,such as kube-apiserver, kube-controller-manager, kube-scheduler.
* `cluster_ipvs` - (Optional, Bool, ForceNew) Indicates whether `ipvs` is enabled. Default is true. False means `iptables` is enabled.
* `cluster_level` - (Optional, String) Specify cluster level, valid for managed cluster, use data source `tencentcloudenterprise_kubernetes_cluster_levels` to query available levels. Available value examples `L5`, `L20`, `L50`, `L100`, etc.
* `cluster_max_pod_num` - (Optional, Int, ForceNew) The maximum number of Pods per node in the cluster. Default is 256. The minimum value is 4. When its power unequal to 2, it will round upward to the closest power of 2.
* `cluster_max_service_num` - (Optional, Int, ForceNew) The maximum number of services in the cluster. Default is 256. The range is from 32 to 32768. When its power unequal to 2, it will round upward to the closest power of 2.
* `cluster_name` - (Optional, String) Name of the cluster.
* `cluster_os_type` - (Optional, String, ForceNew) Image type of the cluster os, the available values include: 'GENERAL'. Default is 'GENERAL'.
* `cluster_os` - (Optional, String, ForceNew) Operating system of the cluster, the available values include: 'centos7.6.0_x64','ubuntu18.04.1x86_64','tlinux2.4x86_64'. Default is 'tlinux2.4x86_64'.
* `cluster_version` - (Optional, String) Version of the cluster, Default is '1.10.5'. Use `tencentcloudenterprise_tke_kubernetes_available_cluster_versions` to get the available versions.
* `container_runtime` - (Optional, String, ForceNew) Runtime type of the cluster, the available values include: 'docker' and 'containerd'.The Kubernetes v1.24 has removed dockershim, so please use containerd in v1.24 or higher.Default is 'docker'.
* `docker_graph_path` - (Optional, String, ForceNew) Docker graph path. Default is `/var/lib/docker`.
* `eni_subnet_ids` - (Optional, List: [`String`]) Subnet Ids for cluster with VPC-CNI network mode. This field can only set when field `network_type` is 'VPC-CNI'. `eni_subnet_ids` can not empty once be set.
* `extra_args` - (Optional, List: [`String`], ForceNew) Custom parameter information related to the node.
* `is_non_static_ip_mode` - (Optional, Bool, ForceNew) Indicates whether non-static ip mode is enabled. Default is false.
* `master_config` - (Optional, List, ForceNew) Deploy the machine configuration information of the 'MASTER_ETCD' service, and create <=7 units for common users.
* `mount_target` - (Optional, String, ForceNew) Mount target. Default is not mounting.
* `network_type` - (Optional, String, ForceNew) Cluster network type, GR or VPC-CNI. Default is GR.
* `node_name_type` - (Optional, String, ForceNew) Node name type of Cluster, the available values include: 'lan-ip' and 'hostname', Default is 'lan-ip'.
* `project_id` - (Optional, Int) Project ID, default value is 0.
* `service_cidr` - (Optional, String, ForceNew) A network address block of the service. Different from vpc cidr and cidr of other clusters within this vpc. Must be in  10./192.168/172.[16-31] segments.
* `tags` - (Optional, Map) The tags of the cluster.
* `unschedulable` - (Optional, Int, ForceNew) Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
* `worker_config` - (Optional, List, ForceNew) Deploy the machine configuration information of the 'WORKER' service, and create <=20 units for common users. The other 'WORK' service are added by 'tencentcloudenterprise_kubernetes_worker'.

The `cluster_extra_args` object supports the following:

* `kube_apiserver` - (Optional, List, ForceNew) The customized parameters for kube-apiserver.
* `kube_controller_manager` - (Optional, List, ForceNew) The customized parameters for kube-controller-manager.
* `kube_scheduler` - (Optional, List, ForceNew) The customized parameters for kube-scheduler.

The `data_disk` object supports the following:

* `auto_format_and_mount` - (Optional, Bool, ForceNew) Indicate whether to auto format and mount or not. Default is `false`.
* `disk_pool_group` - (Optional, String, ForceNew) disk pool group
* `disk_size` - (Optional, Int, ForceNew) Volume of disk in GB. Default is `0`.
* `disk_type` - (Optional, String, ForceNew) Types of disk, available values: `CLOUD_PREMIUM` and `CLOUD_SSD` and `CLOUD_HSSD` and `CLOUD_TSSD`.
* `file_system` - (Optional, String, ForceNew) File system, e.g. `ext3/ext4/xfs`.
* `mount_target` - (Optional, String, ForceNew) Mount target.
* `snapshot_id` - (Optional, String, ForceNew) Data disk snapshot ID.

The `master_config` object supports the following:

* `instance_type` - (Required, String, ForceNew) Specified types of CVM instance.
* `security_group_ids` - (Required, List) Security groups to which a CVM instance belongs.
* `subnet_id` - (Required, String, ForceNew) Private network ID.
* `availability_zone` - (Optional, String, ForceNew) Indicates which availability zone will be used.
* `count` - (Optional, Int, ForceNew) Number of cvm.
* `data_disk` - (Optional, List, ForceNew) Configurations of data disk.
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
* `system_disk_pool_group` - (Optional, String, ForceNew) System disk pool group
* `system_disk_size` - (Optional, Int, ForceNew) Volume of system disk in GB. Default is `50`.
* `system_disk_type` - (Optional, String, ForceNew) System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage. NOTE: `CLOUD_BASIC`, `LOCAL_BASIC` and `LOCAL_SSD` are deprecated.
* `user_data` - (Optional, String, ForceNew) Ase64-encoded User Data text, the length limit is 16KB.

The `worker_config` object supports the following:

* `instance_type` - (Required, String, ForceNew) Specified types of CVM instance.
* `security_group_ids` - (Required, List) Security groups to which a CVM instance belongs.
* `subnet_id` - (Required, String, ForceNew) Private network ID.
* `availability_zone` - (Optional, String, ForceNew) Indicates which availability zone will be used.
* `count` - (Optional, Int, ForceNew) Number of cvm.
* `data_disk` - (Optional, List, ForceNew) Configurations of data disk.
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
* `system_disk_pool_group` - (Optional, String, ForceNew) System disk pool group
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


