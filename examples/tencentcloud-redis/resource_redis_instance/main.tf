resource "tencentcloudenterprise_redis_instance" "redis_instance_test" {
  availability_zone = var.availability_zone
  type_id              = 8
  password          = "waw1111sda"
  mem_size          = 8192
  name              = "terrform_test"
  port              = 6379
  redis_shard_num   = 1
  no_auth           = true
  vpc_id            = "vpc-9u3x42tn"
  security_groups   = ["sg-23th7puy"]
  subnet_id         = "subnet-3xw9ickw"
  redis_replicas_num = 1
  replicas_read_only = false

  tags = {
    "test" = "test"
  }
}

#resource "tencentcloudenterprise_redis_backup_config" "redis_backup_config" {
#  redis_id      = cloud_redis_instance.redis_instance_test.id
#  backup_time   = "01:00-02:00"
#  backup_period = ["Saturday", "Sunday"]
#}
#
#data "tencentcloudenterprise_redis_instances" "redis" {
#  zone       = var.availability_zone
#  search_key = cloud_redis_instance.redis_instance_test.id
#}
#
#data "tencentcloudenterprise_redis_instances" "redis-tags" {
#  zone = var.availability_zone
#  tags = cloud_redis_instance.redis_instance_test.tags
#}
#
#resource "tencentcloudenterprise_redis_instance" "redis_instance_prepaid_test" {
#  availability_zone                   = var.availability_zone
#  type_id                             = 2
#  password                            = "test12345789"
#  mem_size                            = 8192
#  name                                = "terrform_prepaid_test"
#  port                                = 6379
#  charge_type                         = "PREPAID"
#  prepaid_period                      = 1
#  force_delete                        = true
#
#  tags = {
#    "test" = "prepaid test"
#  }
#}
