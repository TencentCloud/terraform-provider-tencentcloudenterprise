/*
Provides a CWP license bind attachment resource.

~> **NOTE:** The license_id is automatically queried from the license order and is not a user input parameter.

Example Usage

```hcl
# Basic CWP license bind attachment
resource "tencentcloudenterprise_cwp_license_bind_attachment" "example" {
  resource_id  = "cwplic-442d44a0"
  license_type = 5
  quuid        = "5c987cf1-b3c9-4b5c-ad45-6787d51f34d7"
}

# Batch bind multiple machines
resource "tencentcloudenterprise_cwp_license_bind_attachment" "batch_bind" {
  for_each = toset(["5c987cf1-b3c9-4b5c-ad45-6787d51f34d7", "another-quuid"])

  resource_id  = "cwplic-442d44a0"
  license_type = 5
  quuid        = each.value
}
```

Import

CWP license bind attachment can be imported using the resource_id#quuid#license_type, e.g.

```
$ terraform import tencentcloudenterprise_cwp_license_bind_attachment.example cwplic-442d44a0#5c987cf1-b3c9-4b5c-ad45-6787d51f34d7#5
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	cwp "terraform-provider-tencentcloudenterprise/sdk/cwp/v20180228"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cwp_license_bind_attachment", CNDescription{
		TerraformTypeCN: "主机安全许可证绑定",
		DescriptionCN:   "提供主机安全许可证绑定资源，用于将许可证绑定到主机。",
		AttributesCN: map[string]string{
			"resource_id":   "资源ID",
			"license_id":    "许可证ID",
			"license_type":  "许可证类型",
			"quuid":         "主机唯一标识",
		},
	})
}

func resourceTencentCloudCwpLicenseBindAttachment() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudCwpLicenseBindAttachmentCreate,
		Read:        resourceTencentCloudCwpLicenseBindAttachmentRead,
		Delete:      resourceTencentCloudCwpLicenseBindAttachmentDelete,
		Description: "Provides a resource to create and manage CWP license bind attachment.",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Resource ID of the license.",
			},
			"license_id": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "License ID (automatically queried from license order).",
			},
			"license_type": {
				Required:     true,
				ForceNew:     true,
				Type:         schema.TypeInt,
				Description:  "License type: 0=CWP Pro Pay-as-you-go, 1=CWP Pro Monthly, 5=CWP Ultimate Monthly.",
			},
			"quuid": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Machine unique identifier (UUID).",
			},
			"machine_name": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Machine name.",
			},
			"machine_wan_ip": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Machine WAN IP.",
			},
			"machine_ip": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Machine IP.",
			},
			"uuid": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Machine UUID.",
			},
			"agent_status": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Agent status.",
			},
			"is_unbind": {
				Computed:    true,
				Type:        schema.TypeBool,
				Description: "Allow unbinding.",
			},
			"is_switch_bind": {
				Computed:    true,
				Type:        schema.TypeBool,
				Description: "Allow switch binding.",
			},
		},
	}
}

func resourceTencentCloudCwpLicenseBindAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cwp_license_bind_attachment.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		service     = CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
		request     = cwp.NewModifyLicenseBindsRequest()
		response    = cwp.NewModifyLicenseBindsResponse()
		taskRequest = cwp.NewDescribeLicenseBindScheduleRequest()
		resourceId  string
		quuid       string
		licenseType string
	)

	if v, ok := d.GetOk("resource_id"); ok {
		request.ResourceId = helper.String(v.(string))
		resourceId = v.(string)
	}

	if v, ok := d.GetOkExists("license_type"); ok {
		request.LicenseType = helper.IntUint64(v.(int))
		licenseTypeInt := v.(int)
		licenseType = strconv.Itoa(licenseTypeInt)
	}

	if v, ok := d.GetOk("quuid"); ok {
		quuid = v.(string)
		request.QuuidList = append(request.QuuidList, &quuid)
	}

	// Set IsAll to false since we're binding specific machines
	request.IsAll = helper.Bool(false)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().ModifyLicenseBinds(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil {
			e = fmt.Errorf("cwp licenseBindAttachment not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cwp licenseBindAttachment failed, reason:%+v", logId, err)
		return err
	}

	// Query license_id from the license order
	licenseOrder, err := service.DescribeCwpLicenseOrderById(ctx, resourceId)
	if err != nil {
		return fmt.Errorf("failed to query license_id from license order: %v", err)
	}
	if licenseOrder == nil || licenseOrder.LicenseId == nil {
		return fmt.Errorf("license order not found or license_id is empty for resource_id: %s", resourceId)
	}
	licenseId := strconv.FormatUint(*licenseOrder.LicenseId, 10)
	log.Printf("[DEBUG]%s queried license_id: %s for resource_id: %s\n", logId, licenseId, resourceId)

	d.SetId(strings.Join([]string{resourceId, quuid, licenseType}, FILED_SP))

	// wait
	taskRequest.TaskId = response.Response.TaskId
	err = resource.Retry(writeRetryTimeout*6, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().DescribeLicenseBindSchedule(taskRequest)
		if e != nil {
			return retryError(e)
		}

		if result == nil {
			e = fmt.Errorf("license bind schedule failed")
			return resource.NonRetryableError(e)
		}

		if *result.Response.List[0].Status == 1 {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("license bind schedule is processing, status: %d", *result.Response.List[0].Status))
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cwp licenseBindAttachment failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCwpLicenseBindAttachmentRead(d, meta)
}

func resourceTencentCloudCwpLicenseBindAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cwp_license_bind_attachment.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service  = CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	resourceId := idSplit[0]
	quuid := idSplit[1]
	licenseType := idSplit[2]

	licenseTypeInt, _ := strconv.ParseUint(licenseType, 10, 64)

	// Query license_id from the license order
	licenseOrder, err := service.DescribeCwpLicenseOrderById(ctx, resourceId)
	if err != nil {
		return err
	}
	if licenseOrder == nil || licenseOrder.LicenseId == nil {
		d.SetId("")
		log.Printf("[WARN]%s license order [%s] not found, please check if it has been deleted.\n", logId, resourceId)
		return nil
	}
	licenseIdInt := *licenseOrder.LicenseId

	licenseBindAttachment, err := service.DescribeCwpLicenseBindAttachmentById(ctx, resourceId, quuid, licenseIdInt, licenseTypeInt)
	if err != nil {
		return err
	}

	if licenseBindAttachment == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CwpLicenseBindAttachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("resource_id", resourceId)
	_ = d.Set("license_id", licenseIdInt)
	_ = d.Set("license_type", licenseTypeInt)
	_ = d.Set("quuid", quuid)

	if licenseBindAttachment.MachineName != nil {
		_ = d.Set("machine_name", licenseBindAttachment.MachineName)
	}

	if licenseBindAttachment.MachineWanIp != nil {
		_ = d.Set("machine_wan_ip", licenseBindAttachment.MachineWanIp)
	}

	if licenseBindAttachment.MachineIp != nil {
		_ = d.Set("machine_ip", licenseBindAttachment.MachineIp)
	}

	if licenseBindAttachment.Uuid != nil {
		_ = d.Set("uuid", licenseBindAttachment.Uuid)
	}

	if licenseBindAttachment.AgentStatus != nil {
		_ = d.Set("agent_status", licenseBindAttachment.AgentStatus)
	}

	if licenseBindAttachment.IsUnBind != nil {
		_ = d.Set("is_unbind", licenseBindAttachment.IsUnBind)
	}

	if licenseBindAttachment.IsSwitchBind != nil {
		_ = d.Set("is_switch_bind", licenseBindAttachment.IsSwitchBind)
	}

	return nil
}

func resourceTencentCloudCwpLicenseBindAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cwp_license_bind_attachment.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service  = CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	resourceId := idSplit[0]
	quuid := idSplit[1]
	licenseType := idSplit[2]

	if err := service.DeleteCwpLicenseBindAttachmentById(ctx, resourceId, quuid, licenseType); err != nil {
		return err
	}

	return nil
}
