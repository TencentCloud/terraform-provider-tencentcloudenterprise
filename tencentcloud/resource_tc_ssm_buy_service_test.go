package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudSsmBuyService_basic tests SSM service activation.
func TestAccTencentCloudSsmBuyService_basic(t *testing.T) {
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
				Config: testAccSsmBuyService,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_ssm_buy_service.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ssm_buy_service.example", "service_enabled", "true"),
				),
			},
		},
	})
}

const testAccSsmBuyService = `
resource "tencentcloudenterprise_ssm_buy_service" "example" {}
`
