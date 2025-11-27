data "tencentcloudenterprise_redis_instance_shards" "redis_shards_test" {
  instance_id = "crs-ppzu5lw8"
  filter_slave = true
}


