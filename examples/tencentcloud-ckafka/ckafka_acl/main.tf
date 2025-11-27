resource "tencentcloudenterprise_ckafka_acl" "terraform_acl" {
  instance_id = "ckafka-agl7fgoz"
  resource_type = "TOPIC"
  resource_name = "terraform_create_third_times"
  operation_type = "WRITE"
  permission_type = "ALLOW"
  host = "203.0.113.12"
  principal = "*"
}