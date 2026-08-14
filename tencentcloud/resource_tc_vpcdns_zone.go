/*
Provide a resource to create a Private Dns Zone.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpcdns_zone" "foo" {
	  domain = "domain.com"
	  tags {
	    "created_by" : "terraform"
	  }
	  vpc_set {
	    region = "ap-guangzhou"
	    uniq_vpc_id = "vpc-xxxxx"
	  }
	  remark = "test"
	  dns_forward_status = "DISABLED"
	  account_vpc_set {
	    uin = "454xxxxxxx"
	    region = "ap-guangzhou"
	    uniq_vpc_id = "vpc-xxxxx"
	    vpc_name = "test-redis"
	  }
	}

```

# Import

Private Dns Zone can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_zone.foo zone_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
)

func resourceTencentCloudVpcDnsZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudVpcDnsZoneCreate,
		Read:   resourceTencentCloudVpcDnsZoneRead,
		Update: resourceTencentCloudVpcDnsZoneUpdate,
		Delete: resourceTencentCloudVpcDnsZoneDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name, which must be in the format of standard TLD.",
			},
			"tag_set": {
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				Description:   "Tags the private domain when it is created.",
				Deprecated:    "It has been deprecated from version 1.72.4. Use `tags` instead.",
				ConflictsWith: []string{"tags"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Key of Tag.",
						},
						"tag_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Value of Tag.",
						},
					},
				},
			},
			"tags": {
				Type:          schema.TypeMap,
				Optional:      true,
				Description:   "Tags of the private dns zone.",
				ConflictsWith: []string{"tag_set"},
			},
			"vpc_set": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Associates the private domain to a VPC when it is created.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uniq_vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "VPC ID.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "VPC REGION.",
						},
					},
				},
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Remarks.",
			},
			"dns_forward_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "DISABLED",
				ValidateFunc: validateAllowedStringValue(PRIVATE_DNS_FORWARD_STATUS),
				Description:  "Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.",
			},
			"cname_speedup_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "ENABLED",
				ValidateFunc: validateAllowedStringValue([]string{"ENABLED", "DISABLED"}),
				Description:  "Whether to enable CNAME speedup. Valid values: ENABLED, DISABLED. Default value: ENABLED.",
			},
			"account_vpc_set": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "List of authorized accounts' VPCs to associate with the private domain.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uin": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "UIN of the VPC account.",
						},
						"uniq_vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "VPC ID.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Region.",
						},
						"vpc_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "VPC NAME.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudVpcDnsZoneCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	request := vpcdns.NewCreatePrivateZoneRequest()

	domain := d.Get("domain").(string)
	request.Domain = &domain

	if v, ok := d.GetOk("tag_set"); ok {
		tagSet := make([]*vpcdns.TagInfo, 0, 10)
		for _, item := range v.([]interface{}) {
			m := item.(map[string]interface{})
			tagInfo := vpcdns.TagInfo{
				TagKey:   helper.String(m["tag_key"].(string)),
				TagValue: helper.String(m["tag_value"].(string)),
			}
			tagSet = append(tagSet, &tagInfo)
		}
		request.TagSet = tagSet
	}

	if v, ok := d.GetOk("vpc_set"); ok {
		vpcSet := make([]*vpcdns.VpcInfo, 0, 10)
		for _, item := range v.([]interface{}) {
			m := item.(map[string]interface{})
			vpcInfo := vpcdns.VpcInfo{
				UniqVpcId: helper.String(m["uniq_vpc_id"].(string)),
				Region:    helper.String(m["region"].(string)),
			}
			vpcSet = append(vpcSet, &vpcInfo)
		}
		request.VpcSet = vpcSet
	}

	if v, ok := d.GetOk("remark"); ok {
		remark := v.(string)
		request.Remark = helper.String(remark)
	}

	if v, ok := d.GetOk("dns_forward_status"); ok {
		request.DnsForwardStatus = helper.String(v.(string))
	}

	if v, ok := d.GetOk("cname_speedup_status"); ok {
		request.CnameSpeedupStatus = helper.String(v.(string))
	}

	if v, ok := d.GetOk("account_vpc_set"); ok {
		accountVpcSet := make([]*vpcdns.AccountVpcInfo, 0, 10)
		for _, item := range v.([]interface{}) {
			m := item.(map[string]interface{})
			accountVpcInfo := vpcdns.AccountVpcInfo{
				Uin:       helper.String(m["uin"].(string)),
				UniqVpcId: helper.String(m["uniq_vpc_id"].(string)),
				Region:    helper.String(m["region"].(string)),
				VpcName:   helper.String(m["vpc_name"].(string)),
			}
			accountVpcSet = append(accountVpcSet, &accountVpcInfo)
		}
		request.AccountVpcSet = accountVpcSet
	}

	var response *vpcdns.CreatePrivateZoneResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().CreatePrivateZone(request)
		if e != nil {
			return retryError(e, vpcdnsSuffixLockRetryableErrors...)
		}
		if result == nil || result.Response == nil || result.Response.ZoneId == nil {
			return resource.NonRetryableError(fmt.Errorf("create PrivateDns zone failed, Response is nil"))
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create PrivateDns failed, reason:%s\n", logId, err.Error())
		return err
	}

	id := *response.Response.ZoneId
	d.SetId(id)

	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		client := meta.(*TencentCloudClient).apiV3Conn
		service := VpcDnsService{client: client}
		zone, err := service.DescribeVpcDnsZoneById(ctx, id)
		if err != nil {
			return err
		}
		if zone == nil {
			return fmt.Errorf("private zone %s not found after creation", id)
		}

		resourceName, err := buildVpcDnsZoneTagResourceName(zone, client.Region)
		if err != nil {
			return err
		}
		tagService := TagService{client: client}
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudVpcDnsZoneRead(d, meta)
}

func resourceTencentCloudVpcDnsZoneRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()
	client := meta.(*TencentCloudClient).apiV3Conn
	service := VpcDnsService{client: client}
	info, err := service.DescribeVpcDnsZoneById(ctx, id)
	if err != nil {
		return err
	}
	if info == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_vpcdns_zone` [%s] not found, please check if it has been deleted.\n", logId, id)
		d.SetId("")
		return nil
	}
	if info.ZoneId == nil {
		return fmt.Errorf("private zone %s has no ZoneId", id)
	}
	d.SetId(*info.ZoneId)

	_ = d.Set("domain", info.Domain)

	tagSets := make([]map[string]interface{}, 0, len(info.Tags))
	for _, item := range info.Tags {
		tagSets = append(tagSets, map[string]interface{}{
			"tag_key":   item.TagKey,
			"tag_value": item.TagValue,
		})
	}
	_ = d.Set("tag_set", tagSets)

	tagService := TagService{client: client}
	tagResourceID, err := vpcDnsZoneTagResourceID(info)
	if err != nil {
		return err
	}
	tags, err := tagService.DescribeResourceTags(ctx, VPCDNS_SERVICE_TYPE, VPCDNS_RESOURCE_TYPE, client.Region, tagResourceID)
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	vpcSet := make([]map[string]interface{}, 0, len(info.VpcSet))
	for _, item := range info.VpcSet {
		vpcSet = append(vpcSet, map[string]interface{}{
			"uniq_vpc_id": item.UniqVpcId,
			"region":      item.Region,
		})
	}
	_ = d.Set("vpc_set", vpcSet)
	_ = d.Set("remark", info.Remark)
	_ = d.Set("dns_forward_status", info.DnsForwardStatus)
	_ = d.Set("cname_speedup_status", info.CnameSpeedupStatus)

	accountVpcSet := make([]map[string]interface{}, 0, len(info.AccountVpcSet))
	for _, item := range info.AccountVpcSet {
		accountVpcSet = append(accountVpcSet, map[string]interface{}{
			"uin":         item.Uin,
			"uniq_vpc_id": item.UniqVpcId,
			"region":      item.Region,
		})
	}
	_ = d.Set("account_vpc_set", accountVpcSet)
	return nil
}

func resourceTencentCloudVpcDnsZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	id := d.Id()

	if d.HasChange("remark") || d.HasChange("dns_forward_status") || d.HasChange("cname_speedup_status") {
		request := vpcdns.NewModifyPrivateZoneRequest()
		request.ZoneId = helper.String(id)
		if v, ok := d.GetOk("remark"); ok {
			request.Remark = helper.String(v.(string))
		}
		if v, ok := d.GetOk("dns_forward_status"); ok {
			request.DnsForwardStatus = helper.String(v.(string))
		}
		if v, ok := d.GetOk("cname_speedup_status"); ok {
			request.CnameSpeedupStatus = helper.String(v.(string))
		}
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			_, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().ModifyPrivateZone(request)
			if e != nil {
				return retryError(e, vpcdnsSuffixLockRetryableErrors...)
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s modify privateDns zone info failed, reason:%s\n", logId, err.Error())
			return err
		}
	}

	if d.HasChange("vpc_set") || d.HasChange("account_vpc_set") {
		request := vpcdns.NewModifyPrivateZoneVpcRequest()
		request.ZoneId = helper.String(id)
		if v, ok := d.GetOk("vpc_set"); ok {
			var vpcSets = make([]*vpcdns.VpcInfo, 0)
			items := v.([]interface{})
			for _, item := range items {
				value := item.(map[string]interface{})
				vpcInfo := &vpcdns.VpcInfo{
					UniqVpcId: helper.String(value["uniq_vpc_id"].(string)),
					Region:    helper.String(value["region"].(string)),
				}
				vpcSets = append(vpcSets, vpcInfo)
			}
			request.VpcSet = vpcSets
		}

		if v, ok := d.GetOk("account_vpc_set"); ok {
			var accVpcSets = make([]*vpcdns.AccountVpcInfo, 0)
			items := v.([]interface{})
			for _, item := range items {
				value := item.(map[string]interface{})
				accVpcInfo := &vpcdns.AccountVpcInfo{
					UniqVpcId: helper.String(value["uniq_vpc_id"].(string)),
					Region:    helper.String(value["region"].(string)),
					Uin:       helper.String(value["uin"].(string)),
					VpcName:   helper.String(value["vpc_name"].(string)),
				}
				accVpcSets = append(accVpcSets, accVpcInfo)
			}
			request.AccountVpcSet = accVpcSets
		}
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			_, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().ModifyPrivateZoneVpc(request)
			if e = ignoreParseJsonError(e); e != nil {
				return retryError(e)
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s modify privateDns zone vpc failed, reason:%s\n", logId, err.Error())
			return err
		}
	}

	if d.HasChange("tag_set") {
		return fmt.Errorf("tag_set do not support change, please use tags instead.")
	}

	if d.HasChange("tags") {
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))

		client := meta.(*TencentCloudClient).apiV3Conn
		service := VpcDnsService{client: client}
		zone, err := service.DescribeVpcDnsZoneById(ctx, id)
		if err != nil {
			return err
		}
		if zone == nil {
			return fmt.Errorf("private zone %s not found while updating tags", id)
		}
		resourceName, err := buildVpcDnsZoneTagResourceName(zone, client.Region)
		if err != nil {
			return err
		}
		tagService := TagService{client: client}
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}

	}

	return resourceTencentCloudVpcDnsZoneRead(d, meta)
}

func vpcDnsZoneTagResourceID(zone *vpcdns.PrivateZone) (string, error) {
	if zone == nil || zone.DomainId == nil {
		return "", fmt.Errorf("private zone has no DomainId for tags")
	}

	return strconv.FormatInt(*zone.DomainId, 10), nil
}

func buildVpcDnsZoneTagResourceName(zone *vpcdns.PrivateZone, region string) (string, error) {
	resourceID, err := vpcDnsZoneTagResourceID(zone)
	if err != nil {
		return "", err
	}
	if zone.OwnerUin == nil {
		return "", fmt.Errorf("private zone %s has no OwnerUin for tags", resourceID)
	}

	return fmt.Sprintf("qcs::%s:%s:uin/%d:%s/%s", VPCDNS_SERVICE_TYPE, region, *zone.OwnerUin, VPCDNS_RESOURCE_TYPE, resourceID), nil
}

func resourceTencentCloudVpcDnsZoneDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone.delete")()

	logId := getLogId(contextNil)

	request := vpcdns.NewDeletePrivateZoneRequest()
	request.ZoneId = helper.String(d.Id())

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().DeletePrivateZone(request)
		if e != nil {
			return retryError(e, vpcdnsSuffixLockRetryableErrors...)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete privateDns zone failed, reason:%s\n", logId, err.Error())
		return err
	}
	return nil
}
