resource "tencentcloudenterprise_as_scaling_config" "launch_configuration" {
  configuration_name = "launch-configuration"
  image_id           = "img-95xgn7er"
  instance_types     = ["S5l.SMALL1"]
  project_id         = 0
  system_disk_type   = "CLOUD_PREMIUM"
  system_disk_size   = "50"

  data_disk {
    disk_type = "CLOUD_PREMIUM"
    disk_size = 50
  }

  internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
  internet_max_bandwidth_out = 5
  public_ip_assigned         = true
  #  password                   = "test123#"
  enhanced_security_service  = false
  enhanced_monitor_service   = false
  user_data                  = "dGVzdA=="
  security_group_ids = ["sg-9s7k6qgw"]

  instance_tags = {
    tag = "as"
  }

}