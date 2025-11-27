/*
Provides a resource to create a clb replace_cert_for_lbs

Example Usage
```hcl

	resource "tencentcloudenterprise_clb_replace_cert" "foo" {
	  old_cert_id  = "76FqSmKT"
	  new_alias    = "INTERNAL"
	  new_cert     = "-----BEGIN CERTIFICATE-----\nMIIGcDCCBNigAwIBAgIQYdY1JwrY35tG69u0An0j9zANBgkqhkiG9w0BAQwFADBZ\nMQswCQYDVQQGEwJDTjElMCMGA1UEChMcVHJ1c3RBc2lhIFRlY2hub2xvZ2llcywg\nSW5jLjEjMCEGA1UEAxMaVHJ1c3RBc2lhIFJTQSBEViBUTFMgQ0EgRzIwHhcNMjMw\nNzE3MDAwMDAwWhcNMjQwNzE2MjM1OTU5WjAWMRQwEgYDVQQDEwticnVjZWxpbi5j\nYzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAIC9Dm3zurJn04Gb1TLa\nHn/2um2hgR5OS/katZVGugSVZAbpONGqzXlSsqhfdu9kqNBIiOjhmRea0wxGQ2GZ\nsfLC/pAxY7IG3Qgek9z1Jb+DMXYkW6JIlxnczqkzqeDsb2xyyujTuep/0SfkjNTY\ntqPexBlUlPApgts++q0Lt/XM2+sKkd/VVkA38agTD4pcTxeuBnKVLmHTYobuXstE\ny0dCuSiJvZ+yhbOA8AxMLFD8zg3S2qc5lshwhFF1CG6coXkkIa+QfsLwPB2oYFs1\nXSvsra1HByTWf1Op4nfXcDCddb/MMNSLqpseza74Gg7Hgj652YpAK9w1aKOG2o4v\ntAcCAwEAAaOCAvUwggLxMB8GA1UdIwQYMBaAFF86fBEQfgxncWHci6O1AANn9Vcc\nMB0GA1UdDgQWBBT4UEcDO5me2uu+Ow+/X5h7/Yf80DAOBgNVHQ8BAf8EBAMCBaAw\nDAYDVR0TAQH/BAIwADAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwSQYD\nVR0gBEIwQDA0BgsrBgEEAbIxAQICMTAlMCMGCCsGAQUFBwIBFhdodHRwczovL3Nl\nY3RpZ28uY29tL0NQUzAIBgZngQwBAgEwfQYIKwYBBQUHAQEEcTBvMEIGCCsGAQUF\nBzAChjZodHRwOi8vY3J0LnRydXN0LXByb3ZpZGVyLmNuL1RydXN0QXNpYVJTQURW\nVExTQ0FHMi5jcnQwKQYIKwYBBQUHMAGGHWh0dHA6Ly9vY3NwLnRydXN0LXByb3Zp\nZGVyLmNuMCcGA1UdEQQgMB6CC2JydWNlbGluLmNjgg93d3cuYnJ1Y2VsaW4uY2Mw\nggF9BgorBgEEAdZ5AgQCBIIBbQSCAWkBZwB1AHb/iD8KtvuVUcJhzPWHujS0pM27\nKdxoQgqf5mdMWjp0AAABiWKyWxsAAAQDAEYwRAIgY1Q4nHmmCVy4ZdJNLjaPyK0J\nf7iYvCgE1CfPYgxd6DQCIA/qC62caUv+zCif+DD5sVVVOzq2KJr4O2BRJUgLXc/y\nAHYA2ra/az+1tiKfm8K7XGvocJFxbLtRhIU0vaQ9MEjX+6sAAAGJYrJbaQAABAMA\nRzBFAiEA7K98sak+X2cofLsnkAZUDaZxsui4gigg/5x3ylcuencCIHfRyfxB6OQS\ns/EzmYHEuLRw6XqgMsgJiXy21ZCMDFOpAHYA7s3QZNXbGs7FXLedtM0TojKHRny8\n7N7DUUhZRnEftZsAAAGJYrJbQwAABAMARzBFAiBvUe5D3FWcFnbBVZuq+SOaP5WG\nSJYp553bNDX9eg6khQIhANRUdxQp1b1IOcjiCoylvAyWyxGhW97jj6HHZ9K23Mmc\nMA0GCSqGSIb3DQEBDAUAA4IBgQA8/aV/QGQ+oycAoe4LtyTqdmW8z3+/g7N00Lho\n32q96Bbo999C1w2ZCZ/A33Jun5HUjktXuWTdENA8zg2JRXECtikO0rElD8Popj6t\n1ZsYsBKNZOfEklSTTHTr9zZti3xuCzi5/2vYezFvV7/3Fvx849xxivK4xgGnv/Zl\neioo2D6MSAxwjBGxqoSgjOM1ozqY1aH9a0OZ1d+Dv2EhHaCIzInOnYTk2qerJc7o\n6fsawwcE+5DuBXXsZksKiTUMgVVO+cbWFV05ldhRzJlOTPkWiBTYIB2+DjafMczK\ntdi9xekiAKYMOMgepl611ofo9REjUfoJQqrVDjMdYDZ6t9rrqsSzNi6YY9EwWVGJ\nm11W4nZzGmqv0jTZeUSrQc+ffXl127cuWISEMJP21z+xgqBgtar3iybMIt22p2AH\n9hpmHu5Uf6eiT9k5QFzPIyeJ2BSqzykcQa4E+Kpnhx9cHLGCRDeQR4WAlAWmx0O1\nHzCs35ma76+3wOlUD36V/ttt5uY=\n-----END CERTIFICATE-----"
	  new_cert_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAgL0ObfO6smfTgZvVMtoef/a6baGBHk5L+Rq1lUa6BJVkBuk4\n0arNeVKyqF9272So0EiI6OGZF5rTDEZDYZmx8sL+kDFjsgbdCB6T3PUlv4MxdiRb\nokiXGdzOqTOp4OxvbHLK6NO56n/RJ+SM1Ni2o97EGVSU8CmC2z76rQu39czb6wqR\n39VWQDfxqBMPilxPF64GcpUuYdNihu5ey0TLR0K5KIm9n7KFs4DwDEwsUPzODdLa\npzmWyHCEUXUIbpyheSQhr5B+wvA8HahgWzVdK+ytrUcHJNZ/U6nid9dwMJ11v8ww\n1Iuqmx7NrvgaDseCPrnZikAr3DVoo4baji+0BwIDAQABAoIBAAJztdmNOJ4BLj6i\nDT6VzTv2CqPToHeETWVHpfPtw9WH8Y3shjPOmK5sQyO3pP/YpWQAR5QLXOh8I7a5\nnKKmERdbyc0om0ek6MYQYzh8xbEL94wDaPrRrziY7i+vr+NwTKQjJUw6C8EujWDZ\nLwWw5dh/74yDxRc0NXy+HMSka2bLmZigSs5Fu/NjShsUxgA2yhmYuE4iHi8Jzryw\nBZIjqkq7trMjjqmavA3CnC5uisIIrfEfbAL646Kyf5SMKt4mag0WejVe634T/FWw\nhXa12dCbtzfmrEm8Mx4ettywVu0CndVcqct6d866T9FiykxCS348K31IrrP+N5Js\n1b8L0nECgYEAtfVbra5pZiiLaPxcV7pS36ko4XapKruCIYxYVqf0DX6TGUoxiMAh\nOsbICj+ridCNU4fwSirI0sjgkWuHETJIIrniQGbHoAmTwbAIQyseWmceRd/uNko+\na25k3M0Sbb9P3CmgEM6YppklobafyBRYZ5lfc5DAIm4BuztIhRxoZokCgYEAtR/E\n01GMh9LXe1eS6ITcatSw/r2a8uXjFEP6snIv7THqRO/Xc4nvRieuWjkFI2n2NSLS\nHDS/fcHQEzR2SMvGpj7wYR2BCeuPD3K2vXOd64nivq6nmYDovznUgBJHVvSYa9Es\n2s0Q8GftTSirjberEftpxKb4Y0H3w2zUCH33og8CgYBYBsgAYQUaZ/jxpEykID6h\ndzuQv5AdXTMaOcQuv/fgY2CUdoE6MACjZ7E0zBKXjG4if/wuVT5sQsPpdgSUvCeL\nrTOYhmCCur9hj2Cf5gc8IvDRSwD6ALbr0C85ZnhI4amnz/dgyiGtTx+WeTwZVkZi\ncB9uUBOzVFbAFnEB/HlBSQKBgQC01uVUAAYJzq+qzMM3OQBCAd4+Wd0NTA4vu7fg\n1zWW49F9xuIcz8mBCDmCh4/jzfYvE3cpBllzHEG+CxFWmW7bqdejfyvJVdHeoLBn\n87nm5CLqM8PO9fBsjTboMFfeMBTHAXCBfWG+RmWeNk8jDhDVwWnXGMbDg6f3DP+f\nAvZubQKBgQCtI45G5gEwvM1gHuRd7gwINB3+oetZiZ1uT96xu6WjQ4uUUm8vtRIP\nO8k9sC0UwOfMwPla6x/8Fqmx5HbKHQ2dkOfq/1IvdwcmeKV9Vlk8jZZ6moIX44QP\nQp9Jc68JTMIOxyZCfSmYp8Sj3D8drrYpyUuXDmlq/YKI+SYS3scXuw==\n-----END RSA PRIVATE KEY-----"
	}

```

```
terraform import cloud_clb_replace_cert.replace_cert_for_lbs replace_cert_for_lbs_id
```
*/
package tencentcloud

import (
	"log"
	"strings"

	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_clb_replace_cert", CNDescription{
		TerraformTypeCN: "更换证书",
		DescriptionCN:   "提供更换证书资源，用于批量更换CLB实例的SSL证书。",
		AttributesCN: map[string]string{
			"new_alias":    "新证书别名",
			"new_cert":     "新的证书",
			"new_cert_key": "证书的私钥",
			"old_cert_id":  "旧证书ID",
		},
	})
}
func resourceTencentCloudClbReplaceCert() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a clb replace_cert_for_lbs",
		Create:      resourceTencentCloudClbReplaceCertCreate,
		Read:        resourceTencentCloudClbReplaceCertRead,
		Update:      resourceTencentCloudClbReplaceCertUpdate,
		Delete:      resourceTencentCloudClbReplaceCertDelete,
		Schema: map[string]*schema.Schema{
			"new_alias": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Name of the SSL certificate.",
			},
			"new_cert": {
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
			"new_cert_key": {
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

			"old_cert_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "ID of the certificate to be replaced, which can be a server certificate or a client certificate.",
			},
		},
	}
}

func resourceTencentCloudClbReplaceCertCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_clb_replace_cert_for_lbs.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = clb.NewReplaceCertRequest()
	)

	// if v, ok := d.GetOk("old_cert_id"); ok {
	// 	request.OldCertId = helper.String(v.(string))
	// }
	oldCertificateID := d.Get("old_cert_id").(string)
	request.OldCertId = helper.String(oldCertificateID)

	if v, ok := d.GetOk("new_alias"); ok {
		request.NewCertAlias = helper.String(v.(string))
	}

	if v, ok := d.GetOk("new_cert"); ok {
		request.NewCertContent = helper.String(v.(string))
	}

	if v, ok := d.GetOk("new_cert_key"); ok {
		request.NewCertKey = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().ReplaceCert(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate clb replaceCert failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(oldCertificateID)

	return resourceTencentCloudClbReplaceCertRead(d, meta)
}

func resourceTencentCloudClbReplaceCertRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_clb_replace_cert_for_lbs.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudClbReplaceCertUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_clb_replace_cert_for_lbs.update")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudClbReplaceCertDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_clb_replace_cert_for_lbs.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
