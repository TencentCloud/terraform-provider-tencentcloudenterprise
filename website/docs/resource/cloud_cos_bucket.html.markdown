---
subcategory: "Cloud Object Storage(COS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cos_bucket"
sidebar_current: "docs-cloud-resource-cos_bucket"
description: |-
  Provides a cos resource to create a COS bucket and set its attributes.
---

# cloud_cos_bucket

Provides a cos resource to create a COS bucket and set its attributes.

## Example Usage

### # Private Bucket

```hcl
resource "cloud_cos_bucket" "mycos" {
  bucket = "mycos-1258798060"
  acl    = "private"
}
```

### # Creation of multiple available zone bucket

```hcl
resource "cloud_cos_bucket" "mycos" {
  bucket            = "mycos-1258798060"
  acl               = "private"
  versioning_enable = true
  force_clean       = true
}
```

### # Static Website

```hcl
resource "cloud_cos_bucket" "mycos" {
  bucket = "mycos-1258798060"

  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}

output "endpoint_test" {
  value = cloud_cos_bucket.mycos.website.0.endpoint
}
```

### # Using CORS

```hcl
resource "cloud_cos_bucket" "mycos" {
  bucket = "mycos-1258798060"
  acl    = "public-read-write"

  cors_rules {
    allowed_origins = ["http://*.abc.com"]
    allowed_methods = ["PUT", "POST"]
    allowed_headers = ["*"]
    max_age_seconds = 300
    expose_headers  = ["Etag"]
  }
}
```

### # Using object lifecycle

```hcl
resource "cloud_cos_bucket" "mycos" {
  bucket = "mycos-1258798060"
  acl    = "public-read-write"

  lifecycle_rules {
    filter_prefix = "path1/"

    transition {
      date          = "2019-06-01"
      storage_class = "STANDARD_IA"
    }

    expiration {
      days = 90
    }
  }
}
```

### # Using object lifecycle with tag filtering

```hcl
resource "cloud_cos_bucket" "mycos" {
  bucket = "mycos-1258798060"
  acl    = "public-read-write"

  lifecycle_rules {
    filter_tags = {
      "test-tag" = "yyy"
      "env"      = "prod"
    }

    expiration {
      days = 180
    }
  }
}
```

### # Using custom origin domain settings

```hcl
resource "cloud_cos_bucket" "with_origin" {
  bucket = "mycos-1258798060"
  acl    = "private"
  origin_domain_rules {
    domain = "abc.example.com"
    type   = "REST"
    status = "ENABLE"
  }
}
```

### Using origin-pull settings

```hcl
resource "cloud_cos_bucket" "with_origin" {
  bucket = "mycos-1258798060"
  acl    = "private"
  origin_pull_rules {
    priority            = 1
    sync_back_to_source = false
    host                = "abc.example.com"
    prefix              = "/"
    protocol            = "FOLLOW" // "HTTP" "HTTPS"
    follow_query_string = true
    follow_redirection  = true
    follow_http_headers = ["origin", "host"]
    custom_http_headers = {
      "x-custom-header" = "custom_value"
    }
  }
}
```

### # Setting log status

```hcl
resource "cloud_cam_role" "cosLogGrant" {
  name     = "CLS_QcsRole"
  document = <<EOF

	{
	  "version": "2.0",
	  "statement": [
	    {
	      "action": [
	        "name/sts:AssumeRole"
	      ],
	      "effect": "allow",
	      "principal": {
	        "service": [
	          "cls.cloud.tencent.com"
	        ]
	      }
	    }
	  ]
	}

EOF

  description = "cos log enable grant"
}

data "cloud_cam_policies" "cosAccess" {
  name = "QcloudCOSAccessForCLSRole"
}

resource "cloud_cam_role_policy_attachment" "cosLogGrant" {
  role_id   = cloud_cam_role.cosLogGrant.id
  policy_id = data.cloud_cam_policies.cosAccess.policy_list.0.policy_id
}

resource "cloud_cos_bucket" "mylog" {
  bucket = "mylog-1258798060"
  acl    = "private"
}

resource "cloud_cos_bucket" "mycos" {
  bucket            = "mycos-1258798060"
  acl               = "private"
  log_enable        = true
  log_target_bucket = "mylog-1258798060"
  log_prefix        = "MyLogPrefix"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required, String, ForceNew) The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
* `acl` - (Optional, String) The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
* `cors_rules` - (Optional, List) A rule of Cross-Origin Resource Sharing (documented below).
* `encryption_algorithm` - (Optional, String) The server-side encryption algorithm to use. Valid value is `AES256`, `SM4`.
* `force_clean` - (Optional, Bool) Force cleanup all objects before delete bucket.
* `lifecycle_rules` - (Optional, List) A configuration of object lifecycle management (documented below).
* `tags` - (Optional, Map) The tags of a bucket.
* `versioning_enable` - (Optional, Bool) Enable bucket versioning.
* `website` - (Optional, List) A website object(documented below).

The `cors_rules` object supports the following:

* `allowed_headers` - (Required, List) Specifies which headers are allowed.
* `allowed_methods` - (Required, List) Specifies which methods are allowed. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
* `allowed_origins` - (Required, List) Specifies which origins are allowed.
* `expose_headers` - (Optional, List) Specifies expose header in the response.
* `max_age_seconds` - (Optional, Int) Specifies time in seconds that browser can cache the response for a preflight request.

The `expiration` object supports the following:

* `days` - (Required, Int) Specifies the number of days after object creation when the specific rule action takes effect.
* `delete_marker` - (Optional, Bool) Indicates whether the delete marker of an expired object will be removed.

The `lifecycle_rules` object supports the following:

* `id` - (Required, String) A unique identifier for the rule. It can be up to 255 characters.
* `expiration` - (Optional, Set) Specifies a period in the object's expire (documented below).
* `filter_prefix` - (Optional, String) Object key prefix identifying one or more objects to which the rule applies.
* `filter_tags` - (Optional, Map) A map of tags to filter objects to which the rule applies.
* `non_current_expiration` - (Optional, Set) Specifies when non current object versions shall expire.
* `non_current_transition` - (Optional, Set) Specifies a period in the non current object's transitions.
* `transition` - (Optional, Set) Specifies a period in the object's transitions (documented below).

The `non_current_expiration` object supports the following:

* `non_current_days` - (Optional, Int) Number of days after non current object creation when the specific rule action takes effect. The maximum value is 3650.

The `non_current_transition` object supports the following:

* `storage_class` - (Required, String) Specifies the storage class to which you want the non current object to transition. Available values include `STANDARD_IA`, `MAZ_STANDARD_IA`, `INTELLIGENT_TIERING`, `MAZ_INTELLIGENT_TIERING`, `ARCHIVE`, `DEEP_ARCHIVE`. For more information, please refer to: https://cloud.tencent.com/document/product/436/33417.
* `non_current_days` - (Optional, Int) Number of days after non current object creation when the specific rule action takes effect.

The `transition` object supports the following:

* `days` - (Required, Int) Specifies the number of days after object creation when the specific rule action takes effect.
* `storage_class` - (Required, String) Specifies the storage class to which you want the object to transition. Available values include `STANDARD_IA`, `MAZ_STANDARD_IA`, `INTELLIGENT_TIERING`, `MAZ_INTELLIGENT_TIERING`, `ARCHIVE`, `DEEP_ARCHIVE`. For more information, please refer to: https://cloud.tencent.com/document/product/436/33417.

The `website` object supports the following:

* `index_document` - (Required, String) COS returns this index document when requests are made to the root domain or any of the subfolders.
* `error_document` - (Optional, String) An absolute path to the document to return in case of a 4XX error.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `cos_bucket_url` - The URL of this cos bucket.


## Import

COS bucket can be imported, e.g.

```
$ terraform import cloud_cos_bucket.bucket bucket-name
```

