# -------------------------------------------------------------------
# Variables
# -------------------------------------------------------------------
variable "number" {
  default = "1"
}

variable "count_format" {
  default = "%02d"
}

variable "image_name_regex_centos" {
  default = "CentOS 7.4"
}

variable "terraform-instance-name" {
  default = "terraform-instance"
}

variable "instance_family" {
  default = "S2"
}

variable "cvm_password" {
  default = "Test@12345"
}

variable "internet_charge_type" {
  default = "TRAFFIC_POSTPAID_BY_HOUR"
}

variable "instance_charge_type" {
  default = "POSTPAID_BY_HOUR"
}

variable "internet_max_bandwidth_out" {
  default = 5
}

variable "disk_type" {
  default = "CLOUD_PREMIUM"
}

variable "disk_size" {
  default = "50"
}

variable "internet_service_provider" {
  default = "CTCC"
}

# -------------------------------------------------------------------
# Datasource Query
# -------------------------------------------------------------------
data "cloud_cvm_instance_types" "instance_type" {
  filter {
    name = "instance-family"
    values = [var.instance_family]
  }
  exclude_sold_out = true
}

data "cloud_cvm_images" "foo" {
  instance_type = data.cloud_cvm_instance_types.instance_type.instance_types[0].instance_type
  image_type = ["PUBLIC_IMAGE"]
  image_name_regex = var.image_name_regex_centos
}


# -------------------------------------------------------------------
# VPC & Subnet
# -------------------------------------------------------------------
resource "cloud_vpc" "vpc" {
  cidr_block = "172.16.0.0/16"
  name       = var.terraform-instance-name
}

resource "cloud_vpc_subnet" "subnet" {
  vpc_id            = cloud_vpc.vpc.id
  cidr_block        = "172.16.0.0/24"
  availability_zone = data.cloud_cvm_instance_types.instance_type.instance_types[0].availability_zone
  name              = "${var.terraform-instance-name}-subnet"
  is_multicast      = false
}

# -------------------------------------------------------------------
# Security Group
# -------------------------------------------------------------------
resource "cloud_vpc_security_group" "group" {
  name        = var.terraform-instance-name
  description = "New security group"
  project_id  = 0
}

resource "cloud_vpc_security_group_rule_set" "group_rule" {
  security_group_id = cloud_vpc_security_group.group.id

  ingress {
    cidr_block  = "0.0.0.0/0"
    protocol    = "TCP"
    port        = "22"
    action      = "ACCEPT"
    description = "Allow SSH"
  }

  ingress {
    cidr_block  = "0.0.0.0/0"
    protocol    = "TCP"
    port        = "80"
    action      = "ACCEPT"
    description = "Allow HTTP"
  }

  ingress {
    cidr_block  = "0.0.0.0/0"
    protocol    = "TCP"
    port        = "443"
    action      = "ACCEPT"
    description = "Allow HTTPS"
  }

  ingress {
    cidr_block  = "0.0.0.0/0"
    protocol    = "ICMP"
    action      = "ACCEPT"
    description = "Allow ICMP"
  }

  egress {
    cidr_block  = "0.0.0.0/0"
    action      = "ACCEPT"
    description = "Allow all egress"
  }
}

# -------------------------------------------------------------------
# CBS Data Disk
# -------------------------------------------------------------------
resource "cloud_cbs_storage" "disk" {
  availability_zone = data.cloud_cvm_instance_types.instance_type.instance_types[0].availability_zone
  storage_type      = var.disk_type
  storage_size      = var.disk_size
  storage_name      = "${var.terraform-instance-name}-disk-${format(var.count_format, count.index + 1)}"
  count             = var.number
}
# -------------------------------------------------------------------
# Elastic IP(EIP)
# -------------------------------------------------------------------
resource "cloud_eip" "eip" {
  count                      = var.number
  name                       = "eip-${format(var.count_format, count.index + 1)}"
  internet_charge_type       = var.internet_charge_type
  internet_max_bandwidth_out = var.internet_max_bandwidth_out
  internet_service_provider = var.internet_service_provider
  type                       = "EIP"

  tags = {
    type = "example"
    env  = "test"
  }
}
# -------------------------------------------------------------------
# CVM Instance
# -------------------------------------------------------------------
resource "cloud_cvm_instance" "instance" {
  instance_name   = "${var.terraform-instance-name}-instance-${format(var.count_format, count.index + 1)}"
  hostname        = "${var.terraform-instance-name}-host-${format(var.count_format, count.index + 1)}"
  image_id        = data.cloud_cvm_images.foo.images[0].image_id
  instance_type   = data.cloud_cvm_instance_types.instance_type.instance_types[0].instance_type
  count           = var.number
  availability_zone = data.cloud_cvm_instance_types.instance_type.instance_types[0].availability_zone

  orderly_security_groups = [cloud_vpc_security_group.group.id]
  vpc_id                  = cloud_vpc.vpc.id
  subnet_id               = cloud_vpc_subnet.subnet.id

  password = var.cvm_password

  instance_charge_type    = var.instance_charge_type
  system_disk_type        = var.disk_type
  system_disk_size        = 50

  tags = {
    type = "example"
    env  = "test"
  }
}
#
# -------------------------------------------------------------------
# Attach EIP to CVM
# -------------------------------------------------------------------
resource "cloud_eip_association" "eip_attachment" {
  count       = var.number
  eip_id      = cloud_eip.eip[count.index].id
  instance_id = cloud_cvm_instance.instance[count.index].id
}

# -------------------------------------------------------------------
# Attach Disk to CVM
# -------------------------------------------------------------------
resource "cloud_cbs_storage_attachment" "instance_attachment" {
  count       = var.number
  storage_id  = cloud_cbs_storage.disk.*.id[count.index]
  instance_id = cloud_cvm_instance.instance.*.id[count.index]
}

# -------------------------------------------------------------------
# Outputs
# -------------------------------------------------------------------
output "hostname_list" {
  value = join(",", cloud_cvm_instance.instance.*.instance_name)
}

output "cvm_ids" {
  value = join(",", cloud_cvm_instance.instance.*.id)
}

output "cvm_public_ip" {
  value = join(",", cloud_eip.eip.*.public_ip)
}

output "eip_ids" {
  value = join(",", cloud_eip.eip.*.id)
}

output "cvm_private_ip" {
  value = join(",", cloud_cvm_instance.instance.*.private_ip)
}

output "tags" {
  value = jsonencode(cloud_cvm_instance.instance.*.tags)
}
