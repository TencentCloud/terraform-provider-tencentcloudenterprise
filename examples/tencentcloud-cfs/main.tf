# Cloud File Storage (CFS) Examples

# ========== Data Sources ==========

# Query CFS access groups
data "tencentcloudenterprise_cfs_access_groups" "groups" {
  access_group_id = "pgroup-xxxxx"
}

# Query CFS access rules
data "tencentcloudenterprise_cfs_access_rules" "rules" {
  access_group_id = "pgroup-xxxxx"
}

# Query CFS available zones
data "tencentcloudenterprise_cfs_available_zone" "zones" {}

# Query CFS file systems
data "tencentcloudenterprise_cfs_file_systems" "fs" {
  file_system_id = "cfs-xxxxx"
}

# Query CFS file system clients
data "tencentcloudenterprise_cfs_file_system_clients" "clients" {
  file_system_id = "cfs-xxxxx"
}

# Query CFS mount targets
data "tencentcloudenterprise_cfs_mount_targets" "targets" {
  file_system_id = "cfs-xxxxx"
}

# ========== Resources ==========

# CFS Sign Up Service
resource "tencentcloudenterprise_cfs_sign_up_cfs_service" "signup" {}

# CFS Access Group
resource "tencentcloudenterprise_cfs_access_group" "group" {
  access_group_name = "example-group"
  description       = "Example CFS access group"
}

# CFS Access Rule
resource "tencentcloudenterprise_cfs_access_rule" "rule" {
  access_group_id = cloud_cfs_access_group.group.id
  auth_client_ip  = "10.0.0.0/24"
  priority        = 1
  rw_permission   = "RO"
  user_permission = "no_root_squash"
}

# CFS File System
resource "tencentcloudenterprise_cfs_file_system" "fs" {
  availability_zone = "ap-guangzhou-3"
  name              = "example-cfs"
  protocol          = "NFS"
  storage_type      = "SD"
  vpc_id            = "vpc-xxxxx"
  subnet_id         = "subnet-xxxxx"
  pgroup_id         = cloud_cfs_access_group.group.id
  
  tags = {
    env = "test"
  }
}

# CFS Auto Snapshot Policy
resource "tencentcloudenterprise_cfs_auto_snapshot_policy" "policy" {
  hour            = "2"
  policy_name     = "example-policy"
  day_of_week     = "1,3,5"
  alive_days      = 7
  is_activated    = 1
}

# CFS Auto Snapshot Policy Attachment
resource "tencentcloudenterprise_cfs_auto_snapshot_policy_attachment" "attachment" {
  auto_snapshot_policy_id = cloud_cfs_auto_snapshot_policy.policy.id
  file_system_ids         = [cloud_cfs_file_system.fs.id]
}

# CFS Snapshot
resource "tencentcloudenterprise_cfs_snapshot" "snapshot" {
  file_system_id = cloud_cfs_file_system.fs.id
  snapshot_name  = "example-snapshot"
  tags = {
    type = "manual"
  }
}
