# Cloud Block Storage (CBS) Examples

# ========== Data Sources ==========

# Query CBS storages
data "tencentcloudenterprise_cbs_storages" "storages" {
  storage_id = "disk-xxxxx"
}

# Query CBS storage set
data "tencentcloudenterprise_cbs_storages_set" "storage_set" {
  disk_ids = ["disk-xxxxx"]
}

# Query CBS snapshots
data "tencentcloudenterprise_cbs_snapshots" "snapshots" {
  snapshot_id = "snap-xxxxx"
}

# Query CBS snapshot policies
data "tencentcloudenterprise_cbs_snapshot_policies" "policies" {
  snapshot_policy_id = "asp-xxxxx"
}

# ========== Resources ==========

# CBS Storage
resource "tencentcloudenterprise_cbs_storage" "storage" {
  storage_name      = "example-cbs"
  storage_type      = "CLOUD_PREMIUM"
  storage_size      = 100
  availability_zone = "ap-guangzhou-3"
  project_id        = 0
  encrypt           = false
  
  tags = {
    env = "test"
  }
}

# CBS Storage Set (for batch creation)
resource "tencentcloudenterprise_cbs_storage_set" "storage_set" {
  disk_count        = 3
  disk_type         = "CLOUD_PREMIUM"
  disk_size         = 100
  availability_zone = "ap-guangzhou-3"
  disk_name         = "example-disk"
  project_id        = 0
}

# CBS Storage Attachment
resource "tencentcloudenterprise_cbs_storage_attachment" "attachment" {
  storage_id  = cloud_cbs_storage.storage.id
  instance_id = "ins-xxxxx"
}

# CBS Storage Set Attachment
resource "tencentcloudenterprise_cbs_storage_set_attachment" "set_attachment" {
  storage_id  = cloud_cbs_storage_set.storage_set.disk_ids[0]
  instance_id = "ins-xxxxx"
}

# CBS Snapshot
resource "tencentcloudenterprise_cbs_snapshot" "snapshot" {
  storage_id    = cloud_cbs_storage.storage.id
  snapshot_name = "example-snapshot"
}

# CBS Snapshot Policy
resource "tencentcloudenterprise_cbs_snapshot_policy" "policy" {
  snapshot_policy_name = "example-policy"
  repeat_weekdays      = [1, 3, 5]
  repeat_hours         = [1]
  retention_days       = 7
}

# CBS Snapshot Policy Attachment
resource "tencentcloudenterprise_cbs_snapshot_policy_attachment" "attachment" {
  snapshot_policy_id = cloud_cbs_snapshot_policy.policy.id
  storage_id         = cloud_cbs_storage.storage.id
}

# CBS Snapshot Share Permission
resource "tencentcloudenterprise_cbs_snapshot_share_permission" "share" {
  snapshot_id         = cloud_cbs_snapshot.snapshot.id
  account_ids         = ["123456789"]
}
