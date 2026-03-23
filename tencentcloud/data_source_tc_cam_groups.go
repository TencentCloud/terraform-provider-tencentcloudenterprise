/*
Use this data source to query detailed information of CAM groups

# Example Usage

```hcl
# query by group_id

	data "tencentcloudenterprise_cam_groups" "foo" {
	  group_id = tencentcloudenterprise_cam_group.foo.id
	}

# query by name

	data "tencentcloudenterprise_cam_groups" "bar" {
	  name = "cam-group-test"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"
	"strconv"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cam_groups", CNDescription{
		TerraformTypeCN: "CAM用户组列表",
		DescriptionCN:   "用于查询 CAM 用户组信息的列表。",
		AttributesCN: map[string]string{
			"name":        "用户组名称",
			"group_id":    "用户组 ID",
			"remark":      "备注",
			"group_list":  "用户组列表",
			"create_time": "创建时间",
		},
	})
}

func dataSourceTencentCloudCamGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamGroupsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the CAM group to be queried.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of CAM group to be queried.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the cam group to be queried.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"group_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of CAM groups. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of CAM group.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of CAM group.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of the CAM group.",
						},
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the CAM group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCamGroupsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cam_groups.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	params := make(map[string]interface{})
	if v, ok := d.GetOk("group_id"); ok {
		groupId, e := strconv.Atoi(v.(string))
		if e != nil {
			return e
		} else {
			params["group_id"] = groupId
		}
	}
	if v, ok := d.GetOk("name"); ok {
		params["name"] = v.(string)
	}
	if v, ok := d.GetOk("remark"); ok {
		params["remark"] = v.(string)
	}

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var groups []*cam.GroupInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := camService.DescribeGroupsByFilter(ctx, params)
		if e != nil {
			return retryError(e)
		}
		groups = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM groups failed, reason:%s\n", logId, err.Error())
		return err
	}
	groupList := make([]map[string]interface{}, 0, len(groups))
	ids := make([]string, 0, len(groups))
	for _, group := range groups {
		mapping := map[string]interface{}{
			"name":        *group.GroupName,
			"create_time": *group.CreateTime,
			"group_id":    strconv.Itoa(int(*group.GroupId)),
		}
		if group.Remark != nil {
			mapping["remark"] = *group.Remark
		}
		groupList = append(groupList, mapping)
		ids = append(ids, strconv.Itoa(int(*group.GroupId)))
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("group_list", groupList); e != nil {
		log.Printf("[CRITAL]%s provider set group list fail, reason:%s\n", logId, e.Error())
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), groupList); e != nil {
			return e
		}
	}

	return nil
}
