package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudKubernetesAvailableClusterVersionsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccKubernetesAvailableClusterVersionsDataSource_basic, defaultTkeClusterId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tke_kubernetes_available_cluster_versions.id"),
					resource.TestCheckResourceAttrSet("data.cloud_tke_kubernetes_available_cluster_versions.id", "versions.#"),
					resource.TestCheckResourceAttrSet("data.cloud_tke_kubernetes_available_cluster_versions.id", "clusters.#"),
				),
			},
			{
				Config: fmt.Sprintf(testAccKubernetesAvailableClusterVersionsDataSource_multiple, defaultTkeClusterId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tke_kubernetes_available_cluster_versions.ids"),
					resource.TestCheckResourceAttrSet("data.cloud_tke_kubernetes_available_cluster_versions.ids", "clusters.#"),
					resource.TestCheckResourceAttr("data.cloud_tke_kubernetes_available_cluster_versions.ids", "clusters.0.cluster_id", defaultTkeClusterId),
				),
			},
		},
	})
}

const testAccKubernetesAvailableClusterVersionsDataSource_basic = `

data "cloud_tke_kubernetes_available_cluster_versions" "id" {
  result_output_file = "versions.json"
  cluster_id = "%s"
}

output "versions"{
  value = data.cloud_tke_kubernetes_available_cluster_versions.id.versions
}

`

const testAccKubernetesAvailableClusterVersionsDataSource_multiple = `

data "cloud_tke_kubernetes_available_cluster_versions" "ids" {
  cluster_ids = ["%s"]
}

`
