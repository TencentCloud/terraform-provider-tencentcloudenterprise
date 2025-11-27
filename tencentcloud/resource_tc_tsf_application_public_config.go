/*
Provides a resource to create a tsf application_public_config

# Example Usage

```hcl

	resource "tencentcloudenterprise_tsf_application_public_config" "application_public_config" {
	  config_name = "my_config"
	  config_version = "1.0"
	  config_value = "test: 1"
	  config_version_desc = "product version"
	  // config_type = "P"
	  encode_with_base64 = true
	  # program_id_list =
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tsf_application_public_config", CNDescription{
		TerraformTypeCN: "全局配置",
		DescriptionCN:   "提供TSF全局配置资源，用于创建和管理全局配置。",
		AttributesCN: map[string]string{
			"config_name":         "配置名称",
			"config_version":      "版本号",
			"config_value":        "配置内容",
			"config_version_desc": "版本描述",
			"encode_with_base64":  "是否base64编码",
		},
	})
}

func resourceTencentCloudTsfApplicationPublicConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tsf application_public_config",
		Create:      resourceTencentCloudTsfApplicationPublicConfigCreate,
		Read:        resourceTencentCloudTsfApplicationPublicConfigRead,
		Delete:      resourceTencentCloudTsfApplicationPublicConfigDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: schema.ImportStatePassthrough,
		// },
		Schema: map[string]*schema.Schema{
			"config_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Config Name.",
			},

			"config_version": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Config version.",
			},

			"config_value": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Config value, only yaml file allowed.",
			},

			"config_version_desc": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Config version description.",
			},

			//"config_type": {
			//	Optional:    true,
			//	ForceNew:    true,
			//	Type:        schema.TypeString,
			//	Description: "Config type.",
			//},

			"encode_with_base64": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeBool,
				Description: "The config value is encoded with base64 or not.",
			},

			//"program_id_list": {
			//	Optional: true,
			//	ForceNew: true,
			//	Type:     schema.TypeSet,
			//	Elem: &schema.Schema{
			//		Type: schema.TypeString,
			//	},
			//	Description: "Datasource for auth.",
			//},
		},
	}
}

func resourceTencentCloudTsfApplicationPublicConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_public_config.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = tsf.NewCreatePublicConfigRequest()
		//response = tsf.NewCreatePublicConfigResponse()
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

	if v, ok := d.GetOk("config_version_desc"); ok {
		request.ConfigVersionDesc = helper.String(v.(string))
	}

	//if v, ok := d.GetOk("config_type"); ok {
	//	request.ConfigType = helper.String(v.(string))
	//}

	if v, ok := d.GetOkExists("encode_with_base64"); ok {
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
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().CreatePublicConfig(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		//response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tsf applicationPublicConfig failed, reason:%+v", logId, err)
		return err
	}

	//configId = *response.Response.Result.ConfigId
	//d.SetId(configId)

	return resourceTencentCloudTsfApplicationPublicConfigRead(d, meta)
}

func resourceTencentCloudTsfApplicationPublicConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_public_config.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId = getLogId(contextNil)
		//configId := d.Id()
		configName              = d.Get("config_name").(string)
		request                 = tsf.NewDescribePublicConfigSummaryRequest()
		applicationPublicConfig *tsf.Config
	)
	request.Offset = helper.Int64(0)
	request.Limit = helper.Int64(20)
	request.SearchWord = &configName
	request.OrderBy = helper.String("creation_time")
	request.OrderType = helper.Int64(1)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().DescribePublicConfigSummary(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		applicationPublicConfig = result.Response.Result.Content[0]
		return nil
	})

	if err != nil {
		return err
	}

	d.SetId(*applicationPublicConfig.ConfigId)

	if applicationPublicConfig == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TsfApplicationPublicConfig` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if applicationPublicConfig.ConfigName != nil {
		_ = d.Set("config_name", applicationPublicConfig.ConfigName)
	}

	if applicationPublicConfig.ConfigVersion != nil {
		_ = d.Set("config_version", applicationPublicConfig.ConfigVersion)
	}

	if applicationPublicConfig.ConfigValue != nil {
		_ = d.Set("config_value", applicationPublicConfig.ConfigValue)
	}

	if applicationPublicConfig.ConfigVersionDesc != nil {
		_ = d.Set("config_version_desc", applicationPublicConfig.ConfigVersionDesc)
	}

	if applicationPublicConfig.ConfigType != nil {
		_ = d.Set("config_type", applicationPublicConfig.ConfigType)
	}

	// if applicationPublicConfig.EncodeWithBase64 != nil {
	// 	_ = d.Set("encode_with_base64", applicationPublicConfig.EncodeWithBase64)
	// }

	// if applicationPublicConfig.ProgramIdList != nil {
	// 	_ = d.Set("program_id_list", applicationPublicConfig.ProgramIdList)
	// }

	return nil
}

func resourceTencentCloudTsfApplicationPublicConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_public_config.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}
	configId := d.Id()

	if err := service.DeleteTsfApplicationPublicConfigById(ctx, configId); err != nil {
		return err
	}

	return nil
}
