/*
Provides a resource to create a VPC routing table.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc" "foo" {
	  name       = "ci-temp-test"
	  cidr_block = "10.0.0.0/16"
	}

	resource "tencentcloudenterprise_vpc_route_table" "foo" {
	  vpc_id = cloud_vpc.foo.id
	  name   = "ci-temp-test-rt"
	}

```

# Import

Vpc routetable instance can be imported, e.g.

```
$ terraform import cloud_route_table.test route_table_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_route_table", CNDescription{
		TerraformTypeCN: "路由表",
		DescriptionCN:   "提供路由表资源，用于创建和管理VPC路由表。",
		AttributesCN: map[string]string{
			"vpc_id":          "VPC实例ID",
			"name":            "路由表名称",
			"subnet_ids":      "子网实例ID",
			"route_entry_ids": "路由策略实例ID",
			"is_default":      "是否默认路由表",
			"create_time":     "创建时间",
			"tags":            "标签",
		},
	})
}
func resourceTencentCloudVpcRouteTable() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a VPC routing table.",
		Create:      resourceTencentCloudVpcRouteTableCreate,
		Read:        resourceTencentCloudVpcRouteTableRead,
		Update:      resourceTencentCloudVpcRouteTableUpdate,
		Delete:      resourceTencentCloudVpcRouteTableDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of VPC to which the route table should be associated.",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "The name of routing table.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags of routing table.",
			},

			// Computed values
			"subnet_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "ID list of the subnets associated with this route table.",
			},
			"route_entry_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "ID list of the routing entries.",
			},
			"is_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether it is the default routing table.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the routing table.",
			},
		},
	}
}

func resourceTencentCloudVpcRouteTableCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_route_table.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	vpcService := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		vpcId string
		name  string
		tags  map[string]string
	)
	if temp, ok := d.GetOk("vpc_id"); ok {
		vpcId = temp.(string)
		if len(vpcId) < 1 {
			return fmt.Errorf("vpc_id should be not empty string")
		}
	}
	if temp, ok := d.GetOk("name"); ok {
		name = temp.(string)
	}

	if temp := helper.GetTags(d, "tags"); len(temp) > 0 {
		tags = temp
	}

	routeTableId, err := vpcService.CreateRouteTable(ctx, name, vpcId, tags)
	if err != nil {
		return err
	}
	d.SetId(routeTableId)

	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}

		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::vpc:%s:uin/:rtb/%s", region, routeTableId)

		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudVpcRouteTableRead(d, meta)
}

func resourceTencentCloudVpcRouteTableRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_route_table.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()

	var (
		info VpcRouteTableBasicInfo
		has  int
		e    error
	)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, has, e = service.DescribeRouteTable(ctx, id)
		if e != nil {
			return retryError(e)
		}
		// deleted
		if has == 0 {
			d.SetId("")
			return nil
		}
		if has != 1 {
			errRet := fmt.Errorf("one route_table_id read get %d route_table info", has)
			log.Printf("[CRITAL]%s %s", logId, errRet.Error())
			return resource.NonRetryableError(errRet)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if has == 0 {
		return nil
	}
	routeEntryIds := make([]string, 0, len(info.entryInfos))
	for _, v := range info.entryInfos {
		tfRouteEntryId := fmt.Sprintf("%d.%s", v.routeEntryId, id)
		routeEntryIds = append(routeEntryIds, tfRouteEntryId)
	}

	tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}

	region := meta.(*TencentCloudClient).apiV3Conn.Region
	tags, err := tagService.DescribeResourceTags(ctx, "vpc", "rtb", region, id)
	if err != nil {
		return err
	}

	_ = d.Set("vpc_id", info.vpcId)
	_ = d.Set("name", info.name)
	_ = d.Set("subnet_ids", info.subnetIds)
	_ = d.Set("route_entry_ids", routeEntryIds)
	_ = d.Set("is_default", info.isDefault)
	_ = d.Set("create_time", info.createTime)
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudVpcRouteTableUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_route_table.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	d.Partial(true)

	if d.HasChange("name") {
		name := d.Get("name").(string)
		err := service.ModifyRouteTableAttribute(ctx, id, name)
		if err != nil {
			return err
		}

	}

	if d.HasChange("tags") {
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))

		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}

		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::vpc:%s:uin/:rtb/%s", region, id)

		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}

	}

	d.Partial(false)

	return resourceTencentCloudVpcRouteTableRead(d, meta)
}

func resourceTencentCloudVpcRouteTableDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_route_table.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if err := service.DeleteRouteTable(ctx, d.Id()); err != nil {
			if sdkErr, ok := err.(*errors.CloudSDKError); ok {
				if sdkErr.Code == VPCNotFound {
					return nil
				}
			}
			return resource.RetryableError(err)
		}
		return nil
	})

	return err
}
