/*
Provides a resource to create a SSL certificate.

# Example Usage

```hcl

	resource "tencentcloudenterprise_clb_certificates" "foo" {
	  alias       = "test-ssl-certificate"
	  cert_type       = "CA"
	  cert       = "-----BEGIN CERTIFICATE-----\nMIIERzCCAq+gAwIBAgIBAjANBgkqhkiG9w0BAQsFADAoMQ0wCwYDVQQDEwR0ZXN0\nMRcwFQYDVQQKEw50ZXJyYWZvcm0gdGVzdDAeFw0xOTA4MTMwMzE5MzlaFw0yOTA4\nMTAwMzE5MzlaMC4xEzARBgNVBAMTCnNlcnZlciBzc2wxFzAVBgNVBAoTDnRlcnJh\nZm9ybS10ZXN0MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA1Ryp+DKK\nSNFKZsPtwfR+jzOnQ8YFieIKYgakV688d8YgpolenbmeEPrzT87tunFD7G9f6ALG\nND8rj7npj0AowxhOL/h/v1D9u0UsIaj5i2GWJrqNAhGLaxWiEB/hy5WOiwxDrGei\ngQqJkFM52Ep7G1Yx7PHJmKFGwN9FhIsFi1cNZfVRopZuCe/RMPNusNVZaIi+qcEf\nfsE1cmfmuSlG3Ap0RKOIyR0ajDEzqZn9/0R7VwWCF97qy8TNYk94K/1tq3zyhVzR\nZ83xOSfrTqEfb3so3AU2jyKgYdwr/FZS72VCHS8IslgnqJW4izIXZqgIKmHaRZtM\nN4jUloi6l/6lktt6Lsgh9xECecxziSJtPMaog88aC8HnMqJJ3kScGCL36GYG+Kaw\n5PnDlWXBaeiDe8z/eWK9+Rr2M+rhTNxosAVGfDJyxAXyiX49LQ0v7f9qzwc/0JiD\nbvsUv1cm6OgpoEMP9SXqqBdwGqeKbD2/2jlP48xlYP6l1SoJG3GgZ8dbAgMBAAGj\ndjB0MAwGA1UdEwEB/wQCMAAwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0PAQH/\nBAUDAweAADAdBgNVHQ4EFgQULwWKBQNLL9s3cb3tTnyPVg+mpCMwHwYDVR0jBBgw\nFoAUKwfrmq791mY831S6UHARHtgYnlgwDQYJKoZIhvcNAQELBQADggGBAMo5RglS\nAHdPgaicWJvmvjjexjF/42b7Rz4pPfMjYw6uYO8He/f4UZWv5CZLrbEe7MywaK3y\n0OsfH8AhyN29pv2x8g9wbmq7omZIOZ0oCAGduEXs/A/qY/hFaCohdkz/IN8qi6JW\nVXreGli3SrpcHFchSwHTyJEXgkutcGAsOvdsOuVSmplOyrkLHc8uUe8SG4j8kGyg\nEzaszFjHkR7g1dVyDVUedc588mjkQxYeAamJgfkgIhljWKMa2XzkVMcVfQHfNpM1\nn+bu8SmqRt9Wma2bMijKRG/Blm756LoI+skY+WRZmlDnq8zj95TT0vceGP0FUWh5\nhKyiocABmpQs9OK9HMi8vgSWISP+fYgkm/bKtKup2NbZBoO5/VL2vCEPInYzUhBO\njCbLMjNjtM5KriCaR7wDARgHiG0gBEPOEW1PIjZ9UOH+LtIxbNZ4eEIIINLHnBHf\nL+doVeZtS/gJc4G4Adr5HYuaS9ZxJ0W2uy0eQlOHzjyxR6Mf/rpnilJlcQ==\n-----END CERTIFICATE-----"
	  cert_key = "XXX"
	}

```

# Import

ssl certificate can be imported using the id, e.g.

```

	$ terraform import cloud_clb_certificates.cert GjTNRoK7

```
*/
package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	sdkError "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_clb_certificates", CNDescription{
		TerraformTypeCN: "SSL证书",
		DescriptionCN:   "提供SSL证书资源，用于创建和管理SSL证书。",
		AttributesCN: map[string]string{
			"alias":           "证书名称",
			"cert_type":       "证书类型",
			"cert":            "证书内容",
			"project_id":      "项目ID",
			"product_zh_name": "证书品牌",
			"domain":          "主域名",
			"status":          "状态",
			"begin_time":      "开始时间",
			"end_time":        "结束时间",
			"create_time":     "创建时间",
			"subject_names":   "证书域名",
			"cert_algorithm":  "证书算法",
			"cert_key":        "证书密钥",
		},
	})
}
func resourceTencentCloudClbCertificate() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a SSL certificate.",
		Create:      resourceTencentCloudClbCertCreate,
		Read:        resourceTencentCloudClbCertRead,
		Update:      resourceTencentCloudClbCertUpdate,
		Delete:      resourceTencentCloudClbCertDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"alias": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Name of the SSL certificate.",
			},
			"cert_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Type of the SSL certificate. Valid values: `CA` and `SVR`.",
			},
			"cert": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Content of the SSL certificate. Not allowed newline at the start and end.",
				ValidateFunc: func(v interface{}, k string) (wss []string, errs []error) {
					value := v.(string)
					if strings.HasPrefix(value, "\n") {
						errs = append(errs, errors.New("cert can't have \\n prefix"))
						return
					}

					if strings.HasSuffix(value, "\n") {
						errs = append(errs, errors.New("cert can't have \\n suffix"))
					}
					return
				},
			},
			"cert_key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.",
				ValidateFunc: func(v interface{}, k string) (wss []string, errs []error) {
					value := v.(string)
					if strings.HasPrefix(value, "\n") {
						errs = append(errs, errors.New("key can't have \\n prefix"))
						return
					}

					if strings.HasSuffix(value, "\n") {
						errs = append(errs, errors.New("key can't have \\n suffix"))
					}
					return
				},
			},

			// computed
			"project_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Project ID of the SSL certificate. Default is `0`.",
			},
			"product_zh_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate authority.",
			},
			"domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Primary domain of the SSL certificate.",
			},
			"status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Status of the SSL certificate.",
			},
			"begin_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Beginning time of the SSL certificate.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Ending time of the SSL certificate.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the SSL certificate.",
			},
			"subject_names": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "ALL domains included in the SSL certificate. Including the primary domain name.",
			},
			"cert_algorithm": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Algorithm of the SSL certificate.",
			},
		},
	}
}

func resourceTencentCloudClbCertCreate(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.cloud_ssl_certificate.create")()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		clbService = ClbService{client: m.(*TencentCloudClient).apiV3Conn}

		outErr, inErr error
		id            string
		//describeResponse *ssl.DescribeCertificateDetailResponse
	)

	request := clb.NewUpLoadCertRequest()
	request.CertContent = helper.String(d.Get("cert").(string))
	request.CertType = helper.String(d.Get("cert_type").(string))
	request.CertAlias = helper.String(d.Get("alias").(string))
	if raw, ok := d.GetOk("cert_key"); ok {
		request.CertKey = helper.String(raw.(string))
	}
	if *request.CertType == "SVR" && (request.CertKey == nil || *request.CertKey == "") {
		return errors.New("when type is SVR, key can't be empty")
	}

	outErr = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		id, inErr = clbService.UploadCertificate(ctx, request)
		if inErr != nil {
			return retryError(inErr)
		}
		return nil
	})
	if outErr != nil {
		log.Printf("[CRITAL]%s create certificate failed, reason: %v", logId, outErr)
		return outErr
	}

	d.SetId(id)

	return resourceTencentCloudClbCertRead(d, m)
}

func resourceTencentCloudClbCertRead(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.cloud_ssl_certificate.read")()
	defer inconsistentCheck(d, m)()

	var (
		logId         = getLogId(contextNil)
		ctx           = context.WithValue(context.TODO(), logIdKey, logId)
		clbService    = ClbService{client: m.(*TencentCloudClient).apiV3Conn}
		outErr, inErr error
		id            = d.Id()
		certs         []*clb.CertList
	)

	describeRequest := clb.NewDescribeCertsRequest()
	describeRequest.CertIdList = []*string{&id}
	describeRequest.WithCert = helper.Bool(true)
	outErr = resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		certs, inErr = clbService.DescribeCert(ctx, describeRequest)
		if inErr != nil {
			return retryError(inErr)
		}
		return nil
	})
	if outErr != nil {
		log.Printf("[CRITAL]%s read certificate failed, reason: %v", logId, outErr)
		d.SetId("")
		return outErr
	}

	if len(certs) < 1 {
		d.SetId("")
		return errors.New("certificate not found")
	}

	certificate := certs[0]
	if nilNames := CheckNil(certificate, map[string]string{
		"Alias":         "alias",
		"CertType":      "cert_type",
		"ProductZhName": "product zh name",
		"Domain":        "domain",
		"Status":        "status",
		"CertBeginTime": "begin time",
		"CertEndTime":   "end time",
		"InsertTime":    "create time",
		"Cert":          "cert",
		"CertAlgorithm": "cert algorithm",
	}); len(nilNames) > 0 {
		return fmt.Errorf("certificate %v are nil", nilNames)
	}

	_ = d.Set("alias", certificate.Alias)
	_ = d.Set("cert_type", certificate.CertType)
	projectId, err := strconv.Atoi(*certificate.ProjectId)
	if err != nil {
		return err
	}
	_ = d.Set("project_id", projectId)
	_ = d.Set("cert", strings.TrimRight(*certificate.Cert, "\n"))
	_ = d.Set("product_zh_name", certificate.ProductZhName)
	_ = d.Set("domain", certificate.Domain)
	_ = d.Set("status", certificate.Status)
	_ = d.Set("begin_time", certificate.CertBeginTime)
	_ = d.Set("end_time", certificate.CertEndTime)
	_ = d.Set("create_time", certificate.InsertTime)
	_ = d.Set("cert_algorithm", certificate.CertAlgorithm)

	subjectAltNames := make([]string, 0, len(certificate.SubjectAltName))
	for _, subjectAltName := range certificate.SubjectAltName {
		subjectAltNames = append(subjectAltNames, *subjectAltName)
	}
	_ = d.Set("subject_names", subjectAltNames)
	return nil
}

func resourceTencentCloudClbCertUpdate(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.cloud_ssl_certificate.update")()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		id         = d.Id()
		clbService = ClbService{client: m.(*TencentCloudClient).apiV3Conn}
		alias      = d.Get("alias").(string)
		cert       = d.Get("cert").(string)
		cert_key   = d.Get("cert_key").(string)
	)

	d.Partial(true)
	request := clb.NewReplaceCertRequest()
	request.OldCertId = helper.String(id)
	request.NewCertKey = helper.String(cert_key)
	request.NewCertContent = helper.String(cert)
	request.NewCertAlias = helper.String(alias)

	if outErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if _, inErr := clbService.ReplaceCert(ctx, request); inErr != nil {
			if sdkErr, ok := inErr.(*sdkError.CloudSDKError); ok {
				code := sdkErr.GetCode()
				if code == InvalidParam || code == CertificateNotFound {
					return resource.NonRetryableError(sdkErr)
				}
			}
			return retryError(inErr)
		}
		return nil
	}); outErr != nil {
		return outErr
	}

	d.Partial(false)
	return resourceTencentCloudClbCertRead(d, m)
}

func resourceTencentCloudClbCertDelete(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.cloud_ssl_certificate.delete")()
	var (
		logId         = getLogId(contextNil)
		ctx           = context.WithValue(context.TODO(), logIdKey, logId)
		clbService    = ClbService{client: m.(*TencentCloudClient).apiV3Conn}
		outErr, inErr error
		id            = d.Id()
	)
	request := clb.NewDeleteCertRequest()
	request.CertIds = []*string{&id}

	outErr = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		inErr = clbService.DeleteCertificate(ctx, request)
		if inErr != nil {
			return retryError(inErr)
		}
		return nil
	})

	if outErr != nil {
		log.Printf("[CRITAL]%s delete SSL certificate failed, reason:%+v", logId, outErr)
		return outErr
	}
	return nil
}
