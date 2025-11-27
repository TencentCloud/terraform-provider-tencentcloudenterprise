/*
Use this data source to query detailed information of clb instance_detail

# Example Usage

```hcl

	data "tencentcloudenterprise_clb_instance_detail" "instance_detail" {
	  target_type = "NODE"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_clb_instance_detail", CNDescription{
		TerraformTypeCN: "负载均衡实例详情",
		DescriptionCN:   "提供负载均衡实例详情数据源，用于查询CLB实例的详细配置信息。",
		AttributesCN: map[string]string{
			"fields":                       "字段列表将只返回指定的字段如果保留为空，则返回“null”默认情况下会添加字段“LoadBalancerId”和“LoadBalancirName”有关字段的详细信息",
			"target_type":                  "目标类型有效值：NODE和GROUP如果字段列表包含“TargetId”、“TargetAddress”、“TargetPort”、“Target Weight”和其他字段，则必须导出目标组或非目标组的“Target”",
			"filters":                      "查询描述CLB实例详细信息列表的过滤条件：负载均衡器id-字符串-必选项：no-（过滤条件）CLB实例id，如lb-12345678；project id-String-必选：no-（过滤条件）项目id，如0、123；network-String-必需：no-（Filter condition）CLB实例的网络类型，例如Public和Private&amp；lt/李；gt&amp；lt；李；gt；vip-String-必选：no-（过滤条件）CLB实例vip，如1.1.1.1、2204:：22:3；target ip-String-必选：no-（过滤条件）目标真实服务器的私有ip，如1.1.1.1和2203:：214:4；vpcid-String-必选：no-（过滤条件）CLB实例所属VPC实例的标识，如VPC-12345678；zone-String-必选：no-（过滤条件）CLB实例所在的可用区，如ap-广州一号；tag key-String-必选：no-（过滤条件）CLB实例的tag key，如name；tag:*-String-必需：no-（Filter condition）CLB实例标记，后面跟着冒号后面的标记键例如，使用{Name:tag:Name，Values:neneneba张三，lisi]}筛选具有标记值`nenenebb张三`和`lisi`的标记键`Name`；模糊搜索-字符串-必需：否-（过滤条件）模糊搜索CLB实例VIP和CLB实例名称，如1",
			"name":                         "过滤器名称",
			"筛选值数组s":                       "筛选值数组",
			"load_balancer_detail_set":     "CLB实例详细信息列表注意：此字段可能返回null，表示无法获得有效值",
			"load_balancer_id":             "CLB实例ID",
			"load_balancer_name":           "CLB实例名称",
			"load_balancer_type":           "CLB实例网络类型：公网：公网；专用：专用网络注意：此字段可能返回null，表示无法获得有效值",
			"status":                       "CLB实例状态，包括：0:正在创建；1：正在运行注意：此字段可能返回null，表示无法获得有效值",
			"address":                      "CLB实例VIP注意：此字段可能返回null，表示无法获得有效值",
			"address_ipv6":                 "CLB实例的IPv6 VIP地址注意：此字段可能返回null，表示无法获得有效值",
			"address_ip_version":           "CLB实例的IP版本有效值：IPv4、IPv6注意：此字段可能返回null，表示无法获得有效值",
			"ipv6_mode":                    "CLB实例的IPv6地址类型有效值：IPv6Nat64、IPv6FullChain注意：此字段可能返回null，表示无法获得有效值",
			"zone":                         "CLB实例所在的可用性区域注意：此字段可能返回null，表示无法获得有效值",
			"address_isp":                  "CLB IP地址所属的ISP注意：此字段可能返回null，表示无法获得有效值",
			"vpc_id":                       "CLB实例所属VPC实例的ID注意：此字段可能返回null，表示无法获得有效值",
			"project_id":                   "CLB实例所属项目的ID0：默认项目注意：此字段可能返回null，表示无法获得有效值",
			"create_time":                  "CLB实例创建时间注意：此字段可能返回null，表示无法获得有效值",
			"charge_type":                  "CLB实例计费模式注意：此字段可能返回null，表示无法获得有效值",
			"network_attributes":           "CLB实例网络属性注意：此字段可能返回null，表示无法获得有效值",
			"internet_charge_type":         "TRAFFIC_POSTPAID_BY_HOUR：按小时计酬；BANDWITH_POSTPAID_BY_HOUR：按带宽按小时计酬；BANDWITH_PACKAGE：按带宽包计费（目前，只有指定ISP时才支持此方法）",
			"internet_max_bandwidth_out":   "最大出站带宽（Mbps），仅适用于公共网络CLB取值范围：0-65535默认值：10",
			"bandwidth_pkg_sub_type":       "带宽包类型，如SINGLIESP注意：此字段可能返回null，表示无法获得有效值",
			"prepaid_attributes":           "CLB实例的现收现付属性注意：此字段可能返回null，表示无法获得有效值",
			"renew_flag":                   "续订类型AUTO_RENEW：自动续费；MANUAL_REQUE:手动续订注意：此字段可能返回null，表示无法获得有效值",
			"period":                       "周期，表示月数（保留字段）注意：此字段可能返回null，表示无法获得有效值",
			"extra_info":                   "保留字段，一般可以忽略注意：此字段可能返回null，表示无法获得有效值",
			"zhi_tong":                     "是否启用VIP直连注意：此字段可能返回null，表示无法获取有效值",
			"tgw_group_name":               "TgwGroup名称注意：此字段可能返回null，表示无法获得有效值",
			"config_id":                    "CLB实例的自定义配置ID多个ID必须用逗号（，）分隔注意：此字段可能返回null，表示无法获得有效值",
			"tags":                         "CLB实例标记信息注意：此字段可能返回null，表示无法获得有效值",
			"tag_key":                      "标记键",
			"tag_标记值":                      "标记值",
			"listener_id":                  "CLB侦听器ID注意：此字段可能返回null，表示无法获得有效值",
			"protocol":                     "侦听器协议注意：此字段可能返回null，表示无法获得有效值",
			"port":                         "侦听器端口注意：此字段可能返回null，表示无法获得有效值",
			"location_id":                  "转发规则ID注意：此字段可能返回null，表示无法获得有效值",
			"domain":                       "转发规则的域名注意：此字段可能返回null，表示无法获得有效值",
			"url":                          "转发规则路径注意：此字段可能返回null，表示无法获得有效值",
			"target_id":                    "目标实际服务器的ID注意：此字段可能返回null，表示无法获得有效值",
			"target_address":               "目标实际服务器的地址注意：此字段可能返回null，表示无法获得有效值",
			"target_port":                  "目标真实服务器的侦听端口注意：此字段可能返回null，表示无法获得有效值",
			"target_weight":                "目标真实服务器的转发权重注意：此字段可能返回null，表示无法获得有效值",
			"isolation":                    "0：未隔离；1：隔离注意：此字段可能返回null，表示无法获得有效值",
			"security_group":               "绑定到CLB实例的安全组的列表注意：此字段可能返回“null”，表示无法获得有效值",
			"load_balancer_pass_to_target": "CLB实例是否按IP计费注意：此字段可能返回“null”，表示无法获得有效值",
			"target_health":                "目标真实服务器的运行状况注意：此字段可能返回“null”，表示无法获得有效值",
			"domains":                      "与转发规则关联的域名列表注意：此字段可能返回“null”，表示无法获得有效值",
			"slave_zone":                   "多AZ CLB实例的辅助区域注意：此字段可能返回“null”，表示无法获得有效值",
			"zones":                        "私有CLB实例的AZ这只适用于测试版用户注意：此字段可能返回“null”，表示无法获得有效值",
			"sni_switch":                   "是否启用SNI此参数仅对HTTPS侦听器有意义注意：此字段可能返回“null”，表示无法获得有效值",
			"load_balancer_domain":         "CLB实例的域名注意：此字段可能返回null，表示无法获得有效值",
		},
	})
}

func dataSourceTencentCloudClbInstanceDetail() *schema.Resource {
	return &schema.Resource{
		Description: "Cloud Clb Instance Detail",
		Read:        dataSourceTencentCloudClbInstanceDetailRead,
		Schema: map[string]*schema.Schema{
			"fields": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of fields. Only fields specified will be returned. If it's left blank, `null` is returned. The fields `LoadBalancerId` and `LoadBalancerName` are added by default. For details about fields.",
			},

			"target_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Target type. Valid values: NODE and GROUP. If the list of fields contains `TargetId`, `TargetAddress`, `TargetPort`, `TargetWeight` and other fields, `Target` of the target group or non-target group must be exported.",
			},

			"filters": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Filter condition of querying lists describing CLB instance details:loadbalancer-id - String - Required: no - (Filter condition) CLB instance ID, such as lb-12345678; project-id - String - Required: no - (Filter condition) Project ID, such as 0 and 123; network - String - Required: no - (Filter condition) Network type of the CLB instance, such as Public and Private.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt; vip - String - Required: no - (Filter condition) CLB instance VIP, such as 1.1.1.1 and 2204::22:3; target-ip - String - Required: no - (Filter condition) Private IP of the target real servers, such as1.1.1.1 and 2203::214:4; vpcid - String - Required: no - (Filter condition) Identifier of the VPC instance to which the CLB instance belongs, such as vpc-12345678; zone - String - Required: no - (Filter condition) Availability zone where the CLB instance resides, such as ap-guangzhou-1; tag-key - String - Required: no - (Filter condition) Tag key of the CLB instance, such as name; tag:* - String - Required: no - (Filter condition) CLB instance tag, followed by tag key after the colon. For example, use {Name: tag:name,Values: [zhangsan, lisi]} to filter the tag key `name` with the tag value `zhangsan` and `lisi`; fuzzy-search - String - Required: no - (Filter condition) Fuzzy search for CLB instance VIP and CLB instance name, such as 1.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Filter name.",
						},
						"values": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required:    true,
							Description: "Filter value array.",
						},
					},
				},
			},

			"load_balancer_detail_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "List of CLB instance details.Note: this field may return null, indicating that no valid values can be obtained.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLB instance ID.",
						},
						"load_balancer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLB instance name.",
						},
						"load_balancer_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLB instance network type:Public: public network; Private: private network.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "CLB instance status, including:0: creating; 1: running.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLB instance VIP.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"address_ipv6": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 VIP address of the CLB instance.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"address_ip_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP version of the CLB instance. Valid values: IPv4, IPv6.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"ipv6_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address type of the CLB instance. Valid values: IPv6Nat64, IPv6FullChain.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Availability zone where the CLB instance resides.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"address_isp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ISP to which the CLB IP address belongs.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the VPC instance to which the CLB instance belongs.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"project_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of the project to which the CLB instance belongs. 0: default project.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLB instance creation time.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLB instance billing mode.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"network_attributes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "CLB instance network attribute.Note: this field may return null, indicating that no valid values can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"internet_charge_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "TRAFFIC_POSTPAID_BY_HOUR: hourly pay-as-you-go by traffic; BANDWIDTH_POSTPAID_BY_HOUR: hourly pay-as-you-go by bandwidth;BANDWIDTH_PACKAGE: billed by bandwidth package (currently, this method is supported only if the ISP is specified).",
									},
									"internet_max_bandwidth_out": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Maximum outbound bandwidth in Mbps, which applies only to public network CLB. Value range: 0-65,535. Default value: 10.",
									},
									"bandwidth_pkg_sub_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Bandwidth package type, such as SINGLEISPNote: This field may return null, indicating that no valid values can be obtained.",
									},
								},
							},
						},
						"prepaid_attributes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Pay-as-you-go attribute of the CLB instance.Note: this field may return null, indicating that no valid values can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"renew_flag": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Renewal type. AUTO_RENEW: automatic renewal; MANUAL_RENEW: manual renewalNote: This field may return null, indicating that no valid values can be obtained.",
									},
									"period": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Cycle, indicating the number of months (reserved field)Note: This field may return null, indicating that no valid values can be obtained.",
									},
								},
							},
						},
						"extra_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Reserved field, which can be ignored generally.Note: this field may return null, indicating that no valid values can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zhi_tong": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable VIP direct connectionNote: This field may return null, indicating that no valid values can be obtained.",
									},
									"tgw_group_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "TgwGroup nameNote: This field may return null, indicating that no valid values can be obtained.",
									},
								},
							},
						},
						"config_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Custom configuration IDs of CLB instances. Multiple IDs must be separated by commas (,).Note: This field may return null, indicating that no valid values can be obtained.",
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "CLB instance tag information.Note: this field may return null, indicating that no valid values can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag key.",
									},
									"tag_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag value.",
									},
								},
							},
						},
						"listener_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLB listener ID.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Listener protocol.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Listener port.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"location_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forwarding rule ID.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name of the forwarding rule.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forwarding rule path.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"target_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of target real servers.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"target_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Address of target real servers.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"target_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Listening port of target real servers.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"target_weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Forwarding weight of target real servers.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"isolation": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "0: not isolated; 1: isolated.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"security_group": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "List of the security groups bound to the CLB instance.Note: this field may return `null`, indicating that no valid values can be obtained.",
						},
						"load_balancer_pass_to_target": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether the CLB instance is billed by IP.Note: this field may return `null`, indicating that no valid values can be obtained.",
						},
						"target_health": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Health status of the target real server.Note: this field may return `null`, indicating that no valid values can be obtained.",
						},
						"domains": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "List o domain names associated with the forwarding ruleNote: This field may return `null`, indicating that no valid values can be obtained.",
						},
						"slave_zone": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "The secondary zone of multi-AZ CLB instanceNote: This field may return `null`, indicating that no valid values can be obtained.",
						},
						"zones": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "The AZ of private CLB instance. This is only available for beta users.Note: This field may return `null`, indicating that no valid values can be obtained.",
						},
						"sni_switch": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether SNI is enabled. This parameter is only meaningful for HTTPS listeners.Note: This field may return `null`, indicating that no valid values can be obtained.",
						},
						"load_balancer_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name of the CLB instance.Note: This field may return null, indicating that no valid values can be obtained.",
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

func dataSourceTencentCloudClbInstanceDetailRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_clb_instance_detail.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("fields"); ok {
		fieldsSet := v.(*schema.Set).List()
		paramMap["Fields"] = helper.InterfacesStringsPoint(fieldsSet)
	}

	if v, ok := d.GetOk("target_type"); ok {
		paramMap["TargetType"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("filters"); ok {
		filtersSet := v.([]interface{})
		tmpSet := make([]*clb.Filter, 0, len(filtersSet))

		for _, item := range filtersSet {
			filter := clb.Filter{}
			filterMap := item.(map[string]interface{})

			if v, ok := filterMap["name"]; ok {
				filter.Name = helper.String(v.(string))
			}
			if v, ok := filterMap["values"]; ok {
				valuesSet := v.(*schema.Set).List()
				filter.Values = helper.InterfacesStringsPoint(valuesSet)
			}
			tmpSet = append(tmpSet, &filter)
		}
		paramMap["Filters"] = tmpSet
	}

	service := ClbService{client: meta.(*TencentCloudClient).apiV3Conn}

	var loadBalancerDetailSet []*clb.LoadBalancer

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeClbInstanceDetailByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		loadBalancerDetailSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(loadBalancerDetailSet))
	tmpList := make([]map[string]interface{}, 0, len(loadBalancerDetailSet))

	if loadBalancerDetailSet != nil {
		for _, loadBalancerDetail := range loadBalancerDetailSet {
			loadBalancerDetailMap := map[string]interface{}{}

			if loadBalancerDetail.LoadBalancerId != nil {
				loadBalancerDetailMap["load_balancer_id"] = loadBalancerDetail.LoadBalancerId
			}

			if loadBalancerDetail.LoadBalancerName != nil {
				loadBalancerDetailMap["load_balancer_name"] = loadBalancerDetail.LoadBalancerName
			}

			if loadBalancerDetail.LoadBalancerType != nil {
				loadBalancerDetailMap["load_balancer_type"] = loadBalancerDetail.LoadBalancerType
			}

			if loadBalancerDetail.Status != nil {
				loadBalancerDetailMap["status"] = loadBalancerDetail.Status
			}

			/*
				if loadBalancerDetail.Address != nil {
					loadBalancerDetailMap["address"] = loadBalancerDetail.Address
				}

			*/

			if loadBalancerDetail.AddressIPv6 != nil {
				loadBalancerDetailMap["address_ipv6"] = loadBalancerDetail.AddressIPv6
			}

			if loadBalancerDetail.AddressIPVersion != nil {
				loadBalancerDetailMap["address_ip_version"] = loadBalancerDetail.AddressIPVersion
			}

			if loadBalancerDetail.IPv6Mode != nil {
				loadBalancerDetailMap["ipv6_mode"] = loadBalancerDetail.IPv6Mode
			}

			/*
				if loadBalancerDetail.Zone != nil {
					loadBalancerDetailMap["zone"] = loadBalancerDetail.Zone
				}

			*/

			/*
				if loadBalancerDetail.AddressIsp != nil {
					loadBalancerDetailMap["address_isp"] = loadBalancerDetail.AddressIsp
				}

			*/

			if loadBalancerDetail.VpcId != nil {
				loadBalancerDetailMap["vpc_id"] = loadBalancerDetail.VpcId
			}

			if loadBalancerDetail.ProjectId != nil {
				loadBalancerDetailMap["project_id"] = loadBalancerDetail.ProjectId
			}

			if loadBalancerDetail.CreateTime != nil {
				loadBalancerDetailMap["create_time"] = loadBalancerDetail.CreateTime
			}

			if loadBalancerDetail.ChargeType != nil {
				loadBalancerDetailMap["charge_type"] = loadBalancerDetail.ChargeType
			}

			if loadBalancerDetail.NetworkAttributes != nil {
				networkAttributesMap := map[string]interface{}{}

				if loadBalancerDetail.NetworkAttributes.InternetChargeType != nil {
					networkAttributesMap["internet_charge_type"] = loadBalancerDetail.NetworkAttributes.InternetChargeType
				}

				if loadBalancerDetail.NetworkAttributes.InternetMaxBandwidthOut != nil {
					networkAttributesMap["internet_max_bandwidth_out"] = loadBalancerDetail.NetworkAttributes.InternetMaxBandwidthOut
				}

				if loadBalancerDetail.NetworkAttributes.BandwidthpkgSubType != nil {
					networkAttributesMap["bandwidth_pkg_sub_type"] = loadBalancerDetail.NetworkAttributes.BandwidthpkgSubType
				}

				loadBalancerDetailMap["network_attributes"] = []interface{}{networkAttributesMap}
			}

			if loadBalancerDetail.PrepaidAttributes != nil {
				prepaidAttributesMap := map[string]interface{}{}

				if loadBalancerDetail.PrepaidAttributes.RenewFlag != nil {
					prepaidAttributesMap["renew_flag"] = loadBalancerDetail.PrepaidAttributes.RenewFlag
				}

				if loadBalancerDetail.PrepaidAttributes.Period != nil {
					prepaidAttributesMap["period"] = loadBalancerDetail.PrepaidAttributes.Period
				}

				loadBalancerDetailMap["prepaid_attributes"] = []interface{}{prepaidAttributesMap}
			}

			if loadBalancerDetail.ExtraInfo != nil {
				extraInfoMap := map[string]interface{}{}

				if loadBalancerDetail.ExtraInfo.ZhiTong != nil {
					extraInfoMap["zhi_tong"] = loadBalancerDetail.ExtraInfo.ZhiTong
				}

				if loadBalancerDetail.ExtraInfo.TgwGroupName != nil {
					extraInfoMap["tgw_group_name"] = loadBalancerDetail.ExtraInfo.TgwGroupName
				}

				loadBalancerDetailMap["extra_info"] = []interface{}{extraInfoMap}
			}

			if loadBalancerDetail.ConfigId != nil {
				loadBalancerDetailMap["config_id"] = loadBalancerDetail.ConfigId
			}

			if loadBalancerDetail.Tags != nil {
				tagsList := []interface{}{}
				for _, tags := range loadBalancerDetail.Tags {
					tagsMap := map[string]interface{}{}

					if tags.TagKey != nil {
						tagsMap["tag_key"] = tags.TagKey
					}

					if tags.TagValue != nil {
						tagsMap["tag_value"] = tags.TagValue
					}

					tagsList = append(tagsList, tagsMap)
				}

				loadBalancerDetailMap["tags"] = tagsList
			}

			/*
				if loadBalancerDetail.ListenerId != nil {
					loadBalancerDetailMap["listener_id"] = loadBalancerDetail.ListenerId
				}

			*/

			/*
				if loadBalancerDetail.Protocol != nil {
					loadBalancerDetailMap["protocol"] = loadBalancerDetail.Protocol
				}

			*/

			/*
				if loadBalancerDetail.Port != nil {
					loadBalancerDetailMap["port"] = loadBalancerDetail.Port
				}

			*/

			/*
				if loadBalancerDetail.LocationId != nil {
					loadBalancerDetailMap["location_id"] = loadBalancerDetail.LocationId
				}

			*/

			if loadBalancerDetail.Domain != nil {
				loadBalancerDetailMap["domain"] = loadBalancerDetail.Domain
			}

			//if loadBalancerDetail.Url != nil {
			//	loadBalancerDetailMap["url"] = loadBalancerDetail.Url
			//}
			//
			//if loadBalancerDetail.TargetId != nil {
			//	loadBalancerDetailMap["target_id"] = loadBalancerDetail.TargetId
			//}
			//
			//if loadBalancerDetail.TargetAddress != nil {
			//	loadBalancerDetailMap["target_address"] = loadBalancerDetail.TargetAddress
			//}
			//
			//if loadBalancerDetail.TargetPort != nil {
			//	loadBalancerDetailMap["target_port"] = loadBalancerDetail.TargetPort
			//}
			//
			//if loadBalancerDetail.TargetWeight != nil {
			//	loadBalancerDetailMap["target_weight"] = loadBalancerDetail.TargetWeight
			//}

			if loadBalancerDetail.Isolation != nil {
				loadBalancerDetailMap["isolation"] = loadBalancerDetail.Isolation
			}

			//if loadBalancerDetail.SecurityGroup != nil {
			//	loadBalancerDetailMap["security_group"] = loadBalancerDetail.SecurityGroup
			//}

			if loadBalancerDetail.LoadBalancerPassToTarget != nil {
				loadBalancerDetailMap["load_balancer_pass_to_target"] = loadBalancerDetail.LoadBalancerPassToTarget
			}

			//if loadBalancerDetail.TargetHealth != nil {
			//	loadBalancerDetailMap["target_health"] = loadBalancerDetail.TargetHealth
			//}

			//if loadBalancerDetail.Domains != nil {
			//	loadBalancerDetailMap["domains"] = loadBalancerDetail.Domains
			//}

			//if loadBalancerDetail.SlaveZone != nil {
			//	loadBalancerDetailMap["slave_zone"] = loadBalancerDetail.SlaveZone
			//}

			if loadBalancerDetail.Zones != nil {
				loadBalancerDetailMap["zones"] = loadBalancerDetail.Zones
			}

			//if loadBalancerDetail.SniSwitch != nil {
			//	loadBalancerDetailMap["sni_switch"] = loadBalancerDetail.SniSwitch
			//}

			//if loadBalancerDetail.LoadBalancerDomain != nil {
			//	loadBalancerDetailMap["load_balancer_domain"] = loadBalancerDetail.LoadBalancerDomain
			//}

			ids = append(ids, *loadBalancerDetail.LoadBalancerId)
			tmpList = append(tmpList, loadBalancerDetailMap)
		}

		_ = d.Set("load_balancer_detail_set", tmpList)
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
