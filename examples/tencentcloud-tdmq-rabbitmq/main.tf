# TDMQ for RabbitMQ Examples

# ========== Data Sources ==========

# Query RabbitMQ node list
data "tencentcloudenterprise_tdmq_rabbitmq_node_list" "nodes" {
  instance_id = "rabbitmq-xxxxx"
}

# Query RabbitMQ VIP instance
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "instance" {
  cluster_id = "rabbitmq-xxxxx"
}

# ========== Resources ==========

# RabbitMQ VIP Instance
resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "instance" {
  name              = "example-rabbitmq"
  spec_type         = "rabbit-vip-basic-1"
  node_num        = 3
  storage_size      = 200
  enable_public     = false
  bandwidth         = 10
  ip_rules          = ["0.0.0.0/0"]
  vpc_id            = "vpc-xxxxx"
  subnet_id         = "subnet-xxxxx"
  tags = {
    env = "test"
  }
}

# RabbitMQ User
resource "tencentcloudenterprise_tdmq_rabbitmq_user" "user" {
  instance_id = cloud_tdmq_rabbitmq_vip_instance.instance.id
  user        = "example-user"
  password    = "Password123!"
  description = "Example RabbitMQ user"
  tags        = ["administrator"]
}

# RabbitMQ Virtual Host
resource "tencentcloudenterprise_tdmq_rabbitmq_virtual_host" "vhost" {
  instance_id  = cloud_tdmq_rabbitmq_vip_instance.instance.id
  virtual_host = "example-vhost"
  description  = "Example virtual host"
}
