/*
Provides a resource to create a vpc end_point

Example Usage

```hcl
resource "tencentcloudenterprise_vpc_end_point" "example" {
  vpc_id               = "vpc-ffwo6rid"
  subnet_id            = "subnet-o7v0wz10"
  end_point_name       = "123tf"
  end_point_service_id = "vpcsvc-o9u88lu5"
  end_point_vip        = "192.168.32.20"
  ip_address_type      = "IPv4"
  security_group_id    = "sg-iz7ipqme"
}
```

Import

vpc end_point can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_vpc_end_point.end_point end_point_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_end_point", CNDescription{
		TerraformTypeCN: "终端节点",
		DescriptionCN:   "提供VPC终端节点资源，用于创建VPC终端节点。",
		AttributesCN: map[string]string{
			"vpc_id":               "VPC ID",
			"subnet_id":            "子网ID",
			"end_point_name":       "终端名称",
			"end_point_service_id": "终端服务ID",
			"end_point_vip":        "终端VIP",
			"security_group_id":    "安全组ID",
			"ip_address_type":      "IP地址类型",
			"end_point_owner":      "终端拥有者",
			"state":                "终端状态",
			"create_time":          "创建时间",
		},
	})
}
func resourceTencentCloudVpcEndPoint() *schema.Resource {
	return &schema.Resource{
		Description: "Create vpc end_point",
		Create:      resourceTencentCloudVpcEndPointCreate,
		Read:        resourceTencentCloudVpcEndPointRead,
		Update:      resourceTencentCloudVpcEndPointUpdate,
		Delete:      resourceTencentCloudVpcEndPointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "ID of vpc instance.",
			},

			"subnet_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "ID of subnet instance.",
			},

			"end_point_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Name of endpoint.",
			},

			"end_point_service_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "ID of endpoint service.",
			},

			"end_point_vip": {
				Computed:	 true,
				Optional:    true,
				Type:        schema.TypeString,
				Description: "VIP of endpoint ip.",
			},

			"security_group_id": {
				Optional:    true,
				Type:        schema.TypeSet,
				Description: "List of security group IDs.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},

			"ip_address_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "IP address type: IPv4/IPv6.",
			},

			"end_point_owner": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "APPID.",
			},

			"state": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "State of end point.",
			},

			"create_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Create Time.",
			},
		},
	}
}

func resourceTencentCloudVpcEndPointCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_end_point.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = vpc.NewCreateVpcEndPointRequest()
		response   = vpc.NewCreateVpcEndPointResponse()
		endPointId string
	)
	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		request.SubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_point_name"); ok {
		request.EndPointName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_point_service_id"); ok {
		request.EndPointServiceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_point_vip"); ok {
		request.EndPointVip = helper.String(v.(string))
	}

	var sgIds []string
	if v, ok := d.GetOk("security_group_id"); ok {
		set := v.(*schema.Set)
		list := set.List()
		s := make([]string, len(list))
		for i, val := range list {
			s[i] = val.(string)
		}
		sgIds = s
		if len(sgIds) > 0 {
			request.SecurityGroupId = helper.String(sgIds[0])
		}
	}

	if v, ok := d.GetOk("ip_address_type"); ok {
		request.IpAddressType = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().CreateVpcEndPoint(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create vpc endPoint failed, reason:%+v", logId, err)
		return err
	}

	endPointId = *response.Response.EndPoint.EndPointId
	d.SetId(endPointId)

	if len(sgIds) > 1 {
		modifyAttrReq := vpc.NewModifyVpcEndPointAttributeRequest()
		modifyAttrReq.EndPointId = &endPointId

		slice := make([]*string, len(sgIds))
		for i, val := range sgIds {
			v := val
			slice[i] = &v
		}
		modifyAttrReq.SecurityGroupIds = slice

		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().ModifyVpcEndPointAttribute(modifyAttrReq)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, modifyAttrReq.GetAction(), modifyAttrReq.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s modify vpc endPoint security groups failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudVpcEndPointRead(d, meta)
}

func resourceTencentCloudVpcEndPointRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_end_point.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	endPointId := d.Id()

	endPoint, err := service.DescribeVpcEndPointById(ctx, endPointId)
	if err != nil {
		return err
	}

	if endPoint == nil {
		d.SetId("")
		return fmt.Errorf("resource `track` %s does not exist", d.Id())
	}

	if endPoint.VpcId != nil {
		_ = d.Set("vpc_id", endPoint.VpcId)
	}

	if endPoint.SubnetId != nil {
		_ = d.Set("subnet_id", endPoint.SubnetId)
	}

	if endPoint.EndPointName != nil {
		_ = d.Set("end_point_name", endPoint.EndPointName)
	}

	if endPoint.EndPointServiceId != nil {
		_ = d.Set("end_point_service_id", endPoint.EndPointServiceId)
	}

	if endPoint.EndPointVip != nil {
		_ = d.Set("end_point_vip", endPoint.EndPointVip)
	}

	if endPoint.EndPointOwner != nil {
		_ = d.Set("end_point_owner", endPoint.EndPointOwner)
	}

	if endPoint.State != nil {
		_ = d.Set("state", endPoint.State)
	}

	if endPoint.CreateTime != nil {
		_ = d.Set("create_time", endPoint.CreateTime)
	}

	if endPoint.IpAddressType != nil {
		_ = d.Set("ip_address_type", endPoint.IpAddressType)
	}

	if endPoint.GroupSet != nil {
		sgIds := make([]string, 0, len(endPoint.GroupSet))
		for _, sgIdPtr := range endPoint.GroupSet {
			if sgIdPtr != nil {
				sgIds = append(sgIds, *sgIdPtr)
			}
		}
		_ = d.Set("security_group_id", sgIds)
	}

	return nil
}

func resourceTencentCloudVpcEndPointUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_end_point.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := vpc.NewModifyVpcEndPointAttributeRequest()

	endPointId := d.Id()

	request.EndPointId = &endPointId

	unsupportedUpdateFields := []string{
		"vpc_id",
		"subnet_id",
		"end_point_service_id",
		"end_point_vip",
		"ip_address_type",
	}
	for _, field := range unsupportedUpdateFields {
		if d.HasChange(field) {
			return fmt.Errorf("tencentcloudenterprise_vpc_end_point update on %s is not support yet", field)
		}
	}

	if d.HasChange("end_point_name") || d.HasChange("security_group_id") {
		if d.HasChange("end_point_name") {
			if v, ok := d.GetOk("end_point_name"); ok {
				request.EndPointName = helper.String(v.(string))
			}
		}

		if d.HasChange("security_group_id") {
			if v, ok := d.GetOk("security_group_id"); ok {
				set := v.(*schema.Set)
				list := set.List()
				s := make([]string, len(list))
				for i, val := range list {
					s[i] = val.(string)
				}

				slice := make([]*string, len(s))
				for i, val := range s {
					v := val
					slice[i] = &v
				}
				request.SecurityGroupIds = slice
			}
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().ModifyVpcEndPointAttribute(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update vpc endPoint failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudVpcEndPointRead(d, meta)
}

func resourceTencentCloudVpcEndPointDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_end_point.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	endPointId := d.Id()

	if err := service.DeleteVpcEndPointById(ctx, endPointId); err != nil {
		return nil
	}

	return nil
}
