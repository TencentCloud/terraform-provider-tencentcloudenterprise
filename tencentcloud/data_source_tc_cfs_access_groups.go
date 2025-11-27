/*
Use this data source to query the detail information of CFS access group.

# Example Usage

```hcl

	data "tencentcloudenterprise_cfs_access_groups" "access_groups" {
	  access_group_id = "pgroup-7nx89k7l"
	  name            = "test"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cfs "terraform-provider-tencentcloudenterprise/sdk/cfs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cfs_access_groups", CNDescription{
		TerraformTypeCN: "访问组ID",
		DescriptionCN:   "提供CFS文件系统访问组数据源，用于查询文件系统访问组的详细信息。",
		AttributesCN: map[string]string{
			"access_group_id":    "用于查询的指定访问组ID",
			"name":               "用于查询的访问组名称",
			"result_output_file": "用于保存结果",
			"access_group_list":  "CFS访问组的信息列表每个元素都包含以下属性：",
			"description":        "访问组的描述",
			"create_time":        "访问组的创建时间",
		},
	})
}

func dataSourceTencentCloudCfsAccessGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query the detail information of CFS access group.",
		Read:        dataSourceTencentCloudCfsAccessGroupsRead,

		Schema: map[string]*schema.Schema{
			"access_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A specified access group ID used to query.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A access group Name used to query.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			"access_group_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An information list of CFS access group. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the access group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the access group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the access group.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the access group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCfsAccessGroupsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cfs_access_groups.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cfsService := CfsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	var accessGroupId string
	var name string
	if v, ok := d.GetOk("access_group_id"); ok {
		accessGroupId = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	var accessGroups []*cfs.PGroupInfo
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		accessGroups, errRet = cfsService.DescribeAccessGroup(ctx, accessGroupId, name)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}

	accessGroupList := make([]map[string]interface{}, 0, len(accessGroups))
	ids := make([]string, 0, len(accessGroups))
	for _, accessGroup := range accessGroups {
		mapping := map[string]interface{}{
			"access_group_id": accessGroup.PGroupId,
			"name":            accessGroup.Name,
			"description":     accessGroup.DescInfo,
			"create_time":     accessGroup.CDate,
		}
		accessGroupList = append(accessGroupList, mapping)
		ids = append(ids, *accessGroup.PGroupId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("access_group_list", accessGroupList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set cfs access group list fail, reason:%s\n ", logId, err.Error())
		return err
	}
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), accessGroupList); err != nil {
			return err
		}
	}
	return nil
}
