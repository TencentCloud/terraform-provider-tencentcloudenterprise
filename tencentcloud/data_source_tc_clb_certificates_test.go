package tencentcloud

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudSslCertificatesDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccDataSourceTencentCloudSslCertificatesBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_ssl_certificates.foo"),
					resource.TestMatchResourceAttr("data.cloud_ssl_certificates.foo", "certificates.#", regexp.MustCompile(`^[1-9]\d*$`)),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.id"),
					resource.TestCheckResourceAttr("data.cloud_ssl_certificates.foo", "certificates.0.name", "keep-ssl-ca"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.type"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.project_id"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.cert"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.product_zh_name"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.domain"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.status"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.begin_time"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.end_time"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.create_time"),
				),
			},
		},
	})
}

func TestAccTencentCloudSslCertificatesDataSource_type(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccDataSourceTencentCloudSslCertificatesType,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_ssl_certificates.foo"),
					resource.TestMatchResourceAttr("data.cloud_ssl_certificates.foo", "certificates.#", regexp.MustCompile(`^[1-9]\d*$`)),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.id"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.name"),
					resource.TestCheckResourceAttr("data.cloud_ssl_certificates.foo", "certificates.0.type", "CA"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.project_id"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.cert"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.product_zh_name"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.domain"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.status"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.begin_time"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.end_time"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.create_time"),
				),
			},
		},
	})
}

func Test_11(t *testing.T) {
	fmt.Println(testAccSslCertificateCA)
}

func TestAccTencentCloudSslCertificatesDataSource_id(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccDataSourceTencentCloudSslCertificatesId,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_ssl_certificates.foo"),
					resource.TestCheckResourceAttr("data.cloud_ssl_certificates.foo", "certificates.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.id"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.name"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.type"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.project_id"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.cert"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.product_zh_name"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.domain"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.status"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.begin_time"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.end_time"),
					resource.TestCheckResourceAttrSet("data.cloud_ssl_certificates.foo", "certificates.0.create_time"),
					resource.TestCheckResourceAttr("data.cloud_ssl_certificates.foo", "subject_names.#", "0"),
				),
			},
		},
	})
}

var TestAccDataSourceTencentCloudSslCertificatesBasic = fmt.Sprintf(`
resource "cloud_ssl_certificate" "foo" {
  type = "CA"
  cert = "%s"
  name = "keep-ssl-ca"
}

data "cloud_ssl_certificates" "foo" {
  name = cloud_ssl_certificate.foo.name
}
`, testAccSslCertificateCA)

var TestAccDataSourceTencentCloudSslCertificatesType = fmt.Sprintf(`
resource "cloud_ssl_certificate" "foo" {
  type = "CA"
  cert = "%s"
  name = "keep-ssl-ca"
}

data "cloud_ssl_certificates" "foo" {
  type = cloud_ssl_certificate.foo.type
}
`, testAccSslCertificateCA)

var TestAccDataSourceTencentCloudSslCertificatesId = fmt.Sprintf(`
resource "cloud_ssl_certificate" "foo" {
  type = "CA"
  cert = "%s"
  name = "keep-ssl-ca"
}

data "cloud_ssl_certificates" "foo" {
  id = cloud_ssl_certificate.foo.id
}
`, testAccSslCertificateCA)
