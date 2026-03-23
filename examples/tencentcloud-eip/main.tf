resource "tencentcloudenterprise_eip_instance" "foo" {
  name = "awesome_eip_example"

  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_eips" "tags" {
  tags = tencentcloudenterprise_eip_instance.foo.tags
}