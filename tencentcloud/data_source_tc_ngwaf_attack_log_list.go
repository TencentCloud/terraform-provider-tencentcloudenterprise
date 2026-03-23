/*
Use this data source to query detailed attack logs from NGWAF (Next-Generation Web Application Firewall).

# Example Usage

```hcl

	data "tencentcloudenterprise_ngwaf_attack_log_list" "example" {
	  start_time   = "2023-01-01 00:00:00"
	  end_time     = "2023-01-31 23:59:59"
	  query_string = "action:block"
	  query_count  = 20
	  page         = 0
	  sort         = "desc"
	}

```
*/
package tencentcloud

import (
	"context"
	"strconv"
	"time"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTencentCloudNgwafAttackLogList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudNgwafAttackLogListRead,
		Schema: map[string]*schema.Schema{

			"start_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Begin time.",
			},
			"end_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "End time.",
			},
			"query_count": {
				Optional:    true,
				Type:        schema.TypeInt,
				Default:     10,
				Description: "Number of queries, default to 10, maximum of 100.",
			},
			"page": {
				Optional:    true,
				Type:        schema.TypeInt,
				Default:     0,
				Description: "Number of pages, starting from 0 by default.",
			},
			"query_string": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Lucene grammar.",
			},
			"sort": {
				Optional:    true,
				Type:        schema.TypeString,
				Default:     "desc",
				Description: "Default desc, support desc, asc.",
			},
			"data": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Attack log array.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The detail of attack log.",
						},
						"file_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Useless.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Useless.",
						},
						"time_stamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time string.",
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

func dataSourceTencentCloudNgwafAttackLogListRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_ngwaf_attack_log_list.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId         = getLogId(contextNil)
		ctx           = context.WithValue(context.TODO(), logIdKey, logId)
		service       = &NgwafService{client: meta.(*TencentCloudClient)}
		attackLogList []*ngwaf.AttackLogInfo
	)

	paramMap := make(map[string]interface{})
	// Domain is always set to "all" to query all domains
	paramMap["Domain"] = helper.String("all")

	if v, ok := d.GetOk("start_time"); ok {
		paramMap["StartTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_time"); ok {
		paramMap["EndTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("query_count"); ok {
		paramMap["Count"] = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("query_string"); ok {
		paramMap["QueryString"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("sort"); ok {
		paramMap["Sort"] = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("page"); ok {
		paramMap["Page"] = helper.IntInt64(v.(int))
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeWafAttackLogListByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		attackLogList = result
		return nil
	})

	if err != nil {
		return err
	}

	tmpList := make([]map[string]interface{}, 0, len(attackLogList))

	if attackLogList != nil {
		for _, attackLogInfo := range attackLogList {
			attackLogInfoMap := map[string]interface{}{}

			if attackLogInfo.Content != nil {
				attackLogInfoMap["content"] = attackLogInfo.Content
			}

			if attackLogInfo.FileName != nil {
				attackLogInfoMap["file_name"] = attackLogInfo.FileName
			}

			if attackLogInfo.Source != nil {
				attackLogInfoMap["source"] = attackLogInfo.Source
			}

			if attackLogInfo.TimeStamp != nil {
				attackLogInfoMap["time_stamp"] = attackLogInfo.TimeStamp
			}

			tmpList = append(tmpList, attackLogInfoMap)
		}

		_ = d.Set("data", tmpList)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}

	return nil
}
