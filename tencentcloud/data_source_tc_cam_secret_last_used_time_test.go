package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamSecretLastUsedTimeDataSource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamSecretLastUsedTimeDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloudenterprise_cam_secret_last_used_time.example"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_secret_last_used_time.example", "secret_id_last_used_rows.#"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_secret_last_used_time.example", "secret_id_last_used_rows.0.secret_id"),
					resource.TestCheckResourceAttr("data.tencentcloudenterprise_cam_secret_last_used_time.example", "secret_id_last_used_rows.0.secret_id", "your-secret-id-here"),
				),
			},
		},
	})
}

const testAccCamSecretLastUsedTimeDataSource = `
data "tencentcloudenterprise_cam_secret_last_used_time" "example" {
  secret_id_list = ["your-secret-id-here"]
}
`
