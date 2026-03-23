/*
Provides a resource to create a cloud firewall (cfw) VPC firewall switch.

Example Usage

```hcl

resource "tencentcloudenterprise_cfw_vpc_firewall_switch" "example" {
  vpc_ins_id = "cfwg-xxxxxxxx"
  switch_id  = "switch-xxxxxxxx"
  enable     = 1
}

```

Import

Cloud firewall VPC firewall switch can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_vpc_firewall_switch.example vpc_ins_id#switch_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudCfwVpcFirewallSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCfwVpcFirewallSwitchCreate,
		Read:   resourceTencentCloudCfwVpcFirewallSwitchRead,
		Update: resourceTencentCloudCfwVpcFirewallSwitchUpdate,
		Delete: resourceTencentCloudCfwVpcFirewallSwitchDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_ins_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Firewall instance id.",
			},
			"switch_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Firewall switch ID.",
			},
			"enable": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Turn the switch on or off. 0: turn off the switch; 1: Turn on the switch.",
			},
		},
	}
}

func resourceTencentCloudCfwVpcFirewallSwitchCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_vpc_firewall_switch.create")()
	defer inconsistentCheck(d, meta)()

	vpcInsId := d.Get("vpc_ins_id").(string)
	switchId := d.Get("switch_id").(string)

	d.SetId(strings.Join([]string{vpcInsId, switchId}, FILED_SP))

	return resourceTencentCloudCfwVpcFirewallSwitchUpdate(d, meta)
}

func resourceTencentCloudCfwVpcFirewallSwitchRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_vpc_firewall_switch.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	vpcInsId := idSplit[0]
	switchId := idSplit[1]

	vpcFirewallSwitch, err := service.DescribeCfwVpcFirewallSwitchById(ctx, vpcInsId, switchId)
	if err != nil {
		return err
	}

	if vpcFirewallSwitch == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CfwVpcFirewallSwitch` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if vpcFirewallSwitch.Enable != nil {
		_ = d.Set("enable", vpcFirewallSwitch.Enable)
	}

	return nil
}

func resourceTencentCloudCfwVpcFirewallSwitchUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_vpc_firewall_switch.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
		request    = cfw.NewModifyFwGroupSwitchRequest()
		switchMode int64
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	vpcInsId := idSplit[0]
	switchId := idSplit[1]

	// get switchMode
	vpcFirewallSwitch, err := service.DescribeCfwVpcFirewallSwitchById(ctx, vpcInsId, switchId)
	if err != nil {
		return err
	}

	if vpcFirewallSwitch == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CfwVpcFirewallSwitch` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if vpcFirewallSwitch.SwitchMode != nil {
		switchMode = *vpcFirewallSwitch.SwitchMode
	}

	if v, ok := d.GetOkExists("enable"); ok {
		request.Enable = helper.IntInt64(v.(int))
	}

	request.AllSwitch = helper.IntInt64(0)
	request.SwitchList = []*cfw.FwGroupSwitch{
		{
			SwitchMode: helper.Int64(switchMode),
			SwitchId:   helper.String(switchId),
		},
	}

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().ModifyFwGroupSwitch(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update cfw vpcFirewallSwitch failed, reason:%+v", logId, err)
		return err
	}

	// wait
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		switchDetail, e := service.DescribeCfwVpcFirewallSwitchById(ctx, vpcInsId, switchId)
		if e != nil {
			return retryError(e)
		}

		if *switchDetail.Status == 0 {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("update cfw vpcFirewallSwitch status is %d", *switchDetail.Status))
	})

	if err != nil {
		return err
	}

	return resourceTencentCloudCfwVpcFirewallSwitchRead(d, meta)
}

func resourceTencentCloudCfwVpcFirewallSwitchDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_vpc_firewall_switch.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfw_vpc_firewall_switch", CNDescription{
		TerraformTypeCN: "VPC防火墙开关",
		DescriptionCN:   "提供云防火墙VPC防火墙开关资源，用于管理VPC防火墙的开关状态。",
		AttributesCN: map[string]string{
			"vpc_ins_id": "VPC防火墙实例ID",
			"switch_id":  "开关ID",
			"enable":     "开关状态，0关闭，1开启",
		},
	})
}
