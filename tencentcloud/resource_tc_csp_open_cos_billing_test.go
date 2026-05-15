package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudCspOpenCosBilling_basic tests COS billing activation.
func TestAccTencentCloudCspOpenCosBilling_basic(t *testing.T) {
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
				Config: testAccCspOpenCosBilling,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_csp_open_cos_billing.example", "id"),
				),
			},
		},
	})
}

const testAccCspOpenCosBilling = `
resource "tencentcloudenterprise_csp_open_cos_billing" "example" {
  cos_region = "kazakhstan-1"
}
`
