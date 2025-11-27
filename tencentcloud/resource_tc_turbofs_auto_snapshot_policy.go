/*
Provides a resource to create a turbofs auto_snapshot_policy

# Example Usage

use day of week

```hcl

	resource "tencentcloudenterprise_turbofs_auto_snapshot_policy" "auto_snapshot_policy" {
	  day_of_week = "1,2"
	  hour = "2,3"
	  policy_name = "policy_name"
	  alive_days = 7
	}

```

use day of month

```hcl

	resource "tencentcloudenterprise_turbofs_auto_snapshot_policy" "auto_snapshot_policy" {
	  hour = "2,3"
	  policy_name = "policy_name"
	  alive_days = 7
	  day_of_month = "2,3,4"
	}

```
Import

turbofs auto_snapshot_policy can be imported using the id, e.g.

```
terraform import cloud_turbofs_auto_snapshot_policy.auto_snapshot_policy auto_snapshot_policy_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_turbofs_auto_snapshot_policy", CNDescription{
		TerraformTypeCN: "自动快照策略",
		DescriptionCN:   "提供TurboFS自动快照策略资源，用于创建和管理TurboFS自动快照策略。",
		AttributesCN: map[string]string{
			"day_of_week":             "自动快照策略产生快照时间，DayOfWeek/DayOfMonth/IntervalDays三选一，DayOfWeek是一周中周几产生快照",
			"day_of_month":            "自动快照策略产生快照时间，DayOfWeek/DayOfMonth/IntervalDays三选一，DayOfMonth是一月中几号产生快照",
			"hour":                    "快照重复时间点",
			"policy_name":             "策略名称",
			"alive_days":              "快照保留时长",
			"auto_snapshot_policy_id": "快照策略ID",
			"interval_days":           "自动快照策略产生快照时间，DayOfWeek/DayOfMonth/IntervalDays三选一，IntervalDays是一年中几号产生快照",
		},
	})
}

func resourceTencentCloudTurbofsAutoSnapshotPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a TurboFS auto snapshot policy.",
		Create:      resourceTencentCloudTurbofsAutoSnapshotPolicyCreate,
		Read:        resourceTencentCloudTurbofsAutoSnapshotPolicyRead,
		Update:      resourceTencentCloudTurbofsAutoSnapshotPolicyUpdate,
		Delete:      resourceTencentCloudTurbofsAutoSnapshotPolicyDelete,
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
				Required:    true,
				Type:        schema.TypeString,
				Description: "Policy name.",
			},
			"interval_days": {
				Optional:     true,
				Type:         schema.TypeInt,
				Description:  "The interval days which to repeat the snapshot operation, conflict with day_of_week and day_of_month.",
				ExactlyOneOf: TURBOFS_SNAPSHOT_POLICY_CYCLE,
			},
			"day_of_week": {
				Optional:     true,
				Type:         schema.TypeString,
				Description:  "The day of the week on which to repeat the snapshot operation, conflict with interval_days and day_of_month.",
				ExactlyOneOf: TURBOFS_SNAPSHOT_POLICY_CYCLE,
			},
			"day_of_month": {
				Optional:     true,
				Type:         schema.TypeString,
				Description:  "The day of the month on which to repeat the snapshot operation, conflict with interval_days and day_of_week.",
				ExactlyOneOf: TURBOFS_SNAPSHOT_POLICY_CYCLE,
			},
			"alive_days": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Snapshot retention period. default value is 0, which means the snapshot will be retained permanently.",
			},
			"auto_snapshot_policy_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The ID of the auto snapshot policy.",
			},
		},
	}
}

func resourceTencentCloudTurbofsAutoSnapshotPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_auto_snapshot_policy.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request              = turbofs.NewCreateAutoSnapshotPolicyRequest()
		response             = turbofs.NewCreateAutoSnapshotPolicyResponse()
		autoSnapshotPolicyId string
	)

	if v, ok := d.GetOk("interval_days"); ok {
		request.IntervalDays = helper.IntUint64(v.(int))
	} else if v, ok := d.GetOk("day_of_week"); ok {
		request.DayOfWeek = helper.String(v.(string))
	} else if v, ok = d.GetOk("day_of_month"); ok {
		request.DayOfMonth = helper.String(v.(string))
	} else {
		return fmt.Errorf("'interval_days', day_of_week' or 'day_of_month' must be specified")
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

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().CreateAutoSnapshotPolicy(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create turbofs autoSnapshotPolicy failed, reason:%+v", logId, err)
		return err
	}

	autoSnapshotPolicyId = *response.Response.AutoSnapshotPolicyId
	d.SetId(autoSnapshotPolicyId)

	return resourceTencentCloudTurbofsAutoSnapshotPolicyRead(d, meta)
}

func resourceTencentCloudTurbofsAutoSnapshotPolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_auto_snapshot_policy.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TurbofsService{client: meta.(*TencentCloudClient).apiV3Conn}
	autoSnapshotPolicyId := d.Id()

	autoSnapshotPolicy, err := service.DescribeCfsAutoSnapshotPolicyById(ctx, autoSnapshotPolicyId)
	if err != nil {
		return err
	}
	if autoSnapshotPolicy == nil {
		d.SetId("")
		return fmt.Errorf("resource `cloud_turbofs_auto_snapshot_policy` %s does not exist", d.Id())
	}

	if autoSnapshotPolicy.IntervalDays != nil {
		_ = d.Set("interval_days", autoSnapshotPolicy.IntervalDays)
	}

	if autoSnapshotPolicy.DayOfWeek != nil {
		_ = d.Set("day_of_week", autoSnapshotPolicy.DayOfWeek)
	}

	if autoSnapshotPolicy.DayOfMonth != nil {
		_ = d.Set("day_of_month", autoSnapshotPolicy.DayOfMonth)
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

	return nil
}

func resourceTencentCloudTurbofsAutoSnapshotPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_auto_snapshot_policy.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := turbofs.NewUpdateAutoSnapshotPolicyRequest()
	autoSnapshotPolicyId := d.Id()
	request.AutoSnapshotPolicyId = &autoSnapshotPolicyId

	if d.HasChange("interval_days") {
		if v, ok := d.GetOk("interval_days"); ok {
			request.IntervalDays = helper.IntUint64(v.(int))
		}
	}

	if d.HasChange("day_of_week") {
		if v, ok := d.GetOk("day_of_week"); ok {
			request.DayOfWeek = helper.String(v.(string))
		}
	}

	if d.HasChange("day_of_month") {
		if v, ok := d.GetOk("day_of_month"); ok {
			request.DayOfMonth = helper.String(v.(string))
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

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().UpdateAutoSnapshotPolicy(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId,
				request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update turbofs autoSnapshotPolicy failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTurbofsAutoSnapshotPolicyRead(d, meta)
}

func resourceTencentCloudTurbofsAutoSnapshotPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_auto_snapshot_policy.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TurbofsService{client: meta.(*TencentCloudClient).apiV3Conn}
	autoSnapshotPolicyId := d.Id()

	if err := service.DeleteCfsAutoSnapshotPolicyById(ctx, autoSnapshotPolicyId); err != nil {
		return err
	}

	return nil
}
