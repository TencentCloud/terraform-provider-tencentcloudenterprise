/*
Use this data source to query detailed information of TCR instances.

Example Usage

```hcl
data "tencentcloudenterprise_tcr_instances" "name" {
  name       = "test"
}
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tcr_instances", CNDescription{
		TerraformTypeCN: "TCR实例",
		DescriptionCN:   "提供TCR实例数据源，用于查询TCR实例的详细信息。",
		AttributesCN: map[string]string{
			"name":               "要查询的TCR实例名称",
			"instance_id":        "要查询的TCR实例ID",
			"result_output_file": "用于保存结果",
			"instance_list":      "TCR实例信息列表",
			"id":                 "实例ID",
			"status":             "实例状态",
			"public_domain":      "实例公网访问域名",
			"internal_end_point": "实例内网访问VPC链接",
			"create_time":        "实例创建时间",
			"instance_type":      "实例类型",
			"instance_charge_type": "实例计费类型",
			"registry_charge_type": "实例计费类型",
		},
	})
}

func dataSourceTencentCloudTCRInstances() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudTCRInstancesRead,
		Description: "Use this data source to query detailed information of TCR instances",

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the TCR instance to query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the TCR instance to query.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			// Computed values
			"instance_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information list of the dedicated TCR instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the TCR instance.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of TCR instance.",
						},

						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the TCR instance.",
						},
						"public_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Public address for access of the TCR instance.",
						},
						"internal_end_point": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Internal address for access of the TCR instance.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance type.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Tags of the TCR instance.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudTCRInstancesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_tcr_instances.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var name, instanceId string
	var filters = make([]*tcr.Filter, 0)
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		filters = append(filters, &tcr.Filter{Name: helper.String("RegistryName"), Values: []*string{&name}})
	}

	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
	}

	if instanceId == "" && name == "" {
		return fmt.Errorf("instance_id or name must be set at least one.")
	}
	tcrService := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}
	var outErr, inErr error
	instances, outErr := tcrService.DescribeTCRInstances(ctx, instanceId, filters)
	if outErr != nil {
		outErr = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			instances, inErr = tcrService.DescribeTCRInstances(ctx, instanceId, filters)
			if inErr != nil {
				return retryError(inErr)
			}
			return nil
		})
	}

	if outErr != nil {
		return outErr
	}

	ids := make([]string, 0, len(instances))
	instanceList := make([]map[string]interface{}, 0, len(instances))
	for _, ins := range instances {
		mapping := map[string]interface{}{
			"id":                 ins.RegistryId,
			"name":               ins.RegistryName,
			"status":             ins.Status,
			"public_domain":      ins.PublicDomain,
			"instance_type":      ins.RegistryType,
			"internal_end_point": ins.InternalEndpoint,
		}
		tags := make(map[string]string, len(ins.TagSpecification.Tags))
		for _, tag := range ins.TagSpecification.Tags {
			tags[*tag.Key] = *tag.Value
		}
		mapping["tags"] = tags
		instanceList = append(instanceList, mapping)
		ids = append(ids, *ins.RegistryId)
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("instance_list", instanceList); e != nil {
		log.Printf("[CRITAL]%s provider set TCR instance list fail, reason:%s\n", logId, e)
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), instanceList); e != nil {
			return e
		}
	}

	return nil

}
