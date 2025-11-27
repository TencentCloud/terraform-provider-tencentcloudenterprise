data "tencentcloudenterprise_cvm_images" "my_favorate_image" {
  image_type = ["PUBLIC_IMAGE"]
  os_name    = "centos"
}

data "tencentcloudenterprise_cvm_instance_types" "my_favorate_instance_types" {
  filter {
    name   = "instance-family"
    values = ["S1"]
  }

  cpu_core_count = 1
  memory_size    = 1
}

data "tencentcloudenterprise_availability_zones" "my_favorate_zones" {}

resource "tencentcloudenterprise_vpc_security_group" "my_sg" {
  name        = "tf_test_sg_name"
  description = "tf_test_sg_desc"
}

resource "tencentcloudenterprise_security_group_rule" "web" {
  security_group_id = cloud_vpc_security_group.my_sg.id
  type              = "ingress"
  cidr_ip           = "203.0.113.225/32"
  ip_protocol       = "tcp"
  port_range        = "80,3000,8080"
  policy            = "accept"
}

resource "tencentcloudenterprise_security_group_rule" "login" {
  security_group_id = cloud_vpc_security_group.my_sg.id
  type              = "ingress"
  cidr_ip           = "203.0.113.93/32"
  ip_protocol       = "tcp"
  port_range        = "22"
  policy            = "accept"
}

resource "tencentcloudenterprise_vpc_security_group" "my_sg2" {
  name        = "tf_test_sg_name2"
  description = "tf_test_sg_desc2"
}

resource "tencentcloudenterprise_security_group_rule" "qortex" {
  security_group_id = cloud_vpc_security_group.my_sg2.id
  type              = "ingress"
  cidr_ip           = "203.0.113.93/32"
  ip_protocol       = "tcp"
  port_range        = "5000"
  policy            = "accept"
}

resource "tencentcloudenterprise_cvm_instance" "instance-without-specified-image-id-example" {
  instance_name     = var.instance_name
  availability_zone = data.cloud_availability_zones.my_favorate_zones.zones.0.name
  image_id          = data.cloud_cvm_images.my_favorate_image.images.0.image_id
  instance_type     = data.cloud_cvm_instance_types.my_favorate_instance_types.instance_types.0.instance_type
  password          = "test1234"
  system_disk_type  = "CLOUD_PREMIUM"

  security_groups = [
    cloud_vpc_security_group.my_sg.id,
    cloud_vpc_security_group.my_sg2.id,
  ]

  internet_max_bandwidth_out = 2
  count                      = 1
}
