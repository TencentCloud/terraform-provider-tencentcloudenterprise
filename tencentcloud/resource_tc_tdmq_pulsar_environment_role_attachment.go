package tencentcloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_pulsar_environment_role_attachment", CNDescription{
		TerraformTypeCN: "TDMQ Pulsar环境",
		DescriptionCN:   "提供TDMQ Pulsar环境角色授权资源，用于为Pulsar角色授权访问指定环境的权限。",
		AttributesCN: map[string]string{
			"environ_id":  "环境ID",
			"role_name":   "角色名称",
			"permissions": "权限列表",
			"cluster_id":  "集群ID",
		},
	})
}
func resourceTencentCloudTdmqPulsarEnvironmentRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a TDMQ Pulsar Environment role attachment.",
		Create:      resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentCreate,
		Read:        resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentRead,
		Update:      resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentUpdate,
		Delete:      resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"environ_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of tdmq namespace.",
			},
			"role_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of tdmq role.",
			},
			"permissions": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				Description: "The permissions of tdmq role.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of tdmq cluster.",
			},
			//compute
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of resource.",
			},
		},
	}
}

func resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_envrionment_role_attachment.create")()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		tdmqService = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		environId   string
		roleName    string
		permissions []*string
		clusterId   string
	)

	if temp, ok := d.GetOk("environ_id"); ok {
		environId = temp.(string)
		if len(environId) < 1 {
			return fmt.Errorf("environ_id should be not empty string")
		}
	}

	if temp, ok := d.GetOk("role_name"); ok {
		roleName = temp.(string)
		if len(roleName) < 1 {
			return fmt.Errorf("role_name should be not empty string")
		}
	}

	if v, ok := d.GetOk("permissions"); ok {
		for _, id := range v.([]interface{}) {
			permissions = append(permissions, helper.String(id.(string)))
		}
	}

	if temp, ok := d.GetOk("cluster_id"); ok {
		clusterId = temp.(string)
		if len(clusterId) < 1 {
			return fmt.Errorf("cluster_id should be not empty string")
		}
	}

	err := tdmqService.CreateTdmqNamespaceRoleAttachment(ctx, environId, roleName, permissions, clusterId)
	if err != nil {
		return err
	}

	d.SetId(environId + FILED_SP + roleName)

	return resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentRead(d, meta)
}

func resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_envrionment_role_attachment.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		tdmqService = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("environment role id is borken, id is %s", d.Id())
	}

	environId := idSplit[0]
	roleName := idSplit[1]
	clusterId := d.Get("cluster_id").(string)

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, has, e := tdmqService.DescribeTdmqNamespaceRoleAttachment(ctx, environId, roleName, clusterId)
		if e != nil {
			return retryError(e)
		}
		if !has {
			d.SetId("")
			return nil
		}

		_ = d.Set("environ_id", info.EnvironmentId)
		_ = d.Set("role_name", info.RoleName)
		_ = d.Set("permissions", info.Permissions)
		_ = d.Set("create_time", info.CreateTime)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_envrionment_role_attachment.update")()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		service     = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		permissions []*string
	)

	immutableArgs := []string{"environ_id", "role_name", "cluster_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("environment role id is borken, id is %s", d.Id())
	}

	environId := idSplit[0]
	roleName := idSplit[1]
	clusterId := d.Get("cluster_id").(string)

	old, now := d.GetChange("permissions")
	if d.HasChange("permissions") {
		for _, id := range now.([]interface{}) {
			permissions = append(permissions, helper.String(id.(string)))
		}
	} else {
		for _, id := range old.([]interface{}) {
			permissions = append(permissions, helper.String(id.(string)))
		}
	}

	d.Partial(true)

	if err := service.ModifyTdmqNamespaceRoleAttachment(ctx, environId, roleName, permissions, clusterId); err != nil {
		return err
	}

	d.Partial(false)
	return resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentRead(d, meta)
}

func resourceTencentCloudTdmqPulsarEnvironmentRoleAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_envrionment_role_attachment.delete")()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("environment role id is borken, id is %s", d.Id())
	}

	environId := idSplit[0]
	roleName := idSplit[1]
	clusterId := d.Get("cluster_id").(string)

	err := service.DeleteTdmqNamespaceRoleAttachment(ctx, environId, roleName, clusterId)
	if err != nil {
		return err
	}

	return nil
}
