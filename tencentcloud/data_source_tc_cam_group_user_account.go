/*
Use this data source to query the list of user groups that a sub-user has joined.

# Example Usage

```hcl
# Query by uid

	data "tencentcloudenterprise_cam_group_user_account" "example" {
	  uid = 4364
	}

# Query by uin

	data "tencentcloudenterprise_cam_group_user_account" "example" {
	  uin = 110000001037
	}

```
*/
package tencentcloud

import (
	"fmt"
	"log"
	"strconv"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cam_group_user_account", CNDescription{
		TerraformTypeCN: "子用户所在用户组列表",
		DescriptionCN:   "用于查询子用户已加入的用户组列表。",
		AttributesCN: map[string]string{
			"uid":        "子用户 UID，与 uin 二选一",
			"uin":        "子用户 UIN，与 uid 二选一",
			"total_num":  "子用户已加入的用户组总数",
			"group_info": "用户组信息列表",
		},
	})
}

func dataSourceTencentCloudCamGroupUserAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamGroupUserAccountRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Optional:      true,
				Type:          schema.TypeInt,
				ConflictsWith: []string{"uin"},
				Description:   "Sub-user uid. Conflicts with uin.",
			},
			"uin": {
				Optional:      true,
				Type:          schema.TypeInt,
				ConflictsWith: []string{"uid"},
				Description:   "Sub-user uin. Conflicts with uid.",
			},
			"total_num": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "The total number of user groups the sub-user has joined.",
			},
			"group_info": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "User group information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "User group ID.",
						},
						"group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User group name.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remark.",
						},
					},
				},
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCamGroupUserAccountRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cam_group_user_account.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// Get uid, if uin is provided, convert it to uid first
	var uid *uint64

	if uidInt, ok := d.GetOk("uid"); ok {
		uid = helper.IntUint64(uidInt.(int))
	} else if uinInt, ok := d.GetOk("uin"); ok {
		// Convert uin to uid using GetUserListByUinList
		uin := helper.IntInt64(uinInt.(int))

		request := cam.NewGetUserListByUinListRequest()
		request.UinList = []*int64{uin}

		var response *cam.GetUserListByUinListResponse
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().GetUserListByUinList(request)
			if e != nil {
				return retryError(e)
			}
			response = result
			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s query sub-account by uin failed, reason:%+v", logId, err)
			return err
		}

		if response == nil || response.Response == nil || len(response.Response.UserList) == 0 {
			return fmt.Errorf("uin %d not found", *uin)
		}

		if response.Response.UserList[0].UserCamUid == nil {
			return fmt.Errorf("uin %d has no uid", *uin)
		}

		uidInt64 := response.Response.UserList[0].UserCamUid
		uidUint64 := uint64(*uidInt64)
		uid = &uidUint64
	} else {
		return fmt.Errorf("either uid or uin must be specified")
	}

	// 自动处理分页，获取所有数据
	var allGroupInfo []*cam.GroupInfo
	page := int64(1)
	rp := int64(100) // 每页100条

	for {
		request := cam.NewGetSubsGroupRequest()
		request.Uid = uid
		request.Rp = &rp
		request.Page = &page

		var response *cam.GetSubsGroupResponse
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().GetSubsGroup(request)
			if e != nil {
				return retryError(e)
			}
			response = result
			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s read CAM group user account failed, reason:%+v", logId, err)
			return err
		}

		if response == nil || response.Response == nil {
			break
		}

		// 添加当前页数据
		if response.Response.GroupInfo != nil {
			allGroupInfo = append(allGroupInfo, response.Response.GroupInfo...)
		}

		// 检查是否还有更多数据
		if response.Response.TotalNum != nil {
			totalNum, _ := strconv.Atoi(*response.Response.TotalNum)
			if len(allGroupInfo) >= totalNum {
				break
			}
		}

		// 如果当前页没有返回数据，说明已经到底了
		if len(response.Response.GroupInfo) == 0 {
			break
		}

		page++
	}

	ids := make([]string, 0, len(allGroupInfo))
	tmpList := make([]map[string]interface{}, 0, len(allGroupInfo))

	// Set total number
	_ = d.Set("total_num", len(allGroupInfo))

	if allGroupInfo != nil {
		for _, groupInfo := range allGroupInfo {
			groupInfoMap := map[string]interface{}{}

			if groupInfo.GroupId != nil {
				groupInfoMap["group_id"] = groupInfo.GroupId
				ids = append(ids, helper.UInt64ToStr(*groupInfo.GroupId))
			}

			if groupInfo.GroupName != nil {
				groupInfoMap["group_name"] = groupInfo.GroupName
			}

			if groupInfo.CreateTime != nil {
				groupInfoMap["create_time"] = groupInfo.CreateTime
			}

			if groupInfo.Remark != nil {
				groupInfoMap["remark"] = groupInfo.Remark
			}

			tmpList = append(tmpList, groupInfoMap)
		}

		_ = d.Set("group_info", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}

	return nil
}
