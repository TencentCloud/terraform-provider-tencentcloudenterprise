resource "tencentcloudenterprise_cbs_snapshot_policy" "snapshot_policy" {
  snapshot_policy_name = "mysnapshotpolicyname"
  repeat_weekdays      = [1, 4]
  repeat_hours         = [1]
  retention_days       = 7
}

resource "tencentcloudenterprise_cbs_snapshot_policy_attachment" "foo" {
  storage_id         = "disk-q5ka14o6"
  snapshot_policy_id = cloud_cbs_snapshot_policy.snapshot_policy.id
}

resource "tencentcloudenterprise_cbs_snapshot_share_permission" "snapshot_share_permission" {
  account_ids = "100004603073"
  snapshot_id = "snap-7o0uk1pt"
}