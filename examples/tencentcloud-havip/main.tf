# Create VPC and Subnet
resource "tencentcloudenterprise_vpc" "example" {
  name       = "example"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "example" {
  name              = "example"
  availability_zone = var.availability_zone
  vpc_id            = tencentcloudenterprise_vpc.example.id
  cidr_block        = "10.0.0.0/24"
  is_multicast      = false
}

resource "tencentcloudenterprise_ha_vip" "example" {
  name      = "example"
  vpc_id    = tencentcloudenterprise_vpc.example.id
  subnet_id = tencentcloudenterprise_vpc_subnet.example.id
  vip       = "10.0.20.5"

}
resource "tencentcloudenterprise_ha_vip_eip_attachment" "example" {
  havip_id   = tencentcloudenterprise_ha_vip.example.id
  address_ip = tencentcloudenterprise_eip_instance.example.public_ip
}

data "tencentcloudenterprise_ha_vips" "example" {
  id = tencentcloudenterprise_ha_vip.example.id
}
data "tencentcloudenterprise_ha_vip_eip_attachments" "example" {
  havip_id = tencentcloudenterprise_ha_vip_eip_attachment.example.havip_id
}
