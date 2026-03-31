/*
Provide a resource to create a DASB bind device resource

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_bind_device_resource" "example" {
	  device_id_set = [1, 2]
	  resource_id   = "example-resource-id"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_bind_device_resource", CNDescription{
		TerraformTypeCN: "DASB 绑定设备资源",
		DescriptionCN:   "提供 DASB 绑定设备资源，用于将设备绑定到堡垒机服务。",
		AttributesCN: map[string]string{
			"device_id_set": "设备ID集合",
			"resource_id":   "堡垒机服务ID",
			"domain_id":     "网络域ID",
		},
	})
}

func ResourceTencentCloudDasbBindDeviceResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbBindDeviceResourceCreate,
		Read:   resourceTencentCloudDasbBindDeviceResourceRead,
		Update: resourceTencentCloudDasbBindDeviceResourceUpdate,
		Delete: resourceTencentCloudDasbBindDeviceResourceDelete,

		Schema: map[string]*schema.Schema{
			"device_id_set": {
				Required:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Asset ID collection.",
			},
			"resource_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Bastion host service ID.",
			},
			"domain_id": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Network Domain ID.",
			},
		},
	}
}

func resourceTencentCloudDasbBindDeviceResourceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_bind_device_resource.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = bhsaas.NewBindDeviceResourceRequest()
		resourceId string
	)

	if v, ok := d.GetOk("device_id_set"); ok {
		deviceIdSetSet := v.(*schema.Set).List()
		for i := range deviceIdSetSet {
			deviceIdSet := deviceIdSetSet[i].(int)
			request.DeviceIdSet = append(request.DeviceIdSet, helper.IntUint64(deviceIdSet))
		}
	}

	if v, ok := d.GetOk("resource_id"); ok {
		request.ResourceId = helper.String(v.(string))
		resourceId = v.(string)
	}

	if v, ok := d.GetOk("domain_id"); ok {
		request.DomainId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().BindDeviceResource(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s operate dasb bindDeviceResource failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(resourceId)
	return resourceTencentCloudDasbBindDeviceResourceRead(d, meta)
}

func resourceTencentCloudDasbBindDeviceResourceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_bind_device_resource.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		resourceId = d.Id()
	)

	deviceSets, err := service.DescribeDasbDeviceByResourceId(ctx, resourceId)
	if err != nil {
		return err
	}

	if deviceSets == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DeviceResource` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("resource_id", resourceId)

	tmpList := make([]interface{}, 0, len(deviceSets))
	for _, item := range deviceSets {
		if item.Id != nil {
			tmpList = append(tmpList, item.Id)
		}

		if item.DomainId != nil {
			_ = d.Set("domain_id", item.DomainId)
		}
	}

	_ = d.Set("device_id_set", tmpList)

	return nil
}

func resourceTencentCloudDasbBindDeviceResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_bind_device_resource.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		resourceId = d.Id()
	)

	if d.HasChange("device_id_set") {
		oldInterface, newInterface := d.GetChange("device_id_set")
		olds := oldInterface.(*schema.Set)
		news := newInterface.(*schema.Set)
		remove := helper.InterfacesIntegers(olds.Difference(news).List())
		add := helper.InterfacesIntegers(news.Difference(olds).List())
		if len(remove) > 0 {
			request := bhsaas.NewBindDeviceResourceRequest()
			for _, item := range remove {
				request.DeviceIdSet = append(request.DeviceIdSet, helper.IntUint64(item))
			}

			request.ResourceId = helper.String("")
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().BindDeviceResource(request)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				}

				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s operate dasb bindDeviceResource failed, reason:%+v", logId, err)
				return err
			}
		}

		if len(add) > 0 {
			request := bhsaas.NewBindDeviceResourceRequest()
			for _, item := range add {
				request.DeviceIdSet = append(request.DeviceIdSet, helper.IntUint64(item))
			}

			request.ResourceId = helper.String(resourceId)
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().BindDeviceResource(request)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				}

				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s operate dasb bindDeviceResource failed, reason:%+v", logId, err)
				return err
			}
		}
	}

	return resourceTencentCloudDasbBindDeviceResourceRead(d, meta)
}

func resourceTencentCloudDasbBindDeviceResourceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_bind_device_resource.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = bhsaas.NewBindDeviceResourceRequest()
	)

	if v, ok := d.GetOk("device_id_set"); ok {
		deviceIdSetSet := v.(*schema.Set).List()
		for i := range deviceIdSetSet {
			deviceIdSet := deviceIdSetSet[i].(int)
			request.DeviceIdSet = append(request.DeviceIdSet, helper.IntUint64(deviceIdSet))
		}
	}

	request.ResourceId = helper.String("")
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().BindDeviceResource(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s operate dasb bindDeviceResource failed, reason:%+v", logId, err)
		return err
	}

	return nil
}
