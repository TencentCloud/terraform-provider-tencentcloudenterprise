/*
Provides a CSP resource to manage COS bucket event notification configuration.

Event notifications allow you to receive messages when certain events happen in your bucket,
such as object creation or deletion. Notifications are delivered to CKafka.

# Example Usage

Without SASL authentication:

```hcl

	resource "tencentcloudenterprise_csp_bucket_notification" "example" {
	  bucket = "est123-1255000115"

	  notification_rule {
	    id                 = "rule-1"
	    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
	    ckafka_instance_id = "ckafka-7k3pve8e"
	    topic              = "my-notification-topic"
	  }
	}

```

With SASL authentication:

```hcl

	resource "tencentcloudenterprise_csp_bucket_notification" "example_sasl" {
	  bucket = "est123-1255000115"

	  notification_rule {
	    id                 = "rule-sasl"
	    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
	    ckafka_instance_id = "ckafka-7k3pve8e"
	    topic              = "my-notification-topic"
	    sasl_user          = "123123"
	    sasl_password      = "Tencent@321"
	  }
	}

```

# Import

CSP bucket notification can be imported using the bucket name, e.g.

```
$ terraform import tencentcloudenterprise_csp_bucket_notification.example mybucket-1258798060
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_csp_bucket_notification", CNDescription{
		TerraformTypeCN: "CSP存储桶事件通知",
		DescriptionCN:   "提供CSP存储桶事件通知资源，用于配置存储桶的事件触发通知到CKafka。",
		AttributesCN: map[string]string{
			"bucket":             "存储桶名称",
			"notification_rule":  "事件通知规则列表",
			"id":                 "规则唯一标识",
			"events":             "触发事件类型列表",
			"ckafka_instance_id": "CKafka实例ID",
			"topic":              "Kafka Topic名称，消息投递的目标Topic",
			"endpoint":           "Kafka接入点地址",
			"filter_prefix":      "对象键前缀过滤，仅匹配该前缀的对象才触发通知",
			"filter_suffix":      "对象键后缀过滤，仅匹配该后缀的对象才触发通知",
			"sasl_user":          "SASL认证用户AppID，provider会自动拼接为 instanceId#appId 格式",
			"sasl_password":      "SASL认证密码",
		},
	})
}

func resourceTencentCloudCspBucketNotification() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a CSP resource to manage COS bucket event notification configuration.",
		Create:      resourceTencentCloudCspBucketNotificationCreate,
		Read:        resourceTencentCloudCspBucketNotificationRead,
		Update:      resourceTencentCloudCspBucketNotificationUpdate,
		Delete:      resourceTencentCloudCspBucketNotificationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateCosBucketName,
				Description:  "The name of the bucket. Bucket format should be [custom name]-[appid], for example `mybucket-1258798060`.",
			},
			"notification_rule": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "A list of notification rules for the bucket.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique identifier for the notification rule.",
						},
						"events": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "List of event types that trigger the notification. Valid values include: `cos:ObjectCreated:*`, `cos:ObjectCreated:Put`, `cos:ObjectCreated:Copy`, `cos:ObjectCreated:Post`, `cos:ObjectCreated:CompleteMultipartUpload`, `cos:ObjectRemove:*`, `cos:ObjectRemove:Delete`, `cos:ObjectRemove:DeleteMarkerCreated`.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
				"ckafka_instance_id": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The CKafka instance ID to deliver notifications to. The provider will automatically resolve the Kafka endpoint from this instance.",
				},
				"topic": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The Kafka topic name to deliver notification messages to.",
				},
				"endpoint": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "The resolved Kafka endpoint address in the format `kafka://host:port`. Automatically determined from the CKafka instance route.",
				},
				"sasl_user": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "SASL username (AppID). When specified, the provider automatically constructs the full user string as `{ckafka_instance_id}#{sasl_user}` for authentication.",
				},
				"sasl_password": {
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
					Description: "SASL password for Kafka authentication.",
				},
				"filter_prefix": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Object key prefix for filtering notifications. Only objects matching this prefix will trigger notifications. For example `adc` means only objects under the `adc/` path.",
				},
				"filter_suffix": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Object key suffix for filtering notifications. Only objects matching this suffix will trigger notifications. For example `.jpg`.",
				},
					},
				},
			},
		},
	}
}

func resourceTencentCloudCspBucketNotificationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_bucket_notification.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	bucket := d.Get("bucket").(string)
	rules := d.Get("notification_rule").([]interface{})

	cosService := CosService{client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true}
	ckafkaService := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}

	// Validate CKafka instances and topics, resolve endpoints, and check Kafka connectivity
	endpoints := make(map[string]string)
	for _, raw := range rules {
		rule := raw.(map[string]interface{})
		instanceId := rule["ckafka_instance_id"].(string)
		topic := rule["topic"].(string)
		saslUser, _ := rule["sasl_user"].(string)
		saslPassword, _ := rule["sasl_password"].(string)
		hasSasl := saslUser != ""

		// Validate ckafka instance exists
		if _, ok := endpoints[instanceId]; !ok {
			_, has, err := ckafkaService.DescribeInstanceById(ctx, instanceId)
			if err != nil {
				return fmt.Errorf("failed to check ckafka instance %s: %s", instanceId, err.Error())
			}
			if !has {
				return fmt.Errorf("ckafka instance %s does not exist", instanceId)
			}

			// Resolve endpoint: AccessType=1 for SASL, AccessType=0 for non-SASL
			var accessType int64
			if hasSasl {
				accessType = 1
			}
			ep, err := cosService.ResolveCkafkaEndpointByAccessType(ctx, instanceId, accessType)
			if err != nil {
				return err
			}
			endpoints[instanceId] = ep
		}

		// Validate topic exists in the ckafka instance
		topicList, err := ckafkaService.DescribeCkafkaTopics(ctx, instanceId, topic)
		if err != nil {
			return fmt.Errorf("failed to check ckafka topic %s in instance %s: %s", topic, instanceId, err.Error())
		}
		topicFound := false
		for _, t := range topicList {
			if t.TopicName != nil && *t.TopicName == topic {
				topicFound = true
				break
			}
		}
		if !topicFound {
			return fmt.Errorf("topic %s does not exist in ckafka instance %s", topic, instanceId)
		}

		// Check Kafka connectivity via CSP
		ep := endpoints[instanceId]
		// endpoint is "kafka://host:port", extract "host:port"
		host := strings.TrimPrefix(ep, "kafka://")
		checkUser := ""
		checkPassword := ""
		if hasSasl {
			checkUser = instanceId + "#" + saslUser
			checkPassword = saslPassword
		}
		if err := cosService.CheckKafkaConnectivity(ctx, host, checkUser, checkPassword); err != nil {
			return err
		}
	}

	config := BuildNotificationConfig(rules, endpoints)
	if putErr := cosService.PutBucketNotification(ctx, bucket, config); putErr != nil {
		return putErr
	}

	d.SetId(bucket)

	// Wait for the notification configuration to propagate before reading back.
	if waitErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := cosService.GetBucketNotification(ctx, bucket)
		if e != nil {
			return retryError(e)
		}
		if result == nil || len(result.TopicConfigurations) == 0 {
			return resource.RetryableError(fmt.Errorf("bucket notification not ready yet"))
		}
		return nil
	}); waitErr != nil {
		log.Printf("[CRITAL]%s wait csp bucket notification ready failed, reason:%s\n", logId, waitErr.Error())
		return waitErr
	}

	return resourceTencentCloudCspBucketNotificationRead(d, meta)
}

func resourceTencentCloudCspBucketNotificationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_bucket_notification.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	bucket := d.Id()
	cosService := CosService{client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true}

	var config *CspNotificationConfiguration
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := cosService.GetBucketNotification(ctx, bucket)
		if e != nil {
			return retryError(e)
		}
		config = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read csp bucket notification failed, reason:%s\n", logId, err.Error())
		return err
	}

	if config == nil || len(config.TopicConfigurations) == 0 {
		d.SetId("")
		return nil
	}

	_ = d.Set("bucket", bucket)
	// Pass existing rules so sasl_user/sasl_password (write-only, not returned by GET) are preserved.
	existingRules := d.Get("notification_rule").([]interface{})
	_ = d.Set("notification_rule", FlattenNotificationRules(config, existingRules))

	return nil
}

func resourceTencentCloudCspBucketNotificationUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_bucket_notification.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	bucket := d.Id()
	cosService := CosService{client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true}
	ckafkaService := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}

	if d.HasChange("notification_rule") {
		rules := d.Get("notification_rule").([]interface{})

		endpoints := make(map[string]string)
		for _, raw := range rules {
			rule := raw.(map[string]interface{})
			instanceId := rule["ckafka_instance_id"].(string)
			topic := rule["topic"].(string)
			saslUser, _ := rule["sasl_user"].(string)
			saslPassword, _ := rule["sasl_password"].(string)
			hasSasl := saslUser != ""

			// Validate ckafka instance exists
			if _, ok := endpoints[instanceId]; !ok {
				_, has, err := ckafkaService.DescribeInstanceById(ctx, instanceId)
				if err != nil {
					return fmt.Errorf("failed to check ckafka instance %s: %s", instanceId, err.Error())
				}
				if !has {
					return fmt.Errorf("ckafka instance %s does not exist", instanceId)
				}

				var accessType int64
				if hasSasl {
					accessType = 1
				}
				ep, err := cosService.ResolveCkafkaEndpointByAccessType(ctx, instanceId, accessType)
				if err != nil {
					return err
				}
				endpoints[instanceId] = ep
			}

			// Validate topic exists in the ckafka instance
			topicList, err := ckafkaService.DescribeCkafkaTopics(ctx, instanceId, topic)
			if err != nil {
				return fmt.Errorf("failed to check ckafka topic %s in instance %s: %s", topic, instanceId, err.Error())
			}
			topicFound := false
			for _, t := range topicList {
				if t.TopicName != nil && *t.TopicName == topic {
					topicFound = true
					break
				}
			}
			if !topicFound {
				return fmt.Errorf("topic %s does not exist in ckafka instance %s", topic, instanceId)
			}

			// Check Kafka connectivity via CSP
			ep := endpoints[instanceId]
			host := strings.TrimPrefix(ep, "kafka://")
			checkUser := ""
			checkPassword := ""
			if hasSasl {
				checkUser = instanceId + "#" + saslUser
				checkPassword = saslPassword
			}
			if err := cosService.CheckKafkaConnectivity(ctx, host, checkUser, checkPassword); err != nil {
				return err
			}
		}

		config := BuildNotificationConfig(rules, endpoints)
		if err := cosService.PutBucketNotification(ctx, bucket, config); err != nil {
			return err
		}
	}

	time.Sleep(3 * time.Second)
	return resourceTencentCloudCspBucketNotificationRead(d, meta)
}

func resourceTencentCloudCspBucketNotificationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_bucket_notification.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	bucket := d.Id()
	cosService := CosService{client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true}

	if err := cosService.DeleteBucketNotification(ctx, bucket); err != nil {
		log.Printf("[CRITAL]%s delete csp bucket notification failed, reason:%s\n", logId, err.Error())
		return err
	}

	time.Sleep(3 * time.Second)
	return nil
}
