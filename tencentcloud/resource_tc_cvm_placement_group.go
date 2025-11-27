/*
Provide a resource to create a placement group.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cvm_placement_group" "foo" {
	  name = "test"
	  type = "HOST"
	}

```

# Import

Placement group can be imported using the id, e.g.

```
$ terraform import cloud_cvm_placement_group.foo ps-ilan8vjf
```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cvm_placement_group", CNDescription{
		TerraformTypeCN: "置放群组",
		DescriptionCN:   "提供置放群组资源，用于创建和管理云服务器的置放群组。",
		AttributesCN: map[string]string{
			"name": "分散置放群组名称，长度1-60个字符，支持中、英文",
			"type": "分散置放群组类型，取值范围：HOST：物理机, SW：交换机, RACK：机架",
			// "strategy": "置放群组类型：分散置放群组；分区置放群组；强亲和置放群组",
			"strategy": "置放群组类型：`SPREAD`：分散置放群组；`WEAK_SPREAD`：弱分散置放群组；`AFFINITY`：亲和置放群组；`WEAK_AFFINITY`：弱亲和置放群组",

			"cvm_quota_total": "分散置放群组内最大容纳云服务器数量",
			"current_num":     "分散置放群组内云服务器当前数量",
			"create_time":     "分散置放群组创建时间",
		},
	})
}
func resourceTencentCloudPlacementGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a placement group.",
		Create:      resourceTencentCloudPlacementGroupCreate,
		Read:        resourceTencentCloudPlacementGroupRead,
		Update:      resourceTencentCloudPlacementGroupUpdate,
		Delete:      resourceTencentCloudPlacementGroupDelete,
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
				ValidateFunc: validateAllowedStringValue(CVM_PLACEMENT_GROUP_TYPE),
				Description:  "Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.",
			},
			"strategy": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(CVM_PLACE_GROUP_STRATEGY),
				Description:  "Types are: `SPREAD`,`WEAK_SPREAD`,`AFFINITY` and `WEAK_AFFINITY`.",
			},

			// computed
			"cvm_quota_total": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Maximum number of hosts in the placement group.",
			},
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

func resourceTencentCloudPlacementGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_placement_group.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		placementName string
		placementType string
		strategy      string
	)

	cvmService := CvmService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	placementName = d.Get("name").(string)
	placementType = d.Get("type").(string)
	if v, ok := d.GetOk("strategy"); ok {
		strategy = v.(string)
	}
	var id string
	var errRet error
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		id, errRet = cvmService.CreatePlacementGroup(ctx, placementName, placementType, strategy)
		if errRet != nil {
			return retryError(errRet)
		}
		return nil
	})
	if err != nil {
		return err
	}
	d.SetId(id)

	return resourceTencentCloudPlacementGroupRead(d, meta)
}

func resourceTencentCloudPlacementGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_placement_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	placementId := d.Id()
	cvmService := CvmService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var placement *cvm.DisasterRecoverGroup
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		placement, errRet = cvmService.DescribePlacementGroupById(ctx, placementId)
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
	_ = d.Set("cvm_quota_total", placement.CvmQuotaTotal)
	_ = d.Set("current_num", placement.CurrentNum)
	_ = d.Set("create_time", placement.CreateTime)

	return nil
}

func resourceTencentCloudPlacementGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_placement_group.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	placementId := d.Id()
	cvmService := CvmService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	if d.HasChange("name") {
		placementName := d.Get("name").(string)
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			err := cvmService.ModifyPlacementGroup(ctx, placementId, placementName)
			if err != nil {
				return retryError(err)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return resourceTencentCloudPlacementGroupRead(d, meta)
}

func resourceTencentCloudPlacementGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_placement_group.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	placementId := d.Id()
	cvmService := CvmService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		err := cvmService.DeletePlacementGroup(ctx, placementId)
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
