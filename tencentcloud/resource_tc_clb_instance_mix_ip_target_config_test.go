package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNeedFixClbInstanceMixIpTargetConfigResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstanceMixIpTargetConfig,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_clb_instance_mix_ip_target_config.instance_mix_ip_target_config", "id")),
			},
			{
				ResourceName:      "cloud_clb_instance_mix_ip_target_config.instance_mix_ip_target_config",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccClbInstanceMixIpTargetConfig = `

resource "cloud_clb_instance_mix_ip_target_config" "instance_mix_ip_target_config" {
  load_balancer_id = "lb-5dnrkgry"
  mix_ip_target = false
}

`
