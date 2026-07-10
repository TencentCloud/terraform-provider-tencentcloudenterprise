package tencentcloud

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccTencentCloudCcnRouteTablesDataSource 验证 data source 能列出 CCN 路由表。
func TestAccTencentCloudCcnRouteTablesDataSource(t *testing.T) {
	t.Parallel()

	key := "data.tencentcloudenterprise_ccn_route_tables.all"
	outFile := "/tmp/ccn_route_tables_check.json"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDataSourceTencentCloudCcnRouteTablesAll, outFile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID(key),
					resource.TestCheckResourceAttrSet(key, "list.#"),
					testAccCheckCcnRouteTablesFormat(outFile),
				),
			},
		},
	})
}

// testAccCheckCcnRouteTablesFormat 读 result_output_file，校验字段格式标准。
func testAccCheckCcnRouteTablesFormat(path string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read result_output_file %s: %s", path, err)
		}
		var list []map[string]interface{}
		if err := json.Unmarshal(data, &list); err != nil {
			return fmt.Errorf("unmarshal result_output_file: %s", err)
		}
		if len(list) == 0 {
			return fmt.Errorf("result_output_file: empty list")
		}
		for i, item := range list {
			id, _ := item["route_table_id"].(string)
			if id == "" {
				return fmt.Errorf("list[%d]: route_table_id 为空", i)
			}
			if !regexp.MustCompile(`^ccnrtb-`).MatchString(id) {
				return fmt.Errorf("list[%d]: route_table_id 不是 ccnrtb- 开头: %s", i, id)
			}
			ccnId, _ := item["ccn_id"].(string)
			if ccnId == "" {
				return fmt.Errorf("list[%d]: ccn_id 为空", i)
			}
			if !regexp.MustCompile(`^ccn-`).MatchString(ccnId) {
				return fmt.Errorf("list[%d]: ccn_id 不是 ccn- 开头: %s", i, ccnId)
			}
		}
		return nil
	}
}

const testAccDataSourceTencentCloudCcnRouteTablesAll = `
data "tencentcloudenterprise_ccn_route_tables" "all" {
  result_output_file = "%s"
}
`

// TestAccTencentCloudCcnRouteTablesByIdDataSource 验证按路由表 ID 精确查询。
func TestAccTencentCloudCcnRouteTablesByIdDataSource(t *testing.T) {
	t.Parallel()

	key := "data.tencentcloudenterprise_ccn_route_tables.by_id"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTencentCloudCcnRouteTablesById,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID(key),
					resource.TestCheckResourceAttr(key, "list.#", "1"),
					resource.TestCheckResourceAttr(key, "list.0.route_table_id", "ccnrtb-l3o19hjh"),
					resource.TestCheckResourceAttr(key, "list.0.ccn_id", "ccn-pe0it21t"),
					resource.TestCheckResourceAttr(key, "list.0.is_default_table", "true"),
				),
			},
		},
	})
}

const testAccDataSourceTencentCloudCcnRouteTablesById = `
data "tencentcloudenterprise_ccn_route_tables" "by_id" {
  route_table_id = "ccnrtb-l3o19hjh"
}
`
