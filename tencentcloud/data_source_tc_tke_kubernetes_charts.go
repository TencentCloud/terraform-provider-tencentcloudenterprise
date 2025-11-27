/*
Use this data source to query detailed information of kubernetes cluster addons.

# Example Usage

```hcl
data "tencentcloudenterprise_tke_kubernetes_charts" "name" {}
```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tke_kubernetes_charts", CNDescription{
		TerraformTypeCN: "集群chart列表",
		DescriptionCN:   "提供TKE集群Charts数据源，用于查询TKE集群Charts的详细信息。",
		AttributesCN: map[string]string{
			"kind":               "有点像应用程序图表可用值：“日志”、“调度程序”、“网络”、“存储”、“监视器”、“dns”、“图像”、“其他”、“不可见”",
			"arch":               "支持操作系统应用程序可用值：`arm32`、`arm64`、`amd64`",
			"cluster_type":       "群集类型可用值：“tke”、“eks”",
			"result_output_file": "用于保存结果",
			"chart_list":         "应用程序图表列表",
			"search":             "搜索图表",
			"list":               "图表列表",
			"label":              "图表标签",
			"latest_version":     "图表最新版本",
			"name":               "图表名称",
		},
	})
}

func dataSourceTencentCloudKubernetesCharts() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of kubernetes cluster addons.",
		Read:        dataSourceTencentCloudKubernetesChartsRead,
		Schema: map[string]*schema.Schema{
			"search": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search of chart.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"list": {
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
		},
	}
}

func dataSourceTencentCloudKubernetesChartsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tke_kubernetes_charts.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	client := meta.(*TencentCloudClient).apiV3Conn
	service := TkeService{client: client}

	var search string
	if v, ok := d.GetOk("search"); ok {
		search = v.(string)
	}

	request := tke.NewDescribeHelmChartRequest()
	request.Search = &search

	response, err := service.DescribeHelmChart(ctx, request)
	if err != nil {
		return err
	}

	err = d.Set("list", response.Response.Results)

	if err != nil {
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		err = writeToFile(output.(string), response.Response.Results)
		if err != nil {
			return err
		}
	}

	//ids := []string{kind, arch, clusterType}
	//d.SetId("app_chart_" + helper.DataResourceIdsHash(ids))
	//
	return nil
}
