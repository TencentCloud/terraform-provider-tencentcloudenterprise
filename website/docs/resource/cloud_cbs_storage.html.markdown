---
subcategory: "Cloud Block Storage(CBS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cbs_storage"
sidebar_current: "docs-cloud-resource-cbs_storage"
description: |-
  Provides a resource to create a CBS.
---

# cloud_cbs_storage

Provides a resource to create a CBS.

## Example Usage

```hcl
resource "cloud_cbs_storage" "storage" {
  storage_name      = "mystorage"
  storage_type      = "CLOUD_SSD"
  storage_size      = 100
  availability_zone = "ap-guangzhou-3"
  encrypt           = false

  tags = {
    test = "tf"
  }
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, String, ForceNew) The available zone that the CBS instance locates at.
* `storage_name` - (Required, String) Name of CBS. The maximum length can not exceed 60 bytes.
* `storage_size` - (Required, Int) Volume of CBS, and unit is GB.
* `storage_type` - (Required, String, ForceNew) Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_SSD: SSD.
* `charge_type` - (Optional, String) The charge type of CBS instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`.
* `disk_storage_pool_group` - (Optional, String) The pool group of CBS instance.
* `encrypt` - (Optional, Bool, ForceNew) Indicates whether CBS is encrypted.
* `force_delete` - (Optional, Bool) Indicate whether to delete CBS instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin.
* `prepaid_period` - (Optional, Int) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.
* `prepaid_renew_flag` - (Optional, String) Auto Renewal flag. Value range: `NOTIFY_AND_AUTO_RENEW`: Notify expiry and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: Notify expiry but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: Neither notify expiry nor renew automatically. Default value range: `NOTIFY_AND_MANUAL_RENEW`: Notify expiry but do not renew automatically. NOTE: it only works when charge_type is set to `PREPAID`.
* `snapshot_id` - (Optional, String) ID of the snapshot. If specified, created the CBS by this snapshot.
* `tags` - (Optional, Map) The available tags within this CBS.
* `throughput_performance` - (Optional, Int) Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `attached` - Indicates whether the CBS is mounted the CVM.
* `storage_id` - The ID of the CBS instance.
* `storage_status` - Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.


## Import

CBS storage can be imported using the id, e.g.

```
$ terraform import cloud_cbs_storage.storage disk-41s6jwy4
```

