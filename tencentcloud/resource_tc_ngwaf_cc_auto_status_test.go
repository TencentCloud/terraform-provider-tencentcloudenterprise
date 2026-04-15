//go:build ignore

package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNgwafCcAutoStatus_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNgwafCcAutoStatusBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_ngwaf_cc_auto_status.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_cc_auto_status.example", "domain", defaultNgwafCcAutoStatusDomain),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_cc_auto_status.example", "edition", defaultNgwafCcAutoStatusEdition),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_cc_auto_status.example", "status", "1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_ngwaf_cc_auto_status.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccNgwafCcAutoStatusBasic = defaultNgwafVariable + `
resource "tencentcloudenterprise_ngwaf_cc_auto_status" "example" {
  domain  = var.ngwaf_cc_auto_status_domain
  edition = var.ngwaf_cc_auto_status_edition
}
`
