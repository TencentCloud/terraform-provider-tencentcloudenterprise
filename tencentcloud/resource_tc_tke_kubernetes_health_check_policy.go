/*
Provide a resource to create a TKE health check policy.

# Example Usage

```hcl
resource "tencentcloudenterprise_tke_kubernetes_health_check_policy" "example" {
  cluster_id = "cls-xxxxxxxx"
  name       = "health-policy-example"

  rules {
    name                = "OOMKilling"
    enabled             = true
    auto_repair_enabled = true
  }

  rules {
    name                = "KubeletUnhealthy"
    enabled             = true
    auto_repair_enabled = false
  }
}
```

# Import

TKE health check policy can be imported using cluster_id#policy_name, e.g.
```
$ terraform import tencentcloudenterprise_tke_kubernetes_health_check_policy.example cls-xxxxxxxx#health-policy-example
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	tke2 "terraform-provider-tencentcloudenterprise/sdk/tke/v20220501"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_health_check_policy", CNDescription{
		TerraformTypeCN: "集群健康检查策略",
		DescriptionCN:   "用于管理 TKE 集群的健康检查策略。",
		AttributesCN: map[string]string{
			"cluster_id":          "集群 ID。",
			"name":                "健康检查策略名称。",
			"rules":               "健康检查规则列表。",
			"enabled":             "是否启用该检查项。",
			"auto_repair_enabled": "是否开启自动修复。",
		},
	})
}

func resourceTencentCloudKubernetesHealthCheckPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a TKE health check policy.",
		Create:      resourceTencentCloudKubernetesHealthCheckPolicyCreate,
		Read:        resourceTencentCloudKubernetesHealthCheckPolicyRead,
		Update:      resourceTencentCloudKubernetesHealthCheckPolicyUpdate,
		Delete:      resourceTencentCloudKubernetesHealthCheckPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the cluster.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Health check policy name.",
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Health check policy rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Health check rule name.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to enable this check item.",
						},
						"auto_repair_enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to enable auto repair.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudKubernetesHealthCheckPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_health_check_policy.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	_ = context.WithValue(context.TODO(), logIdKey, logId)

	client := meta.(*TencentCloudClient).apiV3Conn.UseTke2Client()
	request := tke2.NewCreateHealthCheckPolicyRequest()

	clusterId := d.Get("cluster_id").(string)
	request.ClusterId = helper.String(clusterId)

	name := d.Get("name").(string)
	healthCheckPolicy := tke2.HealthCheckPolicy{}
	healthCheckPolicy.Name = helper.String(name)

	if v, ok := d.GetOk("rules"); ok {
		rulesList := v.([]interface{})
		for _, item := range rulesList {
			ruleMap := item.(map[string]interface{})
			rule := tke2.HealthCheckPolicyRule{}
			if v, ok := ruleMap["name"]; ok {
				rule.Name = helper.String(v.(string))
			}
			if v, ok := ruleMap["enabled"]; ok {
				rule.Enabled = helper.Bool(v.(bool))
			}
			if v, ok := ruleMap["auto_repair_enabled"]; ok {
				rule.AutoRepairEnabled = helper.Bool(v.(bool))
			}
			healthCheckPolicy.Rules = append(healthCheckPolicy.Rules, &rule)
		}
	}

	request.HealthCheckPolicy = &healthCheckPolicy

	log.Printf("[DEBUG]%s api[CreateHealthCheckPolicy] request: %s", logId, request.ToJsonString())

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := client.CreateHealthCheckPolicy(request)
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create health check policy failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(clusterId + FILED_SP + name)

	return resourceTencentCloudKubernetesHealthCheckPolicyRead(d, meta)
}

func resourceTencentCloudKubernetesHealthCheckPolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_health_check_policy.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	_ = context.WithValue(context.TODO(), logIdKey, logId)

	client := meta.(*TencentCloudClient).apiV3Conn.UseTke2Client()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	clusterId := idSplit[0]
	policyName := idSplit[1]

	request := tke2.NewDescribeHealthCheckPoliciesRequest()
	request.ClusterId = helper.String(clusterId)
	request.Filters = []*tke2.Filter{
		{
			Name:   helper.String("HealthCheckPolicyName"),
			Values: []*string{helper.String(policyName)},
		},
	}

	var policy *tke2.HealthCheckPolicy
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		response, e := client.DescribeHealthCheckPolicies(request)
		if e != nil {
			return retryError(e)
		}
		if response == nil || response.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("DescribeHealthCheckPolicies response is nil"))
		}
		if len(response.Response.HealthCheckPolicies) == 0 {
			return nil
		}
		policy = response.Response.HealthCheckPolicies[0]
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read health check policy failed, reason:%+v", logId, err)
		return err
	}

	if policy == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("cluster_id", clusterId)
	_ = d.Set("name", policyName)

	if policy.Rules != nil {
		rules := make([]map[string]interface{}, 0, len(policy.Rules))
		for _, rule := range policy.Rules {
			ruleMap := map[string]interface{}{
				"name":                "",
				"enabled":             false,
				"auto_repair_enabled": false,
			}
			if rule.Name != nil {
				ruleMap["name"] = *rule.Name
			}
			if rule.Enabled != nil {
				ruleMap["enabled"] = *rule.Enabled
			}
			if rule.AutoRepairEnabled != nil {
				ruleMap["auto_repair_enabled"] = *rule.AutoRepairEnabled
			}
			rules = append(rules, ruleMap)
		}
		_ = d.Set("rules", rules)
	}

	return nil
}

func resourceTencentCloudKubernetesHealthCheckPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_health_check_policy.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	_ = context.WithValue(context.TODO(), logIdKey, logId)

	client := meta.(*TencentCloudClient).apiV3Conn.UseTke2Client()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	clusterId := idSplit[0]
	policyName := idSplit[1]

	if d.HasChange("rules") {
		request := tke2.NewModifyHealthCheckPolicyRequest()
		request.ClusterId = helper.String(clusterId)

		healthCheckPolicy := tke2.HealthCheckPolicy{}
		healthCheckPolicy.Name = helper.String(policyName)

		if v, ok := d.GetOk("rules"); ok {
			rulesList := v.([]interface{})
			for _, item := range rulesList {
				ruleMap := item.(map[string]interface{})
				rule := tke2.HealthCheckPolicyRule{}
				if v, ok := ruleMap["name"]; ok {
					rule.Name = helper.String(v.(string))
				}
				if v, ok := ruleMap["enabled"]; ok {
					rule.Enabled = helper.Bool(v.(bool))
				}
				if v, ok := ruleMap["auto_repair_enabled"]; ok {
					rule.AutoRepairEnabled = helper.Bool(v.(bool))
				}
				healthCheckPolicy.Rules = append(healthCheckPolicy.Rules, &rule)
			}
		}

		request.HealthCheckPolicy = &healthCheckPolicy

		log.Printf("[DEBUG]%s api[ModifyHealthCheckPolicy] request: %s", logId, request.ToJsonString())

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			_, e := client.ModifyHealthCheckPolicy(request)
			if e != nil {
				return retryError(e)
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update health check policy failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudKubernetesHealthCheckPolicyRead(d, meta)
}

func resourceTencentCloudKubernetesHealthCheckPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_health_check_policy.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	_ = context.WithValue(context.TODO(), logIdKey, logId)

	client := meta.(*TencentCloudClient).apiV3Conn.UseTke2Client()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	clusterId := idSplit[0]
	policyName := idSplit[1]

	request := tke2.NewDeleteHealthCheckPolicyRequest()
	request.ClusterId = helper.String(clusterId)
	request.HealthCheckPolicyName = helper.String(policyName)

	log.Printf("[DEBUG]%s api[DeleteHealthCheckPolicy] request: %s", logId, request.ToJsonString())

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := client.DeleteHealthCheckPolicy(request)
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete health check policy failed, reason:%+v", logId, err)
		return err
	}

	return nil
}
