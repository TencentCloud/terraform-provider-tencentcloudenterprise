/*
Use this data source to query detailed information of cloud firewall (cfw) VPC firewall switches.

# Example Usage

```hcl

	data "tencentcloudenterprise_cfw_vpc_fw_switches" "example" {
	  vpc_ins_id = "cfwins-xxxxxxxx"
	}

```
*/
package tencentcloud

import (
	"context"

	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTencentCloudCfwVpcFwSwitches() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCfwVpcFwSwitchesRead,
		Schema: map[string]*schema.Schema{
			"vpc_ins_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Firewall instance id.",
			},
			"switch_list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Switch list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switch_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Firewall switch ID.",
						},
						"switch_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Firewall switch name.",
						},
						"switch_mode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "switch mode.",
						},
						"enable": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Switch status 0: off, 1: on.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Switch status 0: normal, 1: switching.",
						},
					},
				},
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCfwVpcFwSwitchesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cfw_vpc_fw_switches.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = CfwService{client: meta.(*TencentCloudClient).apiV3Conn}
		switchList []*cfw.FwGroupSwitchShow
		vpcInsId   string
	)

	if v, ok := d.GetOk("vpc_ins_id"); ok {
		vpcInsId = v.(string)
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCfwVpcFwSwitchesByFilter(ctx, vpcInsId)
		if e != nil {
			return retryError(e)
		}

		switchList = result
		return nil
	})

	if err != nil {
		return err
	}

	tmpList := make([]map[string]interface{}, 0, len(switchList))

	if switchList != nil {
		for _, fwGroupSwitcheshow := range switchList {
			fwGroupSwitcheshowMap := map[string]interface{}{}

			if fwGroupSwitcheshow.SwitchId != nil {
				fwGroupSwitcheshowMap["switch_id"] = fwGroupSwitcheshow.SwitchId
			}

			if fwGroupSwitcheshow.SwitchName != nil {
				fwGroupSwitcheshowMap["switch_name"] = fwGroupSwitcheshow.SwitchName
			}

			if fwGroupSwitcheshow.SwitchMode != nil {
				fwGroupSwitcheshowMap["switch_mode"] = fwGroupSwitcheshow.SwitchMode
			}

			if fwGroupSwitcheshow.Enable != nil {
				fwGroupSwitcheshowMap["enable"] = fwGroupSwitcheshow.Enable
			}

			if fwGroupSwitcheshow.Status != nil {
				fwGroupSwitcheshowMap["status"] = fwGroupSwitcheshow.Status
			}

			tmpList = append(tmpList, fwGroupSwitcheshowMap)
		}

		_ = d.Set("switch_list", tmpList)
	}

	d.SetId(vpcInsId)
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}

	return nil
}

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cfw_vpc_fw_switches", CNDescription{
		TerraformTypeCN: "VPC防火墙开关列表",
		DescriptionCN:   "提供云防火墙VPC防火墙开关列表数据源，用于查询VPC防火墙开关信息。",
		AttributesCN: map[string]string{
			"result_output_file": "输出结果文件路径",
			"vpc_ins_id":         "VPC防火墙实例ID",
			"switch_list":        "VPC防火墙开关数据列表",
			"switch_id":          "开关ID",
			"switch_name":        "开关名称",
			"switch_mode":        "开关模式",
			"enable":             "开关状态，0关闭，1开启",
			"status":             "开关状态，0正常，1正在切换",
		},
	})
}
