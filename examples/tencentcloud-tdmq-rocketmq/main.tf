# TDMQ for RocketMQ Examples

# ========== Data Sources ==========

# Query RocketMQ clusters
data "cloud_tdmq_rocketmq_cluster" "clusters" {
  cluster_id = "rocketmq-xxxxx"
}

# Query RocketMQ namespaces
data "cloud_tdmq_rocketmq_namespace" "namespaces" {
  cluster_id = "rocketmq-xxxxx"
}

# Query RocketMQ topics
data "cloud_tdmq_rocketmq_topic" "topics" {
  cluster_id = "rocketmq-xxxxx"
  namespace_id = "namespace-xxxxx"
}

# Query RocketMQ roles
data "cloud_tdmq_rocketmq_role" "roles" {
  cluster_id = "rocketmq-xxxxx"
}

# Query RocketMQ groups
data "cloud_tdmq_rocketmq_group" "groups" {
  cluster_id   = "rocketmq-xxxxx"
  namespace_id = "namespace-xxxxx"
}

# Query RocketMQ messages
data "cloud_tdmq_rocketmq_messages" "messages" {
  cluster_id = "rocketmq-xxxxx"
  environment_id = "namespace-xxxxx"
  topic_name = "example-topic"
  msg_id     = "msg-xxxxx"
}

# ========== Resources ==========

# RocketMQ Cluster
resource "cloud_tdmq_rocketmq_cluster" "cluster" {
  cluster_name = "example-rocketmq"
  remark       = "Example RocketMQ cluster"
}

# RocketMQ Namespace
resource "cloud_tdmq_rocketmq_namespace" "namespace" {
  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
  namespace_name = "example-namespace"
  ttl            = 3600
  retention_time = 72
  remark         = "Example namespace"
}

# RocketMQ Role
resource "cloud_tdmq_rocketmq_role" "role" {
  role_name  = "example-role"
  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
  remark     = "Example role"
}

# RocketMQ Topic
resource "cloud_tdmq_rocketmq_topic" "topic" {
  topic_name   = "example-topic"
  namespace    = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
  type         = "Normal"
  cluster_id   = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
  remark       = "Example topic"
  partition_num = 3
}

# RocketMQ Group
resource "cloud_tdmq_rocketmq_group" "group" {
  group_name     = "example-group"
  namespace      = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
  read_enable    = true
  broadcast_enable = false
  cluster_id     = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
  remark         = "Example consumer group"
}

# RocketMQ Environment Role
resource "cloud_tdmq_rocketmq_environment_role" "env_role" {
  environment_id = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
  role_name      = cloud_tdmq_rocketmq_role.role.role_name
  permissions    = ["produce", "consume"]
  cluster_id     = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
}

# Send RocketMQ Message
resource "cloud_tdmq_send_rocketmq_message" "message" {
  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
  namespace_id = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
  topic_name = cloud_tdmq_rocketmq_topic.topic.topic_name
  msg_body   = "Hello, RocketMQ!"
  msg_tag    = "example-tag"
  msg_key    = "example-key"
}
