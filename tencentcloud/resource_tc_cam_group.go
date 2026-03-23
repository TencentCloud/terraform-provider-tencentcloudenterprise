/*
Provides a resource to create a CAM group.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_group" "foo" {
	  name   = "cam-group-test"
	  remark = "test"
	}

```

# Import

CAM group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_group.foo 90496
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"log"
	"strconv"
	"strings"
	"time"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_group", CNDescription{
		TerraformTypeCN: "CAM用户组",
		DescriptionCN:   "提供 CAM 用户组资源，用于创建和管理用户组。",
		AttributesCN: map[string]string{
			"name":        "用户组名称",
			"remark":      "备注",
			"channel":     "消息接收渠道",
			"create_time": "创建时间",
			"group_type":  "用户组类型",
		},
	})
}

func resourceTencentCloudCamGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamGroupCreate,
		Read:   resourceTencentCloudCamGroupRead,
		Update: resourceTencentCloudCamGroupUpdate,
		Delete: resourceTencentCloudCamGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of CAM group.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the CAM group.",
			},
			"channel": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Message receiving channel. 1: SMS, 2: Email, 3: SMS+Email.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time of the CAM group.",
			},
			"group_type": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Group type. 0: custom group, 1: preset group.",
			},
		},
	}
}

func resourceTencentCloudCamGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group.create")()

	logId := getLogId(contextNil)

	request := cam.NewCreateGroupRequest()
	// For creating a new group, GroupId must be set to -1 according to the internal API specification
	newGroupId := int64(-1)
	request.GroupId = &newGroupId
	request.GroupName = helper.String(d.Get("name").(string))
	if v, ok := d.GetOk("remark"); ok {
		request.Remark = helper.String(v.(string))
	}
	// Note: Channel is set by the API and cannot be configured by users

	var response *cam.CreateGroupResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().CreateGroup(request)
		if e != nil {
			if ee, ok := e.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				//check if read empty
				if strings.Contains(errCode, "GroupNameInUse") {
					return resource.NonRetryableError(e)
				}
			}
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), e.Error())
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create CAM group failed, reason:%s\n", logId, err.Error())
		return err
	}
	if response.Response.GroupId == nil {
		return fmt.Errorf("CAM group id is nil")
	}
	d.SetId(strconv.Itoa(int(*response.Response.GroupId)))

	//get really instance then read
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	groupId := d.Id()
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instance, e := camService.DescribeGroupById(ctx, groupId)
		if e != nil {
			return retryError(e)
		}
		if instance == nil || instance.GroupId == nil {
			return resource.RetryableError(fmt.Errorf("creation not done"))
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM group failed, reason:%s\n", logId, err.Error())
		return err
	}
	time.Sleep(10 * time.Second)
	return resourceTencentCloudCamGroupRead(d, meta)
}

func resourceTencentCloudCamGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	groupId := d.Id()
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var instance *cam.GroupInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := camService.DescribeGroupById(ctx, groupId)
		if e != nil {
			return retryError(e)
		}
		instance = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM group failed, reason:%s\n", logId, err.Error())
		return err
	}

	if instance == nil || instance.GroupId == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("name", *instance.GroupName)
	_ = d.Set("create_time", *instance.CreateTime)
	if instance.Remark != nil {
		_ = d.Set("remark", *instance.Remark)
	}
	if instance.Channel != nil {
		_ = d.Set("channel", *instance.Channel)
	}
	if instance.GroupType != nil {
		_ = d.Set("group_type", *instance.GroupType)
	}
	return nil
}

func resourceTencentCloudCamGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group.update")()

	logId := getLogId(contextNil)

	groupId := d.Id()
	groupIdInt64, e := strconv.ParseInt(groupId, 10, 64)
	if e != nil {
		return e
	}
	request := cam.NewUpdateGroupRequest()
	request.GroupId = &groupIdInt64
	changeFlag := false

	if d.HasChange("remark") {
		changeFlag = true
	}
	if d.HasChange("name") {
		changeFlag = true
	}

	if changeFlag {
		// Always send all fields in the update request to prevent API from clearing unspecified fields
		request.GroupName = helper.String(d.Get("name").(string))
		request.Remark = helper.String(d.Get("remark").(string))
		// Note: Channel is read-only and managed by the API

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			response, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().UpdateGroup(request)

			if e != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, request.GetAction(), request.ToJsonString(), e.Error())
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update CAM group description failed, reason:%s\n", logId, err.Error())
			return err
		}
	}

	return resourceTencentCloudCamGroupRead(d, meta)
}

func resourceTencentCloudCamGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group.delete")()

	logId := getLogId(contextNil)

	groupId := d.Id()
	groupIdInt64, e := strconv.ParseInt(groupId, 10, 64)
	if e != nil {
		return e
	}
	request := cam.NewDeleteGroupRequest()
	request.GroupId = &groupIdInt64
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().DeleteGroup(request)
		if e != nil {
			log.Printf("[CRITAL]%s reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete CAM group failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}
