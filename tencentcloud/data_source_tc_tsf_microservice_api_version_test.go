package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudTsfMicroserviceApiVersionDataSource_basic -v
func TestAccTencentCloudTsfMicroserviceApiVersionDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfMicroserviceApiVersionDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tsf_microservice_api_version.microservice_api_version"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_microservice_api_version.microservice_api_version", "result.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_microservice_api_version.microservice_api_version", "result.0.application_id"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_microservice_api_version.microservice_api_version", "result.0.application_name"),
					resource.TestCheckResourceAttrSet("data.cloud_tsf_microservice_api_version.microservice_api_version", "result.0.pkg_version"),
				),
			},
		},
	})
}

const testAccTsfMicroserviceApiVersionDataSource = `

data "tencentcloudenterprise_tsf_microservice_api_version" "microservice_api_version" {
  microservice_id = "ms-yq3jo6jd"
  path = "/printRequest"
  method = "get"
}

`
