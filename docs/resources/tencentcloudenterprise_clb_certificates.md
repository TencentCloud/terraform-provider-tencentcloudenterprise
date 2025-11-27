---
subcategory: "Cloud Load Balancer(CLB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_clb_certificates"
sidebar_current: "docs-tencentcloudenterprise-resources-clb_certificates"
description: |-
  Provides a resource to create a SSL certificate.
---

# tencentcloudenterprise_clb_certificates

Provides a resource to create a SSL certificate.

## Example Usage

```hcl
resource "tencentcloudenterprise_clb_certificates" "foo" {
  alias     = "test-ssl-certificate"
  cert_type = "CA"
  cert      = "-----BEGIN CERTIFICATE-----\nMIIERzCCAq+gAwIBAgIBAjANBgkqhkiG9w0BAQsFADAoMQ0wCwYDVQQDEwR0ZXN0\nMRcwFQYDVQQKEw50ZXJyYWZvcm0gdGVzdDAeFw0xOTA4MTMwMzE5MzlaFw0yOTA4\nMTAwMzE5MzlaMC4xEzARBgNVBAMTCnNlcnZlciBzc2wxFzAVBgNVBAoTDnRlcnJh\nZm9ybS10ZXN0MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA1Ryp+DKK\nSNFKZsPtwfR+jzOnQ8YFieIKYgakV688d8YgpolenbmeEPrzT87tunFD7G9f6ALG\nND8rj7npj0AowxhOL/h/v1D9u0UsIaj5i2GWJrqNAhGLaxWiEB/hy5WOiwxDrGei\ngQqJkFM52Ep7G1Yx7PHJmKFGwN9FhIsFi1cNZfVRopZuCe/RMPNusNVZaIi+qcEf\nfsE1cmfmuSlG3Ap0RKOIyR0ajDEzqZn9/0R7VwWCF97qy8TNYk94K/1tq3zyhVzR\nZ83xOSfrTqEfb3so3AU2jyKgYdwr/FZS72VCHS8IslgnqJW4izIXZqgIKmHaRZtM\nN4jUloi6l/6lktt6Lsgh9xECecxziSJtPMaog88aC8HnMqJJ3kScGCL36GYG+Kaw\n5PnDlWXBaeiDe8z/eWK9+Rr2M+rhTNxosAVGfDJyxAXyiX49LQ0v7f9qzwc/0JiD\nbvsUv1cm6OgpoEMP9SXqqBdwGqeKbD2/2jlP48xlYP6l1SoJG3GgZ8dbAgMBAAGj\ndjB0MAwGA1UdEwEB/wQCMAAwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0PAQH/\nBAUDAweAADAdBgNVHQ4EFgQULwWKBQNLL9s3cb3tTnyPVg+mpCMwHwYDVR0jBBgw\nFoAUKwfrmq791mY831S6UHARHtgYnlgwDQYJKoZIhvcNAQELBQADggGBAMo5RglS\nAHdPgaicWJvmvjjexjF/42b7Rz4pPfMjYw6uYO8He/f4UZWv5CZLrbEe7MywaK3y\n0OsfH8AhyN29pv2x8g9wbmq7omZIOZ0oCAGduEXs/A/qY/hFaCohdkz/IN8qi6JW\nVXreGli3SrpcHFchSwHTyJEXgkutcGAsOvdsOuVSmplOyrkLHc8uUe8SG4j8kGyg\nEzaszFjHkR7g1dVyDVUedc588mjkQxYeAamJgfkgIhljWKMa2XzkVMcVfQHfNpM1\nn+bu8SmqRt9Wma2bMijKRG/Blm756LoI+skY+WRZmlDnq8zj95TT0vceGP0FUWh5\nhKyiocABmpQs9OK9HMi8vgSWISP+fYgkm/bKtKup2NbZBoO5/VL2vCEPInYzUhBO\njCbLMjNjtM5KriCaR7wDARgHiG0gBEPOEW1PIjZ9UOH+LtIxbNZ4eEIIINLHnBHf\nL+doVeZtS/gJc4G4Adr5HYuaS9ZxJ0W2uy0eQlOHzjyxR6Mf/rpnilJlcQ==\n-----END CERTIFICATE-----"
  cert_key  = "XXX"
}
```

## Argument Reference

The following arguments are supported:

* `cert_type` - (Required, String, ForceNew) Type of the SSL certificate. Valid values: `CA` and `SVR`.
* `cert` - (Required, String, ForceNew) Content of the SSL certificate. Not allowed newline at the start and end.
* `alias` - (Optional, String) Name of the SSL certificate.
* `cert_key` - (Optional, String, ForceNew) Key of the SSL certificate and required when certificate type is `SVR`. Not allowed newline at the start and end.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `begin_time` - Beginning time of the SSL certificate.
* `cert_algorithm` - Algorithm of the SSL certificate.
* `create_time` - Creation time of the SSL certificate.
* `domain` - Primary domain of the SSL certificate.
* `end_time` - Ending time of the SSL certificate.
* `product_zh_name` - Certificate authority.
* `project_id` - Project ID of the SSL certificate. Default is `0`.
* `status` - Status of the SSL certificate.
* `subject_names` - ALL domains included in the SSL certificate. Including the primary domain name.


## Import

ssl certificate can be imported using the id, e.g.

```

	$ terraform import tencentcloudenterprise_clb_certificates.cert GjTNRoK7

```

