# Direct Connect (DC) Examples

# ========== Data Sources ==========

# Query DC access points
data "tencentcloudenterprise_dc_access_points" "points" {
  region_id = "ap-guangzhou"
}

# Query DC instances
data "tencentcloudenterprise_dc_instances" "instances" {
  dcx_id = "dcx-xxxxx"
}

# ========== Resources ==========

# DC Instance (Direct Connect)
resource "tencentcloudenterprise_dc_instance" "dc" {
  dc_name                = "example-dc"
  access_point_id        = "ap-xxxxx"
  line_operator          = "ChinaTelecom"
  port_type              = "10GBase-LR"
  circuit_code           = "example-circuit"
  location               = "Example IDC"
  bandwidth              = 100
  redundant_dc_id        = ""
  vlan                   = 100
  tencent_address        = "192.168.1.1/30"
  customer_address       = "192.168.1.2/30"
  customer_contact_name  = "Admin"
  customer_contact_number = "13800138000"
}

# DC Dedicated Connection (DCX)
resource "tencentcloudenterprise_dc_dcx" "dcx" {
  dc_id            = cloud_dc_instance.dc.id
  dcx_name         = "example-dcx"
  network_type     = "VPC"
  network_region   = "ap-guangzhou"
  vpc_id           = "vpc-xxxxx"
  route_type       = "BGP"
  bgp_asn          = 65000
  vlan             = 100
  bandwidth        = 50
  customer_address = "192.168.2.2/30"
  tencent_address  = "192.168.2.1/30"
}
