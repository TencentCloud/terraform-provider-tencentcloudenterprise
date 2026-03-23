/*
Provides a resource to create an organization quit_organization_operation

Example Usage

```hcl
resource "tencentcloudenterprise_organization_quit_organization_operation" "quit_organization_operation" {
  org_id = 45155
}
```

Import

organization quit_organization_operation can be imported using the id, e.g.

```
terraform import tencenttencentcloudenterprise_organization_quit_organization_operation.quit_organization_operation quit_organization_operation_id
```
 */
package tencentcloud

import (
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"

	organization "terraform-provider-tencentcloudenterprise/sdk/organization/v20220508"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_organization_quit_organization_operation", CNDescription{
		TerraformTypeCN: "企业组织退出操作",
		DescriptionCN:   "提供企业组织退出操作资源，用于执行退出企业组织的操作。",
		AttributesCN: map[string]string{
			"org_id": "组织ID",
		},
	})
}

func resourceTencentCloudOrganizationQuitOrganizationOperation() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudOrganizationQuitOrganizationOperationCreate,
		Read:        resourceTencentCloudOrganizationQuitOrganizationOperationRead,
		Delete:      resourceTencentCloudOrganizationQuitOrganizationOperationDelete,
		Description: "Provides a resource to execute organization quit operation",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"org_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "Organization ID.",
			},
		},
	}
}

func resourceTencentCloudOrganizationQuitOrganizationOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_quit_organization_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = organization.NewQuitOrganizationRequest()
		orgId   uint64
	)
	if v, _ := d.GetOk("org_id"); v != nil {
		request.OrgId = helper.IntUint64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOrganizationClient().QuitOrganization(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate organization quitOrganizationOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(helper.UInt64ToStr(orgId))

	return resourceTencentCloudOrganizationQuitOrganizationOperationRead(d, meta)
}

func resourceTencentCloudOrganizationQuitOrganizationOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_quit_organization_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudOrganizationQuitOrganizationOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_quit_organization_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
