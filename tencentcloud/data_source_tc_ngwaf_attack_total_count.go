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

func dataSourceTencentCloudNgwafAttackTotalCount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudNgwafAttackTotalCountRead,
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

			"query_string": {
				Optional:    true,
				Type:        schema.TypeString,
				Default:     "",
				Description: "Query conditions.",
			},
			"total_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Total number of attacks.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudNgwafAttackTotalCountRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_ngwaf_attack_total_count.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId            = getLogId(contextNil)
		ctx              = context.WithValue(context.TODO(), logIdKey, logId)
		service          = &NgwafService{client: meta.(*TencentCloudClient)}
		attackTotalCount *ngwaf.GetAttackTotalCountResponse
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("start_time"); ok {
		paramMap["StartTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_time"); ok {
		paramMap["EndTime"] = helper.String(v.(string))
	}

	// Domain is always set to "all" to query all domains
	paramMap["Domain"] = helper.String("all")

	if v, ok := d.GetOkExists("query_string"); ok {
		paramMap["QueryString"] = helper.String(v.(string))
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeWafAttackTotalCountByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		attackTotalCount = result
		return nil
	})

	if err != nil {
		return err
	}

	if attackTotalCount != nil && attackTotalCount.Response != nil && attackTotalCount.Response.TotalCount != nil {
		_ = d.Set("total_count", attackTotalCount.Response.TotalCount)
	}

	// 构造要写入文件的数据映射
	attackTotalCountMap := map[string]interface{}{
		"total_count": 0,
	}
	if attackTotalCount != nil && attackTotalCount.Response != nil && attackTotalCount.Response.TotalCount != nil {
		attackTotalCountMap["total_count"] = *attackTotalCount.Response.TotalCount
	}
	// 添加查询参数
	if v, ok := d.GetOk("start_time"); ok {
		attackTotalCountMap["start_time"] = v.(string)
	}
	if v, ok := d.GetOk("end_time"); ok {
		attackTotalCountMap["end_time"] = v.(string)
	}
	attackTotalCountMap["domain"] = "all"
	if v, ok := d.GetOk("query_string"); ok {
		attackTotalCountMap["query_string"] = v.(string)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), attackTotalCountMap); e != nil {
			return e
		}
	}

	return nil
}
