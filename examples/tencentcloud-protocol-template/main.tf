resource "tencentcloudenterprise_protocol_template" "example" {
  name      = "example"
  protocols = ["udp:all", "tcp:90,110", "icmp", "tcp:1110-1120"]
}

resource "tencentcloudenterprise_protocol_template_group" "example" {
  name         = "example"
  template_ids = [tencentcloudenterprise_protocol_template.example.id]
}

data "tencentcloudenterprise_protocol_templates" "example" {
  id = tencentcloudenterprise_protocol_template.example.id
}

data "tencentcloudenterprise_protocol_template_groups" "example" {
}
