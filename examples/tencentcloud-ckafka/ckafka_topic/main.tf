resource "tencentcloudenterprise_ckafka_topic" "create_topic" {
  instance_id = "ckafka-agl7fgoz"
  topic_name  = "terraform_create"
  partition_num = 1
  replica_num = 2
  note  = "terraform create topic test"
  retention = 86400000
  sync_replica_min_num = 2
  clean_up_policy = "delete"
  unclean_leader_election_enable = true
  max_message_bytes = 1024

}