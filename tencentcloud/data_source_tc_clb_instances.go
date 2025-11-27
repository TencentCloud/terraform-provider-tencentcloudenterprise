/*
Use this data source to query detailed information of CLB

# Example Usage

```hcl

	data "tencentcloudenterprise_clb_instances" "foo" {
	  clb_id             = "lb-k2zjp9lv"
	  network_type       = "OPEN"
	  clb_name           = "myclb"
	  project_id         = 0
	  result_output_file = "mytestpath"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_clb_instances", CNDescription{
		TerraformTypeCN: "负载均衡实例列表",
		DescriptionCN:   "提供负载均衡实例列表数据源，用于查询CLB实例的详细信息。",
		AttributesCN: map[string]string{
			"clb_id":                     "要查询的CLB的ID",
			"network_type":               "CLB实例的类型，可用值包括“OPEN”和“INTERNAL”",
			"clb_name":                   "要查询的CLB的名称",
			"project_id":                 "CLB的项目ID",
			"result_output_file":         "用于保存结果",
			"master_zone":                "主可用区域id",
			"clb_list":                   "云负载平衡器的列表每个元素都包含以下属性：",
			"clb_vips":                   "CLB的虚拟服务地址表",
			"status":                     "CLB的状态",
			"create_time":                "创建CLB的时间",
			"status_time":                "CLB的最新状态转换时间",
			"vpc_id":                     "专有网络的ID",
			"subnet_id":                  "子网的ID",
			"security_groups":            "安全组的ID集",
			"target_region_info_region":  "后端服务的区域信息附在CLB上",
			"target_region_info_vpc_id":  "后端服务的VpcId信息附在CLB上",
			"tags":                       "此CLB中的可用标记",
			"address_ip_version":         "IP版本，仅适用于开放式CLB有效值为“IPV4”、“IPV6”和“IPv6FullChain”",
			"vip_isp":                    "网络运营商，仅适用于开通CLB有效值为“CMCC”（中国移动）、“CTCC”（电信）、“CUCC”（中国联通）和“BGP”如果指定了此ISP，则网络计费方法只能使用带宽包计费（BANDWITH_package）",
			"internet_charge_type":       "互联网收费类型，仅适用于开放CLB有效值为`TRAFFIC_POSTPAID_BY_HOUR`、`BANDWITH_POSTPAID_BY_HOUR'和`BANDWICH_PACKAGE`",
			"internet_bandwidth_max_out": "最大带宽输出，仅适用于开放CLB有效值范围为[12048]单位为MB",
			"zone_id":                    "可用区域唯一id（数字表示），此字段可能为null，表示无法获得有效值",
			"zone":                       "可用区域唯一id（字符串表示），此字段可能为null，表示无法获得有效值",
			"zone_name":                  "可用区域名称，此字段可能为null，表示无法获得有效值",
			"zone_region":                "此可用区域所属的区域，此字段可能为null，表示无法获取有效值",
			"local_zone":                 "此可用区域是否为本地区域，此字段可能为null，表示无法获取有效值",
		},
	})
}

func dataSourceTencentCloudClbInstances() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides the CLB instances of Cloud.",
		Read:        dataSourceTencentCloudClbInstancesRead,

		Schema: map[string]*schema.Schema{
			"clb_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the CLB to be queried.",
			},
			"network_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(CLB_NETWORK_TYPE),
				Description:  "Type of CLB instance, and available values include `OPEN` and `INTERNAL`.",
			},
			"clb_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the CLB to be queried.",
			},
			"project_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Project ID of the CLB.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"master_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Master available zone id.",
			},
			"clb_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of cloud load balancers. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"clb_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of CLB.",
						},
						"clb_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of CLB.",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Types of CLB.",
						},
						"project_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of the project.",
						},
						"clb_vips": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The virtual service address table of the CLB.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The status of CLB.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of the CLB.",
						},
						"status_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Latest state transition time of CLB.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the VPC.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the subnet.",
						},
						"security_groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "ID set of the security groups.",
						},
						"target_region_info_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region information of backend service are attached the CLB.",
						},
						"target_region_info_vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VpcId information of backend service are attached the CLB.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The available tags within this CLB.",
						},
						"address_ip_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP version, only applicable to open CLB. Valid values are `IPV4`, `IPV6` and `IPv6FullChain`.",
						},
						"vip_isp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network operator, only applicable to open CLB. Valid values are `CMCC`(China Mobile), `CTCC`(Telecom), `CUCC`(China Unicom) and `BGP`. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).",
						},
						"internet_charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Internet charge type, only applicable to open CLB. Valid values are `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.",
						},
						"internet_bandwidth_max_out": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2048]. Unit is MB.",
						},
						"zone_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Available zone unique id(numerical representation), This field maybe null, means cannot get a valid value.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Available zone unique id(string representation), This field maybe null, means cannot get a valid value.",
						},
						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Available zone name, This field maybe null, means cannot get a valid value.",
						},
						"zone_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region that this available zone belong to, This field maybe null, means cannot get a valid value.",
						},
						"local_zone": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this available zone is local zone, This field maybe null, means cannot get a valid value.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudClbInstancesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_clb_instances.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	params := make(map[string]interface{})
	if v, ok := d.GetOk("clb_id"); ok {
		params["clb_id"] = v.(string)
	}
	if v, ok := d.GetOk("clb_name"); ok {
		params["clb_name"] = v.(string)
	}
	if v, ok := d.GetOkExists("project_id"); ok {
		params["project_id"] = v.(int)
	}
	if v, ok := d.GetOk("network_type"); ok {
		params["network_type"] = v.(string)
	}
	if v, ok := d.GetOk("master_zone"); ok {
		params["master_zone"] = v.(string)
	}

	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var clbs []*clb.LoadBalancer
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := clbService.DescribeLoadBalancerByFilter(ctx, params)
		if e != nil {
			return retryError(e)
		}
		clbs = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CLB instances failed, reason:%+v", logId, err)
		return err
	}
	clbList := make([]map[string]interface{}, 0, len(clbs))
	ids := make([]string, 0, len(clbs))
	for _, clbInstance := range clbs {
		mapping := map[string]interface{}{
			"clb_id":                    clbInstance.LoadBalancerId,
			"clb_name":                  clbInstance.LoadBalancerName,
			"network_type":              clbInstance.LoadBalancerType,
			"status":                    clbInstance.Status,
			"create_time":               clbInstance.CreateTime,
			"status_time":               clbInstance.StatusTime,
			"project_id":                clbInstance.ProjectId,
			"vpc_id":                    clbInstance.VpcId,
			"subnet_id":                 clbInstance.SubnetId,
			"clb_vips":                  helper.StringsInterfaces(clbInstance.LoadBalancerVips),
			"target_region_info_region": clbInstance.TargetRegionInfo.Region,
			"target_region_info_vpc_id": clbInstance.TargetRegionInfo.VpcId,
			"address_ip_version":        clbInstance.AddressIPVersion,
			"vip_isp":                   clbInstance.VipIsp,
			"security_groups":           helper.StringsInterfaces(clbInstance.SecureGroups),
		}
		if clbInstance.NetworkAttributes != nil {
			mapping["internet_charge_type"] = *clbInstance.NetworkAttributes.InternetChargeType
			mapping["internet_bandwidth_max_out"] = *clbInstance.NetworkAttributes.InternetMaxBandwidthOut
		}
		if clbInstance.MasterZone != nil {
			mapping["zone_id"] = *clbInstance.MasterZone.ZoneId
			mapping["zone"] = *clbInstance.MasterZone.Zone
			mapping["zone_name"] = *clbInstance.MasterZone.ZoneName
			mapping["zone_region"] = *clbInstance.MasterZone.ZoneRegion
			mapping["local_zone"] = *clbInstance.MasterZone.LocalZone
		}

		if clbInstance.Tags != nil {
			tags := make(map[string]interface{}, len(clbInstance.Tags))
			for _, t := range clbInstance.Tags {
				tags[*t.TagKey] = *t.TagValue
			}
			mapping["tags"] = tags
		}
		clbList = append(clbList, mapping)
		ids = append(ids, *clbInstance.LoadBalancerId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("clb_list", clbList); e != nil {
		log.Printf("[CRITAL]%s provider set CLB list fail, reason:%+v", logId, e)
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), clbList); e != nil {
			return e
		}
	}

	return nil
}
