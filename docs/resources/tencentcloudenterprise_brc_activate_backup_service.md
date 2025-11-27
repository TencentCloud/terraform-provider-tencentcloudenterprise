---
subcategory: "Backup and Recovery Center(BRC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_brc_activate_backup_service"
sidebar_current: "docs-tencentcloudenterprise-resources-brc_activate_backup_service"
description: |-
  Provides a resource to activate a brc backup service
---

# tencentcloudenterprise_brc_activate_backup_service

Provides a resource to activate a brc backup service

## Example Usage

```hcl
resource "tencentcloudenterprise_brc_activate_backup_service" "activate_backup_service" {
  resource_type = "COS"
}
```

## Argument Reference

The following arguments are supported:

* `resource_type` - (Required, String, ForceNew) The resource type to be backed up. Valid values: 'CFS','COS', 'CSP', 'DISK', 'INSTANCE', 'MySQL_MariaDB', 'TDSQL_MySQL'.
* `brc_backup_service_is_open` - (Optional, Bool) Current status of the BRC backup service for this user. Valid values: true, false.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



