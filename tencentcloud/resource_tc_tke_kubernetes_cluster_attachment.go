/*
Provide a resource to attach an existing  cvm to kubernetes cluster.

# Example Usage

```hcl

	variable "availability_zone" {
	  default = "ap-guangzhou-3"
	}

	variable "cluster_cidr" {
	  default = "172.16.0.0/16"
	}

	variable "default_instance_type" {
	  default = "S1.SMALL1"
	}

	data "tencentcloudenterprise_cvm_images" "default" {
	  image_type = ["PUBLIC_IMAGE"]
	  os_name    = "centos"
	}

	data "tencentcloudenterprise_vpc_subnets" "vpc" {
	  is_default        = true
	  availability_zone = var.availability_zone
	}

	data "tencentcloudenterprise_cvm_instance_types" "default" {
	  filter {
	    name   = "instance-family"
	    values = ["SA2"]
	  }

	  cpu_core_count = 8
	  memory_size    = 16
	}

	resource "tencentcloudenterprise_cvm_instance" "foo" {
	  instance_name     = "tf-auto-test-1-1"
	  availability_zone = var.availability_zone
	  image_id          = data.cloud_cvm_images.default.images.0.image_id
	  instance_type     = var.default_instance_type
	  system_disk_type  = "CLOUD_PREMIUM"
	  system_disk_size  = 50
	}

	resource "tencentcloudenterprise_tke_kubernetes_cluster" "managed_cluster" {
	  vpc_id                  = data.cloud_vpc_subnets.vpc.instance_list.0.vpc_id
	  cluster_cidr            = "10.1.0.0/16"
	  cluster_max_pod_num     = 32
	  cluster_name            = "keep"
	  cluster_desc            = "test cluster desc"
	  cluster_max_service_num = 32

	  worker_config {
	    count                      = 1
	    availability_zone          = var.availability_zone
	    instance_type              = var.default_instance_type
	    system_disk_type           = "CLOUD_SSD"
	    system_disk_size           = 60
	    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
	    internet_max_bandwidth_out = 100
	    public_ip_assigned         = true
	    subnet_id                  = data.cloud_vpc_subnets.vpc.instance_list.0.subnet_id

	    data_disk {
	      disk_type = "CLOUD_PREMIUM"
	      disk_size = 50
	    }

	    enhanced_security_service = false
	    enhanced_monitor_service  = false
	    user_data                 = "dGVzdA=="
	    password                  = "ZZXXccvv1212"
	  }

	  cluster_deploy_type = "MANAGED_CLUSTER"
	}

	resource "tencentcloudenterprise_tke_kubernetes_cluster_attachment" "test_attach" {
	  cluster_id  = cloud_tke_kubernetes_cluster.managed_cluster.id
	  instance_id = cloud_cvm_instance.foo.id
	  password    = "Lo4wbdit"

	  labels = {
	    "test1" = "test1",
	    "test2" = "test2",
	  }

	  worker_config_overrides {
	    desired_pod_num = 8
	  }
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_attachment", CNDescription{
		TerraformTypeCN: "添加已有节点到集群",
		DescriptionCN:   "提供将现有CVM实例添加到Kubernetes集群的资源。",
		AttributesCN: map[string]string{
			"instance_id":  "实例ID",
			"cluster_id":   "集群ID",
			"password":     "密码",
			"key_ids":      "密钥ID",
			"host_name":    "主机名",
			"woker_config": "节点配置信息",
			//"worker_config_overrides": "节点配置信息",
			"labels":            "标签",
			"unschedulable":     "是否参与调度",
			"security_groups":   "安全组",
			"state":             "状态",
			"hostname":          "主机名",
			"worker_config":     "节点配置信息",
			"docker_graph_path": "Docker 图形路径。 默认为 `/var/lib/docker`.Docker 图形路径。 默认为“/var/lib/docker”。",
			"data_disk":         "数据盘配置信息。数据盘配置信息。",
			"disk_type":         "磁盘类型。可用值：`CLOUD_PREMIUM` 和 `CLOUD_SSD`。",
			"disk_size":         "磁盘容量（单位：GB）。默认为 `0`。",
			//"file_system":           "文件系统，例如 `ext3/ext4/xfs`。",
			//"auto_format_and_mount": "是否自动格式化和挂载。默认为 `false`。",
			"mount_target": "挂载目标。",
			//"disk_partition":        "要挂载的设备或分区的名称。注意：此参数不支持在节点池中设置，否则会导致挂载错误。",
			"extra_args": "与节点相关的自定义参数信息。这是一个白名单参数。",
			"user_data":  "Base64 编码的用户数据文本，长度限制为 16KB。",
			//"is_schedule":           "指示是否调度添加节点。默认为 `true`。",
			//"desired_pod_num": "指示在节点中设置所需的 pod 数。当集群是 podCIDR 时有效。",
			//"gpu_args":        "GPU 驱动程序参数。",
			"custom_driver": "自定义 GPU 驱动程序。格式如：`{address: String}`。`address`：自定义 GPU 驱动程序地址。",
			"mig_enable":    "是否启用 MIG。",
			"driver":        "GPU 驱动程序版本。格式如：`{ version: String, name: String }`。`version`：GPU 驱动程序或 CUDA 的版本；`name`：GPU 驱动程序或 CUDA 的名称。",
			"cuda":          "CUDA 版本。格式如：`{ version: String, name: String }`。`version`：GPU 驱动程序或 CUDA 的版本；`name`：GPU 驱动程序或 CUDA 的名称。",
			"cudnn":         "cuDNN 版本。格式如：`{ version: String, name: String, doc_name: String, dev_name: String }`。`version`：cuDNN 版本；`name`：cuDNN 名称；`doc_name`：cuDNN 的文档名称；`dev_name`：cuDNN 的开发名称。",
		},
	})
}
func TKEGpuArgsSetting() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mig_enable": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether to enable MIG.",
		},
		"driver": {
			Type:         schema.TypeMap,
			Optional:     true,
			ValidateFunc: validateTkeGpuDriverVersion,
			Description:  "GPU driver version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.",
		},
		"cuda": {
			Type:         schema.TypeMap,
			Optional:     true,
			ValidateFunc: validateTkeGpuDriverVersion,
			Description:  "CUDA  version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.",
		},
		"cudnn": {
			Type:         schema.TypeMap,
			Optional:     true,
			ValidateFunc: validateTkeGpuDriverVersion,
			Description: "cuDNN version. Format like: `{ version: String, name: String, doc_name: String, dev_name: String }`." +
				" `version`: cuDNN version; `name`: cuDNN name; `doc_name`: Doc name of cuDNN; `dev_name`: Dev name of cuDNN.",
		},
		"custom_driver": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "Custom GPU driver. Format like: `{address: String}`. `address`: URL of custom GPU driver address.",
		},
	}
}

func TkeInstanceAdvancedSetting() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mount_target": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "Mount target. Default is not mounting.",
		},
		"docker_graph_path": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Default:     "/var/lib/docker",
			Description: "Docker graph path. Default is `/var/lib/docker`.",
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
						Description:  "Types of disk, available values: `CLOUD_PREMIUM` and `CLOUD_SSD`.",
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
					//"disk_partition": {
					//	Type:        schema.TypeString,
					//	ForceNew:    true,
					//	Optional:    true,
					//	Description: "The name of the device or partition to mount. NOTE: this argument doesn't support setting in node pool, or will leads to mount error.",
					//},
				},
			},
		},
		"extra_args": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Custom parameter information related to the node. This is a white-list parameter.",
		},
		"user_data": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Optional:    true,
			Description: "Base64-encoded User Data text, the length limit is 16KB.",
		},
		//"is_schedule": {
		//	Type:        schema.TypeBool,
		//	ForceNew:    true,
		//	Optional:    true,
		//	Default:     true,
		//	Description: "Indicate to schedule the adding node or not. Default is true.",
		//},
		//"desired_pod_num": {
		//	Type:        schema.TypeInt,
		//	ForceNew:    true,
		//	Optional:    true,
		//	Description: "Indicate to set desired pod number in node. valid when the cluster is podCIDR.",
		//},
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
	}
}

func resourceTencentCloudTkeClusterAttachment() *schema.Resource {
	schemaBody := map[string]*schema.Schema{
		"cluster_id": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
			Description: "ID of the cluster.",
		},
		"instance_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "ID of the CVM instance, this cvm will reinstall the system.",
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
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.",
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
		"worker_config": {
			Type:     schema.TypeList,
			ForceNew: true,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: TkeInstanceAdvancedSetting(),
			},
			Description: "Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.",
		},
		//"worker_config_overrides": {
		//	Type:     schema.TypeList,
		//	ForceNew: true,
		//	MaxItems: 1,
		//	Optional: true,
		//	Elem: &schema.Resource{
		//		Schema: TkeInstanceAdvancedSetting(),
		//	},
		//	Description: "Override variable worker_config, commonly used to attach existing instances.",
		//},
		"labels": {
			Type:        schema.TypeMap,
			Optional:    true,
			ForceNew:    true,
			Description: "Labels of tke attachment exits CVM.",
		},
		"unschedulable": {
			Type:        schema.TypeInt,
			Optional:    true,
			ForceNew:    true,
			Default:     0,
			Description: "Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.",
		},
		//compute
		"security_groups": {
			Type:        schema.TypeSet,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Computed:    true,
			Description: "A list of security group IDs after attach to cluster.",
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "State of the node.",
		},
	}

	return &schema.Resource{
		Description: "Provide a resource to attach an existing  cvm to kubernetes cluster.",
		Create:      resourceTencentCloudTkeClusterAttachmentCreate,
		Read:        resourceTencentCloudTkeClusterAttachmentRead,
		Delete:      resourceTencentCloudTkeClusterAttachmentDelete,
		Schema:      schemaBody,
	}
}

func tkeGetInstanceAdvancedPara(dMap map[string]interface{}, meta interface{}) (setting tke.InstanceAdvancedSettings) {
	setting = tke.InstanceAdvancedSettings{}
	if v, ok := dMap["mount_target"]; ok {
		setting.MountTarget = helper.String(v.(string))
	}

	if v, ok := dMap["data_disk"]; ok {

		dataDisks := v.([]interface{})
		setting.DataDisks = make([]*tke.DataDisk, 0, len(dataDisks))

		for _, d := range dataDisks {
			var (
				value    = d.(map[string]interface{})
				diskType = value["disk_type"].(string)
				diskSize = int64(value["disk_size"].(int))
				//fileSystem         = value["file_system"].(string)
				//autoFormatAndMount = value["auto_format_and_mount"].(bool)
				mountTarget = value["mount_target"].(string)
				//diskPartition      = value["disk_partition"].(string)
				dataDisk = tke.DataDisk{
					DiskType: &diskType,
					//FileSystem:         &fileSystem,
					//AutoFormatAndMount: &autoFormatAndMount,
					MountTarget: &mountTarget,
					//DiskPartition:      &diskPartition,
				}
			)
			if diskSize > 0 {
				dataDisk.DiskSize = &diskSize
			}
			setting.DataDisks = append(setting.DataDisks, &dataDisk)
		}
	}
	if v, ok := dMap["is_schedule"]; ok {
		setting.Unschedulable = helper.BoolToInt64Ptr(!v.(bool))
	}

	if v, ok := dMap["user_data"]; ok {
		setting.UserScript = helper.String(v.(string))
	}

	if v, ok := dMap["docker_graph_path"]; ok {
		setting.DockerGraphPath = helper.String(v.(string))
	}

	//if v, ok := dMap["desired_pod_num"]; ok {
	//	setting.DesiredPodNumber = helper.Int64(int64(v.(int)))
	//}

	if temp, ok := dMap["extra_args"]; ok {
		extraArgs := helper.InterfacesStrings(temp.([]interface{}))
		clusterExtraArgs := tke.InstanceExtraArgs{}
		clusterExtraArgs.Kubelet = make([]*string, 0)
		for i := range extraArgs {
			clusterExtraArgs.Kubelet = append(clusterExtraArgs.Kubelet, &extraArgs[i])
		}
		setting.ExtraArgs = &clusterExtraArgs
	}

	// get gpu_args
	//if v, ok := dMap["gpu_args"]; ok && len(v.([]interface{})) > 0 {
	//	gpuArgs := v.([]interface{})[0].(map[string]interface{})
	//
	//	var (
	//		migEnable    = gpuArgs["mig_enable"].(bool)
	//		driver       = gpuArgs["driver"].(map[string]interface{})
	//		cuda         = gpuArgs["cuda"].(map[string]interface{})
	//		cudnn        = gpuArgs["cudnn"].(map[string]interface{})
	//		customDriver = gpuArgs["custom_driver"].(map[string]interface{})
	//	)
	//	tkeGpuArgs := tke.GPUArgs{}
	//	tkeGpuArgs.MIGEnable = &migEnable
	//	if len(driver) > 0 {
	//		tkeGpuArgs.Driver = &tke.DriverVersion{
	//			Version: helper.String(driver["version"].(string)),
	//			Name:    helper.String(driver["name"].(string)),
	//		}
	//	}
	//	if len(cuda) > 0 {
	//		tkeGpuArgs.CUDA = &tke.DriverVersion{
	//			Version: helper.String(cuda["version"].(string)),
	//			Name:    helper.String(cuda["name"].(string)),
	//		}
	//	}
	//	if len(cudnn) > 0 {
	//		tkeGpuArgs.CUDNN = &tke.CUDNN{
	//			Version: helper.String(cudnn["version"].(string)),
	//			Name:    helper.String(cudnn["name"].(string)),
	//		}
	//		if cudnn["doc_name"] != nil {
	//			tkeGpuArgs.CUDNN.DocName = helper.String(cudnn["doc_name"].(string))
	//		}
	//		if cudnn["dev_name"] != nil {
	//			tkeGpuArgs.CUDNN.DevName = helper.String(cudnn["dev_name"].(string))
	//		}
	//	}
	//	if len(customDriver) > 0 {
	//		tkeGpuArgs.CustomDriver = &tke.CustomDriver{
	//			Address: helper.String(customDriver["address"].(string)),
	//		}
	//	}
	//	setting.GPUArgs = &tkeGpuArgs
	//}

	return setting
}
func resourceTencentCloudTkeClusterAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tke_kubernetes_cluster_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	tkeService := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	cvmService := CvmService{client: meta.(*TencentCloudClient).apiV3Conn}
	instanceId, clusterId := "", ""

	if items := strings.Split(d.Id(), "_"); len(items) != 2 {
		return fmt.Errorf("the resource id is corrupted")
	} else {
		instanceId, clusterId = items[0], items[1]
	}

	/*tke has been deleted*/
	_, has, err := tkeService.DescribeCluster(ctx, clusterId)
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			_, has, err = tkeService.DescribeCluster(ctx, clusterId)
			if err != nil {
				return retryError(err, InternalError)
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

	/*cvm has been deleted*/
	var instance *cvm.Instance
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instance, err = cvmService.DescribeInstanceById(ctx, instanceId)
		if err != nil {
			return retryError(err, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if instance == nil {
		d.SetId("")
		return nil
	}

	instanceState := ""
	has = false
	/*attachment has been  deleted*/

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		_, workers, err := tkeService.DescribeClusterInstances(ctx, clusterId)
		if err != nil {
			return retryError(err, InternalError)
		}
		for _, worker := range workers {
			if worker.InstanceId == instanceId {
				has = true
				instanceState = worker.InstanceState
				if worker.InstanceState == "failed" {
					return resource.NonRetryableError(fmt.Errorf("cvm instance %s attach to cluster %s fail,reason:%s",
						worker.InstanceId, clusterId, worker.FailedReason))
				}

				if worker.InstanceState != "running" {
					return resource.RetryableError(fmt.Errorf("cvm instance  %s in tke status is %s, retry...",
						worker.InstanceId, worker.InstanceState))
				}
				_ = d.Set("unschedulable", worker.InstanceAdvancedSettings.Unschedulable)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if !has {
		d.SetId("")
		return nil
	}

	if len(instance.LoginSettings.KeyIds) > 0 {
		_ = d.Set("key_ids", instance.LoginSettings.KeyIds)
	}

	_ = d.Set("security_groups", helper.StringsInterfaces(instance.SecurityGroupIds))
	_ = d.Set("state", instanceState)
	return nil
}

func resourceTencentCloudTkeClusterAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tke_kubernetes_cluster_attachment.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	tkeService := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	cvmService := CvmService{client: meta.(*TencentCloudClient).apiV3Conn}

	request := tke.NewAddExistedInstancesRequest()

	instanceId := helper.String(d.Get("instance_id").(string))
	request.ClusterId = helper.String(d.Get("cluster_id").(string))
	request.InstanceIds = []*string{instanceId}
	request.LoginSettings = &tke.LoginSettings{}

	var loginSettingsNumbers = 0

	if v, ok := d.GetOk("key_ids"); ok {
		request.LoginSettings.KeyIds = helper.Strings(helper.InterfacesStrings(v.([]interface{})))
		loginSettingsNumbers++
	}

	if v, ok := d.GetOk("password"); ok {
		request.LoginSettings.Password = helper.String(v.(string))
		loginSettingsNumbers++
	}

	if loginSettingsNumbers != 1 {
		return fmt.Errorf("parameters `key_ids` and `password` must set and only set one")
	}

	request.InstanceAdvancedSettings = &tke.InstanceAdvancedSettings{}
	if workConfig, ok := d.GetOk("worker_config"); ok {
		workConfigList := workConfig.([]interface{})
		if len(workConfigList) == 1 {
			workConfigPara := workConfigList[0].(map[string]interface{})
			setting := tkeGetInstanceAdvancedPara(workConfigPara, meta)
			request.InstanceAdvancedSettings = &setting
		}
	}

	request.InstanceAdvancedSettings.Labels = GetTkeLabels(d, "labels")
	if hostName, ok := d.GetOk("hostname"); ok {
		hostNameStr := hostName.(string)
		request.HostName = &hostNameStr
	}

	if v, ok := d.GetOk("unschedulable"); ok {
		request.InstanceAdvancedSettings.Unschedulable = helper.Int64(int64(v.(int)))
	}

	//if workConfigOverrides, ok := d.GetOk("worker_config_overrides"); ok {
	//	workConfigOverrideList := workConfigOverrides.([]interface{})
	//	request.InstanceAdvancedSettingsOverrides = make([]*tke.InstanceAdvancedSettings, 0, len(workConfigOverrideList))
	//	for _, conf := range workConfigOverrideList {
	//		workConfigPara := conf.(map[string]interface{})
	//		setting := tkeGetInstanceAdvancedPara(workConfigPara, meta)
	//		request.InstanceAdvancedSettingsOverrides = append(request.InstanceAdvancedSettingsOverrides, &setting)
	//	}
	//}

	/*cvm has been  attached*/
	var err error
	_, workers, err := tkeService.DescribeClusterInstances(ctx, *request.ClusterId)
	if err != nil {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			_, workers, err = tkeService.DescribeClusterInstances(ctx, *request.ClusterId)
			if err != nil {
				return retryError(err, InternalError)
			}
			return nil
		})
	}
	if err != nil {
		return err
	}

	has := false
	for _, worker := range workers {
		if worker.InstanceId == *instanceId {
			has = true
		}
	}
	if has {
		return fmt.Errorf("instance %s has been attached to cluster %s,can not attach again", *instanceId, *request.ClusterId)
	}

	var response *tke.AddExistedInstancesResponse

	if err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, err = tkeService.client.UseTkeClient().AddExistedInstances(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("add existed instance %s to cluster %s error,reason %v", *instanceId, *request.ClusterId, err)
	}
	var success = false
	for _, v := range response.Response.SuccInstanceIds {
		if *v == *instanceId {
			d.SetId(*instanceId + "_" + *request.ClusterId)
			success = true
		}
	}

	if !success {
		return fmt.Errorf("add existed instance %s to cluster %s error, instance not in success instanceIds", *instanceId, *request.ClusterId)
	}

	/*wait for cvm status*/
	if err = resource.Retry(7*readRetryTimeout, func() *resource.RetryError {
		instance, errRet := cvmService.DescribeInstanceById(ctx, *instanceId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if instance != nil && *instance.InstanceState == CVM_STATUS_RUNNING {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("cvm instance %s status is %s, retry...", *instanceId, *instance.InstanceState))
	}); err != nil {
		return err
	}

	/*wait for tke init ok */
	err = resource.Retry(7*readRetryTimeout, func() *resource.RetryError {
		_, workers, err = tkeService.DescribeClusterInstances(ctx, *request.ClusterId)
		if err != nil {
			return retryError(err, InternalError)
		}
		has := false
		for _, worker := range workers {
			if worker.InstanceId == *instanceId {
				has = true
				if worker.InstanceState == "failed" {
					return resource.NonRetryableError(fmt.Errorf("cvm instance %s attach to cluster %s fail,reason:%s",
						*instanceId, *request.ClusterId, worker.FailedReason))
				}

				if worker.InstanceState != "running" {
					return resource.RetryableError(fmt.Errorf("cvm instance  %s in tke status is %s, retry...",
						*instanceId, worker.InstanceState))
				}

			}
		}
		if !has {
			return resource.NonRetryableError(fmt.Errorf("cvm instance %s not exist in tke instance list", *instanceId))
		}
		return nil
	})

	if err != nil {
		return err
	}

	return resourceTencentCloudTkeClusterAttachmentRead(d, meta)
}

func resourceTencentCloudTkeClusterAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tke_kubernetes_cluster_attachment.delete")()

	tkeService := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	instanceId, clusterId := "", ""

	if items := strings.Split(d.Id(), "_"); len(items) != 2 {
		return fmt.Errorf("the resource id is corrupted")
	} else {
		instanceId, clusterId = items[0], items[1]
	}

	request := tke.NewDeleteClusterInstancesRequest()

	request.ClusterId = &clusterId
	request.InstanceIds = []*string{
		&instanceId,
	}
	request.InstanceDeleteMode = helper.String("retain")

	var err error

	if err = resource.Retry(4*writeRetryTimeout, func() *resource.RetryError {
		_, err := tkeService.client.UseTkeClient().DeleteClusterInstances(request)
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
	}); err != nil {
		return err
	}
	return nil
}
