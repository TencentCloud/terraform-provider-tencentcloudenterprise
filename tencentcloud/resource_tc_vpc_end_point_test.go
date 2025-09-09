package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcEndPointResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcEndPoint,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_vpc_end_point.end_point", "id")),
			},
			{
				ResourceName:      "cloud_vpc_end_point.end_point",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcEndPoint = `

resource "cloud_vpc_end_point" "end_point" {
  vpc_id = "vpc-391sv4w3"
  subnet_id = "subnet-ljyn7h30"
  end_point_name = "terraform-test"
  end_point_service_id = "vpcsvc-98jddhcz"
  end_point_vip = "203.0.113.22"
}

`
