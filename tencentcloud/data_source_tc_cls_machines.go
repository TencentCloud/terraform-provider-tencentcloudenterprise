/*
Use this data source to query detailed information of cls machines

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_machine_group" "group" {
	  group_name        = "tf-describe-mg-test"
	  service_logging   = true
	  auto_update       = true
	  update_end_time   = "19:05:00"
	  update_start_time = "17:05:00"

	  machine_group_type {
	    type   = "ip"
	    values = [
	      "203.0.113.101",
	      "203.0.113.102",
	    ]
	  }
	}

	data "tencentcloudenterprise_cls_machines" "machines" {
	  group_id = cloud_cls_machine_group.group.id
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cls_machines", CNDescription{
		TerraformTypeCN: "CLS机器组机器列表",
		DescriptionCN:   "提供CLS机器组机器列表数据源，用于查询指定机器组中的机器状态和详细信息。",
		AttributesCN: map[string]string{
			"group_id":           "机器组ID",
			"result_output_file": "用于保存结果，可视化界面不可用",
			"machines":           "机器列表",
		},
	})
}

func dataSourceTencentCloudClsMachines() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudClsMachinesRead,
		Description: "Queries machine status and details within a CLS machine group",
		Schema: map[string]*schema.Schema{
			"group_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Group id.",
			},

			"machines": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Info of Machines.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Ip of machine.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Status of machine.",
						},
						"offline_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Offline time of machine.",
						},
						"auto_update": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "If open auto update flag.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current machine version.",
						},
						"update_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Machine update status.",
						},
						"err_code": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Code of update operation.",
						},
						"err_msg": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Msg of update operation.",
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

func dataSourceTencentCloudClsMachinesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cls_machines.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("group_id"); ok {
		paramMap["GroupId"] = helper.String(v.(string))
	}

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var machines []*cls.MachineInfo

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeClsMachinesByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		machines = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(machines))
	tmpList := make([]map[string]interface{}, 0, len(machines))

	if machines != nil {
		for _, machineInfo := range machines {
			machineInfoMap := map[string]interface{}{}

			if machineInfo.Ip != nil {
				machineInfoMap["ip"] = machineInfo.Ip
			}

			if machineInfo.Status != nil {
				machineInfoMap["status"] = machineInfo.Status
			}

			if machineInfo.OfflineTime != nil {
				machineInfoMap["offline_time"] = machineInfo.OfflineTime
			}

			if machineInfo.AutoUpdate != nil {
				machineInfoMap["auto_update"] = machineInfo.AutoUpdate
			}

			if machineInfo.Version != nil {
				machineInfoMap["version"] = machineInfo.Version
			}

			if machineInfo.UpdateStatus != nil {
				machineInfoMap["update_status"] = machineInfo.UpdateStatus
			}

			if machineInfo.ErrCode != nil {
				machineInfoMap["err_code"] = machineInfo.ErrCode
			}

			if machineInfo.ErrMsg != nil {
				machineInfoMap["err_msg"] = machineInfo.ErrMsg
			}

			ids = append(ids, *machineInfo.Ip)
			tmpList = append(tmpList, machineInfoMap)
		}

		_ = d.Set("machines", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}
