/*
Use this data source to query detailed information of kubernetes clusters.

# Example Usage

```hcl

	data "tencentcloudenterprise_tke_kubernetes_clusters" "name" {
	  cluster_name = "terraform"
	}

	data "tencentcloudenterprise_tke_kubernetes_clusters" "id" {
	  cluster_id = "cls-godovr32"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func tkeClusterInfo() map[string]*schema.Schema {
	schemaBody := map[string]*schema.Schema{
		"cluster_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of cluster.",
		},
		"cluster_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the cluster.",
		},
		"cluster_desc": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Description of the cluster.",
		},
		"cluster_os": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Operating system of the cluster.",
		},
		"container_runtime": {
			Type:        schema.TypeString,
			Deprecated:  "It has been deprecated from version 1.18.1.",
			Computed:    true,
			Description: "Container runtime of the cluster.",
		},
		"cluster_deploy_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Deployment type of the cluster.",
		},
		"cluster_version": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Version of the cluster.",
		},
		"cluster_ipvs": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicates whether ipvs is enabled.",
		},
		"vpc_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Vpc ID of the cluster.",
		},
		"project_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Project ID of the cluster.",
		},
		"cluster_cidr": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this VPC.",
		},
		"ignore_cluster_cidr_conflict": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicates whether to ignore the cluster cidr conflict error.",
		},
		"cluster_max_pod_num": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The maximum number of Pods per node in the cluster.",
		},
		"cluster_max_service_num": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The maximum number of services in the cluster.",
		},
		"cluster_as_enabled": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicates whether to enable cluster node auto scaler.",
		},
		"node_name_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Node name type of cluster.",
		},
		"cluster_extra_args": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"kube_apiserver": {
						Type:        schema.TypeList,
						Computed:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "The customized parameters for kube-apiserver.",
					},
					"kube_controller_manager": {
						Type:        schema.TypeList,
						Computed:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "The customized parameters for kube-controller-manager.",
					},
					"kube_scheduler": {
						Type:        schema.TypeList,
						Computed:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "The customized parameters for kube-scheduler.",
					},
				},
			},
			Description: "Customized parameters for master component.",
		},
		"network_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Cluster network type.",
		},
		"is_non_static_ip_mode": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicates whether non-static ip mode is enabled.",
		},
		"kube_proxy_mode": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Cluster kube-proxy mode.",
		},

		"service_cidr": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The network address block of the cluster.",
		},
		"eni_subnet_ids": {
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Subnet IDs for cluster with VPC-CNI network mode.",
		},
		"claim_expired_seconds": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The expired seconds to recycle ENI.",
		},
		"deletion_protection": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Indicates whether cluster deletion protection is enabled.",
		},
		"cluster_node_num": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Number of nodes in the cluster.",
		},
		"worker_instances_list": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "An information list of cvm within the WORKER clusters. Each element contains the following attributes.",
			Elem: &schema.Resource{
				Schema: tkeCvmState(),
			},
		},
		"tags": {
			Type:        schema.TypeMap,
			Computed:    true,
			Description: "Tags of the cluster.",
		},
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

	return schemaBody
}

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tke_kubernetes_clusters", CNDescription{
		TerraformTypeCN: "tke集群",
		DescriptionCN:   "提供TKE Kubernetes集群数据源，用于查询TKE Kubernetes集群的详细信息。",
		AttributesCN: map[string]string{
			"cluster_id":                   "群集的ID与cluster_name冲突，不能同时设置",
			"cluster_name":                 "群集的名称与cluster_id冲突，不能同时设置",
			"tags":                         "集群的标签",
			"result_output_file":           "用于保存结果",
			"list":                         "kubernetes集群的信息列表每个元素都包含以下属性：",
			"cluster_extra_args":           "主节点组件的自定义参数",
			"kube_scheduler":               "kube-scheduler的自定义参数",
			"kube_apiserver":               "kube-apiserver的自定义参数",
			"kube_controller_manager":      "kube-controller-manager的自定义参数",
			"certification_authority":      "集群的CA证书",
			"deletion_protection":          "是否开启集群删除保护",
			"cluster_node_num":             "集群节点数量",
			"worker_instances_list":        "工作节点列表",
			"instance_role":                "实例角色",
			"instance_state":               "实例状态",
			"failed_reason":                "失败原因",
			"lan_ip":                       "内网IP",
			"instance_id":                  "实例ID",
			"kube_config":                  "kubernetes配置",
			"cluster_external_endpoint":    "集群外网访问地址",
			"cluster_deploy_type":          "集群部署类型",
			"cluster_ipvs":                 "是否开启ipvs",
			"user_name":                    "用户名",
			"network_type":                 "集群网络类型",
			"kube_proxy_mode":              "kube-proxy模式",
			"password":                     "密码",
			"kube_config_intranet":         "内网kubernetes配置",
			"cluster_os":                   "集群操作系统",
			"project_id":                   "项目ID",
			"cluster_cidr":                 "集群的网络地址块",
			"ignore_cluster_cidr_conflict": "是否忽略集群CIDR冲突错误",
			"service_cidr":                 "集群的网络地址块",
			"cluster_desc":                 "集群描述",
			"cluster_max_pod_num":          "集群每个节点的最大Pod数",
			"is_non_static_ip_mode":        "是否开启非静态IP模式",
			"claim_expired_seconds":        "回收ENI的过期秒数",
			"domain":                       "域名",
			"security_policy":              "安全策略",
			"cluster_as_enabled":           "是否开启集群节点自动扩展",
			"node_name_type":               "集群节点名称类型",
			"eni_subnet_ids":               "VPC-CNI网络模式下的子网ID",
			"container_runtime":            "容器运行时",
			"cluster_version":              "集群版本",
			"vpc_id":                       "集群的VPC ID", "cluster_max_service_num": "集群的最大服务数",
			"pgw_endpoint": "PGW访问地址",
		},
	})
}

func dataSourceTencentCloudKubernetesClusters() *schema.Resource {
	return &schema.Resource{
		Description: "Cloud TKE Cluster",
		Read:        dataSourceTencentCloudKubernetesClustersRead,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:          schema.TypeString,
				ConflictsWith: []string{"cluster_name"},
				Description:   "ID of the cluster. Conflict with cluster_name, can not be set at the same time.",
				Optional:      true,
			},
			"cluster_name": {
				Type:          schema.TypeString,
				ConflictsWith: []string{"cluster_id"},
				Optional:      true,
				Description:   "Name of the cluster. Conflict with cluster_id, can not be set at the same time.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tags of the cluster.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An information list of kubernetes clusters. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: tkeClusterInfo(),
				},
			},
		},
	}

}
func dataSourceTencentCloudKubernetesClustersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tke_kubernetes_clusters.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	var (
		id   string
		name string
	)

	if v, ok := d.GetOk("cluster_id"); ok {
		id = v.(string)
	}

	if v, ok := d.GetOk("cluster_name"); ok {
		name = v.(string)
	}

	tags := helper.GetTags(d, "tags")

	infos, err := service.DescribeClusters(ctx, id, name)
	if err != nil && id == "" {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			infos, err = service.DescribeClusters(ctx, id, name)
			if err != nil {
				return retryError(err)
			}
			return nil
		})
	}

	if err != nil {
		return err
	}

	list := make([]map[string]interface{}, 0, len(infos))

	//var emptyStrFunc = func(ptr *string) string {
	//	if ptr == nil {
	//		return ""
	//	} else {
	//		return *ptr
	//	}
	//}

LOOP:
	for _, info := range infos {
		if len(tags) > 0 {
			for k, v := range tags {
				if info.Tags[k] != v {
					continue LOOP
				}
			}
		}

		var infoMap = map[string]interface{}{}
		infoMap["cluster_id"] = info.ClusterId
		infoMap["cluster_name"] = info.ClusterName
		infoMap["cluster_desc"] = info.ClusterDescription
		infoMap["cluster_os"] = tkeToShowClusterOs(info.ClusterOs)
		infoMap["cluster_deploy_type"] = info.DeployType
		infoMap["cluster_version"] = info.ClusterVersion
		infoMap["cluster_ipvs"] = info.Ipvs
		infoMap["cluster_as_enabled"] = info.AsEnabled
		infoMap["node_name_type"] = info.NodeNameType

		infoMap["cluster_extra_args"] = []map[string]interface{}{{
			"kube_apiserver":          info.ExtraArgs.KubeAPIServer,
			"kube_controller_manager": info.ExtraArgs.KubeControllerManager,
			"kube_scheduler":          info.ExtraArgs.KubeScheduler,
		}}
		infoMap["network_type"] = info.NetworkType
		infoMap["is_non_static_ip_mode"] = info.IsNonStaticIpMode
		infoMap["deletion_protection"] = info.DeletionProtection
		infoMap["kube_proxy_mode"] = info.KubeProxyMode
		infoMap["vpc_id"] = info.VpcId
		infoMap["project_id"] = info.ProjectId
		infoMap["cluster_cidr"] = info.ClusterCidr
		infoMap["ignore_cluster_cidr_conflict"] = info.IgnoreClusterCidrConflict
		infoMap["cluster_max_pod_num"] = info.MaxNodePodNum
		infoMap["cluster_max_service_num"] = info.MaxClusterServiceNum
		infoMap["service_cidr"] = info.ServiceCIDR
		infoMap["eni_subnet_ids"] = info.EniSubnetIds
		infoMap["claim_expired_seconds"] = info.ClaimExpiredSeconds
		infoMap["cluster_node_num"] = info.ClusterNodeNum
		infoMap["tags"] = info.Tags

		_, workers, err := service.DescribeClusterInstances(ctx, info.ClusterId)
		if err != nil {
			_, workers, err = service.DescribeClusterInstances(ctx, info.ClusterId)

		}
		if err != nil {
			log.Printf("[CRITAL]%s cloud_tke_kubernetes_clusters DescribeClusterInstances fail, reason:%s\n ", logId, err.Error())
			//return err
		}

		workerInstancesList := make([]map[string]interface{}, 0, len(workers))
		for _, cvm := range workers {
			tempMap := make(map[string]interface{})
			tempMap["instance_id"] = cvm.InstanceId
			tempMap["instance_role"] = cvm.InstanceRole
			tempMap["instance_state"] = cvm.InstanceState
			tempMap["failed_reason"] = cvm.FailedReason
			workerInstancesList = append(workerInstancesList, tempMap)
		}

		infoMap["worker_instances_list"] = workerInstancesList

		//securityRet, err := service.DescribeClusterSecurity(ctx, info.ClusterId)
		//if err != nil {
		//	securityRet, err = service.DescribeClusterSecurity(ctx, info.ClusterId)
		//}
		//
		//if err != nil {
		//	log.Printf("[CRITAL]%s cloud_tke_kubernetes_clusters DescribeClusterSecurity fail, reason:%s\n ", logId, err.Error())
		//	return err
		//}
		//
		//policies := make([]string, 0, len(securityRet.Response.SecurityPolicy))
		//for _, v := range securityRet.Response.SecurityPolicy {
		//	policies = append(policies, *v)
		//}
		//
		//infoMap["user_name"] = emptyStrFunc(securityRet.Response.UserName)
		//infoMap["password"] = emptyStrFunc(securityRet.Response.Password)
		//infoMap["certification_authority"] = emptyStrFunc(securityRet.Response.CertificationAuthority)
		//infoMap["cluster_external_endpoint"] = emptyStrFunc(securityRet.Response.ClusterExternalEndpoint)
		//infoMap["domain"] = emptyStrFunc(securityRet.Response.Domain)
		//infoMap["pgw_endpoint"] = emptyStrFunc(securityRet.Response.PgwEndpoint)
		//infoMap["security_policy"] = policies

		//config, err := service.DescribeClusterConfig(ctx, info.ClusterId, true)
		//if err != nil {
		//	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		//		config, err = service.DescribeClusterConfig(ctx, d.Id(), true)
		//		if err != nil {
		//			return retryError(err)
		//		}
		//		return nil
		//	})
		//}
		//if err != nil {
		//	log.Printf("[CRITAL]%s cloud_tke_kubernetes_clusters DescribeClusterInstances fail, reason:%s\n ", logId, err.Error())
		//	return err
		//}
		//
		//intranetConfig, err := service.DescribeClusterConfig(ctx, info.ClusterId, false)
		//if err != nil {
		//	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		//		config, err = service.DescribeClusterConfig(ctx, d.Id(), false)
		//		if err != nil {
		//			return retryError(err)
		//		}
		//		return nil
		//	})
		//}
		//if err != nil {
		//	log.Printf("[CRITAL]%s cloud_tke_kubernetes_clusters DescribeClusterInstances fail, reason:%s\n ", logId, err.Error())
		//	return err
		//}

		//infoMap["kube_config"] = config
		//infoMap["kube_config_intranet"] = intranetConfig
		list = append(list, infoMap)
	}

	d.SetId("KubernetesClusters" + name + id)
	err = d.Set("list", list)
	if err != nil {
		log.Printf("[CRITAL]%s provider set cloud_tke_kubernetes_clusters list fail, reason:%s\n ", logId, err.Error())
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), list); err != nil {
			return err
		}
	}
	return nil
}
