/*
Provides a resource to create an organization cic_role_configuration_permission_policy_attachment

Example Usage

```hcl
resource "tencentcloudenterprise_cic_role_configuration_permission_policy_attachment" "cic_role_configuration_permission_policy_attachment" {
    zone_id = "z-xxxxxx"
    role_configuration_id = "rc-xxxxxx"
    role_policy_id = xxxxxx
}
```

Import

organization cic_role_configuration_permission_policy_attachment can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_role_configuration_permission_policy_attachment.cic_role_configuration_permission_policy_attachment ${zoneId}#${roleConfigurationId}#${rolePolicyIdString}
```

 */
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strconv"
	"strings"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_role_configuration_permission_policy_attachment", CNDescription{
		TerraformTypeCN: "身份中心角色配置权限策略关联",
		DescriptionCN:   "提供身份中心角色配置权限策略关联资源，用于将权限策略关联到角色配置。",
		AttributesCN: map[string]string{
			"zone_id":               "空间ID",
			"role_configuration_id": "角色配置ID",
			"role_policy_id":        "角色策略ID",
			"role_policy_name":      "角色策略名称",
			"role_policy_type":      "角色策略类型",
		},
	})
}

func resourceTencentCloudCicRoleConfigurationPermissionPolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Description: "Provide identity center role configuration, permission policy and associated resources for linking permission policies to role configurations.",
		Create: resourceTencentCloudCicRoleConfigurationPermissionPolicyAttachmentCreate,
		Read:   resourceTencentCloudCicRoleConfigurationPermissionPolicyAttachmentRead,
		Delete: resourceTencentCloudCicRoleConfigurationPermissionPolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Space ID.",
			},

			"role_configuration_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Permission configuration ID.",
			},

			"role_policy_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Role policy id.",
			},

			"role_policy_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Role policy name.",
			},

			"role_policy_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role policy type.",
			},

			"role_policy_document": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role policy document.",
			},

			"add_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role policy add time.",
			},
		},
	}
}

func resourceTencentCloudCicRoleConfigurationPermissionPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration_permission_policy_attachment.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	var (
		zoneId              string
		roleConfigurationId string
		rolePolicyId        int64
	)
	var (
		request  = cic.NewAddPermissionPolicyToRoleConfigurationRequest()
		response = cic.NewAddPermissionPolicyToRoleConfigurationResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}
	if v, ok := d.GetOk("role_configuration_id"); ok {
		roleConfigurationId = v.(string)
	}
	if v, ok := d.GetOk("role_policy_id"); ok {
		rolePolicyId = int64(v.(int))
	}

	if v, ok := d.GetOk("zone_id"); ok {
		request.ZoneId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("role_configuration_id"); ok {
		request.RoleConfigurationId = helper.String(v.(string))
	}

	request.RolePolicyType = helper.String("System")

	if v, ok := d.GetOk("role_policy_id"); ok {
		policyDetail := &cic.PolicyDetail{
			PolicyId: helper.IntInt64(v.(int)),
		}
		if vv, innerOk := d.GetOk("role_policy_name"); innerOk {
			policyDetail.PolicyName = helper.String(vv.(string))
		}
		request.RolePolicies = []*cic.PolicyDetail{
			policyDetail,
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().AddPermissionPolicyToRoleConfiguration(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create identity center role configuration permission policy attachment failed, reason:%+v", logId, err)
		return err
	}

	_ = response

	rolePolicyIdString := strconv.FormatInt(rolePolicyId, 10)
	d.SetId(strings.Join([]string{zoneId, roleConfigurationId, rolePolicyIdString}, FILED_SP))

	return resourceTencentCloudCicRoleConfigurationPermissionPolicyAttachmentRead(d, meta)
}

func resourceTencentCloudCicRoleConfigurationPermissionPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration_permission_policy_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	roleConfigurationId := idSplit[1]
	rolePolicyIdString := idSplit[2]
	rolePolicyId, err := strconv.ParseInt(rolePolicyIdString, 10, 64)
	if err != nil {
		return err
	}

	_ = d.Set("zone_id", zoneId)

	_ = d.Set("role_configuration_id", roleConfigurationId)

	_ = d.Set("role_policy_id", rolePolicyId)

	respData, err := service.DescribeCicRoleConfigurationPermissionPolicyAttachmentById(
		ctx, zoneId, roleConfigurationId, "System")
	if err != nil {
		return err
	}

	var rolePolicy *cic.RolePolicie
	if respData != nil {
		for _, candidate := range respData.RolePolicies {
			if candidate != nil && candidate.RolePolicyId != nil && *candidate.RolePolicyId == rolePolicyId {
				rolePolicy = candidate
				break
			}
		}
	}
	if rolePolicy == nil {
		attachmentID := d.Id()
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_role_configuration_permission_policy_attachment` [%s] not found, please check if it has been deleted.\n", logId, attachmentID)
		return nil
	}

	if rolePolicy.RolePolicyName != nil {
		_ = d.Set("role_policy_name", rolePolicy.RolePolicyName)
	}

	if rolePolicy.RolePolicyType != nil {
		_ = d.Set("role_policy_type", rolePolicy.RolePolicyType)
	}

	if rolePolicy.RolePolicyDocument != nil {
		_ = d.Set("role_policy_document", rolePolicy.RolePolicyDocument)
	}

	if rolePolicy.AddTime != nil {
		_ = d.Set("add_time", rolePolicy.AddTime)
	}

	return nil
}

func resourceTencentCloudCicRoleConfigurationPermissionPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration_permission_policy_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	roleConfigurationId := idSplit[1]
	rolePolicyIdString := idSplit[2]
	rolePolicyId, err := strconv.ParseInt(rolePolicyIdString, 10, 64)
	if err != nil {
		return err
	}

	var (
		request  = cic.NewRemovePermissionPolicyFromRoleConfigurationRequest()
		response = cic.NewRemovePermissionPolicyFromRoleConfigurationResponse()
	)

	request.ZoneId = helper.String(zoneId)

	request.RoleConfigurationId = helper.String(roleConfigurationId)

	request.RolePolicyType = helper.String("System")

	request.RolePolicyId = helper.Int64(rolePolicyId)

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().RemovePermissionPolicyFromRoleConfiguration(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete identity center role configuration permission policy attachment failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	return nil
}
