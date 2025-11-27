/*
Use this data source to query detailed information of clb listeners_by_targets

Example Usage
```hcl

	data "tencentcloudenterprise_clb_listeners_by_targets" "listeners_by_targets" {
	  backends {
	    vpc_id     = "vpc-4owdpnwr"
	    private_ip = "203.0.113.10"
	  }
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
	// registerDataDescriptionProvider("tencentcloudenterprise_clb_listeners_by_targets", CNDescription{
	// 	TerraformTypeCN: "过滤信息查询CLB监听器",
	// 	AttributesCN: map[string]string{
	// 		"backends":         "要查询的专用网络IP的列表",
	// 		"vpc_id":           "VPC ID",
	// 		"private_ip":       "要查询的专用网络IP，可以是CVM或ENI",
	// 		"load_balancers":   "CLB实例的详细信息",
	// 		"load_balancer_id": "CLB实例的字符串ID",
	// 		"vip":              "CLB实例的VIP",
	// 		"listeners":        "侦听器规则",
	// 		"listener_id":      "侦听器ID",
	// 		"protocol":         "侦听器协议",
	// 		"port":             "侦听器端口",
	// 		"rules":            "绑定规则注意：此字段可能返回null，表示无法获得有效值",
	// 		"location_id":      "规则ID",
	// 		"domain":           "域名",
	// 		"url":              "url",
	// 		"targets":          "对象绑定到真实服务器",
	// 		"type":             "专用网络IP类型，可以是cvm或eni",
	// 		"weight":           "实际服务器的重量注意：此字段可能返回null，表示无法获得有效值",
	// 		"end_port":         "侦听器的结束端口注意：此字段可能返回null，表示无法获得有效值",
	// 	},
	// })
}

func dataSourceTencentCloudClbListenersByTargets() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of clb listeners_by_targets",
		Read:        dataSourceTencentCloudClbListenersByTargetsRead,
		Schema: map[string]*schema.Schema{
			"backends": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "List of private network IPs to be queried.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "VPC ID.",
						},
						"private_ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Private network IP to be queried, which can be of the CVM or ENI.",
						},
					},
				},
			},

			"load_balancers": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Detail of the CLB instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "String ID of the CLB instance.",
						},
						"vip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VIP of the CLB instance.",
						},
						"listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Listener rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"listener_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Listener ID.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Listener protocol.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Listener port.",
									},
									"rules": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Bound rule. Note: this field may return null, indicating that no valid values can be obtained.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"location_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Rule ID.",
												},
												"domain": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Domain name.",
												},
												"url": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Url.",
												},
												"targets": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Object bound to the real server.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Private network IP type, which can be cvm or eni.",
															},
															"private_ip": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Private network IP of the real server.",
															},
															"port": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Port bound to the real server.",
															},
															"vpc_id": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "VPC ID of the real server. Note: this field may return null, indicating that no valid values can be obtained.",
															},
															"weight": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Weight of the real server. Note: this field may return null, indicating that no valid values can be obtained.",
															},
														},
													},
												},
											},
										},
									},
									"targets": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Object bound to the layer-4 listener. Note: this field may return null, indicating that no valid values can be obtained.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Private network IP type, which can be cvm or eni.",
												},
												"private_ip": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Private network IP of the real server.",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Port bound to the real server.",
												},
												"vpc_id": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "VPC ID of the real server. Note: this field may return null, indicating that no valid values can be obtained.",
												},
												"weight": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Weight of the real server. Note: this field may return null, indicating that no valid values can be obtained.",
												},
											},
										},
									},
									"end_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "End port of the listener. Note: this field may return null, indicating that no valid values can be obtained.",
									},
								},
							},
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CLB instance region.",
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

func dataSourceTencentCloudClbListenersByTargetsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_clb_listeners_by_targets.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("backends"); ok {
		backendsSet := v.([]interface{})
		tmpSet := make([]*clb.LbRsItem, 0, len(backendsSet))

		for _, item := range backendsSet {
			lbRsItem := clb.LbRsItem{}
			lbRsItemMap := item.(map[string]interface{})

			if v, ok := lbRsItemMap["vpc_id"]; ok {
				lbRsItem.VpcId = helper.String(v.(string))
			}
			if v, ok := lbRsItemMap["private_ip"]; ok {
				lbRsItem.PrivateIp = helper.String(v.(string))
			}
			tmpSet = append(tmpSet, &lbRsItem)
		}
		paramMap["Backends"] = tmpSet
	}

	service := ClbService{client: meta.(*TencentCloudClient).apiV3Conn}

	var loadBalancers []*clb.LBItem

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeClbListenersByTargets(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		loadBalancers = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(loadBalancers))
	tmpList := make([]map[string]interface{}, 0, len(loadBalancers))

	if loadBalancers != nil {
		for _, lBItem := range loadBalancers {
			lBItemMap := map[string]interface{}{}

			if lBItem.LoadBalancerId != nil {
				lBItemMap["load_balancer_id"] = lBItem.LoadBalancerId
			}

			if lBItem.Vip != nil {
				lBItemMap["vip"] = lBItem.Vip
			}

			if lBItem.Listeners != nil {
				listenersList := []interface{}{}
				for _, listeners := range lBItem.Listeners {
					listenersMap := map[string]interface{}{}

					if listeners.ListenerId != nil {
						listenersMap["listener_id"] = listeners.ListenerId
					}

					if listeners.Protocol != nil {
						listenersMap["protocol"] = listeners.Protocol
					}

					if listeners.Port != nil {
						listenersMap["port"] = listeners.Port
					}

					if listeners.Rules != nil {
						rulesList := []interface{}{}
						for _, rules := range listeners.Rules {
							rulesMap := map[string]interface{}{}

							if rules.LocationId != nil {
								rulesMap["location_id"] = rules.LocationId
							}

							if rules.Domain != nil {
								rulesMap["domain"] = rules.Domain
							}

							if rules.Url != nil {
								rulesMap["url"] = rules.Url
							}

							if rules.Targets != nil {
								targetsList := []interface{}{}
								for _, targets := range rules.Targets {
									targetsMap := map[string]interface{}{}

									if targets.Type != nil {
										targetsMap["type"] = targets.Type
									}

									if targets.PrivateIp != nil {
										targetsMap["private_ip"] = targets.PrivateIp
									}

									if targets.Port != nil {
										targetsMap["port"] = targets.Port
									}

									if targets.VpcId != nil {
										targetsMap["vpc_id"] = targets.VpcId
									}

									if targets.Weight != nil {
										targetsMap["weight"] = targets.Weight
									}

									targetsList = append(targetsList, targetsMap)
								}

								rulesMap["targets"] = targetsList
							}

							rulesList = append(rulesList, rulesMap)
						}

						listenersMap["rules"] = rulesList
					}

					if listeners.Targets != nil {
						targetsList := []interface{}{}
						for _, targets := range listeners.Targets {
							targetsMap := map[string]interface{}{}

							if targets.Type != nil {
								targetsMap["type"] = targets.Type
							}

							if targets.PrivateIp != nil {
								targetsMap["private_ip"] = targets.PrivateIp
							}

							if targets.Port != nil {
								targetsMap["port"] = targets.Port
							}

							if targets.VpcId != nil {
								targetsMap["vpc_id"] = targets.VpcId
							}

							if targets.Weight != nil {
								targetsMap["weight"] = targets.Weight
							}

							targetsList = append(targetsList, targetsMap)
						}

						listenersMap["targets"] = targetsList
					}

					if listeners.EndPort != nil {
						listenersMap["end_port"] = listeners.EndPort
					}

					listenersList = append(listenersList, listenersMap)
				}

				lBItemMap["listeners"] = listenersList
			}

			if lBItem.Region != nil {
				lBItemMap["region"] = lBItem.Region
			}

			ids = append(ids, *lBItem.LoadBalancerId)
			tmpList = append(tmpList, lBItemMap)
		}

		_ = d.Set("load_balancers", tmpList)
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
