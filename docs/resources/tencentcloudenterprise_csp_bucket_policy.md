---
subcategory: "Cloud Storage Private(CSP)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_csp_bucket_policy"
sidebar_current: "docs-tencentcloudenterprise-resources-csp_bucket_policy"
description: |-
  Provides a csp resource to create a Csp bucket policy and set its attributes.
---

# tencentcloudenterprise_csp_bucket_policy

Provides a csp resource to create a Csp bucket policy and set its attributes.

## Example Usage

```hcl
resource "tencentcloudenterprise_csp_bucket_policy" "csp_policy" {
  bucket = "mycsp-1258798060"

  policy = <<EOF

	{
	  "version": "2.0",
	  "Statement": [
	    {
	      "Principal": {
	        "qcs": [
	          "qcs::cam::uin/<your-account-id>:uin/<your-account-id>"
	        ]
	      },
	      "Action": [
	        "name/csp:DeleteBucket",
	        "name/csp:PutBucketACL"
	      ],
	      "Effect": "allow",
	      "Resource": [
	        "qcs::csp:<bucket region>:uid/<your-account-id>:<bucket name>/*"
	      ]
	    }
	  ]
	}

EOF
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required, String, ForceNew) The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycsp-1258798060`.
* `policy` - (Required, String) The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

csp bucket policy can be imported, e.g.

```
$ terraform import tencentcloudenterprise_csp_bucket_policy.bucket bucket-name
```

