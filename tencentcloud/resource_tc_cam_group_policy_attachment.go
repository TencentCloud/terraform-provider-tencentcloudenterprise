/*
Provides a resource to create a CAM group policy attachment.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_group_policy_attachment" "foo" {
	  group_id  = tencentcloudenterprise_cam_group.foo.id
	  policy_id = tencentcloudenterprise_cam_policy.foo.id
	}

```

# Import

CAM group policy attachment can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_group_policy_attachment.foo 12515263#26800353
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
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_group_policy_attachment", CNDescription{
		TerraformTypeCN: "CAM用户组策略关联",
		DescriptionCN:   "提供一个资源来将策略关联到 CAM 用户组。",
		AttributesCN: map[string]string{
			"group_id":    "关联的 CAM 用户组 ID",
			"policy_id":   "策略 ID",
			"create_mode": "创建模式",
			"policy_type": "策略类型",
			"create_time": "创建时间",
			"policy_name": "策略名称",
		},
	})
}

func resourceTencentCloudCamGroupPolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamGroupPolicyAttachmentCreate,
		Read:   resourceTencentCloudCamGroupPolicyAttachmentRead,
		Delete: resourceTencentCloudCamGroupPolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the attached CAM group.",
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
				Description: "Mode of Creation of the CAM group policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.",
			},
			"policy_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the policy strategy. 'Group' means customer strategy and 'QCS' means preset strategy.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time of the CAM group policy attachment.",
			},
			"policy_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the policy.",
			},
		},
	}
}

func resourceTencentCloudCamGroupPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group_policy_attachment.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	groupId := d.Get("group_id").(string)
	policyId := d.Get("policy_id").(string)
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		e := camService.AddGroupPolicyAttachment(ctx, groupId, policyId)
		if e != nil {
			log.Printf("[CRITAL]%s reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create CAM group policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	d.SetId(groupId + "#" + policyId)

	//get really instance then read
	groupPolicyAttachmentId := d.Id()
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instance, e := camService.DescribeGroupPolicyAttachmentById(ctx, groupPolicyAttachmentId)
		if e != nil {
			return retryError(e)
		}
		if instance == nil {
			return resource.RetryableError(fmt.Errorf("creation not done"))
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM group policy failed, reason:%s\n", logId, err.Error())
		return err
	}
	time.Sleep(10 * time.Second)
	return resourceTencentCloudCamGroupPolicyAttachmentRead(d, meta)
}

func resourceTencentCloudCamGroupPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group_policy_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	groupPolicyAttachmentId := d.Id()

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	// Decode the ID to get group and policy IDs
	groupId, policyId, e := camService.decodeCamPolicyAttachmentId(groupPolicyAttachmentId)
	if e != nil {
		return e
	}

	// Use ListEntitiesForPolicy to check if the attachment exists
	var instance *cam.AttachPolicyInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := camService.DescribeGroupPolicyAttachmentById(ctx, groupPolicyAttachmentId)
		if e != nil {
			return retryError(e)
		}
		instance = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM group policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	if instance == nil {
		d.SetId("")
		return nil
	}

	// Set basic fields
	_ = d.Set("group_id", groupId)
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

func resourceTencentCloudCamGroupPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_group_policy_attachment.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	groupPolicyAttachmentId := d.Id()

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		e := camService.DeleteGroupPolicyAttachmentById(ctx, groupPolicyAttachmentId)
		if e != nil {
			log.Printf("[CRITAL]%s reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete CAM group policy attachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}
