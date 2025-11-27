/*
Use this data source to query detailed information of instance_node_info

# Example Usage

```hcl

	data "tencentcloudenterprise_redis_instance_node_info" "instance_node_info" {
	  instance_id = "crs-c1nl9rpv"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	redis "terraform-provider-tencentcloudenterprise/sdk/redis/v20180412"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_redis_instance_node_info", CNDescription{
		TerraformTypeCN: "Redis®实例节点信息",
		DescriptionCN:   "提供Redis实例节点信息数据源，用于查询Redis实例节点信息的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "实例的ID",
			"proxy_count":        "代理节点数",
			"proxy":              "代理节点信息",
			"node_id":            "节点ID",
			"zone_id":            "区域ID",
			"redis_count":        "redis®节点的数量",
			"redis":              "Redis®节点信息",
			"node_role":          "节点角色",
			"cluster_id":         "碎片ID",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudRedisInstanceNodeInfo() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of instance_node_info",
		Read:        dataSourceTencentCloudRedisInstanceNodeInfoRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The ID of instance.",
			},

			"proxy_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Number of proxy nodes.",
			},

			"proxy": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Proxy node information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node ID.",
						},
						"zone_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Zone ID.",
						},
					},
				},
			},

			"redis_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Number of redis nodes.",
			},

			"redis": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Redis node information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node ID.",
						},
						"node_role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node role.",
						},
						"cluster_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Shard ID.",
						},
						"zone_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Zone ID.",
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

func dataSourceTencentCloudRedisInstanceNodeInfoRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_redis_instance_node_info.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var instanceId string
	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		paramMap["InstanceId"] = helper.String(v.(string))
	}
	service := RedisService{client: meta.(*TencentCloudClient).apiV3Conn}
	var instanceNodeInfo *redis.DescribeInstanceNodeInfoResponse
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeRedisInstanceNodeInfoByFilter(ctx, paramMap)
		if e != nil {
			if sdkerr, ok := e.(*sdkErrors.CloudSDKError); ok {
				if sdkerr.Code == "FailedOperation.SystemError" {
					return nil
				}
			}
			return retryError(e)
		}
		instanceNodeInfo = result
		return nil
	})
	if err != nil {
		return err
	}
	resp := instanceNodeInfo.Response

	if resp.ProxyCount != nil {
		_ = d.Set("proxy_count", resp.ProxyCount)
	}

	if resp.Proxy != nil {
		tmpList := make([]map[string]interface{}, 0, len(resp.Proxy))
		for _, proxyNodes := range resp.Proxy {
			proxyNodesMap := map[string]interface{}{}

			if proxyNodes.NodeId != nil {
				proxyNodesMap["node_id"] = proxyNodes.NodeId
			}

			//if proxyNodes.ZoneId != nil {
			//	proxyNodesMap["zone_id"] = proxyNodes.ZoneId
			//}

			tmpList = append(tmpList, proxyNodesMap)
		}

		_ = d.Set("proxy", tmpList)
	}
	if resp.RedisCount != nil {
		_ = d.Set("redis_count", resp.RedisCount)
	}

	if resp.Redis != nil {
		tmpList := make([]map[string]interface{}, 0, len(resp.Redis))
		for _, redisNodes := range resp.Redis {
			redisNodesMap := map[string]interface{}{}

			if redisNodes.NodeId != nil {
				redisNodesMap["node_id"] = redisNodes.NodeId
			}

			if redisNodes.NodeRole != nil {
				redisNodesMap["node_role"] = redisNodes.NodeRole
			}

			if redisNodes.ClusterId != nil {
				redisNodesMap["cluster_id"] = redisNodes.ClusterId
			}

			if redisNodes.ZoneId != nil {
				redisNodesMap["zone_id"] = redisNodes.ZoneId
			}

			tmpList = append(tmpList, redisNodesMap)
		}

		_ = d.Set("redis", tmpList)
	}

	d.SetId(instanceId)
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		result := make(map[string]interface{})
		result["instance_id"] = instanceId
		result["proxy_count"] = resp.ProxyCount
		result["proxy"] = resp.Proxy
		result["redis_count"] = resp.RedisCount
		result["redis"] = resp.Redis
		if e := writeToFile(output.(string), result); e != nil {
			log.Printf("[ERROR] write to file error: %v", e)
			return e
		}
	}
	return nil
}
