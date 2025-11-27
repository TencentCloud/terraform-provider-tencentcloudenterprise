/*
Use this data source to query detailed information of vpc classic_link_instances

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_classic_link_instances" "classic_link_instances" {
	  filters {
	    name   = "vpc-id"
	    values = ["vpc-lh4nqig9"]
	  }
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_classic_link_instances", CNDescription{
		TerraformTypeCN: "私有网络连接实例",
		DescriptionCN:   "提供私有网络连接实例数据源，用于查询VPC经典网络连接实例的详细信息。",
		AttributesCN: map[string]string{
			"filters": "过滤条件`vpc id`-String-（过滤条件）vpc实例id`vm ip`-String-（过滤情况）CVM在基本网络上的ip地址",
			"name":    "属性名称如果存在多个筛选器，则这些筛选器之间的逻辑关系为“与”",
			"属性值如果一个筛选器有多个值，则同一筛选器下这些值之间的逻辑关系为“OR”s": "属性值如果一个筛选器有多个值，则同一筛选器下这些值之间的逻辑关系为“OR”",
			"classic_link_instance_set": "Classiclink实例",
			"vpc_id":                    "VPC实例ID",
			"instance_id":               "CVM实例的唯一ID",
			"result_output_file":        "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcClassicLinkInstances() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc classic_link_instances",
		Read:        dataSourceTencentCloudVpcClassicLinkInstancesRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Filter conditions.`vpc-id` - String - (Filter condition) The VPC instance ID. `vm-ip` - String - (Filter condition) The IP address of the CVM on the basic network.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The attribute name. If more than one Filter exists, the logical relation between these Filters is `AND`.",
						},
						"values": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required:    true,
							Description: "The attribute value. If there are multiple Values for one Filter, the logical relation between these Values under the same Filter is `OR`.",
						},
					},
				},
			},

			"classic_link_instance_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Classiclink instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC instance ID.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of the CVM instance.",
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

func dataSourceTencentCloudVpcClassicLinkInstancesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_classic_link_instances.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("filters"); ok {
		filtersSet := v.([]interface{})
		tmpSet := make([]*vpc.FilterObject, 0, len(filtersSet))

		for _, item := range filtersSet {
			filterObject := vpc.FilterObject{}
			filterObjectMap := item.(map[string]interface{})

			if v, ok := filterObjectMap["name"]; ok {
				filterObject.Name = helper.String(v.(string))
			}
			if v, ok := filterObjectMap["values"]; ok {
				valuesSet := v.(*schema.Set).List()
				filterObject.Values = helper.InterfacesStringsPoint(valuesSet)
			}
			tmpSet = append(tmpSet, &filterObject)
		}
		paramMap["Filters"] = tmpSet
	}

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var classicLinkInstanceSet []*vpc.ClassicLinkInstance

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcClassicLinkInstancesByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		classicLinkInstanceSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(classicLinkInstanceSet))
	tmpList := make([]map[string]interface{}, 0, len(classicLinkInstanceSet))

	if classicLinkInstanceSet != nil {
		for _, classicLinkInstance := range classicLinkInstanceSet {
			classicLinkInstanceMap := map[string]interface{}{}

			if classicLinkInstance.VpcId != nil {
				classicLinkInstanceMap["vpc_id"] = classicLinkInstance.VpcId
			}

			if classicLinkInstance.InstanceId != nil {
				classicLinkInstanceMap["instance_id"] = classicLinkInstance.InstanceId
			}

			ids = append(ids, *classicLinkInstance.InstanceId)
			tmpList = append(tmpList, classicLinkInstanceMap)
		}

		_ = d.Set("classic_link_instance_set", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}
