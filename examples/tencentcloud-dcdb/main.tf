resource "tencentcloudenterprise_dcdb_instance" "hourdb_instance" {
  instance_name    = "111"
  zones            = ["yfm18", "yfm18"]
  cpu_arch         = "X86"
  ipv6_flag        = 0
  shard_memory     = 2
  shard_storage    = 10
  shard_node_count = 2
  shard_cpu        = 1
  shard_count      = 2
  vpc_id           = "vpc-cs6ffr73"
  subnet_id        = "subnet-mfbxe9zk"
  db_version_id    = "5.7"
  project_id       = "pr-bae40f73"

  resource_tags {
    tag_key   = "createdBy"
    tag_value = "terraform3"
  }

  init_params {

  }
}
