/*
Use this data source to query detailed information of dcdb shards

# Example Usage

```hcl

	data "tencentcloudenterprise_dcdb_shards" "shards" {
	  instance_id = "your_instance_id"
	  shard_instance_ids = ["shard1_id"]
	}

	data "tencentcloudenterprise_dcdb_shards" "shards" {
	  instance_id = "your_instance_id"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_dcdb_shards", CNDescription{
		TerraformTypeCN: "数据库实例的分片信息",
		DescriptionCN:   "提供DCDB分片数据源，用于查询DCDB分片的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "实例id",
			"shard_instance_ids": "碎片实例ID",
			"list":               "碎片列表",
			"shard_serial_id":    "碎片序列号",
			"shard_instance_id":  "碎片实例id",
			"status":             "地位",
			"status_desc":        "状态描述",
			"create_time":        "创造时间",
			"vpc_id":             "vpc id",
			"subnet_id":          "子网id",
			"project_id":         "项目id",
			"region":             "区域",
			"zone":               "区",
			"memory":             "内存，单位为GB",
			"storage":            "内存，单位为GB",
			"period_end_time":    "过期时间",
			"node_count":         "节点计数",
			"storage_usage":      "存储使用情况",
			"memory_usage":       "内存使用情况",
			"proxy_version":      "代理版本",
			"paymode":            "付费模式",
			"shard_master_zone":  "碎片主区域",
			"shard_slave_zones":  "碎片从属区域",
			"cpu":                "cpu核心",
			"range":              "shard密钥的范围",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudDcdbShards() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudDcdbShardsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance id.",
			},

			"shard_instance_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Shard instance ids.",
			},

			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Shard list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance id.",
						},
						"shard_serial_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Shard serial id.",
						},
						"shard_instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Shard instance id.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Status.",
						},
						"status_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status description.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vpc id.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subnet id.",
						},
						"project_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Project id.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Zone.",
						},
						"memory": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory, the unit is GB.",
						},
						"storage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory, the unit is GB.",
						},
						"period_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Expired time.",
						},
						"node_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Node count.",
						},
						"storage_usage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Storage usage.",
						},
						"memory_usage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory usage.",
						},
						"proxy_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Proxy version.",
						},
						"paymode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Pay mode.",
						},
						"shard_master_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Shard master zone.",
						},
						"shard_slave_zones": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "Shard slave zones.",
						},
						"cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Cpu cores.",
						},
						"range": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The range of shard key.",
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
		Description: " query dcdb shards.",
	}
}

func dataSourceTencentCloudDcdbShardsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_dcdb_shards.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("shard_instance_ids"); ok {
		shard_instance_idsSet := v.(*schema.Set).List()
		ids := make([]*string, 0, len(shard_instance_idsSet))
		for _, vv := range shard_instance_idsSet {
			ids = append(ids, helper.String(vv.(string)))
		}
		paramMap["shard_instance_ids"] = ids
	}

	dcdbService := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	var shards []*dcdb.DCDBShardInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := dcdbService.DescribeDcdbShardsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		shards = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Dcdb shards failed, reason:%+v", logId, err)
		return err
	}

	// shardList := []interface{}{}
	shardList := make([]map[string]interface{}, 0, len(shards))
	ids := make([]string, 0, len(shards))
	if shards != nil {
		for _, shard := range shards {
			shardMap := map[string]interface{}{}
			if shard.InstanceId != nil {
				shardMap["instance_id"] = shard.InstanceId
			}
			if shard.ShardSerialId != nil {
				shardMap["shard_serial_id"] = shard.ShardSerialId
			}
			if shard.ShardInstanceId != nil {
				shardMap["shard_instance_id"] = shard.ShardInstanceId
			}
			if shard.Status != nil {
				shardMap["status"] = shard.Status
			}
			if shard.StatusDesc != nil {
				shardMap["status_desc"] = shard.StatusDesc
			}
			if shard.CreateTime != nil {
				shardMap["create_time"] = shard.CreateTime
			}
			if shard.VpcId != nil {
				shardMap["vpc_id"] = shard.VpcId
			}
			if shard.SubnetId != nil {
				shardMap["subnet_id"] = shard.SubnetId
			}
			if shard.ProjectId != nil {
				shardMap["project_id"] = shard.ProjectId
			}
			if shard.Region != nil {
				shardMap["region"] = shard.Region
			}
			if shard.Zone != nil {
				shardMap["zone"] = shard.Zone
			}
			if shard.Memory != nil {
				shardMap["memory"] = shard.Memory
			}
			if shard.Storage != nil {
				shardMap["storage"] = shard.Storage
			}
			if shard.PeriodEndTime != nil {
				shardMap["period_end_time"] = shard.PeriodEndTime
			}
			if shard.NodeCount != nil {
				shardMap["node_count"] = shard.NodeCount
			}
			if shard.StorageUsage != nil {
				shardMap["storage_usage"] = shard.StorageUsage
			}
			if shard.MemoryUsage != nil {
				shardMap["memory_usage"] = shard.MemoryUsage
			}
			if shard.ProxyVersion != nil {
				shardMap["proxy_version"] = shard.ProxyVersion
			}
			if shard.Paymode != nil {
				shardMap["paymode"] = shard.Paymode
			}
			if shard.ShardMasterZone != nil {
				shardMap["shard_master_zone"] = shard.ShardMasterZone
			}
			if shard.ShardSlaveZones != nil {
				shardMap["shard_slave_zones"] = shard.ShardSlaveZones
			}
			if shard.Cpu != nil {
				shardMap["cpu"] = shard.Cpu
			}
			if shard.Range != nil {
				shardMap["range"] = shard.Range
			}

			ids = append(ids, *shard.ShardInstanceId)
			shardList = append(shardList, shardMap)
		}
		d.SetId(helper.DataResourceIdsHash(ids))
		err = d.Set("list", shardList)
		if err != nil {
			log.Printf("[CRITAL]%s set Dcdb shards failed, reason:%+v", logId, err)
			return err
		}
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), shardList); e != nil {
			return e
		}
	}

	return nil
}
