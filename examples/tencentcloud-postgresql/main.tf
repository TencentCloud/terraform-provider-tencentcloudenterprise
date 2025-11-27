# TencentDB for PostgreSQL Examples

# ========== Data Sources ==========

# Query PostgreSQL instances
data "tencentcloudenterprise_postgresql_instances" "instances" {
  db_instance_id = "postgres-xxxxx"
}

# Query PostgreSQL spec infos
data "tencentcloudenterprise_postgresql_specinfos" "specs" {
  zone = "ap-guangzhou-3"
}

# Query PostgreSQL xlogs
data "tencentcloudenterprise_postgresql_xlogs" "xlogs" {
  db_instance_id = "postgres-xxxxx"
  start_time     = "2024-01-01 00:00:00"
  end_time       = "2024-01-31 23:59:59"
}

# Query PostgreSQL parameter templates
data "tencentcloudenterprise_postgresql_parameter_templates" "templates" {}

# Query PostgreSQL readonly groups
data "tencentcloudenterprise_postgresql_readonly_groups" "ro_groups" {
  master_db_instance_id = "postgres-xxxxx"
}

# Query PostgreSQL base backups
data "tencentcloudenterprise_postgresql_base_backups" "base_backups" {
  db_instance_id = "postgres-xxxxx"
}

# Query PostgreSQL log backups
data "tencentcloudenterprise_postgresql_log_backups" "log_backups" {
  db_instance_id = "postgres-xxxxx"
}

# Query PostgreSQL backup download URLs
data "tencentcloudenterprise_postgresql_backup_download_urls" "download" {
  db_instance_id = "postgres-xxxxx"
  backup_id      = "backup-xxxxx"
}

# Query PostgreSQL instance classes
data "tencentcloudenterprise_postgresql_db_instance_classes" "classes" {
  zone             = "ap-guangzhou-3"
  db_engine        = "postgresql"
  db_major_version = "13"
}

# Query PostgreSQL default parameters
data "tencentcloudenterprise_postgresql_default_parameters" "defaults" {
  db_major_version = "13"
  db_engine        = "postgresql"
}

# Query PostgreSQL recovery time
data "tencentcloudenterprise_postgresql_recovery_time" "recovery" {
  db_instance_id = "postgres-xxxxx"
}

# Query PostgreSQL regions
data "tencentcloudenterprise_postgresql_regions" "regions" {}

# Query PostgreSQL versions
data "tencentcloudenterprise_postgresql_db_instance_versions" "versions" {}

# Query PostgreSQL zones
data "tencentcloudenterprise_postgresql_zones" "zones" {}

# ========== Resources ==========

# PostgreSQL Instance
resource "tencentcloudenterprise_postgresql_instance" "instance" {
  name              = "example-postgres"
  availability_zone = "ap-guangzhou-3"
  charge_type       = "POSTPAID_BY_HOUR"
  vpc_id            = "vpc-xxxxx"
  subnet_id         = "subnet-xxxxx"
  db_major_version  = "13"
  db_engine         = "postgresql"
  instance_charge_type = "POSTPAID_BY_HOUR"
  
  db_node_set {
    role = "Primary"
    zone = "ap-guangzhou-3"
  }
  
  storage       = 100
  cpu           = 2
  memory        = 4
  project_id    = 0
  security_group_ids = ["sg-xxxxx"]
  
  tags = {
    env = "test"
  }
}

# PostgreSQL ReadOnly Instance
resource "tencentcloudenterprise_postgresql_readonly_instance" "readonly" {
  master_db_instance_id = cloud_postgresql_instance.instance.id
  name                  = "example-readonly"
  availability_zone     = "ap-guangzhou-3"
  instance_charge_type  = "POSTPAID_BY_HOUR"
  storage               = 100
  cpu                   = 2
  memory                = 4
  vpc_id                = "vpc-xxxxx"
  subnet_id             = "subnet-xxxxx"
  security_group_ids    = ["sg-xxxxx"]
}

# PostgreSQL ReadOnly Group
resource "tencentcloudenterprise_postgresql_readonly_group" "ro_group" {
  master_db_instance_id = cloud_postgresql_instance.instance.id
  name                  = "example-ro-group"
  project_id            = 0
  vpc_id                = "vpc-xxxxx"
  subnet_id             = "subnet-xxxxx"
  
  replay_lag_eliminate = 1
  replay_latency_eliminate = 1
  max_replay_lag       = 100
  max_replay_latency   = 512
  min_delay_eliminate_reserve = 1
}

# PostgreSQL ReadOnly Attachment
resource "tencentcloudenterprise_postgresql_readonly_attachment" "attachment" {
  db_instance_id           = cloud_postgresql_readonly_instance.readonly.id
  readonly_group_id        = cloud_postgresql_readonly_group.ro_group.id
}

# PostgreSQL Parameter Template
resource "tencentcloudenterprise_postgresql_parameter_template" "template" {
  template_name        = "example-template"
  db_major_version     = "13"
  db_engine            = "postgresql"
  template_description = "Example parameter template"
}

# PostgreSQL Backup Plan Config
resource "tencentcloudenterprise_postgresql_backup_plan_config" "backup" {
  db_instance_id           = cloud_postgresql_instance.instance.id
  base_backup_retention_period = 7
  min_backup_start_time    = "00:00:00"
  max_backup_start_time    = "01:00:00"
  backup_period            = ["monday", "wednesday", "friday"]
}

# PostgreSQL Security Group Config
resource "tencentcloudenterprise_postgresql_security_group_config" "sg_config" {
  db_instance_id     = cloud_postgresql_instance.instance.id
  security_group_ids = ["sg-xxxxx", "sg-yyyyy"]
}

# PostgreSQL Backup Download Restriction Config
resource "tencentcloudenterprise_postgresql_backup_download_restriction_config" "restriction" {
  restriction_type = "CUSTOMIZE"
  vpc_restriction_effect = "ALLOW"
  vpc_id_set             = ["vpc-xxxxx"]
  ip_restriction_effect  = "ALLOW"
  ip_set                 = ["1.1.1.1"]
}

# PostgreSQL Restart Instance Operation
resource "tencentcloudenterprise_postgresql_restart_db_instance_operation" "restart" {
  db_instance_id = cloud_postgresql_instance.instance.id
}

# PostgreSQL Renew Instance Operation
resource "tencentcloudenterprise_postgresql_renew_db_instance_operation" "renew" {
  db_instance_id = cloud_postgresql_instance.instance.id
  period         = 1
}

# PostgreSQL Isolate Instance Operation
resource "tencentcloudenterprise_postgresql_isolate_db_instance_operation" "isolate" {
  db_instance_ids = [cloud_postgresql_instance.instance.id]
}

# PostgreSQL Disisolate Instance Operation
resource "tencentcloudenterprise_postgresql_disisolate_db_instance_operation" "disisolate" {
  db_instance_ids = [cloud_postgresql_instance.instance.id]
}

# PostgreSQL Rebalance ReadOnly Group Operation
resource "tencentcloudenterprise_postgresql_rebalance_readonly_group_operation" "rebalance" {
  readonly_group_id = cloud_postgresql_readonly_group.ro_group.id
}

# PostgreSQL Delete Log Backup Operation
resource "tencentcloudenterprise_postgresql_delete_log_backup_operation" "delete_log" {
  db_instance_id = cloud_postgresql_instance.instance.id
  log_backup_id  = "log-xxxxx"
}

# PostgreSQL Base Backup
resource "tencentcloudenterprise_postgresql_base_backup" "backup" {
  db_instance_id = cloud_postgresql_instance.instance.id
}

# PostgreSQL Modify Account Remark Operation
resource "tencentcloudenterprise_postgresql_modify_account_remark_operation" "remark" {
  db_instance_id = cloud_postgresql_instance.instance.id
  user_name      = "postgres"
  remark         = "Administrator account"
}

# PostgreSQL Modify Switch Time Period Operation
resource "tencentcloudenterprise_postgresql_modify_switch_time_period_operation" "switch_time" {
  db_instance_id   = cloud_postgresql_instance.instance.id
  switch_tag       = 0
  switch_start_time = "01:00:00"
  switch_end_time   = "02:00:00"
}
