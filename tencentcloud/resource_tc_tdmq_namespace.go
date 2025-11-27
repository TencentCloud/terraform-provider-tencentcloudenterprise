/*
Provide a resource to create a tdmq namespace.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tdmq_instance" "foo" {
	  cluster_name = "example111"
	  remark = "this is description111."
	  tags = {
	    "createdBy" = "terraform"
	    "test" = "111"
	  }
	  bind_cluster_id = 0
	  bind_cluster_name = "default"
	}

	resource "tencentcloudenterprise_tdmq_namespace" "bar" {
	  environ_name = "example"
	  msg_ttl = 200
	  cluster_id = cloud_tdmq_instance.foo.id
	  remark = "this is description.22222222"
	}

```

# Import

Tdmq namespace can be imported, e.g.

```
$ terraform import cloud_tdmq_instance.test namespace_id
```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_namespace", CNDescription{
		TerraformTypeCN: "命名空间",
		DescriptionCN:   "提供TDMQ命名空间资源，用于创建和管理TDMQ命名空间。",
		AttributesCN: map[string]string{
			"environ_name": "TDMQ命名空间名称",
			"msg_ttl":      "消息过期时间,单位(秒)",
			"remark":       "备注",
			"tags":         "标签",
			"cluster_id":   "集群ID",
			//"retention_policy": "消息保留策略",
		},
	})
}
func RetentionPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"time_in_minutes": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The time of message to retain.",
		},
		"size_in_mb": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The size of message to retain.",
		},
	}
}

func resourceTencentCloudTdmqNamespace() *schema.Resource {
	return &schema.Resource{
		Description:   "Provide a resource to create a tdmq namespace.",
		CreateContext: resourceTencentCloudTdmqNamespaceCreate,
		ReadContext:   resourceTencentCloudTdmqNamespaceRead,
		UpdateContext: resourceTencentCloudTdmqNamespaceUpdate,
		DeleteContext: resourceTencentCloudTdmqNamespaceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"environ_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of namespace to be created.",
			},
			"msg_ttl": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The expiration time of unconsumed message.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Dedicated Cluster Id.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the namespace.",
			},
			"retention_policy": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional:    true,
				Description: "The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `time_in_minutes`: the time of message to retain; `size_in_mb`: the size of message to retain.",
			},
		},
	}
}

func resourceTencentCloudTdmqNamespaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_namespace.create")()

	tdmqService := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		environ_name    string
		msg_ttl         uint64
		remark          string
		clusterId       string
		retentionPolicy = new(tdmq.RetentionPolicy)
	)

	if temp, ok := d.GetOk("environ_name"); ok {
		environ_name = temp.(string)
		if len(environ_name) < 1 {
			return diag.Errorf("environ_name should be not empty string")
		}
	}

	msg_ttl = uint64(d.Get("msg_ttl").(int))

	if temp, ok := d.GetOk("cluster_id"); ok {
		clusterId = temp.(string)
	}

	if temp, ok := d.GetOk("remark"); ok {
		remark = temp.(string)
	}

	if temp, ok := d.GetOk("retention_policy"); ok {
		v := temp.(map[string]interface{})
		timeInMinutes := int64(v["time_in_minutes"].(int))
		sizeInMb := int64(v["size_in_mb"].(int))

		retentionPolicy.TimeInMinutes = &timeInMinutes
		retentionPolicy.SizeInMB = &sizeInMb
	}
	environId, err := tdmqService.CreateTdmqNamespace(ctx, environ_name, msg_ttl, clusterId, remark, *retentionPolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(environId)

	return resourceTencentCloudTdmqNamespaceRead(ctx, d, meta)
}

func resourceTencentCloudTdmqNamespaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_namespace.read")()
	defer inconsistentCheck(d, meta)()

	environId := d.Id()
	clusterId := d.Get("cluster_id").(string)

	tdmqService := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		info, has, e := tdmqService.DescribeTdmqNamespaceById(ctx, environId, clusterId)
		if e != nil {
			return retryError(e)
		}
		if !has {
			d.SetId("")
			return nil
		}

		_ = d.Set("environ_name", info.EnvironmentId)
		_ = d.Set("msg_ttl", info.MsgTTL)
		_ = d.Set("remark", info.Remark)

		retentionPolicy := make(map[string]interface{}, 2)
		retentionPolicy["time_in_minutes"] = *info.RetentionPolicy.TimeInMinutes
		retentionPolicy["size_in_mb"] = *info.RetentionPolicy.SizeInMB
		_ = d.Set("retention_policy", retentionPolicy)
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceTencentCloudTdmqNamespaceUpdate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_instance.update")()

	environId := d.Id()
	clusterId := d.Get("cluster_id").(string)

	service := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		msgTtl       uint64
		remark       string
		retentPolicy = new(tdmq.RetentionPolicy)
	)

	old, now := d.GetChange("msg_ttl")
	if d.HasChange("msg_ttl") {
		msgTtl = uint64(now.(int))
	} else {
		msgTtl = uint64(old.(int))
	}

	old, now = d.GetChange("remark")
	if d.HasChange("remark") {
		remark = now.(string)
	} else {
		remark = old.(string)
	}

	_, now = d.GetChange("retention_policy")
	if d.HasChange("retention_policy") {
		temp := now.(map[string]interface{})
		if timeValue, ok := temp["time_in_minutes"]; ok && timeValue != nil {
			time := timeValue.(int64)
			retentPolicy.TimeInMinutes = &time
		} else {
			retentPolicy.TimeInMinutes = helper.Int64(0)
		}

		if sizeValue, ok := temp["size_in_mb"]; ok && sizeValue != nil {
			size := sizeValue.(int64)
			retentPolicy.SizeInMB = &size
		} else {
			retentPolicy.SizeInMB = helper.Int64(0)
		}

		retentPolicy = &tdmq.RetentionPolicy{
			TimeInMinutes: retentPolicy.TimeInMinutes,
			SizeInMB:      retentPolicy.SizeInMB,
		}
	}

	d.Partial(true)
	if err := service.ModifyTdmqNamespaceAttribute(ctx, environId, msgTtl, remark, clusterId, retentPolicy); err != nil {
		return diag.FromErr(err)
	}

	d.Partial(false)
	return resourceTencentCloudTdmqNamespaceRead(ctx, d, meta)
}

func resourceTencentCloudTdmqNamespaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_instance.delete")()

	environId := d.Id()
	clusterId := d.Get("cluster_id").(string)

	service := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.RetryContext(ctx, writeRetryTimeout, func() *resource.RetryError {
		if err := service.DeleteTdmqNamespace(ctx, environId, clusterId); err != nil {
			if sdkErr, ok := err.(*errors.CloudSDKError); ok {
				if sdkErr.Code == VPCNotFound {
					return nil
				}
			}
			return resource.RetryableError(err)
		}
		return nil
	})

	return diag.FromErr(err)
}
