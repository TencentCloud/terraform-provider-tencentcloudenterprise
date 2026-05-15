package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDcInstanceResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDcInstance,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloudenterprise_dc_instance.instance", "id")),
			},
			{
				ResourceName:      "tencentcloudenterprise_dc_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccTencentCloudDcInstanceResource_isShare verifies that is_share can be set
// during creation (via Modify after Create) and updated without forcing resource recreation.
func TestAccTencentCloudDcInstanceResource_isShare(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Step 1: Create with is_share = true, verify it takes effect in one apply
				Config: testAccDcInstanceIsShareTrue,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dc_instance.is_share_test", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dc_instance.is_share_test", "is_share", "true"),
				),
			},
			{
				// Step 2: Update is_share to false, verify no recreation
				Config: testAccDcInstanceIsShareFalse,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dc_instance.is_share_test", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dc_instance.is_share_test", "is_share", "false"),
				),
			},
		},
	})
}

const testAccDcInstance = `

resource "tencentcloudenterprise_dc_instance" "instance" {
  access_point_id         = "ap-shenzhen-b-ft"
  bandwidth               = 10
  customer_contact_number = "0"
  direct_connect_name     = "terraform-for-test"
  line_operator           = "In-houseWiring"
  port_type               = "10GBase-LR"
  sign_law                = true
  vlan                    = -1
}

`

// TestAccTencentCloudDcInstanceResource_redundantWithIsShare verifies that creating a DC instance
// with redundant_direct_connect_id and is_share=true works in one apply, and updating is_share
// does not trigger recreation of the redundant DC.
func TestAccTencentCloudDcInstanceResource_redundantWithIsShare(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Step 1: Create DC with redundant + is_share=true in one apply
				Config: testAccDcInstanceRedundantIsShareTrue,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dc_instance.redundant", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dc_instance.redundant", "is_share", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dc_instance.redundant", "redundant_direct_connect_id", "dc-gofrv5vl"),
				),
			},
			{
				// Step 2: Update is_share to false, verify no recreation (redundant_direct_connect_id unchanged)
				Config: testAccDcInstanceRedundantIsShareFalse,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dc_instance.redundant", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dc_instance.redundant", "is_share", "false"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dc_instance.redundant", "redundant_direct_connect_id", "dc-gofrv5vl"),
				),
			},
		},
	})
}

const testAccDcInstanceRedundantIsShareTrue = `

resource "tencentcloudenterprise_dc_instance" "redundant" {
  direct_connect_name         = "tf-dc-redundant-isshare"
  access_point_id             = "1"
  line_operator               = "China Telecom"
  tencentcloudenterprise_port_type             = "100Base-T"
  location                    = "test-location"
  bandwidth                   = 100
  customer_name               = "test"
  customer_contact_mail       = "test@example.com"
  customer_contact_number     = "20-12345678901"
  idc_port_type               = "100Base-T"
  idc_city                    = "chengdu"
  redundant_direct_connect_id = "dc-gofrv5vl"
  is_share                    = true
}

`

const testAccDcInstanceRedundantIsShareFalse = `

resource "tencentcloudenterprise_dc_instance" "redundant" {
  direct_connect_name         = "tf-dc-redundant-isshare"
  access_point_id             = "1"
  line_operator               = "China Telecom"
  tencentcloudenterprise_port_type             = "100Base-T"
  location                    = "test-location"
  bandwidth                   = 100
  customer_name               = "test"
  customer_contact_mail       = "test@example.com"
  customer_contact_number     = "20-12345678901"
  idc_port_type               = "100Base-T"
  idc_city                    = "chengdu"
  redundant_direct_connect_id = "dc-gofrv5vl"
  is_share                    = false
}

`

const testAccDcInstanceIsShareTrue = `

resource "tencentcloudenterprise_dc_instance" "is_share_test" {
  direct_connect_name     = "terraform-dc-isshare-test"
  access_point_id         = "1"
  line_operator           = "China Telecom"
  tencentcloudenterprise_port_type         = "100Base-T"
  location                = "test-location"
  bandwidth               = 100
  customer_name           = "test"
  customer_contact_mail   = "test@example.com"
  customer_contact_number = "20-12345678901"
  idc_port_type           = "100Base-T"
  idc_city                = "chengdu"
  is_share                = true
}

`

const testAccDcInstanceIsShareFalse = `

resource "tencentcloudenterprise_dc_instance" "is_share_test" {
  direct_connect_name     = "terraform-dc-isshare-test"
  access_point_id         = "1"
  line_operator           = "China Telecom"
  tencentcloudenterprise_port_type         = "100Base-T"
  location                = "test-location"
  bandwidth               = 100
  customer_name           = "test"
  customer_contact_mail   = "test@example.com"
  customer_contact_number = "20-12345678901"
  idc_port_type           = "100Base-T"
  idc_city                = "chengdu"
  is_share                = false
}

`
