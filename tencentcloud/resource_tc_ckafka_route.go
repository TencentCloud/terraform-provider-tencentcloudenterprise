/*
Provides a resource to create a ckafka route.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ckafka_route" "example" {
	  instance_id = "ckafka-qz8w8rxz"
	  vip_type    = 3
	  vpc_id      = "vpc-k1azc3mv"
	  subnet_id   = "subnet-or7ddsbc"
	  access_type = 0
	}

```

# Import

ckafka route can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ckafka_route.example ckafka-qz8w8rxz#14256
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	ckafkaSDK "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ckafka_route", CNDescription{
		TerraformTypeCN: "CKafka 路由",
		DescriptionCN:   "提供 CKafka 路由资源，用于创建和管理 CKafka 实例的接入路由。",
		AttributesCN: map[string]string{
			"instance_id":  "CKafka 实例 ID",
			"vip_type":     "路由网络类型（3：VPC 路由）",
			"vpc_id":       "VPC 实例 ID",
			"subnet_id":    "子网 ID",
			"access_type":  "接入类型（0：PLAINTEXT；1：SASL_PLAINTEXT；2：SSL；3：SASL_SSL）",
			"auth_flag":    "是否需要权限管理",
			"caller_appid": "调用方 AppId",
			"vip_list":     "虚拟 IP 列表",
		},
	})
}

func resourceTencentCloudCkafkaRoute() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a ckafka route.",
		Create:      resourceTencentCloudCkafkaRouteCreate,
		Read:        resourceTencentCloudCkafkaRouteRead,
		Update:      resourceTencentCloudCkafkaRouteUpdate,
		Delete:      resourceTencentCloudCkafkaRouteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "CKafka instance id.",
			},
			"vip_type": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "Routing network type (3: vpc routing; 4: standard support routing; 7: professional support routing).",
			},
			"vpc_id": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Vpc id.",
			},
			"subnet_id": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Subnet id.",
			},
			"access_type": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeInt,
				Description: "Access type. Valid values:\n" +
					"- 0: PLAINTEXT (in clear text, supported by both the old version and the community version without user information)\n" +
					"- 1: SASL_PLAINTEXT (in clear text, but at the beginning of the data, authentication will be logged in through SASL, which is only supported by the community version)\n" +
					"- 2: SSL (SSL encrypted communication without user information, supported by both older and community versions)\n" +
					"- 3: SASL_SSL (SSL encrypted communication. When the data starts, authentication will be logged in through SASL. Only the community version supports it).",
			},
			"auth_flag": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Auth flag.",
			},
			"caller_appid": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Caller appid.",
			},
			"route_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Route ID, available after route creation is complete.",
			},
			"vip_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Virtual IP list.",
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
							Description: "Virtual port.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudCkafkaRouteCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ckafka_route.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	instanceId := d.Get("instance_id").(string)
	vipType := int64(d.Get("vip_type").(int))
	vpcId := d.Get("vpc_id").(string)
	subnetId := d.Get("subnet_id").(string)
	accessType := int64(d.Get("access_type").(int))

	request := ckafkaSDK.NewCreateRouteRequest()
	request.InstanceId = helper.String(instanceId)
	request.VipType = &vipType
	if vpcId != "" {
		request.VpcId = helper.String(vpcId)
	}
	if subnetId != "" {
		request.SubnetId = helper.String(subnetId)
	}
	request.AccessType = &accessType
	if v, ok := d.GetOkExists("auth_flag"); ok {
		authFlag := int64(v.(int))
		request.AuthFlag = &authFlag
	}
	if v, ok := d.GetOkExists("caller_appid"); ok {
		callerAppid := int64(v.(int))
		request.CallerAppid = &callerAppid
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCkafkaClient().CreateRoute(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create ckafka route failed, reason:%+v", logId, err)
		return err
	}

	// 等待 Processing 从 1 变为 0，并获取真实 RouteId
	// Processing=0 后 DescribeRoute 才会返回真实 RouteId
	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}
	var routeId int64
	appeared := false // 路由是否曾经出现过（Processing=1），出现后消失视为创建失败
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		route, e := service.DescribeCkafkaRouteByKey(ctx, instanceId, vipType, vpcId, subnetId, accessType)
		if e != nil {
			return retryError(e)
		}
		if route == nil {
			if appeared {
				// 路由曾出现但现在消失，说明后端创建失败
				return resource.NonRetryableError(fmt.Errorf("ckafka route disappeared after creation, backend may have failed, instance: %s", instanceId))
			}
			// 还未出现，继续等待
			return resource.RetryableError(fmt.Errorf("ckafka route not found yet, retrying"))
		}
		appeared = true
		if route.Processing != nil && *route.Processing == 1 {
			return resource.RetryableError(fmt.Errorf("ckafka route still processing, retrying"))
		}
		if route.RouteId != nil {
			routeId = *route.RouteId
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("ckafka route did not become ready: %s", err)
	}

	// ID 格式: instanceId#routeId
	d.SetId(fmt.Sprintf("%s%s%d", instanceId, FILED_SP, routeId))

	return resourceTencentCloudCkafkaRouteRead(d, meta)
}

func resourceTencentCloudCkafkaRouteRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ckafka_route.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	// ID 格式: instanceId#routeId
	items := strings.Split(d.Id(), FILED_SP)
	if len(items) < 2 {
		return fmt.Errorf("id is broken, expected format instanceId#routeId, got: %s", d.Id())
	}
	instanceId := items[0]
	routeId, err := strconv.ParseInt(items[1], 10, 64)
	if err != nil {
		return fmt.Errorf("id routeId parse error: %s", err)
	}

	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}
	route, err := service.DescribeCkafkaRouteById(ctx, instanceId, routeId)
	if err != nil {
		return err
	}

	_ = d.Set("instance_id", instanceId)

	if route == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CkafkaRoute` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if route.VipType != nil {
		_ = d.Set("vip_type", route.VipType)
	}
	if route.AccessType != nil {
		_ = d.Set("access_type", route.AccessType)
	}
	if route.VpcId != nil {
		_ = d.Set("vpc_id", route.VpcId)
	}
	if route.Subnet != nil {
		_ = d.Set("subnet_id", route.Subnet)
	}
	if route.RouteId != nil {
		_ = d.Set("route_id", route.RouteId)
	}

	vipList := make([]map[string]interface{}, 0, len(route.VipList))
	for _, vip := range route.VipList {
		vipList = append(vipList, map[string]interface{}{
			"vip":   vip.Vip,
			"vport": vip.Vport,
		})
	}
	_ = d.Set("vip_list", vipList)

	return nil
}

func resourceTencentCloudCkafkaRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ckafka_route.update")()
	defer inconsistentCheck(d, meta)()

	immutableArgs := []string{"instance_id", "vip_type", "vpc_id", "subnet_id", "access_type", "auth_flag", "caller_appid"}
	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}
	return resourceTencentCloudCkafkaRouteRead(d, meta)
}

func resourceTencentCloudCkafkaRouteDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ckafka_route.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	// ID 格式: instanceId#routeId
	items := strings.Split(d.Id(), FILED_SP)
	if len(items) < 2 {
		return fmt.Errorf("id is broken, expected format instanceId#routeId, got: %s", d.Id())
	}
	instanceId := items[0]
	routeId, err := strconv.ParseInt(items[1], 10, 64)
	if err != nil {
		return fmt.Errorf("id routeId parse error: %s", err)
	}

	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}
	if err := service.DeleteCkafkaRouteById(ctx, instanceId, routeId); err != nil {
		return err
	}
	return nil
}
