/*
Provide a resource to create a Private Dns Record.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpcdns_zone_record" "foo" {
	  zone_id      = "zone-rqndjnki"
	  record_type  = "A"
	  record_value = "192.168.1.2"
	  sub_domain   = "www"
	  ttl          = 300
	  weight       = 1
	  mx           = 0
	}

```

# Import

Private Dns Record can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_zone_record.foo zone_id#record_id
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

func resourceTencentCloudVpcDnsZoneRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudVpcDnsZoneRecordCreate,
		Read:   resourceTencentCloudVpcDnsZoneRecordRead,
		Update: resourceTencentCloudVpcDnsZoneRecordUpdate,
		Delete: resourceTencentCloudVpcDnsZoneRecordDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Private domain ID.",
			},
			"record_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Record type. Valid values: \"A\", \"AAAA\", \"CNAME\", \"MX\", \"TXT\", \"PTR\".",
			},
			"sub_domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subdomain, such as \"www\", \"m\", and \"@\".",
			},
			"record_value": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Record value, such as IP: 192.168.10.2," +
					" CNAME: cname.qcloud.com, and MX: mail.qcloud.com..",
			},
			"weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Record weight. Value range: 1~100.",
			},
			"mx": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "MX priority, which is required when the record type is MX." +
					" Valid values: 5, 10, 15, 20, 30, 40, 50.",
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "Record cache time. The smaller the value, the faster the record will take effect." +
					" Value range: 1~86400s.",
			},
		},
	}
}

func resourceTencentCloudVpcDnsZoneRecordCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone_record.create")()

	logId := getLogId(contextNil)

	request := vpcdns.NewCreatePrivateZoneRecordRequest()

	zoneId := d.Get("zone_id").(string)
	request.ZoneId = &zoneId

	recordType := d.Get("record_type").(string)
	request.RecordType = &recordType

	subDomain := d.Get("sub_domain").(string)
	request.SubDomain = &subDomain

	recordValue := d.Get("record_value").(string)
	request.RecordValue = &recordValue

	if v, ok := d.GetOk("weight"); ok {
		request.Weight = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("mx"); ok {
		request.MX = helper.Int64(int64(v.(int)))
	}
	if v, ok := d.GetOk("ttl"); ok {
		request.TTL = helper.Int64(int64(v.(int)))
	}

	result, err := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().CreatePrivateZoneRecord(request)

	if err != nil {
		log.Printf("[CRITAL]%s create PrivateDns record failed, reason:%s\n", logId, err.Error())
		return err
	}

	response := result

	recordId := *response.Response.RecordId
	d.SetId(strings.Join([]string{zoneId, recordId}, FILED_SP))

	return resourceTencentCloudVpcDnsZoneRecordRead(d, meta)
}

func resourceTencentCloudVpcDnsZoneRecordRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone_record.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("record id strategy is can't read, id is borken, id is %s", d.Id())
	}
	zoneId := idSplit[0]
	recordId := idSplit[1]

	records, err := service.DescribeVpcDnsZoneRecordByFilter(ctx, zoneId, "")
	if err != nil {
		return err
	}

	if len(records) < 1 {
		return fmt.Errorf("private dns record not exists.")
	}

	var record *vpcdns.PrivateZoneRecord
	for _, item := range records {
		if *item.RecordId == recordId {
			record = item
		}
	}
	_ = d.Set("zone_id", record.ZoneId)
	_ = d.Set("record_type", record.RecordType)
	_ = d.Set("sub_domain", record.SubDomain)
	_ = d.Set("record_value", record.RecordValue)
	_ = d.Set("weight", record.Weight)
	_ = d.Set("mx", record.MX)
	_ = d.Set("ttl", record.TTL)

	return nil
}

func resourceTencentCloudVpcDnsZoneRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone_record.update")()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("record id strategy is can't read, id is borken, id is %s", d.Id())
	}
	logId := getLogId(contextNil)
	zoneId := idSplit[0]
	recordId := idSplit[1]

	request := vpcdns.NewModifyPrivateZoneRecordRequest()
	request.ZoneId = helper.String(zoneId)
	request.RecordId = helper.String(recordId)

	needModify := false
	if d.HasChange("record_type") {
		needModify = true
	}

	if d.HasChange("sub_domain") {
		needModify = true
	}

	if d.HasChange("record_value") {
		needModify = true
	}

	if d.HasChange("weight") {
		needModify = true
		if v, ok := d.GetOk("weight"); ok {
			request.Weight = helper.Int64(int64(v.(int)))
		}
	}

	if d.HasChange("mx") {
		needModify = true
		if v, ok := d.GetOk("mx"); ok {
			request.MX = helper.Int64(int64(v.(int)))
		}
	}

	if d.HasChange("ttl") {
		needModify = true
		if v, ok := d.GetOk("ttl"); ok {
			request.TTL = helper.Int64(int64(v.(int)))
		}
	}

	if needModify {
		if v, ok := d.GetOk("record_type"); ok {
			request.RecordType = helper.String(v.(string))
		}
		if v, ok := d.GetOk("sub_domain"); ok {
			request.SubDomain = helper.String(v.(string))
		}
		if v, ok := d.GetOk("record_value"); ok {
			request.RecordValue = helper.String(v.(string))
		}
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			_, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().ModifyPrivateZoneRecord(request)
			if e != nil {
				return retryError(e)
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s modify privateDns record info failed, reason:%s\n", logId, err.Error())
			return err
		}
	}

	return resourceTencentCloudVpcDnsZoneRecordRead(d, meta)
}

func resourceTencentCloudVpcDnsZoneRecordDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_zone_record.delete")()

	logId := getLogId(contextNil)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("record id strategy is can't read, id is borken, id is %s", d.Id())
	}
	zoneId := idSplit[0]
	recordId := idSplit[1]

	// unbind
	request := vpcdns.NewDescribePrivateZoneRequest()
	request.ZoneId = helper.String(zoneId)

	var response *vpcdns.DescribePrivateZoneResponse

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().DescribePrivateZone(request)
		if e != nil {
			return retryError(e)
		}

		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read private dns failed, reason:%s\n", logId, err.Error())
		return err
	}

	info := response.Response.PrivateZone
	oldVpcSet := info.VpcSet
	oldAccVpcSet := info.AccountVpcSet

	unBindRequest := vpcdns.NewModifyPrivateZoneVpcRequest()
	unBindRequest.ZoneId = helper.String(zoneId)
	unBindRequest.VpcSet = []*vpcdns.VpcInfo{}
	unBindRequest.AccountVpcSet = []*vpcdns.AccountVpcInfo{}

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().ModifyPrivateZoneVpc(unBindRequest)
		if e = ignoreParseJsonError(e); e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s unbind privateDns zone vpc failed, reason:%s\n", logId, err.Error())
		return err
	}

	// delete
	recordRequest := vpcdns.NewDeletePrivateZoneRecordRequest()
	recordRequest.ZoneId = helper.String(zoneId)
	recordRequest.RecordId = helper.String(recordId)

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().DeletePrivateZoneRecord(recordRequest)
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete privateDns record failed, reason:%s\n", logId, err.Error())
		return err
	}

	// rebind
	unBindRequest = vpcdns.NewModifyPrivateZoneVpcRequest()
	unBindRequest.ZoneId = helper.String(zoneId)
	unBindRequest.VpcSet = oldVpcSet

	accountVpcSet := make([]*vpcdns.AccountVpcInfo, 0, len(oldAccVpcSet))
	for _, item := range oldAccVpcSet {
		info := vpcdns.AccountVpcInfo{
			Uin:       item.Uin,
			UniqVpcId: item.UniqVpcId,
			Region:    item.Region,
		}
		accountVpcSet = append(accountVpcSet, &info)
	}

	unBindRequest.AccountVpcSet = accountVpcSet

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().ModifyPrivateZoneVpc(unBindRequest)
		if e = ignoreParseJsonError(e); e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s rebind privateDns zone vpc failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}
