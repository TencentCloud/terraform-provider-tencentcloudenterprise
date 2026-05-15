package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudCfwBuyService_basic tests CFW service activation.
// NOTE: CFW buy service does not support delete. Once created, it persists.
func TestAccTencentCloudCfwBuyService_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(_ *terraform.State) error {
			// Delete is intentionally unsupported
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCfwBuyService,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cfw_buy_service.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cfw_buy_service.example", "status", "1"),
				),
			},
		},
	})
}

const testAccCfwBuyService = `
resource "tencentcloudenterprise_cfw_buy_service" "example" {
  region_id = "50000001"
  zone_id   = "50010001"
  vpc_spec  = "1"
}
`
