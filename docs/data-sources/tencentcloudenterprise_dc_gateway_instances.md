---
subcategory: "Direct Connect(DC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dc_gateway_instances"
sidebar_current: "docs-tencentcloudenterprise-datasource-dc_gateway_instances"
description: |-
  Use this data source to query detailed information of direct connect gateway instances.
---

# tencentcloudenterprise_dc_gateway_instances

Use this data source to query detailed information of direct connect gateway instances.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn" "main" {
  name        = "ci-temp-test-ccn"
  description = "ci-temp-test-ccn-des"
  qos         = "AG"
}

resource "tencentcloudenterprise_vpc_dc_gateway" "ccn_main" {
  name                = "ci-cdg-ccn-test"
  network_instance_id = tencentcloudenterprise_ccn.main.id
  network_type        = "CCN"
  gateway_type        = "NORMAL"
}

#You need to sleep for a few seconds because there is a cache on the server
data "tencentcloudenterprise_dc_gateway_instances" "name_select" {
  name = tencentcloudenterprise_vpc_dc_gateway.ccn_main.name
}

data "tencentcloudenterprise_dc_gateway_instances" "id_select" {
  dcg_id = tencentcloudenterprise_vpc_dc_gateway.ccn_main.id
}
```

## Argument Reference

The following arguments are supported:

* `dcg_id` - (Optional, String) ID of the DCG to be queried.
* `name` - (Optional, String) Name of the DCG to be queried.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `instance_list` - Information list of the DCG.
  * `cnn_route_type` - Type of CCN route. Valid values: `BGP` and `STATIC`.
  * `create_time` - Creation time of resource.
  * `dcg_id` - ID of the DCG.
  * `dcg_ip` - IP of the DCG.
  * `enable_bgp` - Indicates whether the BGP is enabled.
  * `gateway_type` - Type of the gateway. Valid values: `NORMAL` and `NAT`.
  * `mode_type` - Whether to publish VPC CIDR to CCN. Valid values: `standard` and `exquisite`.
  * `name` - Name of the DCG.
  * `network_instance_id` - ID of the associated network instance.
  * `network_type` - Type of associated network. Valid values: `VPC` and `CCN`.

