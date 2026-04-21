package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var testCicIdentityCenterDataSourceName = "data.tencentcloudenterprise_cic_identity_center.test"

// go test -i; go test -test.run TestAccTencentCloudCicIdentityCenterDataSource_basic -v
func TestAccTencentCloudCicIdentityCenterDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCicIdentityCenterDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(testCicIdentityCenterDataSourceName, "service_status"),
					resource.TestCheckResourceAttrSet(testCicIdentityCenterDataSourceName, "zone_id"),
					resource.TestCheckResourceAttrSet(testCicIdentityCenterDataSourceName, "zone_name"),
					resource.TestCheckResourceAttrSet(testCicIdentityCenterDataSourceName, "create_time"),
					resource.TestCheckResourceAttrSet(testCicIdentityCenterDataSourceName, "update_time"),
					resource.TestCheckResourceAttrSet(testCicIdentityCenterDataSourceName, "scim_sync_status"),
				),
			},
		},
	})
}

const testAccCicIdentityCenterDataSource = `
data "tencentcloudenterprise_cic_identity_center" "test" {
}
`
