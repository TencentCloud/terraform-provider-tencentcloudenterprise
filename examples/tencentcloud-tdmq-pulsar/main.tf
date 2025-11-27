# TDMQ for Pulsar Examples

# ========== Data Sources ==========

# Query Pulsar clusters
data "tencentcloudenterprise_tdmq_pulsar_clusters" "clusters" {
  cluster_id = "pulsar-xxxxx"
}

# Query Pulsar environments (namespaces)
data "tencentcloudenterprise_tdmq_pulsar_environments" "envs" {
  cluster_id = "pulsar-xxxxx"
}

# ========== Resources ==========

# Pulsar Cluster
resource "tencentcloudenterprise_tdmq_pulsar_cluster" "cluster" {
  cluster_name = "example-pulsar"
  remark       = "Example Pulsar cluster"
  
  bind_cluster {
    cluster_name = "example-physical-cluster"
  }
}

# Pulsar Route (HTTP access)
resource "tencentcloudenterprise_tdmq_pulsar_route" "route" {
  cluster_id  = cloud_tdmq_pulsar_cluster.cluster.cluster_id
  route_type  = 0
}

# Pulsar Environment (Namespace)
resource "tencentcloudenterprise_tdmq_pulsar_environment" "env" {
  environment_name = "example-namespace"
  cluster_id       = cloud_tdmq_pulsar_cluster.cluster.cluster_id
  msg_ttl          = 300
  remark           = "Example namespace"
  retention_policy {
    time_in_minutes = 60
    size_in_mb      = 1024
  }
}

# Pulsar Topic
resource "tencentcloudenterprise_tdmq_pulsar_topic" "topic" {
  environment_id = cloud_tdmq_pulsar_environment.env.environment_id
  topic_name     = "example-topic"
  cluster_id     = cloud_tdmq_pulsar_cluster.cluster.cluster_id
  partitions     = 3
  topic_type     = 0
  remark         = "Example topic"
}

# Pulsar Role
resource "tencentcloudenterprise_tdmq_pulsar_role" "role" {
  role_name  = "example-role"
  cluster_id = cloud_tdmq_pulsar_cluster.cluster.cluster_id
  remark     = "Example role"
}

# Pulsar Environment Role Attachment
resource "tencentcloudenterprise_tdmq_pulsar_environment_role_attachment" "attachment" {
  environment_id = cloud_tdmq_pulsar_environment.env.environment_id
  role_name      = cloud_tdmq_pulsar_role.role.role_name
  permissions    = ["produce", "consume"]
  cluster_id     = cloud_tdmq_pulsar_cluster.cluster.cluster_id
}
