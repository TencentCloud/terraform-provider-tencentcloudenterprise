# Bare Metal Server (BMS) Examples

# ========== Data Sources ==========

# Query BMS instances
data "cloud_bms_instances" "instances" {
  instance_ids = ["bms-xxxxx"]
}

# Query BMS placement groups
data "cloud_bms_placement_groups" "groups" {
  placement_group_ids = ["pg-xxxxx"]
}

# Query BMS flavors
data "cloud_bms_flavors" "flavors" {
  zone = "ap-guangzhou-3"
}

# ========== Resources ==========

# BMS Placement Group
resource "cloud_bms_placement_group" "group" {
  placement_group_name = "example-pg"
  placement_group_type = "HOST"
}

# BMS Instance
resource "cloud_bms_instance" "instance" {
  instance_name     = "example-bms"
  zone              = "ap-guangzhou-3"
  instance_type     = "BMS.S5.MEDIUM4"
  image_id          = "img-xxxxx"
  system_disk_type  = "LOCAL_BASIC"
  system_disk_size  = 100
  vpc_id            = "vpc-xxxxx"
  subnet_id         = "subnet-xxxxx"
  placement_group_id = cloud_bms_placement_group.group.id

  data_disks {
    disk_type = "LOCAL_BASIC"
    disk_size = 500
  }

  internet_accessible {
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 10
    public_ip_assigned         = true
  }
}
