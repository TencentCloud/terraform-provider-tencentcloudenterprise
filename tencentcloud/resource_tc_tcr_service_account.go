/*
Provides a resource to create a tcr service account.

Example Usage

Create custom account with specified duration days

```hcl
resource "tencenttencentcloudenterprise_tcr_instance" "example" {
  name          = "tf-example-tcr-instance"
  instance_type = "basic"
  delete_bucket = true
  tags = {
    "createdBy" = "terraform"
  }
}

resource "tencenttencentcloudenterprise_tcr_namespace" "example" {
  instance_id    = tencenttencentcloudenterprise_tcr_instance.example.id
  name           = "tf_test_tcr_namespace"
  is_public      = true
  is_auto_scan   = true
  is_prevent_vul = true
  severity       = "medium"
  cve_whitelist_items {
    cve_id = "tf_example_cve_id"
  }
}

resource "tencenttencentcloudenterprise_tcr_service_account" "example" {
  registry_id = tencenttencentcloudenterprise_tcr_instance.example.id
  name        = "tf_example_account"
  permissions {
    resource = tencenttencentcloudenterprise_tcr_namespace.example.name
    actions  = ["tcr:PushRepository", "tcr:PullRepository"]
  }
  description = "tf example for tcr custom account"
  duration    = 10
  disable     = false
  tags = {
    "createdBy" = "terraform"
  }
}
```

With specified expiration time

```hcl
resource "tencentcloudenterprise_tcr_service_account" "example" {
  registry_id = tencentcloudenterprise_tcr_instance.example.id
  name        = "tf_example_account"
  permissions {
    resource = tencentcloudenterprise_tcr_namespace.example.name
    actions  = ["tcr:PushRepository", "tcr:PullRepository"]
  }
  description = "tf example for tcr custom account"
  expires_at  = 1676897989000 //time stamp
  disable     = false
  tags = {
    "createdBy" = "terraform"
  }
}
```

Import

tcr service_account can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_tcr_service_account.service_account registry_id#account_name
```
 */
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strings"

	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tcr_service_account", CNDescription{
		TerraformTypeCN: "TCR服务账号",
		DescriptionCN:   "提供TCR服务账号资源，用于创建和管理容器镜像仓库的服务账号",
		AttributesCN: map[string]string{
			"registry_id":     "实例ID",
			"name":            "服务账号名称",
			"permissions":     "策略列表",
			"description":     "服务账号描述",
			"duration":        "服务账号有效期",
			"disable":         "是否禁用服务账号",
			"tags":            "标签",
			"password":        "服务账号密码",
		},
	})
}

func resourceTencentCloudTcrServiceAccount() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTcrServiceAccountCreate,
		Read:        resourceTencentCloudTcrServiceAccountRead,
		Update:      resourceTencentCloudTcrServiceAccountUpdate,
		Delete:      resourceTencentCloudTcrServiceAccountDelete,
		Description: "Provides a resource to create and manage TCR service account",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"registry_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "instance id.",
			},

			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Service account name.",
			},

			"permissions": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "strategy list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "resource path, currently only supports Namespace. Note: This field may return null, indicating that no valid value can be obtained.",
						},
						"actions": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required:    true,
							Description: "Actions, currently support: `tcr:PushRepository`, `tcr:PullRepository`, `tcr:CreateRepository`, `tcr:CreateHelmChart`, `tcr:DescribeHelmCharts`. Note: This field may return null, indicating that no valid value can be obtained.",
						},
					},
				},
			},

			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Service account description.",
			},

			"duration": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.",
			},

			"expires_at": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Service account expiration time (time stamp, unit: milliseconds).",
			},

			"disable": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "whether to disable Service accounts.",
			},

			"password": {
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
				Description: "Password of the service account.",
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag description list.",
			},
		},
	}
}

func resourceTencentCloudTcrServiceAccountCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_service_account.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = tcr.NewCreateServiceAccountRequest()
		response   = tcr.NewCreateServiceAccountResponse()
		registryId string
		name       string
	)
	if v, ok := d.GetOk("registry_id"); ok {
		request.RegistryId = helper.String(v.(string))
		registryId = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
		name = v.(string)
	}

	if v, ok := d.GetOk("permissions"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			permission := tcr.Permission{}
			if v, ok := dMap["resource"]; ok {
				permission.Resource = helper.String(v.(string))
			}
			if v, ok := dMap["actions"]; ok {
				actionsSet := v.(*schema.Set).List()
				for i := range actionsSet {
					if actionsSet[i] != nil {
						actions := actionsSet[i].(string)
						permission.Actions = append(permission.Actions, &actions)
					}
				}
			}
			request.Permissions = append(request.Permissions, &permission)
		}
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("duration"); ok {
		request.Duration = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("expires_at"); ok {
		request.ExpiresAt = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("disable"); ok {
		request.Disable = helper.Bool(v.(bool))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTCRClient().CreateServiceAccount(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tcr ServiceAccount failed, reason:%+v", logId, err)
		return err
	}

	if !strings.Contains(*response.Response.Name, name) {
		return fmt.Errorf("the name[%s] return from response is not equal to the name[%s] of tf code",
			*response.Response.Name, name)
	}

	d.SetId(strings.Join([]string{registryId, name}, FILED_SP))

	pw := response.Response.Password
	if pw != nil {
		_ = d.Set("password", *pw)
	}

	ctx := context.WithValue(context.Background(), logIdKey, logId)
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{(meta.(*TencentCloudClient).apiV3Conn)}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::tcr:%s:uin/:instance/%s", region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}
	if v, ok := d.GetOk("password"); ok {
		password, err := service.ModifyServiceAccountPassword(ctx, registryId, name, v.(string))
		if err != nil {
			return err
		}
		_ = d.Set("password", password)
	}

	return resourceTencentCloudTcrServiceAccountRead(d, meta)
}

func resourceTencentCloudTcrServiceAccountRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_service_account.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	registryId := idSplit[0]
	name := idSplit[1]

	ServiceAccount, err := service.DescribeTcrServiceAccountById(ctx, registryId, name)
	if err != nil {
		return err
	}

	if ServiceAccount == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TcrServiceAccount` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("registry_id", registryId)
	_ = d.Set("name", name)

	if ServiceAccount.Permissions != nil {
		permissionsList := []interface{}{}
		for _, permission := range ServiceAccount.Permissions {
			permissionsMap := map[string]interface{}{}

			if permission.Resource != nil {
				permissionsMap["resource"] = permission.Resource
			}

			if len(permission.Actions) > 0 {
				permissionsMap["actions"] = helper.StringsInterfaces(permission.Actions)
			}

			permissionsList = append(permissionsList, permissionsMap)
		}

		_ = d.Set("permissions", permissionsList)

	}

	if ServiceAccount.Description != nil {
		_ = d.Set("description", ServiceAccount.Description)
	}

	if ServiceAccount.ExpiresAt != nil {
		_ = d.Set("expires_at", ServiceAccount.ExpiresAt)
	}

	if ServiceAccount.Disable != nil {
		_ = d.Set("disable", ServiceAccount.Disable)
	}

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := TagService{(tcClient)}
	tags, err := tagService.DescribeResourceTags(ctx, "tcr", "instance", tcClient.Region, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudTcrServiceAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_service_account.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)
	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}

	request := tcr.NewModifyServiceAccountRequest()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	registryId := idSplit[0]
	name := idSplit[1]

	request.RegistryId = &registryId
	request.Name = helper.String(TCR_NAME_PREFIX + name)

	immutableArgs := []string{"registry_id", "name"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	if d.HasChange("permissions") {
		if v, ok := d.GetOk("permissions"); ok {
			for _, item := range v.([]interface{}) {
				permission := tcr.Permission{}
				dMap := item.(map[string]interface{})
				if v, ok := dMap["resource"]; ok {
					permission.Resource = helper.String(v.(string))
				}
				if v, ok := dMap["actions"]; ok {
					actionsSet := v.(*schema.Set).List()
					for i := range actionsSet {
						if actionsSet[i] != nil {
							actions := actionsSet[i].(string)
							permission.Actions = append(permission.Actions, &actions)
						}
					}
				}
				request.Permissions = append(request.Permissions, &permission)
			}
		}
	}

	if d.HasChange("description") {
		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}
	}

	if d.HasChange("duration") {
		if v, ok := d.GetOkExists("duration"); ok {
			request.Duration = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("expires_at") {
		if v, ok := d.GetOkExists("expires_at"); ok {
			request.ExpiresAt = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("disable") {
		if v, ok := d.GetOkExists("disable"); ok {
			request.Disable = helper.Bool(v.(bool))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTCRClient().ModifyServiceAccount(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update tcr ServiceAccount failed, reason:%+v", logId, err)
		return err
	}

	if d.HasChange("tags") {
		ctx := context.WithValue(context.Background(), logIdKey, logId)
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := TagService{tcClient}
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := tagService.DiffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("tcr", "instance", tcClient.Region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	if d.HasChange("password") {
		if v, ok := d.GetOk("password"); ok {
			password, err := service.ModifyServiceAccountPassword(ctx, registryId, name, v.(string))
			if err != nil {
				return err
			}
			_ = d.Set("password", password)
		}
	}
	return resourceTencentCloudTcrServiceAccountRead(d, meta)
}

func resourceTencentCloudTcrServiceAccountDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_service_account.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	registryId := idSplit[0]
	name := TCR_NAME_PREFIX + idSplit[1]

	if err := service.DeleteTcrServiceAccountById(ctx, registryId, name); err != nil {
		return err
	}

	return nil
}
