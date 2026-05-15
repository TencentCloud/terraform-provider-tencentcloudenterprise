package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudClsBuyService_basic tests CLS service activation.
func TestAccTencentCloudClsBuyService_basic(t *testing.T) {
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
				Config: testAccClsBuyService,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cls_buy_service.example", "id"),
				),
			},
		},
	})
}

const testAccClsBuyService = `
resource "tencentcloudenterprise_cls_buy_service" "example" {}
`
