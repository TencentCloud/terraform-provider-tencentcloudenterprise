/*
Use this data source to query bms placement group.

# Example Usage

```hcl

	data "tencentcloudenterprise_bms_flavors" "flavors" {
	  result_output_file = "flavors.json"
	}

```
*/
package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bms "terraform-provider-tencentcloudenterprise/sdk/bms/v20180813"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_bms_flavors", CNDescription{
		TerraformTypeCN: "套餐列表",
		DescriptionCN:   "提供BMS套餐列表数据源，用于查询BMS套餐规格列表信息。",
		AttributesCN: map[string]string{
			"flavor_ids":         "套餐 ID",
			"zone":               "可用区",
			"cpu_arch":           "CPU 架构",
			"result_output_file": "用于保存结果",
		},
	})
}
func dataTencentCloudBmsFlavors() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query bms placement group.",
		Read:        dataTencentCloudBmsFlavorsRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"flavor_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The IDs of the flavors.",
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "The ID of the flavor.",
				},
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone of the flavor.",
			},
			"cpu_arch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The cpu_arch of the flavor, `X86` and `ARM` are optional.",
			},

			// Computed values.
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The file path to output the result.",
			},

			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of flavors.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flavor_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the flavor.",
						},
						"flavor_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the flavor.",
						},
						"flavor_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the flavor.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone of the flavor.",
						},
						"raid_type": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The raid type of the flavor.",
							Elem: &schema.Schema{
								Type:        schema.TypeString,
								Description: "The raid type of the flavor.",
							},
						},
						//"operating_system": {
						//	Type:        schema.TypeList,
						//	Computed:    true,
						//	Description: "The operating system of the flavor.",
						//},
						"cpu": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cpu of the flavor.",
						},
						"cpu_arch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cpu_arch of the flavor.",
						},
						"memory": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The memory of the flavor.",
						},
						"system_disk": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The system disk of the flavor.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the flavor.",
						},
					},
				},
			},
		},
	}
}

func dataTencentCloudBmsFlavorsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_bms_instance.read")()
	defer inconsistentCheck(d, meta)()

	client := meta.(*TencentCloudClient).apiV3Conn

	var flavorSet []*bms.Flavor
	var request = bms.NewDescribeFlavorsRequest()
	request.Filters = []*bms.Filter{}
	if v, ok := d.GetOk("cpu_arch"); ok && v.(string) != "" {
		request.Filters = append(request.Filters, &bms.Filter{
			Name:   helper.String("cpu-arch"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if v, ok := d.GetOk("zone"); ok && v.(string) != "" {
		request.Filters = append(request.Filters, &bms.Filter{
			Name:   helper.String("zone"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		response, err := client.UseBmsClient().DescribeFlavors(request)
		if err != nil {
			return retryError(err, InternalError)
		}
		flavorSet = response.Response.FlavorSet
		return nil
	})
	if err != nil {
		return err
	}

	instanceList := make([]map[string]interface{}, 0, len(flavorSet))
	ids := make([]string, 0, len(flavorSet))

	for _, instance := range flavorSet {
		m := map[string]interface{}{
			"flavor_id":   instance.FlavorId,
			"flavor_name": instance.FlavorName,
			"flavor_type": instance.FlavorType,
			"zone":        instance.Placement.Zone,
			"raid_type":   instance.RaidType,
			//"operating_system": instance.OperatingSystem,
			"cpu":         instance.Cpu,
			"cpu_arch":    instance.CpuArch,
			"memory":      instance.Memory,
			"system_disk": instance.SystemDisk,
			"create_time": instance.CreatedTime,
		}
		instanceList = append(instanceList, m)
		ids = append(ids, *instance.FlavorId)
	}

	_ = d.Set("list", instanceList)
	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), instanceList); err != nil {
			return err
		}
	}

	return nil
}
