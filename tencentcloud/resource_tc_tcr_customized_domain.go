/*
Provides a resource to create a tcr customized_domain

Example Usage

```hcl
resource "tencentcloudenterprise_tcr_customized_domain" "my_domain" {
  registry_id = local.tcr_id
  domain_name = "www.test.com"
  certificate_id = "%s"
  tags = {
    "createdBy" = "terraform"
  }
}
```

Import

tcr customized_domain can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_tcr_customized_domain.customized_domain customized_domain_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tcr_customized_domain", CNDescription{
		TerraformTypeCN: "TCR自定义域名",
		DescriptionCN:   "提供TCR自定义域名资源，用于管理容器镜像仓库的自定义访问域名",
		AttributesCN: map[string]string{
			"registry_id":      "实例ID",
			"domain_name":      "自定义域名",
			"certificate_id":   "证书ID",
			"tags":             "标签",
		},
	})
}

func resourceTencentCloudTcrCustomizedDomain() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTcrCustomizedDomainCreate,
		Read:        resourceTencentCloudTcrCustomizedDomainRead,
		Update:      resourceTencentCloudTcrCustomizedDomainUpdate,
		Delete:      resourceTencentCloudTcrCustomizedDomainDelete,
		Description: "Provides a resource to create and manage TCR customized domain",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"registry_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance id.",
			},

			"domain_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Custom domain name.",
			},

			"certificate_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Certificate id.",
			},

			"tags": {
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "Tag description list.",
			},
		},
	}
}

func resourceTencentCloudTcrCustomizedDomainCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_customized_domain.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = tcr.NewCreateInstanceCustomizedDomainRequest()
		registryId string
		domainName string
	)
	if v, ok := d.GetOk("registry_id"); ok {
		registryId = v.(string)
		request.RegistryId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("domain_name"); ok {
		domainName = v.(string)
		request.DomainName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("certificate_id"); ok {
		request.CertificateId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTCRClient().CreateInstanceCustomizedDomain(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tcr CustomizedDomain failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{registryId, domainName}, FILED_SP))

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}

	conf := BuildStateChangeConf([]string{}, []string{"SUCCESS"}, 2*readRetryTimeout, time.Second, service.TcrCustomizedDomainStateRefreshFunc(registryId, domainName, []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::tcr:%s:uin/:instance/%s", region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudTcrCustomizedDomainRead(d, meta)
}

func resourceTencentCloudTcrCustomizedDomainRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_customized_domain.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	registryId := idSplit[0]
	domainName := idSplit[1]

	domains, err := service.DescribeTcrCustomizedDomainById(ctx, registryId, &domainName)
	if err != nil {
		return err
	}

	if len(domains) == 0 {
		d.SetId("")
		log.Printf("[WARN]%s resource `TcrCustomizedDomain` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	CustomizedDomain := domains[0]

	if CustomizedDomain.RegistryId != nil {
		_ = d.Set("registry_id", CustomizedDomain.RegistryId)
	}

	if CustomizedDomain.DomainName != nil {
		_ = d.Set("domain_name", CustomizedDomain.DomainName)
	}

	if CustomizedDomain.CertId != nil {
		_ = d.Set("certificate_id", CustomizedDomain.CertId)
	}

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := &TagService{client: tcClient}
	tags, err := tagService.DescribeResourceTags(ctx, "tcr", "instance", tcClient.Region, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudTcrCustomizedDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_customized_domain.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	immutableArgs := []string{"registry_id", "domain_name", "certificate_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}
	if d.HasChange("tags") {
		ctx := context.WithValue(context.TODO(), logIdKey, logId)
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := &TagService{client: tcClient}
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("tcr", "instance", tcClient.Region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudTcrCustomizedDomainRead(d, meta)
}

func resourceTencentCloudTcrCustomizedDomainDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_customized_domain.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	registryId := idSplit[0]
	domainName := idSplit[1]

	if err := service.DeleteTcrCustomizedDomainById(ctx, registryId, domainName); err != nil {
		return err
	}

	return nil
}
