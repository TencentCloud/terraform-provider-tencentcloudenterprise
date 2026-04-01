/*
Provide a resource to bindvpc for a Private Dns Zone.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpcdns_zone_vpc_attachment" "foo" {
	  zone_id = tencentcloudenterprise_vpcdns_zone.zone.id
	  vpc_set {
	    uniq_vpc_id = "vpc-xxxxx"
	    region      = "ap-guangzhou"
	  }
	}

```

# Import

Private Dns Zone Vpc Attachment can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_zone_vpc_attachment.foo zone_id#vpc_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudVpcDnsZoneVpcAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudVpcDnsZoneVpcAttachmentCreate,
		Read:   resourceTencentCloudVpcDnsZoneVpcAttachmentRead,
		Delete: resourceTencentCloudVpcDnsZoneVpcAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "PrivateZone ID.",
			},
			"vpc_set": {
				Optional:     true,
				ForceNew:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"account_vpc_set"},
				Type:         schema.TypeList,
				Description:  "New add vpc info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uniq_vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Uniq Vpc Id.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Vpc region.",
						},
					},
				},
			},
			"account_vpc_set": {
				Optional:     true,
				ForceNew:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"vpc_set"},
				Type:         schema.TypeList,
				Description:  "New add account vpc info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uniq_vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Uniq Vpc Id.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Vpc region.",
						},
						"uin": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Vpc owner uin. To grant role authorization to this account.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudVpcDnsZoneVpcAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone_vpc_attachment.create")()

	logId := getLogId(contextNil)
	client := meta.(*TencentCloudClient).apiV3Conn

	zoneId := d.Get("zone_id").(string)
	var uniqVpcId string

	// First, read current zone to get existing VPC list
	describeRequest := vpcdns.NewDescribePrivateZoneRequest()
	describeRequest.ZoneId = helper.String(zoneId)

	var describeResponse *vpcdns.DescribePrivateZoneResponse
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := client.UseVpcDnsClient().DescribePrivateZone(describeRequest)
		if e != nil {
			return retryError(e)
		}
		describeResponse = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s describe private dns zone failed, reason:%s\n", logId, err.Error())
		return err
	}

	info := describeResponse.Response.PrivateZone

	// Build the new full VPC list (existing + new)
	request := vpcdns.NewModifyPrivateZoneVpcRequest()
	request.ZoneId = helper.String(zoneId)

	// Copy existing VPCs
	vpcSet := make([]*vpcdns.VpcInfo, 0)
	if info.VpcSet != nil {
		for _, v := range info.VpcSet {
			vpcSet = append(vpcSet, &vpcdns.VpcInfo{
				UniqVpcId: v.UniqVpcId,
				Region:    v.Region,
			})
		}
	}

	// Copy existing account VPCs
	accountVpcSet := make([]*vpcdns.AccountVpcInfo, 0)
	if info.AccountVpcSet != nil {
		for _, v := range info.AccountVpcSet {
			accountVpcSet = append(accountVpcSet, &vpcdns.AccountVpcInfo{
				UniqVpcId: v.UniqVpcId,
				Region:    v.Region,
				Uin:       v.Uin,
			})
		}
	}

	// Add new VPC
	if v, ok := d.GetOk("vpc_set"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			vpcInfo := &vpcdns.VpcInfo{}
			if val, ok := dMap["uniq_vpc_id"]; ok {
				vpcInfo.UniqVpcId = helper.String(val.(string))
				uniqVpcId = val.(string)
			}
			if val, ok := dMap["region"]; ok {
				vpcInfo.Region = helper.String(val.(string))
			}
			vpcSet = append(vpcSet, vpcInfo)
		}
	}

	// Add new account VPC
	if v, ok := d.GetOk("account_vpc_set"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			accVpcInfo := &vpcdns.AccountVpcInfo{}
			if val, ok := dMap["uniq_vpc_id"]; ok {
				accVpcInfo.UniqVpcId = helper.String(val.(string))
				uniqVpcId = val.(string)
			}
			if val, ok := dMap["region"]; ok {
				accVpcInfo.Region = helper.String(val.(string))
			}
			if val, ok := dMap["uin"]; ok {
				accVpcInfo.Uin = helper.String(val.(string))
			}
			accVpcInfo.VpcName = helper.String("")
			accountVpcSet = append(accountVpcSet, accVpcInfo)
		}
	}

	request.VpcSet = vpcSet
	request.AccountVpcSet = accountVpcSet

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := client.UseVpcDnsClient().ModifyPrivateZoneVpc(request)
		if e = ignoreParseJsonError(e); e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create PrivateDns ZoneVpcAttachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	d.SetId(strings.Join([]string{zoneId, uniqVpcId}, FILED_SP))

	return resourceTencentCloudVpcDnsZoneVpcAttachmentRead(d, meta)
}

func resourceTencentCloudVpcDnsZoneVpcAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone_vpc_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	client := meta.(*TencentCloudClient).apiV3Conn

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	zoneId := idSplit[0]
	uniqVpcId := idSplit[1]

	request := vpcdns.NewDescribePrivateZoneRequest()
	request.ZoneId = helper.String(zoneId)

	var response *vpcdns.DescribePrivateZoneResponse
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := client.UseVpcDnsClient().DescribePrivateZone(request)
		if e != nil {
			return retryError(e)
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s describe private dns zone failed, reason:%s\n", logId, err.Error())
		return err
	}

	if response.Response.PrivateZone == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_private_dns_zone_vpc_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	info := response.Response.PrivateZone

	_ = d.Set("zone_id", info.ZoneId)

	if info.VpcSet != nil {
		vpcSetList := []interface{}{}
		for _, vpcSet := range info.VpcSet {
			if *vpcSet.UniqVpcId == uniqVpcId {
				vpcSetMap := map[string]interface{}{
					"uniq_vpc_id": *vpcSet.UniqVpcId,
					"region":      *vpcSet.Region,
				}
				vpcSetList = append(vpcSetList, vpcSetMap)
				break
			}
		}
		_ = d.Set("vpc_set", vpcSetList)
	}

	if info.AccountVpcSet != nil {
		accountVpcSetList := []interface{}{}
		for _, accountVpcSet := range info.AccountVpcSet {
			if *accountVpcSet.UniqVpcId == uniqVpcId {
				accountVpcSetMap := map[string]interface{}{
					"uniq_vpc_id": *accountVpcSet.UniqVpcId,
					"region":      *accountVpcSet.Region,
					"uin":         *accountVpcSet.Uin,
				}
				accountVpcSetList = append(accountVpcSetList, accountVpcSetMap)
				break
			}
		}
		_ = d.Set("account_vpc_set", accountVpcSetList)
	}

	return nil
}

func resourceTencentCloudVpcDnsZoneVpcAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone_vpc_attachment.delete")()

	logId := getLogId(contextNil)
	_ = context.WithValue(context.TODO(), logIdKey, logId)
	client := meta.(*TencentCloudClient).apiV3Conn

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	zoneId := idSplit[0]
	uniqVpcId := idSplit[1]

	// First, read current zone to get existing VPC list
	describeRequest := vpcdns.NewDescribePrivateZoneRequest()
	describeRequest.ZoneId = helper.String(zoneId)

	var describeResponse *vpcdns.DescribePrivateZoneResponse
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := client.UseVpcDnsClient().DescribePrivateZone(describeRequest)
		if e != nil {
			return retryError(e)
		}
		describeResponse = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s describe private dns zone failed, reason:%s\n", logId, err.Error())
		return err
	}

	if describeResponse.Response.PrivateZone == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_private_dns_zone_vpc_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	info := describeResponse.Response.PrivateZone

	// Build new VPC list excluding the target VPC
	request := vpcdns.NewModifyPrivateZoneVpcRequest()
	request.ZoneId = helper.String(zoneId)

	vpcSet := make([]*vpcdns.VpcInfo, 0)
	if info.VpcSet != nil {
		for _, v := range info.VpcSet {
			if *v.UniqVpcId != uniqVpcId {
				vpcSet = append(vpcSet, &vpcdns.VpcInfo{
					UniqVpcId: v.UniqVpcId,
					Region:    v.Region,
				})
			}
		}
	}

	accountVpcSet := make([]*vpcdns.AccountVpcInfo, 0)
	if info.AccountVpcSet != nil {
		for _, v := range info.AccountVpcSet {
			if *v.UniqVpcId != uniqVpcId {
				accountVpcSet = append(accountVpcSet, &vpcdns.AccountVpcInfo{
					UniqVpcId: v.UniqVpcId,
					Region:    v.Region,
					Uin:       v.Uin,
				})
			}
		}
	}

	request.VpcSet = vpcSet
	request.AccountVpcSet = accountVpcSet

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := client.UseVpcDnsClient().ModifyPrivateZoneVpc(request)
		if e = ignoreParseJsonError(e); e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete PrivateDns ZoneVpcAttachment failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}
