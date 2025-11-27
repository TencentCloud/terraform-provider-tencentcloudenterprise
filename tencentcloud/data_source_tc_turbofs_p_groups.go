/*
Use this data source to query the detail information of Turbofs permission group.

Example Usage
```hcl

	data "tencentcloudenterprise_turbofs_p_groups" "pgroup" {
	  p_group_id = "pgroup-7nx89k7l"
	}

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
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_turbofs_p_groups", CNDescription{
		TerraformTypeCN: "权限组",
		DescriptionCN:   "提供TurboFS权限组数据源，用于查询TurboFS权限组的详细信息。",
		AttributesCN: map[string]string{
			"p_group_id":         "用于查询的指定权限组ID",
			"name":               "用于查询的权限组名称",
			"result_output_file": "用于保存结果，可视化界面不可用",
			"pgroup_list":        "Turbofs权限组的信息列表",
			"desc_info":          "权限组的描述信息",
			"c_date":             "权限组的创建时间",
			"bind_cfs_num":       "关联文件系统个数",
		},
	})
}

func dataSourceTencentCloudTurbofsPGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query the detail information of Turbofs permission group.",
		Read:        dataSourceTencentCloudTurbofsAccessGroupsRead,

		Schema: map[string]*schema.Schema{
			"p_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A specified permission group ID used to query.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"p_group_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An information list of TurboFS permission group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"p_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the permission group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the permission group.",
						},
						"desc_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description info of the permission group.",
						},
						"c_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation datetime of the permission group.",
						},
						"bind_cfs_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of file systems associated with the permission group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudTurbofsAccessGroupsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_turbofs_p_groups.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	var pGroupId string
	if v, ok := d.GetOk("p_group_id"); ok {
		pGroupId = v.(string)
	}

	var pGroups []*turbofs.PGroupInfo
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		pGroups, errRet = turbofsService.DescribeCfsPGroup(ctx, pGroupId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}

	pGroupList := make([]map[string]interface{}, 0, len(pGroups))
	ids := make([]string, 0, len(pGroups))
	for _, pGroup := range pGroups {
		mapping := map[string]interface{}{
			"p_group_id":   pGroup.PGroupId,
			"name":         pGroup.Name,
			"desc_info":    pGroup.DescInfo,
			"c_date":       pGroup.CDate,
			"bind_cfs_num": pGroup.BindCfsNum,
		}
		pGroupList = append(pGroupList, mapping)
		ids = append(ids, *pGroup.PGroupId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("p_group_list", pGroupList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set turbofs permission group list fail, reason:%s\n ", logId, err.Error())
		return err
	}
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), pGroupList); err != nil {
			return err
		}
	}
	return nil
}
