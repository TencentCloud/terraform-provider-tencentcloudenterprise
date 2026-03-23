resource "tencentcloudenterprise_vpc" "my_vpc" {
  name       = "Used to test rtb"
  cidr_block = "10.2.0.0/16"
}

resource "tencentcloudenterprise_route_table" "my_rtb" {
  vpc_id = tencentcloudenterprise_vpc.my_vpc.id
  name   = var.short_name

  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_vpc_route_tables" "tags_instances" {
  tags = tencentcloudenterprise_route_table.my_rtb.tags
}
