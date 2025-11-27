package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudKubernetesChartsDataSource(t *testing.T) {
	t.Parallel()
	dataSourceName := "data.cloud_tke_kubernetes_charts.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKubernetesCharts,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "list.#"),
				),
			},
		},
	})
}

const testAccDataSourceKubernetesCharts = `
data "tencentcloudenterprise_tke_kubernetes_charts" "test" {
	result_output_file = "222.json"
}
`
