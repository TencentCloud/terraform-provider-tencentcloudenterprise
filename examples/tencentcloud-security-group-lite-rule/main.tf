resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"
}

resource "tencentcloudenterprise_security_group_lite_rule" "lite-rule" {
  security_group_id = cloud_vpc_security_group.foo.id

  ingress = [
    "ACCEPT#203.0.113.0/24#80#TCP",
    "DROP#127.0.0.1#80,90#UDP",
    "ACCEPT#0.0.0.0/0#80-90#TCP",
  ]

  egress = [
    "ACCEPT#203.0.113.0/24#ALL#TCP",
    "ACCEPT#127.0.0.0/8#ALL#ICMP",
    "DROP#0.0.0.0/0#ALL#ALL",
  ]
}

resource "tencentcloudenterprise_security_group_lite_rule" "lite-rule-ingress" {
  security_group_id = cloud_vpc_security_group.foo.id

  ingress = [
    "ACCEPT#203.0.113.0/24#80#TCP",
  ]
}

resource "tencentcloudenterprise_security_group_lite_rule" "lite-rule-egress" {
  security_group_id = cloud_vpc_security_group.foo.id

  egress = [
    "ACCEPT#203.0.113.0/24#ALL#TCP",
  ]
}