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
			"endpoint":           "Kafka接入点地址",
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

	// Resolve CKafka endpoints for all rules
	endpoints := make(map[string]string)
	for _, raw := range rules {
		rule := raw.(map[string]interface{})
		instanceId := rule["ckafka_instance_id"].(string)
		if _, ok := endpoints[instanceId]; ok {
			continue
		}
		ep, err := cosService.ResolveCkafkaEndpoint(ctx, instanceId)
		if err != nil {
			return err
		}
		endpoints[instanceId] = ep
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

	if d.HasChange("notification_rule") {
		rules := d.Get("notification_rule").([]interface{})

		endpoints := make(map[string]string)
		for _, raw := range rules {
			rule := raw.(map[string]interface{})
			instanceId := rule["ckafka_instance_id"].(string)
			if _, ok := endpoints[instanceId]; ok {
				continue
			}
			ep, err := cosService.ResolveCkafkaEndpoint(ctx, instanceId)
			if err != nil {
				return err
			}
			endpoints[instanceId] = ep
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
