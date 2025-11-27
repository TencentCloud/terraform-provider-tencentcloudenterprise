# Tencent Service Framework (TSF) Examples

# ========== Data Sources ==========

# Query TSF clusters
data "tencentcloudenterprise_tsf_cluster" "clusters" {
  cluster_id = "cluster-xxxxx"
}

# Query TSF applications
data "tencentcloudenterprise_tsf_application" "apps" {
  application_id = "application-xxxxx"
}

# Query TSF application config
data "tencentcloudenterprise_tsf_application_config" "config" {
  config_id = "config-xxxxx"
}

# Query TSF application file config
data "tencentcloudenterprise_tsf_application_file_config" "file_config" {
  config_id = "config-xxxxx"
}

# Query TSF application public config
data "tencentcloudenterprise_tsf_application_public_config" "public_config" {
  config_id = "config-xxxxx"
}

# Query TSF microservices
data "tencentcloudenterprise_tsf_microservice" "microservices" {
  microservice_id = "ms-xxxxx"
}

# Query TSF config summary
data "tencentcloudenterprise_tsf_config_summary" "summary" {
  application_id = "application-xxxxx"
}

# Query TSF delivery configs
data "tencentcloudenterprise_tsf_delivery_configs" "delivery" {
  search_word = "example"
}

# Query TSF delivery config by group ID
data "tencentcloudenterprise_tsf_delivery_config_by_group_id" "group_delivery" {
  group_id = "group-xxxxx"
}

# Query TSF public config summary
data "tencentcloudenterprise_tsf_public_config_summary" "public_summary" {
  search_word = "example"
}

# Query TSF API groups
data "tencentcloudenterprise_tsf_api_group" "api_groups" {
  group_id = "grp-xxxxx"
}

# Query TSF application attribute
data "tencentcloudenterprise_tsf_application_attribute" "attr" {
  application_id = "application-xxxxx"
}

# Query TSF business log configs
data "tencentcloudenterprise_tsf_business_log_configs" "log_configs" {
  search_word = "example"
}

# Query TSF API detail
data "tencentcloudenterprise_tsf_api_detail" "api" {
  microservice_id = "ms-xxxxx"
  path            = "/api/v1/users"
}

# Query TSF microservice API version
data "tencentcloudenterprise_tsf_microservice_api_version" "api_version" {
  microservice_id = "ms-xxxxx"
}

# Query TSF repositories
data "tencentcloudenterprise_tsf_repository" "repos" {
  search_word = "example"
}

# Query TSF pod instances
data "tencentcloudenterprise_tsf_pod_instances" "pods" {
  group_id = "group-xxxxx"
}

# Query TSF gateway all group APIs
data "tencentcloudenterprise_tsf_gateway_all_group_apis" "gateway_apis" {
  gateway_deploy_group_id = "group-xxxxx"
}

# Query TSF group gateways
data "tencentcloudenterprise_tsf_group_gateways" "group_gateways" {
  gateway_deploy_group_id = "group-xxxxx"
}

# Query TSF usable unit namespaces
data "tencentcloudenterprise_tsf_usable_unit_namespaces" "unit_namespaces" {
  search_word = "example"
}

# Query TSF group instances
data "tencentcloudenterprise_tsf_group_instances" "instances" {
  group_id = "group-xxxxx"
}

# Query TSF container groups
data "tencentcloudenterprise_tsf_container_group" "container_groups" {
  group_id = "group-xxxxx"
}

# Query TSF groups
data "tencentcloudenterprise_tsf_groups" "groups" {
  cluster_id = "cluster-xxxxx"
}

# Query TSF microservice API list
data "tencentcloudenterprise_tsf_ms_api_list" "ms_apis" {
  microservice_id = "ms-xxxxx"
}

# ========== Resources ==========

# TSF Cluster
resource "tencentcloudenterprise_tsf_cluster" "cluster" {
  cluster_name = "example-tsf-cluster"
  cluster_type = "V"  # CVM cluster
  vpc_id       = "vpc-xxxxx"
  cluster_cidr = "172.16.0.0/16"
  cluster_desc = "Example TSF cluster"
  
  tags = {
    env = "test"
  }
}

# TSF Namespace
resource "tencentcloudenterprise_tsf_namespace" "namespace" {
  namespace_name = "example-namespace"
  cluster_id     = cloud_tsf_cluster.cluster.id
  namespace_desc = "Example namespace"
}

# TSF Application
resource "tencentcloudenterprise_tsf_application" "app" {
  application_name = "example-app"
  application_type = "V"  # VM application
  microservice_type = "M"  # SpringCloud
  application_desc = "Example application"
}

# TSF Microservice
resource "tencentcloudenterprise_tsf_microservice" "microservice" {
  namespace_id     = cloud_tsf_namespace.namespace.id
  microservice_name = "example-microservice"
  microservice_desc = "Example microservice"
}

# TSF Group
resource "tencentcloudenterprise_tsf_group" "group" {
  application_id = cloud_tsf_application.app.id
  namespace_id   = cloud_tsf_namespace.namespace.id
  group_name     = "example-group"
  cluster_id     = cloud_tsf_cluster.cluster.id
  group_desc     = "Example deployment group"
}

# TSF Application Config
resource "tencentcloudenterprise_tsf_application_config" "config" {
  config_name     = "example-config"
  config_version  = "1.0"
  config_value    = yamlencode({
    server = {
      port = 8080
    }
    spring = {
      application = {
        name = "example-app"
      }
    }
  })
  application_id  = cloud_tsf_application.app.id
  config_version_desc = "Initial version"
}

# TSF Application Release Config
resource "tencentcloudenterprise_tsf_application_release_config" "release" {
  config_id = cloud_tsf_application_config.config.id
  group_id  = cloud_tsf_group.group.id
  release_desc = "Release to production"
}

# TSF Application File Config
resource "tencentcloudenterprise_tsf_application_file_config" "file_config" {
  config_name      = "application.yml"
  config_version   = "1.0"
  config_file_value = base64encode(yamlencode({
    server = {
      port = 8080
    }
  }))
  application_id   = cloud_tsf_application.app.id
  config_file_path = "/config"
  config_version_desc = "Config file"
}

# TSF Application File Config Release
resource "tencentcloudenterprise_tsf_application_file_config_release" "file_release" {
  config_id = cloud_tsf_application_file_config.file_config.id
  group_id  = cloud_tsf_group.group.id
  release_desc = "Release config file"
}

# TSF Application Public Config
resource "tencentcloudenterprise_tsf_application_public_config" "public_config" {
  config_name     = "public-config"
  config_version  = "1.0"
  config_value    = yamlencode({
    common = {
      timeout = 30
    }
  })
  config_version_desc = "Public configuration"
}

# TSF Application Public Config Release
resource "tencentcloudenterprise_tsf_application_public_config_release" "public_release" {
  config_id    = cloud_tsf_application_public_config.public_config.id
  namespace_id = cloud_tsf_namespace.namespace.id
  release_desc = "Release public config"
}

# TSF Config Template
resource "tencentcloudenterprise_tsf_config_template" "template" {
  template_name = "example-template"
  template_type = "K"
  template_value = yamlencode({
    example_key = "example_value"
  })
  template_desc = "Example config template"
}

# TSF API Group
resource "tencentcloudenterprise_tsf_api_group" "api_group" {
  group_name    = "example-api-group"
  group_context = "/api"
  auth_type     = "none"
  description   = "Example API group"
  group_type    = "ms"
}

# TSF Bind API Group
resource "tencentcloudenterprise_tsf_bind_api_group" "bind" {
  group_gateway_list {
    gateway_deploy_group_id = "group-xxxxx"
    group_id                = cloud_tsf_api_group.api_group.id
  }
}

# TSF API Rate Limit Rule
resource "tencentcloudenterprise_tsf_api_rate_limit_rule" "rate_limit" {
  api_id           = "api-xxxxx"
  max_qps          = 100
  usable_status    = "enabled"
}

# TSF Path Rewrite
resource "tencentcloudenterprise_tsf_path_rewrite" "rewrite" {
  path_rewrite_name = "example-rewrite"
  gateway_group_id  = "group-xxxxx"
  regex             = "^/old/(.*)"
  replacement       = "/new/$1"
  blocked           = "N"
  order             = 0
}

# TSF Lane
resource "tencentcloudenterprise_tsf_lane" "lane" {
  lane_name    = "example-lane"
  remark       = "Example lane for canary deployment"
  lane_group_list {
    group_id     = cloud_tsf_group.group.id
    entrance     = true
  }
}

# TSF Lane Rule
resource "tencentcloudenterprise_tsf_lane_rule" "lane_rule" {
  rule_name = "example-lane-rule"
  remark    = "Example lane rule"
  rule_tag_list {
    tag_id        = "tag-xxxxx"
    tag_name      = "version"
    tag_operator  = "EQUAL"
    tag_value     = "v2"
  }
  lane_id = cloud_tsf_lane.lane.id
  enable  = true
}

# TSF Unit Rule
resource "tencentcloudenterprise_tsf_unit_rule" "unit_rule" {
  gateway_instance_id = "gateway-xxxxx"
  name                = "example-unit-rule"
  description         = "Example unit rule"
  unit_rule_item_list {
    relationship = "AND"
    dest_namespace_id = cloud_tsf_namespace.namespace.id
    priority      = 1
    unit_rule_tag_list {
      tag_type      = "U"
      tag_field     = "user_id"
      tag_operator  = "IN"
      tag_value     = "1001,1002"
    }
  }
}

# TSF Enable Unit Rule
resource "tencentcloudenterprise_tsf_enable_unit_rule" "enable" {
  rule_id = cloud_tsf_unit_rule.unit_rule.id
}

# TSF Task
resource "tencentcloudenterprise_tsf_task" "task" {
  task_name    = "example-task"
  task_content = base64encode("#!/bin/bash\necho 'Hello TSF'")
  task_type    = "execute-script"
  group_id     = cloud_tsf_group.group.id
  time_out     = 60
  retry_count  = 3
  retry_interval = 10
}

# TSF Instances Attachment
resource "tencentcloudenterprise_tsf_instances_attachment" "attachment" {
  cluster_id  = cloud_tsf_cluster.cluster.id
  instance_id = "ins-xxxxx"
}

# TSF Deploy Container Group
resource "tencentcloudenterprise_tsf_deploy_container_group" "deploy" {
  group_id        = cloud_tsf_group.group.id
  server          = "ccr.ccs.tencentyun.com"
  reponame        = "example/app"
  tag_name        = "v1.0.0"
  cpu_limit       = "0.5"
  mem_limit       = "1Gi"
  instance_num    = 2
  update_type     = 0
  update_ivl      = 10
}

# TSF Deploy VM Group
resource "tencentcloudenterprise_tsf_deploy_vm_group" "deploy_vm" {
  group_id    = cloud_tsf_group.group.id
  pkg_id      = "pkg-xxxxx"
  startup_parameters = "-Xms512m -Xmx1024m"
  update_type = 0
  update_ivl  = 10
}

# TSF Release API Group
resource "tencentcloudenterprise_tsf_release_api_group" "release_api" {
  group_id = cloud_tsf_api_group.api_group.id
}

# TSF Operate Container Group
resource "tencentcloudenterprise_tsf_operate_container_group" "operate" {
  group_id  = cloud_tsf_group.group.id
  operate   = "start"
}

# TSF Operate Group (VM)
resource "tencentcloudenterprise_tsf_operate_group" "operate_vm" {
  group_id  = cloud_tsf_group.group.id
  operate   = "start"
}

# TSF Unit Namespace
resource "tencentcloudenterprise_tsf_unit_namespace" "unit_ns" {
  gateway_instance_id = "gateway-xxxxx"
  namespace_id        = cloud_tsf_namespace.namespace.id
  namespace_name      = "example-unit-namespace"
  description         = "Example unit namespace"
}
