/*
Provides a resource to create a vpc end_point_service

Example Usage

```hcl
resource "tencentcloudenterprise_vpc_end_point_service" "end_point_service" {
  vpc_id = "vpc-391sv4w3"
  end_point_service_name = "terraform-endpoint-service"
  auto_accept_flag = false
  service_instance_id = "lb-o5f6x7ke"
}
```

Import

vpc end_point_service can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_vpc_end_point_service.end_point_service end_point_service_id
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
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_end_point_service", CNDescription{
		TerraformTypeCN: "终端节点服务",
		DescriptionCN:   "提供VPC终端节点服务资源，用于创建VPC终端节点服务。",
		AttributesCN: map[string]string{
			"vpc_id":                 "VPC ID",
			"end_point_service_name": "终端节点服务名称",
			"auto_accept_flag":       "是否自动接受",
			"service_instance_id":    "服务实例ID",
			"ip_address_type":        "IP地址类型: IPv4/IPv6",
			"end_point_service":      "终端节点服务对象详细信息",
			"end_point_service_id":   "终端节点服务ID",
			"service_owner":          "用户的APPID",
			"service_name":           "终端节点服务名称",
			"service_vip":            "后端服务的VIP",
			"end_point_count":        "关联的终端节点个数",
			"end_point_id":           "终端节点ID",
			"subnet_id":              "子网ID",
			"end_point_owner":        "用户的APPID",
			"end_point_name":         "终端节点名称",
			"service_vpc_id":         "终端节点服务的VPCID",
			"end_point_vip":          "终端节点的VIP",
			"state":                  "终端节点状态，ACTIVE：可用，PENDING：待接受，ACCEPTING：接受中，REJECTED：已拒绝，FAILED：失败",
			"create_time":            "创建时间",
			"group_set":              "终端节点绑定的安全组实例ID列表",
		},
	})
}
func resourceTencentCloudVpcEndPointService() *schema.Resource {
	return &schema.Resource{
		Description: "Create vpc end_point_service",
		Create:      resourceTencentCloudVpcEndPointServiceCreate,
		Read:        resourceTencentCloudVpcEndPointServiceRead,
		Update:      resourceTencentCloudVpcEndPointServiceUpdate,
		Delete:      resourceTencentCloudVpcEndPointServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "ID of vpc instance.",
			},

			"end_point_service_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Name of end point service.",
			},

			"auto_accept_flag": {
				Required:    true,
				Type:        schema.TypeBool,
				Description: "Whether to automatically accept.",
			},

			"service_instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Id of service instance, like lb-xxx.",
			},

			"ip_address_type": {
				Optional: 	 true,
				Type:        schema.TypeString,
				Description: "Type of the IP address: IPv4/IPv6.",
			},

			"end_point_service": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "End point service details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_point_service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "End point service ID.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC ID.",
						},
						"service_owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service owner (APPID).",
						},
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "End point service name.",
						},
						"service_vip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backend service VIP.",
						},
						"service_instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backend service ID, like lb-xxx.",
						},
						"auto_accept_flag": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to automatically accept.",
						},
						"end_point_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of associated end points.",
						},
						"end_point_service": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "End point object array.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_point_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "End point ID.",
									},
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "VPC ID.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Subnet ID.",
									},
									"end_point_owner": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "End point owner (APPID).",
									},
									"end_point_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "End point name.",
									},
									"service_vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "End point service VPC ID.",
									},
									"service_vip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "End point service VIP.",
									},
									"end_point_service_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "End point service ID.",
									},
									"end_point_vip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "End point VIP.",
									},
									"state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "End point state: ACTIVE, PENDING, ACCEPTING, REJECTED, FAILED.",
									},
									"create_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Create time.",
									},
									"group_set": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Security group instance ID list bound to end point.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"service_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "End point service name.",
									},
								},
							},
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"ip_address_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of IP address: IPv4/IPv6.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudVpcEndPointServiceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_end_point_service.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request           = vpc.NewCreateVpcEndPointServiceRequest()
		response          = vpc.NewCreateVpcEndPointServiceResponse()
		endPointServiceId string
	)
	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_point_service_name"); ok {
		request.EndPointServiceName = helper.String(v.(string))
	}

	if v, _ := d.GetOk("auto_accept_flag"); v != nil {
		request.AutoAcceptFlag = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("service_instance_id"); ok {
		request.ServiceInstanceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("ip_address_type"); ok {
		request.IpAddressType = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().CreateVpcEndPointService(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create vpc endPointService failed, reason:%+v", logId, err)
		return err
	}

	endPointServiceId = *response.Response.EndPointService.EndPointServiceId
	d.SetId(endPointServiceId)

	return resourceTencentCloudVpcEndPointServiceRead(d, meta)
}

func resourceTencentCloudVpcEndPointServiceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_end_point_service.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	endPointServiceId := d.Id()

	endPointService, err := service.DescribeVpcEndPointServiceById(ctx, endPointServiceId)

	if err != nil {
		return err
	}

	if endPointService == nil {
		d.SetId("")
		return fmt.Errorf("resource `track` %s does not exist", d.Id())
	}

	if endPointService.VpcId != nil {
		_ = d.Set("vpc_id", *endPointService.VpcId)
	}

	if endPointService.ServiceName != nil {
		_ = d.Set("end_point_service_name", *endPointService.ServiceName)
	}

	if endPointService.AutoAcceptFlag != nil {
		_ = d.Set("auto_accept_flag", *endPointService.AutoAcceptFlag)
	}

	if endPointService.ServiceInstanceId != nil {
		_ = d.Set("service_instance_id", *endPointService.ServiceInstanceId)
	}

	// Safely build the end_point_service object
	endPointServiceMap := make(map[string]interface{})
	if endPointService.EndPointServiceId != nil {
		endPointServiceMap["end_point_service_id"] = *endPointService.EndPointServiceId
	}
	if endPointService.VpcId != nil {
		endPointServiceMap["vpc_id"] = *endPointService.VpcId
	}
	if endPointService.ServiceOwner != nil {
		endPointServiceMap["service_owner"] = *endPointService.ServiceOwner
	}
	if endPointService.ServiceName != nil {
		endPointServiceMap["service_name"] = *endPointService.ServiceName
	}
	if endPointService.ServiceVip != nil {
		endPointServiceMap["service_vip"] = *endPointService.ServiceVip
	}
	if endPointService.ServiceInstanceId != nil {
		endPointServiceMap["service_instance_id"] = *endPointService.ServiceInstanceId
	}
	if endPointService.AutoAcceptFlag != nil {
		endPointServiceMap["auto_accept_flag"] = *endPointService.AutoAcceptFlag
	}
	if endPointService.EndPointCount != nil {
		endPointServiceMap["end_point_count"] = int(*endPointService.EndPointCount)
	}
	if endPointService.CreateTime != nil {
		endPointServiceMap["create_time"] = *endPointService.CreateTime
	}
	if endPointService.IpAddressType != nil {
		endPointServiceMap["ip_address_type"] = *endPointService.IpAddressType
	}

	// Handle end_point_service array
	if endPointService.EndPointSet != nil {
		endPointServiceList := make([]interface{}, 0, len(endPointService.EndPointSet))
		for _, endPoint := range endPointService.EndPointSet {
			if endPoint == nil {
				continue
			}
			endPointMap := make(map[string]interface{})
			if endPoint.EndPointId != nil {
				endPointMap["end_point_id"] = *endPoint.EndPointId
			}
			if endPoint.VpcId != nil {
				endPointMap["vpc_id"] = *endPoint.VpcId
			}
			if endPoint.SubnetId != nil {
				endPointMap["subnet_id"] = *endPoint.SubnetId
			}
			if endPoint.EndPointOwner != nil {
				endPointMap["end_point_owner"] = *endPoint.EndPointOwner
			}
			if endPoint.EndPointName != nil {
				endPointMap["end_point_name"] = *endPoint.EndPointName
			}
			if endPoint.ServiceVpcId != nil {
				endPointMap["service_vpc_id"] = *endPoint.ServiceVpcId
			}
			if endPoint.ServiceVip != nil {
				endPointMap["service_vip"] = *endPoint.ServiceVip
			}
			if endPoint.EndPointServiceId != nil {
				endPointMap["end_point_service_id"] = *endPoint.EndPointServiceId
			}
			if endPoint.EndPointVip != nil {
				endPointMap["end_point_vip"] = *endPoint.EndPointVip
			}
			if endPoint.State != nil {
				endPointMap["state"] = *endPoint.State
			}
			if endPoint.CreateTime != nil {
				endPointMap["create_time"] = *endPoint.CreateTime
			}
			if endPoint.ServiceName != nil {
				endPointMap["service_name"] = *endPoint.ServiceName
			}

			// Handle group_set array
			if endPoint.GroupSet != nil {
				groupSetList := make([]interface{}, 0, len(endPoint.GroupSet))
				for _, group := range endPoint.GroupSet {
					if group != nil {
						groupSetList = append(groupSetList, *group)
					}
				}
				endPointMap["group_set"] = groupSetList
			}

			endPointServiceList = append(endPointServiceList, endPointMap)
		}
		endPointServiceMap["end_point_service"] = endPointServiceList
	}

	_ = d.Set("end_point_service", []interface{}{endPointServiceMap})

	return nil
}

func resourceTencentCloudVpcEndPointServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_end_point_service.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := vpc.NewModifyVpcEndPointServiceAttributeRequest()

	endPointServiceId := d.Id()

	request.EndPointServiceId = &endPointServiceId

	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
	}

	if d.HasChange("vpc_id") {
		return fmt.Errorf("tencentcloudenterprise_vpc_end_point_service update on vpc_id is not support yet")
	}

	if d.HasChange("end_point_service_name") {
		if v, ok := d.GetOk("end_point_service_name"); ok {
			request.EndPointServiceName = helper.String(v.(string))
		}
	}

	if d.HasChange("auto_accept_flag") {
		if v, _ := d.GetOk("auto_accept_flag"); v != nil {
			request.AutoAcceptFlag = helper.Bool(v.(bool))
		}
	}

	if d.HasChange("service_instance_id") {
		if v, ok := d.GetOk("service_instance_id"); ok {
			request.ServiceInstanceId = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().ModifyVpcEndPointServiceAttribute(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create vpc endPointService failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudVpcEndPointServiceRead(d, meta)
}

func resourceTencentCloudVpcEndPointServiceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_end_point_service.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	endPointServiceId := d.Id()

	if err := service.DeleteVpcEndPointServiceById(ctx, endPointServiceId); err != nil {
		return nil
	}

	return nil
}
