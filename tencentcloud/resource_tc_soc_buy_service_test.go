package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudSocBuyService_basic tests SOC service activation.
func TestAccTencentCloudSocBuyService_basic(t *testing.T) {
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
				Config: testAccSocBuyService,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_soc_buy_service.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_soc_buy_service.example", "buy_status", "true"),
				),
			},
		},
	})
}

const testAccSocBuyService = `
resource "tencentcloudenterprise_soc_buy_service" "example" {
  type = "premium"
  tce_area {
    region_id = 50000001
    zone_id   = 50010001
  }
}
`
