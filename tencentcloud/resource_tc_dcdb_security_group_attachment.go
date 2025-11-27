/*
Provides a resource to create a dcdb security_group_attachment

# Example Usage

```hcl

	resource "tencentcloudenterprise_dcdb_security_group_attachment" "security_group_attachment" {
	  security_group_id = "sg-9s7k6qgw"
	  instance_id = "tdsqlshard-973xatu3"
	}

```
Import

dcdb security_group_attachment can be imported using the id, e.g.
```
$ terraform import cloud_dcdb_security_group_attachment.security_group_attachment securityGroupAttachment_id
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
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dcdb_security_group_attachment", CNDescription{
		TerraformTypeCN: "绑定安全组",
		DescriptionCN:   "提供DCDB安全组绑定资源，用于将安全组绑定到DCDB实例。",
		AttributesCN: map[string]string{
			"security_group_id": "安全组id",
			"instance_id":       "附加的实例id",
		},
	})
}

func resourceTencentCloudDcdbSecurityGroupAttachment() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a dcdb security_group_attachment",
		Create:      resourceTencentCloudDcdbSecurityGroupAttachmentCreate,
		Read:        resourceTencentCloudDcdbSecurityGroupAttachmentRead,
		Delete:      resourceTencentCloudDcdbSecurityGroupAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"security_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Security group id.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Attached instance id.",
			},
		},
	}
}

func resourceTencentCloudDcdbSecurityGroupAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_security_group_attachment.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = dcdb.NewAssociateSecurityGroupsRequest()
		// response        *dcdb.AssociateSecurityGroupsResponse
		instanceId      string
		securityGroupId string
	)

	request.Product = helper.String("dcdb") // api only use this fixed value

	if v, ok := d.GetOk("security_group_id"); ok {
		securityGroupId = v.(string)
		request.SecurityGroupId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceIds = []*string{helper.String(v.(string))}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().AssociateSecurityGroups(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		// response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dcdb securityGroupAttachment failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(instanceId + FILED_SP + securityGroupId)
	return resourceTencentCloudDcdbSecurityGroupAttachmentRead(d, meta)
}

func resourceTencentCloudDcdbSecurityGroupAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_security_group_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := idSplit[0]
	securityGroupId := idSplit[1]

	securityGroupAttachmentResponse, err := service.DescribeDcdbSecurityGroup(ctx, instanceId)

	if err != nil {
		return err
	}
	securityGroupAttachment := securityGroupAttachmentResponse.Response

	if securityGroupAttachment == nil {
		d.SetId("")
		return fmt.Errorf("resource `securityGroupAttachment` %s does not exist", d.Id())
	}

	for _, sg := range securityGroupAttachment.Groups {
		if securityGroupId == *sg.SecurityGroupId {
			_ = d.Set("instance_id", instanceId)
			_ = d.Set("security_group_id", securityGroupId)
			return nil
		}
	}

	return fmt.Errorf("The security group get from api does not match with current instance %v", d.Id())
}

func resourceTencentCloudDcdbSecurityGroupAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_security_group_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := idSplit[0]
	securityGroupId := idSplit[1]

	if err := service.DeleteDcdbSecurityGroupAttachmentById(ctx, instanceId, securityGroupId); err != nil {
		return err
	}

	return nil
}
