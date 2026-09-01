package tencentcloud

import "testing"

func TestPlanNodePoolCapacityUpdate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		currentMin     int64
		currentMax     int64
		currentDesired int64
		newMin         int64
		newMax         int64
		newDesired     int64
		updateDesired  bool
		wantErr        bool
		wantTemp       bool
		tempMin        int64
		tempMax        int64
		wantDesired    bool
		wantFinal      bool
		finalMin       int64
		finalMax       int64
	}{
		{
			name:           "scale down min/max/desired together",
			currentMin:     13,
			currentMax:     13,
			currentDesired: 13,
			newMin:         3,
			newMax:         3,
			newDesired:     3,
			updateDesired:  true,
			wantTemp:       true,
			tempMin:        3,
			tempMax:        13,
			wantDesired:    true,
			wantFinal:      true,
			finalMin:       3,
			finalMax:       3,
		},
		{
			name:           "scale up min and desired",
			currentMin:     1,
			currentMax:     6,
			currentDesired: 2,
			newMin:         3,
			newMax:         6,
			newDesired:     5,
			updateDesired:  true,
			wantTemp:       false,
			tempMin:        1,
			tempMax:        6,
			wantDesired:    true,
			wantFinal:      true,
			finalMin:       3,
			finalMax:       6,
		},
		{
			name:           "desired only",
			currentMin:     1,
			currentMax:     6,
			currentDesired: 2,
			newMin:         1,
			newMax:         6,
			newDesired:     5,
			updateDesired:  true,
			wantTemp:       false,
			tempMin:        1,
			tempMax:        6,
			wantDesired:    true,
			wantFinal:      false,
			finalMin:       1,
			finalMax:       6,
		},
		{
			name:           "expand max only",
			currentMin:     1,
			currentMax:     6,
			currentDesired: 2,
			newMin:         1,
			newMax:         10,
			newDesired:     2,
			updateDesired:  false,
			wantTemp:       true,
			tempMin:        1,
			tempMax:        10,
			wantDesired:    false,
			wantFinal:      false,
			finalMin:       1,
			finalMax:       10,
		},
		{
			name:           "min above current desired without desired change",
			currentMin:     1,
			currentMax:     10,
			currentDesired: 2,
			newMin:         5,
			newMax:         10,
			newDesired:     2,
			updateDesired:  false,
			wantErr:        true,
		},
		{
			name:           "desired outside new range",
			currentMin:     1,
			currentMax:     6,
			currentDesired: 2,
			newMin:         1,
			newMax:         6,
			newDesired:     10,
			updateDesired:  true,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			plan, err := planNodePoolCapacityUpdate(tt.currentMin, tt.currentMax, tt.currentDesired, tt.newMin, tt.newMax, tt.newDesired, tt.updateDesired)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if plan.NeedTempRange != tt.wantTemp || plan.TempMin != tt.tempMin || plan.TempMax != tt.tempMax {
				t.Fatalf("temp range = (%v,%d,%d), want (%v,%d,%d)", plan.NeedTempRange, plan.TempMin, plan.TempMax, tt.wantTemp, tt.tempMin, tt.tempMax)
			}
			if plan.NeedDesired != tt.wantDesired || (tt.wantDesired && plan.Desired != tt.newDesired) {
				t.Fatalf("desired = (%v,%d), want (%v,%d)", plan.NeedDesired, plan.Desired, tt.wantDesired, tt.newDesired)
			}
			if plan.NeedFinalRange != tt.wantFinal || plan.FinalMin != tt.finalMin || plan.FinalMax != tt.finalMax {
				t.Fatalf("final range = (%v,%d,%d), want (%v,%d,%d)", plan.NeedFinalRange, plan.FinalMin, plan.FinalMax, tt.wantFinal, tt.finalMin, tt.finalMax)
			}
		})
	}
}
