/*
Manages a pay-as-you-go MariaDB instance.

# Example Usage

```hcl
variable "zones" { type = set(string) }

	resource "tencentcloudenterprise_mariadb_hour_db_instance" "example" {
	  memory = 2
	  node_count = 2
	  storage = 10
	  zones = var.zones
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
	mariadb "terraform-provider-tencentcloudenterprise/sdk/mariadb/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func resourceTencentCloudMariadbHourDbInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudMariadbHourDbInstanceCreate,
		Read:   resourceTencentCloudMariadbHourDbInstanceRead,
		Update: resourceTencentCloudMariadbHourDbInstanceUpdate,
		Delete: resourceTencentCloudMariadbHourDbInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zones": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				Description: "available zone of instance.",
			},

			"node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "number of node for instance.",
			},

			"memory": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "instance memory.",
			},

			"storage": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "instance disk storage.",
			},

			"project_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "project id.",
			},

			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "vpc id.",
			},

			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "subnet id, it's required when vpc_id is set.",
			},

			"vip": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "vip.",
			},

			"db_version_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Database engine version. Use DescribeDBEngines to find MariaDB versions available in the region.",
			},

			"cpu_arch": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue([]string{"X86", "ARM"}),
				Description:  "CPU architecture of the instance. Valid values: `X86`, `ARM`.",
			},

			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "name of this instance.",
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag description list.",
			},
		},
	}
}

func resourceTencentCloudMariadbHourDbInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_mariadb_hour_db_instance.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = MariadbService{client: meta.(*TencentCloudClient).apiV3Conn}
		request    = mariadb.NewCreateHourDBInstanceRequest()
		response   *mariadb.CreateHourDBInstanceResponse
		instanceId string
	)

	request.Count = helper.IntInt64(1)
	if v, ok := d.GetOk("zones"); ok {
		zonesSet := v.(*schema.Set).List()
		for i := range zonesSet {
			zones := zonesSet[i].(string)
			request.Zones = append(request.Zones, &zones)
		}
	}

	if v, ok := d.GetOk("node_count"); ok {
		request.NodeCount = helper.IntInt64(v.(int))
		// A multi-node deployment in one zone still needs a replica zone in the order.
		if v.(int) > 1 && len(request.Zones) == 1 {
			request.Zones = append(request.Zones, request.Zones[0])
		}
	}

	if v, ok := d.GetOk("memory"); ok {
		request.Memory = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("storage"); ok {
		request.Storage = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("project_id"); ok {
		request.ProjectId = helper.IntInt64(v.(int))
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

	if v, ok := d.GetOk("cpu_arch"); ok {
		request.CpuArch = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseMariadbClient().CreateHourDBInstance(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create mariadb hourDbInstance failed, reason:%+v", logId, err)
		return err
	}

	if response == nil || response.Response == nil || len(response.Response.InstanceIds) != 1 || response.Response.InstanceIds[0] == nil || *response.Response.InstanceIds[0] == "" {
		return fmt.Errorf("MariaDB create response is missing the instance ID")
	}
	instanceId = *response.Response.InstanceIds[0]
	d.SetId(instanceId)

	initParams := []*mariadb.DBParamValue{
		{
			Param: helper.String("character_set_server"),
			Value: helper.String("utf8mb4"),
		},
		{
			Param: helper.String("lower_case_table_names"),
			Value: helper.String("1"),
		},
		{
			Param: helper.String("sync_mode"),
			Value: helper.String("2"),
		},
		{
			Param: helper.String("innodb_page_size"),
			Value: helper.String("16384"),
		},
	}

	initRet, err := service.InitDbInstance(ctx, instanceId, initParams)
	if err != nil {
		return err
	}

	if !initRet {
		return fmt.Errorf("db instance init failed")
	}

	// CreateHourDBInstance does not accept a custom VIP directly; apply it via
	// ModifyInstanceNetwork when the user explicitly configured vip.
	if v, ok := d.GetOk("vip"); ok {
		desiredVip := v.(string)
		detail, err := service.DescribeMariadbDbInstanceDetail(ctx, instanceId)
		if err != nil {
			return err
		}
		if detail == nil || detail.Vip == nil || *detail.Vip != desiredVip {
			vpcId := d.Get("vpc_id").(string)
			subnetId := d.Get("subnet_id").(string)
			if vpcId == "" && detail != nil && detail.VpcId != nil {
				vpcId = *detail.VpcId
			}
			if subnetId == "" && detail != nil && detail.SubnetId != nil {
				subnetId = *detail.SubnetId
			}
			if err := modifyMariadbHourDbInstanceNetwork(ctx, logId, meta.(*TencentCloudClient).apiV3Conn, service, instanceId, desiredVip, vpcId, subnetId); err != nil {
				return err
			}
		}
	}

	// set Tags
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::mariadb:%s:uin/:mariadb-hour-instance/%s", region, instanceId)
		if err = tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudMariadbHourDbInstanceRead(d, meta)
}

func resourceTencentCloudMariadbHourDbInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_mariadb_hour_db_instance.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = MariadbService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId = d.Id()
	)

	hourDbInstance, err := service.DescribeMariadbDbInstanceDetail(ctx, instanceId)
	if err != nil {
		return err
	}

	if hourDbInstance == nil {
		d.SetId("")
		return nil
	}

	if hourDbInstance.Zone != nil {
		var zones []*string
		zones = append(zones, hourDbInstance.Zone)
		if hourDbInstance.SlaveZones != nil {
			zones = append(zones, hourDbInstance.SlaveZones...)
		}
		_ = d.Set("zones", zones)
	}

	if hourDbInstance.NodeCount != nil {
		_ = d.Set("node_count", hourDbInstance.NodeCount)
	}

	if hourDbInstance.Memory != nil {
		_ = d.Set("memory", hourDbInstance.Memory)
	}

	if hourDbInstance.Storage != nil {
		_ = d.Set("storage", hourDbInstance.Storage)
	}

	if hourDbInstance.ProjectId != nil {
		_ = d.Set("project_id", hourDbInstance.ProjectId)
	}

	if hourDbInstance.Vip != nil {
		_ = d.Set("vip", hourDbInstance.Vip)
	}

	if hourDbInstance.VpcId != nil {
		_ = d.Set("vpc_id", hourDbInstance.VpcId)
	}

	if hourDbInstance.SubnetId != nil {
		_ = d.Set("subnet_id", hourDbInstance.SubnetId)
	}

	if hourDbInstance.DbVersionId != nil && *hourDbInstance.DbVersionId != "" {
		_ = d.Set("db_version_id", hourDbInstance.DbVersionId)
	} else if hourDbInstance.DbVersion != nil && *hourDbInstance.DbVersion != "" {
		currentVersion := d.Get("db_version_id").(string)
		if currentVersion == "" || (*hourDbInstance.DbVersion != currentVersion && !strings.HasPrefix(*hourDbInstance.DbVersion, currentVersion+".")) {
			_ = d.Set("db_version_id", hourDbInstance.DbVersion)
		}
	}

	if hourDbInstance.CpuArch != nil && *hourDbInstance.CpuArch != "" {
		_ = d.Set("cpu_arch", strings.ToUpper(*hourDbInstance.CpuArch))
	}

	if hourDbInstance.InstanceName != nil {
		_ = d.Set("instance_name", hourDbInstance.InstanceName)
	}

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := &TagService{client: tcClient}
	tags, err := tagService.DescribeResourceTags(ctx, "mariadb", "mariadb-hour-instance", tcClient.Region, instanceId)
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudMariadbHourDbInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_mariadb_hour_db_instance.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = MariadbService{client: meta.(*TencentCloudClient).apiV3Conn}
		request    = mariadb.NewModifyDBInstanceNameRequest()
		instanceId = d.Id()
	)

	request.InstanceId = &instanceId

	if d.HasChange("zones") {
		return fmt.Errorf("`zones` do not support change now.")
	}

	if d.HasChange("node_count") {
		return fmt.Errorf("`node_count` do not support change now.")
	}

	if d.HasChange("memory") {
		return fmt.Errorf("`memory` do not support change now.")
	}

	if d.HasChange("storage") {
		return fmt.Errorf("`storage` do not support change now.")
	}

	if d.HasChange("vpc_id") {
		return fmt.Errorf("`vpc_id` do not support change now.")
	}

	if d.HasChange("subnet_id") {
		return fmt.Errorf("`subnet_id` do not support change now.")
	}

	if d.HasChange("db_version_id") {
		return fmt.Errorf("`db_version_id` do not support change now.")
	}

	if d.HasChange("instance_name") {
		if v, ok := d.GetOk("instance_name"); ok {
			request.InstanceName = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseMariadbClient().ModifyDBInstanceName(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s create mariadb hourDbInstance failed, reason:%+v", logId, err)
			return err
		}
	}

	if d.HasChange("tags") {
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := &TagService{client: tcClient}
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("mariadb", "mariadb-hour-instance", tcClient.Region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	if d.HasChange("project_id") {
		if v, ok := d.GetOkExists("project_id"); ok {
			projectId := int64(v.(int))
			MPRequest := mariadb.NewModifyDBInstancesProjectRequest()
			MPRequest.InstanceIds = []*string{&instanceId}
			MPRequest.ProjectId = &projectId

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseMariadbClient().ModifyDBInstancesProject(MPRequest)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				}

				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s operate mariadb modifyInstanceProject failed, reason:%+v", logId, err)
				return err
			}
		}
	}

	if d.HasChange("vip") {
		if v, ok := d.GetOk("vip"); ok {
			if err := modifyMariadbHourDbInstanceNetwork(
				ctx,
				logId,
				meta.(*TencentCloudClient).apiV3Conn,
				service,
				instanceId,
				v.(string),
				d.Get("vpc_id").(string),
				d.Get("subnet_id").(string),
			); err != nil {
				return err
			}
		}
	}

	return resourceTencentCloudMariadbHourDbInstanceRead(d, meta)
}

func modifyMariadbHourDbInstanceNetwork(
	ctx context.Context,
	logId string,
	conn *connectivity.TencentCloudClient,
	service MariadbService,
	instanceId, vip, vpcId, subnetId string,
) error {
	var vipFlowId int64
	vipRequest := mariadb.NewModifyInstanceNetworkRequest()
	vipRequest.InstanceId = &instanceId
	vipRequest.Vip = &vip
	if vpcId != "" {
		vipRequest.VpcId = helper.String(vpcId)
	}
	if subnetId != "" {
		vipRequest.SubnetId = helper.String(subnetId)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := conn.UseMariadbClient().ModifyInstanceNetwork(vipRequest)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, vipRequest.GetAction(), vipRequest.ToJsonString(), result.ToJsonString())
		if result != nil && result.Response != nil && result.Response.FlowId != nil {
			vipFlowId = *result.Response.FlowId
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate mariadb network failed, reason:%+v", logId, err)
		return err
	}

	if vipFlowId != NONE_FLOW_TASK {
		err = resource.Retry(10*writeRetryTimeout, func() *resource.RetryError {
			result, e := service.DescribeFlowById(ctx, vipFlowId)
			if e != nil {
				return retryError(e)
			}
			if result == nil || result.Status == nil {
				return resource.RetryableError(fmt.Errorf("operate mariadb network flow %d has no status yet", vipFlowId))
			}
			if *result.Status == MARIADB_TASK_SUCCESS {
				return nil
			} else if *result.Status == MARIADB_TASK_RUNNING {
				return resource.RetryableError(fmt.Errorf("operate mariadb network status is running"))
			} else if *result.Status == MARIADB_TASK_FAIL {
				return resource.NonRetryableError(fmt.Errorf("operate mariadb network status is fail"))
			}
			return resource.NonRetryableError(fmt.Errorf("operate mariadb network status illegal"))
		})
		if err != nil {
			log.Printf("[CRITAL]%s operate mariadb network task failed, reason:%+v", logId, err)
			return err
		}
	}
	return nil
}

func resourceTencentCloudMariadbHourDbInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_mariadb_hour_db_instance.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := MariadbService{client: meta.(*TencentCloudClient).apiV3Conn}

	instanceId := d.Id()

	if err := service.DeleteMariadbHourDbInstanceById(ctx, instanceId); err != nil {
		return err
	}

	return nil
}
