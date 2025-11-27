/*
Use this data source to query bms instances.

# Example Usage

```hcl

	data "tencentcloudenterprise_bms_instances" "bms" {
	  result_output_file = "bms.json"
	}

```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
	bms "terraform-provider-tencentcloudenterprise/sdk/bms/v20180813"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_bms_instances", CNDescription{
		TerraformTypeCN: "裸金属服务器列表",
		DescriptionCN:   "提供裸金属服务器列表数据源，用于查询BMS实例信息。",
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
			"vpc_id":                                  "私有网络ID",
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

			"instance_status":       "实例状态",
			"create_time":           "创建时间",
			"cpu_auch":              "CPU架构",
			"operating_system_type": "操作系统类型",
			"private_ip_address":    "私有IP地址",
			"zone":                  "可用区",
			"instance_state":        "实例状态",
			"result_output_file":    "结果输出文件",
			"cpu_arch":              "CPU架构",
			"instance_list":         "实例列表",
		},
	})
}
func dataTencentCloudBmsInstances() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to query bms instances.",
		Read:        dataTencentCloudBmsInstanceRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the instance.",
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The available zone for the bms instance.",
			},
			"instance_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Terraform-bms-Instance",
				ValidateFunc: validateStringLengthInRange(2, 128),
				Description:  "The name of the instance. The max length of instance_name is 60, and default value is `Terraform-bms-Instance`.",
			},
			"instance_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The state of the instance.",
			},
			"private_ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The private ip address of the instance.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The VPC ID of the instance.",
			},

			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The subnet ID of the instance.",
			},
			/*
				"group_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The group ID of the instance.",
				},
			*/
			"operating_system_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The operating system type of the instance.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The file path to output the result.",
			},

			// Computed values.
			"instance_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the instance.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The available zone for the bms instance.",
						},
						/*
							"group_id": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "The group ID of the instance.",
							},
						*/
						"instance_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current status of the instance.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of the instance.",
						},
						"cpu_arch": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cpu arch of the instance.",
						},
						"operating_system_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The operating system type of the instance.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The vpc id of the instance.",
						},
						"private_ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The private ip address of the instance.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The subnet id of the instance.",
						},
					},
				},
			},
		},
	}
}

func dataTencentCloudBmsInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_bms_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	params := make(map[string]interface{})
	if v, ok := d.GetOk("zone"); ok {
		params["zone"] = v.(string)
	}
	if v, ok := d.GetOk("instance_id"); ok {
		params["instance-id"] = v.(string)
	}
	//if v, ok := d.GetOk("instance_name"); ok {
	//	params["instance-name"] = v.(string)
	//}
	if v, ok := d.GetOk("instance_state"); ok {
		params["instance-state"] = v.(string)
	}
	if v, ok := d.GetOk("private_ip_address"); ok {
		params["private-ip-address"] = v.(string)
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		params["vpc-id"] = v.(string)
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		params["subnet-id"] = v.(string)
	}

	/*
		if v, ok := d.GetOk("group_id"); ok {
			params["group-id"] = v.(string)
		}
	*/

	if v, ok := d.GetOk("cpu_arch"); ok {
		params["cpuArch"] = v.(string)
	}
	if v, ok := d.GetOk("operating_system_type"); ok {
		params["operating-system-type"] = v.(string)
	}

	client := meta.(*TencentCloudClient).apiV3Conn
	bmsService := BmsService{
		client: client,
	}
	var instances []*bms.Instance
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instances, errRet = bmsService.DescribeInstancesByFilter(ctx, params)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}

	instanceList := make([]map[string]interface{}, 0, len(instances))
	ids := make([]string, 0, len(instances))

	for _, instance := range instances {

		strIPs := make([]string, len(instance.PrivateIpAddresses))
		for i := range instance.PrivateIpAddresses {
			if instance.PrivateIpAddresses[i] != nil {
				strIPs[i] = *instance.PrivateIpAddresses[i]
			}
		}

		m := map[string]interface{}{
			"instance_id":           *instance.InstanceId,
			"instance_name":         *instance.InstanceName,
			"create_time":           *instance.CreatedTime,
			"zone":                  *instance.Placement.Zone,
			"instance_status":       *instance.Status,
			"cpu_arch":              *instance.CpuArch,
			"operating_system_type": *instance.OperatingSystemType,
			"vpc_id":                *instance.VirtualPrivateCloud.VpcId,
			"private_ip_address":    strings.Join(strIPs, ","),
			"subnet_id":             *instance.VirtualPrivateCloud.SubnetId,
		}
		instanceList = append(instanceList, m)
		ids = append(ids, *instance.InstanceId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("instance_list", instanceList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set instance list fail, reason:%s\n ", logId, err.Error())
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), instanceList); err != nil {
			return err
		}
	}

	return nil
}
