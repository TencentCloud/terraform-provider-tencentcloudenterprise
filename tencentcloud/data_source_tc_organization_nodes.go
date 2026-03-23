/*
Use this data source to query organization nodes.

# Example Usage

```hcl

	data "tencentcloudenterprise_organization_nodes" "example" {
	}

```
*/
package tencentcloud

import (
	"log"
	"strconv"

	organization "terraform-provider-tencentcloudenterprise/sdk/organization/v20220508"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTencentCloudOrganizationNodes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudOrganizationNodesRead,
		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Organization node list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Organization node ID.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Organization node name.",
						},
						"parent_node_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Parent node ID.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remarks.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time.",
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

func dataSourceTencentCloudOrganizationNodesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_organization_nodes.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		offset   int64 = 0
		pageSize int64 = 50
		nodes    []*organization.OrgNode
	)

	request := organization.NewDescribeOrganizationNodesRequest()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		for {
			request.Offset = &offset
			request.Limit = &pageSize

			response, e := meta.(*TencentCloudClient).apiV3Conn.UseOrganizationClient().DescribeOrganizationNodes(request)
			if e != nil {
				return retryError(e)
			}

			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

			if response == nil || response.Response == nil || len(response.Response.Items) == 0 {
				break
			}

			nodes = append(nodes, response.Response.Items...)

			if len(response.Response.Items) < int(pageSize) {
				break
			}
			offset += pageSize
		}
		return nil
	})

	if err != nil {
		return err
	}

	itemList := make([]map[string]interface{}, 0, len(nodes))
	ids := make([]string, 0, len(nodes))

	for _, item := range nodes {
		itemMap := map[string]interface{}{}

		if item.NodeId != nil {
			itemMap["node_id"] = item.NodeId
			ids = append(ids, strconv.FormatInt(*item.NodeId, 10))
		}

		if item.Name != nil {
			itemMap["name"] = item.Name
		}

		if item.ParentNodeId != nil {
			itemMap["parent_node_id"] = item.ParentNodeId
		}

		if item.Remark != nil {
			itemMap["remark"] = item.Remark
		}

		if item.CreateTime != nil {
			itemMap["create_time"] = item.CreateTime
		}

		if item.UpdateTime != nil {
			itemMap["update_time"] = item.UpdateTime
		}

		itemList = append(itemList, itemMap)
	}

	_ = d.Set("items", itemList)

	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), itemList); e != nil {
			return e
		}
	}

	return nil
}
