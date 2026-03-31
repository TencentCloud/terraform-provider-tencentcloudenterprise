/*
Provide a resource to create a DASB device

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_device" "example" {
	  os_name = "Linux"
	  ip      = "192.168.1.1"
	  port    = 22
	}

```

# Import

DASB device can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_device.example 12345
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
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_device", CNDescription{
		TerraformTypeCN: "DASB 设备",
		DescriptionCN:   "提供 DASB 设备资源，用于创建和管理设备。",
		AttributesCN: map[string]string{
			"os_name":       "操作系统名称",
			"ip":            "IP地址",
			"port":          "管理端口",
			"name":          "主机名",
			"department_id": "部门ID",
			"ip_port_set":   "资产多节点",
		},
	})
}

func ResourceTencentCloudDasbDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbDeviceCreate,
		Read:   resourceTencentCloudDasbDeviceRead,
		Update: resourceTencentCloudDasbDeviceUpdate,
		Delete: resourceTencentCloudDasbDeviceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"os_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue(OS_NAME),
				Description:  "Operating system name, only Linux, Windows or MySQL.",
			},
			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address.",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Management port.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hostname, can be empty.",
			},
			"department_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The department ID to which the device belongs.",
			},
			"ip_port_set": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: "Asset multi-node: fields ip and port.",
			},
		},
	}
}

func resourceTencentCloudDasbDeviceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = bhsaas.NewImportExternalDeviceRequest()
		response = bhsaas.NewImportExternalDeviceResponse()
		deviceId string
	)

	externalDevice := bhsaas.ExternalDevice{}
	if v, ok := d.GetOk("os_name"); ok {
		externalDevice.OsName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("ip"); ok {
		externalDevice.Ip = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("port"); ok {
		externalDevice.Port = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("name"); ok {
		externalDevice.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("department_id"); ok {
		externalDevice.DepartmentId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("ip_port_set"); ok {
		ipPortSetSet := v.(*schema.Set).List()
		for i := range ipPortSetSet {
			ipPortSet := ipPortSetSet[i].(string)
			externalDevice.IpPortSet = append(externalDevice.IpPortSet, &ipPortSet)
		}
	}

	request.DeviceSet = append(request.DeviceSet, &externalDevice)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ImportExternalDevice(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || len(result.Response.DeviceIdSet) != 1 {
			e = fmt.Errorf("dasb device not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb device failed, reason:%+v", logId, err)
		return err
	}

	deviceIdInt := *response.Response.DeviceIdSet[0]
	deviceId = strconv.FormatUint(deviceIdInt, 10)
	d.SetId(deviceId)

	return resourceTencentCloudDasbDeviceRead(d, meta)
}

func resourceTencentCloudDasbDeviceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		ctx      = context.WithValue(context.TODO(), logIdKey, logId)
		service  = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		deviceId = d.Id()
	)

	device, err := service.DescribeDasbDeviceById(ctx, deviceId)
	if err != nil {
		return err
	}

	if device == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DasbDevice` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if device != nil {
		if device.OsName != nil {
			_ = d.Set("os_name", device.OsName)
		}

		if device.PrivateIp != nil {
			_ = d.Set("ip", device.PrivateIp)
		}

		if device.Port != nil {
			_ = d.Set("port", device.Port)
		}

		if device.Name != nil {
			_ = d.Set("name", device.Name)
		}

		if device.Department != nil || device.Department.Id != nil {
			_ = d.Set("department_id", device.Department.Id)
		}

		if device.IpPortSet != nil {
			_ = d.Set("ip_port_set", device.IpPortSet)
		}
	}

	return nil
}

func resourceTencentCloudDasbDeviceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = bhsaas.NewModifyDeviceRequest()
		deviceId = d.Id()
	)

	immutableArgs := []string{"device_set", "os_name", "ip", "name", "ip_port_set"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	request.Id = helper.StrToUint64Point(deviceId)
	if d.HasChange("port") {
		if v, ok := d.GetOkExists("port"); ok {
			request.Port = helper.IntUint64(v.(int))
		}
	}

	if d.HasChange("department_id") {
		if v, ok := d.GetOkExists("department_id"); ok {
			request.DepartmentId = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyDevice(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update dasb device failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudDasbDeviceRead(d, meta)
}

func resourceTencentCloudDasbDeviceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_device.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		ctx      = context.WithValue(context.TODO(), logIdKey, logId)
		service  = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		deviceId = d.Id()
	)

	if err := service.DeleteDasbDeviceById(ctx, deviceId); err != nil {
		return err
	}

	return nil
}
