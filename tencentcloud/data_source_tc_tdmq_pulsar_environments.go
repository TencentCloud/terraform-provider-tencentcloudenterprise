/*
Use this data source to query detailed information of tdmq environment_attributes

# Example Usage

```hcl

	data "tencentcloudenterprise_tdmq_environment_attributes" "environment_attributes" {
	    environment_id = "keep-ns"
	    cluster_id     = "pulsar-9n95ax58b9vn"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_pulsar_environments", CNDescription{
		TerraformTypeCN: "命名空间列表",
		DescriptionCN:   "提供TDMQ Pulsar环境数据源，用于查询TDMQ Pulsar环境的详细信息。",
		AttributesCN: map[string]string{
			"environment_id":     "命名空间名称，模糊搜索",
			"cluster_id":         "Pulsar 集群的ID",
			"result_output_file": "输出文件",
			"list":               "输出列表",
			"remark":             "备注，128个字符以内",
			"msg_ttl":            "未使用消息的过期时间，单位秒，最长1296000（15天）",
			"create_time":        "创建时间",
			"update_time":        "上次修改时间",
			"namespace_id":       "命名空间ID",
			"namespace_name":     "命名空间名称",
			"topic_num":          "主题数量",
			"retention_policy":   "要保留的邮件的策略格式类似：`{time_in_minutes:Int，size_in_mb:Int}``time_in_minutes`：保留消息的时间`size_in_mb`：要保留的消息的大小",
		},
	})
}

func dataSourceTencentCloudTdmqPulsarEnvironments() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tdmq environment_attributes",
		Read:        dataSourceTencentCloudTdmqPulsarEnvironmentsRead,
		Schema: map[string]*schema.Schema{
			// 命名空间名称，模糊搜索
			"environment_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Environment (namespace) name.",
			},
			// Pulsar 集群的ID
			"cluster_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "ID of the Pulsar cluster.",
			},
			// 输出文件
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			// Computed
			// 输出列表
			"list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "List of environments.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// 命名空间名称，模糊搜索
						"environment_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Environment (namespace) name.",
						},
						// 说明
						"remark": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Remarks, within 128 characters.",
						},
						// 未消费消息过期时间，单位：秒，最大1296000（15天）
						"msg_ttl": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "Expiration time of unconsumed messages, unit second, maximum 1296000 (15 days).",
						},
						// 创建时间
						"create_time": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Creation time.",
						},
						// 最近修改时间
						"update_time": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Last modification time.",
						},
						// 命名空间ID
						"namespace_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Namespace ID.",
						},
						// 命名空间名称
						"namespace_name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Namespace name.",
						},
						// Topic数量
						"topic_num": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "Number of topics.",
						},
						// 消息保留策略
						"retention_policy": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Optional:    true,
							Description: "The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `time_in_minutes`: the time of message to retain; `size_in_mb`: the size of message to retain.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudTdmqPulsarEnvironmentsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmq_pulsar_environments.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId        = getLogId(contextNil)
		ctx          = context.WithValue(context.TODO(), logIdKey, logId)
		service      = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		environments []*tdmq.Environment
		request      = tdmq.NewDescribeEnvironmentsRequest()
	)

	request.EnvironmentId = helper.String(d.Get("environment_id").(string))

	if v, ok := d.GetOk("cluster_id"); ok {
		request.ClusterId = helper.String(v.(string))
	}

	err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeEnvironments(ctx, request)
		if e != nil {
			return retryError(e)
		}
		environments = result
		return nil
	})

	if err != nil {
		return err
	}

	ids := make([]string, 0)
	list := make([]map[string]interface{}, 0, len(environments))
	for _, v := range environments {
		m := map[string]interface{}{
			"environment_id": *v.EnvironmentId,
			"remark":         *v.Remark,
			"msg_ttl":        *v.MsgTTL,
			"create_time":    *v.CreateTime,
			"update_time":    *v.UpdateTime,
			"namespace_id":   *v.NamespaceId,
			"namespace_name": *v.NamespaceName,
			"topic_num":      *v.TopicNum,
			"retention_policy": map[string]interface{}{
				"time_in_minutes": *v.RetentionPolicy.TimeInMinutes,
				"size_in_mb":      *v.RetentionPolicy.SizeInMB,
			},
		}
		list = append(list, m)
		ids = append(ids, *v.EnvironmentId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), list); e != nil {
			return e
		}
	}

	return nil
}
