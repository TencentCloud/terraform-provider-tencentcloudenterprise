/*
Manages a CKafka instance.

# Example Usage

```hcl
variable "zone_id" { type = number }
variable "subnet_id" { type = string }
variable "vpc_id" { type = string }

	resource "tencentcloudenterprise_ckafka_instance" "example" {
	  instance_name = "tf-example"
	  zone_id = var.zone_id
	  band_width = 20
	  disk_size = 100
	  subnet_id = var.subnet_id
	  vpc_id = var.vpc_id
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ckafka_instance", CNDescription{
		TerraformTypeCN: "ckafka实例",
		AttributesCN: map[string]string{
			"instance_name":      "实例名称",
			"zone_id":            "可用区域ID",
			"vpc_id":             "vpcId",
			"subnet_id":          "子网id",
			"msg_retention_time": "日志保存时间",
			"renew_flag":         "实例是否续费，int  枚举值：1表示自动续费，2表示明确不自动续费",
			"band_width":         "实例带宽，单位Mbps",
			"disk_size":          "实例的存储大小，单位GB",
			"tag_set":            "标识tag",
			"vip":                "访问实例的vip 信息",
			"vport":              "访问实例的端口信息",
			"instance_id":        "实例id",
		},
	})
}

func resourceTencentCloudCkafkaInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Use this resource to create ckafka instance.",
		Create:      resourceTencentCloudCkafkaInstanceCreate,
		Read:        resourceTencentCloudCkafkaInstanceRead,
		Update:      resourceTencentCloudCkafkaInstanceUpdate,
		Delete:      resourceTencentCLoudCkafkaInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance name.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Ckafka instance ID.",
			},
			"zone_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Available zone id.",
				//DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				//	multiZone := d.Get("multi_zone_flag").(bool)
				//	zoneId := d.Get("zone_id").(int)
				//	v, ok := d.GetOk("zone_ids")
				//
				//	if !multiZone || !ok || old == "" {
				//		return old == new
				//	}
				//	zoneIds := v.(*schema.Set)
				//	return zoneIds.Contains(zoneId)
				//},
			},
			"cluster_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: "Cluster-ID represents  the cluster to which Ckafa belongs, " +
					"associated with cluster_name only needs to select one of them.",
			},
			//"specifications_type": {
			//	Type:         schema.TypeString,
			//	Optional:     true,
			//	Default:      "profession",
			//	ValidateFunc: validateAllowedStringValue([]string{"standard", "profession"}),
			//	Description:  "Specifications type of instance. Allowed values are `standard`, `profession`. Default is `profession`.",
			//},
			//"period": {
			//	Type:        schema.TypeInt,
			//	Optional:    true,
			//	Description: "Prepaid purchase time, such as 1, is one month.",
			//},
			//"instance_type": {
			//	Type:         schema.TypeInt,
			//	Optional:     true,
			//	Computed:     true,
			//	ValidateFunc: validateAllowedIntValue([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			//	Description:  "Description of instance type. `profession`: 1, `standard`:  1(general), 2(standard), 3(advanced), 4(capacity), 5(specialized-1), 6(specialized-2), 7(specialized-3), 8(specialized-4), 9(exclusive).",
			//},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vpc id",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subnet id",
			},
			"msg_retention_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "The maximum retention time of instance logs, in minutes." +
					" the default is 10080 (7 days), the maximum is 30 days, and the default 0 is not filled," +
					" which means that the log retention time recovery policy is not enabled.",
			},
			"renew_flag": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "Prepaid automatic renewal mark, 0 means the default state, the initial state," +
					" 1 means automatic renewal, 2 means clear no automatic renewal (user setting).",
			},
			//"kafka_version": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "Kafka version (0.10.2/1.1.1/2.4.1).",
			//},
			"band_width": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Instance bandwidth as listed by the CKafka type configuration.",
			},
			"disk_size": {
				Type:     schema.TypeInt,
				Required: true,
				Description: "Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. " +
					"If it is not within the interval, the plan will cause a change when first created.",
			},
			//"partition": {
			//	Type:     schema.TypeInt,
			//	Optional: true,
			//	Computed: true,
			//	Description: "Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. " +
			//		"If it is not within the interval, the plan will cause a change when first created.",
			//},
			//"multi_zone_flag": {
			//	Type:        schema.TypeBool,
			//	Optional:    true,
			//	Description: "Indicates whether the instance is multi zones. NOTE: if set to `true`, `zone_ids` must set together.",
			//},
			//"zone_ids": {
			//	Type:         schema.TypeSet,
			//	Optional:     true,
			//	Description:  "List of available zone id. NOTE: this argument must set together with `multi_zone_flag`.",
			//	RequiredWith: []string{"multi_zone_flag"},
			//	Elem:         &schema.Schema{Type: schema.TypeInt},
			//},
			//"tags": {
			//	Type:          schema.TypeList,
			//	Optional:      true,
			//	Computed:      true,
			//	Deprecated:    "It has been deprecated from version 1.78.5, because it do not support change. Use `tag_set` instead.",
			//	ConflictsWith: []string{"tag_set"},
			//	Description:   "Tags of instance. Partition size, the professional version does not need tag.",
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"key": {
			//				Type:        schema.TypeString,
			//				Required:    true,
			//				Description: "Tag key.",
			//			},
			//			"value": {
			//				Type:        schema.TypeString,
			//				Required:    true,
			//				Description: "Tag value.",
			//			},
			//		},
			//	},
			//},
			"tag_set": {
				Type:        schema.TypeMap,
				Optional:    true,
				Computed:    true,
				Description: "Tag set of instance.",
			},
			//"disk_type": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "Type of disk.",
			//},
			//"config": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	MaxItems: 1,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"auto_create_topic_enable": {
			//				Type:        schema.TypeBool,
			//				Required:    true,
			//				Description: "Automatic creation. true: enabled, false: not enabled.",
			//			},
			//			"default_num_partitions": {
			//				Type:     schema.TypeInt,
			//				Required: true,
			//				Description: "If auto.create.topic.enable is set to true and this value is not set, " +
			//					"3 will be used by default.",
			//			},
			//			"default_replication_factor": {
			//				Type:     schema.TypeInt,
			//				Required: true,
			//				Description: "If auto.create.topic.enable is set to true but this value is not set, " +
			//					"2 will be used by default.",
			//			},
			//		},
			//	},
			//	Description: "Instance configuration.",
			//},
			//"dynamic_retention_config": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	MaxItems: 1,
			//	Computed: true,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"enable": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//				Description: "Whether the dynamic message retention time configuration is " +
			//					"enabled. 0: disabled; 1: enabled.",
			//			},
			//			"disk_quota_percentage": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//				Description: "Disk quota threshold (in percentage) for triggering " +
			//					"the message retention time change event.",
			//			},
			//			"step_forward_percentage": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//				Description: "Percentage by which the message retention " +
			//					"time is shortened each time.",
			//			},
			//			"bottom_retention": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Computed:    true,
			//				Description: "Minimum retention time, in minutes.",
			//			},
			//		},
			//	},
			//	Description: "Dynamic message retention policy configuration.",
			//},
			//"rebalance_time": {
			//	Type:        schema.TypeInt,
			//	Optional:    true,
			//	Description: "Modification of the rebalancing time after upgrade.",
			//},
			//"public_network": {
			//	Type:        schema.TypeInt,
			//	Optional:    true,
			//	Computed:    true,
			//	Deprecated:  "It has been deprecated from version 1.81.6. If set public network value, it will cause error.",
			//	Description: "Bandwidth of the public network.",
			//},
			//"dynamic_disk_config": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	MaxItems: 1,
			//	Computed: true,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"enable": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//				Description: "Whether to the dynamic disk expansion configuration is enabled." +
			//					"0: disabled; 1: enabled.",
			//			},
			//			"disk_quota_percentage": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Computed:    true,
			//				Description: "Disk quota threshold (in percentage) for triggering the automatic disk expansion event.",
			//			},
			//			"step_forward_percentage": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Computed:    true,
			//				Description: "Percentage of dynamic disk expansion each time.",
			//			},
			//			"max_disk_space": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Computed:    true,
			//				Description: "Max scale disk size, in GB.",
			//			},
			//		},
			//	},
			//	Description: "Dynamic disk expansion policy configuration.",
			//},
			//"max_message_byte": {
			//	Type:         schema.TypeInt,
			//	Optional:     true,
			//	Computed:     true,
			//	ValidateFunc: validateIntegerInRange(1024, 12*1024*1024),
			//	Description:  "The size of a single message in bytes at the instance level. Value range: `1024 - 12*1024*1024 bytes (i.e., 1KB-12MB).",
			//},
			"vip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Vip of instance.",
			},
			"vport": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of instance.",
			},
		},
	}
}

func resourceTencentCloudCkafkaInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ckafka_instance.create")()
	var (
		logId   = getLogId(contextNil)
		service = CkafkaService{
			client: meta.(*TencentCloudClient).apiV3Conn,
		}
		request  = ckafka.NewCreateInstanceTceRequest()
		response = ckafka.NewCreateInstanceTceResponse()
		ctx      = context.WithValue(context.TODO(), logIdKey, logId)
	)
	instanceName := d.Get("instance_name").(string)
	request.InstanceName = &instanceName

	zoneId := int64(d.Get("zone_id").(int))
	request.ZoneId = &zoneId

	if v, ok := d.GetOk("vpc_id"); ok {
		vpcId := v.(string)
		request.VpcId = helper.String(vpcId)
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		subnetId := v.(string)
		request.SubnetId = helper.String(subnetId)
	}

	if v, ok := d.GetOk("disk_size"); ok {
		diskSize := int64(v.(int))
		request.DiskSize = helper.Int64(diskSize)
	}

	if v, ok := d.GetOk("band_width"); ok {
		bandWidth := int64(v.(int))
		request.BandWidth = helper.Int64(bandWidth)
	}
	if v, ok := d.GetOk("cluster_id"); ok {
		clusterId := v.(int)
		request.ClusterId = helper.Int64(int64(clusterId))
	}

	if v, ok := d.GetOk("msg_retention_time"); ok {
		retentionTime := int64(v.(int))
		request.MsgRetentionTime = helper.Int64(retentionTime)
	}

	result, err := service.client.UseCkafkaClient().CreateInstanceTce(request)
	response = result

	if err != nil {
		log.Printf("[CRITAL]%s create ckafka instance failed, reason:%s\n", logId, err.Error())
		return err
	}

	if response == nil || response.Response == nil || response.Response.Result == nil || len(response.Response.Result.ResourceIds) != 1 || response.Response.Result.ResourceIds[0] == nil || *response.Response.Result.ResourceIds[0] == "" {
		return fmt.Errorf("CKafka create response is missing the instance ID")
	}
	instanceId := response.Response.Result.ResourceIds[0]
	d.SetId(*instanceId)

	err = resource.Retry(5*readRetryTimeout, func() *resource.RetryError {
		has, ready, err := service.CheckCkafkaInstanceReady(ctx, *instanceId)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if !has {
			return resource.RetryableError(fmt.Errorf("CKafka instance %s is not visible yet", *instanceId))
		}
		if ready {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("create ckafka instance task is processing"))
	})
	if err != nil {
		return err
	}

	//if v, ok := d.GetOk("dynamic_disk_config"); ok {
	//	needModify = true
	//	dynamic := make([]*ckafka.DynamicDiskConfig, 0, 10)
	//	for _, item := range v.([]interface{}) {
	//		dMap := item.(map[string]interface{})
	//		dynamicInfo := ckafka.DynamicDiskConfig{}
	//		if enable, ok := dMap["enable"]; ok {
	//			dynamicInfo.Enable = helper.Int64(int64(enable.(int)))
	//		}
	//		if stepForwardPercentage, ok := dMap["step_forward_percentage"]; ok {
	//			dynamicInfo.StepForwardPercentage = helper.Int64(int64(stepForwardPercentage.(int)))
	//		}
	//		if diskQuotaPercentage, ok := dMap["disk_quota_percentage"]; ok {
	//			dynamicInfo.DiskQuotaPercentage = helper.Int64(int64(diskQuotaPercentage.(int)))
	//		}
	//		if maxDiskSpace, ok := dMap["max_disk_space"]; ok {
	//			dynamicInfo.MaxDiskSpace = helper.Int64(int64(maxDiskSpace.(int)))
	//		}
	//		dynamic = append(dynamic, &dynamicInfo)
	//	}
	//	modifyRequest.DynamicDiskConfig = dynamic[0]
	//}

	client := meta.(*TencentCloudClient).apiV3Conn
	tagService := TagService{client: client}
	region := client.Region

	if tags := helper.GetTags(d, "tag_set"); len(tags) > 0 {
		resourceName := BuildTagResourceName("ckafka", "ckafkaId", region, *instanceId)
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudCkafkaInstanceRead(d, meta)
}

func resourceTencentCloudCkafkaInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ckafka_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var service = CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	instanceId := d.Id()

	var info *ckafka.InstanceDetail
	var isExist = true

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		res, has, e := service.DescribeCkafkaInstanceById(ctx, instanceId)
		if e != nil {
			return retryError(e)
		}
		if !has {
			d.SetId("")
			isExist = false
			return nil
		}
		info = res
		return nil
	})
	if err != nil {
		return fmt.Errorf("[API]Describe kafka instance fail, reason:%s", err.Error())
	}

	if !isExist {
		return nil
	}

	fillZone := func() error {
		v, ok := d.GetOk("zone_id")
		if !ok {
			// During terraform import, zone_id is not set in state yet; infer it from API response.
			if info.ZoneId != nil && *info.ZoneId != 0 {
				return d.Set("zone_id", *info.ZoneId)
			}
			for _, zoneId := range info.ZoneIds {
				if zoneId != nil {
					return d.Set("zone_id", *zoneId)
				}
			}
			return fmt.Errorf("zone_id is missing from CKafka instance %s detail", instanceId)
		}
		if len(info.ZoneIds) == 0 && info.ZoneId != nil && *info.ZoneId == int64(v.(int)) {
			return d.Set("zone_id", *info.ZoneId)
		}
		for _, zoneId := range info.ZoneIds {
			if zoneId != nil && *zoneId == int64(v.(int)) {
				_ = d.Set("zone_id", *zoneId)
				return nil
			}
		}
		return fmt.Errorf("zone_id %d not found", v.(int))
	}
	_ = d.Set("instance_name", info.InstanceName)

	if err = fillZone(); err != nil {
		return err
	}
	// calculate period
	//createTime := *info.CreateTime
	//expireTime := *info.ExpireTime
	//period := (expireTime - createTime) / (3600 * 24 * 31)
	//_ = d.Set("period", &period)
	_ = d.Set("vpc_id", info.VpcId)
	_ = d.Set("subnet_id", info.SubnetId)
	_ = d.Set("renew_flag", info.RenewFlag)
	//_ = d.Set("kafka_version", info.Version)
	_ = d.Set("disk_size", info.DiskSize)
	bandWidth := info.Bandwidth
	_ = d.Set("vip", info.Vip)
	_ = d.Set("vport", info.Vport)
	if bandWidth == nil {
		return fmt.Errorf("CKafka instance %s detail is missing the bandwidth", instanceId)
	}
	_ = d.Set("band_width", *bandWidth)

	//if len(info.ZoneIds) > 1 {
	//	_ = d.Set("multi_zone_flag", true)
	//	ids := helper.Int64sInterfaces(info.ZoneIds)
	//	idSet := schema.NewSet(func(i interface{}) int {
	//		return i.(int)
	//	}, ids)
	//	_ = d.Set("zone_ids", idSet)
	//}

	//tagSets := make([]map[string]interface{}, 0, len(info.Tags))
	//for _, item := range info.Tags {
	//	tagSets = append(tagSets, map[string]interface{}{
	//		"key":   item.TagKey,
	//		"value": item.TagValue,
	//	})
	//}
	//_ = d.Set("tags", tagSets)
	_ = d.Set("instance_id", instanceId)

	client := meta.(*TencentCloudClient).apiV3Conn
	tagService := TagService{client: client}
	region := client.Region

	tags, err := tagService.DescribeResourceTags(ctx, "ckafka", "ckafkaId", region, instanceId)
	if err != nil {
		return err
	}
	_ = d.Set("tag_set", tags)

	// query msg_retention_time
	var (
		request  = ckafka.NewDescribeInstanceAttributesRequest()
		response = ckafka.NewDescribeInstanceAttributesResponse()
	)
	request.InstanceId = &instanceId
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.client.UseCkafkaClient().DescribeInstanceAttributes(request)
		if e != nil {
			return retryError(e)
		}
		response = result
		if response != nil && response.Response != nil && response.Response.Result != nil {
			attr := response.Response.Result
			_ = d.Set("msg_retention_time", attr.MsgRetentionTime)
		}

		//if attr.Config != nil {
		//	config := make([]map[string]interface{}, 0)
		//	config = append(config, map[string]interface{}{
		//		"auto_create_topic_enable":   attr.Config.AutoCreateTopicsEnable,
		//		"default_num_partitions":     attr.Config.DefaultNumPartitions,
		//		"default_replication_factor": attr.Config.DefaultReplicationFactor,
		//	})
		//	_ = d.Set("config", config)
		//}

		//dynamicDiskConfig := make([]map[string]interface{}, 0)
		//dynamicDiskConfig = append(dynamicDiskConfig, map[string]interface{}{
		//	"enable":                  attr.DynamicDiskConfig.Enable,
		//	"disk_quota_percentage":   attr.DynamicDiskConfig.DiskQuotaPercentage,
		//	"step_forward_percentage": attr.DynamicDiskConfig.StepForwardPercentage,
		//	"max_disk_space":          attr.DynamicDiskConfig.MaxDiskSpace,
		//})
		//_ = d.Set("dynamic_disk_config", dynamicDiskConfig)

		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Ckafka Instance Attributes failed, reason:%s\n", logId, err.Error())
		return err
	}
	return nil
}

func resourceTencentCloudCkafkaInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ckafka_instance.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	immutableArgs := []string{
		"period", "zone_id", "vpc_id",
		"subnet_id", "renew_flag", "kafka_version",
		"multi_zone_flag", "zone_ids", "disk_type",
		"specifications_type", "instance_type", "cluster_id",
		"msg_retention_time",
	}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	instanceId := d.Id()
	request := ckafka.NewModifyResourceTceRequest()
	request.InstanceId = &instanceId
	//if d.HasChange("instance_name") {
	//	if v, ok := d.GetOk("instance_name"); ok {
	//		request.InstanceName = helper.String(v.(string))
	//	}
	//}
	//
	//if d.HasChange("msg_retention_time") {
	//	if v, ok := d.GetOk("msg_retention_time"); ok {
	//		request.MsgRetentionTime = helper.Int64(int64(v.(int)))
	//	}
	//}
	//
	//if d.HasChange("config") {
	//	if v, ok := d.GetOk("config"); ok {
	//		items := v.([]interface{})
	//		dMap := items[0].(map[string]interface{})
	//		configInfo := ckafka.ModifyInstanceAttributesConfig{}
	//		if autoCreateTopicEnable, ok := dMap["auto_create_topic_enable"]; ok {
	//			configInfo.AutoCreateTopicEnable = helper.Bool(autoCreateTopicEnable.(bool))
	//		}
	//		if defaultNumPartitions, ok := dMap["default_num_partitions"]; ok {
	//			configInfo.DefaultNumPartitions = helper.Int64(int64(defaultNumPartitions.(int)))
	//		}
	//		if defaultReplicationFactor, ok := dMap["default_replication_factor"]; ok {
	//			configInfo.DefaultReplicationFactor = helper.Int64(int64(defaultReplicationFactor.(int)))
	//		}
	//		request.Config = &configInfo
	//	}
	//}
	//
	//if d.HasChange("dynamic_retention_config") {
	//	if v, ok := d.GetOk("dynamic_retention_config"); ok {
	//		items := v.([]interface{})
	//		dMap := items[0].(map[string]interface{})
	//		dynamicInfo := ckafka.DynamicRetentionTime{}
	//		if enable, ok := dMap["enable"]; ok {
	//			dynamicInfo.Enable = helper.Int64(int64(enable.(int)))
	//		}
	//		if diskQuotaPercentage, ok := dMap["disk_quota_percentage"]; ok {
	//			dynamicInfo.DiskQuotaPercentage = helper.Int64(int64(diskQuotaPercentage.(int)))
	//		}
	//		if stepForwardPercentage, ok := dMap["step_forward_percentage"]; ok {
	//			dynamicInfo.StepForwardPercentage = helper.Int64(int64(stepForwardPercentage.(int)))
	//		}
	//		if bottomRetention, ok := dMap["bottom_retention"]; ok {
	//			dynamicInfo.BottomRetention = helper.Int64(int64(bottomRetention.(int)))
	//		}
	//		request.DynamicRetentionConfig = &dynamicInfo
	//	}
	//}
	//
	//if d.HasChange("rebalance_time") {
	//	if v, ok := d.GetOk("rebalance_time"); ok {
	//		request.RebalanceTime = helper.Int64(int64(v.(int)))
	//	}
	//}
	//
	//if d.HasChange("public_network") {
	//	if v, ok := d.GetOk("public_network"); ok {
	//		request.PublicNetwork = helper.Int64(int64(v.(int)))
	//	}
	//}

	//if d.HasChange("dynamic_disk_config") {
	//	if v, ok := d.GetOk("dynamic_disk_config"); ok {
	//		dynamic := make([]*ckafka.DynamicDiskConfig, 0, 10)
	//		for _, item := range v.([]interface{}) {
	//			dMap := item.(map[string]interface{})
	//			dynamicInfo := ckafka.DynamicDiskConfig{}
	//			if enable, ok := dMap["enable"]; ok {
	//				dynamicInfo.Enable = helper.Int64(int64(enable.(int)))
	//			}
	//			if stepForwardPercentage, ok := dMap["step_forward_percentage"]; ok {
	//				dynamicInfo.StepForwardPercentage = helper.Int64(int64(stepForwardPercentage.(int)))
	//			}
	//			if diskQuotaPercentage, ok := dMap["disk_quota_percentage"]; ok {
	//				dynamicInfo.DiskQuotaPercentage = helper.Int64(int64(diskQuotaPercentage.(int)))
	//			}
	//			if maxDiskSpace, ok := dMap["max_disk_space"]; ok {
	//				dynamicInfo.MaxDiskSpace = helper.Int64(int64(maxDiskSpace.(int)))
	//			}
	//			dynamic = append(dynamic, &dynamicInfo)
	//		}
	//		request.DynamicDiskConfig = dynamic[0]
	//	}
	//}

	//if d.HasChange("max_message_byte") {
	//	if v, ok := d.GetOkExists("max_message_byte"); ok {
	//		request.MaxMessageByte = helper.Uint64(uint64(v.(int)))
	//	}
	//}
	//
	//err := service.ModifyCkafkaInstanceAttributes(ctx, request)
	//if err != nil {
	//	return fmt.Errorf("[API]Set kafka instance attributes fail, reason:%s", err.Error())
	//}
	//
	if d.HasChange("band_width") || d.HasChange("disk_size") || d.HasChange("partition") {
		request.InstanceId = helper.String(instanceId)
		if v, ok := d.GetOk("band_width"); ok {
			request.Bandwidth = helper.Int64(int64(v.(int)))
		}
		if v, ok := d.GetOk("disk_size"); ok {
			request.DiskSize = helper.Int64(int64(v.(int)))
		}
		//if v, ok := d.GetOk("partition"); ok {
		//	request.Partition = helper.Int64(int64(v.(int)))
		//}
		//
		_, err := service.client.UseCkafkaClient().ModifyResourceTce(request)
		if err != nil {
			return fmt.Errorf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]", logId,
				request.GetAction(), request.ToJsonString(), err.Error())
		}

		err = resource.Retry(5*readRetryTimeout, func() *resource.RetryError {
			_, ready, err := service.CheckCkafkaInstanceReady(ctx, instanceId)
			if err != nil {
				return resource.NonRetryableError(err)
			}
			if ready {
				return nil
			}
			return resource.RetryableError(fmt.Errorf("upgrade ckafka instance task is processing"))
		})
		if err != nil {
			return err
		}
	}
	//
	if d.HasChange("tag_set") {

		client := meta.(*TencentCloudClient).apiV3Conn
		tagService := TagService{client: client}
		region := client.Region

		oldTags, newTags := d.GetChange("tag_set")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))

		resourceName := BuildTagResourceName("ckafka", "ckafkaId", region, instanceId)
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudCkafkaInstanceRead(d, meta)
}

func resourceTencentCLoudCkafkaInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ckafka_instance.delete")()
	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = CkafkaService{
			client: meta.(*TencentCloudClient).apiV3Conn,
		}
		request = ckafka.NewDeleteInstanceRequest()
	)
	instanceId := d.Id()
	request.InstanceIds = []*string{helper.String(instanceId)}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, err := service.client.UseCkafkaClient().DeleteInstance(request)
		if err != nil {
			if isExpectError(err, []string{CkafkaInstanceNotFound}) {
				return nil
			}
			return retryError(err, "UnsupportedOperation")
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = resource.Retry(5*readRetryTimeout, func() *resource.RetryError {
		has, _, err := service.CheckCkafkaInstanceReady(ctx, instanceId)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if !has {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("delete ckafka instance task is processing"))
	})
	if err != nil {
		return err
	}
	return nil
}
