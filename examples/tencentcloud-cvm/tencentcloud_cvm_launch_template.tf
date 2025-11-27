resource "tencentcloudenterprise_cvm_launch_template" "launch_template" {
  launch_template_name = "test_launch_template"
  placement {
    project_id = 0
    zone       = var.availability_zone
  }
  image_id                            = "img-95xgn7er"
  launch_template_version_description = "test111"
  instance_type                       = "S5l.SMALL1"
  system_disk {
    disk_size = 50
    disk_type = "CLOUD_BASIC"
  }
  virtual_private_cloud {
    vpc_id             = "vpc-cs6ffr73"
    subnet_id          = "subnet-38oi34ta"
    as_vpc_gateway     = false
    ipv6_address_count = 0
  }
  internet_accessible {
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 0
    public_ip_assigned         = false
  }
  instance_count     = 1
  instance_name      = "test_instance_name"
  security_group_ids = ["sg-9s7k6qgw"]
  enhanced_service {
    security_service {
      enabled = true
    }
    monitor_service {
      enabled = true
    }
  }
  instance_charge_type = "POSTPAID_BY_HOUR"
}