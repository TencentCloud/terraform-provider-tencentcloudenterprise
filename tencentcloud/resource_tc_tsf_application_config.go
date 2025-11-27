/*
Provides a resource to create a tsf application_config

# Example Usage

```hcl

	resource "tencentcloudenterprise_tsf_application_config" "application_config" {
		config_name = "tf-test-config"
		config_version = "1.0"
		config_value = "name: \"name\""
		application_id = "application-oydl6xa2"
		config_version_desc = "version desc"
		// config_type = ""
		// encode_with_base64 = false
		// program_id_list =
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"log"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tsf_application_config", CNDescription{
		TerraformTypeCN: "应用配置",
		DescriptionCN:   "提供TSF应用配置资源，用于创建和管理应用配置。",
		AttributesCN: map[string]string{
			"config_name":         "配置项名称",
			"config_version":      "配置项版本",
			"config_value":        "配置项值",
			"application_id":      "应用ID",
			"config_version_desc": "配置项版本描述",
			"config_type":         "配置项值类型",
			"encode_with_base64":  "是否Base64编码",
		},
	})
}

func resourceTencentCloudTsfApplicationConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tsf application_config",
		Create:      resourceTencentCloudTsfApplicationConfigCreate,
		Read:        resourceTencentCloudTsfApplicationConfigRead,
		Update:      resourceTencentCloudTsfApplicationConfigUpdate,
		Delete:      resourceTencentCloudTsfApplicationConfigDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: schema.ImportStatePassthrough,
		// },
		Schema: map[string]*schema.Schema{
			"config_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Configuration item name.",
			},

			"config_version": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Configuration item version.",
			},

			"config_value": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Configuration item value.",
			},

			"application_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Application ID.",
			},

			"config_version_desc": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Configuration item version description.",
			},

			"config_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Configuration item value type.",
			},

			"encode_with_base64": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Base64 encoded configuration items.",
			},

			//"program_id_list": {
			//	Optional: true,
			//	Type:     schema.TypeSet,
			//	Elem: &schema.Schema{
			//		Type: schema.TypeString,
			//	},
			//	Description: "Program id list.",
			//},

			// "tags": {
			// 	Type:        schema.TypeMap,
			// 	Optional:    true,
			// 	Description: "Tag description list.",
			// },
		},
	}
}

func resourceTencentCloudTsfApplicationConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_config.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = tsf.NewCreateConfigRequest()
		//response = tsf.NewCreateConfigResponse()
		//configId string
	)
	if v, ok := d.GetOk("config_name"); ok {
		request.ConfigName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_version"); ok {
		request.ConfigVersion = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_value"); ok {
		request.ConfigValue = helper.String(v.(string))
	}

	if v, ok := d.GetOk("application_id"); ok {
		request.ApplicationId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_version_desc"); ok {
		request.ConfigVersionDesc = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_type"); ok {
		request.ConfigType = helper.String(v.(string))
	}

	if v, _ := d.GetOk("encode_with_base64"); v != nil {
		request.EncodeWithBase64 = helper.Bool(v.(bool))
	}

	//if v, ok := d.GetOk("program_id_list"); ok {
	//	programIdListSet := v.(*schema.Set).List()
	//	for i := range programIdListSet {
	//		programIdList := programIdListSet[i].(string)
	//		request.ProgramIdList = append(request.ProgramIdList, &programIdList)
	//	}
	//}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().CreateConfig(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		//response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tsf applicationConfig failed, reason:%+v", logId, err)
		return err
	}

	//configId = *response.Response.Result.ConfigId
	//d.SetId(configId)

	return resourceTencentCloudTsfApplicationConfigRead(d, meta)
}

func resourceTencentCloudTsfApplicationConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	//configId := d.Id()
	configName := d.Get("config_name").(string)

	applicationConfig, err := service.DescribeTsfApplicationConfigById(ctx, "", configName)
	if err != nil {
		return err
	}

	if applicationConfig == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TsfApplicationConfig` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil

	}

	d.SetId(*applicationConfig.ConfigId)
	if applicationConfig.ConfigName != nil {
		_ = d.Set("config_name", applicationConfig.ConfigName)
	}

	if applicationConfig.ConfigVersion != nil {
		_ = d.Set("config_version", applicationConfig.ConfigVersion)
	}

	if applicationConfig.ConfigValue != nil {
		_ = d.Set("config_value", applicationConfig.ConfigValue)
	}

	if applicationConfig.ApplicationId != nil {
		_ = d.Set("application_id", applicationConfig.ApplicationId)
	}

	if applicationConfig.ConfigVersionDesc != nil {
		_ = d.Set("config_version_desc", applicationConfig.ConfigVersionDesc)
	}

	if applicationConfig.ConfigType != nil {
		_ = d.Set("config_type", applicationConfig.ConfigType)
	}

	// if applicationConfig.EncodeWithBase64 != nil {
	// 	_ = d.Set("encode_with_base64", applicationConfig.EncodeWithBase64)
	// }

	// if applicationConfig.ProgramIdList != nil {
	// 	_ = d.Set("program_id_list", applicationConfig.ProgramIdList)
	// }

	return nil
}

func resourceTencentCloudTsfApplicationConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_microservice.update")()
	defer inconsistentCheck(d, meta)()

	immutableArgs := []string{
		"config_name",
		"config_version",
		"config_value",
		"application_id",
		"config_version_desc",
		"config_type",
		"encode_with_base64",
		"program_id_list",
	}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	return resourceTencentCloudTsfApplicationConfigRead(d, meta)
}

func resourceTencentCloudTsfApplicationConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_config.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}
	configId := d.Id()

	if err := service.DeleteTsfApplicationConfigById(ctx, configId); err != nil {
		return err
	}

	return nil
}
