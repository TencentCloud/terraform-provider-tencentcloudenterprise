/*
Provide a resource to create a placement group.

# Example Usage

```hcl

	resource "tencentcloudenterprise_bms_placement_group" "foo" {
	  name = "test"
	  type = "RACK"
	}

```

# Import

Placement group can be imported using the id, e.g.

```
$ terraform import cloud_bms_placement_group.foo ps-ilan8vjf
```
*/
package tencentcloud

import (
	"context"
	bms "terraform-provider-tencentcloudenterprise/sdk/bms/v20180813"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_bms_placement_group", CNDescription{
		TerraformTypeCN: "置放群组",
		DescriptionCN:   "提供置放群组资源，用于创建和管理裸金属服务器的置放群组。",
		AttributesCN: map[string]string{
			"name":        "置放群组名称",
			"type":        "置放群组类型",
			"current_num": "当前置放群组内云服务器数量",
			"create_time": "置放群组创建时间",
		},
	})
}
func resourceTencentCloudBmsPlacementGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a placement group.",
		Create:      resourceTencentCloudBmsPlacementGroupCreate,
		Read:        resourceTencentCloudBmsPlacementGroupRead,
		Update:      resourceTencentCloudBmsPlacementGroupUpdate,
		Delete:      resourceTencentCloudBmsPlacementGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "Name of the placement group, 1-60 characters in length.",
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue(BMS_PLACEMENT_GROUP_TYPE),
				Description:  "Type of the placement group. Valid values: `RACK_SAME_SW` and `RACK`.",
			},

			// computed
			"current_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of hosts in the placement group.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the placement group.",
			},
		},
	}
}

func resourceTencentCloudBmsPlacementGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_bms_placement_group.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		placementName string
		placementType string
	)

	bmsService := BmsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	placementName = d.Get("name").(string)
	placementType = d.Get("type").(string)

	var id string
	var errRet error
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		errRet = bmsService.CreatePlacementGroup(ctx, placementName, placementType)
		if errRet != nil {
			return retryError(errRet)
		}
		return nil
	})
	if err != nil {
		return err
	}
	res, err := bmsService.DescribePlacementGroupByFilter(ctx, "", placementName)
	if err != nil {
		return err
	}

	id = *res.GroupId
	d.SetId(id)

	return resourceTencentCloudBmsPlacementGroupRead(d, meta)
}

func resourceTencentCloudBmsPlacementGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_bms_placement_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	placementId := d.Id()
	bmsService := BmsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var placement *bms.DisasterRecoverGroup
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		placement, errRet = bmsService.DescribePlacementGroupByFilter(ctx, placementId, "")
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if placement == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("name", placement.Name)
	_ = d.Set("type", placement.Type)
	_ = d.Set("current_num", placement.CurrentNum)
	_ = d.Set("create_time", placement.CreateTime)

	return nil
}

func resourceTencentCloudBmsPlacementGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_bms_placement_group.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	placementId := d.Id()
	bmsService := BmsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	if d.HasChange("name") {
		placementName := d.Get("name").(string)
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			err := bmsService.UpdateDisasterRecoverGroup(ctx, placementId, placementName)
			if err != nil {
				return retryError(err)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return resourceTencentCloudBmsPlacementGroupRead(d, meta)
}

func resourceTencentCloudBmsPlacementGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_bms_placement_group.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	placementId := d.Id()
	bmsService := BmsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		err := bmsService.DeletePlacementGroup(ctx, placementId)
		if err != nil {
			return retryError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
