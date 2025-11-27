/*
Use this data source to query detailed information of clb resources

Example Usage
```hcl

	data "tencentcloudenterprise_clb_resources" "resources" {
	  filters {
	    name = "isp"
	    values = ["BGP"]
	  }
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_clb_resources", CNDescription{
		TerraformTypeCN: "负载均衡资源列表",
		DescriptionCN:   "提供负载均衡资源列表数据源，用于查询CLB资源的详细信息。",
		AttributesCN: map[string]string{
			"filters":            "查询资源列表的过滤器",
			"name":               "过滤名称",
			"values":             "值",
			"resource_set":       "负载均衡实例列表",
			"load_balancer_id":   "负载均衡实例ID",
			"load_balancer_name": "负载均衡实例名称",
			"load_balancer_type": "负载均衡实例的网络类型",
			"domain":             "负载均衡实例的域名",
			"master_zone":        "负载均衡实例的主可用区",
			"zone_id":            "可用区ID",
			"zone":               "可用区名称",
			"zone_name":          "可用区名称",
			"zone_region":        "可用区所属地域",
			"local_zone":         "是否为本地可用区",
			"load_balancer_vips": "负载均衡实例的VIP列表",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudClbResources() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of clb resources",
		Read:        dataSourceTencentCloudClbResourcesRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Filter to query the list of AZ resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Filter name.",
						},
						"values": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required:    true,
							Description: "Filter value array.",
						},
					},
				},
			},

			"resource_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "List of resources supported by the AZ.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id of clb.",
						},
						"load_balancer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of clb.",
						},
						"load_balancer_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of clb.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain of clb.",
						},
						"master_zone": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Master zone of clb.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Zone id of clb.",
									},
									"zone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Zone of clb.",
									},
									"zone_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Zone name of clb.",
									},
									"zone_region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Zone region of clb.",
									},
									"local_zone": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the local zone.",
									},
								},
							},
						},
						"load_balancer_vips": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Vips of clb.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

func dataSourceTencentCloudClbResourcesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_clb_resources.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("filters"); ok {
		filtersSet := v.([]interface{})
		tmpSet := make([]*clb.Filter, 0, len(filtersSet))

		for _, item := range filtersSet {
			filter := clb.Filter{}
			filterMap := item.(map[string]interface{})

			if v, ok := filterMap["name"]; ok {
				filter.Name = helper.String(v.(string))
			}
			if v, ok := filterMap["values"]; ok {
				valuesSet := v.(*schema.Set).List()
				filter.Values = helper.InterfacesStringsPoint(valuesSet)
			}
			tmpSet = append(tmpSet, &filter)
		}
		paramMap["Filters"] = tmpSet
	}

	service := ClbService{client: meta.(*TencentCloudClient).apiV3Conn}

	var resources []*clb.LoadBalancer

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeClbResourcesByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		resources = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(resources))

	resourceSetList := make([]map[string]interface{}, 0, len(resources))
	if resources != nil {
		for _, resource := range resources {
			masterZone := make([]map[string]interface{}, 1)
			if resource.MasterZone != nil {
				mz := map[string]interface{}{
					"zone_id":     resource.MasterZone.ZoneId,
					"zone":        resource.MasterZone.Zone,
					"zone_name":   resource.MasterZone.ZoneName,
					"zone_region": resource.MasterZone.ZoneRegion,
					"local_zone":  resource.MasterZone.LocalZone,
				}
				masterZone[0] = mz
			}

			m := map[string]interface{}{
				"load_balancer_id":   resource.LoadBalancerId,
				"load_balancer_name": resource.LoadBalancerName,
				"load_balancer_type": resource.LoadBalancerType,
				"domain":             resource.Domain,
				"master_zone":        masterZone,
				"load_balancer_vips": resource.LoadBalancerVips,
			}
			resourceSetList = append(resourceSetList, m)
			ids = append(ids, *resource.LoadBalancerId)
		}

		_ = d.Set("resource_set", resourceSetList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), resourceSetList); e != nil {
			return e
		}
	}
	return nil
}
