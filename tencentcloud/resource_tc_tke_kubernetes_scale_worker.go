/*
Provide a resource to increase instance to cluster

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

# Example Usage

```hcl

	variable "availability_zone" {
	  default = "ap-guangzhou-3"
	}

	variable "subnet" {
	  default = "subnet-pqfek0t8"
	}

	variable "scale_instance_type" {
	  default = "S2.LARGE16"
	}

	resource cloud_tke_kubernetes_scale_worker test_scale {
	  cluster_id = "cls-godovr32"
	  desired_pod_num = 16
	  labels = {
	    "test1" = "test1",
	    "test2" = "test2",
	  }
	  worker_config {
	    count                      = 3
	    availability_zone          = var.availability_zone
	    instance_type              = var.scale_instance_type
	    subnet_id                  = var.subnet
	    system_disk_type           = "CLOUD_SSD"
	    system_disk_size           = 50
	    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
	    internet_max_bandwidth_out = 100
	    public_ip_assigned         = true

	    data_disk {
	      disk_type = "CLOUD_PREMIUM"
	      disk_size = 50
	    }

	    enhanced_security_service = false
	    enhanced_monitor_service  = false
	    user_data                 = "dGVzdA=="
	    password                  = "AABBccdd1122"
	  }
	}

```

# Use Kubelet

```hcl

	variable "availability_zone" {
	  default = "ap-guangzhou-3"
	}

	variable "subnet" {
	  default = "subnet-pqfek0t8"
	}

	variable "scale_instance_type" {
	  default = "S2.LARGE16"
	}

	resource cloud_tke_kubernetes_scale_worker test_scale {
	  cluster_id = "cls-godovr32"

	  extra_args = [
	 	"root-dir=/var/lib/kubelet"
	  ]

	   labels = {
	    "test1" = "test1",
	    "test2" = "test2",
	  }

	  worker_config {
	    count                      = 3
	    availability_zone          = var.availability_zone
	    instance_type              = var.scale_instance_type
	    subnet_id                  = var.subnet
	    system_disk_type           = "CLOUD_SSD"
	    system_disk_size           = 50
	    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
	    internet_max_bandwidth_out = 100
	    public_ip_assigned         = true

	    data_disk {
	      disk_type = "CLOUD_PREMIUM"
	      disk_size = 50
	    }

	    enhanced_security_service = false
	    enhanced_monitor_service  = false
	    user_data                 = "dGVzdA=="
	    password                  = "AABBccdd1122"
	  }
	}

```
*/
package tencentcloud

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"strings"
	"time"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_scale_worker", CNDescription{
		TerraformTypeCN: "扩展集群节点",
		DescriptionCN:   "提供向集群增加实例的资源，用于扩展Kubernetes集群节点。",
		AttributesCN: map[string]string{
			"cluster_id":    "集群ID",
			"worker_config": "worker实例配置信息",
			"labels":        "worker实例标签",
			"extra_args":    "worker实例额外参数",
			//"gpu_args":              "worker实例GPU参数",
			"unschedulable": "worker实例是否参与调度",
			//"desired_pod_num":       "worker实例期望pod数量",
			"docker_graph_path":     "worker实例docker graph路径",
			"mount_target":          "worker实例挂载目标",
			"data_disk":             "worker实例数据盘配置",
			"worker_instances_list": "worker实例列表",
			"custom_disk":           "worker实例自定义磁盘",
			"mig_enable":            "worker实例是否开启迁移",
			"driver":                "worker实例驱动",
			"cuda":                  "worker实例cuda",
			"cudnn":                 "worker实例cudnn",
			"system_disk_type":      "worker实例系统盘类型",
			//"cam_role_name":                       "worker实例cam角色名称",
			"disaster_recover_group_id":           "worker实例容灾组ID",
			"instance_charge_type_prepaid_period": "worker实例预付费周期",
			"bandwidth_package_id":                "worker实例带宽包ID",
			"public_ip_assigned":                  "worker实例是否分配公网IP",
			"password":                            "worker实例密码",
			"instance_type":                       "worker实例类型",
			"system_disk_size":                    "worker实例系统盘大小",
			"img_id":                              "worker实例镜像ID",
			"subnet_id":                           "worker实例子网ID",
			"image_max_bandwidth_out":             "worker实例镜像最大出带宽",
			"internet_charge_type":                "worker实例网络计费类型",
			"security_group_id":                   "worker实例安全组ID",
			"count":                               "worker实例数量",
			"availability_zone":                   "worker实例可用区",
			"instance_name":                       "worker实例名称",
			"enhanced_security_service":           "worker实例是否开启安全服务",
			//"instance_charge_type_prepaid_renew_flag": "worker实例预付费续费标识",
			"instance_charge_type": "worker实例计费类型",
			"disk_type":            "worker实例磁盘类型",
			//"file_system":              "worker实例文件系统",
			//"auto_format_and_mount":    "worker实例是否自动格式化和挂载",
			"disk_partition":           "worker实例磁盘分区",
			"disk_size":                "worker实例磁盘大小",
			"snapshot_id":              "worker实例快照ID",
			"encrypt":                  "worker实例是否加密",
			"kms_key_id":               "worker实例KMS密钥ID",
			"key_ids":                  "worker实例密钥ID",
			"enhanced_monitor_service": "worker实例是否开启监控服务",
			"user_data":                "worker实例用户数据",
			"hostname":                 "worker实例主机名",
			//"hpc_cluster_id":           "worker实例HPC集群ID",
			//"disaster_recover_group_ids": "worker实例容灾组ID",
			"security_group_ids":         "worker实例安全组ID",
			"internet_max_bandwidth_out": "worker实例最大出带宽",
			"custom_driver":              "worker实例自定义驱动",
			"lan_ip":                     "worker实例内网IP",
			"instance_id":                "worker实例ID",
			"instance_role":              "worker实例角色",
			"instance_state":             "worker实例状态",
			"failed_reason":              "worker实例失败原因",
		},
	})
}

func resourceTencentCloudTkeScaleWorker() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to increase instance to cluster",
		Create:      resourceTencentCloudTkeScaleWorkerCreate,
		Read:        resourceTencentCloudTkeScaleWorkerRead,
		Delete:      resourceTencentCloudTkeScaleWorkerDelete,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "ID of the cluster.",
			},
			"worker_config": {
				Type:     schema.TypeList,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: TkeCvmCreateInfo(),
				},
				Description: "Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.",
			},
			//advanced instance settings
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Labels of kubernetes scale worker created nodes.",
			},
			"extra_args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Custom parameter information related to the node.",
			},
			//"gpu_args": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	ForceNew: true,
			//	MaxItems: 1,
			//	Elem: &schema.Resource{
			//		Schema: TKEGpuArgsSetting(),
			//	},
			//	Description: "GPU driver parameters.",
			//},
			"unschedulable": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     0,
				Description: "Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.",
			},
			//"desired_pod_num": {
			//	Type:        schema.TypeInt,
			//	ForceNew:    true,
			//	Optional:    true,
			//	Description: "Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.",
			//},
			"docker_graph_path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Docker graph path. Default is `/var/lib/docker`.",
			},
			"mount_target": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Mount target. Default is not mounting.",
			},
			"data_disk": {
				Type:        schema.TypeList,
				ForceNew:    true,
				Optional:    true,
				MaxItems:    11,
				Description: "Configurations of data disk.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_type": {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							Default:      SYSTEM_DISK_TYPE_CLOUD_PREMIUM,
							ValidateFunc: validateAllowedStringValue(SYSTEM_DISK_ALLOW_TYPE),
							Description:  "Types of disk, available values: `CLOUD_PREMIUM` and `CLOUD_SSD` and `CLOUD_HSSD` and `CLOUD_TSSD`.",
						},
						"disk_size": {
							Type:        schema.TypeInt,
							ForceNew:    true,
							Optional:    true,
							Default:     0,
							Description: "Volume of disk in GB. Default is `0`.",
						},
						//"file_system": {
						//	Type:        schema.TypeString,
						//	ForceNew:    true,
						//	Optional:    true,
						//	Default:     "",
						//	Description: "File system, e.g. `ext3/ext4/xfs`.",
						//},
						//"auto_format_and_mount": {
						//	Type:        schema.TypeBool,
						//	Optional:    true,
						//	ForceNew:    true,
						//	Default:     false,
						//	Description: "Indicate whether to auto format and mount or not. Default is `false`.",
						//},
						"mount_target": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Default:     "",
							Description: "Mount target.",
						},
					},
				},
			},
			// Computed values
			"worker_instances_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: tkeCvmState(),
				},
				Description: "An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:",
			},
		},
	}
}

func resourceTencentCloudTkeScaleWorkerCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tke_kubernetes_scale_worker.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var cvms RunInstancesForNode
	var iAdvanced tke.InstanceAdvancedSettings
	cvms.Work = []string{}

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	clusterId := d.Get("cluster_id").(string)
	if clusterId == "" {
		return fmt.Errorf("`cluster_id` is empty.")
	}

	info, has, err := service.DescribeCluster(ctx, clusterId)
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			info, has, err = service.DescribeCluster(ctx, clusterId)
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
		return fmt.Errorf("cluster [%s] is not exist.", clusterId)
	}

	var userData string
	if workers, ok := d.GetOk("worker_config"); ok {
		workerList := workers.([]interface{})
		for index := range workerList {
			worker := workerList[index].(map[string]interface{})
			if v, ok := worker["user_data"]; ok {
				userData = v.(string)
			}
			paraJson, _, err := tkeGetCvmRunInstancesPara(worker, meta, info.VpcId, info.ProjectId)
			if err != nil {
				return err
			}
			cvms.Work = append(cvms.Work, paraJson)
		}
	}
	if len(cvms.Work) != 1 {
		return fmt.Errorf("only one additional configuration of virtual machines is now supported now, " +
			"so len(cvms.Work) should be 1")
	}

	dMap := make(map[string]interface{}, 5)
	//mount_target, docker_graph_path, data_disk, extra_args, desired_pod_num
	//iAdvancedParas := []string{"mount_target", "docker_graph_path", "extra_args", "data_disk", "desired_pod_num", "gpu_args"}
	iAdvancedParas := []string{"mount_target", "docker_graph_path", "extra_args", "data_disk"}
	for _, k := range iAdvancedParas {
		if v, ok := d.GetOk(k); ok {
			dMap[k] = v
		}
	}
	if userData != "" {
		dMap["user_data"] = userData
	}
	iAdvanced = tkeGetInstanceAdvancedPara(dMap, meta)

	iAdvanced.Labels = GetTkeLabels(d, "labels")
	if temp, ok := d.GetOk("unschedulable"); ok {
		iAdvanced.Unschedulable = helper.Int64(int64(temp.(int)))
	}

	instanceIds, err := service.CreateClusterInstances(ctx, clusterId, cvms.Work[0], iAdvanced)
	if err != nil {
		return err
	}

	workerInstancesList := make([]map[string]interface{}, 0, len(instanceIds))
	for _, v := range instanceIds {
		if v == "" {
			return fmt.Errorf("CreateClusterInstances return one instanceId is empty")
		}
		infoMap := make(map[string]interface{})
		infoMap["instance_id"] = v
		infoMap["instance_role"] = TKE_ROLE_WORKER
		workerInstancesList = append(workerInstancesList, infoMap)
	}

	if err = d.Set("worker_instances_list", workerInstancesList); err != nil {
		return err
	}

	md := md5.New()

	if _, err = md.Write([]byte(clusterId)); err != nil {
		return err
	}

	instanceIdJoin := strings.Join(instanceIds, "#")

	if _, err = md.Write([]byte(instanceIdJoin)); err != nil {
		return err
	}

	id := fmt.Sprintf("TkeScaleWorker.%x", md.Sum(nil))

	d.SetId(id)

	//wait for LANIP
	time.Sleep(readRetryTimeout)
	return resourceTencentCloudTkeScaleWorkerRead(d, meta)
}

func resourceTencentCloudTkeScaleWorkerRead(d *schema.ResourceData, meta interface{}) error {

	defer logElapsed("resource.cloud_tke_kubernetes_scale_worker.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	clusterId := d.Get("cluster_id").(string)
	if clusterId == "" {
		return fmt.Errorf("tke.`cluster_id` is empty.")
	}

	_, has, err := service.DescribeCluster(ctx, clusterId)
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			_, has, err = service.DescribeCluster(ctx, clusterId)
			if err != nil {
				return retryError(err)
			}
			return nil
		})
	}

	if err != nil {
		return nil
	}
	// The cluster has been deleted
	if !has {
		d.SetId("")
		return nil
	}

	oldWorkerInstancesList := d.Get("worker_instances_list").([]interface{})

	instanceMap := make(map[string]bool)

	for _, v := range oldWorkerInstancesList {

		infoMap, ok := v.(map[string]interface{})

		if !ok || infoMap["instance_id"] == nil {
			return fmt.Errorf("worker_instances_list is broken.")
		}
		instanceId, ok := infoMap["instance_id"].(string)
		if !ok || instanceId == "" {
			return fmt.Errorf("worker_instances_list.instance_id is broken.")
		}

		if instanceMap[instanceId] {
			log.Printf("[WARN]The same instance id exists in the list")
		}

		instanceMap[instanceId] = true

	}

	_, workers, err := service.DescribeClusterInstances(ctx, clusterId)
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			_, workers, err = service.DescribeClusterInstances(ctx, clusterId)

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

	newWorkerInstancesList := make([]map[string]interface{}, 0, len(workers))
	labelsMap := make(map[string]string)
	for _, cvm := range workers {
		if _, ok := instanceMap[cvm.InstanceId]; !ok {
			continue
		}
		tempMap := make(map[string]interface{})
		tempMap["instance_id"] = cvm.InstanceId
		tempMap["instance_role"] = cvm.InstanceRole
		tempMap["instance_state"] = cvm.InstanceState
		tempMap["failed_reason"] = cvm.FailedReason
		tempMap["lan_ip"] = cvm.LanIp
		newWorkerInstancesList = append(newWorkerInstancesList, tempMap)

		labels := cvm.InstanceAdvancedSettings.Labels

		for _, v := range labels {
			labelsMap[*v.Name] = *v.Value
		}
	}

	// The machines I generated was deleted by others.
	if len(newWorkerInstancesList) == 0 {
		d.SetId("")
		return nil
	}

	_ = d.Set("labels", labelsMap)

	return d.Set("worker_instances_list", newWorkerInstancesList)
}
func resourceTencentCloudTkeScaleWorkerDelete(d *schema.ResourceData, meta interface{}) error {

	defer logElapsed("resource.cloud_tke_kubernetes_scale_worker.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	clusterId := d.Get("cluster_id").(string)

	if clusterId == "" {
		return fmt.Errorf("`cluster_id` is empty.")
	}

	_, has, err := service.DescribeCluster(ctx, clusterId)
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			_, has, err = service.DescribeCluster(ctx, clusterId)
			if err != nil {
				return retryError(err)
			}
			return nil
		})
	}

	if err != nil {
		return nil
	}
	// The cluster has been deleted
	if !has {
		return nil
	}
	workerInstancesList := d.Get("worker_instances_list").([]interface{})

	instanceMap := make(map[string]bool)

	for _, v := range workerInstancesList {

		infoMap, ok := v.(map[string]interface{})

		if !ok || infoMap["instance_id"] == nil {
			return fmt.Errorf("worker_instances_list is broken.")
		}
		instanceId, ok := infoMap["instance_id"].(string)
		if !ok || instanceId == "" {
			return fmt.Errorf("worker_instances_list.instance_id is broken.")
		}

		if instanceMap[instanceId] {
			log.Printf("[WARN]The same instance id exists in the list")
		}

		instanceMap[instanceId] = true

	}

	_, workers, err := service.DescribeClusterInstances(ctx, clusterId)
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			_, workers, err = service.DescribeClusterInstances(ctx, clusterId)

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

	needDeletes := []string{}
	for _, cvm := range workers {
		if _, ok := instanceMap[cvm.InstanceId]; ok {
			needDeletes = append(needDeletes, cvm.InstanceId)
		}
	}
	// The machines I generated was deleted by others.
	if len(needDeletes) == 0 {
		return nil
	}

	err = service.DeleteClusterInstances(ctx, clusterId, needDeletes)
	if err != nil {
		err = resource.Retry(3*writeRetryTimeout, func() *resource.RetryError {
			err = service.DeleteClusterInstances(ctx, clusterId, needDeletes)

			if e, ok := err.(*errors.CloudSDKError); ok {
				if e.GetCode() == "InternalError.ClusterNotFound" {
					return nil
				}

				if e.GetCode() == "InternalError.Param" &&
					strings.Contains(e.GetMessage(), `PARAM_ERROR[some instances []is not in right state`) {
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
