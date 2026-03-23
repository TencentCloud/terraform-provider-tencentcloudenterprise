/*
Provides a resource to create a CCN attachment instance.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ccn_attachment_v2" "example" {
	  ccn_id          = "ccn-cwm743gl"
	  instance_id     = "vpc-l40swg8d"
	  instance_type   = "VPC"
	  instance_region = "ap-qingyuan-region-devtest-ops"
	  route_table_id  = "ccnrtb-gbaaugtl"
	  description     = "test ccn attachment description"
	}

```

# Import

CCN attachment instance can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ccn_attachment_v2.example ccn-cwm743gl#VPC#ap-qingyuan-region-devtest-ops#vpc-l40swg8d
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	ccn "terraform-provider-tencentcloudenterprise/sdk/ccn/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ccn_attachment_v2", CNDescription{
		TerraformTypeCN: "云联网关联实例",
		DescriptionCN:   "用于将网络实例关联到云联网，并可绑定指定路由表。",
		AttributesCN: map[string]string{
			"ccn_id":          "云联网实例 ID",
			"instance_id":     "被关联实例 ID",
			"instance_region": "被关联实例所属地域",
			"instance_type":   "被关联实例类型",
			"description":     "关联备注",
			"route_table_id":  "云联网路由表 ID",
			"ccn_uin":         "云联网所属账号 UIN",
			"state":           "关联状态",
			"attached_time":   "关联时间",
			"cidr_block":      "关联实例网段列表",
			"route_ids":       "云联网路由 ID 列表",
		},
	})
}

func resourceTencentCloudCcnAttachmentV2() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a CCN attachment instance.",
		Create:      resourceTencentCloudCcnAttachmentV2Create,
		Read:        resourceTencentCloudCcnAttachmentV2Read,
		Update:      resourceTencentCloudCcnAttachmentV2Update,
		Delete:      resourceTencentCloudCcnAttachmentV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ccn_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the CCN.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the attached instance.",
			},
			"instance_region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The region where the attached instance is located.",
			},
			"instance_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue([]string{CNN_INSTANCE_TYPE_VPC, CNN_INSTANCE_TYPE_DIRECTCONNECT, CNN_INSTANCE_TYPE_BMVPC, CNN_INSTANCE_TYPE_VPNGW}),
				Description:  "Type of the attached instance. Valid values: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Remark of the attachment.",
			},
			"route_table_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "CCN route table ID bound to the attachment.",
			},
			"ccn_uin": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "UIN of the attached CCN owner. When unset, the current account is used.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of the attachment.",
			},
			"attached_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Attachment time.",
			},
			"cidr_block": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "CIDR blocks of the attached instance.",
			},
			"route_ids": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Route ID list under the CCN.",
			},
		},
	}
}

func resourceTencentCloudCcnAttachmentV2Create(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_attachment_v2.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := ccn.NewAttachCcnInstancesRequest()
	ccnId := d.Get("ccn_id").(string)
	instanceType := d.Get("instance_type").(string)
	instanceRegion := d.Get("instance_region").(string)
	instanceId := d.Get("instance_id").(string)

	request.CcnId = helper.String(ccnId)
	instance := &ccn.CcnInstance{
		InstanceId:     helper.String(instanceId),
		InstanceRegion: helper.String(instanceRegion),
		InstanceType:   helper.String(instanceType),
	}
	if v, ok := d.GetOk("description"); ok {
		instance.Description = helper.String(v.(string))
	}
	if v, ok := d.GetOk("route_table_id"); ok {
		instance.RouteTableId = helper.String(v.(string))
	}
	request.Instances = []*ccn.CcnInstance{instance}

	if v, ok := d.GetOk("ccn_uin"); ok {
		request.CcnUin = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().AttachCcnInstances(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s attach ccn instance failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{ccnId, instanceType, instanceRegion, instanceId}, FILED_SP))
	return resourceTencentCloudCcnAttachmentV2Read(d, meta)
}

func resourceTencentCloudCcnAttachmentV2Read(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_attachment_v2.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	ccnId := idSplit[0]
	instanceType := idSplit[1]
	instanceRegion := idSplit[2]
	instanceId := idSplit[3]

	respData, err := service.DescribeCcnAttachedInstanceByFilter(ctx, ccnId, instanceType, instanceRegion, instanceId)
	if err != nil {
		return err
	}
	if respData == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_ccn_attachment_v2` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	_ = d.Set("ccn_id", ccnId)
	_ = d.Set("instance_id", instanceId)
	_ = d.Set("instance_region", instanceRegion)
	_ = d.Set("instance_type", instanceType)

	if respData.Description != nil {
		_ = d.Set("description", *respData.Description)
	}
	if respData.RouteTableId != nil {
		_ = d.Set("route_table_id", *respData.RouteTableId)
	}
	if respData.CcnUin != nil {
		_ = d.Set("ccn_uin", *respData.CcnUin)
	}
	if respData.State != nil {
		_ = d.Set("state", *respData.State)
	}
	if respData.AttachedTime != nil {
		_ = d.Set("attached_time", *respData.AttachedTime)
	}
	if respData.CidrBlock != nil {
		_ = d.Set("cidr_block", helper.StringsInterfaces(respData.CidrBlock))
	}

	request := ccn.NewDescribeCcnRoutesRequest()
	request.CcnId = helper.String(ccnId)
	routeIDs := make([]interface{}, 0)
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().DescribeCcnRoutes(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())

		routeIDs = routeIDs[:0]
		if result.Response != nil {
			for _, route := range result.Response.RouteSet {
				if route != nil && route.RouteId != nil {
					routeIDs = append(routeIDs, *route.RouteId)
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	_ = d.Set("route_ids", routeIDs)
	return nil
}

func resourceTencentCloudCcnAttachmentV2Update(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_attachment_v2.update")()
	defer inconsistentCheck(d, meta)()

	if !d.HasChange("description") {
		return resourceTencentCloudCcnAttachmentV2Read(d, meta)
	}

	logId := getLogId(contextNil)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	ccnId := idSplit[0]
	instanceType := idSplit[1]
	instanceRegion := idSplit[2]
	instanceId := idSplit[3]

	request := ccn.NewModifyCcnAttachedInstancesAttributeRequest()
	request.CcnId = helper.String(ccnId)

	instance := &ccn.CcnInstance{
		InstanceType:   helper.String(instanceType),
		InstanceRegion: helper.String(instanceRegion),
		InstanceId:     helper.String(instanceId),
	}
	if v, ok := d.GetOk("description"); ok {
		instance.Description = helper.String(v.(string))
	}
	if v, ok := d.GetOk("route_table_id"); ok {
		instance.RouteTableId = helper.String(v.(string))
	}
	request.Instances = []*ccn.CcnInstance{instance}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().ModifyCcnAttachedInstancesAttribute(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s modify ccn instance attribute failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCcnAttachmentV2Read(d, meta)
}

func resourceTencentCloudCcnAttachmentV2Delete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_attachment_v2.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	request := ccn.NewDetachCcnInstancesRequest()
	request.CcnId = helper.String(idSplit[0])
	request.Instances = []*ccn.CcnInstance{
		{
			InstanceType:   helper.String(idSplit[1]),
			InstanceRegion: helper.String(idSplit[2]),
			InstanceId:     helper.String(idSplit[3]),
		},
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().DetachCcnInstances(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s detach ccn instance failed, reason:%+v", logId, err)
		return err
	}

	return nil
}
