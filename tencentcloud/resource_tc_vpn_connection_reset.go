/*
Provides a resource to create a vpc vpn_connection_reset

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpn_connection_reset" "vpn_connection_reset" {
	  vpn_gateway_id    = "vpngw-gt8bianl"
	  vpn_connection_id = "vpnx-kme2tx8m"
	}

```
*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpn_connection_reset", CNDescription{
		TerraformTypeCN: "创建VPN重置通道",
		AttributesCN: map[string]string{
			"vpn_gateway_id":    "VPN网关实例ID",
			"vpn_connection_id": "VPN通道实例ID",
		},
	})
}
func resourceTencentCloudVpnConnectionReset() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a vpc vpn_connection_reset",
		Create:      resourceTencentCloudVpnConnectionResetCreate,
		Read:        resourceTencentCloudVpnConnectionResetRead,
		Delete:      resourceTencentCloudVpnConnectionResetDelete,
		Schema: map[string]*schema.Schema{
			"vpn_gateway_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "VPN GATEWAY INSTANCE ID.",
			},

			"vpn_connection_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "VPN CONNECTION INSTANCE ID.",
			},
		},
	}
}

func resourceTencentCloudVpnConnectionResetCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpn_connection_reset.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request         = vpc.NewResetVpnConnectionRequest()
		vpnGatewayId    string
		vpnConnectionId string
	)
	if v, ok := d.GetOk("vpn_gateway_id"); ok {
		vpnGatewayId = v.(string)
		request.VpnGatewayId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("vpn_connection_id"); ok {
		vpnConnectionId = v.(string)
		request.VpnConnectionId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().ResetVpnConnection(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate vpc vpnConnectionReset failed, reason:%+v", logId, err)
		return nil
	}

	d.SetId(vpnGatewayId + FILED_SP + vpnConnectionId)

	return resourceTencentCloudVpnConnectionResetRead(d, meta)
}

func resourceTencentCloudVpnConnectionResetRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpn_connection_reset.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudVpnConnectionResetDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpn_connection_reset.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
