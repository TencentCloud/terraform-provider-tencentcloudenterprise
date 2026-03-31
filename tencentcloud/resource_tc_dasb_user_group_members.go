/*
Provide a resource to create a DASB user group members

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_user_group_members" "example" {
	  user_group_id = 1
	  member_id_set = [1, 2]
	}

```

# Import

DASB user group members can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_user_group_members.example 1#1,2
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_user_group_members", CNDescription{
		TerraformTypeCN: "DASB 用户组成员",
		DescriptionCN:   "提供 DASB 用户组成员资源，用于管理用户组的成员用户。",
		AttributesCN: map[string]string{
			"user_group_id": "用户组ID",
			"member_id_set": "成员ID集合",
		},
	})
}

func ResourceTencentCloudDasbUserGroupMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbUserGroupMembersCreate,
		Read:   resourceTencentCloudDasbUserGroupMembersRead,
		Delete: resourceTencentCloudDasbUserGroupMembersDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"user_group_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "User Group ID.",
			},
			"member_id_set": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Collection of member user IDs.",
			},
		},
	}
}

func resourceTencentCloudDasbUserGroupMembersCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_user_group_members.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId          = getLogId(contextNil)
		request        = bhsaas.NewAddUserGroupMembersRequest()
		userGroupId    string
		memberIdSetStr string
	)

	if v, ok := d.GetOkExists("user_group_id"); ok {
		request.Id = helper.IntUint64(v.(int))
		userGroupIdInt := v.(int)
		userGroupId = strconv.Itoa(userGroupIdInt)
	}

	if v, ok := d.GetOk("member_id_set"); ok {
		memberIdSetSet := v.(*schema.Set).List()
		tmpList := make([]string, 0)
		for i := range memberIdSetSet {
			memberIdSet := memberIdSetSet[i].(int)
			request.MemberIdSet = append(request.MemberIdSet, helper.IntUint64(memberIdSet))
			tmpList = append(tmpList, strconv.Itoa(memberIdSet))
		}

		memberIdSetStr = strings.Join(tmpList, COMMA_SP)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().AddUserGroupMembers(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil {
			e = fmt.Errorf("dasb UserGroupMembers not exists")
			return resource.NonRetryableError(e)
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb UserGroupMembers failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{userGroupId, memberIdSetStr}, FILED_SP))

	return resourceTencentCloudDasbUserGroupMembersRead(d, meta)
}

func resourceTencentCloudDasbUserGroupMembersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_user_group_members.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	userGroupId := idSplit[0]

	UserGroupMembers, err := service.DescribeDasbUserGroupMembersById(ctx, userGroupId)
	if err != nil {
		return err
	}

	if UserGroupMembers == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DasbUserGroupMembers` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	userGroupIdInt, _ := strconv.Atoi(userGroupId)
	_ = d.Set("user_group_id", userGroupIdInt)
	_ = d.Set("member_id_set", UserGroupMembers)

	return nil
}

func resourceTencentCloudDasbUserGroupMembersDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_user_group_members.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	userGroupId := idSplit[0]
	memberIdSetStr := idSplit[1]

	if err := service.DeleteDasbUserGroupMembersById(ctx, userGroupId, memberIdSetStr); err != nil {
		return err
	}

	return nil
}
