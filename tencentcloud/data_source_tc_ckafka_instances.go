/*
Use this data source to query detailed instance information of Ckafka

# Example Usage

```hcl

	data "tencentcloudenterprise_ckafka_instances" "foo" {
	  instance_ids=["ckafka-vv7wpvae"]
	}

```
*/
package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_instances", CNDescription{
		TerraformTypeCN: "ckafka实例",
		DescriptionCN:   "提供CKafka实例数据源，用于查询CKafka实例的详细信息。",
		AttributesCN: map[string]string{
			"instance_ids":       "（过滤条件）按照实例ID过滤",
			"search_word":        "（过滤条件）按照实例名称过滤，支持模糊查询",
			"tag_key":            "标签的key。",
			"status":             "（过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部",
			"filters":            "查询过滤器>描述键值对过滤器，用于条件过滤查询。例如过滤ID、名称、状态等> * 若存在多个`Filter`时，`Filter`间的关系为逻辑与（`AND`）关系。> * 若同一个`Filter`存在多个`Values`，同一`Filter`下`Values`间的关系为逻辑或（`OR`）关系。>",
			"name":               "需要过滤的字段。",
			"value":              "字段的过滤值。",
			"offset":             "页面起始偏移量，默认值为“0”",
			"limit":              "页数，默认为“10”",
			"result_output_file": "用于保存结果",
			"instance_list":      "符合条件的实例详情列表",
			"instance_id":        "实例ID",
			"instance_name":      "实例名称",
			"vip":                "访问实例的vip 信息",
			"vport":              "访问实例的端口信息",
			"vip_list":           "虚拟IP列表",
			"bandwidth":          "实例带宽，单位Mbps",
			"disk_size":          "实例的存储大小，单位GB",
			"zone_id":            "可用区域ID",
			"vpc_id":             "vpcId，如果为空，说明是基础网络",
			"subnet_id":          "子网id",
			"renew_flag":         "实例是否续费，int  枚举值：1表示自动续费，2表示明确不自动续费",
			"healthy":            "实例状态 int：0表示健康，1表示告警，2 表示实例状态异常",
			"healthy_message":    "实例状态信息",
			"create_time":        "创建实例的时间",
			"expire_time":        "实例过期时间",
			"is_internal":        "是否为内部客户。值为1 表示内部客户",
			"topic_num":          "Topic个数",
			"tags":               "标识tag",
			"tag_value":          "标签的值。",
			"version":            "ckafka版本信息。注意：此字段可能返回null，表示无法检索到有效值",
			"zone_ids":           "跨可用性区。注意：此字段可能返回null，表示无法检索到有效值",
			"cvm":                "ckafka售卖类型。注意：此字段可能返回null，表示无法检索到有效值",
			//"disk_type":                  "磁盘类型注意：此字段可能返回null，表示无法检索到有效值",
			//"max_topic_number":           "当前规范中的最大主题数注意：此字段可能返回null，表示无法检索到有效值。。",
			//"max_partition_number":       "当前规范的最大分区数注意：此字段可能返回null，表示无法检索到有效值",
			//"rebalance_time":             "安排升级配置时间注意：此字段可能返回null，表示无法检索到有效值。。",
			//"partition_number":           "当前实例数注意：此字段可能返回null，表示无法检索到有效值。。",
			//"public_network_charge_type": "互联网带宽的类型注意：此字段可能返回null，表示无法检索到有效值。。",
			//"public_network":             "互联网带宽价值注意：此字段可能返回null，表示无法检索到有效值。。",
		},
	})
}

func dataSourceTencentCloudCkafkaInstances() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed instance information of Ckafka",
		Read:        dataSourceTencentCloudCkafkaInstancesRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by instance ID.",
			},
			"search_word": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by instance name, support fuzzy query.",
			},
			"tag_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Matches the tag key value.",
			},
			"status": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "(Filter Criteria) The status of the instance. 0: Create, 1: Run, 2: Delete, do not fill the default return all.",
			},
			"filters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The field that needs to be filtered.",
						},
						"values": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required:    true,
							Description: "The filtered value of the field.",
						},
					},
				},
				Description: "Filter. filter.name supports ('Ip', 'VpcId', 'SubNetId', 'InstanceType','InstanceId'), filter.values can pass up to 10 values.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The page start offset, default is `0`.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
				Description: "The number of pages, default is `10`.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"instance_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of ckafka users. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance ID.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance name.",
						},
						"vip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Virtual IP.",
						},
						"vport": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Virtual PORT.",
						},
						"vip_list": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Virtual IP.",
									},
									"vport": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Virtual PORT.",
									},
								},
							},
							Description: "Virtual IP entities.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The status of the instance. 0: Created, 1: Running, 2: Delete: 5 Quarantined, -1 Creation failed.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance bandwidth, in Mbps.",
						},
						"disk_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The storage size of the instance, in GB.",
						},
						"zone_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Availability Zone ID.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VpcId, if empty, indicates that it is the underlying network.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subnet id.",
						},
						"renew_flag": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether the instance is renewed, the int enumeration value: 1 indicates auto-renewal, and 2 indicates that it is not automatically renewed.",
						},
						"healthy": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance status int: 1 indicates health, 2 indicates alarm, and 3 indicates abnormal instance status.",
						},
						"healthy_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance status information.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The time when the instance was created.",
						},
						"expire_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The instance expiration time.",
						},
						"is_internal": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether it is an internal customer. A value of 1 indicates an internal customer.",
						},
						"topic_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of topics.",
						},
						"tags": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag Key.",
									},
									"tag_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag Value.",
									},
								},
							},
							Description: "Tag information.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Kafka version information. Note: This field may return null, indicating that a valid value could not be retrieved.",
						},
						"zone_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Description: "Across Availability Zones. Note: This field may return null, indicating that a valid value could not be retrieved.",
						},
						"cvm": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Ckafka sale type. Note: This field may return null, indicating that a valid value could not be retrieved.",
						},
						//"instance_type": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "Ckafka instance type. Note: This field may return null, indicating that a valid value could not be retrieved.",
						//},
						"disk_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Disk Type. Note: This field may return null, indicating that a valid value could not be retrieved.",
						},
						//"max_topic_number": {
						//	Type:        schema.TypeInt,
						//	Computed:    true,
						//	Description: "The maximum number of topics in the current specifications. Note: This field may return null, indicating that a valid value could not be retrieved..",
						//},
						//"max_partition_number": {
						//	Type:        schema.TypeInt,
						//	Computed:    true,
						//	Description: "The maximum number of Partitions for the current specifications. Note: This field may return null, indicating that a valid value could not be retrieved.",
						//},
						//"rebalance_time": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "Schedule the upgrade configuration time. Note: This field may return null, indicating that a valid value could not be retrieved..",
						//},
						//"partition_number": {
						//	Type:        schema.TypeInt,
						//	Computed:    true,
						//	Description: "The current number of instances. Note: This field may return null, indicating that a valid value could not be retrieved..",
						//},
						//"public_network_charge_type": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "The type of Internet bandwidth. Note: This field may return null, indicating that a valid value could not be retrieved..",
						//},
						//"public_network": {
						//	Type:        schema.TypeInt,
						//	Computed:    true,
						//	Description: "The Internet bandwidth value. Note: This field may return null, indicating that a valid value could not be retrieved..",
						//},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCkafkaInstancesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ckafka_instances.read")()

	ckafkaService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	request := ckafka.NewDescribeInstancesDetailRequest()
	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
	}
	if v, ok := d.GetOk("search_word"); ok {
		request.SearchWord = helper.String(v.(string))
	}
	if v, ok := d.GetOk("tag_key"); ok {
		request.TagKey = helper.String(v.(string))
	}
	if v, ok := d.GetOk("status"); ok {
		request.Status = helper.InterfacesIntInt64Point(v.([]interface{}))
	}
	if v, ok := d.GetOk("filters"); ok {
		filterParams := v.([]interface{})
		filters := make([]*ckafka.Filter, 0)
		for _, filterParam := range filterParams {
			filterParamMap := filterParam.(map[string]interface{})
			filters = append(filters, &ckafka.Filter{
				Name:   helper.String(filterParamMap["name"].(string)),
				Values: helper.InterfacesStringsPoint(filterParamMap["values"].([]interface{})),
			})
		}
		request.Filters = filters
	}
	var offset *int64
	if v, ok := d.GetOk("offset"); ok {
		offset = helper.IntInt64(v.(int))
	}
	request.Offset = offset

	if v, ok := d.GetOk("limit"); ok {
		request.Limit = helper.IntInt64(v.(int))
	}

	response, err := ckafkaService.client.UseCkafkaClient().DescribeInstancesDetail(request)
	if err != nil {
		return err
	}
	var kafkaInstanceDetails []*ckafka.InstanceDetail
	if response.Response.Result != nil {
		kafkaInstanceDetails = response.Response.Result.InstanceList
	}
	result := make([]map[string]interface{}, 0)
	ids := make([]string, 0)
	for _, kafkaInstanceDetail := range kafkaInstanceDetails {
		kafkaInstanceDetailMap := make(map[string]interface{})
		ids = append(ids, *kafkaInstanceDetail.InstanceId)
		kafkaInstanceDetailMap["instance_id"] = kafkaInstanceDetail.InstanceId
		kafkaInstanceDetailMap["instance_name"] = kafkaInstanceDetail.InstanceName
		kafkaInstanceDetailMap["vip"] = kafkaInstanceDetail.Vip
		kafkaInstanceDetailMap["vport"] = kafkaInstanceDetail.Vport
		kafkaInstanceDetailMap["status"] = kafkaInstanceDetail.Status
		kafkaInstanceDetailMap["bandwidth"] = kafkaInstanceDetail.Bandwidth
		kafkaInstanceDetailMap["disk_size"] = kafkaInstanceDetail.DiskSize
		kafkaInstanceDetailMap["zone_id"] = kafkaInstanceDetail.ZoneId
		kafkaInstanceDetailMap["vpc_id"] = kafkaInstanceDetail.VpcId
		kafkaInstanceDetailMap["subnet_id"] = kafkaInstanceDetail.SubnetId
		kafkaInstanceDetailMap["renew_flag"] = kafkaInstanceDetail.RenewFlag
		kafkaInstanceDetailMap["healthy"] = kafkaInstanceDetail.Healthy
		kafkaInstanceDetailMap["healthy_message"] = kafkaInstanceDetail.HealthyMessage
		kafkaInstanceDetailMap["create_time"] = kafkaInstanceDetail.CreateTime
		kafkaInstanceDetailMap["expire_time"] = kafkaInstanceDetail.ExpireTime
		kafkaInstanceDetailMap["is_internal"] = kafkaInstanceDetail.IsInternal
		kafkaInstanceDetailMap["topic_num"] = kafkaInstanceDetail.TopicNum
		kafkaInstanceDetailMap["version"] = kafkaInstanceDetail.Version
		kafkaInstanceDetailMap["cvm"] = kafkaInstanceDetail.Cvm
		//kafkaInstanceDetailMap["instance_type"] = kafkaInstanceDetail.InstanceType
		//kafkaInstanceDetailMap["max_topic_number"] = kafkaInstanceDetail.MaxTopicNumber
		//kafkaInstanceDetailMap["max_partition_number"] = kafkaInstanceDetail.MaxPartitionNumber
		//kafkaInstanceDetailMap["rebalance_time"] = kafkaInstanceDetail.RebalanceTime
		//kafkaInstanceDetailMap["partition_number"] = kafkaInstanceDetail.PartitionNumber
		//kafkaInstanceDetailMap["public_network_charge_type"] = kafkaInstanceDetail.PublicNetworkChargeType
		//kafkaInstanceDetailMap["public_network"] = kafkaInstanceDetail.PublicNetwork

		vipList := make([]map[string]string, 0)
		for _, vip := range kafkaInstanceDetail.VipList {
			vipList = append(vipList, map[string]string{
				"vip":   *vip.Vip,
				"vport": *vip.Vport,
			})
		}
		kafkaInstanceDetailMap["vip_list"] = vipList

		tags := make([]map[string]string, 0)
		for _, tag := range kafkaInstanceDetail.Tags {
			tags = append(tags, map[string]string{
				"tag_key":   *tag.TagKey,
				"tag_value": *tag.TagValue,
			})
		}
		kafkaInstanceDetailMap["tags"] = tags

		zoneIds := make([]int64, 0)
		for _, zoneId := range kafkaInstanceDetail.ZoneIds {
			zoneIds = append(zoneIds, *zoneId)
		}
		kafkaInstanceDetailMap["zone_ids"] = zoneIds

		result = append(result, kafkaInstanceDetailMap)
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("instance_list", result)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), result); e != nil {
			return e
		}
	}
	return nil
}
