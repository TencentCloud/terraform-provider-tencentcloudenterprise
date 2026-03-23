resource "tencentcloudenterprise_cbs_storage" "my_storage" {
  storage_type      = var.storage_type
  storage_name      = "tf-test-storage"
  storage_size      = 60
  availability_zone = var.availability_zone
  project_id        = 0
  encrypt           = false

  tags = {
    test = "tf"
  }
}

data "tencentcloudenterprise_cvm_images" "my_favorate_image" {
  image_type = ["PUBLIC_IMAGE"]
  os_name    = "centos"
}

resource "tencentcloudenterprise_cvm_instance" "my_instance" {
  instance_name     = "tf-test-instance"
  availability_zone = var.availability_zone
  image_id          = data.tencentcloudenterprise_cvm_images.my_favorate_image.images.0.image_id
  instance_type     = var.instance_type
  system_disk_type  = var.storage_type
}

resource "tencentcloudenterprise_cbs_storage_attachment" "my_attachment" {
  storage_id  = tencentcloudenterprise_cbs_storage.my_storage.id
  instance_id = tencentcloudenterprise_cvm_instance.my_instance.id
}

resource "tencentcloudenterprise_cbs_snapshot" "my_snapshot" {
  storage_id    = tencentcloudenterprise_cbs_storage.my_storage.id
  snapshot_name = "tf-test-snapshot"
}

resource "tencentcloudenterprise_cbs_snapshot_policy" "snapshot_policy" {
  snapshot_policy_name = "tf-test-snapshot-policy"
  repeat_weekdays      = [1, 4]
  repeat_hours         = [1]
  retention_days       = 7
}

data "tencentcloudenterprise_cbs_storages" "storages" {
  storage_id = tencentcloudenterprise_cbs_storage.my_storage.id
}

data "tencentcloudenterprise_cbs_snapshots" "snapshots" {
  snapshot_id = tencentcloudenterprise_cbs_snapshot.my_snapshot.id
}
