/*
Provides a resource to create a NGWAF CC session.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_cc_session" "location_example" {
  domain           = "example.com"
  source           = "cookie"
  category         = "location"
  key_or_start_mat = "session_id"
  start_offset     = "3"
  end_offset       = "35"
  edition          = "clb-waf"
  session_name     = "session-from-cookie"
}
```

Import

NGWAF CC session can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_cc_session.location_example example.com#clb-waf#session-xxxxxxxx
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func resourceTencentCloudNgwafCcSession() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafCcSessionCreate,
		Read:   resourceTencentCloudNgwafCcSessionRead,
		Update: resourceTencentCloudNgwafCcSessionUpdate,
		Delete: resourceTencentCloudNgwafCcSessionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},
			"source": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Session matching position, Optional locations are get, post, header, cookie.",
			},
			"category": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Session match pattern, Optional patterns are match, location.",
			},
			"key_or_start_mat": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Session identifier.",
			},
			"end_mat": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Session end identifier, when Category is match.",
			},
			"start_offset": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Starting offset position, when Category is location.",
			},
			"end_offset": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "End offset position, when Category is location.",
			},
			"edition": {
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(EDITION_TYPE),
				Description:  "Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.",
			},
			"session_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Session Name.",
			},
			"session_id": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Session ID.",
			},
		},
	}
}

func resourceTencentCloudNgwafCcSessionCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc_session.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId     = getLogId(contextNil)
		request   = ngwaf.NewUpsertSessionRequest()
		response  = ngwaf.NewUpsertSessionResponse()
		domain    string
		edition   string
		sessionID string
	)

	if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
		domain = v.(string)
	}

	if v, ok := d.GetOk("source"); ok {
		request.Source = helper.String(v.(string))
	}

	if v, ok := d.GetOk("category"); ok {
		request.Category = helper.String(v.(string))
	}

	if v, ok := d.GetOk("key_or_start_mat"); ok {
		request.KeyOrStartMat = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_mat"); ok {
		request.EndMat = helper.String(v.(string))
	}

	if v, ok := d.GetOk("start_offset"); ok {
		request.StartOffset = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_offset"); ok {
		request.EndOffset = helper.String(v.(string))
	}

	if v, ok := d.GetOk("edition"); ok {
		request.Edition = helper.String(v.(string))
		edition = v.(string)
	}

	if v, ok := d.GetOk("session_name"); ok {
		request.SessionName = helper.String(v.(string))
	}

	request.SessionID = helper.IntInt64(-1)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().UpsertSession(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response.SessionID == nil {
			e = fmt.Errorf("waf ccSession not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create waf ccSession failed, reason:%+v", logId, err)
		return err
	}

	sessionIDInt := *response.Response.SessionID
	sessionID = strconv.FormatInt(sessionIDInt, 10)
	d.SetId(strings.Join([]string{domain, edition, sessionID}, FILED_SP))

	return resourceTencentCloudNgwafCcSessionRead(d, meta)
}

func resourceTencentCloudNgwafCcSessionRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc_session.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[0]
	edition := idSplit[1]
	sessionID := idSplit[2]

	ccSession, err := service.DescribeWafCcSessionById(ctx, domain, edition, sessionID)
	if err != nil {
		return err
	}

	if ccSession == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `WafCcSession` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("domain", domain)
	_ = d.Set("edition", edition)

	if ccSession.Source != nil {
		_ = d.Set("source", ccSession.Source)
	}

	if ccSession.Category != nil {
		_ = d.Set("category", ccSession.Category)
	}

	if ccSession.KeyOrStartMat != nil {
		_ = d.Set("key_or_start_mat", ccSession.KeyOrStartMat)
	}

	if ccSession.EndMat != nil {
		_ = d.Set("end_mat", ccSession.EndMat)
	}

	if ccSession.StartOffset != nil {
		_ = d.Set("start_offset", ccSession.StartOffset)
	}

	if ccSession.EndOffset != nil {
		_ = d.Set("end_offset", ccSession.EndOffset)
	}

	if ccSession.SessionName != nil {
		_ = d.Set("session_name", ccSession.SessionName)
	}

	if ccSession.SessionId != nil {
		_ = d.Set("session_id", ccSession.SessionId)
	}

	return nil
}

func resourceTencentCloudNgwafCcSessionUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc_session.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewUpsertSessionRequest()
	)

	immutableArgs := []string{"domain", "edition", "session_name"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[0]
	edition := idSplit[1]
	sessionID := idSplit[2]

	request.Domain = &domain
	request.Edition = &edition
	sessionIDInt, _ := strconv.ParseInt(sessionID, 10, 64)
	request.SessionID = &sessionIDInt

	if v, ok := d.GetOk("source"); ok {
		request.Source = helper.String(v.(string))
	}

	if v, ok := d.GetOk("category"); ok {
		request.Category = helper.String(v.(string))
	}

	if v, ok := d.GetOk("key_or_start_mat"); ok {
		request.KeyOrStartMat = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_mat"); ok {
		request.EndMat = helper.String(v.(string))
	}

	if v, ok := d.GetOk("start_offset"); ok {
		request.StartOffset = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_offset"); ok {
		request.EndOffset = helper.String(v.(string))
	}

	if v, ok := d.GetOk("session_name"); ok {
		request.SessionName = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().UpsertSession(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update waf ccSession failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudNgwafCcSessionRead(d, meta)
}

func resourceTencentCloudNgwafCcSessionDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc_session.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[0]
	edition := idSplit[1]
	sessionID := idSplit[2]

	if err := service.DeleteWafCcSessionById(ctx, domain, edition, sessionID); err != nil {
		return err
	}

	return nil
}
