resource "tencentcloudenterprise_ckafka_instance" "ckafka" {
  instance_name = "test_create_ckafka"
  zone_id  = 50010001
  cluster_id = 1
  vpc_id = "vpc-1lubk6tl"
  subnet_id = "subnet-i1ou3a1i"
  msg_retention_time = 20
  band_width = 8
  disk_size = 30
  tag_set = {
    "ckafka": "terraform_create"
  }
}