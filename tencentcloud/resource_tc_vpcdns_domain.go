/*
Provide a resource to create a VPCDNS domain.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpcdns_domain" "foo" {
	  domain      = "brucezylin.cc"
	  dns_forward_status = "ENABLED"
	  tags = {
	    "createdBy" = "terraform3"
	  }
	}

```

# Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import cloud_vpcdns_domain.test domain_id
```
*/
package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	sdkError "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"time"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpcdns_domain", CNDescription{
		TerraformTypeCN: "VPCDNS域名",
		DescriptionCN:   "提供VPCDNS域名资源，用于创建和管理VPC内的私有域名。",
		AttributesCN: map[string]string{
			"domain":             "域名",
			"create_time":        "创建时间",
			"tags":               "标签",
			"dns_forward_status": "DNS转发状态",
			"domain_id":          "域名ID",
		},
	})
}
func resourceTencentCloudVpcDnsDomain() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a VPCDNS domain.",
		Create:      resourceTencentCloudVpcDnsDomainCreate,
		Read:        resourceTencentCloudVpcDnsDomainRead,
		Update:      resourceTencentCloudVpcDnsDomainUpdate,
		Delete:      resourceTencentCloudVpcDnsDomainDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain of the VpcDns.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tags of the VPC.",
			},
			"dns_forward_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "DNS forward status, valid values: ENABLED, DISABLED. Default value: DISABLED.",
				Default:      "DISABLED",
				ValidateFunc: validateAllowedStringValue([]string{"ENABLED", "DISABLED"}),
			},

			// Computed values
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of VpcDns Domain.",
			},
			"domain_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "id of vpcdns domain",
			},
		},
	}
}

func resourceTencentCloudVpcDnsDomainCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentclout_vpcdns_domain.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	vpcDnsService := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		domain string
		status string
		tags   map[string]string
	)
	if temp, ok := d.GetOk("domain"); ok {
		domain = temp.(string)
	}
	if temp, ok := d.GetOk("dns_forward_status"); ok {
		status = temp.(string)
	}

	if temp := helper.GetTags(d, "tags"); len(temp) > 0 {
		tags = temp
	}
	err := vpcDnsService.CreateVpcDnsDomain(ctx, domain, tags, status)
	if err != nil {
		return err
	}

	// protected while tag is not ready, default is 1s
	time.Sleep(waitReadTimeout)

	filters := []*vpcdns.DomainListFilters{
		{
			Name:   helper.String("Domain"),
			Values: []*string{&domain},
		},
	}

	dns, err := vpcDnsService.DescribeVpcDnsDomainByFilters(ctx, filters)
	if err != nil {
		return err
	}
	if dns == nil {
		return fmt.Errorf("vpc dns domain %s not found", domain)
	}
	// convert domain.DomainId to string
	d.SetId(strconv.FormatInt(*dns.DomainId, 10))
	return resourceTencentCloudVpcDnsDomainRead(d, meta)
}

func resourceTencentCloudVpcDnsDomainRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentclout_vpcdns_domain.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	id := d.Id()
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		domain, e := service.DescribeVpcDnsDomainById(ctx, id)
		if e != nil {
			return retryError(e)
		}

		if domain == nil {
			return retryError(errors.New("vpc dns domain not found"))
		}

		_ = d.Set("domain", domain.Domain)
		_ = d.Set("create_time", domain.CreatedOn)

		return nil
	})
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("tags"); ok {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}

		region := meta.(*TencentCloudClient).apiV3Conn.Region
		respTags, err := tagService.DescribeResourceTags(ctx, VPCDNS_SERVICE_TYPE, VPCDNS_RESOURCE_TYPE, region, id)
		if err != nil {
			return err
		}

		_ = d.Set("tags", respTags)

	}
	return nil
}

func resourceTencentCloudVpcDnsDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentclout_vpcdns_domain.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()

	vpcDnsService := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	d.Partial(true)

	var (
		status string
	)

	old, now := d.GetChange("dns_forward_status")
	if d.HasChange("dns_forward_status") {
		status = now.(string)
	} else {
		status = old.(string)
	}

	if err := vpcDnsService.ModifyVpcDnsDomain(ctx, id, status); err != nil {
		return err
	}

	if d.HasChange("tags") {
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))

		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}

		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::vpcdns:%s:uin/:%s/%s", region, VPCDNS_RESOURCE_TYPE, id)

		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}

	}

	d.Partial(false)

	return resourceTencentCloudVpcDnsDomainRead(d, meta)
}

func resourceTencentCloudVpcDnsDomainDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentclout_vpcdns_domain.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if err := service.DeleteVpcDnsDomain(ctx, d.Id()); err != nil {
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
