resource "tencentcloudenterprise_redis_param" "redis_param_test" {
  instance_id = "crs-czzuta4a"
  instance_params = {
    appendonly = "yes"
  }
}


