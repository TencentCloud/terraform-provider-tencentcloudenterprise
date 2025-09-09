package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcSecurityGroupLimitsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcSecurityGroupLimitsDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_vpc_security_group_limits.security_group_limits")),
			},
		},
	})
}

const testAccVpcSecurityGroupLimitsDataSource = `

data "cloud_vpc_security_group_limits" "security_group_limits" {}

`
