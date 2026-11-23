---
subcategory: "MARIADB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_mariadb_db_instances"
sidebar_current: "docs-tencentcloudenterprise-datasource-mariadb_db_instances"
description: |-
  Queries mariadb db instances.
---

# tencentcloudenterprise_mariadb_db_instances

Queries mariadb db instances.

## Example Usage

```hcl
data "tencentcloudenterprise_mariadb_db_instances" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `instance_ids` - (Optional, Set: [`String`]) instance ids.
* `project_ids` - (Optional, Set: [`Int`]) project ids.
* `result_output_file` - (Optional, String) Used to save results.
* `search_name` - (Optional, String) instance name or vip.
* `subnet_id` - (Optional, String) subnet id.
* `vpc_id` - (Optional, String) vpc id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `instances` - instances info.
  * `db_version_id` - db version id.
  * `instance_id` - instance id.
  * `instance_name` - instance name.
  * `memory` - meory of instance.
  * `project_id` - project id.
  * `region` - region.
  * `resource_tags` - resource tags.
    * `tag_key` - tag key.
    * `tag_value` - tag value.
  * `storage` - storage of instance.
  * `subnet_id` - subnet id.
  * `vpc_id` - vpc id.
  * `zone` - available zone.

