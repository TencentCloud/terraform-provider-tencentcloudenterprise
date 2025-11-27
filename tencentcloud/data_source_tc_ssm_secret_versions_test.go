package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudSsmSecretVersionsDataSource(t *testing.T) {
	t.Parallel()
	dataSourceName := "data.cloud_ssm_secret_versions.secret_version"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccTencentCloudSsmSecretVersionsDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "secret_version_list.0.version_id", "v1"),
					resource.TestCheckResourceAttr(dataSourceName, "secret_version_list.0.secret_binary", "MTIzMTIzMTIzMTIzMTIzQQ=="),
				),
			},
		},
	})
}

const TestAccTencentCloudSsmSecretVersionsDataSourceConfig = `
resource "tencentcloudenterprise_ssm_secret" "secret" {
  secret_name = "unit-test-ver-data"
  description = "test secret"

  tags = {
    test-tag = "test"
  }
}

resource "tencentcloudenterprise_ssm_secret_version" "v1" {
  secret_name = cloud_ssm_secret.secret.secret_name
  version_id = "v1"
  secret_binary = "MTIzMTIzMTIzMTIzMTIzQQ=="
}

data "tencentcloudenterprise_ssm_secret_versions" "secret_version" {
  secret_name = cloud_ssm_secret_version.v1.secret_name
  version_id = cloud_ssm_secret_version.v1.version_id
}
`
