package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudSsmSecretsDataSource(t *testing.T) {
	t.Parallel()
	dataSourceName := "data.cloud_ssm_secrets.secret"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccTencentCloudSsmSecretsDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "secret_list.0.secret_name", "unit-test-for-data"),
					resource.TestCheckResourceAttr(dataSourceName, "secret_list.0.description", "test secret"),
					resource.TestCheckResourceAttrSet(dataSourceName, "secret_list.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "secret_list.0.create_uin"),
					resource.TestCheckResourceAttrSet(dataSourceName, "secret_list.0.status"),
					resource.TestCheckResourceAttrSet(dataSourceName, "secret_list.0.create_time"),
				),
			},
		},
	})
}

const TestAccTencentCloudSsmSecretsDataSourceConfig = `
resource "cloud_ssm_secret" "secret" {
  secret_name = "unit-test-for-data"
  description = "test secret"

  tags = {
    test-tag = "test"
  }
}

data "cloud_ssm_secrets" "secret" {
  secret_name = cloud_ssm_secret.secret.secret_name
  state = 1
  
  tags = {
    test-tag = "test"
  }
}
`
