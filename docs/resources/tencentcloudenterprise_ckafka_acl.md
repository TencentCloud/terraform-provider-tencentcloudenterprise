---
subcategory: "Cloud Kafka(CKafka)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ckafka_acl"
sidebar_current: "docs-tencentcloudenterprise-resources-ckafka_acl"
description: |-
  Provides a resource to create a Ckafka Acl.
---

# tencentcloudenterprise_ckafka_acl

Provides a resource to create a Ckafka Acl.

## Example Usage

### # Ckafka Acl

```hcl
resource "tencentcloudenterprise_ckafka_acl" "foo" {
  instance_id        = "ckafka-f9ife4zz"
  resource_type      = "TOPIC"
  resource_type_name = "topic-tf-test"
  operation_type     = "WRITE"
  permission_type    = "ALLOW"
  host               = "*"
  principal          = tencentcloudenterprise_ckafka_user.foo.account_name
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) ID of the ckafka instance.
* `operation_type` - (Required, String, ForceNew) ACL operation mode. Valid values: `READ`, `WRITE`.
* `permission_type` - (Required, String, ForceNew) ACL permission type. Valid values: `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
* `resource_type_name` - (Required, String, ForceNew) ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name;.
* `resource_type` - (Required, String, ForceNew) ACL resource type.  `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
* `host` - (Optional, String, ForceNew) IP address allowed to access. The default value is `*`, which means that any host can access.
* `principal` - (Optional, String, ForceNew) User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

Ckafka acl can be imported using the instance_id#permission_type#principal#host#operation_type#resource_type#resource_type_name, e.g.

```
$ terraform import tencentcloudenterprise_ckafka_acl.foo ckafka-f9ife4zz#ALLOW#test#*#WRITE#TOPIC#topic-tf-test
```

