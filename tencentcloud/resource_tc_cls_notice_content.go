/*
Provides a resource to create a CLS notice content template.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_notice_content" "notice_content" {
	  name = "terraform-notice-content-test"
	  type = 0

	  notice_contents {
	    type = "Email"

	    trigger_content {
	      title   = "Trigger title"
	      content = "Trigger content"
	    }

	    recovery_content {
	      title   = "Recovery title"
	      content = "Recovery content"
	    }
	  }
	}

```

# Import

cls notice_content can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cls_notice_content.notice_content notice_content_id
```
*/
package tencentcloud

import (
	"context"
	"log"

	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_notice_content", CNDescription{
		TerraformTypeCN: "CLS通知内容模板",
		DescriptionCN:   "提供CLS通知内容模板资源，用于创建和管理日志服务通知内容模板。",
		AttributesCN: map[string]string{
			"name":              "通知内容模板名称",
			"type":              "通知内容语言。0：中文；1：英文",
			"notice_contents":   "通知内容模板详细信息",
			"trigger_content":   "告警触发通知内容模板",
			"recovery_content":  "告警恢复通知内容模板",
			"title":             "通知内容模板标题",
			"content":           "通知内容模板正文",
			"headers":           "请求头，仅自定义回调(Http)支持",
			"notice_content_id": "通知内容模板ID",
		},
	})
}

func resourceTencentCloudClsNoticeContent() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsNoticeContentCreate,
		Read:        resourceTencentCloudClsNoticeContentRead,
		Update:      resourceTencentCloudClsNoticeContentUpdate,
		Delete:      resourceTencentCloudClsNoticeContentDelete,
		Description: "Provides a resource to create and manage CLS notice content template",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Notice content template name.",
			},
			"type": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Notice content language. 0: Chinese; 1: English.",
			},
			"notice_contents": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Notice content template details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "Channel type. Valid values: Email, Sms, WeChat, Phone, WeCom, DingTalk, Lark, Http.",
						},
						"trigger_content": {
							Optional:    true,
							Type:        schema.TypeList,
							MaxItems:    1,
							Description: "Alarm trigger notice content.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"title": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Notice content title.",
									},
									"content": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Notice content body.",
									},
									"headers": {
										Type:        schema.TypeSet,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Request headers, supported only for Http callbacks.",
									},
								},
							},
						},
						"recovery_content": {
							Optional:    true,
							Type:        schema.TypeList,
							MaxItems:    1,
							Description: "Alarm recovery notice content.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"title": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Notice content title.",
									},
									"content": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Notice content body.",
									},
									"headers": {
										Type:        schema.TypeSet,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Request headers, supported only for Http callbacks.",
									},
								},
							},
						},
					},
				},
			},
			"notice_content_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Notice content template ID.",
			},
		},
	}
}

func resourceTencentCloudClsNoticeContentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_notice_content.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cls.NewCreateNoticeContentRequest()
	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}
	if v, ok := d.GetOkExists("type"); ok {
		request.Type = helper.IntUint64(v.(int))
	}
	if v, ok := d.GetOk("notice_contents"); ok {
		request.NoticeContents = expandNoticeContents(v.([]interface{}))
	}

	var response *cls.CreateNoticeContentResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateNoticeContent(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cls notice content failed, reason:%+v", logId, err)
		return err
	}

	noticeContentId := *response.Response.NoticeContentId
	d.SetId(noticeContentId)
	return resourceTencentCloudClsNoticeContentRead(d, meta)
}

func resourceTencentCloudClsNoticeContentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_notice_content.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	noticeContentId := d.Id()
	noticeContent, err := service.DescribeClsNoticeContentById(ctx, noticeContentId)
	if err != nil {
		return err
	}

	if noticeContent == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ClsNoticeContent` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if noticeContent.NoticeContentId != nil {
		_ = d.Set("notice_content_id", noticeContent.NoticeContentId)
	}

	if noticeContent.Name != nil {
		_ = d.Set("name", noticeContent.Name)
	}

	if noticeContent.Type != nil {
		_ = d.Set("type", int(*noticeContent.Type))
	}

	if noticeContent.NoticeContents != nil {
		_ = d.Set("notice_contents", flattenNoticeContents(noticeContent.NoticeContents))
	}

	return nil
}

func resourceTencentCloudClsNoticeContentUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_notice_content.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cls.NewModifyNoticeContentRequest()
	request.NoticeContentId = helper.String(d.Id())
	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}
	if v, ok := d.GetOkExists("type"); ok {
		request.Type = helper.IntUint64(v.(int))
	}
	if v, ok := d.GetOk("notice_contents"); ok {
		request.NoticeContents = expandNoticeContents(v.([]interface{}))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyNoticeContent(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		return err
	}

	return resourceTencentCloudClsNoticeContentRead(d, meta)
}

func resourceTencentCloudClsNoticeContentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_notice_content.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	if err := service.DeleteClsNoticeContentById(ctx, d.Id()); err != nil {
		return err
	}

	return nil
}

func expandNoticeContents(items []interface{}) []*cls.NoticeContent {
	results := make([]*cls.NoticeContent, 0, len(items))
	for _, item := range items {
		itemMap := item.(map[string]interface{})
		noticeContent := cls.NoticeContent{}
		if v, ok := itemMap["type"]; ok {
			noticeContent.Type = helper.String(v.(string))
		}
		if v, ok := itemMap["trigger_content"]; ok {
			noticeContent.TriggerContent = expandNoticeContentInfo(v)
		}
		if v, ok := itemMap["recovery_content"]; ok {
			noticeContent.RecoveryContent = expandNoticeContentInfo(v)
		}
		results = append(results, &noticeContent)
	}
	return results
}

func expandNoticeContentInfo(raw interface{}) *cls.NoticeContentInfo {
	if raw == nil {
		return nil
	}
	list, ok := raw.([]interface{})
	if !ok || len(list) == 0 {
		return nil
	}
	itemMap := list[0].(map[string]interface{})
	info := cls.NoticeContentInfo{}
	if v, ok := itemMap["title"]; ok {
		if title, ok := v.(string); ok && title != "" {
			info.Title = helper.String(title)
		}
	}
	if v, ok := itemMap["content"]; ok {
		info.Content = helper.String(v.(string))
	}
	if v, ok := itemMap["headers"]; ok {
		set := v.(*schema.Set)
		for _, header := range set.List() {
			info.Headers = append(info.Headers, helper.String(header.(string)))
		}
	}
	return &info
}

func flattenNoticeContents(contents []*cls.NoticeContent) []interface{} {
	results := make([]interface{}, 0, len(contents))
	for _, content := range contents {
		itemMap := map[string]interface{}{}
		if content.Type != nil {
			itemMap["type"] = *content.Type
		}
		if content.TriggerContent != nil {
			itemMap["trigger_content"] = flattenNoticeContentInfo(content.TriggerContent)
		}
		if content.RecoveryContent != nil {
			itemMap["recovery_content"] = flattenNoticeContentInfo(content.RecoveryContent)
		}
		results = append(results, itemMap)
	}
	return results
}

func flattenNoticeContentInfo(info *cls.NoticeContentInfo) []interface{} {
	if info == nil {
		return nil
	}
	itemMap := map[string]interface{}{}
	if info.Title != nil {
		itemMap["title"] = *info.Title
	}
	if info.Content != nil {
		itemMap["content"] = *info.Content
	}
	if info.Headers != nil {
		headers := make([]string, 0, len(info.Headers))
		for _, header := range info.Headers {
			if header != nil {
				headers = append(headers, *header)
			}
		}
		itemMap["headers"] = headers
	}
	return []interface{}{itemMap}
}
