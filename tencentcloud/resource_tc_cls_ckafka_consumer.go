/*
Provides a resource to create a cls ckafka_consumer

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_ckafka_consumer" "ckafka_consumer" {
	  compression  = 1
	  need_content = true
	  topic_id     = "7e34a3a7-635e-4da8-9005-88106c1fde69"

	  ckafka {
	    instance_id   = "ckafka-qzoeaqx8"
	    instance_name = "ckafka-instance"
	    topic_id      = "topic-c6tm4kpm"
	    topic_name    = "name"
	    vip           = "203.0.113.23"
	    vport         = "9092"
	  }

	  content {
	    enable_tag         = true
	    meta_fields        = [
	      "__FILENAME__",
	      "__HOSTNAME__",
	      "__PKGID__",
	      "__SOURCE__",
	      "__TIMESTAMP__",
	    ]
	    tag_json_not_tiled = true
	    timestamp_accuracy = 2
	  }
	}

```

# Import

cls ckafka_consumer can be imported using the id, e.g.

```
terraform import cloud_cls_ckafka_consumer.ckafka_consumer topic_id
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_ckafka_consumer", CNDescription{
		TerraformTypeCN: "CLS Ckafka消费者",
		DescriptionCN:   "提供CLS Ckafka消费者资源，用于创建和管理日志服务Ckafka消费者。",
		AttributesCN: map[string]string{
			"topic_id":           "投递任务绑定的日志主题 ID",
			"ckafka":             "CKafka的描述",
			"compression":        "投递时压缩方式，取值0，2，3。[0：NONE；2：SNAPPY；3：LZ4]",
			"content":            "如果需要投递元数据信息，元数据信息的描述",
			"need_content":       "是否投递日志的元数据信息，默认为 true。 当NeedContent为true时：字段Content有效。 当NeedContent为false时：字段Content无效。",
			"enable_tag":         "是否投递 TAG 信息。 当EnableTag为true时，表示投递TAG元信息。",
			"meta_fields":        "需要投递的元数据列表，目前仅支持：__SOURCE__，__FILENAME__，__TIMESTAMP__，__HOSTNAME__和__PKGID__",
			"tag_json_not_tiled": "当EnableTag为true时，必须填写TagJsonNotTiled字段。 TagJsonNotTiled用于标识tag信息是否json平铺。  TagJsonNotTiled为true时不平铺，示例： TAG信息：{\"__TAG__\":{\"fieldA\":200,\"fieldB\":\"text\"}} 不平铺：{\"__TAG__\":{\"fieldA\":200,\"fieldB\":\"text\"}}  TagJsonNotTiled为false时平铺，示例： TAG信息：{\"__TAG__\":{\"fieldA\":200,\"fieldB\":\"text\"}} 平铺：{\"__TAG__.fieldA\":200,\"__TAG__.fieldB\":\"text\"}",
			"timestamp_accuracy": "投递时间戳精度，可选项 [1：秒；2：毫秒] ，默认是1。",
			"json_type":          "投递Json格式。 JsonType为0：和原始日志一致，不转义。示例： 日志原文：{\"a\":\"aa\", \"b\":{\"b1\":\"b1b1\", \"c1\":\"c1c1\"}} 投递到Ckafka：{\"a\":\"aa\", \"b\":{\"b1\":\"b1b1\", \"c1\":\"c1c1\"}}  JsonType为1：转义。示例： 日志原文：{\"a\":\"aa\", \"b\":{\"b1\":\"b1b1\", \"c1\":\"c1c1\"}} 投递到Ckafka：{\"a\":\"aa\",\"b\":\"{\\\"b1\\\":\\\"b1b1\\\", \\\"c1\\\":\\\"c1c1\\\"}\"}",
			"vip":                "Ckafka 的 Vip",
			"vport":              "Ckafka 的 Vport",
			"instance_id":        "Ckafka 的 InstanceId",
			"instance_name":      "Ckafka 的 InstanceName",
			"topic_name":         "Ckafka 的 TopicName",
		},
	})
}

func resourceTencentCloudClsCkafkaConsumer() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsCkafkaConsumerCreate,
		Read:        resourceTencentCloudClsCkafkaConsumerRead,
		Update:      resourceTencentCloudClsCkafkaConsumerUpdate,
		Delete:      resourceTencentCloudClsCkafkaConsumerDelete,
		Description: "Provides a resource to create and manage CLS CKafka consumer",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "ID of the log topic bound to the delivery task.",
			},

			"need_content": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether to deliver log metadata information. Default is true. When NeedContent is true, the Content field is valid. When NeedContent is false, the Content field is invalid.",
			},

			"content": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Description of metadata information if metadata needs to be delivered.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_tag": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to deliver TAG information. When EnableTag is true, it means delivering TAG metadata information.",
						},
						"meta_fields": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required:    true,
							Description: "List of metadata to be delivered. Currently only supports: __SOURCE__, __FILENAME__, __TIMESTAMP__, __HOSTNAME__ and __PKGID__.",
						},
						"tag_json_not_tiled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "When EnableTag is true, this field must be filled. TagJsonNotTiled is used to identify whether tag information is JSON flattened. When TagJsonNotTiled is true, it is not flattened. When TagJsonNotTiled is false, it is flattened.",
						},
						"timestamp_accuracy": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delivery timestamp precision. Options: [1: second; 2: millisecond]. Default is 1.",
						},
						"json_type": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delivery JSON format. JsonType 0: consistent with original log, no escaping. JsonType 1: escaped.",
						},
					},
				},
			},

			"ckafka": {
				Required:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "CKafka description.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "CKafka VIP address.",
						},
						"vport": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "CKafka VPort.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "CKafka instance ID.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "CKafka instance name.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "CKafka topic ID.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "CKafka topic name.",
						},
					},
				},
			},

			"compression": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Compression method for delivery. Values: 0 (NONE), 2 (SNAPPY), 3 (LZ4).",
			},
		},
	}
}

func resourceTencentCloudClsCkafkaConsumerCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_ckafka_consumer.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = cls.NewCreateConsumerRequest()
		topicId string
	)
	if v, ok := d.GetOk("topic_id"); ok {
		topicId = v.(string)
		request.TopicId = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("need_content"); ok {
		request.NeedContent = helper.Bool(v.(bool))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "content"); ok {
		consumerContent := cls.ConsumerContent{}
		if v, ok := dMap["enable_tag"]; ok {
			consumerContent.EnableTag = helper.Bool(v.(bool))
		}
		if v, ok := dMap["meta_fields"]; ok {
			metaFieldsList := v.([]interface{})
			for _, metaField := range metaFieldsList {
				consumerContent.MetaFields = append(consumerContent.MetaFields, helper.String(metaField.(string)))
			}
		}
		if v, ok := dMap["tag_json_not_tiled"]; ok {
			consumerContent.TagJsonNotTiled = helper.Bool(v.(bool))
		}
		if v, ok := dMap["timestamp_accuracy"]; ok {
			consumerContent.TimestampAccuracy = helper.IntInt64(v.(int))
		}
		request.Content = &consumerContent
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "ckafka"); ok {
		ckafka := cls.Ckafka{}
		if v, ok := dMap["vip"]; ok {
			ckafka.Vip = helper.String(v.(string))
		}
		if v, ok := dMap["vport"]; ok {
			ckafka.Vport = helper.String(v.(string))
		}
		if v, ok := dMap["instance_id"]; ok {
			ckafka.InstanceId = helper.String(v.(string))
		}
		if v, ok := dMap["instance_name"]; ok {
			ckafka.InstanceName = helper.String(v.(string))
		}
		if v, ok := dMap["topic_id"]; ok {
			ckafka.TopicId = helper.String(v.(string))
		}
		if v, ok := dMap["topic_name"]; ok {
			ckafka.TopicName = helper.String(v.(string))
		}
		request.Ckafka = &ckafka
	}

	if v, ok := d.GetOkExists("compression"); ok {
		request.Compression = helper.IntInt64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateConsumer(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cls ckafkaConsumer failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(topicId)

	return resourceTencentCloudClsCkafkaConsumerRead(d, meta)
}

func resourceTencentCloudClsCkafkaConsumerRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_ckafka_consumer.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	topicId := d.Id()

	ckafkaConsumer, err := service.DescribeClsCkafkaConsumerById(ctx, topicId)
	if err != nil {
		return err
	}

	if ckafkaConsumer == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ClsCkafkaConsumer` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("topic_id", topicId)

	if ckafkaConsumer.NeedContent != nil {
		_ = d.Set("need_content", ckafkaConsumer.NeedContent)
	}

	if ckafkaConsumer.Content != nil {
		contentMap := map[string]interface{}{}

		if ckafkaConsumer.Content.EnableTag != nil {
			contentMap["enable_tag"] = ckafkaConsumer.Content.EnableTag
		}

		if ckafkaConsumer.Content.MetaFields != nil {
			contentMap["meta_fields"] = ckafkaConsumer.Content.MetaFields
		}

		if ckafkaConsumer.Content.TagJsonNotTiled != nil {
			contentMap["tag_json_not_tiled"] = ckafkaConsumer.Content.TagJsonNotTiled
		}

		if ckafkaConsumer.Content.TimestampAccuracy != nil {
			contentMap["timestamp_accuracy"] = ckafkaConsumer.Content.TimestampAccuracy
		}

		_ = d.Set("content", []interface{}{contentMap})
	}

	if ckafkaConsumer.Ckafka != nil {
		ckafkaMap := map[string]interface{}{}

		if ckafkaConsumer.Ckafka.Vip != nil {
			ckafkaMap["vip"] = ckafkaConsumer.Ckafka.Vip
		}

		if ckafkaConsumer.Ckafka.Vport != nil {
			ckafkaMap["vport"] = ckafkaConsumer.Ckafka.Vport
		}

		if ckafkaConsumer.Ckafka.InstanceId != nil {
			ckafkaMap["instance_id"] = ckafkaConsumer.Ckafka.InstanceId
		}

		if ckafkaConsumer.Ckafka.InstanceName != nil {
			ckafkaMap["instance_name"] = ckafkaConsumer.Ckafka.InstanceName
		}

		if ckafkaConsumer.Ckafka.TopicId != nil {
			ckafkaMap["topic_id"] = ckafkaConsumer.Ckafka.TopicId
		}

		if ckafkaConsumer.Ckafka.TopicName != nil {
			ckafkaMap["topic_name"] = ckafkaConsumer.Ckafka.TopicName
		}

		_ = d.Set("ckafka", []interface{}{ckafkaMap})
	}

	if ckafkaConsumer.Compression != nil {
		_ = d.Set("compression", ckafkaConsumer.Compression)
	}

	return nil
}

func resourceTencentCloudClsCkafkaConsumerUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_ckafka_consumer.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cls.NewModifyConsumerRequest()

	topicId := d.Id()

	request.TopicId = &topicId

	needChange := false
	mutableArgs := []string{"need_content", "content", "ckafka", "compression"}

	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {

		if v, ok := d.GetOkExists("need_content"); ok {
			request.NeedContent = helper.Bool(v.(bool))
		}

		if dMap, ok := helper.InterfacesHeadMap(d, "content"); ok {
			consumerContent := cls.ConsumerContent{}
			if v, ok := dMap["enable_tag"]; ok {
				consumerContent.EnableTag = helper.Bool(v.(bool))
			}
			if v, ok := dMap["meta_fields"]; ok {
				metaFieldsList := v.([]interface{})
				for _, metaField := range metaFieldsList {
					consumerContent.MetaFields = append(consumerContent.MetaFields, helper.String(metaField.(string)))
				}
			}
			if v, ok := dMap["tag_json_not_tiled"]; ok {
				consumerContent.TagJsonNotTiled = helper.Bool(v.(bool))
			}
			if v, ok := dMap["timestamp_accuracy"]; ok {
				consumerContent.TimestampAccuracy = helper.IntInt64(v.(int))
			}
			request.Content = &consumerContent
		}

		if dMap, ok := helper.InterfacesHeadMap(d, "ckafka"); ok {
			ckafka := cls.Ckafka{}
			if v, ok := dMap["vip"]; ok {
				ckafka.Vip = helper.String(v.(string))
			}
			if v, ok := dMap["vport"]; ok {
				ckafka.Vport = helper.String(v.(string))
			}
			if v, ok := dMap["instance_id"]; ok {
				ckafka.InstanceId = helper.String(v.(string))
			}
			if v, ok := dMap["instance_name"]; ok {
				ckafka.InstanceName = helper.String(v.(string))
			}
			if v, ok := dMap["topic_id"]; ok {
				ckafka.TopicId = helper.String(v.(string))
			}
			if v, ok := dMap["topic_name"]; ok {
				ckafka.TopicName = helper.String(v.(string))
			}
			request.Ckafka = &ckafka
		}

		if v, ok := d.GetOkExists("compression"); ok {
			request.Compression = helper.IntInt64(v.(int))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyConsumer(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update cls ckafkaConsumer failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudClsCkafkaConsumerRead(d, meta)
}

func resourceTencentCloudClsCkafkaConsumerDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_ckafka_consumer.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	topicId := d.Id()

	if err := service.DeleteClsCkafkaConsumerById(ctx, topicId); err != nil {
		return err
	}

	return nil
}
