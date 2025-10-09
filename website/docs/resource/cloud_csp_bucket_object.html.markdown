---
subcategory: "Cloud Storage Private(CSP)"
layout: "cloud"
page_title: "TencentCloud: cloud_csp_bucket_object"
sidebar_current: "docs-cloud-resource-csp_bucket_object"
description: |-
  Provides a csp object resource to put an object(content or file) to the bucket.
---

# cloud_csp_bucket_object

Provides a csp object resource to put an object(content or file) to the bucket.

## Example Usage

### # Uploading a file to a bucket

```hcl
resource "cloud_csp_bucket_object" "myobject" {
  bucket = "mycsp-1258798060"
  key    = "new_object_key"
  acl    = "public-read"
  source = "path/to/file"
}
```

### # Uploading a content to a bucket

```hcl
resource "cloud_csp_bucket" "mycsp" {
  bucket = "mycsp-1258798060"
  acl    = "public-read"
}

resource "cloud_csp_bucket_object" "myobject" {
  bucket  = cloud_csp_bucket.mycsp.bucket
  key     = "new_object_key"
  content = "the content that you want to upload."
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required, String, ForceNew) The name of a bucket to use. Bucket format should be [custom name]-[appid], for example `mycsp-1258798060`.
* `key` - (Required, String, ForceNew) The name of the object once it is in the bucket.
* `acl` - (Optional, String) The canned ACL to apply. Available values include `private` and `public-read`.
* `cache_control` - (Optional, String) Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.
* `content_disposition` - (Optional, String) Specifies presentational information for the object.
* `content_encoding` - (Optional, String) Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
* `content_type` - (Optional, String) A standard MIME type describing the format of the object data.
* `content` - (Optional, String) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
* `source` - (Optional, String) The path to the source file being uploaded to the bucket.
* `tags` - (Optional, Map) Tag of the object.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



