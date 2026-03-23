resource "tencentcloudenterprise_vpcdns_record" "foo" {
  domain_id   = "745"
  mx          = 0
  record_type = "A"
  sub_domain  = "www"
  value       = "192.168.1.3"
  weight      = "100"
}