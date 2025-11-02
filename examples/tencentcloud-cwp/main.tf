# Cloud Workload Protection (CWP) Examples

# ========== Data Sources ==========

# Query CWP machines (simple)
data "cloud_cwp_machines_simple" "machines" {
  machine_type = "CVM"
  region       = "ap-guangzhou"
}

# ========== Resources ==========

# CWP License Order
resource "cloud_cwp_license_order" "license" {
  license_type = "ProVersion"
  license_num  = 10
  region_id    = 1
  project_id   = 0
  
  tags = {
    env = "production"
  }
}

# CWP License Bind Attachment
resource "cloud_cwp_license_bind_attachment" "bind" {
  resource_id = cloud_cwp_license_order.license.resource_id
  quuid       = "quuid-xxxxx"
}
