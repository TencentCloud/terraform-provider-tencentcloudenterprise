/*
Provides a resource to creating dedicated tunnels instances.

~> **NOTE:** 1. ID of the DC is queried, can only apply for this resource offline.

Example Usage

```hcl

	variable "dc_id" {
	  default = "dc-kax48sg7"
	}

	variable "dcg_id" {
	  default = "dcg-dmbhf7jf"
	}

	variable "vpc_id" {
	  default = "vpc-4h9v4mo3"
	}

	resource "tencentcloudenterprise_dc_dcx" "bgp_main" {
	  bandwidth    = 900
	  dc_id        = var.dc_id
	  dcg_id       = var.dcg_id
	  name         = "bgp_main"
	  network_type = "VPC"
	  route_type   = "BGP"
	  vlan         = 306
	  vpc_id       = var.vpc_id
	}

	resource "tencentcloudenterprise_dc_dcx" "static_main" {
	  bandwidth             = 900
	  dc_id                 = var.dc_id
	  dcg_id                = var.dcg_id
	  name                  = "static_main"
	  network_type          = "VPC"
	  route_type            = "STATIC"
	  vlan                  = 301
	  vpc_id                = var.vpc_id
	  tencentcloudenterprise_address       = "100.93.46.1/30"
	  customer_address      = "100.93.46.2/30"
	  idc_routes = [
      "10.0.0.0/16",
      "10.2.0.0/16"
    ]
	}

```
*/
package tencentcloud

import (
	"context"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dc_dcx", CNDescription{
		TerraformTypeCN: "专线通道",
		AttributesCN: map[string]string{
			"direct_connect_tunnel_name":       "专线通道名称",
			"vpc_id":                           "VPC实例ID",
			"vpc_name":                         "私有网络统一 ID 或者黑石网络统一 ID",
			"route_type":                       "路由类型",
			"dcg_id":                           "专线网关实例ID",
			"bgp_asn":                          "用户侧BGP，Asn",
			"bgp_auth_key":                     "用户侧BGP，AuthKey",
			"vlan":                             "Vlan，范围：11 ~ 4000，需要和用户侧的vlan id一致",
			"tencent_address":                  "腾讯侧互联IP",
			"customer_address":                 "用户侧互联IP",
			"bandwidth":                        "专线带宽",
			"load_mode":                        "通道负载均衡模式：None 非冗余模式LoadBalance：负载均衡MasterSlave：主备",
			"direct_connect_id":                "专线 ID，例如：dc-kd7d06of",
			"direct_connect_name":              "专用通道名称",
			"direct_connect_owner_account":     "物理专线 owner，缺省为当前客户（物理专线 owner）共享专线时这里需要填写共享专线的开发商账号 ID",
			"direct_connect_gateway_id":        "专线网关 ID，例如 dcg-d545ddf",
			"bgp_peer":                         "BgpPeer，用户侧bgp信息，包括Asn和AuthKey",
			"idc_routes":                       "静态路由，用户IDC的网段地址",
			"tencentcloudenterprise_address":                    "云侧互联IP",
			"related_direct_connect_tunnel_id": "关联的冗余通道ID",
			"enable_bfd":                       "是否开启BFD",
			"bfd_interval":                     "BFD协议interval配置",
			"connect_subnet_mask":              "互联地址掩码",
			"network_region":                   "互联地址掩码",
			"ip_type":                          "通道IP协议类型",
			"enable_multicast":					"是否开启组播，仅可在修改时设置",
			"multicast_groups":					"通道组播组地址，当开启组播时可用，仅可在修改时设置",
			"state":                            "专线通道状态",
			"bfd_state":                        "BFD协议状态",
			"created_time":                     "创建时间",
			"net_detect_id":                    "网络探测ID",
			"nat_type":                         "是否为NAT通道",
			"vpc_region":                       "VPC所在地域",
			"direct_connect_gateway_name":      "专线网关名称",
		},
	})
}
func resourceTencentCloudDcxInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to creating dedicated tunnels instances.",
		Create:      resourceTencentCloudDcxInstanceCreate,
		Read:        resourceTencentCloudDcxInstanceRead,
		Update:      resourceTencentCloudDcxInstanceUpdate,
		Delete:      resourceTencentCloudDcxInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"direct_connect_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "ID of the DC to be queried, application deployment offline.",
			},
			"direct_connect_tunnel_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "Name of the dedicated tunnel.",
			},
			"direct_connect_owner_account": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Direct connect owner: The default value is the current customer. " +
					"When sharing the dedicated line (by the physical dedicated line owner), " +
					"the developer account ID of the shared dedicated line must be filled in here.",
			},
			"vpc_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the VPC or BMVPC.",
			},
			"vpc_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the VPC or BMVPC.",
			},
			"route_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedStringValue(DC_ROUTE_TYPES),
				Description:  "Type of the route, and available values include `BGP` and `STATIC`.",
			},
			"direct_connect_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the DC Gateway. Currently only new in the console.",
			},
			"bgp_peer": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp_asn": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "BGP ASN of the user. A required field within BGP.",
						},
						"bgp_auth_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "BGP key of the user.",
						},
					},
				},
			},
			"idc_routes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return helper.HashString(v.(string))
				},
				Description: "Static route, the network segment address of the user's IDC",
			},
			"vlan": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Vlan: Range: 11 ~ 4000; it must be consistent with the VLAN ID on the user side.",
			},
			"tencentcloudenterprise_address": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Interconnect IP of the DC within cloud.",
			},
			"customer_address": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Interconnect IP of the DC within client.",
			},
			"bandwidth": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "Bandwidth of the DC in Mbps.",
			},
			"related_direct_connect_tunnel_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the related redundant DC.",
			},
			"load_mode": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue(DC_LOAD_MODES),
				Description:  "Tunnel Load Balancing Mode: None (Non-redundant Mode), LoadBalance (Load Balancing), MasterSlave (Active-Standby).",
			},
			"enable_bfd": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether enables BFD.",
			},
			"bfd_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "BFD Protocol Interval Configuration.",
			},
			"connect_subnet_mask": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:	 true,
				Description: "Mask of the interconnect address.",
			},
			"network_region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:	 true,
				Description: "Region of the vpc.",
			},
			"ip_type": {
				Type:        schema.TypeString,
				ForceNew:	 true,
				Optional:    true,
				Description: "Protocol type of the tunnel IP.",
			},
			"enable_multicast": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable multicast. Can only be modified after creation via updates.",
			},
			"multicast_groups": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Multicast group addresses supported by the tunnel. Can only be modified after creation via updates. Only valid when enable_multicast is true.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the dedicated tunnel. Possible values: AVAILABLE, APPLYING, ALLOCATING, ALLOCATED, ALTERING, DELETING, DELETED, PENDING, REJECTED.",
			},
			"bfd_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The BFD state of the dedicated tunnel. Possible values: DISABLED, ENABLE, UP, DOWN.",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the direct connect tunnel.",
			},
			"net_detect_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network detection ID.",
			},
			"nat_type": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether it is a NAT tunnel.",
			},
			"vpc_region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPC region where the tunnel is connected.",
			},
			"direct_connect_gateway_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Direct connect gateway name.",
			},
		},
	}
}

func resourceTencentCloudDcxInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_dcx.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		dcId                        string
		dcTunnelName                string
		networkType                 string
		networkRegion               string
		vpcId                       int64
		vpcName                     string
		routeType                   string
		bgpAsn                      int64
		bgpAuthKey                  string
		vlan                        int64
		cloudAddress                string
		customerAddress             string
		bandwidth                   int64
		//routeFilterPrefixes []string
		dcgId                       string
		loadMode                    string
		relatedDirectConnectTunnelId string
		enableBfd                   bool
		bfdInterval                 int64
		connectSubnetMask           uint64
		ipType                      string
		idcRoutes                   string
		ownerAccount                string
	)

	// 获取必需字段 - 使用正确的schema字段名
	if v, ok := d.GetOk("direct_connect_id"); ok {
		dcId = v.(string)
	}
	if v, ok := d.GetOk("direct_connect_tunnel_name"); ok {
		dcTunnelName = v.(string)
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		vpcId = int64(v.(int))
	}
	if v, ok := d.GetOk("vpc_name"); ok {
		vpcName = v.(string)
	}
	if v, ok := d.GetOk("route_type"); ok {
		routeType = strings.ToUpper(v.(string))
	}
	if v, ok := d.GetOk("direct_connect_gateway_id"); ok {
		dcgId = v.(string)
	}
	if v, ok := d.GetOk("load_mode"); ok {
		loadMode = v.(string)
	}
	if v, ok := d.GetOk("related_direct_connect_tunnel_id"); ok {
		relatedDirectConnectTunnelId = v.(string)
	}
	if v, ok := d.GetOk("enable_bfd"); ok {
		enableBfd = v.(bool)
	}
	if v, ok := d.GetOk("connect_subnet_mask"); ok {
		connectSubnetMask = uint64(v.(int))
	}
	if v, ok := d.GetOk("network_region"); ok {
		networkRegion = v.(string)
	} else {
		// 设置默认网络区域
		networkRegion = service.client.Region
	}

	// 获取可选字段
	if v, ok := d.GetOkExists("vlan"); ok {
		vlan = int64(v.(int))
	}
	if v, ok := d.GetOk("tencentcloudenterprise_address"); ok {
		cloudAddress = v.(string)
	}
	if v, ok := d.GetOk("customer_address"); ok {
		customerAddress = v.(string)
	}
	if v, ok := d.GetOk("bandwidth"); ok {
		bandwidth = int64(v.(int))
	}
	if v, ok := d.GetOk("bfd_interval"); ok {
		bfdInterval = int64(v.(int))
	}
	if v, ok := d.GetOk("ip_type"); ok {
		ipType = v.(string)
	}
	if v, ok := d.GetOk("direct_connect_owner_account"); ok {
		ownerAccount = v.(string)
	}

	// 处理BGP参数 - 从bgp_peer嵌套结构中获取
	var bgpPeerExist bool
	if v, ok := d.GetOk("bgp_peer"); ok {
		bgpPeerList := v.([]interface{})
		if len(bgpPeerList) > 0 {
			bgpPeer := bgpPeerList[0].(map[string]interface{})
			if asn, exists := bgpPeer["bgp_asn"]; exists {
				bgpAsn = int64(asn.(int))
			}
			if authKey, exists := bgpPeer["bgp_auth_key"]; exists {
				bgpAuthKey = authKey.(string)
			}
		}
		bgpPeerExist = true
	} else {
		bgpPeerExist = false
	}

	// 处理IDC路由（如果存在）- 将数组转换为换行符分隔的字符串
	if v, ok := d.GetOk("idc_routes"); ok {
		routeSet := v.(*schema.Set)
		var routes []string
		for _, route := range routeSet.List() {
			routes = append(routes, route.(string))
		}
		if len(routes) > 0 {
			idcRoutes = strings.Join(routes, "\n")
		}
	}

	// 设置默认网络类型（如果未指定）
	if networkType == "" {
		networkType = "VPC"
	}

	dcxId, err := service.CreateDirectConnectTunnel(ctx, dcId, dcTunnelName, networkType, networkRegion,
		vpcName, routeType, bgpAuthKey, cloudAddress, customerAddress, dcgId, loadMode, relatedDirectConnectTunnelId,
		ipType, idcRoutes, bgpAsn, vlan, bandwidth, bfdInterval, vpcId, connectSubnetMask, enableBfd, bgpPeerExist, ownerAccount)
	if err != nil {
		return err
	}
	d.SetId(dcxId)

	return resourceTencentCloudDcxInstanceRead(d, meta)
}

func resourceTencentCloudDcxInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_dcx.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		dcxId = d.Id()
	)
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		item, has, e := service.DescribeDirectConnectTunnel(ctx, dcxId)
		if e != nil {
			return retryError(e)
		}

		if has == 0 {
			d.SetId("")
			return nil
		}

		if item.DirectConnectId != nil {
			_ = d.Set("direct_connect_id", *item.DirectConnectId)
		}
		if item.DirectConnectTunnelName != nil {
			_ = d.Set("direct_connect_tunnel_name", *item.DirectConnectTunnelName)
		}
		if item.Bandwidth != nil {
			_ = d.Set("bandwidth", *item.Bandwidth)
		}

		var routeType string
		if item.RouteType != nil {
			routeType = strings.ToUpper(*item.RouteType)
			_ = d.Set("route_type", routeType)
		}

		if item.RouteType != nil && *item.RouteType != DC_ROUTE_TYPE_STATIC && item.BgpPeer != nil {
			bgpPeerList := make([]interface{}, 0, 1)
			bgpPeerMap := make(map[string]interface{})
			if item.BgpPeer.Asn != nil {
				bgpPeerMap["bgp_asn"] = *item.BgpPeer.Asn
			}
			if item.BgpPeer.AuthKey != nil {
				bgpPeerMap["bgp_auth_key"] = *item.BgpPeer.AuthKey
			}
			if len(bgpPeerMap) > 0 {
				bgpPeerList = append(bgpPeerList, bgpPeerMap)
				_ = d.Set("bgp_peer", bgpPeerList)
			}
		}

		// 设置其他字段
		if item.Vlan != nil {
			_ = d.Set("vlan", *item.Vlan)
		}
		if item.CloudAddress != nil {
			_ = d.Set("tencentcloudenterprise_address", *item.CloudAddress)
		}
		if item.CustomerAddress != nil {
			_ = d.Set("customer_address", *item.CustomerAddress)
		}
		if item.DirectConnectGatewayId != nil {
			_ = d.Set("direct_connect_gateway_id", *item.DirectConnectGatewayId)
		}

		// 设置状态字段
		if item.State != nil {
			_ = d.Set("state", *item.State)
		}
		if item.BfdState != nil {
			_ = d.Set("bfd_state", *item.BfdState)
		}

		// 设置其他计算字段
		if item.CreatedTime != nil {
			_ = d.Set("created_time", *item.CreatedTime)
		}
		if item.NetDetectId != nil {
			_ = d.Set("net_detect_id", *item.NetDetectId)
		}
		if item.NatType != nil {
			_ = d.Set("nat_type", *item.NatType)
		}
		if item.VpcRegion != nil {
			_ = d.Set("vpc_region", *item.VpcRegion)
		}
		if item.DirectConnectGatewayName != nil {
			_ = d.Set("direct_connect_gateway_name", *item.DirectConnectGatewayName)
		}

		// 设置IDC路由字段 - 将API返回的字符串数组设置为Set
		if item.IdcRoutes != nil && len(item.IdcRoutes) > 0 {
			var routes []interface{}
			for _, route := range item.IdcRoutes {
				if route != nil {
					routes = append(routes, *route)
				}
			}
			if len(routes) > 0 {
				_ = d.Set("idc_routes", routes)
			}
		}

		// 设置负载均衡模式和关联通道字段
		if item.LoadMode != nil {
			_ = d.Set("load_mode", *item.LoadMode)
		}
		if item.RelatedDirectConnectTunnelId != nil {
			_ = d.Set("related_direct_connect_tunnel_id", *item.RelatedDirectConnectTunnelId)
		}

		// 设置组播相关字段
		if item.EnableMulticast != nil {
			_ = d.Set("enable_multicast", *item.EnableMulticast)
		}
		if item.MulticastGroups != nil {
			_ = d.Set("multicast_groups", *item.MulticastGroups)
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func resourceTencentCloudDcxInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_dcx.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcService{client: meta.(*TencentCloudClient).apiV3Conn}

	dcTunnelId := d.Id()

	// 检查是否需要更新LoadMode或RelatedDirectConnectTunnelId
	if d.HasChange("load_mode") || d.HasChange("related_direct_connect_tunnel_id") {
		loadMode := d.Get("load_mode").(string)
		relatedDirectConnectTunnelId := d.Get("related_direct_connect_tunnel_id").(string)
		
		err := service.UpdateVifAssociated(ctx, dcTunnelId, loadMode, relatedDirectConnectTunnelId)
		if err != nil {
			return err
		}
	}

	// 检查是否有其他字段需要更新
	if d.HasChange("direct_connect_tunnel_name") || d.HasChange("bandwidth") || 
		d.HasChange("bgp_peer") || d.HasChange("enable_bfd") || d.HasChange("bfd_interval") ||
		d.HasChange("idc_routes") || d.HasChange("enable_multicast") || d.HasChange("multicast_groups") {
		
		var (
			dcTunnelName            string
			bandwidth               int64
			bgpAsn                  int64
			bgpAuthKey              string
			enableBfd               bool
			bfdInterval             int64
			enableMulticast         bool
			multicastGroups         string
			oldEnableMulticastRaw   interface{}
			oldMulticastGroupsRaw   interface{}
		)

		if d.HasChange("enable_multicast") {
			oldEnableMulticastRaw, _ = d.GetChange("enable_multicast")
		}
		if d.HasChange("multicast_groups") {
			oldMulticastGroupsRaw, _ = d.GetChange("multicast_groups")
		}

		// 获取名称参数（无条件获取当前值）
		dcTunnelName = d.Get("direct_connect_tunnel_name").(string)

		// 获取带宽参数（无条件获取当前值）
		bandwidth = int64(d.Get("bandwidth").(int))

		// 获取BGP参数（无条件获取当前值，确保API调用时传递正确的当前状态）
		if v, ok := d.GetOk("bgp_peer"); ok {
			bgpPeerList := v.([]interface{})
			if len(bgpPeerList) > 0 {
				bgpPeer := bgpPeerList[0].(map[string]interface{})
				if asn, exists := bgpPeer["bgp_asn"]; exists {
					bgpAsn = int64(asn.(int))
				}
				if authKey, exists := bgpPeer["bgp_auth_key"]; exists {
					bgpAuthKey = authKey.(string)
				}
			}
		}

		// 获取BFD参数（无条件获取当前值，确保API调用时传递正确的当前状态）
		enableBfd = d.Get("enable_bfd").(bool)
		if v, ok := d.GetOk("bfd_interval"); ok {
			bfdInterval = int64(v.(int))
		}

		// 获取IDC路由（如果有变更或需要传递给API）
		var idcRoutes string
		if v, ok := d.GetOk("idc_routes"); ok {
			routeSet := v.(*schema.Set)
			var routes []string
			for _, route := range routeSet.List() {
				routes = append(routes, route.(string))
			}
			if len(routes) > 0 {
				idcRoutes = strings.Join(routes, "\n")
			}
		}

		enableMulticast = d.Get("enable_multicast").(bool)
		if v, ok := d.GetOk("multicast_groups"); ok {
			multicastGroups = v.(string)
		}

		vpcId := d.Get("vpc_name").(string)

		d.Partial(true)
		err := service.ModifyDirectConnectTunnelAttribute(ctx, dcTunnelId, dcTunnelName, bgpAuthKey, idcRoutes,
			bandwidth, bgpAsn, bfdInterval, enableBfd, enableMulticast, multicastGroups, vpcId)
		if err != nil {
			if d.HasChange("enable_multicast") {
				switch v := oldEnableMulticastRaw.(type) {
				case bool:
					_ = d.Set("enable_multicast", v)
				default:
					_ = d.Set("enable_multicast", nil)
				}
			}
			if d.HasChange("multicast_groups") {
				switch v := oldMulticastGroupsRaw.(type) {
				case string:
					_ = d.Set("multicast_groups", v)
				default:
					_ = d.Set("multicast_groups", nil)
				}
			}
			d.Partial(false)
			return err
		}
		d.Partial(false)
	}

	return resourceTencentCloudDcxInstanceRead(d, meta)
}

func resourceTencentCloudDcxInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_dcx.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		dcxId = d.Id()
	)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		e := service.DeleteDirectConnectTunnel(ctx, dcxId)
		if e != nil {
			return retryError(e)
		}
		return nil
	})

	return err
}
