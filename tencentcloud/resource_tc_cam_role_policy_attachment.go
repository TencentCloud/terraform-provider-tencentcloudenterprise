/*
Provides a resource to create a CAM role policy attachment.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_role_policy_attachment" "foo" {
	  role_id   = tencentcloudenterprise_cam_role.foo.id
	  policy_id = tencentcloudenterprise_cam_policy.foo.id
	}

```

# Import

CAM role policy attachment can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_role_policy_attachment.foo 4611686018427922725#26800353
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_role_policy_attachment", CNDescription{
		TerraformTypeCN: "CAM角色策略关联",
		DescriptionCN:   "提供一个资源来将策略关联到 CAM 角色。",
		AttributesCN: map[string]string{
			"role_id":     "关联的 CAM 角色 ID",
			"policy_id":   "策略 ID",
			"create_mode": "创建模式",
			"policy_type": "策略类型",
			"create_time": "创建时间",
			"policy_name": "策略名称",
		},
	})
}

func resourceTencentCloudCamRolePolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamRolePolicyAttachmentCreate,
		Read:   resourceTencentCloudCamRolePolicyAttachmentRead,
		Delete: resourceTencentCloudCamRolePolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"role_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the attached CAM role.",
			},
			"policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the policy.",
			},
			"create_mode": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.",
			},
			"policy_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the CAM role policy attachment.",
			},
			"policy_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the policy.",
			},
		},
	}
}

func resourceTencentCloudCamRolePolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_role_policy_attachment.create")()

	logId := getLogId(contextNil)

	roleIdStr := d.Get("role_id").(string)
	roleIdUint, e := strconv.ParseUint(roleIdStr, 10, 64)
	if e != nil {
		return e
	}
	policyIdStr := d.Get("policy_id").(string)
	policyIdUint, e := strconv.ParseUint(policyIdStr, 10, 64)
	if e != nil {
		return e
	}

	request := cam.NewAttachRolePoliciesRequest()
	request.RoleId = &roleIdUint
	request.PolicyId = []*uint64{&policyIdUint}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().AttachRolePolicies(request)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), e.Error())
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create CAM role policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	d.SetId(roleIdStr + "#" + policyIdStr)

	//get really instance then read
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	rolePolicyAttachmentId := d.Id()
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instance, e := camService.DescribeRolePolicyAttachmentById(ctx, rolePolicyAttachmentId)
		if e != nil {
			return retryError(e)
		}
		if instance == nil {
			return resource.RetryableError(fmt.Errorf("creation not done"))
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM role policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}
	time.Sleep(10 * time.Second)
	return resourceTencentCloudCamRolePolicyAttachmentRead(d, meta)
}

func resourceTencentCloudCamRolePolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_role_policy_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	rolePolicyAttachmentId := d.Id()

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	// Decode the ID to get role and policy IDs
	roleId, policyId, e := camService.decodeCamPolicyAttachmentId(rolePolicyAttachmentId)
	if e != nil {
		return e
	}

	// Use ListEntitiesForPolicy to check if the attachment exists
	var instance *cam.AttachedPolicyOfRole
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := camService.DescribeRolePolicyAttachmentById(ctx, rolePolicyAttachmentId)
		if e != nil {
			return retryError(e)
		}
		instance = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM role policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	if instance == nil {
		d.SetId("")
		return nil
	}

	// Set basic fields
	_ = d.Set("role_id", roleId)
	_ = d.Set("policy_id", strconv.Itoa(int(policyId)))

	// Set policy info from the instance
	if instance.PolicyName != nil {
		_ = d.Set("policy_name", *instance.PolicyName)
	}
	if instance.AddTime != nil {
		_ = d.Set("create_time", *instance.AddTime)
	}
	if instance.CreateMode != nil {
		_ = d.Set("create_mode", int(*instance.CreateMode))
	}
	if instance.PolicyType != nil {
		_ = d.Set("policy_type", *instance.PolicyType)
	} else {
		_ = d.Set("policy_type", "")
	}

	return nil
}

func resourceTencentCloudCamRolePolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_role_policy_attachment.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	rolePolicyAttachmentId := d.Id()

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		e := camService.DeleteRolePolicyAttachmentById(ctx, rolePolicyAttachmentId)
		if e != nil {
			log.Printf("[CRITAL]%s reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete CAM role policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}
