---
subcategory: "BH"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bh_devices"
sidebar_current: "docs-tencentcloudenterprise-datasource-bh_devices"
description: |-
  Use this data source to query detailed information of BH devices
---

# tencentcloudenterprise_bh_devices

Use this data source to query detailed information of BH devices

## Example Usage

```hcl
data "tencentcloudenterprise_bh_devices" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `ap_code_set` - (Optional, Set: [`String`]) Region code collection.
* `authorized_user_id_set` - (Optional, Set: [`Int`]) User ID collection with access to this asset.
* `department_id` - (Optional, String) Filter condition, can filter by department ID.
* `filters` - (Optional, List) Filter array.
* `id_set` - (Optional, Set: [`Int`]) Asset ID collection.
* `ip` - (Optional, String) Not currently used.
* `kind_set` - (Optional, Set: [`Int`]) Can filter by multiple types, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer.
* `kind` - (Optional, Int) Operating system type, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer.
* `managed_account` - (Optional, String) Whether the asset contains managed accounts. 1, contains; 0, does not contain.
* `name` - (Optional, String) Asset name or asset IP, fuzzy search.
* `resource_id_set` - (Optional, Set: [`String`]) Filter condition, asset-bound bastion host service ID collection.
* `result_output_file` - (Optional, String) Used to save results.
* `tag_filters` - (Optional, List) Filter condition, can filter by tag key and tag value. If both tag key and tag value filter conditions are specified, they have an "AND" relationship.

The `filters` object supports the following:

* `name` - (Required, String) Field to filter. Support: BindingStatus, InstanceId, DeviceAccount, VpcId, DomainId, ResourceId, Name, Ip, ManageDimension.
* `values` - (Required, Set) Filter values for the field.

The `tag_filters` object supports the following:

* `tag_key` - (Required, String) Tag key.
* `tag_value` - (Optional, Set) Tag value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `device_set` - Asset information list.
  * `account_count` - Number of accounts bound to the asset.
  * `ap_code` - Region code.
  * `department` - Department to which the asset belongs.
    * `id` - Department ID.
    * `manager_users` - Administrator users.
      * `manager_id` - Administrator ID.
      * `manager_name` - Administrator name.
    * `managers` - Department administrator account ID.
    * `name` - Department name, 1 - 256 characters.
  * `domain_id` - Network domain ID.
  * `domain_name` - Network domain name.
  * `enable_ssl` - Whether SSL is enabled, only supports Redis assets, 0: disabled 1: enabled.
  * `group_set` - Asset group list to which it belongs.
    * `count` - Count.
    * `department` - Department information to which it belongs.
      * `id` - Department ID.
      * `manager_users` - Administrator users.
        * `manager_id` - Administrator ID.
        * `manager_name` - Administrator name.
      * `managers` - Department administrator account ID.
      * `name` - Department name, 1 - 256 characters.
    * `id` - Group ID.
    * `name` - Group name.
  * `id` - Asset ID.
  * `instance_id` - Instance ID, corresponding to CVM, CDB and other instance IDs.
  * `ioa_id` - Resource ID on the IOA side.
  * `ip_port_set` - Multi-node information for database assets.
  * `kind` - Asset type 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer.
  * `manage_account_id` - K8S cluster management account ID.
  * `manage_dimension` - K8S cluster management dimension, 1-cluster, 2-namespace, 3-workload.
  * `name` - Asset name.
  * `namespace` - K8S cluster namespace.
  * `os_name` - Operating system name.
  * `port` - Management port.
  * `private_ip` - Private IP.
  * `public_ip` - Public IP.
  * `resource` - Bastion host service information, note that it is null when no service is bound.
    * `ap_code` - Region code.
    * `cdc_cluster_id` - CDC cluster ID.
    * `cidr_block` - CIDR block of the subnet where the service is deployed.
    * `clb_set` - Bastion host resource load balancer.
      * `clb_ip` - Load balancer IP.
    * `client_access` - 1 default value, client access enabled, 0 client access disabled.
    * `create_time` - Resource creation time.
    * `deploy_model` - Deployment mode, default 0, 0-cvm 1-tke.
    * `deployed` - Whether deployed.
    * `domain_count` - Number of network domains.
    * `expire_time` - Expiration time.
    * `expired` - Whether expired.
    * `extend_points` - Extension points.
    * `external_access` - 1 default value, external access enabled, 0 external access disabled.
    * `intranet_access` - 0 default value, non-intranet access, 1 intranet access, 2 intranet access opening, 3 intranet access closing.
    * `intranet_private_ip_set` - IP addresses for intranet access.
    * `intranet_subnet_id` - Subnet ID for enabling intranet access.
    * `intranet_vpc_cidr` - CIDR block of the VPC for enabling intranet access.
    * `intranet_vpc_id` - VPC for enabling intranet access.
    * `ioa_resource_id` - Zero trust instance ID corresponding to the bastion host instance.
    * `ioa_resource` - 0 default value, 0-free version IOA, 1-paid version IOA.
    * `lb_vip_isp` - ISP information.
    * `log_delivery_args` - Log delivery specification information.
    * `log_delivery` - Log delivery specification information.
    * `module_set` - Advanced feature list enabled for the service.
    * `nodes` - Number of assets corresponding to the service specification.
    * `open_clb_id` - Shared CLB ID.
    * `package_bandwidth` - Number of bandwidth extension packages (4M).
    * `package_ioa_bandwidth` - Number of zero trust bastion host bandwidth extension packages.
    * `package_ioa_user_count` - Number of zero trust bastion host user extension packages.
    * `package_node` - Number of authorization point extension packages (50 points).
    * `pid` - Pricing model ID.
    * `private_ip_set` - Internal IP.
    * `product_code` - Product code.
    * `public_ip_set` - External IP.
    * `renew_flag` - Auto-renewal flag, 0 - default state, 1 - auto-renewal, 2 - explicitly not auto-renewal.
    * `resource_id` - Service instance ID, such as bh-saas-s3ed4r5e.
    * `resource_name` - Service instance name.
    * `share_clb` - Whether to share CLB.
    * `status` - Resource status, 0 - not initialized, 1 - normal, 2 - isolated, 3 - destroyed, 4 - initialization failed, 5 - initializing.
    * `sub_product_code` - Sub-product code.
    * `subnet_id` - Subnet ID where the service is deployed.
    * `subnet_name` - Subnet name where the service is deployed.
    * `sv_args` - Service instance specification information.
    * `trial` - 0 non-trial version, 1 trial version.
    * `tui_cmd_port` - Linux asset command line operation port.
    * `tui_direct_port` - Linux asset direct connection port.
    * `used_domain_count` - Number of network domains already used.
    * `used_nodes` - Number of used authorization points.
    * `vpc_cidr_block` - CIDR block of the VPC where the service is deployed.
    * `vpc_id` - VPC ID.
    * `vpc_name` - VPC name where the service is deployed.
    * `web_access` - 1 default value, web access enabled, 0 web access disabled.
    * `zone` - Availability zone.
  * `ssl_cert_name` - Name of the uploaded SSL certificate.
  * `subnet_id` - Subnet ID.
  * `sync_pod_count` - Number of synchronized pods in K8S cluster.
  * `total_pod_count` - Total number of pods in K8S cluster.
  * `vpc_id` - VPC ID.
  * `workload` - K8S cluster workload.

