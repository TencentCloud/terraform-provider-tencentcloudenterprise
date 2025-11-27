# Virtual Private Cloud DNS (VPCDNS) Examples

# ========== Data Sources ==========

# Query VPCDNS domains
data "tencentcloudenterprise_vpcdns_domains" "domains" {
  zone_id = "zone-xxxxx"
}

# Query VPCDNS records
data "tencentcloudenterprise_vpcdns_records" "records" {
  zone_id = "zone-xxxxx"
}

# Query VPCDNS forward rules
data "tencentcloudenterprise_vpcdns_forward_rules" "rules" {
  rule_id = "rule-xxxxx"
}

# ========== Resources ==========

# VPCDNS Domain (Private Zone)
resource "tencentcloudenterprise_vpcdns_domain" "domain" {
  domain = "example.local"
  remark = "Example private domain"
  
  vpc_set {
    uniq_vpc_id = "vpc-xxxxx"
    region      = "ap-guangzhou"
  }
  
  dns_forward_status = "DISABLED"
  
  tags = {
    env = "test"
  }
}

# VPCDNS Record
resource "tencentcloudenterprise_vpcdns_record" "record" {
  zone_id = cloud_vpcdns_domain.domain.id
  
  sub_domain = "www"
  record_type = "A"
  record_value = "10.0.1.10"
  ttl         = 600
  weight      = 0
}

# VPCDNS Record (CNAME)
resource "tencentcloudenterprise_vpcdns_record" "cname" {
  zone_id = cloud_vpcdns_domain.domain.id
  
  sub_domain = "blog"
  record_type = "CNAME"
  record_value = "www.example.local"
  ttl         = 600
}

# VPCDNS Record (MX)
resource "tencentcloudenterprise_vpcdns_record" "mx" {
  zone_id = cloud_vpcdns_domain.domain.id
  
  sub_domain = "@"
  record_type = "MX"
  record_value = "mail.example.local"
  mx          = 10
  ttl         = 600
}

# VPCDNS Forward Rule
resource "tencentcloudenterprise_vpcdns_forward_rule" "rule" {
  rule_name = "example-forward"
  rule_type = "DOWN"
  zone_id   = cloud_vpcdns_domain.domain.id
  endpoint  = "10.0.1.100"
}
