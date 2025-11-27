/*
Provides a resource to create a cfs auto_snapshot_policy

# Example Usage

use day of week

```hcl

	resource "tencentcloudenterprise_cfs_auto_snapshot_policy" "auto_snapshot_policy" {
	  day_of_week = "1,2"
	  hour = "2,3"
	  policy_name = "policy_name"
	  alive_days = 7
	}

```

use day of month

```hcl

	resource "tencentcloudenterprise_cfs_auto_snapshot_policy" "auto_snapshot_policy" {
	  hour = "2,3"
	  policy_name = "policy_name"
	  alive_days = 7
	  day_of_month = "2,3,4"
	}

```

use interval days

```hcl

	resource "tencentcloudenterprise_cfs_auto_snapshot_policy" "auto_snapshot_policy" {
	  hour = "2,3"
	  policy_name = "policy_name"
	  alive_days = 7
	  interval_days = 1
	}

```

# Import

cfs auto_snapshot_policy can be imported using the id, e.g.

```
terraform import cloud_cfs_auto_snapshot_policy.auto_snapshot_policy auto_snapshot_policy_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cfs "terraform-provider-tencentcloudenterprise/sdk/cfs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfs_auto_snapshot_policy", CNDescription{
		TerraformTypeCN: "自动快照策略",
		DescriptionCN:   "提供CFS自动快照策略资源，用于创建和管理CFS自动快照策略。",
		AttributesCN: map[string]string{
			"day_of_week":             "每周几",
			"hour":                    "每天几点",
			"policy_name":             "策略名称",
			"alive_days":              "保留天数",
			"auto_snapshot_policy_id": "快照策略ID",
			//"day_of_month":  "每月几号",
			//"interval_days": "间隔天数",
		},
	})
}
func resourceTencentCloudCfsAutoSnapshotPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a cfs auto_snapshot_policy",
		Create:      resourceTencentCloudCfsAutoSnapshotPolicyCreate,
		Read:        resourceTencentCloudCfsAutoSnapshotPolicyRead,
		Update:      resourceTencentCloudCfsAutoSnapshotPolicyUpdate,
		Delete:      resourceTencentCloudCfsAutoSnapshotPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"hour": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The time point when to repeat the snapshot operation.",
			},

			"policy_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Policy name.",
			},

			"day_of_week": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The day of the week on which to repeat the snapshot operation.",
			},

			"alive_days": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Snapshot retention period.",
			},
			"auto_snapshot_policy_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The ID of the auto snapshot policy.",
			},

			//"day_of_month": {
			//	Optional:    true,
			//	Type:        schema.TypeString,
			//	Description: "The specific day (day 1 to day 31) of the month on which to create a snapshot.",
			//},

			//"interval_days": {
			//	Optional:    true,
			//	Type:        schema.TypeInt,
			//	Description: "The snapshot interval, in days.",
			//},
		},
	}
}

func resourceTencentCloudCfsAutoSnapshotPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfs_auto_snapshot_policy.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request              = cfs.NewCreateAutoSnapshotPolicyRequest()
		response             = cfs.NewCreateAutoSnapshotPolicyResponse()
		autoSnapshotPolicyId string
	)
	if v, ok := d.GetOk("day_of_week"); ok {
		request.DayOfWeek = helper.String(v.(string))
	}

	if v, ok := d.GetOk("hour"); ok {
		request.Hour = helper.String(v.(string))
	}

	if v, ok := d.GetOk("policy_name"); ok {
		request.PolicyName = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("alive_days"); ok {
		request.AliveDays = helper.IntUint64(v.(int))
	}

	/*
		if v, ok := d.GetOk("day_of_month"); ok {
			request.DayOfMonth = helper.String(v.(string))
		}

		if v, ok := d.GetOkExists("interval_days"); ok {
			request.IntervalDays = helper.IntUint64(v.(int))
		}

	*/

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfsClient().CreateAutoSnapshotPolicy(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cfs autoSnapshotPolicy failed, reason:%+v", logId, err)
		return err
	}

	autoSnapshotPolicyId = *response.Response.AutoSnapshotPolicyId
	d.SetId(autoSnapshotPolicyId)

	return resourceTencentCloudCfsAutoSnapshotPolicyRead(d, meta)
}

func resourceTencentCloudCfsAutoSnapshotPolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfs_auto_snapshot_policy.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CfsService{client: meta.(*TencentCloudClient).apiV3Conn}

	autoSnapshotPolicyId := d.Id()

	autoSnapshotPolicy, err := service.DescribeCfsAutoSnapshotPolicyById(ctx, autoSnapshotPolicyId)
	if err != nil {
		return err
	}

	if autoSnapshotPolicy == nil {
		d.SetId("")
		return fmt.Errorf("resource `cloud_cfs_auto_snapshot_policy` %s does not exist", d.Id())
	}

	if autoSnapshotPolicy.DayOfWeek != nil {
		_ = d.Set("day_of_week", autoSnapshotPolicy.DayOfWeek)
	}

	if autoSnapshotPolicy.Hour != nil {
		_ = d.Set("hour", autoSnapshotPolicy.Hour)
	}

	if autoSnapshotPolicy.PolicyName != nil {
		_ = d.Set("policy_name", autoSnapshotPolicy.PolicyName)
	}

	if autoSnapshotPolicy.AliveDays != nil {
		_ = d.Set("alive_days", autoSnapshotPolicy.AliveDays)
	}

	if autoSnapshotPolicy.AutoSnapshotPolicyId != nil {
		_ = d.Set("auto_snapshot_policy_id", autoSnapshotPolicy.AutoSnapshotPolicyId)
	}

	/*
		if autoSnapshotPolicy.DayOfMonth != nil {
			_ = d.Set("day_of_month", autoSnapshotPolicy.DayOfMonth)
		}

		if autoSnapshotPolicy.IntervalDays != nil {
			_ = d.Set("interval_days", autoSnapshotPolicy.IntervalDays)
		}

	*/

	return nil
}

func resourceTencentCloudCfsAutoSnapshotPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfs_auto_snapshot_policy.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cfs.NewUpdateAutoSnapshotPolicyRequest()

	autoSnapshotPolicyId := d.Id()

	request.AutoSnapshotPolicyId = &autoSnapshotPolicyId
	if d.HasChange("day_of_week") {
		if v, ok := d.GetOk("day_of_week"); ok {
			request.DayOfWeek = helper.String(v.(string))
		}
	}

	if d.HasChange("hour") {
		if v, ok := d.GetOk("hour"); ok {
			request.Hour = helper.String(v.(string))
		}
	}

	if d.HasChange("policy_name") {
		if v, ok := d.GetOk("policy_name"); ok {
			request.PolicyName = helper.String(v.(string))
		}
	}

	if d.HasChange("alive_days") {
		if v, ok := d.GetOkExists("alive_days"); ok {
			request.AliveDays = helper.IntUint64(v.(int))
		}
	}

	/*
		if d.HasChange("day_of_month") {
			if v, ok := d.GetOk("day_of_month"); ok {
				request.DayOfMonth = helper.String(v.(string))
			}
		}

		if d.HasChange("interval_days") {
			if v, ok := d.GetOkExists("interval_days"); ok {
				request.IntervalDays = helper.IntUint64(v.(int))
			}
		}

	*/

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfsClient().UpdateAutoSnapshotPolicy(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cfs autoSnapshotPolicy failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCfsAutoSnapshotPolicyRead(d, meta)
}

func resourceTencentCloudCfsAutoSnapshotPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfs_auto_snapshot_policy.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CfsService{client: meta.(*TencentCloudClient).apiV3Conn}
	autoSnapshotPolicyId := d.Id()

	if err := service.DeleteCfsAutoSnapshotPolicyById(ctx, autoSnapshotPolicyId); err != nil {
		return err
	}

	return nil
}
