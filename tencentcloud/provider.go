/*
The TencentCloud provider is used to interact with many resources supported by [TencentCloud](https://intl.cloud.tencent.com).
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.

-> **Note:** From version 1.9.0 (June 18, 2019), the provider start to support Terraform 0.12.x.

Example Usage
```hcl

	terraform {
	  required_providers {
	    tencentcloudenterprise = {
	      source = "TencentCloud/tencentcloudenterprise"
	    }
	  }
	}

# Configure the TencentCloud Enterprise Provider

	provider "tencentcloudenterprise" {
	  secret_id  = var.secret_id
	  secret_key = var.secret_key
	  region     = var.region
	}

#Configure the TencentCloud Enterprise Provider with STS

	provider "tencentcloudenterprise" {
	  secret_id  = var.secret_id
	  secret_key = var.secret_key
	  region     = var.region
	  assume_role {
	    role_arn         = var.assume_role_arn
	    session_name     = var.session_name
	    session_duration = var.session_duration
	    policy           = var.policy
	  }
	}

```

# Resources List

Provider Data Sources

	tencentcloudenterprise_availability_regions
	tencentcloudenterprise_availability_zones

Auto Scaling(AS)

	 Data Source
	   tencentcloudenterprise_as_scaling_configs
	   tencentcloudenterprise_as_scaling_groups
	   tencentcloudenterprise_as_scaling_policies
	tencentcloudenterprise_as_instances
	tencentcloudenterprise_as_last_activity

	 Resource
	   tencentcloudenterprise_as_scaling_config
	   tencentcloudenterprise_as_scaling_group
	   tencentcloudenterprise_as_attachment
	   tencentcloudenterprise_as_scaling_policy
	   tencentcloudenterprise_as_schedule
	   tencentcloudenterprise_as_lifecycle_hook
	   tencentcloudenterprise_as_notification
	tencentcloudenterprise_as_remove_instances
	   tencentcloudenterprise_as_protect_instances

Bare Metal Server(BMS)

	 Data Source
	tencentcloudenterprise_bms_instances
	tencentcloudenterprise_bms_placement_groups
	tencentcloudenterprise_bms_flavors

	 Resource
	tencentcloudenterprise_bms_instance
	tencentcloudenterprise_bms_placement_group

Cloud Kafka(ckafka)

	 Data Source
	   tencentcloudenterprise_ckafka_users
	   tencentcloudenterprise_ckafka_acls
	   tencentcloudenterprise_ckafka_topics
	   tencentcloudenterprise_ckafka_instances
	tencentcloudenterprise_ckafka_connect_resource
	tencentcloudenterprise_ckafka_region
	tencentcloudenterprise_ckafka_datahub_topic
	tencentcloudenterprise_ckafka_datahub_group_offsets
	tencentcloudenterprise_ckafka_datahub_task
	tencentcloudenterprise_ckafka_group
	tencentcloudenterprise_ckafka_group_offsets
	tencentcloudenterprise_ckafka_group_info
	tencentcloudenterprise_ckafka_task_status
	tencentcloudenterprise_ckafka_topic_flow_ranking
	tencentcloudenterprise_ckafka_topic_produce_connection
	tencentcloudenterprise_ckafka_topic_subscribe_group
	tencentcloudenterprise_ckafka_topic_sync_replica
	tencentcloudenterprise_ckafka_zone

	 Resource
	tencentcloudenterprise_ckafka_instance
	   tencentcloudenterprise_ckafka_user
	   tencentcloudenterprise_ckafka_acl
	   tencentcloudenterprise_ckafka_topic
	tencentcloudenterprise_ckafka_datahub_topic
	tencentcloudenterprise_ckafka_connect_resource
	tencentcloudenterprise_ckafka_renew_instance
	tencentcloudenterprise_ckafka_acl_rule
	tencentcloudenterprise_ckafka_consumer_group
	tencentcloudenterprise_ckafka_consumer_group_modify_offset
	tencentcloudenterprise_ckafka_datahub_task

Cloud Block Storage(CBS)

	 Data Source
	   tencentcloudenterprise_cbs_snapshots
	   tencentcloudenterprise_cbs_storages
	tencentcloudenterprise_cbs_storages_set
	   tencentcloudenterprise_cbs_snapshot_policies

	 Resource
	   tencentcloudenterprise_cbs_storage
	tencentcloudenterprise_cbs_storage_set
	   tencentcloudenterprise_cbs_storage_attachment
	tencentcloudenterprise_cbs_storage_set_attachment
	   tencentcloudenterprise_cbs_snapshot
	   tencentcloudenterprise_cbs_snapshot_policy
	   tencentcloudenterprise_cbs_snapshot_policy_attachment
	tencentcloudenterprise_cbs_snapshot_share_permission

Cloud Load Balancer(CLB)

	 Data Source
	   tencentcloudenterprise_clb_attachments
	   tencentcloudenterprise_clb_instances
	   tencentcloudenterprise_clb_listener_rules
	   tencentcloudenterprise_clb_listeners
	   tencentcloudenterprise_clb_redirections
	tencentcloudenterprise_clb_instance_by_cert_id
	tencentcloudenterprise_clb_instance_detail
	tencentcloudenterprise_clb_resources
	tencentcloudenterprise_clb_target_health
	   tencentcloudenterprise_clb_certificates

	 Resource
	   tencentcloudenterprise_clb_instance
	   tencentcloudenterprise_clb_listener
	   tencentcloudenterprise_clb_listener_rule
	   tencentcloudenterprise_clb_attachment
	   tencentcloudenterprise_clb_redirection
	tencentcloudenterprise_clb_customized_config
	tencentcloudenterprise_clb_security_group_attachment
	tencentcloudenterprise_clb_certificates

Cloud Object Storage(COS)

	Data Source
	  tencentcloudenterprise_cos_bucket_object
	  tencentcloudenterprise_cos_buckets

	Resource
	  tencentcloudenterprise_cos_bucket
	  tencentcloudenterprise_cos_bucket_object
	  tencentcloudenterprise_cos_bucket_policy

Cloud Object Storage(CSP)

	Data Source
	  tencentcloudenterprise_csp_bucket_object
	  tencentcloudenterprise_csp_buckets

	Resource
	  tencentcloudenterprise_csp_bucket
	  tencentcloudenterprise_csp_bucket_object
	  tencentcloudenterprise_csp_bucket_policy

Cloud Virtual Machine(CVM)

	 Data Source
	   tencentcloudenterprise_cvm_image
	   tencentcloudenterprise_cvm_images
	   tencentcloudenterprise_cvm_instance_types
	   tencentcloudenterprise_cvm_instances
	tencentcloudenterprise_cvm_instances_set
	   tencentcloudenterprise_cvm_key_pairs
	   tencentcloudenterprise_cvm_placement_groups
	tencentcloudenterprise_cvm_instances_modification
	tencentcloudenterprise_cvm_instance_vnc_url
	tencentcloudenterprise_cvm_disaster_recover_group_quota
	tencentcloudenterprise_cvm_image_quota
	tencentcloudenterprise_cvm_image_share_permission

	 Resource
	   tencentcloudenterprise_cvm_instance
	tencentcloudenterprise_cvm_instance_set

	   tencentcloudenterprise_cvm_key_pair
	   tencentcloudenterprise_cvm_placement_group
	   tencentcloudenterprise_cvm_image
	tencentcloudenterprise_cvm_launch_template
	tencentcloudenterprise_cvm_security_group_attachment
	tencentcloudenterprise_cvm_reboot_instance
	tencentcloudenterprise_cvm_renew_instance
	tencentcloudenterprise_cvm_sync_image
	tencentcloudenterprise_cvm_image_share_permission

Cloud Elastic IP(EIP)

	 Data Source
	tencentcloudenterprise_eips
	tencentcloudenterprise_eip_address_quota
	 Resource
	   tencentcloudenterprise_eip
	   tencentcloudenterprise_eip_association
	tencentcloudenterprise_eip_address_transform
	tencentcloudenterprise_eip_public_address_adjust
	tencentcloudenterprise_eip_normal_address_return

Direct Connect(DC)

	Data Source


	Resource

Elasticsearch Service(ES)

	Data Source

	Resource

Key Management Service(KMS)

	Data Source
	  tencentcloudenterprise_kms_keys

	Resource
	  tencentcloudenterprise_kms_key
	  tencentcloudenterprise_kms_external_key

Corporate Identity Center(CIC)

		  Data Source

		  Resource
	        tencentcloudenterprise_cic_external_saml_identity_provider
	        tencentcloudenterprise_cic_group
	        tencentcloudenterprise_cic_role_assignment
	        tencentcloudenterprise_cic_role_configuration
	        tencentcloudenterprise_cic_role_configuration_permission_custom_policy_attachment
	        tencentcloudenterprise_cic_role_configuration_permission_policy_attachment
	        tencentcloudenterprise_cic_scim_credential
	        tencentcloudenterprise_cic_scim_synchronization_status
	        tencentcloudenterprise_cic_user
	        tencentcloudenterprise_cic_user_group_attachment
	        tencentcloudenterprise_cic_user_sync_provisioning

Tencent Kubernetes Engine(TKE)

	 Data Source
	   tencentcloudenterprise_tke_kubernetes_clusters
	   tencentcloudenterprise_tke_kubernetes_charts
	   tencentcloudenterprise_tke_kubernetes_cluster_common_names
	tencentcloudenterprise_tke_kubernetes_available_cluster_versions

	 Resource
	   tencentcloudenterprise_tke_kubernetes_cluster
	   tencentcloudenterprise_tke_kubernetes_scale_worker
	   tencentcloudenterprise_tke_kubernetes_cluster_attachment
	tencentcloudenterprise_tke_kubernetes_cluster_endpoint

TencentDB for PostgreSQL(PostgreSQL)

	 Data Source
	tencentcloudenterprise_postgresql_instances
	tencentcloudenterprise_postgresql_specinfos
	tencentcloudenterprise_postgresql_xlogs
	tencentcloudenterprise_postgresql_parameter_templates
	tencentcloudenterprise_postgresql_readonly_groups
	tencentcloudenterprise_postgresql_base_backups
	tencentcloudenterprise_postgresql_log_backups
	tencentcloudenterprise_postgresql_backup_download_urls
	tencentcloudenterprise_postgresql_db_instance_classes
	tencentcloudenterprise_postgresql_default_parameters
	tencentcloudenterprise_postgresql_recovery_time
	tencentcloudenterprise_postgresql_regions
	tencentcloudenterprise_postgresql_db_instance_versions
	tencentcloudenterprise_postgresql_zones

	 Resource
	tencentcloudenterprise_postgresql_instance
	tencentcloudenterprise_postgresql_readonly_instance
	tencentcloudenterprise_postgresql_readonly_group
	tencentcloudenterprise_postgresql_readonly_attachment
	tencentcloudenterprise_postgresql_parameter_template
	tencentcloudenterprise_postgresql_backup_plan_config
	tencentcloudenterprise_postgresql_security_group_config
	tencentcloudenterprise_postgresql_backup_download_restriction_config
	tencentcloudenterprise_postgresql_restart_db_instance_operation
	tencentcloudenterprise_postgresql_renew_db_instance_operation
	tencentcloudenterprise_postgresql_isolate_db_instance_operation
	tencentcloudenterprise_postgresql_disisolate_db_instance_operation
	tencentcloudenterprise_postgresql_rebalance_readonly_group_operation
	tencentcloudenterprise_postgresql_delete_log_backup_operation
	tencentcloudenterprise_postgresql_modify_account_remark_operation
	tencentcloudenterprise_postgresql_modify_switch_time_period_operation
	tencentcloudenterprise_postgresql_base_backup

Virtual Private Cloud(VPC)

	 Data Source
	   tencentcloudenterprise_vpc_security_groups
	tencentcloudenterprise_vpc_address_templates
	tencentcloudenterprise_vpc_address_template_groups
	   tencentcloudenterprise_vpc_acls
	tencentcloudenterprise_vpc_account_attributes
	tencentcloudenterprise_vpc_classic_link_instances
	tencentcloudenterprise_vpc_gateway_flow_qos
	tencentcloudenterprise_vpc_cvm_instances
	tencentcloudenterprise_vpc_net_detect_states
	tencentcloudenterprise_vpc_net_detect_state_check
	tencentcloudenterprise_vpc_private_ip_addresses
	tencentcloudenterprise_vpc_resource_dashboard
	tencentcloudenterprise_vpc_security_group_limits
	tencentcloudenterprise_vpc_security_group_references
	tencentcloudenterprise_vpc_template_limits
	tencentcloudenterprise_vpc_limits
	   tencentcloudenterprise_vpc_instances
	   tencentcloudenterprise_vpc_route_tables
	   tencentcloudenterprise_vpc_subnets
	   tencentcloudenterprise_vpc_dnats
	   tencentcloudenterprise_vpc_enis
	   tencentcloudenterprise_vpc_ha_vip_eip_attachments
	   tencentcloudenterprise_vpc_ha_vips
	   tencentcloudenterprise_vpc_nat_gateways
	tencentcloudenterprise_vpc_bandwidth_package_quota

	 Resource
	   tencentcloudenterprise_vpc_eni
	   tencentcloudenterprise_vpc_eni_attachment
	tencentcloudenterprise_vpc_eni_sg_attachment
	   tencentcloudenterprise_vpc
	tencentcloudenterprise_vpc_acl
	tencentcloudenterprise_vpc_acl_attachment
	tencentcloudenterprise_vpc_dc_gateway
	tencentcloudenterprise_vpc_net_detect
	tencentcloudenterprise_vpc_ipv6_cidr_block
	tencentcloudenterprise_vpc_ipv6_subnet_cidr_block
	tencentcloudenterprise_vpc_ipv6_eni_address
	   tencentcloudenterprise_vpc_subnet
	   tencentcloudenterprise_vpc_security_group
	   tencentcloudenterprise_vpc_security_group_rule
	   tencentcloudenterprise_vpc_security_group_rule_set
	   tencentcloudenterprise_vpc_security_group_lite_rule
	tencentcloudenterprise_vpc_address_template
	tencentcloudenterprise_vpc_address_template_group
	   tencentcloudenterprise_vpc_route_table
	   tencentcloudenterprise_vpc_route_table_entry
	   tencentcloudenterprise_vpc_dnat
	   tencentcloudenterprise_vpc_nat_gateway
	   tencentcloudenterprise_vpc_ha_vip
	   tencentcloudenterprise_vpc_ha_vip_eip_attachment
	tencentcloudenterprise_vpc_bandwidth_package
	tencentcloudenterprise_vpc_bandwidth_package_attachment
	tencentcloudenterprise_vpc_ipv6_address_bandwidth

Virtual Private Cloud DNS(VPCDNS)

	 Data Source
	tencentcloudenterprise_vpcdns_domains
	   tencentcloudenterprise_vpcdns_records

	 Resource
	tencentcloudenterprise_vpcdns_domain
	   tencentcloudenterprise_vpcdns_record

TDSQL for MySQL(DCDB)

	 Data Source
	tencentcloudenterprise_dcdb_instances
	tencentcloudenterprise_dcdb_accounts
	tencentcloudenterprise_dcdb_databases
	tencentcloudenterprise_dcdb_parameters
	tencentcloudenterprise_dcdb_shards
	tencentcloudenterprise_dcdb_security_groups
	tencentcloudenterprise_dcdb_database_objects
	tencentcloudenterprise_dcdb_database_tables

	 Resource
	tencentcloudenterprise_dcdb_account
	tencentcloudenterprise_dcdb_instance
	tencentcloudenterprise_dcdb_security_group_attachment
	tencentcloudenterprise_dcdb_account_privileges
	tencentcloudenterprise_dcdb_db_parameters
	tencentcloudenterprise_dcdb_db_sync_mode_config
	tencentcloudenterprise_dcdb_encrypt_attributes_config
	tencentcloudenterprise_dcdb_instance_config
	tencentcloudenterprise_dcdb_cancel_dcn_job_operation
	tencentcloudenterprise_dcdb_activate_hour_instance_operation
	tencentcloudenterprise_dcdb_isolate_hour_instance_operation
	tencentcloudenterprise_dcdb_flush_binlog_operation
	tencentcloudenterprise_dcdb_switch_db_instance_ha_operation

TDSQL PostgreSQL (Tbase)

	 Data Source
	   tencentcloudenterprise_tbase_instances
	   tencentcloudenterprise_tbase_pg_instances

	 Resource
	tencentcloudenterprise_tbase_instance
	tencentcloudenterprise_tbase_pg_instance
	tencentcloudenterprise_tbase_pg_instance_vip

TencentDB for Redis(crs)

	 Data Source
	tencentcloudenterprise_redis_zone_config
	   tencentcloudenterprise_redis_instances
	   tencentcloudenterprise_redis_backup
	   tencentcloudenterprise_redis_backup_download_info
	   tencentcloudenterprise_redis_param_records
	   tencentcloudenterprise_redis_instance_shards
	   tencentcloudenterprise_redis_instance_task_list
	   tencentcloudenterprise_redis_instance_node_info

	 Resource
	   tencentcloudenterprise_redis_instance
	tencentcloudenterprise_redis_backup_config
	tencentcloudenterprise_redis_param
	tencentcloudenterprise_redis_clear_instance_operation
	tencentcloudenterprise_redis_startup_instance_operation
	tencentcloudenterprise_redis_replica_readonly

TDMQ for Pulsar(tpulsar)

	 Data Source
	tencentcloudenterprise_tdmq_pulsar_clusters
	tencentcloudenterprise_tdmq_pulsar_environments

	 Resource
	   tencentcloudenterprise_tdmq_pulsar_cluster
	tencentcloudenterprise_tdmq_pulsar_route
	tencentcloudenterprise_tdmq_pulsar_environment
	tencentcloudenterprise_tdmq_pulsar_topic
	tencentcloudenterprise_tdmq_pulsar_role
	tencentcloudenterprise_tdmq_pulsar_environment_role_attachment

TDMQ for RabbitMQ(trabbit)

	 Data Source
	tencentcloudenterprise_tdmq_rabbitmq_node_list
	tencentcloudenterprise_tdmq_rabbitmq_vip_instance

	 Resource
	tencentcloudenterprise_tdmq_rabbitmq_user
	tencentcloudenterprise_tdmq_rabbitmq_vip_instance
	tencentcloudenterprise_tdmq_rabbitmq_virtual_host

TDMQ for RocketMQ(trocket)

	 Data Source
	tencentcloudenterprise_tdmq_rocketmq_cluster
	tencentcloudenterprise_tdmq_rocketmq_namespace
	tencentcloudenterprise_tdmq_rocketmq_topic
	tencentcloudenterprise_tdmq_rocketmq_role
	tencentcloudenterprise_tdmq_rocketmq_group
	tencentcloudenterprise_tdmq_rocketmq_messages

	 Resource
	tencentcloudenterprise_tdmq_rocketmq_cluster
	tencentcloudenterprise_tdmq_rocketmq_namespace
	tencentcloudenterprise_tdmq_rocketmq_role
	tencentcloudenterprise_tdmq_rocketmq_topic
	tencentcloudenterprise_tdmq_rocketmq_group
	tencentcloudenterprise_tdmq_rocketmq_environment_role
	tencentcloudenterprise_tdmq_send_rocketmq_message

Tencent Service Framework(TSF)

	Data Source
	  tencentcloudenterprise_tsf_application
	  tencentcloudenterprise_tsf_application_config
	  tencentcloudenterprise_tsf_application_file_config
	  tencentcloudenterprise_tsf_application_public_config
	  tencentcloudenterprise_tsf_cluster
	  tencentcloudenterprise_tsf_microservice
	  tencentcloudenterprise_tsf_config_summary
	  tencentcloudenterprise_tsf_delivery_config_by_group_id
	  tencentcloudenterprise_tsf_delivery_configs
	  tencentcloudenterprise_tsf_public_config_summary
	  tencentcloudenterprise_tsf_api_group
	  tencentcloudenterprise_tsf_application_attribute
	  tencentcloudenterprise_tsf_business_log_configs
	  tencentcloudenterprise_tsf_api_detail
	  tencentcloudenterprise_tsf_microservice_api_version
	  tencentcloudenterprise_tsf_repository
	  tencentcloudenterprise_tsf_pod_instances
	  tencentcloudenterprise_tsf_gateway_all_group_apis
	  tencentcloudenterprise_tsf_group_gateways
	  tencentcloudenterprise_tsf_usable_unit_namespaces
	  tencentcloudenterprise_tsf_group_instances
	  tencentcloudenterprise_tsf_container_group
	  tencentcloudenterprise_tsf_groups
	  tencentcloudenterprise_tsf_ms_api_list

	Resource
	  tencentcloudenterprise_tsf_cluster
	  tencentcloudenterprise_tsf_microservice
	  tencentcloudenterprise_tsf_application_config
	  tencentcloudenterprise_tsf_api_group
	  tencentcloudenterprise_tsf_namespace
	  tencentcloudenterprise_tsf_path_rewrite
	  tencentcloudenterprise_tsf_unit_rule
	  tencentcloudenterprise_tsf_task
	  tencentcloudenterprise_tsf_config_template
	  tencentcloudenterprise_tsf_api_rate_limit_rule
	  tencentcloudenterprise_tsf_application_release_config
	  tencentcloudenterprise_tsf_lane
	  tencentcloudenterprise_tsf_lane_rule
	  tencentcloudenterprise_tsf_group
	  tencentcloudenterprise_tsf_application
	  tencentcloudenterprise_tsf_application_public_config_release
	  tencentcloudenterprise_tsf_application_public_config
	  tencentcloudenterprise_tsf_application_file_config_release
	  tencentcloudenterprise_tsf_instances_attachment
	  tencentcloudenterprise_tsf_bind_api_group
	  tencentcloudenterprise_tsf_application_file_config
	  tencentcloudenterprise_tsf_enable_unit_rule
	  tencentcloudenterprise_tsf_deploy_container_group
	  tencentcloudenterprise_tsf_deploy_vm_group
	  tencentcloudenterprise_tsf_release_api_group
	  tencentcloudenterprise_tsf_operate_container_group
	  tencentcloudenterprise_tsf_operate_group
	  tencentcloudenterprise_tsf_unit_namespace

Cloud Workload Protection Platform (CWP)

	 Data Source
	tencentcloudenterprise_tdmq_rabbitmq_node_list
	tencentcloudenterprise_tdmq_rabbitmq_vip_instance

	 Resource
	tencentcloudenterprise_tdmq_rabbitmq_user
	tencentcloudenterprise_tdmq_rabbitmq_vip_instance
	tencentcloudenterprise_tdmq_rabbitmq_virtual_host
*/
package tencentcloud

import (
	"context"
	"net/url"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"

	common2 "terraform-provider-tencentcloudenterprise/sdk/common"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

const (
	PROVIDER_SECRET_ID                    = "TENCENTCLOUDENTERPRISE_SECRET_ID"
	PROVIDER_SECRET_KEY                   = "TENCENTCLOUDENTERPRISE_SECRET_KEY"
	PROVIDER_SECURITY_TOKEN               = "TENCENTCLOUDENTERPRISE_SECURITY_TOKEN"
	PROVIDER_REGION                       = "TENCENTCLOUDENTERPRISE_REGION"
	PROVIDER_PROTOCOL                     = "TENCENTCLOUDENTERPRISE_PROTOCOL"
	PROVIDER_DOMAIN                       = "TENCENTCLOUDENTERPRISE_DOMAIN"
	PROVIDER_ASSUME_ROLE_ARN              = "TENCENTCLOUDENTERPRISE_ASSUME_ROLE_ARN"
	PROVIDER_ASSUME_ROLE_SESSION_NAME     = "TENCENTCLOUDENTERPRISE_ASSUME_ROLE_SESSION_NAME"
	PROVIDER_ASSUME_ROLE_SESSION_DURATION = "TENCENTCLOUDENTERPRISE_ASSUME_ROLE_SESSION_DURATION"
	PROVIDER_CSP_DOMAIN                   = "TENCENTCLOUDENTERPRISE_CSP_DOMAIN"
	PROVIDER_COS_DOMAIN                   = "TENCENTCLOUDENTERPRISE_COS_DOMAIN"
)

type TencentCloudClient struct {
	apiV3Conn *connectivity.TencentCloudClient
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_ID, nil),
				Description: "This is the TencentCloud access key. It must be provided, but it can also be sourced from the `TENCENTCLOUDENTERPRISE_SECRET_ID` environment variable.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_KEY, nil),
				Description: "This is the TencentCloud secret key. It must be provided, but it can also be sourced from the `TENCENTCLOUDENTERPRISE_SECRET_KEY` environment variable.",
				Sensitive:   true,
			},
			"security_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECURITY_TOKEN, nil),
				Description: "TencentCloud Security Token of temporary access credentials. It can be sourced from the `TENCENTCLOUDENTERPRISE_SECURITY_TOKEN` environment variable. Notice: for supported products, please refer to: [temporary key supported products](https://intl.cloud.tencent.com/document/product/598/10588).",
				Sensitive:   true,
			},
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_REGION, nil),
				Description:  "This is the TencentCloud region. It must be provided, but it can also be sourced from the `TENCENTCLOUDENTERPRISE_REGION` environment variables. The default input value is ap-guangzhou.",
				InputDefault: "ap-guangzhou",
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_PROTOCOL, "HTTP"),
				ValidateFunc: validateAllowedStringValue([]string{"HTTP", "HTTPS"}),
				Description:  "The protocol of the API request. Valid values: `HTTP` and `HTTPS`. Default is `HTTPS`.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_DOMAIN, nil),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"csp_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_CSP_DOMAIN, ""),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"cos_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_COS_DOMAIN, ""),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"assume_role": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "The `assume_role` block. If provided, terraform will attempt to assume this role using the supplied credentials.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_arn": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_ARN, nil),
							Description: "The ARN of the role to assume. It can be sourced from the `TENCENTCLOUDENTERPRISE_ASSUME_ROLE_ARN`.",
						},
						"session_name": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_SESSION_NAME, nil),
							Description: "The session name to use when making the AssumeRole call. It can be sourced from the `TENCENTCLOUDENTERPRISE_ASSUME_ROLE_SESSION_NAME`.",
						},
						"session_duration": {
							Type:     schema.TypeInt,
							Required: true,
							DefaultFunc: func() (interface{}, error) {
								if v := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); v != "" {
									return strconv.Atoi(v)
								}
								return 7200, nil
							},
							ValidateFunc: validateIntegerInRange(0, 43200),
							Description:  "The duration of the session when making the AssumeRole call. Its value ranges from 0 to 43200(seconds), and default is 7200 seconds. It can be sourced from the `TENCENTCLOUDENTERPRISE_ASSUME_ROLE_SESSION_DURATION`.",
						},
						"policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A more restrictive policy when making the AssumeRole call. Its content must not contains `principal` elements. Notice: more syntax references, please refer to: [policies syntax logic](https://intl.cloud.tencent.com/document/product/598/10603).",
						},
					},
				},
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"tencentcloudenterprise_as_instances":                              dataSourceTencentCloudAsInstances(),
			"tencentcloudenterprise_as_last_activity":                          dataSourceTencentCloudAsLastActivity(),
			"tencentcloudenterprise_as_scaling_configs":                        dataSourceTencentCloudAsScalingConfigs(),
			"tencentcloudenterprise_as_scaling_groups":                         dataSourceTencentCloudAsScalingGroups(),
			"tencentcloudenterprise_as_scaling_policies":                       dataSourceTencentCloudAsScalingPolicies(),
			"tencentcloudenterprise_bms_instances":                             dataTencentCloudBmsInstances(),
			"tencentcloudenterprise_bms_placement_groups":                      dataTencentCloudBmsPlacementGroup(),
			"tencentcloudenterprise_bms_flavors":                               dataTencentCloudBmsFlavors(),
			"tencentcloudenterprise_brc_autobackup_policies":                   dataSourceTencentCloudBrcAutoBackupPolicies(),
			"tencentcloudenterprise_brc_backups":                               dataSourceTencentCloudBrcBackups(),
			"tencentcloudenterprise_brc_resource_backup_overview":              dataSourceTencentCloudBrcResourceBackupOverview(),
			"tencentcloudenterprise_cbs_snapshot_policies":                     dataSourceTencentCloudCbsSnapshotPolicies(),
			"tencentcloudenterprise_cbs_snapshots":                             dataSourceTencentCloudCbsSnapshots(),
			"tencentcloudenterprise_cbs_storages":                              dataSourceTencentCloudCbsStorages(),
			"tencentcloudenterprise_cbs_storages_set":                          dataSourceTencentCloudCbsStoragesSet(),
			"tencentcloudenterprise_cfs_access_groups":                         dataSourceTencentCloudCfsAccessGroups(),
			"tencentcloudenterprise_cfs_access_rules":                          dataSourceTencentCloudCfsAccessRules(),
			"tencentcloudenterprise_cfs_available_zone":                        dataSourceTencentCloudCfsAvailableZone(),
			"tencentcloudenterprise_cfs_file_system_clients":                   dataSourceTencentCloudCfsFileSystemClients(),
			"tencentcloudenterprise_cfs_file_systems":                          dataSourceTencentCloudCfsFileSystems(),
			"tencentcloudenterprise_cfs_mount_targets":                         dataSourceTencentCloudCfsMountTargets(),
			"tencentcloudenterprise_ckafka_acls":                               dataSourceTencentCloudCkafkaAcls(),
			"tencentcloudenterprise_ckafka_group":                              dataSourceTencentCloudCkafkaGroup(),
			"tencentcloudenterprise_ckafka_group_info":                         dataSourceTencentCloudCkafkaGroupInfo(),
			"tencentcloudenterprise_ckafka_group_offsets":                      dataSourceTencentCloudCkafkaGroupOffsets(),
			"tencentcloudenterprise_ckafka_instances":                          dataSourceTencentCloudCkafkaInstances(),
			"tencentcloudenterprise_ckafka_task_status":                        dataSourceTencentCloudCkafkaTaskStatus(),
			"tencentcloudenterprise_ckafka_topics":                             dataSourceTencentCloudCkafkaTopics(),
			"tencentcloudenterprise_ckafka_users":                              dataSourceTencentCloudCkafkaUsers(),
			"tencentcloudenterprise_ckafka_zone":                               dataSourceTencentCloudCkafkaZone(),
			"tencentcloudenterprise_clb_attachments":                           dataSourceTencentCloudClbServerAttachments(),
			"tencentcloudenterprise_clb_certificates":                          dataSourceTencentCloudClbCertificates(),
			"tencentcloudenterprise_clb_instance_by_cert_id":                   dataSourceTencentCloudClbInstanceByCertId(),
			"tencentcloudenterprise_clb_instance_detail":                       dataSourceTencentCloudClbInstanceDetail(),
			"tencentcloudenterprise_clb_instances":                             dataSourceTencentCloudClbInstances(),
			"tencentcloudenterprise_clb_listener_rules":                        dataSourceTencentCloudClbListenerRules(),
			"tencentcloudenterprise_clb_listeners":                             dataSourceTencentCloudClbListeners(),
			"tencentcloudenterprise_clb_redirections":                          dataSourceTencentCloudClbRedirections(),
			"tencentcloudenterprise_clb_resources":                             dataSourceTencentCloudClbResources(),
			"tencentcloudenterprise_clb_target_health":                         dataSourceTencentCloudClbTargetHealth(),
			"tencentcloudenterprise_cls_machine_group_configs":                 dataSourceTencentCloudClsMachineGroupConfigs(),
			"tencentcloudenterprise_cos_bucket_object":                         dataSourceTencentCloudCosBucketObject(),
			"tencentcloudenterprise_cos_buckets":                               dataSourceTencentCloudCosBuckets(),
			"tencentcloudenterprise_csp_bucket_object":                         dataSourceTencentCloudCspBucketObject(),
			"tencentcloudenterprise_csp_buckets":                               dataSourceTencentCloudCspBuckets(),
			"tencentcloudenterprise_cvm_disaster_recover_group_quota":          dataSourceTencentCloudCvmDisasterRecoverGroupQuota(),
			"tencentcloudenterprise_cvm_image":                                 dataSourceTencentCloudImage(),
			"tencentcloudenterprise_cvm_image_quota":                           dataSourceTencentCloudCvmImageQuota(),
			"tencentcloudenterprise_cvm_image_share_permission":                dataSourceTencentCloudCvmImageSharePermission(),
			"tencentcloudenterprise_cvm_images":                                dataSourceTencentCloudImages(),
			"tencentcloudenterprise_cvm_instance_types":                        dataSourceInstanceTypes(),
			"tencentcloudenterprise_cvm_instance_vnc_url":                      dataSourceTencentCloudCvmInstanceVncUrl(),
			"tencentcloudenterprise_cvm_instances":                             dataSourceTencentCloudInstances(),
			"tencentcloudenterprise_cvm_instances_modification":                dataSourceTencentCloudCvmInstancesModification(),
			"tencentcloudenterprise_cvm_instances_set":                         dataSourceTencentCloudInstancesSet(),
			"tencentcloudenterprise_cvm_key_pairs":                             dataSourceTencentCloudKeyPairs(),
			"tencentcloudenterprise_cvm_placement_groups":                      dataSourceTencentCloudPlacementGroups(),
			"tencentcloudenterprise_dc_access_points":                          dataSourceTencentCloudDcAccessPoints(),
			"tencentcloudenterprise_dc_instances":                              dataSourceTencentCloudDcInstances(),
			"tencentcloudenterprise_dcdb_accounts":                             dataSourceTencentCloudDcdbAccounts(),
			"tencentcloudenterprise_dcdb_database_objects":                     dataSourceTencentCloudDcdbDatabaseObjects(),
			"tencentcloudenterprise_dcdb_database_tables":                      dataSourceTencentCloudDcdbDatabaseTables(),
			"tencentcloudenterprise_dcdb_databases":                            dataSourceTencentCloudDcdbDatabases(),
			"tencentcloudenterprise_dcdb_instances":                            dataSourceTencentCloudDcdbInstances(),
			"tencentcloudenterprise_dcdb_parameters":                           dataSourceTencentCloudDcdbParameters(),
			"tencentcloudenterprise_dcdb_security_groups":                      dataSourceTencentCloudDcdbSecurityGroups(),
			"tencentcloudenterprise_dcdb_shards":                               dataSourceTencentCloudDcdbShards(),
			"tencentcloudenterprise_eip_address_quota":                         dataSourceTencentCloudEipAddressQuota(),
			"tencentcloudenterprise_eips":                                      dataSourceTencentCloudEips(),
			"tencentcloudenterprise_kms_keys":                                  dataSourceTencentCloudKmsKeys(),
			"tencentcloudenterprise_redis_backup":                              dataSourceTencentCloudRedisBackup(),
			"tencentcloudenterprise_redis_backup_download_info":                dataSourceTencentCloudRedisBackupDownloadInfo(),
			"tencentcloudenterprise_redis_instance_node_info":                  dataSourceTencentCloudRedisInstanceNodeInfo(),
			"tencentcloudenterprise_redis_instance_shards":                     dataSourceTencentCloudRedisInstanceShards(),
			"tencentcloudenterprise_redis_instance_task_list":                  dataSourceTencentCloudRedisInstanceTaskList(),
			"tencentcloudenterprise_redis_instances":                           dataSourceTencentRedisInstances(),
			"tencentcloudenterprise_redis_param_records":                       dataSourceTencentCloudRedisRecordsParam(),
			"tencentcloudenterprise_redis_zone_config":                         dataSourceTencentRedisZoneConfig(),
			"tencentcloudenterprise_ssm_secret_versions":                       dataSourceTencentCloudSsmSecretVersions(),
			"tencentcloudenterprise_ssm_secrets":                               dataSourceTencentCloudSsmSecrets(),
			"tencentcloudenterprise_tbase_instances":                           dataSourceTencentCloudTbaseInstances(),
			"tencentcloudenterprise_tbase_pg_instances":                        dataSourceTencentCloudTbasePGInstances(),
			"tencentcloudenterprise_tdmq_environments":                         dataSourceTencentCloudTdmqEnvironments(),
			"tencentcloudenterprise_tdmq_pulsar_environments":                  dataSourceTencentCloudTdmqPulsarEnvironments(),
			"tencentcloudenterprise_tdmq_pulsar_clusters":                      dataSourceTencentCloudTdmqPulsarCluster(),
			"tencentcloudenterprise_tdmq_rabbitmq_node_list":                   dataSourceTencentCloudTdmqRabbitmqNodeList(),
			"tencentcloudenterprise_tdmq_rabbitmq_vip_instance":                dataSourceTencentCloudTdmqRabbitmqVipInstance(),
			"tencentcloudenterprise_tdmq_rocketmq_cluster":                     dataSourceTencentCloudTdmqRocketmqCluster(),
			"tencentcloudenterprise_tdmq_rocketmq_group":                       dataSourceTencentCloudTdmqRocketmqGroup(),
			"tencentcloudenterprise_tdmq_rocketmq_messages":                    dataSourceTencentCloudTdmqRocketmqMessages(),
			"tencentcloudenterprise_tdmq_rocketmq_namespace":                   dataSourceTencentCloudTdmqRocketmqNamespace(),
			"tencentcloudenterprise_tdmq_rocketmq_role":                        dataSourceTencentCloudTdmqRocketmqRole(),
			"tencentcloudenterprise_tdmq_rocketmq_topic":                       dataSourceTencentCloudTdmqRocketmqTopic(),
			"tencentcloudenterprise_tke_kubernetes_available_cluster_versions": dataSourceTencentCloudKubernetesAvailableClusterVersions(),
			"tencentcloudenterprise_tke_kubernetes_charts":                     dataSourceTencentCloudKubernetesCharts(),
			"tencentcloudenterprise_tke_kubernetes_cluster_common_names":       datasourceTencentCloudKubernetesClusterCommonNames(),
			"tencentcloudenterprise_tke_kubernetes_clusters":                   dataSourceTencentCloudKubernetesClusters(),
			"tencentcloudenterprise_tsf_api_detail":                            dataSourceTencentCloudTsfApiDetail(),
			"tencentcloudenterprise_tsf_api_group":                             dataSourceTencentCloudTsfApiGroup(),
			"tencentcloudenterprise_tsf_application":                           dataSourceTencentCloudTsfApplication(),
			"tencentcloudenterprise_tsf_application_attribute":                 dataSourceTencentCloudTsfApplicationAttribute(),
			"tencentcloudenterprise_tsf_application_config":                    dataSourceTencentCloudTsfApplicationConfig(),
			"tencentcloudenterprise_tsf_application_file_config":               dataSourceTencentCloudTsfApplicationFileConfig(),
			"tencentcloudenterprise_tsf_application_public_config":             dataSourceTencentCloudTsfApplicationPublicConfig(),
			"tencentcloudenterprise_tsf_business_log_configs":                  dataSourceTencentCloudTsfBusinessLogConfigs(),
			"tencentcloudenterprise_tsf_cluster":                               dataSourceTencentCloudTsfCluster(),
			"tencentcloudenterprise_tsf_config_summary":                        dataSourceTencentCloudTsfConfigSummary(),
			"tencentcloudenterprise_tsf_delivery_config_by_group_id":           dataSourceTencentCloudTsfDeliveryConfigByGroupId(),
			"tencentcloudenterprise_tsf_delivery_configs":                      dataSourceTencentCloudTsfDeliveryConfigs(),
			"tencentcloudenterprise_tsf_gateway_all_group_apis":                dataSourceTencentCloudTsfGatewayAllGroupApis(),
			"tencentcloudenterprise_tsf_group_gateways":                        dataSourceTencentCloudTsfGroupGateways(),
			"tencentcloudenterprise_tsf_group_instances":                       dataSourceTencentCloudTsfGroupInstances(),
			"tencentcloudenterprise_tsf_microservice":                          dataSourceTencentCloudTsfMicroservice(),
			"tencentcloudenterprise_tsf_microservice_api_version":              dataSourceTencentCloudTsfMicroserviceApiVersion(),
			"tencentcloudenterprise_tsf_pod_instances":                         dataSourceTencentCloudTsfPodInstances(),
			"tencentcloudenterprise_tsf_public_config_summary":                 dataSourceTencentCloudTsfPublicConfigSummary(),
			"tencentcloudenterprise_tsf_repository":                            dataSourceTencentCloudTsfRepository(),
			"tencentcloudenterprise_tsf_usable_unit_namespaces":                dataSourceTencentCloudTsfUsableUnitNamespaces(),
			"tencentcloudenterprise_turbofs_p_groups":                          dataSourceTencentCloudTurbofsPGroups(),
			"tencentcloudenterprise_turbofs_rules":                             dataSourceTencentCloudTurbofsRules(),
			"tencentcloudenterprise_turbofs_file_systems":                      dataSourceTencentCloudTurbofsFileSystems(),
			"tencentcloudenterprise_turbofs_mount_targets":                     dataSourceTencentCloudTurbofsMountTargets(),
			"tencentcloudenterprise_vpc_account_attributes":                    dataSourceTencentCloudVpcAccountAttributes(),
			"tencentcloudenterprise_vpc_acls":                                  dataSourceTencentCloudVpcAcls(),
			"tencentcloudenterprise_vpc_address_template_groups":               dataSourceTencentCloudAddressTemplateGroups(),
			"tencentcloudenterprise_vpc_address_templates":                     dataSourceTencentCloudAddressTemplates(),
			"tencentcloudenterprise_vpc_classic_link_instances":                dataSourceTencentCloudVpcClassicLinkInstances(),
			"tencentcloudenterprise_vpc_cvm_instances":                         dataSourceTencentCloudVpcCvmInstances(),
			"tencentcloudenterprise_vpc_dnats":                                 dataSourceTencentCloudDnats(),
			"tencentcloudenterprise_vpc_enis":                                  dataSourceTencentCloudEnis(),
			"tencentcloudenterprise_vpc_gateway_flow_qos":                      dataSourceTencentCloudVpcGatewayFlowQos(),
			"tencentcloudenterprise_vpc_ha_vip_eip_attachments":                dataSourceTencentCloudHaVipEipAttachments(),
			"tencentcloudenterprise_vpc_ha_vips":                               dataSourceTencentCloudHaVips(),
			"tencentcloudenterprise_vpc_instances":                             dataSourceTencentCloudVpcInstances(),
			"tencentcloudenterprise_vpc_limits":                                dataSourceTencentCloudVpcLimits(),
			"tencentcloudenterprise_vpc_nat_gateways":                          dataSourceTencentCloudNatGateways(),
			"tencentcloudenterprise_vpc_net_detect_state_check":                dataSourceTencentCloudVpcNetDetectStateCheck(),
			"tencentcloudenterprise_vpc_net_detect_states":                     dataSourceTencentCloudVpcNetDetectStates(),
			"tencentcloudenterprise_vpc_private_ip_addresses":                  dataSourceTencentCloudVpcPrivateIpAddresses(),
			"tencentcloudenterprise_vpc_resource_dashboard":                    dataSourceTencentCloudVpcResourceDashboard(),
			"tencentcloudenterprise_vpc_route_tables":                          dataSourceTencentCloudVpcRouteTables(),
			"tencentcloudenterprise_vpc_security_group_limits":                 dataSourceTencentCloudVpcSecurityGroupLimits(),
			"tencentcloudenterprise_vpc_security_group_references":             dataSourceTencentCloudVpcSecurityGroupReferences(),
			"tencentcloudenterprise_vpc_security_groups":                       dataSourceTencentCloudSecurityGroups(),
			"tencentcloudenterprise_vpc_subnets":                               dataSourceTencentCloudVpcSubnets(),
			"tencentcloudenterprise_vpc_template_limits":                       dataSourceTencentCloudVpcTemplateLimits(),
			"tencentcloudenterprise_vpcdns_domains":                            dataSourceTencentCloudVpcDnsDomain(),
			"tencentcloudenterprise_vpcdns_records":                            dataSourceTencentCloudVpcDnsRecord(),
			"tencentcloudenterprise_vpcdns_forward_rules":                      dataSourceTencentCloudVpcDnsForwardRules(),
			"tencentcloudenterprise_cwp_machines_simple":                       dataSourceTencentCloudCwpMachinesSimple(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"tencentcloudenterprise_as_attachment":                           resourceTencentCloudAsAttachment(),
			"tencentcloudenterprise_as_lifecycle_hook":                       resourceTencentCloudAsLifecycleHook(),
			"tencentcloudenterprise_as_notification":                         resourceTencentCloudAsNotification(),
			"tencentcloudenterprise_as_protect_instances":                    resourceTencentCloudAsProtectInstances(),
			"tencentcloudenterprise_as_remove_instances":                     resourceTencentCloudAsRemoveInstances(),
			"tencentcloudenterprise_as_scaling_config":                       resourceTencentCloudAsScalingConfig(),
			"tencentcloudenterprise_as_scaling_group":                        resourceTencentCloudAsScalingGroup(),
			"tencentcloudenterprise_as_scaling_policy":                       resourceTencentCloudAsScalingPolicy(),
			"tencentcloudenterprise_as_schedule":                             resourceTencentCloudAsSchedule(),
			"tencentcloudenterprise_bms_instance":                            resourceTencentCloudBmsInstance(),
			"tencentcloudenterprise_bms_placement_group":                     resourceTencentCloudBmsPlacementGroup(),
			"tencentcloudenterprise_brc_activate_backup_service":             resourceTencentCloudBrcActivateBackupService(),
			"tencentcloudenterprise_brc_auto_backup_policy":                  resourceTencentCloudBrcAutoBackupPolicy(),
			"tencentcloudenterprise_brc_auto_backup_policy_binding":          resourceTencentCloudBrcAutoBackupPolicyBinding(),
			"tencentcloudenterprise_brc_backup_group":                        resourceTencentCloudBrcBackupGroup(),
			"tencentcloudenterprise_brc_backup_disk":                         resourceTencentCloudBrcBackupDisk(),
			"tencentcloudenterprise_brc_backup_cfs":                          resourceTencentCloudBrcBackupCfs(),
			"tencentcloudenterprise_brc_backup_resource":                     resourceTencentCloudBrcBackupResource(),
			"tencentcloudenterprise_cbs_snapshot":                            resourceTencentCloudCbsSnapshot(),
			"tencentcloudenterprise_cbs_snapshot_policy":                     resourceTencentCloudCbsSnapshotPolicy(),
			"tencentcloudenterprise_cbs_snapshot_policy_attachment":          resourceTencentCloudCbsSnapshotPolicyAttachment(),
			"tencentcloudenterprise_cbs_snapshot_share_permission":           resourceTencentCloudCbsSnapshotSharePermission(),
			"tencentcloudenterprise_cbs_storage":                             resourceTencentCloudCbsStorage(),
			"tencentcloudenterprise_cbs_storage_attachment":                  resourceTencentCloudCbsStorageAttachment(),
			"tencentcloudenterprise_cbs_storage_set":                         resourceTencentCloudCbsStorageSet(),
			"tencentcloudenterprise_cfs_access_group":                        resourceTencentCloudCfsAccessGroup(),
			"tencentcloudenterprise_cfs_access_rule":                         resourceTencentCloudCfsAccessRule(),
			"tencentcloudenterprise_cfs_auto_snapshot_policy":                resourceTencentCloudCfsAutoSnapshotPolicy(),
			"tencentcloudenterprise_cfs_auto_snapshot_policy_attachment":     resourceTencentCloudCfsAutoSnapshotPolicyAttachment(),
			"tencentcloudenterprise_cfs_file_system":                         resourceTencentCloudCfsFileSystem(),
			"tencentcloudenterprise_cfs_sign_up_cfs_service":                 resourceTencentCloudCfsSignUpCfsService(),
			"tencentcloudenterprise_cfs_snapshot":                            resourceTencentCloudCfsSnapshot(),
			"tencentcloudenterprise_cfw_nat_instance":                        resourceTencentCloudCfwNatInstance(),
			"tencentcloudenterprise_cfw_vpc_instance":                        resourceTencentCloudCfwVpcInstance(),
			"tencentcloudenterprise_cfw_nat_policy":                          resourceTencentCloudCfwNatPolicy(),
			"tencentcloudenterprise_cfw_vpc_policy":                          resourceTencentCloudCfwVpcPolicy(),
			"tencentcloudenterprise_cfw_block_ignore":                        resourceTencentCloudCfwBlockIgnore(),
			"tencentcloudenterprise_ckafka_acl":                              resourceTencentCloudCkafkaAcl(),
			"tencentcloudenterprise_ckafka_instance":                         resourceTencentCloudCkafkaInstance(),
			"tencentcloudenterprise_ckafka_topic":                            resourceTencentCloudCkafkaTopic(),
			"tencentcloudenterprise_ckafka_user":                             resourceTencentCloudCkafkaUser(),
			"tencentcloudenterprise_clb_attachment":                          resourceTencentCloudClbServerAttachment(),
			"tencentcloudenterprise_clb_certificates":                        resourceTencentCloudClbCertificate(),
			"tencentcloudenterprise_clb_customized_config":                   resourceTencentCloudClbCustomizedConfig(),
			"tencentcloudenterprise_clb_instance":                            resourceTencentCloudClbInstance(),
			"tencentcloudenterprise_clb_listener":                            resourceTencentCloudClbListener(),
			"tencentcloudenterprise_clb_listener_rule":                       resourceTencentCloudClbListenerRule(),
			"tencentcloudenterprise_clb_redirection":                         resourceTencentCloudClbRedirection(),
			"tencentcloudenterprise_clb_replace_cert":                        resourceTencentCloudClbReplaceCert(),
			"tencentcloudenterprise_clb_security_group_attachment":           resourceTencentCloudClbSecurityGroupAttachment(),
			"tencentcloudenterprise_cls_ckafka_consumer":                     resourceTencentCloudClsCkafkaConsumer(),
			"tencentcloudenterprise_cls_config":                              resourceTencentCloudClsConfig(),
			"tencentcloudenterprise_cls_config_attachment":                   resourceTencentCloudClsConfigAttachment(),
			"tencentcloudenterprise_cls_cos_recharge":                        resourceTencentCloudClsCosRecharge(),
			"tencentcloudenterprise_cls_cos_shipper":                         resourceTencentCloudClsCosShipper(),
			"tencentcloudenterprise_cls_export":                              resourceTencentCloudClsExport(),
			"tencentcloudenterprise_cls_index":                               resourceTencentCloudClsIndex(),
			"tencentcloudenterprise_cls_logset":                              resourceTencentCloudClsLogset(),
			"tencentcloudenterprise_cls_machine_group":                       resourceTencentCloudClsMachineGroup(),
			"tencentcloudenterprise_cls_topic":                               resourceTencentCloudClsTopic(),
			"tencentcloudenterprise_cos_bucket":                              resourceTencentCloudCosBucket(),
			"tencentcloudenterprise_cos_bucket_object":                       resourceTencentCloudCosBucketObject(),
			"tencentcloudenterprise_cos_bucket_policy":                       resourceTencentCloudCosBucketPolicy(),
			"tencentcloudenterprise_csp_bucket":                              resourceTencentCloudCspBucket(),
			"tencentcloudenterprise_csp_bucket_object":                       resourceTencentCloudCspBucketObject(),
			"tencentcloudenterprise_csp_bucket_policy":                       resourceTencentCloudCspBucketPolicy(),
			"tencentcloudenterprise_cvm_image":                               resourceTencentCloudImage(),
			"tencentcloudenterprise_cvm_image_share_permission":              resourceTencentCloudCvmImageSharePermission(),
			"tencentcloudenterprise_cvm_instance":                            resourceTencentCloudInstance(),
			"tencentcloudenterprise_cvm_instance_set":                        resourceTencentCloudInstanceSet(),
			"tencentcloudenterprise_cvm_key_pair":                            resourceTencentCloudKeyPair(),
			"tencentcloudenterprise_cvm_launch_template":                     resourceTencentCloudCvmLaunchTemplate(),
			"tencentcloudenterprise_cvm_placement_group":                     resourceTencentCloudPlacementGroup(),
			"tencentcloudenterprise_cvm_reboot_instance":                     resourceTencentCloudCvmRebootInstance(),
			"tencentcloudenterprise_cvm_renew_instance":                      resourceTencentCloudCvmRenewInstance(),
			"tencentcloudenterprise_cvm_security_group_attachment":           resourceTencentCloudCvmSecurityGroupAttachment(),
			"tencentcloudenterprise_cvm_sync_image":                          resourceTencentCloudCvmSyncImage(),
			"tencentcloudenterprise_dc_dcx":                                  resourceTencentCloudDcxInstance(),
			"tencentcloudenterprise_dc_instance":                             resourceTencentCloudDcInstance(),
			"tencentcloudenterprise_dcdb_account":                            resourceTencentCloudDcdbAccount(),
			"tencentcloudenterprise_dcdb_account_privileges":                 resourceTencentCloudDcdbAccountPrivileges(),
			"tencentcloudenterprise_dcdb_activate_hour_instance_operation":   resourceTencentCloudDcdbActivateHourInstanceOperation(),
			"tencentcloudenterprise_dcdb_cancel_dcn_job_operation":           resourceTencentCloudDcdbCancelDcnJobOperation(),
			"tencentcloudenterprise_dcdb_db_parameters":                      resourceTencentCloudDcdbDbParameters(),
			"tencentcloudenterprise_dcdb_db_sync_mode_config":                resourceTencentCloudDcdbDbSyncModeConfig(),
			"tencentcloudenterprise_dcdb_encrypt_attributes_config":          resourceTencentCloudDcdbEncryptAttributesConfig(),
			"tencentcloudenterprise_dcdb_flush_binlog_operation":             resourceTencentCloudDcdbFlushBinlogOperation(),
			"tencentcloudenterprise_dcdb_instance":                           resourceTencentCloudDcdbdbInstance(),
			"tencentcloudenterprise_dcdb_instance_config":                    resourceTencentCloudDcdbInstanceConfig(),
			"tencentcloudenterprise_dcdb_isolate_hour_instance_operation":    resourceTencentCloudDcdbIsolateHourInstanceOperation(),
			"tencentcloudenterprise_dcdb_security_group_attachment":          resourceTencentCloudDcdbSecurityGroupAttachment(),
			"tencentcloudenterprise_dcdb_switch_db_instance_ha_operation":    resourceTencentCloudDcdbSwitchDbInstanceHaOperation(),
			"tencentcloudenterprise_eip":                                     resourceTencentCloudEip(),
			"tencentcloudenterprise_eip_address_transform":                   resourceTencentCloudEipAddressTransform(),
			"tencentcloudenterprise_eip_association":                         resourceTencentCloudEipAssociation(),
			"tencentcloudenterprise_eip_normal_address_return":               resourceTencentCloudEipNormalAddressReturn(),
			"tencentcloudenterprise_kms_external_key":                        resourceTencentCloudKmsExternalKey(),
			"tencentcloudenterprise_kms_key":                                 resourceTencentCloudKmsKey(),
			"tencentcloudenterprise_redis_backup_config":                     resourceTencentCloudRedisBackupConfig(),
			"tencentcloudenterprise_redis_clear_instance_operation":          resourceTencentCloudRedisClearInstanceOperation(),
			"tencentcloudenterprise_redis_instance":                          resourceTencentCloudRedisInstance(),
			"tencentcloudenterprise_redis_param":                             resourceTencentCloudRedisParam(),
			"tencentcloudenterprise_redis_replica_readonly":                  resourceTencentCloudRedisReplicaReadonly(),
			"tencentcloudenterprise_redis_startup_instance_operation":        resourceTencentCloudRedisStartupInstanceOperation(),
			"tencentcloudenterprise_ssm_secret":                              resourceTencentCloudSsmSecret(),
			"tencentcloudenterprise_ssm_secret_version":                      resourceTencentCloudSsmSecretVersion(),
			"tencentcloudenterprise_tbase_instance":                          resourceTencentCloudTbaseInstance(),
			"tencentcloudenterprise_tbase_pg_instance":                       resourceTencentCloudTbasePGInstance(),
			"tencentcloudenterprise_tbase_pg_instance_vip":                   resourceTencentCloudTbasePGInstanceVip(),
			"tencentcloudenterprise_tdmq_instance":                           resourceTencentCloudTdmqInstance(),
			"tencentcloudenterprise_tdmq_pulsar_cluster":                     resourceTencentCloudTdmqPulsarCluster(),
			"tencentcloudenterprise_tdmq_pulsar_route":                       resourceTencentCloudTdmqPulsarRoute(),
			"tencentcloudenterprise_tdmq_pulsar_environment":                 resourceTencentCloudTdmqPulsarEnvironment(),
			"tencentcloudenterprise_tdmq_pulsar_topic":                       resourceTencentCloudTdmqPulsarTopic(),
			"tencentcloudenterprise_tdmq_pulsar_role":                        resourceTencentCloudTdmqPulsarRole(),
			"tencentcloudenterprise_tdmq_pulsar_environment_role_attachment": resourceTencentCloudTdmqPulsarEnvironmentRoleAttachment(),
			"tencentcloudenterprise_tdmq_namespace":                          resourceTencentCloudTdmqNamespace(),
			"tencentcloudenterprise_tdmq_namespace_role_attachment":          resourceTencentCloudTdmqNamespaceRoleAttachment(),
			"tencentcloudenterprise_tdmq_rabbitmq_user":                      resourceTencentCloudTdmqRabbitmqUser(),
			"tencentcloudenterprise_tdmq_rabbitmq_vip_instance":              resourceTencentCloudTdmqRabbitmqVipInstance(),
			"tencentcloudenterprise_tdmq_rabbitmq_virtual_host":              resourceTencentCloudTdmqRabbitmqVirtualHost(),
			"tencentcloudenterprise_tdmq_rocketmq_cluster":                   resourceTencentCloudTdmqRocketmqCluster(),
			"tencentcloudenterprise_tdmq_rocketmq_environment_role":          resourceTencentCloudTdmqRocketmqEnvironmentRole(),
			"tencentcloudenterprise_tdmq_rocketmq_group":                     resourceTencentCloudTdmqRocketmqGroup(),
			"tencentcloudenterprise_tdmq_rocketmq_namespace":                 resourceTencentCloudTdmqRocketmqNamespace(),
			"tencentcloudenterprise_tdmq_rocketmq_role":                      resourceTencentCloudTdmqRocketmqRole(),
			"tencentcloudenterprise_tdmq_rocketmq_topic":                     resourceTencentCloudTdmqRocketmqTopic(),
			"tencentcloudenterprise_tdmq_role":                               resourceTencentCloudTdmqRole(),
			"tencentcloudenterprise_tdmq_route":                              resourceTencentCloudTdmqRoute(),
			"tencentcloudenterprise_tdmq_send_rocketmq_message":              resourceTencentCloudTdmqSendRocketmqMessage(),
			"tencentcloudenterprise_tdmq_subscription_attachment":            resourceTencentCloudTdmqSubscriptionAttachment(),
			"tencentcloudenterprise_tdmq_topic":                              resourceTencentCloudTdmqTopic(),
			"tencentcloudenterprise_tke_kubernetes_cluster":                  resourceTencentCloudTkeCluster(),
			"tencentcloudenterprise_tke_kubernetes_cluster_namespace":        resourceTencentCloudTkeClusterNamespace(),
			"tencentcloudenterprise_tke_kubernetes_cluster_plugin":           resourceTencentCloudTkeClusterPlugin(),
			"tencentcloudenterprise_tke_kubernetes_cluster_secret":           resourceTencentCloudTkeClusterSecret(),
			"tencentcloudenterprise_tke_kubernetes_cluster_pv":               resourceTencentCloudTkeClusterPv(),
			"tencentcloudenterprise_tke_kubernetes_cluster_pvc":              resourceTencentCloudTkeClusterPvc(),
			"tencentcloudenterprise_tke_kubernetes_cluster_deploy":           resourceTencentCloudTkeClusterDeploy(),
			"tencentcloudenterprise_tke_kubernetes_cluster_affinity":         resourceTencentCloudTkeClusterAffinity(),
			"tencentcloudenterprise_tke_kubernetes_cluster_ing":              resourceTencentCloudTkeClusterIng(),
			"tencentcloudenterprise_tke_kubernetes_cluster_attachment":       resourceTencentCloudTkeClusterAttachment(),
			"tencentcloudenterprise_tke_kubernetes_cluster_endpoint":         resourceTencentCloudTkeClusterEndpoint(),
			"tencentcloudenterprise_tke_kubernetes_scale_worker":             resourceTencentCloudTkeScaleWorker(),
			"tencentcloudenterprise_tsf_cluster":                             resourceTencentCloudTsfCluster(),
			"tencentcloudenterprise_tsf_namespace":                           resourceTencentCloudTsfNamespace(),
			"tencentcloudenterprise_tsf_group":                               resourceTencentCloudTsfGroup(),
			"tencentcloudenterprise_tsf_application":                         resourceTencentCloudTsfApplication(),
			"tencentcloudenterprise_tsf_application_config":                  resourceTencentCloudTsfApplicationConfig(),
			"tencentcloudenterprise_tsf_application_release_config":          resourceTencentCloudTsfApplicationReleaseConfig(),
			"tencentcloudenterprise_tsf_application_file_config":             resourceTencentCloudTsfApplicationFileConfig(),
			"tencentcloudenterprise_tsf_application_file_config_release":     resourceTencentCloudTsfApplicationFileConfigRelease(),
			"tencentcloudenterprise_tsf_application_public_config":           resourceTencentCloudTsfApplicationPublicConfig(),
			"tencentcloudenterprise_tsf_application_public_config_release":   resourceTencentCloudTsfApplicationPublicConfigRelease(),
			"tencentcloudenterprise_tsf_config_template":                     resourceTencentCloudTsfConfigTemplate(),
			"tencentcloudenterprise_tsf_lane":                                resourceTencentCloudTsfLane(),
			"tencentcloudenterprise_tsf_lane_rule":                           resourceTencentCloudTsfLaneRule(),
			"tencentcloudenterprise_tsf_api_group":                           resourceTencentCloudTsfApiGroup(),
			"tencentcloudenterprise_tsf_bind_api_group":                      resourceTencentCloudTsfBindApiGroup(),
			"tencentcloudenterprise_tsf_microservice":                        resourceTencentCloudTsfMicroservice(),
			"tencentcloudenterprise_tsf_api_rate_limit_rule":                 resourceTencentCloudTsfApiRateLimitRule(),
			"tencentcloudenterprise_tsf_enable_unit_rule":                    resourceTencentCloudTsfEnableUnitRule(),
			"tencentcloudenterprise_tsf_instances_attachment":                resourceTencentCloudTsfInstancesAttachment(),
			"tencentcloudenterprise_tsf_path_rewrite":                        resourceTencentCloudTsfPathRewrite(),
			"tencentcloudenterprise_tsf_task":                                resourceTencentCloudTsfTask(),
			"tencentcloudenterprise_tsf_unit_rule":                           resourceTencentCloudTsfUnitRule(),
			"tencentcloudenterprise_turbofs_sign_up_service":                 resourceTencentCloudTurbofsSignUpService(),
			"tencentcloudenterprise_turbofs_file_system":                     resourceTencentCloudTurbofsFileSystem(),
			"tencentcloudenterprise_turbofs_p_group":                         resourceTencentCloudTurbofsPGroup(),
			"tencentcloudenterprise_turbofs_rule":                            resourceTencentCloudTurbofsRule(),
			"tencentcloudenterprise_turbofs_auto_snapshot_policy":            resourceTencentCloudTurbofsAutoSnapshotPolicy(),
			"tencentcloudenterprise_turbofs_auto_snapshot_policy_attachment": resourceTencentCloudTurbofsAutoSnapshotPolicyAttachment(),
			"tencentcloudenterprise_vpc":                                     resourceTencentCloudVpcInstance(),
			"tencentcloudenterprise_vpc_acl":                                 resourceTencentCloudVpcACL(),
			"tencentcloudenterprise_vpc_acl_attachment":                      resourceTencentCloudVpcAclAttachment(),
			"tencentcloudenterprise_vpc_address_template":                    resourceTencentCloudAddressTemplate(),
			"tencentcloudenterprise_vpc_address_template_group":              resourceTencentCloudAddressTemplateGroup(),
			"tencentcloudenterprise_vpc_classic_link_attachment":             resourceTencentCloudVpcClassicLinkAttachment(),
			"tencentcloudenterprise_vpc_dc_gateway":                          resourceTencentCloudDcGatewayInstance(),
			"tencentcloudenterprise_vpc_peer_connect_manager":                resourceTencentCloudVpcPeerConnectManager(),
			"tencentcloudenterprise_vpc_peer_connect_ex_manager":             resourceTencentCloudVpcPeerConnectExManager(),
			"tencentcloudenterprise_vpc_dnat":                                resourceTencentCloudDnat(),
			"tencentcloudenterprise_vpc_enable_end_point_connect":            resourceTencentCloudVpcEnableEndPointConnect(),
			"tencentcloudenterprise_vpc_end_point":                           resourceTencentCloudVpcEndPoint(),
			"tencentcloudenterprise_vpc_end_point_service":                   resourceTencentCloudVpcEndPointService(),
			"tencentcloudenterprise_vpc_eni":                                 resourceTencentCloudEni(),
			"tencentcloudenterprise_vpc_eni_attachment":                      resourceTencentCloudEniAttachment(),
			"tencentcloudenterprise_vpc_eni_sg_attachment":                   resourceTencentCloudEniSgAttachment(),
			"tencentcloudenterprise_vpc_ha_vip":                              resourceTencentCloudHaVip(),
			"tencentcloudenterprise_vpc_ha_vip_eip_attachment":               resourceTencentCloudHaVipEipAttachment(),
			"tencentcloudenterprise_vpc_ipv6_address_bandwidth":              resourceTencentCloudIpv6AddressBandwidth(),
			"tencentcloudenterprise_vpc_ipv6_cidr_block":                     resourceTencentCloudVpcIpv6CidrBlock(),
			"tencentcloudenterprise_vpc_ipv6_eni_address":                    resourceTencentCloudVpcIpv6EniAddress(),
			"tencentcloudenterprise_vpc_ipv6_subnet_cidr_block":              resourceTencentCloudVpcIpv6SubnetCidrBlock(),
			"tencentcloudenterprise_vpc_nat_gateway":                         resourceTencentCloudNatGateway(),
			"tencentcloudenterprise_vpc_net_detect":                          resourceTencentCloudVpcNetDetect(),
			"tencentcloudenterprise_vpc_route_table":                         resourceTencentCloudVpcRouteTable(),
			"tencentcloudenterprise_vpc_route_table_entry":                   resourceTencentCloudVpcRouteEntry(),
			"tencentcloudenterprise_vpc_security_group":                      resourceTencentCloudSecurityGroup(),
			"tencentcloudenterprise_vpc_security_group_lite_rule":            resourceTencentCloudSecurityGroupLiteRule(),
			"tencentcloudenterprise_vpc_security_group_rule":                 resourceTencentCloudSecurityGroupRule(),
			"tencentcloudenterprise_vpc_security_group_rule_set":             resourceTencentCloudSecurityGroupRuleSet(),
			"tencentcloudenterprise_vpc_subnet":                              resourceTencentCloudVpcSubnet(),
			"tencentcloudenterprise_vpcdns_domain":                           resourceTencentCloudVpcDnsDomain(),
			"tencentcloudenterprise_vpcdns_forward_rule":                     resourceTencentCloudVpcDnsForwardRule(),
			"tencentcloudenterprise_vpcdns_record":                           resourceTencentCloudVpcDnsRecord(),
			"tencentcloudenterprise_cwp_license_order":                       ResourceTencentCloudCwpLicenseOrder(),
			"tencentcloudenterprise_cwp_license_bind_attachment":             ResourceTencentCloudCwpLicenseBindAttachment(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	secretId := d.Get("secret_id").(string)
	secretKey := d.Get("secret_key").(string)
	securityToken := d.Get("security_token").(string)
	region := d.Get("region").(string)
	protocol := d.Get("protocol").(string)
	domain := d.Get("domain").(string)
	cspDomain := d.Get("csp_domain").(string)
	cosDomain := d.Get("cos_domain").(string)
	// standard client
	var tcClient TencentCloudClient
	tcClient.apiV3Conn = &connectivity.TencentCloudClient{
		Credential: common.NewTokenCredential(
			secretId,
			secretKey,
			securityToken,
		),
		CredentialTce: common2.NewTokenCredential(
			secretId,
			secretKey,
			securityToken,
		),
		Region:    region,
		Protocol:  protocol,
		Domain:    domain,
		CspDomain: cspDomain,
		CosDomain: cosDomain,
	}

	envRoleArn := os.Getenv(PROVIDER_ASSUME_ROLE_ARN)
	envSessionName := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_NAME)

	// get assume role from env
	if envRoleArn != "" && envSessionName != "" {
		var assumeRoleSessionDuration int
		if envSessionDuration := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); envSessionDuration != "" {
			var err error
			assumeRoleSessionDuration, err = strconv.Atoi(envSessionDuration)
			if err != nil {
				return nil, diag.FromErr(err)
			}
		}
		if assumeRoleSessionDuration == 0 {
			assumeRoleSessionDuration = 7200
		}

		_ = genClientWithSTS(&tcClient, envRoleArn, envSessionName, assumeRoleSessionDuration, "")
	}

	// get assume role from tf config
	assumeRoleList := d.Get("assume_role").(*schema.Set).List()
	if len(assumeRoleList) == 1 {
		assumeRole := assumeRoleList[0].(map[string]interface{})
		assumeRoleArn := assumeRole["role_arn"].(string)
		assumeRoleSessionName := assumeRole["session_name"].(string)
		assumeRoleSessionDuration := assumeRole["session_duration"].(int)
		assumeRolePolicy := assumeRole["policy"].(string)

		_ = genClientWithSTS(&tcClient, assumeRoleArn, assumeRoleSessionName, assumeRoleSessionDuration, assumeRolePolicy)
	}
	return &tcClient, nil
}

func genClientWithSTS(tcClient *TencentCloudClient, assumeRoleArn, assumeRoleSessionName string,
	assumeRoleSessionDuration int, assumeRolePolicy string) error {
	// applying STS credentials
	request := sts.NewAssumeRoleRequest()
	request.RoleArn = helper.String(assumeRoleArn)
	request.RoleSessionName = helper.String(assumeRoleSessionName)
	request.DurationSeconds = helper.IntUint64(assumeRoleSessionDuration)
	if assumeRolePolicy != "" {
		request.Policy = helper.String(url.QueryEscape(assumeRolePolicy))
	}
	ratelimit.Check(request.GetAction())
	response, err := tcClient.apiV3Conn.UseStsClient().AssumeRole(request)
	if err != nil {
		return err
	}
	// using STS credentials
	tcClient.apiV3Conn.Credential = common.NewTokenCredential(
		*response.Response.Credentials.TmpSecretId,
		*response.Response.Credentials.TmpSecretKey,
		*response.Response.Credentials.Token,
	)
	return nil
}
