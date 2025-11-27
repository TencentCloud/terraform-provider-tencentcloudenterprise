resource "tencentcloudenterprise_vpcdns_record" "foo" {
  domain_id   = "745"
  mx          = 0
  record_type = "A"
  sub_domain  = "www"
  value       = "203.0.113.3"
  weight      = "100"
}