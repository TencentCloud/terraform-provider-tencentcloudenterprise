package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcEndPointServiceWhiteListResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcEndPointServiceWhiteList,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_vpc_end_point_service_white_list.end_point_service_white_list", "id")),
			},
			{
				ResourceName:      "cloud_vpc_end_point_service_white_list.end_point_service_white_list",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcEndPointServiceWhiteList = `

resource "cloud_vpc_end_point_service_white_list" "end_point_service_white_list" {
  user_uin = "100020512675"
  end_point_service_id = "vpcsvc-98jddhcz"
  description = "terraform for test"
}

`
