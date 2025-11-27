/*
Provides a resource to create a tsf application_release_config

# Example Usage

```hcl

	resource "tencentcloudenterprise_tsf_application_release_config" "application_release_config" {
	  config_id = "dcfg-nalqbqwv"
	  group_id = "group-yxmz72gv"
	  release_desc = "terraform-test"
	}

```

# Import

tsf application_release_config can be imported using the configId#groupId#configReleaseId, e.g.

```
terraform import cloud_tsf_application_release_config.application_release_config dcfg-nalqbqwv#group-yxmz72gv#dcfgr-maeeq2ea
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tsf_application_release_config", CNDescription{
		TerraformTypeCN: "应用配置发布",
		DescriptionCN:   "提供TSF应用配置发布资源，用于创建和管理应用配置发布。",
		AttributesCN: map[string]string{
			"config_id":         "配置项ID",
			"group_id":          "部署组ID",
			"release_desc":      "发布描述",
			"config_release_id": "配置项发布ID",
			"config_name":       "配置项名称",
			"config_version":    "配置项版本",
			"release_time":      "发布时间",
			"group_name":        "部署组名称",
			"namespace_id":      "命名空间ID",
			"namespace_name":    "命名空间名称",
			"cluster_id":        "集群ID",
			"cluster_name":      "集群名称",
			"application_id":    "应用ID",
		},
	})
}

func resourceTencentCloudTsfApplicationReleaseConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tsf application_release_config",
		Create:      resourceTencentCloudTsfApplicationReleaseConfigCreate,
		Read:        resourceTencentCloudTsfApplicationReleaseConfigRead,
		Delete:      resourceTencentCloudTsfApplicationReleaseConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"config_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Configuration ID.",
			},

			"group_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Deployment group ID.",
			},

			"release_desc": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Release description.",
			},

			"config_release_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Configuration item release ID.",
			},

			"config_name": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Configuration item name.",
			},

			"config_version": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Configuration item version.",
			},

			"release_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Release time.",
			},

			"group_name": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Deployment group name.",
			},

			"namespace_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Namespace ID.",
			},

			"namespace_name": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Namespace name.",
			},

			"cluster_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Cluster ID.",
			},

			"cluster_name": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Cluster name.",
			},

			"application_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Application ID.",
			},
		},
	}
}

func resourceTencentCloudTsfApplicationReleaseConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_release_config.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = tsf.NewReleaseConfigRequest()
		configId string
		groupId  string
	)
	if v, ok := d.GetOk("config_id"); ok {
		configId = v.(string)
		request.ConfigId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("group_id"); ok {
		groupId = v.(string)
		request.GroupId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("release_desc"); ok {
		request.ReleaseDesc = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().ReleaseConfig(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tsf applicationReleaseConfig failed, reason:%+v", logId, err)
		return err
	}
	req := tsf.NewDescribeConfigReleasesRequest()
	req.ConfigId = helper.String(configId)
	req.GroupId = helper.String(groupId)
	rsp, err := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().DescribeConfigReleases(req)
	if err != nil {
		log.Printf("[CRITAL]%s create tsf applicationReleaseConfig failed, reason:%+v", logId, err)
		return err
	}

	configReleaseId := *rsp.Response.Result.Content[0].ConfigReleaseId
	d.SetId(configId + FILED_SP + groupId + FILED_SP + configReleaseId)
	return resourceTencentCloudTsfApplicationReleaseConfigRead(d, meta)
}

func resourceTencentCloudTsfApplicationReleaseConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_release_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	configId := idSplit[0]
	groupId := idSplit[1]

	applicationReleaseConfig, err := service.DescribeTsfApplicationReleaseConfigById(ctx, configId, groupId)
	if err != nil {
		return err
	}

	if applicationReleaseConfig == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TsfApplicationReleaseConfig` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if applicationReleaseConfig.ConfigId != nil {
		_ = d.Set("config_id", applicationReleaseConfig.ConfigId)
	}

	if applicationReleaseConfig.GroupId != nil {
		_ = d.Set("group_id", applicationReleaseConfig.GroupId)
	}

	if applicationReleaseConfig.ReleaseDesc != nil {
		_ = d.Set("release_desc", applicationReleaseConfig.ReleaseDesc)
	}

	if applicationReleaseConfig.ConfigReleaseId != nil {
		_ = d.Set("config_release_id", applicationReleaseConfig.ConfigReleaseId)
	}

	if applicationReleaseConfig.ConfigName != nil {
		_ = d.Set("config_name", applicationReleaseConfig.ConfigName)
	}

	if applicationReleaseConfig.ConfigVersion != nil {
		_ = d.Set("config_version", applicationReleaseConfig.ConfigVersion)
	}

	if applicationReleaseConfig.ReleaseTime != nil {
		_ = d.Set("release_time", applicationReleaseConfig.ReleaseTime)
	}

	if applicationReleaseConfig.GroupName != nil {
		_ = d.Set("group_name", applicationReleaseConfig.GroupName)
	}

	if applicationReleaseConfig.NamespaceId != nil {
		_ = d.Set("namespace_id", applicationReleaseConfig.NamespaceId)
	}

	if applicationReleaseConfig.NamespaceName != nil {
		_ = d.Set("namespace_name", applicationReleaseConfig.NamespaceName)
	}

	if applicationReleaseConfig.ClusterId != nil {
		_ = d.Set("cluster_id", applicationReleaseConfig.ClusterId)
	}

	if applicationReleaseConfig.ClusterName != nil {
		_ = d.Set("cluster_name", applicationReleaseConfig.ClusterName)
	}

	if applicationReleaseConfig.ApplicationId != nil {
		_ = d.Set("application_id", applicationReleaseConfig.ApplicationId)
	}

	return nil
}

func resourceTencentCloudTsfApplicationReleaseConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_application_release_config.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	configId := idSplit[2]

	if err := service.DeleteTsfApplicationReleaseConfigById(ctx, configId); err != nil {
		return err
	}

	return nil
}
