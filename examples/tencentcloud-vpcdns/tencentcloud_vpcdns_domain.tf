resource "tencentcloudenterprise_vpcdns_domain" "foo" {
  domain      = "brucezylin.cc"
  dns_forward_status = "ENABLED"
  tags = {
    "createdBy" = "terraform3"
  }
}
