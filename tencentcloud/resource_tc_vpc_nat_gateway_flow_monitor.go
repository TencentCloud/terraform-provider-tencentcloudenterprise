/*
Provides a resource to manage the flow monitor of a NAT gateway.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc_nat_gateway_flow_monitor" "example" {
	  gateway_id = "nat-xxxxxxxx"
	  enable     = true
	}

```

# Import

NAT gateway flow monitor can be imported using the gateway id, e.g.

```
$ terraform import tencentcloudenterprise_vpc_nat_gateway_flow_monitor.example nat-xxxxxxxx
```
*/
package tencentcloud

import (
	"log"

	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_nat_gateway_flow_monitor", CNDescription{
		TerraformTypeCN: "NAT网关流量监控",
		DescriptionCN:   "提供NAT网关流量监控资源，用于开启或关闭NAT网关的流量监控功能。",
		AttributesCN: map[string]string{
			"gateway_id": "NAT网关实例ID",
			"enable":     "是否开启流量监控",
			"bandwidth":  "流量监控带宽",
		},
	})
}

func resourceTencentCloudVpcNatGatewayFlowMonitor() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to manage the flow monitor of a NAT gateway.",
		Create:      resourceTencentCloudVpcNatGatewayFlowMonitorCreate,
		Read:        resourceTencentCloudVpcNatGatewayFlowMonitorRead,
		Update:      resourceTencentCloudVpcNatGatewayFlowMonitorUpdate,
		Delete:      resourceTencentCloudVpcNatGatewayFlowMonitorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"gateway_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "ID of Gateway.",
			},
			"enable": {
				Required:    true,
				Type:        schema.TypeBool,
				Description: "Whether to enable flow monitor.",
			},
			"bandwidth": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Bandwidth of flow monitor.",
			},
		},
	}
}

func resourceTencentCloudVpcNatGatewayFlowMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_nat_gateway_flow_monitor.create")()
	defer inconsistentCheck(d, meta)()

	gatewayId := d.Get("gateway_id").(string)
	d.SetId(gatewayId)

	return resourceTencentCloudVpcNatGatewayFlowMonitorUpdate(d, meta)
}

func resourceTencentCloudVpcNatGatewayFlowMonitorRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_nat_gateway_flow_monitor.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	gatewayId := d.Id()

	var (
		checkRequest = vpc.NewCheckGatewayFlowMonitorRequest()
		response     = vpc.NewCheckGatewayFlowMonitorResponse()
	)

	checkRequest.GatewayId = &gatewayId
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().CheckGatewayFlowMonitor(checkRequest)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, checkRequest.GetAction(), checkRequest.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s check nat gateway flow monitor failed, reason:%+v", logId, err)
		return err
	}

	_ = d.Set("gateway_id", gatewayId)

	if response.Response != nil {
		if response.Response.Enabled != nil {
			_ = d.Set("enable", *response.Response.Enabled)
		}
		if response.Response.Bandwidth != nil {
			_ = d.Set("bandwidth", int(*response.Response.Bandwidth)) // bandwidth in Mbps, safe to convert uint64 to int
		}
	}

	return nil
}

func resourceTencentCloudVpcNatGatewayFlowMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_nat_gateway_flow_monitor.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	gatewayId := d.Id()
	enable := d.Get("enable").(bool)

	if enable {
		enableRequest := vpc.NewEnableGatewayFlowMonitorRequest()
		enableRequest.GatewayId = &gatewayId
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().EnableGatewayFlowMonitor(enableRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, enableRequest.GetAction(), enableRequest.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s enable nat gateway flow monitor failed, reason:%+v", logId, err)
			return err
		}
	} else {
		disableRequest := vpc.NewDisableGatewayFlowMonitorRequest()
		disableRequest.GatewayId = &gatewayId
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().DisableGatewayFlowMonitor(disableRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, disableRequest.GetAction(), disableRequest.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s disable nat gateway flow monitor failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudVpcNatGatewayFlowMonitorRead(d, meta)
}

func resourceTencentCloudVpcNatGatewayFlowMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_nat_gateway_flow_monitor.delete")()
	defer inconsistentCheck(d, meta)()

	// Flow monitor is an attribute of NAT gateway and cannot be physically deleted.
	// Removing this resource only clears the Terraform state.
	return nil
}
