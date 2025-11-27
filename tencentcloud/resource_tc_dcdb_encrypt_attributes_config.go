/*
Provides a resource to create a dcdb encrypt_attributes_config

~> **NOTE:**  This resource currently only supports the newly created MySQL 8.0.24 version.

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_security_groups" "internal" {
		name = "default"
	}

	data "tencentcloudenterprise_vpc_instances" "vpc" {
		name ="Default-VPC"
	}

	data "tencentcloudenterprise_vpc_subnets" "subnet" {
		vpc_id = data.cloud_vpc_instances.vpc.instance_list.0.vpc_id
	}

	locals {
		vpc_id = data.cloud_vpc_subnets.subnet.instance_list.0.vpc_id
		subnet_id = data.cloud_vpc_subnets.subnet.instance_list.0.subnet_id
		sg_id = data.cloud_vpc_security_groups.internal.security_groups.0.security_group_id
	}

	resource "tencentcloudenterprise_dcdb_db_instance" "prepaid_instance" {
		instance_name = "test_dcdb_db_post_instance"
		zones = [var.default_az]
		period = 1
		shard_memory = "2"
		shard_storage = "10"
		shard_node_count = "2"
		shard_count = "2"
		vpc_id = local.vpc_id
		subnet_id = local.subnet_id
		db_version_id = "8.0"
		resource_tags {
		  tag_key = "aaa"
		  tag_value = "bbb"
		}
		security_group_ids = [local.sg_id]
	}

	resource "tencentcloudenterprise_dcdb_instance" "hourdb_instance" {
		instance_name = "test_dcdb_db_hourdb_instance"
		zones = [var.default_az]
		shard_memory = "2"
		shard_storage = "10"
		shard_node_count = "2"
		shard_count = "2"
		vpc_id = local.vpc_id
		subnet_id = local.subnet_id
		security_group_id = local.sg_id
		db_version_id = "8.0"
		resource_tags {
		  tag_key = "aaa"
		  tag_value = "bbb"
		}
	}

	locals {
		prepaid_dcdb_id = cloud_dcdb_db_instance.prepaid_instance.id
		hourdb_dcdb_id = cloud_dcdb_instance.hourdb_instance.id
	}

// for postpaid instance

	resource "tencentcloudenterprise_dcdb_encrypt_attributes_config" "config_hourdb" {
	  instance_id = local.hourdb_dcdb_id
	  encrypt_enabled = 1
	}

// for prepaid instance

	resource "tencentcloudenterprise_dcdb_encrypt_attributes_config" "config_prepaid" {
	  instance_id = local.prepaid_dcdb_id
	  encrypt_enabled = 1
	}

```

# Import

dcdb encrypt_attributes_config can be imported using the id, e.g.

```
terraform import cloud_dcdb_encrypt_attributes_config.encrypt_attributes_config encrypt_attributes_config_id
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
	registerResourceDescriptionProvider("tencentcloudenterprise_dcdb_encrypt_attributes_config", CNDescription{
		TerraformTypeCN: "修改实例数据加密",
		DescriptionCN:   "提供DCDB实例数据加密配置资源，用于启用或配置DCDB实例的数据加密功能。",
		AttributesCN: map[string]string{
			"instance_id":     "实例id",
			"encrypt_enabled": "是否启用数据加密注意：开启后不支持关闭，可选值：0-禁用，1-启用",
		},
	})
}

func resourceTencentCloudDcdbEncryptAttributesConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a dcdb encrypt_attributes_config",
		Create:      resourceTencentCloudDcdbEncryptAttributesConfigCreate,
		Read:        resourceTencentCloudDcdbEncryptAttributesConfigRead,
		Update:      resourceTencentCloudDcdbEncryptAttributesConfigUpdate,
		Delete:      resourceTencentCloudDcdbEncryptAttributesConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Instance id.",
			},

			"encrypt_enabled": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.",
			},
		},
	}
}

func resourceTencentCloudDcdbEncryptAttributesConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_encrypt_attributes_config.create")()
	defer inconsistentCheck(d, meta)()

	var instanceId string
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
	}
	d.SetId(instanceId)

	return resourceTencentCloudDcdbEncryptAttributesConfigUpdate(d, meta)
}

func resourceTencentCloudDcdbEncryptAttributesConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_encrypt_attributes_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	instanceId := d.Id()

	encryptAttributesConfig, err := service.DescribeDcdbEncryptAttributesConfigById(ctx, instanceId)
	if err != nil {
		return err
	}

	if encryptAttributesConfig == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DcdbEncryptAttributesConfig` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("instance_id", instanceId)

	if encryptAttributesConfig.Response.EncryptStatus != nil {
		_ = d.Set("encrypt_enabled", encryptAttributesConfig.Response.EncryptStatus)
	}

	return nil
}

func resourceTencentCloudDcdbEncryptAttributesConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_encrypt_attributes_config.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := dcdb.NewModifyDBEncryptAttributesRequest()

	instanceId := d.Id()

	request.InstanceId = &instanceId

	if d.HasChange("encrypt_enabled") {
		if v, ok := d.GetOkExists("encrypt_enabled"); ok {
			request.EncryptEnabled = helper.IntInt64(v.(int))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().ModifyDBEncryptAttributes(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update dcdb encryptAttributesConfig failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudDcdbEncryptAttributesConfigRead(d, meta)
}

func resourceTencentCloudDcdbEncryptAttributesConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_encrypt_attributes_config.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
