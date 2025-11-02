# Cloud Firewall (CFW) Examples

# ========== Resources ==========

# CFW NAT Instance
resource "cloud_cfw_nat_instance" "nat" {
  name          = "example-nat-firewall"
  width         = 20
  mode          = 1
  new_mode_items {
    vpc_list = ["vpc-xxxxx"]
    eips     = ["eip-xxxxx"]
  }
  cross_a_zone = 0
  zone_set     = ["ap-guangzhou-3"]
}

# CFW VPC Instance
resource "cloud_cfw_vpc_instance" "vpc" {
  name  = "example-vpc-firewall"
  mode  = 1
  vpc_fw_instances {
    fw_vpc_id = "vpc-xxxxx"
    fw_subnet_id = "subnet-xxxxx"
  }
  fw_vpc_cidr = "10.0.0.0/16"
}

# CFW NAT Policy
resource "cloud_cfw_nat_policy" "policy" {
  source_content = "10.0.1.0/24"
  source_type    = "net"
  target_content = "0.0.0.0/0"
  target_type    = "net"
  protocol       = "TCP"
  rule_action    = "log"
  port           = "80"
  direction      = 1
  enable         = "true"
  description    = "Example NAT firewall policy"
}

# CFW VPC Policy
resource "cloud_cfw_vpc_policy" "policy" {
  source_content = "10.0.1.0/24"
  source_type    = "net"
  dest_content   = "10.0.2.0/24"
  dest_type      = "net"
  protocol       = "TCP"
  rule_action    = "log"
  port           = "443"
  enable         = "true"
  description    = "Example VPC firewall policy"
}

# CFW Block Ignore
resource "cloud_cfw_block_ignore" "ignore" {
  ip        = "1.1.1.1"
  direction = 1
  rule_type = 1
  comment   = "Example block ignore rule"
}
