resource "tencentcloudenterprise_vpcdns_forward_rule" "foo" {
  remark     = "forward_rule_foo"
  domain_id = "my_domain_id1"
  forward_address = ["8.8.8.8:88", "1.1.1.1:88"]
}