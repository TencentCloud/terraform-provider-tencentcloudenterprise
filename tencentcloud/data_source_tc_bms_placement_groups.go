/*
Use this data source to query bms placement group.

# Example Usage

```hcl

	data "tencentcloudenterprise_bms_placement_groups" "bms" {
	  result_output_file = "bms.json"
	}

```
*/
package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bms "terraform-provider-tencentcloudenterprise/sdk/bms/v20180813"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_bms_placement_groups", CNDescription{
		TerraformTypeCN: "置放群组信息列表",
		DescriptionCN:   "提供BMS置放群组信息列表数据源，用于查询BMS置放群组信息。",
		AttributesCN: map[string]string{
			"instance_id":                         "实例ID",
			"instance_name":                       "实例名称",
			"instance_type":                       "实例机型",
			"image_id":                            "镜像ID",
			"hostname":                            "主机名",
			"availability_zone":                   "可用区",
			"project_id":                          "项目ID",
			"instance_charge_type":                "计费模式",
			"instance_charge_type_prepaid_period": "预付费时长",
			"instance_charge_type_prepaid_renew_flag": "自动续费标识",
			"spot_instance_type":                      "竞价实例类型",
			"spot_max_price":                          "竞价实例最高价格",
			"cdh_instance_type":                       "CDH实例类型",
			"cdh_host_id":                             "CDH实例ID",
			"internet_charge_type":                    "网络计费类型",
			"bandwidth_package_id":                    "共享带宽包ID",
			"internet_max_bandwidth_out":              "公网出带宽",
			"allocate_public_ip":                      "是否分配公网IP",
			"vpc_id":                                  "VPC ID",
			"subnet_id":                               "子网ID",
			"private_ip":                              "私有IP",
			"security_groups":                         "安全组ID",
			"orderly_security_groups":                 "安全组ID",
			"system_disk_type":                        "系统盘类型",
			"system_disk_size":                        "系统盘大小",
			//"system_disk_id":                          "系统盘快照ID",
			"data_disks":               "数据盘",
			"disable_security_service": "是否关闭安全服务",
			"disable_monitor_service":  "是否关闭云监控服务",
			"key_name":                 "密钥对名称",
			"key_ids":                  "密钥对ID",
			"password":                 "密码",
			"keep_image_login":         "保留镜像原始设置",
			"tags":                     "标签",
			"force_delete":             "是否强制删除",

			"instance_status":    "实例状态",
			"create_time":        "创建时间",
			"group_id":           "置放群组ID",
			"result_output_file": "结果输出文件",
		},
	})
}
func dataTencentCloudBmsPlacementGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query bms placement group.",
		Read:        dataTencentCloudBmsPlacementGroupRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the placement group.",
			},

			// Computed values.
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The file path to output the result.",
			},

			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of placement group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The group ID of the instance.",
						},

						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the instance.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the instance.",
						},
						"current_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The current num of the instance.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of the instance.",
						},

						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of the instance.",
						},
					},
				},
			},
		},
	}
}

func dataTencentCloudBmsPlacementGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_bms_instance.read")()
	defer inconsistentCheck(d, meta)()

	client := meta.(*TencentCloudClient).apiV3Conn

	var groups []*bms.DisasterRecoverGroup
	var request = bms.NewDescribeDisasterRecoverGroupsRequest()
	if v, ok := d.GetOk("group_id"); ok && v.(string) != "" {
		request.GroupIds = []*string{helper.String(v.(string))}
	}
	// 显式指定分页大小和偏移量
	request.Limit = helper.Int64(100)
	request.Offset = helper.Int64(0)

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		// 循环获取所有分页数据
		for {
			response, err := client.UseBmsClient().DescribeDisasterRecoverGroups(request)
			if err != nil {
				return retryError(err, InternalError)
			}

			groups = append(groups, response.Response.GroupSet...)

			if len(response.Response.GroupSet) < int(*request.Limit) {
				break
			}

			request.Offset = helper.Int64(*request.Offset + *request.Limit)
		}
		return nil
	})
	if err != nil {
		return err
	}

	instanceList := make([]map[string]interface{}, 0, len(groups))
	ids := make([]string, 0, len(groups))

	for _, instance := range groups {
		m := map[string]interface{}{
			"group_id":    instance.GroupId,
			"name":        instance.Name,
			"type":        instance.Type,
			"current_num": instance.CurrentNum,
			"create_time": instance.CreateTime,
			//"update_time": instance.UpdateTime,
		}
		instanceList = append(instanceList, m)
		ids = append(ids, *instance.GroupId)
	}

	_ = d.Set("list", instanceList)
	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), instanceList); err != nil {
			return err
		}
	}

	return nil
}
