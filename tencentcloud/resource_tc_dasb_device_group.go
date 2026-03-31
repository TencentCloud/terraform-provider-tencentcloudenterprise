/*
Provide a resource to create a DASB device group

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_device_group" "example" {
	  name = "example-device-group"
	}

```

# Import

DASB device group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_device_group.example 12345
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_device_group", CNDescription{
		TerraformTypeCN: "DASB 设备组",
		DescriptionCN:   "提供 DASB 设备组资源，用于创建和管理设备组。",
		AttributesCN: map[string]string{
			"name":          "名称",
			"department_id": "部门ID",
		},
	})
}

func ResourceTencentCloudDasbDeviceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbDeviceGroupCreate,
		Read:   resourceTencentCloudDasbDeviceGroupRead,
		Update: resourceTencentCloudDasbDeviceGroupUpdate,
		Delete: resourceTencentCloudDasbDeviceGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Device group name, the maximum length is 32 characters.",
			},
			"department_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The ID of the department to which the asset group belongs, such as: 1.2.3 name, with a maximum length of 32 characters.",
			},
		},
	}
}

func resourceTencentCloudDasbDeviceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device_group.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId         = getLogId(contextNil)
		request       = bhsaas.NewCreateDeviceGroupRequest()
		response      = bhsaas.NewCreateDeviceGroupResponse()
		deviceGroupId string
	)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("department_id"); ok {
		request.DepartmentId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateDeviceGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response.Id == nil {
			e = fmt.Errorf("dasb DeviceGroup not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb DeviceGroup failed, reason:%+v", logId, err)
		return err
	}

	deviceGroupIdInt := *response.Response.Id
	deviceGroupId = strconv.FormatUint(deviceGroupIdInt, 10)
	d.SetId(deviceGroupId)

	return resourceTencentCloudDasbDeviceGroupRead(d, meta)
}

func resourceTencentCloudDasbDeviceGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device_group.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId         = getLogId(contextNil)
		ctx           = context.WithValue(context.TODO(), logIdKey, logId)
		service       = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		deviceGroupId = d.Id()
	)

	DeviceGroup, err := service.DescribeDasbDeviceGroupById(ctx, deviceGroupId)
	if err != nil {
		return err
	}

	if DeviceGroup == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DasbDeviceGroup` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if DeviceGroup.Name != nil {
		_ = d.Set("name", DeviceGroup.Name)
	}

	if DeviceGroup.Department != nil {
		departmentId := DeviceGroup.Department.Id
		if *departmentId != "1" {
			_ = d.Set("department_id", departmentId)
		}
	}

	return nil
}

func resourceTencentCloudDasbDeviceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device_group.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId         = getLogId(contextNil)
		request       = bhsaas.NewModifyDeviceGroupRequest()
		deviceGroupId = d.Id()
	)

	request.Id = helper.StrToUint64Point(deviceGroupId)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("department_id"); ok {
		request.DepartmentId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyDeviceGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update dasb DeviceGroup failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudDasbDeviceGroupRead(d, meta)
}

func resourceTencentCloudDasbDeviceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device_group.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId         = getLogId(contextNil)
		ctx           = context.WithValue(context.TODO(), logIdKey, logId)
		service       = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		deviceGroupId = d.Id()
	)

	if err := service.DeleteDasbDeviceGroupById(ctx, deviceGroupId); err != nil {
		return err
	}

	return nil
}
