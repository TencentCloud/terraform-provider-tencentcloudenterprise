package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcNetDetectResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcNetDetect,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_net_detect.net_detect", "id")),
			},
			{
				Config: testAccVpcNetDetectUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_net_detect.net_detect", "net_detect_name", "terraform-for-test"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_net_detect.net_detect",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcNetDetect = `

resource "tencentcloudenterprise_vpc_net_detect" "net_detect" {
  net_detect_name       = "terrform-test"
  vpc_id                = "vpc-4owdpnwr"
  subnet_id             = "subnet-c1l35990"
  next_hop_destination  = "203.0.113.57"
  next_hop_type         = "NORMAL_CVM"
  detect_destination_ip = [
    "203.0.113.1",
    "203.0.113.2",
  ]
}

`

const testAccVpcNetDetectUpdate = `

resource "tencentcloudenterprise_vpc_net_detect" "net_detect" {
  net_detect_name       = "terraform-for-test"
  vpc_id                = "vpc-4owdpnwr"
  subnet_id             = "subnet-c1l35990"
  next_hop_destination  = "203.0.113.57"
  next_hop_type         = "NORMAL_CVM"
  detect_destination_ip = [
    "203.0.113.1",
    "203.0.113.2",
  ]
}

`
