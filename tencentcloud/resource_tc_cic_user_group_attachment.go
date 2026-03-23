/*
Provides a resource to create an identity center user group attachment

Example Usage

```hcl
resource "tencentcloudenterprise_cic_user_group_attachment" "cic_user_group_attachment" {
    zone_id = "z-xxxxxx"
    user_id = "u-xxxxxx"
    group_id = "g-xxxxxx"
}
```

Import

organization cic_user_group_attachment can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_user_group_attachment.cic_user_group_attachment ${zoneId}#${groupId}#${userId}
```

 */
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strings"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_user_group_attachment", CNDescription{
		TerraformTypeCN: "身份中心用户组关联",
		DescriptionCN:   "提供身份中心用户组关联资源，用于将用户添加到用户组。",
		AttributesCN: map[string]string{
			"zone_id":  "空间ID",
			"user_id":  "用户ID",
			"group_id": "用户组ID",
		},
	})
}

func resourceTencentCloudCicUserGroupAttachment() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create an identity center user group attachment",
		Create: resourceTencentCloudCicUserGroupAttachmentCreate,
		Read:   resourceTencentCloudCicUserGroupAttachmentRead,
		Delete: resourceTencentCloudCicUserGroupAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone id.",
			},

			"group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "User group ID.",
			},

			"user_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "User ID.",
			},
		},
	}
}

func resourceTencentCloudCicUserGroupAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user_group_attachment.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	var (
		zoneId  string
		groupId string
		userId  string
	)
	var (
		request  = cic.NewAddUserToGroupRequest()
		response = cic.NewAddUserToGroupResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}
	if v, ok := d.GetOk("group_id"); ok {
		groupId = v.(string)
	}
	if v, ok := d.GetOk("user_id"); ok {
		userId = v.(string)
	}

	if v, ok := d.GetOk("zone_id"); ok {
		request.ZoneId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("group_id"); ok {
		request.GroupId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("user_id"); ok {
		request.UserId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().AddUserToGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create identity center user group attachment failed, reason:%+v", logId, err)
		return err
	}

	_ = response

	d.SetId(strings.Join([]string{zoneId, groupId, userId}, FILED_SP))

	return resourceTencentCloudCicUserGroupAttachmentRead(d, meta)
}

func resourceTencentCloudCicUserGroupAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user_group_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	groupId := idSplit[1]
	userId := idSplit[2]

	respData, err := service.DescribeCicUserGroupAttachmentById(ctx, zoneId, groupId, userId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_user_group_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if respData.GroupId != nil {
		_ = d.Set("group_id", respData.GroupId)
	}

	_ = d.Set("zone_id", zoneId)
	_ = d.Set("user_id", userId)
	return nil
}

func resourceTencentCloudCicUserGroupAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user_group_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	groupId := idSplit[1]
	userId := idSplit[2]

	var (
		request  = cic.NewRemoveUserFromGroupRequest()
		response = cic.NewRemoveUserFromGroupResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		request.ZoneId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("group_id"); ok {
		request.GroupId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("user_id"); ok {
		request.UserId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().RemoveUserFromGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete identity center user group attachment failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	_ = zoneId
	_ = groupId
	_ = userId
	return nil
}
