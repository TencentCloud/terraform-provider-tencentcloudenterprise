/*
Provide a datasource to query cluster CommonNames.

# Example Usage

```hcl

	data "tencentcloudenterprise_tke_kubernetes_cluster_common_names" "foo" {
	  cluster_id = "cls-tddgs8fl"
	  subaccount_uins = ["100004610790", "0987654321"]
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_common_names", CNDescription{
		TerraformTypeCN: "集群CommonNames",
		DescriptionCN:   "提供TKE集群CommonNames数据源，用于查询TKE集群CommonNames的详细信息。",
		AttributesCN: map[string]string{
			"cluster_id":         "群集ID",
			"subaccount_uins":    "子账户列表一次最多可以传入50个子帐户",
			"result_output_file": "用于保存结果",
			"list":               "子账户UIN对应的客户端证书中的CommonName列表",
			"subaccount_uin":     "用户UIN",
			"common_names":       "子账户对应的客户端证书中的CommonName",
			"cn":                 "子账户UIN对应的客户端证书中的CommonName",
		},
	})
}

func datasourceTencentCloudKubernetesClusterCommonNames() *schema.Resource {
	return &schema.Resource{
		Description: "Cloud TKE Cluster CommonNames",
		Read:        datasourceTencentCloudKubernetesClusterCommonNamesRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster ID.",
			},
			"subaccount_uins": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of sub-account. Up to 50 sub-accounts can be passed in at a time.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used for save result.",
			},
			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of the CommonName in the certificate of the client corresponding to the sub-account UIN.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subaccount_uin": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User UIN.",
						},
						"cn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cn of the client certificate corresponding to the sub-account UIN.",
						},
					},
				},
			},
		},
	}
}

func datasourceTencentCloudKubernetesClusterCommonNamesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("datasource.cloud_tke_kubernetes_cluster_common_names.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	client := meta.(*TencentCloudClient).apiV3Conn
	service := TkeService{client}

	clusterId := d.Get("cluster_id").(string)
	request := tke.NewDescribeClusterCommonNamesRequest()
	request.ClusterId = &clusterId

	if v, ok := d.GetOk("subaccount_uins"); ok {
		request.SubaccountUins = helper.InterfacesStringsPoint(v.([]interface{}))
	}
	//if v, ok := d.GetOk("role_ids"); ok {
	//	request.RoleIds = helper.InterfacesStringsPoint(v.([]interface{}))
	//}

	names, err := service.DescribeClusterCommonNames(ctx, request)

	if err != nil {
		return err
	}

	list := make([]map[string]interface{}, 0, len(names))
	for _, v := range names {
		mapping := map[string]interface{}{
			"subaccount_uin": v.SubaccountUin,
			"cn":             v.CN,
		}
		list = append(list, mapping)
	}

	if err = d.Set("list", list); err != nil {
		return err
	}

	d.SetId(clusterId)

	if output, ok := d.GetOk("result_output_file"); ok {
		return writeToFile(output.(string), list)
	}

	return nil
}
