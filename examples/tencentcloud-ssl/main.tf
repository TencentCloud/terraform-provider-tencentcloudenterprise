resource "tencentcloudenterprise_ssl_certificate" "ca" {
  name = "ssl-ca"
  type = "CA"
  cert = var.ca
}

resource "tencentcloudenterprise_ssl_certificate" "svr" {
  name = "ssl-svr"
  type = "SVR"
  cert = var.cert
  key  = var.key
}

data "tencentcloudenterprise_ssl_certificates" "ca" {
  name = tencentcloudenterprise_ssl_certificate.ca.name
}

data "tencentcloudenterprise_ssl_certificates" "svr" {
  type = tencentcloudenterprise_ssl_certificate.svr.type
}