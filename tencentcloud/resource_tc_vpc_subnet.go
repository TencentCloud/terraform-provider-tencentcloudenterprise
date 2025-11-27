/*
Provide a resource to create a VPC subnet.

# Example Usage

```hcl

	variable "availability_zone" {
	  default = "ap-guangzhou-3"
	}

	resource "tencentcloudenterprise_vpc" "foo" {
	  name       = "guagua-ci-temp-test"
	  cidr_block = "10.0.0.0/16"
	}

	resource "tencentcloudenterprise_vpc_subnet" "subnet" {
	  availability_zone = var.availability_zone
	  name              = "guagua-ci-temp-test"
	  vpc_id            = cloud_vpc.foo.id
	  cidr_block        = "10.0.20.0/28"
	  is_multicast      = false
	}

```

# Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import cloud_vpc_subnet.test subnet_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_subnet", CNDescription{
		TerraformTypeCN: "私有网络子网",
		DescriptionCN:   "提供VPC子网资源，用于创建和管理VPC子网。",
		AttributesCN: map[string]string{
			"vpc_id":             "VPC实例ID",
			"availability_zone":  "可用区",
			"name":               "子网名称",
			"cidr_block":         "子网网段",
			"is_multicast":       "是否开启广播",
			"route_table_id":     "路由表实例ID",
			"tags":               "标签",
			"subnet_type":        "子网类型",
			"subnet_id":          "子网实例ID",
			"create_time":        "创建时间",
			"is_default":         "是否默认子网",
			"available_ip_count": "可用IP数",
		},
	})
}
func resourceTencentCloudVpcSubnet() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a VPC subnet.",
		Create:      resourceTencentCloudVpcSubnetCreate,
		Read:        resourceTencentCloudVpcSubnetRead,
		Update:      resourceTencentCloudVpcSubnetUpdate,
		Delete:      resourceTencentCloudVpcSubnetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the VPC to be associated.",
			},
			"availability_zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The availability zone within which the subnet should be created.",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "The name of subnet to be created.",
			},
			"cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateCIDRNetworkAddress,
				Description:  "A network address block of the subnet.",
			},
			"is_multicast": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicates whether multicast is enabled. The default value is 'true'.",
			},
			"route_table_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "ID of a routing table to which the subnet should be associated.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tags of the subnet.",
			},
			"subnet_type": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateAllowedIntValue(Subnet_Type_Set),
				Default:      Subnet_Type_Normal,
				Description:  "subnet type. 0: Normal VPC; 9: BMS. Default is 0.",
			},

			// Computed values
			"is_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether it is the default VPC for this region.",
			},
			"available_ip_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of available IPs.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of subnet resource.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the subnet.",
			},
		},
	}
}

func resourceTencentCloudVpcSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_subnet.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	vpcService := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		vpcId            string
		availabilityZone string
		name             string
		cidrBlock        string
		//isMulticast      bool
		routeTableId string
		subnetType   uint64
		tags         map[string]string
	)
	if temp, ok := d.GetOk("vpc_id"); ok {
		vpcId = temp.(string)
		if len(vpcId) < 1 {
			return fmt.Errorf("vpc_id should be not empty string")
		}
	}
	if temp, ok := d.GetOk("availability_zone"); ok {
		availabilityZone = temp.(string)
		if len(availabilityZone) < 1 {
			return fmt.Errorf("availability_zone should be not empty string")
		}
	}
	if temp, ok := d.GetOk("name"); ok {
		name = temp.(string)
	}
	if temp, ok := d.GetOk("cidr_block"); ok {
		cidrBlock = temp.(string)
	}
	if temp, ok := d.GetOk("subnet_type"); ok {
		subnetType = uint64(temp.(int))
	}

	//isMulticast = d.Get("is_multicast").(bool)

	if temp, ok := d.GetOk("route_table_id"); ok {
		routeTableId = temp.(string)
		if len(routeTableId) < 1 {
			return fmt.Errorf("route_table_id should be not empty string")
		}
	}

	if routeTableId != "" {
		_, has, err := vpcService.IsRouteTableInVpc(ctx, routeTableId, vpcId)
		if err != nil {
			return err
		}
		if has != 1 {
			err = fmt.Errorf("error,route_table [%s]  not found in vpc [%s]", routeTableId, vpcId)
			log.Printf("[CRITAL]%s %s", logId, err.Error())
			return err
		}
	}

	if temp := helper.GetTags(d, "tags"); len(temp) > 0 {
		tags = temp
	}

	subnetId, err := vpcService.CreateSubnet(ctx, vpcId, name, cidrBlock, availabilityZone, tags, subnetType)
	if err != nil {
		return err
	}
	d.SetId(subnetId)

	//err = vpcService.ModifySubnetAttribute(ctx, subnetId, name, isMulticast)
	//if err != nil {
	//	return err
	//}

	if routeTableId != "" {
		err = vpcService.ReplaceRouteTableAssociation(ctx, subnetId, routeTableId)
		if err != nil {
			return err
		}
	}

	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}

		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::vpc:%s:uin/:subnet/%s", region, subnetId)

		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}
	return resourceTencentCloudVpcSubnetRead(d, meta)
}

func resourceTencentCloudVpcSubnetRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_subnet.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	subnetId := d.Id()

	vpcService := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	region := meta.(*TencentCloudClient).apiV3Conn.Region
	var (
		info VpcSubnetBasicInfo
		has  int
		e    error
	)
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, has, e = vpcService.DescribeSubnet(ctx, subnetId, nil, "", "")
		if e != nil {
			return retryError(e)
		}
		log.Printf("[INFO] info %v, has %d", info, has)

		// deleted
		if has == 0 {
			d.SetId("")
			return nil
		}

		if has != 1 {
			errRet := fmt.Errorf("one subnet_id read get %d subnet info", has)
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
	tags, err := tagService.DescribeResourceTags(ctx, "vpc", "subnet", region, subnetId)
	if err != nil {
		return err
	}

	_ = d.Set("vpc_id", info.vpcId)
	_ = d.Set("availability_zone", info.zone)
	_ = d.Set("name", info.name)
	_ = d.Set("cidr_block", info.cidr)
	_ = d.Set("is_multicast", info.isMulticast)
	_ = d.Set("route_table_id", info.routeTableId)
	_ = d.Set("is_default", info.isDefault)
	_ = d.Set("available_ip_count", info.availableIpCount)
	_ = d.Set("create_time", info.createTime)
	_ = d.Set("tags", tags)
	_ = d.Set("subnet_id", info.subnetId)

	return nil
}

func resourceTencentCloudVpcSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_subnet.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		name        string
		isMulticast bool
	)
	old, now := d.GetChange("name")
	if d.HasChange("name") {
		name = now.(string)
	} else {
		name = old.(string)
	}

	old, now = d.GetChange("is_multicast")
	if d.HasChange("is_multicast") {
		isMulticast = now.(bool)
	} else {
		isMulticast = old.(bool)
	}

	d.Partial(true)

	if err := service.ModifySubnetAttribute(ctx, id, name, isMulticast); err != nil {
		return err
	}

	if d.HasChange("route_table_id") {
		routeTableId := d.Get("route_table_id").(string)
		if len(routeTableId) < 1 {
			return fmt.Errorf("route_table_id should be not empty string")
		}

		_, has, err := service.IsRouteTableInVpc(ctx, routeTableId, d.Get("vpc_id").(string))
		if err != nil {
			return err
		}
		if has != 1 {
			err = fmt.Errorf("error,route_table [%s]  not found in vpc [%s]", routeTableId, d.Get("vpc_id").(string))
			log.Printf("[CRITAL]%s %s", logId, err.Error())
			return err
		}

		if err := service.ReplaceRouteTableAssociation(ctx, id, routeTableId); err != nil {
			return err
		}
	}

	if d.HasChange("tags") {
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))

		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}

		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::vpc:%s:uin/:subnet/%s", region, id)

		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}

	}

	d.Partial(false)

	return resourceTencentCloudVpcSubnetRead(d, meta)
}

func resourceTencentCloudVpcSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_subnet.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if err := service.DeleteSubnet(ctx, d.Id()); err != nil {
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
