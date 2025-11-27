/*
Provides a resource to create a TurboFS permission group.

# Example Usage

```hcl

	resource "tencentcloudenterprise_turbofs_p_group" "foo" {
	  name        = "test_p_group"
	  desc_info = "test"
	}

```

# Import

TurboFS permission group can be imported using the id, e.g.

```
$ terraform import cloud_turbofs_p_group.foo pgroup-7nx89k7l
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_turbofs_p_group", CNDescription{
		TerraformTypeCN: "权限组",
		DescriptionCN:   "提供TurboFS权限组资源，用于创建和管理TurboFS权限组。",
		AttributesCN: map[string]string{
			"name":         "权限组名称",
			"desc_info":    "权限组描述信息",
			"create_time":  "创建时间",
			"p_group_id":   "分组Id",
			"bind_cfs_num": "权限组关联文件系统个数",
			"c_date":       "权限组创建时间",
		},
	})
}
func resourceTencentCloudTurbofsPGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a TurboFS permission group.",
		Create:      resourceTencentCloudTurbofsPGroupCreate,
		Read:        resourceTencentCloudTurbofsPGroupRead,
		Update:      resourceTencentCloudTurbofsPGroupUpdate,
		Delete:      resourceTencentCloudTurbofsPGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the permission group",
			},
			"desc_info": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Description of the permission group.",
			},
			"c_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create datetime of the permission group.",
			},
			"p_group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id of the permission group.",
			},
			"bind_cfs_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of file systems associated with the permission group.",
			},
		},
	}
}

func resourceTencentCloudTurbofsPGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_p_group.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var name, descInfo, pGroupId *string

	if v, ok := d.GetOk("name"); ok {
		name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("desc_info"); ok {
		descInfo = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		id, errRet := turbofsService.CreateCfsPGroup(ctx, name, descInfo)
		if errRet != nil {
			return retryError(errRet)
		}
		pGroupId = id
		return nil
	})
	if err != nil {
		return err
	}
	d.SetId(*pGroupId)

	return resourceTencentCloudTurbofsPGroupRead(d, meta)
}

func resourceTencentCloudTurbofsPGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_p_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	id := d.Id()
	var pGroup *turbofs.PGroupInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		pGroupResult, errRet := turbofsService.DescribeCfsPGroup(ctx, id)
		if errRet != nil {
			return retryError(errRet)
		}
		if pGroupResult != nil && len(pGroupResult) > 0 {
			pGroup = pGroupResult[0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if pGroup == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("name", pGroup.Name)
	_ = d.Set("desc_info", pGroup.DescInfo)
	_ = d.Set("c_date", pGroup.CDate)
	_ = d.Set("p_group_id", pGroup.PGroupId)
	_ = d.Set("bind_cfs_num", pGroup.BindCfsNum)

	return nil
}

func resourceTencentCloudTurbofsPGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_p_group.update")()
	logId := getLogId(contextNil)

	request := turbofs.NewUpdateCfsPGroupRequest()
	if d.HasChange("name") {
		request.Name = helper.String(d.Get("name").(string))
	}
	if d.HasChange("desc_info") {
		request.DescInfo = helper.String(d.Get("desc_info").(string))
	}

	request.PGroupId = helper.String(d.Id())

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, err := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().UpdateCfsPGroup(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			return retryError(err)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		return nil
	})
	if err != nil {
		return err
	}

	return resourceTencentCloudTurbofsPGroupRead(d, meta)
}

func resourceTencentCloudTurbofsPGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_p_group.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := helper.String(d.Id())
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		errRet := turbofsService.DeleteCfsPGroup(ctx, id)
		if errRet != nil {
			return retryError(errRet)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
