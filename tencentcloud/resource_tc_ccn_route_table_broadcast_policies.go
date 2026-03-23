/*
Provides a resource to manage all broadcast policies of a CCN route table.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ccn_route_table_broadcast_policies" "example" {
	  ccn_id         = "ccn-cwm743gl"
	  route_table_id = "ccnrtb-gbaaugtl"

	  policies {
	    action      = "accept"
	    description = "test broadcast policy"

	    route_conditions {
	      name          = "instance-region"
	      match_pattern = 1
	      values        = ["ap-qingyuan-region-devtest-ops"]
	    }

	    broadcast_conditions {
	      name          = "instance-type"
	      match_pattern = 1
	      values        = ["DIRECTCONNECT"]
	    }
	  }
	}

```

# Import

CCN route table broadcast policies can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ccn_route_table_broadcast_policies.example ccn-cwm743gl#ccnrtb-gbaaugtl
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	ccn "terraform-provider-tencentcloudenterprise/sdk/ccn/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ccn_route_table_broadcast_policies", CNDescription{
		TerraformTypeCN: "云联网路由表广播策略",
		DescriptionCN:   "用于整体管理云联网路由表下的路由广播策略。",
		AttributesCN: map[string]string{
			"ccn_id":               "云联网实例 ID",
			"route_table_id":       "云联网路由表 ID",
			"policies":             "广播策略列表",
			"action":               "策略动作",
			"description":          "策略描述",
			"route_conditions":     "路由匹配条件",
			"broadcast_conditions": "广播目标条件",
			"name":                 "条件名称",
			"values":               "条件值列表",
			"match_pattern":        "匹配模式",
		},
	})
}

func resourceTencentCloudCcnRouteTableBroadcastPolicies() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to manage all broadcast policies of a CCN route table.",
		Create:      resourceTencentCloudCcnRouteTableBroadcastPoliciesCreate,
		Read:        resourceTencentCloudCcnRouteTableBroadcastPoliciesRead,
		Update:      resourceTencentCloudCcnRouteTableBroadcastPoliciesUpdate,
		Delete:      resourceTencentCloudCcnRouteTableBroadcastPoliciesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ccn_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CCN instance ID.",
			},
			"route_table_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CCN route table ID.",
			},
			"policies": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Broadcast policy list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Routing behavior. `accept` allows and `drop` rejects.",
						},
						"description": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Policy description.",
						},
						"route_conditions": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Route matching conditions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Condition type.",
									},
									"values": {
										Type:        schema.TypeList,
										Required:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Condition values.",
									},
									"match_pattern": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Matching mode. `1` for exact match and `0` for fuzzy match.",
									},
								},
							},
						},
						"broadcast_conditions": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Broadcast target conditions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Condition type.",
									},
									"values": {
										Type:        schema.TypeList,
										Required:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Condition values.",
									},
									"match_pattern": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Matching mode. `1` for exact match and `0` for fuzzy match.",
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

func resourceTencentCloudCcnRouteTableBroadcastPoliciesCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table_broadcast_policies.create")()
	defer inconsistentCheck(d, meta)()

	ccnId := d.Get("ccn_id").(string)
	routeTableId := d.Get("route_table_id").(string)

	if err := replaceCcnRouteTableBroadcastPolicies(d, meta, ccnId, routeTableId); err != nil {
		return err
	}

	d.SetId(strings.Join([]string{ccnId, routeTableId}, FILED_SP))
	return resourceTencentCloudCcnRouteTableBroadcastPoliciesRead(d, meta)
}

func resourceTencentCloudCcnRouteTableBroadcastPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table_broadcast_policies.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	items := strings.Split(d.Id(), FILED_SP)
	if len(items) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	ccnId := items[0]
	routeTableId := items[1]

	policySet, err := service.DescribeVpcReplaceCcnRouteTableBroadcastPolicysById(ctx, ccnId, routeTableId)
	if err != nil {
		return err
	}
	if policySet == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `tencentcloudenterprise_ccn_route_table_broadcast_policies` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("ccn_id", ccnId)
	_ = d.Set("route_table_id", routeTableId)
	_ = d.Set("policies", flattenCcnRouteTableBroadcastPolicies(policySet.Policys))
	return nil
}

func resourceTencentCloudCcnRouteTableBroadcastPoliciesUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table_broadcast_policies.update")()
	defer inconsistentCheck(d, meta)()

	items := strings.Split(d.Id(), FILED_SP)
	if len(items) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	if err := replaceCcnRouteTableBroadcastPolicies(d, meta, items[0], items[1]); err != nil {
		return err
	}
	return resourceTencentCloudCcnRouteTableBroadcastPoliciesRead(d, meta)
}

func resourceTencentCloudCcnRouteTableBroadcastPoliciesDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table_broadcast_policies.delete")()
	defer inconsistentCheck(d, meta)()

	items := strings.Split(d.Id(), FILED_SP)
	if len(items) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	logId := getLogId(contextNil)
	request := ccn.NewReplaceCcnRouteTableBroadcastPolicysRequest()
	request.CcnId = helper.String(items[0])
	request.RouteTableId = helper.String(items[1])
	request.Policys = []*ccn.CcnRouteTableBroadcastPolicy{}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().ReplaceCcnRouteTableBroadcastPolicys(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s clear broadcast policies failed, reason:%+v", logId, err)
		return err
	}

	return nil
}

func replaceCcnRouteTableBroadcastPolicies(d *schema.ResourceData, meta interface{}, ccnId, routeTableId string) error {
	logId := getLogId(contextNil)
	request := ccn.NewReplaceCcnRouteTableBroadcastPolicysRequest()
	request.CcnId = helper.String(ccnId)
	request.RouteTableId = helper.String(routeTableId)
	request.Policys = expandCcnRouteTableBroadcastPolicies(d.Get("policies").([]interface{}))

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().ReplaceCcnRouteTableBroadcastPolicys(request)
		if e != nil {
			return retryError(e)
		}
		if result == nil {
			return resource.NonRetryableError(fmt.Errorf("replace ccn route table broadcast policies failed"))
		}

		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s replace broadcast policies failed, reason:%+v", logId, err)
	}
	return err
}

func expandCcnRouteTableBroadcastPolicies(items []interface{}) []*ccn.CcnRouteTableBroadcastPolicy {
	policies := make([]*ccn.CcnRouteTableBroadcastPolicy, 0, len(items))
	for _, item := range items {
		policyMap := item.(map[string]interface{})
		policy := &ccn.CcnRouteTableBroadcastPolicy{}

		if v, ok := policyMap["action"]; ok {
			policy.Action = helper.String(v.(string))
		}
		if v, ok := policyMap["description"]; ok {
			policy.Description = helper.String(v.(string))
		}
		if v, ok := policyMap["route_conditions"]; ok {
			policy.RouteConditions = expandCcnRouteTableBroadcastPolicyConditions(v.([]interface{}))
		}
		if v, ok := policyMap["broadcast_conditions"]; ok {
			policy.BroadcastConditions = expandCcnRouteTableBroadcastPolicyConditions(v.([]interface{}))
		}

		policies = append(policies, policy)
	}
	return policies
}

func flattenCcnRouteTableBroadcastPolicies(items []*ccn.CcnRouteTableBroadcastPolicy) []interface{} {
	policies := make([]interface{}, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}

		policyMap := map[string]interface{}{
			"route_conditions":     flattenCcnRouteTableBroadcastPolicyConditions(item.RouteConditions),
			"broadcast_conditions": flattenCcnRouteTableBroadcastPolicyConditions(item.BroadcastConditions),
		}
		if item.Action != nil {
			policyMap["action"] = *item.Action
		}
		if item.Description != nil {
			policyMap["description"] = *item.Description
		}

		policies = append(policies, policyMap)
	}
	return policies
}

func expandCcnRouteTableBroadcastPolicyConditions(items []interface{}) []*ccn.CcnRouteBroadcastPolicyRouteCondition {
	conditions := make([]*ccn.CcnRouteBroadcastPolicyRouteCondition, 0, len(items))
	for _, item := range items {
		conditionMap := item.(map[string]interface{})
		condition := &ccn.CcnRouteBroadcastPolicyRouteCondition{}

		if v, ok := conditionMap["name"]; ok {
			condition.Name = helper.String(v.(string))
		}
		if v, ok := conditionMap["values"]; ok {
			condition.Values = helper.InterfacesStringsPoint(v.([]interface{}))
		}
		if v, ok := conditionMap["match_pattern"]; ok {
			condition.MatchPattern = helper.IntUint64(v.(int))
		}

		conditions = append(conditions, condition)
	}
	return conditions
}

func flattenCcnRouteTableBroadcastPolicyConditions(items []*ccn.CcnRouteBroadcastPolicyRouteCondition) []interface{} {
	conditions := make([]interface{}, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}

		conditionMap := map[string]interface{}{}
		if item.Name != nil {
			conditionMap["name"] = *item.Name
		}
		if item.Values != nil {
			conditionMap["values"] = helper.StringsInterfaces(item.Values)
		}
		if item.MatchPattern != nil {
			conditionMap["match_pattern"] = int(*item.MatchPattern)
		}

		conditions = append(conditions, conditionMap)
	}
	return conditions
}
