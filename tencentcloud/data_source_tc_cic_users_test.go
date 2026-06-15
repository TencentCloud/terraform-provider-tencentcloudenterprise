package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var testCicUsersDataSourceName = "data.tencentcloudenterprise_cic_users.test"

// go test -i; go test -test.run TestAccTencentCloudCicUsersDataSource_basic -v
func TestAccTencentCloudCicUsersDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCicUsersDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(testCicUsersDataSourceName, "zone_id"),
					resource.TestCheckResourceAttrSet(testCicUsersDataSourceName, "users.#"),
				),
			},
		},
	})
}

const testAccCicUsersDataSource = `
data "tencentcloudenterprise_cic_identity_center" "test" {
}

data "tencentcloudenterprise_cic_users" "test" {
  zone_id = data.tencentcloudenterprise_cic_identity_center.test.zone_id
}
`
