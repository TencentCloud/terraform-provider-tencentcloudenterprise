package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudTdmqEnvironmentAttributesDataSource_basic -v
func TestAccTencentCloudTdmqPulsarEnvironmentsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTdmqPulsarEnvironmentsDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tdmq_pulsar_environments.example"),
				),
			},
		},
	})
}

const testAccTdmqPulsarEnvironmentsDataSource = `
data "tencentcloudenterprise_tdmq_pulsar_environments" "example" {
    environment_id = "keep-ns"
    cluster_id     = "pulsar-9n95ax58b9vn"
}
`
