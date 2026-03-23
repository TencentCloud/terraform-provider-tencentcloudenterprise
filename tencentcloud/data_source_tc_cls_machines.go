/*
Use this data source to query detailed information of cls machines

Example Usage

```hcl
data "tencentcloudenterprise_cls_machines" "machines" {
  group_id = "80278829-74ed-4bde-a4bc-e9837fd8e0ef"
}

# Output all machines information
output "machines_list" {
  value = data.tencentcloudenterprise_cls_machines.machines.machines
}

# Output specific machine details
output "machine_details" {
  value = [for machine in data.tencentcloudenterprise_cls_machines.machines.machines : {
    ip          = machine.ip
    status      = machine.status
    version     = machine.version
    instance_id = machine.instance_id
  }]
}
```
*/
package tencentcloud

import (
	"context"

	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cls_machines", CNDescription{
		TerraformTypeCN: "CLS机器组机器列表",
		DescriptionCN:   "提供CLS机器组机器列表数据源，用于查询指定机器组中的机器状态和详细信息。",
		AttributesCN: map[string]string{
			"group_id":             "查询的机器组ID",
			"result_output_file":   "用于保存结果，可视化界面不可用",
			"machines":             "机器状态信息组",
			"ip":                   "机器的IP",
			"status":               "机器状态，0:异常，1:正常",
			"offline_time":         "机器离线时间，空为正常，异常返回具体时间",
			"auto_update":          "机器是否开启自动升级。0:关闭，1:开启",
			"version":              "机器当前版本号",
			"update_status":        "机器升级功能状态。0：升级成功；1：升级中；-1：升级失败",
			"err_code":             "机器升级结果标识。0：成功；1200：升级成功；其他值表示异常",
			"err_msg":              "机器升级结果信息。\"ok\"：成功；\"update success\"：升级成功；其他值为失败原因",
			"instance_id":          "机器实例ID",
			"group_auto_update":    "机器组是否开启自动升级功能。0：未开启自动升级；1：开启了自动升级",
			"update_start_time":    "机器组自动升级功能预设开始时间",
			"update_end_time":      "机器组自动升级功能预设结束时间",
			"latest_agent_version": "当前用户可用最新的Loglistener版本",
			"service_logging":      "是否开启服务日志。true表示开启服务日志，false表示不开启服务日志",
			"flag":                 "TKE标志位，默认值为空字符串。空字符串表示日志不是来自于TKE，label_k8s表示日志来自于TKE",
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
				Description: "Queried machine group ID.",
			},

			"machines": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Machine status information group.",
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
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Machine instance ID.",
						},
					},
				},
			},

			"group_auto_update": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Whether the machine group has enabled automatic upgrade function. 0: Automatic upgrade not enabled; 1: Automatic upgrade enabled.",
			},

			"update_start_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Scheduled start time for the automatic upgrade feature of the machine group.",
			},

			"update_end_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Scheduled end time for the auto-upgrade feature of the machine group.",
			},

			"latest_agent_version": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The latest LogListener version available to the current user.",
			},

			"service_logging": {
				Computed:    true,
				Type:        schema.TypeBool,
				Description: "Whether service logs are enabled. true: Service logs are enabled. false: Service logs are not enabled.",
			},

			"flag": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "TKE flag. The default value is an empty string. An empty string indicates logs are not from TKE. label_k8s indicates logs are from TKE.",
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
	defer logElapsed("data_source.tencentcloudenterprise_cls_machines.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("group_id"); ok {
		paramMap["GroupId"] = helper.String(v.(string))
	}

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var machines []*cls.MachineInfo
	var response *cls.DescribeMachinesResponse

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, resp, e := service.DescribeClsMachinesByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		machines = result
		response = resp
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

			if machineInfo.InstanceID != nil {
				machineInfoMap["instance_id"] = machineInfo.InstanceID
			}

			ids = append(ids, *machineInfo.Ip)
			tmpList = append(tmpList, machineInfoMap)
		}

		_ = d.Set("machines", tmpList)
	}

	// Set additional response fields
	if response != nil && response.Response != nil {
		if response.Response.AutoUpdate != nil {
			_ = d.Set("group_auto_update", response.Response.AutoUpdate)
		}
		if response.Response.UpdateStartTime != nil {
			_ = d.Set("update_start_time", response.Response.UpdateStartTime)
		}
		if response.Response.UpdateEndTime != nil {
			_ = d.Set("update_end_time", response.Response.UpdateEndTime)
		}
		if response.Response.LatestAgentVersion != nil {
			_ = d.Set("latest_agent_version", response.Response.LatestAgentVersion)
		}
		if response.Response.ServiceLogging != nil {
			_ = d.Set("service_logging", response.Response.ServiceLogging)
		}
		if response.Response.Flag != nil {
			_ = d.Set("flag", response.Response.Flag)
		}
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
