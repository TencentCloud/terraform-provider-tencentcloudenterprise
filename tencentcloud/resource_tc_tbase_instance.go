/*
Provides a resource to create a tbase instance.

~> **NOTE:** This resource is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tbase_instance" "instance1" {
	  admin_password    = "Tcdn@2007"
	  charset           = "UTF8"
	  engine_type        = "TbaseV5"
	  engine_version = "5.06.1.1"
	  goods_num = 1
	  net_type = "vpc"
	  security_group_id_list = ["sg-h45zkk4g"]
	  subnet_id = "subnet-ebbhkayi"
	  vpc_id = "vpc-1dz8zfpr"
	  zone = "yfdb1"
	  dn_node {
	    node_count = 4
	    set_count = 2
	    spec_code = "dn.v5.1g"
	    storage = 10
	    sync_type = "sync2async"
	  }
	  cn_node {
	    node_count = 4
	    set_count = 2
	    spec_code = "cn.v5.1g"
	    storage = 10
	  }
	}

```
Import

tbase instance can be imported using the id, e.g.
```
$ terraform import cloud_tbase_instance.instance cluster_id#instance_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	tbase "terraform-provider-tencentcloudenterprise/sdk/tbase/v20190107"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tbase_instance", CNDescription{
		TerraformTypeCN: "分布式PGSQL实例",
		DescriptionCN:   "提供TBase分布式PostgreSQL实例资源，用于创建和管理TBase分布式数据库实例。",
		AttributesCN: map[string]string{
			"zone":                   "可用区",
			"vpc_id":                 "VPC ID",
			"subnet_id":              "子网ID",
			"engine_version":         "引擎版本",
			"admin_password":         "管理员密码",
			"charset":                "字符集",
			"engine_type":            "引擎类型",
			"instance_name":          "实例名称",
			"net_type":               "网络类型",
			"cn_node":                "实例cn节点信息",
			"dn_node":                "实例dn节点信息",
			"goods_num":              "购买数量",
			"security_group_id_list": "安全组ID列表",
			"recovery":               "恢复实例信息",
			"tags":                   "标签列表",
			"project_id":             "项目ID",
			"slave_zones":            "备机可用区列表",
			"node_count":             "要创建的节点数",
			"recovery_instance_id":   "实例的恢复实例id",
			"recovery_time":          "实例的恢复时间",
			"recovery_time_point":    "实例的恢复时间点",
			"set_count":              "要创建的节点数",
			"spec_code":              "要创建的节点的规范代码",
			"storage":                "要创建的节点的存储",
			"sync_type":              "要创建的节点的同步类型",
		},
	})
}
func resourceTencentCloudTbaseInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tbase instance.",
		Create:      resourceTencentCloudTbaseInstanceCreate,
		Read:        resourceTencentCloudTbaseInstanceRead,
		Update:      resourceTencentCloudTbaseInstanceUpdate,
		Delete:      resourceTencentCloudTbaseInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone.",
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
			"engine_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Engine version.",
			},
			"admin_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Admin pass.",
				Sensitive:   true,
			},
			"charset": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Character set, based on DescribeAvailableCharset.",
				ValidateFunc: validateAllowedStringValue(TBASE_CHARACTER_SET),
			},
			"engine_type": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Engine type, optional value is `TbaseV2`, `TbaseV5`, `TbaseV5_C`, `TbaseH`,	`TbaseHC`.",
				ValidateFunc: validateAllowedStringValue(TBASE_ENGINE_TYPE),
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance name.",
			},
			"net_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Net type, optional value is VPC or BASIC.",
			},
			"cn_node": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Example Create specifications of the instance cn node.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The number of cn nodes to be created.",
						},
						"set_count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The number of cn nodes to be created.",
						},
						"storage": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The stroage of cn nodes to be created.",
						},
						"spec_code": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The spec code of cn nodes to be created.",
						},
						"sync_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The sync type of cn nodes to be created.",
						},
					},
				},
			},
			"dn_node": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Create the specification information related to the dn node of the instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The number of nodes to be created.",
						},
						"set_count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The number of nodes to be created.",
						},
						"storage": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The storage of nodes to be created.",
						},
						"spec_code": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The spec code of nodes to be created.",
						},
						"sync_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The sync type of nodes to be created.",
						},
					},
				},
			},

			"goods_num": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "Purchase quantity, (0,10), no default create 1.",
				ValidateFunc: validateIntegerInRange(1, 10),
			},
			"security_group_id_list": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The security group id list of the instance.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"recovery": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Mandatory parameter for creating a backfile instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"recovery_time_point": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The recovery time point of the instance.",
						},
						"recovery_time": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The recovery time of the instance.",
						},
						"recovery_instance_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The recovery instance id of the instance.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag list.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Project id.",
			},
			"slave_zones": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of slave zones.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceTencentCloudTbaseInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tbase_instance.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = tbase.NewCreateInstanceHourRequest()
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		instanceId string
	)
	if v, ok := d.GetOk("admin_password"); ok {
		request.AdminPassword = helper.String(v.(string))
	}

	if v, ok := d.GetOk("zone"); ok {
		request.Zone = helper.String(v.(string))
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		request.SubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("net_type"); ok {
		request.NetType = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "cn_node"); ok {
		cnCode := &tbase.NodeInfo{}
		if v, ok := dMap["node_count"]; ok {
			cnCode.NodeCount = helper.Int64(int64(v.(int)))
		}
		if v, ok := dMap["set_count"]; ok {
			cnCode.SetCount = helper.Int64(int64(v.(int)))
		}
		if v, ok := dMap["storage"]; ok {
			cnCode.Storage = helper.Int64(int64(v.(int)))
		}
		if v, ok := dMap["spec_code"]; ok {
			cnCode.SpecCode = helper.String(v.(string))
		}
		if v, ok := dMap["sync_type"]; ok {
			if v.(string) != "" {
				cnCode.SyncType = helper.String(v.(string))
			}
		}
		request.CnNode = cnCode
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "dn_node"); ok {
		dnCode := &tbase.NodeInfo{}
		if v, ok := dMap["node_count"]; ok {
			dnCode.NodeCount = helper.Int64(int64(v.(int)))
		}
		if v, ok := dMap["set_count"]; ok {
			dnCode.SetCount = helper.Int64(int64(v.(int)))
		}
		if v, ok := dMap["storage"]; ok {
			dnCode.Storage = helper.Int64(int64(v.(int)))
		}
		if v, ok := dMap["spec_code"]; ok {
			dnCode.SpecCode = helper.String(v.(string))
		}
		if v, ok := dMap["sync_type"]; ok {
			if v.(string) != "" {
				dnCode.SyncType = helper.String(v.(string))
			}
		}
		request.DnNode = dnCode
	}

	if v, ok := d.GetOk("goods_num"); ok {
		request.GoodsNum = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("security_group_id_list"); ok {
		securityGroupIdsSet := v.(*schema.Set).List()
		for i := range securityGroupIdsSet {
			securityGroupIds := securityGroupIdsSet[i].(string)
			request.SecurityGroupIdList = append(request.SecurityGroupIdList, &securityGroupIds)
		}
		request.SecurityGroupIdList = helper.InterfacesStringsPoint(securityGroupIdsSet)
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "recovery"); ok {
		recovery := &tbase.RecoveryInfo{}
		if v, ok := dMap["recovery_time"]; ok {
			recovery.RecoveryTime = helper.String(v.(string))
		}
		if v, ok := dMap["recovery_time_point"]; ok {
			recovery.RecoveryTimePoint = helper.String(v.(string))
		}
		if v, ok := dMap["recovery_instance_id"]; ok {
			recovery.RecoveryInstanceId = helper.String(v.(string))
		}
		request.Recovery = recovery
	}

	if v, ok := d.GetOk("tags"); ok {
		for key, value := range v.(map[string]interface{}) {
			resourceTag := &tbase.TagInfo{
				TagKey:   helper.String(key),
				TagValue: helper.String(value.(string)),
			}
			request.Tags = append(request.Tags, resourceTag)
		}
	}

	if v, ok := d.GetOk("project_id"); ok {
		request.ProjectId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("slave_zones"); ok {
		slaveZones := v.([]interface{})
		for i := range slaveZones {
			slaveZone := slaveZones[i].(string)
			request.SlaveZones = append(request.SlaveZones, &slaveZone)
		}
	}

	engineType := d.Get("engine_type").(string)
	request.EngineType = &engineType

	engineVersion := d.Get("engine_version").(string)
	request.EngineVersion = &engineVersion

	charset := d.Get("charset").(string)
	request.Charset = &charset

	instanceId, err := service.CreateInstanceHour(ctx, request)
	if err != nil {
		return err
	}
	d.SetId(instanceId)
	return resourceTencentCloudTbaseInstanceRead(d, meta)
}

func resourceTencentCloudTbaseInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tbase_instance.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instance   *tbase.Instance
		instanceId = d.Id()
	)

	// query the instance
	err := resource.RetryContext(ctx, 3*readRetryTimeout, func() *resource.RetryError {
		instances, e := service.DescribeTbaseInstance(ctx, &instanceId)
		if e != nil {
			return retryError(e)
		}

		if len(instances) != 0 {
			status := *instances[0].Status

			if status == "online" {
				instance = instances[0]
				return nil
			}

			if status == "creating" || status == "recovering" {
				return resource.RetryableError(fmt.Errorf(
					"tbase instance[%s] status is still creating or recovering, retry", d.Id()))
			}
			return resource.NonRetryableError(fmt.Errorf("tbase instance[%s] status is invalid, exit", d.Id()))
		}
		return resource.RetryableError(fmt.Errorf("can not get tbase instance[%s] status, retry", d.Id()))
	})
	if err != nil {
		return err
	}

	if instance == nil {
		d.SetId("")
		return fmt.Errorf("resource `instance` %s does not exist", instanceId)
	}

	if instance.InstanceName != nil {
		_ = d.Set("instance_name", instance.InstanceName)
	}

	return nil
}

func resourceTencentCloudTbaseInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tbase_instance.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId = d.Id()
	)

	err := resource.RetryContext(ctx, writeRetryTimeout, func() *resource.RetryError {
		if d.HasChange("instance_name") {
			req := tbase.NewModifyInstanceNameRequest()
			req.InstanceId = &instanceId
			req.InstanceName = helper.String(d.Get("instance_name").(string))

			e := service.ModifyInstanceName(ctx, &instanceId, helper.String(d.Get("instance_name").(string)))
			if e != nil {
				return retryError(e)
			}
		}

		if d.HasChange("security_group_id_list") {
			list := helper.InterfacesStringsPoint(d.Get("security_group_id_list").(*schema.Set).List())
			e := service.ModifyDBInstanceSecurityGroups(ctx, &instanceId, list)
			if e != nil {
				return retryError(e)
			}
		}

		if d.HasChange("admin_password") {
			req := tbase.NewResetAccountPasswordRequest()
			req.InstanceId = &instanceId
			req.Username = helper.String(TBASE_DATABASE_USERNAME)
			req.NewPassword = helper.String(d.Get("admin_password").(string))

			e := service.ResetAccountPassword(ctx, req)
			if e != nil {
				return retryError(e)
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITICAL]%s modify tbase instance name failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTbaseInstanceRead(d, meta)
}

func resourceTencentCloudTbaseInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tbase.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId = d.Id()
	)

	if err := service.TerminateInstanceById(ctx, &instanceId); err != nil {
		return err
	}

	return nil
}
