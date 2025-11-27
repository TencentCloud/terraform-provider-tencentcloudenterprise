/*
Provides a resource to create a tsf application_file_config

# Example Usage

```hcl

	resource "tencentcloudenterprise_tsf_application_file_config" "application_file_config" {
	  config_name = "terraform-test"
	  config_version = "1.0"
	  config_file_name = "application.yaml"
	  config_file_value = "test: 1"
	  application_id = "application-a24x29xv"
	  config_file_path = "/etc/nginx"
	  config_version_desc = "1.0"
	  config_file_code = "UTF-8"
	  config_post_cmd = "source .bashrc"
	  encode_with_base64 = true
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
	registerResourceDescriptionProvider("tencentcloudenterprise_tsf_application_file_config", CNDescription{
		TerraformTypeCN: "文件配置",
		DescriptionCN:   "提供TSF文件配置资源，用于创建和管理文件配置。",
		AttributesCN: map[string]string{
			"config_name":         "配置项名称",
			"config_version":      "配置项版本",
			"config_file_name":    "配置文件名称",
			"config_file_value":   "配置文件内容",
			"application_id":      "应用ID",
			"config_file_path":    "配置文件路径",
			"config_version_desc": "配置项版本描述",
			"config_file_code":    "配置文件编码",
			"config_post_cmd":     "配置文件后置命令",
			"encode_with_base64":  "是否Base64编码",
		},
	})
}

func resourceTencentCloudTsfApplicationFileConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tsf application_file_config",
		Create:      resourceTencentCloudTsfApplicationFileConfigCreate,
		Read:        resourceTencentCloudTsfApplicationFileConfigRead,
		Delete:      resourceTencentCloudTsfApplicationFileConfigDelete,

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

			"config_file_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Config file name.",
			},

			"config_file_value": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Configuration file content (the original content encoding needs to be in utf-8 format, if the ConfigFileCode is gbk, it will be converted in the background).",
			},

			"application_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Config file associated application ID.",
			},

			"config_file_path": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Config release path.",
			},

			"config_version_desc": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Config version description.",
			},

			"config_file_code": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Configuration file encoding, utf-8 or gbk. Note: If you choose gbk, you need the support of a new version of tsf-consul-template (public cloud virtual machines need to use 1.32 tsf-agent, and containers need to obtain the latest tsf-consul-template-docker.tar.gz from the documentation).",
			},

			"config_post_cmd": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Post command.",
			},

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

func resourceTencentCloudTsfApplicationFileConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_file_config.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = tsf.NewCreateFileConfigRequest()
		//response = tsf.NewCreateFileConfigResponse()
		configId string
	)
	if v, ok := d.GetOk("config_name"); ok {
		request.ConfigName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_version"); ok {
		request.ConfigVersion = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_file_name"); ok {
		request.ConfigFileName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_file_value"); ok {
		request.ConfigFileValue = helper.String(v.(string))
	}

	if v, ok := d.GetOk("application_id"); ok {
		request.ApplicationId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_file_path"); ok {
		request.ConfigFilePath = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_version_desc"); ok {
		request.ConfigVersionDesc = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_file_code"); ok {
		request.ConfigFileCode = helper.String(v.(string))
	}

	if v, ok := d.GetOk("config_post_cmd"); ok {
		request.ConfigPostCmd = helper.String(v.(string))
	}

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
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().CreateFileConfig(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		//response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tsf applicationFileConfig failed, reason:%+v", logId, err)
		return err
	}

	//configId = *response.Response.Result.ConfigId
	d.SetId(configId)

	return resourceTencentCloudTsfApplicationFileConfigRead(d, meta)
}

func resourceTencentCloudTsfApplicationFileConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_file_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	//configId := d.Id()
	confitName := d.Get("config_name").(string)

	applicationFileConfig, err := service.DescribeTsfApplicationFileConfigById(ctx, "", confitName)
	if err != nil {
		return err
	}

	if applicationFileConfig == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TsfApplicationFileConfig` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	d.SetId(*applicationFileConfig.ConfigId)

	if applicationFileConfig.ConfigName != nil {
		_ = d.Set("config_name", applicationFileConfig.ConfigName)
	}

	if applicationFileConfig.ConfigVersion != nil {
		_ = d.Set("config_version", applicationFileConfig.ConfigVersion)
	}

	if applicationFileConfig.ConfigFileName != nil {
		_ = d.Set("config_file_name", applicationFileConfig.ConfigFileName)
	}

	if applicationFileConfig.ConfigFileValue != nil {
		_ = d.Set("config_file_value", applicationFileConfig.ConfigFileValue)
	}

	if applicationFileConfig.ApplicationId != nil {
		_ = d.Set("application_id", applicationFileConfig.ApplicationId)
	}

	if applicationFileConfig.ConfigFilePath != nil {
		_ = d.Set("config_file_path", applicationFileConfig.ConfigFilePath)
	}

	if applicationFileConfig.ConfigVersionDesc != nil {
		_ = d.Set("config_version_desc", applicationFileConfig.ConfigVersionDesc)
	}

	if applicationFileConfig.ConfigFileCode != nil {
		_ = d.Set("config_file_code", applicationFileConfig.ConfigFileCode)
	}

	if applicationFileConfig.ConfigPostCmd != nil {
		_ = d.Set("config_post_cmd", applicationFileConfig.ConfigPostCmd)
	}

	return nil
}

func resourceTencentCloudTsfApplicationFileConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_file_config.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}
	configId := d.Id()

	if err := service.DeleteTsfApplicationFileConfigById(ctx, configId); err != nil {
		return err
	}

	return nil
}
