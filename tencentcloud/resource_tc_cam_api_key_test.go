package tencentcloud

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCamApiKey_basic(t *testing.T) {
	t.Parallel()

	apiUin := "110000000285"
	rName := fmt.Sprintf("tf-cam-key-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamApiKeyDestroy(apiUin),
		Steps: []resource.TestStep{
			{
				Config: testAccCamApiKeyBasic(apiUin, rName, camAPIKeyEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamApiKeyExists("tencentcloudenterprise_cam_api_key.test", apiUin),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_api_key.test", "api_uin", apiUin),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_api_key.test", "status", camAPIKeyEnabled),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_api_key.test", "secret_id"),
				),
			},
			{
				Config: testAccCamApiKeyBasic(apiUin, rName, camAPIKeyDisabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamApiKeyExists("tencentcloudenterprise_cam_api_key.test", apiUin),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_api_key.test", "status", camAPIKeyDisabled),
				),
			},
			{
				ResourceName:            "tencentcloudenterprise_cam_api_key.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret_key"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources["tencentcloudenterprise_cam_api_key.test"]
					if !ok {
						return "", fmt.Errorf("resource not found in state")
					}
					return fmt.Sprintf("%s#%s", apiUin, rs.Primary.ID), nil
				},
			},
		},
	})
}

func testAccCheckCamApiKeyExists(n, apiUin string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource %s not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("resource %s id not set", n)
		}

		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		uin, err := strconv.ParseUint(apiUin, 10, 64)
		if err != nil {
			return err
		}
		keys, err := camService.QueryApiKey(context.Background(), uin)
		if err != nil {
			return err
		}
		for _, k := range keys {
			if k.SecretId != nil && *k.SecretId == rs.Primary.ID {
				return nil
			}
		}
		return fmt.Errorf("api key %s not found", rs.Primary.ID)
	}
}

func testAccCheckCamApiKeyDestroy(apiUin string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		uin, err := strconv.ParseUint(apiUin, 10, 64)
		if err != nil {
			return err
		}
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "tencentcloudenterprise_cam_api_key" {
				continue
			}
			keys, err := camService.QueryApiKey(context.Background(), uin)
			if err != nil {
				return err
			}
			for _, k := range keys {
				if k.SecretId != nil && *k.SecretId == rs.Primary.ID {
					return fmt.Errorf("api key still exists: %s", rs.Primary.ID)
				}
			}
		}
		return nil
	}
}

func testAccCamApiKeyBasic(apiUin, name, status string) string {
	_ = name
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_api_key" "test" {
  api_uin = "%s"
  status  = "%s"
}
`, apiUin, status)
}
