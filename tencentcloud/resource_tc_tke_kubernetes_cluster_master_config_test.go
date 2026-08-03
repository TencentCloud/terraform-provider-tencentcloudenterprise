package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestTkeMasterConfigPreStartUserScriptSchema(t *testing.T) {
	masterSchema := TkeMasterCvmCreateInfo()
	field, ok := masterSchema["pre_start_user_script"]
	if !ok {
		t.Fatal("master_config must expose pre_start_user_script")
	}
	if field.Type != schema.TypeString || !field.Optional || field.ForceNew {
		t.Fatalf("unexpected pre_start_user_script schema: %#v", field)
	}
	if _, ok := TkeCvmCreateInfo()["pre_start_user_script"]; ok {
		t.Fatal("pre_start_user_script must not be exposed by worker_config")
	}
}

func TestExpandMasterConfigInstanceAdvancedSettings(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]interface{}
		hasOverride bool
		wantScript  string
		wantDesired int64
		wantErr     bool
	}{
		{name: "nil block", input: nil},
		{name: "empty values", input: map[string]interface{}{"pre_start_user_script": "", "desired_pod_num": 0}},
		{name: "script", input: map[string]interface{}{"pre_start_user_script": "c2NyaXB0"}, hasOverride: true, wantScript: "c2NyaXB0"},
		{name: "desired pod number", input: map[string]interface{}{"desired_pod_num": 32}, hasOverride: true, wantDesired: 32},
		{name: "bad script type", input: map[string]interface{}{"pre_start_user_script": 1}, wantErr: true},
		{name: "bad desired pod number type", input: map[string]interface{}{"desired_pod_num": "32"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, hasOverride, err := expandMasterConfigInstanceAdvancedSettings(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr = %t", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if hasOverride != tt.hasOverride {
				t.Fatalf("hasOverride = %t, want %t", hasOverride, tt.hasOverride)
			}
			if got.PreStartUserScript != nil && *got.PreStartUserScript != tt.wantScript {
				t.Fatalf("pre_start_user_script = %q, want %q", *got.PreStartUserScript, tt.wantScript)
			}
			if got.DesiredPodNumber != nil && *got.DesiredPodNumber != tt.wantDesired {
				t.Fatalf("desired_pod_num = %d, want %d", *got.DesiredPodNumber, tt.wantDesired)
			}
		})
	}
}

func TestNormalizeMasterConfigBlockPreStartUserScript(t *testing.T) {
	normalizeMasterConfigBlock(nil)
	block := map[string]interface{}{}
	normalizeMasterConfigBlock(block)
	if got, ok := block["pre_start_user_script"].(string); !ok || got != "" {
		t.Fatalf("normalized pre_start_user_script = %#v, want empty string", block["pre_start_user_script"])
	}
}
