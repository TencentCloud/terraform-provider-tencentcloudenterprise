package tencentcloud

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// 测试基础场景 - 创建、更新和删除 OIDC SSO 配置
func TestAccTencentCloudCamOIDCSSO_basic(t *testing.T) {
	t.Parallel()

	testIdpName := getTestOidcIdpName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamOIDCSSODestroy,
		Steps: []resource.TestStep{
			// Step 1: 创建 OIDC SSO 配置
			{
				Config: testAccCamOIDCSSOBasic(testIdpName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamOIDCSSOExists("tencentcloudenterprise_cam_oidc_sso.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_oidc_sso.test", "idp_name", testIdpName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_oidc_sso.test", "protocol", "oidc"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_oidc_sso.test", "response_type", "id_token"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_oidc_sso.test", "response_mode", "form_post"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_oidc_sso.test", "idp_id"),
				),
			},
			// Step 2: 更新 OIDC SSO 配置
			{
				Config: testAccCamOIDCSSOUpdate(testIdpName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamOIDCSSOExists("tencentcloudenterprise_cam_oidc_sso.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_oidc_sso.test", "idp_name", testIdpName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_oidc_sso.test", "remark", "Updated OIDC SSO configuration"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_oidc_sso.test", "is_sync_idp_user", "1"),
				),
			},
			// Step 3: 导入测试
			{
				ResourceName:            "tencentcloudenterprise_cam_oidc_sso.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"identity_key"}, // Sensitive field
			},
		},
	})
}

func getTestOidcIdpName() string {
	testIdpName := os.Getenv("TENCENTCLOUD_TEST_OIDC_IDP_NAME")
	if testIdpName == "" {
		// 默认使用测试 IdP 名称
		testIdpName = "terraform-test-oidc-idp"
	}
	return testIdpName
}

func testAccCheckCamOIDCSSOExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM oidc_sso][Exists] resource %s not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM oidc_sso][Exists] ID is not set")
		}

		openService := OpenService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		oidcConfig, err := openService.DescribeOidcConfigByName(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}

		if oidcConfig == nil {
			return fmt.Errorf("[CHECK][CAM oidc_sso][Exists] config not found")
		}

		return nil
	}
}

func testAccCheckCamOIDCSSODestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_oidc_sso" {
			continue
		}

		openService := OpenService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		oidcConfig, err := openService.DescribeOidcConfigByName(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}

		if oidcConfig != nil {
			return fmt.Errorf("[CHECK][CAM oidc_sso][Destroy] config still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCamOIDCSSOBasic(idpName string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_oidc_sso" "test" {
  idp_name                = "%s"
  protocol                = "oidc"
  identity_url            = "https://login.microsoftonline.com/test-tenant/v2.0"
  identity_key            = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN0RENDQWZ5Z0F3SUJBZ0lKQUl6Z0xhL3dkZzlxTUEwR0NTcUdTSWIzRFFFQkN3VUFNQkV4RHpBTkJnTlYKQkFNTUJtTmxjblJ6TUI0WERUSXpNRE13TVRFME1EQXdPRm9YRFRNek1ESXlOakUwTURBd09Gb3dFVEVQTUEwRwpBMVVFQXd3R1kyVnlkSE13Z1o4d0RRWUpLb1pJaHZjTkFRRUJCUUFEZ1kwQU1JR0pBb0dCQUt2Yk5YMFR2TkV4CkdqcHN2YmE1Y0VQRWJCVE9rNzBLdnVlcXkwTUp1UFJNQlNtNGRRM0p5NGxjSWxqRldNTWs4VlhIZjBJTzYxV2gKU1ZMNkpHb3EzM09LRWlMMy9GdUk5aFJ3cmhVdVBJbXBmKzJzREhhMHJGL0c1Q3JnYXYvZXhQWHQ5cjhqTTBCUgpoZUhDakRld1dpUVEyNlF0THJMNnJsS2dNSjNSQWdNQkFBR2pkVEJ6TUFrR0ExVWRFd1FDTUFBd0N3WURWUjBQCkJBUURBZ1hnTUIwR0ExVWREZ1FXQkJRN2dwaWhHVkRIaXhRZ2Q1V01YMTV4UVNjNHFqQWZCZ05WSFNNRUdEQVcKZ0JRN2dwaWhHVkRIaXhRZ2Q1V01YMTV4UVNjNHFqQVJCZ05WSFJFRUNqQUloZ1pqWlhKMGN6QU5CZ2txaGtpRwo5dzBCQVFzRkFBT0JnUUFsWUgwK1dGcFhvTk8rZjMvdjFTNDlHWDZGaVU1clFaTWtqS1hHMFNUWVdOdE1VWEFMCmlWMXhMcW9TdDlxcDUxZTJmb3FuQ3VrL3VWbGhJV0RTaUNiVGxRZUU5L3VlNXV0dWdIQWprSTZIM0lJdDN1c28KR1hBenZocWFpVzErWjdTWkNnYXVBUTE3ektIaU1XNTlVUGUvT2p1amFBa1ZZNWNDNThqMXdFdGxCUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
  client_id               = "12345678-1234-1234-1234-123456789012"
  authorization_endpoint  = "https://login.microsoftonline.com/test-tenant/oauth2/v2.0/authorize"
  response_type           = "id_token"
  response_mode           = "form_post"
  scope                   = ["openid", "email", "profile"]
  remark                  = "Test OIDC SSO configuration"
}
`, idpName)
}

func testAccCamOIDCSSOUpdate(idpName string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_oidc_sso" "test" {
  idp_name                = "%s"
  protocol                = "oidc"
  identity_url            = "https://login.microsoftonline.com/test-tenant/v2.0"
  identity_key            = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN0RENDQWZ5Z0F3SUJBZ0lKQUl6Z0xhL3dkZzlxTUEwR0NTcUdTSWIzRFFFQkN3VUFNQkV4RHpBTkJnTlYKQkFNTUJtTmxjblJ6TUI0WERUSXpNRE13TVRFME1EQXdPRm9YRFRNek1ESXlOakUwTURBd09Gb3dFVEVQTUEwRwpBMVVFQXd3R1kyVnlkSE13Z1o4d0RRWUpLb1pJaHZjTkFRRUJCUUFEZ1kwQU1JR0pBb0dCQUt2Yk5YMFR2TkV4CkdqcHN2YmE1Y0VQRWJCVE9rNzBLdnVlcXkwTUp1UFJNQlNtNGRRM0p5NGxjSWxqRldNTWs4VlhIZjBJTzYxV2gKU1ZMNkpHb3EzM09LRWlMMy9GdUk5aFJ3cmhVdVBJbXBmKzJzREhhMHJGL0c1Q3JnYXYvZXhQWHQ5cjhqTTBCUgpoZUhDakRld1dpUVEyNlF0THJMNnJsS2dNSjNSQWdNQkFBR2pkVEJ6TUFrR0ExVWRFd1FDTUFBd0N3WURWUjBQCkJBUURBZ1hnTUIwR0ExVWREZ1FXQkJRN2dwaWhHVkRIaXhRZ2Q1V01YMTV4UVNjNHFqQWZCZ05WSFNNRUdEQVcKZ0JRN2dwaWhHVkRIaXhRZ2Q1V01YMTV4UVNjNHFqQVJCZ05WSFJFRUNqQUloZ1pqWlhKMGN6QU5CZ2txaGtpRwo5dzBCQVFzRkFBT0JnUUFsWUgwK1dGcFhvTk8rZjMvdjFTNDlHWDZGaVU1clFaTWtqS1hHMFNUWVdOdE1VWEFMCmlWMXhMcW9TdDlxcDUxZTJmb3FuQ3VrL3VWbGhJV0RTaUNiVGxRZUU5L3VlNXV0dWdIQWprSTZIM0lJdDN1c28KR1hBenZocWFpVzErWjdTWkNnYXVBUTE3ektIaU1XNTlVUGUvT2p1amFBa1ZZNWNDNThqMXdFdGxCUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
  client_id               = "12345678-1234-1234-1234-123456789012"
  authorization_endpoint  = "https://login.microsoftonline.com/test-tenant/oauth2/v2.0/authorize"
  response_type           = "id_token"
  response_mode           = "form_post"
  scope                   = ["openid", "email", "profile"]
  remark                  = "Updated OIDC SSO configuration"
  is_sync_idp_user        = 1
  email_field             = "email"
  nick_name_field         = "name"
  phone_num_field         = "phone"
  login_account_field     = "preferred_username"
}
`, idpName)
}
