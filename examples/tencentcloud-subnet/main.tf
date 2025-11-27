resource "tencentcloudenterprise_vpc" "test_vpc" {
  name       = "Used for testing subnets"
  cidr_block = "203.0.113.0/24"
}

resource "tencentcloudenterprise_route_table" "foo" {
  vpc_id = cloud_vpc.test_vpc.id
  name   = "ci-temp-test-rt"
}

resource "tencentcloudenterprise_vpc_subnet" "test_subnet" {
  vpc_id            = cloud_vpc.test_vpc.id
  name              = "terraform test subnet"
  cidr_block        = "203.0.113.0/28"
  availability_zone = var.availability_zone
  route_table_id    = cloud_route_table.foo.id
}
