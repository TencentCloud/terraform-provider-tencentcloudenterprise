/*
Provides a resource to create a tdmq subscription_attachment

# Example Usage

```hcl

	resource "tencentcloudenterprise_tdmq_subscription_attachment" "subscription_attachment" {
	  environment_id           = "keep-ns"
	  topic_name               = "keep-topic"
	  subscription_name        = "test-subcription"
	  remark                   = "test"
	  cluster_id               = "pulsar-9n95ax58b9vn"
	  auto_create_policy_topic = true
	  is_idempotent            = true
	}

```

# Import

tdmq subscription_attachment can be imported using the id, e.g.

```
terraform import cloud_tdmq_subscription_attachment.subscription_attachment subscription_attachment_id
```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_subscription_attachment", CNDescription{
		TerraformTypeCN: "TDMQ 订阅关系",
		DescriptionCN:   "提供TDMQ订阅关系资源，用于创建和管理TDMQ订阅关系。",
		AttributesCN: map[string]string{
			"environment_id":           "环境（命名空间）名称",
			"topic_name":               "主题名称",
			"subscription_name":        "订阅者名称",
			"remark":                   "备注",
			"cluster_id":               "Pulsar 集群的ID",
			"auto_create_policy_topic": "是否自动创建死信和重试主题",
			//"post_fix_pattern":         "指定死信和重试主题名称规范",
			"is_idempotent": "是否幂等创建，若否不允许创建同名的订阅关系",
		},
	})
}
func resourceTencentCloudTdmqSubscriptionAttachment() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a resource to create a tdmq subscription_attachment",
		CreateContext: resourceTencentCloudTdmqSubscriptionAttachmentCreate,
		ReadContext:   resourceTencentCloudTdmqSubscriptionAttachmentRead,
		UpdateContext: resourceTencentCloudTdmqSubscriptionAttachmentUpdate,
		DeleteContext: resourceTencentCloudTdmqSubscriptionAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			// 环境（命名空间）名称
			"environment_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Environment (namespace) name.",
			},
			// 主题名称
			"topic_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Topic name.",
			},
			// 订阅者名称，不支持中字以及除了短线和下划线外的特殊字符且不超过150个字符
			"subscription_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Subscriber name, no more than 128 characters.",
			},
			// 备注，128个字符以内
			"remark": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Remarks, within 128 characters.",
			},
			// 是否幂等创建，若否不允许创建同名的订阅关系
			"is_idempotent": {
				Required: true,
				Type:     schema.TypeBool,
				Description: "Whether to create idempotent, if not, " +
					"it is not allowed to create a subscription relationship with the same name.",
			},
			// Pulsar 集群的ID
			"cluster_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "ID of the Pulsar cluster.",
			},
			// 是否自动创建死信和重试主题，True 表示创建，False表示不创建，默认自动创建死信和重试主题
			"auto_create_policy_topic": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether to automatically create dead letters and retry topics, True means to create, False means not to create, the default is to automatically create dead letters and retry topics.",
			},
			// 指定死信和重试主题名称规范，LEGACY表示历史命名规则，COMMUNITY表示Pulsar社区命名规范
			"post_fix_pattern": {
				Optional: true,
				Type:     schema.TypeString,
				Description: "Specify the naming convention for dead letters and retry topics, " +
					"LEGACY means historical naming rules, COMMUNITY means Pulsar community naming conventions.",
				ValidateFunc: validateAllowedStringValue(PostFixPattern),
			},
		},
	}
}

func resourceTencentCloudTdmqSubscriptionAttachmentCreate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_subscription_attachment.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId            = getLogId(contextNil)
		request          = tdmq.NewCreateSubscriptionRequest()
		environmentId    string
		Topic            string
		subscriptionName string
		clusterId        string
	)

	if v, ok := d.GetOk("environment_id"); ok {
		request.EnvironmentId = helper.String(v.(string))
		environmentId = v.(string)
	}

	if v, ok := d.GetOk("topic_name"); ok {
		request.TopicName = helper.String(v.(string))
		Topic = v.(string)
	}

	if v, ok := d.GetOk("subscription_name"); ok {
		request.SubscriptionName = helper.String(v.(string))
		subscriptionName = v.(string)
	}

	if v, ok := d.GetOk("is_idempotent"); ok {
		request.IsIdempotent = helper.Bool(v.(bool))
	} else {
		flag := helper.Bool(v.(bool))
		request.IsIdempotent = flag
	}

	if v, ok := d.GetOk("remark"); ok {
		request.Remark = helper.String(v.(string))
	}

	if v, ok := d.GetOk("cluster_id"); ok {
		request.ClusterId = helper.String(v.(string))
		clusterId = v.(string)
	}

	if v, ok := d.GetOk("auto_create_policy_topic"); ok {
		request.AutoCreatePolicyTopic = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("post_fix_pattern"); ok {
		request.PostFixPattern = helper.String(v.(string))
	}

	err := resource.RetryContext(ctx, writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().CreateSubscription(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmq subscriptionAttachment failed, reason:%+v", logId, err)
		return diag.FromErr(err)
	}

	d.SetId(strings.Join([]string{environmentId, Topic, subscriptionName, clusterId}, FILED_SP))

	return resourceTencentCloudTdmqSubscriptionAttachmentRead(ctx, d, meta)
}

func resourceTencentCloudTdmqSubscriptionAttachmentRead(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_subscription_attachment.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		service = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 4 {
		return diag.Errorf("id is broken, id is %s", d.Id())
	}
	environmentId := idSplit[0]
	Topic := idSplit[1]
	subscriptionName := idSplit[2]
	clusterId := idSplit[3]

	subscriptionAttachment, err := service.DescribeTdmqSubscriptionAttachmentById(ctx, environmentId, Topic, subscriptionName, clusterId)
	if err != nil {
		return diag.FromErr(err)
	}

	if subscriptionAttachment == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TdmqSubscriptionAttachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if subscriptionAttachment.EnvironmentId != nil {
		_ = d.Set("environment_id", subscriptionAttachment.EnvironmentId)
	}

	if subscriptionAttachment.TopicName != nil {
		_ = d.Set("topic_name", subscriptionAttachment.TopicName)
	}

	if subscriptionAttachment.SubscriptionName != nil {
		_ = d.Set("subscription_name", subscriptionAttachment.SubscriptionName)
	}

	if subscriptionAttachment.Remark != nil {
		_ = d.Set("remark", subscriptionAttachment.Remark)
	}
	_ = d.Set("cluster_id", clusterId)

	// Get Topics Status For auto_create_policy_topic
	has, err := service.GetTdmqTopicsAttachmentById(ctx, environmentId, Topic, subscriptionName, clusterId)
	if err != nil {
		return diag.FromErr(err)
	}

	_ = d.Set("auto_create_policy_topic", has)

	return nil
}

func resourceTencentCloudTdmqSubscriptionAttachmentUpdate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_subscription_attachment.update")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudTdmqSubscriptionAttachmentDelete(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_subscription_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		service               = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		autoCreatePolicyTopic bool
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 4 {
		return diag.Errorf("id is broken, id is %s", d.Id())
	}

	environmentId := idSplit[0]
	Topic := idSplit[1]
	subscriptionName := idSplit[2]
	clusterId := idSplit[3]

	// Delete Subscription
	if err := service.DeleteTdmqSubscriptionAttachmentById(ctx, environmentId, Topic, subscriptionName, clusterId); err != nil {
		return diag.FromErr(err)
	}

	if v, ok := d.GetOk("auto_create_policy_topic"); ok {
		autoCreatePolicyTopic = v.(bool)
		if autoCreatePolicyTopic {
			// Delete Topics
			if err := service.DeleteTdmqTopicsAttachmentById(ctx, environmentId, Topic, subscriptionName, clusterId); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}
