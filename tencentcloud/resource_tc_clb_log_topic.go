/*
Provides a resource to create a CLB instance topic.

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
$ terraform import tencentcloudenterprise_clb_log_topic.topic lb-7a0t6zqb
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"sync"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"

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

var clsActionMu = &sync.Mutex{}

func resourceTencentCloudClbLogTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudClbInstanceTopicCreate,
		Read:   resourceTencentCloudClbInstanceTopicRead,
		Update: resourceTencentCloudClbInstanceTopicUpdate,
		Delete: resourceTencentCloudClbInstanceTopicDelete,
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

func resourceTencentCloudClbInstanceTopicCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_log_topic.create")()
	defer inconsistentCheck(d, meta)()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clsService := ClsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	if v, ok := d.GetOk("log_set_id"); ok {
		info, err := clsService.DescribeClsLogset(ctx, v.(string))
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf("resource `log_set` %s does not exist", v.(string))
		}
	}

	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	params := make(map[string]interface{})
	if topicName, ok := d.GetOk("topic_name"); ok {
		params["topic_name"] = topicName
	}
	resp, err := clbService.CreateTopic(ctx, params)
	if err != nil {
		log.Printf("[CRITAL]%s create clb topic failed, reason:%+v", logId, err)
		return err
	}

	topicId := *resp.Response.TopicId
	d.SetId(topicId)

	if v, ok := d.GetOkExists("status"); ok {
		if !v.(bool) {
			request := cls.NewModifyTopicRequest()
			request.TopicId = &topicId
			request.Status = helper.Bool(false)
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyTopic(request)
				if e != nil {
					return retryError(e)
				}
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				return nil
			})

			if err != nil {
				return err
			}
		}
	}

	return resourceTencentCloudClbInstanceTopicRead(d, meta)
}

func resourceTencentCloudClbInstanceTopicRead(d *schema.ResourceData, meta interface{}) error {
	clsActionMu.Lock()
	defer clsActionMu.Unlock()
	defer inconsistentCheck(d, meta)()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()
	clsService := ClsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
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
	return nil
}

func resourceTencentCloudClbInstanceTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_log_topic.update")()

	var (
		logId   = getLogId(contextNil)
		topicId = d.Id()
	)

	if d.HasChange("status") {
		if v, ok := d.GetOkExists("status"); ok {
			request := cls.NewModifyTopicRequest()
			request.TopicId = &topicId
			request.Status = helper.Bool(v.(bool))
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyTopic(request)
				if e != nil {
					return retryError(e)
				}
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				return nil
			})

			if err != nil {
				return err
			}
		}
	}

	return resourceTencentCloudClbInstanceTopicRead(d, meta)
}

func resourceTencentCloudClbInstanceTopicDelete(d *schema.ResourceData, meta interface{}) error {
	clsActionMu.Lock()
	defer clsActionMu.Unlock()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()
	clsService := ClsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := clsService.DeleteClsTopic(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
