package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudKmsBuyService_basic tests KMS service activation.
func TestAccTencentCloudKmsBuyService_basic(t *testing.T) {
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
				Config: testAccKmsBuyService,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_kms_buy_service.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_kms_buy_service.example", "service_enabled", "true"),
				),
			},
		},
	})
}

const testAccKmsBuyService = `
resource "tencentcloudenterprise_kms_buy_service" "example" {}
`
