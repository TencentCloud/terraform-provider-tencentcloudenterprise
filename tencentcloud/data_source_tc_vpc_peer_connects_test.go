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

// TestAccTencentCloudVpcPeerConnectsDataSource 验证 data source 能列出对等连接。
// 无过滤条件时列出全部；校验 list 字段格式标准。
func TestAccTencentCloudVpcPeerConnectsDataSource(t *testing.T) {
	t.Parallel()

	key := "data.tencentcloudenterprise_vpc_peer_connects.all"
	outFile := "/tmp/vpc_peer_connects_check.json"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDataSourceTencentCloudVpcPeerConnectsAll, outFile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID(key),
					resource.TestCheckResourceAttrSet(key, "list.#"),
					testAccCheckVpcPeerConnectsFormat(outFile),
				),
			},
		},
	})
}

// testAccCheckVpcPeerConnectsFormat 读 result_output_file，
// 校验每个对等连接的必填字段格式标准（非空、类型正确）。
func testAccCheckVpcPeerConnectsFormat(path string) resource.TestCheckFunc {
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
			id, _ := item["peering_connection_id"].(string)
			if id == "" {
				return fmt.Errorf("list[%d]: peering_connection_id 为空", i)
			}
			if !regexp.MustCompile(`^pcx-`).MatchString(id) {
				return fmt.Errorf("list[%d]: peering_connection_id 不是 pcx- 开头: %s", i, id)
			}
			state, _ := item["state"].(string)
			if state == "" {
				return fmt.Errorf("list[%d]: state 为空", i)
			}
			vpcId, _ := item["vpc_id"].(string)
			if vpcId == "" {
				return fmt.Errorf("list[%d]: vpc_id 为空", i)
			}
		}
		return nil
	}
}

const testAccDataSourceTencentCloudVpcPeerConnectsAll = `
data "tencentcloudenterprise_vpc_peer_connects" "all" {
  result_output_file = "%s"
}
`

// TestAccTencentCloudVpcPeerConnectsByIdDataSource 验证按 ID 精确查询。
func TestAccTencentCloudVpcPeerConnectsByIdDataSource(t *testing.T) {
	t.Parallel()

	key := "data.tencentcloudenterprise_vpc_peer_connects.by_id"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTencentCloudVpcPeerConnectsById,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID(key),
					resource.TestCheckResourceAttr(key, "list.#", "1"),
					resource.TestCheckResourceAttr(key, "list.0.peering_connection_id", "pcx-3j0yk582"),
					resource.TestCheckResourceAttr(key, "list.0.state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(key, "list.0.vpc_id"),
					resource.TestCheckResourceAttrSet(key, "list.0.peer_vpc_id"),
				),
			},
		},
	})
}

const testAccDataSourceTencentCloudVpcPeerConnectsById = `
data "tencentcloudenterprise_vpc_peer_connects" "by_id" {
  peering_connection_id = "pcx-3j0yk582"
}
`
