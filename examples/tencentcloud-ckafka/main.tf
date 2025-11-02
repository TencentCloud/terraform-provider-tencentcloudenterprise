# Cloud Kafka (CKafka) Examples

# ========== Data Sources ==========

# Query CKafka instances
data "cloud_ckafka_instances" "instances" {
  instance_id = "ckafka-xxxxx"
}

# Query CKafka topics
data "cloud_ckafka_topics" "topics" {
  instance_id = "ckafka-xxxxx"
  topic_name  = "example-topic"
}

# Query CKafka users
data "cloud_ckafka_users" "users" {
  instance_id = "ckafka-xxxxx"
}

# Query CKafka ACLs
data "cloud_ckafka_acls" "acls" {
  instance_id   = "ckafka-xxxxx"
  resource_type = "TOPIC"
  resource_name = "example-topic"
}

# Query CKafka connect resources
data "cloud_ckafka_connect_resource" "connect" {
  resource_id = "resource-xxxxx"
}

# Query CKafka regions
data "cloud_ckafka_region" "regions" {}

# Query CKafka zones
data "cloud_ckafka_zone" "zones" {}

# Query CKafka DataHub topics
data "cloud_ckafka_datahub_topic" "datahub" {
  search_word = "example"
}

# Query CKafka groups
data "cloud_ckafka_group" "groups" {
  instance_id = "ckafka-xxxxx"
  search_word = "example-group"
}

# Query CKafka group info
data "cloud_ckafka_group_info" "group_info" {
  instance_id = "ckafka-xxxxx"
  group_list  = ["example-group"]
}

# Query CKafka group offsets
data "cloud_ckafka_group_offsets" "offsets" {
  instance_id = "ckafka-xxxxx"
  group       = "example-group"
}

# ========== Resources ==========

# CKafka Instance
resource "cloud_ckafka_instance" "instance" {
  instance_name      = "example-ckafka"
  zone_id            = 100003
  period             = 1
  vpc_id             = "vpc-xxxxx"
  subnet_id          = "subnet-xxxxx"
  msg_retention_time = 168
  renew_flag         = 0
  kafka_version      = "2.4.2"
  disk_type          = "CLOUD_BASIC"
  disk_size          = 500
  partition          = 500
  topic_num          = 50

  config {
    auto_create_topic_enable   = true
    default_num_partitions     = 3
    default_replication_factor = 2
  }
}

# CKafka User
resource "cloud_ckafka_user" "user" {
  instance_id = cloud_ckafka_instance.instance.id
  account_name = "example-user"
  password     = "Password123!"
}

# CKafka Topic
resource "cloud_ckafka_topic" "topic" {
  instance_id                     = cloud_ckafka_instance.instance.id
  topic_name                      = "example-topic"
  note                            = "Example topic"
  replica_num                     = 2
  partition_num                   = 3
  enable_white_list               = false
  ip_white_list                   = []
  clean_up_policy                 = "delete"
  min_insync_replica_num          = 1
  unclean_leader_election_enable  = false
  retention                       = 168
  segment                         = 168
  max_message_bytes              = 1024000
}

# CKafka ACL
resource "cloud_ckafka_acl" "acl" {
  instance_id     = cloud_ckafka_instance.instance.id
  resource_type   = "TOPIC"
  resource_name   = cloud_ckafka_topic.topic.topic_name
  operation_type  = "WRITE"
  permission_type = "ALLOW"
  host            = "*"
  principal       = cloud_ckafka_user.user.account_name
}

# CKafka ACL Rule
resource "cloud_ckafka_acl_rule" "rule" {
  instance_id     = cloud_ckafka_instance.instance.id
  resource_type   = "TOPIC"
  pattern_type    = "LITERAL"
  rule_name       = "example-rule"
  rule_list {
    operation       = "Read"
    permission_type = "Allow"
    host            = "*"
    principal       = cloud_ckafka_user.user.account_name
  }
}

# CKafka Consumer Group
resource "cloud_ckafka_consumer_group" "group" {
  instance_id = cloud_ckafka_instance.instance.id
  group_name  = "example-group"
}

# CKafka DataHub Task
resource "cloud_ckafka_datahub_task" "task" {
  task_name     = "example-task"
  task_type     = "SOURCE"
  source_resource {
    type = "MYSQL"
  }
  target_resource {
    type = "CKAFKA"
    kafka_param {
      instance_id = cloud_ckafka_instance.instance.id
      topic_name  = cloud_ckafka_topic.topic.topic_name
    }
  }
}

# CKafka Renew Instance
resource "cloud_ckafka_renew_instance" "renew" {
  instance_id = cloud_ckafka_instance.instance.id
  time_span   = 1
}
