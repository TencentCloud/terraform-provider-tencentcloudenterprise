/*
Provides a resource to create a cls logset

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_logset" "logset" {
	  logset_name = "demo"
	  tags = {
	    "createdBy" = "terraform"
	  }
	}

```
Import

cls logset can be imported using the id, e.g.
```
$ terraform import cloud_cls_logset.logset logset_id
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
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_logset", CNDescription{
		TerraformTypeCN: "CLS日志集",
		DescriptionCN:   "提供CLS日志集资源，用于创建和管理日志服务日志集。",
		AttributesCN: map[string]string{
			"logset_name": "日志集名称",
			"period":      "生命周期，单位天；可取值范围1～366",
			"logset_id":   "作为入参：日志集ID，格式为：用户自定义部分-用户appid，用户自定义部分仅支持小写字母、数字和-，且不能以-开头和结尾，长度为3至40字符，尾部需要使用-拼接用户appid；作为出参：日志集ID",
			"tags":        "标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。",
		},
	})
}

func resourceTencentCloudClsLogset() *schema.Resource {
	return &schema.Resource{
		Read:        resourceTencentCloudClsLogsetRead,
		Create:      resourceTencentCloudClsLogsetCreate,
		Update:      resourceTencentCloudClsLogsetUpdate,
		Delete:      resourceTencentCloudClsLogsetDelete,
		Description: "Provides a resource to create and manage CLS logset.",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"logset_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Logset name.",
			},

			"period": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Lifecycle in days. Value range: 1-366.",
			},

			"logset_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Logset ID. Format: {user-defined-part}-{user-appid}. The user-defined part only supports lowercase letters, numbers and '-', cannot start or end with '-', and the length is 3 to 40 characters. The user appid needs to be appended at the end.",
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "List of tag descriptions. By specifying this parameter, tags can be bound to the corresponding logset. A maximum of 10 tag key-value pairs is supported, and each tag key must be unique to the resource.",
			},
		},
	}
}

func resourceTencentCloudClsLogsetCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_logset.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cls.NewCreateLogsetRequest()

	if v, ok := d.GetOk("logset_name"); ok {
		request.LogsetName = helper.String(v.(string))
	}
	if v, ok := d.GetOk("period"); ok {
		request.Period = helper.Int64(int64(v.(int)))
	}
	if v, ok := d.GetOk("logset_id"); ok {
		request.LogsetId = helper.String(v.(string))
	}

	var response *cls.CreateLogsetResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateLogset(request)
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
		log.Printf("[CRITAL]%s create cls logset failed, reason:%+v", logId, err)
		return err
	}

	var logsetId string
	if v, ok := d.GetOk("logset_id"); ok {
		logsetId = v.(string)
	} else {
		logsetId = *response.Response.LogsetId
		_ = d.Set("logset_id", logsetId)
	}
	d.SetId(logsetId)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::cls:%s:uin/:logset/%s", region, logsetId)
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudClsLogsetRead(d, meta)
}

func resourceTencentCloudClsLogsetRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_logset.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	logsetId := d.Id()

	logset, err := service.DescribeClsLogset(ctx, logsetId)

	if err != nil {
		return err
	}

	if logset == nil {
		d.SetId("")
		return fmt.Errorf("resource `logset` %s does not exist", logsetId)
	}

	if err := d.Set("logset_name", logset.LogsetName); err != nil {
		return err
	}
	if err := d.Set("logset_id", logset.LogsetId); err != nil {
		return err
	}
	if err := d.Set("period", logset.Period); err != nil {
		return err
	}

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := &TagService{client: tcClient}
	tags, err := tagService.DescribeResourceTags(ctx, "cls", "logset", tcClient.Region, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudClsLogsetUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_logset.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	request := cls.NewModifyLogsetRequest()

	request.LogsetId = helper.String(d.Id())

	if d.HasChange("logset_name") {
		if v, ok := d.GetOk("logset_name"); ok {
			request.LogsetName = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyLogset(request)
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
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := &TagService{client: tcClient}
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("cls", "logset", tcClient.Region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudClsLogsetRead(d, meta)
}

func resourceTencentCloudClsLogsetDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_logset.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	logsetId := d.Id()

	if err := service.DeleteClsLogsetById(ctx, logsetId); err != nil {
		return err
	}

	return nil
}
