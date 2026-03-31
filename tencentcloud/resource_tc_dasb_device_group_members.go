/*
Provide a resource to create a DASB device group members

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_device_group_members" "example" {
	  device_group_id = 1
	  member_id_set   = [1, 2]
	}

```

# Import

DASB device group members can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_device_group_members.example 1#1,2
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
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_device_group_members", CNDescription{
		TerraformTypeCN: "DASB 设备组成员",
		DescriptionCN:   "提供 DASB 设备组成员资源，用于管理设备组的成员设备。",
		AttributesCN: map[string]string{
			"device_group_id": "设备组ID",
			"member_id_set":   "成员ID集合",
		},
	})
}

func ResourceTencentCloudDasbDeviceGroupMembers() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbDeviceGroupMembersCreate,
		Read:   resourceTencentCloudDasbDeviceGroupMembersRead,
		Delete: resourceTencentCloudDasbDeviceGroupMembersDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"device_group_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "Device Group ID.",
			},
			"member_id_set": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "A collection of device IDs that need to be added to the device group.",
			},
		},
	}
}

func resourceTencentCloudDasbDeviceGroupMembersCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device_group_members.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId          = getLogId(contextNil)
		request        = bhsaas.NewAddDeviceGroupMembersRequest()
		deviceGroupId  string
		memberIdSetStr string
	)

	if v, ok := d.GetOkExists("device_group_id"); ok {
		request.Id = helper.IntUint64(v.(int))
		deviceGroupIdInt := v.(int)
		deviceGroupId = strconv.Itoa(deviceGroupIdInt)
	}

	if v, ok := d.GetOk("member_id_set"); ok {
		memberIdSetList := v.(*schema.Set).List()
		tmpList := make([]string, 0)
		for i := range memberIdSetList {
			memberIdSet := memberIdSetList[i].(int)
			request.MemberIdSet = append(request.MemberIdSet, helper.IntUint64(memberIdSet))
			tmpList = append(tmpList, strconv.Itoa(memberIdSet))
		}

		memberIdSetStr = strings.Join(tmpList, COMMA_SP)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().AddDeviceGroupMembers(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil {
			e = fmt.Errorf("dasb DeviceGroupMembers not exists")
			return resource.NonRetryableError(e)
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb DeviceGroupMembers failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{deviceGroupId, memberIdSetStr}, FILED_SP))

	return resourceTencentCloudDasbDeviceGroupMembersRead(d, meta)
}

func resourceTencentCloudDasbDeviceGroupMembersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device_group_members.read")()
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
	deviceGroupId := idSplit[0]

	DeviceGroupMembers, err := service.DescribeDasbDeviceGroupMembersById(ctx, deviceGroupId)
	if err != nil {
		return err
	}

	if DeviceGroupMembers == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DasbDeviceGroupMembers` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	deviceGroupIdInt, _ := strconv.Atoi(deviceGroupId)
	_ = d.Set("device_group_id", deviceGroupIdInt)
	_ = d.Set("member_id_set", DeviceGroupMembers)

	return nil
}

func resourceTencentCloudDasbDeviceGroupMembersDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device_group_members.delete")()
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
	deviceGroupId := idSplit[0]
	memberIdSetStr := idSplit[1]

	if err := service.DeleteDasbDeviceGroupMembersById(ctx, deviceGroupId, memberIdSetStr); err != nil {
		return err
	}

	return nil
}
