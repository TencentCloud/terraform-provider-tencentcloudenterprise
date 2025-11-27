/*
Use this resource to create ckafka instance.

~> **NOTE:** It only support create prepaid ckafka instance.

Example Usage
Basic Instance
```hcl

	variable "vpc_id" {
	  default = "vpc-68vi2d3h"
	}

	variable "subnet_id" {
	  default = "subnet-ob6clqwk"
	}

	variable "product_info_list_standard" {
	  type = list(object({
	    name = string
	    value = string
	  }))
	  default = [
	    {
	      name  = "Region"
	      value = "shenzhen"
	    },
	    {
	      name  = "Cluster"
	      value = "ap-shenzhen-region-jcctest-ops-1 cluster"
	    },
	    {
	      name  = "AZ"
	      value = "ap-shenzhen-region-jcctest-ops-1 & ap-shenzhen-region-jcctest-ops-2"
	    },
	    {
	      name  = "Instance Name"
	      value = "Not named"
	    },
	    {
	      name  = "Specs Type"
	      value = "Standard Edition"
	    },
	    {
	      name  = "Product Model"
	      value = "Basic"
	    },
	    {
	      name  = "Peak Bandwidth"
	      value = "40MB/s"
	    },
	    {
	      name  = "Disk Capacity"
	      value = "300GB"
	    },
	    {
	      name  = "Message Retention Period"
	      value = "24 hours"
	    },
	    {
	      name  = "Network"
	      value = "vpc-hqa3zhut"
	    },
	    {
	      name  = "Subnet"
	      value = "subnet-dot6weka",
	    }
	  ]
	}

	variable "product_info_list_profession" {
	  type = list(object({
	    name = string
	    value = string
	  }))
	  default = [
	    {
	      name  = "Region"
	      value = "shenzhen"
	    },
	    {
	      name  = "Multi-AZ"
	      value = "ap-shenzhen-region-jcctest-ops-1 & ap-shenzhen-region-jcctest-ops-2"
	    },
	    {
	      name  = "Instance Name"
	      value = "Not named"
	    },
	    {
	      name  = "Specs Type"
	      value = "Pro Edition"
	    },
	    {
	      name  = "Kafka Version"
	      value = "2.4.1"
	    },
	    {
	      name  = "Peak Bandwidth"
	      value = "1600MB/s"
	    },
	    {
	      name  = "Disk Capacity"
	      value = "10000GB"
	    },
	    {
	      name  = "Topic"
	      value = 2000
	    },
	    {
	      name  = "Partition"
	      value = 4000
	    },
	    {
	      name  = "Message Retention Period"
	      value = "24 hours"
	    },
	    {
	      name  = "Network"
	      value = "vpc-hqa3zhut"
	    },
	    {
	      name  = "Subnet"
	      value = "subnet-dot6weka"
	    }
	  ]
	}

	data "tencentcloudenterprise_availability_zones" "gz" {
	  name    = "ap-guangzhou-3"
	  product = "ckafka"
	}

	resource "tencentcloudenterprise_ckafka_instance" "kafka_instance" {
	  instance_name      = "ckafka-instance-type-tf-test"
	  zone_id            = data.cloud_availability_zones.gz.zones.0.id
	  region_id          = 80000052
	  region_name        = "深圳"
	  pid                = 1004667
	  vpc_id             = var.vpc_id
	  subnet_id          = var.subnet_id
	  msg_retention_time = 1300
	  renew_flag         = 0
	  kafka_version      = "2.4.1"
	  disk_size          = 1000
	  disk_type          = "SSD"
	  instance_type      = "Basic"
	  topic              = 20
	  partition          = 4
	  dynamic "product_info" {
	    for_each = var.product_info_list_profession
	    content {
	      name = product_info.value["name"]
	      value = product_info.value["value"]
	    }
	  }
	}

```

Multi zone Instance
```hcl

	variable "vpc_id" {
	  default = "vpc-68vi2d3h"
	}

	variable "subnet_id" {
	  default = "subnet-ob6clqwk"
	}

	data "tencentcloudenterprise_availability_zones" "gz3" {
	  name    = "ap-guangzhou-3"
	  product = "ckafka"
	}

	data "tencentcloudenterprise_availability_zones" "gz6" {
	  name    = "ap-guangzhou-6"
	  product = "ckafka"
	}

	resource "tencentcloudenterprise_ckafka_instance" "kafka_instance" {
	  instance_name   = "ckafka-instance-maz-tf-test"
	  zone_id         = data.cloud_availability_zones.gz3.zones.0.id
	  multi_zone_flag = true
	  zone_ids        = [
	    data.cloud_availability_zones.gz3.zones.0.id,
	    data.cloud_availability_zones.gz6.zones.0.id
	  ]
	  period             = 1
	  vpc_id             = var.vpc_id
	  subnet_id          = var.subnet_id
	  msg_retention_time = 1300
	  renew_flag         = 0
	  kafka_version      = "1.1.1"
	  disk_size          = 500
	  disk_type          = "CLOUD_BASIC"
	  region_id          = 80000052
	  region_name        = "深圳"
	  pid                = 1
	}

```

# Import

ckafka instance can be imported using the instance_id, e.g.

```
$ terraform import cloud_ckafka_instance.foo ckafka-f9ife4zz
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

const (
	// CKafka 产品Code
	CKafkaProductCode = "p_ckafka"
	// Ckafka 子产品Code 这里默认是标准版
	CkafkaSubCode_Standard   = "sp_ckafka_standard"
	CkafkaSubCode_Profession = "sp_ckafka_profession"
	CkafkaProfessionPid      = 1004667
)

var (
	CKafkaStandardDiskSizeMap = map[int]int64{
		1: 300,
		2: 1000,
		3: 2500,
		4: 4000,
		5: 6000,
		6: 6000,
		7: 6000,
		8: 8000,
		9: 24000,
	}
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ckafka_instance", CNDescription{
		TerraformTypeCN: "ckafka实例",
		DescriptionCN:   "提供CKafka实例资源，用于创建和管理消息队列CKafka实例。",
		AttributesCN: map[string]string{
			"instance_name":      "实例名称",
			"zone_id":            "可用区域ID",
			"vpc_id":             "vpcId",
			"subnet_id":          "子网id",
			"msg_retention_time": "日志保存时间",
			"renew_flag":         "实例是否续费，int  枚举值：1表示自动续费，2表示明确不自动续费",
			"band_width":         "实例带宽，单位Mbps",
			"disk_size":          "实例的存储大小，单位GB",
			"tag_set":            "标识tag",
			"vip":                "访问实例的vip 信息",
			"vport":              "访问实例的端口信息",
			"instance_id":        "实例id",
			"cluster_id":         "集群id",
			"pid":                "实例类型",
			"region_id":          "地域ID",
			"region_name":        "地域名称",
			"goods_num":          "数量",
			"zone_ids":           "可用区信息",
			"product_info": "产品信息，\n" +
				" 当'规格类型'为'标准版'时'实例名'非必填，当'规格类型'为'专业版'时'实例名'、'产品型号'非必填，其余字段必填:\n" +
				"  - name: 地域, value: 2R3AZ仲裁区集成测试环境北京\n" +
				"  - name: 集群, value: cqyfm7 cluster\n" +
				"  - name: 可用区, value: 重庆云福M7\n" +
				"  - name: 实例名, value: test111\n" +
				"  - name: 规格类型, value: 标准版\n" +
				"  - name: 产品型号, value: 入门型\n" +
				"  - name: 峰值带宽, value: 40MB/s\n" +
				"  - name: 磁盘容量, value: 300GB\n" +
				"  - name: 信息保留时长, value: 72小时\n" +
				"  - name: 网络, value: vpc-kltzarib\n" +
				"  - name: 子网, value: subnet-7qt1q9h6",
		},
	})
}
func resourceTencentCloudCkafkaInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Use this resource to create ckafka instance.",
		Create:      resourceTencentCloudCkafkaInstanceCreate,
		Read:        resourceTencentCloudCkafkaInstanceRead,
		Update:      resourceTencentCloudCkafkaInstanceUpdate,
		Delete:      resourceTencentCLoudCkafkaInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "Instance name.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Ckafka instance ID.",
			},
			"zone_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Available zone id.",
				//DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				//	multiZone := d.Get("multi_zone_flag").(bool)
				//	zoneId := d.Get("zone_id").(int)
				//	v, ok := d.GetOk("zone_ids")
				//
				//	if !multiZone || !ok || old == "" {
				//		return old == new
				//	}
				//	zoneIds := v.(*schema.Set)
				//	return zoneIds.Contains(zoneId)
				//},
			},
			"cluster_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: "Cluster-ID represents  the cluster to which Ckafa belongs, " +
					"associated with cluster_name only needs to select one of them.",
			},
			//"specifications_type": {
			//	Type:         schema.TypeString,
			//	Optional:     true,
			//	Default:      "profession",
			//	ValidateFunc: validateAllowedStringValue([]string{"standard", "profession"}),
			//	Description:  "Specifications type of instance. Allowed values are `standard`, `profession`. Default is `profession`.",
			//},
			// "period": {
			// 	Type:        schema.TypeInt,
			// 	Optional:    true,
			// 	Description: "Prepaid purchase time, such as 1, is one month.",
			// },
			"instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"Basic", "High Performance"}),
				// Description:  "Description of instance type. `profession`: 1, `standard`:  1(general), 2(standard), 3(advanced), 4(capacity), 5(specialized-1), 6(specialized-2), 7(specialized-3), 8(specialized-4), 9(exclusive).",
				Description: "Description of instance type. [Basic, High Performance].",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vpc id.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subnet id.",
			},
			"msg_retention_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "The maximum retention time of instance logs, in minutes." +
					" the default is 10080 (7 days), the maximum is 30 days, and the default 0 is not filled," +
					" which means that the log retention time recovery policy is not enabled.",
			},
			"renew_flag": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "Prepaid automatic renewal mark, 0 means the default state, the initial state," +
					" 1 means automatic renewal, 2 means clear no automatic renewal (user setting).",
			},
			"kafka_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Kafka version (2.4.1/2.8.1).",
			},
			"band_width": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Instance bandwidth in MBps.",
			},
			"disk_size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				Description: "Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. " +
					"If it is not within the interval, the plan will cause a change when first created.",
			},
			"partition": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. " +
					"If it is not within the interval, the plan will cause a change when first created.",
			},
			"topic": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Topic Size",
			},
			"multi_zone_flag": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether the instance is multi zones. NOTE: if set to `true`, `zone_ids` must set together.",
			},
			"zone_ids": {
				Type:         schema.TypeSet,
				Optional:     true,
				Description:  "List of available zone id. NOTE: this argument must set together with `multi_zone_flag`.",
				RequiredWith: []string{"multi_zone_flag"},
				Elem:         &schema.Schema{Type: schema.TypeInt},
			},
			//"tags": {
			//	Type:          schema.TypeList,
			//	Optional:      true,
			//	Computed:      true,
			//	Deprecated:    "It has been deprecated from version 1.78.5, because it do not support change. Use `tag_set` instead.",
			//	ConflictsWith: []string{"tag_set"},
			//	Description:   "Tags of instance. Partition size, the professional version does not need tag.",
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"key": {
			//				Type:        schema.TypeString,
			//				Required:    true,
			//				Description: "Tag key.",
			//			},
			//			"value": {
			//				Type:        schema.TypeString,
			//				Required:    true,
			//				Description: "Tag value.",
			//			},
			//		},
			//	},
			//},
			"tag_set": {
				Type:        schema.TypeMap,
				Optional:    true,
				Computed:    true,
				Description: "Tag set of instance.",
			},
			"disk_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Type of disk. [SSD]",
			},
			//"config": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	MaxItems: 1,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"auto_create_topic_enable": {
			//				Type:        schema.TypeBool,
			//				Required:    true,
			//				Description: "Automatic creation. true: enabled, false: not enabled.",
			//			},
			//			"default_num_partitions": {
			//				Type:     schema.TypeInt,
			//				Required: true,
			//				Description: "If auto.create.topic.enable is set to true and this value is not set, " +
			//					"3 will be used by default.",
			//			},
			//			"default_replication_factor": {
			//				Type:     schema.TypeInt,
			//				Required: true,
			//				Description: "If auto.create.topic.enable is set to true but this value is not set, " +
			//					"2 will be used by default.",
			//			},
			//		},
			//	},
			//	Description: "Instance configuration.",
			//},
			//"dynamic_retention_config": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	MaxItems: 1,
			//	Computed: true,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"enable": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//				Description: "Whether the dynamic message retention time configuration is " +
			//					"enabled. 0: disabled; 1: enabled.",
			//			},
			//			"disk_quota_percentage": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//				Description: "Disk quota threshold (in percentage) for triggering " +
			//					"the message retention time change event.",
			//			},
			//			"step_forward_percentage": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//				Description: "Percentage by which the message retention " +
			//					"time is shortened each time.",
			//			},
			//			"bottom_retention": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Computed:    true,
			//				Description: "Minimum retention time, in minutes.",
			//			},
			//		},
			//	},
			//	Description: "Dynamic message retention policy configuration.",
			//},
			//"rebalance_time": {
			//	Type:        schema.TypeInt,
			//	Optional:    true,
			//	Description: "Modification of the rebalancing time after upgrade.",
			//},
			//"public_network": {
			//	Type:        schema.TypeInt,
			//	Optional:    true,
			//	Computed:    true,
			//	Deprecated:  "It has been deprecated from version 1.81.6. If set public network value, it will cause error.",
			//	Description: "Bandwidth of the public network.",
			//},
			//"dynamic_disk_config": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	MaxItems: 1,
			//	Computed: true,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"enable": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Computed: true,
			//				Description: "Whether to the dynamic disk expansion configuration is enabled." +
			//					"0: disabled; 1: enabled.",
			//			},
			//			"disk_quota_percentage": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Computed:    true,
			//				Description: "Disk quota threshold (in percentage) for triggering the automatic disk expansion event.",
			//			},
			//			"step_forward_percentage": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Computed:    true,
			//				Description: "Percentage of dynamic disk expansion each time.",
			//			},
			//			"max_disk_space": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Computed:    true,
			//				Description: "Max scale disk size, in GB.",
			//			},
			//		},
			//	},
			//	Description: "Dynamic disk expansion policy configuration.",
			//},
			//"max_message_byte": {
			//	Type:         schema.TypeInt,
			//	Optional:     true,
			//	Computed:     true,
			//	ValidateFunc: validateIntegerInRange(1024, 12*1024*1024),
			//	Description:  "The size of a single message in bytes at the instance level. Value range: `1024 - 12*1024*1024 bytes (i.e., 1KB-12MB).",
			//},
			"vip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Vip of instance.",
			},
			"vport": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of instance.",
			},
			"region_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Region ID.",
			},
			"region_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region Name.",
			},
			"pid": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Pricing formula ID. 1-9",
			},
			"goods_num": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Quantity.",
			},
			"product_info": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Item name.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Item value.",
						},
					},
				},
				Description: "Product information, when 规格类型=标准版 :\n" +
					" When '规格类型' is '标准版','实例名' is optional; when '规格类型' is '专业版','实例名' and '产品型号'\n" +
					" are optional, all other fields are required:\n" +
					"  - name: 地域, value: 2R3AZ仲裁区集成测试环境北京\n" +
					"  - name: 集群, value: cqyfm7 cluster\n" +
					"  - name: 可用区, value: 重庆云福M7\n" +
					"  - name: 实例名, value: test111\n" +
					"  - name: 规格类型, value: 标准版\n" +
					"  - name: 产品型号, value: 入门型\n" +
					"  - name: 峰值带宽, value: 40MB/s\n" +
					"  - name: 磁盘容量, value: 300GB\n" +
					"  - name: 信息保留时长, value: 72小时\n" +
					"  - name: 网络, value: vpc-kltzarib\n" +
					"  - name: 子网, value: subnet-7qt1q9h6",
			},
		},
	}
}

func buildGoodsInfo(d *schema.ResourceData) *ckafka.GoodsInfo {
	// 默认付费方式后付费
	var payMode int64 = 0
	// 产品类型默认p_ckafka
	var productType string = CKafkaProductCode

	goodsInfo := &ckafka.GoodsInfo{
		GoodsNum: helper.IntInt64(d.Get("goods_num").(int)),
		PayMode:  &payMode,
		Type:     &productType,
	}

	goodsInfo.RegionId = helper.IntInt64(d.Get("region_id").(int))

	goodsInfo.ZoneId = helper.IntInt64(d.Get("zone_id").(int))

	goodsDetail := buildGoodsDetail(d)

	goodsDetailByte, err := goodsDetail.MarshalJSON()
	if err != nil {
		log.Printf("[ERROR] buildGoodsDetail failed, err: %v", err)
		panic(err)
	}

	log.Printf("[DEBUG] goodsDetail: %v", string(goodsDetailByte))
	goodsInfo.GoodsDetail = helper.String(string(goodsDetailByte))

	return goodsInfo
}

func buildGoodsDetail(d *schema.ResourceData) *ckafka.GoodsDetail {
	goodsDetail := &ckafka.GoodsDetail{}
	goodsDetail.ProductCode = CKafkaProductCode
	goodsDetail.TimeUnit = "h"
	// 购买时长这里默认是1小时， 通过购买数量来设置实际时间
	goodsDetail.TimeSpan = 1

	clusterId := d.Get("cluster_id").(int)
	goodsDetail.ClusterId = int64(clusterId)

	renewFlag := d.Get("renew_flag").(int)
	goodsDetail.AutoRenewFlag = int64(renewFlag)

	instanceName := d.Get("instance_name").(string)
	goodsDetail.InstanceName = instanceName

	vpcId := d.Get("vpc_id").(string)
	goodsDetail.VpcId = vpcId

	subnet_id := d.Get("subnet_id").(string)
	goodsDetail.SubnetId = subnet_id

	goodsDetail.ProductInfo = buildProductInfo(d)

	pid := d.Get("pid").(int)
	goodsDetail.Pid = int64(pid)

	diskSize := d.Get("disk_size").(int)

	if v, ok := d.GetOk("msg_retention_time"); ok {
		goodsDetail.MsgRetentionTime = int64(v.(int))
	}

	goodsDetail.ExData = make(map[string]interface{})

	// 专业版
	if pid == CkafkaProfessionPid {
		goodsDetail.SubProductCode = CkafkaSubCode_Profession
		goodsDetail.ExData["sv_ckafka_profession_partition_package_s1"] = 0
		goodsDetail.ExData["sv_ckafka_profession_cloud_disk_disk_ssd"] = d.Get("disk_size").(int)
		if v, ok := d.GetOk("disk_type"); ok {
			goodsDetail.ExData["diskType"] = v.(string)
		}
		goodsDetail.ExData["sv_ckafka_profession_topic_package_topic"] = 0
		if v, ok := d.GetOk("instance_type"); ok {
			instanceType := v.(string)
			if instanceType == "Basic" {
				goodsDetail.ExData["sv_ckafka_profession_instance_smallpackage"] = 1
				goodsDetail.ExData["sv_ckafka_profession_instance_extra_package_b1"] = 0
			} else if instanceType == "High Performance" {
				goodsDetail.ExData["sv_ckafka_profession_instance_peaktraffic_package_s1"] = 1
				goodsDetail.ExData["sv_ckafka_profession_instance_extrapackage_200s1"] = 0
			}
		}
		goodsDetail.ExData["resourceTags"] = make(map[string]interface{})
		if v, ok := d.GetOk("kafka_version"); ok {
			goodsDetail.ExData["brokerVersion"] = v.(string)
		}

		exParams := map[string]string{
			"band_width": "bandwidth",
			"topic":      "topic",
			"disk_size":  "disk",
			"partition":  "partition",
		}
		for ticK, pK := range exParams {
			if v, ok := d.GetOk(ticK); ok {
				goodsDetail.ExData[pK] = v.(int)
			}
		}
		productDetail := ckafka.ProductsDetail{
			CVM:           1,
			Tags:          []string{},
			IPType:        4,
			MultiZoneFlag: true,
		}
		if v, ok := d.GetOk("zone_ids"); ok {
			zoneIdsSet := v.(*schema.Set).List()
			zoneIds := make([]int, 0)
			for _, zoneId := range zoneIdsSet {
				zoneIds = append(zoneIds, zoneId.(int))
			}
			productDetail.ZoneIdList = zoneIds
		}
		if v, ok := d.GetOk("multi_zone_flag"); ok {
			productDetail.MultiZoneFlag = v.(bool)
		}
		if pid == CkafkaProfessionPid {
			productDetail.InstanceType = "profession"
		}
		if v, ok := d.GetOk("kafka_version"); ok {
			productDetail.BrokerVersion = v.(string)
		}
		goodsDetail.ExData["productsDetail"] = productDetail

	} else {
		// 标准版 diskSize 为额外磁盘大小
		goodsDetail.SubProductCode = CkafkaSubCode_Standard

		goodsDetail.ExData["CKafka-Standard-Disk"] = func() int64 {
			ret := int64(diskSize) - CKafkaStandardDiskSizeMap[pid]
			if ret < 0 {
				ret = 0
			}
			return ret
		}()

		goodsDetail.ExData[fmt.Sprintf("CKafka-Standard-Instance-S1_%d", pid)] = 1
	}

	return goodsDetail
}

func buildProductInfo(d *schema.ResourceData) []ckafka.ProductInfo {
	productInfoList := make([]ckafka.ProductInfo, 0)
	rawData := d.Get("product_info").([]interface{})
	for data := range rawData {
		productInfo := ckafka.ProductInfo{}
		tmpMap := rawData[data].(map[string]interface{})
		productInfo.Name = tmpMap["name"].(string)
		productInfo.Value = tmpMap["value"].(string)
		productInfoList = append(productInfoList, productInfo)
	}
	return productInfoList
}

func resourceTencentCloudCkafkaInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_instance.create")()
	var (
		logId   = getLogId(contextNil)
		service = CkafkaService{
			client: meta.(*TencentCloudClient).apiV3Conn,
		}
		request  = ckafka.NewHourCreateRequest()
		response = ckafka.NewHourCreateResponse()
		ctx      = context.WithValue(context.TODO(), logIdKey, logId)
	)

	request.GoodsInfoList = make([]*ckafka.GoodsInfo, 0)
	goodsInfo := buildGoodsInfo(d)
	request.GoodsInfoList = append(request.GoodsInfoList, goodsInfo)

	result, err := service.client.UseCkafkaClient().HourCreate(request)
	response = result

	if err != nil {
		log.Printf("[CRITAL]%s create ckafka instance failed, reason:%s\n", logId, err.Error())
		return err
	}

	instanceId := response.Response.Result.ResourceIds[0]

	err = resource.Retry(5*readRetryTimeout, func() *resource.RetryError {
		has, ready, err := service.CheckCkafkaInstanceReady(ctx, *instanceId)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if !has {
			return resource.NonRetryableError(fmt.Errorf("ckafka instance not exists."))
		}
		if ready {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("create ckafka instance task is processing"))
	})
	if err != nil {
		return err
	}
	d.SetId(*instanceId)

	//if v, ok := d.GetOk("dynamic_disk_config"); ok {
	//	needModify = true
	//	dynamic := make([]*ckafka.DynamicDiskConfig, 0, 10)
	//	for _, item := range v.([]interface{}) {
	//		dMap := item.(map[string]interface{})
	//		dynamicInfo := ckafka.DynamicDiskConfig{}
	//		if enable, ok := dMap["enable"]; ok {
	//			dynamicInfo.Enable = helper.Int64(int64(enable.(int)))
	//		}
	//		if stepForwardPercentage, ok := dMap["step_forward_percentage"]; ok {
	//			dynamicInfo.StepForwardPercentage = helper.Int64(int64(stepForwardPercentage.(int)))
	//		}
	//		if diskQuotaPercentage, ok := dMap["disk_quota_percentage"]; ok {
	//			dynamicInfo.DiskQuotaPercentage = helper.Int64(int64(diskQuotaPercentage.(int)))
	//		}
	//		if maxDiskSpace, ok := dMap["max_disk_space"]; ok {
	//			dynamicInfo.MaxDiskSpace = helper.Int64(int64(maxDiskSpace.(int)))
	//		}
	//		dynamic = append(dynamic, &dynamicInfo)
	//	}
	//	modifyRequest.DynamicDiskConfig = dynamic[0]
	//}

	client := meta.(*TencentCloudClient).apiV3Conn
	tagService := TagService{client: client}
	region := client.Region

	if tags := helper.GetTags(d, "tag_set"); len(tags) > 0 {
		resourceName := BuildTagResourceName("ckafka", "ckafkaId", region, *instanceId)
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudCkafkaInstanceRead(d, meta)
}

func resourceTencentCloudCkafkaInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var service = CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	instanceId := d.Id()

	var info *ckafka.InstanceDetail
	var isExist = true

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		res, has, e := service.DescribeCkafkaInstanceById(ctx, instanceId)
		if e != nil {
			return retryError(e)
		}
		if !has {
			d.SetId("")
			isExist = false
			return nil
		}
		info = res
		return nil
	})
	if err != nil {
		return fmt.Errorf("[API]Describe kafka instance fail, reason:%s", err.Error())
	}

	if !isExist {
		return nil
	}

	fillZone := func() error {
		if v, ok := d.GetOk("zone_id"); !ok {
			return fmt.Errorf("zone_id is required")
		} else {
			for _, zoneId := range info.ZoneIds {
				if *zoneId == int64(v.(int)) {
					_ = d.Set("zone_id", zoneId)
					return nil
				}
			}
			return fmt.Errorf("zone_id %d not found", v.(int))
		}
	}
	_ = d.Set("instance_name", info.InstanceName)

	if err = fillZone(); err != nil {
		return err
	}
	// calculate period
	//createTime := *info.CreateTime
	//expireTime := *info.ExpireTime
	//period := (expireTime - createTime) / (3600 * 24 * 31)
	//_ = d.Set("period", &period)
	_ = d.Set("vpc_id", info.VpcId)
	_ = d.Set("subnet_id", info.SubnetId)
	_ = d.Set("renew_flag", info.RenewFlag)
	//_ = d.Set("kafka_version", info.Version)
	_ = d.Set("disk_size", info.DiskSize)
	bandWidth := info.Bandwidth
	_ = d.Set("vip", info.Vip)
	_ = d.Set("vport", info.Vport)
	_ = d.Set("band_width", *bandWidth/8)

	//if len(info.ZoneIds) > 1 {
	//	_ = d.Set("multi_zone_flag", true)
	//	ids := helper.Int64sInterfaces(info.ZoneIds)
	//	idSet := schema.NewSet(func(i interface{}) int {
	//		return i.(int)
	//	}, ids)
	//	_ = d.Set("zone_ids", idSet)
	//}

	//tagSets := make([]map[string]interface{}, 0, len(info.Tags))
	//for _, item := range info.Tags {
	//	tagSets = append(tagSets, map[string]interface{}{
	//		"key":   item.TagKey,
	//		"value": item.TagValue,
	//	})
	//}
	//_ = d.Set("tags", tagSets)
	_ = d.Set("instance_id", instanceId)

	client := meta.(*TencentCloudClient).apiV3Conn
	tagService := TagService{client: client}
	region := client.Region

	tags, err := tagService.DescribeResourceTags(ctx, "ckafka", "ckafkaId", region, instanceId)
	if err != nil {
		return err
	}
	_ = d.Set("tag_set", tags)

	// query msg_retention_time
	var (
		request  = ckafka.NewDescribeInstanceAttributesRequest()
		response = ckafka.NewDescribeInstanceAttributesResponse()
	)
	request.InstanceId = &instanceId
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.client.UseCkafkaClient().DescribeInstanceAttributes(request)
		if e != nil {
			return retryError(e)
		}
		response = result
		attr := response.Response.Result
		_ = d.Set("msg_retention_time", attr.MsgRetentionTime)

		//if attr.Config != nil {
		//	config := make([]map[string]interface{}, 0)
		//	config = append(config, map[string]interface{}{
		//		"auto_create_topic_enable":   attr.Config.AutoCreateTopicsEnable,
		//		"default_num_partitions":     attr.Config.DefaultNumPartitions,
		//		"default_replication_factor": attr.Config.DefaultReplicationFactor,
		//	})
		//	_ = d.Set("config", config)
		//}

		//dynamicDiskConfig := make([]map[string]interface{}, 0)
		//dynamicDiskConfig = append(dynamicDiskConfig, map[string]interface{}{
		//	"enable":                  attr.DynamicDiskConfig.Enable,
		//	"disk_quota_percentage":   attr.DynamicDiskConfig.DiskQuotaPercentage,
		//	"step_forward_percentage": attr.DynamicDiskConfig.StepForwardPercentage,
		//	"max_disk_space":          attr.DynamicDiskConfig.MaxDiskSpace,
		//})
		//_ = d.Set("dynamic_disk_config", dynamicDiskConfig)

		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Ckafka Instance Attributes failed, reason:%s\n", logId, err.Error())
		return err
	}
	return nil
}

func resourceTencentCloudCkafkaInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_instance.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	immutableArgs := []string{
		"period", "vpc_id",
		"subnet_id", "renew_flag", "kafka_version",
		"multi_zone_flag", "zone_ids", "disk_type",
		"specifications_type", "instance_type", "cluster_id",
	}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	instanceId := d.Id()
	request := ckafka.NewModifyInstanceAttributesRequest()
	//if d.HasChange("instance_name") {
	//	if v, ok := d.GetOk("instance_name"); ok {
	//		request.InstanceName = helper.String(v.(string))
	//	}
	//}
	//
	//if d.HasChange("msg_retention_time") {
	//	if v, ok := d.GetOk("msg_retention_time"); ok {
	//		request.MsgRetentionTime = helper.Int64(int64(v.(int)))
	//	}
	//}
	//
	//if d.HasChange("config") {
	//	if v, ok := d.GetOk("config"); ok {
	//		items := v.([]interface{})
	//		dMap := items[0].(map[string]interface{})
	//		configInfo := ckafka.ModifyInstanceAttributesConfig{}
	//		if autoCreateTopicEnable, ok := dMap["auto_create_topic_enable"]; ok {
	//			configInfo.AutoCreateTopicEnable = helper.Bool(autoCreateTopicEnable.(bool))
	//		}
	//		if defaultNumPartitions, ok := dMap["default_num_partitions"]; ok {
	//			configInfo.DefaultNumPartitions = helper.Int64(int64(defaultNumPartitions.(int)))
	//		}
	//		if defaultReplicationFactor, ok := dMap["default_replication_factor"]; ok {
	//			configInfo.DefaultReplicationFactor = helper.Int64(int64(defaultReplicationFactor.(int)))
	//		}
	//		request.Config = &configInfo
	//	}
	//}
	//
	//if d.HasChange("dynamic_retention_config") {
	//	if v, ok := d.GetOk("dynamic_retention_config"); ok {
	//		items := v.([]interface{})
	//		dMap := items[0].(map[string]interface{})
	//		dynamicInfo := ckafka.DynamicRetentionTime{}
	//		if enable, ok := dMap["enable"]; ok {
	//			dynamicInfo.Enable = helper.Int64(int64(enable.(int)))
	//		}
	//		if diskQuotaPercentage, ok := dMap["disk_quota_percentage"]; ok {
	//			dynamicInfo.DiskQuotaPercentage = helper.Int64(int64(diskQuotaPercentage.(int)))
	//		}
	//		if stepForwardPercentage, ok := dMap["step_forward_percentage"]; ok {
	//			dynamicInfo.StepForwardPercentage = helper.Int64(int64(stepForwardPercentage.(int)))
	//		}
	//		if bottomRetention, ok := dMap["bottom_retention"]; ok {
	//			dynamicInfo.BottomRetention = helper.Int64(int64(bottomRetention.(int)))
	//		}
	//		request.DynamicRetentionConfig = &dynamicInfo
	//	}
	//}
	//
	//if d.HasChange("rebalance_time") {
	//	if v, ok := d.GetOk("rebalance_time"); ok {
	//		request.RebalanceTime = helper.Int64(int64(v.(int)))
	//	}
	//}
	//
	//if d.HasChange("public_network") {
	//	if v, ok := d.GetOk("public_network"); ok {
	//		request.PublicNetwork = helper.Int64(int64(v.(int)))
	//	}
	//}

	//if d.HasChange("dynamic_disk_config") {
	//	if v, ok := d.GetOk("dynamic_disk_config"); ok {
	//		dynamic := make([]*ckafka.DynamicDiskConfig, 0, 10)
	//		for _, item := range v.([]interface{}) {
	//			dMap := item.(map[string]interface{})
	//			dynamicInfo := ckafka.DynamicDiskConfig{}
	//			if enable, ok := dMap["enable"]; ok {
	//				dynamicInfo.Enable = helper.Int64(int64(enable.(int)))
	//			}
	//			if stepForwardPercentage, ok := dMap["step_forward_percentage"]; ok {
	//				dynamicInfo.StepForwardPercentage = helper.Int64(int64(stepForwardPercentage.(int)))
	//			}
	//			if diskQuotaPercentage, ok := dMap["disk_quota_percentage"]; ok {
	//				dynamicInfo.DiskQuotaPercentage = helper.Int64(int64(diskQuotaPercentage.(int)))
	//			}
	//			if maxDiskSpace, ok := dMap["max_disk_space"]; ok {
	//				dynamicInfo.MaxDiskSpace = helper.Int64(int64(maxDiskSpace.(int)))
	//			}
	//			dynamic = append(dynamic, &dynamicInfo)
	//		}
	//		request.DynamicDiskConfig = dynamic[0]
	//	}
	//}

	//if d.HasChange("max_message_byte") {
	//	if v, ok := d.GetOkExists("max_message_byte"); ok {
	//		request.MaxMessageByte = helper.Uint64(uint64(v.(int)))
	//	}
	//}
	//
	//err := service.ModifyCkafkaInstanceAttributes(ctx, request)
	//if err != nil {
	//	return fmt.Errorf("[API]Set kafka instance attributes fail, reason:%s", err.Error())
	//}
	//
	if d.HasChange("instance_name") || d.HasChange("msg_retention_time") {
		request.InstanceId = helper.String(instanceId)
		if v, ok := d.GetOk("instance_name"); ok {
			request.InstanceName = helper.String(v.(string))
		}
		if v, ok := d.GetOk("msg_retention_time"); ok {
			request.MsgRetentionTime = helper.Int64(int64(v.(int)))
		}
		//if v, ok := d.GetOk("partition"); ok {
		//	request.Partition = helper.Int64(int64(v.(int)))
		//}
		//
		_, err := service.client.UseCkafkaClient().ModifyInstanceAttributes(request)
		if err != nil {
			return fmt.Errorf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]", logId,
				request.GetAction(), request.ToJsonString(), err.Error())
		}

		err = resource.Retry(5*readRetryTimeout, func() *resource.RetryError {
			_, ready, err := service.CheckCkafkaInstanceReady(ctx, instanceId)
			if err != nil {
				return resource.NonRetryableError(err)
			}
			if ready {
				return nil
			}
			return resource.RetryableError(fmt.Errorf("upgrade ckafka instance task is processing"))
		})
	}
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	if d.HasChange("tag_set") {

		client := meta.(*TencentCloudClient).apiV3Conn
		tagService := TagService{client: client}
		region := client.Region

		oldTags, newTags := d.GetChange("tag_set")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))

		resourceName := BuildTagResourceName("ckafka", "ckafkaId", region, instanceId)
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudCkafkaInstanceRead(d, meta)
}

func resourceTencentCLoudCkafkaInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_instance.delete")()
	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = CkafkaService{
			client: meta.(*TencentCloudClient).apiV3Conn,
		}
		request = ckafka.NewDeleteInstanceRequest()
	)
	instanceId := d.Id()
	request.InstanceIds = []*string{helper.String(instanceId)}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, err := service.client.UseCkafkaClient().DeleteInstance(request)
		if err != nil {
			return retryError(err, "UnsupportedOperation")
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		has, _, err := service.CheckCkafkaInstanceReady(ctx, instanceId)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if !has {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("delete ckafka instance task is processing"))
	})
	if err != nil {
		return err
	}
	return nil
}
