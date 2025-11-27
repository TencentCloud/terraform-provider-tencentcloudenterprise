/*
Provides a resource to create a vpc vpn_customer_gateway_configuration_download

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpn_customer_gateway_configuration_download" "vpn_customer_gateway_configuration_download" {
	  vpn_gateway_id    = "vpngw-gt8bianl"
	  vpn_connection_id = "vpnx-kme2tx8m"
	  customer_gateway_vendor {
	    platform         = "comware"
	    software_version = "V1.0"
	    vendor_name      = "h3c"
	  }
	  interface_name    = "test"
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
	registerResourceDescriptionProvider("tencentcloudenterprise_vpn_customer_gateway_configuration_download", CNDescription{
		TerraformTypeCN: "创建VPN客户网关配置下载",
		AttributesCN: map[string]string{
			"vpn_gateway_id":                 "VPN网关实例ID",
			"vpn_connection_id":              "VPN通道实例ID",
			"customer_gateway_vendor":        "客户网关厂商信息",
			"interface_name":                 "VPN通道接入设备物理接口名",
			"customer_gateway_configuration": "xml配置",
		},
	})
}
func resourceTencentCloudVpnCustomerGatewayConfigurationDownload() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a vpc vpn_customer_gateway_configuration_download",
		Create:      resourceTencentCloudVpnCustomerGatewayConfigurationDownloadCreate,
		Read:        resourceTencentCloudVpnCustomerGatewayConfigurationDownloadRead,
		Delete:      resourceTencentCloudVpnCustomerGatewayConfigurationDownloadDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpn_gateway_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "VPN Gateway Instance ID.",
			},

			"vpn_connection_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "VPN Connection Instance id.",
			},

			"customer_gateway_vendor": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Customer Gateway Vendor Info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"platform": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Platform.",
						},
						"software_version": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "SoftwareVersion.",
						},
						"vendor_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "VendorName.",
						},
					},
				},
			},

			"interface_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "VPN connection access device physical interface name.",
			},

			"customer_gateway_configuration": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Xml configuration.",
			},
		},
	}
}

func resourceTencentCloudVpnCustomerGatewayConfigurationDownloadCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpn_customer_gateway_configuration_download.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request         = vpc.NewDownloadCustomerGatewayConfigurationRequest()
		response        = vpc.NewDownloadCustomerGatewayConfigurationResponse()
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

	if dMap, ok := helper.InterfacesHeadMap(d, "customer_gateway_vendor"); ok {
		customerGatewayVendor := vpc.CustomerGatewayVendor{}
		if v, ok := dMap["platform"]; ok {
			customerGatewayVendor.Platform = helper.String(v.(string))
		}
		if v, ok := dMap["software_version"]; ok {
			customerGatewayVendor.SoftwareVersion = helper.String(v.(string))
		}
		if v, ok := dMap["vendor_name"]; ok {
			customerGatewayVendor.VendorName = helper.String(v.(string))
		}
		request.CustomerGatewayVendor = &customerGatewayVendor
	}

	if v, ok := d.GetOk("interface_name"); ok {
		request.InterfaceName = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().DownloadCustomerGatewayConfiguration(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate vpc vpnCustomerGatewayConfigurationDownload failed, reason:%+v", logId, err)
		return nil
	}

	d.SetId(vpnGatewayId + FILED_SP + vpnConnectionId)

	_ = d.Set("customer_gateway_configuration", response.Response.CustomerGatewayConfiguration)

	return resourceTencentCloudVpnCustomerGatewayConfigurationDownloadRead(d, meta)
}

func resourceTencentCloudVpnCustomerGatewayConfigurationDownloadRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpn_customer_gateway_configuration_download.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudVpnCustomerGatewayConfigurationDownloadDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpn_customer_gateway_configuration_download.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
