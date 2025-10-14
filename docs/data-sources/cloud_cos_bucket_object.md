---
subcategory: "Cloud Object Storage(COS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cos_bucket_object"
sidebar_current: "docs-cloud-datasource-cos_bucket_object"
description: |-
  Use this data source to query the metadata of an object stored inside a bucket.
---

# cloud_cos_bucket_object

Use this data source to query the metadata of an object stored inside a bucket.

## Example Usage

```hcl
data "cloud_cos_bucket_object" "mycos" {
  bucket             = "mycos-test-1258798060"
  prefix             = "hello-world.py"
  result_output_file = "TfCosResults"
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required, String) Name of the bucket that contains the objects to query.
* `delimiter` - (Optional, String) bucket object key delimiter.
* `encoding_type` - (Optional, String) Specifies the encoding method of the returned key. Valid values: url. If this parameter is not specified, the returned key is not encoded.
* `marker` - (Optional, String) Specifies the object name to start with when listing objects in a bucket. The object name in the response must be greater than this parameter in lexicographic order.
* `max_keys` - (Optional, Int) Specifies the maximum number of objects returned in the response. The default value is 1000. The maximum value is 1000.
* `prefix` - (Optional, String) bucket object key prefix.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `common_prefixes` - Specifies the common prefix of the returned object list.
* `contents` - Specifies the object list.
* `is_truncated` - Specifies whether the returned object list is truncated.
* `name` - Specifies the object name.
* `next_marker` - Specifies the object name to start with when listing objects in a bucket. If the returned value of IsTruncated is true, the value of NextMarker is the object name that starts with the value of Marker in the request. If the returned value of IsTruncated is false, the value of NextMarker is null.


