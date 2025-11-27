data "tencentcloudenterprise_cvm_images" "my_favorite_image" {
  image_type = ["PUBLIC_IMAGE"]
  os_name    = "centos"
}

data "tencentcloudenterprise_cvm_instance_types" "my_favorite_instance_types" {
  filter {
    name   = "instance-family"
    values = ["S3"]
  }

  cpu_core_count = 1
  memory_size    = 1
}

data "tencentcloudenterprise_availability_zones" "my_favorite_zones" {}

// Create VPC resource
resource "tencentcloudenterprise_vpc" "app" {
  cidr_block = "203.0.113.0/24"
  name       = "awesome_app_vpc"
}

resource "tencentcloudenterprise_vpc_subnet" "app" {
  vpc_id            = cloud_vpc.app.id
  availability_zone = data.cloud_availability_zones.my_favorite_zones.zones.0.name
  name              = "awesome_app_subnet"
  cidr_block        = "203.0.113.0/28"
}

// Create 2 CVM instances to host awesome_app
resource "tencentcloudenterprise_cvm_instance" "my_awesome_app" {
  instance_name     = "awesome_app"
  availability_zone = data.cloud_availability_zones.my_favorite_zones.zones.0.name
  image_id          = data.cloud_cvm_images.my_favorite_image.images.0.image_id
  instance_type     = data.cloud_cvm_instance_types.my_favorite_instance_types.instance_types.0.instance_type
  system_disk_type  = "CLOUD_PREMIUM"
  system_disk_size  = 60
  hostname          = "user"
  vpc_id            = cloud_vpc.app.id
  subnet_id         = cloud_vpc_subnet.app.id
  count             = 2

  data_disks {
    data_disk_type = "CLOUD_PREMIUM"
    data_disk_size = 60
    # encrypt        = false
  }

  # tags = {
  #   tagKey = "tagValue"
  # }
}

