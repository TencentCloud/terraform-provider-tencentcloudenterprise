/*
Provide a resource to create a DASB command template

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_cmd_template" "example" {
	  name     = "example-cmd-template"
	  cmd_list = "rm -rf*"
	}

```

# Import

DASB command template can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_cmd_template.example 12345
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_cmd_template", CNDescription{
		TerraformTypeCN: "DASB 命令模板",
		DescriptionCN:   "提供 DASB 命令模板资源，用于创建和管理高危命令模板。",
		AttributesCN: map[string]string{
			"name":     "名称",
			"cmd_list": "命令列表",
		},
	})
}

func ResourceTencentCloudDasbCmdTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbCmdTemplateCreate,
		Read:   resourceTencentCloudDasbCmdTemplateRead,
		Update: resourceTencentCloudDasbCmdTemplateUpdate,
		Delete: resourceTencentCloudDasbCmdTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Template name, maximum length 32 characters, cannot contain blank characters.",
			},
			"cmd_list": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Command list, n separated, maximum length 32768 bytes.",
			},
		},
	}
}

func resourceTencentCloudDasbCmdTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_cmd_template.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = bhsaas.NewCreateCmdTemplateRequest()
		response   = bhsaas.NewCreateCmdTemplateResponse()
		templateId string
	)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("cmd_list"); ok {
		request.CmdList = helper.String(v.(string))
	}

	request.Encoding = helper.IntUint64(0)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateCmdTemplate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response.Id == nil {
			e = fmt.Errorf("dasb CmdTemplate not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb CmdTemplate failed, reason:%+v", logId, err)
		return err
	}

	templateIdInt := *response.Response.Id
	templateId = strconv.FormatUint(templateIdInt, 10)
	d.SetId(templateId)

	return resourceTencentCloudDasbCmdTemplateRead(d, meta)
}

func resourceTencentCloudDasbCmdTemplateRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_cmd_template.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		templateId = d.Id()
	)

	CmdTemplate, err := service.DescribeDasbCmdTemplateById(ctx, templateId)
	if err != nil {
		return err
	}

	if CmdTemplate == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DasbCmdTemplate` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if CmdTemplate.Name != nil {
		_ = d.Set("name", CmdTemplate.Name)
	}

	if CmdTemplate.CmdList != nil {
		_ = d.Set("cmd_list", CmdTemplate.CmdList)
	}

	return nil
}

func resourceTencentCloudDasbCmdTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_cmd_template.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = bhsaas.NewModifyCmdTemplateRequest()
		templateId = d.Id()
	)

	request.Id = helper.StrToUint64Point(templateId)
	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("cmd_list"); ok {
		request.CmdList = helper.String(v.(string))
	}

	request.Encoding = helper.IntUint64(0)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyCmdTemplate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update dasb CmdTemplate failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudDasbCmdTemplateRead(d, meta)
}

func resourceTencentCloudDasbCmdTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_cmd_template.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		templateId = d.Id()
	)

	if err := service.DeleteDasbCmdTemplateById(ctx, templateId); err != nil {
		return err
	}

	return nil
}
