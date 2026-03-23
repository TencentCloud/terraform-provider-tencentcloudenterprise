/*
Provides a resource to creating direct connect gateway instance.

Example Usage

```hcl
resource "tencentcloudenterprise_vpc" "main" {
  name       = "ci-vpc-instance-test"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_dc_gateway" "vpc_main" {
  name                = "ci-cdg-vpc-test"
  network_instance_id = tencentcloudenterprise_vpc.main.id
  network_type        = "VPC"
  gateway_type        = "NAT"
}
```

Import

Direct connect gateway instance can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_dc_gateway.instance dcg-id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_dc_gateway", CNDescription{
		TerraformTypeCN: "专线网关",
		DescriptionCN:   "提供专线网关资源，用于创建专线网关实例。",
		AttributesCN: map[string]string{
			"name":                "专线网关名称",
			"network_type":        "网络类型",
			"network_instance_id": "网络实例ID，当network_type为VPC时填写VPC ID",
			"gateway_type":        "网关类型：NORMAL（普通网关）、NAT（NAT网关）",
			"bandwidth":           "带宽",
			"cnn_route_type":              "CCN路由类型",
			"enable_bgp":                  "是否启用BGP",
			"create_time":                 "创建时间",
			"vpc_id":                      "VPC ID",
			"ccn_id":                      "CCN ID",
			"direct_connect_gateway_ip":   "专线网关IP地址",
			"enable_bgp_community":        "是否启用BGP Community属性",
		},
	})
}

func resourceTencentCloudDcGatewayInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to creating direct connect gateway instance.",
		Create:      resourceTencentCloudDcGatewayCreate,
		Read:        resourceTencentCloudDcGatewayRead,
		Update:      resourceTencentCloudDcGatewayUpdate,
		Delete:      resourceTencentCloudDcGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "Name of the direct connect gateway.",
			},
			"network_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				// ValidateFunc: validateAllowedStringValue(DCG_NETWORK_TYPES),
				Description:  "Type of associated network. Valid value: `VPC`.",
			},
			"network_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "If the `network_type` value is `VPC`, the available value is VPC ID.", 
			},
			"gateway_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:	  true,
				ValidateFunc: validateAllowedStringValue(DCG_GATEWAY_TYPES),
				Description:  "Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`.",
			},
			"bandwidth": {
				Type:        	schema.TypeInt,
				Optional:    	true,
				ValidateFunc:	validation.IntAtLeast(1),
				Description: 	"The bandwidth speed limit of the gateway.",
			},

			//compute
			"cnn_route_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP enabled.",
			},
			"enable_bgp": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the BGP is enabled.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of resource.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPC ID when network_type is VPC.",
			},
			"ccn_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CCN ID when network_type is CCN.",
			},
			"direct_connect_gateway_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Direct connect gateway IP address.",
			},
			"enable_bgp_community": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether BGP community attribute is enabled.",
			},
		},
	}
}

func resourceTencentCloudDcGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_dc_gateway.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		name              = d.Get("name").(string)
		networkType       = d.Get("network_type").(string)
		networkInstanceId = d.Get("network_instance_id").(string)
		gatewayType       = d.Get("gateway_type").(string)
		bandwidth          = int64(d.Get("bandwidth").(int))
	)

	if networkType == DCG_NETWORK_TYPE_VPC &&
		!strings.HasPrefix(networkInstanceId, "vpc") {

		return fmt.Errorf("if `network_type` is '%s', the field `network_instance_id` must be a VPC resource",
			DCG_NETWORK_TYPE_VPC)
	}

	if networkType == DCG_NETWORK_TYPE_CCN &&
		!strings.HasPrefix(networkInstanceId, "ccn") {

		return fmt.Errorf("if `network_type` is '%s', the field `network_instance_id` must be a CCN resource",
			DCG_NETWORK_TYPE_CCN)
	}

	if networkType == DCG_NETWORK_TYPE_CCN && gatewayType != DCG_GATEWAY_TYPE_NORMAL {

		return fmt.Errorf("if `network_type` is '%s', the field `gateway_type` must be '%s'",
			DCG_NETWORK_TYPE_CCN,
			DCG_GATEWAY_TYPE_NORMAL)
	}

	dcgId, err := service.CreateDirectConnectGateway(ctx, name, gatewayType, networkInstanceId, networkType, bandwidth)
	if err != nil {
		return err
	}

	d.SetId(dcgId)

	// add sleep protect, either network_instance_id will be set "".
	time.Sleep(1 * time.Second)

	return resourceTencentCloudDcGatewayRead(d, meta)
}

func resourceTencentCloudDcGatewayRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_dc_gateway.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, e := service.DescribeDirectConnectGatewayById(ctx, d.Id())
		if e != nil {
			return retryError(e)
		}

		if info == nil {
			return resource.NonRetryableError(fmt.Errorf("direct connect gateway %s not found", d.Id()))
		}

		if info.DirectConnectGatewayName != nil {
			_ = d.Set("name", *info.DirectConnectGatewayName)
		}
		if info.NetworkType != nil {
			_ = d.Set("network_type", *info.NetworkType)
		}
		if info.NetworkInstanceId != nil {
			_ = d.Set("network_instance_id", *info.NetworkInstanceId)
		}
		if info.GatewayType != nil {
			_ = d.Set("gateway_type", *info.GatewayType)
		}
		if info.CcnRouteType != nil {
			_ = d.Set("cnn_route_type", *info.CcnRouteType)
		}
		if info.EnableBGP != nil {
			_ = d.Set("enable_bgp", *info.EnableBGP)
		}
		if info.CreateTime != nil {
			_ = d.Set("create_time", *info.CreateTime)
		}
		if info.VpcId != nil {
			_ = d.Set("vpc_id", *info.VpcId)
		}
		if info.CcnId != nil {
			_ = d.Set("ccn_id", *info.CcnId)
		}
		if info.DirectConnectGatewayIp != nil {
			_ = d.Set("direct_connect_gateway_ip", *info.DirectConnectGatewayIp)
		}
		if info.EnableBGPCommunity != nil {
			_ = d.Set("enable_bgp_community", *info.EnableBGPCommunity)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func resourceTencentCloudDcGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_dc_gateway.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	if d.HasChange("gateway_type") {
		return fmt.Errorf("argument `gateway_type` cannot be changed")
	}

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	if d.HasChange("name") || d.HasChange("bandwidth") {
		var name = d.Get("name").(string)
		var bandwidth = int64(d.Get("bandwidth").(int))
		return service.ModifyDirectConnectGatewayAttribute(ctx, d.Id(), name, bandwidth)
	}

	return resourceTencentCloudDcGatewayRead(d, meta)
}

func resourceTencentCloudDcGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_dc_gateway.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, e := service.DescribeDirectConnectGatewayById(ctx, d.Id())
		if e != nil {
			return retryError(e)
		}

		if info == nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}
	return service.DeleteDirectConnectGateway(ctx, d.Id())
}
