---
subcategory: "Direct Connect(DC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dc_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-dc_instance"
description: |-
  Provides a resource to create a dc instance
---

# tencentcloudenterprise_dc_instance

Provides a resource to create a dc instance

## Example Usage

```hcl
resource "tencentcloudenterprise_dc_instance" "instance" {
  access_point_id                  = "ap-shenzhen-b-ft"
  bandwidth                        = 10
  customer_contact_number          = "0"
  direct_connect_name              = "terraform-for-test"
  line_operator                    = "In-houseWiring"
  tencentcloudenterprise_port_type = "10GBase-LR"
  sign_law                         = true
  vlan                             = -1
}
```

## Argument Reference

The following arguments are supported:

* `access_point_id` - (Required, String, ForceNew) Access point of connection.The selected access point must exist and be available.
* `bandwidth` - (Required, Int) Connection port bandwidth in Mbps. Value range: [2,10240]. Default value: 1000.
* `customer_contact_mail` - (Required, String) Email address of connection applicant, which is obtained from the account system by default.
* `customer_contact_number` - (Required, String) Contact number of connection applicant. Format：Area code: 1-3 digits, phone number: 5-15 digits (e.g.: 1-5551234567).
* `customer_name` - (Required, String) Name of connection applicant, which is obtained from the account system by default.
* `direct_connect_name` - (Required, String) Connection name.
* `idc_city` - (Required, String, ForceNew) City where the local data center is located.
* `idc_port_type` - (Required, String, ForceNew) IDC-side port type for physical dedicated line access. Values: 100Base-T (100M electrical port), 1000Base-T (default, 1000M electrical port), 1000Base-LX (1000M single-mode optical port, 10km), 10GBase-T (10G electrical port), 10GBase-LR (10G single-mode optical port, 10km, default value).
* `line_operator` - (Required, String, ForceNew) ISP that provides connections.
* `location` - (Required, String, ForceNew) Local IDC location.
* `tencentcloudenterprise_port_type` - (Required, String, ForceNew) Port type of connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface), 1000Base-LX (1-Gigabit single-module optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-module optical Ethernet interface; 10 KM). Default value: 1000Base-LX.
* `is_share` - (Optional, Bool) Whether the direct connect instance is shared. Can only be modified after creation via updates.
* `redundant_direct_connect_id` - (Optional, String, ForceNew) ID of redundant connection.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `apply_id` - Application ID of the direct connect.
* `created_time` - Creation time of the direct connect.
* `enabled_time` - Enabled time of the direct connect.
* `fault_report_contact_number` - Fault report contact number.
* `fault_report_contact_person` - Fault report contact person.
* `state` - Direct connect state. Possible values: PENDING, REJECTED, ALLOCATED, AVAILABLE, DELETING, DELETED.

## Import

tencentcloudenterprise_dc_instance can be imported using the id, e.g.

```
dc instance can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_dc_instance.instance dc_id
```
```

