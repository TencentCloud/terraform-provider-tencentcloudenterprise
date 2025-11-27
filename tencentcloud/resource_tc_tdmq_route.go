/*
Provide a resource to create a TDMQ Route.

Example Usage
```hcl

	resource "tencentcloudenterprise_tdmq_Route" "foo" {
	  remark = "this is description111."
	  cluster_id = 0
	  net_type = 2
	}

```

Import
Tdmq Route can be imported, e.g.

```
$ terraform import cloud_tdmq_Route.test tdmq_id
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_route", CNDescription{
		TerraformTypeCN: "TDMQ 路由",
		DescriptionCN:   "提供TDMQ路由资源，用于创建和管理TDMQ路由。",
		AttributesCN: map[string]string{
			"net_type":         "网络类型",
			"cluster_id":       "集群ID",
			"remark":           "备注",
			"unique_vpc_id":    "唯一VPC ID",
			"unique_subnet_id": "唯一子网 ID",
			"router_id":        "路由ID",
			"vip":              "路由VIP",
			"vport":            "路由VPORT",
			"encrypt_type":     "是否加密",
		},
	})
}
func resourceTencentCloudTdmqRoute() *schema.Resource {
	return &schema.Resource{
		Description:   "Provide a resource to create a TDMQ Route.",
		CreateContext: resourceTencentCloudTdmqRouteCreate,
		ReadContext:   resourceTencentCloudTdmqRouteRead,
		UpdateContext: resourceTencentCloudTdmqRouteUpdate,
		DeleteContext: resourceTencentCloudTdmqRouteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Dedicated Cluster Id.",
			},
			"net_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     2,
				Description: "Net type.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Remark of the route.",
			},
			"encrypt_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Encrypt type of route. 0: no encryption, 1: TLS encrypt. Default: 0.",
			},

			// Computed
			"unique_vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Uniq vpc id.",
			},
			"unique_subnet_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 128),
				Description:  "uniq subnet id, the value contains a maximum of 128 characters.",
			},
			"router_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the route.",
			},
			"vip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The vip of the route.",
			},
			"vport": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The vport of the route.",
			},
		},
	}
}

func resourceTencentCloudTdmqRouteCreate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_Route.create")()

	logId := getLogId(contextNil)

	var (
		request  = tdmq.NewCreateRouteRequest()
		response *tdmq.CreateRouteResponse
	)
	request.ClusterId = helper.String(d.Get("cluster_id").(string))

	request.NetType = helper.IntUint64(d.Get("net_type").(int))

	if v, ok := d.GetOk("remark"); ok {
		request.Remark = helper.String(v.(string))
	}

	if v, ok := d.GetOk("unique_vpc_id"); ok {
		request.UniqueVpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("unique_subnet_id"); ok {
		request.UniqueSubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("encrypt_type"); ok {
		request.EncryptType = helper.IntUint64(v.(int))
	} else {
		request.EncryptType = helper.IntUint64(0)
	}
	log.Printf(" request body, %#v\n", request)

	err := resource.RetryContext(ctx, writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().CreateRoute(request)
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
		log.Printf("[CRITAL]%s create tdmq Route failed, reason:%+v", logId, err)
		return diag.FromErr(err)
	}

	routerId := *response.Response.RouterId

	d.SetId(routerId)

	return resourceTencentCloudTdmqRouteRead(ctx, d, meta)
}

func resourceTencentCloudTdmqRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_Route.read")()
	defer inconsistentCheck(d, meta)()

	var (
		clusterId = d.Get("cluster_id").(string)
		id        = d.Id()
	)

	tdmqService := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	info, has, err := tdmqService.DescribeTdmqRouteById(ctx, clusterId, id)
	if err != nil {
		return diag.FromErr(err)
	}
	if !has {
		d.SetId("")
		return nil
	}

	_ = d.Set("remark", info.Remark)
	_ = d.Set("vip", info.Vip)
	_ = d.Set("vport", info.Vport)
	_ = d.Set("encrypt_type", info.EncryptType)
	_ = d.Set("unique_vpc_id", info.VpcId)
	_ = d.Set("unique_subnet_id", info.SubnetId)

	return nil
}

func resourceTencentCloudTdmqRouteUpdate(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_Route.update")()

	return nil
}

func resourceTencentCloudTdmqRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("resource.cloud_tdmq_Route.delete")()

	var (
		request = tdmq.NewDeleteRouteRequest()
	)

	service := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	request.RouteId = helper.String(d.Id())
	request.NetType = helper.IntUint64(d.Get("net_type").(int))
	request.ClusterId = helper.String(d.Get("cluster_id").(string))
	request.UniqueVpcId = helper.String(d.Get("unique_vpc_id").(string))
	request.UniqueSubnetId = helper.String(d.Get("unique_subnet_id").(string))

	if err := service.DeleteTdmqRoute(ctx, request); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
