/*
Provides a resource to create a CLB log topic.

Example Usage

```hcl
resource "tencentcloudenterprise_clb_log_topic" "topic" {
  log_set_id = tencentcloudenterprise_clb_log_set.set.id
  topic_name = "clb-topic"
}
```

Import

CLB log topic can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_clb_log_topic.topic 439b0e84-c5dc-4382-b4c2-937a5a4d8245
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_clb_log_topic", CNDescription{
		TerraformTypeCN: "CLB 日志主题",
		DescriptionCN:   "提供 CLB 实例日志主题资源，用于创建和管理 CLB 的 CLS 日志主题。",
		AttributesCN: map[string]string{
			"log_set_id":  "CLB 实例关联的日志集 ID",
			"topic_name":  "CLB 实例日志主题名称",
			"status":      "日志主题状态，true 表示启用，false 表示禁用，默认为 true",
			"create_time": "日志主题创建时间",
		},
	})
}

func resourceTencentCloudClbLogTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudClbLogTopicCreate,
		Read:   resourceTencentCloudClbLogTopicRead,
		Update: resourceTencentCloudClbLogTopicUpdate,
		Delete: resourceTencentCloudClbLogTopicDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"log_set_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Log set ID of CLB instance.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Log topic name of CLB instance.",
			},
			"status": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "The status of log topic. true: enable; false: disable. Default is true.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Log topic creation time.",
			},
		},
	}
}

func resourceTencentCloudClbLogTopicCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_log_topic.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clsService := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	logSetId := d.Get("log_set_id").(string)
	// verify logset exists
	info, err := clsService.DescribeClsLogset(ctx, logSetId)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("log_set %s does not exist", logSetId)
	}

	topicName := d.Get("topic_name").(string)
	topicId, err := clsService.CreateClsTopic(ctx, logSetId, topicName, 0)
	if err != nil {
		log.Printf("[CRITAL]%s create clb topic failed, reason:%+v", logId, err)
		return err
	}
	if topicId == "" {
		return fmt.Errorf("[CRITAL]%s create clb topic failed, topicId is empty", logId)
	}

	d.SetId(topicId)

	if v, ok := d.GetOkExists("status"); ok {
		if !v.(bool) {
			err := clsService.ModifyClsTopic(ctx, topicId, helper.Bool(false))
			if err != nil {
				return err
			}
		}
	}

	return resourceTencentCloudClbLogTopicRead(d, meta)
}

func resourceTencentCloudClbLogTopicRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_log_topic.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()
	clsService := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	res, err := clsService.DescribeClsTopicById(ctx, id)
	if err != nil {
		return err
	}
	if res == nil {
		d.SetId("")
		return nil
	}
	_ = d.Set("log_set_id", res.LogsetId)
	_ = d.Set("topic_name", res.TopicName)
	_ = d.Set("create_time", res.CreateTime)
	_ = d.Set("status", res.Status)
	log.Printf("[DEBUG]%s read clb log topic success, id: %s", logId, id)
	return nil
}

func resourceTencentCloudClbLogTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_log_topic.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		topicId = d.Id()
	)

	clsService := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	if d.HasChange("status") {
		if v, ok := d.GetOkExists("status"); ok {
			ctx := context.WithValue(context.TODO(), logIdKey, logId)
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				e := clsService.ModifyClsTopic(ctx, topicId, helper.Bool(v.(bool)))
				if e != nil {
					return retryError(e)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	return resourceTencentCloudClbLogTopicRead(d, meta)
}

func resourceTencentCloudClbLogTopicDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_log_topic.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()
	clsService := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	err := clsService.DeleteClsTopic(ctx, id)
	if err != nil {
		log.Printf("[CRITAL]%s delete clb log topic failed, reason:%+v", logId, err)
		return err
	}
	return nil
}
