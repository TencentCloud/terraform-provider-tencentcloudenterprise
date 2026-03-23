package tencentcloud

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// 测试基础场景 - 创建、更新和删除 SAML 配置
func TestAccTencentCloudCamUserSamlConfig_basic(t *testing.T) {
	t.Parallel()

	testIdpName := getTestSamlIdpName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamUserSamlConfigDestroy,
		Steps: []resource.TestStep{
			// Step 1: 创建 SAML 配置
			{
				Config: testAccCamUserSamlConfigBasic(testIdpName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamUserSamlConfigExists("tencentcloudenterprise_cam_user_saml_config.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_user_saml_config.test", "idp_name", testIdpName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_user_saml_config.test", "protocol", "saml"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_user_saml_config.test", "idp_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_user_saml_config.test", "saml_metadata"),
				),
			},
			// Step 2: 更新 SAML 配置
			{
				Config: testAccCamUserSamlConfigUpdate(testIdpName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamUserSamlConfigExists("tencentcloudenterprise_cam_user_saml_config.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_user_saml_config.test", "idp_name", testIdpName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_user_saml_config.test", "remark", "Updated SAML configuration"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_user_saml_config.test", "is_sync_idp_user", "1"),
				),
			},
			// Step 3: 导入测试
			{
				ResourceName:      "tencentcloudenterprise_cam_user_saml_config.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func getTestSamlIdpName() string {
	testIdpName := os.Getenv("TENCENTCLOUD_TEST_SAML_IDP_NAME")
	if testIdpName == "" {
		// 默认使用测试 IdP 名称
		testIdpName = "terraform-test-saml-idp"
	}
	return testIdpName
}

func testAccCheckCamUserSamlConfigExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM user_saml_config][Exists] resource %s not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM user_saml_config][Exists] ID is not set")
		}

		openService := OpenService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		samlConfig, err := openService.DescribeSamlConfigByName(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}

		if samlConfig == nil {
			return fmt.Errorf("[CHECK][CAM user_saml_config][Exists] config not found")
		}

		return nil
	}
}

func testAccCheckCamUserSamlConfigDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_user_saml_config" {
			continue
		}

		openService := OpenService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		samlConfig, err := openService.DescribeSamlConfigByName(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}

		if samlConfig != nil {
			return fmt.Errorf("[CHECK][CAM user_saml_config][Destroy] config still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCamUserSamlConfigBasic(idpName string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_user_saml_config" "test" {
  idp_name      = "%s"
  protocol      = "saml"
  saml_metadata = <<-EOT
<?xml version="1.0" encoding="UTF-8"?>
<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://test-idp.example.com">
  <md:IDPSSODescriptor WantAuthnRequestsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
    <md:KeyDescriptor use="signing">
      <ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
        <ds:X509Data>
          <ds:X509Certificate>MIICmzCCAgQCCQDOg5QJ1Ms39jANBgkqhkiG9w0BAQsFADCBjzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCldhc2hpbmd0b24xEDAOBgNVBAcMB1NlYXR0bGUxEjAQBgNVBAoMCU15Q29tcGFueTETMBEGA1UECwwKSVQgU3VwcG9ydDEQMA4GA1UEAwwHdGVzdC1pZHAxHjAcBgkqhkiG9w0BCQEWD3Rlc3RAZXhhbXBsZS5jb20wHhcNMjMwMTAxMTIwMDAwWhcNMjQwMTAxMTIwMDAwWjCBjzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCldhc2hpbmd0b24xEDAOBgNVBAcMB1NlYXR0bGUxEjAQBgNVBAoMCU15Q29tcGFueTETMBEGA1UECwwKSVQgU3VwcG9ydDEQMA4GA1UEAwwHdGVzdC1pZHAxHjAcBgkqhkiG9w0BCQEWD3Rlc3RAZXhhbXBsZS5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBALCkJE3Gv3dKRqpOGqGqhNq0rPqmjF4QKnI7qT8fKXvhF0gCqT7bPqmVxQxZx2kGHI4dF6lPqm8QfKLqrTwGqM8bFqrQxP7vGqF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8AgMBAAEwDQYJKoZIhvcNAQELBQADgYEAkH1T8q7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxE=</ds:X509Certificate>
        </ds:X509Data>
      </ds:KeyInfo>
    </md:KeyDescriptor>
    <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://test-idp.example.com/saml/sso"/>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://test-idp.example.com/saml/sso"/>
  </md:IDPSSODescriptor>
</md:EntityDescriptor>
EOT
  remark        = "Test SAML configuration"
}
`, idpName)
}

func testAccCamUserSamlConfigUpdate(idpName string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_user_saml_config" "test" {
  idp_name          = "%s"
  protocol          = "saml"
  saml_metadata     = <<-EOT
<?xml version="1.0" encoding="UTF-8"?>
<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://test-idp.example.com">
  <md:IDPSSODescriptor WantAuthnRequestsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
    <md:KeyDescriptor use="signing">
      <ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
        <ds:X509Data>
          <ds:X509Certificate>MIICmzCCAgQCCQDOg5QJ1Ms39jANBgkqhkiG9w0BAQsFADCBjzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCldhc2hpbmd0b24xEDAOBgNVBAcMB1NlYXR0bGUxEjAQBgNVBAoMCU15Q29tcGFueTETMBEGA1UECwwKSVQgU3VwcG9ydDEQMA4GA1UEAwwHdGVzdC1pZHAxHjAcBgkqhkiG9w0BCQEWD3Rlc3RAZXhhbXBsZS5jb20wHhcNMjMwMTAxMTIwMDAwWhcNMjQwMTAxMTIwMDAwWjCBjzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCldhc2hpbmd0b24xEDAOBgNVBAcMB1NlYXR0bGUxEjAQBgNVBAoMCU15Q29tcGFueTETMBEGA1UECwwKSVQgU3VwcG9ydDEQMA4GA1UEAwwHdGVzdC1pZHAxHjAcBgkqhkiG9w0BCQEWD3Rlc3RAZXhhbXBsZS5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBALCkJE3Gv3dKRqpOGqGqhNq0rPqmjF4QKnI7qT8fKXvhF0gCqT7bPqmVxQxZx2kGHI4dF6lPqm8QfKLqrTwGqM8bFqrQxP7vGqF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8AgMBAAEwDQYJKoZIhvcNAQELBQADgYEAkH1T8q7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxF8wQ7qrTqrQxE=</ds:X509Certificate>
        </ds:X509Data>
      </ds:KeyInfo>
    </md:KeyDescriptor>
    <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://test-idp.example.com/saml/sso"/>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://test-idp.example.com/saml/sso"/>
  </md:IDPSSODescriptor>
</md:EntityDescriptor>
EOT
  remark            = "Updated SAML configuration"
  is_sync_idp_user  = 1
  assist_domain     = "example.com"
}
`, idpName)
}
