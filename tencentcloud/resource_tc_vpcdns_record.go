/*
Provide a resource to create a VPCDNS record.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpcdns_record" "foo" {
	  domain_id   = "745"
	  mx          = 0
	  record_type = "A"
	  sub_domain  = "www"
	  value       = "192.168.1.3"
	  weight      = "100"
	}

```

# Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import cloud_vpcdns_record.test record_id
```
*/
package tencentcloud

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	sdkError "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"time"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpcdns_record", CNDescription{
		TerraformTypeCN: "VPCDNS记录",
		DescriptionCN:   "提供VPCDNS记录资源，用于创建和管理VPC内的DNS解析记录。",
		AttributesCN: map[string]string{
			"domain_id":   "域名ID",
			"sub_domain":  "子域名",
			"record_type": "记录类型",
			"value":       "记录值",
			"mx":          "优先级",
			"weight":      "权重",
			"record_id":   "记录ID",
			"ttl":         "生存时间",
			"enabled":     "是否启用",
			"status":      "状态",
			"extra":       "额外信息",
			"created_on":  "创建时间",
			"updated_on":  "更新时间",
		},
	})
}
func resourceTencentCloudVpcDnsRecord() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a VPCDNS record.",
		Create:      resourceTencentCloudVpcDnsRecordCreate,
		Read:        resourceTencentCloudVpcDnsRecordRead,
		Update:      resourceTencentCloudVpcDnsRecordUpdate,
		Delete:      resourceTencentCloudVpcDnsRecordDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Id of the VpcDns.",
			},
			"sub_domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Sub domain.",
			},
			"record_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validateAllowedStringValue([]string{
					"A", "CNAME", "TXT", "MX", "SRV", "AAAA", "SPF",
				}),
				Description: "Record type, valid values: A, CNAME, TXT, MX, SRV, AAAA, SPF. Default value: A.",
			},

			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value of the VpcDns Record.",
			},
			"mx": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The priority of the MX record. Value range: 1-50. Default value: 10.",
			},
			"weight": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The weight of the SRV record. Value range: 1-100. Default value: 10.",
			},

			// Computed values
			"record_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the VpcDns Record.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The TTL of the VpcDns Record.",
			},
			"enabled": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Whether the VpcDns Record is enabled. Valid values: 0 (disabled), 1 (enabled).",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the VpcDns Record. Valid values: 1, 0.",
			},
			"extra": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The extra information of the VpcDns Record.",
			},
			"created_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the VpcDns Record.",
			},
			"updated_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last update time of the VpcDns Record.",
			},
		},
	}
}

func resourceTencentCloudVpcDnsRecordCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentclout_vpcdns_record.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	vpcDnsService := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		domainId   int
		subDomain  string
		recordType string
		value      string
		mx         int
		weight     string
	)
	if temp, ok := d.GetOk("domain_id"); ok {
		domainId = temp.(int)
	}
	if temp, ok := d.GetOk("sub_domain"); ok {
		subDomain = temp.(string)
	}
	if temp, ok := d.GetOk("record_type"); ok {
		recordType = temp.(string)
	}
	if temp, ok := d.GetOk("value"); ok {
		value = temp.(string)
	}
	if temp, ok := d.GetOk("mx"); ok {
		mx = temp.(int)
	}
	if temp, ok := d.GetOk("weight"); ok {
		weight = temp.(string)
	}

	req := vpcdns.NewCreateVpcDnsRecordRequest()
	req.DomainId = helper.IntUint64(domainId)
	req.SubDomain = &subDomain
	req.RecordType = &recordType
	req.Value = &value
	req.Mx = helper.IntUint64(mx)
	req.Weight = &weight

	id, err := vpcDnsService.CreateVpcDnsRecord(ctx, req)
	if err != nil {
		return err
	}
	// protected while tag is not ready, default is 1s
	time.Sleep(waitReadTimeout)

	d.SetId(strconv.FormatInt(id, 10))

	return resourceTencentCloudVpcDnsRecordRead(d, meta)
}

func resourceTencentCloudVpcDnsRecordRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpcdns_record.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		filters := []*vpcdns.RecordListFilters{
			{
				Name: helper.String("SubDomain"),
				Values: []*string{
					helper.String(d.Get("sub_domain").(string)),
				},
			},
		}
		record, e := service.DescribeVpcDnsRecordByFilters(ctx, d.Get("domain_id").(int), filters)
		if e != nil {
			return retryError(e)
		}

		if record == nil {
			return retryError(errors.New("vpc dns record not found"))
		}

		_ = d.Set("record_id", record.RecordId)
		_ = d.Set("ttl", record.Ttl)
		_ = d.Set("enabled", record.Enabled)
		_ = d.Set("status", record.Status)
		_ = d.Set("extra", record.Extra)
		_ = d.Set("created_on", record.CreatedOn)
		_ = d.Set("updated_on", record.CreatedOn)

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func resourceTencentCloudVpcDnsRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpcdns_record.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()

	vpcDnsService := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	d.Partial(true)

	var (
		domainId   int
		subDomain  string
		recordType string
		value      string
		mx         int
		weight     string
	)
	domainId = d.Get("domain_id").(int)

	old, ne := d.GetChange("sub_domain")
	if d.HasChange("sub_domain") {
		subDomain = ne.(string)
	} else {
		subDomain = old.(string)
	}

	old, ne = d.GetChange("record_type")
	if d.HasChange("record_type") {
		recordType = ne.(string)
	} else {
		recordType = old.(string)
	}

	old, ne = d.GetChange("value")
	if d.HasChange("value") {
		value = ne.(string)
	} else {
		value = old.(string)
	}

	old, ne = d.GetChange("mx")
	if d.HasChange("mx") {
		mx = ne.(int)
	} else {
		mx = old.(int)
	}

	old, ne = d.GetChange("weight")
	if d.HasChange("weight") {
		weight = ne.(string)
	} else {
		weight = old.(string)
	}

	req := vpcdns.NewModifyVpcDnsRecordRequest()

	req.DomainId = helper.IntUint64(domainId)
	record, _ := strconv.ParseUint(id, 10, 64)
	req.RecordId = helper.Uint64(record)
	req.SubDomain = &subDomain
	req.RecordType = &recordType
	req.Value = &value
	req.Mx = helper.IntUint64(mx)
	req.Weight = &weight

	if err := vpcDnsService.ModifyVpcDnsRecord(ctx, req); err != nil {
		return err
	}

	d.Partial(false)

	return resourceTencentCloudVpcDnsRecordRead(d, meta)
}

func resourceTencentCloudVpcDnsRecordDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpcdns_Record.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	domainId := d.Get("domain_id").(int)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if err := service.DeleteVpcDnsRecord(ctx, uint64(domainId), d.Id()); err != nil {
			if sdkErr, ok := err.(*sdkError.CloudSDKError); ok {
				if sdkErr.Code == VPCNotFound {
					return nil
				}
			}
			return resource.RetryableError(err)
		}
		return nil
	})

	return err
}
