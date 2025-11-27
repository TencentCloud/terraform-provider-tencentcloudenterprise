/*
Provides a resource to create a TurboFS permission group rule.

# Example Usage

```hcl

	resource "tencentcloudenterprise_turbofs_rule" "foo" {
	  p_group_id = "pgroup-7nx89k7l"
	  auth_client_ip  = "10.10.1.0/24"
	  priority        = 1
	  rw_permission   = "ro"
	  user_permission = "root_squash"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_turbofs_rule", CNDescription{
		TerraformTypeCN: "权限规则",
		DescriptionCN:   "提供TurboFS权限规则资源，用于创建和管理TurboFS权限规则。",
		AttributesCN: map[string]string{
			"p_group_id":      "权限组ID",
			"auth_client_ip":  "允许访问的客户端IP地址或地址段",
			"rw_permission":   "读写权限, 可选参数：ro, rw。ro为只读，rw为读写，不填默认为读写",
			"user_permission": "用户权限，可选参数：all_squash，no_all_squash，root_squash，no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为no_root_squash。",
			"rule_id":         "规则ID",
			"priority":        "规则优先级，参数范围1-100。 其中 1 为最高，100为最低",
		},
	})
}
func resourceTencentCloudTurbofsRule() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a TurboFS permission group rule.",
		Create:      resourceTencentCloudTurbofsRuleCreate,
		Read:        resourceTencentCloudTurbofsRuleRead,
		Update:      resourceTencentCloudTurbofsRuleUpdate,
		Delete:      resourceTencentCloudTurbofsRuleDelete,

		Schema: map[string]*schema.Schema{
			"p_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of a permission group.",
			},
			"auth_client_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.",
			},
			"priority": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateIntegerInRange(1, 100),
				Description:  "The priority level of rule. Valid value ranges: (1~100). `1` indicates the highest priority.",
			},
			"rw_permission": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     TURBOFS_RW_PERMISSION_RW,
				Description: "Read and write permissions. Valid values are `ro` and `rw`, and default is `rw`.",
			},
			"user_permission": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     TURBOFS_USER_PERMISSION_NO_ROOT_SQUASH,
				Description: "The permissions of users. Valid values are `all_squash`, `no_all_squash`, `root_squash` and `no_root_squash`. and default is `root_squash`. `all_squash` indicates that all users are mapped as anonymous users or user groups; `no_all_squash` indicates that users will match local users first and be mapped to anonymous users or user groups after matching failed; `root_squash` indicates that map root users to anonymous users or user groups; `no_root_squash` indicates that root users keep root account permission.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of rule",
			},
		},
	}
}

func resourceTencentCloudTurbofsRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_rule.create")()
	logId := getLogId(contextNil)

	var ruleId *string

	request := turbofs.NewCreateCfsRuleRequest()
	if v, ok := d.GetOk("p_group_id"); ok {
		request.PGroupId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("auth_client_ip"); ok {
		request.AuthClientIp = helper.String(v.(string))
	}

	if v, ok := d.GetOk("priority"); ok {
		request.Priority = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("rw_permission"); ok {
		request.RWPermission = helper.String(v.(string))
	}

	if v, ok := d.GetOk("user_permission"); ok {
		request.UserPermission = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, err := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().CreateCfsRule(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			return retryError(err)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response.Response == nil || response.Response.RuleId == nil {
			return resource.NonRetryableError(fmt.Errorf("turbofs rule id is nil"))
		}
		ruleId = response.Response.RuleId
		return nil
	})
	if err != nil {
		return err
	}
	d.SetId(*ruleId)

	return resourceTencentCloudTurbofsRuleRead(d, meta)
}

func resourceTencentCloudTurbofsRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_rule.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	ruleId := d.Id()
	groupId := d.Get("p_group_id").(string)
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var rule *turbofs.PGroupRuleInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		rules, errRet := turbofsService.DescribeCfsRule(ctx, groupId, ruleId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if len(rules) > 0 {
			rule = rules[0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if rule == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("auth_client_ip", rule.AuthClientIp)
	_ = d.Set("user_permission", rule.UserPermission)
	_ = d.Set("priority", rule.Priority)
	if rule.RWPermission != nil {
		_ = d.Set("rw_permission", strings.ToUpper(*rule.RWPermission))
	}

	return nil
}

func resourceTencentCloudTurbofsRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_rule.update")()
	logId := getLogId(contextNil)

	request := turbofs.NewUpdateCfsRuleRequest()
	request.RuleId = helper.String(d.Id())
	request.PGroupId = helper.String(d.Get("p_group_id").(string))
	if d.HasChange("auth_client_ip") {
		request.AuthClientIp = helper.String(d.Get("auth_client_ip").(string))
	}
	if d.HasChange("rw_permission") {
		request.RWPermission = helper.String(d.Get("rw_permission").(string))
	}
	if d.HasChange("user_permission") {
		request.UserPermission = helper.String(d.Get("user_permission").(string))
	}
	if d.HasChange("priority") {
		request.Priority = helper.IntInt64(d.Get("priority").(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, err := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().UpdateCfsRule(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			return retryError(err)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		return nil
	})
	if err != nil {
		return err
	}

	return resourceTencentCloudTurbofsRuleRead(d, meta)
}

func resourceTencentCloudTurbofsRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_rule.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	ruleId := d.Id()
	groupId := d.Get("p_group_id").(string)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		errRet := turbofsService.DeleteAccessRule(ctx, groupId, ruleId)
		if errRet != nil {
			return retryError(errRet)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
