/*
Provide a resource to create a DASB ACL

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_acl" "example" {
	  name                = "example-acl"
	  allow_disk_redirect = true
	  allow_any_account   = true
	}

```

# Import

DASB ACL can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_acl.example 12345
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
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_acl", CNDescription{
		TerraformTypeCN: "DASB 访问权限",
		DescriptionCN:   "提供 DASB 访问权限资源，用于创建和管理访问控制策略。",
		AttributesCN: map[string]string{
			"name":                    "名称",
			"allow_disk_redirect":     "允许磁盘重定向",
			"allow_any_account":       "允许任意账号",
			"allow_clip_file_up":      "允许剪贴板文件上传",
			"allow_clip_file_down":    "允许剪贴板文件下载",
			"allow_clip_text_up":      "允许剪贴板文本上传",
			"allow_clip_text_down":    "允许剪贴板文本下载",
			"allow_file_up":           "允许SFTP文件上传",
			"max_file_up_size":        "文件上传大小限制",
			"allow_file_down":         "允许SFTP文件下载",
			"max_file_down_size":      "文件下载大小限制",
			"user_id_set":             "用户ID集合",
			"user_group_id_set":       "用户组ID集合",
			"device_id_set":           "设备ID集合",
			"device_group_id_set":     "设备组ID集合",
			"account_set":             "账号集合",
			"cmd_template_id_set":     "命令模板ID集合",
			"ac_template_id_set":      "高危DB模板ID集合",
			"allow_disk_file_up":      "允许磁盘文件上传",
			"allow_disk_file_down":    "允许磁盘文件下载",
			"allow_shell_file_up":     "允许Shell文件上传",
			"allow_shell_file_down":   "允许Shell文件下载",
			"allow_file_del":          "允许SFTP文件删除",
			"validate_from":           "生效时间",
			"validate_to":             "失效时间",
			"department_id":           "部门ID",
			"allow_access_credential": "允许访问凭证",
		},
	})
}

func ResourceTencentCloudDasbAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbAclCreate,
		Read:   resourceTencentCloudDasbAclRead,
		Update: resourceTencentCloudDasbAclUpdate,
		Delete: resourceTencentCloudDasbAclDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Acl name.",
			},
			"allow_disk_redirect": {
				Required:    true,
				Type:        schema.TypeBool,
				Description: "Allow disk redirect.",
			},
			"allow_any_account": {
				Required:    true,
				Type:        schema.TypeBool,
				Description: "Allow any account.",
			},
			"allow_clip_file_up": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow clip file up.",
			},
			"allow_clip_file_down": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow clip file down.",
			},
			"allow_clip_text_up": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow clip text up.",
			},
			"allow_clip_text_down": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow clip text down.",
			},
			"allow_file_up": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow sftp up file.",
			},
			"max_file_up_size": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "File upload transfer size limit (artifact parameter, currently unused).",
			},
			"allow_file_down": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow sftp file download.",
			},
			"max_file_down_size": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "File transfer download size limit (reserved parameter, currently unused).",
			},
			"user_id_set": {
				Optional:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Associated set of user IDs.",
			},
			"user_group_id_set": {
				Optional:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Associated user group ID.",
			},
			"device_id_set": {
				Optional:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Associated collection of device IDs.",
			},
			"device_group_id_set": {
				Optional:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Associated device group ID.",
			},
			"account_set": {
				Optional:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Associated accounts.",
			},
			"cmd_template_id_set": {
				Optional:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Associated high-risk command template ID.",
			},
			"ac_template_id_set": {
				Optional:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Associate high-risk DB template IDs.",
			},
			"allow_disk_file_up": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow disk file upload.",
			},
			"allow_disk_file_down": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow disk file download.",
			},
			"allow_shell_file_up": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow shell file upload.",
			},
			"allow_shell_file_down": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow shell file download.",
			},
			"allow_file_del": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow sftp file delete.",
			},
			"validate_from": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Access permission effective time, such as: 2021-09-22T00:00:00+08:00If the effective and expiry time are not filled in, the access rights will be valid for a long time.",
			},
			"validate_to": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Access permission expiration time, such as: 2021-09-23T00:00:00+08:00If the effective and expiry time are not filled in, the access rights will be valid for a long time.",
			},
			"department_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Department id.",
			},
			"allow_access_credential": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Allow access credential,default allow.",
			},
		},
	}
}

func resourceTencentCloudDasbAclCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_acl.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = bhsaas.NewCreateAclRequest()
		response = bhsaas.NewCreateAclResponse()
		aclId    string
	)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("allow_disk_redirect"); ok {
		request.AllowDiskRedirect = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_any_account"); ok {
		request.AllowAnyAccount = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_clip_file_up"); ok {
		request.AllowClipFileUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_clip_file_down"); ok {
		request.AllowClipFileDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_clip_text_up"); ok {
		request.AllowClipTextUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_clip_text_down"); ok {
		request.AllowClipTextDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_file_up"); ok {
		request.AllowFileUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("max_file_up_size"); ok {
		request.MaxFileUpSize = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("allow_file_down"); ok {
		request.AllowFileDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("max_file_down_size"); ok {
		request.MaxFileDownSize = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("user_id_set"); ok {
		userIdSetSet := v.(*schema.Set).List()
		for i := range userIdSetSet {
			userIdSet := userIdSetSet[i].(int)
			request.UserIdSet = append(request.UserIdSet, helper.IntUint64(userIdSet))
		}
	}

	if v, ok := d.GetOk("user_group_id_set"); ok {
		userGroupIdSetSet := v.(*schema.Set).List()
		for i := range userGroupIdSetSet {
			userGroupIdSet := userGroupIdSetSet[i].(int)
			request.UserGroupIdSet = append(request.UserGroupIdSet, helper.IntUint64(userGroupIdSet))
		}
	}

	if v, ok := d.GetOk("device_id_set"); ok {
		deviceIdSetSet := v.(*schema.Set).List()
		for i := range deviceIdSetSet {
			deviceIdSet := deviceIdSetSet[i].(int)
			request.DeviceIdSet = append(request.DeviceIdSet, helper.IntUint64(deviceIdSet))
		}
	}

	if v, ok := d.GetOk("device_group_id_set"); ok {
		deviceGroupIdSetSet := v.(*schema.Set).List()
		for i := range deviceGroupIdSetSet {
			deviceGroupIdSet := deviceGroupIdSetSet[i].(int)
			request.DeviceGroupIdSet = append(request.DeviceGroupIdSet, helper.IntUint64(deviceGroupIdSet))
		}
	}

	if v, ok := d.GetOk("account_set"); ok {
		accountSetSet := v.(*schema.Set).List()
		for i := range accountSetSet {
			accountSet := accountSetSet[i].(string)
			request.AccountSet = append(request.AccountSet, &accountSet)
		}
	}

	if v, ok := d.GetOk("cmd_template_id_set"); ok {
		cmdTemplateIdSetSet := v.(*schema.Set).List()
		for i := range cmdTemplateIdSetSet {
			cmdTemplateIdSet := cmdTemplateIdSetSet[i].(int)
			request.CmdTemplateIdSet = append(request.CmdTemplateIdSet, helper.IntUint64(cmdTemplateIdSet))
		}
	}

	if v, ok := d.GetOk("ac_template_id_set"); ok {
		aCTemplateIdSetSet := v.(*schema.Set).List()
		for i := range aCTemplateIdSetSet {
			aCTemplateIdSet := aCTemplateIdSetSet[i].(string)
			request.ACTemplateIdSet = append(request.ACTemplateIdSet, &aCTemplateIdSet)
		}
	}

	if v, ok := d.GetOkExists("allow_disk_file_up"); ok {
		request.AllowDiskFileUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_disk_file_down"); ok {
		request.AllowDiskFileDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_shell_file_up"); ok {
		request.AllowShellFileUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_shell_file_down"); ok {
		request.AllowShellFileDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_file_del"); ok {
		request.AllowFileDel = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("validate_from"); ok {
		request.ValidateFrom = helper.String(v.(string))
	}

	if v, ok := d.GetOk("validate_to"); ok {
		request.ValidateTo = helper.String(v.(string))
	}

	if v, ok := d.GetOk("department_id"); ok {
		request.DepartmentId = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("allow_access_credential"); ok {
		request.AllowAccessCredential = helper.Bool(v.(bool))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateAcl(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response.Id == nil {
			e = fmt.Errorf("dasb acl not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb acl failed, reason:%+v", logId, err)
		return err
	}

	aclIdInt := *response.Response.Id
	aclId = strconv.FormatUint(aclIdInt, 10)
	d.SetId(aclId)

	return resourceTencentCloudDasbAclRead(d, meta)
}

func resourceTencentCloudDasbAclRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_acl.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		aclId   = d.Id()
	)

	acl, err := service.DescribeDasbAclById(ctx, aclId)
	if err != nil {
		return err
	}

	if acl == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DasbAcl` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if acl.Name != nil {
		_ = d.Set("name", acl.Name)
	}

	if acl.AllowDiskRedirect != nil {
		_ = d.Set("allow_disk_redirect", acl.AllowDiskRedirect)
	}

	if acl.AllowAnyAccount != nil {
		_ = d.Set("allow_any_account", acl.AllowAnyAccount)
	}

	if acl.AllowClipFileUp != nil {
		_ = d.Set("allow_clip_file_up", acl.AllowClipFileUp)
	}

	if acl.AllowClipFileDown != nil {
		_ = d.Set("allow_clip_file_down", acl.AllowClipFileDown)
	}

	if acl.AllowClipTextUp != nil {
		_ = d.Set("allow_clip_text_up", acl.AllowClipTextUp)
	}

	if acl.AllowClipTextDown != nil {
		_ = d.Set("allow_clip_text_down", acl.AllowClipTextDown)
	}

	if acl.AllowFileUp != nil {
		_ = d.Set("allow_file_up", acl.AllowFileUp)
	}

	if acl.MaxFileUpSize != nil {
		_ = d.Set("max_file_up_size", acl.MaxFileUpSize)
	}

	if acl.AllowFileDown != nil {
		_ = d.Set("allow_file_down", acl.AllowFileDown)
	}

	if acl.MaxFileDownSize != nil {
		_ = d.Set("max_file_down_size", acl.MaxFileDownSize)
	}

	if acl.UserSet != nil {
		tmpList := make([]*uint64, 0)
		for _, item := range acl.UserSet {
			if item.Id != nil {
				tmpList = append(tmpList, item.Id)
			}
		}

		_ = d.Set("user_id_set", tmpList)
	}

	if acl.UserGroupSet != nil {
		tmpList := make([]*uint64, 0)
		for _, item := range acl.UserGroupSet {
			if item.Id != nil {
				tmpList = append(tmpList, item.Id)
			}
		}

		_ = d.Set("user_group_id_set", tmpList)
	}

	if acl.DeviceSet != nil {
		tmpList := make([]*uint64, 0)
		for _, item := range acl.DeviceSet {
			if item.Id != nil {
				tmpList = append(tmpList, item.Id)
			}
		}

		_ = d.Set("device_id_set", tmpList)
	}

	if acl.DeviceGroupSet != nil {
		tmpList := make([]*uint64, 0)
		for _, item := range acl.DeviceGroupSet {
			if item.Id != nil {
				tmpList = append(tmpList, item.Id)
			}
		}

		_ = d.Set("device_group_id_set", tmpList)
	}

	if acl.AccountSet != nil {
		_ = d.Set("account_set", acl.AccountSet)
	}

	if acl.CmdTemplateSet != nil {
		tmpList := make([]*uint64, 0)
		for _, item := range acl.CmdTemplateSet {
			if item.Id != nil {
				tmpList = append(tmpList, item.Id)
			}
		}

		_ = d.Set("cmd_template_id_set", tmpList)
	}

	if acl.ACTemplateSet != nil {
		tmpList := make([]*string, 0)
		for _, item := range acl.ACTemplateSet {
			if item.TemplateId != nil {
				tmpList = append(tmpList, item.TemplateId)
			}
		}

		_ = d.Set("ac_template_id_set", tmpList)
	}

	if acl.AllowDiskFileUp != nil {
		_ = d.Set("allow_disk_file_up", acl.AllowDiskFileUp)
	}

	if acl.AllowDiskFileDown != nil {
		_ = d.Set("allow_disk_file_down", acl.AllowDiskFileDown)
	}

	if acl.AllowShellFileUp != nil {
		_ = d.Set("allow_shell_file_up", acl.AllowShellFileUp)
	}

	if acl.AllowShellFileDown != nil {
		_ = d.Set("allow_shell_file_down", acl.AllowShellFileDown)
	}

	if acl.AllowFileDel != nil {
		_ = d.Set("allow_file_del", acl.AllowFileDel)
	}

	if acl.ValidateFrom != nil {
		_ = d.Set("validate_from", acl.ValidateFrom)
	}

	if acl.ValidateTo != nil {
		_ = d.Set("validate_to", acl.ValidateTo)
	}

	if acl.Department != nil {
		departmentId := acl.Department.Id
		if *departmentId != "1" {
			_ = d.Set("department_id", departmentId)
		}
	}

	if acl.AllowAccessCredential != nil {
		_ = d.Set("allow_access_credential", acl.AllowAccessCredential)
	}

	return nil
}

func resourceTencentCloudDasbAclUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_acl.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = bhsaas.NewModifyAclRequest()
		aclId   = d.Id()
	)

	request.Id = helper.StrToUint64Point(aclId)
	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("allow_disk_redirect"); ok {
		request.AllowDiskRedirect = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_any_account"); ok {
		request.AllowAnyAccount = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_clip_file_up"); ok {
		request.AllowClipFileUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_clip_file_down"); ok {
		request.AllowClipFileDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_clip_text_up"); ok {
		request.AllowClipTextUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_clip_text_down"); ok {
		request.AllowClipTextDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_file_up"); ok {
		request.AllowFileUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("max_file_up_size"); ok {
		request.MaxFileUpSize = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("allow_file_down"); ok {
		request.AllowFileDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("max_file_down_size"); ok {
		request.MaxFileDownSize = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("user_id_set"); ok {
		userIdSetSet := v.(*schema.Set).List()
		for i := range userIdSetSet {
			userIdSet := userIdSetSet[i].(int)
			request.UserIdSet = append(request.UserIdSet, helper.IntUint64(userIdSet))
		}
	}

	if v, ok := d.GetOk("user_group_id_set"); ok {
		userGroupIdSetSet := v.(*schema.Set).List()
		for i := range userGroupIdSetSet {
			userGroupIdSet := userGroupIdSetSet[i].(int)
			request.UserGroupIdSet = append(request.UserGroupIdSet, helper.IntUint64(userGroupIdSet))
		}
	}

	if v, ok := d.GetOk("device_id_set"); ok {
		deviceIdSetSet := v.(*schema.Set).List()
		for i := range deviceIdSetSet {
			deviceIdSet := deviceIdSetSet[i].(int)
			request.DeviceIdSet = append(request.DeviceIdSet, helper.IntUint64(deviceIdSet))
		}
	}

	if v, ok := d.GetOk("device_group_id_set"); ok {
		deviceGroupIdSetSet := v.(*schema.Set).List()
		for i := range deviceGroupIdSetSet {
			deviceGroupIdSet := deviceGroupIdSetSet[i].(int)
			request.DeviceGroupIdSet = append(request.DeviceGroupIdSet, helper.IntUint64(deviceGroupIdSet))
		}
	}

	if v, ok := d.GetOk("account_set"); ok {
		accountSetSet := v.(*schema.Set).List()
		for i := range accountSetSet {
			accountSet := accountSetSet[i].(string)
			request.AccountSet = append(request.AccountSet, &accountSet)
		}
	}

	if v, ok := d.GetOk("cmd_template_id_set"); ok {
		cmdTemplateIdSetSet := v.(*schema.Set).List()
		for i := range cmdTemplateIdSetSet {
			cmdTemplateIdSet := cmdTemplateIdSetSet[i].(int)
			request.CmdTemplateIdSet = append(request.CmdTemplateIdSet, helper.IntUint64(cmdTemplateIdSet))
		}
	}

	if v, ok := d.GetOk("ac_template_id_set"); ok {
		aCTemplateIdSetSet := v.(*schema.Set).List()
		for i := range aCTemplateIdSetSet {
			aCTemplateIdSet := aCTemplateIdSetSet[i].(string)
			request.ACTemplateIdSet = append(request.ACTemplateIdSet, &aCTemplateIdSet)
		}
	}

	if v, ok := d.GetOkExists("allow_disk_file_up"); ok {
		request.AllowDiskFileUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_disk_file_down"); ok {
		request.AllowDiskFileDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_shell_file_up"); ok {
		request.AllowShellFileUp = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_shell_file_down"); ok {
		request.AllowShellFileDown = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_file_del"); ok {
		request.AllowFileDel = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("validate_from"); ok {
		request.ValidateFrom = helper.String(v.(string))
	}

	if v, ok := d.GetOk("validate_to"); ok {
		request.ValidateTo = helper.String(v.(string))
	}

	if v, ok := d.GetOk("department_id"); ok {
		request.DepartmentId = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("allow_access_credential"); ok {
		request.AllowAccessCredential = helper.Bool(v.(bool))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyAcl(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update dasb acl failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudDasbAclRead(d, meta)
}

func resourceTencentCloudDasbAclDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_acl.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		aclId   = d.Id()
	)

	if err := service.DeleteDasbAclById(ctx, aclId); err != nil {
		return err
	}

	return nil
}
