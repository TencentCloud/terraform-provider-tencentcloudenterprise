package tencentcloud

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
)

func TestCicProvisionRoleConfigurationOperationSchema(t *testing.T) {
	operationResource := resourceTencentCloudCicProvisionRoleConfigurationOperation()
	if err := operationResource.InternalValidate(nil, true); err != nil {
		t.Fatalf("resource schema validation failed: %v", err)
	}
	resourceSchema := operationResource.Schema
	for _, name := range []string{"zone_id", "role_configuration_id", "target_type", "target_uin"} {
		field, ok := resourceSchema[name]
		if !ok {
			t.Fatalf("schema field %q is missing", name)
		}
		if !field.Optional {
			t.Fatalf("schema field %q must be optional", name)
		}
	}
	if !resourceSchema["zone_id"].Computed {
		t.Fatal("zone_id must be computed when it is discovered automatically")
	}
	if _, ok := resourceSchema["deployment_status"]; ok {
		t.Fatal("deployment_status must not be exposed as a resource argument")
	}
	for _, name := range []string{"provisioned_count", "provisioned"} {
		if field := resourceSchema[name]; field == nil || !field.Computed {
			t.Fatalf("schema field %q must be computed", name)
		}
	}
}

func TestValidateCicProvisionRoleConfigurationOperationTarget(t *testing.T) {
	resourceSchema := resourceTencentCloudCicProvisionRoleConfigurationOperation().Schema
	tests := []struct {
		name    string
		values  map[string]interface{}
		wantErr bool
	}{
		{name: "automatic deployment", values: map[string]interface{}{}},
		{name: "targeted deployment", values: map[string]interface{}{
			"role_configuration_id": "rc-test",
			"target_type":           "MemberUin",
			"target_uin":            110000000001,
		}},
		{name: "role only", values: map[string]interface{}{"role_configuration_id": "rc-test"}, wantErr: true},
		{name: "target type only", values: map[string]interface{}{"target_type": "MemberUin"}, wantErr: true},
		{name: "target uin only", values: map[string]interface{}{"target_uin": 110000000001}, wantErr: true},
		{name: "missing target uin", values: map[string]interface{}{
			"role_configuration_id": "rc-test",
			"target_type":           "MemberUin",
		}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, resourceSchema, tt.values)
			err := validateCicProvisionRoleConfigurationOperationTarget(d)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validate target error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCicProvisioningTargetKey(t *testing.T) {
	roleConfigurationID := "rc-test"
	memberType := "MemberUin"
	memberUin := int64(110000000001)
	target := &cic.RoleConfigurationProvisionings{RoleConfigurationId: &roleConfigurationID, TargetType: &memberType, TargetUin: &memberUin}
	if got, want := cicProvisioningTargetKey(target), "rc-test#MemberUin#110000000001"; got != want {
		t.Fatalf("unexpected provisioning target key: got %q, want %q", got, want)
	}
}

func TestAccTencentCloudCicProvisionRoleConfigurationOperationResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCicProvisionRoleConfigurationOperation,
				Check:  resource.ComposeTestCheckFunc(testAccCicProvisionRoleConfigurationOperationChecks()...),
			},
		},
	})
}

func testAccCicProvisionRoleConfigurationOperationChecks() []resource.TestCheckFunc {
	const resourceName = "tencentcloudenterprise_cic_provision_role_configuration_operation.example"
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttrSet(resourceName, "id"),
		resource.TestCheckResourceAttrSet(resourceName, "zone_id"),
		resource.TestCheckResourceAttrSet(resourceName, "provisioned_count"),
	}
	if expected := os.Getenv("TCE_CIC_ACC_EXPECTED_PROVISIONED_COUNT"); expected != "" {
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "provisioned_count", expected))
		if expected != "0" {
			checks = append(checks, resource.TestCheckResourceAttr(resourceName, "provisioned.0.deployment_status", "Deployed"))
		}
	}
	return checks
}

const testAccCicProvisionRoleConfigurationOperation = `
resource "tencentcloudenterprise_cic_provision_role_configuration_operation" "example" {
}
`
