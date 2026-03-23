resource "tencentcloudenterprise_vpc" "foo" {
  name       = "ci-test-eni-vpc"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "foo" {
  availability_zone = var.availability_zone
  name              = "ci-test-eni-subnet"
  vpc_id            = tencentcloudenterprise_vpc.foo.id
  cidr_block        = "10.0.0.0/16"
  is_multicast      = false
}

resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "test-ci-eni-sg1"
}

resource "tencentcloudenterprise_vpc_security_group" "bar" {
  name = "test-ci-eni-sg2"
}

resource "tencentcloudenterprise_eni" "foo" {
  name            = "ci-test-eni"
  vpc_id          = tencentcloudenterprise_vpc.foo.id
  subnet_id       = tencentcloudenterprise_vpc_subnet.foo.id
  description     = "eni desc"
  security_groups = [tencentcloudenterprise_vpc_security_group.foo.id, tencentcloudenterprise_vpc_security_group.bar.id]

  ipv4s {
    ip          = "10.0.0.10"
    primary     = true
    description = "new desc"
  }

  ipv4s {
    ip      = "10.0.0.11"
    primary = false
  }

  ipv4s {
    ip      = "10.0.0.12"
    primary = false
  }

  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_cvm_images" "my_favorate_image" {
  image_type = ["PUBLIC_IMAGE"]
  os_name    = "centos"
}

data "tencentcloudenterprise_cvm_instance_types" "my_favorite_instance_types" {
  filter {
    name   = "instance-family"
    values = ["S3"]
  }

  cpu_core_count = 4
  memory_size    = 8
}

resource "tencentcloudenterprise_cvm_instance" "foo" {
  instance_name            = "ci-test-eni-attach"
  availability_zone        = "ap-guangzhou-3"
  image_id                 = data.tencentcloudenterprise_cvm_images.my_favorate_image.images.0.image_id
  instance_type            = data.tencentcloudenterprise_cvm_instance_types.my_favorite_instance_types.instance_types.0.instance_type
  system_disk_type         = "CLOUD_PREMIUM"
  disable_security_service = true
  disable_monitor_service  = true
  vpc_id                   = tencentcloudenterprise_vpc.foo.id
  subnet_id                = tencentcloudenterprise_vpc_subnet.foo.id
}

resource "tencentcloudenterprise_eni_attachment" "foo" {
  eni_id      = tencentcloudenterprise_eni.foo.id
  instance_id = tencentcloudenterprise_cvm_instance.foo.id
}

data "tencentcloudenterprise_enis" "subnet" {
  subnet_id      = tencentcloudenterprise_eni.foo.subnet_id
  security_group = tencentcloudenterprise_vpc_security_group.foo.id
}
