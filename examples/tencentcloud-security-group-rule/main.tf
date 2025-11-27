resource "tencentcloudenterprise_vpc_security_group" "default" {
  name        = var.security_group_name
  description = "New security group"
}

resource "tencentcloudenterprise_vpc_security_group" "default2" {
  name        = var.security_group_name
  description = "Anthor security group"
}

resource "tencentcloudenterprise_security_group_rule" "http-in" {
  security_group_id = cloud_vpc_security_group.default.id
  type              = "ingress"
  cidr_ip           = "0.0.0.0/0"
  ip_protocol       = "tcp"
  port_range        = "80,8080"
  policy            = "accept"
}

resource "tencentcloudenterprise_security_group_rule" "ssh-in" {
  security_group_id = cloud_vpc_security_group.default.id
  type              = "ingress"
  cidr_ip           = "0.0.0.0/0"
  ip_protocol       = "tcp"
  port_range        = "22"
  policy            = "accept"
}

resource "tencentcloudenterprise_security_group_rule" "egress-drop" {
  security_group_id = cloud_vpc_security_group.default.id
  type              = "egress"
  cidr_ip           = "203.0.113.0/24"
  ip_protocol       = "udp"
  port_range        = "3000-4000"
  policy            = "drop"
}

resource "tencentcloudenterprise_security_group_rule" "sourcesgid-in" {
  security_group_id = cloud_vpc_security_group.default.id
  type              = "ingress"
  source_sgid       = cloud_vpc_security_group.default2.id
  ip_protocol       = "tcp"
  port_range        = "80,8080"
  policy            = "accept"
}
