/*
Use this data source to query detailed information of VPC peering connections.

# Example Usage

```hcl

data "tencentcloudenterprise_vpc_peer_connects" "by_id" {
  peering_connection_id = "pcx-xxxxxxxx"
}

data "tencentcloudenterprise_vpc_peer_connects" "by_vpc" {
  vpc_id = "vpc-xxxxxxxx"
}

data "tencentcloudenterprise_vpc_peer_connects" "by_name" {
  peering_connection_name = "my-peer"
}

data "tencentcloudenterprise_vpc_peer_connects" "by_state" {
  state = "ACTIVE"
}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func tkePeerConnectionInfo() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"peering_connection_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the peering connection.",
		},
		"peering_connection_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the peering connection.",
		},
		"vpc_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "VPC ID of the local end.",
		},
		"vpc_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "VPC name of the local end.",
		},
		"vpc_cidr_block": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "IPv4 CIDR of the local VPC.",
		},
		"peer_vpc_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "VPC ID of the peer end.",
		},
		"peer_vpc_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "VPC name of the peer end.",
		},
		"peer_vpc_cidr_block": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "IPv4 CIDR of the peer VPC.",
		},
		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "State of the peering connection: PENDING / ACTIVE / REJECTED / DELETED / FAILED.",
		},
		"type": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "0: basic network interconnection; 1: VPC-to-VPC; 2: VPC to BM network.",
		},
		"bandwidth": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Bandwidth of the peering connection (Mbps).",
		},
		"src_region": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Region of the local end.",
		},
		"dst_region": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Region of the peer end.",
		},
		"create_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Creation time of the peering connection.",
		},
		"app_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "AppId of the local end.",
		},
		"peer_app_id": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "AppId of the peer end.",
		},
		"uin": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "UIN of the local end.",
		},
		"peer_uin": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "UIN of the peer end.",
		},
		"charge_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Charge type of the peering connection.",
		},
		"tags": {
			Type:        schema.TypeMap,
			Computed:    true,
			Description: "Tags of the peering connection.",
		},
	}
}

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_peer_connects", CNDescription{
		TerraformTypeCN: "vpc对等连接",
		DescriptionCN:   "提供VPC对等连接数据源，用于查询VPC对等连接的详细信息。支持按ID/名称/本端VPC/状态过滤。",
		AttributesCN: map[string]string{
			"peering_connection_id":   "对等连接ID",
			"peering_connection_name": "对等连接名称（过滤条件，模糊匹配）",
			"vpc_id":                  "本端VPC ID（过滤条件）",
			"state":                   "对等连接状态（过滤条件）：PENDING/ACTIVE/REJECTED/DELETED/FAILED",
			"peer_vpc_id":             "对端VPC ID",
			"peer_vpc_name":           "对端VPC名称",
			"peer_vpc_cidr_block":     "对端VPC的IPv4 CIDR",
			"vpc_name":                "本端VPC名称",
			"vpc_cidr_block":          "本端VPC的IPv4 CIDR",
			"type":                    "0:基础网络互通;1:VPC间互通;2:VPC与黑石网络互通",
			"bandwidth":               "对等连接带宽值",
			"src_region":              "本端地域",
			"dst_region":              "对端地域",
			"create_time":             "创建时间",
			"app_id":                  "本端APPID",
			"peer_app_id":             "对端APPID",
			"uin":                     "本端UIN",
			"peer_uin":                "对端UIN",
			"charge_type":             "计费类型",
			"tags":                    "标签",
			"result_output_file":      "用于保存结果",
			"list":                    "对等连接信息列表",
		},
	})
}

func dataSourceTencentCloudVpcPeerConnects() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a data source to query VPC peering connections.",
		Read:        dataSourceTencentCloudVpcPeerConnectsRead,
		Schema: map[string]*schema.Schema{
			"peering_connection_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query by exact peering connection ID. Cannot be used with other filters.",
			},
			"peering_connection_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by peering connection name (fuzzy match via Filters).",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by local VPC ID.",
			},
			"state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by state: PENDING / ACTIVE / EXPIRED / REJECTED.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of peering connections. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: tkePeerConnectionInfo(),
				},
			},
		},
	}
}

func dataSourceTencentCloudVpcPeerConnectsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_vpc_peer_connects.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	var (
		peerId   string
		name     string
		vpcId    string
		state    string
	)

	if v, ok := d.GetOk("peering_connection_id"); ok {
		peerId = v.(string)
	}
	if v, ok := d.GetOk("peering_connection_name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		vpcId = v.(string)
	}
	if v, ok := d.GetOk("state"); ok {
		state = v.(string)
	}

	infos, err := service.DescribeVpcPeerConnects(ctx, peerId, name, vpcId, state)
	if err != nil && peerId == "" {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			infos, err = service.DescribeVpcPeerConnects(ctx, peerId, name, vpcId, state)
			if err != nil {
				return retryError(err)
			}
			return nil
		})
	}
	if err != nil {
		return err
	}

	list := make([]map[string]interface{}, 0, len(infos))
	for _, info := range infos {
		infoMap := map[string]interface{}{}
		infoMap["peering_connection_id"] = info.PeeringConnectionId
		infoMap["peering_connection_name"] = info.PeeringConnectionName
		infoMap["vpc_id"] = info.VpcId
		infoMap["vpc_name"] = info.VpcName
		infoMap["vpc_cidr_block"] = info.VpcCidrBlock
		infoMap["peer_vpc_id"] = info.PeerVpcId
		infoMap["peer_vpc_name"] = info.PeerVpcName
		infoMap["peer_vpc_cidr_block"] = info.PeerVpcCidrBlock
		infoMap["state"] = info.State
		infoMap["type"] = info.Type
		infoMap["bandwidth"] = info.Bandwidth
		infoMap["src_region"] = info.SrcRegion
		infoMap["dst_region"] = info.DstRegion
		infoMap["create_time"] = info.CreateTime
		infoMap["app_id"] = info.AppId
		infoMap["peer_app_id"] = info.PeerAppId
		if info.Uin != nil {
			infoMap["uin"] = helper.UInt64ToStr(*info.Uin)
		}
		if info.PeerUin != nil {
			infoMap["peer_uin"] = helper.UInt64ToStr(*info.PeerUin)
		}
		infoMap["charge_type"] = info.ChargeType

		tags := map[string]string{}
		for _, tag := range info.TagSet {
			if tag != nil && tag.Key != nil && tag.Value != nil {
				tags[*tag.Key] = *tag.Value
			}
		}
		infoMap["tags"] = tags

		list = append(list, infoMap)
	}

	d.SetId("VpcPeerConnects" + peerId + name + vpcId + state)
	err = d.Set("list", list)
	if err != nil {
		log.Printf("[CRITAL]%s provider set tencentcloudenterprise_vpc_peer_connects list fail, reason:%s\n ", logId, err.Error())
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), list); err != nil {
			return err
		}
	}
	return nil
}
