# TurboFS (High Performance File Storage) Examples

# ========== Data Sources ==========

# Query TurboFS permission groups
data "tencentcloudenterprise_turbofs_p_groups" "groups" {
  p_group_id = "pgroup-xxxxx"
}

# Query TurboFS rules
data "tencentcloudenterprise_turbofs_rules" "rules" {
  p_group_id = "pgroup-xxxxx"
}

# Query TurboFS file systems
data "tencentcloudenterprise_turbofs_file_systems" "fs" {
  file_system_id = "turbofs-xxxxx"
}

# Query TurboFS mount targets
data "tencentcloudenterprise_turbofs_mount_targets" "targets" {
  file_system_id = "turbofs-xxxxx"
}

# ========== Resources ==========

# TurboFS Sign Up Service
resource "tencentcloudenterprise_turbofs_sign_up_service" "signup" {}

# TurboFS Permission Group
resource "tencentcloudenterprise_turbofs_p_group" "group" {
  p_group_name = "example-pgroup"
  description  = "Example TurboFS permission group"
}

# TurboFS Rule
resource "tencentcloudenterprise_turbofs_rule" "rule" {
  p_group_id      = cloud_turbofs_p_group.group.id
  auth_client_ip  = "10.0.0.0/24"
  priority        = 1
  rw_permission   = "RO"
  user_permission = "no_root_squash"
}

# TurboFS File System
resource "tencentcloudenterprise_turbofs_file_system" "fs" {
  availability_zone = "ap-guangzhou-3"
  capacity          = 20480
  name              = "example-turbofs"
  protocol          = "NFS"
  storage_type      = "TB"
  vpc_id            = "vpc-xxxxx"
  subnet_id         = "subnet-xxxxx"
  p_group_id        = cloud_turbofs_p_group.group.id
  
  tags = {
    env = "production"
  }
}

# TurboFS Auto Snapshot Policy
resource "tencentcloudenterprise_turbofs_auto_snapshot_policy" "policy" {
  hour         = "2"
  policy_name  = "example-policy"
  day_of_week  = "1,3,5"
  alive_days   = 7
  is_activated = 1
}

# TurboFS Auto Snapshot Policy Attachment
resource "tencentcloudenterprise_turbofs_auto_snapshot_policy_attachment" "attachment" {
  auto_snapshot_policy_id = cloud_turbofs_auto_snapshot_policy.policy.id
  file_system_ids         = [cloud_turbofs_file_system.fs.id]
}
