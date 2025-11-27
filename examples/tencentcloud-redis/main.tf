# TencentDB for Redis Examples
# Based on actual resource schema from resource_tc_redis_instance.go

# ========== Data Sources ==========

# Query Redis zone config (recommended to use before creating instance)
data "tencentcloudenterprise_redis_zone_config" "config" {
  region = "ap-guangzhou"
}

# Query Redis instances
data "tencentcloudenterprise_redis_instances" "instances" {
  instance_id = "crs-xxxxx"
}

# Query Redis backup
data "tencentcloudenterprise_redis_backup" "backup" {
  instance_id = "crs-xxxxx"
}

# Query Redis backup download info
data "tencentcloudenterprise_redis_backup_download_info" "download" {
  instance_id = "crs-xxxxx"
  backup_id   = "backup-xxxxx"
}

# Query Redis param records
data "tencentcloudenterprise_redis_param_records" "params" {
  instance_id = "crs-xxxxx"
}

# Query Redis instance shards
data "tencentcloudenterprise_redis_instance_shards" "shards" {
  instance_id = "crs-xxxxx"
}

# Query Redis instance task list
data "tencentcloudenterprise_redis_instance_task_list" "tasks" {
  instance_id = "crs-xxxxx"
}

# Query Redis instance node info
data "tencentcloudenterprise_redis_instance_node_info" "nodes" {
  instance_id = "crs-xxxxx"
}

# ========== Resources ==========

# Redis Instance (Standard version)
# Required fields: availability_zone, type_id, mem_size, vpc_id, subnet_id, security_groups
resource "tencentcloudenterprise_redis_instance" "standard" {
  # Required fields
  availability_zone = "ap-guangzhou-3"
  type_id           = 6  # Redis 5.0 Standard, refer to cloud_redis_zone_config
  mem_size          = 2048  # Memory size in MB
  vpc_id            = "vpc-xxxxx"
  subnet_id         = "subnet-xxxxx"
  security_groups   = ["sg-xxxxx"]  # Required! At least one security group
  
  # Optional but recommended
  name              = "example-redis-standard"
  password          = "Test12345"  # 8-16 characters
  port              = 6379
  
  # For standard version, these are optional
  redis_replicas_num = 1  # Default is 1
  
  tags = {
    env = "test"
  }
}

# Redis Instance (Cluster version with replicas)
resource "tencentcloudenterprise_redis_instance" "cluster" {
  # Required fields
  availability_zone = "ap-guangzhou-3"
  type_id           = 7  # Redis Cluster, refer to cloud_redis_zone_config
  mem_size          = 2048
  vpc_id            = "vpc-xxxxx"
  subnet_id         = "subnet-xxxxx"
  security_groups   = ["sg-xxxxx"]
  
  # Cluster specific
  redis_shard_num    = 3  # Number of shards
  redis_replicas_num = 2  # Number of replicas per shard
  
  # Optional
  name     = "example-redis-cluster"
  password = "Test12345"
  port     = 6379
  
  tags = {
    env = "production"
  }
}

# Redis Instance with multi-zone replicas
resource "tencentcloudenterprise_redis_instance" "multi_zone" {
  availability_zone = "ap-guangzhou-3"
  type_id           = 6
  mem_size          = 2048
  vpc_id            = "vpc-xxxxx"
  subnet_id         = "subnet-xxxxx"
  security_groups   = ["sg-xxxxx"]
  
  # Multi-zone configuration
  redis_replicas_num = 3
  replica_zone_ids   = [100003, 100004, 100006]  # Must match replicas_num
  
  name     = "example-redis-multi-zone"
  password = "Test12345"
  
  tags = {
    env = "production"
  }
}

# Redis Backup Config
resource "tencentcloudenterprise_redis_backup_config" "backup" {
  redis_id      = cloud_redis_instance.standard.id
  backup_time   = "02:00-03:00"
  backup_period = ["Monday", "Wednesday", "Friday"]
}

# Redis Param
resource "tencentcloudenterprise_redis_param" "param" {
  instance_id = cloud_redis_instance.standard.id
  instance_params = {
    "timeout"                    = "300"
    "maxmemory-policy"           = "allkeys-lru"
    "notify-keyspace-events"     = "Ex"
  }
}

# Redis Replica Readonly
resource "tencentcloudenterprise_redis_replica_readonly" "readonly" {
  instance_id     = cloud_redis_instance.multi_zone.id
  readonly_policy = ["master", "slave"]
  operate         = "enable"
}

# Redis Clear Instance Operation
resource "tencentcloudenterprise_redis_clear_instance_operation" "clear" {
  instance_id = cloud_redis_instance.standard.id
  password    = "Test12345"
}

# Redis Startup Instance Operation
resource "tencentcloudenterprise_redis_startup_instance_operation" "startup" {
  instance_id = cloud_redis_instance.standard.id
}
