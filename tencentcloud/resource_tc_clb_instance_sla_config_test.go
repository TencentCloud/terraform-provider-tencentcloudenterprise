package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNeedFixClbInstanceSlaConfigResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstanceSlaConfig,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_clb_instance_sla_config.instance_sla_config", "id")),
			},
			{
				ResourceName:      "cloud_clb_instance_sla_config.instance_sla_config",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccClbInstanceSlaConfig = `


resource "cloud_clb_instance_sla_config" "instance_sla_config" {
  load_balancer_id = "lb-5dnrkgry"
  sla_type         = "SLA"
}

`
