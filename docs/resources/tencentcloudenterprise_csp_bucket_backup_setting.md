---
subcategory: "Cloud Storage Platform(CSP)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_csp_bucket_backup_setting"
sidebar_current: "docs-tencentcloudenterprise-resource-csp_bucket_backup_setting"
description: |-
  Provides a resource to manage CSP bucket backup setting.
---

# tencentcloudenterprise_csp_bucket_backup_setting

Provides a resource to manage CSP bucket backup setting.

## Example Usage

```hcl
resource "tencentcloudenterprise_csp_bucket" "src" {
  bucket = "src-bucket-1258798060"
  acl    = "private"
}

resource "tencentcloudenterprise_csp_bucket_backup_setting" "example" {
  bucket             = tencentcloudenterprise_csp_bucket.src.bucket
  region             = "kazakhstan-1"
  backup_enable      = true
  backup_bucket_name = "dst-bucket-1258798060"
  backup_endpoint    = "cos.kazakhstan-1.csp.example.com"
  access_key         = "AKIDxxxxxxxxxxxxxxxx"
  secret_key         = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  backsource_enabled = false
  use_https          = false
}
```

## Argument Reference

The following arguments are supported:

* `backup_enable` - (Required, Bool) Whether to enable bucket backup.
* `bucket` - (Required, String, ForceNew) Source bucket name. Format: [custom name]-[appid], e.g. `mybucket-1258798060`.
* `region` - (Required, String, ForceNew) Region of the source bucket, e.g. `kazakhstan-1`.
* `access_key` - (Optional, String) SecretId (AccessKey) for accessing the backup bucket. Required when `backup_enable` is true.
* `backsource_enabled` - (Optional, Bool) Whether to enable image origin-pull (back-to-source) from the backup bucket.
* `backup_bucket_name` - (Optional, String) Backup destination bucket name. Required when `backup_enable` is true. Format: [custom name]-[appid].
* `backup_endpoint` - (Optional, String) COS domain name suffix for accessing the backup bucket. Required when `backup_enable` is true. Example: `cos.ap-region.mydomain.com`.
* `secret_key` - (Optional, String) SecretKey for accessing the backup bucket. Required when `backup_enable` is true.
* `use_https` - (Optional, Bool) Whether to use HTTPS when connecting to the backup bucket.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_csp_bucket_backup_setting can be imported using the id, e.g.

```
csp bucket backup setting can be imported using the id (bucket#region), e.g.

```
$ terraform import tencentcloudenterprise_csp_bucket_backup_setting.example src-bucket-1258798060#kazakhstan-1
```
```

