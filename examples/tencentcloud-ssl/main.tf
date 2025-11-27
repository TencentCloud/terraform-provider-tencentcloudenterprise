# SSL Certificate Examples

# Note: SSL certificate resources are managed through CLB
# See examples/tencentcloud-clb/main.tf for certificate examples

# CLB Certificate creation example:
# resource "tencentcloudenterprise_clb_certificates" "cert" {
#   certificate_name    = "example-cert"
#   certificate_type    = "SERVER"
#   certificate_content = file("./server.crt")
#   certificate_key     = file("./server.key")
# }

# For standalone SSL certificate management, please refer to:
# - CLB Certificates (cloud_clb_certificates)
# - CDN Certificate operations
