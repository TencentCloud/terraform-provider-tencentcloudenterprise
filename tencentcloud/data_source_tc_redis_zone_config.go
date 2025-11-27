/*
Use this data source to query which instance types of Redis® are available in a specific region.

# Example Usage

```hcl

	data "tencentcloudenterprise_redis_zone_config" "redislab" {
	  region             = "ap-hongkong"
	  result_output_file = "/temp/mytestpath"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	redis "terraform-provider-tencentcloudenterprise/sdk/redis/v20180412"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_redis_zone_config", CNDescription{
		TerraformTypeCN: "Redis®可用区配置",
		DescriptionCN:   "提供Redis可用区配置数据源，用于查询Redis可用区配置的详细信息。",
		AttributesCN: map[string]string{
			"region":              "区域的名称如果未设置此值，则将使用从提供程序配置中获取的当前区域",
			"type_id":             "实例类型ID",
			"result_output_file":  "结果输出文件",
			"list":                "区域列表每个元素都包含以下属性：",
			"zone":                "可用区域的ID",
			"type":                "实例类型可用值：“master_slave_redis”、“master_slave_ckv'”、“cluster_ckv'、“cluster_redis’”和“standalone_redis‘”",
			"version":             "可用实例的版本描述可能的值：“Redis 3.2”、“Redis 4.0”",
			"mem_sizes":           "可用实例的内存容量（MB）",
			"shard_memories":      "可用实例碎片的内存卷列表（以MB为单位）",
			"redis_shard_nums":    "实例shard的支持号",
			"redis_replicas_nums": "实例副本的支持数量",
		},
	})
}

func dataSourceTencentRedisZoneConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query which instance types of Redis® are available in a specific region.",
		Read:        dataSourceTencentRedisZoneConfigRead,
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of a region. If this value is not set, the current region getting from provider's configuration will be used.",
			},
			"type_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerMin(2),
				Description:  "Instance type ID.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			// Computed values
			"list": {Type: schema.TypeList,
				Description: "A list of zone. Each element contains the following attributes:",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of available zone.",
						},
						"type_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance type. Which redis type supports in this zone.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Deprecated:  "It has been deprecated from version 1.33.1. Please use 'type_id' instead.",
							Description: "Instance type. Available values: `master_slave_redis`, `master_slave_ckv`, `cluster_ckv`, `cluster_redis` and `standalone_redis`.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version description of an available instance. Possible values: `Redis 3.2`, `Redis 4.0`.",
						},
						"mem_sizes": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Computed:    true,
							Deprecated:  "It has been deprecated from version 1.26.0. Use `shard_memories` instead.",
							Description: "The memory volume of an available instance(in MB).",
						},
						"shard_memories": {
							Type:        schema.TypeList,
							Description: "The memory volume list of an available instance shard(in MB).",
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
						"redis_shard_nums": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Description: "The support numbers of instance shard.",
						},
						"redis_replicas_nums": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Description: "The support numbers of instance copies.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentRedisZoneConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_redis_zone_config.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := RedisService{client: meta.(*TencentCloudClient).apiV3Conn}
	region := meta.(*TencentCloudClient).apiV3Conn.Region

	if regionInterface, ok := d.GetOk("region"); ok {
		region = regionInterface.(string)
	} else {
		log.Printf("[INFO]%s region is not set,so we use [%s] from env\n ", logId, region)
	}

	typeId := int64(d.Get("type_id").(int))

	sellConfigures, err := service.DescribeRedisZoneConfig(ctx)
	if err != nil {
		return fmt.Errorf("api[DescribeRedisZoneConfig]fail, return %s", err.Error())
	}

	var regionItem *redis.RegionConf
	for _, regionItem = range sellConfigures {
		if *regionItem.RegionId == region {
			break
		}
	}
	if regionItem == nil {
		return nil
	}
	var allZonesConfigs []interface{}

	for _, zones := range regionItem.ZoneSet {
		zoneName := *zones.ZoneId

		for _, products := range zones.ProductSet {
			//1- package year and month, 0- billing according to quantity.
			if *products.PayMode != "0" {
				continue
			}

			if typeId != 0 && typeId != *products.Type {
				continue
			}

			zoneConfigures := map[string]interface{}{}
			zoneConfigures["zone"] = zoneName
			zoneConfigures["version"] = *products.Version
			zoneConfigures["type_id"] = products.Type
			//this products sale out.
			if *products.Saleout {
				continue
			}

			memSizes := make([]int64, 0, len(products.TotalSize))

			for _, size := range products.TotalSize {
				temp, err := strconv.ParseInt(*size, 10, 64)
				if err != nil {
					continue
				}
				memSizes = append(memSizes, temp*1024)
			}

			shardMemories := make([]int64, 0, len(products.ShardSize))
			for _, size := range products.ShardSize {
				temp, err := strconv.ParseInt(*size, 10, 64)
				if err != nil {
					continue
				}
				shardMemories = append(shardMemories, temp*1024)
			}

			zoneConfigures["mem_sizes"] = memSizes
			zoneConfigures["shard_memories"] = shardMemories
			zoneConfigures["type"] = REDIS_NAMES[*products.Type]

			var redisShardNums []int64
			var redisReplicasNums []int64

			for _, v := range products.ShardNum {
				int64Value, err := strconv.ParseInt(*v, 10, 64)
				if err != nil {
					return fmt.Errorf("api[DescribeRedisZoneConfig]return error `redis_shard_nums`,%s", err.Error())
				}
				redisShardNums = append(redisShardNums, int64Value)
			}
			zoneConfigures["redis_shard_nums"] = redisShardNums

			for _, v := range products.ReplicaNum {
				int64Value, err := strconv.ParseInt(*v, 10, 64)
				if err != nil {
					return fmt.Errorf("api[DescribeRedisZoneConfig]return error `redis_replicas_nums`,%s", err.Error())
				}
				redisReplicasNums = append(redisReplicasNums, int64Value)
			}
			zoneConfigures["redis_replicas_nums"] = redisReplicasNums

			allZonesConfigs = append(allZonesConfigs, zoneConfigures)
		}
	}

	if err := d.Set("list", allZonesConfigs); err != nil {
		log.Printf("[CRITAL]%s provider set  redis zoneConfigs fail, reason:%s\n ", logId, err.Error())
		return err
	}

	id := "redis_zoneconfig" + region
	if typeId != 0 {
		id += fmt.Sprintf("%d", typeId)
	}
	d.SetId(id)

	if output, ok := d.GetOk("result_output_file"); ok && output.(string) != "" {

		if err := writeToFile(output.(string), allZonesConfigs); err != nil {
			log.Printf("[CRITAL]%s output file[%s] fail, reason[%s]\n",
				logId, output.(string), err.Error())
			return err
		}

	}
	return nil
}
