data "tencentcloudenterprise_cvm_images" "my_favorate_image" {
  image_type = ["PUBLIC_IMAGE"]
  os_name    = "centos"
}

data "tencentcloudenterprise_cvm_instance_types" "my_favorate_instance_types" {
  filter {
    name   = "instance-family"
    values = ["S2"]
  }

  cpu_core_count = 1
  memory_size    = 2
}

data "tencentcloudenterprise_availability_zones" "my_favorate_zones" {}

resource "tencentcloudenterprise_cvm_instance" "my_instance" {
  instance_name     = "terraform_automation_test_kuruk"
  availability_zone = data.cloud_availability_zones.my_favorate_zones.zones.0.name
  image_id          = data.cloud_cvm_images.my_favorate_image.images.0.image_id
  instance_type     = data.cloud_cvm_instance_types.my_favorate_instance_types.instance_types.0.instance_type

  system_disk_type = "CLOUD_PREMIUM"

  data_disks {
    data_disk_type = "CLOUD_PREMIUM"
    data_disk_size = 70
  }

  disable_security_service = true
  disable_monitor_service  = true
}

resource "tencentcloudenterprise_eip_instance" "my_eip" {
  name = "tf_auto_test"
}

resource "tencentcloudenterprise_eip_association" "foo" {
  eip_id      = cloud_eip_instance.my_eip.id
  instance_id = cloud_cvm_instance.my_instance.id
}
