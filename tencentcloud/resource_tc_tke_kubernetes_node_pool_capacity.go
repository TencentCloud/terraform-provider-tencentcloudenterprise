package tencentcloud

import (
	"fmt"

	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"

	"github.com/pkg/errors"
)

type nodePoolCapacityPlan struct {
	NeedTempRange  bool
	TempMin        int64
	TempMax        int64
	NeedDesired    bool
	Desired        int64
	NeedFinalRange bool
	FinalMin       int64
	FinalMax       int64
}

type nodePoolWaitAction int

const (
	nodePoolWaitReady nodePoolWaitAction = iota
	nodePoolWaitRetry
	nodePoolWaitFail
)

func minInt64(values ...int64) int64 {
	minValue := values[0]
	for _, value := range values[1:] {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func maxInt64(values ...int64) int64 {
	maxValue := values[0]
	for _, value := range values[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// planNodePoolCapacityUpdate keeps AS `MinSize <= DesiredCapacity <= MaxSize` while
// min/max/desired change together. Scale-down 13/13/13 -> 3/3/3 becomes
// [3,13] -> desired 3 -> [3,3].
func planNodePoolCapacityUpdate(currentMin, currentMax, currentDesired, newMin, newMax, newDesired int64, updateDesired bool) (*nodePoolCapacityPlan, error) {
	if newMin > newMax {
		return nil, fmt.Errorf("constraints `min_size <= desired_capacity <= max_size` must be established,")
	}

	effectiveDesired := currentDesired
	if updateDesired {
		effectiveDesired = newDesired
		if newDesired < newMin || newDesired > newMax {
			return nil, fmt.Errorf("constraints `min_size <= desired_capacity <= max_size` must be established,")
		}
	} else if currentDesired < newMin || currentDesired > newMax {
		return nil, fmt.Errorf("cannot set min_size=%d max_size=%d while desired_capacity is %d; keep min_size <= desired_capacity <= max_size", newMin, newMax, currentDesired)
	}

	tempMin := minInt64(currentMin, newMin, currentDesired, effectiveDesired)
	tempMax := maxInt64(currentMax, newMax, currentDesired, effectiveDesired)

	return &nodePoolCapacityPlan{
		NeedTempRange:  tempMin != currentMin || tempMax != currentMax,
		TempMin:        tempMin,
		TempMax:        tempMax,
		NeedDesired:    updateDesired && newDesired != currentDesired,
		Desired:        newDesired,
		NeedFinalRange: newMin != tempMin || newMax != tempMax,
		FinalMin:       newMin,
		FinalMax:       newMax,
	}, nil
}

func nodePoolWaitDecision(has bool, lifeState *string) (nodePoolWaitAction, string) {
	if !has {
		return nodePoolWaitFail, "not found"
	}
	state := "unknown"
	if lifeState != nil {
		state = *lifeState
	}
	switch state {
	case "normal":
		return nodePoolWaitReady, state
	case "abnormal":
		return nodePoolWaitFail, state
	default:
		return nodePoolWaitRetry, state
	}
}

func isBareOperationDenied(err error) bool {
	if err == nil {
		return false
	}
	sdkErr, ok := err.(*sdkErrors.CloudSDKError)
	if !ok {
		sdkErr, ok = errors.Cause(err).(*sdkErrors.CloudSDKError)
	}
	return ok && sdkErr.Code == "OperationDenied"
}
