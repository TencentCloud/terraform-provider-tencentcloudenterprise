/*
Provides a resource to create an identity center group

Example Usage

```hcl
resource "tencentcloudenterprise_cic_group" "cic_group" {
    zone_id = "z-xxxxxx"
    group_name = "test-group"
    description = "test"
}
```

Import

tencentcloudenterprise_cic_group can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_group.cic_group ${zoneId}#${groupId}
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
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_group", CNDescription{
		TerraformTypeCN: "身份中心用户组",
		DescriptionCN:   "提供身份中心用户组资源，用于创建和管理身份中心用户组。",
		AttributesCN: map[string]string{
			"zone_id":     "空间ID",
			"group_name":  "用户组名称",
			"description": "用户组描述",
			"group_type":  "用户组类型",
			"group_id":    "用户组ID",
			"create_time": "创建时间",
			"update_time": "更新时间",
			"member_count": "成员数量",
		},
	})
}

func resourceTencentCloudCicGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create an identity center group",
		Create: resourceTencentCloudCicGroupCreate,
		Read:   resourceTencentCloudCicGroupRead,
		Update: resourceTencentCloudCicGroupUpdate,
		Delete: resourceTencentCloudCicGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone id.",
			},

			"group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the user group. Format: Allow English letters, numbers and special characters-. Length: Maximum 128 characters.",
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the user group.",
			},
			"group_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Type of user group. `Manual`: manual creation, `Synchronized`: external import.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Modification time for the user group.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the user group.",
			},
			"member_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of team members.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "A description of the user group. Length: Maximum 1024 characters.",
			},
		},
	}
}

func resourceTencentCloudCicGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_group.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	var (
		zoneId  string
		groupId string
	)
	var (
		request  = cic.NewCreateGroupRequest()
		response = cic.NewCreateGroupResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}

	if v, ok := d.GetOk("zone_id"); ok {
		request.ZoneId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("group_name"); ok {
		request.GroupName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOk("group_type"); ok {
		request.GroupType = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().CreateGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create identity center group failed, reason:%+v", logId, err)
		return err
	}

	groupId = *response.Response.GroupInfo.GroupId

	d.SetId(strings.Join([]string{zoneId, groupId}, FILED_SP))

	return resourceTencentCloudCicGroupRead(d, meta)
}

func resourceTencentCloudCicGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	groupId := idSplit[1]

	_ = d.Set("zone_id", zoneId)

	respData, err := service.DescribeCicGroupById(ctx, zoneId, groupId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_group` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if respData.GroupName != nil {
		_ = d.Set("group_name", respData.GroupName)
	}

	if respData.Description != nil {
		_ = d.Set("description", respData.Description)
	}

	if respData.CreateTime != nil {
		_ = d.Set("create_time", respData.CreateTime)
	}

	if respData.GroupType != nil {
		_ = d.Set("group_type", respData.GroupType)
	}

	if respData.UpdateTime != nil {
		_ = d.Set("update_time", respData.UpdateTime)
	}

	if respData.GroupId != nil {
		_ = d.Set("group_id", respData.GroupId)
	}

	if respData.MemberCount != nil {
		_ = d.Set("member_count", respData.MemberCount)
	}

	return nil
}

func resourceTencentCloudCicGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_group.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	immutableArgs := []string{"zone_id", "group_type"}
	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	groupId := idSplit[1]

	needChange := false
	mutableArgs := []string{"group_name", "description"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := cic.NewUpdateGroupRequest()

		request.ZoneId = helper.String(zoneId)

		request.GroupId = helper.String(groupId)

		if v, ok := d.GetOk("group_name"); ok {
			request.NewGroupName = helper.String(v.(string))
		}

		if v, ok := d.GetOk("description"); ok {
			request.NewDescription = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().UpdateGroup(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update identity center group failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCicGroupRead(d, meta)
}

func resourceTencentCloudCicGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_group.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	groupId := idSplit[1]

	var (
		request  = cic.NewDeleteGroupRequest()
		response = cic.NewDeleteGroupResponse()
	)

	request.ZoneId = helper.String(zoneId)

	request.GroupId = helper.String(groupId)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().DeleteGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete identity center group failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	return nil
}
