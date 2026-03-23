/*
Provides a resource to create a CAM user policy attachment.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_user_policy_attachment" "foo" {
	  user_id   = tencentcloudenterprise_cam_user.foo.id
	  policy_id = tencentcloudenterprise_cam_policy.foo.id
	}

```

# Import

CAM user policy attachment can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_user_policy_attachment.foo cam-test#26800353
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_user_policy_attachment", CNDescription{
		TerraformTypeCN: "CAM用户策略关联",
		DescriptionCN:   "提供一个资源来将策略关联到 CAM 子用户。",
		AttributesCN: map[string]string{
			"user_id":     "关联的 CAM 子用户 ID",
			"policy_id":   "策略 ID",
			"create_mode": "创建模式",
			"policy_type": "策略类型",
			"create_time": "创建时间",
			"policy_name": "策略名称",
		},
	})
}

func resourceTencentCloudCamUserPolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamUserPolicyAttachmentCreate,
		Read:   resourceTencentCloudCamUserPolicyAttachmentRead,
		Delete: resourceTencentCloudCamUserPolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				// Import ID format: <user_id>#<policy_id>
				parts := strings.Split(d.Id(), "#")
				if len(parts) != 2 {
					return nil, fmt.Errorf("invalid import id format, expected <user_id>#<policy_id>")
				}
				userPart := parts[0]
				policyPart := parts[1]
				if err := d.Set("user_id", userPart); err != nil {
					return nil, err
				}
				if err := d.Set("policy_id", policyPart); err != nil {
					return nil, err
				}
				d.SetId(fmt.Sprintf("%s#%s", userPart, policyPart))
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the attached CAM user.",
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
				Description: "Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.",
			},
			"policy_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time of the CAM user policy attachment.",
			},
			"policy_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the policy.",
			},
		},
	}
}

func resourceTencentCloudCamUserPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user_policy_attachment.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	userId, _, err := getUserId(d)
	if err != nil {
		return err
	}
	policyId := d.Get("policy_id").(string)
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		e := camService.AddUserPolicyAttachment(ctx, userId, policyId)
		if e != nil {
			log.Printf("[CRITAL]%s reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create CAM user policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	d.SetId(userId + "#" + policyId)

	//get really instance then read

	userPolicyAttachmentId := d.Id()
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instance, e := camService.DescribeUserPolicyAttachmentById(ctx, userPolicyAttachmentId)
		if e != nil {
			return retryError(e)
		}
		if instance == nil {
			return resource.RetryableError(fmt.Errorf("creation not done"))
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM user policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}
	time.Sleep(10 * time.Second)
	return resourceTencentCloudCamUserPolicyAttachmentRead(d, meta)
}

func resourceTencentCloudCamUserPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user_policy_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	userPolicyAttachmentId := d.Id()

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	// Decode the ID to get user and policy IDs
	userId, policyId, e := camService.decodeCamPolicyAttachmentId(userPolicyAttachmentId)
	if e != nil {
		return e
	}

	// Use ListEntitiesForPolicy to check if the attachment exists
	var instance *cam.AttachPolicyInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := camService.DescribeUserPolicyAttachmentById(ctx, userPolicyAttachmentId)
		if e != nil {
			return retryError(e)
		}
		instance = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM user policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	if instance == nil {
		d.SetId("")
		return nil
	}

	// Set basic fields
	_ = d.Set("user_id", userId)
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

	// Set policy_type (empty for now, can be enhanced if needed)
	_ = d.Set("policy_type", "")

	return nil
}

func resourceTencentCloudCamUserPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user_policy_attachment.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	userPolicyAttachmentId := d.Id()

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		e := camService.DeleteUserPolicyAttachmentById(ctx, userPolicyAttachmentId)
		if e != nil {
			log.Printf("[CRITAL]%s reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete CAM user policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}

func getUserId(d *schema.ResourceData) (value string, usingName bool, err error) {
	// Check user_name first (preferred)
	if name, hasName := d.GetOk("user_name"); hasName {
		return name.(string), true, nil
	}
	// Fall back to user_id (deprecated)
	if id, hasId := d.GetOk("user_id"); hasId {
		return id.(string), false, nil
	}
	return "", false, fmt.Errorf("no user_id or user_name provided")
}
