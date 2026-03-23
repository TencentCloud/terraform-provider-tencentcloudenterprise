/*
Provides a resource to create a cloud firewall (cfw) NAT firewall switch.

Example Usage

```hcl

resource "tencentcloudenterprise_cfw_nat_firewall_switch" "example" {
  nat_ins_id = "cfwnat-xxxxxxxx"
  subnet_id  = "subnet-xxxxxxxx"
  enable     = 1
}

```

Import

Cloud firewall NAT firewall switch can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_nat_firewall_switch.example nat_ins_id#subnet_id
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

func resourceTencentCloudCfwNatFirewallSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCfwNatFirewallSwitchCreate,
		Read:   resourceTencentCloudCfwNatFirewallSwitchRead,
		Update: resourceTencentCloudCfwNatFirewallSwitchUpdate,
		Delete: resourceTencentCloudCfwNatFirewallSwitchDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"nat_ins_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Firewall instance id.",
			},
			"subnet_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "subnet id.",
			},
			"enable": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Switch, 0: off, 1: on.",
			},
		},
	}
}

func resourceTencentCloudCfwNatFirewallSwitchCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_nat_firewall_switch.create")()
	defer inconsistentCheck(d, meta)()

	var (
		natInsId string
		subnetId string
	)

	if v, ok := d.GetOk("nat_ins_id"); ok {
		natInsId = v.(string)
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		subnetId = v.(string)
	}

	d.SetId(strings.Join([]string{natInsId, subnetId}, FILED_SP))
	return resourceTencentCloudCfwNatFirewallSwitchUpdate(d, meta)
}

func resourceTencentCloudCfwNatFirewallSwitchRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_nat_firewall_switch.read")()
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
	natInsId := idSplit[0]
	subnetId := idSplit[1]

	natFirewallSwitch, err := service.DescribeCfwNatFirewallFwSwitchById(ctx, natInsId, subnetId)
	if err != nil {
		return err
	}

	if natFirewallSwitch == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CfwNatFirewallSwitch` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if natFirewallSwitch.NatInsId != nil {
		_ = d.Set("nat_ins_id", natFirewallSwitch.NatInsId)
	}

	if natFirewallSwitch.SubnetId != nil {
		_ = d.Set("subnet_id", natFirewallSwitch.SubnetId)
	}

	if natFirewallSwitch.Enable != nil {
		_ = d.Set("enable", natFirewallSwitch.Enable)
	}

	return nil
}

func resourceTencentCloudCfwNatFirewallSwitchUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_nat_firewall_switch.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
		request = cfw.NewModifyNatFwSwitchRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	natInsId := idSplit[0]
	subnetId := idSplit[1]

	if v, ok := d.GetOkExists("enable"); ok {
		request.Enable = helper.IntInt64(v.(int))
	}

	request.SubnetIdList = helper.Strings([]string{subnetId})
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().ModifyNatFwSwitch(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update cfw natFirewallSwitch failed, reason:%+v", logId, err)
		return err
	}

	// wait
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		switchDetail, e := service.DescribeCfwNatFirewallFwSwitchById(ctx, natInsId, subnetId)
		if e != nil {
			return retryError(e)
		}

		if *switchDetail.Status == 0 {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("update cfw natFirewallSwitch status is %d", *switchDetail.Status))
	})

	if err != nil {
		return err
	}

	return resourceTencentCloudCfwNatFirewallSwitchRead(d, meta)
}

func resourceTencentCloudCfwNatFirewallSwitchDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_nat_firewall_switch.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfw_nat_firewall_switch", CNDescription{
		TerraformTypeCN: "NAT防火墙开关",
		DescriptionCN:   "提供云防火墙NAT防火墙开关资源，用于管理NAT防火墙的开关状态。",
		AttributesCN: map[string]string{
			"nat_ins_id": "NAT防火墙实例ID",
			"subnet_id":  "子网ID",
			"enable":     "开关状态，0关闭，1开启",
		},
	})
}
