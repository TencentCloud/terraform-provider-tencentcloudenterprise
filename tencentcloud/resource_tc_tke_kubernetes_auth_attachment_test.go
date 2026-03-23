package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudTkeKubernetesAuthAttachmentResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTkeAuthAttachDefault,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "cluster_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "use_tke_default", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "auto_create_discovery_anonymous_auth", "true"),
				),
			},
			{
				Config: testAccTkeAuthAttachNonDefault,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "cluster_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "auto_create_discovery_anonymous_auth", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "jwks_uri", "https://ap-guangzhou-oidc.tke.tencentcs.com/id/7cbe7ca92eba3abc76a17de1/openid/v1/jwks"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "issuer", "https://ap-guangzhou-oidc.tke.tencentcs.com/id/7cbe7ca92eba3abc76a17de1"),
				),
			},
			{
				Config: testAccTkeAuthAttachOidcConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "cluster_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "use_tke_default", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "auto_create_discovery_anonymous_auth", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "auto_create_oidc_config", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_auth_attachment.test", "auto_install_pod_identity_webhook_addon", "true"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tke_kubernetes_auth_attachment.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_create_discovery_anonymous_auth",
				},
			},
		},
	})
}

const testAccTkeAuthAttachDefault = TkeDataSource + `

resource "tencentcloudenterprise_tke_kubernetes_auth_attachment" "test" {
  cluster_id                           = local.cluster_id
  use_tke_default                      = true
  auto_create_discovery_anonymous_auth = true
}
`

const testAccTkeAuthAttachNonDefault = TkeDataSource + `

resource "tencentcloudenterprise_tke_kubernetes_auth_attachment" "test" {
  cluster_id                           = local.cluster_id
  auto_create_discovery_anonymous_auth = true
  issuer                               = "https://ap-guangzhou-oidc.tke.tencentcs.com/id/7cbe7ca92eba3abc76a17de1"
  jwks_uri                             = "https://ap-guangzhou-oidc.tke.tencentcs.com/id/7cbe7ca92eba3abc76a17de1/openid/v1/jwks"
}
`

const testAccTkeAuthAttachOidcConfig = TkeDataSource + `

resource "tencentcloudenterprise_tke_kubernetes_auth_attachment" "test" {
  cluster_id                              = local.cluster_id
  use_tke_default                         = true
  auto_create_discovery_anonymous_auth    = true
  auto_create_oidc_config                 = true
  auto_create_client_id                   = ["xxx"]
  auto_install_pod_identity_webhook_addon = true
}
`
