/*
Use this data source to query detailed information of tdmqRocketmq namespace

# Example Usage

```hcl

	resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
		cluster_name = "test_rocketmq_namespace_sdatasource"
		remark = "test recket mq"
	}

	resource "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespacedata" {
		cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
		namespace_name = "test_namespace_datasource"
		ttl = 65000
		retention_time = 65000
		remark = "test namespace"
	}

	data "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespace" {
		cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
		name_keyword = cloud_tdmq_rocketmq_namespace.namespacedata.namespace_name
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmqRocketmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rocketmq_namespace", CNDescription{
		TerraformTypeCN: "RocketMQ 命名空间",
		DescriptionCN:   "提供TDMQ RocketMQ命名空间数据源，用于查询TDMQ RocketMQ命名空间的详细信息。",
		AttributesCN: map[string]string{
			"cluster_id":         "群集ID",
			"name_keyword":       "按名称搜索",
			"namespaces":         "命名空间列表",
			"namespace_id":       "命名空间名称，可以包含3-64个字母、数字、连字符和下划线",
			"ttl":                "未使用邮件的保留时间（以毫秒为单位）取值范围：60秒-15天",
			"retention_time":     "持久化消息的保留时间（以毫秒为单位）",
			"remark":             "备注（最多128个字符）",
			"public_endpoint":    "公共网络接入点地址",
			"vpc_endpoint":       "VPC接入点地址",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudTdmqRocketmqNamespace() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tdmqRocketmq namespace",
		Read:        dataSourceTencentCloudTdmqRocketmqNamespaceRead,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster ID.",
			},

			"name_keyword": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search by name.",
			},

			"namespaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of namespaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Namespace name, which can contain 3-64 letters, digits, hyphens, and underscores.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Retention time of unconsumed messages in milliseconds. Value range: 60 seconds-15 days.",
						},
						"retention_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Retention time of persisted messages in milliseconds.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remarks (up to 128 characters).",
						},
						"public_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Public network access point address.",
						},
						"vpc_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC access point address.",
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

func dataSourceTencentCloudTdmqRocketmqNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmqRocketmq_namespace.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("cluster_id"); ok {
		paramMap["cluster_id"] = v.(string)
	}

	if v, ok := d.GetOk("name_keyword"); ok {
		paramMap["name_keyword"] = v.(string)
	}

	tdmqRocketmqService := TdmqRocketmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	ids := make([]string, 0)
	var namespaces []*tdmqRocketmq.RocketMQNamespace
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := tdmqRocketmqService.DescribeTdmqRocketmqNamespaceByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		namespaces = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read TdmqRocketmq namespaces failed, reason:%+v", logId, err)
		return err
	}

	namespaceList := []interface{}{}

	for _, namespace := range namespaces {
		namespaceMap := map[string]interface{}{}
		ids = append(ids, *namespace.NamespaceId)
		namespaceMap["namespace_id"] = namespace.NamespaceId
		if namespace.Ttl != nil {
			namespaceMap["ttl"] = namespace.Ttl
		}
		if namespace.RetentionTime != nil {
			namespaceMap["retention_time"] = namespace.RetentionTime
		}
		if namespace.Remark != nil {
			namespaceMap["remark"] = namespace.Remark
		}
		if namespace.PublicEndpoint != nil {
			namespaceMap["public_endpoint"] = namespace.PublicEndpoint
		}
		if namespace.VpcEndpoint != nil {
			namespaceMap["vpc_endpoint"] = namespace.VpcEndpoint
		}

		namespaceList = append(namespaceList, namespaceMap)
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("namespaces", namespaceList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), namespaceList); e != nil {
			return e
		}
	}

	return nil
}
