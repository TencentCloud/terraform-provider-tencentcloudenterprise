/*
Use this data source to query detailed information of kubernetes cluster addons.

# Example Usage

```hcl

	data "tencentcloudenterprise_tke_kubernetes_charts" "name" {
	  kind         = "network"
	  arch         = "amd64"
	  cluster_type = "tke"
	}

```
*/
package tencentcloud

import (
	"context"
	"encoding/json"
	"strings"

	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tke_kubernetes_charts", CNDescription{
		TerraformTypeCN: "集群chart列表",
		DescriptionCN:   "提供TKE集群Charts数据源，用于查询TKE集群Charts的详细信息。",
		AttributesCN: map[string]string{
			"kind":               "应用类型，可选值：log、scheduler、network、storage、monitor、dns、image、other、invisible",
			"arch":               "应用支持的操作系统，可选值：arm32、arm64、amd64",
			"cluster_type":       "集群类型，可选值：tke、eks",
			"result_output_file": "用于保存结果",
			"chart_list":         "应用程序图表列表",
			"name":               "图表名称",
			"label":              "图表标签",
			"latest_version":     "图表最新版本",
		},
	})
}

func dataSourceTencentCloudKubernetesCharts() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of kubernetes cluster addons.",
		Read:        dataSourceTencentCloudKubernetesChartsRead,
		Schema: map[string]*schema.Schema{
			"kind": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Kind of app chart. Available values: `log`, `scheduler`, `network`, `storage`, `monitor`, `dns`, `image`, `other`, `invisible`.",
			},
			"arch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Operation system app supported. Available values: `arm32`, `arm64`, `amd64`.",
			},
			"cluster_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster type. Available values: `tke`, `eks`.",
			},
			"chart_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "App chart list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of chart.",
						},
						"label": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Label of chart.",
						},
						"latest_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Chart latest version.",
						},
					},
				},
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudKubernetesChartsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_tke_kubernetes_charts.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	client := meta.(*TencentCloudClient).apiV3Conn
	service := TkeService{client: client}

	var (
		kind        string
		arch        string
		clusterType string
	)
	if v, ok := d.GetOk("kind"); ok {
		kind = v.(string)
	}
	if v, ok := d.GetOk("arch"); ok {
		arch = v.(string)
	}
	if v, ok := d.GetOk("cluster_type"); ok {
		clusterType = v.(string)
	}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("kind"); ok {
		paramMap["Kind"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("arch"); ok {
		paramMap["Arch"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("cluster_type"); ok {
		paramMap["ClusterType"] = helper.String(v.(string))
	}

	var respData []*tke.AppChart
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeKubernetesChartsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		respData = result
		return nil
	})
	if err != nil {
		return err
	}

	appChartsList := make([]map[string]interface{}, 0, len(respData))
	if respData != nil {
		for _, appCharts := range respData {
			appChartsMap := map[string]interface{}{}

			if appCharts.Name != nil {
				appChartsMap["name"] = appCharts.Name
			}

			if appCharts.LatestVersion != nil {
				appChartsMap["latest_version"] = appCharts.LatestVersion
			}

			if appCharts.Label != nil {
				tmpMap := make(map[string]interface{})
				if err := json.Unmarshal([]byte(*appCharts.Label), &tmpMap); err != nil {
					return err
				}
				appChartsMap["label"] = tmpMap
			}

			appChartsList = append(appChartsList, appChartsMap)
		}

		_ = d.Set("chart_list", appChartsList)
	}

	d.SetId(strings.Join([]string{kind, arch, clusterType}, FILED_SP))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), appChartsList); e != nil {
			return e
		}
	}

	return nil
}
