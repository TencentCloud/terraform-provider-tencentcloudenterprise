/*
Use this data source to query detailed information of dcdb instances

# Example Usage

```hcl

	data "tencentcloudenterprise_dcdb_instances" "instances1" {
	  instance_ids = ["your_dcdb_instance1_id"]
	  search_name = "instancename"
	  search_key = "search_key"
	  project_ids = [0]
	  excluster_type = 0
	  is_filter_excluster = true
	  excluster_type = 0
	  is_filter_vpc = true
	  vpc_id = "your_vpc_id"
	  subnet_id = "your_subnet_id"
	}

	data "tencentcloudenterprise_dcdb_instances" "instances2" {
	  instance_ids = ["your_dcdb_instance2_id"]
	}

	data "tencentcloudenterprise_dcdb_instances" "instances3" {
	  search_name = "instancename"
	  search_key = "instances3"
	  is_filter_excluster = false
	  excluster_type = 2
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
	registerDataDescriptionProvider("tencentcloudenterprise_dcdb_instances", CNDescription{
		TerraformTypeCN: "数据库实例",
		DescriptionCN:   "提供DCDB实例数据源，用于查询DCDB实例的详细信息。",
		AttributesCN: map[string]string{
			"instance_ids":        "实例ID",
			"search_name":         "搜索名称，支持instancename，vip，all",
			"search_key":          "搜索关键字，支持模糊查询",
			"project_ids":         "项目ID",
			"excluster_type":      "簇外型",
			"is_filter_excluster": "根据聚类排除器类型进行搜索",
			"is_filter_vpc":       "根据vpc进行搜索",
			"vpc_id":              "vpc id，当IsFilterVpc为true时有效",
			"subnet_id":           "子网id，当IsFilterVpc为true时有效",
			"list":                "实例列表",
			"instance_id":         "实例id",
			"instance_name":       "实例名称",
			"app_id":              "应用程序id",
			"project_id":          "项目id",
			"region":              "区域",
			"status_desc":         "状态描述",
			"status":              "地位",
			"vip":                 "vip",
			"vport":               "vport",
			"create_time":         "创造时间",
			"auto_renew_flag":     "自动续订标志",
			"memory":              "内存，单位为GB",
			"storage":             "内存，单位为GB",
			"shard_count":         "碎片计数",
			"period_end_time":     "过期时间",
			"isolated_timestamp":  "隔离时间",
			"uin":                 "账户uin",
			"shard_detail":        "碎片细节",
			"shard_instance_id":   "碎片实例id",
			"shard_serial_id":     "碎片序列号",
			"createtime":          "碎片创建时间",
			"shard_id":            "碎片id",
			"node_count":          "节点计数",
			"cpu":                 "cpu核心",
			"is_tmp":              "tmp实例标记",
			"wan_domain":          "wan域",
			"wan_vip":             "wan vip",
			"wan_port":            "湾港口",
			"update_time":         "更新时间",
			"db_engine":           "db引擎",
			"db_version":          "db引擎版本",
			"paymode":             "付费模式",
			"wan_status":          "wan状态，0:未激活，1:激活，2:关闭，3:激活",
			"is_audit_supported":  "aduit支持，0:支持，1:不支持",
			"instance_type":       "实例类型",
			"resource_tags":       "资源标签",
			"tag_key":             "标记键",
			"tag_标记值":             "标记值",
			"result_output_file":  "用于保存结果",
		},
	})
}

func dataSourceTencentCloudDcdbInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudDcdbInstancesRead,
		Schema: map[string]*schema.Schema{
			"instance_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Instance ids.",
			},

			"search_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search name, support instancename, vip, all.",
			},

			"search_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search key, support fuzzy query.",
			},

			"project_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional:    true,
				Description: "Project ids.",
			},

			"excluster_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Cluster excluster type.",
			},

			"is_filter_excluster": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Search according to the cluster excluter type.",
			},

			"is_filter_vpc": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Search according to the vpc.",
			},

			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Vpc id, valid when IsFilterVpc is true.",
			},

			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Subnet id, valid when IsFilterVpc is true.",
			},

			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Instance list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance id.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance name.",
						},
						"app_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "App id.",
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
						"vpc_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Vpc id.",
						},
						"subnet_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Subnet id.",
						},
						"status_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status description.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Status.",
						},
						"vip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vip.",
						},
						"vport": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Vport.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"auto_renew_flag": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Auto renew flag.",
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
						"shard_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Shard count.",
						},
						"period_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Expired time.",
						},
						"isolated_timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Isolated time.",
						},
						"uin": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account uin.",
						},
						"shard_detail": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Shard detail.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"shard_instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Shard instance id.",
									},
									"shard_serial_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Shard serial id.",
									},
									"status": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Shard status.",
									},
									"createtime": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Shard create time.",
									},
									"memory": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Memory.",
									},
									"storage": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Storage.",
									},
									"shard_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Shard id.",
									},
									"node_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Node count.",
									},
									"cpu": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Cpu cores.",
									},
								},
							},
						},
						"node_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Node count.",
						},
						"is_tmp": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Tmp instance mark.",
						},
						"wan_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Wan domain.",
						},
						"wan_vip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Wan vip.",
						},
						"wan_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Wan port.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time.",
						},
						"db_engine": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Db engine.",
						},
						"db_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Db engine version.",
						},
						"paymode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Pay mode.",
						},
						"wan_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Wan status, 0:nonactivated, 1:activated, 2:closed, 3:activating.",
						},
						"is_audit_supported": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Aduit support, 0:support, 1:unsupport.",
						},
						"instance_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance type.",
						},
						"resource_tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Resource tags.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag key.",
									},
									"tag_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag value.",
									},
								},
							},
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
		Description: "Query dcdb instance list.",
	}
}

func dataSourceTencentCloudDcdbInstancesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_dcdb_hour_instances.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_ids"); ok {
		instance_idsSet := v.(*schema.Set).List()
		ids := make([]*string, 0, len(instance_idsSet))
		for _, vv := range instance_idsSet {
			ids = append(ids, helper.String(vv.(string)))
		}
		paramMap["instance_ids"] = ids
	}

	if v, ok := d.GetOk("search_name"); ok {
		paramMap["search_name"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("search_key"); ok {
		paramMap["search_key"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("project_ids"); ok {
		project_idsSet := v.(*schema.Set).List()
		ids := make([]*int64, 0, len(project_idsSet))
		for _, vv := range project_idsSet {
			ids = append(ids, helper.IntInt64(vv.(int)))
		}
		paramMap["project_ids"] = ids
	}

	if v, ok := d.GetOk("excluster_type"); ok {
		paramMap["excluster_type"] = helper.IntInt64(v.(int))
	}

	if v, _ := d.GetOk("is_filter_excluster"); v != nil {
		paramMap["is_filter_excluster"] = helper.Bool(v.(bool))
	}

	if v, _ := d.GetOk("is_filter_vpc"); v != nil {
		paramMap["is_filter_vpc"] = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		paramMap["vpc_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		paramMap["subnet_id"] = helper.String(v.(string))
	}

	dcdbService := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	var instances []*dcdb.DCDBInstanceInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := dcdbService.DescribeDcdbInstancesByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		instances = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Dcdb instances failed, reason:%+v", logId, err)
		return err
	}

	ids := make([]string, 0, len(instances))
	instanceList := make([]map[string]interface{}, 0, len(instances))
	if instances != nil {
		for _, instance := range instances {
			instanceMap := map[string]interface{}{}
			if instance.InstanceId != nil {
				instanceMap["instance_id"] = instance.InstanceId
			}
			if instance.InstanceName != nil {
				instanceMap["instance_name"] = instance.InstanceName
			}
			if instance.AppId != nil {
				instanceMap["app_id"] = instance.AppId
			}
			if instance.ProjectId != nil {
				instanceMap["project_id"] = instance.ProjectId
			}
			if instance.Region != nil {
				instanceMap["region"] = instance.Region
			}
			if instance.VpcId != nil {
				instanceMap["vpc_id"] = instance.VpcId
			}
			if instance.SubnetId != nil {
				instanceMap["subnet_id"] = instance.SubnetId
			}
			if instance.StatusDesc != nil {
				instanceMap["status_desc"] = instance.StatusDesc
			}
			if instance.Status != nil {
				instanceMap["status"] = instance.Status
			}
			if instance.Vip != nil {
				instanceMap["vip"] = instance.Vip
			}
			if instance.Vport != nil {
				instanceMap["vport"] = instance.Vport
			}
			if instance.CreateTime != nil {
				instanceMap["create_time"] = instance.CreateTime
			}
			if instance.AutoRenewFlag != nil {
				instanceMap["auto_renew_flag"] = instance.AutoRenewFlag
			}
			if instance.Memory != nil {
				instanceMap["memory"] = instance.Memory
			}
			if instance.Storage != nil {
				instanceMap["storage"] = instance.Storage
			}
			if instance.ShardCount != nil {
				instanceMap["shard_count"] = instance.ShardCount
			}
			if instance.PeriodEndTime != nil {
				instanceMap["period_end_time"] = instance.PeriodEndTime
			}
			if instance.IsolatedTimestamp != nil {
				instanceMap["isolated_timestamp"] = instance.IsolatedTimestamp
			}
			if instance.Uin != nil {
				instanceMap["uin"] = instance.Uin
			}
			if instance.ShardDetail != nil {
				shardDetailList := []interface{}{}
				for _, shardDetail := range instance.ShardDetail {
					shardDetailMap := map[string]interface{}{}
					if shardDetail.ShardInstanceId != nil {
						shardDetailMap["shard_instance_id"] = shardDetail.ShardInstanceId
					}
					if shardDetail.ShardSerialId != nil {
						shardDetailMap["shard_serial_id"] = shardDetail.ShardSerialId
					}
					if shardDetail.Status != nil {
						shardDetailMap["status"] = shardDetail.Status
					}
					if shardDetail.Createtime != nil {
						shardDetailMap["createtime"] = shardDetail.Createtime
					}
					if shardDetail.Memory != nil {
						shardDetailMap["memory"] = shardDetail.Memory
					}
					if shardDetail.Storage != nil {
						shardDetailMap["storage"] = shardDetail.Storage
					}
					if shardDetail.ShardId != nil {
						shardDetailMap["shard_id"] = shardDetail.ShardId
					}
					if shardDetail.NodeCount != nil {
						shardDetailMap["node_count"] = shardDetail.NodeCount
					}
					if shardDetail.Cpu != nil {
						shardDetailMap["cpu"] = shardDetail.Cpu
					}

					shardDetailList = append(shardDetailList, shardDetailMap)
				}
				instanceMap["shard_detail"] = shardDetailList
			}
			if instance.NodeCount != nil {
				instanceMap["node_count"] = instance.NodeCount
			}
			if instance.IsTmp != nil {
				instanceMap["is_tmp"] = instance.IsTmp
			}
			if instance.WanDomain != nil {
				instanceMap["wan_domain"] = instance.WanDomain
			}
			if instance.WanVip != nil {
				instanceMap["wan_vip"] = instance.WanVip
			}
			if instance.WanPort != nil {
				instanceMap["wan_port"] = instance.WanPort
			}
			if instance.UpdateTime != nil {
				instanceMap["update_time"] = instance.UpdateTime
			}
			if instance.DbEngine != nil {
				instanceMap["db_engine"] = instance.DbEngine
			}
			if instance.DbVersion != nil {
				instanceMap["db_version"] = instance.DbVersion
			}
			if instance.Paymode != nil {
				instanceMap["paymode"] = instance.Paymode
			}
			if instance.WanStatus != nil {
				instanceMap["wan_status"] = instance.WanStatus
			}
			if instance.IsAuditSupported != nil {
				instanceMap["is_audit_supported"] = instance.IsAuditSupported
			}
			if instance.InstanceType != nil {
				instanceMap["instance_type"] = instance.InstanceType
			}
			//if instance.ResourceTags != nil {
			//	resourceTagsList := []interface{}{}
			//	for _, resourceTags := range instance.ResourceTags {
			//		resourceTagsMap := map[string]interface{}{}
			//		if resourceTags.TagKey != nil {
			//			resourceTagsMap["tag_key"] = resourceTags.TagKey
			//		}
			//		if resourceTags.TagValue != nil {
			//			resourceTagsMap["tag_value"] = resourceTags.TagValue
			//		}
			//
			//		resourceTagsList = append(resourceTagsList, resourceTagsMap)
			//	}
			//	instanceMap["resource_tags"] = resourceTagsList
			//}
			ids = append(ids, *instance.InstanceId)
			instanceList = append(instanceList, instanceMap)
		}
		d.SetId(helper.DataResourceIdsHash(ids))
		_ = d.Set("list", instanceList)
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), instanceList); e != nil {
			return e
		}
	}

	return nil
}
