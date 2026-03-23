/*
Provides a resource to create a CAM group membership.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_group_membership" "foo" {
	  group_id = tencentcloudenterprise_cam_group.foo.id
	  user_names = [tencentcloudenterprise_cam_user.foo.name, tencentcloudenterprise_cam_user.bar.name]
	}

```

# Import

CAM group membership can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_group_membership.foo 12515263
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_group_membership", CNDescription{
		TerraformTypeCN: "CAM用户组成员",
		DescriptionCN:   "提供 CAM 用户组成员资源，用于管理用户组内的子用户。",
		AttributesCN: map[string]string{
			"group_id":   "用户组 ID",
			"user_names": "要加入用户组的用户名列表",
		},
	})
}

func resourceTencentCloudCamGroupMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamGroupMembershipCreate,
		Read:   resourceTencentCloudCamGroupMembershipRead,
		Update: resourceTencentCloudCamGroupMembershipUpdate,
		Delete: resourceTencentCloudCamGroupMembershipDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of CAM group.",
			},
			"user_names": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User name set as ID of the CAM group members.",
			},
		},
	}
}

func resourceTencentCloudCamGroupMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group_membership.create")()

	logId := getLogId(contextNil)

	groupId := d.Get("group_id").(string)
	members, _, err := getUserIds(d)
	if err != nil {
		return err
	}
	err = addUsersToGroup(members.List(), groupId, meta)
	if err != nil {
		log.Printf("[CRITAL]%s create CAM group membership failed, reason:%s\n", logId, err.Error())
		return err
	}
	d.SetId(groupId)

	//get really instance then read
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instance, e := camService.DescribeGroupMembershipById(ctx, groupId)
		if e != nil {
			return retryError(e)
		}
		if len(instance) == 0 {
			return resource.RetryableError(fmt.Errorf("creation not done"))
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM group membership failed, reason:%s\n", logId, err.Error())
		return err
	}
	time.Sleep(10 * time.Second)
	return resourceTencentCloudCamGroupMembershipRead(d, meta)
}

func resourceTencentCloudCamGroupMembershipRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group_membership.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	groupId := d.Id()
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var members []*string
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := camService.DescribeGroupMembershipById(ctx, groupId)
		if e != nil {
			return retryError(e)
		}
		members = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM group membership failed, reason:%s\n", logId, err.Error())
		return err
	}

	if len(members) == 0 {
		d.SetId("")
		return nil
	}

	// Filter members based on what's in state
	stateMembers, _, err := getUserIds(d)
	if err != nil {
		stateMembers = &schema.Set{}
	}

	var memberResult []*string
	if stateMembers.Len() != 0 {
		// The old state exists - only include members that are in state
		exactMembers := make([]*string, 0)
		for _, v := range members {
			if stateMembers.Contains(*v) {
				exactMembers = append(exactMembers, v)
			}
		}
		memberResult = exactMembers
	} else {
		memberResult = members
	}

	_ = d.Set("user_names", memberResult)
	_ = d.Set("group_id", groupId)

	return nil
}

func resourceTencentCloudCamGroupMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group_membership.update")()

	logId := getLogId(contextNil)

	groupId := d.Id()

	if err := processChange(d, groupId, logId, meta); err != nil {
		return err
	}

	return resourceTencentCloudCamGroupMembershipRead(d, meta)
}

func resourceTencentCloudCamGroupMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group_membership.delete")()

	logId := getLogId(contextNil)
	groupId := d.Get("group_id").(string)
	userIds, _, err := getUserIds(d)
	if err != nil {
		return err
	}
	members := userIds.List()
	err = removeUsersFromGroup(members, groupId, meta)
	if err != nil {
		log.Printf("[CRITAL]%s delete CAM group failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}

func getUidFromName(name string, meta interface{}) (uid *string, errRet error) {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var uidUint64 *uint64
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := camService.DescribeUserById(ctx, name)
		if e != nil {
			return retryError(e)
		}
		if result == nil || result.Uid == nil {
			return nil
		}
		uidUint64 = result.Uid
		return nil
	})
	if err != nil {
		errRet = err
		return
	}
	if uidUint64 != nil {
		uidStr := strconv.FormatUint(*uidUint64, 10)
		uid = &uidStr
	}
	return
}

func addUsersToGroup(members []interface{}, groupId string, meta interface{}) error {
	logId := getLogId(contextNil)

	request := cam.NewAddUserToGroupRequest()
	// UpdateType: 1 = Add, 2 = Remove
	updateType := int64(1)
	request.UpdateType = &updateType
	request.Info = make([]*cam.GroupMember, 0)
	for _, member := range members {
		var info cam.GroupMember
		// Get uid from name
		uId, e := getUidFromName(member.(string), meta)
		if e != nil {
			return e
		}
		if uId == nil {
			continue
		}
		info.Uid = uId
		info.GroupId = &groupId
		request.Info = append(request.Info, &info)
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().AddUserToGroup(request)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), e.Error())
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create CAM group membership failed, reason:%s\n", logId, err.Error())
		return err
	}
	return nil
}

func removeUsersFromGroup(members []interface{}, groupId string, meta interface{}) error {
	logId := getLogId(contextNil)

	request := cam.NewRemoveUserFromGroupRequest()
	// UpdateType: 2 = Remove (according to SDK comment "传2")
	updateType := int64(2)
	request.UpdateType = &updateType
	request.Info = make([]*cam.GroupMember, 0)
	for _, member := range members {
		var info cam.GroupMember
		uId, e := getUidFromName(member.(string), meta)
		if e != nil {
			// Notice case when user is deleted, the uin is not found, and the membership is removed in the user module when deleted
			ee, ok := e.(*sdkErrors.CloudSDKError)
			if !ok {
				return e
			}
			if ee.Code == "ResourceNotFound.UserNotExist" {
				continue
			} else {
				return e
			}
		}
		if uId == nil {
			continue
		}
		info.Uid = uId
		info.GroupId = &groupId
		request.Info = append(request.Info, &info)
	}
	// No exist user need to remove, then return
	if len(request.Info) == 0 {
		return nil
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().RemoveUserFromGroup(request)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), e.Error())
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete CAM group membership failed, reason:%s\n", logId, err.Error())
		return err
	}
	return nil
}

func getUserIds(d *schema.ResourceData) (data *schema.Set, usingNames bool, errRet error) {
	names, hasNames := d.GetOk("user_names")
	if hasNames {
		return names.(*schema.Set), true, nil
	}
	return nil, true, fmt.Errorf("no user names provided")
}

func processChange(d *schema.ResourceData, groupId string, logId string, meta interface{}) error {
	if !d.HasChange("user_names") {
		return nil
	}

	o, n := d.GetChange("user_names")
	os := o.(*schema.Set)
	ns := n.(*schema.Set)
	add := ns.Difference(os).List()
	remove := os.Difference(ns).List()

	if len(remove) > 0 {
		oErr := removeUsersFromGroup(remove, groupId, meta)
		if oErr != nil {
			log.Printf("[CRITAL]%s update CAM group membership failed, reason:%s\n", logId, oErr.Error())
			return oErr
		}
	}
	if len(add) > 0 {
		nErr := addUsersToGroup(add, groupId, meta)
		if nErr != nil {
			log.Printf("[CRITAL]%s update CAM group membership failed, reason:%s\n", logId, nErr.Error())
			return nErr
		}
	}
	return nil
}
