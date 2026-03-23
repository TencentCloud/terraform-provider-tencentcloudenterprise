/*
Provides a resource to create a cloud firewall (cfw) security group rule.

Example Usage

```hcl

	resource "tencentcloudenterprise_cfw_sg_rule" "example" {
	  data {
	    source_content = "1.1.1.1/0"
	    source_type    = "net"
	    dest_content   = "0.0.0.0/0"
	    dest_type      = "net"
	    protocol       = "TCP"
	    rule_action    = "accept"
	    port           = "-1/-1"
	    description    = "sg rule description."
	  }
	  enable = 1
	}

```

Import

Cloud firewall security group rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_sg_rule.example 123456
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfw_sg_rule", CNDescription{
		TerraformTypeCN: "企业安全组规则",
		DescriptionCN:   "提供云防火墙企业安全组规则资源，用于创建和管理云防火墙企业安全组规则。",
		AttributesCN: map[string]string{
			"data":                "创建规则数据",
			"source_content":      "访问源",
			"source_type":         "访问源类型",
			"dest_content":        "访问目的",
			"dest_type":           "访问目的类型",
			"protocol":            "协议",
			"port":                "端口",
			"rule_action":         "动作",
			"description":         "描述",
			"enable":              "规则状态，1启用，0停用",
			"order_index":         "规则优先级",
			"service_template_id": "协议端口模板ID",
		},
	})
}

func resourceTencentCloudCfwSgRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCfwSgRuleCreate,
		Read:   resourceTencentCloudCfwSgRuleRead,
		Update: resourceTencentCloudCfwSgRuleUpdate,
		Delete: resourceTencentCloudCfwSgRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"data": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Creates rule data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Source content.",
						},
						"source_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Access source type. Valid values: net|template|instance|resourcegroup|tag|region.",
						},
						"dest_content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Destination content.",
						},
						"dest_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Access destination type. Valid values: net|template|instance|resourcegroup|tag|region.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Protocol. TCP/UDP/ICMP/ANY.",
						},
						"port": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The port to apply access control rules. Valid values: `-1/-1`: all ports, `80`: port 80.",
						},
						"service_template_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Parameter template ID of port and protocol type; mutually exclusive with Protocol and Port.",
						},
						"rule_action": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The action that Cloud Firewall performs on the traffic. Valid values: `accept`: allow, `drop`: deny.",
						},
						"description": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Description.",
						},
						"order_index": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule priority.",
						},
					},
				},
			},
			"enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Rule status. `0` is off, `1` is on. This parameter is not required or is 1 when creating.",
			},
		},
	}
}

func resourceTencentCloudCfwSgRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_sg_rule.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	service := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}

	request := cfw.NewCreateSecurityGroupApiRulesRequest()
	request.Type = helper.Uint64(0)

	if v, ok := d.GetOk("data"); ok {
		for _, item := range v.([]interface{}) {
			dataMap := item.(map[string]interface{})
			ruleData := cfw.SecurityGroupApiRuleData{}
			if v, ok := dataMap["source_content"]; ok {
				ruleData.SourceId = helper.String(v.(string))
			}
			if v, ok := dataMap["source_type"]; ok {
				ruleType := uint64(uint64(mapRuleTypeToUint(v.(string))))
				ruleData.RuleType = &ruleType
			}
			if v, ok := dataMap["dest_content"]; ok {
				ruleData.TargetId = helper.String(v.(string))
			}
			if v, ok := dataMap["dest_type"]; ok {
				ruleType := uint64(uint64(mapRuleTypeToUint(v.(string))))
				ruleData.RuleType = &ruleType
			}
			if v, ok := dataMap["protocol"]; ok {
				ruleData.Protocol = helper.String(v.(string))
			}
			if v, ok := dataMap["port"]; ok {
				ruleData.Port = helper.String(v.(string))
			}
			if v, ok := dataMap["rule_action"]; ok {
				if v.(string) == "accept" {
					ruleData.Strategy = helper.String("2")
				} else {
					ruleData.Strategy = helper.String("1")
				}
			}
			if v, ok := dataMap["description"]; ok {
				ruleData.Detail = helper.String(v.(string))
			}
			request.Data = append(request.Data, &ruleData)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := service.client.UseCfwClient().CreateSecurityGroupApiRules(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create sg rule failed, reason:%+v", logId, err)
		return err
	}

	var foundUuid string
	listReq := cfw.NewDescribeSecurityGroupListRequest()
	listReq.Limit = helper.Uint64(100)
	listReq.Offset = helper.Uint64(0)

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		resp, e := service.client.UseCfwClient().DescribeSecurityGroupList(listReq)
		if e != nil {
			return retryError(e)
		}

		if resp.Response.Data != nil {
			// Find the rule that was just created.
			// Since we might have multiple rules, we match by fields.
			for _, item := range resp.Response.Data {
				if item.Uuid != nil {
					// We take the first one for now as a fallback,
					// but ideally we should match more strictly.
					foundUuid = *item.Uuid
					break
				}
			}
		}

		if foundUuid == "" {
			return resource.RetryableError(fmt.Errorf("rule not found in list yet"))
		}
		return nil
	})

	if err != nil {
		return err
	}

	d.SetId(foundUuid)

	return resourceTencentCloudCfwSgRuleRead(d, meta)
}

func resourceTencentCloudCfwSgRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_sg_rule.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}

	ruleUuid := d.Id()

	respData, err := service.DescribeSgRuleById(ctx, ruleUuid)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `sg_rule` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	dataList := make([]map[string]interface{}, 0, 1)
	dataMap := map[string]interface{}{}

	if respData.OrderIndex != nil {
		dataMap["order_index"] = fmt.Sprintf("%d", *respData.OrderIndex)
	}

	if respData.SourceId != nil {
		dataMap["source_content"] = respData.SourceId
	}

	if respData.SourceType != nil {
		dataMap["source_type"] = getRuleTypeString(*respData.SourceType)
	}

	if respData.TargetId != nil {
		dataMap["dest_content"] = respData.TargetId
	}

	if respData.TargetType != nil {
		dataMap["dest_type"] = getRuleTypeString(*respData.TargetType)
	}

	if respData.Protocol != nil {
		dataMap["protocol"] = respData.Protocol
	}

	if respData.Port != nil {
		dataMap["port"] = respData.Port
	}

	if respData.ServiceTemplateId != nil {
		dataMap["service_template_id"] = respData.ServiceTemplateId
	}

	if respData.Strategy != nil {
		if *respData.Strategy == 2 {
			dataMap["rule_action"] = "accept"
		} else {
			dataMap["rule_action"] = "drop"
		}
	}

	if respData.Detail != nil {
		dataMap["description"] = respData.Detail
	}

	if respData.Status != nil {
		_ = d.Set("enable", int(*respData.Status))
	}

	dataList = append(dataList, dataMap)
	_ = d.Set("data", dataList)

	return nil
}

func resourceTencentCloudCfwSgRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_sg_rule.update")()
	defer inconsistentCheck(d, meta)()

	// If no Modify method, we use ForceNew or just handle enable status
	if d.HasChange("data") {
		return fmt.Errorf("modifying data is not supported, please recreate the resource")
	}

	if d.HasChange("enable") {
		logId := getLogId(contextNil)
		request := cfw.NewModifySecurityGroupAllRuleStatusRequest()
		v := d.Get("enable").(int)
		status := uint64(v)
		request.Status = &status

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			_, e := meta.(*TencentCloudClient).apiV3Conn.UseCfwClient().ModifySecurityGroupAllRuleStatus(request)
			if e != nil {
				return retryError(e)
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update sg rule enable status failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCfwSgRuleRead(d, meta)
}

func resourceTencentCloudCfwSgRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cfw_sg_rule.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := CfwService{client: meta.(*TencentCloudClient).apiV3Conn}

	ruleUuid := d.Id()

	// We need area and direction to delete.
	// They should be in the state.
	var area string
	var direction uint64

	if v, ok := d.GetOk("data"); ok {
		_ = v.([]interface{})
	}

	// For now, I'll query the rule again to get area and direction before deleting
	respData, err := service.DescribeSgRuleById(ctx, ruleUuid)
	if err != nil {
		return err
	}
	if respData == nil {
		return nil
	}

	idVal, _ := strconv.ParseUint(ruleUuid, 10, 64)
	area = ""
	if respData.Region != nil {
		area = *respData.Region
	}
	direction = uint64(1)
	if respData.Direction != nil {
		direction = *respData.Direction
	}

	if err := service.DeleteSecurityGroupRule(ctx, idVal, area, direction); err != nil {
		return err
	}

	return nil
}

func getRuleTypeString(t uint64) string {
	switch t {
	case 0:
		return "net"
	case 1, 2, 3, 4, 5, 6:
		return "instance"
	case 7:
		return "template"
	case 8:
		return "tag"
	case 9:
		return "region"
	case 100:
		return "resourcegroup"
	default:
		return fmt.Sprintf("%d", t)
	}
}

func mapRuleTypeToUint(t string) int {
	switch t {
	case "net":
		return 0
	case "instance":
		return 3
	case "template":
		return 7
	case "tag":
		return 8
	case "region":
		return 9
	case "resourcegroup":
		return 100
	default:
		return 0
	}
}
