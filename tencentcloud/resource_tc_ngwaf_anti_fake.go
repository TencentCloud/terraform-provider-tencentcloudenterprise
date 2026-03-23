/*
Provides a resource to create a NGWAF anti fake URL rule.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_anti_fake" "example" {
  domain = "example.com"
  name   = "anti-fake-rule"
  uri    = "/index.html"
  status = 1
}
```

Import

NGWAF anti fake rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_anti_fake.example rule_id#example.com
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strconv"
	"strings"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudNgwafAntiFake() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafAntiFakeCreate,
		Read:   resourceTencentCloudNgwafAntiFakeRead,
		Update: resourceTencentCloudNgwafAntiFakeUpdate,
		Delete: resourceTencentCloudNgwafAntiFakeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Rule Name.",
			},
			"uri": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Uri.",
			},
			"status": {
				Optional:     true,
				Type:         schema.TypeInt,
				Default:      ANTI_FAKE_URL_STATUS_1,
				ValidateFunc: validateAllowedIntValue(ANTI_FAKE_URL_STATUS),
				Description:  "Status. 0: Turn off rules and log switches, 1: Turn on the rule switch and Turn off the log switch; 2: Turn off the rule switch and turn on the log switch; 3: Turn on the rule switch and turn on the log switch.",
			},
			"rule_id": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Rule ID.",
			},
			"protocol": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Protocol.",
			},
		},
	}
}

func resourceTencentCloudNgwafAntiFakeCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_anti_fake.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = ngwaf.NewAddAntiFakeUrlRequest()
		response = ngwaf.NewAddAntiFakeUrlResponse()
		id       string
		domain   string
	)

	if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
		domain = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("uri"); ok {
		request.Uri = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().AddAntiFakeUrl(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create waf antiFake failed, reason:%+v", logId, err)
		return err
	}

	if response == nil || response.Response == nil || response.Response.Id == nil {
		return fmt.Errorf("create waf antiFake failed: invalid response")
	}

	id = *response.Response.Id
	d.SetId(strings.Join([]string{id, domain}, FILED_SP))

	// set status
	if v, ok := d.GetOkExists("status"); ok {
		status := v.(int)
		if status != ANTI_FAKE_URL_STATUS_1 {
			modifyAntiFakeUrlStatusRequest := ngwaf.NewModifyAntiFakeUrlStatusRequest()
			idUInt, _ := strconv.ParseUint(id, 10, 64)
			modifyAntiFakeUrlStatusRequest.Ids = []*uint64{helper.Uint64(idUInt)}
			modifyAntiFakeUrlStatusRequest.Domain = &domain
			modifyAntiFakeUrlStatusRequest.Status = helper.Uint64(uint64(v.(int)))
			err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyAntiFakeUrlStatus(modifyAntiFakeUrlStatusRequest)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
						logId, modifyAntiFakeUrlStatusRequest.GetAction(), modifyAntiFakeUrlStatusRequest.ToJsonString(), result.ToJsonString())
				}
				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s update waf antiFake status failed, reason:%+v", logId, err)
				return err
			}
		}
	}

	return resourceTencentCloudNgwafAntiFakeRead(d, meta)
}

func resourceTencentCloudNgwafAntiFakeRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_anti_fake.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	id := idSplit[0]
	domain := idSplit[1]

	antiFake, err := service.DescribeWafAntiFakeById(ctx, id, domain)
	if err != nil {
		return err
	}

	if antiFake == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `WafAntiFake` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if antiFake.Domain != nil {
		_ = d.Set("domain", antiFake.Domain)
	}

	if antiFake.Name != nil {
		_ = d.Set("name", antiFake.Name)
	}

	if antiFake.Uri != nil {
		_ = d.Set("uri", antiFake.Uri)
	}

	if antiFake.Status != nil {
		_ = d.Set("status", antiFake.Status)
	}

	if antiFake.Id != nil {
		_ = d.Set("rule_id", antiFake.Id)
	}

	if antiFake.Protocol != nil {
		_ = d.Set("protocol", antiFake.Protocol)
	}

	return nil
}

func resourceTencentCloudNgwafAntiFakeUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_anti_fake.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId                          = getLogId(contextNil)
		modifyAntiFakeUrlRequest       = ngwaf.NewModifyAntiFakeUrlRequest()
		modifyAntiFakeUrlStatusRequest = ngwaf.NewModifyAntiFakeUrlStatusRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	id := idSplit[0]
	domain := idSplit[1]

	immutableArgs := []string{"domain"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	idInt, _ := strconv.ParseInt(id, 10, 64)
	modifyAntiFakeUrlRequest.Id = &idInt
	modifyAntiFakeUrlRequest.Domain = &domain

	if v, ok := d.GetOk("name"); ok {
		modifyAntiFakeUrlRequest.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("uri"); ok {
		modifyAntiFakeUrlRequest.Uri = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyAntiFakeUrl(modifyAntiFakeUrlRequest)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, modifyAntiFakeUrlRequest.GetAction(), modifyAntiFakeUrlRequest.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update waf antiFake failed, reason:%+v", logId, err)
		return err
	}

	if d.HasChange("status") {
		if v, ok := d.GetOkExists("status"); ok {
			modifyAntiFakeUrlStatusRequest.Status = helper.Uint64(uint64(v.(int)))
		}

		idUInt, _ := strconv.ParseUint(id, 10, 64)
		modifyAntiFakeUrlStatusRequest.Ids = []*uint64{helper.Uint64(idUInt)}
		modifyAntiFakeUrlStatusRequest.Domain = &domain
		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyAntiFakeUrlStatus(modifyAntiFakeUrlStatusRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, modifyAntiFakeUrlStatusRequest.GetAction(), modifyAntiFakeUrlStatusRequest.ToJsonString(), result.ToJsonString())
			}
			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s update waf antiFake status failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudNgwafAntiFakeRead(d, meta)
}

func resourceTencentCloudNgwafAntiFakeDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_anti_fake.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	id := idSplit[0]
	domain := idSplit[1]

	if err := service.DeleteWafAntiFakeById(ctx, id, domain); err != nil {
		return err
	}

	return nil
}
