data "tencentcloudenterprise_ckafka_acls" "acl_test" {
  instance_id = "ckafka-agl7fgoz"
  resource_type = "TOPIC"
  resource_name = "terraform_create_third_times"
  host = "203.0.113.15"
}