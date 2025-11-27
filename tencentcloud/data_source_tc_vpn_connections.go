/*
Use this data source to query detailed information of VPN connections.

# Example Usage

```hcl

	data "tencentcloudenterprise_vpn_connections" "foo" {
	  name                = "main"
	  id                  = "vpnx-xfqag"
	  vpn_gateway_id      = "vpngw-8ccsnclt"
	  vpc_id              = "cgw-xfqag"
	  customer_gateway_id = ""
	  tags = {
	    test = "tf"
	  }
	}

```
*/
package tencentcloud

import (
	"context"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpn_connections", CNDescription{
		TerraformTypeCN: "VPN连接",
		AttributesCN: map[string]string{
			"name":                       "VPN连接的名称字符长度限制为1-60",
			"id":                         "VPN连接的ID",
			"vpn_gateway_id":             "VPN连接的VPN网关ID",
			"customer_gateway_id":        "VPN连接的客户网关ID",
			"vpc_id":                     "专有网络的ID",
			"tags":                       "要查询的VPN连接的标签",
			"result_output_file":         "用于保存结果",
			"connection_list":            "专用连接的信息列表",
			"pre_share_key":              "VPN连接的预共享密钥",
			"security_group_policy":      "VPN连接的安全组策略",
			"local_cidr_block":           "局部cidr块",
			"remote_cidr_block":          "远程cidr阻止列表",
			"ike_proto_encry_algorithm":  "IKE操作规范的协议加密算法",
			"ike_proto_authen_algorithm": "IKE操作规范的Proto-Authentication算法",
			"ike_exchange_mode":          "IKE操作规范的交换模式",
			"ike_local_identity":         "IKE操作规范的本地标识",
			"ike_remote_identity":        "IKE操作规范的远程标识",
			"ike_local_address":          "IKE操作规范的本地地址",
			"ike_remote_address":         "IKE操作规范的远程地址",
			"ike_local_fqdn_name":        "IKE操作规范的本地FQDN名称",
			"ike_remote_fqdn_name":       "IKE操作规范的远程FQDN名称",
			"ike_dh_group_name":          "IKE操作规范的DH组名称",
			"ike_sa_lifetime_seconds":    "IKE操作规范的SA生存期，单位为“秒”",
			"ike_version":                "IKE操作规范的版本",
			"ipsec_encrypt_algorithm":    "IPSEC操作规范的加密算法",
			"ipsec_integrity_algorithm":  "IPSEC操作规范的完整性算法",
			"ipsec_sa_lifetime_seconds":  "IPSEC操作规范的SA寿命，单位为“秒”",
			"ipsec_pfs_dh_group":         "IPSEC操作规范的PFS DH组名称",
			"ipsec_sa_lifetime_traffic":  "IPSEC操作规范的SA生存期流量，单位为“KB”",
			"create_time":                "创建VPN连接的时间",
			"vpn_proto":                  "Vpn连接的Vpn proto",
			"encrypt_proto":              "加密VPN连接的原型",
			"route_type":                 "VPN连接的路由类型",
			"state":                      "VPN连接的状态",
			"net_status":                 "VPN连接的网络状态",
		},
	})
}

func dataSourceTencentCloudVpnConnections() *schema.Resource {
	return &schema.Resource{
		Description: "Tencent Cloud VPN Connections",
		Read:        dataSourceTencentCloudVpnConnectionsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "Name of the VPN connection. The length of character is limited to 1-60.",
			},
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the VPN connection.",
			},
			"vpn_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "VPN gateway ID of the VPN connection.",
			},
			"customer_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Customer gateway ID of the VPN connection.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the VPC.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tags of the VPN connection to be queried.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			// Computed values
			"connection_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information list of the dedicated connections.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the VPN connection.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the VPN connection.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the VPC.",
						},
						"customer_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the customer gateway.",
						},
						"vpn_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the VPN gateway.",
						},
						"pre_share_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Pre-shared key of the VPN connection.",
						},
						"security_group_policy": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Security group policy of the VPN connection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_cidr_block": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Local cidr block.",
									},
									"remote_cidr_block": {
										Type:        schema.TypeSet,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Remote cidr block list.",
									},
								},
							},
						},
						"ike_proto_encry_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Proto encrypt algorithm of the IKE operation specification.",
						},
						"ike_proto_authen_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Proto authenticate algorithm of the IKE operation specification.",
						},
						"ike_exchange_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Exchange mode of the IKE operation specification.",
						},
						"ike_local_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Local identity of the IKE operation specification.",
						},
						"ike_remote_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remote identity of the IKE operation specification.",
						},
						"ike_local_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Local address of the IKE operation specification.",
						},
						"ike_remote_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remote address of the IKE operation specification.",
						},
						"ike_local_fqdn_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Local FQDN name of the IKE operation specification.",
						},
						"ike_remote_fqdn_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remote FQDN name of the IKE operation specification.",
						},
						"ike_dh_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DH group name of the IKE operation specification.",
						},
						"ike_sa_lifetime_seconds": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "SA lifetime of the IKE operation specification, unit is `second`.",
						},
						"ike_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version of the IKE operation specification.",
						},
						"ipsec_encrypt_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Encrypt algorithm of the IPSEC operation specification.",
						},
						"ipsec_integrity_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integrity algorithm of the IPSEC operation specification.",
						},
						"ipsec_sa_lifetime_seconds": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "SA lifetime of the IPSEC operation specification, unit is `second`.",
						},
						"ipsec_pfs_dh_group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PFS DH group name of the IPSEC operation specification.",
						},
						"ipsec_sa_lifetime_traffic": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "SA lifetime traffic of the IPSEC operation specification, unit is `KB`.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "A list of tags used to associate different resources.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of the VPN connection.",
						},
						"vpn_proto": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vpn proto of the VPN connection.",
						},
						"encrypt_proto": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Encrypt proto of the VPN connection.",
						},
						"route_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Route type of the VPN connection.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State of the VPN connection.",
						},
						"net_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Net status of the VPN connection.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudVpnConnectionsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpn_connections.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	region := meta.(*TencentCloudClient).apiV3Conn.Region

	request := vpc.NewDescribeVpnConnectionsRequest()

	params := make(map[string]string)
	if v, ok := d.GetOk("id"); ok {
		params["vpn-connection-id"] = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		params["vpn-connection-name"] = v.(string)
	}
	if v, ok := d.GetOk("vpn_gateway_id"); ok {
		params["vpn-gateway-id"] = v.(string)
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		params["vpc-id"] = v.(string)
	}
	if v, ok := d.GetOk("customer_gateway_id"); ok {
		params["customer-gateway-id"] = v.(string)
	}

	tags := helper.GetTags(d, "tags")

	request.Filters = make([]*vpc.Filter, 0, len(params))
	for k, v := range params {
		filter := &vpc.Filter{
			Name:   helper.String(k),
			Values: []*string{helper.String(v)},
		}
		request.Filters = append(request.Filters, filter)
	}
	offset := uint64(0)
	request.Offset = &offset
	result := make([]*vpc.VpnConnection, 0)
	limit := uint64(VPN_DESCRIBE_LIMIT)
	request.Limit = &limit
	for {
		var response *vpc.DescribeVpnConnectionsResponse
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().DescribeVpnConnections(request)
			if e != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, request.GetAction(), request.ToJsonString(), e.Error())
				return retryError(e)
			}
			response = result
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s read VPN connection failed, reason:%s\n", logId, err.Error())
			return err
		} else {
			result = append(result, response.Response.VpnConnectionSet...)
			if len(response.Response.VpnConnectionSet) < VPN_DESCRIBE_LIMIT {
				break
			} else {
				offset = offset + limit
			}
		}
	}
	ids := make([]string, 0, len(result))
	connectionList := make([]map[string]interface{}, 0, len(result))
	for _, connection := range result {
		//tags
		respTags, err := tagService.DescribeResourceTags(ctx, "vpc", "vpnx", region, *connection.VpnConnectionId)
		if err != nil {
			return err
		}
		if len(tags) > 0 {
			if !reflect.DeepEqual(respTags, tags) {
				continue
			}
		}

		mapping := map[string]interface{}{
			"id":                         *connection.VpnConnectionId,
			"name":                       *connection.VpnConnectionName,
			"vpn_gateway_id":             *connection.VpnGatewayId,
			"customer_gateway_id":        *connection.CustomerGatewayId,
			"ike_proto_authen_algorithm": *connection.IKEOptionsSpecification.PropoAuthenAlgorithm,
			"ike_proto_encry_algorithm":  *connection.IKEOptionsSpecification.PropoEncryAlgorithm,
			"ike_exchange_mode":          *connection.IKEOptionsSpecification.ExchangeMode,
			"ike_dh_group_name":          *connection.IKEOptionsSpecification.DhGroupName,
			"ike_sa_lifetime_seconds":    int(*connection.IKEOptionsSpecification.IKESaLifetimeSeconds),
			"ike_version":                *connection.IKEOptionsSpecification.IKEVersion,
			"ike_local_identity":         *connection.IKEOptionsSpecification.LocalIdentity,
			"ike_local_address":          *connection.IKEOptionsSpecification.LocalAddress,
			"ike_local_fqdn_name":        *connection.IKEOptionsSpecification.LocalFqdnName,
			"ike_remote_identity":        *connection.IKEOptionsSpecification.RemoteIdentity,
			"ike_remote_address":         *connection.IKEOptionsSpecification.RemoteAddress,
			"ike_remote_fqdn_name":       *connection.IKEOptionsSpecification.RemoteFqdnName,
			"ipsec_sa_lifetime_seconds":  int(*connection.IPSECOptionsSpecification.IPSECSaLifetimeSeconds),
			"ipsec_encrypt_algorithm":    *connection.IPSECOptionsSpecification.EncryptAlgorithm,
			"ipsec_integrity_algorithm":  *connection.IPSECOptionsSpecification.IntegrityAlgorith,
			"ipsec_pfs_dh_group":         *connection.IPSECOptionsSpecification.PfsDhGroup,
			"ipsec_sa_lifetime_traffic":  int(*connection.IPSECOptionsSpecification.IPSECSaLifetimeTraffic),
			"security_group_policy":      flattenVpnSPDList(connection.SecurityPolicyDatabaseSet),
			"net_status":                 *connection.NetStatus,
			"state":                      *connection.State,
			"create_time":                *connection.CreatedTime,
			"vpn_proto":                  *connection.VpnProto,
			"encrypt_proto":              *connection.EncryptProto,
			"route_type":                 *connection.RouteType,
			"tags":                       respTags,
		}
		connectionList = append(connectionList, mapping)
		ids = append(ids, *connection.VpnConnectionId)
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("connection_list", connectionList); e != nil {
		log.Printf("[CRITAL]%s provider set connection list fail, reason:%s\n", logId, e.Error())
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), connectionList); e != nil {
			return e
		}
	}

	return nil

}
