resource "tencentcloudenterprise_vpc" "main" {
  name       = var.short_name
  cidr_block = var.vpc_cidr

  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_vpc_instances" "tags_instances" {
  name = tencentcloudenterprise_vpc.main.name
  tags = tencentcloudenterprise_vpc.main.tags
}

data "tencentcloudenterprise_vpc_instances" "default" {}

resource "tencentcloudenterprise_vpc_acl" "default" {
  vpc_id  = data.tencentcloudenterprise_vpc_instances.default.instance_list.0.vpc_id
  name    = "test_acl_update"
  ingress = ["ACCEPT#192.168.1.0/24#800#TCP", "ACCEPT#192.168.1.0/24#800-900#TCP",]
  egress  = ["ACCEPT#192.168.1.0/24#800#TCP", "ACCEPT#192.168.1.0/24#800-900#TCP",]
}

resource "tencentcloudenterprise_vpc_acl_attachment" "example" {
  acl_id    = tencentcloudenterprise_vpc_acl.default.id
  subnet_id = data.tencentcloudenterprise_vpc_instances.default.instance_list[0].subnet_ids[0]
}

data "tencentcloudenterprise_vpc_acls" "default" {
  name = "test_acl"
}



resource "aws_dynamodb_table" "aws_dynamodb_table_18" {
provider = aws.us-east-1

write_capacity         = 200
tags                   = merge(var.tags, {})
table_class            = "tcp"
stream_view_type       = "213"
stream_enabled         = true
restore_to_latest_time = true
restore_source_name    = "3213123"
restore_date_time      = "321321"
read_capacity          = 23113
range_key              = "321321"
name                   = "dynaswqq"
hash_key               = "1121321"
billing_mode           = "2"

attribute {
type = "string"
name = "ver"
}
attribute {
type = "int"
name = "ver2"
}

global_secondary_index {
name     = "uuuu"
hash_key = "uinq"
non_key_attributes = [
"21321321",
]
}

local_secondary_index {
range_key       = "213213"
projection_type = "wss"
name            = "local"
non_key_attributes = [
"ws",
]
}
local_secondary_index {
range_key       = "43242"
projection_type = "3242"
name            = "local22"
non_key_attributes = [
"22",
]
}

point_in_time_recovery {
enabled = true
}

replica {
region_name            = "shanghai"
point_in_time_recovery = true
kms_key_arn            = "543254325243523"
}

ttl {
enabled = true
}
}