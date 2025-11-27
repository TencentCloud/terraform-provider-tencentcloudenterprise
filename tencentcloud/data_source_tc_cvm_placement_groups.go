/*
Use this data source to query placement groups.

# Example Usage

```hcl

	data "tencentcloudenterprise_cvm_placement_groups" "foo" {
	  placement_group_id = "ps-21q9ibvr"
	  name               = "test"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	// 注册资源中文描述提供者映射
	registerDataDescriptionProvider("tencentcloudenterprise_cvm_placement_groups", CNDescription{
		TerraformTypeCN: "分散置放群组列表",
		DescriptionCN:   "提供分散置放群组列表数据源，用于查询置放群组信息。",
		AttributesCN: map[string]string{
			"placement_group_id":   "分散置放群组 ID",
			"name":                 "分散置放群组的名称",
			"result_output_file":   "用于保存结果的文件",
			"placement_group_list": "分散置放群组信息列表，其中每个元素包含以下属性:",
			"type":                 "分散置放群组的类型",
			"cvm_quota_total":      "分散置放群组内允许容纳云服务器的最大数量",
			"current_num":          "分散置放群组内云服务器的当前数量",
			"instance_ids":         "分散置放群组内云服务器的 ID 列表",
			"create_time":          "分散置放群组的创建时间",
			"strategies":           "查询策略属于策略列表中的放置群组",
			"strategy":             "分散置放群组的策略",
		},
	})
}

func dataSourceTencentCloudPlacementGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query placement groups.",
		Read:        dataSourceTencentCloudPlacementGroupsRead,

		Schema: map[string]*schema.Schema{
			"placement_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the placement group to be queried.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the placement group to be queried.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"strategies": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validateAllowedStringValue(CVM_PLACE_GROUP_STRATEGY),
				},
				Description: "Placement group strategies. Valid values: `SPREAD`, `WEAK_SPREAD`, `AFFINITY`, and `WEAK_AFFINITY`.",
			},
			"placement_group_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An information list of placement group. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"placement_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the placement group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the placement group.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the placement group.",
						},
						"cvm_quota_total": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of hosts in the placement group.",
						},
						"current_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of hosts in the placement group.",
						},
						"instance_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Host IDs in the placement group.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the placement group.",
						},
						"strategy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Placement group strategy.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudPlacementGroupsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cvm_placement_groups.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cvmService := CvmService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	var placementGroupId string
	var name string
	if v, ok := d.GetOk("placement_group_id"); ok {
		placementGroupId = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	strategies := make([]*string, 0)
	if vs, ok := d.GetOk("strategies"); ok {
		strategys := vs.([]interface{})
		for _, strategy := range strategys {
			strategies = append(strategies, helper.String(strategy.(string)))
		}
	}

	var placementGroups []*cvm.DisasterRecoverGroup
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		placementGroups, errRet = cvmService.DescribePlacementGroupByFilter(ctx, placementGroupId, name, strategies)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}

	placementGroupList := make([]map[string]interface{}, 0, len(placementGroups))
	ids := make([]string, 0, len(placementGroups))
	for _, placement := range placementGroups {
		mapping := map[string]interface{}{
			"placement_group_id": placement.DisasterRecoverGroupId,
			"name":               placement.Name,
			"type":               placement.Type,
			"cvm_quota_total":    placement.CvmQuotaTotal,
			"current_num":        placement.CurrentNum,
			"instance_ids":       helper.StringsInterfaces(placement.InstanceIds),
			"create_time":        placement.CreateTime,
			"strategy":           placement.Strategy,
		}
		placementGroupList = append(placementGroupList, mapping)
		ids = append(ids, *placement.DisasterRecoverGroupId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("placement_group_list", placementGroupList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set placement group list fail, reason:%s\n ", logId, err.Error())
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), placementGroupList); err != nil {
			return err
		}
	}
	return nil
}
