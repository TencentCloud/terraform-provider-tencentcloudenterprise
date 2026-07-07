package tencentcloud

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudKubernetesClusterDataSource(t *testing.T) {
	t.Parallel()

	key := "data.tencentcloudenterprise_tke_kubernetes_clusters.name"

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

// TestAccTencentCloudKubernetesClusterKubeconfigDataSource 验证 kube_config / kube_config_intranet
// 能从 DescribeClusterKubeconfig (Type=INTERNETLB/INNERLB) 正确拉取并填充。
// 修复前：data source 注释掉了获取逻辑，两字段始终为空 → yamldecode("") 报 "missing start of document"。
func TestAccTencentCloudKubernetesClusterKubeconfigDataSource(t *testing.T) {
	t.Parallel()

	key := "data.tencentcloudenterprise_tke_kubernetes_clusters.kubeconfig"
	outFile := "/tmp/tke_kubeconfig_check.json"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDataSourceTencentCloudTkeKubeconfigStr, outFile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID(key),
					resource.TestCheckResourceAttrSet(key, "list.#"),
					resource.TestCheckResourceAttrSet(key, "list.0.kube_config"),
					resource.TestCheckResourceAttrSet(key, "list.0.kube_config_intranet"),
					// kubeconfig 必须是合法 YAML（含 clusters 段），而非空字符串
					resource.TestMatchResourceAttr(key, "list.0.kube_config", regexp.MustCompile(`clusters:`)),
					resource.TestMatchResourceAttr(key, "list.0.kube_config_intranet", regexp.MustCompile(`clusters:`)),
					// 验证 kubeconfig 是 helm/kubernetes provider 可用的标准格式：
					// server 字段非空 + 三段 base64 能 decode 出合法 PEM
					testAccCheckTkeKubeconfigHelmCompatible(outFile),
				),
			},
		},
	})
}

// testAccCheckTkeKubeconfigHelmCompatible 读 result_output_file，
// 校验 kube_config / kube_config_intranet 是 helm/kubernetes provider 可用的标准格式。
// 不实际连接集群，只做格式校验：
//  1. server 字段非空且 https:// 开头
//  2. certificate-authority-data / client-certificate-data / client-key-data 三段 base64 能 decode 出含 "BEGIN " 的 PEM
//  3. 注：内外网 server 是否相同取决于集群是否开启 LB，不做差异断言。
func testAccCheckTkeKubeconfigHelmCompatible(path string) resource.TestCheckFunc {
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
		first := list[0]
		publicCfg, _ := first["kube_config"].(string)
		intranetCfg, _ := first["kube_config_intranet"].(string)

		for _, cfg := range []struct{ name, body string }{
			{"kube_config", publicCfg},
			{"kube_config_intranet", intranetCfg},
		} {
			if cfg.body == "" {
				return fmt.Errorf("%s 为空", cfg.name)
			}
			server := extractYAMLField(cfg.body, "server:")
			if server == "" {
				return fmt.Errorf("%s: server 字段缺失", cfg.name)
			}
			if !strings.HasPrefix(server, "https://") {
				return fmt.Errorf("%s: server 不是 https:// 开头: %s", cfg.name, server)
			}
			for _, field := range []string{
				"certificate-authority-data:",
				"client-certificate-data:",
				"client-key-data:",
			} {
				b64 := extractYAMLField(cfg.body, field)
				if b64 == "" {
					return fmt.Errorf("%s: %s 缺失", cfg.name, field)
				}
				decoded, err := base64.StdEncoding.DecodeString(b64)
				if err != nil {
					return fmt.Errorf("%s: %s base64 decode 失败: %s", cfg.name, field, err)
				}
				if !strings.Contains(string(decoded), "BEGIN ") {
					return fmt.Errorf("%s: %s decode 后不是 PEM", cfg.name, field)
				}
			}
		}

		// 注：内外网 server 是否相同取决于集群是否开启了 LB，
		// 未开启时 Type=INNERLB/INTERNETLB 都会 fallback 到 master IP，
		// 这是集群 LB 状态决定的，不是 provider bug，故不做差异断言。

		return nil
	}
}

// extractYAMLField 从 kubeconfig YAML 文本里提取 "field: value" 的 value。
func extractYAMLField(kubeconfig, fieldWithColon string) string {
	re := regexp.MustCompile(regexp.QuoteMeta(fieldWithColon) + `\s*(\S+)`)
	m := re.FindStringSubmatch(kubeconfig)
	if len(m) < 2 {
		return ""
	}
	return strings.TrimSpace(m[1])
}

func TestAccTencentCloudKubernetesClusterTagsDataSource(t *testing.T) {
	t.Parallel()

	key := "data.tencentcloudenterprise_tke_kubernetes_clusters.tags"

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

// cls-2fjdorvc 是 v120 (kazakhstan-1) 环境里现存的 TKE 集群。
// 此用例验证修复后 kube_config / kube_config_intranet 能从
// DescribeClusterKubeconfig (Type=INTERNETLB/INNERLB) 正确拉取并填充。
// %s 由 result_output_file 路径填充，用于在 Check 阶段读回完整 kubeconfig 做格式校验。
const testAccDataSourceTencentCloudTkeKubeconfigStr = `
data "tencentcloudenterprise_tke_kubernetes_clusters" "kubeconfig" {
  cluster_id = "cls-2fjdorvc"
  result_output_file = "%s"
}
`
