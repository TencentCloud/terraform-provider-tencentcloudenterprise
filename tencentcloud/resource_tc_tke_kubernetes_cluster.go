/*
Provide a resource to create a kubernetes cluster.

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

~> **NOTE:** We recommend this usage that uses the `tencentcloudenterprise_tke_kubernetes_cluster` resource to create a cluster without any `worker_config`, then adds nodes by the `tencentcloudenterprise_kubernetes_node_pool` resource.
It's more flexible than managing worker config directly with `tencentcloudenterprise_tke_kubernetes_cluster`, `tencentcloudenterprise_tke_kubernetes_scale_worker`, or existing node management of `tencentcloudenterprise_kubernetes_attachment`. The reason is that `worker_config` is unchangeable and may cause the whole cluster resource to `ForceNew`.

# Example Usage

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
	    key_ids                   = ["skey-11112222"]
	    security_group_ids        = ["sg-xxxxxxxx"]
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
	    key_ids                   = ["skey-11112222"]
	    security_group_ids        = ["sg-xxxxxxxx"]
	  }

	  tags = {
	    "test1" = "test1"
	    "test2" = "test2"
	  }
	}

```

# Use Kubelet

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
	    key_ids                   = ["skey-11112222"]
	    security_group_ids        = ["sg-xxxxxxxx"]
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
	    key_ids                   = ["skey-11112222"]
	    security_group_ids        = ["sg-xxxxxxxx"]
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

Using VPC-CNI network type
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
	    key_ids                   = ["skey-11112222"]
	    security_group_ids        = ["sg-xxxxxxxx"]
	  }

	  tags = {
	    "test1" = "test1"
	    "test2" = "test2"
	  }
	}

```

# Independent Cluster with master_config

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
	    instance_type     = "S5.LARGE8"
	    availability_zone = var.availability_zone
	    subnet_id         = var.subnet_id
	    system_disk_type  = "CLOUD_SSD"
	    system_disk_size  = 100
	    security_group_ids = ["sg-xxxxxxxx"]
	    password          = "YourPassword123!"

	    data_disk {
	      disk_type = "CLOUD_SSD"
	      disk_size = 100
	    }
	  }

	  master_config {
	    instance_type     = "S5.LARGE8"
	    availability_zone = var.availability_zone
	    subnet_id         = var.subnet_id
	    system_disk_type  = "CLOUD_SSD"
	    system_disk_size  = 100
	    security_group_ids = ["sg-xxxxxxxx"]
	    password          = "YourPassword123!"

	    data_disk {
	      disk_type = "CLOUD_SSD"
	      disk_size = 100
	    }
	  }

	  master_config {
	    instance_type     = "S5.LARGE8"
	    availability_zone = var.availability_zone
	    subnet_id         = var.subnet_id
	    system_disk_type  = "CLOUD_SSD"
	    system_disk_size  = 100
	    security_group_ids = ["sg-xxxxxxxx"]
	    password          = "YourPassword123!"

	    data_disk {
	      disk_type = "CLOUD_SSD"
	      disk_size = 100
	    }
	  }
	}

```

# Scale out master nodes (add 2 more masters to existing 3-master cluster)

To scale out master nodes, simply add more `master_config` blocks. To scale in, remove the blocks.
The `node_role` field defaults to `MASTER_ETCD`. Use `MASTER` or `ETCD` for split-role deployments.

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster" "independent_cluster" {
	  vpc_id              = var.vpc_id
	  cluster_cidr        = "172.16.0.0/16"
	  cluster_name        = "my-independent-cluster"
	  cluster_deploy_type = "INDEPENDENT_CLUSTER"

	  # Existing 3 masters
	  master_config {
	    instance_type     = "S5.LARGE8"
	    availability_zone = "ap-guangzhou-3"
	    subnet_id         = "subnet-xxxxxxxx"
	    system_disk_type  = "CLOUD_SSD"
	    system_disk_size  = 100
	    security_group_ids = ["sg-xxxxxxxx"]
	    password          = "YourPassword123!"
	  }

	  master_config {
	    instance_type     = "S5.LARGE8"
	    availability_zone = "ap-guangzhou-3"
	    subnet_id         = "subnet-xxxxxxxx"
	    system_disk_type  = "CLOUD_SSD"
	    system_disk_size  = 100
	    security_group_ids = ["sg-xxxxxxxx"]
	    password          = "YourPassword123!"
	  }

	  master_config {
	    instance_type     = "S5.LARGE8"
	    availability_zone = "ap-guangzhou-3"
	    subnet_id         = "subnet-xxxxxxxx"
	    system_disk_type  = "CLOUD_SSD"
	    system_disk_size  = 100
	    security_group_ids = ["sg-xxxxxxxx"]
	    password          = "YourPassword123!"
	  }

	  # Scale out: add 2 more master nodes
	  master_config {
	    instance_type     = "S5.LARGE8"
	    availability_zone = "ap-guangzhou-4"
	    subnet_id         = "subnet-yyyyyyyy"
	    system_disk_type  = "CLOUD_SSD"
	    system_disk_size  = 100
	    security_group_ids = ["sg-xxxxxxxx"]
	    password          = "YourPassword123!"
	    node_role         = "MASTER_ETCD"
	  }

	  master_config {
	    instance_type     = "S5.LARGE8"
	    availability_zone = "ap-guangzhou-4"
	    subnet_id         = "subnet-yyyyyyyy"
	    system_disk_type  = "CLOUD_SSD"
	    system_disk_size  = 100
	    security_group_ids = ["sg-xxxxxxxx"]
	    password          = "YourPassword123!"
	    node_role         = "MASTER_ETCD"
	  }
	}

```
*/
package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster", CNDescription{
		TerraformTypeCN: "Tke集群",
		DescriptionCN:   "提供Kubernetes集群资源，用于创建和管理TKE集群。",
		AttributesCN: map[string]string{
			"cluster_id":           "集群ID",
			"cluster_arch":         "集群CPU架构",
			"cluster_name":         "集群名称",
			"cluster_desc":         "集群描述",
			"cluster_os":           "集群版本",
			"cluster_os_type":      "集群系统类型",
			"auto_upgrade_cluster": "是否开启集群自动升级",
			"network_type":         "网络类型",
			//"managed_cluster_internet_security_policies": "集群公网访问白名单",
			"cluster_deploy_type": "集群部署类型",
			//"cluster_intranet_subnet_id":                 "集群内网子网ID",
			"docker_graph_path":               "Docker数据目录",
			"is_non_static_ip_mode":           "是否开启非静态IP模式",
			"cluster_internal_service_group":  "集群内网服务组",
			"tags":                            "标签",
			"label":                           "标签",
			"upgrade_instance_follow_cluster": "是否开启节点自动升级",
			"eni_subnet_ids":                  "ENI子网ID",
			"vpc_cni_type":                    "VPC-CNI类型",
			"mount_target":                    "挂载目标",
			"container_runtime":               "容器运行时",
			"runtime_version":                 "运行时版本",
			"enable_customized_pod_cidr":      "是否开启自定义Pod CIDR",
			"deletion_protection":             "是否开启删除保护",
			"kube_proxy_mode":                 "kube-proxy模式",
			"audit_enabled":                   "是否开启审计",
			"audit_logset_id":                 "审计日志集ID",
			"audit_log_topic_id":              "审计日志主题ID",
			"data_plane_v2":                   "是否使用数据平面V2",
			"qgpu_share_enable":               "是否开启QGPU共享",
			"is_dual_stack":                   "是否双栈集群",
			"extension_addon":                 "扩展组件",
			"disable_addons":                  "禁用的组件列表",
			//"cluster_intranet_domain":                    "集群内网域名",
			"node_pool_id": "节点池ID",
			//"acquire_cluster_admin_role": "是否获取集群管理员角色",
			"claim_expired_seconds": "认领过期时间",
			//"cluster_intranet":      "集群内网",
			"cluster_cidr":                 "集群CIDR",
			"ignore_cluster_cidr_conflict": "是否忽略CIDR冲突",
			"ignore_service_cidr_conflict": "是否忽略ServiceCIDR冲突",
			"globe_desired_pod_num":        "全局期望Pod数",
			"base_pod_num":                 "基础Pod数",
			"project_id":                   "项目ID",
			"need_work_security_group":     "是否开启默认节点安全组",
			"unschedulable":                "是否不可调度",
			"cluster_max_service_num":      "集群最大服务数",
			"cluster_ipvs":                 "集群IPVS",
			//"cluster_internet":             "集群公网",
			"service_cidr": "服务CIDR",
			//"cluster_as_enabled":              "是否开启集群自动伸缩",
			"vpc_id": "VPC ID",
			//"cluster_internet_domain":    "集群公网域名",
			"extra_args":             "额外参数",
			"cluster_level":          "集群等级",
			"instance_data_disks":    "节点数据盘配置",
			"run_instances_for_node": "RunInstancesForNode配置",
			"node_name_type":         "节点名称类型",
			"cluster_max_pod_num":    "集群最大Pod数",
			"log_agent":              "日志代理",
			"master_config":          "主节点配置",
			//"node_pool_global_config": "节点池全局配置",
			"worker_config":     "工作节点配置",
			"event_persistence": "事件持久化",
			//"exist_instance":    "已有实例",
			"cluster_extra_args": "集群额外参数",
			//"auth_options":            "认证选项",
			"cluster_audit":              "集群审计",
			"cluster_node_num":           "集群节点数",
			"user_name":                  "用户名",
			"password":                   "密码",
			"certification_authority":    "证书",
			"cluseter_external_endpoint": "集群外网地址",
			"security_policy":            "安全策略",
			"kube_config":                "kube配置",
			"pgw_endpoint":               "内网地址",
			"domain":                     "域名",
			"worker_instances_list":      "工作节点实例列表",
			//"auto_upgrade_cluster_level":      "自动升级集群等级",
			//"cluster_internet_security_group": "集群公网安全组",
			//"labels":                          "标签",
			"cluster_version": "集群版本",
			//"upgrade_instances_follow_cluster":        "升级实例跟随集群",
			"cluster_external_endpoint":           "集群外网地址",
			"kube_config_intranet":                "内网kube配置",
			"enabled":                             "是否启用",
			"kubelet_root_dir":                    "kubelet数据目录",
			"subnet_id":                           "子网ID",
			"system_disk_type":                    "系统盘类型",
			"system_disk_pool_group":              "System存储资源池分组",
			"internet_charge_type":                "公网计费类型",
			"internet_max_bandwidth_out":          "公网出带宽",
			"hostname":                            "主机名",
			"instance_charge_type_prepaid_period": "预付费时长",
			"disaster_recover_group_ids":          "容灾组ID",
			"cam_role_name":                       "CAM角色名",
			"img_id":                              "镜像ID",
			"bandwidth_package_id":                "带宽包ID",
			"enhanced_security_service":           "增强安全服务",
			"enhanced_monitor_service":            "增强监控服务",
			"data_disk":                           "数据盘",
			"disk_type":                           "磁盘类型",
			"disk_pool_group":                     "存储资源池分组",
			"disk_size":                           "磁盘大小",
			"snapshot_id":                         "快照ID",
			"encrypt":                             "是否加密",
			"kms_key_id":                          "KMS密钥ID",
			"auto_format_and_mount":               "自动格式化和挂载",
			"disk_partition":                      "磁盘分区",
			"file_system":                         "文件系统",
			"public_ip_assigned":                  "是否分配公网IP",
			//"desired_pod_num":                     "期望Pod数",
			"instance_type": "实例类型",
			"instance_charge_type_prepaid_renew_flag": "预付费续费标识",
			"system_disk_size":                        "系统盘大小",
			"availability_zone":                       "可用区",
			"hpc_cluster_id":                          "HPC集群ID",
			"key_ids":                                 "密钥ID",
			"count":                                   "数量",
			"instance_charge_type":                    "实例计费类型",
			"instance_name":                           "实例名称",
			"security_group_ids":                      "安全组ID",
			"user_data":                               "用户数据",
			"pre_start_user_script":                   "节点初始化前脚本",
			"user_script":                             "节点初始化后脚本",
			"kube_apiserver":                          "K8S API服务器",
			"kube_controller_manager":                 "K8S控制器管理器",
			"kube_scheduler":                          "K8S调度器",
			"scale_in_unneeded_time":                  "缩容不需要时间",
			"ignore_daemon_sets_utilization":          "忽略守护进程集利用率",
			"skip_nodes_with_local_storage":           "跳过具有本地存储的节点",
			"skip_nodes_with_system_pods":             "跳过具有系统Pods的节点",
			"is_scale_in_enabled":                     "是否启用缩容",
			"expander":                                "扩展器",
			"scale_in_utilization_threshold":          "缩容利用率阈值",
			"max_concurrent_scale_in":                 "最大并发缩容",
			"scale_in_delay":                          "缩容延迟",
			"issuer":                                  "发行者",
			"auto_create_discovery_anonymous_auth":    "自动创建发现匿名认证",
			"use_tke_default":                         "使用TKE默认",
			"jwks_uri":                                "指定 service-account-jwks-uri。 如果 use_tke_default 设置为“true”，请不要设置此字段，无论如何它都会被忽略。",
			"desired_pod_numbers":                     "期望Pod数",
			"node_role":                               "节点角色",
			"instances_para":                          "实例参数",
			"instance_ids":                            "实例ID",
			"log_set_id":                              "日志集ID",
			"topic_id":                                "主题ID",
			"delete_audit_log_and_topic":              "删除审计日志和主题",
			"delete_event_log_and_topic":              "删除事件日志和主题",
			"name":                                    "名称",
			"param":                                   "参数",
			"instance_id":                             "实例ID",
			"instance_role":                           "实例角色",
			"instance_state":                          "实例状态",
			"failed_reason":                           "失败原因",
			"lan_ip":                                  "LAN IP",
		},
	})
}

func expandInstanceDataDisks(raw []interface{}) ([]*tke.DataDisk, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	dataDisks := make([]*tke.DataDisk, 0, len(raw))
	for _, item := range raw {
		diskMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("instance_data_disks should be a map")
		}
		dataDisk := &tke.DataDisk{}
		if v, ok := diskMap["disk_type"].(string); ok && v != "" {
			dataDisk.DiskType = helper.String(v)
		}
		if v, ok := diskMap["disk_size"]; ok {
			dataDisk.DiskSize = helper.IntInt64(v.(int))
		}
		if v, ok := diskMap["disk_id"].(string); ok && v != "" {
			dataDisk.DiskId = helper.String(v)
		}
		if v, ok := diskMap["disk_partition"].(string); ok && v != "" {
			dataDisk.DiskPartition = helper.String(v)
		}
		if v, ok := diskMap["file_system"].(string); ok && v != "" {
			dataDisk.FileSystem = helper.String(v)
		}
		if v, ok := diskMap["auto_format_and_mount"]; ok {
			dataDisk.AutoFormatAndMount = helper.Bool(v.(bool))
		}
		if v, ok := diskMap["mount_target"].(string); ok && v != "" {
			dataDisk.MountTarget = helper.String(v)
		}
		dataDisks = append(dataDisks, dataDisk)
	}
	return dataDisks, nil
}

func expandInstanceAdvancedSettingsOverride(raw map[string]interface{}) (*tke.InstanceAdvancedSettings, error) {
	override := &tke.InstanceAdvancedSettings{}
	if v, ok := raw["docker_graph_path"].(string); ok && v != "" {
		override.DockerGraphPath = helper.String(v)
	}
	if v, ok := raw["unschedulable"]; ok {
		override.Unschedulable = helper.IntInt64(v.(int))
	}
	if v, ok := raw["pre_start_user_script"].(string); ok && v != "" {
		override.PreStartUserScript = helper.String(v)
	}
	if v, ok := raw["user_script"].(string); ok && v != "" {
		override.UserScript = helper.String(v)
	}
	if v, ok := raw["labels"]; ok {
		labelList := v.([]interface{})
		labels := make([]*tke.Label, 0, len(labelList))
		for _, item := range labelList {
			labelMap, ok := item.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("instance_advanced_settings_overrides.labels should be a map")
			}
			name, _ := labelMap["name"].(string)
			value, _ := labelMap["value"].(string)
			labels = append(labels, &tke.Label{Name: helper.String(name), Value: helper.String(value)})
		}
		if len(labels) > 0 {
			override.Labels = labels
		}
	}
	if v, ok := raw["taints"]; ok {
		taintList := v.([]interface{})
		taints := make([]*tke.Taint, 0, len(taintList))
		for _, item := range taintList {
			taintMap, ok := item.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("instance_advanced_settings_overrides.taints should be a map")
			}
			key, _ := taintMap["key"].(string)
			value, _ := taintMap["value"].(string)
			effect, _ := taintMap["effect"].(string)
			taints = append(taints, &tke.Taint{Key: helper.String(key), Value: helper.String(value), Effect: helper.String(effect)})
		}
		if len(taints) > 0 {
			override.Taints = taints
		}
	}
	if v, ok := raw["extra_args"]; ok {
		argsList := v.([]interface{})
		if len(argsList) > 0 {
			extraArgs := tke.InstanceExtraArgs{}
			for _, arg := range argsList {
				if s, ok := arg.(string); ok && s != "" {
					extraArgs.Kubelet = append(extraArgs.Kubelet, helper.String(s))
				}
			}
			override.ExtraArgs = &extraArgs
		}
	}
	if v, ok := raw["data_disks"]; ok {
		dataDisks, err := expandInstanceDataDisks(v.([]interface{}))
		if err != nil {
			return nil, err
		}
		if len(dataDisks) > 0 {
			override.DataDisks = dataDisks
		}
	}
	if v, ok := raw["desired_pod_number"]; ok {
		override.DesiredPodNumber = helper.IntInt64(v.(int))
	}
	if v, ok := raw["gpu_args"]; ok && len(v.([]interface{})) > 0 {
		rawElem := v.([]interface{})[0]
		if rawElem != nil {
			if gpuArgs, ok := rawElem.(map[string]interface{}); ok {
				tkeGpuArgs := &tke.GPUArgs{}
				if migEnable, ok := gpuArgs["mig_enable"]; ok {
					migEnabled := migEnable.(bool)
					tkeGpuArgs.MIGEnable = &migEnabled
				}
				if raw, ok := gpuArgs["driver"]; ok && raw != nil {
					if driver, ok := raw.(map[string]interface{}); ok && len(driver) > 0 {
						version, hasVersion := driver["version"].(string)
						name, hasName := driver["name"].(string)
						if hasVersion || hasName {
							tkeGpuArgs.Driver = &tke.DriverVersion{
								Version: helper.String(version),
								Name:    helper.String(name),
							}
						}
					}
				}
				if raw, ok := gpuArgs["cuda"]; ok && raw != nil {
					if cuda, ok := raw.(map[string]interface{}); ok && len(cuda) > 0 {
						version, hasVersion := cuda["version"].(string)
						name, hasName := cuda["name"].(string)
						if hasVersion || hasName {
							tkeGpuArgs.CUDA = &tke.DriverVersion{
								Version: helper.String(version),
								Name:    helper.String(name),
							}
						}
					}
				}
				if raw, ok := gpuArgs["cudnn"]; ok && raw != nil {
					if cudnn, ok := raw.(map[string]interface{}); ok && len(cudnn) > 0 {
						version, hasVersion := cudnn["version"].(string)
						name, hasName := cudnn["name"].(string)
						if hasVersion || hasName {
							tkeGpuArgs.CUDNN = &tke.CUDNN{
								Version: helper.String(version),
								Name:    helper.String(name),
							}
						}
						if docName, ok := cudnn["doc_name"].(string); ok && tkeGpuArgs.CUDNN != nil {
							tkeGpuArgs.CUDNN.DocName = helper.String(docName)
						}
						if devName, ok := cudnn["dev_name"].(string); ok && tkeGpuArgs.CUDNN != nil {
							tkeGpuArgs.CUDNN.DevName = helper.String(devName)
						}
					}
				}
				if raw, ok := gpuArgs["custom_driver"]; ok && raw != nil {
					if customDriver, ok := raw.(map[string]interface{}); ok && len(customDriver) > 0 {
						if address, ok := customDriver["address"].(string); ok {
							tkeGpuArgs.CustomDriver = &tke.CustomDriver{
								Address: helper.String(address),
							}
						}
					}
				}
				override.GPUArgs = tkeGpuArgs
			}
		}
	}
	return override, nil
}

func expandRunInstancesForNode(raw []interface{}) ([]*tke.RunInstancesForNode, error) {
	if len(raw) == 0 {
		return nil, fmt.Errorf("run_instances_for_node must not be empty")
	}
	nodes := make([]*tke.RunInstancesForNode, 0, len(raw))
	for _, item := range raw {
		nodeMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("run_instances_for_node should be a map")
		}
		node := &tke.RunInstancesForNode{}
		if v, ok := nodeMap["node_role"].(string); ok {
			node.NodeRole = helper.String(v)
		}
		if v, ok := nodeMap["run_instances_para"]; ok {
			paraList := v.([]interface{})
			node.RunInstancesPara = make([]*string, 0, len(paraList))
			for _, para := range paraList {
				if s, ok := para.(string); ok {
					node.RunInstancesPara = append(node.RunInstancesPara, helper.String(s))
				}
			}
		}
		if len(node.RunInstancesPara) == 0 {
			return nil, fmt.Errorf("run_instances_para must be set when using run_instances_for_node")
		}
		if v, ok := nodeMap["instance_advanced_settings_overrides"]; ok {
			overrideList := v.([]interface{})
			for _, overrideItem := range overrideList {
				overrideMap, ok := overrideItem.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("instance_advanced_settings_overrides should be a map")
				}
				override, err := expandInstanceAdvancedSettingsOverride(overrideMap)
				if err != nil {
					return nil, err
				}
				node.InstanceAdvancedSettingsOverrides = append(node.InstanceAdvancedSettingsOverrides, override)
			}
			if len(node.InstanceAdvancedSettingsOverrides) > 0 && len(node.InstanceAdvancedSettingsOverrides) != len(node.RunInstancesPara) {
				return nil, fmt.Errorf("instance_advanced_settings_overrides length must match run_instances_para length")
			}
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func tkeCvmState() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Tke.V20170312.CreateCluster.InstanceId
		"instance_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the cvm.",
		},
		"instance_role": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Role of the cvm.",
		},
		"instance_state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "State of the cvm.",
		},
		"failed_reason": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Information of the cvm when it is failed.",
		},
		"lan_ip": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "LAN IP of the cvm.",
		},
	}
}

func tkeSecurityInfo() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "User name of account.",
		},
		"password": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Password of account.",
		},
		"certification_authority": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The certificate used for access.",
		},
		"cluster_external_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "External network address to access.",
		},
		"domain": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Domain name for access.",
		},
		"pgw_endpoint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The Intranet address used for access.",
		},
		"security_policy": {
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Access policy.",
		},
	}
}

func TkeMasterCvmCreateInfo() map[string]*schema.Schema {
	res := TkeCvmCreateInfo()
	masterRes := make(map[string]*schema.Schema)
	for k, v := range res {
		if k == "count" {
			// master_config block always represents 1 master node; count is handled by block count
			continue
		}
		newSchema := *v
		newSchema.ForceNew = false
		masterRes[k] = &newSchema
	}
	masterRes["instance_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The ID of the master CVM instance.",
	}
	masterRes["node_role"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Default:      "MASTER_ETCD",
		ValidateFunc: validateAllowedStringValue([]string{"MASTER_ETCD", "MASTER", "ETCD"}),
		Description:  "The role of the node. Valid values: `MASTER_ETCD` (default), `MASTER`, `ETCD`.",
	}
	masterRes["pre_start_user_script"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    false,
		Description: "Base64-encoded user script executed before TKE initializes the node. All master configurations must use the same value during cluster creation; newly added masters may use different values during scale-out.",
	}
	masterRes["user_script"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    false,
		Description: "Base64-encoded user script executed after TKE initializes the node. All master configurations must use the same value during cluster creation; newly added masters may use different values during scale-out.",
	}
	return masterRes
}

// normalizeMasterConfigBlock fills in default zero values for fields that Read
// cannot recover from the CVM API (password) or that may end up as nil in state
// (user_data, security_group_ids). Without this, Terraform sees state as null
// but the schema default as "" during plan, producing a plan drift on every
// invocation and blocking scale-in/scale-out through masterConfigValueEqual.
// TKE user scripts are also not returned by the CVM API, so imported clusters
// must get the same empty-string state value.
func normalizeMasterConfigBlock(m map[string]interface{}) {
	if m == nil {
		return
	}
	if _, ok := m["password"]; !ok || m["password"] == nil {
		m["password"] = ""
	}
	if _, ok := m["user_data"]; !ok || m["user_data"] == nil {
		m["user_data"] = ""
	}
	if v, ok := m["security_group_ids"]; !ok || v == nil {
		m["security_group_ids"] = []interface{}{}
	}
	if _, ok := m["pre_start_user_script"]; !ok || m["pre_start_user_script"] == nil {
		m["pre_start_user_script"] = ""
	}
	if _, ok := m["user_script"]; !ok || m["user_script"] == nil {
		m["user_script"] = ""
	}
}

func bindMasterConfigInstance(existing map[string]interface{}, instance *cvm.Instance) map[string]interface{} {
	bound := make(map[string]interface{}, len(existing)+1)
	for key, value := range existing {
		bound[key] = value
	}
	if instance != nil && instance.InstanceId != nil {
		bound["instance_id"] = helper.PString(instance.InstanceId)
	}
	normalizeMasterConfigBlock(bound)
	return bound
}

func newMasterConfigFromCvm(instance *cvm.Instance, masterInfo InstanceInfo) map[string]interface{} {
	mapping := map[string]interface{}{
		"instance_charge_type_prepaid_period":     1,
		"instance_charge_type_prepaid_renew_flag": CVM_PREPAID_RENEW_FLAG_NOTIFY_AND_MANUAL_RENEW,
		"system_disk_pool_group":                  "",
		"cam_role_name":                           "",
		"desired_pod_num":                         DefaultDesiredPodNum,
		"node_role":                               "MASTER_ETCD",
		"enhanced_security_service":               true,
		"enhanced_monitor_service":                true,
		"enhanced_automation_service":             true,
		"public_ip_assigned":                      false,
		"internet_charge_type":                    INTERNET_CHARGE_TYPE_TRAFFIC_POSTPAID_BY_HOUR,
	}
	if instance == nil {
		normalizeMasterConfigBlock(mapping)
		return mapping
	}
	if instance.InstanceId != nil {
		mapping["instance_id"] = helper.PString(instance.InstanceId)
	}
	if instance.InstanceName != nil {
		mapping["instance_name"] = helper.PString(instance.InstanceName)
	}
	if instance.InstanceType != nil {
		mapping["instance_type"] = helper.PString(instance.InstanceType)
	}
	if instance.InstanceChargeType != nil {
		mapping["instance_charge_type"] = helper.PString(instance.InstanceChargeType)
	}
	if instance.ImageId != nil {
		mapping["img_id"] = helper.PString(instance.ImageId)
	}
	if instance.VirtualPrivateCloud != nil && instance.VirtualPrivateCloud.SubnetId != nil {
		mapping["subnet_id"] = helper.PString(instance.VirtualPrivateCloud.SubnetId)
	}
	if instance.Placement != nil && instance.Placement.Zone != nil {
		mapping["availability_zone"] = helper.PString(instance.Placement.Zone)
	}
	if instance.SystemDisk != nil {
		if instance.SystemDisk.DiskType != nil {
			mapping["system_disk_type"] = helper.PString(instance.SystemDisk.DiskType)
		}
		if instance.SystemDisk.DiskSize != nil {
			mapping["system_disk_size"] = helper.PInt64(instance.SystemDisk.DiskSize)
		}
	}
	if instance.InternetAccessible != nil {
		if instance.InternetAccessible.InternetChargeType != nil {
			mapping["internet_charge_type"] = helper.PString(instance.InternetAccessible.InternetChargeType)
		}
		if instance.InternetAccessible.InternetMaxBandwidthOut != nil {
			bandwidth := helper.PInt64(instance.InternetAccessible.InternetMaxBandwidthOut)
			mapping["internet_max_bandwidth_out"] = bandwidth
			mapping["public_ip_assigned"] = bandwidth > 0
		}
	}
	if instance.SecurityGroupIds != nil {
		mapping["security_group_ids"] = helper.StringsInterfaces(instance.SecurityGroupIds)
	}
	if instance.RenewFlag != nil && helper.PString(instance.InstanceChargeType) == "PREPAID" {
		mapping["instance_charge_type_prepaid_renew_flag"] = helper.PString(instance.RenewFlag)
	}
	if instance.CamRoleName != nil {
		mapping["cam_role_name"] = helper.PString(instance.CamRoleName)
	}
	if instance.LoginSettings != nil {
		mapping["key_ids"] = helper.StringsInterfaces(instance.LoginSettings.KeyIds)
	}
	if instance.DataDisks != nil {
		dataDisks := make([]interface{}, 0, len(instance.DataDisks))
		for _, disk := range instance.DataDisks {
			if disk == nil {
				continue
			}
			diskConfig := make(map[string]interface{})
			if disk.DiskType != nil {
				diskConfig["disk_type"] = helper.PString(disk.DiskType)
			}
			if disk.DiskSize != nil {
				diskConfig["disk_size"] = helper.PInt64(disk.DiskSize)
			}
			dataDisks = append(dataDisks, diskConfig)
		}
		mapping["data_disk"] = dataDisks
	}
	if masterInfo.InstanceRole != "" {
		mapping["node_role"] = masterInfo.InstanceRole
	}
	if masterInfo.InstanceAdvancedSettings != nil {
		if masterInfo.InstanceAdvancedSettings.DesiredPodNumber != nil {
			mapping["desired_pod_num"] = helper.PInt64(masterInfo.InstanceAdvancedSettings.DesiredPodNumber)
		}
		if masterInfo.InstanceAdvancedSettings.PreStartUserScript != nil {
			mapping["pre_start_user_script"] = helper.PString(masterInfo.InstanceAdvancedSettings.PreStartUserScript)
		}
		if masterInfo.InstanceAdvancedSettings.UserScript != nil {
			mapping["user_script"] = helper.PString(masterInfo.InstanceAdvancedSettings.UserScript)
		}
	}
	normalizeMasterConfigBlock(mapping)
	return mapping
}

func masterConfigMatchesCvmSpec(config map[string]interface{}, instance *cvm.Instance) bool {
	if instance == nil {
		return false
	}
	matchedField := false
	if value, _ := config["instance_type"].(string); value != "" {
		matchedField = true
		if value != helper.PString(instance.InstanceType) {
			return false
		}
	}
	if value, _ := config["subnet_id"].(string); value != "" {
		matchedField = true
		if instance.VirtualPrivateCloud == nil || value != helper.PString(instance.VirtualPrivateCloud.SubnetId) {
			return false
		}
	}
	if value, _ := config["availability_zone"].(string); value != "" {
		matchedField = true
		if instance.Placement == nil || value != helper.PString(instance.Placement.Zone) {
			return false
		}
	}
	return matchedField
}

func matchMasterConfigToCvm(config map[string]interface{}, instances []*cvm.Instance, matched map[string]bool) *cvm.Instance {
	instanceName, _ := config["instance_name"].(string)
	var idCandidate *cvm.Instance
	if instanceId, _ := config["instance_id"].(string); instanceId != "" {
		for _, instance := range instances {
			if instance != nil && instance.InstanceId != nil && helper.PString(instance.InstanceId) == instanceId && !matched[instanceId] {
				// A TypeList plan may carry the computed instance_id from the old
				// positional element after HCL reorders the blocks. Trust the ID only
				// when it still agrees with the configured instance_name; otherwise
				// keep it as a fallback and rebind by name below.
				if instanceName == "" || helper.PString(instance.InstanceName) == instanceName {
					return instance
				}
				idCandidate = instance
				break
			}
		}
	}

	var nameCandidate *cvm.Instance
	for _, instance := range instances {
		if instance == nil || instance.InstanceId == nil || matched[helper.PString(instance.InstanceId)] {
			continue
		}
		if instanceName != "" && helper.PString(instance.InstanceName) == instanceName {
			if masterConfigMatchesCvmSpec(config, instance) {
				return instance
			}
			if nameCandidate == nil {
				nameCandidate = instance
			}
		}
	}
	if nameCandidate != nil {
		return nameCandidate
	}
	if idCandidate != nil {
		return idCandidate
	}
	if instanceName != "" {
		return nil
	}

	for _, instance := range instances {
		if instance == nil || instance.InstanceId == nil || matched[helper.PString(instance.InstanceId)] {
			continue
		}
		if masterConfigMatchesCvmSpec(config, instance) {
			return instance
		}
	}
	return nil
}

func reconcileMasterConfigState(existing []interface{}, instances []*cvm.Instance, masters []InstanceInfo) []interface{} {
	orderedInstances := make([]*cvm.Instance, 0, len(instances))
	for _, instance := range instances {
		if instance != nil && instance.InstanceId != nil {
			orderedInstances = append(orderedInstances, instance)
		}
	}
	sort.SliceStable(orderedInstances, func(i, j int) bool {
		leftName := helper.PString(orderedInstances[i].InstanceName)
		rightName := helper.PString(orderedInstances[j].InstanceName)
		if leftName == rightName {
			return helper.PString(orderedInstances[i].InstanceId) < helper.PString(orderedInstances[j].InstanceId)
		}
		return leftName < rightName
	})

	masterById := make(map[string]InstanceInfo, len(masters))
	for _, master := range masters {
		masterById[master.InstanceId] = master
	}
	matched := make(map[string]bool, len(orderedInstances))
	reconciled := make([]interface{}, 0, len(orderedInstances))

	// Preserve the configuration order and all create-only fields for masters
	// that still exist. This is the normal create, scale-out, and scale-in path.
	for _, raw := range existing {
		config, ok := raw.(map[string]interface{})
		if !ok || config == nil {
			continue
		}
		instance := matchMasterConfigToCvm(config, orderedInstances, matched)
		if instance == nil || instance.InstanceId == nil {
			continue
		}
		instanceId := helper.PString(instance.InstanceId)
		matched[instanceId] = true
		reconciled = append(reconciled, bindMasterConfigInstance(config, instance))
	}

	// Unmatched CVMs are masters created outside the current state (including
	// import and externally added nodes). Append them in deterministic order.
	for _, instance := range orderedInstances {
		if instance == nil || instance.InstanceId == nil {
			continue
		}
		instanceId := helper.PString(instance.InstanceId)
		if matched[instanceId] {
			continue
		}
		reconciled = append(reconciled, newMasterConfigFromCvm(instance, masterById[instanceId]))
	}
	return reconciled
}

func validateMasterConfigStateSchema(configs []interface{}) error {
	masterSchema := TkeMasterCvmCreateInfo()
	for index, raw := range configs {
		config, ok := raw.(map[string]interface{})
		if !ok || config == nil {
			continue
		}
		for key := range config {
			if _, exists := masterSchema[key]; !exists {
				return fmt.Errorf("master_config.%d contains unsupported state field %q", index, key)
			}
		}
	}
	return nil
}

func masterInstancesNotReadyError(masters []InstanceInfo) error {
	notReady := make([]string, 0)
	for _, master := range masters {
		if strings.EqualFold(master.InstanceState, "running") {
			continue
		}
		reason := strings.TrimSpace(master.FailedReason)
		if reason == "" {
			reason = "reason not reported by TKE"
		}
		notReady = append(notReady, fmt.Sprintf("%s (state: %s, reason: %s)", master.InstanceId, master.InstanceState, reason))
	}
	if len(notReady) == 0 {
		return nil
	}
	sort.Strings(notReady)
	return fmt.Errorf("TKE reported master node(s) that are not ready: %s. Terraform will not submit another master scaling request or automatically delete/recreate control-plane nodes; wait for the TKE operation to finish and all masters to become healthy, then run Terraform again. If the condition persists, investigate the reported TKE reason", strings.Join(notReady, ", "))
}

type masterInventoryDirection int

const (
	masterInventoryGrowing   masterInventoryDirection = 1
	masterInventoryShrinking masterInventoryDirection = -1
)

func masterConfigInstanceIDs(configs []interface{}) (map[string]bool, error) {
	ids := make(map[string]bool, len(configs))
	for _, raw := range configs {
		config, ok := raw.(map[string]interface{})
		if !ok || config == nil {
			return nil, fmt.Errorf("master_config contains an invalid state block")
		}
		instanceId, _ := config["instance_id"].(string)
		if instanceId == "" {
			name, _ := config["instance_name"].(string)
			return nil, fmt.Errorf("master_config block %q has no instance_id; refresh state before scaling", name)
		}
		if ids[instanceId] {
			return nil, fmt.Errorf("master_config contains duplicate instance_id %q; refresh and repair state before scaling", instanceId)
		}
		ids[instanceId] = true
	}
	return ids, nil
}

func masterInventoryConvergenceError(masters []InstanceInfo, expectedCount int, requiredInstanceIds map[string]bool, direction masterInventoryDirection, operation string) *resource.RetryError {
	actualCount := len(masters)
	if actualCount != expectedCount {
		err := fmt.Errorf("%s is still converging: TKE returned %d master node(s), want %d", operation, actualCount, expectedCount)
		if (direction == masterInventoryGrowing && actualCount < expectedCount) || (direction == masterInventoryShrinking && actualCount > expectedCount) {
			return resource.RetryableError(err)
		}
		return resource.NonRetryableError(fmt.Errorf("%s; the master inventory moved in an unexpected direction, so Terraform stopped without submitting another scaling request or deleting nodes automatically", err.Error()))
	}

	liveIds := make(map[string]bool, actualCount)
	for _, master := range masters {
		liveIds[master.InstanceId] = true
	}
	for instanceId := range requiredInstanceIds {
		if !liveIds[instanceId] {
			return resource.NonRetryableError(fmt.Errorf("%s reached the target count, but existing master %s is missing; Terraform stopped without deleting or recreating any master", operation, instanceId))
		}
	}

	if err := masterInstancesNotReadyError(masters); err != nil {
		return resource.RetryableError(fmt.Errorf("%s is still converging: %w", operation, err))
	}
	return nil
}

func validateMasterScalingPreflight(masters []InstanceInfo, sourceCount, targetCount int, requiredInstanceIds map[string]bool, clusterState, operation string) (bool, error) {
	if !strings.EqualFold(clusterState, "Running") {
		return false, fmt.Errorf("cluster is currently in %q before %s; Terraform did not submit a new master scaling request. Wait for TKE to finish the existing operation and run Terraform again", clusterState, operation)
	}

	liveIds := make(map[string]bool, len(masters))
	for _, master := range masters {
		liveIds[master.InstanceId] = true
	}
	for instanceId := range requiredInstanceIds {
		if !liveIds[instanceId] {
			return false, fmt.Errorf("master inventory changed after planning %s: required master %s is missing. Terraform did not submit a scaling request; refresh and review the plan", operation, instanceId)
		}
	}

	if len(masters) == targetCount {
		return true, nil
	}
	if len(masters) != sourceCount {
		return false, fmt.Errorf("master inventory changed after planning %s: TKE now has %d node(s), while the plan was based on %d and targets %d. Terraform did not submit a scaling request; refresh and review the plan", operation, len(masters), sourceCount, targetCount)
	}
	if err := masterInstancesNotReadyError(masters); err != nil {
		return false, fmt.Errorf("cannot start %s while the current master inventory is not healthy: %w", operation, err)
	}
	return false, nil
}

func readMasterScalingPreflight(ctx context.Context, service TkeService, clusterId string, sourceCount, targetCount int, requiredInstanceIds map[string]bool, operation string) (bool, error) {
	var masters []InstanceInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		var inErr error
		masters, _, inErr = service.DescribeClusterInstances(ctx, clusterId)
		if inErr != nil {
			return retryError(inErr)
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	var status *tke.ClusterStatus
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		var inErr error
		status, inErr = service.DescribeClusterStatus(ctx, clusterId)
		if inErr != nil {
			return retryError(inErr)
		}
		if status == nil || status.ClusterState == nil {
			return resource.RetryableError(fmt.Errorf("cluster %s has not reported a state before %s", clusterId, operation))
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return validateMasterScalingPreflight(masters, sourceCount, targetCount, requiredInstanceIds, helper.PString(status.ClusterState), operation)
}

// expandMasterConfigInstanceAdvancedSettings expands the TKE per-node settings
// supported by master_config. The returned bool is true only when a non-empty
// override must be sent to TKE. A zero-value settings object is still useful to
// preserve the position of other master/worker entries once any override exists.
func expandMasterConfigInstanceAdvancedSettings(raw map[string]interface{}) (tke.InstanceAdvancedSettings, bool, error) {
	var override tke.InstanceAdvancedSettings
	if raw == nil {
		return override, false, nil
	}

	if value, exists := raw["desired_pod_num"]; exists && value != nil {
		desiredPodNum, ok := value.(int)
		if !ok {
			return override, false, fmt.Errorf("master_config.desired_pod_num must be an integer")
		}
		if int64(desiredPodNum) != DefaultDesiredPodNum {
			override.DesiredPodNumber = helper.Int64(int64(desiredPodNum))
		}
	}

	if value, exists := raw["pre_start_user_script"]; exists && value != nil {
		script, ok := value.(string)
		if !ok {
			return override, false, fmt.Errorf("master_config.pre_start_user_script must be a string")
		}
		if script != "" {
			override.PreStartUserScript = helper.String(script)
		}
	}

	if value, exists := raw["user_script"]; exists && value != nil {
		script, ok := value.(string)
		if !ok {
			return override, false, fmt.Errorf("master_config.user_script must be a string")
		}
		if script != "" {
			override.UserScript = helper.String(script)
		}
	}

	return override, override.DesiredPodNumber != nil || override.PreStartUserScript != nil || override.UserScript != nil, nil
}

func expandDesiredPodNumberOverride(raw map[string]interface{}) (tke.InstanceAdvancedSettings, bool, error) {
	var override tke.InstanceAdvancedSettings
	if raw == nil {
		return override, false, nil
	}
	value, exists := raw["desired_pod_num"]
	if !exists || value == nil {
		return override, false, nil
	}
	desiredPodNum, ok := value.(int)
	if !ok {
		return override, false, fmt.Errorf("desired_pod_num must be an integer")
	}
	if int64(desiredPodNum) == DefaultDesiredPodNum {
		return override, false, nil
	}
	override.DesiredPodNumber = helper.Int64(int64(desiredPodNum))
	return override, true, nil
}

// masterConfigFieldsEquivalent reports whether two field values should be
// treated as equal for the purpose of detecting user-authored modifications.
// It returns true only when both values are "empty" in one of the shapes that
// Terraform routinely conflates (nil, "", 0, false, empty list/map). This lets
// scale-in/scale-out proceed when Read has populated state with a zero value
// but the user's HCL never set the field at all (or vice versa).
func masterConfigFieldsEquivalent(a, b interface{}) bool {
	return isMasterConfigZero(a) && isMasterConfigZero(b)
}

// masterConfigFieldDoesNotRequireMutation reports whether a field difference
// can be reconciled in state without changing an existing master.
//
// hostname is write-only in the current CVM/TKE read APIs. Once another stable
// identity (instance_name and the old state's instance_id) has matched the
// block, an omitted hostname means "unknown", not a request to rename the
// machine. A change between two non-empty hostnames remains unsupported.
//
// desired_pod_num uses DefaultDesiredPodNum as "no per-instance override".
// Older Read implementations persisted TKE's resolved value (for example 128)
// instead, so a later scale-in could incorrectly see 128 -> 0 as a mutation of
// every retained master. The default value must inherit the existing remote
// setting during comparison.
//
// user_data is also write-only. Preserve the existing recovery behavior that
// allows an empty state value to be hydrated from configuration after a remote
// scale operation completed but Terraform could not persist the planned state.
func masterConfigFieldDoesNotRequireMutation(field string, oldValue, newValue interface{}) bool {
	switch field {
	case "hostname":
		return isMasterConfigZero(oldValue) || isMasterConfigZero(newValue)
	case "desired_pod_num":
		value, ok := newValue.(int)
		return ok && int64(value) == DefaultDesiredPodNum
	case "user_data":
		return isMasterConfigZero(oldValue) && !isMasterConfigZero(newValue)
	default:
		return false
	}
}

func isMasterConfigZero(v interface{}) bool {
	switch t := v.(type) {
	case nil:
		return true
	case string:
		return t == ""
	case int:
		return t == 0
	case bool:
		return !t
	case []interface{}:
		return len(t) == 0
	case map[string]interface{}:
		return len(t) == 0
	}
	return false
}

// masterConfigValueEqual compares two master_config field values for equality.
// For slice types (security_group_ids, data_disk, key_ids, disaster_recover_group_ids),
// order is normalized before comparison to avoid false positives from reflect.DeepEqual.
func masterConfigValueEqual(a, b interface{}) bool {
	// Normalize slices: sort both sides then compare
	switch va := a.(type) {
	case []interface{}:
		vb, ok := b.([]interface{})
		if !ok {
			return false
		}
		if len(va) != len(vb) {
			return false
		}
		// For simple string slices, sort before compare
		strA := make([]string, 0, len(va))
		allStrings := true
		for _, v := range va {
			s, ok := v.(string)
			if !ok {
				allStrings = false
				break
			}
			strA = append(strA, s)
		}
		if allStrings {
			strB := make([]string, 0, len(vb))
			for _, v := range vb {
				s, ok := v.(string)
				if !ok {
					allStrings = false
					break
				}
				strB = append(strB, s)
			}
			if allStrings {
				sort.Strings(strA)
				sort.Strings(strB)
				return reflect.DeepEqual(strA, strB)
			}
		}
		// For complex slices (e.g. data_disk blocks), fall back to DeepEqual
		return reflect.DeepEqual(va, vb)
	case string:
		vb, ok := b.(string)
		return ok && va == vb
	case int:
		vb, ok := b.(int)
		return ok && va == vb
	case bool:
		vb, ok := b.(bool)
		return ok && va == vb
	case map[string]interface{}:
		vb, ok := b.(map[string]interface{})
		return ok && reflect.DeepEqual(va, vb)
	default:
		return reflect.DeepEqual(a, b)
	}
}

// masterConfigMutationField returns the first deterministic field that would
// require mutating an existing master. Computed identity and sensitive values
// are excluded because scale operations resolve the authoritative instance_id
// from the old state and never update a retained master's password.
func masterConfigMutationField(oldConfig, newConfig map[string]interface{}) string {
	skipFields := map[string]bool{
		"instance_id": true,
		"password":    true,
	}
	fields := make([]string, 0, len(newConfig))
	for field := range newConfig {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		if skipFields[field] {
			continue
		}
		newValue := newConfig[field]
		oldValue := oldConfig[field]
		if masterConfigFieldDoesNotRequireMutation(field, oldValue, newValue) {
			continue
		}
		if masterConfigFieldsEquivalent(newValue, oldValue) {
			continue
		}
		if !masterConfigValueEqual(newValue, oldValue) {
			return field
		}
	}
	return ""
}

// masterConfigRetainedMutation returns an unsupported retained-block change
// only when the master inventory itself is unchanged. During a pure scale
// operation, TypeList nested values are positional and cannot be compared
// safely after blocks have been inserted or removed.
func masterConfigRetainedMutation(oldMap, newMap map[string]map[string]interface{}, inventoryChanged bool) (string, string) {
	if inventoryChanged {
		return "", ""
	}
	names := make([]string, 0, len(newMap))
	for name := range newMap {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		oldConfig, exists := oldMap[name]
		if !exists {
			continue
		}
		if field := masterConfigMutationField(oldConfig, newMap[name]); field != "" {
			return name, field
		}
	}
	return "", ""
}

// masterConfigStateForInventoryChange keeps authoritative old-state values for
// retained masters and configuration values for newly added masters. This
// prevents Terraform's positional TypeList merge from persisting shifted
// create-only values after a successful scale operation; Read subsequently
// binds every block to the live CVM instance_id by instance_name.
func masterConfigStateForInventoryChange(newList []interface{}, oldMap map[string]map[string]interface{}) []interface{} {
	state := make([]interface{}, 0, len(newList))
	for _, raw := range newList {
		config, ok := raw.(map[string]interface{})
		if !ok || config == nil {
			continue
		}
		name, _ := config["instance_name"].(string)
		if oldConfig, exists := oldMap[name]; exists {
			state = append(state, oldConfig)
			continue
		}
		state = append(state, config)
	}
	return state
}

type masterConfigInventoryChange struct {
	oldMap        map[string]map[string]interface{}
	newMap        map[string]map[string]interface{}
	addedBlocks   []map[string]interface{}
	removedBlocks []map[string]interface{}
}

// planMasterConfigInventoryChange interprets the complete old/new TypeList
// values produced by Terraform Core. Identity is based on instance_name;
// computed or create-only fields are deliberately not used to decide which
// master is retained or removed.
func planMasterConfigInventoryChange(oldList, newList []interface{}) (*masterConfigInventoryChange, error) {
	if len(newList) < 3 || len(newList) > 7 {
		return nil, fmt.Errorf("TKE independent cluster master_config node count must be between 3 and 7, got %d", len(newList))
	}

	change := &masterConfigInventoryChange{
		oldMap: make(map[string]map[string]interface{}, len(oldList)),
		newMap: make(map[string]map[string]interface{}, len(newList)),
	}
	for _, raw := range oldList {
		config, ok := raw.(map[string]interface{})
		if !ok || config == nil {
			return nil, fmt.Errorf("master_config contains an invalid state block")
		}
		name, _ := config["instance_name"].(string)
		if name == "" {
			return nil, fmt.Errorf("instance_name in master_config must be set and unique")
		}
		if _, exists := change.oldMap[name]; exists {
			return nil, fmt.Errorf("duplicate instance_name %q found in master_config state", name)
		}
		change.oldMap[name] = config
	}
	for _, raw := range newList {
		config, ok := raw.(map[string]interface{})
		if !ok || config == nil {
			return nil, fmt.Errorf("master_config contains an invalid configuration block")
		}
		name, _ := config["instance_name"].(string)
		if name == "" {
			return nil, fmt.Errorf("instance_name in master_config must be set and unique")
		}
		if _, exists := change.newMap[name]; exists {
			return nil, fmt.Errorf("duplicate instance_name %q found in master_config", name)
		}
		change.newMap[name] = config
	}

	newNames := make([]string, 0, len(change.newMap))
	for name := range change.newMap {
		newNames = append(newNames, name)
	}
	sort.Strings(newNames)
	for _, name := range newNames {
		if _, exists := change.oldMap[name]; !exists {
			change.addedBlocks = append(change.addedBlocks, change.newMap[name])
		}
	}
	oldNames := make([]string, 0, len(change.oldMap))
	for name := range change.oldMap {
		oldNames = append(oldNames, name)
	}
	sort.Strings(oldNames)
	for _, name := range oldNames {
		if _, exists := change.newMap[name]; !exists {
			change.removedBlocks = append(change.removedBlocks, change.oldMap[name])
		}
	}

	if len(change.addedBlocks) > 0 && len(change.removedBlocks) > 0 {
		addedNames := make([]string, 0, len(change.addedBlocks))
		for _, config := range change.addedBlocks {
			addedNames = append(addedNames, config["instance_name"].(string))
		}
		removedNames := make([]string, 0, len(change.removedBlocks))
		for _, config := range change.removedBlocks {
			removedNames = append(removedNames, config["instance_name"].(string))
		}
		sort.Strings(removedNames)
		return nil, fmt.Errorf("master_config: cannot add and remove blocks in the same apply (adding %v, removing %v). Split into two applies: scale out first, then scale in (or vice versa)", addedNames, removedNames)
	}

	inventoryChanged := len(change.addedBlocks) > 0 || len(change.removedBlocks) > 0
	if name, field := masterConfigRetainedMutation(change.oldMap, change.newMap, inventoryChanged); field != "" {
		return nil, fmt.Errorf("modifying field %q of existing master_config block %q is not supported. Please delete the block and add a new one instead", field, name)
	}
	return change, nil
}

func buildScaleOutMasterRequest(addedBlocks []map[string]interface{}, meta interface{}, vpcId string, projectId int64) ([]*tke.RunInstancesForNode, error) {
	roleGroups := make(map[string][]map[string]interface{})
	for _, block := range addedBlocks {
		role := "MASTER_ETCD"
		if value, ok := block["node_role"].(string); ok && value != "" {
			role = value
		}
		roleGroups[role] = append(roleGroups[role], block)
	}

	roles := make([]string, 0, len(roleGroups))
	for role := range roleGroups {
		roles = append(roles, role)
	}
	sort.Strings(roles)
	runInstancesForNodeList := make([]*tke.RunInstancesForNode, 0, len(roles))
	for _, role := range roles {
		blocks := roleGroups[role]
		runInstancesParaList := make([]*string, 0, len(blocks))
		overrideSettings := make([]*tke.InstanceAdvancedSettings, 0, len(blocks))
		hasInstanceOverrides := false
		for _, block := range blocks {
			paraJSON, _, err := tkeGetCvmRunInstancesPara(block, meta, vpcId, projectId)
			if err != nil {
				return nil, err
			}
			runInstancesParaList = append(runInstancesParaList, &paraJSON)

			override, hasOverride, err := expandMasterConfigInstanceAdvancedSettings(block)
			if err != nil {
				return nil, err
			}
			overrideSettings = append(overrideSettings, &override)
			hasInstanceOverrides = hasInstanceOverrides || hasOverride
		}
		if !hasInstanceOverrides {
			overrideSettings = nil
		}
		nodeRole := role
		runInstancesForNodeList = append(runInstancesForNodeList, &tke.RunInstancesForNode{
			NodeRole:                          &nodeRole,
			RunInstancesPara:                  runInstancesParaList,
			InstanceAdvancedSettingsOverrides: overrideSettings,
		})
	}
	return runInstancesForNodeList, nil
}

func buildScaleInMasterRequest(removedBlocks []map[string]interface{}) ([]*tke.ScaleInMaster, error) {
	scaleInMasters := make([]*tke.ScaleInMaster, 0, len(removedBlocks))
	for _, block := range removedBlocks {
		instanceId, _ := block["instance_id"].(string)
		instanceName, _ := block["instance_name"].(string)
		if instanceId == "" {
			return nil, fmt.Errorf("instance_id for master %q not found in state, cannot safely perform scale-in. Please refresh state or check configuration", instanceName)
		}

		role := "MASTER_ETCD"
		if value, ok := block["node_role"].(string); ok && value != "" {
			role = value
		}
		deleteMode := "terminate"
		scaleInMasters = append(scaleInMasters, &tke.ScaleInMaster{
			InstanceId:         &instanceId,
			NodeRole:           &role,
			InstanceDeleteMode: &deleteMode,
		})
	}
	return scaleInMasters, nil
}

func TkeCvmCreateInfo() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"count": {
			Type:        schema.TypeInt,
			Optional:    true,
			ForceNew:    true,
			Default:     1,
			Description: "Number of cvm.",
		},
		"availability_zone": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Description: "Indicates which availability zone will be used.",
		},
		"instance_name": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Default:     "sub machine of tke",
			Description: "Name of the CVMs.",
		},
		"instance_type": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
			Description: "Specified types of CVM instance.",
		},
		// payment
		"instance_charge_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			Default:      CVM_CHARGE_TYPE_POSTPAID,
			ValidateFunc: validateAllowedStringValue(TKE_INSTANCE_CHARGE_TYPE),
			Description:  "The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`. Note: Cloud International only supports `POSTPAID_BY_HOUR`, `PREPAID` instance will not terminated after cluster deleted, and may not allow to delete before expired.",
		},
		"instance_charge_type_prepaid_period": {
			Type:         schema.TypeInt,
			Optional:     true,
			ForceNew:     true,
			Default:      1,
			ValidateFunc: validateAllowedIntValue(CVM_PREPAID_PERIOD),
			Description:  "The tenancy (time unit is month) of the prepaid instance. NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.",
		},
		"instance_charge_type_prepaid_renew_flag": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			Default:      CVM_PREPAID_RENEW_FLAG_NOTIFY_AND_MANUAL_RENEW,
			ValidateFunc: validateAllowedStringValue(CVM_PREPAID_RENEW_FLAG),
			Description:  "Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.",
		},
		"subnet_id": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Required:     true,
			ValidateFunc: validateStringLengthInRange(4, 100),
			Description:  "Private network ID.",
		},
		"system_disk_type": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Optional:     true,
			Default:      SYSTEM_DISK_TYPE_CLOUD_PREMIUM,
			ValidateFunc: validateAllowedStringValue(SYSTEM_DISK_ALLOW_TYPE),
			Description:  "System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage. NOTE: `CLOUD_BASIC`, `LOCAL_BASIC` and `LOCAL_SSD` are deprecated.",
		},
		"system_disk_pool_group": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
			Description: "System disk pool group",
		},
		"system_disk_size": {
			Type:         schema.TypeInt,
			ForceNew:     true,
			Optional:     true,
			Default:      50,
			ValidateFunc: validateIntegerInRange(50, 500),
			Description:  "Volume of system disk in GB. Default is `50`.",
		},
		"data_disk": {
			Type:        schema.TypeList,
			ForceNew:    true,
			Optional:    true,
			MaxItems:    11,
			Description: "Configurations of data disk.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"disk_id": {
						Type:         schema.TypeString,
						ForceNew:     true,
						Required:     true,
						ValidateFunc: validateNotEmpty,
						Description:  "Data disk ID.",
					},
					"disk_partition": {
						Type:         schema.TypeString,
						ForceNew:     true,
						Required:     true,
						ValidateFunc: validateNotEmpty,
						Description:  "The device or partition name to mount.",
					},
					"disk_type": {
						Type:         schema.TypeString,
						ForceNew:     true,
						Optional:     true,
						Default:      SYSTEM_DISK_TYPE_CLOUD_PREMIUM,
						ValidateFunc: validateAllowedStringValue(SYSTEM_DISK_ALLOW_TYPE),
						Description:  "Types of disk, available values: `CLOUD_PREMIUM` and `CLOUD_SSD` and `CLOUD_HSSD` and `CLOUD_TSSD`.",
					},
					"disk_pool_group": {
						Type:        schema.TypeString,
						ForceNew:    true,
						Optional:    true,
						Description: "disk pool group",
					},
					"disk_size": {
						Type:        schema.TypeInt,
						ForceNew:    true,
						Optional:    true,
						Default:     0,
						Description: "Volume of disk in GB. Default is `0`.",
					},
					"snapshot_id": {
						Type:        schema.TypeString,
						ForceNew:    true,
						Optional:    true,
						Description: "Data disk snapshot ID.",
					},
					// "encrypt": {
					// 	Type:        schema.TypeBool,
					// 	Optional:    true,
					// 	Description: "Indicates whether to encrypt data disk, default `false`.",
					// },
					// "kms_key_id": {
					// 	Type:        schema.TypeString,
					// 	Optional:    true,
					// 	Description: "ID of the custom CMK in the format of UUID or `kms-abcd1234`. This parameter is used to encrypt cloud disks.",
					// },
					"file_system": {
						Type:        schema.TypeString,
						ForceNew:    true,
						Optional:    true,
						Description: "File system, e.g. `ext3/ext4/xfs`.",
					},
					"auto_format_and_mount": {
						Type:        schema.TypeBool,
						ForceNew:    true,
						Optional:    true,
						Default:     false,
						Description: "Indicate whether to auto format and mount or not. Default is `false`.",
					},
					"mount_target": {
						Type:         schema.TypeString,
						ForceNew:     true,
						Required:     true,
						ValidateFunc: validateNotEmpty,
						Description:  "Mount target.",
					},
					// "disk_partition": {
					// 	Type:        schema.TypeString,
					// 	ForceNew:    true,
					// 	Optional:    true,
					// 	Description: "The name of the device or partition to mount.",
					// },
				},
			},
		},
		"internet_charge_type": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Optional:     true,
			Default:      INTERNET_CHARGE_TYPE_TRAFFIC_POSTPAID_BY_HOUR,
			ValidateFunc: validateAllowedStringValue(INTERNET_CHARGE_ALLOW_TYPE),
			Description:  "Charge types for network traffic. Available values include `TRAFFIC_POSTPAID_BY_HOUR`.",
		},
		"internet_max_bandwidth_out": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     0,
			Description: "Max bandwidth of Internet access in Mbps. Default is 0.",
		},
		// "bandwidth_package_id": {
		// 	Type:        schema.TypeString,
		// 	Optional:    true,
		// 	Description: "bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.",
		// },
		"public_ip_assigned": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Description: "Specify whether to assign an Internet IP address.",
		},
		"password": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Optional:     true,
			Sensitive:    true,
			ValidateFunc: validateAsConfigPassword,
			Description:  "Password to access, should be set if `key_ids` not set.",
		},
		"key_ids": {
			MaxItems:    1,
			Type:        schema.TypeList,
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "ID list of keys, should be set if `password` not set.",
		},
		"security_group_ids": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Security groups to which a CVM instance belongs.",
		},
		"enhanced_security_service": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     true,
			Description: "To specify whether to enable cloud security service. Default is TRUE.",
		},
		"enhanced_monitor_service": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     true,
			Description: "To specify whether to enable cloud monitor service. Default is TRUE.",
		},
		"enhanced_automation_service": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     true,
			Description: "To specify whether to enable automation service. Default is TRUE.",
		},
		"user_data": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Description: "Ase64-encoded User Data text, the length limit is 16KB.",
		},
		"cam_role_name": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Computed:    true,
			Description: "CAM role name authorized to access.",
		},
		"hostname": {
			Type:     schema.TypeString,
			ForceNew: true,
			Optional: true,
			Description: "The host name of the attached instance. " +
				"Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. " +
				"Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. " +
				"Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).",
		},
		//"disaster_recover_group_ids": {
		//	Type:        schema.TypeList,
		//	ForceNew:    true,
		//	Optional:    true,
		//	MaxItems:    1,
		//	Elem:        &schema.Schema{Type: schema.TypeString},
		//	Description: "Disaster recover groups to which a CVM instance belongs. Only support maximum 1.",
		//},
		"img_id": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateImageID,
			Description:  "The valid image id, format of img-xxx.",
		},
		// InstanceAdvancedSettingsOverrides
		"desired_pod_num": {
			Type:     schema.TypeInt,
			ForceNew: true,
			Optional: true,
			Default:  DefaultDesiredPodNum,
			Description: "Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, " +
				"and it override `[globe_]desired_pod_num` for current node. Either all the fields `desired_pod_num` or none.",
		},
		// "hpc_cluster_id": {
		// 	Type:        schema.TypeString,
		// 	Optional:    true,
		// 	Description: "Id of cvm hpc cluster.",
		// },
	}
}

func TkeExistCvmCreateInfo() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"node_role": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Optional:     true,
			ValidateFunc: validateAllowedStringValue([]string{TKE_ROLE_WORKER, TKE_ROLE_MASTER_ETCD}),
			Description:  "Role of existed node. value:MASTER_ETCD or WORKER.",
		},
		"instances_para": {
			Type:     schema.TypeList,
			ForceNew: true,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"instance_ids": {
						Type:        schema.TypeList,
						ForceNew:    true,
						Required:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "Cluster IDs.",
					},
					"security_group_ids": {
						Type:        schema.TypeList,
						ForceNew:    true,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "Security groups to which a CVM instance belongs.",
					},
					"enhanced_service": {
						Type:     schema.TypeList,
						ForceNew: true,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"security_service": {
									Type:        schema.TypeBool,
									ForceNew:    true,
									Optional:    true,
									Description: "Enable cloud security service.",
								},
								"monitor_service": {
									Type:        schema.TypeBool,
									ForceNew:    true,
									Optional:    true,
									Description: "Enable cloud monitor service.",
								},
								"automation_service": {
									Type:        schema.TypeBool,
									ForceNew:    true,
									Optional:    true,
									Description: "Enable automation service.",
								},
							},
						},
						Description: "Enhanced service settings.",
					},
					"host_name": {
						Type:        schema.TypeString,
						ForceNew:    true,
						Optional:    true,
						Description: "Host name of the instance.",
					},
					"skip_options": {
						Type:        schema.TypeList,
						ForceNew:    true,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "Skip options for existing instances.",
					},
				},
			},
			Description: "Reinstallation parameters of an existing instance.",
		},
		"desired_pod_numbers": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Elem:        &schema.Schema{Type: schema.TypeInt},
			Description: "Custom mode cluster, you can specify the number of pods for each node. corresponding to the existed_instances_para.instance_ids parameter.",
		},
	}
}

func TkeNodePoolGlobalConfig() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"is_scale_in_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates whether to enable scale-in.",
		},
		"expander": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Indicates which scale-out method will be used when there are multiple scaling groups. Valid values: `random` - select a random scaling group, `most-pods` - select the scaling group that can schedule the most pods, `least-waste` - select the scaling group that can ensure the fewest remaining resources after Pod scheduling.",
		},
		"max_concurrent_scale_in": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Max concurrent scale-in volume.",
		},
		"scale_in_delay": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Number of minutes after cluster scale-out when the system starts judging whether to perform scale-in.",
		},
		"scale_in_unneeded_time": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Number of consecutive minutes of idleness after which the node is subject to scale-in.",
		},
		"scale_in_utilization_threshold": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Percentage of node resource usage below which the node is considered to be idle.",
		},
		"ignore_daemon_sets_utilization": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Whether to ignore DaemonSet pods by default when calculating resource usage.",
		},
		"skip_nodes_with_local_storage": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "During scale-in, ignore nodes with local storage pods.",
		},
		"skip_nodes_with_system_pods": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "During scale-in, ignore nodes with pods in the kube-system namespace that are not managed by DaemonSet.",
		},
	}
}

func resourceTencentCloudTkeCluster() *schema.Resource {
	schemaBody := map[string]*schema.Schema{
		"cluster_arch": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cpu arch of the cluster.",
		},
		// "node_pool_id": {
		// 	Type:        schema.TypeString,
		// 	Required:    true,
		// 	Description: "ID of the node pool.",
		// },
		"cluster_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the cluster.",
		},
		"cluster_desc": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of the cluster.",
		},
		"cluster_os": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Default:     TKE_CLUSTER_OS_LINUX24,
			Description: "Cluster operating system, supports setting public images (image Name) and custom images (image ID).",
		},
		"cluster_os_type": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Optional:     true,
			Default:      TKE_CLUSTER_OS_TYPE_GENERAL,
			ValidateFunc: validateAllowedStringValue(TKE_CLUSTER_OS_TYPES),
			Description: "Image type of the cluster os, the available values include: '" + strings.Join(TKE_CLUSTER_OS_TYPES, "','") +
				"'. Default is '" + TKE_CLUSTER_OS_TYPE_GENERAL + "'.",
		},
		"cluster_subnet_id": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Description: "Control Plane Subnet Information. Required for some network plugins (for example, CiliumOverlay).",
		},
		"container_runtime": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Optional:     true,
			Default:      TKE_RUNTIME_DOCKER,
			ValidateFunc: validateAllowedStringValue(TKE_RUNTIMES),
			Description: "Runtime type of the cluster, the available values include: 'docker' and 'containerd'." +
				"The Kubernetes v1.24 has removed dockershim, so please use containerd in v1.24 or higher." +
				"Default is 'docker'.",
		},
		"cluster_deploy_type": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Optional:     true,
			Default:      TKE_DEPLOY_TYPE_MANAGED,
			ValidateFunc: validateAllowedStringValue(TKE_DEPLOY_TYPES),
			Description:  "Deployment type of the cluster, the available values include: 'MANAGED_CLUSTER' and 'INDEPENDENT_CLUSTER'. Default is 'MANAGED_CLUSTER'.",
		},
		"cluster_version": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "1.10.5",
			Description: "Version of the cluster, Default is '1.10.5'. Use `tencentcloudenterprise_tke_kubernetes_available_cluster_versions` to get the available versions.",
		},
		//"upgrade_instances_follow_cluster": {
		//	Type:        schema.TypeBool,
		//	Optional:    true,
		//	Default:     false,
		//	Description: "Indicates whether upgrade all instances when cluster_version change. Default is false.",
		//},
		"cluster_ipvs": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     true,
			Description: "Indicates whether `ipvs` is enabled. Default is true. False means `iptables` is enabled.",
		},
		//"cluster_as_enabled": {
		//	Type:        schema.TypeBool,
		//	ForceNew:    true,
		//	Optional:    true,
		//	Default:     false,
		//	Deprecated:  "This argument is deprecated because the TKE auto-scaling group was no longer available.",
		//	Description: "Indicates whether to enable cluster node auto scaling. Default is false.",
		//},
		"cluster_level": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Specify cluster level, valid for managed cluster, use data source `tencentcloudenterprise_kubernetes_cluster_levels` to query available levels. Available value examples `L5`, `L20`, `L50`, `L100`, etc.",
		},
		"auto_upgrade_cluster_level": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether the cluster level auto upgraded, valid for managed cluster.",
		},
		//"acquire_cluster_admin_role": {
		//	Type:        schema.TypeBool,
		//	Optional:    true,
		//	Description: "If set to true, it will acquire the ClusterRole tke:admin. NOTE: this arguments cannot revoke to `false` after acquired.",
		//},
		//"node_pool_global_config": {
		//	Type:     schema.TypeList,
		//	Optional: true,
		//	Computed: true,
		//	Elem: &schema.Resource{
		//		Schema: TkeNodePoolGlobalConfig(),
		//	},
		//	Description: "Global config effective for all node pools.",
		//},
		"cluster_extra_args": {
			Type:     schema.TypeList,
			ForceNew: true,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"kube_apiserver": {
						Type:        schema.TypeList,
						ForceNew:    true,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "The customized parameters for kube-apiserver.",
					},
					"kube_controller_manager": {
						Type:        schema.TypeList,
						ForceNew:    true,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "The customized parameters for kube-controller-manager.",
					},
					"kube_scheduler": {
						Type:        schema.TypeList,
						ForceNew:    true,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "The customized parameters for kube-scheduler.",
					},
				},
			},
			Description: "Customized parameters for master component,such as kube-apiserver, kube-controller-manager, kube-scheduler.",
		},
		"node_name_type": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Optional:     true,
			Default:      "lan-ip",
			Description:  "Node name type of Cluster, the available values include: 'lan-ip' and 'hostname', Default is 'lan-ip'.",
			ValidateFunc: validateAllowedStringValue(TKE_CLUSTER_NODE_NAME_TYPE),
		},
		"network_type": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateAllowedStringValue(TKE_CLUSTER_NETWORK_TYPE),
			Description:  "Cluster network type, GR or VPC-CNI. Default is GR.",
		},
		//"enable_customized_pod_cidr": {
		//	Type: schema.TypeBool,
		//	//ForceNew:    true,
		//	Optional:    true,
		//	Default:     false,
		//	Description: "Whether to enable the custom mode of node podCIDR size. Default is false.",
		//},
		"base_pod_num": {
			Type:        schema.TypeInt,
			ForceNew:    true,
			Optional:    true,
			Description: "The number of basic pods. valid when enable_customized_pod_cidr=true.",
		},
		"is_non_static_ip_mode": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     false,
			Description: "Indicates whether non-static ip mode is enabled. Default is false.",
		},
		"deletion_protection": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Indicates whether cluster deletion protection is enabled. Default is false.",
		},
		"kube_proxy_mode": {
			Type:     schema.TypeString,
			ForceNew: true,
			Optional: true,
			Default:  "",
			Description: "Cluster kube-proxy mode, the available values include: 'kube-proxy-bpf'. Default is not set." +
				"When set to kube-proxy-bpf, cluster version greater than 1.14 and with Linux 2.4 is required.",
		},
		"runtime_version": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Description: "Container runtime version.",
		},
		"enable_customized_pod_cidr": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     false,
			Description: "Whether to enable the custom mode of node podCIDR size. Default is false.",
		},
		"audit_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates whether cluster audit is enabled. Default is false.",
		},
		"audit_logset_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Audit log logset ID.",
		},
		"audit_log_topic_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Audit log topic ID.",
		},
		"data_plane_v2": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     false,
			Description: "Whether to use data plane V2 (cilium v2). Default is false.",
		},
		"qgpu_share_enable": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     false,
			Description: "Indicates whether QGPU sharing is enabled. Default is false.",
		},
		"is_dual_stack": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     false,
			Description: "Indicates whether the cluster is dual stack (IPv4/IPv6). Default is false.",
		},
		"vpc_id": {
			Type:         schema.TypeString,
			ForceNew:     true,
			Required:     true,
			ValidateFunc: validateStringLengthInRange(4, 100),
			Description:  "Vpc Id of the cluster.",
		},
		//"cluster_internet": {
		//	Type:     schema.TypeBool,
		//	Default:  false,
		//	Optional: true,
		//	Description: "Open internet access or not." +
		//		" If this field is set 'true', the field below `worker_config` must be set." +
		//		" Because only cluster with node is allowed enable access endpoint.",
		//},
		//"cluster_internet_domain": {
		//	Type:     schema.TypeString,
		//	Optional: true,
		//	Description: "Domain name for cluster Kube-apiserver internet access." +
		//		" Be careful if you modify value of this parameter, the cluster_external_endpoint value may be changed automatically too.",
		//},
		//"cluster_intranet": {
		//	Type:     schema.TypeBool,
		//	Default:  false,
		//	Optional: true,
		//	Description: "Open intranet access or not." +
		//		" If this field is set 'true', the field below `worker_config` must be set." +
		//		" Because only cluster with node is allowed enable access endpoint.",
		//},
		//"cluster_intranet_domain": {
		//	Type:     schema.TypeString,
		//	Optional: true,
		//	Description: "Domain name for cluster Kube-apiserver intranet access." +
		//		" Be careful if you modify value of this parameter, the pgw_endpoint value may be changed automatically too.",
		//},
		//"cluster_internet_security_group": {
		//	Type:        schema.TypeString,
		//	Optional:    true,
		//	Description: "Specify security group, NOTE: This argument must not be empty if cluster internet enabled.",
		//},
		//"managed_cluster_internet_security_policies": {
		//	Type:       schema.TypeList,
		//	Optional:   true,
		//	Elem:       &schema.Schema{Type: schema.TypeString},
		//	Deprecated: "this argument was deprecated, use `cluster_internet_security_group` instead.",
		//	Description: "Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all." +
		//		" This field can only set when field `cluster_deploy_type` is 'MANAGED_CLUSTER' and `cluster_internet` is true." +
		//		" `managed_cluster_internet_security_policies` can not delete or empty once be set.",
		//},
		//"cluster_intranet_subnet_id": {
		//	Type:     schema.TypeString,
		//	Optional: true,
		//	Description: "Subnet id who can access this independent cluster, this field must and can only set  when `cluster_intranet` is true." +
		//		" `cluster_intranet_subnet_id` can not modify once be set.",
		//},
		"project_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Project ID, default value is 0.",
		},
		"need_work_security_group": {
			Type:        schema.TypeBool,
			Optional:    true,
			ForceNew:    true,
			Default:     false,
			Description: "Indicates whether to enable the default node security group. Default is false.",
		},
		"cluster_cidr": {
			Type: schema.TypeString,
			//ForceNew:    true,
			Optional:    true,
			Computed:    true,
			Description: "A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this vpc. Must be in  10./192.168/172.[16-31] segments.",
			ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
				value := v.(string)
				if value == "" {
					return
				}
				_, ipnet, err := net.ParseCIDR(value)
				if err != nil {
					errors = append(errors, fmt.Errorf("%q must contain a valid CIDR, got error parsing: %s", k, err))
					return
				}
				if ipnet == nil || value != ipnet.String() {
					errors = append(errors, fmt.Errorf("%q must contain a valid network CIDR, expected %q, got %q", k, ipnet, value))
					return
				}
				if !strings.Contains(value, "/") {
					errors = append(errors, fmt.Errorf("%q must be a network segment", k))
					return
				}
				if !strings.HasPrefix(value, "9.") && !strings.HasPrefix(value, "10.") && !strings.HasPrefix(value, "192.168.") && !strings.HasPrefix(value, "172.") {
					errors = append(errors, fmt.Errorf("%q must in 9. | 10. | 192.168. | 172.[16-31]", k))
					return
				}

				if strings.HasPrefix(value, "172.") {
					nextNo := strings.Split(value, ".")[1]
					no, _ := strconv.ParseInt(nextNo, 10, 64)
					if no < 16 || no > 31 {
						errors = append(errors, fmt.Errorf("%q must in 9.0 | 10. | 192.168. | 172.[16-31]", k))
						return
					}
				}
				return
			},
		},
		"ignore_cluster_cidr_conflict": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     false,
			Description: "Indicates whether to ignore the cluster cidr conflict error. Default is false.",
		},
		"cluster_max_pod_num": {
			Type:        schema.TypeInt,
			ForceNew:    true,
			Optional:    true,
			Default:     256,
			Description: "The maximum number of Pods per node in the cluster. Default is 256. The minimum value is 4. When its power unequal to 2, it will round upward to the closest power of 2.",
		},
		"cluster_max_service_num": {
			Type:        schema.TypeInt,
			ForceNew:    true,
			Optional:    true,
			Default:     256,
			Description: "The maximum number of services in the cluster. Default is 256. The range is from 32 to 32768. When its power unequal to 2, it will round upward to the closest power of 2.",
		},
		"service_cidr": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Description: "A network address block of the service. Different from vpc cidr and cidr of other clusters within this vpc. Must be in  10./192.168/172.[16-31] segments.",
			ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
				value := v.(string)
				if value == "" {
					return
				}
				_, ipnet, err := net.ParseCIDR(value)
				if err != nil {
					errors = append(errors, fmt.Errorf("%q must contain a valid CIDR, got error parsing: %s", k, err))
					return
				}
				if ipnet == nil || value != ipnet.String() {
					errors = append(errors, fmt.Errorf("%q must contain a valid network CIDR, expected %q, got %q", k, ipnet, value))
					return
				}
				if !strings.Contains(value, "/") {
					errors = append(errors, fmt.Errorf("%q must be a network segment", k))
					return
				}
				if !strings.HasPrefix(value, "9.") && !strings.HasPrefix(value, "10.") && !strings.HasPrefix(value, "192.168.") && !strings.HasPrefix(value, "172.") {
					errors = append(errors, fmt.Errorf("%q must in 9. | 10. | 192.168. | 172.[16-31]", k))
					return
				}

				if strings.HasPrefix(value, "172.") {
					nextNo := strings.Split(value, ".")[1]
					no, _ := strconv.ParseInt(nextNo, 10, 64)
					if no < 16 || no > 31 {
						errors = append(errors, fmt.Errorf("%q must in 9. | 10. | 192.168. | 172.[16-31]", k))
						return
					}
				}
				return
			},
		},
		"ignore_service_cidr_conflict": {
			Type:        schema.TypeBool,
			ForceNew:    true,
			Optional:    true,
			Default:     false,
			Description: "Indicates whether to ignore the service cidr conflict error. Only valid in VPC-CNI mode. Default is false.",
		},
		"eni_subnet_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Description: "Subnet Ids for cluster with VPC-CNI network mode." +
				" This field can only set when field `network_type` is 'VPC-CNI'." +
				" `eni_subnet_ids` can not empty once be set.",
		},
		"vpc_cni_type": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Distinguish between shared network card multi-IP mode and independent network card mode. Fill in `tke-route-eni` for shared network card multi-IP mode and `tke-direct-eni` for independent network card mode. The default is shared network card mode.",
			ValidateFunc: validateAllowedStringValue([]string{"tke-route-eni", "tke-direct-eni"}),
		},
		"claim_expired_seconds": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Description: "Claim expired seconds to recycle ENI." +
				" This field can only set when field `network_type` is 'VPC-CNI'." +
				" `claim_expired_seconds` must greater or equal than 300 and less than 15768000.",
			ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
				value := v.(int)
				if value < 300 || value > 15768000 {
					errors = append(errors, fmt.Errorf("%q must greater or equal than 300 and less than 15768000", k))
					return
				}
				return
			},
		},
		"master_config": {
			Type:     schema.TypeList,
			ForceNew: false,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: TkeMasterCvmCreateInfo(),
			},
			Description: "Deploy the machine configuration information of the 'MASTER_ETCD' service, and create <=7 units for common users.",
		},
		"worker_config": {
			Type:     schema.TypeList,
			ForceNew: true,
			Optional: true,
			Elem: &schema.Resource{
				Schema: TkeCvmCreateInfo(),
			},
			Description: "Deploy the machine configuration information of the 'WORKER' service, and create <=20 units for common users. The other 'WORK' service are added by 'tencentcloudenterprise_kubernetes_worker'.",
		},
		"exist_instance": {
			Type:     schema.TypeList,
			ForceNew: true,
			Optional: true,
			Elem: &schema.Resource{
				Schema: TkeExistCvmCreateInfo(),
			},
			Description: "Create tke cluster by existed instances.",
		},
		"run_instances_for_node": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Description: "RunInstancesForNode settings to build CreateCluster request directly.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"node_role": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Node role, value: MASTER_ETCD or WORKER.",
					},
					"run_instances_para": {
						Type:        schema.TypeList,
						Required:    true,
						Description: "CVM run instances parameters in JSON string format.",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"instance_advanced_settings_overrides": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Per-node instance advanced settings overrides. Order must match run_instances_para.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"docker_graph_path": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Docker graph path.",
								},
								"unschedulable": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Sets whether the joining node participates in the schedule.",
								},
								"pre_start_user_script": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Base64-encoded user script, executed before initializing the node.",
								},
								"user_script": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Base64-encoded user script executed after initializing the node.",
								},
								"labels": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "Node label list.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"name": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Label name.",
											},
											"value": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Label value.",
											},
										},
									},
								},
								"taints": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "Node taint.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"key": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Key of the taint.",
											},
											"value": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Value of the taint.",
											},
											"effect": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Effect of the taint.",
											},
										},
									},
								},
								"extra_args": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "Custom parameter information related to the node.",
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"data_disks": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "Data disks for instance advanced settings override.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"disk_type": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Types of disk.",
											},
											"disk_size": {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "Volume of disk in GB.",
											},
											"disk_id": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Data disk ID.",
											},
											"disk_partition": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "The name of the device or partition to mount.",
											},
											"file_system": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "File system, e.g. `ext3/ext4/xfs`.",
											},
											"auto_format_and_mount": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Indicate whether to auto format and mount or not.",
											},
											"mount_target": {
												Type:        schema.TypeString,
												Required:    true,
												Description: "Mount target.",
											},
										},
									},
								},
								"desired_pod_number": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Indicate to set desired pod number in node. valid when the cluster is podCIDR.",
								},
								"gpu_args": {
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "GPU driver parameters.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"mig_enable": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Whether to enable MIG.",
											},
											"driver": {
												Type:        schema.TypeMap,
												Optional:    true,
												Description: "GPU driver version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.",
											},
											"cuda": {
												Type:        schema.TypeMap,
												Optional:    true,
												Description: "CUDA  version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.",
											},
											"cudnn": {
												Type:        schema.TypeMap,
												Optional:    true,
												Description: "cuDNN version. Format like: `{ version: String, name: String, doc_name: String, dev_name: String }`. `version`: cuDNN version; `name`: cuDNN name; `doc_name`: Doc name of cuDNN; `dev_name`: Dev name of cuDNN.",
											},
											"custom_driver": {
												Type:        schema.TypeMap,
												Optional:    true,
												Description: "Custom GPU driver. Format like: `{address: String}`. `address`: URL of custom GPU driver address.",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		//"extension_addon": {
		//	Type:        schema.TypeList,
		//	Optional:    true,
		//	Description: "Information of the add-on to be installed.",
		//	Elem: &schema.Resource{
		//		Schema: map[string]*schema.Schema{
		//			"name": {
		//				Type:        schema.TypeString,
		//				Required:    true,
		//				Description: "Add-on name.",
		//			},
		//			"param": {
		//				Type:             schema.TypeString,
		//				Required:         true,
		//				DiffSuppressFunc: helper.DiffSupressJSON,
		//				Description:      "Parameter of the add-on resource object in JSON string format, please check the example at the top of page for reference.",
		//			},
		//		},
		//	},
		//},
		"log_agent": {
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "Specify cluster log agent config.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:        schema.TypeBool,
						Required:    true,
						Description: "Whether the log agent enabled.",
					},
					"kubelet_root_dir": {
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
						Description: "Kubelet root directory as the literal.",
					},
				},
			},
		},
		"event_persistence": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify cluster Event Persistence config. NOTE: Please make sure your TKE CamRole have permission to access CLS service.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:        schema.TypeBool,
						Required:    true,
						Description: "Specify weather the Event Persistence enabled.",
					},
					"log_set_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Specify id of existing CLS log set, or auto create a new set by leave it empty.",
					},
					"topic_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Specify id of existing CLS log topic, or auto create a new topic by leave it empty.",
					},
					"delete_event_log_and_topic": {
						Type:     schema.TypeBool,
						Optional: true,
						Description: "When you want to close the cluster event persistence or delete the cluster, you can use this parameter to determine " +
							"whether the event persistence log set and topic created by default will be deleted.",
					},
				},
			},
		},
		"cluster_audit": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specify Cluster Audit config. NOTE: Please make sure your TKE CamRole have permission to access CLS service.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:        schema.TypeBool,
						Required:    true,
						Description: "Specify weather the Cluster Audit enabled. NOTE: Enable Cluster Audit will also auto install Log Agent.",
					},
					"log_set_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Specify id of existing CLS log set, or auto create a new set by leave it empty.",
					},
					"topic_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Specify id of existing CLS log topic, or auto create a new topic by leave it empty.",
					},
					"delete_audit_log_and_topic": {
						Type:     schema.TypeBool,
						Optional: true,
						Description: "When you want to close the cluster audit log or delete the cluster, you can use " +
							"this parameter to determine whether the audit log set and topic created by default will" +
							" be deleted.",
					},
				},
			},
		},
		"tags": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "The tags of the cluster.",
		},
		"extension_addon": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Add-on name.",
					},
					"param": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Add-on parameters (JSON string format), please check the example at the top of page for reference.",
					},
				},
			},
			Description: "List of extension add-ons to be installed.",
		},

		"disable_addons": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of add-on names to disable during cluster creation. For example: [\"ip-masq-agent\"].",
		},

		// Computed values
		"cluster_node_num": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of nodes in the cluster.",
		},
		"worker_instances_list": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: tkeCvmState(),
			},
			Description: "An information list of cvm within the 'WORKER' clusters. Each element contains the following attributes:",
		},
		//advanced instance setting
		"labels": {
			Type:        schema.TypeMap,
			Optional:    true,
			ForceNew:    true,
			Description: "Labels of tke cluster nodes.",
		},
		"unschedulable": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if new == "0" && old == "" {
					return true
				} else {
					return old == new
				}
			},
			Default:     0,
			Description: "Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.",
		},
		"mount_target": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "Mount target. Default is not mounting.",
		},
		"globe_desired_pod_num": {
			Type:        schema.TypeInt,
			ForceNew:    true,
			Optional:    true,
			Description: "Indicate to set desired pod number in node. valid when enable_customized_pod_cidr=true, and it takes effect for all nodes.",
		},
		"docker_graph_path": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
			Default:  "/var/lib/docker",
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if new == "/var/lib/docker" && old == "" || old == "/var/lib/docker" && new == "" {
					return true
				} else {
					return old == new
				}
			},
			ValidateFunc: validateNotEmpty,
			Description:  "Docker graph path. Default is `/var/lib/docker`.",
		},
		"pre_start_user_script": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "Base64-encoded user script, executed before initializing the node, currently only effective for adding existing nodes.",
		},
		"user_script": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "Base64-encoded user script executed after initializing the node.",
		},
		"instance_data_disks": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Description: "Data disks for instance advanced settings.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"disk_type": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Types of disk.",
					},
					"disk_size": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Volume of disk in GB.",
					},
					"disk_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Data disk ID.",
					},
					"disk_partition": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The name of the device or partition to mount.",
					},
					"file_system": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "File system, e.g. `ext3/ext4/xfs`.",
					},
					"auto_format_and_mount": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Indicate whether to auto format and mount or not.",
					},
					"mount_target": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Mount target.",
					},
				},
			},
		},
		"extra_args": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Custom parameter information related to the node.",
		},
		"taints": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"key": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Key of the taint.",
					},
					"value": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Value of the taint.",
					},
					"effect": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Effect of the taint. Valid values are: `NoSchedule`, `PreferNoSchedule`, `NoExecute`.",
					},
				},
			},
			Description: "Node taint.",
		},
		//"runtime_version": {
		//	Type:        schema.TypeString,
		//	Optional:    true,
		//	Description: "Container Runtime version.",
		//},

		"kube_config": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Kubernetes config.",
		},
		"kube_config_intranet": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Kubernetes config of private network.",
		},
	}

	for k, v := range tkeSecurityInfo() {
		schemaBody[k] = v
	}

	return &schema.Resource{
		Description: "Provides a cloud tke cluster resource.",
		Create:      resourceTencentCloudTkeClusterCreate,
		Read:        resourceTencentCloudTkeClusterRead,
		Update:      resourceTencentCloudTkeClusterUpdate,
		Delete:      resourceTencentCloudTkeClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: schemaBody,
	}
}

func tkeGetCvmRunInstancesPara(dMap map[string]interface{}, meta interface{},
	vpcId string, projectId int64) (cvmJson string, count int64, errRet error) {

	request := cvm.NewRunInstancesRequest()

	var place cvm.Placement
	request.Placement = &place

	place.ProjectId = &projectId

	if v, ok := dMap["availability_zone"]; ok {
		place.Zone = helper.String(v.(string))
	}

	if v, ok := dMap["instance_type"]; ok {
		request.InstanceType = helper.String(v.(string))
	} else {
		errRet = fmt.Errorf("instance_type must be set.")
		return
	}

	subnetId := ""

	if v, ok := dMap["subnet_id"]; ok {
		subnetId = v.(string)
	}

	if (vpcId == "" && subnetId != "") ||
		(vpcId != "" && subnetId == "") {
		errRet = fmt.Errorf("Parameters cvm.`subnet_id` and cluster.`vpc_id` are both set or neither")
		return
	}

	if vpcId != "" {
		request.VirtualPrivateCloud = &cvm.VirtualPrivateCloud{
			VpcId:    &vpcId,
			SubnetId: &subnetId,
		}
	}

	if v, ok := dMap["system_disk_type"]; ok {
		if request.SystemDisk == nil {
			request.SystemDisk = &cvm.SystemDisk{}
		}
		request.SystemDisk.DiskType = helper.String(v.(string))
	}

	if v, ok := dMap["system_disk_size"]; ok {
		if request.SystemDisk == nil {
			request.SystemDisk = &cvm.SystemDisk{}
		}
		request.SystemDisk.DiskSize = helper.Int64(int64(v.(int)))

	}

	if v, ok := dMap["system_disk_pool_group"]; ok {
		if request.SystemDisk == nil {
			request.SystemDisk = &cvm.SystemDisk{}
		}
		request.SystemDisk.DiskStoragePoolGroup = helper.String(v.(string))
	}

	camRoleName := ""
	if v, ok := dMap["cam_role_name"]; ok {
		camRoleName = v.(string)
	}

	if v, ok := dMap["data_disk"]; ok {

		dataDisks := v.([]interface{})
		request.DataDisks = make([]*cvm.DataDisk, 0, len(dataDisks))

		for _, d := range dataDisks {

			var (
				value         = d.(map[string]interface{})
				diskType      = value["disk_type"].(string)
				diskSize      = int64(value["disk_size"].(int))
				snapshotId    = value["snapshot_id"].(string)
				diskPoolGroup = value["disk_pool_group"].(string)
				//encrypt    = value["encrypt"].(bool)
				//kmsKeyId   = value["kms_key_id"].(string)
				dataDisk = cvm.DataDisk{
					DiskType:             &diskType,
					DiskStoragePoolGroup: &diskPoolGroup,
				}
			)
			if diskSize > 0 {
				dataDisk.DiskSize = &diskSize
			}
			if snapshotId != "" {
				dataDisk.SnapshotId = &snapshotId
			}
			/*
				if encrypt {
					dataDisk.Encrypt = &encrypt
				}
				if kmsKeyId != "" {
					dataDisk.KmsKeyId = &kmsKeyId
				}

			*/
			request.DataDisks = append(request.DataDisks, &dataDisk)
		}
	}

	if v, ok := dMap["internet_charge_type"]; ok {

		if request.InternetAccessible == nil {
			request.InternetAccessible = &cvm.InternetAccessible{}
		}
		request.InternetAccessible.InternetChargeType = helper.String(v.(string))
	}

	if v, ok := dMap["internet_max_bandwidth_out"]; ok {
		if request.InternetAccessible == nil {
			request.InternetAccessible = &cvm.InternetAccessible{}
		}
		request.InternetAccessible.InternetMaxBandwidthOut = helper.Int64(int64(v.(int)))
	}

	/*
		if v, ok := dMap["bandwidth_package_id"]; ok {
			if v.(string) != "" {
				request.InternetAccessible.BandwidthPackageId = helper.String(v.(string))
			}
		}

	*/

	if v, ok := dMap["public_ip_assigned"]; ok {
		publicIpAssigned := v.(bool)
		request.InternetAccessible.PublicIpAssigned = &publicIpAssigned
	}

	if v, ok := dMap["password"]; ok {
		if request.LoginSettings == nil {
			request.LoginSettings = &cvm.LoginSettings{}
		}

		if v.(string) != "" {
			request.LoginSettings.Password = helper.String(v.(string))
		}
	}

	if v, ok := dMap["instance_name"]; ok {
		request.InstanceName = helper.String(v.(string))
	}

	if v, ok := dMap["key_ids"]; ok {
		if request.LoginSettings == nil {
			request.LoginSettings = &cvm.LoginSettings{}
		}
		keyIds := v.([]interface{})

		if len(keyIds) != 0 {
			request.LoginSettings.KeyIds = make([]*string, 0, len(keyIds))
			for i := range keyIds {
				keyId := keyIds[i].(string)
				request.LoginSettings.KeyIds = append(request.LoginSettings.KeyIds, &keyId)
			}
		}
	}

	if request.LoginSettings.Password == nil && len(request.LoginSettings.KeyIds) == 0 {
		errRet = fmt.Errorf("Parameters cvm.`key_ids` and cluster.`password` should be set one")
		return
	}

	if request.LoginSettings.Password != nil && len(request.LoginSettings.KeyIds) != 0 {
		errRet = fmt.Errorf("Parameters cvm.`key_ids` and cluster.`password` can only be supported one")
		return
	}

	if v, ok := dMap["security_group_ids"]; ok {
		securityGroups := v.([]interface{})
		request.SecurityGroupIds = make([]*string, 0, len(securityGroups))
		for i := range securityGroups {
			securityGroup := securityGroups[i].(string)
			request.SecurityGroupIds = append(request.SecurityGroupIds, &securityGroup)
		}
	}

	if v, ok := dMap["disaster_recover_group_ids"]; ok {
		disasterGroups := v.([]interface{})
		request.DisasterRecoverGroupIds = make([]*string, 0, len(disasterGroups))
		for i := range disasterGroups {
			disasterGroup := disasterGroups[i].(string)
			request.DisasterRecoverGroupIds = append(request.DisasterRecoverGroupIds, &disasterGroup)
		}
	}

	if v, ok := dMap["enhanced_security_service"]; ok {

		if request.EnhancedService == nil {
			request.EnhancedService = &cvm.EnhancedService{}
		}

		securityService := v.(bool)
		request.EnhancedService.SecurityService = &cvm.RunSecurityServiceEnabled{
			Enabled: &securityService,
		}
	}
	if v, ok := dMap["enhanced_monitor_service"]; ok {
		if request.EnhancedService == nil {
			request.EnhancedService = &cvm.EnhancedService{}
		}
		monitorService := v.(bool)
		request.EnhancedService.MonitorService = &cvm.RunMonitorServiceEnabled{
			Enabled: &monitorService,
		}
	}
	if v, ok := dMap["enhanced_automation_service"]; ok {
		if request.EnhancedService == nil {
			request.EnhancedService = &cvm.EnhancedService{}
		}
		automationService := v.(bool)
		request.EnhancedService.AutomationService = &cvm.AutomationServiceEnabled{
			Enabled: &automationService,
		}
	}
	if v, ok := dMap["user_data"]; ok {
		request.UserData = helper.String(v.(string))
	}
	if v, ok := dMap["instance_charge_type"]; ok {
		instanceChargeType := v.(string)
		request.InstanceChargeType = &instanceChargeType
		if instanceChargeType == CVM_CHARGE_TYPE_PREPAID {
			request.InstanceChargePrepaid = &cvm.InstanceChargePrepaid{}
			if period, ok := dMap["instance_charge_type_prepaid_period"]; ok {
				periodInt64 := int64(period.(int))
				request.InstanceChargePrepaid.Period = &periodInt64
			} else {
				errRet = fmt.Errorf("instance charge type prepaid period can not be empty when charge type is %s",
					instanceChargeType)
				return
			}
			if renewFlag, ok := dMap["instance_charge_type_prepaid_renew_flag"]; ok {
				request.InstanceChargePrepaid.RenewFlag = helper.String(renewFlag.(string))
			}
		}
	}
	if v, ok := dMap["count"]; ok {
		count = int64(v.(int))
	} else {
		count = 1
	}
	request.InstanceCount = &count

	if v, ok := dMap["hostname"]; ok {
		hostname := v.(string)
		if hostname != "" {
			request.HostName = &hostname
		}
	}

	if v, ok := dMap["img_id"]; ok && v.(string) != "" {
		request.ImageId = helper.String(v.(string))
	}

	/*
		if v, ok := dMap["hpc_cluster_id"]; ok && v.(string) != "" {
			request.HpcClusterId = helper.String(v.(string))
		}

	*/

	cvmJson = request.ToJsonString()

	cvmJson = strings.Replace(cvmJson, `"Password":"",`, "", -1)

	if camRoleName != "" {
		var cvmMap map[string]interface{}
		if err := json.Unmarshal([]byte(cvmJson), &cvmMap); err == nil {
			cvmMap["CamRoleName"] = camRoleName
			if b, err := json.Marshal(cvmMap); err == nil {
				cvmJson = string(b)
			}
		}
	}

	return
}

func tkeGetCvmExistInstancesPara(dMap map[string]interface{}) (tke.ExistedInstancesForNode, error) {

	inst := tke.ExistedInstancesForNode{}

	if temp, ok := dMap["instances_para"]; ok {
		paras := temp.([]interface{})
		if len(paras) > 0 {
			paraMap := paras[0].(map[string]interface{})
			instanceIds := paraMap["instance_ids"].([]interface{})
			inst.ExistedInstancesPara = &tke.ExistedInstancesPara{}
			inst.ExistedInstancesPara.InstanceIds = make([]*string, 0)
			for _, v := range instanceIds {
				inst.ExistedInstancesPara.InstanceIds = append(inst.ExistedInstancesPara.InstanceIds, helper.String(v.(string)))
			}
			if temp, ok := paraMap["security_group_ids"]; ok {
				sgs := temp.([]interface{})
				inst.ExistedInstancesPara.SecurityGroupIds = make([]*string, 0, len(sgs))
				for _, v := range sgs {
					inst.ExistedInstancesPara.SecurityGroupIds = append(inst.ExistedInstancesPara.SecurityGroupIds, helper.String(v.(string)))
				}
			}
			if temp, ok := paraMap["host_name"]; ok {
				hostName := temp.(string)
				if hostName != "" {
					inst.ExistedInstancesPara.HostName = helper.String(hostName)
				}
			}
			if temp, ok := paraMap["skip_options"]; ok {
				options := temp.([]interface{})
				inst.ExistedInstancesPara.SkipOptions = make([]*string, 0, len(options))
				for _, v := range options {
					inst.ExistedInstancesPara.SkipOptions = append(inst.ExistedInstancesPara.SkipOptions, helper.String(v.(string)))
				}
			}
			if temp, ok := paraMap["login_settings"]; ok {
				loginList := temp.([]interface{})
				if len(loginList) > 0 {
					loginMap := loginList[0].(map[string]interface{})
					login := &tke.LoginSettings{}
					if v, ok := loginMap["password"]; ok && v.(string) != "" {
						login.Password = helper.String(v.(string))
					}
					if v, ok := loginMap["key_ids"]; ok {
						keyIds := v.([]interface{})
						login.KeyIds = make([]*string, 0, len(keyIds))
						for _, key := range keyIds {
							login.KeyIds = append(login.KeyIds, helper.String(key.(string)))
						}
					}
					if v, ok := loginMap["keep_image_login"]; ok {
						if v.(bool) {
							login.KeepImageLogin = helper.String(CVM_IMAGE_LOGIN)
						} else {
							login.KeepImageLogin = helper.String(CVM_IMAGE_LOGIN_NOT)
						}
					}
					inst.ExistedInstancesPara.LoginSettings = login
				}
			}
			if temp, ok := paraMap["enhanced_service"]; ok {
				serviceList := temp.([]interface{})
				if len(serviceList) > 0 {
					serviceMap := serviceList[0].(map[string]interface{})
					enhanced := &tke.EnhancedService{}
					if v, ok := serviceMap["security_service"]; ok {
						enabled := v.(bool)
						enhanced.SecurityService = &tke.RunSecurityServiceEnabled{Enabled: &enabled}
					}
					if v, ok := serviceMap["monitor_service"]; ok {
						enabled := v.(bool)
						enhanced.MonitorService = &tke.RunMonitorServiceEnabled{Enabled: &enabled}
					}
					if v, ok := serviceMap["automation_service"]; ok {
						enabled := v.(bool)
						enhanced.AutomationService = &tke.RunAutomationServiceEnabled{Enabled: &enabled}
					}
					inst.ExistedInstancesPara.EnhancedService = enhanced
				}
			}
		}
	}
	//if temp, ok := dMap["desired_pod_numbers"]; ok {
	//	inst.DesiredPodNumbers = make([]*int64, 0)
	//	podNums := temp.([]interface{})
	//	for _, v := range podNums {
	//		inst.DesiredPodNumbers = append(inst.DesiredPodNumbers, helper.Int64(int64(v.(int))))
	//	}
	//}
	if temp, ok := dMap["node_role"]; ok {
		nodeRole := temp.(string)
		inst.NodeRole = &nodeRole
	}

	return inst, nil
}

func tkeGetNodePoolGlobalConfig(d *schema.ResourceData) *tke.ModifyClusterAsGroupOptionAttributeRequest {
	request := tke.NewModifyClusterAsGroupOptionAttributeRequest()
	request.ClusterId = helper.String(d.Id())

	clusterAsGroupOption := &tke.ClusterAsGroupOption{}
	if v, ok := d.GetOkExists("node_pool_global_config.0.is_scale_in_enabled"); ok {
		clusterAsGroupOption.IsScaleDownEnabled = helper.Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("node_pool_global_config.0.expander"); ok {
		clusterAsGroupOption.Expander = helper.String(v.(string))
	}
	if v, ok := d.GetOkExists("node_pool_global_config.0.max_concurrent_scale_in"); ok {
		clusterAsGroupOption.MaxEmptyBulkDelete = helper.IntInt64(v.(int))
	}
	if v, ok := d.GetOkExists("node_pool_global_config.0.scale_in_delay"); ok {
		clusterAsGroupOption.ScaleDownDelay = helper.IntInt64(v.(int))
	}
	if v, ok := d.GetOkExists("node_pool_global_config.0.scale_in_unneeded_time"); ok {
		clusterAsGroupOption.ScaleDownUnneededTime = helper.IntInt64(v.(int))
	}
	if v, ok := d.GetOkExists("node_pool_global_config.0.scale_in_utilization_threshold"); ok {
		clusterAsGroupOption.ScaleDownUtilizationThreshold = helper.IntInt64(v.(int))
	}
	if v, ok := d.GetOkExists("node_pool_global_config.0.ignore_daemon_sets_utilization"); ok {
		clusterAsGroupOption.IgnoreDaemonSetsUtilization = helper.Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("node_pool_global_config.0.skip_nodes_with_local_storage"); ok {
		clusterAsGroupOption.SkipNodesWithLocalStorage = helper.Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("node_pool_global_config.0.skip_nodes_with_system_pods"); ok {
		clusterAsGroupOption.SkipNodesWithSystemPods = helper.Bool(v.(bool))
	}

	request.ClusterAsGroupOption = clusterAsGroupOption
	return request
}

// upgradeClusterInstances upgrade instances, upgrade type try seq:major, hot.
func upgradeClusterInstances(tkeService TkeService, ctx context.Context, id string) error {
	// get all available instances for upgrade
	upgradeType := "major"
	instanceIds, err := tkeService.CheckInstancesUpgradeAble(ctx, id, upgradeType)
	if err != nil {
		return err
	}
	if len(instanceIds) == 0 {
		upgradeType = "hot"
		instanceIds, err = tkeService.CheckInstancesUpgradeAble(ctx, id, upgradeType)
		if err != nil {
			return err
		}
	}
	log.Println("instancesIds for upgrade:", instanceIds)
	instNum := len(instanceIds)
	if instNum == 0 {
		return nil
	}

	// upgrade instances
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		inErr := tkeService.UpgradeClusterInstances(ctx, id, upgradeType, instanceIds)
		if inErr != nil {
			return retryError(inErr)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// check update status: upgrade instance one by one, so timeout depend on instance number.
	timeout := readRetryTimeout * time.Duration(instNum)
	err = resource.Retry(timeout, func() *resource.RetryError {
		done, inErr := tkeService.GetUpgradeInstanceResult(ctx, id)
		if inErr != nil {
			return retryError(inErr)
		}
		if done {
			return nil
		} else {
			return resource.RetryableError(fmt.Errorf("cluster %s, retry...", id))
		}
	})
	if err != nil {
		return err
	}

	return nil
}

func resourceTencentCloudTkeClusterCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_cluster.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		basic               ClusterBasicSetting
		advanced            ClusterAdvancedSettings
		cvms                RunInstancesForNode
		iAdvanced           InstanceAdvancedSettings
		iDiskMountSettings  []*tke.InstanceDataDiskMountSetting
		cidrSet             tke.ClusterCIDRSettings
		extensionAddons     []*tke.ExtensionAddon
		runInstancesForNode []*tke.RunInstancesForNode
		clusterArch         string
		//clusterInternet              = d.Get("cluster_internet").(bool)
		//clusterIntranet              = d.Get("cluster_intranet").(bool)
		//intranetSubnetId             = d.Get("cluster_intranet_subnet_id").(string)
		//clusterInternetSecurityGroup = d.Get("cluster_internet_security_group").(string)
		//clusterInternetDomain        = d.Get("cluster_internet_domain").(string)
		//clusterIntranetDomain        = d.Get("cluster_intranet_domain").(string)
		// nodePoolId                   = d.Get("node_pool_id").(string)
	)
	if v, ok := d.GetOk("cluster_arch"); ok {
		clusterArch = v.(string)
	}

	clusterDeployType := d.Get("cluster_deploy_type").(string)
	runInstancesForNodeRaw, runInstancesForNodeOk := d.GetOk("run_instances_for_node")
	if runInstancesForNodeOk {
		if _, ok := d.GetOk("master_config"); ok {
			return fmt.Errorf("`run_instances_for_node` can not be used with `master_config`")
		}
		if _, ok := d.GetOk("worker_config"); ok {
			return fmt.Errorf("`run_instances_for_node` can not be used with `worker_config`")
		}
		if _, ok := d.GetOk("exist_instance"); ok {
			return fmt.Errorf("`run_instances_for_node` can not be used with `exist_instance`")
		}
		nodes, err := expandRunInstancesForNode(runInstancesForNodeRaw.([]interface{}))
		if err != nil {
			return err
		}
		runInstancesForNode = nodes
	}

	//if clusterIntranet && intranetSubnetId == "" {
	//	return fmt.Errorf("`cluster_intranet_subnet_id` must set when `cluster_intranet` is true")
	//}
	//if !clusterIntranet && intranetSubnetId != "" {
	//	return fmt.Errorf("`cluster_intranet_subnet_id` can only set when `cluster_intranet` is true")
	//}

	vpcId := d.Get("vpc_id").(string)
	if vpcId != "" {
		basic.VpcId = vpcId
	}

	basic.ProjectId = int64(d.Get("project_id").(int))
	basic.NeedWorkSecurityGroup = d.Get("need_work_security_group").(bool)

	cluster_os := d.Get("cluster_os").(string)

	if v, ok := tkeClusterOsMap[cluster_os]; ok {
		basic.ClusterOs = v
	} else {
		basic.ClusterOs = cluster_os
	}

	if tkeClusterOsMap[cluster_os] != "" {
		basic.ClusterOs = tkeClusterOsMap[cluster_os]
	} else {
		basic.ClusterOs = cluster_os
	}

	basic.ClusterOsType = d.Get("cluster_os_type").(string)

	basic.ClusterVersion = d.Get("cluster_version").(string)
	if v, ok := d.GetOk("cluster_name"); ok {
		basic.ClusterName = v.(string)
	}
	if v, ok := d.GetOk("cluster_desc"); ok {
		basic.ClusterDescription = v.(string)
	}
	if v, ok := d.GetOk("cluster_subnet_id"); ok {
		basic.SubnetId = v.(string)
	}

	if v, ok := d.GetOk("cluster_level"); ok {
		basic.ClusterLevel = helper.String(v.(string))
	}
	if v, ok := d.GetOkExists("auto_upgrade_cluster_level"); ok {
		basic.AutoUpgradeClusterLevel = helper.Bool(v.(bool))
	}

	advanced.Ipvs = d.Get("cluster_ipvs").(bool)
	//advanced.AsEnabled = d.Get("cluster_as_enabled").(bool)
	advanced.ContainerRuntime = d.Get("container_runtime").(string)
	if v, ok := d.GetOk("runtime_version"); ok {
		advanced.RuntimeVersion = v.(string)
	}
	advanced.NodeNameType = d.Get("node_name_type").(string)
	advanced.NetworkType = d.Get("network_type").(string)
	if advanced.NetworkType == "" {
		advanced.NetworkType = TKE_CLUSTER_NETWORK_TYPE_GR
	}
	advanced.IsNonStaticIpMode = d.Get("is_non_static_ip_mode").(bool)
	if advanced.NetworkType == TKE_CLUSTER_NETWORK_TYPE_VPC_CNI {
		if v, ok := d.GetOk("vpc_cni_type"); ok {
			advanced.VpcCniType = v.(string)
		} else {
			advanced.VpcCniType = "tke-route-eni"
		}
	}
	advanced.DeletionProtection = d.Get("deletion_protection").(bool)
	if v, ok := d.GetOk("kube_proxy_mode"); ok {
		advanced.KubeProxyMode = v.(string)
	}
	advanced.EnableCustomizedPodCIDR = d.Get("enable_customized_pod_cidr").(bool)
	advanced.AuditEnabled = d.Get("audit_enabled").(bool)
	if v, ok := d.GetOk("audit_logset_id"); ok {
		advanced.AuditLogsetId = v.(string)
	}
	if v, ok := d.GetOk("audit_log_topic_id"); ok {
		advanced.AuditLogTopicId = v.(string)
	}
	advanced.DataPlaneV2 = d.Get("data_plane_v2").(bool)
	advanced.QGPUShareEnable = d.Get("qgpu_share_enable").(bool)
	advanced.IsDualStack = d.Get("is_dual_stack").(bool)
	if v, ok := d.GetOk("base_pod_num"); ok {
		advanced.BasePodNumber = int64(v.(int))
	}

	if extraArgs, ok := d.GetOk("cluster_extra_args"); ok {
		extraArgList := extraArgs.([]interface{})
		for index := range extraArgList {
			raw := extraArgList[index]
			if raw == nil {
				continue
			}
			extraArg, ok := raw.(map[string]interface{})
			if !ok || extraArg == nil {
				continue
			}
			if apiserverArgs, exist := extraArg["kube_apiserver"]; exist {
				args := apiserverArgs.([]interface{})
				for index := range args {
					advanced.ExtraArgs.KubeAPIServer = append(advanced.ExtraArgs.KubeAPIServer, args[index].(string))
				}
			}
			if cmArgs, exist := extraArg["kube_controller_manager"]; exist {
				args := cmArgs.([]interface{})
				for index := range args {
					advanced.ExtraArgs.KubeControllerManager = append(advanced.ExtraArgs.KubeControllerManager, args[index].(string))
				}
			}
			if schedulerArgs, exist := extraArg["kube_scheduler"]; exist {
				args := schedulerArgs.([]interface{})
				for index := range args {
					advanced.ExtraArgs.KubeScheduler = append(advanced.ExtraArgs.KubeScheduler, args[index].(string))
				}
			}
		}
	}

	// Parse extension_addon
	if v, ok := d.GetOk("extension_addon"); ok {
		addonList := v.([]interface{})
		for _, addon := range addonList {
			addonMap := addon.(map[string]interface{})
			extensionAddon := &tke.ExtensionAddon{
				AddonName: helper.String(addonMap["name"].(string)),
			}
			if addonParam, ok := addonMap["param"]; ok && addonParam.(string) != "" {
				extensionAddon.AddonParam = helper.String(addonParam.(string))
			}
			extensionAddons = append(extensionAddons, extensionAddon)
		}
	}

	// Parse disable_addons
	var disableAddons []*string
	if v, ok := d.GetOk("disable_addons"); ok {
		for _, item := range v.([]interface{}) {
			disableAddons = append(disableAddons, helper.String(item.(string)))
		}
	}

	clusterCIDR := d.Get("cluster_cidr").(string)
	if clusterCIDR != "" {
		cidrSet.ClusterCIDR = helper.String(clusterCIDR)
	}
	if v, ok := d.GetOkExists("ignore_cluster_cidr_conflict"); ok {
		cidrSet.IgnoreClusterCIDRConflict = helper.Bool(v.(bool))
	}
	maxNodePodNum := d.Get("cluster_max_pod_num").(int)
	cidrSet.MaxNodePodNum = helper.Uint64(uint64(maxNodePodNum))
	serviceCIDR := d.Get("service_cidr").(string)
	if serviceCIDR != "" {
		cidrSet.ServiceCIDR = helper.String(serviceCIDR)
	}
	if v, ok := d.GetOkExists("cluster_max_service_num"); ok {
		cidrSet.MaxClusterServiceNum = helper.Uint64(uint64(v.(int)))
	}
	if v, ok := d.GetOkExists("ignore_service_cidr_conflict"); ok {
		cidrSet.IgnoreServiceCIDRConflict = helper.Bool(v.(bool))
	}
	cidrSet.ClaimExpiredSeconds = helper.Int64(int64(d.Get("claim_expired_seconds").(int)))

	if advanced.NetworkType == TKE_CLUSTER_NETWORK_TYPE_VPC_CNI {
		// VPC-CNI cluster need to set eni subnet and service cidr.
		eniSubnetIdList := d.Get("eni_subnet_ids").([]interface{})
		for index := range eniSubnetIdList {
			subnetId := eniSubnetIdList[index].(string)
			cidrSet.EniSubnetIds = append(cidrSet.EniSubnetIds, helper.String(subnetId))
		}
		if serviceCIDR == "" || len(cidrSet.EniSubnetIds) == 0 {
			return fmt.Errorf("`service_cidr` must be set and `eni_subnet_ids` must be set when cluster `network_type` is VPC-CNI.")
		}
	} else {
		// GR cluster
		if clusterCIDR == "" {
			return fmt.Errorf("`cluster_cidr` must be set when cluster `network_type` is GR")
		}
		items := strings.Split(clusterCIDR, "/")
		if len(items) != 2 {
			return fmt.Errorf("`cluster_cidr` must be network segment ")
		}

		bitNumber, err := strconv.ParseInt(items[1], 10, 64)

		if err != nil {
			return fmt.Errorf("`cluster_cidr` must be network segment ")
		}

		if math.Pow(2, float64(32-bitNumber)) <= float64(maxNodePodNum) {
			return fmt.Errorf("`cluster_cidr` Network segment range is too small, can not cover cluster_max_service_num")
		}
	}

	if version, ok := d.GetOk("runtime_version"); ok {
		advanced.RuntimeVersion = version.(string)
	}

	overrideSettings := OverrideSettings{
		Master: make([]tke.InstanceAdvancedSettings, 0),
		Work:   make([]tke.InstanceAdvancedSettings, 0),
	}
	hasInstanceOverrides := false
	masterPreStartUserScript := ""
	masterPreStartUserScriptCount := 0
	masterUserScript := ""
	masterUserScriptCount := 0
	masterConfigCount := 0
	if !runInstancesForNodeOk {
		if masters, ok := d.GetOk("master_config"); ok {
			if clusterDeployType == TKE_DEPLOY_TYPE_MANAGED {
				return fmt.Errorf("if `cluster_deploy_type` is `MANAGED_CLUSTER` , You don't need define the master yourself")
			}
			var masterCount int64 = 0
			masterList := masters.([]interface{})
			masterConfigCount = len(masterList)
			for index := range masterList {
				master := masterList[index].(map[string]interface{})
				masterCvmConfig := master
				if userScript, _ := master["user_script"].(string); userScript != "" {
					if userData, _ := master["user_data"].(string); userData == "" {
						masterCvmConfig = make(map[string]interface{}, len(master)+1)
						for key, value := range master {
							masterCvmConfig[key] = value
						}
						masterCvmConfig["user_data"] = userScript
					}
				}
				paraJson, count, err := tkeGetCvmRunInstancesPara(masterCvmConfig, meta, vpcId, basic.ProjectId)
				if err != nil {
					return err
				}

				cvms.Master = append(cvms.Master, paraJson)
				masterCount += count

				override, hasOverride, err := expandMasterConfigInstanceAdvancedSettings(master)
				if err != nil {
					return err
				}
				overrideSettings.Master = append(overrideSettings.Master, override)
				hasInstanceOverrides = hasInstanceOverrides || hasOverride
				if override.PreStartUserScript != nil {
					masterPreStartUserScriptCount++
					if masterPreStartUserScript == "" {
						masterPreStartUserScript = *override.PreStartUserScript
					} else if masterPreStartUserScript != *override.PreStartUserScript {
						return fmt.Errorf("master_config.pre_start_user_script must be identical for all master nodes during cluster creation")
					}
				}
				if override.UserScript != nil {
					masterUserScriptCount++
					if masterUserScript == "" {
						masterUserScript = *override.UserScript
					} else if masterUserScript != *override.UserScript {
						return fmt.Errorf("master_config.user_script must be identical for all master nodes during cluster creation")
					}
				}
			}
			if masterPreStartUserScriptCount > 0 && masterPreStartUserScriptCount != masterConfigCount {
				return fmt.Errorf("master_config.pre_start_user_script must be configured for all master nodes during cluster creation")
			}
			if masterUserScriptCount > 0 && masterUserScriptCount != masterConfigCount {
				return fmt.Errorf("master_config.user_script must be configured for all master nodes during cluster creation")
			}
			if masterCount < 3 {
				return fmt.Errorf("if `cluster_deploy_type` is `TKE_DEPLOY_TYPE_INDEPENDENT` len(master_config) should >=3")
			}
		} else if clusterDeployType == TKE_DEPLOY_TYPE_INDEPENDENT {
			return fmt.Errorf("if `cluster_deploy_type` is `TKE_DEPLOY_TYPE_INDEPENDENT` , You need define the master yourself")
		}

		if workers, ok := d.GetOk("worker_config"); ok {
			workerList := workers.([]interface{})
			for index := range workerList {
				worker := workerList[index].(map[string]interface{})
				paraJson, _, err := tkeGetCvmRunInstancesPara(worker, meta, vpcId, basic.ProjectId)
				if err != nil {
					return err
				}
				cvms.Work = append(cvms.Work, paraJson)

				override, hasOverride, err := expandDesiredPodNumberOverride(worker)
				if err != nil {
					return err
				}
				overrideSettings.Work = append(overrideSettings.Work, override)
				hasInstanceOverrides = hasInstanceOverrides || hasOverride

				if v, ok := worker["data_disk"]; ok {
					var (
						instanceType = worker["instance_type"].(string)
						zone         = worker["availability_zone"].(string)
					)
					iDiskMountSetting := &tke.InstanceDataDiskMountSetting{
						InstanceType: &instanceType,
						Zone:         &zone,
					}

					diskList := v.([]interface{})
					for _, d := range diskList {
						var (
							disk               = d.(map[string]interface{})
							diskId             = disk["disk_id"].(string)
							diskPartition      = disk["disk_partition"].(string)
							diskType           = disk["disk_type"].(string)
							diskSize           = int64(disk["disk_size"].(int))
							fileSystem         = disk["file_system"].(string)
							autoFormatAndMount = disk["auto_format_and_mount"].(bool)
							mountTarget        = disk["mount_target"].(string)
						)

						dataDisk := &tke.DataDisk{
							DiskType:           &diskType,
							DiskSize:           &diskSize,
							AutoFormatAndMount: &autoFormatAndMount,
						}

						if diskId != "" {
							dataDisk.DiskId = &diskId
						}

						if diskPartition != "" {
							dataDisk.DiskPartition = &diskPartition
						}

						if fileSystem != "" {
							dataDisk.FileSystem = &fileSystem
						}

						if mountTarget != "" {
							dataDisk.MountTarget = &mountTarget
						}

						iDiskMountSetting.DataDisks = append(iDiskMountSetting.DataDisks, dataDisk)
					}

					iDiskMountSettings = append(iDiskMountSettings, iDiskMountSetting)
				}
			}
		}
		if !hasInstanceOverrides {
			overrideSettings.Master = nil
			overrideSettings.Work = nil
		}
	}

	tags := helper.GetTags(d, "tags")

	iAdvanced.Labels = GetTkeLabels(d, "labels")

	if temp, ok := d.GetOk("extra_args"); ok {
		extraArgs := helper.InterfacesStrings(temp.([]interface{}))
		for i := range extraArgs {
			iAdvanced.ExtraArgs.Kubelet = append(iAdvanced.ExtraArgs.Kubelet, &extraArgs[i])
		}
	}
	if temp, ok := d.GetOk("unschedulable"); ok {
		iAdvanced.Unschedulable = int64(temp.(int))
	}
	if temp, ok := d.GetOk("docker_graph_path"); ok {
		iAdvanced.DockerGraphPath = temp.(string)
	}
	if temp, ok := d.GetOk("pre_start_user_script"); ok {
		iAdvanced.PreStartUserScript = temp.(string)
	}
	if masterPreStartUserScript != "" {
		if iAdvanced.PreStartUserScript != "" && iAdvanced.PreStartUserScript != masterPreStartUserScript {
			return fmt.Errorf("master_config.pre_start_user_script conflicts with the cluster-level pre_start_user_script")
		}
		iAdvanced.PreStartUserScript = masterPreStartUserScript
	}
	if temp, ok := d.GetOk("user_script"); ok {
		iAdvanced.UserScript = temp.(string)
	}
	if masterUserScript != "" {
		if iAdvanced.UserScript != "" && iAdvanced.UserScript != masterUserScript {
			return fmt.Errorf("master_config.user_script conflicts with the cluster-level user_script")
		}
		iAdvanced.UserScript = masterUserScript
	}
	if temp, ok := d.GetOk("mount_target"); ok {
		iAdvanced.MountTarget = temp.(string)
	}
	if temp, ok := d.GetOkExists("globe_desired_pod_num"); ok {
		iAdvanced.DesiredPodNum = int64(temp.(int))
		iAdvanced.DesiredPodNumSet = true
	}
	if v, ok := d.GetOk("instance_data_disks"); ok {
		dataDisks, err := expandInstanceDataDisks(v.([]interface{}))
		if err != nil {
			return err
		}
		iAdvanced.DataDisks = dataDisks
	}
	iAdvanced.Taints = GetTkeTaints(d, "taints")

	// ExistedInstancesForNode
	existInstances := make([]*tke.ExistedInstancesForNode, 0)
	if !runInstancesForNodeOk {
		if instances, ok := d.GetOk("exist_instance"); ok {
			instanceList := instances.([]interface{})
			for index := range instanceList {
				instance := instanceList[index].(map[string]interface{})
				existedInstance, _ := tkeGetCvmExistInstancesPara(instance)
				existInstances = append(existInstances, &existedInstance)
			}
		}
	}

	// RunInstancesForNode（master_config+worker_config) 和 ExistedInstancesForNode 不能同时存在
	if len(cvms.Master)+len(cvms.Work) > 0 && len(existInstances) > 0 {
		return fmt.Errorf("master_config+worker_config and exist_instance can not exist at the same time")
	}

	//if v, ok := d.GetOk("extension_addon"); ok {
	//	for _, i := range v.([]interface{}) {
	//		dMap := i.(map[string]interface{})
	//		name := dMap["name"].(string)
	//		param := dMap["param"].(string)
	//		addon := &tke.ExtensionAddon{
	//			AddonName:  helper.String(name),
	//			AddonParam: helper.String(param),
	//		}
	//		extensionAddons = append(extensionAddons, addon)
	//	}
	//}

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	id, err := service.CreateCluster(ctx, basic, advanced, cvms, runInstancesForNode, clusterDeployType, iAdvanced, cidrSet, tags, existInstances,
		&overrideSettings, iDiskMountSettings, clusterArch, extensionAddons, "", disableAddons)
	if err != nil {
		return err
	}

	d.SetId(id)

	_, _, err = service.DescribeClusterInstances(ctx, d.Id())

	if err != nil {
		// create often cost more than 20 Minutes.
		err = resource.Retry(10*readRetryTimeout, func() *resource.RetryError {
			_, _, err = service.DescribeClusterInstances(ctx, d.Id())

			if e, ok := err.(*errors.CloudSDKError); ok {
				if e.GetCode() == "InternalError.ClusterNotFound" {
					return nil
				}
			}

			if err != nil {
				return resource.RetryableError(err)
			}
			return nil
		})
	}

	if err != nil {
		return err
	}

	// err = service.CheckOneOfClusterNodeReady(ctx, d.Id(), nodePoolId, clusterInternet || clusterIntranet)

	// if err != nil {
	// 	return err
	// }

	//intranet
	//if clusterIntranet {
	//	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
	//		inErr := service.CreateClusterEndpoint(ctx, id, intranetSubnetId, clusterInternetSecurityGroup, false, clusterIntranetDomain, "")
	//		if inErr != nil {
	//			return retryError(inErr)
	//		}
	//		return nil
	//	})
	//	if err != nil {
	//		return err
	//	}
	//	err = resource.Retry(2*readRetryTimeout, func() *resource.RetryError {
	//		status, message, inErr := service.DescribeClusterEndpointStatus(ctx, id, false)
	//		if inErr != nil {
	//			return retryError(inErr)
	//		}
	//		if status == TkeInternetStatusCreating {
	//			return resource.RetryableError(
	//				fmt.Errorf("%s create intranet cluster endpoint status still is %s", id, status))
	//		}
	//		if status == TkeInternetStatusNotfound || status == TkeInternetStatusCreated {
	//			return nil
	//		}
	//		return resource.NonRetryableError(
	//			fmt.Errorf("%s create intranet cluster endpoint error ,status is %s,message is %s", id, status, message))
	//	})
	//	if err != nil {
	//		return err
	//	}
	//}

	//if clusterInternet {
	//	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
	//		inErr := service.CreateClusterEndpoint(ctx, id, "", clusterInternetSecurityGroup, true, clusterInternetDomain, "")
	//		if inErr != nil {
	//			return retryError(inErr)
	//		}
	//		return nil
	//	})
	//	if err != nil {
	//		return err
	//	}
	//	err = resource.Retry(2*readRetryTimeout, func() *resource.RetryError {
	//		status, message, inErr := service.DescribeClusterEndpointStatus(ctx, id, true)
	//		if inErr != nil {
	//			return retryError(inErr)
	//		}
	//		if status == TkeInternetStatusCreating {
	//			return resource.RetryableError(
	//				fmt.Errorf("%s create cluster internet endpoint status still is %s", id, status))
	//		}
	//		if status == TkeInternetStatusNotfound || status == TkeInternetStatusCreated {
	//			return nil
	//		}
	//		return resource.NonRetryableError(
	//			fmt.Errorf("%s create cluster internet endpoint error ,status is %s,message is %s", id, status, message))
	//	})
	//	if err != nil {
	//		return err
	//	}
	//}

	err = resource.Retry(20*readRetryTimeout, func() *resource.RetryError {
		info, has, inErr := service.DescribeCluster(ctx, d.Id())
		if inErr != nil {
			return retryError(inErr)
		}

		if info.ClusterStatus == "Creating" || info.ClusterStatus == "Initializing" {
			return resource.RetryableError(
				fmt.Errorf("%s create cluster internet endpoint status still is %s", id, info.ClusterStatus))
		}

		if info.ClusterStatus == "Abnormal" {
			return retryError(inErr, "create cluster %s failed", id)
		}

		if info.ClusterStatus == "Running" {
			return nil
		}

		return resource.NonRetryableError(
			fmt.Errorf("%s create cluster error ,status is %s,message is %v", id, info.ClusterStatus, has))
	})

	if err != nil {
		return err
	}

	if clusterDeployType == TKE_DEPLOY_TYPE_INDEPENDENT {
		masterConfigs, _ := d.Get("master_config").([]interface{})
		expectedMasterCount := len(masterConfigs)
		err = resource.Retry(10*readRetryTimeout, func() *resource.RetryError {
			masters, _, inErr := service.DescribeClusterInstances(ctx, id)
			if inErr != nil {
				return retryError(inErr)
			}
			return masterInventoryConvergenceError(masters, expectedMasterCount, nil, masterInventoryGrowing, "master creation")
		})
		if err != nil {
			return fmt.Errorf("cluster was created but its master inventory did not converge to %d healthy node(s): %w. Terraform will not delete or recreate control-plane nodes automatically; wait for TKE to stabilize and run Terraform again", expectedMasterCount, err)
		}
	}

	if v, ok := helper.InterfacesHeadMap(d, "log_agent"); ok {
		enabled := v["enabled"].(bool)
		rootDir := v["kubelet_root_dir"].(string)

		if enabled {
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				err := service.SwitchLogAgent(ctx, id, rootDir, enabled)
				if err != nil {
					return retryError(err, "FailedOperation.ClusterNotFound")
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	if v, ok := helper.InterfacesHeadMap(d, "event_persistence"); ok {
		enabled := v["enabled"].(bool)
		logSetId := v["log_set_id"].(string)
		topicId := v["topic_id"].(string)
		if enabled {
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				err := service.SwitchEventPersistence(ctx, id, logSetId, topicId, enabled, false)
				if err != nil {
					return retryError(err, "FailedOperation.ClusterNotFound")
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	if v, ok := helper.InterfacesHeadMap(d, "cluster_audit"); ok {
		enabled := v["enabled"].(bool)
		logSetId := v["log_set_id"].(string)
		topicId := v["topic_id"].(string)
		if enabled {
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				err := service.SwitchClusterAudit(ctx, id, logSetId, topicId, enabled, false)
				if err != nil {
					return retryError(err, "FailedOperation.ClusterNotFound")
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	if err = resourceTencentCloudTkeClusterRead(d, meta); err != nil {
		log.Printf("[WARN]%s resource.kubernetes_cluster.read after create fail , %s", logId, err.Error())
		return err
	}
	return nil
}

func resourceTencentCloudTkeClusterRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_cluster.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	info, has, err := service.DescribeCluster(ctx, d.Id())
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			info, has, err = service.DescribeCluster(ctx, d.Id())
			if err != nil {
				return retryError(err)
			}
			return nil
		})
	}

	if err != nil {
		return nil
	}

	if !has {
		d.SetId("")
		return nil
	}

	// 兼容旧的 cluster_os 的 key, 由于 cluster_os有默认值，所以不大可能为空
	oldOs := d.Get("cluster_os").(string)
	newOs := tkeToShowClusterOs(info.ClusterOs)

	if (oldOs == TkeClusterOsCentOS76 && newOs == TKE_CLUSTER_OS_CENTOS76) ||
		(oldOs == TkeClusterOsUbuntu18 && newOs == TKE_CLUSTER_OS_UBUNTU18) {
		newOs = oldOs
	}

	_ = d.Set("cluster_name", info.ClusterName)
	_ = d.Set("cluster_desc", info.ClusterDescription)
	_ = d.Set("cluster_os", newOs)
	_ = d.Set("cluster_deploy_type", info.DeployType)
	_ = d.Set("cluster_version", info.ClusterVersion)
	_ = d.Set("cluster_ipvs", info.Ipvs)
	_ = d.Set("network_type", info.NetworkType)
	_ = d.Set("vpc_id", info.VpcId)
	_ = d.Set("project_id", info.ProjectId)
	_ = d.Set("cluster_cidr", info.ClusterCidr)
	_ = d.Set("ignore_cluster_cidr_conflict", info.IgnoreClusterCidrConflict)
	_ = d.Set("cluster_max_pod_num", info.MaxNodePodNum)
	_ = d.Set("cluster_max_service_num", info.MaxClusterServiceNum)
	_ = d.Set("cluster_node_num", info.ClusterNodeNum)
	_ = d.Set("tags", info.Tags)

	if _, ok := d.GetOk("cluster_level"); ok && info.ClusterLevel != nil {
		_ = d.Set("cluster_level", info.ClusterLevel)
	}

	if _, ok := d.GetOkExists("auto_upgrade_cluster_level"); ok && info.AutoUpgradeClusterLevel != nil {
		_ = d.Set("auto_upgrade_cluster_level", *info.AutoUpgradeClusterLevel)
	}

	config, err := service.DescribeClusterConfig(ctx, d.Id(), true)
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			config, err = service.DescribeClusterConfig(ctx, d.Id(), true)
			if err != nil {
				return retryError(err)
			}
			return nil
		})
	}

	if err != nil {
		return nil
	}

	_ = d.Set("kube_config", config)

	intranetConfig, err := service.DescribeClusterConfig(ctx, d.Id(), false)
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			intranetConfig, err = service.DescribeClusterConfig(ctx, d.Id(), false)
			if err != nil {
				return retryError(err)
			}
			return nil
		})
	}

	if err != nil {
		return nil
	}

	_ = d.Set("kube_config_intranet", intranetConfig)

	masters, workers, err := service.DescribeClusterInstances(ctx, d.Id())
	if err != nil {
		err = resource.Retry(10*readRetryTimeout, func() *resource.RetryError {
			masters, workers, err = service.DescribeClusterInstances(ctx, d.Id())

			if e, ok := err.(*errors.CloudSDKError); ok {
				if e.GetCode() == "InternalError.ClusterNotFound" {
					return nil
				}
			}
			if err != nil {
				return resource.RetryableError(err)
			}
			return nil
		})
	}
	if err != nil {
		return err
	}

	workerInstancesList := make([]map[string]interface{}, 0, len(workers))
	for _, worker := range workers {
		tempMap := make(map[string]interface{})
		tempMap["instance_id"] = worker.InstanceId
		tempMap["instance_role"] = worker.InstanceRole
		tempMap["instance_state"] = worker.InstanceState
		tempMap["failed_reason"] = worker.FailedReason
		tempMap["lan_ip"] = worker.LanIp
		workerInstancesList = append(workerInstancesList, tempMap)
	}

	_ = d.Set("worker_instances_list", workerInstancesList)

	var masterHealthErr error
	// 回读 master_config（仅 INDEPENDENT_CLUSTER）
	if info.DeployType == TKE_DEPLOY_TYPE_INDEPENDENT {
		masterInstanceIds := make([]*string, 0, len(masters))
		for _, m := range masters {
			id := m.InstanceId
			masterInstanceIds = append(masterInstanceIds, &id)
		}
		// An independent cluster cannot legitimately have zero masters. During
		// creation the TKE instance list may briefly be empty, so keep the
		// configured blocks until TKE exposes the authoritative instance set.
		if len(masterInstanceIds) > 0 {
			masterConfigs, _ := d.Get("master_config").([]interface{})
			cvmService := CvmService{client: meta.(*TencentCloudClient).apiV3Conn}
			var cvmInstances []*cvm.Instance
			err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
				var inErr error
				cvmInstances, inErr = cvmService.DescribeInstanceByFilter(ctx, masterInstanceIds, nil)
				if inErr != nil {
					return retryError(inErr)
				}
				if len(cvmInstances) != len(masterInstanceIds) {
					return resource.RetryableError(fmt.Errorf("master_config: TKE returned %d masters but CVM returned %d instances", len(masterInstanceIds), len(cvmInstances)))
				}
				return nil
			})
			if err != nil {
				// Keep known configuration during transient CVM read/count mismatch.
				// Import has no usable master state to preserve, so it must fail.
				if len(masterConfigs) == 0 {
					return err
				}
				log.Printf("[WARN] master_config: keeping existing state because CVM instances could not be read completely: %s", err.Error())
			} else {
				masterList := reconcileMasterConfigState(masterConfigs, cvmInstances, masters)
				if schemaErr := validateMasterConfigStateSchema(masterList); schemaErr != nil {
					return schemaErr
				}
				if setErr := d.Set("master_config", masterList); setErr != nil {
					return setErr
				}
			}
		}
		if len(masters) == 0 {
			masterHealthErr = fmt.Errorf("TKE returned no master nodes for an independent cluster. Terraform kept the existing master_config state and will not submit a scaling request; wait for TKE inventory to recover or investigate the cluster before running Terraform again")
		} else if info.ClusterStatus != "" && !strings.EqualFold(info.ClusterStatus, "Running") {
			masterHealthErr = fmt.Errorf("TKE cluster is currently in %q. Terraform will not submit another master scaling request while the previous operation may still be converging; wait for the cluster and all masters to become healthy, then run Terraform again", info.ClusterStatus)
		} else {
			masterHealthErr = masterInstancesNotReadyError(masters)
		}
	}

	securityRet, err := service.DescribeClusterSecurity(ctx, d.Id())

	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			securityRet, err = service.DescribeClusterSecurity(ctx, d.Id())
			if e, ok := err.(*errors.CloudSDKError); ok {
				if e.GetCode() == "InternalError.ClusterNotFound" {
					return nil
				}
			}
			if err != nil {
				return resource.RetryableError(err)
			}
			return nil
		})
	}
	if err != nil {
		return err
	}
	var emptyStrFunc = func(ptr *string) string {
		if ptr == nil {
			return ""
		} else {
			return *ptr
		}
	}

	policies := make([]string, 0, len(securityRet.Response.SecurityPolicy))
	for _, v := range securityRet.Response.SecurityPolicy {
		policies = append(policies, *v)
	}

	_ = d.Set("user_name", emptyStrFunc(securityRet.Response.UserName))
	_ = d.Set("password", emptyStrFunc(securityRet.Response.Password))
	_ = d.Set("certification_authority", emptyStrFunc(securityRet.Response.CertificationAuthority))
	_ = d.Set("cluster_external_endpoint", emptyStrFunc(securityRet.Response.ClusterExternalEndpoint))
	_ = d.Set("domain", emptyStrFunc(securityRet.Response.Domain))
	_ = d.Set("pgw_endpoint", emptyStrFunc(securityRet.Response.PgwEndpoint))
	_ = d.Set("security_policy", policies)

	//if v, ok := d.GetOk("worker_config"); ok && len(v.([]interface{})) > 0 {
	//	if emptyStrFunc(securityRet.Response.ClusterExternalEndpoint) == "" {
	//		_ = d.Set("cluster_internet", false)
	//	} else {
	//		_ = d.Set("cluster_internet", true)
	//	}
	//
	//	if emptyStrFunc(securityRet.Response.PgwEndpoint) == "" {
	//		_ = d.Set("cluster_intranet", false)
	//	} else {
	//		_ = d.Set("cluster_intranet", true)
	//	}
	//}

	var globalConfig *tke.ClusterAsGroupOption
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		globalConfig, err = service.DescribeClusterNodePoolGlobalConfig(ctx, d.Id())
		if e, ok := err.(*errors.CloudSDKError); ok {
			if e.GetCode() == "InternalError.ClusterNotFound" {
				return nil
			}
		}
		if err != nil {
			return resource.RetryableError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	if globalConfig != nil {
		temp := make(map[string]interface{})
		temp["is_scale_in_enabled"] = globalConfig.IsScaleDownEnabled
		temp["expander"] = globalConfig.Expander
		temp["max_concurrent_scale_in"] = globalConfig.MaxEmptyBulkDelete
		temp["scale_in_delay"] = globalConfig.ScaleDownDelay
		temp["scale_in_unneeded_time"] = globalConfig.ScaleDownUnneededTime
		temp["scale_in_utilization_threshold"] = globalConfig.ScaleDownUtilizationThreshold
		temp["ignore_daemon_sets_utilization"] = globalConfig.IgnoreDaemonSetsUtilization
		temp["skip_nodes_with_local_storage"] = globalConfig.SkipNodesWithLocalStorage
		temp["skip_nodes_with_system_pods"] = globalConfig.SkipNodesWithSystemPods

		//_ = d.Set("node_pool_global_config", []map[string]interface{}{temp})
	}
	return masterHealthErr
}

func resourceTencentCloudTkeClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_cluster.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()

	client := meta.(*TencentCloudClient).apiV3Conn
	service := TagService{client: client}
	tkeService := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	region := client.Region
	d.Partial(true)

	if d.HasChange("tags") {
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))

		resourceName := BuildTagResourceName("ccs", "cluster", region, id)
		if err := service.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}

	}

	if d.HasChange("project_id") || d.HasChange("cluster_name") || d.HasChange("cluster_desc") || d.HasChange("cluster_level") || d.HasChange("auto_upgrade_cluster_level") {
		projectId := int64(d.Get("project_id").(int))
		clusterName := d.Get("cluster_name").(string)
		clusterDesc := d.Get("cluster_desc").(string)
		clusterLevel := d.Get("cluster_level").(string)
		autoUpgradeClusterLevel := d.Get("auto_upgrade_cluster_level").(bool)

		ins, _, err := tkeService.DescribeCluster(ctx, id)
		if err != nil {
			return err
		}

		//ignore same cluster level if same
		if ins.ClusterLevel != nil && *ins.ClusterLevel == clusterLevel {
			clusterLevel = ""
		}

		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			err := tkeService.ModifyClusterAttribute(ctx, id, projectId, clusterName, clusterDesc, clusterLevel, autoUpgradeClusterLevel)
			if err != nil {
				// create and update immediately may cause cluster level syntax error, this error can wait until cluster level state normal
				return retryError(err, tke.INTERNALERROR_UNEXPECTEDINTERNAL, tke.RESOURCEUNAVAILABLE)
			}
			return nil
		})

		if err != nil {
			return err
		}
	}

	//upgrade k8s cluster version
	if d.HasChange("cluster_version") {
		newVersion := d.Get("cluster_version").(string)
		isOk, err := tkeService.CheckClusterVersion(ctx, id, newVersion)
		if err != nil {
			return err
		}
		if !isOk {
			return fmt.Errorf("version %s is unsupported", newVersion)
		}
		extraArgs, ok := d.GetOk("cluster_extra_args")
		if !ok {
			extraArgs = nil
		}
		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			inErr := tkeService.ModifyClusterVersion(ctx, id, newVersion, extraArgs)
			if inErr != nil {
				return retryError(inErr)
			}
			return nil
		})
		if err != nil {
			return err
		}
		//check status
		err = resource.Retry(3*readRetryTimeout, func() *resource.RetryError {
			ins, has, inErr := tkeService.DescribeCluster(ctx, id)
			if inErr != nil {
				return retryError(inErr)
			}
			if !has {
				return resource.NonRetryableError(fmt.Errorf("Cluster %s is not exist", id))
			}
			if ins.ClusterStatus == "Running" {
				return nil
			} else {
				return resource.RetryableError(fmt.Errorf("cluster %s status %s, retry...", id, ins.ClusterStatus))
			}
		})
		if err != nil {
			return err
		}

		// upgrade instances version
		upgrade := false
		if v, ok := d.GetOk("upgrade_instances_follow_cluster"); ok {
			upgrade = v.(bool)
		}
		if upgrade {
			err := upgradeClusterInstances(tkeService, ctx, id)
			if err != nil {
				return err
			}
		}
	}

	// update node pool global config
	if d.HasChange("node_pool_global_config") {
		request := tkeGetNodePoolGlobalConfig(d)
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			inErr := tkeService.ModifyClusterNodePoolGlobalConfig(ctx, request)
			if inErr != nil {
				return retryError(inErr)
			}
			return nil
		})
		if err != nil {
			return err
		}

	}

	//if d.HasChange("deletion_protection") {
	//	enable := d.Get("deletion_protection").(bool)
	//	if err := tkeService.ModifyDeletionProtection(ctx, id, enable); err != nil {
	//		return err
	//	}
	//
	//}

	if d.HasChange("acquire_cluster_admin_role") {
		o, n := d.GetChange("acquire_cluster_admin_role")
		if o.(bool) && !n.(bool) {
			return fmt.Errorf("argument `acquire_cluster_admin_role` cannot set to false")
		}
		err := tkeService.AcquireClusterAdminRole(ctx, id)
		if err != nil {
			return err
		}
	}

	if d.HasChange("log_agent") {
		v, ok := helper.InterfacesHeadMap(d, "log_agent")
		enabled := false
		rootDir := ""
		if ok {
			rootDir = v["kubelet_root_dir"].(string)
			enabled = v["enabled"].(bool)
		}
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			err := tkeService.SwitchLogAgent(ctx, id, rootDir, enabled)
			if err != nil {
				return retryError(err, "FailedOperation.ClusterNotFound")
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	if d.HasChange("event_persistence") {
		v, ok := helper.InterfacesHeadMap(d, "event_persistence")
		enabled := false
		logSetId := ""
		topicId := ""
		deleteEventLog := false
		if ok {
			enabled = v["enabled"].(bool)
			logSetId = v["log_set_id"].(string)
			topicId = v["topic_id"].(string)
			deleteEventLog = v["delete_event_log_and_topic"].(bool)
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			err := tkeService.SwitchEventPersistence(ctx, id, logSetId, topicId, enabled, deleteEventLog)
			if err != nil {
				return retryError(err, "FailedOperation.ClusterNotFound")
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	if d.HasChange("cluster_audit") {
		v, ok := helper.InterfacesHeadMap(d, "cluster_audit")
		enabled := false
		logSetId := ""
		topicId := ""
		deleteAuditLog := false
		if ok {
			enabled = v["enabled"].(bool)
			logSetId = v["log_set_id"].(string)
			topicId = v["topic_id"].(string)
			deleteAuditLog = v["delete_audit_log_and_topic"].(bool)
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			err := tkeService.SwitchClusterAudit(ctx, id, logSetId, topicId, enabled, deleteAuditLog)
			if err != nil {
				return retryError(err, "FailedOperation.ClusterNotFound")
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	//if d.HasChange("extension_addon") {
	//	o, n := d.GetChange("extension_addon")
	//	adds, removes, changes := resourceTkeGetAddonsDiffs(o.([]interface{}), n.([]interface{}))
	//	updates := append(adds, changes...)
	//	for i := range updates {
	//		var err error
	//		addon := updates[i].(map[string]interface{})
	//		param := addon["param"].(string)
	//		name, err := tkeService.GetAddonNameFromJson(param)
	//		if err != nil {
	//			return err
	//		}
	//		_, has, _ := tkeService.PollingAddonsPhase(ctx, id, name, nil)
	//		if has {
	//			err = tkeService.UpdateExtensionAddon(ctx, id, name, param)
	//		} else {
	//			err = tkeService.CreateExtensionAddon(ctx, id, param)
	//		}
	//		if err != nil {
	//			return err
	//		}
	//		_, _, err = tkeService.PollingAddonsPhase(ctx, id, name, nil)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//
	//	for i := range removes {
	//		addon := removes[i].(map[string]interface{})
	//		param := addon["param"].(string)
	//		name, err := tkeService.GetAddonNameFromJson(param)
	//		if err != nil {
	//			return err
	//		}
	//		_, has, _ := tkeService.PollingAddonsPhase(ctx, id, name, nil)
	//		if !has {
	//			continue
	//		}
	//		err = tkeService.DeleteExtensionAddon(ctx, id, name)
	//		if err != nil {
	//			return err
	//		}
	//		_, has, _ = tkeService.PollingAddonsPhase(ctx, id, name, nil)
	//		if has {
	//			return fmt.Errorf("addon %s still exists", name)
	//		}
	//	}
	//
	//}

	if d.HasChange("master_config") {
		oldM, newM := d.GetChange("master_config")
		oldList := oldM.([]interface{})
		newList := newM.([]interface{})

		finalCount := len(newList)
		vpcId := d.Get("vpc_id").(string)
		projectId := int64(d.Get("project_id").(int))

		change, err := planMasterConfigInventoryChange(oldList, newList)
		if err != nil {
			return err
		}
		oldMap := change.oldMap
		newMap := change.newMap
		addedBlocks := change.addedBlocks
		removedBlocks := change.removedBlocks

		// 1. Scale Out (Addition)
		if len(addedBlocks) > 0 {
			requiredInstanceIds, err := masterConfigInstanceIDs(oldList)
			if err != nil {
				return err
			}
			alreadyAtTarget, err := readMasterScalingPreflight(ctx, tkeService, id, len(oldList), finalCount, requiredInstanceIds, "master scale-out")
			if err != nil {
				return err
			}

			if !alreadyAtTarget {
				runInstancesForNodeList, err := buildScaleOutMasterRequest(addedBlocks, meta, vpcId, projectId)
				if err != nil {
					return err
				}

				// ScaleOutClusterMaster has no idempotency token. Do not automatically
				// retry an ambiguous response, because the first request may already
				// have been accepted and a retry could create duplicate masters.
				if err := tkeService.ScaleOutClusterMaster(ctx, id, runInstancesForNodeList); err != nil {
					return fmt.Errorf("master scale-out request returned an error and its remote outcome is unknown: %w. Terraform did not retry the request; refresh TKE state before running Terraform again", err)
				}
			}

			// ScaleOutClusterMaster only acknowledges an asynchronous operation.
			// The remote cluster may now contain some or all requested masters. Keep
			// full-state mode; a later Read reconciles remote IDs by stable identity
			// even when Terraform Core cannot persist state after an update error.
			d.Partial(false)

			// Poll until the complete desired master inventory is visible and every
			// node is running. Any non-running state may be an intermediate service
			// state, so neither the state name nor its reason is treated as terminal
			// while this request is converging. A persistent condition returns the
			// last detailed status when the bounded wait expires.
			err = resource.Retry(10*readRetryTimeout, func() *resource.RetryError {
				masters, _, inErr := tkeService.DescribeClusterInstances(ctx, id)
				if inErr != nil {
					return retryError(inErr)
				}
				return masterInventoryConvergenceError(masters, finalCount, requiredInstanceIds, masterInventoryGrowing, "master scale-out")
			})
			if err != nil {
				return fmt.Errorf("master scale-out did not converge to %d healthy node(s): %w. Terraform will not automatically delete/recreate masters or submit another scaling request; wait for TKE to become healthy, then run Terraform again so state is reconciled first", finalCount, err)
			}

			// After nodes reach running, wait for the cluster to exit MasterScaling so
			// the control plane and etcd finish the async member-join flow before we return.
			// ScaleOutClusterMaster only acknowledges the request (no TaskId); a node being
			// "running" does not guarantee the cluster has left the MasterScaling state.
			err = resource.Retry(20*readRetryTimeout, func() *resource.RetryError {
				status, inErr := tkeService.DescribeClusterStatus(ctx, id)
				if inErr != nil {
					return retryError(inErr)
				}
				if status == nil || status.ClusterState == nil {
					return resource.RetryableError(fmt.Errorf("cluster %s has not reported a state after master scale-out", id))
				}
				if !strings.EqualFold(*status.ClusterState, "Running") {
					return resource.RetryableError(fmt.Errorf("cluster %s still in %s after master scale-out, waiting for control plane to stabilize", id, *status.ClusterState))
				}
				return nil
			})
			if err != nil {
				return fmt.Errorf("all requested masters became healthy, but the cluster did not return to Running after scale-out: %w. Terraform will not submit another scaling request; wait for TKE control-plane convergence and run Terraform again", err)
			}
		}

		// 2. Scale In (Deletion)
		if len(removedBlocks) > 0 {
			// Computed values in a TypeList plan can remain positional after HCL
			// removes a block. Resolve retained IDs from the old state by stable
			// instance_name instead of trusting instance_id values in newList.
			retainedStateBlocks := make([]interface{}, 0, len(newMap))
			for name := range newMap {
				if oldBlock, exists := oldMap[name]; exists {
					retainedStateBlocks = append(retainedStateBlocks, oldBlock)
				}
			}
			requiredInstanceIds, err := masterConfigInstanceIDs(retainedStateBlocks)
			if err != nil {
				return err
			}
			alreadyAtTarget, err := readMasterScalingPreflight(ctx, tkeService, id, len(oldList), finalCount, requiredInstanceIds, "master scale-in")
			if err != nil {
				return err
			}

			scaleInMasters, err := buildScaleInMasterRequest(removedBlocks)
			if err != nil {
				return err
			}

			if !alreadyAtTarget {
				// As with scale-out, do not retry a mutating request whose remote
				// outcome may be ambiguous.
				if err := tkeService.ScaleInClusterMaster(ctx, id, scaleInMasters); err != nil {
					return fmt.Errorf("master scale-in request returned an error and its remote outcome is unknown: %w. Terraform did not retry the request; refresh TKE state before running Terraform again", err)
				}
			}

			// As with scale-out, the API acknowledgement means the remote state may
			// already have changed even if the subsequent convergence check fails.
			d.Partial(false)

			// Wait for the exact retained inventory and for every retained master
			// to be healthy. A removed node may temporarily report failed while TKE
			// drains it, so presence is treated as pending rather than terminal.
			err = resource.Retry(10*readRetryTimeout, func() *resource.RetryError {
				masters, _, inErr := tkeService.DescribeClusterInstances(ctx, id)
				if inErr != nil {
					return retryError(inErr)
				}
				return masterInventoryConvergenceError(masters, finalCount, requiredInstanceIds, masterInventoryShrinking, "master scale-in")
			})
			if err != nil {
				return fmt.Errorf("master scale-in did not converge to %d healthy node(s): %w. Terraform will not automatically retry the mutation or change the delete mode; wait for TKE to stabilize, then run Terraform again so state is reconciled first", finalCount, err)
			}

			// After nodes are removed from the instance list, the cluster stays in
			// MasterScaling while the backend finishes etcd member removal and resource
			// cleanup. ScaleInClusterMaster only acknowledges the request (no TaskId);
			// a node disappearing from DescribeClusterInstances does not mean the
			// async scale-in is done. Wait for the cluster to return to Running so
			// subsequent operations (refresh, plan, further scaling) hit a stable cluster.
			err = resource.Retry(20*readRetryTimeout, func() *resource.RetryError {
				status, inErr := tkeService.DescribeClusterStatus(ctx, id)
				if inErr != nil {
					return retryError(inErr)
				}
				if status == nil || status.ClusterState == nil {
					return resource.RetryableError(fmt.Errorf("cluster %s has not reported a state after master scale-in", id))
				}
				if !strings.EqualFold(*status.ClusterState, "Running") {
					return resource.RetryableError(fmt.Errorf("cluster %s still in %s after master scale-in, waiting for etcd metadata convergence", id, *status.ClusterState))
				}
				return nil
			})
			if err != nil {
				return fmt.Errorf("the retained masters became healthy, but the cluster did not return to Running after scale-in: %w. Terraform will not retry the mutation; wait for TKE control-plane convergence and run Terraform again", err)
			}
		}

		if len(addedBlocks) > 0 || len(removedBlocks) > 0 {
			if err := d.Set("master_config", masterConfigStateForInventoryChange(newList, oldMap)); err != nil {
				return fmt.Errorf("reconciling master_config state after scaling: %w", err)
			}
		}
	}

	d.Partial(false)
	if err := resourceTencentCloudTkeClusterRead(d, meta); err != nil {
		log.Printf("[WARN]%s resource.kubernetes_cluster.read after update fail , %s", logId, err.Error())
	}

	return nil
}

func resourceTencentCloudTkeClusterDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_cluster.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	deleteEventLogSetAndTopic := false
	enableEventLog := false
	deleteAuditLogSetAndTopic := false
	if v, ok := helper.InterfacesHeadMap(d, "event_persistence"); ok {
		deleteEventLogSetAndTopic = v["delete_event_log_and_topic"].(bool)
		// get cluster current enabled status
		enableEventLog = v["enabled"].(bool)
	}

	if v, ok := helper.InterfacesHeadMap(d, "cluster_audit"); ok {
		deleteAuditLogSetAndTopic = v["delete_audit_log_and_topic"].(bool)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if deleteEventLogSetAndTopic && enableEventLog {
			err := service.SwitchEventPersistence(ctx, d.Id(), "", "", false, true)
			if e, ok := err.(*errors.CloudSDKError); ok {
				if e.GetCode() != "ResourceNotFound.ClusterNotFound" {
					return retryError(err, InternalError)
				}
			} else if err != nil {
				return retryError(err, InternalError)
			}
		}
		if deleteAuditLogSetAndTopic {
			err := service.SwitchClusterAudit(ctx, d.Id(), "", "", false, true)
			if e, ok := err.(*errors.CloudSDKError); ok {
				if e.GetCode() != "ResourceNotFound.ClusterNotFound" {
					return retryError(err, InternalError)
				}
			} else if err != nil {
				return retryError(err, InternalError)
			}
		}
		err := service.DeleteCluster(ctx, d.Id(), "", nil)

		if e, ok := err.(*errors.CloudSDKError); ok {
			if e.GetCode() == "InternalError.ClusterNotFound" {
				return nil
			}
		}

		if err != nil {
			return retryError(err, InternalError)
		}
		return nil
	})

	if err != nil {
		return err
	}
	_, _, err = service.DescribeClusterInstances(ctx, d.Id())

	if err != nil {
		err = resource.Retry(10*readRetryTimeout, func() *resource.RetryError {
			_, _, err = service.DescribeClusterInstances(ctx, d.Id())
			if e, ok := err.(*errors.CloudSDKError); ok {
				if e.GetCode() == "InvalidParameter.ClusterNotFound" {
					return nil
				}
			}
			if err != nil {
				return retryError(err, InternalError)
			}
			return nil
		})
	}
	return err

}

func resourceTkeGetAddonsDiffs(o, n []interface{}) (adds, removes, changes []interface{}) {
	indexByName := func(i interface{}) int {
		v := i.(map[string]interface{})
		return helper.HashString(v["name"].(string))
	}
	indexAll := func(i interface{}) int {
		v := i.(map[string]interface{})
		name := v["name"].(string)
		param := v["param"].(string)
		return helper.HashString(fmt.Sprintf("%s#%s", name, param))
	}

	os := schema.NewSet(indexByName, o)
	ns := schema.NewSet(indexByName, n)

	adds = ns.Difference(os).List()
	removes = os.Difference(ns).List()

	fullIndexedKeeps := schema.NewSet(indexAll, ns.Intersection(os).List())
	fullIndexedOlds := schema.NewSet(indexAll, o)

	changes = fullIndexedKeeps.Difference(fullIndexedOlds).List()
	return
}
