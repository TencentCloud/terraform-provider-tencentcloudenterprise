/*
Provides a resource to create an organization cic_role_configuration_permission_custom_policy_attachment

Example Usage

```hcl
resource "tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment" "cic_role_configuration_permission_custom_policy_attachment" {
    zone_id = "z-xxxxxx"
    role_configuration_id = "rc-xxxxxx"
    role_policy_name = "CustomPolicy"
	role_policy_document = <<-EOF
{
    "version": "2.0",
    "statement": [
        {
            "effect": "allow",
            "action": [
                "vpc:AcceptAttachCcnInstances"
            ],
            "resource": [
                "*"
            ]
        }
    ]
}
EOF
}
```

Import

organization cic_role_configuration_permission_custom_policy_attachment can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment.cic_role_configuration_permission_custom_policy_attachment ${zoneId}#${roleConfigurationId}#${rolePolicyName}
```

 */
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strings"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment", CNDescription{
		TerraformTypeCN: "身份中心角色配置权限自定义策略关联",
		DescriptionCN:   "提供身份中心角色配置权限自定义策略关联资源，用于将权限策略关联到角色配置。",
		AttributesCN: map[string]string{
			"zone_id":               "空间ID",
			"role_configuration_id": "角色配置ID",
			"role_policy_name":      "角色策略名称",
			"role_policy_document":  "角色策略文档",
			"role_policy_type":      "角色策略类型",
			"add_time":              "角色策略添加时间",
		},
	})
}

func resourceTencentCloudCicRoleConfigurationPermissionCustomPolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCicRoleConfigurationPermissionCustomPolicyAttachmentCreate,
		Read:   resourceTencentCloudCicRoleConfigurationPermissionCustomPolicyAttachmentRead,
		Delete: resourceTencentCloudCicRoleConfigurationPermissionCustomPolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Description: "Provide cic role configuration permission custom policy associated resources to associate permission policies with role configurations.",
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

			"role_policy_document": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Role policy document.",
			},

			"role_policy_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Role policy name.",
			},

			"role_policy_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role policy type.",
			},

			"add_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role policy add time.",
			},
		},
	}
}

func resourceTencentCloudCicRoleConfigurationPermissionCustomPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	var (
		zoneId              string
		roleConfigurationId string
		rolePolicyName      string
	)
	var (
		request  = cic.NewAddPermissionPolicyToRoleConfigurationRequest()
		response = cic.NewAddPermissionPolicyToRoleConfigurationResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
		request.ZoneId = helper.String(zoneId)
	}

	if v, ok := d.GetOk("role_configuration_id"); ok {
		roleConfigurationId = v.(string)
		request.RoleConfigurationId = helper.String(roleConfigurationId)
	}

	if v, ok := d.GetOk("role_policy_name"); ok {
		rolePolicyName = v.(string)
		request.RolePolicyNames = []*string{helper.String(rolePolicyName)}
	}

	if v, ok := d.GetOk("role_policy_document"); ok {
		request.CustomPolicyDocument = helper.String(v.(string))
	}

	request.RolePolicyType = helper.String("Custom")

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
		log.Printf("[CRITAL]%s create identity center role configuration permission policy attachment failed, reason:%+v",
			logId, err)
		return err
	}

	_ = response

	d.SetId(strings.Join([]string{zoneId, roleConfigurationId, rolePolicyName}, FILED_SP))

	return resourceTencentCloudCicRoleConfigurationPermissionCustomPolicyAttachmentRead(d, meta)
}

func resourceTencentCloudCicRoleConfigurationPermissionCustomPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment.read")()
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
	rolePolicyName := idSplit[2]

	_ = d.Set("zone_id", zoneId)

	_ = d.Set("role_configuration_id", roleConfigurationId)

	_ = d.Set("role_policy_name", rolePolicyName)

	respData, err := service.DescribeCicRoleConfigurationPermissionPolicyAttachmentById(
		ctx, zoneId, roleConfigurationId, "Custom")
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_role_configuration_permission_policy_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if respData.RolePolicies != nil {
		var rolePolicie *cic.RolePolicie
		for _, r := range respData.RolePolicies {
			if *r.RolePolicyName == rolePolicyName {
				rolePolicie = r
				break
			}
		}

		if rolePolicie == nil {
			log.Printf("[WARN]%s resource `cic_role_configuration_permission_policy_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
			return fmt.Errorf("RolePolicy %s is not exist", d.Id())
		}

		if rolePolicie.RolePolicyName != nil {
			_ = d.Set("role_policy_name", rolePolicie.RolePolicyName)
		}

		if rolePolicie.RolePolicyType != nil {
			_ = d.Set("role_policy_type", rolePolicie.RolePolicyType)
		}

		if rolePolicie.RolePolicyDocument != nil {
			_ = d.Set("role_policy_document", rolePolicie.RolePolicyDocument)
		}

		if rolePolicie.AddTime != nil {
			_ = d.Set("add_time", rolePolicie.AddTime)
		}

	}

	return nil
}

func resourceTencentCloudCicRoleConfigurationPermissionCustomPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	roleConfigurationId := idSplit[1]
	rolePolicyName := idSplit[2]

	var (
		request  = cic.NewRemovePermissionPolicyFromRoleConfigurationRequest()
		response = cic.NewRemovePermissionPolicyFromRoleConfigurationResponse()
	)

	request.ZoneId = helper.String(zoneId)

	request.RoleConfigurationId = helper.String(roleConfigurationId)

	request.RolePolicyType = helper.String("Custom")

	request.RolePolicyId = helper.Int64(0)

	request.RolePolicyName = helper.String(rolePolicyName)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
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
		log.Printf("[CRITAL]%s delete identity center role configuration permission policy attachment failed, reason:%+v",
			logId, err)
		return err
	}

	_ = response
	return nil
}