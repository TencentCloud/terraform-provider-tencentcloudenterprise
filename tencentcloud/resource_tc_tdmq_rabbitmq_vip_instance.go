/*
Provides a resource to create and manage TDMQ RabbitMQ VIP instance

Example Usage

### Create a basic RabbitMQ VIP instance with single node

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "example" {
  cluster_name = "rabbitmq-cluster"
  zone_ids     = ["ap-chongqing-1"]
  vpc_id       = "vpc-xxxxxxxx"
  subnet_id    = "subnet-xxxxxxxx"
  node_spec    = "rabbit-vip-basic-1"
  node_num     = 3
  storage_size = 200
  cluster_version = "3.8.30"
}
```

### Create a high-availability RabbitMQ VIP instance with multi-zone deployment

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "ha_instance" {
  cluster_name = "rabbitmq-ha-cluster"
  zone_ids     = ["ap-chongqing-1", "ap-chongqing-2", "ap-chongqing-3"]
  vpc_id       = "vpc-xxxxxxxx"
  subnet_id    = "subnet-xxxxxxxx"
  node_spec    = "rabbit-vip-basic-2"
  node_num     = 3
  storage_size = 500
  enable_create_default_ha_mirror_queue = true
  cluster_version = "3.11.8"
}
```

### Create a production RabbitMQ VIP instance with enhanced resources

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "production" {
  cluster_name = "rabbitmq-prod"
  zone_ids     = ["ap-chongqing-1"]
  vpc_id       = "vpc-xxxxxxxx"
  subnet_id    = "subnet-xxxxxxxx"
  node_spec    = "rabbit-vip-basic-4"
  node_num     = 3
  storage_size = 1000
  cluster_version = "3.11.8"
}
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_rabbitmq_vip_instance", CNDescription{
		TerraformTypeCN: "TDMQ RabbitMQ实例",
		DescriptionCN:   "提供TDMQ RabbitMQ实例资源，用于创建和管理TDMQ RabbitMQ专享实例。",
		AttributesCN: map[string]string{
			"zone_ids":                              "可用区ID列表,支持多可用区部署",
			"vpc_id":                                "私有网络VPC ID",
			"subnet_id":                             "私有网络子网ID",
			"cluster_name":                          "集群名称,3-64个字符,只能包含字母、数字、\"-\" 及 \"_\"",
			"node_spec":                             "节点规格,可选值:rabbit-vip-basic-1(4C8G)、rabbit-vip-basic-2(8C16G)、rabbit-vip-basic-4(16C32G)等",
			"node_num":                              "节点数量,必填且必须大于0,单可用区通常配置1个节点,多可用区至少需要3个节点以保证高可用,请根据实际环境配置",
			"storage_size":                          "单节点存储容量,默认200GB",
			"enable_create_default_ha_mirror_queue": "是否创建默认的HA镜像队列,一般为true，请根据需要配置",
			"auto_renew_flag":                       "自动续费标志,默认为true",
			"time_span":                             "购买时长,默认为1(月)",
			"pay_mode":                              "付费模式,0表示后付费,1表示预付费,默认为预付费",
			"cluster_version":                       "集群版本,支持3.8.30和3.11.8,默认为3.8.30",
			"public_access_endpoint":                "公网接入点地址",
			"vpcs":                                  "VPC接入点列表",
			"vpc_endpoint":                          "VPC访问端点地址",
			"vpc_data_stream_endpoint_status":       "VPC端点状态",
		},
	})
}

func resourceTencentCloudTdmqRabbitmqVipInstance() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTdmqRabbitmqVipInstanceCreate,
		Read:        resourceTencentCloudTdmqRabbitmqVipInstanceRead,
		Update:      resourceTencentCloudTdmqRabbitmqVipInstanceUpdate,
		Delete:      resourceTencentCloudTdmqRabbitmqVipInstanceDelete,
		Description: "Provides a resource to create and manage TDMQ RabbitMQ VIP instance",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_ids": {
				Required:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Availability zone ID list. For single availability zone deployment, provide one zone ID; for multi-availability zone deployment, provide multiple zone IDs. Multi-availability zone instances require at least 3 nodes.",
			},
			"vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "VPC (Virtual Private Cloud) ID where the RabbitMQ instance will be deployed. Format: vpc-xxxxxxxx.",
			},
			"subnet_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Subnet ID within the specified VPC where the instance will be deployed. Format: subnet-xxxxxxxx.",
			},
			"cluster_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "RabbitMQ cluster name. Length must be between 3-64 characters. Only letters, numbers, hyphens (-), and underscores (_) are allowed.",
			},
			"node_spec": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Node specification. Valid values: `rabbit-vip-basic-5` (2C4G), `rabbit-vip-profession-2c8g` (2C8G), `rabbit-vip-basic-1` (4C8G, default), `rabbit-vip-profession-4c16g` (4C16G), `rabbit-vip-basic-2` (8C16G), `rabbit-vip-profession-8c32g` (8C32G), `rabbit-vip-basic-4` (16C32G), `rabbit-vip-profession-16c64g` (16C64G). Note: Some specifications may be unavailable due to stock limitations.",
			},
			"node_num": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateIntegerMin(1),
				Description:  "Number of nodes in the cluster. Must be greater than 0. For single availability zone deployment, typically use 1 node; for multi-availability zone deployment, minimum 3 nodes are required for high availability. Please configure based on your actual environment and requirements.",
			},
			"storage_size": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Storage capacity per node in GB. Default is 200GB.",
			},
			"enable_create_default_ha_mirror_queue": {
				Required:    true,
				Type:        schema.TypeBool,
				Description: "Whether to create a default HA (High Availability) mirrored queue. When enabled, queues will be automatically mirrored across nodes for high availability. Default is true.",
			},
			"cluster_version": {
				Optional:    true,
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RabbitMQ cluster version. Valid values: `3.8.30` (default), `3.11.8`. Different versions may have different features and performance characteristics.",
			},
			"public_access_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public network access endpoint address. Used to access the RabbitMQ instance from the internet.",
			},
			"vpcs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of VPC access points. Contains VPC network endpoint information for accessing the RabbitMQ instance from within VPC.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC ID where the access endpoint is located.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subnet ID where the access endpoint is located.",
						},
						"vpc_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC private network access endpoint address. Use this address to connect to RabbitMQ from within the VPC.",
						},
						"vpc_data_stream_endpoint_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the VPC endpoint. Indicates the availability status of the VPC access point.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudTdmqRabbitmqVipInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		request    = tdmq.NewCreateRabbitMQVipInstanceRequest()
		response   = tdmq.NewCreateRabbitMQVipInstanceResponse()
		instanceId string
	)

	if v, ok := d.GetOk("zone_ids"); ok {
		zoneIdsSet := v.(*schema.Set).List()
		for i := range zoneIdsSet {
			zoneIds := zoneIdsSet[i].(int)
			request.ZoneIds = append(request.ZoneIds, helper.IntInt64(zoneIds))
		}
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		request.SubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("cluster_name"); ok {
		request.ClusterName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("node_spec"); ok {
		request.NodeSpec = helper.String(v.(string))
	}

	if v, ok := d.GetOk("node_num"); ok {
		request.NodeNum = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("storage_size"); ok {
		request.StorageSize = helper.IntInt64(v.(int))
	}

	// Get enable_create_default_ha_mirror_queue with default value applied
	request.EnableCreateDefaultHaMirrorQueue = helper.Bool(d.Get("enable_create_default_ha_mirror_queue").(bool))

	if v, ok := d.GetOk("cluster_version"); ok {
		request.ClusterVersion = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().CreateRabbitMQVipInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmq rabbitmqVipInstance failed, reason:%+v", logId, err)
		return err
	}

	instanceId = *response.Response.InstanceId

	// wait
	paramMap := make(map[string]interface{})
	tmpSet := make([]*tdmq.Filter, 0)
	filter := tdmq.Filter{}
	filter.Name = helper.String("instanceIds")
	filter.Values = helper.Strings([]string{instanceId})
	tmpSet = append(tmpSet, &filter)
	paramMap["filters"] = tmpSet
	err = resource.Retry(readRetryTimeout*10, func() *resource.RetryError {
		result, e := service.DescribeTdmqRabbitmqVipInstanceByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		if result == nil {
			return resource.NonRetryableError(fmt.Errorf("resource `tencentcloudenterprise_tdmq_rabbitmq_vip_instance` %s does not exist", instanceId))
		}

		if len(result) != 1 {
			return resource.NonRetryableError(fmt.Errorf("resource `tencentcloudenterprise_tdmq_rabbitmq_vip_instance` %s id error", instanceId))
		}

		switch *result[0].Status {
		case RabbitMQVipInstanceRunning:
			return resource.RetryableError(fmt.Errorf("rabbitmq_vip_instance status is creating"))
		case RabbitMQVipInstanceSuccess:
			return nil
		default:
			return resource.NonRetryableError(fmt.Errorf("rabbitmq_vip_instance status illegal"))

		}
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmq rabbitmqVipInstance failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(instanceId)

	return resourceTencentCloudTdmqRabbitmqVipInstanceRead(d, meta)
}

func resourceTencentCloudTdmqRabbitmqVipInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId = d.Id()
	)

	rabbitmqVipInstanceResponse, err := service.DescribeTdmqRabbitmqVipInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}

	if rabbitmqVipInstanceResponse == nil || rabbitmqVipInstanceResponse.Response == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TdmqRabbitmqVipInstance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	rabbitmqVipInstance := rabbitmqVipInstanceResponse.Response

	if rabbitmqVipInstance.ClusterInfo.ZoneIds != nil {
		_ = d.Set("zone_ids", rabbitmqVipInstance.ClusterInfo.ZoneIds)
	}

	if rabbitmqVipInstance.ClusterInfo.Vpcs != nil {
		_ = d.Set("vpc_id", rabbitmqVipInstance.ClusterInfo.Vpcs[0].VpcId)
		_ = d.Set("subnet_id", rabbitmqVipInstance.ClusterInfo.Vpcs[0].SubnetId)
	}

	if rabbitmqVipInstance.ClusterSpecInfo.NodeCount != nil {
		_ = d.Set("node_num", rabbitmqVipInstance.ClusterSpecInfo.NodeCount)
	}

	if rabbitmqVipInstance.ClusterSpecInfo.MaxStorage != nil {
		_ = d.Set("storage_size", rabbitmqVipInstance.ClusterSpecInfo.MaxStorage)
	}

	if rabbitmqVipInstance.ClusterInfo.ClusterVersion != nil {
		_ = d.Set("cluster_version", rabbitmqVipInstance.ClusterInfo.ClusterVersion)
	}

	// Set enable_create_default_ha_mirror_queue from MirrorQueuePolicyFlag
	// MirrorQueuePolicyFlag: 1 = enabled, 0 = disabled
	if rabbitmqVipInstance.ClusterInfo.MirrorQueuePolicyFlag != nil {
		_ = d.Set("enable_create_default_ha_mirror_queue", *rabbitmqVipInstance.ClusterInfo.MirrorQueuePolicyFlag == 1)
	}

	paramMap := make(map[string]interface{})
	tmpSet := make([]*tdmq.Filter, 0)
	filter := tdmq.Filter{}
	filter.Name = helper.String("instanceIds")
	filter.Values = helper.Strings([]string{instanceId})
	tmpSet = append(tmpSet, &filter)
	paramMap["filters"] = tmpSet
	err = resource.Retry(readRetryTimeout*10, func() *resource.RetryError {
		result, e := service.DescribeTdmqRabbitmqVipInstanceByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		if result[0].SpecName != nil {
			_ = d.Set("node_spec", result[0].SpecName)
		}

		if result[0].InstanceName != nil {
			_ = d.Set("cluster_name", result[0].InstanceName)
		}

		if result[0].PublicAccessEndpoint != nil {
			_ = d.Set("public_access_endpoint", result[0].PublicAccessEndpoint)
		}

		if result[0].Vpcs != nil {
			tmpList := make([]map[string]interface{}, 0, len(result[0].Vpcs))
			for _, vpc := range result[0].Vpcs {
				vpcMap := map[string]interface{}{}
				if vpc.VpcId != nil {
					vpcMap["vpc_id"] = vpc.VpcId
				}
				if vpc.SubnetId != nil {
					vpcMap["subnet_id"] = vpc.SubnetId
				}
				if vpc.VpcEndpoint != nil {
					vpcMap["vpc_endpoint"] = vpc.VpcEndpoint
				}
				if vpc.VpcDataStreamEndpointStatus != nil {
					vpcMap["vpc_data_stream_endpoint_status"] = vpc.VpcDataStreamEndpointStatus
				}
				tmpList = append(tmpList, vpcMap)
			}
			_ = d.Set("vpcs", tmpList)
		}

		return nil
	})

	if err != nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TdmqRabbitmqVipInstance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	return nil
}

func resourceTencentCloudTdmqRabbitmqVipInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = tdmq.NewModifyRabbitMQVipInstanceRequest()
		instanceId = d.Id()
	)

	immutableArgs := []string{
		"zone_ids", "vpc_id", "subnet_id", "node_spec", "node_num", "storage_size",
		"enable_create_default_ha_mirror_queue", "cluster_version",
	}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	request.InstanceId = &instanceId

	if d.HasChange("cluster_name") {
		if v, ok := d.GetOk("cluster_name"); ok {
			request.ClusterName = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().ModifyRabbitMQVipInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update tdmq rabbitmqVipInstance failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTdmqRabbitmqVipInstanceRead(d, meta)
}

func resourceTencentCloudTdmqRabbitmqVipInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId = d.Id()
	)

	if err := service.DeleteTdmqRabbitmqVipInstanceById(ctx, instanceId); err != nil {
		return err
	}

	return nil
}
