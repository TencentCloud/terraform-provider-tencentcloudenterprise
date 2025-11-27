package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudKubernetesClusterDataSource(t *testing.T) {
	t.Parallel()

	key := "data.cloud_tke_kubernetes_clusters.name"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTencentCloudTkeStr,
				Check: resource.ComposeTestCheckFunc(
					// name filter
					testAccCheckTencentCloudDataSourceID(key),
					resource.TestCheckResourceAttr(key, "cluster_name", "test11"),
					resource.TestCheckResourceAttrSet(key, "list.#"),
				),
			},
		},
	})
}

func TestAccTencentCloudKubernetesClusterTagsDataSource(t *testing.T) {
	t.Parallel()

	key := "data.cloud_tke_kubernetes_clusters.tags"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTencentCloudTkeTags,
				Check: resource.ComposeTestCheckFunc(
					// tags filter
					testAccCheckTencentCloudDataSourceID(key),
					resource.TestCheckResourceAttrSet(key, "list.#"),
				),
			},
		},
	})
}

const testAccDataSourceTencentCloudTkeStr = `
data "tencentcloudenterprise_tke_kubernetes_clusters" "name" {
  #examples have been created to serve other resources
  cluster_name = "test11"
  result_output_file = "111.json"

  #tags = {
  #  "test" = "test"
  #}
}
`

const testAccDataSourceTencentCloudTkeTags = `
data "tencentcloudenterprise_tke_kubernetes_clusters" "tags" {
  #examples have been created to serve other resources
  tags = {
    "test" = "test"
  }
}
`
