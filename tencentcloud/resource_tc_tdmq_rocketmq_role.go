/*
Provides a resource to create a tdmqRocketmq role

# Example Usage

```hcl

	resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
		cluster_name = "test_rocketmq"
		remark = "test recket mq"
	}

	resource "tencentcloudenterprise_tdmq_rocketmq_role" "role" {
	  role_name = "test_rocketmq_role"
	  remark = "test rocketmq role"
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	}

```
Import

tdmqRocketmq role can be imported using the id, e.g.
```
$ terraform import cloud_tdmq_rocketmq_role.role role_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmqRocketmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	// 注册资源到资源中文描述提供者映射里
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_rocketmq_role", CNDescription{
		TerraformTypeCN: "Rocketmq 角色管理",
		DescriptionCN:   "提供Rocketmq角色资源，用于创建和管理Rocketmq角色。",
		AttributesCN: map[string]string{
			"role_name":   "角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32",
			"remark":      "备注说明，长度必须大等于0且小等于128",
			"cluster_id":  "必填字段，集群Id",
			"token":       "角色token",
			"create_time": "创建时间",
			"update_time": "更新时间",
		}},
	)
}

func resourceTencentCloudTdmqRocketmqRole() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tdmqRocketmq role",
		Create:      resourceTencentCloudTdmqRocketmqRoleCreate,
		Read:        resourceTencentCloudTdmqRocketmqRoleRead,
		Update:      resourceTencentCloudTdmqRocketmqRoleUpdate,
		Delete:      resourceTencentCloudTdmqRocketmqRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32
			"role_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Role name, which can contain up to 32 letters, digits, hyphens, and underscores.",
			},
			// 备注说明，长度必须大等于0且小等于128
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Remarks (up to 128 characters).",
			},
			// 必填字段，集群Id
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster ID (required).",
			},
			// 角色token
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Value of the role token.",
			},
			// 创建时间
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time.",
			},
			// 更新时间
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update time.",
			},
		},
	}
}

func resourceTencentCloudTdmqRocketmqRoleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmqRocketmq_role.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request   = tdmqRocketmq.NewCreateRocketMQRoleRequest()
		clusterId string
		roleName  string
	)

	if v, ok := d.GetOk("role_name"); ok {
		roleName = v.(string)
		request.RoleName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("remark"); ok {
		request.Remark = helper.String(v.(string))
	}

	if v, ok := d.GetOk("cluster_id"); ok {
		clusterId = v.(string)
		request.ClusterId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().CreateRocketMQRole(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmqRocketmq role failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(clusterId + FILED_SP + roleName)
	return resourceTencentCloudTdmqRocketmqRoleRead(d, meta)
}

func resourceTencentCloudTdmqRocketmqRoleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmqRocketmq_role.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TdmqRocketmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	roleName := idSplit[1]

	role, err := service.DescribeTdmqRocketmqRole(ctx, clusterId, roleName)

	if err != nil {
		return err
	}

	if role == nil {
		d.SetId("")
		return fmt.Errorf("resource `role` %s does not exist", roleName)
	}

	_ = d.Set("role_name", role.RoleName)
	_ = d.Set("remark", role.Remark)
	_ = d.Set("cluster_id", clusterId)
	_ = d.Set("token", role.Token)
	_ = d.Set("create_time", role.CreateTime)
	_ = d.Set("update_time", role.UpdateTime)

	return nil
}

func resourceTencentCloudTdmqRocketmqRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmqRocketmq_role.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := tdmqRocketmq.NewModifyRocketMQRoleRequest()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	roleName := idSplit[1]

	request.ClusterId = &clusterId
	request.RoleName = &roleName

	if d.HasChange("role_name") {

		return fmt.Errorf("`role_name` do not support change now.")

	}

	if d.HasChange("remark") {
		if v, ok := d.GetOk("remark"); ok {
			request.Remark = helper.String(v.(string))
		}

	}

	if d.HasChange("cluster_id") {

		return fmt.Errorf("`cluster_id` do not support change now.")

	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().ModifyRocketMQRole(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmqRocketmq role failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTdmqRocketmqRoleRead(d, meta)
}

func resourceTencentCloudTdmqRocketmqRoleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmqRocketmq_role.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TdmqRocketmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	roleName := idSplit[1]

	if err := service.DeleteTdmqRocketmqRoleById(ctx, clusterId, roleName); err != nil {
		return err
	}

	return nil
}
