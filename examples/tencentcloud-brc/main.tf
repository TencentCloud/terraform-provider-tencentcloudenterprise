# Backup and Recovery (BRC) Examples

# ========== Data Sources ==========

# Query BRC auto backup policies
data "cloud_brc_autobackup_policies" "policies" {
  auto_backup_policy_id = "policy-xxxxx"
}

# Query BRC backups
data "cloud_brc_backups" "backups" {
  backup_id = "backup-xxxxx"
}

# Query BRC resource backup overview
data "cloud_brc_resource_backup_overview" "overview" {
  resource_type = "CVM"
}

# ========== Resources ==========

# BRC Activate Backup Service
resource "cloud_brc_activate_backup_service" "activate" {}

# BRC Auto Backup Policy
resource "cloud_brc_auto_backup_policy" "policy" {
  backup_policy_name = "example-policy"
  backup_method      = "SNAPSHOT"
  
  timing_strategy {
    execute_time = "02:00"
    backup_period = ["Monday", "Wednesday", "Friday"]
  }
  
  retention_strategy {
    retention_days = 7
  }
}

# BRC Auto Backup Policy Binding
resource "cloud_brc_auto_backup_policy_binding" "binding" {
  auto_backup_policy_id = cloud_brc_auto_backup_policy.policy.id
  resource_type         = "CVM"
  resource_ids          = ["ins-xxxxx"]
}

# BRC Backup Group
resource "cloud_brc_backup_group" "group" {
  backup_group_name = "example-group"
  resource_type     = "CVM"
  resource_ids      = ["ins-xxxxx"]
}

# BRC Backup Disk
resource "cloud_brc_backup_disk" "disk" {
  disk_id       = "disk-xxxxx"
  backup_name   = "example-disk-backup"
  backup_method = "SNAPSHOT"
}

# BRC Backup CFS
resource "cloud_brc_backup_cfs" "cfs" {
  file_system_id = "cfs-xxxxx"
  backup_name    = "example-cfs-backup"
}

# BRC Backup Resource
resource "cloud_brc_backup_resource" "resource" {
  resource_type = "CVM"
  resource_id   = "ins-xxxxx"
  backup_name   = "example-backup"
  backup_method = "SNAPSHOT"
}
