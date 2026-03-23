/*
Provides a resource to create an organization manage policy

# Example Usage

```hcl

	resource "tencentcloudenterprise_organization_org_manage_policy" "example" {
	  name        = "example-policy"
	  type        = "SERVICE_CONTROL_POLICY"
	  description = "example policy description"
	  content     = <<EOF
	{
	    "version": "2.0",
	    "statement": [
	        {
	            "effect": "deny",
	            "action": [
	                "account:*"
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
*/
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strings"

	organization "terraform-provider-tencentcloudenterprise/sdk/organization/v20220508"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudOrganizationOrgManagePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudOrganizationOrgManagePolicyCreate,
		Read:   resourceTencentCloudOrganizationOrgManagePolicyRead,
		Update: resourceTencentCloudOrganizationOrgManagePolicyUpdate,
		Delete: resourceTencentCloudOrganizationOrgManagePolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Policy name.\nThe length is 1~128 characters, which can include Chinese characters, English letters, numbers, and underscores.",
			},

			"content": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Policy content. Refer to the CAM policy syntax.",
			},

			"type": {
				Optional:    true,
				Default:     ServiceControlPolicyType,
				Type:        schema.TypeString,
				Description: "Policy type. Default value is SERVICE_CONTROL_POLICY.\nValid values:\n  - `SERVICE_CONTROL_POLICY`: Service control policy.\n  - `TAG_POLICY`: Tag policy.",
			},

			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Policy description.",
			},

			"policy_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Policy Id.",
			},
		},
	}
}

func resourceTencentCloudOrganizationOrgManagePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_org_manage_policy.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		policyType string
		request    = organization.NewCreatePolicyRequest()
		response   = organization.NewCreatePolicyResponse()
	)
	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("content"); ok {
		request.Content = helper.String(v.(string))
	}

	if v, ok := d.GetOk("type"); ok {
		policyType = v.(string)
		request.Type = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOrganizationClient().CreatePolicy(request)
		if e != nil {
			return resource.NonRetryableError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create organization OrgManagePolicy failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{helper.UInt64ToStr(*response.Response.PolicyId), policyType}, FILED_SP))
	return resourceTencentCloudOrganizationOrgManagePolicyRead(d, meta)
}

func resourceTencentCloudOrganizationOrgManagePolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_org_manage_policy.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := OrganizationService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	policyId := idSplit[0]
	policyType := idSplit[1]

	OrgManagePolicy, err := service.DescribeOrganizationOrgManagePolicyById(ctx, policyId, policyType)
	if err != nil {
		return err
	}

	if OrgManagePolicy == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `OrganizationOrgManagePolicy` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if OrgManagePolicy.PolicyName != nil {
		_ = d.Set("name", OrgManagePolicy.PolicyName)
	}

	if OrgManagePolicy.PolicyDocument != nil {
		_ = d.Set("content", OrgManagePolicy.PolicyDocument)
	}

	if OrgManagePolicy.Type != nil {
		_ = d.Set("type", policyType)
	}

	if OrgManagePolicy.Description != nil {
		_ = d.Set("description", OrgManagePolicy.Description)
	}
	_ = d.Set("policy_id", policyId)

	return nil
}

func resourceTencentCloudOrganizationOrgManagePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_org_manage_policy.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := organization.NewUpdatePolicyRequest()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	policyId := idSplit[0]

	request.PolicyId = helper.StrToInt64Point(policyId)

	needChange := false
	mutableArgs := []string{"name", "content", "type", "description"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		if v, ok := d.GetOk("name"); ok {
			request.Name = helper.String(v.(string))
		}
		if v, ok := d.GetOk("content"); ok {
			request.Content = helper.String(v.(string))
		}
		if v, ok := d.GetOk("type"); ok {
			request.Type = helper.String(v.(string))
		}
		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseOrganizationClient().UpdatePolicy(request)
			if e != nil {
				return resource.NonRetryableError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update organization OrgManagePolicy failed, reason:%+v", logId, err)
			return err
		}

	}
	return resourceTencentCloudOrganizationOrgManagePolicyRead(d, meta)
}

func resourceTencentCloudOrganizationOrgManagePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_org_manage_policy.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := OrganizationService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	policyId := idSplit[0]
	policyType := idSplit[1]

	if err := service.DeleteOrganizationOrgManagePolicyById(ctx, policyId, policyType); err != nil {
		return err
	}

	return nil
}
