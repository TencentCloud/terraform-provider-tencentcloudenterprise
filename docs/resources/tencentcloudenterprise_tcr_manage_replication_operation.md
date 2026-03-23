---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_manage_replication_operation"
sidebar_current: "docs-tencentcloudenterprise-resource-tcr_manage_replication_operation"
description: |-
  Provides a resource to create a tcr manage_replication_operation
---

# tencentcloudenterprise_tcr_manage_replication_operation

Provides a resource to create a tcr manage_replication_operation

## Example Usage

```hcl
resource "tencentcloudenterprise_tcr_instance" "mytcr_dest" {
  name          = "tf-test-tcr-%s"
  instance_type = "premium"
  delete_bucket = true
}

resource "tencentcloudenterprise_tcr_namespace" "myns_dest" {
  instance_id    = tencentcloudenterprise_tcr_instance.mytcr_dest.id
  name           = "tf_test_ns_dest"
  is_public      = true
  is_auto_scan   = true
  is_prevent_vul = true
  severity       = "medium"
  cve_whitelist_items {
    cve_id = "cve-xxxxx"
  }
}

resource "tencentcloudenterprise_tcr_manage_replication_operation" "my_replica" {
  source_registry_id      = local.tcr_id
  destination_registry_id = tencentcloudenterprise_tcr_instance.mytcr_dest.id
  rule {
    name           = "test_sync_%d"
    dest_namespace = tencentcloudenterprise_tcr_namespace.myns_dest.name
    override       = true
    filters {
      type  = "name"
      value = join("/", [var.tcr_namespace, "**"])
    }
    filters {
      type  = "tag"
      value = ""
    }
    filters {
      type  = "resource"
      value = ""
    }
  }
  description           = "this is the tcr sync operation"
  destination_region_id = 1 // "ap-region"
  peer_replication_option {
    peer_registry_uin       = ""
    peer_registry_token     = ""
    enable_peer_replication = false
  }
}
```

## Argument Reference

The following arguments are supported:

* `destination_registry_id` - (Required, String, ForceNew) Copy destination instance Id.
* `rule` - (Required, List, ForceNew) Synchronization rules.
* `source_registry_id` - (Required, String, ForceNew) Copy source instance Id.
* `description` - (Optional, String, ForceNew) Rule description.
* `destination_region_id` - (Optional, Int, ForceNew) The region ID of the target instance.
* `peer_replication_option` - (Optional, List, ForceNew) Enable synchronization of configuration items across master account instances.

The `filters` object supports the following:

* `type` - (Required, String) Type (name, tag, and resource).
* `value` - (Optional, String) Empty by default.

The `peer_replication_option` object supports the following:

* `enable_peer_replication` - (Required, Bool) Whether to enable cross-master account instance synchronization.
* `peer_registry_token` - (Required, String) Access permanent token of the instance to be synchronized.
* `peer_registry_uin` - (Required, String) Uin of the instance to be synchronized.

The `rule` object supports the following:

* `dest_namespace` - (Required, String) Target namespace.
* `filters` - (Required, List) Sync filters.
* `name` - (Required, String) Synchronization rule names.
* `override` - (Required, Bool) Whether to cover.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


