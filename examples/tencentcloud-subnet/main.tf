resource "tencentcloudenterprise_vpc" "test_vpc" {
  name       = "Used for testing subnets"
  cidr_block = "10.1.0.0/16"
}

resource "tencentcloudenterprise_route_table" "foo" {
  vpc_id = tencentcloudenterprise_vpc.test_vpc.id
  name   = "ci-temp-test-rt"
}

resource "tencentcloudenterprise_vpc_subnet" "test_subnet" {
  vpc_id            = tencentcloudenterprise_vpc.test_vpc.id
  name              = "terraform test subnet"
  cidr_block        = "10.1.1.0/24"
  availability_zone = var.availability_zone
  route_table_id    = tencentcloudenterprise_route_table.foo.id
}
