package tencentcloud

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateNodePoolCapacityWait(wait, autoScale, hasTolerance bool) error {
	if wait && autoScale {
		return fmt.Errorf("wait_node_ready requires enable_auto_scale=false")
	}
	if hasTolerance && !wait {
		return fmt.Errorf("scale_tolerance requires wait_node_ready=true")
	}
	return nil
}

// Do not mistake an optional field's zero value in state for an explicit setting.
func nodePoolToleranceConfigured(config cty.Value, fallback bool) bool {
	if !config.IsKnown() || config.IsNull() || !config.Type().IsObjectType() {
		return fallback
	}
	return config.Type().HasAttribute("scale_tolerance") && !config.GetAttr("scale_tolerance").IsNull()
}

func customizeNodePoolCapacityWaitDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	if !d.NewValueKnown("wait_node_ready") {
		return nil
	}
	wait := d.Get("wait_node_ready").(bool)
	_, hasTolerance := d.GetOkExists("scale_tolerance")
	hasTolerance = nodePoolToleranceConfigured(d.GetRawConfig(), hasTolerance)
	if !wait {
		return validateNodePoolCapacityWait(false, false, hasTolerance)
	}
	if !d.NewValueKnown("enable_auto_scale") {
		return nil
	}
	return validateNodePoolCapacityWait(true, d.Get("enable_auto_scale").(bool), hasTolerance)
}

func validateNodePoolCapacityWaitData(d *schema.ResourceData) error {
	_, hasTolerance := d.GetOkExists("scale_tolerance")
	hasTolerance = nodePoolToleranceConfigured(d.GetRawConfig(), hasTolerance)
	return validateNodePoolCapacityWait(d.Get("wait_node_ready").(bool), d.Get("enable_auto_scale").(bool), hasTolerance)
}

// A Computed desired_capacity in state is not necessarily an explicit HCL target.
func configuredNodePoolDesired(d *schema.ResourceData) *int64 {
	config := d.GetRawConfig()
	if !config.IsKnown() || config.IsNull() || !config.Type().IsObjectType() || !config.Type().HasAttribute("desired_capacity") {
		return nil
	}
	value := config.GetAttr("desired_capacity")
	if !value.IsKnown() || value.IsNull() {
		return nil
	}
	desired := int64(d.Get("desired_capacity").(int))
	return &desired
}

type nodePoolCapacityWaitOptions struct {
	desired   *int64
	min, max  int64
	tolerance int
	timeout   time.Duration
}

func waitForNodePoolCapacity(ctx context.Context, service TkeService, clusterId, nodePoolId string, d *schema.ResourceData, deadline time.Time) error {
	if !d.Get("wait_node_ready").(bool) {
		return nil
	}
	tolerance := 100
	if v, ok := d.GetOkExists("scale_tolerance"); nodePoolToleranceConfigured(d.GetRawConfig(), ok) {
		tolerance = v.(int)
	}
	return waitNodePoolCapacity(ctx, service, clusterId, nodePoolId, nodePoolCapacityWaitOptions{
		desired: configuredNodePoolDesired(d), min: int64(d.Get("min_size").(int)), max: int64(d.Get("max_size").(int)),
		tolerance: tolerance, timeout: time.Until(deadline),
	})
}

func nodePoolCapacityConverged(target, total, normal int64, tolerance int) bool {
	if target < 0 || total < 0 || normal < 0 || normal > total || total > target || tolerance < 0 || tolerance > 100 {
		return false
	}
	if target == 0 {
		return total == 0 && normal == 0
	}
	// ceil(target*tolerance/100), without multiplying the full target first.
	required := target/100*int64(tolerance) + (target%100*int64(tolerance)+99)/100
	return normal >= required
}

func waitNodePoolCapacity(ctx context.Context, service TkeService, clusterId, nodePoolId string, options nodePoolCapacityWaitOptions) error {
	var observationMu sync.Mutex
	targetText, desiredText, totalText, normalText := "unknown", "unknown", "unknown", "unknown"
	var target int64
	targetKnown, targetObserved := false, false
	if options.desired != nil {
		target, targetKnown = *options.desired, true
		targetText = fmt.Sprintf("%d", target)
	}
	wrapError := func(err error) error {
		observationMu.Lock()
		defer observationMu.Unlock()
		return fmt.Errorf("waiting for node pool %s capacity failed (target=%s, observed_desired=%s, total=%s, normal=%s, scale_tolerance=%d%%, remaining_budget=%s): %w",
			nodePoolId, targetText, desiredText, totalText, normalText, options.tolerance, options.timeout, err)
	}
	if options.timeout <= 0 {
		return wrapError(fmt.Errorf("create/update timeout budget exhausted before capacity verification"))
	}
	waitCtx, cancel := context.WithTimeout(ctx, options.timeout)
	defer cancel()
	started := time.Now()
	err := resource.RetryContext(waitCtx, options.timeout, func() *resource.RetryError {
		pool, has, err := service.DescribeNodePool(waitCtx, clusterId, nodePoolId)
		if err != nil {
			return retryError(err, InternalError)
		}
		observationMu.Lock()
		defer observationMu.Unlock()
		desiredText, totalText, normalText = "unknown", "unknown", "unknown"
		if !has || pool == nil {
			return resource.RetryableError(fmt.Errorf("node pool details are not yet visible"))
		}
		if pool.DesiredNodesNum != nil {
			desiredText = fmt.Sprintf("%d", *pool.DesiredNodesNum)
		}
		if pool.NodeCountSummary != nil && pool.NodeCountSummary.AutoscalingAdded != nil {
			counts := pool.NodeCountSummary.AutoscalingAdded
			if counts.Total != nil {
				totalText = fmt.Sprintf("%d", *counts.Total)
			}
			if counts.Normal != nil {
				normalText = fmt.Sprintf("%d", *counts.Normal)
			}
		}
		if pool.LifeState != nil && *pool.LifeState == "abnormal" {
			return resource.NonRetryableError(fmt.Errorf("node pool is abnormal"))
		}
		if pool.DesiredNodesNum == nil {
			return resource.RetryableError(fmt.Errorf("desired node count is not yet available"))
		}
		observed := *pool.DesiredNodesNum
		if targetKnown && targetObserved && observed != target {
			return resource.NonRetryableError(fmt.Errorf("cloud desired node count changed from %d to %d during capacity verification", target, observed))
		}
		if pool.LifeState == nil || *pool.LifeState != "normal" {
			return resource.RetryableError(fmt.Errorf("node pool is not normal"))
		}
		if observed < options.min || observed > options.max {
			return resource.RetryableError(fmt.Errorf("cloud desired node count %d is outside the final configured range [%d,%d]", observed, options.min, options.max))
		}
		if !targetKnown {
			target, targetKnown = observed, true
			targetText = fmt.Sprintf("%d", target)
		}
		if observed != target {
			return resource.RetryableError(fmt.Errorf("waiting for cloud desired node count %d, currently %d", target, observed))
		}
		targetObserved = true
		if pool.NodeCountSummary == nil || pool.NodeCountSummary.AutoscalingAdded == nil || pool.NodeCountSummary.AutoscalingAdded.Total == nil || pool.NodeCountSummary.AutoscalingAdded.Normal == nil {
			return resource.RetryableError(fmt.Errorf("AS-added node statistics are not yet available"))
		}
		counts := pool.NodeCountSummary.AutoscalingAdded
		if !nodePoolCapacityConverged(target, *counts.Total, *counts.Normal, options.tolerance) {
			return resource.RetryableError(fmt.Errorf("AS-added node counts have not converged"))
		}
		return nil
	})
	if err != nil {
		if waitCtx.Err() != nil {
			return wrapError(fmt.Errorf("capacity verification %w; last observation: %v", waitCtx.Err(), err))
		}
		if time.Since(started) >= options.timeout {
			return wrapError(fmt.Errorf("capacity verification timed out; last observation: %w", err))
		}
		return wrapError(err)
	}
	return nil
}
