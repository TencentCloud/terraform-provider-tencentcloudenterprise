/*
Provides a resource to create a tag attachment

# Example Usage

```hcl

	resource "tencentcloudenterprise_tag_attachment" "example" {
	  tag_key   = "example-key"
	  tag_value = "example-value"
	  resource  = "qcs::cvm:ap-guangzhou:uin/123456789:instance/ins-xxxxxxxx"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	tag "terraform-provider-tencentcloudenterprise/sdk/tag/v20180813"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudTagAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudTagAttachmentCreate,
		Read:   resourceTencentCloudTagAttachmentRead,
		Delete: resourceTencentCloudTagAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"tag_key": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Tag key.",
			},

			"tag_value": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Tag value.",
			},

			"resource": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Six-segment description of resources.",
			},
		},
	}
}

func resourceTencentCloudTagAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tag_attachment.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = tag.NewModifyResourceTagsRequest()
		tagKey     string
		tagValue   string
		resourceId string
	)

	if v, ok := d.GetOk("tag_key"); ok {
		tagKey = v.(string)
	}

	if v, ok := d.GetOk("tag_value"); ok {
		tagValue = v.(string)
	}

	if v, ok := d.GetOk("resource"); ok {
		resourceId = v.(string)
		request.Resource = helper.String(v.(string))
	}

	request.ReplaceTags = []*tag.Tag{
		{
			TagKey:   helper.String(tagKey),
			TagValue: helper.String(tagValue),
		},
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTagClient().ModifyResourceTags(request)
		if e != nil {
			return resource.NonRetryableError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tag attachment failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(tagKey + FILED_SP + tagValue + FILED_SP + resourceId)

	return resourceTencentCloudTagAttachmentRead(d, meta)
}

func resourceTencentCloudTagAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tag_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TagService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	tagKey := idSplit[0]
	tagValue := idSplit[1]
	resourceId := idSplit[2]

	found, err := service.DescribeTagAttachmentById(ctx, tagKey, tagValue, resourceId)
	if err != nil {
		return err
	}

	if !found {
		d.SetId("")
		log.Printf("[WARN]%s resource `tencentcloudenterprise_tag_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("tag_key", tagKey)
	_ = d.Set("tag_value", tagValue)
	_ = d.Set("resource", resourceId)

	return nil
}

func resourceTencentCloudTagAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tag_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	tagKey := idSplit[0]
	resourceId := idSplit[2]

	if err := service.DeleteTagAttachmentById(ctx, tagKey, resourceId); err != nil {
		return err
	}

	return nil
}
