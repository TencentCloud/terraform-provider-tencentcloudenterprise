package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudApmInstanceResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccApmInstance,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloudenterprise_apm_instance.instance", "id")),
			},
			{
				Config: testAccApmInstanceUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_apm_instance.instance", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_apm_instance.instance", "name", "terraform-for-test"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_apm_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccApmInstance = `

resource "tencentcloudenterprise_apm_instance" "instance" {
  name = "terraform-test"
  description = "for terraform test"
  trace_duration = 15
  span_daily_counters = 20
}

`

const testAccApmInstanceUpdate = `

resource "tencentcloudenterprise_apm_instance" "instance" {
  name = "terraform-for-test"
  description = "for terraform test"
  trace_duration = 15
  span_daily_counters = 20
}

`
