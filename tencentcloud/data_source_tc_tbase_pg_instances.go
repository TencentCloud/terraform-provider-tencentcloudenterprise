/*
Use this data source to query detailed information of dcdb instances

# Example Usage

```hcl

	data "tencentcloudenterprise_tbase_pg_instances" "foo" {
	  instance_ids = ["tdpg-8xg0uenw", "tdpg-4uq7ufzw"]
	  result_output_file = "foo.json"
	}

```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	tbase "terraform-provider-tencentcloudenterprise/sdk/tbase/v20190107"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tbase_pg_instances", CNDescription{
		TerraformTypeCN: "PGSQL实例列表",
		DescriptionCN:   "提供Tbase PG实例数据源，用于查询Tbase PG实例的详细信息。",
		AttributesCN: map[string]string{
			"instance_ids":       "实例id列表， 用于过滤",
			"result_output_file": "用于保存结果",
			"list":               "实例列表",
			"instance_id":        "实例id",
			"instance_name":      "实例名称",
			"region":             "区域",
			"vpc_id":             "私有网络id",
			"subnet_id":          "子网id",
			"status_desc":        "状态描述",
			"status":             "状态",
			"vip":                "vip",
			"vport":              "端口",
			"create_time":        "创建时间",
			"project_id":         "项目id",
		},
	})
}
func dataSourceTencentCloudTbasePGInstances() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tbase pg instances.",
		ReadContext: dataSourceTencentCloudTbasePGInstancesRead,
		Schema: map[string]*schema.Schema{
			"instance_ids": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Instance ids, if not specified, all instances will be returned.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of tbase pg instances. Each element contains the following attributes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vpc_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_desc": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vport": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudTbasePGInstancesRead(ctx context.Context, d *schema.ResourceData,
	meta interface{}) diag.Diagnostics {
	defer logElapsed("data_source.cloud_dcdb_hour_instances.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	request := tbase.NewDescribePGInstancesRequest()

	if v, ok := d.GetOk("instance_ids"); ok {
		instanceIDsSet := v.(*schema.Set).List()
		ids := make([]*string, 0, len(instanceIDsSet))
		for _, vv := range instanceIDsSet {
			ids = append(ids, helper.String(vv.(string)))
		}
		request.InstanceIds = ids
	}

	service := TbaseService{client: meta.(*TencentCloudClient).apiV3Conn}

	var instances []*tbase.PGInstanceSet
	err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		results, e := service.DescribePGInstances(ctx, request)
		if e != nil {
			return retryError(e)
		}
		instances = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Dcdb instances failed, reason:%+v", logId, err)
		return diag.FromErr(err)
	}

	ids := make([]string, 0, len(instances))
	instanceList := make([]map[string]interface{}, 0, len(instances))
	if instances != nil {
		for _, instance := range instances {
			instanceMap := make(map[string]interface{})

			if instance.InstanceId != nil {
				instanceMap["instance_id"] = *instance.InstanceId
			}
			if instance.InstanceName != nil {
				instanceMap["instance_name"] = *instance.InstanceName
			}
			if instance.Region != nil {
				instanceMap["region"] = *instance.Region
			}
			if instance.VpcId != nil {
				instanceMap["vpc_id"] = *instance.VpcId
			}
			if instance.SubnetId != nil {
				instanceMap["subnet_id"] = *instance.SubnetId
			}
			if instance.StatusDesc != nil {
				instanceMap["status_desc"] = *instance.StatusDesc
			}
			if instance.Status != nil {
				instanceMap["status"] = *instance.Status
			}
			if instance.Vip != nil {
				instanceMap["vip"] = *instance.Vip
			}
			if instance.Vport != nil {
				instanceMap["vport"] = *instance.Vport
			}
			if instance.CreatedAt != nil {
				instanceMap["create_time"] = *instance.CreatedAt
			}
			if instance.ProjectId != nil {
				instanceMap["project_id"] = *instance.ProjectId
			}

			ids = append(ids, *instance.InstanceId)
			instanceList = append(instanceList, instanceMap)
		}
		d.SetId(helper.DataResourceIdsHash(ids))
		_ = d.Set("list", instanceList)
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), instanceList); e != nil {
			return diag.FromErr(e)
		}
	}

	return nil
}
