/*
Provides a resource to create a tbase instance.

~> **NOTE:** This resource is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tbase_pg_instance_vip" "foo" {
	  instance_id = cloud_tbase_pg_instance.example.id
	  uniq_subnet_id = "subnet-38oi34ta"
	  uniq_vpc_id = "vpc-cs6ffr73"
	}

```
Import

tbase instance can be imported using the id, e.g.
```
$ terraform import cloud_tbase_pg_instance_vip.instance cluster_id#instance_id
```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tbase "terraform-provider-tencentcloudenterprise/sdk/tbase/v20190107"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"time"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tbase_pg_instance_vip", CNDescription{
		TerraformTypeCN: "PGSQL实例只读vip",
		DescriptionCN:   "提供TBase PostgreSQL实例只读VIP资源，用于创建和管理PostgreSQL实例的只读VIP。",
		AttributesCN: map[string]string{
			"instance_id":    "实例ID",
			"uniq_subnet_id": "子网ID",
			"uniq_vpc_id":    "VPC ID",
			"vip":            "只读vip",
			"vport":          "只读vport",
		},
	})
}

func resourceTencentCloudTbasePGInstanceVip() *schema.Resource {
	return &schema.Resource{
		Description:   "Create read only vip for pg instance.",
		CreateContext: resourceTencentCloudTbasePGInstanceVipCreate,
		ReadContext:   resourceTencentCloudTbasePGInstanceVipRead,
		UpdateContext: resourceTencentCloudTbasePGInstanceVipUpdate,
		DeleteContext: resourceTencentCloudTbasePGInstanceVipDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance id.",
			},
			"uniq_subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subnet id.",
			},
			"uniq_vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vpc id.",
			},

			// Computed
			"vip": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Vip.",
			},
			"vport": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Vport.",
			},
		},
	}
}

func resourceTencentCloudTbasePGInstanceVipCreate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tbase_instance_vip.create")()
	defer inconsistentCheck(d, meta)()

	var (
		request    = tbase.NewCreatePGReadOnlyVipRequest()
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId = d.Get("instance_id")
	)
	vpcId := d.Get("uniq_vpc_id")
	request.UniqVpcId = helper.String(vpcId.(string))

	subnetId := d.Get("uniq_subnet_id")
	request.UniqSubnetId = helper.String(subnetId.(string))

	request.InstanceId = helper.String(instanceId.(string))

	if v, ok := d.GetOk("vip"); ok {
		request.Vip = helper.String(v.(string))
	}

	taskId, err := service.CreatePGReadOnlyVip(ctx, request)
	if err != nil {
		return diag.FromErr(err)
	}
	if taskId != 0 {
		// 任务的执行状态，有“success”，“running”，“failure”三种状态
		conf := BuildStateChangeConf([]string{}, []string{"running"}, 3*readRetryTimeout, time.Second,
			service.InstanceStateRefreshFunc(taskId, []string{}))
		if _, e := conf.WaitForStateContext(ctx); e != nil {
			return diag.FromErr(e)
		}
	}
	d.SetId(d.Get("instance_id").(string))
	return resourceTencentCloudTbasePGInstanceVipRead(ctx, d, meta)
}

func resourceTencentCloudTbasePGInstanceVipRead(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tbase_instance_vip.read")()
	defer inconsistentCheck(d, meta)()

	var (
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instance   *tbase.DescribePGInstanceDetailsResponse
		instanceId = d.Id()
	)

	// query the instance
	err := resource.RetryContext(ctx, 3*readRetryTimeout, func() *resource.RetryError {
		instances, e := service.DescribePGInstanceDetails(ctx, instanceId)
		if e != nil {
			return retryError(e)
		}

		instance = instances
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	info := instance.Response.RoVpc

	if info.Vip != nil {
		_ = d.Set("vip", info.Vip)
	}

	if info.Vport != nil {
		_ = d.Set("vport", info.Vport)
	}

	return nil
}

func resourceTencentCloudTbasePGInstanceVipUpdate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tbase_instance_vip.update")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudTbasePGInstanceVipDelete(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tbase_pg_instance_vip.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		service    = TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}
		instanceId = d.Id()
		request    = tbase.NewDeletePGReadOnlyVipRequest()
	)
	request.InstanceId = &instanceId
	request.Vip = helper.String(d.Get("vip").(string))
	request.UniqVpcId = helper.String(d.Get("uniq_vpc_id").(string))
	request.UniqSubnetId = helper.String(d.Get("uniq_subnet_id").(string))

	if err := service.DeletePGReadOnlyVip(ctx, request); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
