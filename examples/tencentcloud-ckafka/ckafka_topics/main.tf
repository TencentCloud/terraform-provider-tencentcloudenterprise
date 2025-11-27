data "tencentcloudenterprise_ckafka_topics" "test_topics" {
  instance_id = "ckafka-agl7fgoz"
  topic_name = "terraform_create_third_times"
}