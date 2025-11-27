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

resource "tencentcloudenterprise_vpc" "my_vpc" {
  cidr_block = "203.0.113.0/24"
  name       = "tf_vpc_test"
}

resource "tencentcloudenterprise_vpc_subnet" "my_subnet" {
  vpc_id = cloud_vpc.my_vpc.id

  //  vpc_id     = "vpc-csybef02"
  availability_zone = data.cloud_availability_zones.my_favorate_zones.zones.0.name
  name              = "tf_test_subnet"
  cidr_block        = "203.0.113.0/28"
}

resource "tencentcloudenterprise_cvm_instance" "instance-vpc-example" {
  instance_name     = var.instance_name
  availability_zone = data.cloud_availability_zones.my_favorate_zones.zones.0.name
  image_id          = data.cloud_cvm_images.my_favorate_image.images.0.image_id
  instance_type     = data.cloud_cvm_instance_types.my_favorate_instance_types.instance_types.0.instance_type
  system_disk_type  = "CLOUD_PREMIUM"

  vpc_id    = cloud_vpc.my_vpc.id
  subnet_id = cloud_vpc_subnet.my_subnet.id

  internet_max_bandwidth_out = 1
}
