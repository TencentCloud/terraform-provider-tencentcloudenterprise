/*
Provide a resource to create a TDMQ topic.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tdmq_instance" "foo" {
	  cluster_name = "example"
	  remark = "this is description."
	}

	resource "tencentcloudenterprise_tdmq_namespace" "bar" {
	  environ_name = "example"
	  msg_ttl = 300
	  cluster_id = "tencentcloudenterprise_tdmq_instance.foo.id"
	  remark = "this is description."
	}

	resource "tencentcloudenterprise_tdmq_topic" "bar" {
	  environ_id = cloud_tdmq_namespace.bar.id
	  topic_name = "example"
	  partitions = 6
	  topic_type = 0
	  cluster_id = cloud_tdmq_instance.foo.id
	  remark = "this is description."
	}

```

# Import

Tdmq Topic can be imported, e.g.

```
$ terraform import cloud_tdmq_topic.test topic_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_topic", CNDescription{
		TerraformTypeCN: "TDMQ Topic",
		DescriptionCN:   "提供TDMQ Topic资源，用于创建和管理TDMQ Topic。",
		AttributesCN: map[string]string{
			"environ_id":  "TDMQ命名空间名称",
			"topic_name":  "主题名称",
			"partitions":  "分区数",
			"topic_type":  "主题类型",
			"remark":      "备注",
			"tags":        "标签",
			"cluster_id":  "Pulsar 集群的ID",
			"create_time": "创建时间",
		},
	})
}
func resourceTencentCloudTdmqTopic() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a TDMQ topic.",
		Create:      resourceTencentCloudTdmqTopicCreate,
		Read:        resourceTencentCloudTdmqTopicRead,
		Update:      resourceTencentCloudTdmqTopicUpdate,
		Delete:      resourceTencentCloudTdmqTopicDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// 环境（命名空间）名称
			"environ_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of tdmq namespace.",
			},
			// 主题名，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符
			"topic_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 64),
				ForceNew:     true,
				Description:  "The name of topic to be created.",
			},
			// 0：非分区topic，无分区；非0：具体分区topic的分区数，最大不允许超过128
			"partitions": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateIntegerInRange(0, 128),
				Description:  "The partitions of topic.",
			},
			// 备注，128字符以内
			"remark": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringLengthInRange(1, 128),
				Description:  "Description of the namespace.",
			},
			// 0： 普通消息；1 ：全局顺序消息；2 ：局部顺序消息；3 ：重试队列；4 ：死信队列
			"topic_type": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateIntegerInRange(0, 4),
				Description:  "The type of topic.",
			},
			// Pulsar 集群的ID
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Dedicated Cluster Id.",
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

func resourceTencentCloudTdmqTopicCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_topic.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	tdmqService := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
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

	pulsarTopicType = int64(d.Get("topic_type").(int))

	err := tdmqService.CreateTdmqTopic(ctx, environId, topicName, partitions, topicType, remark, clusterId, pulsarTopicType)
	if err != nil {
		return err
	}
	d.SetId(topicName)

	return resourceTencentCloudTdmqTopicRead(d, meta)
}

func resourceTencentCloudTdmqTopicRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	topicName := d.Id()
	environId := d.Get("environ_id").(string)
	clusterId := d.Get("cluster_id").(string)

	tdmqService := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
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
		_ = d.Set("remark", info.Remark)
		_ = d.Set("create_time", info.CreateTime)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func resourceTencentCloudTdmqTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_topic.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	if d.HasChange("topic_type") {
		return fmt.Errorf("`topic_type` do not support change now.")
	}

	topicName := d.Id()
	environId := d.Get("environ_id").(string)
	clusterId := d.Get("cluster_id").(string)

	service := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		partitions uint64
		remark     string
	)
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
	return resourceTencentCloudTdmqTopicRead(d, meta)
}

func resourceTencentCloudTdmqTopicDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_instance.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	topicName := d.Id()
	environId := d.Get("environ_id").(string)
	clusterId := d.Get("cluster_id").(string)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if err := service.DeleteTdmqTopic(ctx, environId, topicName, clusterId); err != nil {
			if sdkErr, ok := err.(*errors.CloudSDKError); ok {
				if sdkErr.Code == VPCNotFound {
					return nil
				}
			}
			return resource.RetryableError(err)
		}
		return nil
	})

	return err
}
