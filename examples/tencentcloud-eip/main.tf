# Elastic IP (EIP) Examples

# ========== Data Sources ==========

# Query EIPs
data "tencentcloudenterprise_eips" "eips" {
  address_id = "eip-xxxxx"
}

# Query EIP address quota
data "tencentcloudenterprise_eip_address_quota" "quota" {}

# ========== Resources ==========

# EIP
resource "tencentcloudenterprise_eip" "eip" {
  name                       = "example-eip"
  type                       = "EIP"
  internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
  internet_max_bandwidth_out = 10
  
  tags = {
    env = "test"
  }
}

# EIP Association
resource "tencentcloudenterprise_eip_association" "association" {
  eip_id      = cloud_eip.eip.id
  instance_id = "ins-xxxxx"
}

# EIP Address Transform (Convert ordinary IP to EIP)
resource "tencentcloudenterprise_eip_address_transform" "transform" {
  instance_id = "ins-xxxxx"
}

# EIP Public Address Adjust (Adjust bandwidth)
resource "tencentcloudenterprise_eip_public_address_adjust" "adjust" {
  address_id                 = cloud_eip.eip.id
  internet_max_bandwidth_out = 20
}

# EIP Normal Address Return (Release EIP)
resource "tencentcloudenterprise_eip_normal_address_return" "return" {
  address_id = cloud_eip.eip.id
}
