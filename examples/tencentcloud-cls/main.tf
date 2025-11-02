# Cloud Log Service (CLS) Examples

# ========== Data Sources ==========

# Query CLS machine group configs
data "cloud_cls_machine_group_configs" "configs" {
  group_id = "group-xxxxx"
}

# ========== Resources ==========

# CLS Logset
resource "cloud_cls_logset" "logset" {
  logset_name = "example-logset"
  period      = 7
  
  tags = {
    env = "test"
  }
}

# CLS Topic
resource "cloud_cls_topic" "topic" {
  topic_name           = "example-topic"
  logset_id            = cloud_cls_logset.logset.id
  auto_split           = true
  max_split_partitions = 20
  partition_count      = 1
  period               = 7
  storage_type         = "hot"
  
  tags = {
    env = "test"
  }
}

# CLS Index
resource "cloud_cls_index" "index" {
  topic_id = cloud_cls_topic.topic.id
  status   = true
  
  rule {
    full_text {
      case_sensitive = false
      tokenizer      = "!@#%^&*()_=\"', <>/?|\\;:\\n\\t\\r[]{}."
      contain_z_h    = false
    }
    
    key_value {
      case_sensitive = false
      
      key_values {
        key = "level"
        value {
          type          = "text"
          tokenizer     = ""
          sql_flag      = true
          contain_z_h   = false
        }
      }
      
      key_values {
        key = "timestamp"
        value {
          type     = "long"
          sql_flag = true
        }
      }
    }
  }
}

# CLS Machine Group
resource "cloud_cls_machine_group" "group" {
  group_name        = "example-machine-group"
  service_logging   = true
  auto_update       = true
  update_start_time = "02:00"
  update_end_time   = "06:00"
  
  machine_group_type {
    type = "ip"
    values = ["10.0.1.10", "10.0.1.11"]
  }
  
  tags = {
    env = "test"
  }
}

# CLS Config
resource "cloud_cls_config" "config" {
  name = "example-config"
  output = cloud_cls_topic.topic.id
  path = "/var/log/nginx"
  
  log_type = "json_log"
  
  extract_rule {
    time_key   = "time"
    time_format = "%Y-%m-%d %H:%M:%S"
  }
}

# CLS Config Attachment
resource "cloud_cls_config_attachment" "attachment" {
  config_id = cloud_cls_config.config.id
  group_id  = cloud_cls_machine_group.group.id
}

# CLS COS Shipper
resource "cloud_cls_cos_shipper" "shipper" {
  topic_id       = cloud_cls_topic.topic.id
  bucket         = "example-bucket-123456"
  prefix         = "logs/"
  shipper_name   = "example-shipper"
  interval       = 300
  max_size       = 256
  partition      = "%Y/%m/%d"
  compress_type  = "gzip"
  
  content {
    format = "json"
  }
}

# CLS COS Recharge
resource "cloud_cls_cos_recharge" "recharge" {
  topic_id    = cloud_cls_topic.topic.id
  name        = "example-recharge"
  bucket      = "example-bucket-123456"
  bucket_region = "ap-guangzhou"
  prefix      = "logs/"
  log_type    = "json_log"
}

# CLS CKafka Consumer
resource "cloud_cls_ckafka_consumer" "consumer" {
  topic_id = cloud_cls_topic.topic.id
  need_content = true
  
  content {
    format = "json"
  }
  
  ckafka {
    instance_id   = "ckafka-xxxxx"
    instance_name = "example-ckafka"
    topic_id      = "topic-xxxxx"
    topic_name    = "example-topic"
  }
}

# CLS Export
resource "cloud_cls_export" "export" {
  topic_id    = cloud_cls_topic.topic.id
  log_count   = 1000
  query       = "level:ERROR"
  from        = 1609459200000  # Unix timestamp in ms
  to          = 1609545600000  # Unix timestamp in ms
  order       = "desc"
  format      = "json"
}
