/*
Provides a resource to create a cloud firewall (cfw) edge firewall switch.

Example Usage

```hcl

resource "tencentcloudenterprise_cfw_edge_firewall_switch" "example" {
  public_ip   = "1.1.1.1"
  subnet_id   = "subnet-xxxxxxxx"
  switch_mode = 1
  enable      = 1
}

```

Import

Cloud firewall edge firewall switch can be imported using the public_ip, e.g.

```
$ terraform import tencentcloudenterprise_cfw_edge_firewall_switch.example 1.1.1.1
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudCfwEdgeFirewallSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCfwEdgeFirewallSwitchCreate,
		Read:   resourceTencentCloudCfwEdgeFirewallSwitchRead,
		Update: resourceTencentCloudCfwEdgeFirewallSwitchUpdate,
		Delete: resourceTencentCloudCfwEdgeFirewallSwitchDelete,

		Schema: map[string]*schema.Schema{
			"public_ip": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Public Ip.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The first EIP switch in the vpc is turned on, and you need to specify a subnet to create a private connection. If `switch_mode` is 1 and `enable` is 1, this field is required.",
			},
			"switch_mode": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "0: bypass; 1: serial.",
			},
			"enable": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Switch, 0: off, 1: on.",
			},
		},
	}
}

func resourceTencentCloudCfwEdgeFirewallSwitchCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_edge_firewall_switch.create")()
	defer inconsistentCheck(d, meta)()

	publicIp := d.Get("public_ip").(string)
	d.SetId(publicIp)

	return resourceTencentCloudCfwEdgeFirewallSwitchUpdate(d, meta)
}

func resourceTencentCloudCfwEdgeFirewallSwitchRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_edge_firewall_switch.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		ctx      = context.WithValue(context.TODO(), logIdKey, logId)
		service  = CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
		publicIp = d.Id()
	)

	edgeFirewallSwitch, err := service.DescribeCfwEdgeFirewallSwitchById(ctx, publicIp)
	if err != nil {
		return err
	}

	if edgeFirewallSwitch == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CfwEdgeFirewallSwitch` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if edgeFirewallSwitch.PublicIp != nil {
		_ = d.Set("public_ip", edgeFirewallSwitch.PublicIp)
	}

	if edgeFirewallSwitch.SwitchMode != nil {
		_ = d.Set("switch_mode", edgeFirewallSwitch.SwitchMode)
	}

	if edgeFirewallSwitch.Status != nil {
		_ = d.Set("enable", edgeFirewallSwitch.Status)
	}

	return nil
}

func resourceTencentCloudCfwEdgeFirewallSwitchUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_edge_firewall_switch.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId        = getLogId(contextNil)
		ctx          = context.WithValue(context.TODO(), logIdKey, logId)
		service      = CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
		request      = cfw.NewModifyEdgeIpSwitchRequest()
		edgeIpSwitch = cfw.EdgeIpSwitch{}
		publicIp     = d.Id()
	)

	edgeIpSwitch.PublicIp = &publicIp

	if v, ok := d.GetOk("subnet_id"); ok {
		edgeIpSwitch.SubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("switch_mode"); ok {
		edgeIpSwitch.SwitchMode = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("enable"); ok {
		request.Enable = helper.IntInt64(v.(int))
	}

	request.EdgeIpSwitchLst = append(request.EdgeIpSwitchLst, &edgeIpSwitch)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().ModifyEdgeIpSwitch(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update cfw edgeFirewallSwitch failed, reason:%+v", logId, err)
		return err
	}

	// wait
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		switchDetail, e := service.DescribeCfwEdgeFirewallSwitchById(ctx, publicIp)
		if e != nil {
			return retryError(e)
		}

		if *switchDetail.Status == 0 || *switchDetail.Status == 1 {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("update cfw edgeFirewallSwitch status is %d", *switchDetail.Status))
	})

	if err != nil {
		return err
	}

	return resourceTencentCloudCfwEdgeFirewallSwitchRead(d, meta)
}

func resourceTencentCloudCfwEdgeFirewallSwitchDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_edge_firewall_switch.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfw_edge_firewall_switch", CNDescription{
		TerraformTypeCN: "边界防火墙开关",
		DescriptionCN:   "提供云防火墙边界防火墙开关资源，用于管理边界防火墙的开关状态。",
		AttributesCN: map[string]string{
			"public_ip":   "公网IP地址",
			"enable":      "开关状态，0关闭，1开启",
			"subnet_id":   "子网ID",
			"switch_mode": "交换模式，0：旁路，1：串联",
		},
	})
}
