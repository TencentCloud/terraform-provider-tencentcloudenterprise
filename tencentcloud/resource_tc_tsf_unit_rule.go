/*
Provides a resource to create a tsf unit_rule

# Example Usage

```hcl

	resource "tencentcloudenterprise_tsf_unit_rule" "unit_rule" {
	  gateway_instance_id = "gw-ins-rug79a70"
	  name = "terraform-test"
	  description = "terraform-desc"
	  unit_rule_item_list {
			relationship = "AND"
			dest_namespace_id = "namespace-y8p88eka"
			dest_namespace_name = "garden-test_default"
			name = "Rule1"
			description = "rule1-desc"
			unit_rule_tag_list {
				tag_type = "U"
				tag_field = "aaa"
				tag_operator = "IN"
				tag_value = "1"
			}

	  }
	}

```

# Import

tsf unit_rule can be imported using the id, e.g.

```
terraform import cloud_tsf_unit_rule.unit_rule unit-rl-zbywqeca
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tsf_unit_rule", CNDescription{
		TerraformTypeCN: "单元化规则",
		DescriptionCN:   "提供TSF单元化规则资源，用于创建和管理单元化规则。",
		AttributesCN: map[string]string{
			"unit_rule_id":        "单元化规则ID",
			"gateway_instance_id": "网关实体ID",
			"name":                "规则名称",
			"status":              "使用状态：enabled/disabled",
			"description":         "规则描述",
			"unit_rule_item_list": "规则项列表",
			"rule_id":             "规则ID",
			"dest_namespace_id":   "目标命名空间ID",
			"dest_namespace_name": "目标命名空间名称",
			"priority":            "规则顺序越小，优先级越高，默认值为0",
			"relationship":        "逻辑关系和/或",
			"tag_field":           "标签名称",
			"tag_operator":        "运算符IN/NOT_IN/EQUAL/NOT_EQUAL/REGEX",
			"tag_type":            "标签类型U（用户标签）",
			"tag_value":           "标签值",
			"unit_rule_item_id":   "单位化规则项ID",
			"unit_rule_tag_list":  "规则标签列表",
		},
	})
}

func resourceTencentCloudTsfUnitRule() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tsf unit_rule",
		Create:      resourceTencentCloudTsfUnitRuleCreate,
		Read:        resourceTencentCloudTsfUnitRuleRead,
		Update:      resourceTencentCloudTsfUnitRuleUpdate,
		Delete:      resourceTencentCloudTsfUnitRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"gateway_instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Gateway entity ID.",
			},

			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Rule name.",
			},

			"rule_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Rule ID.",
			},

			"status": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Usage status: enabled/disabled.",
			},

			"description": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Rule description.",
			},

			"unit_rule_item_list": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "List of rule items.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"relationship": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Logical relationship: AND/OR.",
						},
						"dest_namespace_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Destination namespace ID.",
						},
						"dest_namespace_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Destination namespace name.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Rule item name.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Rule item ID.",
						},
						"unit_rule_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Unitization rule ID.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Rule order, the smaller the higher the priority: the default is 0.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Rule description.",
						},
						"unit_rule_tag_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of rule labels.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Tag Type: U(User Tag).",
									},
									"tag_field": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Label name.",
									},
									"tag_operator": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Operator: IN/NOT_IN/EQUAL/NOT_EQUAL/REGEX.",
									},
									"tag_value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Tag value.",
									},
									"unit_rule_item_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Unitization rule item ID.",
									},
									"rule_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Rule ID.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudTsfUnitRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_unit_rule.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = tsf.NewCreateUnitRuleRequest()
		//response = tsf.NewCreateUnitRuleResponse()
		ruleId string
	)
	if v, ok := d.GetOk("gateway_instance_id"); ok {
		request.GatewayInstanceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOk("unit_rule_item_list"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			unitRuleItem := tsf.UnitRuleItem{}
			if v, ok := dMap["relationship"]; ok {
				unitRuleItem.Relationship = helper.String(v.(string))
			}
			if v, ok := dMap["dest_namespace_id"]; ok {
				unitRuleItem.DestNamespaceId = helper.String(v.(string))
			}
			if v, ok := dMap["dest_namespace_name"]; ok {
				unitRuleItem.DestNamespaceName = helper.String(v.(string))
			}
			if v, ok := dMap["name"]; ok {
				unitRuleItem.Name = helper.String(v.(string))
			}
			if v, ok := dMap["rule_id"]; ok {
				unitRuleItem.Id = helper.String(v.(string))
			}
			if v, ok := dMap["unit_rule_id"]; ok {
				unitRuleItem.UnitRuleId = helper.String(v.(string))
			}
			if v, ok := dMap["priority"]; ok {
				unitRuleItem.Priority = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["description"]; ok {
				unitRuleItem.Description = helper.String(v.(string))
			}
			if v, ok := dMap["unit_rule_tag_list"]; ok {
				for _, item := range v.([]interface{}) {
					unitRuleTagListMap := item.(map[string]interface{})
					unitRuleTag := tsf.UnitRuleTag{}
					if v, ok := unitRuleTagListMap["tag_type"]; ok {
						unitRuleTag.TagType = helper.String(v.(string))
					}
					if v, ok := unitRuleTagListMap["tag_field"]; ok {
						unitRuleTag.TagField = helper.String(v.(string))
					}
					if v, ok := unitRuleTagListMap["tag_operator"]; ok {
						unitRuleTag.TagOperator = helper.String(v.(string))
					}
					if v, ok := unitRuleTagListMap["tag_value"]; ok {
						unitRuleTag.TagValue = helper.String(v.(string))
					}
					if v, ok := unitRuleTagListMap["unit_rule_item_id"]; ok {
						unitRuleTag.UnitRuleItemId = helper.String(v.(string))
					}
					if v, ok := unitRuleTagListMap["rule_id"]; ok {
						unitRuleTag.Id = helper.String(v.(string))
					}
					unitRuleItem.UnitRuleTagList = append(unitRuleItem.UnitRuleTagList, &unitRuleTag)
				}
			}
			request.UnitRuleItemList = append(request.UnitRuleItemList, &unitRuleItem)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().CreateUnitRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		//response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tsf unitRule failed, reason:%+v", logId, err)
		return err
	}

	// 创建接口没有返回实例id ，通过查询接口获取
	req := tsf.NewDescribeUnitRulesV2Request()
	req.Offset = helper.Int64(0)
	req.Limit = helper.Int64(20)
	req.SearchWord = request.Name
	req.GatewayInstanceId = request.GatewayInstanceId
	result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().DescribeUnitRulesV2(req)
	if e != nil {
		return e
	}
	if *result.Response.Result.TotalCount != 1 {
		return fmt.Errorf("query unitRule by name [%s] failed", *request.Name)
	}
	ruleId = *result.Response.Result.Content[0].Id
	//ruleId = *response.Response.Result.Id
	d.SetId(ruleId)

	return resourceTencentCloudTsfUnitRuleRead(d, meta)
}

func resourceTencentCloudTsfUnitRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_unit_rule.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	id := d.Id()

	unitRule, err := service.DescribeTsfUnitRuleById(ctx, id)
	if err != nil {
		return err
	}

	if unitRule == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TsfUnitRule` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if unitRule.GatewayInstanceId != nil {
		_ = d.Set("gateway_instance_id", unitRule.GatewayInstanceId)
	}

	if unitRule.Name != nil {
		_ = d.Set("name", unitRule.Name)
	}

	if unitRule.Id != nil {
		_ = d.Set("rule_id", unitRule.Id)
	}

	if unitRule.Status != nil {
		_ = d.Set("status", unitRule.Status)
	}

	if unitRule.Description != nil {
		_ = d.Set("description", unitRule.Description)
	}

	if unitRule.UnitRuleItemList != nil {
		unitRuleItemListList := []interface{}{}
		for _, unitRuleItemList := range unitRule.UnitRuleItemList {
			unitRuleItemListMap := map[string]interface{}{}

			if unitRuleItemList.Relationship != nil {
				unitRuleItemListMap["relationship"] = unitRuleItemList.Relationship
			}

			if unitRuleItemList.DestNamespaceId != nil {
				unitRuleItemListMap["dest_namespace_id"] = unitRuleItemList.DestNamespaceId
			}

			if unitRuleItemList.DestNamespaceName != nil {
				unitRuleItemListMap["dest_namespace_name"] = unitRuleItemList.DestNamespaceName
			}

			if unitRuleItemList.Name != nil {
				unitRuleItemListMap["name"] = unitRuleItemList.Name
			}

			if unitRuleItemList.Id != nil {
				unitRuleItemListMap["rule_id"] = unitRuleItemList.Id
			}

			if unitRuleItemList.UnitRuleId != nil {
				unitRuleItemListMap["unit_rule_id"] = unitRuleItemList.UnitRuleId
			}

			if unitRuleItemList.Priority != nil {
				unitRuleItemListMap["priority"] = unitRuleItemList.Priority
			}

			if unitRuleItemList.Description != nil {
				unitRuleItemListMap["description"] = unitRuleItemList.Description
			}

			if unitRuleItemList.UnitRuleTagList != nil {
				unitRuleTagListList := []interface{}{}
				for _, unitRuleTagList := range unitRuleItemList.UnitRuleTagList {
					unitRuleTagListMap := map[string]interface{}{}

					if unitRuleTagList.TagType != nil {
						unitRuleTagListMap["tag_type"] = unitRuleTagList.TagType
					}

					if unitRuleTagList.TagField != nil {
						unitRuleTagListMap["tag_field"] = unitRuleTagList.TagField
					}

					if unitRuleTagList.TagOperator != nil {
						unitRuleTagListMap["tag_operator"] = unitRuleTagList.TagOperator
					}

					if unitRuleTagList.TagValue != nil {
						unitRuleTagListMap["tag_value"] = unitRuleTagList.TagValue
					}

					if unitRuleTagList.UnitRuleItemId != nil {
						unitRuleTagListMap["unit_rule_item_id"] = unitRuleTagList.UnitRuleItemId
					}

					if unitRuleTagList.Id != nil {
						unitRuleTagListMap["rule_id"] = unitRuleTagList.Id
					}

					unitRuleTagListList = append(unitRuleTagListList, unitRuleTagListMap)
				}

				unitRuleItemListMap["unit_rule_tag_list"] = unitRuleTagListList
			}

			unitRuleItemListList = append(unitRuleItemListList, unitRuleItemListMap)
		}

		err = d.Set("unit_rule_item_list", unitRuleItemListList)
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceTencentCloudTsfUnitRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_unit_rule.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := tsf.NewUpdateUnitRuleRequest()

	id := d.Id()

	request.Id = &id

	immutableArgs := []string{"gateway_instance_id", "rule_id", "status"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if d.HasChange("description") {
		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}
	}

	if d.HasChange("unit_rule_item_list") {
		if v, ok := d.GetOk("unit_rule_item_list"); ok {
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				unitRuleItem := tsf.UnitRuleItem{}
				if v, ok := dMap["relationship"]; ok {
					unitRuleItem.Relationship = helper.String(v.(string))
				}
				if v, ok := dMap["dest_namespace_id"]; ok {
					unitRuleItem.DestNamespaceId = helper.String(v.(string))
				}
				if v, ok := dMap["dest_namespace_name"]; ok {
					unitRuleItem.DestNamespaceName = helper.String(v.(string))
				}
				if v, ok := dMap["name"]; ok {
					unitRuleItem.Name = helper.String(v.(string))
				}
				if v, ok := dMap["rule_id"]; ok {
					unitRuleItem.Id = helper.String(v.(string))
				}
				if v, ok := dMap["unit_rule_id"]; ok {
					unitRuleItem.UnitRuleId = helper.String(v.(string))
				}
				if v, ok := dMap["priority"]; ok {
					unitRuleItem.Priority = helper.IntInt64(v.(int))
				}
				if v, ok := dMap["description"]; ok {
					unitRuleItem.Description = helper.String(v.(string))
				}
				if v, ok := dMap["unit_rule_tag_list"]; ok {
					for _, item := range v.([]interface{}) {
						unitRuleTagListMap := item.(map[string]interface{})
						unitRuleTag := tsf.UnitRuleTag{}
						if v, ok := unitRuleTagListMap["tag_type"]; ok {
							unitRuleTag.TagType = helper.String(v.(string))
						}
						if v, ok := unitRuleTagListMap["tag_field"]; ok {
							unitRuleTag.TagField = helper.String(v.(string))
						}
						if v, ok := unitRuleTagListMap["tag_operator"]; ok {
							unitRuleTag.TagOperator = helper.String(v.(string))
						}
						if v, ok := unitRuleTagListMap["tag_value"]; ok {
							unitRuleTag.TagValue = helper.String(v.(string))
						}
						if v, ok := unitRuleTagListMap["unit_rule_item_id"]; ok {
							unitRuleTag.UnitRuleItemId = helper.String(v.(string))
						}
						if v, ok := unitRuleTagListMap["rule_id"]; ok {
							unitRuleTag.Id = helper.String(v.(string))
						}
						unitRuleItem.UnitRuleTagList = append(unitRuleItem.UnitRuleTagList, &unitRuleTag)
					}
				}
				request.UnitRuleItemList = append(request.UnitRuleItemList, &unitRuleItem)
			}
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().UpdateUnitRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update tsf unitRule failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTsfUnitRuleRead(d, meta)
}

func resourceTencentCloudTsfUnitRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_unit_rule.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}
	id := d.Id()

	if err := service.DeleteTsfUnitRuleById(ctx, id); err != nil {
		return err
	}

	return nil
}
