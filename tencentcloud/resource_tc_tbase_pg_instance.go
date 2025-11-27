/*
Provides a resource to create a tbase instance.

~> **NOTE:** This resource is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tbase_pg_instance" "example" {
	  vpc_id                 = "vpc-cs6ffr73"
	  subnet_id              = "subnet-38oi34ta"
	  net_type               = "vpc"
	  pool_id                = 1
	  instance_name          = "tf-test"
	  engine_type            = "PostgreSQL"
	  engine_version         = "10.5.2"
	  memory                 = 1
	  cpu                    = 1
	  admin_password         = "Xx5kjgw%RXKf"
	  charset                = "UTF8"
	  instance_count         = 1
	  instance_role          = "master"
	  master_node_zone       = "yfm18"
	  follow_nodes_zones     = ["yfm18"]
	  security_group_id_list = ["sg-9s7k6qgw", "sg-rrpst23s"]
	  pg_node {
	    storage   = 10
	    spec_code = "pg.dn.1g"
	  }
	  project_id = "pr-bae40f73"
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	tbase "terraform-provider-tencentcloudenterprise/sdk/tbase/v20190107"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tbase_pg_instance", CNDescription{
		TerraformTypeCN: "PGSQL实例",
		DescriptionCN:   "提供TBase PostgreSQL实例资源，用于创建和管理TBase PostgreSQL数据库实例。",
		AttributesCN: map[string]string{
			"vpc_id":                 "VPC实例ID",
			"subnet_id":              "子网实例ID",
			"net_type":               "网络类型，可选值为VPC或BASIC",
			"pool_id":                "资源池ID",
			"engine_type":            "引擎类型",
			"engine_version":         "引擎版本",
			"memory":                 "内存",
			"cpu":                    "中央处理器",
			"admin_password":         "管理员密码",
			"charset":                "字符集",
			"instance_count":         "实例数量",
			"instance_name":          "实例名称",
			"instance_role":          "实例类型",
			"master_node_zone":       "主节点的可用区",
			"follow_nodes_zones":     "从节点的可用区",
			"security_group_id_list": "实例的安全组ID列表",
			"pg_node":                "节点规格信息",
			"project_id":             "项目ID",
			"spec_code":              "要创建的pg节点的规范代码",
			"storage":                "要创建的pg节点的同步类型",
		},
	})
}
func resourceTencentCloudTbasePGInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a resource to create a tbase instance.",
		CreateContext: resourceTencentCloudTbasePGInstanceCreate,
		ReadContext:   resourceTencentCloudTbasePGInstanceRead,
		UpdateContext: resourceTencentCloudTbasePGInstanceUpdate,
		DeleteContext: resourceTencentCloudTbasePGInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
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
			"net_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Net type, optional value is VPC or BASIC.",
			},
			"pool_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Pool id.",
			},
			"engine_type": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "engine type. `PostgreSQL`, `TbaseXC`",
				ValidateFunc: validateAllowedStringValue(TBASE_PG_ENGINE_TYPE),
			},
			"engine_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Engine version.",
			},
			"memory": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Memory.",
			},
			"cpu": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Cpu.",
			},
			"admin_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Admin pass.",
				Sensitive:   true,
			},
			// 字符集，根据DescribeAvailableCharset查询
			"charset": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "charset.",
				ValidateFunc: validateAllowedStringValue(TBASE_CHARACTER_SET),
			},
			"instance_count": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateIntegerInRange(1, 10),
				Description:  "instance count.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance name.",
			},
			// 实例类型，当前只支持master类型
			"instance_role": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "instance role, current only support `master` type.",
				ValidateFunc: validateAllowedStringValue([]string{"master"}),
			},
			"master_node_zone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The zone of the master node.",
			},
			"follow_nodes_zones": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The zone of the follow nodes.",
			},
			"security_group_id_list": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The security group id list of the instance.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"pg_node": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Node specification information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spec_code": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The spec code of pg nodes to be created.",
						},
						"storage": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The sync type of pg nodes to be created.",
						},
					},
				},
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Project id.",
			},
		},
	}
}

func resourceTencentCloudTbasePGInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tbase_instance.create")()
	defer inconsistentCheck(d, meta)()

	var (
		request    = tbase.NewCreatePGInstanceHourRequest()
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId string
	)
	vpcId := d.Get("vpc_id")
	request.VpcId = helper.String(vpcId.(string))

	subnetId := d.Get("subnet_id")
	request.SubnetId = helper.String(subnetId.(string))

	netType := d.Get("net_type")
	request.NetType = helper.String(netType.(string))

	poolId := d.Get("pool_id")
	request.PoolId = helper.Int64(int64(poolId.(int)))

	engineType := d.Get("engine_type").(string)
	request.EngineType = &engineType

	engineVersion := d.Get("engine_version").(string)
	request.EngineVersion = &engineVersion

	memory := int64(d.Get("memory").(int))
	request.Memory = &memory

	cpu := int64(d.Get("cpu").(int))
	request.Cpu = &cpu

	if v, ok := d.GetOk("admin_password"); ok {
		request.AdminPassword = helper.String(v.(string))
	}
	charset := d.Get("charset").(string)
	request.Charset = &charset

	instanceCount := d.Get("instance_count")
	request.Count = helper.Int64(int64(instanceCount.(int)))

	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_role"); ok {
		request.InstanceRole = helper.String(v.(string))
	}

	masterNodeZone := d.Get("master_node_zone")
	request.MasterNodeZone = helper.String(masterNodeZone.(string))

	if v, ok := d.GetOk("follow_nodes_zones"); ok {
		followNodesZoneSet := v.(*schema.Set).List()
		followNodesZone := helper.InterfacesStringsPoint(followNodesZoneSet)
		request.FollowNodesZones = followNodesZone
	}

	group := d.Get("security_group_id_list")
	securityGroupIdsSet := group.(*schema.Set).List()
	for i := range securityGroupIdsSet {
		securityGroupIds := securityGroupIdsSet[i].(string)
		request.SecurityGroupIdList = append(request.SecurityGroupIdList, &securityGroupIds)
	}
	request.SecurityGroupIdList = helper.InterfacesStringsPoint(securityGroupIdsSet)

	if dMap, ok := helper.InterfacesHeadMap(d, "pg_node"); ok {
		pgNode := &tbase.PgNodeInfo{}
		if v, ok := dMap["storage"]; ok {
			pgNode.Storage = helper.Int64(int64(v.(int)))
		}
		if v, ok := helper.GetStringValue(dMap, "spec_code"); ok {
			pgNode.SpecCode = helper.String(v)
		}
		request.PgNode = pgNode
	}

	if v, ok := d.GetOk("project_id"); ok {
		request.ProjectId = helper.String(v.(string))
	}

	instanceId, err := service.CreatePGInstanceHour(ctx, request)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(instanceId)
	return resourceTencentCloudTbasePGInstanceRead(ctx, d, meta)
}

func resourceTencentCloudTbasePGInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tbase_instance.read")()
	defer inconsistentCheck(d, meta)()

	var (
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instance   *tbase.PGInstanceSet
		request    = tbase.NewDescribePGInstancesRequest()
		instanceId = d.Id()
	)

	// query the instance
	err := resource.RetryContext(ctx, 3*readRetryTimeout, func() *resource.RetryError {
		request.InstanceIds = []*string{&instanceId}
		instances, e := service.DescribePGInstances(ctx, request)
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
		return diag.FromErr(err)
	}

	if instance == nil {
		d.SetId("")
		return diag.Errorf("resource `instance` %s does not exist", instanceId)
	}

	if instance.InstanceName != nil {
		_ = d.Set("instance_name", *instance.InstanceName)
	}

	return nil
}

func resourceTencentCloudTbasePGInstanceUpdate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tbase_instance.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId = d.Id()
	)

	err := resource.RetryContext(ctx, writeRetryTimeout, func() *resource.RetryError {
		if d.HasChange("instance_name") {
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
		log.Printf("[CRITICAL]%s modify tbase failed, reason:%+v", logId, err)
		return diag.FromErr(err)
	}

	return resourceTencentCloudTbasePGInstanceRead(ctx, d, meta)
}

func resourceTencentCloudTbasePGInstanceDelete(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tbase.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId = d.Id()
		request    = tbase.NewIsolatePGInstanceHourRequest()
	)
	request.InstanceId = &instanceId

	if err := service.IsolatePGInstanceHour(ctx, request); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
