---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_account"
sidebar_current: "docs-cloud-resource-dcdb_account"
description: |-
  Provides a resource to create a dcdb account
---

# cloud_dcdb_account

Provides a resource to create a dcdb account

## Example Usage

```hcl
resource "cloud_dcdb_account" "account" {
  instance_id = "tdsqlshard-973xatu3"
  user_name   = "mysql"
  host        = "127.0.0.1"
  password    = "===password==="
  read_only   = 0
  description = "test modify ---this is a test account"
}
```

## Argument Reference

The following arguments are supported:

* `host` - (Required, String) Db host.
* `instance_id` - (Required, String) Instance id.
* `password` - (Required, String) Password.
* `user_name` - (Required, String) Account name.
* `delay_thresh` - (Optional, Int) If the delay of the backup machine exceeds the value set in this parameter, the system will consider the backup machine to have a failure. It is recommended to set the parameter value greater than 10. This parameter takes effect when ReadOnly is set to 1 or 2.
* `description` - (Optional, String) Description for account.
* `read_only` - (Optional, Int) Whether the account is readonly. 0 means not a readonly account.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

dcdb account can be imported using the id, e.g.
```
$ terraform import cloud_dcdb_account.account account_id
```

