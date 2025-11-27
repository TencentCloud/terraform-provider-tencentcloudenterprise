/*
Provides a resource to create a cls topic.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_topic" "topic" {
	  topic_name           = "topic"
	  logset_id            = "5cd3a17e-fb0b-418c-afd7-77b365397426"
	  auto_split           = false
	  max_split_partitions = 20
	  partition_count      = 1
	  period               = 10
	  storage_type         = "hot"
	  tags                 = {
	    "test" = "test",
	  }
	}

```

# Import

cls topic can be imported using the id, e.g.

```
$ terraform import cloud_cls_topic.topic 2f5764c1-c833-44c5-84c7-950979b2a278
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_topic", CNDescription{
		TerraformTypeCN: "CLS日志主题",
		DescriptionCN:   "提供CLS日志主题资源，用于创建和管理日志服务日志主题。",
		AttributesCN: map[string]string{
			"topic_name":           "日志主题名称",
			"logset_id":            "日志集ID",
			"auto_split":           "是否开启自动分裂，是否开启自动分裂，默认值为true。true表示开启自动分裂，false表示关闭自动分裂。",
			"max_split_partitions": "最大分裂分区数，开启自动分裂后，每个主题能够允许的最大分区数，默认值为50",
			"partition_count":      "日志主题分区个数。默认创建1个，最大支持创建10个分区。",
			"period":               "生命周期，单位天；可取值范围1~90。默认30天。",
			"storage_type":         "日志主题的存储类型，可选值 hot（实时存储）；默认为hot。",
			"tags":                 "标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。",
			"sub_assumer_name":     "二级产品标识。",
			"describes":            "日志主题描述。",
			"hot_period":           "0：关闭日志沉降。 非0：开启日志沉降后标准存储的天数。HotPeriod需要大于等于7，且小于Period。仅在StorageType为 hot 时生效",
			"encryption":           "加密相关参数。 支持加密地域并且开白用户可以传此参数，其他场景不能传递该参数。  0或者不传： 不加密 1：kms-cls 云产品秘钥加密",
			"biz_type":             "0: 日志主题topic 1: 时序主题topic",
			"topic_id":             "作为入参：日志主题ID，格式为：用户自定义部分-用户appid，用户自定义部分仅支持小写字母、数字和-，且不能以-开头和结尾，长度为3至40字符，结尾使用-拼接用户appid；作为出参：日志主题ID",
			"is_web_tracking":      "webtracking开关； false: 关闭 true： 开启",
		},
	})
}

func resourceTencentCloudClsTopic() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsTopicCreate,
		Read:        resourceTencentCloudClsTopicRead,
		Delete:      resourceTencentCloudClsTopicDelete,
		Update:      resourceTencentCloudClsTopicUpdate,
		Description: "Provides a resource to create and manage CLS topic",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"logset_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Logset ID.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Log topic name.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Log topic ID is in the format of [user-defined part (3-40 characters, supporting lowercase letters, numbers and \"-\", not starting or ending with \"-\")]-[user appid], with the user-defined part concatenated with the user appid via \"-\" at the end.",
			},
			"partition_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of partitions for a log topic. One partition is created by default, and a maximum of 10 partitions can be created.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "A list of tags. By specifying this parameter, you can bind tags to the corresponding log topic at the same time. A maximum of 10 tag key-value pairs are supported, and the same resource can only be bound to the same tag key.",
			},
			"auto_split": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable automatic split. Default value is true. true means enabled, false means disabled.",
			},
			"max_split_partitions": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximum number of partitions allowed for each topic after auto-split is enabled. The default value is 50.",
			},
			"storage_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Storage type of the log topic. The optional value is hot (real-time storage); the default is hot.",
			},
			"period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Lifecycle in days; value range: 1-90. Default: 30 days.",
			},
			"sub_assumer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secondary product identifier.",
			},
			"describes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Log topic description.",
			},
			"hot_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "0: Disable log sinking. Non-zero: The number of days for standard storage after enabling log sinking. HotPeriod needs to be greater than or equal to 7 and less than Period. It only takes effect when StorageType is hot.",
			},
			"encryption": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Encryption-related parameters. This parameter can be passed by whitelisted users in regions that support encryption. It cannot be passed in other scenarios. 0 or not passed: no encryption; 1: kms-cls cloud product key encryption.",
			},
			"biz_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "0: Log topic, 1: Time series topic.",
			},
			"is_web_tracking": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Webtracking switch; false: off, true: on.",
			},
		},
	}
}

func resourceTencentCloudClsTopicCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_topic.create")()

	logId := getLogId(contextNil)

	var (
		request  = cls.NewCreateTopicRequest()
		response *cls.CreateTopicResponse
	)

	if v, ok := d.GetOk("logset_id"); ok {
		request.LogsetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("topic_name"); ok {
		request.TopicName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("partition_count"); ok {
		request.PartitionCount = helper.IntInt64(v.(int))
	}

	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		for k, v := range tags {
			key := k
			value := v
			request.Tags = append(request.Tags, &cls.Tag{
				Key:   &key,
				Value: &value,
			})
		}
	}

	if v, ok := d.GetOkExists("auto_split"); ok {
		request.AutoSplit = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("max_split_partitions"); ok {
		request.MaxSplitPartitions = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("storage_type"); ok {
		request.StorageType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("period"); ok {
		request.Period = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("sub_assumer_name"); ok {
		request.SubAssumerName = helper.String(v.(string))
	}
	if v, ok := d.GetOk("describes"); ok {
		request.Describes = helper.String(v.(string))
	}

	if v, ok := d.GetOk("hot_period"); ok {
		request.HotPeriod = helper.Uint64(uint64(v.(int)))
	}

	if v, ok := d.GetOk("encryption"); ok {
		request.Encryption = helper.Uint64(v.(uint64))
	}

	if v, ok := d.GetOk("biz_type"); ok {
		request.BizType = helper.Uint64(v.(uint64))
	}

	if v, ok := d.GetOk("is_web_tracking"); ok {
		request.IsWebTracking = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("topic_id"); ok {
		request.TopicId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateTopic(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cls topic failed, reason:%+v", logId, err)
		return err
	}

	id := *response.Response.TopicId
	d.SetId(id)
	return resourceTencentCloudClsTopicRead(d, meta)
}

func resourceTencentCloudClsTopicRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_topic.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	id := d.Id()

	topic, err := service.DescribeClsTopicById(ctx, id)

	if err != nil {
		return err
	}

	if topic == nil {
		d.SetId("")
		return fmt.Errorf("resource `Topic` %s does not exist", id)
	}

	_ = d.Set("logset_id", topic.LogsetId)
	_ = d.Set("topic_name", topic.TopicName)
	_ = d.Set("partition_count", topic.PartitionCount)

	tags := make(map[string]string, len(topic.Tags))
	for _, tag := range topic.Tags {
		tags[*tag.Key] = *tag.Value
	}
	_ = d.Set("tags", tags)
	_ = d.Set("auto_split", topic.AutoSplit)
	_ = d.Set("max_split_partitions", topic.MaxSplitPartitions)
	_ = d.Set("storage_type", topic.StorageType)
	_ = d.Set("period", topic.Period)
	_ = d.Set("sub_assumer_name", topic.SubAssumerName)
	_ = d.Set("describes", topic.Describes)
	_ = d.Set("hot_period", topic.HotPeriod)
	_ = d.Set("encryption", topic.Encryption)
	_ = d.Set("biz_type", topic.BizType)
	_ = d.Set("topic_id", topic.TopicId)
	_ = d.Set("is_web_tracking", topic.IsWebTracking)

	return nil
}

func resourceTencentCloudClsTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_topic.update")()
	logId := getLogId(contextNil)
	request := cls.NewModifyTopicRequest()

	request.TopicId = helper.String(d.Id())

	if d.HasChange("partition_count") {
		return fmt.Errorf("`partition_count` do not support change now.")
	}

	if d.HasChange("storage_type") {
		return fmt.Errorf("`storage_type` do not support change now")
	}

	if d.HasChange("is_web_tracking") {
		return fmt.Errorf("`is_web_tracking` do not support change now")
	}

	if d.HasChange("sub_assumer_name") {
		return fmt.Errorf("`sub_assumer_name` do not support change now")
	}

	if d.HasChange("hot_period") {
		return fmt.Errorf("`hot_period` do not support change now")
	}

	attributeModified := false

	if d.HasChange("auto_split") {
		attributeModified = true
		request.AutoSplit = helper.Bool(d.Get("auto_split").(bool))
	}

	if d.HasChange("max_split_partitions") {
		attributeModified = true
		request.MaxSplitPartitions = helper.IntInt64(d.Get("max_split_partitions").(int))
	}

	if d.HasChange("period") {
		attributeModified = true
		request.Period = helper.IntInt64(d.Get("period").(int))
	}

	if d.HasChange("topic_name") {
		attributeModified = true
		request.TopicName = helper.String(d.Get("topic_name").(string))
	}

	if d.HasChange("describes") {
		attributeModified = true
		request.Describes = helper.String(d.Get("describes").(string))
	}

	if attributeModified {
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyTopic(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})

		if err != nil {
			return err
		}
	}

	if d.HasChange("tags") {

		tags := d.Get("tags").(map[string]interface{})
		request.Tags = make([]*cls.Tag, 0, len(tags))
		for k, v := range tags {
			key := k
			value := v
			request.Tags = append(request.Tags, &cls.Tag{
				Key:   &key,
				Value: helper.String(value.(string)),
			})
		}
	}

	return resourceTencentCloudClsTopicRead(d, meta)
}

func resourceTencentCloudClsTopicDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_topic.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	id := d.Id()

	if err := service.DeleteClsTopic(ctx, id); err != nil {
		return err
	}

	return nil
}
