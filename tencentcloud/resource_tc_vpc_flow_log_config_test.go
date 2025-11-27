package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcFlowLogConfigResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcFlowLogConfig,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_flow_log_config.flow_log_config", "id")),
			},
			{
				Config: testAccVpcFlowLogConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_flow_log_config.flow_log_config", "enable", "true"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_flow_log_config.flow_log_config",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcFlowLogConfig = `

resource "tencentcloudenterprise_vpc_flow_log_config" "flow_log_config" {
  flow_log_id = "fl-geg2keoj"
  enable = false
}

`

const testAccVpcFlowLogConfigUpdate = `

resource "tencentcloudenterprise_vpc_flow_log_config" "flow_log_config" {
  flow_log_id = "fl-geg2keoj"
  enable = true
}

`
