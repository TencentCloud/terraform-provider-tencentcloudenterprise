# Cloud Virtual Machine (CVM) Examples

# ========== Data Sources ==========

# Query CVM images
data "tencentcloudenterprise_cvm_images" "images" {
  image_type = ["PUBLIC_IMAGE"]
  os_name    = "ubuntu"
}

# Query single CVM image
data "tencentcloudenterprise_cvm_image" "image" {
  image_id = "img-xxxxx"
}

# Query CVM instance types
data "tencentcloudenterprise_cvm_instance_types" "types" {
  availability_zone = "ap-guangzhou-3"
  cpu_core_count    = 2
  memory_size       = 4
}

# Query CVM instances
data "tencentcloudenterprise_cvm_instances" "instances" {
  instance_id = "ins-xxxxx"
}

# Query CVM instance set
data "tencentcloudenterprise_cvm_instances_set" "instance_set" {
  instance_ids = ["ins-xxxxx", "ins-yyyyy"]
}

# Query CVM key pairs
data "tencentcloudenterprise_cvm_key_pairs" "keys" {
  key_id = "skey-xxxxx"
}

# Query CVM placement groups
data "tencentcloudenterprise_cvm_placement_groups" "groups" {
  placement_group_id = "ps-xxxxx"
}

# Query CVM instances modification
data "tencentcloudenterprise_cvm_instances_modification" "modification" {
  instance_ids = ["ins-xxxxx"]
}

# Query CVM instance VNC URL
data "tencentcloudenterprise_cvm_instance_vnc_url" "vnc" {
  instance_id = "ins-xxxxx"
}

# Query disaster recover group quota
data "tencentcloudenterprise_cvm_disaster_recover_group_quota" "quota" {}

# Query image quota
data "tencentcloudenterprise_cvm_image_quota" "quota" {}

# Query image share permission
data "tencentcloudenterprise_cvm_image_share_permission" "share" {
  image_id = "img-xxxxx"
}

# ========== Resources ==========

# CVM Key Pair
resource "tencentcloudenterprise_cvm_key_pair" "key" {
  key_name   = "example-key"
  public_key = file("~/.ssh/id_rsa.pub")
}

# CVM Placement Group
resource "tencentcloudenterprise_cvm_placement_group" "group" {
  name = "example-pg"
  type = "HOST"
}

# CVM Instance
resource "tencentcloudenterprise_cvm_instance" "instance" {
  instance_name              = "example-cvm"
  availability_zone          = "ap-guangzhou-3"
  image_id                   = "img-xxxxx"
  instance_type              = "S5.MEDIUM4"
  system_disk_type           = "CLOUD_PREMIUM"
  system_disk_size           = 50
  allocate_public_ip         = true
  internet_max_bandwidth_out = 10
  instance_charge_type       = "POSTPAID_BY_HOUR"
  vpc_id                     = "vpc-xxxxx"
  subnet_id                  = "subnet-xxxxx"
  key_ids                    = [cloud_cvm_key_pair.key.id]
  placement_group_id         = cloud_cvm_placement_group.group.id
  
  data_disks {
    data_disk_type = "CLOUD_PREMIUM"
    data_disk_size = 100
  }
  
  tags = {
    env = "test"
  }
}

# CVM Instance Set (batch creation)
resource "tencentcloudenterprise_cvm_instance_set" "instance_set" {
  instance_count             = 3
  instance_name              = "example-batch"
  availability_zone          = "ap-guangzhou-3"
  image_id                   = "img-xxxxx"
  instance_type              = "S5.MEDIUM4"
  system_disk_type           = "CLOUD_PREMIUM"
  system_disk_size           = 50
  allocate_public_ip         = true
  internet_max_bandwidth_out = 10
  vpc_id                     = "vpc-xxxxx"
  subnet_id                  = "subnet-xxxxx"
}

# CVM Launch Template
resource "tencentcloudenterprise_cvm_launch_template" "template" {
  launch_template_name = "example-template"
  placement {
    zone       = "ap-guangzhou-3"
    project_id = 0
  }
  image_id      = "img-xxxxx"
  instance_type = "S5.MEDIUM4"
  
  system_disk {
    disk_type = "CLOUD_PREMIUM"
    disk_size = 50
  }
  
  data_disks {
    disk_type = "CLOUD_PREMIUM"
    disk_size = 100
  }
  
  internet_accessible {
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 10
    public_ip_assigned         = true
  }
  
  virtual_private_cloud {
    vpc_id     = "vpc-xxxxx"
    subnet_id  = "subnet-xxxxx"
  }
}

# CVM Image
resource "tencentcloudenterprise_cvm_image" "image" {
  image_name        = "example-image"
  instance_id       = cloud_cvm_instance.instance.id
  image_description = "Example custom image"
  force_poweroff    = true
  sysprep           = false
  
  tags = {
    type = "custom"
  }
}

# CVM Security Group Attachment
resource "tencentcloudenterprise_cvm_security_group_attachment" "sg_attachment" {
  security_group_ids = ["sg-xxxxx"]
  instance_id        = cloud_cvm_instance.instance.id
}

# CVM Reboot Instance
resource "tencentcloudenterprise_cvm_reboot_instance" "reboot" {
  instance_id = cloud_cvm_instance.instance.id
  stop_type   = "SOFT"
}

# CVM Renew Instance
resource "tencentcloudenterprise_cvm_renew_instance" "renew" {
  instance_id    = cloud_cvm_instance.instance.id
  renew_portable_data_disk = true
}

# CVM Sync Image
resource "tencentcloudenterprise_cvm_sync_image" "sync" {
  image_id            = cloud_cvm_image.image.id
  destination_regions = ["ap-shanghai", "ap-beijing"]
}

# CVM Image Share Permission
resource "tencentcloudenterprise_cvm_image_share_permission" "share" {
  image_id    = cloud_cvm_image.image.id
  account_ids = ["123456789"]
}
