package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudCfwBuyService_import tests CFW buy_service Read by importing
// an already-activated service. The resource ID is the region_id (e.g. "50000001").
func TestAccTencentCloudCfwBuyService_import(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:       testAccCfwBuyServicePlaceholder,
				ResourceName: "tencentcloudenterprise_cfw_buy_service.test",
				ImportState:  true,
				ImportStateId: "50000001",
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cfw_buy_service.test", "id"),
				),
			},
		},
	})
}

// Placeholder config — the resource is imported, not created
const testAccCfwBuyServicePlaceholder = `
# CFW buy_service is imported, not created
resource "tencentcloudenterprise_cfw_buy_service" "test" {
  region_id = "50000001"
  zone_id   = "50010001"
  vpc_spec  = "2"
}
`

// TestAccTencentCloudKmsBuyService_import tests KMS buy_service Read by importing.
func TestAccTencentCloudKmsBuyService_import(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(_ *terraform.State) error {
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccKmsBuyServicePlaceholder,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_kms_buy_service.test", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_kms_buy_service.test", "service_enabled", "true"),
				),
			},
		},
	})
}

const testAccKmsBuyServicePlaceholder = `
resource "tencentcloudenterprise_kms_buy_service" "test" {}
`

// TestAccTencentCloudSsmBuyService_import tests SSM buy_service Read by importing.
func TestAccTencentCloudSsmBuyService_import(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(_ *terraform.State) error {
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSsmBuyServicePlaceholder,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_ssm_buy_service.test", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ssm_buy_service.test", "service_enabled", "true"),
				),
			},
		},
	})
}

const testAccSsmBuyServicePlaceholder = `
resource "tencentcloudenterprise_ssm_buy_service" "test" {}
`

// TestAccTencentCloudSocBuyService_import tests SOC buy_service Read.
func TestAccTencentCloudSocBuyService_import(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(_ *terraform.State) error {
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccSocBuyServicePlaceholder,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_soc_buy_service.test", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_soc_buy_service.test", "buy_status", "true"),
				),
			},
		},
	})
}

const testAccSocBuyServicePlaceholder = `
resource "tencentcloudenterprise_soc_buy_service" "test" {
  type = "premium"
  tce_area {
    region_id = 50000001
    zone_id   = 50010001
  }
}
`

// TestAccTencentCloudCspOpenCosBilling_import tests CSP COS billing Read.
func TestAccTencentCloudCspOpenCosBilling_import(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(_ *terraform.State) error {
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCspOpenCosBillingPlaceholder,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_csp_open_cos_billing.test", "id"),
				),
			},
		},
	})
}

const testAccCspOpenCosBillingPlaceholder = `
resource "tencentcloudenterprise_csp_open_cos_billing" "test" {
  cos_region = "kazakhstan-1"
}
`
