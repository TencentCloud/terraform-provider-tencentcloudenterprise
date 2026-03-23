/*
Provides a resource to create an identity center scim credential

Example Usage

```hcl
resource "tencentcloudenterprise_cic_scim_credential" "cic_scim_credential" {
  zone_id = "z-xxxxxx"
}
```

Import

organization cic_scim_credential can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_scim_credential.cic_scim_credential ${zone_id}#${credential_id}
```

 */
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strings"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_scim_credential", CNDescription{
		TerraformTypeCN: "身份中心SCIM凭证",
		DescriptionCN:   "提供身份中心SCIM凭证资源，用于创建和管理SCIM同步凭证。",
		AttributesCN: map[string]string{
			"zone_id":       "空间ID",
			"credential_id": "凭证ID",
			"credential":    "凭证内容",
			"status":        "凭证状态",
			"create_time":   "创建时间",
		},
	})
}

func resourceTencentCloudCicScimCredential() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create an identity center scim credential",
		Create: resourceTencentCloudCicScimCredentialCreate,
		Read:   resourceTencentCloudCicScimCredentialRead,
		Delete: resourceTencentCloudCicScimCredentialDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Space ID. z-prefix starts with 12 random digits/lowercase letters.",
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SCIM key status, Enabled-On, Disabled-Closed.",
			},

			"credential_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SCIM key ID. scimcred-prefix and followed by 12 random digits/lowercase letters.",
			},

			"credential_secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "SCIM key.",
			},

			"credential_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SCIM credential type.",
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SCIM create time.",
			},

			"expire_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SCIM expire time.",
			},
		},
	}
}

func resourceTencentCloudCicScimCredentialCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_scim_credential.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	var (
		zoneId       string
		credentialId string
	)
	var (
		request  = cic.NewCreateSCIMCredentialRequest()
		response = cic.NewCreateSCIMCredentialResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}

	request.ZoneId = helper.String(zoneId)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().CreateSCIMCredential(request)
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
		log.Printf("[CRITAL]%s create identity center scim credential failed, reason:%+v", logId, err)
		return err
	}

	credentialId = *response.Response.CredentialId

	d.SetId(strings.Join([]string{zoneId, credentialId}, FILED_SP))
	if response.Response != nil && response.Response.CredentialSecret != nil {
		_ = d.Set("credential_secret", *response.Response.CredentialSecret)
	}

	return resourceTencentCloudCicScimCredentialRead(d, meta)
}

func resourceTencentCloudCicScimCredentialRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_scim_credential.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	credentialId := idSplit[1]

	_ = d.Set("zone_id", zoneId)

	respData, err := service.DescribeCicScimCredentialById(ctx, zoneId, credentialId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_scim_credential` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if respData.ZoneId != nil {
		_ = d.Set("zone_id", respData.ZoneId)
	}

	if respData.Status != nil {
		_ = d.Set("status", respData.Status)
	}

	if respData.CredentialId != nil {
		_ = d.Set("credential_id", respData.CredentialId)
	}

	if respData.CredentialType != nil {
		_ = d.Set("credential_type", respData.CredentialType)
	}

	if respData.CreateTime != nil {
		_ = d.Set("create_time", respData.CreateTime)
	}

	if respData.ExpireTime != nil {
		_ = d.Set("expire_time", respData.ExpireTime)
	}

	return nil
}

func resourceTencentCloudCicScimCredentialDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_scim_credential.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	credentialId := idSplit[1]

	var (
		request  = cic.NewDeleteSCIMCredentialRequest()
		response = cic.NewDeleteSCIMCredentialResponse()
	)

	request.ZoneId = helper.String(zoneId)

	request.CredentialId = helper.String(credentialId)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().DeleteSCIMCredential(request)
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
		log.Printf("[CRITAL]%s delete identity center scim credential failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	return nil
}
