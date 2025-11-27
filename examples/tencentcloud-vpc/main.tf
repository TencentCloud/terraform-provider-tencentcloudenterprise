# Virtual Private Cloud (VPC) Examples

# ========== Data Sources ==========

# Query VPC instances
data "tencentcloudenterprise_vpc_instances" "vpcs" {
  vpc_id = "vpc-xxxxx"
}

# Query VPC subnets
data "tencentcloudenterprise_vpc_subnets" "subnets" {
  vpc_id    = "vpc-xxxxx"
  subnet_id = "subnet-xxxxx"
}

# Query VPC security groups
data "tencentcloudenterprise_vpc_security_groups" "sgs" {
  security_group_id = "sg-xxxxx"
}

# Query VPC ACLs
data "tencentcloudenterprise_vpc_acls" "acls" {
  vpc_id = "vpc-xxxxx"
}

# Query VPC route tables
data "tencentcloudenterprise_vpc_route_tables" "route_tables" {
  vpc_id = "vpc-xxxxx"
}

# Query VPC ENIs
data "tencentcloudenterprise_vpc_enis" "enis" {
  vpc_id = "vpc-xxxxx"
}

# Query VPC NAT gateways
data "tencentcloudenterprise_vpc_nat_gateways" "nats" {
  vpc_id = "vpc-xxxxx"
}

# Query VPC HA VIPs
data "tencentcloudenterprise_vpc_ha_vips" "vips" {
  vpc_id = "vpc-xxxxx"
}

# Query VPC DNATs
data "tencentcloudenterprise_vpc_dnats" "dnats" {
  vpc_id     = "vpc-xxxxx"
  nat_id     = "nat-xxxxx"
}

# Query VPC address templates
data "tencentcloudenterprise_vpc_address_templates" "templates" {
  name = "example-template"
}

# Query VPC address template groups
data "tencentcloudenterprise_vpc_address_template_groups" "groups" {
  name = "example-group"
}

# ========== Resources ==========

# VPC
resource "tencentcloudenterprise_vpc" "vpc" {
  name         = "example-vpc"
  cidr_block   = "10.0.0.0/16"
  dns_servers  = ["183.60.83.19", "183.60.82.98"]
  is_multicast = false
  
  tags = {
    env = "test"
  }
}

# VPC Subnet
resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  name              = "example-subnet"
  vpc_id            = cloud_vpc.vpc.id
  availability_zone = "ap-guangzhou-3"
  cidr_block        = "10.0.1.0/24"
  is_multicast      = false
  
  tags = {
    env = "test"
  }
}

# VPC Security Group
resource "tencentcloudenterprise_vpc_security_group" "sg" {
  name        = "example-sg"
  description = "Example security group"
  project_id  = 0
  
  tags = {
    env = "test"
  }
}

# VPC Security Group Rule
resource "tencentcloudenterprise_vpc_security_group_rule" "rule" {
  security_group_id = cloud_vpc_security_group.sg.id
  type              = "ingress"
  cidr_ip           = "0.0.0.0/0"
  ip_protocol       = "TCP"
  port_range        = "80,443"
  policy            = "ACCEPT"
  description       = "Allow HTTP/HTTPS"
}

# VPC Security Group Rule Set (batch rules)
resource "tencentcloudenterprise_vpc_security_group_rule_set" "rule_set" {
  security_group_id = cloud_vpc_security_group.sg.id
  
  ingress {
    action      = "ACCEPT"
    cidr_block  = "10.0.0.0/16"
    protocol    = "TCP"
    port        = "22"
    description = "Allow SSH from VPC"
  }
  
  ingress {
    action      = "ACCEPT"
    cidr_block  = "0.0.0.0/0"
    protocol    = "ICMP"
    description = "Allow ICMP"
  }
  
  egress {
    action     = "ACCEPT"
    cidr_block = "0.0.0.0/0"
    protocol   = "ALL"
    description = "Allow all outbound"
  }
}

# VPC Security Group Lite Rule
resource "tencentcloudenterprise_vpc_security_group_lite_rule" "lite_rule" {
  security_group_id = cloud_vpc_security_group.sg.id
  
  ingress = [
    "ACCEPT#0.0.0.0/0#80#TCP",
    "ACCEPT#0.0.0.0/0#443#TCP",
  ]
  
  egress = [
    "ACCEPT#0.0.0.0/0#ALL#ALL",
  ]
}

# VPC ACL
resource "tencentcloudenterprise_vpc_acl" "acl" {
  vpc_id  = cloud_vpc.vpc.id
  name    = "example-acl"
  
  ingress = [
    "ACCEPT#0.0.0.0/0#80#TCP",
    "ACCEPT#0.0.0.0/0#443#TCP",
  ]
  
  egress = [
    "ACCEPT#0.0.0.0/0#ALL#ALL",
  ]
}

# VPC ACL Attachment
resource "tencentcloudenterprise_vpc_acl_attachment" "acl_attachment" {
  acl_id    = cloud_vpc_acl.acl.id
  subnet_id = cloud_vpc_subnet.subnet.id
}

# VPC Route Table
resource "tencentcloudenterprise_vpc_route_table" "route_table" {
  vpc_id = cloud_vpc.vpc.id
  name   = "example-route-table"
  
  tags = {
    env = "test"
  }
}

# VPC Route Table Entry
resource "tencentcloudenterprise_vpc_route_table_entry" "route" {
  route_table_id         = cloud_vpc_route_table.route_table.id
  destination_cidr_block = "192.168.0.0/16"
  next_type              = "VPN"
  next_hub               = "vpngw-xxxxx"
  description            = "Route to VPN"
}

# VPC NAT Gateway
resource "tencentcloudenterprise_vpc_nat_gateway" "nat" {
  name             = "example-nat"
  vpc_id           = cloud_vpc.vpc.id
  bandwidth        = 100
  max_concurrent   = 1000000
  assigned_eip_set = ["eip-xxxxx"]
  
  tags = {
    env = "test"
  }
}

# VPC DNAT
resource "tencentcloudenterprise_vpc_dnat" "dnat" {
  vpc_id       = cloud_vpc.vpc.id
  nat_id       = cloud_vpc_nat_gateway.nat.id
  protocol     = "TCP"
  elastic_ip   = "1.1.1.1"
  elastic_port = "80"
  private_ip   = "10.0.1.10"
  private_port = "8080"
  description  = "DNAT for web service"
}

# VPC ENI (Elastic Network Interface)
resource "tencentcloudenterprise_vpc_eni" "eni" {
  name        = "example-eni"
  vpc_id      = cloud_vpc.vpc.id
  subnet_id   = cloud_vpc_subnet.subnet.id
  description = "Example ENI"
  
  ipv4_count = 1
  
  security_groups = [cloud_vpc_security_group.sg.id]
  
  tags = {
    env = "test"
  }
}

# VPC ENI Attachment
resource "tencentcloudenterprise_vpc_eni_attachment" "attachment" {
  eni_id      = cloud_vpc_eni.eni.id
  instance_id = "ins-xxxxx"
}

# VPC ENI Security Group Attachment
resource "tencentcloudenterprise_vpc_eni_sg_attachment" "sg_attachment" {
  network_interface_ids = [cloud_vpc_eni.eni.id]
  security_group_ids    = [cloud_vpc_security_group.sg.id]
}

# VPC HA VIP (High Availability Virtual IP)
resource "tencentcloudenterprise_vpc_ha_vip" "vip" {
  name      = "example-vip"
  vpc_id    = cloud_vpc.vpc.id
  subnet_id = cloud_vpc_subnet.subnet.id
  vip       = "10.0.1.100"
}

# VPC HA VIP EIP Attachment
resource "tencentcloudenterprise_vpc_ha_vip_eip_attachment" "vip_eip" {
  ha_vip_id      = cloud_vpc_ha_vip.vip.id
  address_ip     = "eip-xxxxx"
}

# VPC Address Template
resource "tencentcloudenterprise_vpc_address_template" "template" {
  name      = "example-template"
  addresses = ["10.0.0.0/16", "192.168.0.0/16"]
}

# VPC Address Template Group
resource "tencentcloudenterprise_vpc_address_template_group" "group" {
  name         = "example-group"
  template_ids = [cloud_vpc_address_template.template.id]
}

# VPC IPv6 CIDR Block
resource "tencentcloudenterprise_vpc_ipv6_cidr_block" "ipv6" {
  vpc_id = cloud_vpc.vpc.id
}

# VPC IPv6 Subnet CIDR Block
resource "tencentcloudenterprise_vpc_ipv6_subnet_cidr_block" "ipv6_subnet" {
  vpc_id            = cloud_vpc.vpc.id
  ipv6_cidr_blocks {
    subnet_id         = cloud_vpc_subnet.subnet.id
    ipv6_cidr_block   = "2402:4e00:1000:4200::/64"
  }
}

# VPC IPv6 ENI Address
resource "tencentcloudenterprise_vpc_ipv6_eni_address" "ipv6_eni" {
  network_interface_id = cloud_vpc_eni.eni.id
  ipv6_addresses {
    address     = "2402:4e00:1000:4200::1"
    primary     = true
    description = "Primary IPv6"
  }
}

# VPC IPv6 Address Bandwidth
resource "tencentcloudenterprise_vpc_ipv6_address_bandwidth" "ipv6_bw" {
  ipv6_address         = "2402:4e00:1000:4200::1"
  internet_charge_type = "TRAFFIC_POSTPAID_BY_HOUR"
  internet_max_bandwidth_out = 10
}

# VPC DC Gateway (Direct Connect Gateway)
resource "tencentcloudenterprise_vpc_dc_gateway" "dc_gw" {
  name                = "example-dc-gateway"
  vpc_id              = cloud_vpc.vpc.id
  network_type        = "VPC"
  network_instance_id = cloud_vpc.vpc.id
}

# VPC Net Detect
resource "tencentcloudenterprise_vpc_net_detect" "detect" {
  vpc_id          = cloud_vpc.vpc.id
  subnet_id       = cloud_vpc_subnet.subnet.id
  net_detect_name = "example-detect"
  detect_destination_ip = ["10.0.2.1"]
  next_hop_type   = "NORMAL_CVM"
  next_hop_destination = "10.0.1.10"
}

# VPC Bandwidth Package
resource "tencentcloudenterprise_vpc_bandwidth_package" "bw_pkg" {
  network_type = "BGP"
  charge_type  = "TOP5_POSTPAID_BY_MONTH"
  bandwidth_package_name = "example-bw-package"
  
  tags = {
    env = "test"
  }
}

# VPC Bandwidth Package Attachment
resource "tencentcloudenterprise_vpc_bandwidth_package_attachment" "bw_attachment" {
  resource_id          = "eip-xxxxx"
  bandwidth_package_id = cloud_vpc_bandwidth_package.bw_pkg.id
  resource_type        = "Address"
}
