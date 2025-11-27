package tencentcloud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_pulsar_topic", CNDescription{
		TerraformTypeCN: "TDMQ Pulsar主题",
		DescriptionCN:   "提供TDMQ Pulsar主题资源，用于创建和管理TDMQ Pulsar主题。",
		AttributesCN: map[string]string{
			"environ_id":        "环境ID",
			"topic_name":        "主题名称",
			"partitions":        "分区数",
			"topic_type":        "主题类型",
			"cluster_id":        "集群ID",
			"pulsar_topic_type": "Pulsar主题类型",
			"remark":            "主题描述",
			"create_time":       "创建时间",
		},
	})
}

func resourceTencentCloudTdmqPulsarTopic() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTdmqPulsarTopicCreate,
		Read:        resourceTencentCloudTdmqPulsarTopicRead,
		Update:      resourceTencentCloudTdmqPulsarTopicUpdate,
		Delete:      resourceTencentCloudTdmqPulsarTopicDelete,
		Description: "Provides a resource to create and manage TDMQ Pulsar topic",

		Schema: map[string]*schema.Schema{
			"environ_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of tdmq namespace.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of topic to be created.",
			},
			"partitions": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The partitions of topic.",
			},
			"topic_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Deprecated:  "This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue.",
				Description: "The type of topic.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Dedicated Cluster Id.",
			},
			"pulsar_topic_type": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"topic_type"},
				Description:   "Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the namespace.",
			},

			//compute
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of resource.",
			},
		},
	}
}

func resourceTencentCloudTdmqPulsarTopicCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_topic.create")()

	var (
		logId           = getLogId(contextNil)
		ctx             = context.WithValue(context.TODO(), logIdKey, logId)
		tdmqService     = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		environId       string
		topicName       string
		partitions      uint64
		topicType       int64
		remark          string
		clusterId       string
		pulsarTopicType int64
	)

	if temp, ok := d.GetOk("environ_id"); ok {
		environId = temp.(string)
		if len(environId) < 1 {
			return fmt.Errorf("environ_id should be not empty string")
		}
	}

	if temp, ok := d.GetOk("topic_name"); ok {
		topicName = temp.(string)
		if len(topicName) < 1 {
			return fmt.Errorf("topic_name should be not empty string")
		}
	}

	partitions = uint64(d.Get("partitions").(int))

	if temp, ok := d.GetOk("remark"); ok {
		remark = temp.(string)
	}

	if temp, ok := d.GetOk("cluster_id"); ok {
		clusterId = temp.(string)
	}

	if v, ok := d.GetOkExists("pulsar_topic_type"); ok {
		pulsarTopicType = int64(v.(int))
	} else {
		pulsarTopicType = NonePulsarTopicType
		if v, ok := d.GetOkExists("topic_type"); ok {
			topicType = int64(v.(int))
		} else {
			topicType = NoneTopicType
		}
	}

	err := tdmqService.CreateTdmqTopic(ctx, environId, topicName, partitions, topicType, remark, clusterId, pulsarTopicType)
	if err != nil {
		return err
	}

	d.SetId(topicName)
	return resourceTencentCloudTdmqPulsarTopicRead(d, meta)
}

func resourceTencentCloudTdmqPulsarTopicRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_instance.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		tdmqService = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	topicName := d.Id()
	environId := d.Get("environ_id").(string)
	clusterId := d.Get("cluster_id").(string)

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, has, e := tdmqService.DescribeTdmqTopicById(ctx, environId, topicName, clusterId)
		if e != nil {
			return retryError(e)
		}

		if !has {
			d.SetId("")
			return nil
		}

		_ = d.Set("partitions", info.Partitions)
		_ = d.Set("topic_type", info.TopicType)
		_ = d.Set("pulsar_topic_type", info.PulsarTopicType)
		_ = d.Set("remark", info.Remark)
		_ = d.Set("create_time", info.CreateTime)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceTencentCloudTdmqPulsarTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_topic.update")()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		partitions uint64
		remark     string
	)

	immutableArgs := []string{"topic_type", "pulsar_topic_type"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	topicName := d.Id()
	environId := d.Get("environ_id").(string)
	clusterId := d.Get("cluster_id").(string)

	old, now := d.GetChange("partitions")
	if d.HasChange("partitions") {
		partitions = uint64(now.(int))
	} else {
		partitions = uint64(old.(int))
	}

	old, now = d.GetChange("remark")
	if d.HasChange("remark") {
		remark = now.(string)
	} else {
		remark = old.(string)
	}

	d.Partial(true)

	if err := service.ModifyTdmqTopicAttribute(ctx, environId, topicName,
		partitions, remark, clusterId); err != nil {
		return err
	}

	d.Partial(false)
	return resourceTencentCloudTdmqPulsarTopicRead(d, meta)
}

func resourceTencentCloudTdmqPulsarTopicDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_instance.delete")()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	topicName := d.Id()
	environId := d.Get("environ_id").(string)
	clusterId := d.Get("cluster_id").(string)

	err := service.DeleteTdmqTopic(ctx, environId, topicName, clusterId)
	if err != nil {
		return err
	}
	return nil
}
