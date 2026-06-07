package tencentcloud

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const defaultTkeEncryptionProtectionClusterId = "cls-pwyy5m0s"

func TestAccTencentCloudTkeKubernetesEncryptionProtection_basic(t *testing.T) {
	t.Parallel()

	resourceName := "tencentcloudenterprise_tke_kubernetes_encryption_protection.test"
	alias := fmt.Sprintf("tf-testacc-tke-enc-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTkeKubernetesEncryptionProtectionBasic(defaultTkeEncryptionProtectionClusterId, alias),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_id", defaultTkeEncryptionProtectionClusterId),
					resource.TestCheckResourceAttr(resourceName, "kms_configuration.#", "1"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"kms_configuration"},
			},
		},
	})
}

func testAccTkeKubernetesEncryptionProtectionBasic(clusterId, alias string) string {
	kmsRegion := os.Getenv(PROVIDER_REGION)
	if kmsRegion == "" {
		kmsRegion = defaultRegion
	}

	return fmt.Sprintf(`
resource "tencentcloudenterprise_kms_key" "test" {
  alias = %q
}

resource "tencentcloudenterprise_tke_kubernetes_encryption_protection" "test" {
  cluster_id = %q

  kms_configuration {
    key_id     = tencentcloudenterprise_kms_key.test.id
    kms_region = %q
  }
}
`, alias, clusterId, kmsRegion)
}
