/*
Provides a resource to create a dcdb hourdb_instance

# Example Usage

```hcl

	resource "tencentcloudenterprise_dcdb_instance" "hourdb_instance" {
	  instance_name    = "111"
	  zones            = ["yfm18", "yfm18"]
	  cpu_arch         = "X86"
	  ipv6_flag        = 0
	  shard_memory     = 2
	  shard_storage    = 10
	  shard_node_count = 2
	  shard_cpu        = 1
	  shard_count      = 2
	  vpc_id           = "vpc-cs6ffr73"
	  subnet_id        = "subnet-mfbxe9zk"
	  db_version_id    = "5.7"
	  project_id       = "pr-bae40f73"

	  resource_tags {
	    tag_key   = "createdBy"
	    tag_value = "terraform3"
	  }

	  init_params {

	  }
	}

```
Import

dcdb hourdb_instance can be imported using the id, e.g.
```
$ terraform import cloud_dcdb_instance.hourdb_instance hourdbInstance_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dcdb_instance", CNDescription{
		TerraformTypeCN: "云数据库DCDB实例",
		DescriptionCN:   "提供云数据库DCDB实例资源，用于创建和管理分布式数据库DCDB实例。",
		AttributesCN: map[string]string{
			"zones":                  "可用区",
			"shard_cpu":              "每个分片的CPU核数",
			"cpu_arch":               "CPU架构",
			"shard_memory":           "每个分片的内存",
			"shard_storage":          "每个分片的存储",
			"shard_node_count":       "每个分片的节点数",
			"shard_count":            "实例分片数",
			"instance_count":         "实例数量",
			"vpc_id":                 "VPC ID",
			"subnet_id":              "子网 ID",
			"db_version_id":          "数据库引擎版本",
			"security_group_id":      "安全组 ID",
			"project_id":             "项目 ID",
			"instance_name":          "实例名称",
			"dcn_region":             "DCN源地域",
			"dcn_instance_id":        "DCN源实例ID",
			"ipv6_flag":              "是否支持IPv6",
			"extranet_access":        "是否开启外网访问",
			"vip":                    "内网访问IP",
			"vipv6":                  "IPv6访问IP",
			"vport":                  "内网访问端口",
			"resource_tags":          "资源标签",
			"init_params":            "初始化参数",
			"character_set_server":   "实例的可用区域ID",
			"lower_case_table_names": "（表名区分大小写，必填，0-敏感；1-不敏感）",
			"sync_mode":              "（同步模式0-异步；1-强同步；2-可降解强同步。默认为强同步）",
			"tag_key":                "标签密钥",
			"tag_value":              "标签值",
		},
	})
}
func resourceTencentCloudDcdbdbInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDcdbdbInstanceCreate,
		Read:   resourceTencentCloudDcdbdbInstanceRead,
		Update: resourceTencentCloudDcdbdbInstanceUpdate,
		Delete: resourceTencentCloudDcdbdbInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zones": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "Available zone.",
			},
			"shard_cpu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Cpu core for each shard. It can be obtained by querying api DescribeShardSpec.",
			},

			"cpu_arch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cpu architecture, default to x86_64.",
				Default:     "x86",
			},

			"shard_memory": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Memory(GB) for each shard. It can be obtained by querying api DescribeShardSpec.",
			},

			"shard_storage": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Storage(GB) for each shard. It can be obtained by querying api DescribeShardSpec.",
			},

			"shard_node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Node count for each shard. It can be obtained by querying api DescribeShardSpec.",
			},

			"shard_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Instance shard count.",
			},

			"instance_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Instance count, default to 1.",
			},

			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vpc id.",
			},

			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subnet id, its required when vpcId is set.",
			},

			"db_version_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Db engine version, default to Percona 5.7.17.",
			},

			"security_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Security group id.",
			},

			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Project id.",
			},

			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of this instance.",
			},

			"dcn_region": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "DCN source region.",
			},

			"dcn_instance_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "DCN source instance ID.",
			},

			"ipv6_flag": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Whether to support IPv6.",
			},

			"extranet_access": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether to open the extranet access.",
			},

			"vip": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The field is required to specify VIP.",
			},

			"vipv6": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The field is required to specify VIPv6.",
			},

			"vport": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Intranet port.",
			},

			"resource_tags": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Resource tags.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Tag key.",
						},
						"tag_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Tag value.",
						},
					},
				},
			},

			"init_params": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "The optional values for this interface are: character_set_server (character set,required), lower_case_table_names (case sensitivity of table names, required, 0 - sensitive; 1 - insensitive), innodb_page_size (innodb data page, default 16K), sync_mode (synchronization mode: 0 - asynchronous; 1 - strong synchronization; 2 - degradable strong synchronization. Default is strong synchronization).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"character_set_server": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The available zone ID of the instance.",
							Default:      DCDB_CHARACTER_SET_UTF8,
							ValidateFunc: validateAllowedStringValue(DCDB_CHARACTER_SET),
						},
						"lower_case_table_names": {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "(case sensitivity of table names, required, 0 - sensitive; 1 - insensitive).",
							ValidateFunc: validateAllowedIntValue(DCDB_LOWER_CASE_TABLE_NAMES),
							Default:      DCDB_LOWER_CASE_TABLE_NAMES_0,
						},
						"sync_mode": {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  "(synchronization mode: 0 - asynchronous; 1 - strong synchronization; 2 - degradable strong synchronization. Default is strong synchronization).",
							ValidateFunc: validateIntegerInRange(0, 2),
							Default:      0,
						},
					},
				},
			},
		},
		Description: "Provides a dcdb instance resource.",
	}
}

func resourceTencentCloudDcdbdbInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_instance.create")()
	defer inconsistentCheck(d, meta)()

	var (
		request    = dcdb.NewCreateHourDCDBInstanceRequest()
		response   *dcdb.CreateHourDCDBInstanceResponse
		instanceId string
		ipv6Flag   int
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	if v, ok := d.GetOk("zones"); ok {
		zonesSet := v.(*schema.Set).List()
		for i := range zonesSet {
			zones := zonesSet[i].(string)
			request.Zones = append(request.Zones, &zones)
		}
	}

	if v, ok := d.GetOk("instance_count"); ok {
		request.Count = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("ipv6_flag"); ok {
		ipv6Flag = v.(int)
		request.Ipv6Flag = helper.Int64(int64(ipv6Flag))
	}

	if v, ok := d.GetOk("cpu_arch"); ok {
		request.CpuArch = helper.String(v.(string))
	}

	if v, ok := d.GetOk("shard_cpu"); ok {
		request.ShardCpu = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("shard_memory"); ok {
		request.ShardMemory = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("shard_storage"); ok {
		request.ShardStorage = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("shard_node_count"); ok {
		request.ShardNodeCount = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("shard_count"); ok {
		request.ShardCount = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))

	}

	if v, ok := d.GetOk("subnet_id"); ok {
		request.SubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("db_version_id"); ok {
		request.DbVersionId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("security_group_id"); ok {
		request.SecurityGroupId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("project_id"); ok {
		request.PlatformProjectId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	if v, _ := d.GetOk("ipv6_flag"); v != nil {
		ipv6Flag = v.(int)
		request.Ipv6Flag = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("resource_tags"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			resourceTag := dcdb.ResourceTag{}
			if v, ok := dMap["tag_key"]; ok {
				resourceTag.TagKey = helper.String(v.(string))
			}
			if v, ok := dMap["tag_value"]; ok {
				resourceTag.TagValue = helper.String(v.(string))
			}
			request.ResourceTags = append(request.ResourceTags, &resourceTag)
		}
	}

	if v, ok := d.GetOk("dcn_region"); ok {
		request.DcnRegion = helper.String(v.(string))
	}

	if v, ok := d.GetOk("dcn_instance_id"); ok {
		request.DcnInstanceId = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "init_params"); ok {
		var initParams []*dcdb.DBParamValue
		if v, ok := dMap["character_set_server"]; ok {
			initParams = append(initParams, &dcdb.DBParamValue{
				Param: helper.String("character_set_server"),
				Value: helper.String(v.(string)),
			})
		}

		if v, ok := dMap["lower_case_table_names"]; ok {
			initParams = append(initParams, &dcdb.DBParamValue{
				Param: helper.String("lower_case_table_names"),
				Value: helper.String(strconv.Itoa(v.(int))),
			})
		}

		if v, ok := dMap["sync_mode"]; ok {
			initParams = append(initParams, &dcdb.DBParamValue{
				Param: helper.String("sync_mode"),
				Value: helper.String(strconv.Itoa(v.(int))),
			})
		}

		request.InitParams = initParams
	}

	result, err := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().CreateHourDCDBInstance(request)
	if err != nil {
		log.Printf("[CRITAL]%s create dcdb hourdbInstance failed, reason:%+v", logId, err)
		return err
	} else {
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
	}
	if result.Response == nil || len(result.Response.InstanceIds) == 0 {
		return fmt.Errorf("create dcdb hourdbInstance failed, response body [%s]", result.ToJsonString())
	}
	response = result

	instanceId = *response.Response.InstanceIds[0]

	d.SetId(instanceId)

	if instanceId != "" {
		// need to wait dcn init processing complete
		// 0:none; 1:creating, 2:running
		conf := BuildStateChangeConf([]string{}, []string{"2"}, 20*readRetryTimeout, time.Second,
			service.DcdbInstanceStateRefreshFunc(instanceId, []string{}))
		if _, e := conf.WaitForStateContext(ctx); e != nil {
			return e
		}
	}

	//var (
	//	vip   string
	//	vipv6 string
	//)

	//if v, ok := d.GetOk("vip"); ok {
	//	vip = v.(string)
	//}
	//if v, ok := d.GetOk("vipv6"); ok {
	//	vipv6 = v.(string)
	//}

	/*
		if vip != "" || vipv6 != "" {
			if vpcId == "" || subnetId == "" {
				return fmt.Errorf("`vpc_id` and `subnet_id` cannot be empty when setting `vip` or `vipv6` fields!")
			}

			err := service.SetNetworkVip(ctx, instanceId, vpcId, subnetId, vip, vipv6)
			if err != nil {
				return err
			}
		}
	*/

	return resourceTencentCloudDcdbdbInstanceRead(d, meta)
}

func resourceTencentCloudDcdbdbInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	instanceId := d.Id()

	hourdbInstances, err := service.DescribeDcdbHourdbInstance(ctx, instanceId)

	if err != nil {
		return err
	}

	if hourdbInstances == nil {
		d.SetId("")
		return fmt.Errorf("resource `hourdbInstance` %s does not exist", instanceId)
	}

	if len(hourdbInstances) > 1 {
		d.SetId("")
		return fmt.Errorf("the count of resource `hourdbInstance` shoud not beyond one. count: %v\n",
			len(hourdbInstances))
	}

	hourdbInstance := hourdbInstances[0]

	if hourdbInstance.ShardDetail[0] != nil { // Memory and Storage is params for one shard
		shard := hourdbInstance.ShardDetail[0]

		if shard.Memory != nil {
			_ = d.Set("shard_memory", shard.Memory)
		}

		if shard.Storage != nil {
			_ = d.Set("shard_storage", shard.Storage)
		}
	}

	if hourdbInstance.NodeCount != nil {
		_ = d.Set("shard_node_count", hourdbInstance.NodeCount)
	}

	if hourdbInstance.ShardCount != nil {
		_ = d.Set("shard_count", hourdbInstance.ShardCount)
	}

	if hourdbInstance.UniqueVpcId != nil {
		_ = d.Set("vpc_id", hourdbInstance.UniqueVpcId)
	}

	if hourdbInstance.UniqueSubnetId != nil {
		_ = d.Set("subnet_id", hourdbInstance.UniqueSubnetId)
	}

	if hourdbInstance.DbVersion != nil {
		_ = d.Set("db_version_id", hourdbInstance.DbVersion)
	}

	if hourdbInstance.ProjectId != nil {
		_ = d.Set("project_id", hourdbInstance.ProjectId)
	}

	if hourdbInstance.InstanceName != nil {
		_ = d.Set("instance_name", hourdbInstance.InstanceName)
	}

	if hourdbInstance.Ipv6Flag != nil {
		_ = d.Set("ipv6_flag", hourdbInstance.Ipv6Flag)
	}

	if hourdbInstance.Status != nil {
		// 2 运行中
		if *hourdbInstance.Status == DCDB_STATUS_OPENED {
			_ = d.Set("extranet_access", true)
		} else {
			_ = d.Set("extranet_access", false)
		}
	}

	if sgResponse, err := service.DescribeDcdbSecurityGroup(ctx, instanceId); err == nil {
		sg := sgResponse.Response
		sgId := ""
		if sg != nil && len(sg.Groups) > 0 {
			sgId = *sg.Groups[0].SecurityGroupId
		}
		_ = d.Set("security_group_id", sgId)
	} else {
		return err
	}

	// set vip, vipv6 and vport dcn id and region
	if detailResponse, err := service.DescribeDcdbDbInstanceDetailById(ctx, instanceId); err == nil {
		detail := detailResponse.Response
		if detail != nil {
			_ = d.Set("vip", detail.Vip)
			_ = d.Set("vipv6", detail.Vip6)
			_ = d.Set("vport", detail.Vport)

			_ = d.Set("dcn_instance_id", detail.InstanceId)
			_ = d.Set("dcn_region", detail.Region)

			if detail.MasterZone != nil {
				zones := []*string{detail.MasterZone}
				if detail.SlaveZones != nil {
					zones = append(zones, detail.SlaveZones...)
				}
				_ = d.Set("zones", zones)
			}

			if detail.ResourceTags != nil {
				tags := make(map[string]string)
				for _, tag := range detail.ResourceTags {
					tags[*tag.TagKey] = *tag.TagValue
				}
				_ = d.Set("resource_tags", tags)
			}
		}
	} else {
		return err
	}

	return nil
}

func resourceTencentCloudDcdbdbInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_instance.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		request = dcdb.NewModifyDBInstanceNameRequest()
		service = DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	instanceId := d.Id()

	request.InstanceId = &instanceId

	if d.HasChange("zones") {
		return fmt.Errorf("`zones` do not support change now.")
	}

	if d.HasChange("shard_memory") {
		return fmt.Errorf("`shard_memory` do not support change now.")
	}

	if d.HasChange("shard_storage") {
		return fmt.Errorf("`shard_storage` do not support change now.")
	}

	if d.HasChange("shard_node_count") {
		return fmt.Errorf("`shard_node_count` do not support change now.")
	}

	if d.HasChange("shard_count") {
		return fmt.Errorf("`shard_count` do not support change now.")
	}

	if d.HasChange("db_version_id") {
		return fmt.Errorf("`db_version_id` do not support change now.")
	}

	//wait for ModifyDBInstanceSecurityGroups
	//if d.HasChange("security_group_id") {
	//	return fmt.Errorf("`security_group_id` do not support change now.")
	//}

	if d.HasChange("vpc_id") || d.HasChange("subnet_id") || d.HasChange("vip") || d.HasChange("vipv6") {
		var (
			vip      string
			vipv6    string
			vpcId    string
			subnetId string
		)
		if v, ok := d.GetOk("vip"); ok {
			vip = v.(string)
		}
		if v, ok := d.GetOk("vipv6"); ok {
			vipv6 = v.(string)
		}
		if v, ok := d.GetOk("vpc_id"); ok {
			vpcId = v.(string)
		}
		if v, ok := d.GetOk("subnet_id"); ok {
			subnetId = v.(string)
		}

		if vpcId == "" || subnetId == "" {
			return fmt.Errorf("`vpc_id` and `subnet_id` cannot be empty when updating network configs!")
		}

		err := service.SetNetworkVip(ctx, instanceId, vpcId, subnetId, vip, vipv6)
		if err != nil {
			return err
		}
	}

	if d.HasChange("dcn_region") {
		return fmt.Errorf("`dcn_region` do not support change now.")
	}
	if d.HasChange("dcn_instance_id") {
		return fmt.Errorf("`dcn_instance_id` do not support change now.")
	}

	if d.HasChange("resource_tags") {
		return fmt.Errorf("`resource_tags` do not support change now.")
	}

	if d.HasChange("instance_name") {
		if v, ok := d.GetOk("instance_name"); ok {
			request.InstanceName = helper.String(v.(string))
		}
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().ModifyDBInstanceName(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s create dcdb hourdbInstance failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudDcdbdbInstanceRead(d, meta)
}

func resourceTencentCloudDcdbdbInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_instance.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	instanceId := d.Id()

	if err := service.DeleteDcdbHourdbInstanceById(ctx, instanceId); err != nil {
		return err
	}

	return nil
}
