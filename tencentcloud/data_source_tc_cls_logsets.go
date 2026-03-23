/*
Use this data source to query detailed information of cls logsets.

# Example Usage

# Query all cls logsets

```hcl
data "tencentcloudenterprise_cls_logsets" "logsets" {}
```

# Query by filters

```hcl
# Query by `logsetName`

	data "tencentcloudenterprise_cls_logsets" "logsets" {
	  filters {
	    key    = "logsetName"
	    values = ["cls_service_logging"]
	  }
	}

# Query by `logsetId`

	data "tencentcloudenterprise_cls_logsets" "logsets" {
	  filters {
	    key    = "logsetId"
	    values = ["50d499a8-c4c0-4442-aa04-e8aa8a02437d"]
	  }
	}

# Query by `tagKey`

	data "tencentcloudenterprise_cls_logsets" "logsets" {
	  filters {
	    key    = "tagKey"
	    values = ["createdBy"]
	  }
	}

# Query by `tag:tagKey`

	data "tencentcloudenterprise_cls_logsets" "logsets" {
	  filters {
	    key    = "tag:createdBy"
	    values = ["terraform"]
	  }
	}

```
*/
package tencentcloud

import (
	"log"

	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTencentCloudClsLogsets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudClsLogsetsRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Optional:    true,
				MaxItems:    10,
				Type:        schema.TypeList,
				Description: "Query by filter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Fields that need to be filtered. Support: `logsetName`, `logsetId`, `tagKey`, `tag:tagKey`.",
						},
						"values": {
							Type:        schema.TypeSet,
							Required:    true,
							MaxItems:    5,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The values that need to be filtered.",
						},
					},
				},
			},
			// computed
			"logsets": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "logset lists.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logset_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Logset Id.",
						},
						"logset_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Logset name.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"assumer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud product identification, when the log set is created by another cloud product, this field will display the cloud product name, such as CDN, TKE.",
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Tags.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag key.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag value.",
									},
								},
							},
						},
						"topic_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Topic count.",
						},
						"role_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "If `assumer_name` is not empty, it indicates the service role that created the log set.",
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

func dataSourceTencentCloudClsLogsetsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cls_logsets.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cls.NewDescribeLogsetsRequest()

	if v, ok := d.GetOk("filters"); ok {
		filtersSet := v.([]interface{})
		tmpSet := make([]*cls.Filter, 0, len(filtersSet))
		for _, item := range filtersSet {
			filter := cls.Filter{}
			filterMap := item.(map[string]interface{})
			if v, ok := filterMap["key"]; ok {
				filter.Key = helper.String(v.(string))
			}

			if v, ok := filterMap["values"]; ok {
				valuesSet := v.(*schema.Set).List()
				filter.Values = helper.InterfacesStringsPoint(valuesSet)
			}

			tmpSet = append(tmpSet, &filter)
		}
		request.Filters = tmpSet
	}

	var logsets []*cls.LogsetInfo
	var offset int64 = 0
	var pageSize int64 = 50

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		for {
			request.Offset = &offset
			request.Limit = &pageSize

			response, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().DescribeLogsets(request)
			if e != nil {
				return retryError(e)
			}

			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

			if response == nil || response.Response == nil || len(response.Response.Logsets) == 0 {
				break
			}

			logsets = append(logsets, response.Response.Logsets...)

			if len(response.Response.Logsets) < int(pageSize) {
				break
			}
			offset += pageSize
		}
		return nil
	})

	if err != nil {
		return err
	}

	ids := make([]string, 0, len(logsets))
	tmpList := make([]map[string]interface{}, 0, len(logsets))

	for _, logsetInfo := range logsets {
		logsetInfoMap := map[string]interface{}{}

		if logsetInfo.LogsetId != nil {
			logsetInfoMap["logset_id"] = logsetInfo.LogsetId
			ids = append(ids, *logsetInfo.LogsetId)
		}

		if logsetInfo.LogsetName != nil {
			logsetInfoMap["logset_name"] = logsetInfo.LogsetName
		}

		if logsetInfo.CreateTime != nil {
			logsetInfoMap["create_time"] = logsetInfo.CreateTime
		}

		if logsetInfo.AssumerName != nil {
			logsetInfoMap["assumer_name"] = logsetInfo.AssumerName
		}

		if logsetInfo.Tags != nil {
			tagsList := []interface{}{}
			for _, tags := range logsetInfo.Tags {
				tagsMap := map[string]interface{}{}
				if tags.Key != nil {
					tagsMap["key"] = tags.Key
				}

				if tags.Value != nil {
					tagsMap["value"] = tags.Value
				}

				tagsList = append(tagsList, tagsMap)
			}
			logsetInfoMap["tags"] = tagsList
		}

		if logsetInfo.TopicCount != nil {
			logsetInfoMap["topic_count"] = logsetInfo.TopicCount
		}

		if logsetInfo.RoleName != nil {
			logsetInfoMap["role_name"] = logsetInfo.RoleName
		}

		tmpList = append(tmpList, logsetInfoMap)
	}

	_ = d.Set("logsets", tmpList)

	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}

	return nil
}
