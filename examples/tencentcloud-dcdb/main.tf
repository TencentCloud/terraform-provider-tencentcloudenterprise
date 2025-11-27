# TDSQL for MySQL (DCDB) Examples

# ========== Data Sources ==========

# Query DCDB instances
data "tencentcloudenterprise_dcdb_instances" "instances" {
  instance_id = "dcdb-xxxxx"
}

# Query DCDB accounts
data "tencentcloudenterprise_dcdb_accounts" "accounts" {
  instance_id = "dcdb-xxxxx"
}

# Query DCDB databases
data "tencentcloudenterprise_dcdb_databases" "databases" {
  instance_id = "dcdb-xxxxx"
}

# Query DCDB parameters
data "tencentcloudenterprise_dcdb_parameters" "parameters" {
  instance_id = "dcdb-xxxxx"
}

# Query DCDB shards
data "tencentcloudenterprise_dcdb_shards" "shards" {
  instance_id = "dcdb-xxxxx"
}

# Query DCDB security groups
data "tencentcloudenterprise_dcdb_security_groups" "sgs" {
  instance_id = "dcdb-xxxxx"
}

# Query DCDB database objects
data "tencentcloudenterprise_dcdb_database_objects" "objects" {
  instance_id = "dcdb-xxxxx"
  db_name     = "example_db"
}

# Query DCDB database tables
data "tencentcloudenterprise_dcdb_database_tables" "tables" {
  instance_id = "dcdb-xxxxx"
  db_name     = "example_db"
}

# ========== Resources ==========

# DCDB Instance
resource "tencentcloudenterprise_dcdb_instance" "instance" {
  instance_name = "example-dcdb"
  zones         = ["ap-guangzhou-3", "ap-guangzhou-4"]
  period        = 1
  shard_memory  = 2
  shard_storage = 100
  shard_node_count = 2
  shard_count   = 2
  vpc_id        = "vpc-xxxxx"
  subnet_id     = "subnet-xxxxx"
  db_version_id = "8.0"
  
  resource_tags {
    tag_key   = "env"
    tag_value = "test"
  }
}

# DCDB Account
resource "tencentcloudenterprise_dcdb_account" "account" {
  instance_id = cloud_dcdb_instance.instance.id
  user_name   = "example_user"
  host        = "%"
  password    = "Password123!"
  description = "Example DCDB account"
  read_only   = 0
}

# DCDB Account Privileges
resource "tencentcloudenterprise_dcdb_account_privileges" "privileges" {
  instance_id = cloud_dcdb_instance.instance.id
  user_name   = cloud_dcdb_account.account.user_name
  host        = cloud_dcdb_account.account.host
  
  privileges = ["SELECT", "INSERT", "UPDATE", "DELETE"]
  
  database {
    db_name = "example_db"
  }
}

# DCDB Security Group Attachment
resource "tencentcloudenterprise_dcdb_security_group_attachment" "sg_attachment" {
  instance_id        = cloud_dcdb_instance.instance.id
  security_group_ids = ["sg-xxxxx"]
}

# DCDB DB Parameters
resource "tencentcloudenterprise_dcdb_db_parameters" "params" {
  instance_id = cloud_dcdb_instance.instance.id
  params {
    param = "max_connections"
    value = "1000"
  }
  params {
    param = "innodb_buffer_pool_size"
    value = "1073741824"
  }
}

# DCDB DB Sync Mode Config
resource "tencentcloudenterprise_dcdb_db_sync_mode_config" "sync_mode" {
  instance_id = cloud_dcdb_instance.instance.id
  sync_mode   = 2  # Strong sync
}

# DCDB Encrypt Attributes Config
resource "tencentcloudenterprise_dcdb_encrypt_attributes_config" "encrypt" {
  instance_id    = cloud_dcdb_instance.instance.id
  encrypt_enabled = 1
}

# DCDB Instance Config
resource "tencentcloudenterprise_dcdb_instance_config" "config" {
  instance_id = cloud_dcdb_instance.instance.id
  
  rs_access_strategy = 1
  extra_config {
    rs_access_strategy = 1
  }
}

# DCDB Activate Hour Instance Operation
resource "tencentcloudenterprise_dcdb_activate_hour_instance_operation" "activate" {
  instance_id = cloud_dcdb_instance.instance.id
}

# DCDB Isolate Hour Instance Operation
resource "tencentcloudenterprise_dcdb_isolate_hour_instance_operation" "isolate" {
  instance_id = cloud_dcdb_instance.instance.id
}

# DCDB Flush Binlog Operation
resource "tencentcloudenterprise_dcdb_flush_binlog_operation" "flush_binlog" {
  instance_id = cloud_dcdb_instance.instance.id
}

# DCDB Switch DB Instance HA Operation
resource "tencentcloudenterprise_dcdb_switch_db_instance_ha_operation" "switch_ha" {
  instance_id = cloud_dcdb_instance.instance.id
  zone        = "ap-guangzhou-4"
}

# DCDB Cancel DCN Job Operation
resource "tencentcloudenterprise_dcdb_cancel_dcn_job_operation" "cancel_dcn" {
  instance_id = cloud_dcdb_instance.instance.id
}
