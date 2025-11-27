/*
Use this data source to query detailed information of ckafka group

# Example Usage

```hcl

	data "tencentcloudenterprise_ckafka_group" "group" {
	  instance_id = "ckafka-xxxxxxx"
	  search_word = "xxxxxx"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_group", CNDescription{
		TerraformTypeCN: "ckafka 消费分组",
		DescriptionCN:   "提供Ckafka消费组数据源，用于查询Ckafka消费组的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "实例Id",
			"search_word":        "搜索关键字",
			"group_list":         "GroupList",
			"group":              "groupId",
			"protocol":           "该 group 使用的协议。",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudCkafkaGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of ckafka group",
		Read:        dataSourceTencentCloudCkafkaGroupRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "InstanceId.",
			},

			"search_word": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Search for the keyword.",
			},

			"group_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "GroupList.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "GroupId.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol used by this group.",
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

func dataSourceTencentCloudCkafkaGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ckafka_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("search_word"); ok {
		paramMap["search_word"] = helper.String(v.(string))
	}

	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}

	var groups []*ckafka.DescribeGroup

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCkafkaGroupByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		groups = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(groups))
	groupMapList := []interface{}{}
	for _, group := range groups {
		groupMap := map[string]interface{}{}

		if group.Group != nil {
			groupMap["group"] = group.Group
			ids = append(ids, *group.Group)
		}

		if group.Protocol != nil {
			groupMap["protocol"] = group.Protocol
		}

		groupMapList = append(groupMapList, groupMap)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("group_list", groupMapList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), groupMapList); e != nil {
			return e
		}
	}
	return nil
}
