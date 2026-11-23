---
subcategory: "Cloud Load Balancer(CLB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_clb_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-clb_instance"
description: |-
  Manages a CLB instance.
---

# tencentcloudenterprise_clb_instance

Manages a CLB instance.

## Example Usage

```hcl
resource "tencentcloudenterprise_clb_instance" "example" {
  network_type = "INTERNAL"
}
```

## Argument Reference

The following arguments are supported:

* `network_type` - (Required, String, ForceNew) Type of CLB instance. Valid values: `OPEN` and `INTERNAL`.
* `address_ip_version` - (Optional, String) IP version, only applicable to open CLB. Valid values are `IPV4`, `IPV6` and `IPv6FullChain`, Default is IPV4.
* `clb_name` - (Optional, String) Name of the CLB. The name can only contain English letters, numbers, underscore and hyphen '-'.
* `eip_address_id` - (Optional, String) The unique ID of EIP, e.g. `eip-11112222`. Only available for INTERNAL CLB instance to bind EIP.
* `internet_bandwidth_max_out` - (Optional, Int) Max bandwidth out, only applicable to open CLB. Valid value ranges is [1, 2000]. Unit is Mbps.
* `internet_charge_type` - (Optional, String) Internet charge type, only applicable to open CLB. Valid values are only `TRAFFIC_POSTPAID_BY_HOUR`.
* `log_set_id` - (Optional, String) The id of log set.
* `log_topic_id` - (Optional, String) The id of log topic.
* `master_zone_id` - (Optional, String) Setting master zone id of cross available zone disaster recovery.
* `project_id` - (Optional, Int, ForceNew) ID of the project within the CLB instance, `0` - Default Project.
* `security_groups` - (Optional, List: [`String`]) Security groups of the CLB instance. Supports both `OPEN` and `INTERNAL` CLBs.
* `slave_zone_id` - (Optional, String) Setting slave zone id of cross available zone disaster recovery. this zone will undertake traffic when the master is down.
* `subnet_id` - (Optional, String, ForceNew) Subnet ID of the CLB. Effective only for CLB within the VPC.
* `tags` - (Optional, Map) The available tags within this CLB.
* `target_region_info_region` - (Optional, String) Region of the target region for cross-region CLB.
* `target_region_info_vpc_id` - (Optional, String) VPC ID of the target region for cross-region CLB.
* `vpc_id` - (Optional, String, ForceNew) VPC ID of the CLB.
* `zone_id` - (Optional, String) Available zone id, only applicable to open CLB.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `clb_vips` - The virtual service address table of the CLB.
* `instance_id` - CLB instance ID.
* `vip_isp` - Network operator, only applicable to open CLB. If this ISP is specified, network billing method can only use the bandwidth package billing (BANDWIDTH_PACKAGE).

