package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudTsfApplicationFileConfigResource_basic -v
func TestAccTencentCloudTsfApplicationFileConfigResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		//PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTsfApplicationFileConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfApplicationFileConfig2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsfApplicationFileConfigExists("cloud_tsf_application_file_config.application_file_config"),
					resource.TestCheckResourceAttrSet("cloud_tsf_application_file_config.application_file_config", "id"),
					//resource.TestCheckResourceAttr("cloud_tsf_application_file_config.application_file_config", "application_id", defaultTsfApplicationId),
					resource.TestCheckResourceAttr("cloud_tsf_application_file_config.application_file_config", "config_name", "terraform-test"),
					resource.TestCheckResourceAttr("cloud_tsf_application_file_config.application_file_config", "config_file_code", "UTF-8"),
					resource.TestCheckResourceAttr("cloud_tsf_application_file_config.application_file_config", "config_file_name", "application.yaml"),
					resource.TestCheckResourceAttr("cloud_tsf_application_file_config.application_file_config", "config_file_path", "/etc/nginx"),
					resource.TestCheckResourceAttr("cloud_tsf_application_file_config.application_file_config", "config_file_value", "test: 1"),
					resource.TestCheckResourceAttr("cloud_tsf_application_file_config.application_file_config", "config_post_cmd", "source .bashrc"),
					resource.TestCheckResourceAttr("cloud_tsf_application_file_config.application_file_config", "config_version", "1.0"),
					resource.TestCheckResourceAttr("cloud_tsf_application_file_config.application_file_config", "config_version_desc", "1.0"),
				),
			},
		},
	})
}

func testAccCheckTsfApplicationFileConfigDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloud_tsf_application_file_config" {
			continue
		}

		res, err := service.DescribeTsfApplicationFileConfigById(ctx, rs.Primary.ID, "")
		if err != nil {
			return err
		}

		if res != nil {
			return fmt.Errorf("tsf ApplicationFileConfig %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckTsfApplicationFileConfigExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		res, err := service.DescribeTsfApplicationFileConfigById(ctx, rs.Primary.ID, "")
		if err != nil {
			return err
		}

		if res == nil {
			return fmt.Errorf("tsf ApplicationFileConfig %s is not found", rs.Primary.ID)
		}

		return nil
	}
}

const testAccTsfApplicationFileConfigVar = `
variable "application_id" {
	default = "` + defaultTsfApplicationId + `"
}

`

const testAccTsfApplicationFileConfig = testAccTsfApplicationFileConfigVar + `

resource "cloud_tsf_application_file_config" "application_file_config" {
	config_name = "terraform-test"
	config_version = "1.0"
	config_file_name = "application.yaml"
	config_file_value = "test: 1"
	application_id = var.application_id
	config_file_path = "/etc/nginx"
	config_version_desc = "1.0"
	config_file_code = "UTF-8"
	config_post_cmd = "source .bashrc"
	encode_with_base64 = true
}

`
const testAccTsfApplicationFileConfig2 = `
resource "cloud_tsf_application_file_config" "application_file_config" {
	config_name = "terraform-test"
	config_version = "1.0"
	config_file_name = "application.yaml"
	config_file_value = "test: 1"
	application_id = "application-dapbqkyo"
	config_file_path = "/etc/nginx"
	config_version_desc = "1.0"
	config_file_code = "UTF-8"
	config_post_cmd = "source .bashrc"
	encode_with_base64 = true
}
`
