/*
Provides a resource to create a vpc net_detect

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc_net_detect" "net_detect" {
	  net_detect_name       = "terrform-test"
	  vpc_id                = "vpc-4owdpnwr"
	  subnet_id             = "subnet-c1l35990"
	  next_hop_destination  = "203.0.113.57"
	  next_hop_type         = "NORMAL_CVM"
	  detect_destination_ip = [
	    "203.0.113.1",
	    "203.0.113.2",
	  ]
	}

```

# Import

vpc net_detect can be imported using the id, e.g.

```
terraform import cloud_vpc_net_detect.net_detect net_detect_id
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_net_detect", CNDescription{
		TerraformTypeCN: "网络探测",
		DescriptionCN:   "提供网络探测资源，用于创建网络探测实例。",
		AttributesCN: map[string]string{
			"vpc_id":                 "VPC实例ID",
			"subnet_id":              "子网实例ID",
			"net_detect_name":        "网络探测名称",
			"detect_destination_ip":  "探测目的IPv4地址数组",
			"next_hop_type":          "下一跳类型",
			"next_hop_destination":   "下一跳目的网关",
			"net_detect_description": "网络探测描述",
		},
	})
}

func resourceTencentCloudVpcNetDetect() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a vpc net_detect",
		Create:      resourceTencentCloudVpcNetDetectCreate,
		Read:        resourceTencentCloudVpcNetDetectRead,
		Update:      resourceTencentCloudVpcNetDetectUpdate,
		Delete:      resourceTencentCloudVpcNetDetectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "`VPC` instance `ID`. Such as:`vpc-12345678`.",
			},

			"subnet_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Subnet instance ID. Such as:subnet-12345678.",
			},

			"net_detect_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Network probe name, the maximum length cannot exceed 60 bytes.",
			},

			"detect_destination_ip": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "An array of probe destination IPv4 addresses. Up to two.",
			},

			"next_hop_type": {
				Required: true,
				Type:     schema.TypeString,
				ValidateFunc: validateAllowedStringValue([]string{"VPN", "DIRECTCONNECT", "PEERCONNECTION", "NAT",
					"NORMAL_CVM"}),
				Description: "NextHopType: indicates the next hop type.Currently, the following types are supported: " +
					"VPN: indicates the VPN gateway. " +
					"DIRECTCONNECT: private line gateway; " +
					"PEERCONNECTION: peer connection; " +
					"NAT: indicates the NAT gateway. " +
					"NORMAL_CVM: common cloud server.",
			},

			"next_hop_destination": {
				Required: true,
				Type:     schema.TypeString,
				Description: "The destination gateway of the next hop, the value is related to the next hop type." +
					"If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678;" +
					"If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; " +
					"If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678;" +
					"If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; " +
					"If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 203.0.113.12.",
			},

			"net_detect_description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Network probe description.",
			},
		},
	}
}

func resourceTencentCloudVpcNetDetectCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_net_detect.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request     = vpc.NewCreateNetDetectRequest()
		response    = vpc.NewCreateNetDetectResponse()
		netDetectId string
	)
	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		request.SubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("net_detect_name"); ok {
		request.NetDetectName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("detect_destination_ip"); ok {
		detectDestinationIpSet := v.(*schema.Set).List()
		for i := range detectDestinationIpSet {
			detectDestinationIp := detectDestinationIpSet[i].(string)
			request.DetectDestinationIp = append(request.DetectDestinationIp, &detectDestinationIp)
		}
	}

	if v, ok := d.GetOk("next_hop_type"); ok {
		request.NextHopType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("next_hop_destination"); ok {
		request.NextHopDestination = helper.String(v.(string))
	}

	if v, ok := d.GetOk("net_detect_description"); ok {
		request.NetDetectDescription = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().CreateNetDetect(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create vpc netDetect failed, reason:%+v", logId, err)
		return err
	}

	netDetectId = *response.Response.NetDetect.NetDetectId
	d.SetId(netDetectId)

	return resourceTencentCloudVpcNetDetectRead(d, meta)
}

func resourceTencentCloudVpcNetDetectRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_net_detect.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	netDetectId := d.Id()

	netDetect, err := service.DescribeVpcNetDetectById(ctx, netDetectId)
	if err != nil {
		return err
	}

	if netDetect == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `VpcNetDetect` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if netDetect.VpcId != nil {
		_ = d.Set("vpc_id", netDetect.VpcId)
	}

	if netDetect.SubnetId != nil {
		_ = d.Set("subnet_id", netDetect.SubnetId)
	}

	if netDetect.NetDetectName != nil {
		_ = d.Set("net_detect_name", netDetect.NetDetectName)
	}

	if netDetect.DetectDestinationIp != nil {
		_ = d.Set("detect_destination_ip", netDetect.DetectDestinationIp)
	}

	if netDetect.NextHopType != nil {
		_ = d.Set("next_hop_type", netDetect.NextHopType)
	}

	if netDetect.NextHopDestination != nil {
		_ = d.Set("next_hop_destination", netDetect.NextHopDestination)
	}

	if netDetect.NetDetectDescription != nil {
		_ = d.Set("net_detect_description", netDetect.NetDetectDescription)
	}

	return nil
}

func resourceTencentCloudVpcNetDetectUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_net_detect.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := vpc.NewModifyNetDetectRequest()

	netDetectId := d.Id()

	request.NetDetectId = &netDetectId

	mutableArgs := []string{
		"net_detect_name", "detect_destination_ip", "next_hop_type",
		"next_hop_destination", "net_detect_description",
	}
	needChange := false
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {

		if v, ok := d.GetOk("net_detect_name"); ok {
			request.NetDetectName = helper.String(v.(string))
		}

		if v, ok := d.GetOk("detect_destination_ip"); ok {
			detectDestinationIpSet := v.(*schema.Set).List()
			for i := range detectDestinationIpSet {
				detectDestinationIp := detectDestinationIpSet[i].(string)
				request.DetectDestinationIp = append(request.DetectDestinationIp, &detectDestinationIp)
			}
		}

		if v, ok := d.GetOk("next_hop_type"); ok {
			request.NextHopType = helper.String(v.(string))
		}

		if v, ok := d.GetOk("next_hop_destination"); ok {
			request.NextHopDestination = helper.String(v.(string))
		}

		if v, ok := d.GetOk("net_detect_description"); ok {
			request.NetDetectDescription = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().ModifyNetDetect(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update vpc netDetect failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudVpcNetDetectRead(d, meta)
}

func resourceTencentCloudVpcNetDetectDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_net_detect.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	netDetectId := d.Id()

	if err := service.DeleteVpcNetDetectById(ctx, netDetectId); err != nil {
		return err
	}

	return nil
}
