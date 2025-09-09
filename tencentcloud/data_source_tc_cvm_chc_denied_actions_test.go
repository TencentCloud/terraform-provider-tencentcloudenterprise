package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCvmChcDeniedActionsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCvmChcDeniedActionsDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_cvm_chc_denied_actions.chc_denied_actions")),
			},
		},
	})
}

const testAccCvmChcDeniedActionsDataSource = `

data "cloud_cvm_chc_denied_actions" "chc_denied_actions" {
  chc_ids = ["chc-0brmw3wl"]
}
`
