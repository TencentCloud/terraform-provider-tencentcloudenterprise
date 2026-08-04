/*
Provides a resource to create a CLB instance.

Example Usage
INTERNAL CLB

```hcl

	resource "tencentcloudenterprise_clb_instance" "internal_clb" {
	  network_type = "INTERNAL"
	  clb_name     = "myclb"
	  project_id   = 0
	  vpc_id       = "vpc-7007ll7q"
	  subnet_id    = "subnet-12rastkr"

	  tags = {
	    test = "tf"
	  }
	}

```

OPEN CLB
```hcl

	resource "tencentcloudenterprise_clb_instance" "open_clb" {
	  network_type              = "OPEN"
	  clb_name                  = "myclb"
	  project_id                = 0
	  vpc_id                    = "vpc-da7ffa61"
	  security_groups           = ["sg-o0ek7r93"]

	  tags = {
	    test = "tf"
	  }
	}

```

Default enable
```hcl

	resource "tencentcloudenterprise_vpc_subnet" "subnet" {
	  availability_zone = "ap-guangzhou-1"
	  name              = "sdk-feature-test"
	  vpc_id            = tencentcloudenterprise_vpc.foo.id
	  cidr_block        = "10.0.20.0/28"
	  is_multicast      = false
	}

	resource "tencentcloudenterprise_vpc_security_group" "sglab" {
	  name        = "sg_o0ek7r93"
	  description = "favourite sg"
	  project_id  = 0
	}

	resource "tencentcloudenterprise_vpc" "foo" {
	  name         = "for-my-open-clb"
	  cidr_block   = "10.0.0.0/16"

	  tags = {
	    "test" = "mytest"
	  }
	}

	resource "tencentcloudenterprise_clb_instance" "open_clb" {
	  network_type                 = "OPEN"
	  clb_name                     = "my-open-clb"
	  project_id                   = 0
	  vpc_id                       = tencentcloudenterprise_vpc.foo.id
	  security_groups              = [tencentcloudenterprise_vpc_security_group.sglab.id]

	  tags = {
	    test = "open"
	  }
	}

```

CREATE multiple instance
```hcl

	resource "tencentcloudenterprise_clb_instance" "open_clb1" {
	  network_type              = "OPEN"
	  clb_name = "hello"
	  master_zone_id = "ap-guangzhou-3"
	}

```

Import
CLB instance can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_clb_instance.foo lb-7a0t6zqb
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

var clbActionMu = &sync.Mutex{}

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_clb_instance", CNDescription{
		TerraformTypeCN: "CLB实例",
		DescriptionCN:   "提供负载均衡CLB实例资源，用于创建和管理腾讯云负载均衡实例。",
		AttributesCN: map[string]string{
			"network_type":                 "网络类型",
			"clb_name":                     "CLB名称",
			"clb_vips":                     "CLB的VIP",
			"project_id":                   "项目ID",
			"vpc_id":                       "VPC ID",
			"subnet_id":                    "子网 ID",
			"address_ip_version":           "IP版本",
			"internet_charge_type":         "网络计费类型",
			"internet_bandwidth_max_out":   "最大出带宽",
			"security_groups":              "安全组",
			"snat_pro":                     "是否开启SnatPro",
			"snat_ips":                     "Snat IP列表",
			"tags":                         "标签",
			"vip_isp":                      "运营商",
			"master_zone_id":               "主可用区ID",
			"zone_id":                      "可用区ID",
			"slave_zone_id":                "备可用区ID",
			"log_set_id":                   "日志集ID",
			"log_topic_id":                 "日志主题ID",
			"instance_id":                  "CLB实例ID",
			"target_region_info_vpc_id":    "目标VPC ID",
			"target_region_info_region":    "目标地域",
			"eip_address_id":               "EIP的唯一ID，仅IPv4内网CLB",
			"forward":                      "负载均衡实例类型",
			"number":                       "创建负载均衡的个数",
			"anycast_zone":                 "Anycast的发布域",
			"bandwidth_package_id":         "带宽包ID",
			"tgw_group_name":               "TGW独占集群名称",
			"vip":                          "指定VIP申请负载均衡",
			"tgw_set_labels":               "四层集群标签",
			"stgw_set_labels":              "七层集群标签",
			"cluster_ids":                  "本地专用集群ID",
			"ipv6_decouple_vpc":            "是否从VPC侧申请v6地址",
			"domain":                       "负载均衡域名",
			"status":                       "实例状态",
			"create_time":                  "创建时间",
			"status_time":                  "状态转换时间",
			"snat":                         "是否开启snat",
			"isolation":                    "是否被隔离",
			"isolated_time":                "隔离时间",
			"expire_time":                  "过期时间",
			"charge_type":                  "计费类型",
			"address_ipv6":                 "IPv6地址",
			"is_ddos":                      "是否可绑定高防包",
			"config_id":                    "个性化配置ID",
			"load_balancer_pass_to_target": "是否放通LB流量",
			"ipv6_mode":                    "IPv6模式",
			"sla_type":                     "性能保障规格",
			"is_block":                     "VIP是否被封堵",
			"is_block_time":                "封堵时间",
			"local_bgp":                    "是否本地BGP",
			"cluster_tag":                  "7层独占标签",
			"mix_ip_target":                "是否支持混绑IPv4/IPv6",
			"zones":                        "可用区列表",
			"nfv_info":                     "是否为NFV",
			"health_log_set_id":            "健康检查日志集ID",
			"health_log_topic_id":          "健康检查日志主题ID",
		},
	})
}
func resourceTencentCloudClbInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a CLB instance.",
		Create:      resourceTencentCloudClbInstanceCreate,
		Read:        resourceTencentCloudClbInstanceRead,
		Update:      resourceTencentCloudClbInstanceUpdate,
		Delete:      resourceTencentCloudClbInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"network_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue(CLB_NETWORK_TYPE),
				Description:  "Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.",
			},
			"clb_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of the CLB. The name can only contain English letters, numbers, underscore and hyphen '-'.",
			},

			// Computed
			"clb_vips": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The virtual service address table of the CLB.",
			},
			"project_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     0,
				Description: "ID of the project within the CLB instance, `0` - Default Project.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "VPC ID of the CLB.",
			},
			"subnet_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateStringLengthInRange(2, 60),
				Description:  "Subnet ID of the CLB. Effective only for CLB within the VPC.",
			},
			"address_ip_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Ignore case differences: ipv4 == IPV4 == IPv4
					if strings.EqualFold(old, new) {
						return true
					}
					// Suppress diff between ipv6 and IPv6FullChain
					if (strings.EqualFold(old, "ipv6") && strings.EqualFold(new, "IPv6FullChain")) ||
						(strings.EqualFold(old, "IPv6FullChain") && strings.EqualFold(new, "ipv6")) {
						return true
					}
					return false
				},
				ValidateFunc: validateAllowedStringValue(CLB_ADDRESS_IP_VERSION),
				Description: "IP version, only applicable to open CLB. Valid values are `IPV4`, " +
					"`IPV6` and `IPv6FullChain`, Default is IPV4.",
			},
			"internet_charge_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{INTERNET_CHARGE_TYPE_TRAFFIC_POSTPAID_BY_HOUR}),
				Description: "Internet charge type, only applicable to open CLB. " +
					"Valid values are only `TRAFFIC_POSTPAID_BY_HOUR`.",
			},
			"internet_bandwidth_max_out": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateIntegerInRange(1, 2000),
				Description: "Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, " +
					"2000]. Unit is Mbps.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CLB instance ID.",
			},
			"security_groups": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.",
			},
			// "snat_pro": {
			// 	Type:        schema.TypeBool,
			// 	Optional:    true,
			// 	Description: "Indicates whether Binding IPs of other VPCs feature switch.",
			// },
			// "snat_ips": {
			// 	Type:        schema.TypeList,
			// 	Optional:    true,
			// 	Description: "Snat Ip List, required with `snat_pro=true`. NOTE: This argument cannot be read and modified here because dynamic ip is untraceable, please import resource `tencentcloudenterprise_clb_snat_ip` to handle fixed ips.",
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"ip": {
			// 				Type:        schema.TypeString,
			// 				Optional:    true,
			// 				Description: "Snat IP address, If set to empty will auto allocated.",
			// 			},
			// 			"subnet_id": {
			// 				Type:        schema.TypeString,
			// 				Required:    true,
			// 				Description: "Snat subnet ID.",
			// 			},
			// 		},
			// 	},
			// },
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The available tags within this CLB.",
			},
			"vip_isp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network operator, only applicable to open CLB. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).",
			},
			"master_zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Setting master zone id of cross available zone disaster recovery.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Available zone id, only applicable to open CLB.",
			},
			"slave_zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Setting slave zone id of cross available zone disaster recovery. this zone will undertake traffic when the master is down.",
			},
			"log_set_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of log set.",
			},
			"log_topic_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of log topic.",
			},
			"target_region_info_vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "VPC ID of the target region for cross-region CLB.",
			},
			"target_region_info_region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Region of the target region for cross-region CLB.",
			},
			"vip": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ValidateFunc:  validateIp,
				ConflictsWith: []string{"eip_address_id"},
				Description:   "Specified VIP for INTERNAL CLB instance. The IP must be available in the selected subnet.",
			},
			"eip_address_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"vip"},
				Description:   "The unique ID of EIP, e.g. `eip-11112222`. Only available for INTERNAL CLB instance to bind EIP.",
			},
			"tgw_set_labels": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Layer-4 cluster labels used to create an exclusive INTERNAL CLB instance.",
			},
			"stgw_set_labels": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Layer-7 cluster labels used to create an exclusive INTERNAL CLB instance.",
			},
		},
	}
}

func resourceTencentCloudClbInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_instance.create")()

	clbActionMu.Lock()
	defer clbActionMu.Unlock()

	logId := getLogId(contextNil)

	networkType := d.Get("network_type").(string)

	request := clb.NewCreateLoadBalancerRequest()
	request.LoadBalancerType = helper.String(networkType)

	if v, ok := d.GetOk("clb_name"); ok {
		clbName := v.(string)
		flag, e := checkSameName(clbName, meta)
		if e != nil {
			return e
		}
		if flag {
			return fmt.Errorf("[CHECK][CLB instance][Create] check: Same CLB name %s exists!", clbName)
		}
		request.LoadBalancerName = helper.String(clbName)
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
	}
	if v, ok := d.GetOk("project_id"); ok {
		projectId := int64(v.(int))
		request.ProjectId = &projectId
	}

	//vip
	if v, ok := d.GetOk("vip_isp"); ok {
		if networkType == CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance][Create] check: INTERNAL network_type do not support vip ISP setting")
		}
		request.VipIsp = helper.String(v.(string))
	}

	//ip version
	if v, ok := d.GetOk("address_ip_version"); ok {
		//if networkType == CLB_NETWORK_TYPE_INTERNAL {
		//	return fmt.Errorf("[CHECK][CLB instance][Create] check: INTERNAL network_type do not support IP version setting")
		//}
		request.AddressIPVersion = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		if networkType == CLB_NETWORK_TYPE_OPEN &&
			(request.AddressIPVersion != nil && *request.AddressIPVersion == CLB_ADDRESS_IP_VERSION_IPV4) {
			return fmt.Errorf("[CHECK][CLB instance][Create] check: " +
				"OPEN network_type do not support this operation with subnet_id")
		}
		request.SubnetId = helper.String(v.(string))
	}

	v, ok := d.GetOk("internet_charge_type")
	bv, bok := d.GetOk("internet_bandwidth_max_out")

	chargeType := v.(string)

	//internet charge type
	if ok || bok {
		if networkType == CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance][Create] check: INTERNAL network_type do not support internet charge type setting")
		}
		request.InternetAccessible = &clb.InternetAccessible{}
		if ok {
			request.InternetAccessible.InternetChargeType = helper.String(chargeType)
		}
		if bok {
			request.InternetAccessible.InternetMaxBandwidthOut = helper.IntInt64(bv.(int))
		}

	}

	if v, ok := d.GetOk("master_zone_id"); ok {
		// if networkType == CLB_NETWORK_TYPE_INTERNAL {
		// 	return fmt.Errorf("[CHECK][CLB instance][Create] check: " +
		// 		"INTERNAL network_type do not support master zone id setting")
		// }
		request.MasterZoneId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("vip"); ok {
		if networkType != CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance][Create] check: vip only supports INTERNAL network_type")
		}
		request.Vip = helper.String(v.(string))
	}

	if v, ok := d.GetOk("eip_address_id"); ok {
		if networkType != CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance][Create] check: eip_address_id only supports INTERNAL network_type")
		}
		request.EipAddressId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("tgw_set_labels"); ok {
		if networkType != CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance][Create] check: tgw_set_labels only supports INTERNAL network_type")
		}
		request.TgwSetLabels = helper.InterfacesStringsPoint(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("stgw_set_labels"); ok {
		if networkType != CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance][Create] check: stgw_set_labels only supports INTERNAL network_type")
		}
		request.StgwSetLabels = helper.InterfacesStringsPoint(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("zone_id"); ok {
		if networkType == CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance][Create] check: INTERNAL network_type do not support zone id setting")
		}
		request.ZoneId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("slave_zone_id"); ok {
		// if networkType == CLB_NETWORK_TYPE_INTERNAL {
		// 	return fmt.Errorf("[CHECK][CLB instance][Create] check:" +
		// 		" INTERNAL network_type do not support slave zone id setting")
		// }
		request.SlaveZoneId = helper.String(v.(string))
	}

	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		for k, v := range tags {
			tmpKey := k
			tmpValue := v
			request.Tags = append(request.Tags, &clb.TagInfo{
				TagKey:   &tmpKey,
				TagValue: &tmpValue,
			})
		}
	}

	clbId := ""
	var response *clb.CreateLoadBalancerResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().CreateLoadBalancer(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			requestId := *result.Response.RequestId
			retryErr := waitForTaskFinish(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
			if retryErr != nil {
				return retryError(errors.WithStack(retryErr))
			}
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create CLB instance failed, reason:%+v", logId, err)
		return err
	}
	if len(response.Response.LoadBalancerIds) < 1 {
		return fmt.Errorf("[CHECK][CLB instance][Create] check: response error, load balancer id is nil")
	}
	clbId = *response.Response.LoadBalancerIds[0]
	d.SetId(clbId)

	if v, ok := d.GetOk("security_groups"); ok {
		sgRequest := clb.NewSetLoadBalancerSecurityGroupsRequest()
		sgRequest.LoadBalancerId = helper.String(clbId)
		securityGroups := v.([]interface{})
		sgRequest.SecurityGroups = make([]*string, 0, len(securityGroups))
		for i := range securityGroups {
			if securityGroups[i] != nil {
				securityGroup := securityGroups[i].(string)
				if securityGroup != "" {
					sgRequest.SecurityGroups = append(sgRequest.SecurityGroups, &securityGroup)
				}
			}
		}
		if len(sgRequest.SecurityGroups) > 0 {
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				sgResponse, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().SetLoadBalancerSecurityGroups(sgRequest)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
						logId, sgRequest.GetAction(), sgRequest.ToJsonString(), sgResponse.ToJsonString())
					requestID := *sgResponse.Response.RequestId
					retryErr := waitForTaskFinish(requestID, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
					if retryErr != nil {
						return retryError(errors.WithStack(retryErr))
					}
				}
				return nil
			})
			if err != nil {
				log.Printf("[CRITAL]%s create CLB instance security_groups failed, reason:%+v", logId, err)
				return err
			}
		}
	}

	if v, ok := d.GetOk("log_set_id"); ok {
		if u, ok := d.GetOk("log_topic_id"); ok {
			logRequest := clb.NewSetLoadBalancerClsLogRequest()
			logRequest.LoadBalancerId = helper.String(clbId)
			logRequest.LogSetId = helper.String(v.(string))
			logRequest.LogTopicId = helper.String(u.(string))
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				logResponse, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().SetLoadBalancerClsLog(logRequest)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
						logId, logRequest.GetAction(), logRequest.ToJsonString(), logResponse.ToJsonString())
					requestId := *logResponse.Response.RequestId

					retryErr := waitForTaskFinish(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
					if retryErr != nil {
						return retryError(errors.WithStack(retryErr))
					}
				}
				return nil
			})
			if err != nil {
				log.Printf("[CRITAL]%s set CLB instance log failed, reason:%+v", logId, err)
				return err
			}

		} else {
			return fmt.Errorf("log_topic_id and log_set_id must be set together.")
		}
	}

	return resourceTencentCloudClbInstanceRead(d, meta)
}

func resourceTencentCloudClbInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clbId := d.Id()
	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var instance *clb.LoadBalancer
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := clbService.DescribeLoadBalancerById(ctx, clbId)
		if e != nil {
			return retryError(e)
		}
		instance = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CLB instance failed, reason:%+v", logId, err)
		return err
	}

	if instance == nil {
		d.SetId("")
		return nil
	}

	if instance.LoadBalancerType != nil {
		_ = d.Set("network_type", instance.LoadBalancerType)
	}
	if instance.LoadBalancerName != nil {
		_ = d.Set("clb_name", instance.LoadBalancerName)
	}
	if instance.LoadBalancerVips != nil {
		_ = d.Set("clb_vips", helper.StringsInterfaces(instance.LoadBalancerVips))
		if _, ok := d.GetOk("vip"); ok && len(instance.LoadBalancerVips) > 0 {
			_ = d.Set("vip", instance.LoadBalancerVips[0])
		}
	}
	if instance.SubnetId != nil {
		_ = d.Set("subnet_id", instance.SubnetId)
	}
	if instance.VpcId != nil {
		_ = d.Set("vpc_id", instance.VpcId)
	}
	if instance.TargetRegionInfo != nil && instance.TargetRegionInfo.Region != nil {
		_ = d.Set("target_region_info_region", instance.TargetRegionInfo.Region)
	}
	if instance.TargetRegionInfo != nil && instance.TargetRegionInfo.VpcId != nil {
		_ = d.Set("target_region_info_vpc_id", instance.TargetRegionInfo.VpcId)
	}
	if instance.ProjectId != nil {
		_ = d.Set("project_id", instance.ProjectId)
	}
	if instance.SecureGroups != nil {
		_ = d.Set("security_groups", helper.StringsInterfaces(instance.SecureGroups))
	}
	if instance.TgwSetLabels != nil {
		_ = d.Set("tgw_set_labels", helper.StringsInterfaces(instance.TgwSetLabels))
	}
	if instance.StgwSetLabels != nil {
		_ = d.Set("stgw_set_labels", helper.StringsInterfaces(instance.StgwSetLabels))
	}
	if instance.LoadBalancerId != nil {
		_ = d.Set("instance_id", instance.LoadBalancerId)
	}

	if instance.VipIsp != nil {
		_ = d.Set("vip_isp", instance.VipIsp)
	}
	if instance.AddressIPVersion != nil {
		_ = d.Set("address_ip_version", instance.AddressIPVersion)
	}
	if instance.NetworkAttributes != nil {
		_ = d.Set("internet_bandwidth_max_out", instance.NetworkAttributes.InternetMaxBandwidthOut)
		_ = d.Set("internet_charge_type", instance.NetworkAttributes.InternetChargeType)
	}

	//_ = d.Set("master_zone_id", instance.MasterZone.ZoneId)
	//_ = d.Set("zone_id", instance.Zones)
	//_ = d.Set("slave_zone_id", instance.MasterZone)
	if instance.LogSetId != nil {
		_ = d.Set("log_set_id", instance.LogSetId)
	}
	if instance.LogTopicId != nil {
		_ = d.Set("log_topic_id", instance.LogTopicId)
	}
	// if _, ok := d.GetOk("snat_pro"); ok {
	// 	_ = d.Set("snat_pro", instance.SnatPro)
	// }

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := &TagService{client: tcClient}
	tags, err := tagService.DescribeResourceTags(ctx, "clb", "clb", tcClient.Region, d.Id())
	if err != nil {
		return err
	}

	_ = d.Set("tags", tags)

	// Query EIP binding information
	vpcService := VpcService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var eipId *string
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		eips, e := vpcService.DescribeAddressesByFilter(ctx, map[string]string{
			"instance-id": clbId,
		})
		if e != nil {
			return retryError(e)
		}
		if len(eips) > 0 {
			eipId = eips[0].AddressId
		}
		return nil
	})
	if err != nil {
		log.Printf("[WARN]%s read CLB instance EIP binding failed, reason:%+v", logId, err)
	}
	if eipId != nil {
		_ = d.Set("eip_address_id", eipId)
	}

	return nil
}

func resourceTencentCloudClbInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_instance.update")()

	clbActionMu.Lock()
	defer clbActionMu.Unlock()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	d.Partial(true)

	clbId := d.Id()
	request := clb.NewModifyLoadBalancerAttributesRequest()
	request.LoadBalancerId = helper.String(clbId)
	clbName := ""
	targetRegionInfo := clb.TargetRegionInfo{}
	internet := clb.InternetAccessible{}
	changed := false
	// snatPro := d.Get("snat_pro").(bool)

	if d.HasChange("clb_name") {
		changed = true
		clbName = d.Get("clb_name").(string)
		flag, err := checkSameName(clbName, meta)
		if err != nil {
			return err
		}
		if flag {
			return fmt.Errorf("[CHECK][CLB instance][Update] check: Same CLB name %s exists!", clbName)
		}
		request.LoadBalancerName = helper.String(clbName)
	}

	if d.HasChange("target_region_info_region") || d.HasChange("target_region_info_vpc_id") {
		if d.Get("network_type") == CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance %s][Update] check: INTERNAL network_type do not support this operation with target_region_info", clbId)
		}
		changed = true
		region := d.Get("target_region_info_region").(string)
		vpcId := d.Get("target_region_info_vpc_id").(string)
		targetRegionInfo = clb.TargetRegionInfo{
			Region: &region,
			VpcId:  &vpcId,
		}
		request.TargetRegionInfo = &targetRegionInfo
	}

	if d.HasChange("internet_charge_type") || d.HasChange("internet_bandwidth_max_out") {
		if d.Get("network_type") == CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance %s][Update] check: INTERNAL network_type do not support this operation with internet setting", clbId)
		}
		changed = true
		chargeType := d.Get("internet_charge_type").(string)
		bandwidth := d.Get("internet_bandwidth_max_out").(int)
		if chargeType != "" {
			internet.InternetChargeType = &chargeType
		}
		if bandwidth > 0 {
			internet.InternetMaxBandwidthOut = helper.IntInt64(bandwidth)
		}
		request.InternetChargeInfo = &internet
	}

	// if d.HasChange("snat_pro") {
	// 	changed = true
	// 	request.SnatPro = &snatPro
	// }

	// if d.HasChange("snat_ips") {
	// 	return fmt.Errorf("`snat_ips`")
	// }

	if changed {
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			response, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().ModifyLoadBalancerAttributes(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
				requestId := *response.Response.RequestId
				retryErr := waitForTaskFinish(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
				if retryErr != nil {
					return retryError(retryErr)
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update CLB instance failed, reason:%+v", logId, err)
			return err
		}
	}

	if d.HasChange("security_groups") {

		sgRequest := clb.NewSetLoadBalancerSecurityGroupsRequest()
		sgRequest.LoadBalancerId = helper.String(clbId)
		securityGroups := d.Get("security_groups").([]interface{})
		sgRequest.SecurityGroups = make([]*string, 0, len(securityGroups))
		for i := range securityGroups {
			securityGroup := securityGroups[i].(string)
			sgRequest.SecurityGroups = append(sgRequest.SecurityGroups, &securityGroup)
		}
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			sgResponse, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().SetLoadBalancerSecurityGroups(sgRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, sgRequest.GetAction(), sgRequest.ToJsonString(), sgResponse.ToJsonString())
				requestId := *sgResponse.Response.RequestId
				retryErr := waitForTaskFinish(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
				if retryErr != nil {
					return retryError(errors.WithStack(retryErr))
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update CLB instance security_group failed, reason:%+v", logId, err)
			return err
		}

	}

	if d.HasChange("log_set_id") || d.HasChange("log_topic_id") {
		logSetId := d.Get("log_set_id")
		logTopicId := d.Get("log_topic_id")
		logRequest := clb.NewSetLoadBalancerClsLogRequest()
		logRequest.LoadBalancerId = helper.String(clbId)
		logRequest.LogSetId = helper.String(logSetId.(string))
		logRequest.LogTopicId = helper.String(logTopicId.(string))
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			logResponse, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().SetLoadBalancerClsLog(logRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, logRequest.GetAction(), logRequest.ToJsonString(), logResponse.ToJsonString())
				requestId := *logResponse.Response.RequestId

				retryErr := waitForTaskFinish(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
				if retryErr != nil {
					return retryError(errors.WithStack(retryErr))
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s set CLB instance log failed, reason:%+v", logId, err)
			return err
		}
	}

	if d.HasChange("eip_address_id") {
		if d.Get("network_type") != CLB_NETWORK_TYPE_INTERNAL {
			return fmt.Errorf("[CHECK][CLB instance %s][Update] check: eip_address_id only supports INTERNAL network_type", clbId)
		}

		oldEipId, newEipId := d.GetChange("eip_address_id")

		// 如果有旧的 EIP,需要先解绑
		if oldEipId.(string) != "" {
			vpcService := VpcService{
				client: meta.(*TencentCloudClient).apiV3Conn,
			}
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				e := vpcService.UnattachEip(ctx, oldEipId.(string))
				if e != nil {
					return retryError(e, "DesOperation.MutexTaskRunning")
				}
				return nil
			})
			if err != nil {
				log.Printf("[CRITAL]%s unbind EIP from CLB instance failed, reason:%+v", logId, err)
				return err
			}
			log.Printf("[DEBUG]%s unbind EIP %s from CLB instance %s success", logId, oldEipId.(string), clbId)
		}

		// 如果有新的 EIP,需要绑定
		if newEipId.(string) != "" {
			vpcService := VpcService{
				client: meta.(*TencentCloudClient).apiV3Conn,
			}
			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				e := vpcService.AttachEip(ctx, newEipId.(string), clbId)
				if e != nil {
					return retryError(e)
				}
				return nil
			})
			if err != nil {
				log.Printf("[CRITAL]%s bind EIP to CLB instance failed, reason:%+v", logId, err)
				return err
			}
			log.Printf("[DEBUG]%s bind EIP %s to CLB instance %s success", logId, newEipId.(string), clbId)
		}
	}

	if d.HasChange("tags") {

		oldValue, newValue := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldValue.(map[string]interface{}), newValue.(map[string]interface{}))

		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := &TagService{client: tcClient}
		resourceName := BuildTagResourceName("clb", "clb", tcClient.Region, d.Id())
		err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags)
		if err != nil {
			return err
		}

	}
	d.Partial(false)

	return nil
}

func resourceTencentCloudClbInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_instance.delete")()

	clbActionMu.Lock()
	defer clbActionMu.Unlock()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clbId := d.Id()
	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		e := clbService.DeleteLoadBalancerById(ctx, clbId)
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete CLB instance failed, reason:%+v", logId, err)
		return err
	}

	return nil
}

func checkSameName(name string, meta interface{}) (flag bool, errRet error) {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	flag = false
	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	params := make(map[string]interface{})
	params["clb_name"] = name
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		clbs, e := clbService.DescribeLoadBalancerByFilter(ctx, params)
		if e != nil {
			return retryError(e)
		}
		if len(clbs) > 0 {
			//this describe function is a fuzzy query
			// so take a further check
			for _, clb := range clbs {
				if *clb.LoadBalancerName == name {
					flag = true
					return nil
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CLB instance failed, reason:%+v", logId, err)
	}
	errRet = err
	return
}
