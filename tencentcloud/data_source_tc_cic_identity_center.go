/*
Use this data source to query detailed information of cic identity center

# Example Usage

```hcl

	data "tencentcloudenterprise_cic_identity_center" "identity_center" {
	}

	output "zone_id" {
	  value = data.tencentcloudenterprise_cic_identity_center.identity_center.zone_id
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cic_identity_center", CNDescription{
		TerraformTypeCN: "CIC 身份中心服务信息",
		DescriptionCN:   "提供 CIC 身份中心服务信息数据源，用于查询身份中心的服务状态、空间信息等。",
		AttributesCN: map[string]string{
			"create_time":      "创建时间",
			"scim_sync_status": "SCIM 同步状态。Enabled：启用。Disabled：禁用",
			"service_status":   "服务开启状态，Disabled 代表未开通，Enabled 代表已开通",
			"update_time":      "更新时间",
			"zone_id":             "空间 ID。z-前缀开头，后面是 12 位随机数字/小写字母",
			"zone_name":           "空间名，必须全局唯一",
			"result_output_file": "用于保存查询结果",
		},
	})
}

func dataSourceTencentCloudCicIdentityCenter() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudCicIdentityCenterRead,
		Description: "Use this data source to query detailed information of cic identity center.",
		Schema: map[string]*schema.Schema{
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time.",
			},
			"scim_sync_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SCIM synchronization status. Enabled: enabled; Disabled: disabled.",
			},
			"service_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Service status. Disabled: not opened; Enabled: opened.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update time.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Space ID. It starts with the z- prefix, followed by 12 random digits/lowercase letters.",
			},
			"zone_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Space name, which must be globally unique. It contains lowercase letters, digits, and hyphens (-). It cannot start or end with a hyphen (-), and cannot have two consecutive hyphens (-). Length: 2-64 characters.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCicIdentityCenterRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cic_identity_center.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	identityCenter, err := service.DescribeCicIdentityCenter(ctx)
	if err != nil {
		return err
	}

	if identityCenter == nil || identityCenter.Response == nil {
		d.SetId("")
		return nil
	}

	resp := identityCenter.Response
	_ = d.Set("create_time", resp.CreateTime)
	_ = d.Set("scim_sync_status", resp.ScimSyncStatus)
	_ = d.Set("service_status", resp.ServiceStatus)
	_ = d.Set("update_time", resp.UpdateTime)
	_ = d.Set("zone_id", resp.ZoneId)
	_ = d.Set("zone_name", resp.ZoneName)

	d.SetId("cic_identity_center")

	return nil
}
