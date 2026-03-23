---
subcategory: "Cloud Workload Protection Platform(CWP)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cwp_license_bind_attachment"
sidebar_current: "docs-tencentcloudenterprise-resource-cwp_license_bind_attachment"
description: |-
  Provides a CWP license bind attachment resource.
---

# tencentcloudenterprise_cwp_license_bind_attachment

Provides a CWP license bind attachment resource.

~> **NOTE:** The license_id is automatically queried from the license order and is not a user input parameter.

## Example Usage

```hcl
# Basic CWP license bind attachment
resource "tencentcloudenterprise_cwp_license_bind_attachment" "example" {
  resource_id  = "cwplic-442d44a0"
  license_type = 5
  quuid        = "5c987cf1-b3c9-4b5c-ad45-6787d51f34d7"
}

# Batch bind multiple machines
resource "tencentcloudenterprise_cwp_license_bind_attachment" "batch_bind" {
  for_each = toset(["5c987cf1-b3c9-4b5c-ad45-6787d51f34d7", "another-quuid"])

  resource_id  = "cwplic-442d44a0"
  license_type = 5
  quuid        = each.value
}
```

## Argument Reference

The following arguments are supported:

* `license_type` - (Required, Int, ForceNew) License type: 0=CWP Pro Pay-as-you-go, 1=CWP Pro Monthly, 5=CWP Ultimate Monthly.
* `quuid` - (Required, String, ForceNew) Machine unique identifier (UUID).
* `resource_id` - (Required, String, ForceNew) Resource ID of the license.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `agent_status` - Agent status.
* `is_switch_bind` - Allow switch binding.
* `is_unbind` - Allow unbinding.
* `license_id` - License ID (automatically queried from license order).
* `machine_ip` - Machine IP.
* `machine_name` - Machine name.
* `machine_wan_ip` - Machine WAN IP.
* `uuid` - Machine UUID.

## Import

tencentcloudenterprise_cwp_license_bind_attachment can be imported using the id, e.g.

```
CWP license bind attachment can be imported using the resource_id#quuid#license_type, e.g.

```
$ terraform import tencentcloudenterprise_cwp_license_bind_attachment.example cwplic-442d44a0#5c987cf1-b3c9-4b5c-ad45-6787d51f34d7#5
```
```

