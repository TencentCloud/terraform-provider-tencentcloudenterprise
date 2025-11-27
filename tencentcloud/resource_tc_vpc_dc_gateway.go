/*
Provides a resource to creating direct connect gateway instance.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc" "main" {
	  name       = "ci-vpc-instance-test"
	  cidr_block = "10.0.0.0/16"
	}

	resource "tencentcloudenterprise_vpc_dc_gateway" "vpc_main" {
	  name                = "ci-cdg-vpc-test"
	  network_instance_id = cloud_vpc.main.id
	  network_type        = "VPC"
	  gateway_type        = "NAT"
	}

```

# Import

Direct connect gateway instance can be imported, e.g.

```
$ terraform import cloud_vpc_dc_gateway.instance dcg-id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
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
			"network_instance_id": "网络实例ID",
			"gateway_type":        "网关类型",
			"band_with":           "带宽",
			"cnn_route_type":      "CCN路由类型",
			"enable_bgp":          "是否启用BGP",
			"create_time":         "创建时间",
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
				Description:  "Name of the DCG.",
			},
			"network_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// ValidateFunc: validateAllowedStringValue(DCG_NETWORK_TYPES),
				Description: "Type of associated network. Valid value: `VPC`.",
			},
			"network_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "If the `network_type` value is `VPC`, the available value is VPC ID.",
			},
			"gateway_type": {
				Type:     schema.TypeString,
				Required: true,
				// Default:      DCG_GATEWAY_TYPE_NORMAL,
				ValidateFunc: validateAllowedStringValue(DCG_GATEWAY_TYPES),
				Description:  "Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`.",
			},
			"band_with": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The bandwith speed limit of the gateway.",
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
		},
	}
}

func resourceTencentCloudDcGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_dc_gateway.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		name              = d.Get("name").(string)
		networkType       = d.Get("network_type").(string)
		networkInstanceId = d.Get("network_instance_id").(string)
		gatewayType       = d.Get("gateway_type").(string)
		bandwith          = int64(d.Get("band_with").(int))
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

	dcgId, err := service.CreateDirectConnectGateway(ctx, name, gatewayType, networkInstanceId, networkType, bandwith)
	if err != nil {
		return err
	}

	d.SetId(dcgId)

	// add sleep protect, either network_instance_id will be set "".
	time.Sleep(1 * time.Second)

	return resourceTencentCloudDcGatewayRead(d, meta)
}

func resourceTencentCloudDcGatewayRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_dc_gateway.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, e := service.DescribeDirectConnectGatewayById(ctx, d.Id())
		if e != nil {
			return retryError(e)
		}

		_ = d.Set("name", info.DirectConnectGatewayName)
		_ = d.Set("network_type", info.NetworkType)
		_ = d.Set("network_instance_id", info.NetworkInstanceId)
		_ = d.Set("gateway_type", info.GatewayType)
		_ = d.Set("cnn_route_type", info.CcnRouteType)
		_ = d.Set("enable_bgp", info.EnableBGP)
		_ = d.Set("create_time", info.CreateTime)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func resourceTencentCloudDcGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_dc_gateway.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	if d.HasChange("name") || d.HasChanges("band_with") {
		var name = d.Get("name").(string)
		var bandwith = int64(d.Get("band_with").(int))
		return service.ModifyDirectConnectGatewayAttribute(ctx, d.Id(), name, bandwith)
	}

	return resourceTencentCloudDcGatewayRead(d, meta)
}

func resourceTencentCloudDcGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_dc_gateway.delete")()

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
