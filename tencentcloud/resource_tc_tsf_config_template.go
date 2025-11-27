/*
Provides a resource to create a tsf config_template

# Example Usage

```hcl

	resource "tencentcloudenterprise_tsf_config_template" "config_template" {
	  config_template_name = "terraform-template-name"
	  config_template_type = "Ribbon"
	  config_template_value = <<-EOT
	    ribbon.ReadTimeout: 5000
	    ribbon.ConnectTimeout: 2000
	    ribbon.MaxAutoRetries: 0
	    ribbon.MaxAutoRetriesNextServer: 1
	    ribbon.OkToRetryOnAllOperations: true
	  EOT
	  config_template_desc = "terraform-test"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tsf_config_template", CNDescription{
		TerraformTypeCN: "配置模板",
		DescriptionCN:   "提供TSF配置模板资源，用于创建和管理TSF配置模板。",
		AttributesCN: map[string]string{
			"config_template_name":  "配置模板名称",
			"config_template_type":  "配置模板类型",
			"config_template_value": "配置模板值",
			"config_template_desc":  "配置模板描述",
			"program_id_list":       "程序包ID列表",
			"create_time":           "创建时间",
			"update_time":           "更新时间",
			"config_template_id":    "配置模板ID",
		},
	})
}

func resourceTencentCloudTsfConfigTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tsf config_template",
		Create:      resourceTencentCloudTsfConfigTemplateCreate,
		Read:        resourceTencentCloudTsfConfigTemplateRead,
		Update:      resourceTencentCloudTsfConfigTemplateUpdate,
		Delete:      resourceTencentCloudTsfConfigTemplateDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: schema.ImportStatePassthrough,
		// },
		Schema: map[string]*schema.Schema{
			"config_template_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Configuration template name.",
			},

			"config_template_type": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Configure the microservice framework corresponding to the template.",
			},

			"config_template_value": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Configure template data.",
			},

			"config_template_desc": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Configuration template description.",
			},

			"program_id_list": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Program id list.",
			},

			"create_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Creation time.",
			},

			"update_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Update time.",
			},

			"config_template_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Template Id.",
			},
		},
	}
}

func resourceTencentCloudTsfConfigTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_config_template.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = tsf.NewCreateConfigTemplateRequest()
		//response   = tsf.NewCreateConfigTemplateResponse()
		templateId string
	)
	if v, ok := d.GetOk("config_template_name"); ok {
		request.ConfigTemplateName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_template_type"); ok {
		request.ConfigTemplateType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_template_value"); ok {
		request.ConfigTemplateValue = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_template_desc"); ok {
		request.ConfigTemplateDesc = helper.String(v.(string))
	}

	//if v, ok := d.GetOk("program_id_list"); ok {
	//	programIdListSet := v.(*schema.Set).List()
	//	for i := range programIdListSet {
	//		programIdList := programIdListSet[i].(string)
	//		request.ProgramIdList = append(request.ProgramIdList, &programIdList)
	//	}
	//}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().CreateConfigTemplate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		//response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tsf configTemplate failed, reason:%+v", logId, err)
		return err
	}

	//templateId = *response.Response.Result.ConfigTemplateId
	d.SetId(templateId)

	return resourceTencentCloudTsfConfigTemplateRead(d, meta)
}

func resourceTencentCloudTsfConfigTemplateRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_config_template.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	//ctx := context.WithValue(context.TODO(), logIdKey, logId)

	//service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	//templateId := d.Id()
	templateName := d.Get("config_template_name").(string)
	request := tsf.NewDescribeConfigTemplatesRequest()
	request.Offset = helper.Uint64(0)
	request.Limit = helper.Uint64(20)
	request.OrderBy = helper.String("create_time")
	request.OrderType = helper.String("1")
	request.SearchWord = helper.String(templateName)

	response, err := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().DescribeConfigTemplates(request)
	if err != nil {
		return err
	}
	configTemplate := response.Response.Result.Content[0]

	if configTemplate == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TsfConfigTemplate` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	d.SetId(*configTemplate.ConfigTemplateId)

	if configTemplate.ConfigTemplateName != nil {
		_ = d.Set("config_template_name", configTemplate.ConfigTemplateName)
	}

	if configTemplate.ConfigTemplateType != nil {
		_ = d.Set("config_template_type", configTemplate.ConfigTemplateType)
	}

	if configTemplate.ConfigTemplateValue != nil {
		_ = d.Set("config_template_value", configTemplate.ConfigTemplateValue)
	}

	if configTemplate.ConfigTemplateDesc != nil {
		_ = d.Set("config_template_desc", configTemplate.ConfigTemplateDesc)
	}

	// if configTemplate.ProgramIdList != nil {
	// 	_ = d.Set("program_id_list", configTemplate.ProgramIdList)
	// }

	if configTemplate.CreateTime != nil {
		_ = d.Set("create_time", configTemplate.CreateTime)
	}

	if configTemplate.UpdateTime != nil {
		_ = d.Set("update_time", configTemplate.UpdateTime)
	}

	if configTemplate.ConfigTemplateId != nil {
		_ = d.Set("config_template_id", configTemplate.ConfigTemplateId)
	}

	return nil
}

func resourceTencentCloudTsfConfigTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_config_template.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := tsf.NewUpdateConfigTemplateRequest()

	templateId := d.Id()

	request.ConfigTemplateId = &templateId

	immutableArgs := []string{"program_id_list", "create_time", "update_time", "config_template_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	if v, ok := d.GetOk("config_template_name"); ok {
		request.ConfigTemplateName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_template_type"); ok {
		request.ConfigTemplateType = helper.String(v.(string))
	}

	if d.HasChange("config_template_value") {
		if v, ok := d.GetOk("config_template_value"); ok {
			request.ConfigTemplateValue = helper.String(v.(string))
		}
	}

	if d.HasChange("config_template_desc") {
		if v, ok := d.GetOk("config_template_desc"); ok {
			request.ConfigTemplateDesc = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().UpdateConfigTemplate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update tsf configTemplate failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTsfConfigTemplateRead(d, meta)
}

func resourceTencentCloudTsfConfigTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_config_template.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}
	templateId := d.Id()

	if err := service.DeleteTsfConfigTemplateById(ctx, templateId); err != nil {
		return err
	}

	return nil
}
