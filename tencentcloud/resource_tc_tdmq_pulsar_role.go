package tencentcloud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_pulsar_role", CNDescription{
		TerraformTypeCN: "TDMQ Pulsar角色",
		DescriptionCN:   "提供TDMQ Pulsar角色资源，用于创建和管理TDMQ Pulsar角色。",
		AttributesCN: map[string]string{
			"role_name":  "角色名称",
			"cluster_id": "集群ID",
			"remark":     "角色描述",
		},
	})
}

func resourceTencentCloudTdmqPulsarRole() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTdmqPulsarRoleCreate,
		Read:        resourceTencentCloudTdmqPulsarRoleRead,
		Update:      resourceTencentCloudTdmqPulsarRoleUpdate,
		Delete:      resourceTencentCloudTdmqPulsarRoleDelete,
		Description: "Provides a resource to create and manage TDMQ Pulsar role",

		Schema: map[string]*schema.Schema{
			"role_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of tdmq role.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of tdmq cluster.",
			},
			"remark": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The description of tdmq role.",
			},
		},
	}
}

func resourceTencentCloudTdmqPulsarRoleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_role.create")()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		tdmqService = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		roleName    string
		clusterId   string
		remark      string
	)

	if temp, ok := d.GetOk("role_name"); ok {
		roleName = temp.(string)
		if len(roleName) < 1 {
			return fmt.Errorf("`role_name` should be not empty string")
		}
	}

	if temp, ok := d.GetOk("cluster_id"); ok {
		clusterId = temp.(string)
		if len(clusterId) < 1 {
			return fmt.Errorf("`cluster_id` should be not empty string")
		}
	}

	if temp, ok := d.GetOk("remark"); ok {
		remark = temp.(string)
	}

	respRoleName, err := tdmqService.CreateTdmqRole(ctx, roleName, clusterId, remark)
	if err != nil {
		return err
	}

	d.SetId(respRoleName)
	return resourceTencentCloudTdmqPulsarRoleRead(d, meta)
}

func resourceTencentCloudTdmqPulsarRoleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_role.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId     = getLogId(contextNil)
		ctx       = context.WithValue(context.TODO(), logIdKey, logId)
		roleName  = d.Id()
		clusterId = d.Get("cluster_id").(string)
	)

	tdmqService := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, has, e := tdmqService.DescribeTdmqRoleById(ctx, roleName, clusterId)
		if e != nil {
			return retryError(e)
		}

		if !has {
			d.SetId("")
			return nil
		}

		_ = d.Set("role_name", info.RoleName)
		_ = d.Set("remark", info.Remark)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceTencentCloudTdmqPulsarRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_role.update")()

	var (
		logId     = getLogId(contextNil)
		ctx       = context.WithValue(context.TODO(), logIdKey, logId)
		service   = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		roleName  = d.Id()
		clusterId = d.Get("cluster_id").(string)
		remark    string
	)

	immutableArgs := []string{"role_name", "cluster_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	old, now := d.GetChange("remark")
	if d.HasChange("remark") {
		remark = now.(string)
	} else {
		remark = old.(string)
	}

	d.Partial(true)

	if err := service.ModifyTdmqRoleAttribute(ctx, roleName, clusterId, remark); err != nil {
		return err
	}

	d.Partial(false)
	return resourceTencentCloudTdmqPulsarRoleRead(d, meta)
}

func resourceTencentCloudTdmqPulsarRoleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_role.delete")()

	var (
		logId     = getLogId(contextNil)
		ctx       = context.WithValue(context.TODO(), logIdKey, logId)
		service   = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		roleName  = d.Id()
		clusterId = d.Get("cluster_id").(string)
	)

	err := service.DeleteTdmqRole(ctx, roleName, clusterId)
	if err != nil {
		return nil
	}

	return nil
}
