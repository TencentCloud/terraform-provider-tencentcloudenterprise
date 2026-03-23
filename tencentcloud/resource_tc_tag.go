/*
Provides a resource to create a tag

# Example Usage

```hcl

	resource "tencentcloudenterprise_tag" "example" {
	  tag_key   = "example-key"
	  tag_value = "example-value"
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

func resourceTencentCloudTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudTagCreate,
		Read:   resourceTencentCloudTagRead,
		Delete: resourceTencentCloudTagDelete,
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
		},
	}
}

func resourceTencentCloudTagCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tag.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = tag.NewCreateTagRequest()
		tagKey   string
		tagValue string
	)

	if v, ok := d.GetOk("tag_key"); ok {
		tagKey = v.(string)
		request.TagKey = helper.String(v.(string))
	}

	if v, ok := d.GetOk("tag_value"); ok {
		tagValue = v.(string)
		request.TagValue = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTagClient().CreateTag(request)
		if e != nil {
			return resource.NonRetryableError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tag failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(tagKey + FILED_SP + tagValue)
	return resourceTencentCloudTagRead(d, meta)
}

func resourceTencentCloudTagRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tag.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	tagKey := idSplit[0]
	tagValue := idSplit[1]

	service := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	tagRes, err := service.DescribeTagById(ctx, tagKey, tagValue)
	if err != nil {
		return err
	}

	if tagRes == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_tag` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	if tagRes.TagKey != nil {
		_ = d.Set("tag_key", tagRes.TagKey)
	}

	if tagRes.TagValue != nil {
		_ = d.Set("tag_value", tagRes.TagValue)
	}

	return nil
}

func resourceTencentCloudTagDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tag.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TagService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	tagKey := idSplit[0]
	tagValue := idSplit[1]

	if err := service.DeleteTagById(ctx, tagKey, tagValue); err != nil {
		return err
	}

	return nil
}
