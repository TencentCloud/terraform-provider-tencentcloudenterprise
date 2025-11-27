package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNeedFixDcxExtraConfigResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDcxExtraConfig,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcx_extra_config.dcx_extra_config", "id")),
			},
			{
				ResourceName:      "tencentcloudenterprise_dcx_extra_config.dcx_extra_config",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccDcxExtraConfig = `

resource "tencentcloudenterprise_dcx_extra_config" "dcx_extra_config" {
  direct_connect_tunnel_id = "dcx-4z49tnws"
  vlan                     = 123
  bgp_peer {
    asn      = 65101
    auth_key = "test123"

  }
  route_filter_prefixes {
    cidr = "192.168.0.0/24"
  }
  tencent_address        = "203.0.113.101"
  tencent_backup_address = "203.0.113.102"
  customer_address       = "203.0.113.104"
  bandwidth              = 10
  enable_bgp_community   = false
  bfd_enable             = 0
  nqa_enable             = 1
  bfd_info {
    probe_failed_times = 3
    interval           = 100

  }
  nqa_info {
    probe_failed_times = 3
    interval           = 100
    destination_ip     = "203.0.113.202"

  }
  ipv6_enable = 0
  jumbo_enable = 0
}

`
