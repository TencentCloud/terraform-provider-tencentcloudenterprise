/*
Provides a resource to create a cvm security_group_attachment

# Example Usage

```hcl

	resource "tencentcloudenterprise_cvm_security_group_attachment" "security_group_attachment" {
	  security_group_id = "sg-xxxxxxx"
	  instance_id = "ins-xxxxxxxx"
	}

```

# Import

cvm security_group_attachment can be imported using the id, e.g.

```
terraform import cloud_cvm_security_group_attachment.security_group_attachment ${instance_id}#${security_group_id}
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
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cvm_security_group_attachment", CNDescription{
		TerraformTypeCN: "云服务器安全组绑定",
		DescriptionCN:   "提供云服务器安全组绑定资源，用于将安全组绑定到云服务器实例。",
		AttributesCN: map[string]string{
			"security_group_id": "安全组ID",
			"instance_id":       "实例ID",
		},
	})
}
func resourceTencentCloudCvmSecurityGroupAttachment() *schema.Resource {
	return &schema.Resource{
		Description: "Create cvm security_group_attachment",
		Create:      resourceTencentCloudCvmSecurityGroupAttachmentCreate,
		Read:        resourceTencentCloudCvmSecurityGroupAttachmentRead,
		Delete:      resourceTencentCloudCvmSecurityGroupAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"security_group_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Security group id.",
			},

			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance id.",
			},
		},
	}
}

func resourceTencentCloudCvmSecurityGroupAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_security_group_attachment.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cvm.NewAssociateSecurityGroupsRequest()
	securityGroupId := d.Get("security_group_id").(string)
	instanceId := d.Get("instance_id").(string)
	request.SecurityGroupIds = []*string{}

	request.SecurityGroupIds = []*string{&securityGroupId}
	request.InstanceIds = []*string{&instanceId}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().AssociateSecurityGroups(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cvm securityGroupAttachment failed, reason:%+v", logId, err)
		return err
	}
	d.SetId(instanceId + FILED_SP + securityGroupId)

	return resourceTencentCloudCvmSecurityGroupAttachmentRead(d, meta)
}

func resourceTencentCloudCvmSecurityGroupAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_security_group_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CvmService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := idSplit[0]
	securityGroupId := idSplit[1]

	instanceInfo, err := service.DescribeInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}

	if instanceInfo == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CvmSecurityGroupAttachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	for _, sgId := range instanceInfo.SecurityGroupIds {
		if *sgId == securityGroupId {
			_ = d.Set("instance_id", instanceId)
			_ = d.Set("security_group_id", securityGroupId)
			return nil

		}
	}
	return fmt.Errorf("The security group get from api does not match with current instance %v", d.Id())
}

func resourceTencentCloudCvmSecurityGroupAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_security_group_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := idSplit[0]
	securityGroupId := idSplit[1]

	request := cvm.NewDisassociateSecurityGroupsRequest()
	request.SecurityGroupIds = []*string{&securityGroupId}
	request.InstanceIds = []*string{&instanceId}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().DisassociateSecurityGroups(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete cvm securityGroupAttachment failed, reason:%+v", logId, err)
		return err
	}

	return nil
}
